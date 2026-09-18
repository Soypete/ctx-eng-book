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

## ch09.01 — Sources queued for review

- Stephen E. Robertson and Hugo Zaragoza (2009). *The Probabilistic Relevance Framework: BM25 and Beyond*. Foundations and Trends in Information Retrieval. https://doi.org/10.1561/1500000019
  - **Why it matters here:** provides the primary account of the probabilistic relevance framework and BM25 assumptions, supporting the module's bounded claims about term frequency, document frequency, and length normalization.
- Gordon V. Cormack, Charles L. A. Clarke, and Stefan Büttcher (2009). *Reciprocal Rank Fusion Outperforms Condorcet and Individual Rank Learning Methods*. Proceedings of SIGIR. https://doi.org/10.1145/1571941.1572114
  - **Why it matters here:** evaluates a rank-fusion method, supporting the module's distinction between candidate generation and later fusion when lexical and other result lists are combined.
- *COIL: Revisit Exact Lexical Match in Information Retrieval with Contextualized Inverted List* (2021). Proceedings of NAACL. https://aclanthology.org/2021.naacl-main.241/
  - **Why it matters here:** revisits exact lexical matching with contextualized representations, supporting the module's warning that “lexical” behavior depends on the configured index and scoring design.
