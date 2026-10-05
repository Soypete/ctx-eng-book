# Data engineer review — Your Agent Doesn't Need a Better Prompt. It Needs a Stop Condition.

**Would I finish it?** yes. The opener is a schema-drift bug I've personally paged on, so you had me at `user_id` vs `userid`.
**Would I share it?** yes, with the platform team that's being asked to "expose the warehouse to the agent." But the code example needs fixing first.

## What lands
- "Missing schema context started the loop. Missing state and enforcement are why it never ends." This is the right split. Catalog problem first, harness problem second.
- The three-failure table. "Over-scoped context: access check passes, but far more data enters the working set than the task needs" is exactly the gap between table grants and row/column policies.
- Splitting episodic, semantic, and working state reads like events / dimension table / query result. I can map that.

## Where I got lost or rolled my eyes
- "guesses a new column name, retries, gets a different error, guesses again" and then later "it returns `stuck-same-error` on the second identical error." Those are two different scenarios. Your opener says the errors *vary*, and `err.Error() == last` only catches consecutive byte-identical strings. Real database errors embed the guessed column name, a query ID, sometimes a timestamp. Your detector wouldn't fire on your own example, and the run ends `budget-exhausted`, not `stuck`.
- "each with a subject and a rule for replacing stale values." That's the whole freshness problem in nine words. Where's the timestamp, the source, the TTL? Semantic state without provenance is a cache nobody can invalidate.
- "record cost per request and attribute it to users and queries." Attribute it to what grain, stored where? That's a fact table someone has to own.

## Missing for me
- Who owns the fix for the root cause. The opener is a missing data contract. One line pointing to "put the catalog/schema in the working set" would close the loop the post opens.
- A sketch of a semantic-state record: `subject, value, source, observed_at, supersedes`.
- What "enforce at retrieval" means on a real platform: row-level security or a policy-filtered view, not a filter in the prompt.

## Top 3 edits (ranked)
1. Make the code match the story. Either normalize errors to a class (e.g. `undefined_column`) before comparing, or rewrite the opener so the error repeats. Right now the example doesn't do what the next paragraph says it does.
2. Add provenance and freshness fields to the semantic-state definition, even if it's just a one-line record shape.
3. Say once that the schema mismatch is a data-contract/catalog failure the harness can stop but can't fix.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
