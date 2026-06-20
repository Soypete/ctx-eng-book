---
name: plan
description: Maintain and reshape the book's structure — module outlines, ordering, splits/merges — and map each module to the research evidence that supports it. Use when planning what to write, reordering chapters/modules, or checking a module has the evidence it needs before writing.
---

# Planning Skill

## Purpose
Own the *structure* of the book. Decide what each module argues, in what order, and whether the research evidence exists to support it. This is the bridge between `research` (raw evidence) and `write` (drafting).

This skill does NOT write prose and does NOT write detailed evidence notes (that's `research`). It produces outlines and structural decisions.

## Trigger
Use this skill when the user:
- Asks "what should we write next / today?"
- Wants to reorder, split, merge, add, or cut chapters/modules.
- Wants to plan a module's internal beats before writing.
- Wants to check whether a module is "ready to write" (has evidence + a clear argument).

## Inputs
- `chapters.md` — the canonical Part I–V chapter/module outline.
- `research/five-pillars-outline.md` — pillars, reliability goals, failure modes.
- `research/_evidence-ledger.md` — what evidence exists (from `research`).
- The existing draft tree: `cmd/authorpedro/books/ctx-eng-book/chapters/{chapter}/modules/{module}.md`.

## Workflow

### 1. Establish current state
- Read `chapters.md` for the intended structure.
- List existing module files to see what's drafted vs. only outlined.
- Read the evidence ledger to know what's supportable.

### 2. For structural work (reorder / split / merge)
- Propose the change as a diff to `chapters.md` with a one-line rationale per change.
- Check downstream impact: does moving a module orphan a dependency or break a narrative build-up? Flag it.
- Only edit `chapters.md` after the user confirms the reshape.

### 3. For module planning (the common case)
Produce a **module outline** at:
`cmd/authorpedro/books/ctx-eng-book/chapters/{chapter}/modules/{module}.outline.md`

Each module outline contains:
```
# {Module Title}

## Argument (one sentence)
The single claim this module makes.

## Why it matters (reliability goal)
Tie to the pillar's reliability goal + a failure mode it addresses.

## Beats (the path to the conclusion)
1. {beat} — evidence: {ledger row or "GAP"}
2. {beat} — evidence: {ledger row or "GAP"}
...

## Conclusion to reach
What the reader should believe by the end — this is the target `write` drives toward.

## Examples needed
- {what code/diagram example would teach beat N} (feeds `code-examples`)

## Readiness
ready | needs-research ({list gaps}) | needs-decision
```

### 4. Recommend the next action
When asked "what next?", return a ranked shortlist: which module is most ready to write, what's blocking the others, and whether the next step is `research`, `plan`, `write`, or `code-examples`.

## Outputs
- Updated `chapters.md` (only on confirmed structural changes)
- `{module}.outline.md` files
- A ranked "what to do next" recommendation

## Hand-off
- `write` reads `{module}.outline.md` (esp. "Conclusion to reach" + "Beats") to drive Socratic drafting.
- `research` is invoked to close any "GAP" evidence before a module is marked `ready`.
- `code-examples` reads "Examples needed".
