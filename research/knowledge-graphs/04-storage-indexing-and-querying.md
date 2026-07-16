# Storage, Indexing, and Querying

## Core Insight

The storage engine is less important than the indexing strategy.

Repeated discussions while reading the book led to the conclusion that retrieval quality dominates storage choice.

## Observations

- Graph databases are useful but not universally required.
- PostgreSQL can represent graph structures effectively.
- Indexes frequently matter more than database type.
- Graph algorithms remain valuable regardless of storage.

## Information Retrieval

Knowledge graph research repeatedly intersects with information retrieval.

Important technologies include:

- Lucene
- Elasticsearch
- BM25
- SQL indexes
- Unix filesystem lookup

These systems solved retrieval problems long before LLMs.

## Graph Algorithms

Potential applications:

- Dijkstra for retrieval planning
- Tarjan for dependency analysis
- graph traversal for authorization
- shortest-path reasoning

## When Is a Graph the Right Tool?

Before reaching for a knowledge graph, first ask: is this a graph problem?

Not all data benefits from graph representation. The book implicitly addresses this, but it deserves explicit treatment.

### Graph-Shaped Problems

Graphs excel when you need to **traverse many nodes at variable depth** — the number of hops is unknown at query time. Signs you have a graph problem:

- **Deep traversal** — queries must walk chains of unknown length (e.g., "find all dependencies of this component recursively")
- **Relationship attributes** — edges carry meaningful metadata (confidence, provenance, temporal bounds)
- **Network analysis** — centrality, community detection, dependency chains
- **Authorization as traversal** — access decisions depend on walking a relationship graph

### Simpler Relationships Don't Need Graphs

For bounded relationships known at write time, you don't need a graph database:

- **Tool call semantics** — function descriptions and parameters can encode relationships. The LLM learns which tool to call in sequence.
- **Business logic** — explicit code can handle joins. Two JOIN statements don't require Neo4j.
- **Prompt engineering** — relationship semantics can live in context, not storage.

A well-normalized relational database with proper foreign keys and an ERD is sufficient for most known relationships. The graph becomes valuable when **discovery** is the problem — when relationships must be inferred, not asserted.

### Graphs Are Not Free

Practical concerns that the book doesn't emphasize:

- **LLMs already understand semantics.** You don't need a graph to teach an LLM about relationships. The model learned relationships from training data. Graphs are for **engineering** those relationships, not for making them intelligible to the model.
- **Maintenance burden.** Graphs require ongoing curation. Schema changes, entity resolution, and relationship validation are continuous work.
- **Ingestion latency.** Neo4j and similar graph databases can be slow to ingest. For high-volume streaming data, the overhead is significant.
- **Live data mismatch.** If your data changes frequently in real-time, maintaining an up-to-date graph adds latency and complexity. The graph becomes stale, or the sync pipeline becomes a second system to maintain.

### The Decision Framework

For context engineering specifically:

1. **Start with the query pattern.** If queries traverse variable-length paths through relationships, a graph helps. If queries retrieve by ID or similarity, they don't.
2. **Consider write frequency.** Graphs require maintaining relationship integrity on writes. High-volume inserts without complex traversal may not justify the overhead.
3. **Ask about uncertainty.** Graphs shine when relationships are discovered (entity resolution, relation extraction). They add less value when relationships are known at write time.
4. **Think about governance.** Graphs make relationships visible and auditable—a feature for compliance, but overhead if relationships are simple and static.

## Open Questions

- When should graph traversal replace SQL joins?
- When are vectors preferable?
- What is the optimal hybrid retrieval architecture?
- Can graph problems be identified automatically from query logs?