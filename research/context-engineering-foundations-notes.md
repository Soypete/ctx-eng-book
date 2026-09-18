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

## ch01.02 — Missing State and Constraints

### Sources queued for review

- **A Survey on Large Language Model based Autonomous Agents** — Lei Wang et al., arXiv v7 (2025). https://arxiv.org/abs/2308.11432
  - Surveys agent construction, memory, planning, tool use, and evaluation; useful for situating workflow state as a system component rather than model memory.
- **Generative Agents: Interactive Simulacra of Human Behavior** — Joon Sung Park et al., arXiv v2 (2023). https://arxiv.org/abs/2304.03442
  - Describes experience records, higher-level reflections, and dynamic retrieval, supporting the distinction between episodic records, semantic summaries, and working context.
- **MemGPT: Towards LLMs as Operating Systems** — Charles Packer et al., arXiv v2 (2024). https://arxiv.org/abs/2310.08560
  - Describes virtual context management across memory tiers and control interrupts for contexts that exceed the active window.

## ch01.03 — Context Failure Case Studies

### Sources queued for review

- **τ-bench: A Benchmark for Tool-Agent-User Interaction in Real-World Domains** — Shunyu Yao et al., arXiv v1 (2024). https://arxiv.org/abs/2406.12045
  - Evaluates tool-agent-user trajectories against final database state and introduces pass^k for repeated reliability.
- **Judging LLM-as-a-Judge with MT-Bench and Chatbot Arena** — Lianmin Zheng et al., arXiv v4 (2023), NeurIPS 2023. https://arxiv.org/abs/2306.05685
  - Measures judge agreement and documents biases relevant to treating evaluation results as evidence rather than intuition.
- **Formalizing and Benchmarking Prompt Injection Attacks and Defenses** — Yupei Liu et al., arXiv v5 (2025), USENIX Security 2024. https://arxiv.org/abs/2310.12815
  - Formalizes prompt-injection vectors and evaluates attacks and defenses across models and tasks, supporting attack provenance and regression testing.

## ch01.04 — Personalization Failures

### Sources queued for review

- **When Large Language Models Meet Personalization: Perspectives of Challenges and Opportunities** — Jin Chen et al., arXiv v1 (2023). https://arxiv.org/abs/2307.16376
  - Reviews personalization challenges and the role of LLMs as interfaces to user-specific services and external tools.
- **Personalization of Large Language Models: A Survey** — Zhehao Zhang et al., arXiv v1 (2024). https://arxiv.org/abs/2411.00027
  - Provides a taxonomy of personalized LLM usage, techniques, evaluation, and open challenges; useful for distinguishing personalized context from model-level adaptation.
- **A Survey of Personalized Large Language Models: Progress and Future Directions** — Jiahong Liu et al., arXiv v1 (2025). https://arxiv.org/abs/2502.11528
  - Surveys prompting, adapters, and alignment approaches, supporting the module's distinction between query-time context and changes to model behavior.
- **Large Language Models as Recommender Systems: A Study of Popularity Bias** — Jan Malte Lichtenberg et al., arXiv v1 (2024). https://arxiv.org/abs/2406.01285
  - Studies a concrete personalization failure mode—popularity bias—and how prompting changes the accuracy/bias trade-off.

## ch02.01 — Context Infrastructure Is the System

### Sources queued for review

- **Harnessing the Power of LLMs in Practice: A Survey on ChatGPT and Beyond** — Jingfeng Yang et al., arXiv v1 (2023). https://arxiv.org/abs/2304.13712
  - Surveys LLM use across tasks and discusses where conventional NLP components remain preferable, supporting the module's model-agnostic caveat.
- **Large Language Models for Information Retrieval: A Survey** — Yutao Zhu et al., arXiv v1 (2023). https://arxiv.org/abs/2308.07107
  - Covers query rewriting, retrieval, reranking, readers, and search agents, supporting the claim that context infrastructure is broader than prompt wording.
- **Retrieval Augmented Generation Evaluation in the Era of Large Language Models: A Comprehensive Survey** — Aoran Gan et al., arXiv v1 (2025). https://arxiv.org/abs/2504.14891
  - Organizes evaluation across system performance, factuality, safety, and computational efficiency, supporting the module's insistence on measuring the complete system.

## ch02.02 — The Production Context Stack

### Sources queued for review

- **AgentBench: Evaluating LLMs as Agents** — Xiao Liu et al., arXiv v1 (2023). https://arxiv.org/abs/2308.03688
  - Evaluates agents across interactive environments and identifies reasoning, decision-making, and instruction-following failures that belong to the system boundary rather than a final-answer-only score.
- **On Evaluating the Integration of Reasoning and Action in LLM Agents** — arXiv v1 (2023). https://arxiv.org/abs/2311.09721
  - Compares no-interaction, sequential, and iterative tool-use strategies, supporting the orchestration layer's role in production behavior.
- **Preble: Efficient Distributed Prompt Scheduling for LLM Serving** — arXiv v1 (2024). https://arxiv.org/abs/2407.00023
  - Measures throughput and tail latency under distributed prompt scheduling, supporting the execution-layer claim that serving choices affect behavior even when prompts are unchanged.

## ch07.01 — Sources of Context

### Sources queued for review

- **Dense Passage Retrieval for Open-Domain Question Answering** — Vladimir Karpukhin et al., EMNLP (2020). https://aclanthology.org/2020.emnlp-main.550/
  - Provides a primary dense-retrieval formulation and evaluation, relevant to distinguishing candidate selection from authority and answer generation.
- **Reciprocal Rank Fusion outperforms Condorcet and Individual Rank Learning Methods** — Gordon V. Cormack et al., SIGIR (2009). https://doi.org/10.1145/1571941.1572114
  - Supports the claim that rank fusion combines retrieval lists but does not itself establish source truth, scope, or freshness.
- **ColBERT: Efficient and Effective Passage Search via Contextualized Late Interaction over BERT** — Omar Khattab and Matei Zaharia, SIGIR (2020). https://arxiv.org/abs/2004.12832
  - Provides a primary late-interaction retrieval design, relevant to the module's distinction among lexical, dense, and reranking selection operations.

## ch07.02 — Context Assembly Pipelines

### Sources queued for review

- **Lost in the Middle: How Language Models Use Long Contexts** — Nelson F. Liu et al., arXiv (2023). https://arxiv.org/abs/2307.03172
  - Measures position-sensitive use of long inputs, relevant to ranking, truncation, and the claim that a context window is not a uniform evidence budget.
- **Never Lost in the Middle: Mastering Long-Context Question Answering with Position-Agnostic Decompositional Training** — Junqing He et al., arXiv (2023). https://arxiv.org/abs/2311.09198
  - Tests a method for reducing positional retrieval failures, relevant to evaluating assembly order rather than assuming model access is uniform.
- **Lost in the Middle, and In-Between: Enhancing Language Models' Ability to Reason Over Long Contexts in Multi-Hop QA** — George Arthur Baker et al., arXiv (2024). https://arxiv.org/abs/2412.10079
  - Extends positional failure analysis to multi-hop evidence, relevant to preserving related evidence during budget allocation and truncation.

## ch07.03 — Freshness, Consistency, and Partial Failure

### Sources queued for review

- **DRAGged into Conflicts: Detecting and Addressing Conflicting Sources in Search-Augmented LLMs** — Arie Cattan et al., arXiv (2025). https://arxiv.org/abs/2506.08500
  - Studies conflicting retrieved sources, relevant to resolving disagreement before serialization and distinguishing attribution from authority.
- **Retrieval-Augmented Generation with Conflicting Evidence** — authors not captured in this pass, arXiv (2025). https://arxiv.org/abs/2504.13079
  - Introduces conflict-oriented evaluation data and methods, relevant to testing whether a pipeline should answer, qualify, or abstain.
- **Resolving Conflicting Evidence in Automated Fact-Checking: A Study on Retrieval-Augmented LLMs** — arXiv (2025). https://arxiv.org/abs/2505.17762
  - Evaluates source-aware handling of conflicting evidence, relevant to authority rules and incorrect-proceed measurements.

## ch07.04 — Hydration Coverage and Retrieval Success

### Sources queued for review

- **RAGAS: Automated Evaluation of Retrieval Augmented Generation** — Shahul Es et al., arXiv (2023). https://arxiv.org/abs/2309.15217
  - Separates retrieval relevance, context use/faithfulness, and generation quality, relevant to keeping hydration coverage diagnostic rather than an all-purpose score.
- **BEIR: A Heterogeneous Benchmark for Zero-shot Evaluation of Information Retrieval Models** — Nandan Thakur et al., arXiv (2021). https://arxiv.org/abs/2104.08663
  - Evaluates retrieval across diverse tasks and model families, relevant to measuring within-slot retrieval success rather than assuming one corpus result generalizes.
- **Evaluation of Retrieval-Augmented Generation: A Survey** — Hao Yu et al., arXiv (2024). https://arxiv.org/abs/2405.07437
  - Organizes retrieval, generation, relevance, accuracy, and faithfulness measures, relevant to the module's limitation boundaries.

## ch02.03 — Engineering the Context Boundaries

### Sources queued for review

- **AutoGen: Enabling Next-Gen LLM Applications via Multi-Agent Conversation** — Qingyun Wu et al., arXiv v1 (2023). https://arxiv.org/abs/2308.08155
  - Describes configurable multi-agent, human, and tool interactions, supporting ownership of capability interfaces and workflow boundaries.
- **MetaGPT: Meta Programming for A Multi-Agent Collaborative Framework** — Sirui Hong et al., arXiv v1 (2023). https://arxiv.org/abs/2308.00352
  - Encodes standardized procedures and role-based collaboration, supporting explicit handoffs, verification, and boundary ownership.
- **LLM Agents for Interactive Workflow Provenance: Reference Architecture and Evaluation Methodology** — Renan Souza et al., arXiv v1 (2025). https://arxiv.org/abs/2509.13978
  - Connects agent architecture to workflow provenance and evaluation, supporting the module's emphasis on inspectable ownership and evidence.

## ch06.03 — User, Session, and Workflow State

### Sources queued for review

- **AgentR: A Stateful and Recovery-Aware Software Architecture for LLM-based Auditable Workflows** — Riya Samanta et al., arXiv (2026). https://arxiv.org/abs/2608.15264
  - Describes durable workflow artifacts, explicit state transitions, retries, and orphan-job handling, relevant to treating workflow state as a recovery contract.
- **AgentRewind: Recoverable Execution for Long-Horizon LLM Agents** — Yu Zhuang et al., arXiv (2026). https://arxiv.org/abs/2608.14380
  - Uses aligned checkpoints of agent context and environment state for recovery, relevant to distinguishing session context from resumable workflow state.
- **Isolated but Exposed: Persistence-Based Memory Extraction Attack on LLM Agents** — Xinyu Gao et al., arXiv (2026). https://arxiv.org/abs/2607.23444
  - Shows that user-level memory isolation does not eliminate tool-side exfiltration risk, relevant to testing state boundaries and downstream capability use.
