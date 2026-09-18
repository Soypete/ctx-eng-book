# Local Models and Model Routing Notes

## ch14.04 — Sources queued for review

- Tim Dettmers et al. (2022). *LLM.int8(): 8-bit Matrix Multiplication for Transformers at Scale*. arXiv. https://arxiv.org/abs/2208.07339
  - **Why it matters here:** demonstrates a way to reduce inference memory while retaining model performance under a specified quantization procedure; relevant to evaluating local serving as a resource tradeoff rather than assuming smaller hardware is free.
- Guangxuan Xiao et al. (2022). *SmoothQuant: Accurate and Efficient Post-Training Quantization for Large Language Models*. arXiv. https://arxiv.org/abs/2211.10438
  - **Why it matters here:** studies memory and speed effects of post-training quantization; relevant to the module's warning that precision changes must be evaluated on the deployed workload.
- Lingjiao Chen, Matei Zaharia, and James Zou (2023). *FrugalGPT: How to Use Large Language Models While Reducing Cost and Improving Performance*. arXiv. https://arxiv.org/abs/2305.05176
  - **Why it matters here:** formalizes prompt adaptation, approximation, and cascades as cost/quality strategies; relevant to model routing and escalation accounting.
- Isaac Ong et al. (2024). *RouteLLM: Learning to Route LLMs with Preference Data*. arXiv (version 4, 2025). https://arxiv.org/abs/2406.18665
  - **Why it matters here:** evaluates learned routers for choosing between stronger and weaker models under quality/cost tradeoffs; relevant to calibration, route-level evaluation, and routing regret.
