# Full-stack engineer review — Context Engineering Is Not Prompt Engineering

**Would I finish it?** yes. The comp-spreadsheet opener is a bug I could actually ship, and the Go snippet showed up before I got bored.
**Would I share it?** yes, with the team that owns our RAG search box, after the fixes below.

## What lands
- "Nobody decided whether that spreadsheet was allowed to compete for a spot in the answer." That's a good framing. It's an authz bug, not a prompt bug.
- The posture table works. "Model proposes, host authorizes" is the sentence I'd paste into a design doc.
- "Notice where the model's tool call goes: back through Authorize, not around it." I'd stop scrolling for that diagram.
- Guideline 3 (tool calls go through the same gate) is a change I could make on Monday.

## Where I got lost or rolled my eyes
- "The salary sheet had the best relevance score and still never reached the ranker." In my stack the vector DB *is* the ranker. `cands` comes back from a top-k similarity query that already ran over everything. So `Compile` is filtering after the ranking, not before it. If the k=10 results are mostly finance docs, I'm left with one eligible doc and a thin answer. The real fix is a metadata filter or Postgres RLS inside the retrieval query. The post never says this, so it teaches the post-filter pattern it says it's against.
- "That's the bug from the top of this post, fixed in four lines." It isn't four lines, and it isn't the fix (see above). Readers will check.
- "The support bot failed at the Lexicon layer." Two paragraphs earlier, "what is this actor allowed to say or do" was **Pragmatics**. So which layer is authorization in? The linguistics terms already cost me effort, and here they contradict each other.
- The definition says "for an authorized purpose," but the code is `allowed(user, owner)`. Purpose is never passed in. Where does purpose come from in a real request: a route, a scope, a header?
- Mermaid block plus a PNG. Substack won't render the mermaid, so readers get a code dump and then the same diagram again.

## Missing for me
- Where `allowed` comes from. Can I reuse my existing auth middleware or session claims, or do I need a new policy service?
- Where the manifest goes (request log, trace span, a table?) and what it costs in latency and storage per request.
- A test: one table-driven test showing the unauthorized doc never shows up would sell this better than any paragraph.

## Top 3 edits (ranked)
1. Fix the "authorize before ranking" claim. Show the filter pushed into the retrieval query (vector metadata filter or `WHERE owner IN (...)`/RLS), and say why post-filtering top-k starves results. Drop "four lines."
2. Settle which layer authorization belongs to (Lexicon or Pragmatics) and say it once.
3. Remove the raw mermaid block and keep the PNG. Then add two sentences on where the manifest lives in a normal request/trace setup.

## Miriah's notes

- Keep the opening failure as a **Lexicon** problem: it is a data problem. The spreadsheet's ownership/sensitivity and the user's access scope are Lexicon facts (ch00 defines Lexicon as sources, owners, sensitivity, and the boundaries of what is available). Make that explicit in the draft so it doesn't read as Pragmatics, and use Lexicon consistently everywhere authorization-of-data comes up.

- **Code fix:** keep `Compile` as a filter that runs after search (candidates arrive already scored). Drop "never reached the ranker" and any claim the doc never competed. Name the tradeoff plainly: post-filtering after top-k can starve the answer of eligible documents, so production systems push the scope into the retrieval query (metadata filter / row-level security). Keep "Authorize before you rank" as the guideline, framed as where you want to end up.

- **Rewrite 2 — voice:** cut the hedging and caveats back to my voice. Short, direct, first person, a clear stance, light humor (see blog-posts/how-everyone-is-using-ai-wrong.md). Keep the corrections (post-filter tradeoff, Lexicon framing, versioned manifest) but say each once, plainly. Trim the replay/retention and measurement sections to a few sentences each. Aim for ~1,100–1,300 words.
- **Rewrite 2 — code:** don't put all the code in the post. The full runnable code now lives at https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch00-what-we-mean-by-context-engineering (gofmt'd, with tests). Show only the key lines (≤15, the authorize-then-budget loop) and link to the repo for the rest.

- **Rewrite 3 — apply editor-review.md.** Fix every finding in editor-review.md, including #1 (my own edits on the lines starting "When an AI feature misbehaves", "This applies to", and "I organize the work"): repair the broken book URL (https://github.com/Soypete/ctx-eng-book/blob/main/book/chapters/ch00-what-we-mean-by-context-engineering.md), fix the typos, and keep my intent: the paragraph should end on the question "So what is context engineering?" leading into the definition. Call Lexicon/Semantics/Pragmatics a lens, not a theory.
- Add the agents → harness setup before "Who owns selection?": agents are models calling tools in a loop; the harness is the code around the model that builds its input and runs its tools; that's where context engineering happens. Use "harness" everywhere (no "host"), and one name for the authorization step (use "Authorize", matching the diagram).
- Introduce the runnable Go example and its link before talking about `Compile`, and say what compiling context means. Never say "companion".
- Finding 6: add one sentence after the Lexicon mapping saying this post works the Lexicon boundary in code; Semantics and Pragmatics get their own posts. Finding 9: make the Pragmatics example an action the support bot could take (e.g., issuing a discount), and say which posture the bot used (prompt-time compilation with no access scope in the query).
- Add one early sentence that this is post 1 of a series drawn from my book.

- **Final polish — fix the remaining editor-review.md cohesion findings, small edits only.** Rewrite from the current post.md and change as little as possible; don't restructure, don't add new sections, keep my voice and the 15-line excerpt.
  - Most important: say in one sentence that the diagram shows the production target (Authorize scopes the search itself) and the Go example is the recheck inside `Compile` after search.
  - Don't name "the Authorize step" before the pipeline is introduced; one name for scope ("access scope") and one for the selected output ("compiled context").
  - Define "actor" on first use; make "an oversized note" part of the example's setup.
  - Shrink the data-contract section: explain the heading in plain words or rename it, and cut index jargon a full-stack reader won't know.
  - Author findings, my call: bring "authority" from the subtitle back once in the body; make the agents bridge concrete; smooth the choppy lines after the series sentence. Embed stays: keep the scene-1.png image line exactly as is.

<!-- Add your notes for the rewrite here. Astra reads this section. -->
