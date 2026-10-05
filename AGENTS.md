# AGENTS.md

## Repository Structure

This is a book manuscript and research repository. There is no application,
build system, package manager, or CI configuration.

- `book/` — manuscript chapters, module outlines, and book examples
- `research/` — reading lists, source notes, and the evidence ledger
- `blog-posts/` — related essays and public-facing drafts
- `chapters.md` — working book outline and thesis

## Working with This Repo

- Edit manuscript and research files directly
- Do not initialize an application or package manager
- Add new manuscript content under `book/`
- Add new research files under `research/`

## Conventions

- Use clear, hierarchical markdown structure
- Include URLs for references
- Group related topics into sections

## Wiki (personal knowledge base)

You have a `wiki` command for a personal LLM-Wiki knowledge base. It prints
results to stdout, so run it and read the output.

- **Before answering** about a topic, search for prior notes/claims:
  `wiki search <query> [--top-k N] [--json]`
- **When you learn something durable**, capture it:
  `wiki capture --title "..." --type claim --content "..." [--link predicate:target]`
- **Stats:** `wiki stats`
- **Reconcile the inbox into the graph** (usually a human does this, not you):
  `wiki organize`

Rules:
- `--type` must be one of: `claim`, `contradiction`, `decision`, `entity`,
  `source`.
- Link `predicate` must be one of: `derived_from`, `contradicts`, `supports`,
  `about`, `relates_to`.
- Invalid types/predicates are rejected — do not guess or coerce a value.
- Captures land in an inbox; do **not** edit wiki pages directly.
- If `wiki` is not on PATH, run
  `python3 /Users/soypete/code/herdr-wiki-plugin/bin/wiki <args>`.
