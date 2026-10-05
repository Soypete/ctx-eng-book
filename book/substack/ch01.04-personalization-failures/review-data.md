# Data engineer review — Your Model Doesn't Know Your Users. Your Retrieval Does.

**Would I finish it?** yes. "A framework might call your preference store 'memory.' It's still a table" got me, and the rest stays on ground I own.
**Would I share it?** yes. I'd send it to the platform team lead and our data governance person when the AI team asks for "a user profile API."

## What lands
- Calling memory just a table, plus the four questions (who, task, access, as-of-now). That's a data contract, and I can actually review against it.
- Per-type max age ("Code style can probably be a month old... current project... an hour"). That's a freshness SLA, and I already set those.
- Declared vs inferred vs imposed provenance with an override order. Most teams squash all three into one `user_profile` row.
- Pointing at Postgres RLS. Enforcement in the store, not the app, is the right instinct.

## Where I got lost or rolled my eyes
- "a cache in front of the preference store never got invalidated." That's the only freshness mechanism the post names. Inferred preferences usually come out of a batch job ("opened a lot of Python files last week"). The stale data is often the nightly dbt model, not a cache, and the age that matters is when it was computed, not when it was written. Which timestamp does "Timestamp each one" mean?
- `for _, r := range s.Rows`. A full scan over an in-memory slice reads like a toy. I know it's an example, but say that pushdown (`WHERE owner = $subject AND field = ANY($fields)`) is what a real store does, or a reader will copy the loop.
- "log the user, the capability issued, and each preference's source and timestamp" for every response. That's a new high-volume table holding PII. Who owns retention and access on it? The post creates a governance problem and never mentions it.

## Missing for me
- A schema. Something like `preference(owner, field, value, source ENUM(declared,inferred,org), observed_at, computed_at, pipeline_run_id)`. Five lines of DDL would tell me what this asks of my warehouse.
- Lineage for inferred prefs: which job produced the value, and can I trace "Python" back to its input events?
- Deletion and correction. When a user corrects an inference, does that write back upstream or just shadow the inferred value?
- What happens to the task→fields policy when a field is renamed (schema evolution).

## Top 3 edits (ranked)
1. Add a short preference-row schema with `source` and distinct observed/computed timestamps, and tie Failure 3 to batch-pipeline staleness, not only caches.
2. Add one sentence on the context log: it's PII, it needs retention and access policy, and someone has to own it.
3. Comment on the Go loop that production filtering happens in the query (RLS or predicate pushdown), not in app memory.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
