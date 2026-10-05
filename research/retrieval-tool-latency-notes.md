# Retrieval, Tool, and Latency Notes

## ch14.02 — Sources queued for review

- Jeffrey Dean and Luiz André Barroso (2013). *The Tail at Scale*. Communications of the ACM. https://doi.org/10.1145/2408776.2408794
  - **Why it matters here:** analyzes tail latency in fan-out services and techniques for tolerating latency variability; relevant to p95/p99 budgets and the difference between serial critical paths and parallel work.
- Vladimir Karpukhin et al. (2020). *Dense Passage Retrieval for Open-Domain Question Answering*. Proceedings of EMNLP 2020. https://arxiv.org/abs/2004.04906
  - **Why it matters here:** provides a primary dense-retrieval baseline and evaluates passage-retrieval accuracy; relevant to comparing retrieval quality against latency and cost rather than treating “semantic search” as a free upgrade.
- Gordon V. Cormack, Charles L. A. Clarke, and Stefan Büttcher (2009). *Reciprocal Rank Fusion Outperforms Condorcet and Individual Rank Learning Methods*. Proceedings of SIGIR 2009. https://doi.org/10.1145/1571941.1572114
  - **Why it matters here:** evaluates a simple rank-fusion method across information-retrieval systems; relevant to the cost/quality tradeoff of adding fusion before reranking.
- Omar Khattab and Matei Zaharia (2020). *ColBERT: Efficient and Effective Passage Search via Contextualized Late Interaction over BERT*. Proceedings of SIGIR 2020. https://arxiv.org/abs/2004.12832
  - **Why it matters here:** separates document encoding from query-time interaction to reduce reranking cost; relevant to the module's retrieval ladder and latency budget.
