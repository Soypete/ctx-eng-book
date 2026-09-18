# Semantics and Ontologies

## ch08.01 — Sources queued for review

- W3C OWL Working Group (2012). *OWL 2 Web Ontology Language Primer (Second Edition)*. W3C Recommendation. https://www.w3.org/TR/owl2-primer/
  - **Why it matters here:** defines OWL as a logic-based language for representing knowledge and reasoning about implicit knowledge, supporting the module's distinction between ontology semantics and ordinary schema validation.
- W3C RDF Data Shapes Working Group (2017). *Shapes Constraint Language (SHACL)*. W3C Recommendation. https://www.w3.org/TR/shacl/
  - **Why it matters here:** defines a separate validation mechanism for RDF graphs, supporting the module's claim that operational acceptance constraints need a boundary distinct from ontology entailment.
- C. Maria Keet and Paolo D. D. Ferrario (2013). *Evaluating Ontologies with Competency Questions*. IEEE/WIC/ACM International Joint Conferences on Web Intelligence and Intelligent Agent Technology. https://doi.org/10.1109/WI-IAT.2013.199
  - **Why it matters here:** treats competency questions as ontology requirements and evaluation prompts, supporting the module's recommendation to model only what changes a decision.

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
