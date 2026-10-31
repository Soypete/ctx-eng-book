You are writing a technical Substack post for Miriah Peterson (@soypetetech),
condensed from one module of her book on context engineering.
Work in the repository at /Users/soypete/code/misc/ctx-eng-book. Write the post to book/substack/ch03.02-context-windows/post.md.

Read, in this order:
1. book/substack/GUIDELINES.md — style and required elements; hard requirements.
2. blog-posts/how-everyone-is-using-ai-wrong.md and
   book/substack/ch01.03-context-failure-case-studies/post.md — calibration for
   Miriah's voice and the series' house style (bold term definitions, a failure as the
   hook, decision records, concrete guidelines).
3. book/chapters/ch03-attention-is-all-you-need/modules/ch03.02-context-windows.md and its .outline.md if present — the source. Pick the module's ONE
   strongest argument; cut everything that doesn't serve it. Keep the module's own
   evidence, and keep anything the module marks as inference worded as Miriah's argument.
4. book/substack/ch03.01-tokens-embeddings-attention/post.md if it exists — the previous post, for a one-sentence bridge.

Write for three readers at once: a full-stack engineer (path to code), an AI engineer
(evidence, failure modes) and a data engineer (provenance, freshness, governance).
The post must stand alone for a reader who hasn't seen the book.

Drew Breunig's context-engineering writing is now an approved source:
research/drew-breunig-context-engineering.md lists his posts with URLs and the claims
each one supports. Cite one of his posts (inline link to the post URL) where it
genuinely strengthens the argument — at most two per post, never as decoration.
Use only the claims listed in that file, worded as he states them. For everything
else, keep numbers and citations exactly as the module supports them; add no others.

Thumbnail: write book/substack/ch03.02-context-windows/thumbnail.html, a self-contained 1456x816 page that
headless Chrome screenshots as the Substack header. Requirements:
- SoyPeteTech brand: copy the :root tokens and Google Fonts link from
  book/substack/brand/thumbnail.html (plum #6f2d5f, cyan #00cfff, teal #00c89a,
  sky #26abe7, PedroBot orange #f18d2b, ink #1b0b17, white); Fredoka for the title,
  JetBrains Mono for the eyebrow; 6px solid black outlines and hard 16px black drop
  shadows on shapes; no gradients, no blur.
- Must show: eyebrow "Context Engineering · 10"; the post title (3 lines max,
  at least 54px); the logo ../brand/spt-logo-standard-bluepurple.png with
  "@soypetetech"; PedroBot ../brand/pedrobot-confused.png (other moods available:
  professor, happy, confused, broken, mean).
- Do NOT reuse the shared template's composition (title left, PedroBot in a cyan disc
  right, teal band at the bottom). This post's art direction: white ground (ink text); a thick outlined window frame on the right cropping faded lines of text (some lines cut off at the edges); PedroBot peeking around the frame; title left. Build any props
  with plain HTML/CSS or inline SVG. Keep text clear of overlaps and inside a 64px
  safe margin; body { margin:0; width:1456px; height:816px; overflow:hidden }.

Diagram: if book/substack/ch03.02-context-windows/diagram-1.mmd does not exist, create it: a Mermaid
flowchart of the post's single argument, with the contents of
book/substack/mermaid-classes.mmd appended and every node given a role class
(system, source, model, ok, neutral, denied, at most one focus). Reference it in the
post as https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch03.02-context-windows/diagram-1.png
(I will export the PNG). Never put a ```mermaid block in the post.

Code: the post's excerpt (15 lines or fewer, fenced, language-tagged, commented for
why) must come from book/examples/ch03.02-context-windows/, and the post links to
https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch03.02-context-windows. Create that folder: a small runnable program that proves the post's claim. Go by default (its own go.mod with module path github.com/Soypete/ctx-eng-book/book/examples/ch03.02-context-windows, a _test.go that asserts the claim, gofmt and go vet clean, run go test); Python only when the point is data/ML-shaped (standard library only, a test runnable with python3 -m unittest). Add a short README.md. 

Close with one paragraph and a pointer to the next post in the series (Compaction, Attention, and the Scaffolding Tax — module ch03.03-compaction-scaffolding-tax),
linking https://substack.com/@soypetetech. Use absolute URLs only.

Start post.md with:
---
title: <title, at most 72 characters>
subtitle: <one sentence>
module: ch03.02-context-windows
mood: confused
scheduled: 2026-11-13
---
Then the body in Markdown, 900–1,600 words. You may create or edit only:
book/substack/ch03.02-context-windows/post.md, book/substack/ch03.02-context-windows/diagram-1.mmd,
book/substack/ch03.02-context-windows/thumbnail.html, and files under book/examples/ch03.02-context-windows/.
When done, reply with a 3-line summary of what you wrote.
