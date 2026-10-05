# One-Shot, Loop, and Delegation Notes

## ch14.03 — Sources queued for review

- Qingyun Wu et al. (2023). *AutoGen: Enabling Next-Gen LLM Applications via Multi-Agent Conversation*. arXiv. https://arxiv.org/abs/2308.08155
  - **Why it matters here:** describes a framework for agents that converse, use tools, and involve humans; relevant to distinguishing a coordination mechanism from a reliability guarantee.
- Chen Qian et al. (2023). *ChatDev: Communicative Agents for Software Development*. arXiv. https://arxiv.org/abs/2307.07924
  - **Why it matters here:** presents role-based multi-agent workflows with staged communication; relevant to typed handoffs, phase boundaries, and the cost of coordinating workers.
- Junyi Ao et al. (2024). *Mixture-of-Agents Enhances Large Language Model Capabilities*. International Conference on Learning Representations. https://arxiv.org/abs/2406.04692
  - **Why it matters here:** evaluates layered agents whose outputs become context for later agents; relevant to measuring whether added calls improve outcomes enough to justify multiplied context and inference cost.
- Anthropic. *How we built our multi-agent research system*. Anthropic Engineering. https://www.anthropic.com/engineering/multi-agent-research-system
  - **Why it matters here:** reports a production orchestrator-worker pattern with parallel subagents, retries, and checkpoints; relevant as engineering evidence for the benefits and operational controls of delegation.
