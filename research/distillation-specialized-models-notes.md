# Distillation and Specialized Models Notes

## ch15.03 — Sources queued for review

- Victor Sanh, Lysandre Debut, Julien Chaumond, and Thomas Wolf (2019). *DistilBERT, a distilled version of BERT: smaller, faster, cheaper and lighter*. arXiv. https://arxiv.org/abs/1910.01108
  - **Why it matters here:** reports a concrete teacher-to-student compression result with measured size, speed, and task performance rather than treating “smaller” as an abstract benefit.
- Xiaoqi Jiao et al. (2019). *TinyBERT: Distilling BERT for Natural Language Understanding*. arXiv. https://arxiv.org/abs/1909.10351
  - **Why it matters here:** separates general-domain and task-specific distillation stages, supporting explicit target behavior and evaluation.
- Zhiqing Sun et al. (2020). *MobileBERT: a Compact Task-Agnostic BERT for Resource-Limited Devices*. arXiv. https://arxiv.org/abs/2004.02984
  - **Why it matters here:** evaluates compression against device latency and task quality, supporting deployment-specific benchmarking.
- Cheng-Yu Hsieh et al. (2023). *Distilling Step-by-Step! Outperforming Larger Language Models with Less Training Data and Smaller Model Sizes*. arXiv. https://arxiv.org/abs/2305.02301
  - **Why it matters here:** shows that teacher-generated rationales and task-specific supervision can change the data and quality tradeoff, while remaining a measured result for particular tasks.
