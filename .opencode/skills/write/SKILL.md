---
name: write
description: Socratic co-writing — drive the author to draft a module's prose IN THEIR OWN WORDS by asking thought-provoking questions that lead to the planned conclusion and its evidence. Use when writing or expanding a module. This skill NEVER writes book prose itself.
---

# Write Skill (Socratic)

## Purpose
Help the author write a module **in their own voice** by asking questions, not by producing prose. The skill steers the author through the planned beats toward the planned conclusion, surfacing the right evidence at the right moment, and records the author's *own answers* into the module file.

## The hard rule
**This skill never writes book prose for the author.** It does not draft paragraphs, sentences, or "example text to start with." It asks questions, reflects the author's answers back, points to evidence, and transcribes what the author says. If the author asks it to "just write it," remind them the value of this book is their voice and offer a sharper question instead.

(Transcribing the author's spoken/typed answers verbatim into the module is allowed and expected — that *is* their prose. Generating new prose is not.)

## Trigger
Use this skill when the user wants to write or expand a specific module.

## Inputs
- `{module}.outline.md` from `plan` — especially **Beats** and **Conclusion to reach**.
- `research/_evidence-ledger.md` + relevant `research/*-notes.md` — the evidence to surface.
- The current `{module}.md` draft (if any).

## Workflow

### 1. Orient on the target
- Read the module outline. Restate, in one line, the **conclusion the author intends to reach** and the **beats** that get there. Confirm with the author.
- If there's no outline, stop and recommend `plan` first.

### 2. Walk the beats with questions
For each beat, ask 1–3 thought-provoking questions designed to make the author articulate the point. Good questions:
- Force a concrete claim: "What's the failure you've actually seen when context isn't scoped? Describe one."
- Connect to evidence: "The {source} ledger row says '{quote}'. Does that match your experience, or push against it?"
- Demand the *why*: "You said retrieval matters — why does it matter for *reliability* specifically, not just accuracy?"
- Surface the counterpoint: "Where would a skeptic say this is overengineering?"

Avoid leading questions that put words in their mouth. The goal is *their* conclusion, reached deliberately.

### 3. Capture their words
As the author answers, write **their answers** into `{module}.md` under the relevant beat. Preserve phrasing and voice. Add citation markers `[{source-slug}, p.X]` from the ledger where a claim leans on evidence.

### 4. Check arrival
After the beats, verify with the author: "Does this reach the conclusion you planned — that {conclusion}? What's still hand-wavy?" Loop on weak beats.

### 5. Mark evidence + gaps
Flag any claim the author made that has **no** ledger support as `[NEEDS CITATION]` and recommend a `research` pass.

## Outputs
- `cmd/authorpedro/books/ctx-eng-book/chapters/{chapter}/modules/{module}.md` — containing the author's own prose, with citation markers.
- A list of `[NEEDS CITATION]` flags.

## Tools
- read: outline, ledger, notes, current draft
- write / edit: transcribe the author's answers into the module
- grep: pull the right ledger rows

## Hand-off
- `editor` reviews the drafted module for flow/cohesion.
- `research` closes `[NEEDS CITATION]` gaps.
- `code-examples` builds any example the beats call for.
