# Knowledge Graph Quality Metrics for Context Engineering

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