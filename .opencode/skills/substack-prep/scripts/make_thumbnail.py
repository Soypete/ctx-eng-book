#!/usr/bin/env python3
"""Render a SoyPeteTech-branded Substack thumbnail for one post.

Fills book/substack/brand/thumbnail.html with the post's title, series number,
and a PedroBot mood, then screenshots it with Puppeteer's headless Chrome at
1456x816 (Substack's 16:9 header size).

Usage:
  make_thumbnail.py MODULE [--mood professor|happy|confused|broken|mean|gopher]
                           [--repo DIR] [--title "..."]

Title comes from post.md front matter (falling back to draft.md) unless given.
Mood comes from the front matter `mood:` key, else --mood, else professor.
Writes book/substack/MODULE/thumbnail.png.
"""
import argparse
import html
import re
import subprocess
import sys
from pathlib import Path

MOODS = {"professor", "happy", "confused", "broken", "mean", "gopher"}


def front_matter(path: Path) -> dict:
    if not path.exists():
        return {}
    m = re.match(r"---\n(.*?)\n---", path.read_text(), re.S)
    if not m:
        return {}
    return dict(re.findall(r"^(\w+):\s*(.+)$", m.group(1), re.M))


def series_number(repo: Path, module: str) -> str:
    sched = repo / "book/substack/SCHEDULE.md"
    for line in sched.read_text().splitlines() if sched.exists() else []:
        cells = [c.strip() for c in line.strip("|").split("|")]
        if len(cells) == 6 and cells[2].strip("`") == module:
            return cells[0]
    return ""


def find_chrome() -> str:
    roots = sorted((Path.home() / ".cache/puppeteer/chrome-headless-shell").glob("*/*/chrome-headless-shell"))
    if not roots:
        sys.exit("headless Chrome not found; run `npx -y @mermaid-js/mermaid-cli --help` once to install it")
    return str(roots[-1])


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("module")
    ap.add_argument("--repo", default=".")
    ap.add_argument("--mood")
    ap.add_argument("--title")
    a = ap.parse_args()

    repo = Path(a.repo).resolve()
    post_dir = repo / "book/substack" / a.module
    fm = front_matter(post_dir / "post.md") or front_matter(post_dir / "draft.md")
    title = a.title or fm.get("title")
    if not title:
        sys.exit(f"no title: add front matter to {post_dir}/draft.md or pass --title")
    mood = (a.mood or fm.get("mood") or "professor").lower()
    if mood not in MOODS:
        sys.exit(f"mood must be one of {sorted(MOODS)}")
    character = "gopher-question.png" if mood == "gopher" else f"pedrobot-{mood}.png"

    n = series_number(repo, a.module)
    eyebrow = f"Context Engineering · {int(n):02d}" if n.isdigit() else "Context Engineering"
    size = "76px" if len(title) <= 48 else "64px" if len(title) <= 72 else "54px"

    brand = repo / "book/substack/brand"
    page = (brand / "thumbnail.html").read_text()
    for k, v in {"EYEBROW": eyebrow, "TITLE": title, "TITLE_SIZE": size, "CHARACTER": character}.items():
        page = page.replace("{{" + k + "}}", html.escape(v))
    filled = brand / f".thumb-{a.module}.html"
    filled.write_text(page)
    out = post_dir / "thumbnail.png"
    try:
        subprocess.run([find_chrome(), "--headless", "--hide-scrollbars", "--force-device-scale-factor=1",
                        "--window-size=1456,816", "--virtual-time-budget=4000",
                        f"--screenshot={out}", filled.as_uri()],
                       check=True, capture_output=True)
    finally:
        filled.unlink(missing_ok=True)
    print(f"wrote {out} ({mood}, {eyebrow})")


if __name__ == "__main__":
    main()
