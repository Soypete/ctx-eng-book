# Data engineer review — Context Engineering Is Not Prompt Engineering

**Would I finish it?** yes. The comp-spreadsheet opener is an incident I've actually had to write up, and "authorize before you rank" is the right instinct.
**Would I share it?** yes, with our platform and governance leads, but only after the replay and index-time points below are fixed. They'd catch both right away.

## What lands
- "Nobody decided whether that spreadsheet was allowed to compete." That's a governance failure, and it's described as one.
- The posture table. "Model proposes, host authorizes" is a contract I can enforce.
- Asking "when it was true and whether the authorization is still valid" puts freshness and entitlement expiry next to each other, which is where they belong.

## Where I got lost or rolled my eyes
- "The salary sheet had the best relevance score and still never reached the ranker." It did reach a ranker: the vector search that produced `cands`. If the index returns top-k and `Compile` filters afterward, unauthorized docs have already pushed eligible ones out of the k. Push it far enough and the user gets zero results. That's post-filtering. The fix the opener needs is upstream: why was comp data embedded into a shared index at all?
- "the manifest makes the selection replayable." It stores doc IDs, not versions. Once the pricing doc is re-ingested, or `allowed()` changes because someone left the finance group, replaying gives you a different answer. A manifest that can actually be replayed needs a source snapshot or version, the policy version, the ranker version, and a timestamp.
- `allowed(user, d.Owner)`. Owner isn't an access policy. Real entitlements are row- and column-level and live in the catalog. Treating "who owns it" as "who can read it" is how the comp sheet ended up indexed in the first place.
- "fixed in four lines." It's patched at query time. Fixing it means classification tags and ACLs that carry through into the index.

## Missing for me
- What this asks of my platform: sensitivity tags in the catalog, ACL metadata stored alongside embeddings, pre-filtered search, and lineage from source table to chunk to manifest.
- One line on cost and freshness: how stale an index is acceptable, and who pays to re-embed after a permission change.
- Add `Version`/`AsOf` to `Doc`, so the "Time" section shows up in the code too.

## Top 3 edits (ranked)
1. Say where authorization happens: filter inside or before the index query, not after top-k. Then reword "never reached the ranker."
2. Put source version, policy version, and timestamp in `Manifest`, or stop calling it replayable.
3. Add a short "what this asks of your data platform" paragraph (catalog tags, ACL propagation, lineage) so the Lexicon layer reads as data the platform owns.

## Miriah's notes

- **Code fix:** keep `Compile` as a filter that runs after search (candidates arrive already scored). Drop "never reached the ranker" and any claim the doc never competed. Name the tradeoff plainly: post-filtering after top-k can starve the answer of eligible documents, so production systems push the scope into the retrieval query (metadata filter / row-level security). Keep "Authorize before you rank" as the guideline, framed as where you want to end up.
- **Manifest:** add source, policy, and ranker version fields to the `Manifest` struct so the "replayable" claim holds. Keep the code block at 40 lines or fewer.

<!-- Add your notes for the rewrite here. Astra reads this section. -->
