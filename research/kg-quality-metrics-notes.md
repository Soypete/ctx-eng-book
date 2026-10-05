# Knowledge Graph Quality Metrics for Context Engineering

## ch08.09 — Sources queued for review

- Sewon Min, Kalpesh Krishna, Xinxi Lyu, Mike Lewis, Wen-tau Yih, Pang Wei Koh, Mohit Iyyer, Luke Zettlemoyer, and Hannaneh Hajishirzi (2023). *FActScore: Fine-grained Atomic Evaluation of Factual Precision in Long Form Text Generation*. EMNLP. https://arxiv.org/abs/2305.14251
  - **Why it matters here:** decomposes generated content into atomic facts and evaluates support, supporting claim-level source entailment and the module's rejection of a single undifferentiated confidence score.
- Potsawee Manakul, Adian Liusie, and Mark J. F. Gales (2023). *SelfCheckGPT: Zero-Resource Black-Box Hallucination Detection for Generative Large Language Models*. EMNLP. https://arxiv.org/abs/2303.08896
  - **Why it matters here:** demonstrates a sampling-based consistency signal for black-box outputs, useful as one uncertainty feature but not a replacement for source and authority validation.
- Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv:2309.15217. https://arxiv.org/abs/2309.15217
  - **Why it matters here:** separates faithfulness, answer relevance, and context relevance, supporting the module's recommendation to measure validation and downstream escape as separate outcomes.

## ch09.05 — Sources queued for review

- Nandan Thakur, Nils Reimers, Andreas Rücklé, Abhishek Srivastava, and Iryna Gurevych (2021, arXiv v4 read 2026-09-18). *BEIR: A Heterogenous Benchmark for Zero-shot Evaluation of Information Retrieval Models*. NeurIPS. https://arxiv.org/abs/2104.08663
  - **Why it matters here:** evaluates lexical, sparse, dense, late-interaction, and reranking systems across heterogeneous tasks, supporting bounded judged corpora and cross-task caution.
- Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv:2309.15217. https://arxiv.org/abs/2309.15217
  - **Why it matters here:** separates context relevance, answer relevance, and faithfulness, supporting the module's distinction between retrieval/assembly and answer support.
- Sewon Min, Kalpesh Krishna, Xinxi Lyu, Mike Lewis, Wen-tau Yih, Pang Wei Koh, Mohit Iyyer, Luke Zettlemoyer, and Hannaneh Hajishirzi (2023). *FActScore: Fine-grained Atomic Evaluation of Factual Precision in Long Form Text Generation*. EMNLP. https://arxiv.org/abs/2305.14251
  - **Why it matters here:** evaluates atomic claim support, supporting claim-to-source measurement after retrieval and context assembly.

## ch08.06 — Sources queued for review

- Axel-Cyrille Ngonga Ngomo, Irini Fundulaki, Anastasia Krithara, Mohammad Rashid, Marco Torchiano, Giuseppe Rizzo, Nandana Mihindukulasooriya, and Oscar Corcho (2019). *A Quality Assessment Approach for Evolving Knowledge Bases*. Semantic Web. https://doi.org/10.3233/SW-180324
  - **Why it matters here:** distinguishes schema, property, population, and interlinking completeness and defines property completeness relative to a class and release, matching the module's requirement-specific denominator.
- Subhi Issa, Onaopepo Adekunle, Fayçal Hamdi, Samira Si-Said Cherfi, Michel Dumontier, and Amrapali Zaveri (2021). *Knowledge Graph Completeness: A Systematic Literature Review*. IEEE Access. https://doi.org/10.1109/ACCESS.2021.3056622
  - **Why it matters here:** surveys completeness as a distinct quality dimension and supports keeping applicability, scope, and intended use explicit.
- Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv:2309.15217. https://arxiv.org/abs/2309.15217
  - **Why it matters here:** separates context relevance, answer relevance, and faithfulness, supporting the module's warning that graph property fill rates do not uniformly predict downstream retrieval or task quality.

## ch08.05 — Sources queued for review

- Pascal Hitzler, Amrapali Zaveri, Anisa Rula, Andrea Maurino, Ricardo Pietrobon, Jens Lehmann, and Sören Auer (2016). *Quality Assessment for Linked Data: A Survey*. Semantic Web. https://doi.org/10.3233/SW-150175
  - **Why it matters here:** organizes quality dimensions including completeness, accuracy, consistency, timeliness, provenance, and accessibility, supporting the module's warning that population coverage is only one part of graph readiness.
- Subhi Issa, Onaopepo Adekunle, Fayçal Hamdi, Samira Si-Said Cherfi, Michel Dumontier, and Amrapali Zaveri (2021). *Knowledge Graph Completeness: A Systematic Literature Review*. IEEE Access. https://doi.org/10.1109/ACCESS.2021.3056622
  - **Why it matters here:** surveys completeness as a distinct knowledge-graph quality dimension and supports treating denominators and completeness claims as explicit evaluation choices.
- Philipp Cimiano and Heiko Paulheim (2016). *Knowledge Graph Refinement: A Survey of Approaches and Evaluation Methods*. Semantic Web. https://doi.org/10.3233/SW-160218
  - **Why it matters here:** connects graph quality to refinement and evaluation workflows, supporting correction propagation, rebuilds, and competency-question testing.

Research notes on measuring knowledge graph quality as a systems discipline.

> **Convention:** This note follows the principles-first paradigm.
> - Core principles (timeless) are marked with "## Core Principle"
> - Current technology snapshots are marked with "## Current Implementation"

---

## Core Principle: Measurement as Missing Discipline (Timeless)

## Core Insight: Measurement as Missing Discipline

One of the missing disciplines in context engineering is **measurement**.

Databases have metrics. Distributed systems have metrics. Networks have metrics. Search systems have precision and recall.

Knowledge graphs, however, are often discussed qualitatively instead of quantitatively, especially when used as semantic layers for AI systems.

If context engineering is going to become a systems discipline, it needs measurable engineering metrics.

---

## Precision vs Coverage vs Completeness

### Information Retrieval

Precision means:

> Of the retrieved documents, how many were relevant?

This metric evaluates search systems. It does **not** describe a knowledge graph itself.

### Wikidata

Wikidata defines precision differently. Example:

```json
{
  "time": "+2026-07-14",
  "precision": 11
}
```

Precision indicates temporal granularity (century, decade, year, month, day, hour, minute, second). This is **value precision**, not graph quality.

### Knowledge Graphs

Knowledge graphs introduce another meaning. The precision of the ontology determines how well the graph represents concepts and supports inference.

Examples:

- Poor precision: Employee
- Better precision: Faculty Member → Adjunct Professor → Research Professor

These distinctions enable richer reasoning.

---

## A Context Engineering Perspective

The graph should **not** be the database. The graph should be the semantic layer over a distributed data mesh.

Instead of storing every fact inside Neo4j, the graph stores:
- ontology
- semantic relationships
- identifiers
- provenance
- authorization boundaries
- source mappings

Operational data remains in: PostgreSQL, Iceberg, Delta Lake, APIs, Event Streams, Object Storage, Search indexes.

The graph becomes the map rather than the territory.

---

## Slowly Changing Dimensions

Traditional dimensional modeling introduces Slowly Changing Dimensions (SCDs) that map well onto graph storage because they change infrequently.

Examples of excellent graph nodes:
- organizations
- people
- departments
- products
- taxonomies
- permissions
- locations

High-volume operational data should remain in relational or analytical storage:
- transactions
- chat history
- telemetry
- events
- logs
- mutable business records

The graph links to these systems rather than duplicating them.

---

## Graph as Semantic Router

An agent should traverse the graph to discover:
- where information lives
- who owns it
- authorization requirements
- retrieval method
- semantic relationships

Rather than asking "Give me all customer information", the graph answers:
- "The CRM owns customer identity."
- "Billing owns invoices."
- "Support owns tickets."

The graph routes retrieval.

---

## Measuring Knowledge Graph Quality

### 1. Semantic Coverage

Question: How much of the domain has been modeled?

Measures:
- ontology coverage
- relationship coverage
- property coverage
- source coverage

This is a design-time metric.

### 2. Instance Coverage

Question: How much of the modeled domain contains actual data?

Example: Customer nodes → Orders → Invoices → Products

A populated ontology has higher instance coverage.

### 3. Hydration Coverage

Question: How much of the required context was successfully retrieved?

Example:

```
Customer
Known sources:
  ✓ CRM
  ✓ Billing
  ✓ Identity
  ✓ Support

Retrieved:
  ✓ CRM
  ✓ Identity
  ✗ Billing
  ✗ Support

Hydration coverage: 2/4
```

This becomes a runtime metric.

### 4. Property Completeness

Question: How many expected semantic properties exist?

Person - Expected: employer, education, nationality, advisor

If only employer exists, inference quality decreases.

### 5. Inferential Reach

Question: How much reasoning can the graph support?

Metrics:
- multi-hop traversal success
- competency questions answered
- reachable semantic depth
- average traversal depth

This may become one of the most important metrics for AI systems.

### 6. Context Precision

Not information retrieval precision. Instead: Of the retrieved context, how much was actually useful?

Too much irrelevant context increases:
- cost
- latency
- attention competition
- hallucination risk

Conceptually:

```
Context Precision = Useful Context / Retrieved Context
```

### 7. Context Recall

Question: Did the retrieval system miss important context?

Conceptually:

```
Relevant Context Retrieved / Relevant Context Available
```

This will be difficult to measure because the denominator is rarely known.

### 8. Provenance Coverage

Question: How much retrieved information includes provenance?

Examples: source, owner, timestamp, confidence, authorization

Reliable systems require provenance.

### 9. Authorization Coverage

Question: How much retrieved context was both authorized and necessary rather than simply retrieved?

This aligns directly with least-privilege retrieval.

### 10. Context Efficiency

Question: How efficiently was context produced?

Metrics:
- tokens
- retrieval latency
- retrieval cost
- hydration time
- API calls
- storage lookups

If two systems answer equally well, the cheaper system is objectively better.

---

## Relationship to Reliability

Traditional AI evaluation focuses on outputs:
- accuracy
- hallucination rate
- helpfulness

Context engineering should also evaluate inputs.

Questions become:
- Did the system provide enough semantic context?
- Was the context relevant?
- Was it fresh?
- Was it authorized?
- Was provenance preserved?
- Was retrieval efficient?

These become leading indicators of AI reliability.

---

## Open Research Questions

- Can semantic coverage predict downstream model accuracy?
- Can hydration coverage predict hallucination rates?
- Can inferential reach predict agent capability?
- Can context precision predict token efficiency?
- Can provenance coverage predict user trust?
- Can these metrics become standardized for evaluating production AI systems?

---

## Working Hypothesis

Reliable AI systems should not be evaluated solely by model outputs. They should also be evaluated by measurable properties of the engineered context supplied to the model.

Knowledge graphs, retrieval systems, authorization layers, and semantic indexes should all expose engineering metrics analogous to those used for databases and distributed systems.
