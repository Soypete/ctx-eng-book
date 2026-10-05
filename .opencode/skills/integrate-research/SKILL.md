---
name: integrate-research
description: Thread a research note from research/ into the book — pick the chapter modules whose arguments it actually changes, draft short evidence-backed subsections in each, update the matching .outline.md beats, add evidence-ledger and RESEARCH_LEDGER rows, and capture the durable claims to the wiki. Use whenever the user wants a research note, paper, or source "worked into", "threaded through", "woven into", or "applied to" the chapters, asks which modules a note should change, or says something like "now put this research into the book". Use this rather than `write` when the author wants drafted prose they will edit afterwards; use `research` first if the source has no note yet.
---

# Integrate Research Skill

## Purpose
Take a research note that already exists under `research/` and make the book use it. The
output is a set of small, targeted manuscript edits — not a summary of the note pasted into
several chapters — plus the bookkeeping that lets the other skills (`draft-plan`, `write`,
`editor`) see what changed and what still needs the author's attention.

This skill sits between `research` (which produces the note) and `editor` (which reviews flow).
Unlike `write`, it drafts prose. The author has chosen this: the draft is a starting point
they will revise, so write in the book's voice, keep it short, and make every judgment call
visible in the final report so it is easy for them to push back on.

## Inputs
- The research note (path given by the user, or the most recently added file in `research/`).
- `chapters.md` — the book's chapter/module structure and thesis.
- `book/chapters/**/modules/*.md` and their `*.outline.md` partners.
- `research/_evidence-ledger.md` — claim rows (format below).
- `book/RESEARCH_LEDGER.md` — per-module status table, research queue, open questions.

## Workflow

### 1. Read the note and separate evidence from interpretation
Read the whole note. List the claims it can support and tag each one:
- **evidence** — the source itself demonstrates or states it (keep the locator: section,
  page, arXiv id, URL).
- **inference** — the note's or the book's interpretation of the source.
- **unsupported** — asserted in the note with no source, locator, or link (common in older
  notes: e.g. a benchmark number with no paper attached).

Why: the book's credibility rests on not inflating what sources show. Chapter prose may use
inferences, but must word them as the book's argument ("this suggests", "the transferable
pattern is"), never as the source's finding. Unsupported claims do not go into chapters as
fact — flag them for the author instead.

When a claim rests on a volatile external figure (a project's self-reported benchmark,
a version number, a market ranking) and you have web access, check the current source
before relying on it — such numbers are often revised or retracted. Record what you found
and the date checked.

Also note the source's stated limitations and failure modes; these usually make the most
useful chapter content, because they tell the reader where the pattern stops working.

### 2. Choose modules by argument, not by keyword
Search for candidate modules (`grep -ril` over `book/chapters` for the note's key concepts,
plus `chapters.md`). Then read each candidate's outline and the relevant section of its
manuscript. Keep a module only if the note does at least one of:
- supplies evidence for a claim the module currently makes without support,
- sharpens or corrects a boundary the module draws,
- adds a failure mode or limitation the module's reader needs,
- resolves an "open gap" listed for the module in `book/RESEARCH_LEDGER.md`.

Before adding to a module, check whether the material was deliberately cut: read the
outline's `## Readiness` line and search `book-edit-audit.md` for the module. If an
editorial pass removed vendor inventory, volatile numbers, or a specific source, adding it
back undoes the author's decision — leave it out and mention it in the report instead.

A module that merely mentions the same topic is not a target. Most notes justify 2–6 modules;
a narrow or purely historical note may justify one or none, and saying so is a valid outcome.
Give each chosen module a **different angle** of the note — if two modules would get the same
paragraph, the claim belongs in the earlier one, and the later one should reference it.

Before editing, write the placement map down (in your response or working notes):

| Module | Section / beat it attaches to | Claim it adds | Evidence or inference | Why this module |

### 3. Draft the manuscript edits
For each module, read the full module file (and skim the neighbors in `chapters.md` order) so
the addition fits the existing argument rather than restating it.

- Prefer a new `###` subsection inside the existing `##` section it extends, or a few
  paragraphs appended to that section. Do not create a new `##` section unless the outline
  gets a matching new beat.
- Keep each addition short: typically 1–4 paragraphs, optionally a table if it carries a
  real comparison. The note holds the detail; the chapter holds the argument.
- Match the module's existing style: heading levels, line wrapping (some modules hard-wrap
  around 100 columns, some use one line per paragraph — follow the file), tone, terminology.
- Name the source in prose where it is evidence ("SIFT reports…", "Bradley and Terry's model…")
  and link the research note with a correct **relative** path from the module file, e.g.
  `[the ranked context selection research note](../../../../research/ranked-context-selection.md)`
  from `book/chapters/chNN-…/modules/`. Chapter-level files (`book/chapters/chNN-….md`) are two
  levels shallower: `../../research/…`.
- Do not invent quotes, numbers, page numbers, or citations. If the note lacks a locator, say
  so in the report rather than fabricating one.
- End with the operational consequence for the reader (what to build, test, or refuse), since
  this book is for practitioners building reliable AI systems.

### 4. Keep outlines aligned
Repo convention: the `.outline.md` is the plan and the prose file "should not introduce an
unplanned section". For every module edited, update its outline so the new material appears
as a beat/sequence item in the right position, renumbering as needed. Keep the outline's
existing format (some use a numbered `## Beats` list, some a one-line `## Sequence`). Leave
`## Readiness` alone unless the edit changes readiness. Chapter-level files without an outline
(e.g. `ch00`) are the documented exception — do not invent one.

### 5. Update the ledgers
**`research/_evidence-ledger.md`** — append one row per claim the chapters now rely on:

```
## {Pillar or chapter theme} — {Claim the book makes}

- **Source:** {authors} ({year}) — "{title}"
- **Quote:** "{verbatim text from the source or note}"   ← only if verbatim text exists
- **Locator:** {section/page/arXiv id/URL, or research note path}
- **Supports:** {Chapter N — module name; …}
- **Strength:** strong | suggestive | anecdotal
- **Counterpoint:** {limitation or tension, if any}

---
```

If there is no verbatim text, write `**Claim:**` instead of `**Quote:**` — never fabricate a
quote. Use `suggestive` for preprints and single-benchmark results, and for book-level
inferences.

**`book/RESEARCH_LEDGER.md`** — for each module touched, update its row in the `## Modules`
table: set `Last pass` to today's date and edit `Open gaps` to mention the new research
note and anything still unresolved (unsupported claims, missing locators, sources that should
be read in full). If a cited primary source has not been reviewed, add it under
`## Research queue` in that module's subsection using the existing entry format (unreviewed
checkbox, why it matters, claim it would support, notes file, empty `Miriah's notes:`).
If something needs the author's decision, add a bullet to `## Questions for Miriah`.

### 6. Capture to the wiki
Capture the durable results — not every sentence. Typically: one `source` per primary source
newly relied on, one `claim` per core claim the book now makes, and one `decision` recording
which modules changed and why.

```bash
wiki capture --title "..." --type claim --content "..." --link supports:<target> --link derived_from:<source>
```

Valid `--type`: `claim`, `contradiction`, `decision`, `entity`, `source`. Valid link
predicates: `derived_from`, `contradicts`, `supports`, `about`, `relates_to`. Invalid values
are rejected — do not guess. If `wiki` is not on PATH, use
`python3 /Users/soypete/code/herdr-wiki-plugin/bin/wiki <args>`. If the user or the task says
not to write to the wiki (e.g. a dry run), write the exact commands to a file instead.

### 7. Verify
- Run `python3 .opencode/skills/integrate-research/scripts/check_links.py <edited files…>`
  to confirm every relative link resolves.
- For each edited module, confirm each new subsection corresponds to an outline beat.
- `git diff --stat` and reread the diff once with fresh eyes: cut repetition across modules
  and any sentence that overstates the source.

### 8. Report
End with:
1. The placement map (final version), including modules considered and rejected, with a reason.
2. Files changed.
3. **For the author** — the judgment calls to review: inferences worded as argument, claims
   flagged as unsupported, any place where the new text tensions with existing text.
4. Wiki captures made (or the file of commands, if dry run).

Do not commit or push unless asked.

## Hand-off
- `editor` reviews the edited modules for flow and cohesion with neighbors.
- `write` can take any drafted subsection the author wants to rework in their own words.
- `research` handles queued primary sources that need full reading.
