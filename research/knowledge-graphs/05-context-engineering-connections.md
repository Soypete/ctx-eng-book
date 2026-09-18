# Context Engineering Connections

This document captures ideas inspired by the book rather than statements made by the authors.

## Knowledge Graphs Are Not Context Engineering

Knowledge graphs are one implementation of semantic context.

Context engineering also includes:

- retrieval
- authorization
- memory/state
- orchestration
- evaluation
- observability

## Memory Is Retrieval

One recurring conclusion is that AI memory is better understood as structured retrieval over governed state.

The challenge is organizing information so that the correct context can be selected when needed.

## Attention Does Not Replace Data Engineering

Attention selects among available tokens.

It cannot organize enterprise information, enforce authorization, or predict future retrieval requirements.

Those remain engineering problems.

## Reliable AI

Reliable AI systems depend on:

- engineered context
- semantic constraints
- governed retrieval
- structured state
- evaluation

Knowledge graphs provide techniques that strengthen each of these areas without becoming the entire architecture.
# ch08.04 — Sources queued for review

- Haoyu Han, Harry Shomer, Yu Wang, Yongjia Lei, Kai Guo, Zhigang Hua, Bo Long, Hui Liu, and Jiliang Tang (2025). *RAG vs. GraphRAG: A Systematic Evaluation and Key Insights*. arXiv:2502.11371. https://arxiv.org/abs/2502.11371
  - **Why it matters here:** compares RAG and GraphRAG across question answering and query-focused summarization, finding task-specific strengths and motivating selection or integration rather than a universal winner.
- *Knowledge graph quality control: A survey* (2021). Future Internet. https://doi.org/10.1016/j.fmre.2021.09.003
  - **Why it matters here:** surveys quality dimensions and their differences, supporting the module's broader cost model beyond answer accuracy.
- Philipp Cimiano and Heiko Paulheim (2016). *Knowledge Graph Refinement: A Survey of Approaches and Evaluation Methods*. Semantic Web. https://doi.org/10.3233/SW-160218
  - **Why it matters here:** frames refinement and evaluation as ongoing knowledge-graph operations, supporting reversible pilots, correction workflows, and rebuild requirements.
