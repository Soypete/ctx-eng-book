#!/usr/bin/env python3
"""List candidate terms of art in a markdown draft with where each first appears.

This is a helper for the editor's concept-introduction trace, not a judge. It surfaces
terms a reader might not know, so the reviewer can check each one was explained before
(or where) it is first used. It favours recall: expect some noise.

Candidates:
  - *emphasized* and **bold** phrases
  - `code` identifiers used in prose (outside fenced blocks)
  - Capitalized multi-word names and acronyms (e.g. "Forge", "MCP", "RAG")
  - "the X" references to something the reader is assumed to know
    (the companion, the curriculum, the harness, the repo ...)

For each, prints: first line number, how many times it appears, and whether a nearby
defining phrase ("I mean", "means", "is a", "refers to", "—", ":") appears on the
first-use line. A definition flag is a hint, not proof.

Usage: term_trace.py FILE [--min-count 1]
"""
import argparse
import re
from collections import OrderedDict
from pathlib import Path

DEFINING = re.compile(r"^\W{0,3}(,? (which|that) (is|means)|\s*(—|--|\()|:? (I mean|means|is an?|is the|refers to|stands for|is what)\b)", re.I)
MAX_WORDS = 5  # longer emphasized spans are sentences, not terms
STOP = {"The", "This", "That", "These", "Those", "When", "What", "Who", "Why", "How", "If", "It",
        "In", "On", "For", "And", "But", "Or", "A", "An", "I", "We", "You", "Here", "There", "Then",
        "Use", "Run", "Keep", "Make", "Do", "Don't", "Not", "Next", "Picture", "Notice", "Each",
        "Every", "Some", "Most", "One", "Two", "Three", "Ask", "Try", "Put", "Set", "Check"}
ASSUMED = re.compile(r"\bthe (companion|curriculum|repo|repository|harness|host|manifest|ledger|"
                     r"framework|pipeline|gate|loop|example|spine|lens|series|book|course|workshop|"
                     r"deck|slides?|module|chapter)\b", re.I)


def candidates(line):
    line = re.sub(r"\]\([^)]*\)", "]", line)  # drop link targets
    for m in re.finditer(r"\*\*([^*]+)\*\*|(?<!\*)\*([^*\s][^*]*?)\*(?!\*)|_([^_\s][^_]*?)_", line):
        t = (m.group(1) or m.group(2) or m.group(3)).strip().rstrip(":.")
        if len(t.split()) <= MAX_WORDS:
            yield t
    for m in re.finditer(r"`([^`]+)`", line):
        yield "`" + m.group(1) + "`"
    for m in re.finditer(r"\b([A-Z][a-zA-Z0-9]+(?:[ -][A-Z][a-zA-Z0-9]+)*)\b", line):
        t = m.group(1)
        if t.split()[0] in STOP or len(t) < 3:
            continue
        if m.start() == 0 or line[:m.start()].rstrip().endswith((".", "!", "?", "#", "|", "-", "*")):
            if " " not in t and not t.isupper():
                continue  # sentence-initial single word: probably not a term
        yield t
    for m in ASSUMED.finditer(line):
        yield m.group(0).lower()


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("file")
    ap.add_argument("--min-count", type=int, default=1)
    a = ap.parse_args()

    seen = OrderedDict()
    in_fence = False
    in_front = False
    for n, raw in enumerate(Path(a.file).read_text().splitlines(), 1):
        if n == 1 and raw.strip() == "---":
            in_front = True
            continue
        if in_front:
            in_front = raw.strip() != "---"
            continue
        if raw.lstrip().startswith("```"):
            in_fence = not in_fence
            continue
        if in_fence or raw.lstrip().startswith(("![", "|---", "| ---")):
            continue
        for t in candidates(raw):
            key = t.strip("`").lower()
            if key not in seen:
                i = raw.lower().find(key)
                after = raw[i + len(key):] if i >= 0 else ""
                seen[key] = {"term": t, "line": n, "count": 0, "defined_here": bool(DEFINING.search(after)),
                             "text": raw.strip()[:110]}
            seen[key]["count"] += 1

    rows = [v for v in seen.values() if v["count"] >= a.min_count]
    print(f"{'line':>4}  {'uses':>4}  def?  term  —  first-use context")
    for v in rows:
        flag = " yes" if v["defined_here"] else "  - "
        print(f"{v['line']:>4}  {v['count']:>4}  {flag}  {v['term']}  —  {v['text']}")


if __name__ == "__main__":
    main()
