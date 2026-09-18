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

## ch09.02 — Sources queued for review

- Patrick Lewis, Ethan Perez, Aleksandara Piktus, Fabio Petroni, Vladimir Karpukhin, Naman Goyal, Heinrich Küttler, Mike Lewis, Wen-tau Yih, Tim Rocktäschel, Sebastian Riedel, and Douwe Kiela (2020). *Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks*. arXiv:2005.11401. https://arxiv.org/abs/2005.11401
  - **Why it matters here:** defines retrieval-augmented generation as a retriever-plus-generator system and supplies the primary source for treating retrieved passages as context rather than authority.
- Vladimir Karpukhin, Barlas Oğuz, Sewon Min, Patrick Lewis, Ledell Wu, Sergey Edunov, Danqi Chen, and Wen-tau Yih (2020). *Dense Passage Retrieval for Open-Domain Question Answering*. arXiv:2004.04906. https://arxiv.org/abs/2004.04906
  - **Why it matters here:** provides a primary dual-encoder dense-retrieval account and evaluates dense passage candidates against a BM25 baseline.
- Nils Reimers and Iryna Gurevych (2019). *Sentence-BERT: Sentence Embeddings using Siamese BERT-Networks*. Proceedings of EMNLP. https://arxiv.org/abs/1908.10084
  - **Why it matters here:** grounds the module's explanation of reusable sentence embeddings and cosine-based semantic comparison.
- Yu. A. Malkov and D. A. Yashunin (2018). *Efficient and robust approximate nearest neighbor search using Hierarchical Navigable Small World graphs*. IEEE Transactions on Pattern Analysis and Machine Intelligence. https://arxiv.org/abs/1603.09320
  - **Why it matters here:** supplies the primary ANN index account for the module's latency/accuracy tradeoff and index-parameter evaluation claim.
