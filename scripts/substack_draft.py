#!/usr/bin/env python3
"""
Create a Substack draft from a book/substack/{module}/post.md file.

Substack's draft API wants the body as a ProseMirror JSON document in
`draft_body`. Sending markdown (or any other field) creates a draft whose body
is silently empty, so this script converts the markdown, uploads local images,
creates the draft, then reads it back and fails if the body did not stick.

Usage:
    # Inspect the converted body without touching Substack:
    python3 scripts/substack_draft.py book/substack/ch00-what-we-mean-by-context-engineering --dry-run

    # Create the draft (substack.sid cookie from a logged-in browser session):
    SUBSTACK_SESSION_COOKIE="..." python3 scripts/substack_draft.py book/substack/ch00-...

This uses Substack's unofficial API and may break.
"""

import argparse
import base64
import http.cookiejar
import json
import mimetypes
import os
import re
import sys
import urllib.error
import urllib.parse
import urllib.request
from pathlib import Path
from typing import Any, Optional

REPO_ROOT = Path(__file__).parent.parent


def _load_dotenv(path: Path) -> None:
    """Load KEY=value lines from the repo's git-ignored .env without overriding the environment."""
    if not path.is_file():
        return
    for line in path.read_text().splitlines():
        key, sep, value = line.strip().partition("=")
        if sep and not key.startswith("#"):
            os.environ.setdefault(key.strip(), value.strip().strip("\"'"))


_load_dotenv(REPO_ROOT / ".env")

SUBSTACK_SESSION_COOKIE = os.environ.get("SUBSTACK_SESSION_COOKIE", "")
SUBSTACK_PUBLICATION = os.environ.get("SUBSTACK_PUBLICATION", "ctxengineering")
PUBLICATION_BASE = f"https://{SUBSTACK_PUBLICATION}.substack.com/api/v1"
SUBSTACK_BASE = "https://substack.com/api/v1"
RAW_GITHUB_PREFIX = "https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/"

INLINE_PATTERN = re.compile(
    r"(?P<code>`[^`]+`)"
    r"|(?P<link>\[(?P<link_text>[^\]]+)\]\((?P<href>[^)\s]+)\))"
    r"|(?P<bold>\*\*(?P<bold_text>.+?)\*\*)"
    r"|(?P<italic>(?<![*\w])\*(?P<italic_text>[^*\s][^*]*?)\*(?![*\w]))"
)


# --- markdown -> ProseMirror -------------------------------------------------


def parse_post(path: Path) -> tuple[dict[str, str], str]:
    """Split post.md into front matter fields and markdown body."""
    text = path.read_text(encoding="utf-8")
    meta: dict[str, str] = {}
    if text.startswith("---\n"):
        header, _, text = text[4:].partition("\n---\n")
        for line in header.splitlines():
            key, sep, value = line.partition(":")
            if sep:
                meta[key.strip()] = value.strip()
    return meta, text


def inline_nodes(text: str, marks: Optional[list] = None) -> list[dict[str, Any]]:
    """Convert inline markdown (code, links, bold, italic) to text nodes with marks."""
    marks = marks or []
    nodes: list[dict[str, Any]] = []

    def add_text(value: str, node_marks: list) -> None:
        if not value:
            return
        node: dict[str, Any] = {"type": "text", "text": value}
        if node_marks:
            node["marks"] = node_marks
        nodes.append(node)

    pos = 0
    for match in INLINE_PATTERN.finditer(text):
        add_text(text[pos : match.start()], marks)
        if match.group("code"):
            add_text(match.group("code")[1:-1], marks + [{"type": "code"}])
        elif match.group("link"):
            link = {"type": "link", "attrs": {"href": match.group("href")}}
            nodes.extend(inline_nodes(match.group("link_text"), marks + [link]))
        elif match.group("bold"):
            nodes.extend(inline_nodes(match.group("bold_text"), marks + [{"type": "strong"}]))
        else:
            nodes.extend(inline_nodes(match.group("italic_text"), marks + [{"type": "em"}]))
        pos = match.end()
    add_text(text[pos:], marks)
    return nodes


def paragraph(text: str) -> dict[str, Any]:
    return {"type": "paragraph", "content": inline_nodes(text)}


def table_to_list(rows: list[str]) -> dict[str, Any]:
    """Substack posts have no tables; render each row as a bullet keyed by its first cell."""
    cells = [[c.strip() for c in row.strip().strip("|").split("|")] for row in rows]
    header, body = cells[0], [r for r in cells[1:] if not set("".join(r)) <= set("-: ")]
    items = []
    for row in body:
        parts = [f"**{row[0]}**"] + [
            f"{header[i]}: {cell}" for i, cell in enumerate(row[1:], start=1) if i < len(header)
        ]
        items.append({"type": "list_item", "content": [paragraph(" — ".join(parts))]})
    return {"type": "bullet_list", "content": items}


def markdown_to_doc(markdown: str, image_src) -> dict[str, Any]:
    """Convert post markdown to a Substack ProseMirror document.

    image_src maps a markdown image URL to the src Substack should display.
    """
    blocks: list[dict[str, Any]] = []
    lines = markdown.split("\n")
    i = 0

    def collect(predicate) -> list[str]:
        nonlocal i
        group = []
        while i < len(lines) and predicate(lines[i]):
            group.append(lines[i])
            i += 1
        return group

    while i < len(lines):
        line = lines[i]
        stripped = line.strip()

        if not stripped:
            i += 1
        elif stripped.startswith("```"):
            lang = stripped[3:].strip()
            i += 1
            code = collect(lambda l: not l.strip().startswith("```"))
            i += 1  # closing fence
            node: dict[str, Any] = {"type": "code_block", "attrs": {"language": lang or None}}
            if code:
                node["content"] = [{"type": "text", "text": "\n".join(code)}]
            blocks.append(node)
        elif heading := re.match(r"(#{1,6}) (.*)", stripped):
            blocks.append(
                {
                    "type": "heading",
                    "attrs": {"level": len(heading.group(1))},
                    "content": inline_nodes(heading.group(2)),
                }
            )
            i += 1
        elif image := re.fullmatch(r"!\[([^\]]*)\]\(([^)\s]+)\)", stripped):
            alt, url = image.groups()
            blocks.append(
                {
                    "type": "captionedImage",
                    "content": [{"type": "image2", "attrs": {"src": image_src(url), "alt": alt}}],
                }
            )
            i += 1
        elif stripped.startswith("|"):
            blocks.append(table_to_list(collect(lambda l: l.strip().startswith("|"))))
        elif stripped.startswith(">"):
            quote = collect(lambda l: l.strip().startswith(">"))
            text = " ".join(l.strip()[1:].strip() for l in quote)
            blocks.append({"type": "blockquote", "content": [paragraph(text)]})
        elif re.match(r"[-*] ", stripped):
            items = collect(lambda l: re.match(r"\s*[-*] ", l) is not None)
            blocks.append(
                {
                    "type": "bullet_list",
                    "content": [
                        {"type": "list_item", "content": [paragraph(l.strip()[2:])]} for l in items
                    ],
                }
            )
        elif re.match(r"\d+\. ", stripped):
            items = collect(lambda l: re.match(r"\s*\d+\. ", l) is not None)
            blocks.append(
                {
                    "type": "ordered_list",
                    "attrs": {"start": int(stripped.split(".")[0])},
                    "content": [
                        {
                            "type": "list_item",
                            "content": [paragraph(l.strip().split(". ", 1)[1])],
                        }
                        for l in items
                    ],
                }
            )
        else:
            text = collect(
                lambda l: l.strip() != ""
                and not re.match(r"(```|#{1,6} |!\[|\||>|[-*] |\d+\. )", l.strip())
            )
            blocks.append(paragraph(" ".join(l.strip() for l in text)))

    return {"type": "doc", "content": blocks}


# --- Substack API -------------------------------------------------------------


def _opener() -> urllib.request.OpenerDirector:
    jar = http.cookiejar.CookieJar()
    for domain in (".substack.com", f"{SUBSTACK_PUBLICATION}.substack.com"):
        jar.set_cookie(
            http.cookiejar.Cookie(
                0, "substack.sid", SUBSTACK_SESSION_COOKIE, None, False, domain, True,
                domain.startswith("."), "/", True, True, None, False, None, None, {},
            )
        )
    opener = urllib.request.build_opener(urllib.request.HTTPCookieProcessor(jar))
    opener.addheaders = [
        ("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36"),
    ]
    return opener


def api(opener, method: str, url: str, payload: Optional[dict] = None, form: bool = False) -> Any:
    """Call the API with a JSON body, or a form-encoded one where Substack expects a form."""
    data, headers = None, {}
    if payload is not None and form:
        data = urllib.parse.urlencode(payload).encode()
        headers["Content-Type"] = "application/x-www-form-urlencoded"
    elif payload is not None:
        data = json.dumps(payload).encode()
        headers["Content-Type"] = "application/json"
    request = urllib.request.Request(url, data=data, headers=headers, method=method)
    try:
        with opener.open(request, timeout=60) as response:
            return json.loads(response.read() or b"null")
    except urllib.error.HTTPError as e:
        sys.exit(f"{method} {url} failed: {e.code} {e.read()[:500]!r}")


def upload_image(opener, path: Path) -> str:
    """Upload a local image to Substack's CDN and return its URL."""
    mime = mimetypes.guess_type(path.name)[0] or "image/png"
    data_uri = f"data:{mime};base64,{base64.b64encode(path.read_bytes()).decode()}"
    return api(opener, "POST", f"{PUBLICATION_BASE}/image", {"image": data_uri}, form=True)["url"]


def local_image_path(url: str, module_dir: Path) -> Optional[Path]:
    """Resolve an image URL to the repo file it was published from, if any."""
    if url.startswith(RAW_GITHUB_PREFIX):
        path = REPO_ROOT / url[len(RAW_GITHUB_PREFIX) :]
    else:
        path = module_dir / url
    return path if path.is_file() else None


def body_text_length(doc: dict[str, Any]) -> int:
    """Count visible characters, so an empty-looking body is caught."""
    if doc.get("type") == "text":
        return len(doc.get("text", ""))
    if doc.get("type") == "image2":
        return 1
    return sum(body_text_length(child) for child in doc.get("content", []))


def create_draft(module_dir: Path, dry_run: bool) -> None:
    meta, markdown = parse_post(module_dir / "post.md")
    opener = None if dry_run else _opener()

    def image_src(url: str) -> str:
        path = local_image_path(url, module_dir)
        if dry_run or path is None:
            return url
        print(f"  uploading image {path.name}")
        return upload_image(opener, path)

    doc = markdown_to_doc(markdown, image_src)
    if body_text_length(doc) == 0:
        sys.exit("Converted body is empty; refusing to create a blank draft.")

    thumbnail = module_dir / "thumbnail.png"
    if dry_run:
        cover = str(thumbnail) if thumbnail.is_file() else None
        print(json.dumps({"title": meta.get("title"), "subtitle": meta.get("subtitle"), "cover_image": cover, "body": doc}, indent=2))
        return

    payload: dict[str, Any] = {
        "draft_title": meta.get("title", module_dir.name),
        "draft_subtitle": meta.get("subtitle", ""),
        "draft_body": json.dumps(doc),
        "audience": "everyone",
        "type": "newsletter",
    }
    if thumbnail.is_file():
        print(f"  uploading cover image {thumbnail.name}")
        payload["cover_image"] = upload_image(opener, thumbnail)
    else:
        print(f"  warning: no {thumbnail.name}; draft will have no cover image")

    user_id = api(opener, "GET", f"{SUBSTACK_BASE}/user/profile/self")["id"]
    payload["draft_bylines"] = [{"id": user_id, "is_guest": False}]
    draft = api(opener, "POST", f"{PUBLICATION_BASE}/drafts", payload)

    # Read the draft back: a wrong body format still returns 200 with an empty body.
    saved = api(opener, "GET", f"{PUBLICATION_BASE}/drafts/{draft['id']}")
    saved_body = json.loads(saved.get("draft_body") or "{}")
    if body_text_length(saved_body) == 0:
        sys.exit(f"Draft {draft['id']} was created but its body is empty on Substack.")
    if "cover_image" in payload and not saved.get("cover_image"):
        print(f"  warning: cover image did not stick on draft {draft['id']}; set it in the editor")

    print(f"Draft {draft['id']} created: https://{SUBSTACK_PUBLICATION}.substack.com/publish/post/{draft['id']}")
    print(f"  {body_text_length(saved_body)} characters saved")


def main() -> None:
    parser = argparse.ArgumentParser(description="Create a Substack draft from a post.md file.")
    parser.add_argument("module_dir", type=Path, help="book/substack/{module} directory containing post.md")
    parser.add_argument("--dry-run", action="store_true", help="print the converted body and exit")
    args = parser.parse_args()

    if not (args.module_dir / "post.md").is_file():
        sys.exit(f"No post.md in {args.module_dir}")
    if not args.dry_run and not SUBSTACK_SESSION_COOKIE:
        sys.exit("Set SUBSTACK_SESSION_COOKIE (env or .env in the repo root) to your substack.sid cookie value.")

    create_draft(args.module_dir, args.dry_run)


if __name__ == "__main__":
    main()
