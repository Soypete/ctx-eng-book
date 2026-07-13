#!/usr/bin/env python3
"""
Upload chapters from ctx-eng-book to Notion.

Usage:
    # First run (will create root page):
    op run -- python scripts/upload_to_notion.py

    # Subsequent runs (uses existing parent):
    export NOTION_PARENT_PAGE_ID="..."
    op run -- python scripts/upload_to_notion.py
"""

import os
import sys
import json
import subprocess
from pathlib import Path
from typing import Optional
from dataclasses import dataclass


def get_secret_from_1password(item: str, field: str = "credential") -> Optional[str]:
    """Get a secret from 1Password."""
    try:
        result = subprocess.run(
            ["op", "read", f"op://{item}/{field}"],
            capture_output=True,
            text=True,
            timeout=10,
        )
        if result.returncode == 0:
            return result.stdout.strip()
    except Exception as e:
        print(f"Warning: Could not read from 1Password: {e}")
    return None


# Try environment variable first, then 1Password
NOTION_API_KEY = os.environ.get("NOTION_API_KEY") or get_secret_from_1password(
    "notion_api_key"
)
NOTION_PARENT_PAGE_ID = os.environ.get("NOTION_PARENT_PAGE_ID")

BOOK_ROOT = (
    Path(__file__).parent.parent / "cmd" / "authorpedro" / "books" / "ctx-eng-book"
)

PART_STRUCTURE = [
    (
        "Part I — Why AI Systems Fail",
        [
            "Every Failure Is a Context Failure",
            "AI Is a Systems Problem",
        ],
    ),
    (
        "Part II — How Models Actually Use Context",
        [
            "Attention Is All You Need (But We Stopped Paying Attention)",
            "In-Context Learning and Pragmatics",
            "Tool Use Is Structured Context",
        ],
    ),
    (
        "Part III — Context Is Data",
        [
            "Memory Is a Database Problem",
            "Context Is a Query Over Distributed State",
            "Knowledge Graphs and Semantic Context",
            "Retrieval Beyond Vector Databases",
        ],
    ),
    (
        "Part IV — Context Must Be Governed",
        [
            "Personalization Is Governed Data Access",  # Authorization Across Stores
            "Stop Giving Agents Permissions",
            "The UNIX Philosophy of AI Systems",
        ],
    ),
    (
        "Part V — Orchestration and Cost",
        [
            "Agents Are Workflows",
            "The Cost of Context",
            "When Context Engineering Stops Working",
        ],
    ),
    (
        "Part VI — Reliability Engineering for AI",
        [
            "Observability for Context Systems",
            "Evaluating AI Systems",
            "Building a Context Engineering Platform",
        ],
    ),
]

CHAPTER_FILE_MAP = {
    "Every Failure Is a Context Failure": "chapters/ch01-every-failure-is-a-context-failure/modules/ch01-every-failure-is-a-context-failure.md",
    "AI Is a Systems Problem": "chapters/ch02-ai-is-a-systems-problem/modules/ch02-ai-is-a-systems-problem.md",
    "Attention Is All You Need (But We Stopped Paying Attention)": "chapters/ch03-attention-is-all-you-need/modules/ch03-attention-is-all-you-need.md",
    "In-Context Learning and Pragmatics": "chapters/ch04-in-context-learning-and-pragmatics/modules/ch04-in-context-learning.md",
    "Tool Use Is Structured Context": "chapters/ch05-tool-use-is-structured-context/modules/ch05-tool-use-is-structured-context.md",
    "Memory Is a Database Problem": "chapters/ch06-memory-is-a-database-problem/modules/ch06-memory-is-a-database-problem.md",
    "Context Is a Query Over Distributed State": "chapters/ch07-context-is-a-query/modules/ch07-context-is-a-query.md",
    "Knowledge Graphs and Semantic Context": "chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08-knowledge-graphs-and-semantic-context.md",
    "Retrieval Beyond Vector Databases": "chapters/ch09-retrieval-beyond-vector-databases/modules/ch09-retrieval-beyond-vector-databases.md",
    "Personalization Is Governed Data Access": "chapters/ch10-authorization-across-stores/modules/ch10-authorization-across-stores.md",
    "Stop Giving Agents Permissions": "chapters/ch11-stop-giving-agents-permissions/modules/ch11-stop-giving-agents-permissions.md",
    "The UNIX Philosophy of AI Systems": "chapters/ch12-the-unix-philosophy-of-ai-systems/modules/ch12-the-unix-philosophy-of-ai-systems.md",
    "Agents Are Workflows": "chapters/ch13-agents-are-workflows/modules/ch13-agents-are-workflows.md",
    "The Cost of Context": "chapters/ch14-the-cost-of-context/modules/ch14-the-cost-of-context.md",
    "When Context Engineering Stops Working": "chapters/ch15-when-context-engineering-stops-working/modules/ch15-when-context-engineering-stops-working.md",
    "Observability for Context Systems": "chapters/ch16-observability-for-context-systems/modules/ch16-observability-for-context-systems.md",
    "Evaluating AI Systems": "chapters/ch17-evaluating-ai-systems/modules/ch17-evaluating-ai-systems.md",
    "Building a Context Engineering Platform": "chapters/ch18-building-a-context-engineering-platform/modules/ch18-building-a-context-engineering-platform.md",
}


def get_notion_client():
    """Initialize Notion client."""
    try:
        from notion_client import Client
    except ImportError:
        print("Installing notion-client...")
        subprocess.run(
            [sys.executable, "-m", "pip", "install", "notion-client"], check=True
        )
        from notion_client import Client

    if not NOTION_API_KEY:
        raise ValueError(
            "NOTION_API_KEY not found in 1Password (notion_api_key/credential)"
        )

    return Client(auth=NOTION_API_KEY)


def create_page(client, title: str, parent_id: str, content: str = "") -> str:
    """Create a page in Notion and return its ID."""
    properties = {"title": {"title": [{"text": {"content": title}}]}}

    response = client.pages.create(
        parent={"page_id": parent_id},
        properties=properties,
    )
    return response["id"]


def create_page_with_content(client, title: str, parent_id: str, content: str) -> str:
    """Create a page with content blocks in Notion."""
    # Create the page first
    page_id = create_page(client, title, parent_id, content)

    # Convert markdown to blocks and append
    blocks = markdown_to_blocks(content)

    # Append blocks in batches of 100
    for i in range(0, len(blocks), 100):
        batch = blocks[i : i + 100]
        client.blocks.children.append(block_id=page_id, children=batch)

    return page_id


def markdown_to_blocks(markdown: str) -> list:
    """Convert markdown to Notion blocks."""
    blocks = []
    lines = markdown.split("\n")
    i = 0

    while i < len(lines):
        line = lines[i]

        # Skip empty lines
        if not line.strip():
            i += 1
            continue

        # Heading 1
        if line.startswith("# "):
            blocks.append(
                {
                    "object": "block",
                    "type": "heading_1",
                    "heading_1": {"rich_text": [{"text": {"content": line[2:]}}]},
                }
            )

        # Heading 2
        elif line.startswith("## "):
            blocks.append(
                {
                    "object": "block",
                    "type": "heading_2",
                    "heading_2": {"rich_text": [{"text": {"content": line[3:]}}]},
                }
            )

        # Heading 3
        elif line.startswith("### "):
            blocks.append(
                {
                    "object": "block",
                    "type": "heading_3",
                    "heading_3": {"rich_text": [{"text": {"content": line[4:]}}]},
                }
            )

        # Code block
        elif line.startswith("```"):
            lang = line[3:].strip().lower()
            valid_langs = {
                "abap",
                "abc",
                "agda",
                "arduino",
                "ascii art",
                "assembly",
                "bash",
                "basic",
                "bnf",
                "c",
                "c#",
                "c++",
                "clojure",
                "coffeescript",
                "coq",
                "css",
                "dart",
                "dhall",
                "diff",
                "docker",
                "ebnf",
                "elixir",
                "elm",
                "erlang",
                "f#",
                "flow",
                "fortran",
                "gherkin",
                "glsl",
                "go",
                "graphql",
                "groovy",
                "haskell",
                "hcl",
                "html",
                "idris",
                "java",
                "javascript",
                "json",
                "julia",
                "kotlin",
                "latex",
                "less",
                "lisp",
                "livescript",
                "llvm ir",
                "lua",
                "makefile",
                "markdown",
                "markup",
                "matlab",
                "mathematica",
                "mermaid",
                "nix",
                "notion formula",
                "objective-c",
                "ocaml",
                "pascal",
                "perl",
                "php",
                "plain text",
                "powershell",
                "prolog",
                "protobuf",
                "purescript",
                "python",
                "r",
                "racket",
                "reason",
                "ruby",
                "rust",
                "sass",
                "scala",
                "scheme",
                "scss",
                "shell",
                "smalltalk",
                "solidity",
                "sql",
                "swift",
                "toml",
                "typescript",
                "vb.net",
                "verilog",
                "vhdl",
                "visual basic",
                "webassembly",
                "xml",
                "yaml",
                "java/c/c++/c#",
            }
            if not lang or lang not in valid_langs:
                lang = "plain text"
            code_lines = []
            i += 1
            while i < len(lines) and not lines[i].startswith("```"):
                code_lines.append(lines[i])
                i += 1
            blocks.append(
                {
                    "object": "block",
                    "type": "code",
                    "code": {
                        "language": lang,
                        "rich_text": [{"text": {"content": "\n".join(code_lines)}}],
                    },
                }
            )

        # Bullet list
        elif line.strip().startswith("- ") or line.strip().startswith("* "):
            blocks.append(
                {
                    "object": "block",
                    "type": "bulleted_list_item",
                    "bulleted_list_item": {
                        "rich_text": [{"text": {"content": line.strip()[2:]}}]
                    },
                }
            )

        # Numbered list
        elif line.strip()[0].isdigit() and ". " in line.strip():
            content = line.strip()
            idx = content.index(". ")
            blocks.append(
                {
                    "object": "block",
                    "type": "numbered_list_item",
                    "numbered_list_item": {
                        "rich_text": [{"text": {"content": content[idx + 2 :]}}]
                    },
                }
            )

        # Quote
        elif line.startswith("> "):
            blocks.append(
                {
                    "object": "block",
                    "type": "quote",
                    "quote": {"rich_text": [{"text": {"content": line[2:]}}]},
                }
            )

        # Paragraph (default)
        else:
            blocks.append(
                {
                    "object": "block",
                    "type": "paragraph",
                    "paragraph": {"rich_text": [{"text": {"content": line}}]},
                }
            )

        i += 1

    return blocks


def read_chapter_content(chapter_file: str) -> str:
    """Read chapter content from file."""
    file_path = BOOK_ROOT / chapter_file
    if not file_path.exists():
        print(f"Warning: File not found: {file_path}")
        return ""
    return file_path.read_text(encoding="utf-8")


def upload_book_to_notion():
    """Main function to upload the entire book."""
    if not NOTION_API_KEY:
        print("Error: NOTION_API_KEY environment variable not set")
        print("Run: export NOTION_API_KEY='secret_...'")
        sys.exit(1)

    client = get_notion_client()

    # Create root book page or use existing parent
    if NOTION_PARENT_PAGE_ID:
        root_id = NOTION_PARENT_PAGE_ID
        print(f"Using existing parent page: {root_id}")
    else:
        print("Creating root book page...")
        root_id = create_page(
            client, "Context Engineering: Building Reliable AI Systems", "", ""
        )
        print(f"Created root page: {root_id}")
        print(f"Set NOTION_PARENT_PAGE_ID={root_id} for future runs")

    # Upload preface if exists
    preface_path = BOOK_ROOT / "preface.md"
    if preface_path.exists():
        print("\nUploading preface...")
        content = preface_path.read_text(encoding="utf-8")
        create_page_with_content(client, "Preface", root_id, content)
        print("Preface uploaded")

    # Upload each part and chapter
    for part_name, chapters in PART_STRUCTURE:
        print(f"\n{'=' * 50}")
        print(f"Creating Part: {part_name}")
        print(f"{'=' * 50}")

        part_id = create_page(client, part_name, root_id, "")
        print(f"  Created: {part_id}")

        for chapter_title in chapters:
            print(f"\n  Uploading: {chapter_title}")

            if chapter_title not in CHAPTER_FILE_MAP:
                print(f"    Warning: No file mapping for '{chapter_title}'")
                continue

            chapter_file = CHAPTER_FILE_MAP[chapter_title]
            content = read_chapter_content(chapter_file)

            if content:
                create_page_with_content(client, chapter_title, part_id, content)
                print(f"    Uploaded: {len(content)} chars")
            else:
                print(f"    Warning: No content found")

    print("\n" + "=" * 50)
    print("Upload complete!")
    print("=" * 50)


if __name__ == "__main__":
    upload_book_to_notion()
