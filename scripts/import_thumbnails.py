#!/usr/bin/env python3
"""
Render the Claude Design thumbnail boards into book/substack/{module}/thumbnail.png.

The canvas (book/substack/brand/design-assets.json "canvas") keeps one board per
module, titled with the module name. Pull its project/canvas.json and *.dc.html
files into a folder, then:

    python3 scripts/import_thumbnails.py PULLED_DIR [--module ch00-...]

Each board's markup is rendered as a static page in the same headless Chrome
make_thumbnail.py uses. Images the board loads from the canvas asset store
(/_blob/<id>) are swapped for the brand files listed in design-assets.json; add
a row there when a new image is uploaded to the canvas.
"""

import argparse
import json
import re
import struct
import subprocess
import sys
import tempfile
from pathlib import Path

REPO_ROOT = Path(__file__).parent.parent
BRAND = REPO_ROOT / "book/substack/brand"
WIDTH, HEIGHT = 1456, 816


def find_chrome() -> str:
    roots = sorted((Path.home() / ".cache/puppeteer/chrome-headless-shell").glob("*/*/chrome-headless-shell"))
    if roots:
        return str(roots[-1])
    desktop = Path("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome")
    if desktop.exists():
        return str(desktop)
    sys.exit("headless Chrome not found; run `npx -y @mermaid-js/mermaid-cli --help` once to install it")


def static_page(board_html: str, assets: dict[str, str]) -> str:
    """Turn a .dc.html board into a plain page: its <helmet> into <head>, its <x-dc> markup into <body>."""
    helmet = re.search(r"<helmet>(.*?)</helmet>", board_html, re.S)
    body = re.search(r"<x-dc>(.*?)</x-dc>", board_html, re.S)
    if not body:
        raise ValueError("board has no <x-dc> markup")
    markup = body.group(1).replace(helmet.group(0), "") if helmet else body.group(1)

    def local_asset(match: re.Match) -> str:
        blob_id = match.group(1)
        if blob_id not in assets:
            raise ValueError(f"/_blob/{blob_id} is not in design-assets.json")
        return (BRAND / assets[blob_id]).as_uri()

    page = (
        f'<!doctype html><html><head><meta charset="utf-8">{helmet.group(1) if helmet else ""}'
        f"<style>html,body{{margin:0;width:{WIDTH}px;height:{HEIGHT}px;overflow:hidden}}</style>"
        f"</head><body>{markup}</body></html>"
    )
    return re.sub(r"/_blob/([0-9a-f]{32})", local_asset, page)


def png_size(path: Path) -> tuple[int, int]:
    return struct.unpack(">II", path.read_bytes()[16:24])


def render(page: str, out: Path) -> None:
    with tempfile.NamedTemporaryFile("w", suffix=".html", dir=BRAND, delete=False) as f:
        f.write(page)
        filled = Path(f.name)
    try:
        subprocess.run(
            [find_chrome(), "--headless=new", "--hide-scrollbars", "--force-device-scale-factor=1",
             f"--window-size={WIDTH},{HEIGHT}", "--virtual-time-budget=4000",
             "--allow-file-access-from-files", f"--screenshot={out}", filled.as_uri()],
            check=True, capture_output=True,
        )
    finally:
        filled.unlink(missing_ok=True)
    if png_size(out) != (WIDTH, HEIGHT):
        sys.exit(f"{out} rendered at {png_size(out)}, expected {(WIDTH, HEIGHT)}")


def main() -> None:
    parser = argparse.ArgumentParser(description="Render Claude Design thumbnail boards into module thumbnails.")
    parser.add_argument("pulled_dir", type=Path, help="folder holding project/canvas.json and the boards")
    parser.add_argument("--module", help="render only this module's board")
    args = parser.parse_args()

    project = args.pulled_dir / "project" if (args.pulled_dir / "project").is_dir() else args.pulled_dir
    canvas = json.loads((project / "canvas.json").read_text())
    assets = json.loads((BRAND / "design-assets.json").read_text())["assets"]

    for board, frame in canvas["boards"].items():
        module = frame.get("title", "")
        if args.module and module != args.module:
            continue
        module_dir = REPO_ROOT / "book/substack" / module
        if not module_dir.is_dir():
            print(f"skip {board}: no book/substack/{module}")
            continue
        out = module_dir / "thumbnail.png"
        render(static_page((project / board).read_text(encoding="utf-8"), assets), out)
        print(f"wrote {out.relative_to(REPO_ROOT)}")


if __name__ == "__main__":
    main()
