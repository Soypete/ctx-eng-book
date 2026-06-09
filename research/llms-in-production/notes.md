# Notes: LLMs in Production, Attention, Pragmatics, and Few-Shot Learning

**Source:** *LLMs in Production* by Chris Fregly and Antje Barth (Manning)

**Date:** June 2026

**Purpose:** Capture observations relevant to the Context Engineering book and connect them to reliability, retrieval, semantics, and agent systems.

---

# Key Observation

Many of the concepts now being discussed under the banner of **Context Engineering** were already present in earlier discussions of:

- computational pragmatics
- retrieval systems
- data engineering
- in-context learning
- transformer architectures

The terminology is new.

Many of the underlying ideas are not.

---

# Pragmatics and Context Engineering

Chris describes the role of pragmatics as:

> introducing pragmatics into a conversation with an LLM via prompting and data engineering.

This appears to predate widespread use of the term **context engineering**.

## Interpretation

Semantics answers:

> What does this mean?

Pragmatics answers:

> What should this mean in this specific context?

Prompts, examples, retrieved data, and tool descriptions all serve as mechanisms for introducing pragmatic information into the model's decision process.

---

## Context Engineering Perspective

Prompting and retrieval are not context engineering themselves.

They are mechanisms used by context engineering.

Possible definition:

> Context engineering is the discipline of designing, governing, retrieving, and assembling the information, state, constraints, and interactions required for an AI system to behave reliably within a specific operational context.

---

# Semantics vs Pragmatics

## Semantics

Encoding meaning.

Examples:

- ontologies
- RDF
- OWL
- knowledge graphs

Questions:

- What is a student?
- What is a teacher?
- What relationships exist?

---

## Pragmatics

Encoding contextual interpretation.

Examples:

- prompts
- tool descriptions
- role instructions
- examples
- agent workflows

Questions:

- What does the user intend?
- Which tool should be used?
- What information is relevant right now?

---

# Competency Questions Beyond Ontologies

Traditional ontology engineering uses competency questions.

Example:

> Which courses is this student enrolled in?

The ontology succeeds if it can answer the question.

---

## Extension to Context Engineering

The same concept applies to AI systems.

Instead of asking:

> Can the ontology answer this question?

Ask:

> Does the system have access to the information required to perform this task?

Examples:

### Task

Generate a quarterly sales report.

### Competency Questions

- Can the system access sales data?
- Can it identify the reporting period?
- Can it locate the correct template?
- Can it determine the audience?

If the answer is "no," the failure is a context failure rather than a model failure.

---

## Potential Definition

> Context engineering is the discipline of ensuring that a system can answer its competency questions.

---

# Attention and Transformers

## Why Revisit Attention Is All You Need

Need to reread:

**Attention Is All You Need (2017)**

with focus on:

- what problem attention solved
- why transformers replaced RNNs
- why transformers replaced CNN-based sequence models
- relevance determination
- context engineering implications

---

## Current Understanding

### RNN / LSTM

Process information sequentially.

Problems:

- difficult to parallelize
- long-range dependencies
- hidden-state bottlenecks

---

### CNN Approaches

Process local relationships through convolutions.

Problems:

- limited receptive fields
- increasingly complex architectures for long-range relationships

---

### Transformers

Use self-attention.

No recurrence.

No convolution.

Each token can attend directly to other tokens.

---

## Important Correction

Initial misunderstanding:

> Transformers are feedforward only and therefore do not use backward passes.

Correction:

Transformers absolutely use backpropagation during training.

The distinction is:

### Training

```text
Forward Pass
Backward Pass
Gradient Update
```

### Inference

```text
Forward Pass Only
```

The transformer innovation was eliminating recurrence, not eliminating backpropagation.

---

## Attention as Relevance Weighting

Attention is not exactly a filter.

A better description:

> Attention is a learned relevance-weighting mechanism.

Instead of:

```text
Keep
Discard
```

Attention performs:

```text
Very Relevant
Somewhat Relevant
Slightly Relevant
Not Relevant
```

for every token relationship.

---

## Context Engineering Connection

Transformer:

```text
Large Context
      ↓
Attention
      ↓
Prediction
```

Context Engineering:

```text
Task
      ↓
Classification
      ↓
Retrieval
      ↓
Semantic Constraints
      ↓
Context Assembly
      ↓
Transformer
```

Potential thesis:

> Context engineering attempts to solve relevance before inference, reducing the burden placed on the model's attention mechanism.

---

# Large Context Windows

Large context windows reduce retrieval failures.

However, they may increase relevance failures.

Potential observation:

> Large context windows reduce missing-information problems but increase information-selection problems.

This appears highly relevant to:

- agent systems
- tool calling
- memory systems
- reasoning chains

---

# BERT as Historical Context

## Timeline

```text
2017  Attention Is All You Need
2018  BERT
2020  GPT-3
2022  ChatGPT
2023+ Agent Systems
```

BERT is particularly important because it appeared around the time I entered the software industry.

It provides historical context for understanding the rise of transformer-based semantic representations.

---

## Why BERT Matters

BERT is an encoder model.

Primary focus:

- semantic representations
- entity relationships
- classification
- retrieval
- semantic search

Not text generation.

---

# KG-BERT and Knowledge Graph Compression

KG-BERT suggests an interesting architectural alternative.

Traditional:

```text
Knowledge Graph
       ↓
Traversal
       ↓
Answer
```

KG-BERT:

```text
Knowledge Graph
       ↓
Training
       ↓
Embeddings
       ↓
Prediction
```

---

## Important Observation

The model becomes a lossy semantic compression of the graph.

Not a graph itself.

Not a replacement for the graph.

A compressed approximation of graph relationships.

---

## Tradeoff

### Knowledge Graph

Advantages:

- provenance
- governance
- explainability
- updates without retraining

---

### Model

Advantages:

- compression
- inference speed
- generalization

---

## Open Question

When should knowledge live:

- in weights?
- in retrieval systems?
- in governed knowledge stores?

This appears central to context engineering.

---

# Scaling Laws and Compression

Scaling laws become more intuitive when viewed through compression.

More data does not eliminate loss.

Instead:

> More data reduces approximation error.

The model remains lossy.

It simply becomes a better approximation of the underlying distribution.

---

# Few-Shot Learning vs Few-Shot Prompting

Important distinction.

These terms are frequently confused.

---

## Traditional Few-Shot Learning

Small training dataset.

Weights change.

Example:

```text
5 cat images
5 dog images
       ↓
Training
       ↓
Updated Model
```

---

## Few-Shot Prompting

Examples are placed in the prompt.

Weights do not change.

Example:

```text
Example 1
Example 2
Example 3

Now perform the task.
```

---

## In-Context Learning

Behavior changes due to context.

Weights remain unchanged.

This is what GPT-3 demonstrated.

---

## Important Clarification

The paper:

> Language Models are Few-Shot Learners

is largely demonstrating in-context learning rather than traditional few-shot training.

---

# Harnesses and Few-Shot Prompting

Few-shot prompting teaches behavior through examples.

Harnesses teach behavior through environments.

---

## Few-Shot Prompting

```text
Examples
      ↓
Prompt
      ↓
Model
```

---

## Harness

```text
Task
      ↓
Context Assembly
      ↓
Tool Definitions
      ↓
State
      ↓
Examples
      ↓
Model
      ↓
Validation
      ↓
Retry
```

---

## Observation

Harnesses generalize few-shot prompting.

Examples become one component of a larger behavioral system.

---

# Index and Retrieval Engineering

## Observation

A historical job title:

> Index and Retrieval Engineer

may better describe much of modern AI infrastructure work than "Data Engineer."

---

## Traditional Analytics Engineering

```text
Source Data
      ↓
Transform
      ↓
Metrics
      ↓
Dashboard
```

---

## Index and Retrieval Engineering

```text
Source Data
      ↓
Modeling
      ↓
Indexing
      ↓
Retrieval
      ↓
Inference
```

---

## Questions

- Can the system find the correct information?
- Can it retrieve it efficiently?
- Can it retrieve it securely?
- Can it retrieve it with sufficient semantic context?

---

## Potential Book Insight

Context engineering may be creating a new specialization within data engineering:

> Index and Retrieval Engineering

Core responsibilities:

- semantic modeling
- indexing
- retrieval
- authorization-aware retrieval
- provenance
- context assembly

---

# Emerging Book Thesis

Reliable AI systems require:

- semantics
- pragmatics
- governance
- retrieval
- structured state
- evaluation

Prompt engineering is only one mechanism among many.

The actual objective is:

> Produce predictable behavior under varying operational conditions.

Context engineering is the systems discipline responsible for achieving that goal.