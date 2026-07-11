# Knowledge Graphs Reading Notes
## Semantics, Knowledge Engineering, and Context Engineering

> Working hypothesis:
>
> Large language models have changed the role of knowledge graphs. Their primary value is no longer teaching machines semantics—it is providing deterministic, governed structure for systems that already understand semantics.

---

# 1. Knowledge Graphs Existed Because Machines Could Not Understand Language

One of the most important observations from *Knowledge Graphs* is that the proliferation of knowledge graphs was driven by the difficulty of getting machines to understand natural language and other human modalities.

Historically, this made perfect sense.

Classical software systems cannot infer meaning from text. If we wanted machines to reason about "employees," "departments," or "products," we had to explicitly model those concepts and their relationships.

Knowledge graphs became that representation.

LLMs fundamentally change this assumption.

Modern language models already possess a semantic understanding of natural language. They do not require RDF to understand that an employee works for a department or that a purchase order belongs to a project.

Instead, what they lack is deterministic access to authoritative information.

This changes the role of the knowledge graph.

Rather than replacing natural language with symbolic representations, knowledge graphs provide structured, governed, and explainable access to enterprise data.

---

# 2. The Semantic Web Was Right About Semantics, Wrong About the Consumer

The Semantic Web envisioned computers directly reasoning over RDF and ontologies.

That largely failed in production.

Instead, the software industry standardized around:

- REST APIs
- JSON
- SQL
- application-specific interfaces

These approaches were easier to engineer, easier to scale, and easier to maintain than generalized semantic reasoning.

Interestingly, AI changes the original assumption.

The consumer of semantic information is no longer another deterministic software system.

The consumer is an LLM.

This suggests that the Semantic Web may have arrived decades before machines were capable of benefiting from it.

---

# 3. Wikipedia Is Not a Knowledge Graph

An important distinction emerged while thinking about Wikipedia.

Wikipedia is often described as a knowledge graph because it contains linked information.

However, the links themselves are not semantic relationships.

They are navigational hyperlinks.

Humans infer meaning from the surrounding text.

Machines historically could not.

Projects like Wikidata and DBpedia exist precisely because Wikipedia itself is not a semantic graph. Both use RDF triple stores.

The graph must be extracted from the documents.

This distinction matters because it separates:

- documents
- navigation
- semantic relationships

Those are different engineering problems.

---

# 4. Wikis Are Human Navigation Systems

A useful mental model is:

> A wiki is a navigable filesystem designed for humans.

Humans browse pages.

Humans infer relationships.

Humans understand context.

The links primarily improve discoverability.

This is fundamentally different from a knowledge graph where relationships have explicit meaning.

For example:

Employee
WORKS_FOR
Department

contains significantly more information than simply linking one page to another.

---

# 5. Obsidian Is an Authoring Tool, Not a Runtime Data Store

Obsidian graph visualizations are useful.

However, they should not be confused with production knowledge graphs.

Markdown files optimize for:

- writing
- editing
- version control
- human readability

Production AI systems optimize for:

- retrieval
- indexing
- authorization
- provenance
- latency
- graph traversal

Those are different workloads.

This suggests a useful architecture.

```text
Markdown / Wiki
        ↓
Extraction
        ↓
Ontology
        ↓
Entity Resolution
        ↓
Knowledge Graph
        ↓
Indexes
        ↓
Governed Retrieval
        ↓
LLM
```

The wiki becomes the authoring experience.

The graph becomes the runtime representation.

---

# 6. AI Reunites Knowledge Engineering and Software Engineering

Historically there have been two largely separate disciplines.

Knowledge engineering focused on:

- ontologies
- semantics
- taxonomy
- information science
- library science

Software engineering focused on:

- storage
- indexing
- APIs
- distributed systems
- databases

AI forces these disciplines back together.

Reliable AI systems require both.

A graph database without an ontology is simply another database.

An ontology without efficient retrieval remains largely academic.

Context engineering becomes the systems discipline that connects semantic modeling with production infrastructure.

---

# 7. Star Schema Often Embeds Semantic Layers

Star and snowflake schemas often embed significant semantic information within their structure:

- Dimension tables encode hierarchies (time, geography, product category)
- Foreign key relationships encode business logic
- Slowly changing dimensions (SCD) encode temporal semantics

These schemas function as implicit semantic graphs—they encode meaning through table relationships rather than explicit RDF/OWL.

The semantic reasoning happens through SQL joins and aggregations rather than graph traversal, but the information content is equivalent.

This suggests semantic reasoning doesn't require graph databases. It requires semantic structure—which can exist in relational, document, or graph stores. The storage layer is less important than explicit relationships and governed retrieval.

---

# 7b. Third Normal Form Achieves Basic Graph Semantics

3NF modeling with a well-designed query layer can achieve basic graph semantics:

- Foreign keys encode relationships (the "edges")
- Join tables encode many-to-many relationships
- Primary keys provide node identity
- Referential integrity enforces relationship constraints

A good query layer (GraphQL, REST with include params, or even smart ORM) can traverse these relationships declaratively, effectively presenting a graph interface over normalized relational data.

This challenges the assumption that you need a graph database to have a knowledge graph. A properly normalized relational model with good relationship traversal is already a primitive property graph.

Any RDBMS can achieve this—PostgreSQL, MySQL, SQLite, Snowflake. The semantic layer comes from the modeling, not the product.

Properties in property graphs can be viewed as metadata or attribute columns in RDBMS—key-value pairs on edges and nodes are semantically equivalent to nullable columns in a relational table. The graph structure emerges from the foreign key relationships, not from specialized storage.

The practical implication: you don't need a graph database to have graph-like semantics. A normalized 3NF schema with proper foreign keys IS a property graph, and JSON/JSONB columns can capture flexible properties without schema migration.

---

# 8. Engineers Often Start With Tools Instead of Models

One observation is that many modern AI discussions begin with technology selection.

Questions such as:

- Should I use Neo4j?
- Should I use GraphRAG?
- Which vector database should I choose?

appear before anyone asks:

- What entities exist?
- What relationships matter?
- What constraints should govern retrieval?
- What ontology describes this domain?

The original Semantic Web research started with semantics.

Modern engineering discussions often start with implementation.

Reliable AI systems likely require reversing that order.

---

# 9. Knowledge Graphs Become More Valuable Because LLMs Understand Semantics

> "Log does not have semantic safeguards."

The graph no longer needs to teach the model language.

The model already understands language.

Instead, the graph provides:

- authoritative entities
- explicit relationships
- governance
- permissions
- provenance
- deterministic retrieval

The semantic understanding comes from the model.

The reliability comes from the data platform.

---

# 10. Enterprise AI Will Likely Consist of Many Topical Graphs

Rather than a single universal enterprise graph, production systems will likely consist of multiple domain-specific graphs.

Examples include:

- HR
- Finance
- Infrastructure
- Customers
- Products
- Security

Each graph can expose:

- its own ontology
- its own governance
- its own authorization
- its own lifecycle

Agents discover the appropriate ontology before retrieving information.

This aligns well with data mesh and bounded context principles.

---

# 11. Classical Graph Algorithms Become Runtime Tools

One surprising observation from the bibliography is the inclusion of classical graph algorithms such as Dijkstra and Tarjan.

These are not tools for ontology design.

They are tools for operating on graphs.

This distinction is important.

Ontology defines meaning.

Graph algorithms optimize traversal.

---

## Dijkstra

Dijkstra minimizes traversal cost.

In AI systems, "cost" could represent:

- latency
- authorization
- confidence
- semantic distance
- token cost
- retrieval expense

Rather than asking:

> What is the shortest path?

An AI system asks:

> What is the cheapest or highest-quality path to assemble the context required for this inference?

This reframes Dijkstra as a context assembly algorithm.

---

## Tarjan

Tarjan identifies strongly connected components.

These may naturally correspond to:

- bounded contexts
- communities
- organizational structures
- dependency clusters
- semantic neighborhoods

Such clusters may become ideal retrieval units for context hydration.

---

# 11. Context Engineering Is Graph Traversal

One emerging hypothesis is that context engineering can be viewed as graph traversal under constraints.

The ontology defines:

- what relationships exist

The storage layer defines:

- where data lives

Graph algorithms determine:

- what information should be retrieved

Authorization determines:

- what information may be retrieved

The LLM determines:

- how retrieved information is interpreted

This cleanly separates responsibilities between semantic modeling, infrastructure, governance, and inference.

---

# Key Takeaways

- Knowledge graphs originally compensated for machines' inability to understand language.
- LLMs already understand semantics; they require reliable, governed access to structured information.
- The Semantic Web may have been conceptually correct but technologically premature.
- Wikis are navigable document systems, not semantic knowledge graphs.
- Markdown is an excellent authoring format but a poor runtime datastore.
- Knowledge engineering and software engineering are converging because of AI.
- Engineers often optimize tooling before defining semantic models.
- Enterprise AI is likely to consist of multiple governed topical graphs rather than a single global graph.
- Classical graph algorithms such as Dijkstra and Tarjan become valuable for context retrieval and graph traversal, not ontology creation.
- Context engineering is increasingly about assembling the smallest, highest-quality, authorized subgraph required for a given inference.

---

# 12. Property Graphs vs Triple Stores

From Chapter 2.3 — Property Graphs and Triple Stores

Property graphs model edge/node relations as key-value pairs, typically foreign-key driven rather than URI-driven.

Triple stores (RDF) are the alternative—modeling data as subject-predicate-object triples with global URIs.

Tradeoffs:
- **Property graphs**: flexible schema, native graph traversal, popular in industry (Neo4j, Amazon Neptune)
- **Triple stores**: global identity, inferencing standards (OWL), linked data principles, academic legacy

Both can represent the same information. The choice is about tooling ecosystem and whether you need URI-driven global identity or flexible local schemas.

---

## RDF on PostgreSQL Scales

RDF triples stored in property tables (subject, predicate, object) on PostgreSQL can scale well for enterprise workloads. This combines:
- RDF semantics and standards (SPARQL, OWL inferencing potential)
- PostgreSQL's proven scalability, partitioning, and operational maturity

This argues for enterprise adoption—you get graph semantics without specialized graph database infrastructure.

Existing tools support this: Apache Jena and Apache AGE provide RDF/SPARQL on top of PostgreSQL, enabling graph querying over proven relational infrastructure.

Dedicated RDF triplestores (category: GraphDB—e.g., GraphDB, BlazeGraph, AllegroGraph) offer optimized reasoning and SPARQL, but add operational complexity compared to PostgreSQL-based solutions.

This creates an interesting observation:

> A star schema or snowflake schema in a relational database is essentially a flat property graph.

The difference is normalization. Relational schemas normalize data across tables. Property graphs denormalize into vertex and edge tables with flexible key-value properties.

Neither representation is inherently superior—they represent different tradeoffs in the same design space.

---

## Star Schema as Implicit Semantic Graph

Star and snowflake schemas often embed significant semantic information within their structure:

- Dimension tables encode hierarchies (time, geography, product category)
- Foreign key relationships encode business logic
- Slowly changing dimensions (SCD) encode temporal semantics

These schemas effectively function as semantic graphs—they encode meaning through structure rather than through explicit ontology languages.

The semantic layer exists, just normalized across table relationships rather than expressed as RDF or OWL.

This suggests that many "relational" data warehouses already contain semantic graphs—they just aren't labeled as such. The semantic reasoning happens through SQL joins and aggregations rather than graph traversal, but the information content is equivalent.

The implication: semantic reasoning doesn't require graph databases. It requires semantic structure—which can exist in relational, document, or graph stores. The storage layer is less important than the presence of explicit relationships and governed retrieval.

---

# 13. Reasoning in Property Graphs Happens in the Model

One critical distinction between property graphs and classical knowledge graphs:

Classical RDF/OWL reasoning happens in the graph itself—transitive closure, inferencing, rule engines.

Property graph reasoning happens in the LLM.

The graph provides structure and retrieval.

The model provides interpretation and inference.

This is a fundamental shift in where reasoning occurs—moving from symbolic computation to neural computation.

---

# 14. Encoding and Cipher Move to the Model

Related to reasoning, encoding and cipher also shift to the model.

In classical knowledge graphs, encoding schemes, controlled vocabularies, and cipher relationships were embedded in the graph schema itself.

In property graphs with LLMs, the model handles:
- entity resolution
- schema interpretation
- relationship inference
- encoding translation

The graph becomes a retrieval substrate. The model performs the interpretative work that previously required explicit graph reasoning.

---

# Potential Book Themes

- **The Semantic Web Was Twenty Years Too Early**
- **Knowledge Graphs Don't Replace Language—They Govern It**
- **A Wiki Is Not a Knowledge Graph**
- **Author Once, Retrieve Structurally**
- **Context Engineering Is Graph Traversal Under Constraints**
- **Reliable AI Requires Both Semantics and Systems Engineering**
- **Ontologies Define Meaning; Context Engineering Delivers Meaning**