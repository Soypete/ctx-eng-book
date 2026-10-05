# Cost-Aware Extraction Pipeline Notes

## ch14.08 — Sources queued for review

- D. Sculley et al. (2015). *Hidden Technical Debt in Machine Learning Systems*. Advances in Neural Information Processing Systems 28. https://papers.nips.cc/paper/2015/hash/86df7dcfd896fcaf2674f757a2463eba-Abstract.html
  - **Why it matters here:** identifies data dependencies, entanglement, undeclared consumers, and feedback loops as system-level maintenance risks beyond model inference cost.
- Neoklis Polyzotis, Sudip Roy, Steven Euijong Whang, and Martin Zinkevich (2017). *Data Management Challenges in Production Machine Learning*. Proceedings of the 2017 ACM International Conference on Management of Data. https://doi.org/10.1145/3035918.3054782
  - **Why it matters here:** frames validation, debugging, cleaning, and enrichment of production data as pipeline responsibilities, supporting a stage ledger and rebuild discipline.
- Saleema Amershi et al. (2019). *Software Engineering for Machine Learning: A Case Study*. Proceedings of ICSE (SEIP). https://doi.org/10.1109/ICSE-SEIP.2019.00042
  - **Why it matters here:** reports that data discovery, management, and versioning are unusually difficult in ML systems, supporting explicit artifact lineage and version keys.
- Eric Breck, Shanqing Cai, Eric Nielsen, Michael Salib, and D. Sculley (2017). *The ML Test Score: A Rubric for ML Production Readiness and Technical Debt Reduction*. IEEE. https://research.google.com/pubs/archive/aad9f93b86b7addfea4c419b9100c6cdd26cacea.pdf
  - **Why it matters here:** provides actionable production tests for monitoring, rollback, and data/model readiness, supporting release evidence beyond throughput or accuracy.
