# Semantic and Retrieval Infrastructure Notes

## ch18.02 — Sources queued for review

- Vladimir Karpukhin, Barlas Oğuz, Sewon Min, Patrick Lewis, Ledell Wu, Sergey Edunov, Danqi Chen, and Wen-tau Yih (2020). *Dense Passage Retrieval for Open-Domain Question Answering*. EMNLP. <https://arxiv.org/abs/2004.04906>
  - **Why it matters here:** gives a primary dense-retrieval design and compares it with a lexical BM25 baseline, supporting route-specific evaluation rather than one undifferentiated search score.
  - **Claim it would support:** “A system can combine relational, lexical, vector, graph, event, or application-specific routes without pretending that one index is the meaning layer.”
- Stephen E. Robertson and Hugo Zaragoza (2009). *The Probabilistic Relevance Framework: BM25 and Beyond*. Foundations and Trends in Information Retrieval. <https://doi.org/10.1561/1500000019>
  - **Why it matters here:** grounds lexical ranking as a particular scoring model with assumptions, not a universal relevance scale.
  - **Claim it would support:** “Do not compare BM25, vector similarity, graph distance, and business priority as if they shared a scale.”
- Gordon V. Cormack, Charles L. A. Clarke, and Stefan Büttcher (2009). *Reciprocal Rank Fusion Outperforms Condorcet and Individual Rank Learning Methods*. SIGIR. <https://doi.org/10.1145/1571941.1572114>
  - **Why it matters here:** provides a primary rank-fusion method for combining result lists, supporting an explicit fusion policy and route contribution accounting.
  - **Claim it would support:** “Normalize through an explicit fusion or reranking policy and preserve each route’s contribution.”
- Patrick Lewis et al. (2020). *Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks*. NeurIPS. <https://arxiv.org/abs/2005.11401>
  - **Why it matters here:** separates retrieval from generation and evaluates retrieved evidence as part of a larger task pipeline.
  - **Claim it would support:** “The retrieval service returns evidence and declared uncertainty—not a prompt-sized pile of text.”
