# Memstore Analysis: Synthesis Notes

**Source**: Soypete/memstore-analysis + personal communication
**Date**: July 2026

## Core Thesis: Domain-First Memory Architecture

The key insight from memstore-analysis is that building a knowledge graph before understanding your domain is premature optimization. The recommendation:

1. **Start with a database (read-only)**
   - Specifically curated agent data
   - APIs for user/agent access
   - Simple, well-understood structure

2. **Put semantics into business logic of tools**
   - As you build tools, you learn the domain
   - Semantics emerge from tool relationships, not upfront modeling

3. **Then build your knowledge store**
   - Once you have a set of tools you like
   - Better understanding of the domain
   - Can now model relationships meaningfully

4. **Why this order matters**
   - Without domain knowledge, RDF modeling is "basically impossible"
   - Can't reason about relationships you don't understand
   - KG is the end state, not the starting point

## Storage Layer Independence (Aligns with KG Book)

From the knowledge graph book (early chapters):
- RDF is the best ontology modeling language
- The store doesn't matter — triple store, LPG, or RDBMS
- As long as you have a model and can reason, you have a knowledge graph

This supports the domain-first approach: start with any store (RDBMS is simplest), build tools, learn domain, then evolve to KG when relationships are understood.

## Tool Categories & Scale Behavior

| Tool | Approach | Scale Behavior |
|------|----------|----------------|
| **MemPalace** | Database-style indexed retrieval | Millisecond latency, megabyte-scale — best at scale |
| **LLMWiki** | Markdown-based, domain-less search | Works at moderate scale, falls apart at scale (recursive queries) |
| **Graphify** | Full graph materialization | Second-level latency, gigabyte-scale — expensive at scale |

**Key insight**: Domain-less approaches (LLMWiki, MemPalace) work initially but fail at scale due to recursive queries. Models that understand how tools work can traverse paths via SQL or unix commands to get small data sets efficiently.

## Working Assumption: ~/code/pedro/pedro-tag

This approach informs the design of pedro-tag — starting with simple indexed retrieval, evolving toward graph-based reasoning as domain understanding grows.

## Alignment with Book Thesis

| Book Thesis | Memstore Analysis Alignment |
|-------------|----------------------------|
| "Reliable AI = engineered context" | Database-first provides explicit context before KG inference |
| "Relationships given to agent, no inference needed" | Domain-first approach ensures relationships are understood before being encoded |
| Pragmatics → Data → Semantics | Tools (pragmatics) → Database (data) → Domain model (semantics) |

## Sources

- github.com/Soypete/memstore-analysis
- Personal communication (July 2026)