# Reliability Metrics and Failure Budgets Notes

## ch17.04 — Sources queued for review

- Betsy Beyer et al. (2016). *Site Reliability Engineering: How Google Runs Production Systems*. Google. <https://sre.google/sre-book/service-level-objectives/>
  - **Why it matters here:** defines SLIs, SLOs, and error budgets as an operational decision system rather than a model-quality score.
  - **Claim it would support:** “An objective states the acceptable level over a window. A failure budget is the permitted shortfall from that objective.”
- Eric Breck, Marty Zinkevich, Neoklis Polyzotis, Steven Whang, and Sudip Roy (2019). *Data Validation for Machine Learning*. Proceedings of SysML. <https://research.google/pubs/data-validation-for-machine-learning/>
  - **Why it matters here:** presents production validation for incoming ML data and reports early detection of data problems as a reliability control.
  - **Claim it would support:** “Production AI needs metrics that connect system boundaries to user-visible and consequential outcomes.”
- Bradley Eck, Duygu Kabakci-Zorlu, Yan Chen, France Savard, and Xiaowei Bao (2022). *A monitoring framework for deployed machine learning models with supply chain examples*. IEEE Big Data. <https://arxiv.org/abs/2211.06239>
  - **Why it matters here:** compares drift signals with measured model performance and shows why distribution movement is not itself an outcome failure.
  - **Claim it would support:** “Treat metric movement as a trigger for investigation, not as an explanation.”
- Doris Xin, Hui Miao, Aditya Parameswaran, and Neoklis Polyzotis (2021). *Production Machine Learning Pipelines: Empirical Analysis and Optimization Opportunities*. ACM SIGMOD. <https://research.google/pubs/production-machine-learning-pipelines-empirical-analysis-and-optimization-opportunities/>
  - **Why it matters here:** documents the complexity and repeated components of production ML pipelines, supporting provenance and cost-aware operational metrics.
  - **Claim it would support:** “Version datasets, prompts, judges, labels, and scoring code.”
