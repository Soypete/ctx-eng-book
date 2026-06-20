---
name: code-examples
description: Author code examples that teach a module's specific point, then hand off to code-audit for validation. Use when a module needs a worked example, snippet, or runnable demo to illustrate a concept. Authors examples; does NOT validate them (code-audit does that).
---

# Code Examples Skill

## Purpose
Write code examples that **teach the specific point** a module is making — not generic boilerplate. Each example is tied to a beat in the module's outline and is built to be idiomatic, minimal, and pedagogically clear. After authoring, this skill hands the example to the `code-audit` skill for syntax/build validation.

Authoring and validation are deliberately separate: this skill creates examples; `code-audit` checks them.

## Trigger
Use this skill when:
- A `{module}.outline.md` lists items under "Examples needed".
- The user asks for a code example, snippet, or runnable demo for a concept.
- A `write` session surfaces a point that's clearer with code.

## Inputs
- The module's outline ("Examples needed") and draft.
- The teaching goal: which beat/claim the example must illustrate.
- Book conventions: examples should be syntactically correct, idiomatic, well-commented for teaching, and consistent with any existing examples directory.

## Workflow

### 1. Pin the teaching goal
State in one line: "This example must make the reader understand {X}." If the goal is fuzzy, ask — a code example with no teaching target is noise.

### 2. Choose the smallest example that teaches it
- Prefer the minimal snippet that demonstrates the point; cut unrelated setup.
- Pick the language that fits the book's running examples (Go is the project's primary; match neighbors).
- Comment for *why*, not just *what* — the comments carry the lesson.

### 3. Write the example into the module
Place the fenced code block in `{module}.md` at the beat it supports, with a one-line lead-in and a one-line takeaway after it. Keep prose minimal — the `write` skill owns surrounding narrative.

### 4. Make it real if it should run
If the example claims to run, create a buildable version under an examples directory (e.g. `cmd/authorpedro/books/ctx-eng-book/examples/{chapter}/{slug}/`) so it can actually compile/execute, not just look right.

### 5. Hand to code-audit
Invoke the `code-audit` skill on the touched module(s) to validate every block (gofmt/go build, etc.). Fix anything it flags. Do not consider the example done until `code-audit` passes.

## Outputs
- Code block(s) in `{module}.md`, each tied to a teaching goal.
- Optional runnable source under the examples directory.
- A clean `code-audit` report.

## Tools
- read: outline, draft, existing examples
- write / edit: add examples and runnable sources
- bash: run/compile to confirm behavior
- skill: `code-audit` for validation

## Hand-off
- `code-audit` validates syntax/build (always run it).
- `editor` checks the example earns its place and the lead-in/takeaway flow.
