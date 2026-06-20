# Author Pedro — OpenCode Skills

Skills for authoring *Context Engineering: Building Reliable AI Systems*. This
replaces the author-pedro Go TUI agent: the same jobs are now opencode skills.

## The pipeline

```
research ──> plan ──> write ──> editor
   │           │        │          │
   │           └─ maps evidence ── │
   │                    │          │
   └─ evidence ledger ──┘          │
                        │          │
                  code-examples ───┘
                        └─ code-audit (validate)
```

## Skills

| Skill | Job | Writes prose? |
|-------|-----|---------------|
| `research` | Source material → citation notes + `research/_evidence-ledger.md` | No (notes only) |
| `plan` | Module outlines, ordering, reorders; map evidence → modules | No |
| `write` | **Socratic** co-writing — drives the author to write in their own voice | **No — never** |
| `editor` | Review drafts + plan for flow, cohesion, coverage | No (diagnoses) |
| `code-examples` | Author teaching examples, then hand to `code-audit` | Code only |
| `code-audit` | Validate code blocks (syntax/build) | No |

## Core principle
`write` never produces book prose. The book's value is the author's voice, so
`write` only asks thought-provoking questions and transcribes the author's own
answers toward a conclusion that `plan` set and `research` supports.

## Key artifacts
- `chapters.md` — canonical outline (owned by `plan`)
- `research/_evidence-ledger.md` — citable claims/evidence (owned by `research`)
- `.../chapters/{chapter}/modules/{module}.outline.md` — per-module plan
- `.../chapters/{chapter}/modules/{module}.md` — the draft (author's words)

## Note on format
All skills use folder-per-skill `SKILL.md` with frontmatter.
