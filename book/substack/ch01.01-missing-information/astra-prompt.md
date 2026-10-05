You are rewriting a technical Substack post for Miriah Peterson (@soypetetech).
Work in the repository at /Users/soypete/code/misc/ctx-eng-book. Write the final post to book/substack/ch01.01-missing-information/post.md. The only other file you may
edit is book/substack/ch01.01-missing-information/diagram-1.mmd, and only if the diagram must change.

Read, in this order:
1. book/substack/GUIDELINES.md — the style and required elements. These are hard
   requirements.
2. book/substack/ch01.01-missing-information/draft.md — the current draft (if post.md already exists,
   rewrite from post.md and use draft.md as the reference for Miriah's voice).
3. book/substack/ch01.01-missing-information/review-fullstack.md, review-ai.md, review-data.md — three
   reader reviews. The "## Miriah's notes" section at the end of each review is the
   author's direction and overrides the reviewer when they conflict.
4. book/substack/ch01.01-missing-information/editor-review.md (if present) — flow and cohesion findings;
   fix every one, and introduce each concept before it is used.
5. book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.01-missing-information.md — the book module the post is condensed from. Use it to check
   facts; do not copy it wholesale.

Rewrite the draft so it works for all three readers at once: keep the full-stack
reader's path to code, the AI engineer's demand for evidence and failure modes, and
the data engineer's concern for provenance, freshness and governance. Apply every
item in Miriah's notes. Where reviewers disagree and Miriah's notes are silent,
prefer the change that makes the post more concrete.

The post must keep: at least one diagram embedded as a Markdown image link to its
absolute raw GitHub URL (never a ```mermaid block) with a one-line caption; if you
change what the diagram shows, edit book/substack/ch01.01-missing-information/diagram-1.mmd too; a short code
excerpt (15 lines or fewer, fenced and language-tagged) taken from
book/examples/ch01.01-missing-information/ with a link to
https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.01-missing-information for the full
program (never paste the whole program), and a "## Guidelines"
section of 3–7 imperative rules. Keep any numbers and citations exactly as the
module supports them; do not add new ones. Use absolute URLs only.

Start the file with:
---
title: Before You Blame the Model, Check Three Gates
subtitle: <one sentence>
module: ch01.01-missing-information
mood: confused
scheduled: 2026-10-09
---
Then the post body in Markdown. When done, reply with a 3-line summary of what you
changed.
