---
name: research
description: Process source material (papers, articles, books, URLs) into citation-ready notes and a claims/evidence ledger for the book. Use when reading a new source, taking notes on research material, or extracting quotes/evidence that a chapter will cite.
---

# Research Skill

## Purpose
Turn raw source material into two durable artifacts:
1. **Citation-ready notes** — a per-source markdown note in `research/` capturing the argument in the author's own framing.
2. **A claims/evidence ledger** — extracted, quotable evidence mapped to the book's pillars, so the `write` and `plan` skills know what we can actually support.

This skill produces *knowledge and citations*. It does not write book prose. It feeds `plan` (what evidence exists) and `write` (what to cite).

## Trigger
Use this skill when the user:
- Wants to read or take notes on a new source (paper, chapter, blog post, URL).
- Asks to extract citations, quotes, or evidence from material.
- Wants to know "what do we have to support X?"

## Inputs
- Source material: a file under `research/`, a PDF, or a URL.
- The pillar/chapter the source is relevant to (`chapters.md`, `research/five-pillars-outline.md`).
- Existing notes in `research/` (match their voice and structure).

## Workflow

### 1. Locate and read the source
- If it's a URL, fetch it. If a local file, read it fully.
- Skim `chapters.md` and `five-pillars-outline.md` to know which pillar(s) it serves.

### 2. Write a source note (the author's framing)
Create or update `research/{source-slug}-notes.md`. Match the existing note style:
- Title + attribution line (author, year).
- An "Initial Reaction" / framing section in the author's own words.
- Section per key idea, each with the verbatim quote (blockquoted) **and** the author's interpretation beneath it.
- Connect the idea forward to the book thesis ("reliability via engineered context").

> The note is the author thinking on the page, not a neutral summary. Preserve voice.

### 3. Extract a claims/evidence ledger
Append (or maintain) a structured ledger so downstream skills can query evidence. Use `research/_evidence-ledger.md`:

```
## {Pillar} — {Claim the book will make}
- **Source:** {source-slug} ({author}, {year})
- **Quote:** "{verbatim quote}"
- **Locator:** {page / section}
- **Supports:** {which chapter/module argument}
- **Strength:** strong | suggestive | anecdotal
- **Counterpoint:** {any tension / opposing evidence, if noted}
```

### 4. Flag gaps
End your turn with: claims the book wants to make that this source does **not** support, so research planning can target them.

## Outputs
- `research/{source-slug}-notes.md` (voice-preserving note)
- Updated `research/_evidence-ledger.md` (citation/evidence rows)
- A short list of remaining evidence gaps

## Tools
- read / webfetch: read sources
- grep: find related existing notes and ledger entries
- write / edit: create notes and ledger rows

## Hand-off
- `plan` consumes the ledger to map evidence → modules.
- `write` cites ledger rows when driving the author to support a conclusion.
