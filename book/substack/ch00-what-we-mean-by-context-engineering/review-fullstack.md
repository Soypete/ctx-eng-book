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

<!-- Add your notes for the rewrite here. Astra reads this section. -->
