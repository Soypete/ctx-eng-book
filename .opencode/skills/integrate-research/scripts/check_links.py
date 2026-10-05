#!/usr/bin/env python3
"""Check that relative markdown links in the given files resolve to existing paths.

Usage: check_links.py FILE [FILE ...]
Exits non-zero if any relative link is broken. External links (http, mailto) and
pure anchors are skipped.
"""
import re
import sys
from pathlib import Path

LINK = re.compile(r"\[[^\]]*\]\(([^)\s]+)(?:\s+\"[^\"]*\")?\)")


def main(paths):
    broken = 0
    for p in paths:
        path = Path(p)
        if not path.is_file():
            print(f"missing file: {p}")
            broken += 1
            continue
        for lineno, line in enumerate(path.read_text().splitlines(), 1):
            for target in LINK.findall(line):
                if re.match(r"^[a-z]+:", target) or target.startswith("#"):
                    continue
                target_path = target.split("#", 1)[0]
                if not target_path:
                    continue
                if not (path.parent / target_path).exists():
                    print(f"{p}:{lineno}: broken link -> {target}")
                    broken += 1
    if broken:
        print(f"{broken} broken link(s)")
        return 1
    print(f"ok: all relative links resolve in {len(paths)} file(s)")
    return 0


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(__doc__)
        sys.exit(2)
    sys.exit(main(sys.argv[1:]))
