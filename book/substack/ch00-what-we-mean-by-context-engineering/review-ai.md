# AI engineer review — Context Engineering Is Not Prompt Engineering

**Would I finish it?** yes. The posture table and "the system, not the model, owns the boundary" are worth the read, though the code oversells itself.
**Would I share it?** yes, with a PM or a new hire on the RAG team. I wouldn't send it to my evals lead yet.

## What lands
- The "who owns selection" framing, with the "Model proposes, host authorizes" row. Most writing on tool-calling skips that column completely.
- "Ranking can be as probabilistic as you like... Scope, cutoff, budget, and admission should not be." That's a line I'd quote in a design review.
- Routing tool calls back through Authorize, as the diagram shows, is the right call. Too many MCP setups bypass it.

## Where I got lost or rolled my eyes
- "The salary sheet had the best relevance score and still never reached the ranker." It did reach the ranker. `Score` is already filled in when `Compile` runs, so the retriever indexed and scored the comp sheet. This is post-filtering, and it has a failure mode you skip: if your top-k=20 comes back with 18 unauthorized hits, the user gets thin context and the answer degrades without any error. The real fix is a pre-filter in the index (ACL metadata filters on the vector query).
- "the manifest makes the selection replayable." `sort.Slice` isn't stable, so documents with tied scores can come back in a different order from run to run. Also, a list of IDs isn't enough to replay anything. You'd need the ranker/embedding version, the scores, the budget, and an index snapshot.
- "The support bot failed at the Lexicon layer." Your own definition puts "what is this actor allowed to say or do" under Pragmatics. A failed permission check sounds like Pragmatics to me. If the taxonomy can't sort the opening example cleanly, readers won't trust it.
- "fixed in four lines." The authorize loop is seven lines. It's a small thing, but my kind of reader counts.

## Missing for me
- Evals. The definition says "within measurable limits," but nothing gets measured. How do I test recall under authorization, or catch a manifest regression?
- Prompt injection through scoped tool results. Mid-loop retrieval also lets untrusted text in, not just the risk of missing documents.
- Is this new? Lexicon/Semantics/Pragmatics looks like Morris's syntax/semantics/pragmatics with the first term relabeled, and then "Structure" comes back as syntax. Say where the framing comes from.

## Top 3 edits (ranked)
1. Fix the code claim. Either show authorization as an index pre-filter, or say plainly that this is post-filtering and name the top-k starvation tradeoff. Drop "never reached the ranker."
2. Make the manifest actually replayable: `sort.SliceStable` with an ID tiebreak, and record the ranker version and scores.
3. Settle whether the opening failure is Lexicon or Pragmatics, and add one sentence on how you'd eval this pipeline.

## Miriah's notes

- Keep the opening failure as a **Lexicon** problem: it is a data problem. The spreadsheet's ownership/sensitivity and the user's access scope are Lexicon facts (ch00 defines Lexicon as sources, owners, sensitivity, and the boundaries of what is available). Make that explicit in the draft so it doesn't read as Pragmatics, and use Lexicon consistently everywhere authorization-of-data comes up.

- **Code fix:** keep `Compile` as a filter that runs after search (candidates arrive already scored). Drop "never reached the ranker" and any claim the doc never competed. Name the tradeoff plainly: post-filtering after top-k can starve the answer of eligible documents, so production systems push the scope into the retrieval query (metadata filter / row-level security). Keep "Authorize before you rank" as the guideline, framed as where you want to end up.
- **Manifest:** add source, policy, and ranker version fields to the `Manifest` struct so the "replayable" claim holds. Keep the code block at 40 lines or fewer.

<!-- Add your notes for the rewrite here. Astra reads this section. -->
