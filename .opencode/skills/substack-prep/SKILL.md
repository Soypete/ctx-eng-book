---
name: substack-prep
description: Turn book modules into Substack posts for @soypetetech — keep a three-times-weekly (Mon/Wed/Fri) publishing schedule (one post per module), condense a module into a standalone draft with a diagram, code, and a Guidelines section, get three reader reviews (full-stack, AI, and data engineer personas) for Miriah to annotate, then one-shot the rewrite with Codex on the gpt-6-astra model and verify the result. Use whenever the user mentions Substack, the newsletter, posting chapters, a publishing/post schedule, "get chapter N ready to post", persona or audience reviews of a post, or rewriting a post with Astra/Codex — even if they don't say "substack-prep".
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
├── brand/                 # logo, PedroBot moods, thumbnail.html template
├── mermaid-*.{json,css,mmd} # brand diagram theme and role classes
└── {module}/
    ├── draft.md           # condensed post (Claude)
    ├── review-fullstack.md
    ├── review-ai.md
    ├── review-data.md     # each ends with "## Miriah's notes"
    ├── thumbnail.png      # 1456x816 header, rendered from brand/thumbnail.html
    ├── diagram-N.mmd      # Mermaid source (role classes appended)
    ├── diagram-N.png      # exported PNG, linked from the post by raw GitHub URL
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
   `--start YYYY-MM-DD` if the user gives a start date; default is the next Mon/Wed/Fri).
   It lists every module in book order — currently ~87 posts, roughly 29 weeks at
   three posts a week. Rerunning it later is safe: statuses/notes persist and unscheduled
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
     `npx -y @mermaid-js/mermaid-cli -i diagram-1.mmd -o book/substack/{module}/diagram-1.png -c book/substack/mermaid-theme.json -C book/substack/mermaid-css.css -b white -s 2`
     First append `book/substack/mermaid-classes.mmd` and give every node a role class
     (system, source, model, ok, neutral, denied, one focus) per the SoyPeteTech design
     system's Diagrams section (https://claude.ai/artifact/HfZp11Wm9mGJDXQ4LdKq5H).
     If export fails, keep the Mermaid source and note "PNG export pending" in the
     schedule row rather than blocking.
   - *Code*: put the full, runnable code in `book/examples/{module}/` (reuse the
     module's or existing examples where possible; Go by default, Python/SQL when the
     point is data-shaped) with a test or `main` that proves the post's claim and a short
     README. Validate it the way the `code-audit` skill does (gofmt, `go vet`,
     `go test`; `python -m py_compile`). The post shows only the key lines (≤15) and
     links to `https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/{module}`.
   - *Guidelines*: a `## Guidelines` section of 3–7 imperative, testable rules, each
     traceable to something the post argues.
   - *Thumbnail*: pick a PedroBot mood that fits the post's point and put it in front
     matter as `mood:` — `professor` (explaining), `happy` (a working system), `confused`
     (missing/ambiguous context), `broken` (failures, loops, outages), `mean`
     (adversarial or misbehaving models), or `gopher` (Go-specific posts). Render:
     `python3 .opencode/skills/substack-prep/scripts/make_thumbnail.py {module}`
     → `book/substack/{module}/thumbnail.png` (1456×816, SoyPeteTech design system:
     plum ground, title in Fredoka, PedroBot on the cyan logo disc). Look at the PNG; if
     the title wraps past three lines, shorten the title rather than shrinking type.
4. **Editor pass** before the personas see it: follow the repo's `editor` skill
   (`.claude/skills/editor/SKILL.md` in Claude Code, `.opencode/skills/editor/SKILL.md`
   elsewhere; Post mode) against the draft and write
   `book/substack/{module}/editor-review.md`. For a post, cohesion means: every concept
   is introduced before it is used (trace the order: list each term of art — e.g.
   agent, harness, manifest, top-k, the companion example — and the first sentence
   that explains it); the post stands alone for a reader who hasn't seen the book or
   the previous post; and it bridges from the previous post and to the next. Fix flow
   and cohesion findings in the draft before step 5. Readers can't ask what "harness"
   means; if the draft can't explain it in a sentence, it isn't ready for review.
5. **Review** with the three personas in `references/personas.md` — run them as three
   independent subagents in parallel when possible (each sees only the draft, the
   persona, and the template; not the module and not the other reviews). Each writes
   `review-{fullstack,ai,data}.md` ending with an empty `## Miriah's notes` section.
6. **Update** the schedule row to `awaiting-notes` and stop. Tell Miriah which files to
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
Run the editor pass again on `post.md` (Astra can reintroduce undefined terms or drop a
bridge); update `editor-review.md` and resolve cohesion findings before marking ready.
Re-render the thumbnail (`make_thumbnail.py {module}` reads `post.md`'s title), then
check `post.md` against GUIDELINES.md and the module:
- front matter present; length in range; post stands alone (no "as we saw in Chapter N").
- ≥1 diagram as a Markdown image link to its absolute raw GitHub URL
  (`https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/{module}/diagram-N.png`)
  with a caption; no ```mermaid block in the post; the `.mmd` and PNG both exist and are
  committed (re-export if Astra changed the diagram).
- the code excerpt is ≤15 lines and matches `book/examples/{module}/`, which passes
  gofmt/vet/tests; the post links to that folder rather than pasting the full program.
- `## Guidelines` section with 3–7 imperative rules.
- every link is absolute; every number/citation also appears in the module (no new
  claims crept in during rewrite); each item in Miriah's notes was applied.
- `git diff --stat` shows Astra changed only `post.md`, `astra-summary.md`, and (if the diagram changed) `diagram-N.mmd`.

Fix small mechanical issues yourself (a relative link, a PNG export). For substantive
misses (a note ignored, a new unsupported claim), list them and ask whether to rerun
Astra with a sharper prompt or hand-fix. When it passes, set status `ready` and give
Miriah the path to `post.md`, `thumbnail.png` (Substack header / social image), and the
diagram PNGs.

Moving a post to `scheduled`/`published` is Miriah's call; update the row when she says
it's on Substack's calendar (dates for those rows are then locked by the schedule
script).

## Notes
- Do not commit or push unless asked.
- Keep GUIDELINES.md as the single source of style truth; if Miriah's notes keep
  repeating the same correction across posts, suggest adding it to GUIDELINES.md.
- Related skills: `code-audit` (validate snippets), `editor` (if a module itself has
  flow problems the reviews expose — fix the book, not just the post).
