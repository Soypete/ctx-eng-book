# Token and Context Economics Notes

## ch14.01 — Sources queued for review

- Nelson F. Liu et al. (2023). *Lost in the Middle: How Language Models Use Long Contexts*. Transactions of the Association for Computational Linguistics. https://arxiv.org/abs/2307.03172
  - **Why it matters here:** measures how the position of relevant information affects long-context use; relevant to the distinction between fitting context and using it effectively.
- In Gim et al. (2024). *Prompt Cache: Modular Attention Reuse for Low-Latency Inference*. arXiv. https://arxiv.org/abs/2311.04934
  - **Why it matters here:** studies reusing attention states for repeated prompt segments; relevant to prefix stability, latency, and the economics of repeatedly sending stable context.
- Woosuk Kwon et al. (2023). *Efficient Memory Management for Large Language Model Serving with PagedAttention*. Proceedings of the 29th ACM Symposium on Operating Systems Principles. https://doi.org/10.1145/3600006.3613165
  - **Why it matters here:** connects sequence length and KV-cache management to serving memory and throughput; relevant to treating context length as an operational cost, not only a provider limit.
- Patrick Lewis et al. (2020). *Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks*. Advances in Neural Information Processing Systems. https://arxiv.org/abs/2005.11401
  - **Why it matters here:** provides the primary retrieval-versus-parametric-memory framing; relevant to pricing retrieval and context assembly as part of the full pipeline rather than assuming the model call is the whole cost.
