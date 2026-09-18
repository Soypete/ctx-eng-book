# Entity Resolution and Relation Extraction

## ch08.03 — Sources queued for review

- Olivier Binette and Rebecca C. Steorts (2020). *(Almost) All of Entity Resolution*. arXiv:2008.04443. https://arxiv.org/abs/2008.04443
  - **Why it matters here:** reviews record linkage, deduplication, clustering, canonicalization, and the practical difficulty of integrating records without unique identifiers.
- Matt Barnes (2015). *A Practitioner's Guide to Evaluating Entity Resolution Results*. arXiv:1509.04238. https://arxiv.org/abs/1509.04238
  - **Why it matters here:** surveys entity-resolution evaluation metrics and warns that rankings can conflict, supporting consequence-specific evaluation rather than a single match score.
- Rostislav Nedelchev, Debanjan Chaudhuri, Jens Lehmann, and Asja Fischer (2020). *End-to-End Entity Linking and Disambiguation leveraging Word and Knowledge Graph Embeddings*. arXiv:2002.11143. https://arxiv.org/abs/2002.11143
  - **Why it matters here:** studies entity linking as connecting mentions to graph entities and uses relational context for disambiguation, matching the curriculum's entity-linking and graph-neighborhood handoff.

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
