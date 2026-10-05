# Source and Ingestion Architecture Notes

## ch18.01 — Sources queued for review

- Eric Breck, Marty Zinkevich, Neoklis Polyzotis, Steven Whang, and Sudip Roy (2019). *Data Validation for Machine Learning*. Proceedings of SysML. <https://research.google/pubs/data-validation-for-machine-learning/>
  - **Why it matters here:** presents production validation for incoming ML data and the operational consequences of schema-free data and training/serving skew.
  - **Claim it would support:** “An ingestion record should identify the source and version, observed time, content digest, parser and schema versions, tenant and policy scope, deletion obligations, and processing outcome.”
- Martín Abadi et al. (2016). *TensorFlow: A System for Large-Scale Machine Learning*. arXiv. <https://arxiv.org/abs/1605.08695>
  - **Why it matters here:** describes a large-scale ML system whose graph and execution abstractions clarify why derived computation and deployment boundaries need explicit contracts.
  - **Claim it would support:** “A production context platform is not a required stack of databases.”
- Alkis Polyzotis, Martin A. Zinkevich, Steven Whang, and Sudip Roy (2017). *Data Management Challenges in Production Machine Learning*. SIGMOD. <https://research.google/pubs/data-management-challenges-in-production-machine-learning/>
  - **Why it matters here:** frames validation, debugging, cleaning, understanding, and enrichment of production ML data as data-management concerns.
  - **Claim it would support:** “Derived data retains provenance and can be rebuilt when its source, schema, policy, or model changes.”
- Patrick Lewis et al. (2020). *Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks*. NeurIPS. <https://arxiv.org/abs/2005.11401>
  - **Why it matters here:** provides a primary example of separating parametric generation from retrieved evidence, supporting the distinction between source records, indexes, and task-time context.
  - **Claim it would support:** “The platform should distinguish authoritative source records from derived chunks, embeddings, entities, summaries, and memories.”
