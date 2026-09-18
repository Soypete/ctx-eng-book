# Guardrails and Validation Notes

## ch10.00 — Sources queued for review

- Saibo Geng, Martin Josifoski, Maxime Peyrard, and Robert West (2023). *Grammar-Constrained Decoding for Structured NLP Tasks without Finetuning*. arXiv:2305.13971. https://arxiv.org/abs/2305.13971
  - **Why it matters here:** supports the distinction between structural validity enforced during decoding and semantic, authorization, and policy checks that still belong to the host.
- Eric Wallace, Kai Xiao, Reimar Leike, Lilian Weng, Johannes Heidecke, and Alex Beutel (2024). *The Instruction Hierarchy: Training LLMs to Prioritize Privileged Instructions*. arXiv:2404.13208. https://arxiv.org/abs/2404.13208
  - **Why it matters here:** evaluates instruction-priority training against prompt injection and over-refusal, supporting the module's distinction between model behavior and enforcement boundaries.
- Yen-Shan Chen, Sian-Yao Huang, Cheng-Lin Yang, and Yun-Nung Chen (2026). *TraceSafe: A Systematic Assessment of LLM Guardrails on Multi-Step Tool-Calling Trajectories*. arXiv:2604.07223. https://arxiv.org/abs/2604.07223
  - **Why it matters here:** evaluates guardrails over intermediate tool-use trajectories, supporting checks before side effects rather than only final-output filtering.
