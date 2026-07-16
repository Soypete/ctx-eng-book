# Entity Resolution and Relation Extraction

## Core Insight

Building a knowledge graph is primarily a data engineering problem.

The most difficult work is not storing relationships but identifying entities correctly and extracting reliable relationships.

## Topics

### Entity Resolution

Challenges include:

- duplicate entities
- aliases
- canonical identifiers
- provenance
- confidence scoring

Reliable AI begins with stable identities.

### Relation Extraction

The book discusses both traditional supervised approaches and distant supervision.

One particularly interesting observation is that an existing knowledge graph can automatically label text, allowing the graph itself to generate training data.

This reverses the common assumption that NLP always produces the graph.

## Positional Embeddings

Older relation extraction systems relied on positional embeddings because the distance between words often indicated relationships.

Transformers moved much of this capability into attention, but explicit structure still has value for smaller models.

## Context Engineering Connections

- Stable identities improve retrieval.
- Explicit relationships reduce ambiguity.
- Semantic extraction can feed governed knowledge stores.