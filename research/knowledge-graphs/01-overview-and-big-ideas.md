# Overview and Big Ideas

> Working Notes from *Knowledge Graphs* (MIT Press)
>
> These are not chapter summaries. They are the ideas, observations, and hypotheses that emerged while reading the book and discussing its implications for Context Engineering.

---

# Why I Read This Book

I did not read *Knowledge Graphs* because I intend to build every AI system around a graph database.

I read it because I wanted to understand one of the oldest attempts at making information understandable by machines.

The central question isn't whether knowledge graphs are the future of AI.

The question is:

> **What can decades of semantic web research teach us about building reliable AI systems?**

That framing completely changes how the book should be interpreted.

Rather than looking for a technology recommendation, I am looking for engineering principles that improve retrieval, reasoning, reliability, and governance.

---

# The Biggest Realization

Knowledge graphs are not an AI invention.

They predate modern machine learning by decades.

Most of the concepts discussed throughout the book—semantic relationships, ontologies, linked data, graph traversal, inference, and contextual retrieval—were developed long before transformers existed.

This was an important realization because much of today's AI discourse treats graphs as a new discovery.

They are not.

LLMs simply made many of these older ideas relevant again.

---

# Context Engineering Is Bigger Than Knowledge Graphs

One of the strongest conclusions from reading the book is that context engineering should not become synonymous with knowledge graphs.

Knowledge graphs are one implementation of semantic context.

They are not context itself.

Reliable AI systems may use:

- relational databases
- graph databases
- document stores
- object storage
- vector databases
- search indexes
- APIs

or combinations of all of them.

Context engineering is the discipline of organizing and retrieving information regardless of where it is stored.

Knowledge graphs provide one possible semantic layer.

---

# Engineers Think About Tools

One recurring observation throughout the book is that most historical work was performed by researchers, librarians, and ontologists.

Their focus is on:

- knowledge representation
- ontology design
- semantics
- logical inference

Modern software engineers often approach the same problem differently.

Instead of asking:

> "How should this domain be represented?"

they ask:

> "Should I use Neo4j?"

The distinction matters.

A graph database is an implementation.

An ontology is a model of meaning.

The ontology survives migrations between storage systems.

The database does not.

---

# Semantics Matter More Than Storage

The more I read, the less convinced I became that graph databases are the important innovation.

Instead, the real innovation is semantic modeling.

Storage technologies change.

Indexes improve.

Query languages evolve.

Meaning remains.

This reinforces one of the central ideas developing throughout Context Engineering:

> **Reliable AI depends more on preserving semantics than preserving storage formats.**

---

# The Semantic Web Didn't Fail

Popular engineering culture often describes the Semantic Web as a failed project.

The book paints a different picture.

Many of its ideas survived.

Examples include:

- Wikidata
- linked identifiers
- RDF
- ontology languages
- structured metadata
- graph representations

Rather than disappearing, semantic technologies became specialized infrastructure instead of consumer-facing products.

That is an important historical distinction.

---

# AI Did Not Invent Retrieval

Reading this book repeatedly reinforced another observation.

Retrieval has been a solved engineering problem for decades.

Search engines.

Lucene.

Elasticsearch.

Database indexes.

Unix file systems.

SQL query planners.

These technologies already solved many of the problems that modern AI systems are rediscovering.

This raises an important question.

Why are so many AI systems attempting to solve retrieval inside an LLM instead of leveraging the systems already designed for retrieval?

---

# Data Engineering Is Still the Foundation

The more I read about graphs, the stronger my belief became that reliable AI systems are fundamentally data engineering systems.

Every capability eventually depends on data that is:

- organized
- indexed
- governed
- versioned
- queryable
- semantically meaningful

The language model is only one component.

Without reliable data organization, the model has nothing reliable to reason over.

---

# Attention Does Not Replace Semantics

Several chapters discussed earlier NLP systems that required engineered positional embeddings and explicit relationship modeling.

Transformers changed that landscape by introducing attention.

However, attention does not eliminate semantics.

It changes where semantics are represented.

The model still requires meaningful information.

Attention simply provides a mechanism for selecting that information during inference.

The engineering problem therefore shifts from designing features to designing context.

---

# Graph Algorithms Are Still Relevant

One unexpected observation came from references to classic graph algorithms such as Dijkstra and Tarjan.

These are not ontology algorithms.

They are graph algorithms.

Their appearance suggests something important.

Once information is represented as a graph, decades of graph theory immediately become applicable.

That has implications far beyond knowledge graphs.

Graph algorithms could support:

- retrieval planning
- dependency analysis
- semantic navigation
- authorization traversal
- context selection

This deserves further investigation.

---

# Knowledge Graphs Are About Relationships

Traditional databases organize records.

Knowledge graphs organize relationships.

That distinction appears throughout the book.

Many AI tasks are relationship problems rather than storage problems.

Examples include:

- entity linking
- recommendation
- provenance
- dependency analysis
- authorization
- contextual retrieval

The graph is valuable because it makes relationships explicit.

---

# Reliable AI Requires Better Context, Not Bigger Models

Perhaps the strongest conclusion from these discussions is that scaling model size is not the only path toward better AI.

Smaller models supplied with:

- better retrieval
- richer semantics
- cleaner context
- governed access
- structured state

may outperform larger models operating over poorly organized information.

This idea has become one of the central hypotheses of the Context Engineering book.

---

# Emerging Thesis

As I progressed through the book, my working thesis evolved into something much broader than knowledge graphs.

Reliable AI systems require engineered context.

Engineered context requires structured information.

Structured information requires semantics.

Semantics require intentional modeling.

Knowledge graphs represent one mature approach to semantic modeling, but they are only one part of a much larger engineering discipline that also includes retrieval systems, authorization, data engineering, state management, evaluation, and observability.

That broader discipline is what I call **Context Engineering**.

---

# Questions Worth Investigating

The book also generated a number of research questions that remain unanswered.

## Reliability

- How should semantic systems be evaluated?
- What metrics indicate a reliable knowledge graph?
- Can graph quality predict AI performance?

## Retrieval

- When should retrieval use graphs?
- When should it use SQL?
- When should it use lexical search?
- When should it use vectors?
- What does an optimal hybrid retrieval pipeline look like?

## Context Engineering

- Which ideas from semantic web research transfer directly to production AI?
- Which ideas were abandoned because the technology of the time was insufficient?
- Which ideas become practical again because of modern language models?

## Production Systems

- How should semantic modeling integrate with existing data platforms?
- Should every enterprise knowledge platform expose graph semantics even if it stores data relationally?
- How should authorization interact with semantic traversal?

---

# Book Implications

These notes reinforce several directions for *Context Engineering*.

- Treat knowledge graphs as one implementation of semantic context, not the definition of context engineering.
- Focus on semantics rather than graph databases.
- Explain why retrieval technologies remain foundational to AI.
- Connect semantic modeling directly to reliability.
- Show how decades of information retrieval and semantic web research solve problems now being rediscovered by LLM systems.
- Position context engineering as the systems discipline that integrates these ideas into production AI architectures.