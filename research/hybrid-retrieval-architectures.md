# Hybrid Retrieval Architectures

## Core Concept

Hybrid retrieval combines multiple search strategies to improve context quality.

### The Four Strategies

1. **Vector Search**: Semantic similarity via embeddings
2. **BM25/Keyword**: Traditional text matching
3. **RRF (Reciprocal Rank Fusion)**: Combine rankings from multiple methods
4. **Graph Traversal**: Navigate knowledge graph relationships

### Key Finding

> Four-strategy hybrid achieves **49.1 P@5** (vs 18 without graph)

Graph retrieval significantly boosts performance because it retrieves connected context that vector search cannot find.

### Why Graph Helps

- Vector search finds similar content
- Graph traversal finds *related* content (via relationships)
- Combines "what's similar" with "what's connected"

### Connection to Web IE

Web IE produces the **structured data** that enables graph retrieval:

```
IE Output → Entities + Relations → Knowledge Graph → Graph Traversal → Context
```

Without IE, there's no graph to traverse.

---

## Evidence Ledger

See `research/_evidence-ledger.md`:
- Graph & Hybrid Retrieval — Four-strategy approach
- Context Precision/Recall — Retrieval quality metrics

---

## Cross-References

- **Chapter 9**: Retrieval Beyond Vector Databases
- **Chapter 8**: Knowledge Graphs enable graph traversal
- **Chapter 7**: Context assembly uses hybrid retrieval
- **webinformationextraction.md**: IE produces graph-ready data