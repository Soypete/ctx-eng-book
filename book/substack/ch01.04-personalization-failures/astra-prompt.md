You are rewriting a technical Substack post for Miriah Peterson (@soypetetech).
Work in the repository at /Users/soypete/code/misc/ctx-eng-book. Write the final post to book/substack/ch01.04-personalization-failures/post.md.

Read, in this order:
1. book/substack/GUIDELINES.md — style and required elements; hard requirements.
2. book/substack/ch01.04-personalization-failures/draft.md — the current draft and the reference for Miriah's voice.
3. book/substack/ch01.04-personalization-failures/review-fullstack.md, review-ai.md, review-data.md — three reader
   reviews. Their "## Miriah's notes" sections are empty; Miriah approved rewriting
   without notes. Where reviewers disagree, prefer the change that makes the post
   more concrete.
4. book/substack/ch01.04-personalization-failures/editor-review.md — fix every flow and cohesion finding; introduce
   each concept before it is used. Where it lists an open author decision, take the
   more conservative option (weaker claim, clearly labeled as the author's argument).
5. book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.04-personalization-failures.md — the book module; check facts against it, don't copy it wholesale.
6. book/substack/ch01.03-context-failure-case-studies/post.md — the previous post, for a one-sentence bridge.

Rewrite so it works for a full-stack engineer (path to code), an AI engineer
(evidence, failure modes) and a data engineer (provenance, freshness, governance).

Drew Breunig's context-engineering writing is now an approved source:
research/drew-breunig-context-engineering.md lists his posts with URLs and the claims
each one supports. Cite one of his posts (inline link to the post URL) where it
genuinely strengthens the argument — at most two per post, never as decoration.
Use only the claims listed in that file, worded as he states them. For everything
else, keep numbers and citations exactly as the module supports them; add no others.

Thumbnail: write book/substack/ch01.04-personalization-failures/thumbnail.html, a self-contained 1456x816 page that
headless Chrome screenshots as the Substack header. Requirements:
- SoyPeteTech brand: copy the :root tokens and Google Fonts link from
  book/substack/brand/thumbnail.html (plum #6f2d5f, cyan #00cfff, teal #00c89a,
  sky #26abe7, PedroBot orange #f18d2b, ink #1b0b17, white); Fredoka for the title,
  JetBrains Mono for the eyebrow; 6px solid black outlines and hard 16px black drop
  shadows on shapes; no gradients, no blur.
- Must show: eyebrow "Context Engineering · 05"; the post title (3 lines max,
  at least 54px); the logo ../brand/spt-logo-standard-bluepurple.png with
  "@soypetetech"; PedroBot ../brand/pedrobot-confused.png (other moods available:
  professor, happy, confused, broken, mean).
- Do NOT reuse the shared template's composition (title left, PedroBot in a cyan disc
  right, teal band at the bottom). This post's art direction: cyan ground (ink text); title in a white card along the bottom-left; PedroBot top-right handing a user the WRONG profile card (two CSS profile cards, one marked with a red 'not you' tag). Build any props
  with plain HTML/CSS or inline SVG. Keep text clear of overlaps and inside a 64px
  safe margin; body { margin:0; width:1456px; height:816px; overflow:hidden }.

Diagram: if book/substack/ch01.04-personalization-failures/diagram-1.mmd does not exist, create it: a Mermaid
flowchart of the post's single argument, with the contents of
book/substack/mermaid-classes.mmd appended and every node given a role class
(system, source, model, ok, neutral, denied, at most one focus). Reference it in the
post as https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch01.04-personalization-failures/diagram-1.png
(I will export the PNG). Never put a ```mermaid block in the post.

Code: the post's excerpt (15 lines or fewer, fenced, language-tagged, commented for
why) must come from book/examples/ch01.04-personalization-failures/, and the post links to
https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.04-personalization-failures. Keep and reuse the existing example in that folder; you may tighten it but it must keep passing its tests.

Close with one paragraph and a pointer to the next post in the series (Stop Saying "The AI Failed" — module ch02.01-ai-marketing),
linking https://substack.com/@soypetetech. Use absolute URLs only.

Start post.md with:
---
title: <title, at most 72 characters>
subtitle: <one sentence>
module: ch01.04-personalization-failures
mood: confused
scheduled: 2026-11-02
---
Then the body in Markdown, 900–1,600 words. You may create or edit only:
book/substack/ch01.04-personalization-failures/post.md, book/substack/ch01.04-personalization-failures/diagram-1.mmd,
book/substack/ch01.04-personalization-failures/thumbnail.html, and files under book/examples/ch01.04-personalization-failures/.
When done, reply with a 3-line summary of what you wrote.
