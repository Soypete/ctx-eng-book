---
name: editor
description: Review drafted modules, Substack posts, and the plan for narrative flow and cohesion, above all concepts used before they're introduced, plus argument soundness and coverage gaps. Use when reviewing a draft or post, checking that chapters connect, validating the outline tells a coherent story, or before considering a module or post done.
---

# Editor Skill

## Purpose
Review the book's material — both drafts and plans — for **narrative flow and cohesion**. The editor does not rewrite in its own voice; it diagnoses problems and proposes structural/argument fixes for the author to make (often by re-entering `write` or `plan`).

## Trigger
Use this skill when the user:
- Wants a drafted module reviewed.
- Asks whether chapters/modules connect and build on each other.
- Wants to check the outline tells a coherent story end to end.
- Is about to mark something "done."

## Two review modes

### A. Draft review (a `{module}.md`, or a Substack `draft.md`/`post.md`)
Start with the **concept-introduction trace**: run
`python3 .claude/skills/editor/scripts/term_trace.py <file>` for candidate terms, add the
ordinary words used as terms of art (agent, harness, manifest, top-k), and for each record
the line of first use and the line where it's explained in plain words. Flag terms used
before they're explained, never explained, two names for one thing, and internal
references leaking to readers ("the curriculum", "the Forge repo", "the repository
notes"). A module may rely on a concept an earlier module (in `chapters.md` order) really
introduced; grep to check. A Substack post may not: every post stands alone.

Then check, in order:
1. **Argument soundness** — does the module reach its planned conclusion (from `{module}.outline.md`)? Is each beat actually established, or asserted?
2. **Evidence** — are claims cited? Any `[NEEDS CITATION]` or unsupported leaps? Cross-check against `research/_evidence-ledger.md`.
3. **Flow** — do paragraphs follow logically? Are there abrupt jumps, orphaned points, or repetition?
4. **Cohesion with neighbors** — does it pick up where the previous module left off and set up the next? Check the surrounding modules in `chapters.md` order.
5. **Voice consistency** — does it sound like the author's other notes/modules? Flag passages that read generic/AI-like.
6. **Altitude** — right level of abstraction for the book's audience (practitioners building reliable AI systems).

### B. Plan review (the outline / `chapters.md`)
Check:
1. **Narrative arc** — does Part I→V build an argument, or is it a list of topics?
2. **Coverage** — every pillar's reliability goal + failure modes addressed somewhere?
3. **Ordering** — does any module depend on a concept introduced later?
4. **Redundancy** — overlapping modules that should merge.
5. **Gaps** — promised-but-missing material.

## Output format
```
## Editor Review — {target}

### Verdict
ship | revise | needs-rework

### Concept trace
| term | first use | explained at | status |

### Findings (ordered by severity)
1. [flow|evidence|cohesion|voice|argument] {what's wrong} → {specific fix, and which skill: write/plan/research}
...

### What's working
- {keep these}

### Recommended next step
{write | plan | research | code-examples} on {target}
```

## The rule
Diagnose and direct; don't rewrite prose. When a fix needs new wording, hand back to `write` so the author supplies it in their voice. When it's structural, hand back to `plan`.

## Tools
- read: drafts, outlines, ledger, neighboring modules
- grep: trace concept introduction order across modules
- (no prose generation)

## Hand-off
- `write` for wording/argument fixes (author re-drafts).
- `plan` for reorders/splits/merges.
- `research` for evidence gaps.
