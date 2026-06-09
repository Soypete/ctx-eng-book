# LLMs in Production — Chapter 3 Notes
## Parallelism, Quantization, Retrieval, and Context Engineering

**Date:** 2026-06-08
**Source:** *LLMs in Production* Chapter 3
**Related Research Areas:**
- Context Engineering
- Distributed Systems
- Self-Hosted AI
- Inference Infrastructure
- Retrieval Architectures
- Quantization
- Knowledge Graphs
- Vector Databases

---

# Key Insight

Chapter 3 is primarily about:

> How do we serve large models efficiently?

It is **not primarily about context engineering**, but it provides useful background on the infrastructure constraints that context engineering operates within.

The chapter explains how inference systems scale through:

- Data Parallelism
- Tensor Parallelism
- Pipeline Parallelism
- Quantization
- Precision Tradeoffs

These concepts affect:

- latency
- throughput
- memory consumption
- infrastructure cost

They do **not directly solve retrieval, grounding, memory, or hallucination problems.**

---

# Parallelism and Context Engineering

## Data Parallelism

Multiple copies of the same model run on different GPUs.

```text
Request A → GPU 1
Request B → GPU 2
Request C → GPU 3
```

### Purpose

- Scale throughput
- Handle more users
- Improve availability

### Context Engineering Relevance

Minimal.

The context packet is assembled before inference.

If retrieval is wrong:

```text
Bad Context
→ Replica 1
→ Replica 2
→ Replica 3
```

All replicas fail equally.

---

## Tensor Parallelism

Model weights are split across multiple GPUs.

```text
Layer Computation
GPU1 + GPU2 + GPU3
```

### Purpose

- Fit larger models into memory
- Enable serving models larger than a single GPU

### Context Engineering Relevance

Indirect.

Larger models often support larger context windows.

However:

> Larger context windows do not solve context engineering.

A 1M-token context window still suffers from:

- irrelevant information
- authorization problems
- retrieval mistakes
- evaluation challenges
- token waste

---

## Pipeline Parallelism

Different layers execute on different devices.

```text
Stage 1 → GPU 1
Stage 2 → GPU 2
Stage 3 → GPU 3
```

### Purpose

- Increase utilization
- Run larger models

### Context Engineering Relevance

Very little.

Changes inference architecture, not retrieval architecture.

---

# Important Context Engineering Observation

Infrastructure research often asks:

```text
How do we make models bigger?
```

Context engineering asks:

```text
How do we provide better information?
```

These are fundamentally different optimization strategies.

---

# Context Engineering as Query Optimization

One emerging insight:

> Context engineering is query planning for inference.

Bad systems:

```text
Request A → 2k tokens
Request B → 50k tokens
Request C → 200k tokens
```

Results:

- unpredictable latency
- KV cache pressure
- inconsistent throughput
- higher infrastructure costs

Better systems:

```text
Task A → 4k tokens
Task B → 6k tokens
Task C → 5k tokens
```

Results:

- predictable latency
- predictable memory usage
- improved throughput
- easier scaling

This mirrors database query optimization.

---

# Quantization and Precision

## What Quantization Does

Reduces numerical precision.

Examples:

```text
FP16
Q8
Q6
Q4
Q3
Q2
```

As precision decreases:

- weight accuracy decreases
- logits become noisier
- reasoning quality degrades
- tool calling becomes less reliable
- factual recall weakens

---

## Precision vs Hallucination

Quantization itself can increase hallucinations.

The cause is:

```text
Lower Precision
→ Less Accurate Internal Representations
→ More Inference Error
```

This is different from retrieval failures.

---

# Context Engineering as a Compensation Mechanism

Context engineering cannot eliminate quantization error.

However:

> It can reduce the amount of knowledge the model must remember.

Example:

Without retrieval:

```text
Question
→ Model Memory
→ Answer
```

With retrieval:

```text
Question
→ Retrieve Facts
→ Inject Facts
→ Answer
```

The model becomes:

- less dependent on stored parameters
- less dependent on perfect recall
- less dependent on reasoning over incomplete knowledge

---

# Emerging Hypothesis

A possible book thesis:

> Better context engineering reduces the precision requirements of the model.

More specifically:

```text
Better Context
→ Smaller Search Space
→ Less Ambiguity
→ Less Reasoning Burden
→ Smaller Models Become Viable
```

This shifts responsibility from:

```text
Model Parameters
```

to:

```text
System Architecture
```

---

# Retrieval-Augmented Generation (RAG)

## Revised Definition

Industry often assumes:

```text
RAG = Vector Database
```

This is too narrow.

A broader systems definition:

> Retrieval-Augmented Generation is any architecture that retrieves external information and injects it into inference.

---

# What Counts as RAG?

## SQL

```text
Question
→ SQL Query
→ Inject Results
→ Generate
```

RAG.

---

## Knowledge Graph

```text
Question
→ Graph Traversal
→ Inject Entities
→ Generate
```

RAG.

---

## API Calls

```text
Question
→ API Call
→ Inject Results
→ Generate
```

RAG.

---

## Web Search

```text
Question
→ Search Engine
→ Retrieve Pages
→ Generate
```

RAG.

---

## Vector Search

```text
Question
→ Embedding Search
→ Retrieve Chunks
→ Generate
```

Also RAG.

---

# Retrieval Categories

## Direct Retrieval

Examples:

- SQL
- Key-value stores
- Cache lookups
- APIs

Characteristics:

- deterministic
- fast
- structured

---

## Relational Retrieval

Examples:

- foreign keys
- graph traversal
- entity resolution

Characteristics:

- relationship-aware
- deterministic
- schema-driven

---

## Similarity Retrieval

Examples:

- embeddings
- vector databases
- ANN search

Characteristics:

- probabilistic
- semantic
- approximate

---

## Search Retrieval

Examples:

- BM25
- Elasticsearch
- Web Search

Characteristics:

- ranking-based
- text-oriented
- hybridizable

---

# Knowledge Graphs vs Vector Databases

## Vector Databases

Answer:

```text
What is similar?
```

Based on:

- embeddings
- distance metrics
- nearest neighbors

Useful when:

- structure is unknown
- language is fuzzy
- exploration is required

---

## Knowledge Graphs

Answer:

```text
What is related?
```

Based on:

- entities
- relationships
- graph traversal

Useful when:

- relationships matter
- structure exists
- correctness matters

---

# DDSO Implications

DDSOs introduce a third retrieval model:

```text
Task
→ Authorization
→ Scoped Ontology
→ Context Assembly
→ Inference
```

Instead of:

```text
Search Everything
```

DDSOs support:

```text
Retrieve Only What Matters
```

---

# Similarity Metrics

Embeddings require a distance function.

---

## Cosine Similarity

Measures:

```text
Angle Between Vectors
```

Useful for:

- semantic similarity
- direction-based meaning

Most common embedding metric.

---

## Dot Product

Measures:

```text
Direction + Magnitude
```

Important because transformer attention itself is based on dot products.

Many ANN systems optimize around dot-product search.

---

## Euclidean Distance

Measures:

```text
Straight-Line Distance
```

Traditional geometric metric.

Less common in modern high-dimensional embedding systems.

---

# Important Context Engineering Question

The industry often asks:

> Which similarity metric is best?

Context engineering asks:

> Why are we performing similarity search at all?

Many production retrieval tasks are actually:

- direct lookups
- graph traversals
- relational queries
- authorization checks

not similarity problems.

---

# Speed Matters

Every retrieval mechanism has different costs.

## Direct Lookup

```sql
SELECT *
FROM users
WHERE id = 123
```

Fast.

Deterministic.

Cheap.

---

## Graph Traversal

```text
Teacher
→ Course
→ Assignment
```

Usually fast.

Deterministic.

Relationship-aware.

---

## Vector Search

```text
Embed Query
→ ANN Search
→ Rerank
→ Inject
```

More expensive.

Approximate.

Additional infrastructure.

---

# Emerging Context Engineering Position

Possible thesis:

> Vector databases are not retrieval.
> They are one retrieval strategy.

More broadly:

> Retrieval should be selected based on data shape, reliability requirements, authorization constraints, and latency needs.

Not because:

```text
AI = Vector Database
```

---

# Relationship to the Book Thesis

This discussion reinforces the broader argument:

> Context is a query over distributed state.

The primary engineering challenge is not:

```text
How do we make the model larger?
```

It is:

```text
How do we retrieve the correct information,
for the correct user,
at the correct time,
with the correct permissions,
at the lowest possible cost?
```

That is the core systems problem of Context Engineering.

---

# Open Research Questions

1. How much can high-quality retrieval compensate for quantization?
2. At what point does better context outperform larger models?
3. When should retrieval use:
   - direct lookup
   - graph traversal
   - vector search
   - hybrid search
4. Can DDSOs reduce token consumption enough to materially reduce infrastructure costs?
5. Does context engineering reduce the precision requirements of deployed models?
6. How should retrieval systems be evaluated independent of model quality?
7. What percentage of production AI retrieval tasks are actually similarity problems?

---

# Connection to Reading List

Relevant resources:

- Attention Is All You Need
- Language Models Are Few-Shot Learners
- Computational Pragmatics
- Designing Data-Intensive Applications
- Semantic Web
- RDF Storage and Indexing
- Knowledge Graphs
- PostgreSQL Indexing
- BM25
- Vector Search Basics

Potential future chapter:

**Retrieval Is Not Vector Search: Choosing the Right Context Access Pattern**

This is probably the most important conceptual takeaway from the session:

Context engineering is not primarily a model optimization discipline. It is a data access discipline.

Models consume context. Systems decide what context exists.