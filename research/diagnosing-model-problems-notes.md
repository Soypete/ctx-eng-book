# Diagnosing Model Problems Notes

## ch15.01 — Sources queued for review

- Marco Tulio Ribeiro, Tongshuang Wu, Carlos Guestrin, and Sameer Singh (2020). *Beyond Accuracy: Behavioral Testing of NLP Models with CheckList*. Proceedings of ACL. https://aclanthology.org/2020.acl-main.442/
  - **Why it matters here:** demonstrates that held-out accuracy can miss actionable behavioral failures and provides a way to test capabilities with controlled cases.
- Pang Wei Koh et al. (2021). *WILDS: A Benchmark of in-the-Wild Distribution Shifts*. Proceedings of ICML. https://arxiv.org/abs/2012.07421
  - **Why it matters here:** shows that in-distribution performance can overstate deployment performance under naturally occurring shifts, supporting slice and distribution diagnostics.
- Margaret Mitchell et al. (2019). *Model Cards for Model Reporting*. Proceedings of FAT*. https://arxiv.org/abs/1810.03993
  - **Why it matters here:** proposes documenting intended use and performance across relevant conditions, supporting a complete model-boundary diagnosis rather than a single score.
- Yaniv Ovadia et al. (2019). *Can You Trust Your Model's Uncertainty? Evaluating Predictive Uncertainty Under Dataset Shift*. arXiv. https://arxiv.org/abs/1906.02530
  - **Why it matters here:** tests calibration and uncertainty under distribution shift, supporting the warning that confidence and abstention require their own evidence.
