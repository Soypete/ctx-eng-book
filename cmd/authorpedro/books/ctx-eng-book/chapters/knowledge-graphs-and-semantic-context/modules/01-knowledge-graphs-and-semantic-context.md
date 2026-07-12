# Knowledge Graphs and Semantic Context

## Beat 1: The Reliability Question

How do we represent meaning instead of text?

This is the reliability question for knowledge graphs. A vector database stores embeddings—numerical representations of token sequences. It knows that "purchase order" and "order" are semantically similar. It does not know that a purchase order belongs to a vendor, contains line items, and requires approval before fulfillment.

LLMs understand language. They comprehend that purchase orders relate to vendors, that line items aggregate to totals, and that approvals represent workflow states. But they infer these relationships probabilistically, and inference varies with context. Today's inference might connect vendor to purchase order. Tomorrow's might not.

Knowledge graphs answer a simple question: why should the model guess when we can tell it?

The graph does not replace the model's semantic understanding. It provides deterministic, explicit relationships that the model can rely on. This is the shift from hoping the model infers to encoding relationships explicitly.

## Beat 2: What the Graph Actually Represents

A knowledge graph encodes entities and relationships as data, not text.

The fundamental unit is the triple: subject-predicate-object. "Vendor A supplies Product X" becomes three nodes with two edges. The relationship is first-class data, not a string in a document field.

This matters for reliability because relationships in the graph have:

- **Identity**: Every node and edge has a unique identifier. No ambiguity about which vendor or which supply relationship.
- **Type**: Edges are typed. "supplies" is distinct from "manufactures" or "distributes."
- **Constraints**: Relationships follow rules. A vendor supplies a product. A vendor does not supply a vendor. The graph enforces this.
- **Provenance**: Every triple can track its source. The model knows whether a relationship came from the ERP system, a manual annotation, or an inferred correlation.

These properties do not exist in document stores or vector databases. They exist in graph representations because the graph models meaning as structure.

The reliability implication is direct: when the model traverses the graph, it follows explicit paths with known semantics. It does not interpret text. It traverses relationships.

## Beat 3: RDF, OWL, and the Semantic Web Legacy

The Resource Description Framework (RDF) is the foundational standard. Data is triples. Every entity is a URI. Every relationship is a URI. Global identity is the core principle—if two systems use the same URI, they refer to the same entity.

RDF Schema (RDFS) adds typing: classes, properties, domain and range constraints. An entity can be declared a "Vendor" class. A property can be declared to have "Vendor" as domain and "Product" as range. The schema provides meta-information about what the triples mean.

Web Ontology Language (OWL) extends this further. OWL adds expressivity: cardinality constraints, property characteristics (transitive, symmetric, functional), and class relationships (equivalent, disjoint, subclass). OWL is where formal reasoning becomes possible.

Consider the statement: "Every vendor has at least one product." In RDFS, you might express this as a domain/range constraint. In OWL, you express it as a cardinality restriction on the hasProduct property. The graph can then infer that a vendor without products violates the constraint.

This is reasoning in the model. OWL reasoners compute the logical closure—the set of all triples that follow from the axioms. A reasoner can infer that if Vendor A supplies Product X, and Product X is a Product, then Vendor A supplies some Product. The inference is explicit, auditable, and deterministic.

SPARQL is the query language for RDF graphs. It is to graphs what SQL is to relational databases. A SPARQL query traverses patterns in the triple store, filters by constraints, and returns results. "Find all vendors who supply products in category 'electronics'" becomes a pattern match over triples, not a full-text search over documents.

The practical reality is that most production systems use a subset of these standards. Full OWL reasoning is computationally expensive and often unnecessary. But even basic RDF—typed relationships with URIs—provides value beyond document storage.

## Beat 4: Property Graphs and the Storage Question

Property graphs offer an alternative to RDF. In property graphs, nodes and edges carry key-value properties. An edge representing "supplies" might have properties like quantity, effective_date, and contract_id.

The structural difference is subtle. RDF models everything as triples. Property graphs model nodes, edges, and properties. Practically, both represent the same information.

The choice between them is a systems decision, not a semantic one:

- **RDF** provides global identity (URIs), standardized reasoning (OWL), and SPARQL queries. It integrates with the broader semantic web ecosystem.
- **Property graphs** provide flexible schemas, native graph traversal, and operational tooling. They are popular in industry (Neo4j, Amazon Neptune) and often easier to operationalize.

Neither representation is inherently superior. Both can encode the same domain model. The decision depends on your tooling ecosystem and whether you need global identity or flexible local schemas.

An important observation: you may not need a graph database at all. Third normal form (3NF) in a relational database achieves basic graph semantics. Foreign keys encode relationships. Join tables encode many-to-many edges. Primary keys provide node identity. Referential integrity enforces relationship constraints.

A well-designed relational schema with a query layer (GraphQL, REST with include parameters, or a smart ORM) can present a graph interface over normalized data. PostgreSQL, MySQL, or Snowflake become graph databases with the right abstraction layer.

More specifically, RDF triples can be stored in property tables on PostgreSQL. Apache Jena and Apache AGE provide SPARQL query capability over PostgreSQL. You get RDF semantics with PostgreSQL's operational maturity. This is not theoretical—enterprise systems use this approach.

The implication: graph semantics are a modeling choice, not a product choice. The knowledge graph exists in your schema design and your retrieval patterns, not just in your database product.

## Beat 5: Reasoning Shifts to the Model

Here is the critical distinction that changes everything:

Classical knowledge graphs performed reasoning in the graph. OWL reasoners computed transitive closure. Rule engines applied logical inference. The graph itself was the computational engine.

Modern knowledge graphs with LLMs perform reasoning in the model. The graph provides structure and retrieval. The model provides interpretation and inference.

This is a fundamental architectural shift. The graph does not need to infer that a vendor supplies a product because the product belongs to a category. The model already understands that relationship. The graph simply needs to provide the specific vendor, product, and category so the model has explicit grounding.

This means reasoning in property graphs happens in the model. The graph provides explicit relationships. The model interprets them in context.

This is more reliable than graph-based reasoning because:

- The model adapts to context. The same relationship means different things in different prompts.
- The model handles ambiguity. "_supplies" could mean raw materials, logistics, or distribution. The model resolves this based on surrounding information.
- The model generalizes. The graph provides explicit relationships for known entities. The model infers relationships for novel situations.

The graph is the retrieval substrate. The model is the interpretation engine. This separation is cleaner than trying to encode all possible inference in the graph itself.

## Beat 5b: KG-Shaped Retrieval in Practice

A distinction worth noting: some systems use graph structure for **retrieval** without implementing **reasoning** in the graph itself. Tools like GBrain, Mem0, and Letta use knowledge graph structure to improve retrieval—typed edges, schema-defined relationships, graph traversal as a retrieval strategy—but perform reasoning in the LLM synthesis layer.

GBrain demonstrates this clearly. It uses **schema packs** to define explicit page types (person, company, meeting, deal, email) and typed relationships (attended, works_at, invested_in, founded, advises). Every write operation runs `extractEntityRefs`—pattern matching on markdown links and Obsidian wikilinks—to create typed edges with **zero LLM calls**. The graph grows on every write at near-zero cost.

The benchmark results are striking:

| Strategy | P@5 | R@5 |
|----------|-----|-----|
| ripgrep BM25 only | 18 | 75 |
| vector-only RAG | 18 | 80 |
| hybrid + RRF (no graph) | 18 | 85 |
| GBrain full stack | 49.1 | 97.9 |

The graph adds **+31 P@5 points**—this is the critical finding. Vector search alone underdelivers on relational queries. GBrain's graph traversal retrieval strategy provides explicit paths that vector similarity cannot capture.

GBrain has **typed edges, schema packs, and graph traversal for retrieval**, but it does **not** have classical reasoning. No OWL. No transitive closure. No formal ontology enforcement. The synthesis layer—an LLM—generates answers from retrieved context, including explicit gap analysis (what the brain doesn't know).

This is not a failure of these tools. It's a pragmatic evolution: graph structure for retrieval reliability, LLM for interpretation. The book thesis holds—relationships are given to the agent through the retrieval structure, no inference needed for traversal. But the reasoning about those relationships still happens in the model.

## Beat 6: Enterprise Architecture—Multiple Topical Graphs

Enterprise AI will not run on a single universal knowledge graph. The data is too diverse, the governance too varied, and the access patterns too distinct.

Instead, production systems will consist of multiple topical graphs:

- An HR graph covering employees, departments, roles, and organizational hierarchy
- A finance graph covering accounts, transactions, budgets, and approvals
- An infrastructure graph covering services, dependencies, and deployments
- A customer graph covering accounts, contacts, interactions, and opportunities
- A product graph covering catalogs, SKUs, categories, and relationships

Each graph has its own ontology, governance, authorization, and lifecycle. An agent working on infrastructure does not need access to HR data. An agent working on finance does not need to see product catalogs.

This aligns with bounded context and data mesh principles. Each domain owns its graph. Each graph exposes controlled interfaces. Agents discover the appropriate graph before retrieving information.

### The Domain-First Approach

Building knowledge graphs before understanding your domain is premature optimization. The recommended approach:

1. **Start with a database** — read-only, curated agent data with simple, well-understood structure.
2. **Put semantics into business logic of tools** — as you build tools, you learn the domain. Semantics emerge from tool relationships, not upfront modeling.
3. **Then build your knowledge store** — once you have a set of tools you understand, you have better domain knowledge. You can now model relationships meaningfully.

Without domain knowledge, RDF modeling is "basically impossible"—you cannot reason about relationships you do not understand. The knowledge graph is the end state, not the starting point. Start with tools, learn the domain, then evolve to graph-based reasoning.

### Structured Domain Classification vs Wikis

Wikis work well for personal knowledge management but struggle with enterprise domain expansion. LLMWiki-style stores (markdown-based, domain-less search) work at moderate scale but fall apart at scale due to recursive query performance.

Structured stores with explicit classification (GBrain's schema packs, MemPalace's Wing/Room/Drawer hierarchy) handle domain expansion better because:

- **Explicit types** constrain what entities can exist
- **Typed relationships** prevent semantic drift
- **Domain-specific schemas** enable intent-aware retrieval
- **Classification hierarchies** allow agents to navigate to the right subgraph

Wing-based domain classification outperforms wiki-based approaches for production systems because the taxonomy is explicit, not emergent. Wikis are fine for personal use. For enterprise AI, you need structural guarantees.

The practical implication: build topical graphs, not a universal graph. Start with the domain that needs explicit relationships most. Expand incrementally.

## Beat 7: Graph Algorithms as Runtime Tools

Ontology defines meaning. Graph algorithms optimize traversal.

This distinction matters. Ontology design (defining entities, relationships, constraints) is a modeling activity. Graph algorithms (finding paths, identifying clusters, computing centrality) are operational activities.

Two classical algorithms become particularly relevant for context engineering:

**Dijkstra's algorithm** finds the shortest path between nodes. In AI systems, "distance" can represent latency, authorization constraints, confidence scores, semantic distance, or retrieval cost. Rather than asking "what is the shortest path," the system asks "what is the cheapest or highest-quality path to assemble the context required for this inference?"

This reframes Dijkstra as a context assembly algorithm. The graph becomes a cost matrix. The retrieval engine computes the optimal subgraph to provide to the model.

**Tarjan's algorithm** identifies strongly connected components—clusters of nodes where each node can reach every other node. These clusters naturally correspond to bounded contexts, organizational units, dependency graphs, or semantic neighborhoods.

Such clusters become ideal retrieval units. Instead of traversing arbitrary paths, the system retrieves entire connected components that represent coherent contexts. This is more reliable than ad-hoc path traversal because the retrieval boundaries are structurally meaningful.

These algorithms are not part of ontology design. They are part of runtime context assembly. The ontology defines what relationships exist. The algorithms determine which subgraphs to retrieve.

## Beat 8: The Reliability Argument

We return to the reliability question: how do we represent meaning instead of text?

Vector databases store text as embeddings. They know that words are similar. They do not know what words mean.

LLMs understand meaning probabilistically. They infer relationships from training data and context. But inference is not guaranteed. The same model might infer a relationship in one prompt and miss it in another.

Knowledge graphs make meaning explicit. Relationships are first-class data. Constraints are enforced. Provenance is tracked. The model does not guess—the graph tells it.

This is the engineering argument for knowledge graphs in AI systems:

- **Determinism**: Relationships are explicit, not inferred.
- **Governance**: Access control applies to edges, not just nodes.
- **Explainability**: The retrieval path is visible and auditable.
- **Reliability**: The model uses known relationships, not guessed ones.

The model's semantic understanding is not replaced. It is grounded. The graph tells the model exactly which relationships are authoritative. The model interprets those relationships in context.

This is the shift from hoping the model infers to encoding relationships explicitly. It is the difference between probabilistic retrieval and deterministic assembly. It is the difference between "the model probably understands" and "the system guarantees it."

Context engineering builds that guarantee. The knowledge graph is the structure that makes it possible.