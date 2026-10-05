#!/usr/bin/env python3
"""Build or refresh book/substack/SCHEDULE.md: one Substack post per book module.

Posts follow book order (ch00 → ch18, modules in numeric order) on a twice-weekly
cadence. Existing rows keep their status, title, and notes; only new modules are
appended, and dates are reassigned to unpublished rows in order so the schedule
stays gap-free when posts slip.

Usage:
  build_schedule.py [--repo DIR] [--start YYYY-MM-DD] [--days tue,thu]

--start defaults to the first publishing day after today.
"""
import argparse
import datetime as dt
import re
from pathlib import Path

DAYS = {"mon": 0, "tue": 1, "wed": 2, "thu": 3, "fri": 4, "sat": 5, "sun": 6}
STATUSES = ["planned", "drafted", "awaiting-notes", "rewriting", "ready", "scheduled", "published"]
HEADER = "| # | Date | Module | Post title | Status | Notes |"
SEP = "|---|------|--------|------------|--------|-------|"


def module_files(repo: Path):
    chapters = repo / "book" / "chapters"
    files = [p for p in chapters.glob("*.md")]
    files += [p for p in chapters.glob("*/*.md")]
    files += [p for p in chapters.glob("*/modules/*.md") if not p.name.endswith(".outline.md")]

    def key(p):
        return [int(n) if n.isdigit() else n for n in re.split(r"(\d+)", p.stem)]

    return sorted(files, key=key)


def title_of(path: Path):
    for line in path.read_text().splitlines():
        if line.startswith("# "):
            return line[2:].strip()
    return path.stem


def publish_days(start: dt.date, weekdays):
    d = start
    while True:
        if d.weekday() in weekdays:
            yield d
        d += dt.timedelta(days=1)


def parse_existing(path: Path):
    rows = {}
    if not path.exists():
        return rows
    for line in path.read_text().splitlines():
        cells = [c.strip() for c in line.strip().strip("|").split("|")]
        if len(cells) == 6 and cells[0].isdigit():
            rows[cells[2].strip("`")] = {"date": cells[1], "title": cells[3], "status": cells[4], "notes": cells[5]}
    return rows


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--repo", default=".")
    ap.add_argument("--start")
    ap.add_argument("--days", default="tue,thu")
    a = ap.parse_args()

    repo = Path(a.repo).resolve()
    out = repo / "book" / "substack" / "SCHEDULE.md"
    out.parent.mkdir(parents=True, exist_ok=True)
    weekdays = {DAYS[d.strip().lower()[:3]] for d in a.days.split(",")}
    start = dt.date.fromisoformat(a.start) if a.start else dt.date.today() + dt.timedelta(days=1)

    existing = parse_existing(out)
    days = publish_days(start, weekdays)
    lines = []
    for i, f in enumerate(module_files(repo), 1):
        mod = f.stem
        row = existing.get(mod, {"title": title_of(f), "status": "planned", "notes": ""})
        if row["status"] in ("scheduled", "published") and row.get("date"):
            date = row["date"]  # locked once it is on Substack's calendar
        else:
            date = next(days).isoformat()
        lines.append(f"| {i} | {date} | `{mod}` | {row['title']} | {row['status']} | {row['notes']} |")

    body = [
        "# Substack Schedule",
        "",
        f"One post per module, {a.days} cadence. Regenerate with",
        "`python3 .opencode/skills/substack-prep/scripts/build_schedule.py` — statuses and notes are kept;",
        "dates of unscheduled posts shift to stay gap-free.",
        "",
        "Statuses: " + " → ".join(STATUSES),
        "",
        HEADER,
        SEP,
        *lines,
        "",
    ]
    out.write_text("\n".join(body))
    print(f"wrote {out} ({len(lines)} posts, first {lines[0].split('|')[2].strip()})")


if __name__ == "__main__":
    main()
