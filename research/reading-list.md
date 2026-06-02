# Context Engineering — Research & Reading List

## Goal of the Research

Build the foundational knowledge to support the hypothesis:

> Reliable AI systems require engineered context, structured state, semantic constraints, and governed retrieval.
>

This reading list is organized around:

- semantics
- distributed systems
- LLM behavior
- retrieval systems
- ontologies
- lakehouse infrastructure
- authorization
- memory/state management
- semantic indexing

The goal is not to become an academic specialist in every area.

The goal is:

> understand enough first principles to build reliable systems and explain them authoritatively.
>

---

# 1. Semantic Web + Ontologies

## Why This Matters

This is the foundation for:

- semantic modeling
- machine-readable relationships
- linked data
- ontologies
- constrained retrieval

Without this, "context engineering" becomes prompt engineering with extra steps.

---

## Read First

### The Semantic Web — Tim Berners-Lee

Foundational paper introducing machine-readable semantics.

https://www-sop.inria.fr/acacia/cours/essi2006/Scientific American_ Feature Article_ The Semantic Web_ May 2001.pdf

---

### W3C Semantic Web Overview

https://www.w3.org/2001/sw/

---

### RDF / Turtle

Focus on:

- namespaces
- triples
- identifiers
- linked data

https://www.w3.org/TR/turtle/

---

### OWL (Web Ontology Language)

Focus on:

- classes
- relationships
- constraints
- reasoning

https://www.w3.org/TR/owl2-overview/

---

### SPARQL

Focus on:

- graph traversal
- semantic querying
- relationship-aware retrieval

https://www.w3.org/TR/sparql11-overview/

---

# 2. Knowledge Representation + Knowledge Graphs

## Why This Matters

You are not building "graph hype."

You are learning:

- relationship modeling
- semantic traversal
- entity resolution
- contextual retrieval

---

### MIT Press — Knowledge Graphs

https://mitpress.mit.edu/9780262048330/knowledge-graphs/

---

### Linked Data Principles

https://www.w3.org/DesignIssues/LinkedData.html

---

# 3. LLM Foundations

## Why This Matters

You need to understand:

- attention
- tokenization
- context windows
- hallucinations
- tool use
- in-context learning

Not to train models.

To understand failure modes.

---

### Attention Is All You Need

https://arxiv.org/abs/1706.03762

---

### Model Context Protocol (MCP)

https://arxiv.org/abs/2412.17159

---

### Toolformer

https://arxiv.org/abs/2302.04761

---

### Language Models are Few-Shot Learners

https://arxiv.org/abs/2005.14165

Computational pragmatics

https://web.stanford.edu/~jurafsky/prag.pdf

---

### LLMs in Production

[Follow up with Chris for reference]

---

# 4. Distributed Systems + Reliability

## Why This Matters

This is your:

- determinism
- failure handling
- consistency
- observability
- state management

foundation.

---

### Designing Data-Intensive Applications

Probably the single most important systems book for this project.

Focus on:

- replication
- consistency
- streams
- event logs
- state
- distributed systems tradeoffs

https://dataintensive.net/

---

### CAP Twelve Years Later

https://www.infoq.com/articles/cap-twelve-years-later-how-the-rules-have-changed/

---

### Martin Fowler — Event Sourcing

Important for:

- memory
- replayability
- auditability
- conversational state

https://martinfowler.com/eaaDev/EventSourcing.html

---

# 5. Data Engineering + Lakehouse Architecture

## Why This Matters

This grounds your ideas in real production infrastructure.

You are arguing:

> Context is a query over distributed state.
>

---

### Apache Iceberg

Focus on:

- snapshot isolation
- schema evolution
- time travel
- partition pruning

https://iceberg.apache.org/docs/latest/

---

### Delta Lake

Focus on:

- ACID guarantees
- reliability
- versioned tables

https://docs.delta.io/latest/index.html

---

### Amazon S3

Research:

- consistency
- object semantics
- metadata access
- storage/compute separation

https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html

---

# 6. Auth, Security, and Permissioning

## Why This Matters

This is one of the biggest differentiators in your thesis.

Most AI systems:

- over-scope access
- ignore authorization
- trust the model too much

---

### OAuth 2.0

https://oauth.net/2/

---

### OpenID Connect

https://openid.net/developers/how-connect-works/

---

### NIST RBAC Model

Research:

- scoped access
- policy-driven retrieval
- least privilege

https://csrc.nist.gov/projects/role-based-access-control

---

### The UNIX Programming Environment

Focus on:

- composability
- pipelines
- small focused systems

https://archive.org/details/UnixProgrammingEnviornment

---

# 7. RDF Storage + Semantic Indexing

## Why This Matters

Directly supports:

- URI design
- semantic indexing
- compact identifiers
- SPO indexing
- retrieval performance

---

### Survey of RDF Stores and SPARQL Engines

https://arxiv.org/abs/2102.13027

---

### Storing and Indexing Massive RDF Data Sets

https://www.csd.uoc.gr/~hy561/papers/storageaccess/optimization/Storing and Indexing Massive RDF Data Sets.pdf

---

### PostgreSQL Index Types

Focus on:

- B-tree
- hash indexes
- operator classes
- text indexing

https://www.postgresql.org/docs/current/indexes-types.html

---

# 8. Retrieval + Search Systems

## Why This Matters

You need to understand:

- retrieval ranking
- vector search limitations
- lexical vs semantic retrieval
- hybrid search

---

### BM25 Explained

https://www.elastic.co/blog/practical-bm25-part-1-how-shards-affect-relevance-scoring-in-elasticsearch

---

### Vector Search Basics

https://www.pinecone.io/learn/vector-search-basics/

---

# 9. Memory + State Management

## Why This Matters

This supports:

> Memory as structured retrieval over persistent state.
>

Not:

> "the model remembers."
>

---

### Event Sourcing

(repeated intentionally — extremely important)

https://martinfowler.com/eaaDev/EventSourcing.html

---

### Google Recommender Systems Guide

Useful for:

- personalization
- retrieval ranking
- user state

https://developers.google.com/machine-learning/recommendation

---

# 10. Reliability + Evaluation

## Why This Matters

Without evaluation:

- you cannot prove improvement
- your hypothesis stays philosophical

---

### OpenTelemetry

https://opentelemetry.io/docs/

---

### LangSmith Evaluation Concepts

https://docs.smith.langchain.com/evaluation

---

# Suggested Reading Order

## Phase 1 — Foundations

1. Semantic Web
2. Designing Data-Intensive Applications
3. Attention Is All You Need
4. OAuth / RBAC basics

---

## Phase 2 — Structured Systems

1. Knowledge Graphs
2. RDF storage/indexing
3. Iceberg / Delta Lake
4. Retrieval systems

---

## Phase 3 — Reliability + Context

1. Event sourcing
2. Memory systems
3. Evaluation frameworks
4. Tool-calling architectures

---

# Questions to Ask While Reading

For every source ask:

1. What failure mode does this explain?
2. What engineering constraint does this imply?
3. How would this improve retrieval?
4. How would this improve reliability?
5. How would this reduce hallucination?
6. How would this reduce token waste?
7. How would this change a real production system?

That's how this becomes:

- a systems book
- a framework
- and eventually corporate training material.