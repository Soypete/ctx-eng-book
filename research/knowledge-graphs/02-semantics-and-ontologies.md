# Semantics and Ontologies

## Core Insight

The primary lesson from the book is that **semantics describe meaning while schemas describe storage**.

A relational schema tells us where data lives. An ontology explains what the data means and how concepts relate to one another.

This distinction is fundamental to Context Engineering because reliable retrieval depends on preserving meaning rather than preserving implementation.

## Major Takeaways

- RDF represents knowledge as subject–predicate–object triples.
- URIs provide globally unique identities.
- Ontologies define classes, properties, and constraints.
- Reasoning engines infer new knowledge from existing relationships.
- Semantic modeling is portable across storage technologies.

## Context Engineering Connections

Rather than treating ontologies as an academic exercise, they can be viewed as semantic contracts for AI systems.

They define:

- allowed concepts
- valid relationships
- retrieval boundaries
- inference constraints

The implementation may be PostgreSQL, Neo4j, or object storage. The semantic layer remains stable.

## Questions

- Which semantic concepts belong in enterprise ontologies?
- Can ontologies be generated from existing data catalogs?
- How much semantic richness is necessary before diminishing returns?