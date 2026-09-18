# Context Engineering Foundations — Research Notes

*Reading Context: Context Engineering Book Research — ch00*

## Primary Sources Queued for Review

### Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks

- **Authors:** Patrick Lewis et al.
- **Version read:** arXiv v4, revised 12 April 2021
- **Venue:** NeurIPS 2020
- **Source:** https://arxiv.org/abs/2005.11401
- **Why it matters:** establishes retrieval-augmented generation as a combination of parametric model memory and an explicit non-parametric memory accessed through retrieval.

### Toolformer: Language Models Can Teach Themselves to Use Tools

- **Authors:** Timo Schick et al.
- **Version read:** arXiv v1, submitted 9 February 2023
- **Source:** https://arxiv.org/abs/2302.04761
- **Why it matters:** describes a model learning when to call APIs, what arguments to pass, and how to incorporate returned results, supporting the distinction between prompt-time injection and agent-directed context acquisition.

### ReAct: Synergizing Reasoning and Acting in Language Models

- **Authors:** Shunyu Yao et al.
- **Version read:** arXiv v3, revised 10 March 2023
- **Venue:** ICLR camera-ready version
- **Source:** https://arxiv.org/abs/2210.03629
- **Why it matters:** gives a primary account of interleaving reasoning traces and actions that query external sources, useful for the chapter's agent-directed retrieval distinction.

### Lost in the Middle: How Language Models Use Long Contexts

- **Authors:** Nelson F. Liu et al.
- **Version read:** arXiv v3, revised 20 November 2023
- **Venue:** Transactions of the Association for Computational Linguistics, 2023
- **Source:** https://arxiv.org/abs/2307.03172
- **Why it matters:** evaluates how the position of relevant information in a long input affects performance, supporting the chapter's bounded-working-set and context-selection argument.

## Notes

These sources are queued as `unreviewed` in `book/RESEARCH_LEDGER.md`. No inline manuscript citation should be added until Miriah marks an item reviewed.

## ch01.01 — Missing and Incorrect Information

### Sources queued for review

- **Survey of Hallucination in Natural Language Generation** — Ziwei Ji et al., arXiv v7 (2024), ACM Computing Surveys (2022). https://arxiv.org/abs/2202.03629
  - Surveys hallucination definitions, measurement, and mitigation across NLG tasks.
- **SelfCheckGPT: Zero-Resource Black-Box Hallucination Detection for Generative Large Language Models** — Potsawee Manakul, Adian Liusie, and Mark J. F. Gales, arXiv v3 (2023), EMNLP 2023. https://arxiv.org/abs/2303.08896
  - Studies consistency across sampled outputs as a black-box signal for factuality.
- **Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks** — Patrick Lewis et al., arXiv v4 (2021), NeurIPS 2020. https://arxiv.org/abs/2005.11401
  - Provides a primary account of combining parametric generation with explicit retrieved memory.
- **ReAct: Synergizing Reasoning and Acting in Language Models** — Shunyu Yao et al., arXiv v3 (2023), ICLR camera-ready. https://arxiv.org/abs/2210.03629
  - Describes interleaving actions and observations, supporting the distinction between a requested tool step and verified evidence.
