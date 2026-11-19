# Substack Post Guidelines

Seed for `book/substack/GUIDELINES.md`. On first run, copy this file there; after that
the repo copy is canonical and Miriah may edit it — always read the repo copy.

## Voice
- First person, direct, opinionated, practitioner-to-practitioner. Match
  `blog-posts/how-everyone-is-using-ai-wrong.md` and the published Substack
  (https://substack.com/@soypetetech): short paragraphs, concrete failures, a clear
  stance, light humor, no hype words ("revolutionary", "game-changer", "unlock").
- The post stands alone. A reader who has never seen the book must follow it.
- Introduce every concept before you use it. Terms of art (agent, harness, manifest,
  top-k, the companion example) get a plain one-sentence explanation the first time
  they appear. Use one name per thing (not "harness" in one paragraph and "host" in the
  next).
  Replace "as we saw in Chapter 4" with one sentence of context or a link.

## Shape (900–1,600 words)
1. **Hook** — a concrete failure, surprise, or claim in the first 3 sentences.
2. **The problem** — why the obvious approach breaks, in the reader's terms.
3. **The idea** — the module's core argument, one level of abstraction above code.
4. **Visual** — at least one, in the form that fits the argument (see below).
5. **Code** — at least one runnable or near-runnable snippet (see below).
6. **Guidelines** — a `## Guidelines` section: 3–7 imperative, testable rules the
   reader can apply this week ("Authorize before you rank", not "Think about
   security").
7. **Close** — one-paragraph takeaway plus a pointer to the series/next post.

## Required elements
- **Visual:** at least one image that shows the post's argument, in the form that
  fits it. Choose the form from `.opencode/skills/substack-prep/references/visual-forms.md`
  and name it with one sentence of why in `editor-review.md`:
  - **Scene** — the failure as a user would see it (a chat, a terminal, a dashboard).
  - **Annotated artifact** — a real object (JSON, a log, a prompt, a tool result)
    with callouts on the lines that matter.
  - **Timeline** — when timing or ordering is the bug (retries, timeouts, staleness).
  - **Comparison** — two panels side by side when the argument is "X is not Y".
  - **Comic** — 2–3 panels with PedroBot when the point is a sequence with a punchline.
  - **Chart** — only when the post has a real number the module supports (use the
    `dataviz` skill).
  - **Flowchart** — only when the argument really is a pipeline of decisions.

  **No two consecutive posts lead with the same form**, and a flowchart needs a reason
  the other forms can't serve. Every visual has a one-line caption saying what to
  notice, invented or redacted data only, and alt text that states the point.

  Files: `book/substack/{module}/visual-N.png` (posts before 2026-11-18 used
  `diagram-N.png`; leave those). Embed with the absolute raw GitHub URL
  (`![alt](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/{module}/visual-N.png)`),
  never a ```mermaid block (Substack can't render it).
  - *Illustrated forms* (scene, artifact, timeline, comparison, comic) are 1456×816
    boards on the Claude Design canvas
    (https://claude.ai/artifact/DeaWCkASLMUsV8NJb9VWJr), titled `{module}/visual-N`,
    rendered with `python3 scripts/import_thumbnails.py <pulled dir> --module {module}`.
    They use the SoyPeteTech design system
    (https://claude.ai/artifact/HfZp11Wm9mGJDXQ4LdKq5H): white `surface-raised` or plum
    ground, Fredoka for any headline, JetBrains Mono for every label, flat fills,
    6px black outlines, one hard 16px drop shadow per image, `danger` red always paired
    with a word or ✕. The same role colors as diagrams: plum = the system/harness, sky
    = data sources, PedroBot orange = the model, teal = allowed/included, red = denied.
  - *Flowcharts* keep the Mermaid flow: source in `visual-N.mmd` with
    `book/substack/mermaid-classes.mmd` appended and every node given a role (`system`,
    `source`, `model`, `ok`, `neutral`, `denied`, at most one `focus`), rendered with
    `npx -y @mermaid-js/mermaid-cli -i visual-N.mmd -o visual-N.png -c book/substack/mermaid-theme.json -C book/substack/mermaid-css.css -b white -s 2`.
  - Commit the PNG and its source (`.mmd`, or the board, which lives on the canvas).
- **Code:** the full, runnable code lives in the book repo at
  `book/examples/{module}/` (its own Go module or package, gofmt/go vet clean, with a
  test or `main` that proves what the post claims, and a short README). The post shows
  only the key lines (≤15, fenced and language-tagged, commented for *why*) and links
  to the folder: `https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/{module}`.
  Never paste the whole program into Substack. Go is the book's primary language; use
  Python or SQL when the audience point is data/ML.
- **Guidelines:** the section above. Each rule must trace to an argument in the post.
- **Thumbnail:** `book/substack/{module}/thumbnail.png`, 1456×816, designed as a board
  titled `{module}` on the Claude Design canvas
  (https://claude.ai/artifact/DeaWCkASLMUsV8NJb9VWJr) and rendered with
  `scripts/import_thumbnails.py`. Brand tokens are fixed (plum ground with cyan,
  Fredoka/JetBrains Mono, 6px black outlines, a hard drop shadow, PedroBot, logo), but
  the props and PedroBot's pose are designed per post around its argument so
  consecutive covers don't look alike. Front matter carries `mood:` (professor, happy,
  confused, broken, mean, gopher) to choose the PedroBot. Titles stay at three lines
  or fewer on the thumbnail.
- **Links:** absolute URLs only (Substack can't resolve repo-relative paths). Cite
  sources inline with a link; no bibliography dump.

## Don'ts
- No unsupported numbers. If the module marks a claim as inference, keep it worded as
  the author's argument.
- No "In this post we will…" openers, no "In conclusion".
- Don't paste the module. Cut sections that don't serve the single argument of the post.
