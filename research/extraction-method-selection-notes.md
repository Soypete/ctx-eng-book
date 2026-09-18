# Extraction Method Selection Notes

## ch14.07 — Sources queued for review

- Ying Lin, Heng Ji, Fei Huang, and Lingfei Wu (2020). *A Joint Neural Model for Information Extraction with Global Features*. Proceedings of ACL. https://aclanthology.org/2020.acl-main.713/
  - **Why it matters here:** provides a primary example of jointly modeling entities, triggers, and links, supporting route selection by assertion structure rather than by document format alone.
- Chenguang Wang, Xiao Liu, Zui Chen, Haoyun Hong, Jie Tang, and Dawn Song (2021). *Zero-Shot Information Extraction as a Unified Text-to-Triple Translation*. Proceedings of EMNLP. https://aclanthology.org/2021.emnlp-main.94/
  - **Why it matters here:** tests a unified extraction representation across open extraction and relation tasks, useful for comparing a shared contract with task-specific routes.
- Tingyu Xie, Qi Li, Jian Zhang, Yan Zhang, Zuozhu Liu, and Hongwei Wang (2023). *Empirical Study of Zero-Shot NER with ChatGPT*. Proceedings of EMNLP. https://aclanthology.org/2023.emnlp-main.493/
  - **Why it matters here:** reports error types and consistency strategies for zero-shot NER, supporting explicit validation and escalation rather than unexamined model agreement.
- Lingjiao Chen, Matei Zaharia, and James Zou (2023). *FrugalGPT: How to Use Large Language Models While Reducing Cost and Improving Performance*. arXiv. https://arxiv.org/abs/2305.05176
  - **Why it matters here:** studies cascades and routing across models, supporting the claim that a router must earn its complexity on measured quality and cost.
