---
name: substack-prep
description: Turn book modules into Substack posts for @soypetetech — keep a twice-weekly publishing schedule (one post per module), condense a module into a standalone draft with a diagram, code, and a Guidelines section, get three reader reviews (full-stack, AI, and data engineer personas) for Miriah to annotate, then one-shot the rewrite with Codex on the gpt-6-astra model and verify the result. Use whenever the user mentions Substack, the newsletter, posting chapters, a publishing/post schedule, "get chapter N ready to post", persona or audience reviews of a post, or rewriting a post with Astra/Codex — even if they don't say "substack-prep".
---

# Substack Prep Skill

## Purpose
Ship the book as a Substack series, one post per module, without losing the author's
voice or the book's evidence discipline. The skill does the mechanical and editorial
legwork; Miriah stays in the loop at exactly one point — adding notes to the three
reader reviews — and Astra (Codex `gpt-6-astra`) does the final rewrite from those notes.

Everything lives under `book/substack/`:

```
book/substack/
├── SCHEDULE.md            # all posts, dates, statuses (generated, then maintained)
├── GUIDELINES.md          # style + required elements (seeded from references/)
└── {module}/
    ├── draft.md           # condensed post (Claude)
    ├── review-fullstack.md
    ├── review-ai.md
    ├── review-data.md     # each ends with "## Miriah's notes"
    ├── diagram-N.png      # exported Mermaid
    ├── astra-prompt.md    # self-contained rewrite prompt
    ├── astra-summary.md   # Astra's final message
    └── post.md            # the post to paste into Substack
```

## Work out where we are
The work is resumable and status-driven; read `book/substack/SCHEDULE.md` first.
- No schedule yet → **Phase 0**.
- User names a module/chapter, or says "next" → find its row and run the phase its
  status calls for: `planned` → Phase 1; `awaiting-notes` with notes filled in → Phase 2;
  `rewriting` → check whether `post.md` exists, then Phase 3.
- "Prep the next N" → Phase 1 for the next N `planned` rows in date order.
- If notes are still empty for an `awaiting-notes` post, say so and stop — the rewrite
  without Miriah's notes is the thing this workflow exists to avoid. Proceed only if she
  explicitly says to rewrite without notes.

## Phase 0 — Set up
1. `python3 .opencode/skills/substack-prep/scripts/build_schedule.py` (add
   `--start YYYY-MM-DD` if the user gives a start date; default is the next Tue/Thu).
   It lists every module in book order — currently ~87 posts, roughly 44 weeks at
   twice weekly. Rerunning it later is safe: statuses/notes persist and unscheduled
   dates shift so a slipped post doesn't leave a gap.
2. If `book/substack/GUIDELINES.md` doesn't exist, copy
   `references/post-guidelines.md` there.
3. Show the user the first few weeks and the end date; ask if the cadence/start is right
   only if they haven't said.

## Phase 1 — Draft and review (per module)
1. **Read** the module, its `.outline.md`, `book/substack/GUIDELINES.md`, and the
   previous post's `draft.md`/`post.md` if one exists (continuity, no repeated hook).
2. **Draft** `draft.md` following GUIDELINES.md: one argument, standalone, 900–1,600
   words, in Miriah's voice (read `blog-posts/how-everyone-is-using-ai-wrong.md` once
   per session for calibration). Front matter: `title`, `module`, `scheduled`.
3. **Make sure the three required elements exist** — this is the part most likely to be
   missing, because most modules have Mermaid but few have code:
   - *Diagram*: reuse the module's Mermaid where it serves the post's single argument;
     otherwise draw a simpler one. Export each to PNG:
     `npx -y @mermaid-js/mermaid-cli -i diagram-1.mmd -o book/substack/{module}/diagram-1.png -b white -w 1400`
     If export fails, keep the Mermaid source and note "PNG export pending" in the
     schedule row rather than blocking.
   - *Code*: reuse code from the module or `book/examples/`; if none, write the smallest
     snippet that makes the post's point concrete (Go by default; Python/SQL when the
     point is data-shaped). Validate it the way the `code-audit` skill does (gofmt /
     `go vet` in a temp module, `python -m py_compile`, etc.).
   - *Guidelines*: a `## Guidelines` section of 3–7 imperative, testable rules, each
     traceable to something the post argues.
4. **Review** with the three personas in `references/personas.md` — run them as three
   independent subagents in parallel when possible (each sees only the draft, the
   persona, and the template; not the module and not the other reviews). Each writes
   `review-{fullstack,ai,data}.md` ending with an empty `## Miriah's notes` section.
5. **Update** the schedule row to `awaiting-notes` and stop. Tell Miriah which files to
   annotate, and give a 3-line digest per persona (verdict + top edit) so she can decide
   where to spend her notes.

## Phase 2 — Astra rewrite (one-shot)
1. Confirm at least one `## Miriah's notes` section has content (see above).
2. Fill `references/astra-prompt.md` into `book/substack/{module}/astra-prompt.md`. The
   prompt must be self-contained: Astra sees nothing from this conversation.
3. Set status `rewriting`, then run the `codex exec -m gpt-6-astra …` command from the
   template. It can take several minutes: run it in the background (or with a 10-minute
   timeout) and do Phase 1 for another post meanwhile if the user asked for more.
4. If Codex fails (auth, model unavailable, sandbox), report the exact error and stop —
   don't silently rewrite it yourself, since the user chose Astra for this step.

## Phase 3 — Verify and mark ready
Check `post.md` against GUIDELINES.md and the module:
- front matter present; length in range; post stands alone (no "as we saw in Chapter N").
- ≥1 Mermaid diagram with caption and its PNG exists (re-export if Astra changed it).
- ≥1 language-tagged code block ≤40 lines that passes the code-audit checks.
- `## Guidelines` section with 3–7 imperative rules.
- every link is absolute; every number/citation also appears in the module (no new
  claims crept in during rewrite); each item in Miriah's notes was applied.
- `git diff --stat` shows Astra changed only `post.md` (and `astra-summary.md`).

Fix small mechanical issues yourself (a relative link, a PNG export). For substantive
misses (a note ignored, a new unsupported claim), list them and ask whether to rerun
Astra with a sharper prompt or hand-fix. When it passes, set status `ready` and give
Miriah the path to `post.md` plus the image files to upload.

Moving a post to `scheduled`/`published` is Miriah's call; update the row when she says
it's on Substack's calendar (dates for those rows are then locked by the schedule
script).

## Notes
- Do not commit or push unless asked.
- Keep GUIDELINES.md as the single source of style truth; if Miriah's notes keep
  repeating the same correction across posts, suggest adding it to GUIDELINES.md.
- Related skills: `code-audit` (validate snippets), `editor` (if a module itself has
  flow problems the reviews expose — fix the book, not just the post).
