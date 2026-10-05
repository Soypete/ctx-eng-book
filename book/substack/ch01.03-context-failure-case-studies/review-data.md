# Data engineer review — Every Context Failure Has an Address

**Would I finish it?** yes. The opening, a stale index next to a real-time CRM, is a bug I've been paged for, so it had me from the first paragraph.
**Would I share it?** yes. I'd send it to the platform team that keeps asking us to "just expose the warehouse to the agent."

## What lands
- "Authorization is not relevance." This is the clearest case against `SELECT *` into a prompt that I've seen. "Trusted code then selects the task-specific rows and fields" is projection pushdown, and it's the right instinct.
- "A stale answer looks exactly like a fresh one" and refusing to fall back to the cached value. Fail closed instead of serving stale data. Good.
- Per-question-type max age ("deal-stage… an hour; product-docs… a week") is a freshness SLA, and I can put that in a contract.
- Decision records at every boundary are lineage for context.

## Where I got lost or rolled my eyes
- "Every retrieved fact carries when the source last vouched for it" vs. `now.Sub(cached.ObservedAt)`. Is that the source's `updated_at`, the CDC commit time, or the time the indexer ran? Event time and processing time give different ages. If the index ran at 09:00 but read a replica that was 6 hours behind, `ObservedAt` makes the data look fresh when it isn't.
- "The CRM updated in real time; the index refreshed daily; nothing compared the two." The only fix offered sits on the read side. The upstream fix is CDC or event-driven invalidation, and so is the producer's freshness SLA. Refetching straight from the CRM on every miss also puts unbounded load on an OLTP system. That's a fourth budget you don't mention.
- "The context says what the data means: field definitions, source, freshness." That's a data catalog and a data contract. Say so. Otherwise readers will hand-write field descriptions into prompt templates, and those will drift.
- The Rashidi stat measures coding-agent command denylists. It's a stretch as support for a data-authorization point, and the hedge sentence after it admits as much.

## Missing for me
- **Schema evolution.** When a source renames `stage` to `deal_stage`, what does the context boundary do? Nothing here covers contract breaks.
- **Where decision records live.** Table schema, retention, and PII handling: "which fields were selected… where they came from" could log customer data. Is this OpenLineage-shaped?
- How a "capability" maps to the row- and column-level policies my warehouse already enforces. Is it a replacement or a pass-through?
- A diagram of source → CDC → index → freshness check → context, with each timestamp labeled.

## Top 3 edits (ranked)
1. Define which timestamp `ObservedAt` is (source commit time, not index run time), and add one line on fixing staleness upstream with CDC or invalidation, not only at read time.
2. Name the catalog and data contract as where "what the data means" comes from, and add one paragraph on schema changes.
3. Say where decision records are stored and how PII in them is governed. Otherwise the audit trail turns into a leak.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
