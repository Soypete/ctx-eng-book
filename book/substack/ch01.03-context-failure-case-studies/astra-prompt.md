You are rewriting a technical Substack post for Miriah Peterson (@soypetetech).
Work in the repository at /Users/soypete/code/misc/ctx-eng-book. Write the final post to book/substack/ch01.03-context-failure-case-studies/post.md. The only other file you may
edit is book/substack/ch01.03-context-failure-case-studies/diagram-1.mmd, and only if the diagram must change.

Read, in this order:
1. book/substack/GUIDELINES.md — the style and required elements. These are hard
   requirements.
2. book/substack/ch01.03-context-failure-case-studies/draft.md — the current draft.
3. book/substack/ch01.03-context-failure-case-studies/review-fullstack.md, review-ai.md, review-data.md — three
   reader reviews. The "## Miriah's notes" sections are empty for this post; Miriah's
   direction is given below instead and overrides the reviewers when they conflict.
4. book/substack/ch01.03-context-failure-case-studies/editor-review.md — flow and cohesion findings;
   fix every one, and introduce each concept before it is used.
5. book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.03-context-failure-case-studies.md — the book module the post is condensed from. Use it to check
   facts; do not copy it wholesale.

Miriah's direction for this rewrite:
- Keep the current code example exactly as scoped: it covers only the freshness
  boundary. Do not add code for the other four boundaries. Make that scope explicit
  in one sentence so readers know the example is the freshness case and the other
  boundaries follow the same pattern (check at the boundary, record the decision).
- Where reviewers ask for capability "shown in code", answer in prose with a concrete
  example of what the check and the decision record would contain, not new code.
- Move the definition of agent/harness above the bridge sentence that first uses
  "agent" (editor finding 4).

Rewrite the draft so it works for all three readers at once: keep the full-stack
reader's path to code, the AI engineer's demand for evidence and failure modes, and
the data engineer's concern for provenance, freshness and governance. Where reviewers
disagree and the direction above is silent, prefer the change that makes the post
more concrete.

The post must keep: at least one diagram embedded as a Markdown image link to its
absolute raw GitHub URL (never a ```mermaid block) with a one-line caption; if you
change what the diagram shows, edit book/substack/ch01.03-context-failure-case-studies/diagram-1.mmd too; a short code
excerpt (15 lines or fewer, fenced and language-tagged) taken from
book/examples/ch01.03-context-failure-case-studies/ with a link to
https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.03-context-failure-case-studies for the full
program (never paste the whole program), and a "## Guidelines"
section of 3–7 imperative rules. Keep any numbers and citations exactly as the
module supports them; do not add new ones. Use absolute URLs only.

Start the file with:
---
title: Every Context Failure Has an Address
subtitle: <one sentence>
module: ch01.03-context-failure-case-studies
mood: broken
scheduled: 2026-10-30
---
Then the post body in Markdown. When done, reply with a 3-line summary of what you
changed.
