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
4. **Diagram** — at least one (see below).
5. **Code** — at least one runnable or near-runnable snippet (see below).
6. **Guidelines** — a `## Guidelines` section: 3–7 imperative, testable rules the
   reader can apply this week ("Authorize before you rank", not "Think about
   security").
7. **Close** — one-paragraph takeaway plus a pointer to the series/next post.

## Required elements
- **Diagram:** Source in `book/substack/{module}/diagram-N.mmd`, exported to
  `book/substack/{module}/diagram-N.png`. Posts embed the committed PNG with a Markdown image link to its absolute raw
  GitHub URL (`![alt](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/{module}/diagram-N.png)`), never a ```mermaid block (Substack can't render it).
  Keep the Mermaid source beside it as `diagram-N.mmd`, and commit both. Every
  diagram has a one-line caption saying what to notice.
  Style every diagram with the SoyPeteTech design system
  (https://claude.ai/artifact/HfZp11Wm9mGJDXQ4LdKq5H, Diagrams section): append
  `book/substack/mermaid-classes.mmd` and assign each node a role — `system` (plum:
  harness, policy), `source` (sky: data stores), `model` (PedroBot orange: the LLM),
  `ok` (teal: included/allowed), `neutral` (white: inputs), `denied` (red, labeled),
  and at most one `focus` (cyan). Render with the brand theme:
  `npx -y @mermaid-js/mermaid-cli -i diagram-N.mmd -o diagram-N.png -c book/substack/mermaid-theme.json -C book/substack/mermaid-css.css -b white -s 2`
- **Code:** the full, runnable code lives in the book repo at
  `book/examples/{module}/` (its own Go module or package, gofmt/go vet clean, with a
  test or `main` that proves what the post claims, and a short README). The post shows
  only the key lines (≤15, fenced and language-tagged, commented for *why*) and links
  to the folder: `https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/{module}`.
  Never paste the whole program into Substack. Go is the book's primary language; use
  Python or SQL when the audience point is data/ML.
- **Guidelines:** the section above. Each rule must trace to an argument in the post.
- **Thumbnail:** `book/substack/{module}/thumbnail.png`, 1456×816, rendered from
  `book/substack/brand/thumbnail.html` with `make_thumbnail.py`. Front matter carries
  `mood:` (professor, happy, confused, broken, mean, gopher) to choose the PedroBot.
  Titles stay at three lines or fewer on the thumbnail.
- **Links:** absolute URLs only (Substack can't resolve repo-relative paths). Cite
  sources inline with a link; no bibliography dump.

## Don'ts
- No unsupported numbers. If the module marks a claim as inference, keep it worded as
  the author's argument.
- No "In this post we will…" openers, no "In conclusion".
- Don't paste the module. Cut sections that don't serve the single argument of the post.
