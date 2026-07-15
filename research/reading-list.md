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

→ [Notes](semantic-web-paper-notes.md)

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

### Knowledge Graphs: Fundamentals, Techniques, and Applications

By Mayank Kerjriwal, Craig A. Knoblock, and Pedro Szekely (2021)

Comprehensive textbook on knowledge graph construction, focused crawling, and applications.

Key chapters:
- Learning graphs (focused crawling, context graphs)
- Construction techniques
- Ontology alignment

→ [Notes](wordnet-anchor-text-notes.md) — includes context graph quotes

---

### Diligenti et al. (2000) — Context Graphs

Original paper on context graph focused crawling.

> "The structure of paths leading to relevant pages can be an important factor in focused crawling."

Key ideas:
- Build classifiers for pages at distance 1-2 from relevant pages
- Use Hidden Markov Models for browsing
- Work backward from relevant pages via backlinks

**Need to locate this paper.**

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

→ [Notes](attention-is-all-you-need-notes.md)

---

### Model Context Protocol (MCP)

https://arxiv.org/abs/2412.17159

---

### Toolformer

https://arxiv.org/abs/2302.04761

→ [Notes](toolformer-notes.md)

---

### Forge — Tool-Calling Reliability

A Python framework for self-hosted LLM tool-calling. Provides guardrails, rescue parsing, retry loops, and response validation. Available as:

- **Proxy server** — drop-in layer for OpenAI/Anthropic APIs
- **WorkflowRunner** — direct Python integration
- **Guardrails middleware** — composable reliability stack

Ported to Go in [incode-agentware](https://github.com/incode-agentware/forge).

https://github.com/antoinezambelli/forge

→ [Notes](forge-notes.md)

---

### Language Models are Few-Shot Learners

https://arxiv.org/abs/2005.14165

→ [Notes](language-models-few-shot-learners-notes.md)

---

### Megatron-LM: Training Multi-Billion Parameter Language Models Using Model Parallelism

https://arxiv.org/abs/1909.08053

---

### Efficient Large-Scale Language Model Training on GPU Clusters

https://arxiv.org/abs/2104.04473

---

### DSPy: Compiling Declarative Language Model Calls

https://arxiv.org/abs/2312.08382

---

### Computational pragmatics

https://web.stanford.edu/~jurafsky/prag.pdf

→ [Notes](computational-pragmatics-notes.md)

---

### LLMs in Production

→ [Notes](llms-in-production/notes.md)
→ [Chapter 3](llms-in-production/chapter-3.md)
→ [Chapter 5](llms-in-production/chapter-5.md)
→ [Chapter 6](llms-in-production/chapter-6.md)
→ [Chapter 7](llms-in-production/chapter-7.md)
→ [Chapter 12 — Ethics (Appendix)](llms-in-production/chapter-12.md)

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

→ [Notes](unix-programming-environment-notes.md)

---

### The Unix Time-Sharing System (Berkeley)

https://dsf.berkeley.edu/cs262/unix.pdf

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

# 11. Knowledge Graph Quality Metrics

## Why This Matters

Context engineering needs measurable engineering metrics, just like databases and distributed systems. This section establishes a quantitative framework for evaluating knowledge graphs as semantic layers for AI systems.

---

### Zaveri et al. — Quality Assessment for Linked Data: A Survey

Comprehensive survey establishing core quality dimensions for knowledge graphs.

https://doi.org/10.3233/sw-150175

---

### Färber et al. — Linked Data Quality of DBpedia, Freebase, OpenCyc, Wikidata, and YAGO

Comparative empirical analysis of major knowledge graphs.

https://doi.org/10.3233/SW-170285

---

### Paulheim — Knowledge Graph Refinement: A Survey of Approaches and Evaluation Methods

Survey of KG error detection, completion, and evaluation methodologies.

https://doi.org/10.3233/SW-170267

---

### HATEOAS and Wikipedia Navigation

Hypermedia patterns in Wikipedia that LLMs can navigate.

https://en.wikipedia.org/wiki/HATEOAS

---

### Wikidata Data Model

Defines statement representation including temporal precision, qualifiers, references, and ranks.

https://www.mediawiki.org/wiki/Wikibase/DataModel

---

### WordNet

Princeton University's lexical database of English words. Provides synonym sets, hypernyms (broader concepts), hyponyms (narrower concepts), and structured semantic relationships.

Useful for: constrained semantic expansion in web search agents, guardrails against hallucinated connections, topic disambiguation.

https://wordnet.princeton.edu/

---

### Wikipedia API

Programmatic access to Wikipedia content for context retrieval.

https://en.wikipedia.org/api/rest_v1/

---

### Google Custom Search API

Public web indexing for context retrieval.

https://developers.google.com/custom-search/v1/overview

---

→ [Notes](kg-quality-metrics-notes.md)

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