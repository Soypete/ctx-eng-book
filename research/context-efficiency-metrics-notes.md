# Context Efficiency Metrics Notes

## ch14.05 — Sources queued for review

- Percy Liang et al. (2022). *Holistic Evaluation of Language Models*. arXiv. https://arxiv.org/abs/2211.09110
  - **Why it matters here:** defines scenario-based, multi-metric evaluation and exposes trade-offs beyond accuracy; relevant to rejecting a universal context-efficiency scalar.
- Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv. https://arxiv.org/abs/2309.15217
  - **Why it matters here:** separates retrieval context quality, faithfulness, and generation quality; relevant to preserving resource and outcome dimensions across the context pipeline.
- Xiao Liu et al. (2023). *AgentBench: Evaluating LLMs as Agents*. arXiv. https://arxiv.org/abs/2308.03688
  - **Why it matters here:** evaluates agents across multiple interactive environments rather than a single answer score; relevant to task slices, trajectory outcomes, and representative evaluation sets.
- Shunyu Yao, Noah Shinn, Pedram Razavi, and Karthik Narasimhan (2024). *τ-bench: A Benchmark for Tool-Agent-User Interaction in Real-World Domains*. arXiv. https://arxiv.org/abs/2406.12045
  - **Why it matters here:** evaluates policy-following tool interactions against terminal database state and introduces pass^k; relevant to reliability, repeated-run variation, and terminal-outcome measurement.
