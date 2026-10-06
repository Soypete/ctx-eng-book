# Data engineer review — Stop Saying "The AI Failed"

**Would I finish it?** yes. It's short, and "a catalog with an owner, a version, and an as-of date" is how I already talk, so I kept reading.
**Would I share it?** yes. I'd send it to the ML/platform team that keeps asking me for "a feed for the LLM," so they can see the data contract is the real work.

## What lands
- "A catalog with an owner, a version, and an as-of date." That's a data contract, and the post treats it as the fix. Good.
- The freshness gate runs *before* the model call (`now.Sub(s.Catalog.AsOf) > s.Catalog.MaxAge`). Fail closed on stale reference data is exactly right.
- Stamping `CatalogVersion` on every `Result` gives me lineage: I can tell which snapshot produced a label.
- The Postgres analogy works for my audience.

## Where I got lost or rolled my eyes
- "Does someone accountable vouch for them, and are they current and authorized for this request?" Vouch how? That's ownership, an SLA, and row-level access policy. Name them. "Vouch" means nothing in a catalog tool.
- "The model labels a ticket with a product that was retired last quarter." If the catalog is stale, the membership check `s.Catalog.Products[l.Product]` *passes* for the retired product. The freshness gate is doing all the work here, and the post doesn't say that, or what happens between refreshes when a product retires inside `MaxAge`.
- "Refuse to run on a catalog past its max age." So every ticket goes to human triage when my Airflow DAG is late? That's an outage triggered by a pipeline SLA miss. Say who gets paged, and whether stale-but-recent is better than nothing.
- "Schemas and taxonomies hold how your domain defines its terms." Fine, but taxonomies evolve. What happens to historical labels when "sev2" gets redefined? No word on schema evolution.

## Missing for me
- Where `Catalog.AsOf` comes from: the source system's last-modified time, or load time? Those differ, and freshness checks built on load time lie.
- A small sketch of the upstream: source of truth -> dbt model/snapshot -> versioned artifact the service loads. Even three boxes in the diagram.
- How the severity taxonomy is versioned alongside the catalog. Is it a contract with tests, or a string in a prompt?
- Cost/volume: is the catalog re-sent into context on every request? At what size does that stop working?

## Top 3 edits (ranked)
1. Define where `AsOf` and `Version` come from and who owns the pipeline that sets them. One sentence on "event time vs. load time" earns data engineers' trust.
2. Explain the stale-catalog failure mode honestly: the membership check doesn't catch retired products, the freshness gate does, and failing closed has an operational cost (triage flood). Say how to handle that.
3. Replace "vouch for" with concrete governance terms: owner, SLA, access policy.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
