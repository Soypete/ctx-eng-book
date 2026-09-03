# Chapter 8: Knowledge Graphs and Semantic Context

## Semantics Is the Middle of the Context Pipeline

Chapter 7 established the Lexicon side of the problem: context is assembled
from distributed sources, records, documents, APIs, indexes, and event
streams. Those inputs are available only as far as their source identity,
authority, scope, and freshness can be established. This chapter addresses
the next question:

> What do those entities, attributes, and relationships mean in this domain,
> for this task, and at this time?

That is the Semantics layer of the book's Lexicon → Semantics → Pragmatics
spine. It is not a fifth pillar, a graph product category, or a requirement
that every application adopt a knowledge graph. Semantics can be carried by a
relational schema, API contract, taxonomy, controlled vocabulary, policy
model, event model, ontology, graph, or a deliberate combination of them.
Choose the least elaborate representation that makes the decision-relevant
meaning explicit and testable.

## What This Chapter Builds

The chapter moves from semantic design to operational use:

```text
Lexicon inputs
  -> schemas, taxonomies, ontologies, and identifiers
  -> relationships, provenance, and temporal semantics
  -> population, extraction, resolution, and validation
  -> scoped semantic context
  -> Pragmatic retrieval, answers, classifications, and actions
```

The center of gravity is not graph storage. It is the semantic infrastructure
that lets a system preserve distinctions a model would otherwise have to
reconstruct from prose:

- schemas define shape and structural expectations;
- taxonomies define controlled classification;
- ontologies define concepts, relationships, and formal entailments;
- identity resolution decides which records refer to the same thing;
- provenance and temporal semantics say who asserted what, from which source,
  and when it applies;
- extraction turns source material into proposals rather than unquestioned
  facts; and
- validation decides what may become durable, retrievable state.

Knowledge graphs are particularly useful when relationship-centric questions,
shared identifiers, variable-depth traversal, or relationship-level evidence
justify their operational cost. They are not universal memory, a replacement
for authoritative stores, or a guarantee of truth. A graph projection can be
stale, incomplete, incorrectly resolved, unauthorized, or wrong. The tradeoff
module makes that decision explicit.

## The Downstream Contract

Semantics exists to support Pragmatics. Chapter 4's prompts and structured
outputs, Chapter 5's tools, Chapter 9's retrieval strategies, and Chapter 10's
authorization and validation boundaries all consume the distinctions made
here. A semantic representation should therefore expose not only a label or
path, but the evidence, identity, time interval, authority, uncertainty, and
scope needed to decide whether an output is appropriate for an actor and
purpose.

The model is still a participant in this process. It may interpret a bounded
semantic context, propose an extraction, or formulate a query. It does not get
to silently resolve ambiguous identities, turn absence into negation, invent
missing qualifiers, or promote a plausible proposal into shared state. Those
responsibilities belong at the surrounding system's semantic and pragmatic
boundaries.

## Module Map

1. **Schemas, taxonomies, and ontologies** establish the vocabulary and
   modeling decisions.
2. **RDF, OWL, SPARQL, and shapes** distinguish representation, entailment,
   querying, and operational validation.
3. **Entity resolution and relationship traversal** make identity and scoped
   paths reliable.
4. **Knowledge-graph tradeoffs** define when a graph earns its cost.
5. **Instance coverage and ontology population** separate populated state from
   successful context assembly.
6. **Property completeness and schema quality** make missingness meaningful.
7. **Ontology-guided extraction** compiles semantics into a bounded proposal
   contract.
8. **Extraction methods** choose a field- and consequence-specific pipeline.
9. **Extraction validation** sets the promotion boundary.
10. **Multilingual extraction** evaluates the complete semantic path by
    language and failure consequence.

Chapter 9 then asks how to retrieve the right evidence under a context budget.
The result is not “put everything in a graph.” It is a governed semantic
surface that helps downstream language and information processes interpret
the right things, with the right qualifications, for the right purpose.
