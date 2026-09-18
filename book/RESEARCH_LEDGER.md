## Conventions

- Research notes use a descriptive `#` title, followed by italic metadata lines such as `*Source: <URL>*`, `*Reading Context: ...*`, and `*Reading List Reference: ...*` when applicable.
- Notes are organized with hierarchical Markdown headings, usually beginning with an overview or reaction, followed by named sections, tables, fenced examples, observations, and explicit research questions.
- References are written as Markdown links or bare URLs in bullet lists; primary sources are labeled under headings such as `## Primary Source`, with related readings grouped separately.
- Module files are prose manuscripts with `#`/`##` headings, explanatory paragraphs, tables, and fenced examples. Their `.outline.md` partners are concise heading-and-bullet plans for the same sections; the prose file is the expanded version and should not introduce an unplanned section.
- Repository convention: “Use clear, hierarchical markdown structure”, “Include URLs for references”, and “Group related topics into sections.”
- Conflict recorded: `AGENTS.md` says “Add new research files under `research/`”, while this task explicitly requires the loop ledger at `book/RESEARCH_LEDGER.md`. The task-specific ledger location is used; no additional research file is created during bootstrap.

## Modules

| Chapter | Module | Maps to curriculum | Status | Last pass | Open gaps |
|---------|--------|--------------------|--------|-----------|-----------|
| ch00 | ch00-what-we-mean-by-context-engineering | mapping: unsure — no single numbered module | done | 2026-09-18 | open: promote reviewed sources; confirm whether day1 deck should carry the Lexicon → Semantics → Pragmatics diagram; standalone chapter has no outline |
| ch01 | ch01.01-missing-information | mapping: unsure — 03, 05 | done | 2026-09-18 | open: promote reviewed sources; confirm whether the day1 source/retrieval/generation slide should be linked from the chapter; existing hallucination research marker remains until review |
| ch01 | ch01.02-missing-state | mapping: unsure — 03, 10 | done | 2026-09-18 | open: promote reviewed agent-memory sources; slide alignment deferred; mapping unsure |
| ch01 | ch01.03-context-failure-case-studies | mapping: unsure — 04, 09, 20 | done | 2026-09-18 | open: promote reviewed sources; confirm whether day3/day5 decks should link the case-summary table; scenario examples remain constructed, not incident reports |
| ch01 | ch01.04-personalization-failures | mapping: unsure — 10, 13 | done | 2026-09-18 | open: promote reviewed personalization sources; executable cross-tenant retrieval example/tests remain deferred because implementation belongs in source/app repo; slide alignment deferred; mapping unsure |
| ch02 | ch02.01-ai-marketing | mapping: unsure — 02, 03 | done | 2026-09-18 | open: promote reviewed systems sources; add the Mermaid context-to-outcome flow and curriculum anchors to day1 slides in a slide-specific pass; mapping unsure |
| ch02 | ch02.02-production-ai-stack | 03, 05, 09, 22, 23 | done | 2026-09-18 | open: promote reviewed systems sources; slide alignment deferred to a day1/day2 slide pass; mapping confirmation deferred |
| ch02 | ch02.03-future-ai-engineering | mapping: unsure — 11, 12, 17, 22 | done | 2026-09-18 | open: promote reviewed multi-agent sources; slide alignment deferred to a day1/day3 pass; mapping confirmation deferred |
| ch03 | ch03.01-tokens-embeddings-attention | mapping: unsure — 02 | done | 2026-09-18 | open: promote reviewed embedding/long-context sources; day1 slide alignment deferred; mapping confirmation deferred |
| ch03 | ch03.02-context-windows | mapping: unsure — 02, 05 | done | 2026-09-18 | open: promote reviewed long-context/serving sources; live capacity and pricing tables deferred; slide alignment and mapping confirmation deferred |
| ch03 | ch03.03-compaction-scaffolding-tax | mapping: unsure — 03, 11, 22 | done | 2026-09-18 | open: promote reviewed compression sources; scaffolding-tax attribution remains recollected; slide alignment and mapping confirmation deferred |
| ch04 | ch04.01-in-context-learning | 02 | done | 2026-09-18 | open: promote reviewed ICL and multilingual extraction sources; slide alignment and mapping confirmation deferred |
| ch04 | ch04.02-computational-pragmatics | mapping: unsure — 02, 03 | done | 2026-09-18 | open: promote reviewed clarification sources; slide alignment and mapping confirmation deferred |
| ch04 | ch04.03-examples-instructions-structured-outputs | 02, 03 | done | 2026-09-18 | open: promote reviewed structured-output/tool-use sources; slide alignment and mapping confirmation deferred |
| ch05 | ch05.01-toolformer-and-react | 03, 07, 11 | done | 2026-09-18 | open: promote reviewed tool-use sources; slide alignment and mapping confirmation deferred |
| ch05 | ch05.02-tool-schemas-and-function-calling | 03, 11 | done | 2026-09-18 | open: promote reviewed tool-schema sources; slide alignment and mapping confirmation deferred |
| ch05 | ch05.03-tool-selection-routing-validation | 03, 11, 13 | done | 2026-09-18 | open: promote reviewed routing/safety sources; slide alignment and mapping confirmation deferred |
| ch05 | ch05.04-tool-usage-pattern-detection | mapping: unsure — 03, 09, 22 | done | 2026-09-18 | open: review four queued trajectory/provenance sources; confirm mapping; align day4 trace/eval handoff slides in a slide pass |
| ch06 | ch06.01-the-myth-of-model-memory | 10 | done | 2026-09-18 | open: review four queued memory sources; promote or replace the existing Orogat citation; align any memory/compaction slide material in a later slide pass |
| ch06 | ch06.02-persistent-state-and-retrieval | 10, 11 | done | 2026-09-18 | open: review three queued persistence/provenance sources; promote or replace the existing Orogat citation; align day3 memory-platform slides in a later slide pass |
| ch06 | ch06.03-user-session-workflow-state | 10, 11, 23 | pending | — | — |
| ch07 | ch07.01-sources-of-context | 05, 06, 10, 11 | pending | — | — |
| ch07 | ch07.02-context-assembly-pipelines | 05, 06, 07, 11 | pending | — | — |
| ch07 | ch07.03-freshness-consistency-and-partial-failure | mapping: unsure — 05, 09, 22 | pending | — | — |
| ch07 | ch07.04-hydration-coverage-and-retrieval-success | 05, 06, 08, 21 | pending | — | — |
| ch07 | ch07.05-public-data-sources-wikipedia-web | mapping: unsure — 05, 17 | pending | — | — |
| ch07 | ch07.06-information-extraction-pipelines | 05, 08, 16 | pending | — | — |
| ch08 | ch08-knowledge-graphs-and-semantic-context | mapping: unsure — 16 | pending | — | — |
| ch08 | ch08.01-schemas-taxonomies-and-ontologies | 16 | pending | — | — |
| ch08 | ch08.02-rdf-owl-and-sparql | mapping: unsure — 16 | pending | — | — |
| ch08 | ch08.03-entity-resolution-and-relationship-traversal | 16 | pending | — | — |
| ch08 | ch08.04-knowledge-graph-tradeoffs | mapping: unsure — 05, 16 | pending | — | — |
| ch08 | ch08.05-instance-coverage-and-ontology-population | 16, 21 | pending | — | — |
| ch08 | ch08.06-property-completeness-and-schema-quality | 16, 21 | pending | — | — |
| ch08 | ch08.07-ontology-guided-information-extraction | 16 | pending | — | — |
| ch08 | ch08.08-knowledge-extraction-methods | 08, 16 | pending | — | — |
| ch08 | ch08.09-guardrails-for-extraction-validation | 13, 18, 21 | pending | — | — |
| ch08 | ch08.10-multilingual-extraction-with-llms | mapping: unsure — 05, 16 | pending | — | — |
| ch09 | ch09.01-lexical-and-relational-retrieval | 05, 06 | pending | — | — |
| ch09 | ch09.02-vector-and-semantic-retrieval | 05, 06 | pending | — | — |
| ch09 | ch09.03-graph-and-hybrid-retrieval | 06, 16 | pending | — | — |
| ch09 | ch09.04-ranking-reranking-and-query-planning | 06, 07 | pending | — | — |
| ch09 | ch09.05-context-precision-and-context-recall | 05, 06, 08, 21 | pending | — | — |
| ch10 | ch10.00-guardrails-and-ontology-based-validation | 13, 16, 18 | pending | — | — |
| ch10 | ch10.01-personalization-as-retrieval | 10, 11 | pending | — | — |
| ch10 | ch10.02-scoped-hydration | 10, 13 | pending | — | — |
| ch10 | ch10.03-provenance-and-derived-context | 10, 22 | pending | — | — |
| ch10 | ch10.04-policy-aware-user-context | 13, 19, 20 | pending | — | — |
| ch10 | ch10.05-provenance-coverage-metrics | 09, 21, 22 | pending | — | — |
| ch11 | ch11.01-least-privilege | 13, 19, 20 | pending | — | — |
| ch11 | ch11.02-rbac-abac-capability-based-access | mapping: unsure — 13, 19 | pending | — | — |
| ch11 | ch11.03-scoped-credentials-knowledge-stores | 11, 13, 20 | pending | — | — |
| ch11 | ch11.04-retrieval-execution-boundaries | 07, 11, 13 | pending | — | — |
| ch11 | ch11.05-authorization-coverage-and-necessary-access | 13, 19, 20 | pending | — | — |
| ch12 | ch12.01-small-composable-systems | 01, 11 | pending | — | — |
| ch12 | ch12.02-pipes-files-explicit-interfaces | 01, 11 | pending | — | — |
| ch12 | ch12.03-mounts-namespaces-isolation | mapping: unsure — 01, 13, 20 | pending | — | — |
| ch12 | ch12.04-task-workspaces-secret-management | 01, 11, 13 | pending | — | — |
| ch13 | ch13.01-planning-and-react | 03, 11 | pending | — | — |
| ch13 | ch13.02-harnesses-and-state-machines | 11, 23 | pending | — | — |
| ch13 | ch13.03-durable-and-event-driven-execution | mapping: unsure — 11, 23 | pending | — | — |
| ch13 | ch13.04-loops-retries-and-bounded-autonomy | 03, 11, 13, 23 | pending | — | — |
| ch14 | ch14.01-token-and-context-economics | 02, 05 | pending | — | — |
| ch14 | ch14.02-retrieval-tool-and-latency-costs | 05, 06, 07 | pending | — | — |
| ch14 | ch14.03-one-shot-execution-loops-and-subagents | 03, 11, 12 | pending | — | — |
| ch14 | ch14.04-local-models-and-model-routing | mapping: unsure — 03, 15 | pending | — | — |
| ch14 | ch14.05-context-efficiency-metrics | 08, 09, 21, 22 | pending | — | — |
| ch14 | ch14.06-ner-vs-llm-extraction-costs | 08, 16 | pending | — | — |
| ch14 | ch14.07-extraction-method-selection | 08, 16 | pending | — | — |
| ch14 | ch14.08-cost-aware-extraction-pipeline-design | 08, 16, 21 | pending | — | — |
| ch15 | ch15.01-diagnosing-model-problems | 04, 09, 21 | pending | — | — |
| ch15 | ch15.02-fine-tuning-and-lora | mapping: unsure — 02, 15 | pending | — | — |
| ch15 | ch15.03-distillation-and-specialized-models | mapping: unsure — 02, 15 | pending | — | — |
| ch15 | ch15.04-context-engineering-as-the-research-phase | 04, 09, 21 | pending | — | — |
| ch16 | ch16.01-tracing-context-assembly | 22 | pending | — | — |
| ch16 | ch16.02-prompt-retrieval-tool-lineage | 22 | pending | — | — |
| ch16 | ch16.03-state-cost-latency-observability | 22, 23 | pending | — | — |
| ch17 | ch17.01-evals-and-benchmarks | 04, 09, 21 | pending | — | — |
| ch17 | ch17.02-retrieval-and-tool-evaluation | 08, 09, 21 | pending | — | — |
| ch17 | ch17.03-regression-and-scenario-testing | 09, 20, 21, 23 | pending | — | — |
| ch17 | ch17.04-reliability-metrics-and-failure-budgets | 09, 21, 22, 23 | pending | — | — |
| ch17 | ch17.05-qa-driven-srl-benchmarks | mapping: unsure — 08, 21 | pending | — | — |
| ch17 | ch17.06-openie-evaluation-relvis | mapping: unsure — 08, 16, 21 | pending | — | — |
| ch18 | ch18.01-source-and-ingestion-architecture | 05, 06, 16 | pending | — | — |
| ch18 | ch18.02-semantic-and-retrieval-infrastructure | 05, 06, 16 | pending | — | — |
| ch18 | ch18.03-authorization-state-and-tooling | 10, 11, 13, 19 | pending | — | — |
| ch18 | ch18.04-observability-evaluation-cost-control | 09, 21, 22, 23 | pending | — | — |

## Research queue

### ch00-what-we-mean-by-context-engineering

- [ ] **unreviewed** — Patrick Lewis et al. (2021). *Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks*. NeurIPS 2020. <https://arxiv.org/abs/2005.11401>
  - **Why it matters here:** establishes retrieval-augmented generation as a way to provide explicit non-parametric memory to a language process; relevant to “Three ways context reaches a language process”.
  - **Claim it would support:** “The same discipline appears under different retrieval postures” and the description of prompt-time retrieval and assembly.
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Timo Schick et al. (2023). *Toolformer: Language Models Can Teach Themselves to Use Tools*. arXiv v1. <https://arxiv.org/abs/2302.04761>
  - **Why it matters here:** describes selecting APIs, arguments, and returned results during generation; relevant to the agent-directed retrieval distinction.
  - **Claim it would support:** “In agent-directed retrieval, the model receives a search or fetch capability and decides what to request during a loop.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shunyu Yao et al. (2023). *ReAct: Synergizing Reasoning and Acting in Language Models*. ICLR camera-ready. <https://arxiv.org/abs/2210.03629>
  - **Why it matters here:** provides a primary account of interleaving reasoning and external actions to gather information; relevant to agent-directed retrieval and its added turns.
  - **Claim it would support:** “Each result becomes a new observation that may lead to another request.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Nelson F. Liu et al. (2023). *Lost in the Middle: How Language Models Use Long Contexts*. Transactions of the Association for Computational Linguistics. <https://arxiv.org/abs/2307.03172>
  - **Why it matters here:** evaluates degradation when relevant information moves within long inputs; relevant to the bounded working set and selection/budget argument.
  - **Claim it would support:** “More accessible data can add cost and distract from the evidence that matters.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch01.01-missing-information

- [ ] **unreviewed** — Ziwei Ji et al. (2024). *Survey of Hallucination in Natural Language Generation*. ACM Computing Surveys. <https://arxiv.org/abs/2202.03629>
  - **Why it matters here:** supplies a research taxonomy and measurement/mitigation overview for the chapter's operational discussion of hallucination.
  - **Claim it would support:** “A hallucination is ... an output that asserts unsupported or contradictory information as though it were grounded.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Potsawee Manakul, Adian Liusie, and Mark J. F. Gales (2023). *SelfCheckGPT: Zero-Resource Black-Box Hallucination Detection for Generative Large Language Models*. EMNLP 2023. <https://arxiv.org/abs/2303.08896>
  - **Why it matters here:** tests a black-box signal for distinguishing fluent output from inconsistent factual claims.
  - **Claim it would support:** “The model still produces a continuation when the supplied evidence is insufficient ... Better context can reduce unsupported answers. It cannot guarantee truth.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Patrick Lewis et al. (2021). *Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks*. NeurIPS 2020. <https://arxiv.org/abs/2005.11401>
  - **Why it matters here:** provides primary evidence for retrieval-grounded generation and the distinction between model memory and explicit retrieved evidence.
  - **Claim it would support:** “A missing source needs ingestion ... A retrieval miss needs a better query, index, ranking policy, or scope.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shunyu Yao et al. (2023). *ReAct: Synergizing Reasoning and Acting in Language Models*. ICLR camera-ready. <https://arxiv.org/abs/2210.03629>
  - **Why it matters here:** gives a primary account of action/observation loops where tool calls produce evidence for subsequent steps.
  - **Claim it would support:** “A tool name, description, and input schema tell the model what capability it may request; they do not prove ... that an action succeeded.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch01.02-missing-state

- [ ] **unreviewed** — Lei Wang et al. (2025). *A Survey on Large Language Model based Autonomous Agents*. arXiv v7. <https://arxiv.org/abs/2308.11432>
  - **Why it matters here:** surveys agent construction, memory, planning, tool use, and evaluation; relevant to treating workflow state as a system component.
  - **Claim it would support:** “Durable state is therefore not ‘the model remembering’; it is a harness deciding what to record, retain, retrieve, and place back into the working context.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Joon Sung Park et al. (2023). *Generative Agents: Interactive Simulacra of Human Behavior*. arXiv v2. <https://arxiv.org/abs/2304.03442>
  - **Why it matters here:** describes experience records, reflections, and dynamic retrieval; relevant to the episodic/semantic/working-state distinction.
  - **Claim it would support:** “Episodic state records what happened ... Semantic state stores current facts ... Working state is the bounded context assembled for this turn.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Charles Packer et al. (2024). *MemGPT: Towards LLMs as Operating Systems*. arXiv v2. <https://arxiv.org/abs/2310.08560>
  - **Why it matters here:** describes virtual context management across memory tiers and control interrupts for active windows with limited capacity.
  - **Claim it would support:** “The root set ... must not be trimmed merely because a context budget is tight.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch01.03-context-failure-case-studies

- [ ] **unreviewed** — Shunyu Yao et al. (2024). *τ-bench: A Benchmark for Tool-Agent-User Interaction in Real-World Domains*. arXiv v1. <https://arxiv.org/abs/2406.12045>
  - **Why it matters here:** evaluates complete tool-agent-user trajectories against final state and repeated-run reliability.
  - **Claim it would support:** “Run the same task repeatedly, record the complete trajectory rather than only the final answer.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Lianmin Zheng et al. (2023). *Judging LLM-as-a-Judge with MT-Bench and Chatbot Arena*. NeurIPS 2023. <https://arxiv.org/abs/2306.05685>
  - **Why it matters here:** documents judge biases and agreement limits relevant to interpreting case-study evaluation.
  - **Claim it would support:** “A scenario is a debugging hypothesis until its boundary, expected outcome, and telemetry are measured.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yupei Liu et al. (2025). *Formalizing and Benchmarking Prompt Injection Attacks and Defenses*. USENIX Security 2024. <https://arxiv.org/abs/2310.12815>
  - **Why it matters here:** formalizes attack vectors and evaluates defenses across models and tasks, supporting provenance-aware regression testing.
  - **Claim it would support:** “A successful attack should become a regression case, with its vector recorded as user input, retrieved content, or a tool description.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch01.04-personalization-failures

- [ ] **unreviewed** — Jin Chen et al. (2023). *When Large Language Models Meet Personalization: Perspectives of Challenges and Opportunities*. arXiv. <https://arxiv.org/abs/2307.16376>
  - **Why it matters here:** grounds the distinction between personalization as user-specific service behavior and generic model capability in the opening definition and missing-preferences section.
  - **Claim it would support:** “Personalization is the process of binding authenticated identity + relevant preferences + permissions + task constraints to context at query time.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Zhehao Zhang et al. (2024). *Personalization of Large Language Models: A Survey*. arXiv. <https://arxiv.org/abs/2411.00027>
  - **Why it matters here:** supplies a taxonomy for separating personalized context assembly from model-level personalization techniques.
  - **Claim it would support:** “Persistent preferences may live in a state store; they do not become reliable merely because a framework calls that store ‘memory.’”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Jiahong Liu et al. (2025). *A Survey of Personalized Large Language Models: Progress and Future Directions*. arXiv. <https://arxiv.org/abs/2502.11528>
  - **Why it matters here:** compares prompting, adapter, and alignment approaches, supporting the module's model-parameters versus query-time-context distinction.
  - **Claim it would support:** “Treat personalization as context assembly at query time, not as model memory.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Jan Malte Lichtenberg et al. (2024). *Large Language Models as Recommender Systems: A Study of Popularity Bias*. arXiv. <https://arxiv.org/abs/2406.01285>
  - **Why it matters here:** gives a concrete evaluation failure mode for personalized recommendations and a measurable accuracy/bias trade-off.
  - **Claim it would support:** “Measure whether personalization works with task relevance, preference freshness, provenance coverage, and user-correction rates.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch02.01-ai-marketing

- [ ] **unreviewed** — Jingfeng Yang et al. (2023). *Harnessing the Power of LLMs in Practice: A Survey on ChatGPT and Beyond*. arXiv. <https://arxiv.org/abs/2304.13712>
  - **Why it matters here:** supports the model-quality caveat and the choice to keep classical NLP components in the system map where they fit the task.
  - **Claim it would support:** “Production behavior cannot be delegated to the model alone.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yutao Zhu et al. (2023). *Large Language Models for Information Retrieval: A Survey*. arXiv. <https://arxiv.org/abs/2308.07107>
  - **Why it matters here:** covers the retrieval and search components that sit between source data and model input.
  - **Claim it would support:** “Retrieval quality, durable state, authorization boundaries, operational feedback, and evaluation encode deeper knowledge about how the system works.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Aoran Gan et al. (2025). *Retrieval Augmented Generation Evaluation in the Era of Large Language Models: A Comprehensive Survey*. arXiv. <https://arxiv.org/abs/2504.14891>
  - **Why it matters here:** provides an evidence path for evaluating the retrieval-generation system across effectiveness, factuality, safety, and efficiency.
  - **Claim it would support:** “The practical method is to assign each failure to a component boundary, define the contract at that boundary, and measure whether the complete request succeeds.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch02.02-production-ai-stack

- [ ] **unreviewed** — Xiao Liu et al. (2023). *AgentBench: Evaluating LLMs as Agents*. arXiv. <https://arxiv.org/abs/2308.03688>
  - **Why it matters here:** evaluates interactive agent environments and failure modes across orchestration and evaluation boundaries.
  - **Claim it would support:** “The practical method is to assign each failure to a component boundary ... and measure whether the complete request succeeds.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — *On Evaluating the Integration of Reasoning and Action in LLM Agents* (2023). arXiv. <https://arxiv.org/abs/2311.09721>
  - **Why it matters here:** compares interaction strategies, making the orchestration layer's effect on outcomes measurable.
  - **Claim it would support:** “In an agent it manages loops; in a classifier or extraction pipeline it may simply coordinate retrieval, inference, validation, and routing.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — *Preble: Efficient Distributed Prompt Scheduling for LLM Serving* (2024). arXiv. <https://arxiv.org/abs/2407.00023>
  - **Why it matters here:** supplies serving-level latency and throughput evidence for the execution layer.
  - **Claim it would support:** “Serving choices can affect observed behavior even when the application prompt is unchanged.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch02.03-future-ai-engineering

- [ ] **unreviewed** — Qingyun Wu et al. (2023). *AutoGen: Enabling Next-Gen LLM Applications via Multi-Agent Conversation*. arXiv. <https://arxiv.org/abs/2308.08155>
  - **Why it matters here:** describes configurable multi-agent, human, and tool interactions, supporting ownership of capability interfaces and workflow boundaries.
  - **Claim it would support:** “The durable unit is the boundary and its owner, not a proposed title.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Sirui Hong et al. (2023). *MetaGPT: Meta Programming for A Multi-Agent Collaborative Framework*. arXiv. <https://arxiv.org/abs/2308.00352>
  - **Why it matters here:** encodes standardized procedures and role-based collaboration, supporting explicit handoffs, verification, and boundary ownership.
  - **Claim it would support:** “As scale, risk, and workload diversity grow, ownership should follow failure boundaries.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Renan Souza et al. (2025). *LLM Agents for Interactive Workflow Provenance: Reference Architecture and Evaluation Methodology*. arXiv. <https://arxiv.org/abs/2509.13978>
  - **Why it matters here:** connects agent architecture to workflow provenance and evaluation, supporting inspectable ownership and evidence.
  - **Claim it would support:** “A boundary without an owner becomes a boundary nobody tests.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch03.01-tokens-embeddings-attention

- [ ] **unreviewed** — Nils Reimers and Iryna Gurevych (2019). *Sentence-BERT: Sentence Embeddings using Siamese BERT-Networks*. arXiv. <https://arxiv.org/abs/1908.10084>
  - **Why it matters here:** grounds the distinction between retrieval-oriented sentence embeddings and token-level representations.
  - **Claim it would support:** “A retrieval embedding compresses a passage into a vector optimized for a retrieval objective; it is not simply the model's token embedding copied into a vector index.”
  - **Notes file:** [attention-is-all-you-need-notes.md](../research/attention-is-all-you-need-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Nelson F. Liu et al. (2023). *Lost in the Middle: How Language Models Use Long Contexts*. TACL. <https://arxiv.org/abs/2307.03172>
  - **Why it matters here:** provides primary evidence for position-sensitive utilization in long contexts.
  - **Claim it would support:** “Long or distracting context can degrade performance empirically; the effect depends on model, task, placement, and content.”
  - **Notes file:** [attention-is-all-you-need-notes.md](../research/attention-is-all-you-need-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Nina Poerner, Ulli Waltinger, and Hinrich Schütze (2019). *Sentence Meta-Embeddings for Unsupervised Semantic Textual Similarity*. arXiv. <https://arxiv.org/abs/1911.03700>
  - **Why it matters here:** demonstrates that embedding geometry and similarity depend on representation choices and objectives.
  - **Claim it would support:** “Similarity is a learned proxy for relevance, not evidence that two records are identical, authorized, current, or suitable for the same action.”
  - **Notes file:** [attention-is-all-you-need-notes.md](../research/attention-is-all-you-need-notes.md)
  - **Miriah's notes:**

### ch03.02-context-windows

- [ ] **unreviewed** — Yucheng Li et al. (2024). *SCBench: A KV Cache-Centric Analysis of Long-Context Methods*. arXiv. <https://arxiv.org/abs/2412.10319>
  - **Why it matters here:** evaluates KV-cache lifecycle and long-context methods across memory and computation costs.
  - **Claim it would support:** “KV cache memory grows with context ... throughput drops as context grows.”
  - **Notes file:** [attention-is-all-you-need-notes.md](../research/attention-is-all-you-need-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Hanshi Sun et al. (2024). *ShadowKV: KV Cache in Shadows for High-Throughput Long-Context LLM Inference*. arXiv. <https://arxiv.org/abs/2410.21465>
  - **Why it matters here:** measures serving trade-offs when long-context KV state is compressed or offloaded.
  - **Claim it would support:** “The practical question isn't just ‘what fits in the window’ but ‘what's the right signal-to-cost ratio?’”
  - **Notes file:** [attention-is-all-you-need-notes.md](../research/attention-is-all-you-need-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — *In-Context Learning with Long-Context Models* (2024). arXiv. <https://arxiv.org/abs/2405.00200>
  - **Why it matters here:** evaluates long-context example counts and ordering, supporting a measured rather than universal placement policy.
  - **Claim it would support:** “Do not infer a universal ordering rule from one model family.”
  - **Notes file:** [attention-is-all-you-need-notes.md](../research/attention-is-all-you-need-notes.md)
  - **Miriah's notes:**

### ch03.03-compaction-scaffolding-tax

- [ ] **unreviewed** — Huiqiang Jiang et al. (2023). *LLMLingua: Compressing Prompts for Accelerated Inference of Large Language Models*. arXiv. <https://arxiv.org/abs/2310.05736>
  - **Why it matters here:** evaluates prompt compression and semantic-integrity controls as an explicit transformation.
  - **Claim it would support:** “A smaller prompt is not automatically an improvement.”
  - **Notes file:** [attention-is-all-you-need-notes.md](../research/attention-is-all-you-need-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Huiqiang Jiang et al. (2023). *LongLLMLingua: Accelerating and Enhancing LLMs in Long Context Scenarios via Prompt Compression*. arXiv. <https://arxiv.org/abs/2310.06839>
  - **Why it matters here:** measures quality, cost, latency, and position effects under long-context compression.
  - **Claim it would support:** “Keep a layer when controlled evaluation shows that it improves reliability enough to justify its operational cost.”
  - **Notes file:** [attention-is-all-you-need-notes.md](../research/attention-is-all-you-need-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Zhuoshi Pan et al. (2024). *LLMLingua-2: Data Distillation for Efficient and Faithful Task-Agnostic Prompt Compression*. arXiv. <https://arxiv.org/abs/2403.12968>
  - **Why it matters here:** evaluates faithfulness and latency for a learned compression transformation, supporting explicit information-loss tests.
  - **Claim it would support:** “A context manifest should preserve ... the items deferred for rehydration.”
  - **Notes file:** [attention-is-all-you-need-notes.md](../research/attention-is-all-you-need-notes.md)
  - **Miriah's notes:**

### ch04.01-in-context-learning

- [ ] **unreviewed** — Hao Peng et al. (2023). *When does In-context Learning Fall Short and Why? A Study on Specification-Heavy Tasks*. arXiv. <https://arxiv.org/abs/2311.08993>
  - **Why it matters here:** provides a primary failure boundary for complex task specifications and schema-heavy in-context learning.
  - **Claim it would support:** “Hidden, unconstrained inference makes failures harder to reproduce.”
  - **Notes file:** [llms-in-production/notes.md](../research/llms-in-production/notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Hanjun Luo et al. (2024). *GEIC: Universal and Multilingual Named Entity Recognition with Large Language Models*. arXiv. <https://arxiv.org/abs/2409.11022>
  - **Why it matters here:** evaluates few-shot and zero-shot multilingual, fine-grained entity extraction.
  - **Claim it would support:** “Quality must be evaluated per language and domain.”
  - **Notes file:** [llms-in-production/notes.md](../research/llms-in-production/notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Marco Naguib, Xavier Tannier, and Aurélie Névéol (2024). *Few-shot clinical entity recognition in English, French and Spanish: masked language models outperform generative model prompting*. arXiv. <https://arxiv.org/abs/2402.12801>
  - **Why it matters here:** gives a counterexample to prompt-only extraction in a specialized multilingual domain.
  - **Claim it would support:** “Prompted extraction is a prototyping advantage, not a production guarantee.”
  - **Notes file:** [llms-in-production/notes.md](../research/llms-in-production/notes.md)
  - **Miriah's notes:**

### ch04.02-computational-pragmatics

- [ ] **unreviewed** — Michael J. Q. Zhang and Eunsol Choi (2023). *Clarify When Necessary: Resolving Ambiguity Through Interaction with LMs*. arXiv. <https://arxiv.org/abs/2311.09469>
  - **Why it matters here:** evaluates when a system should ask a clarifying question and how to choose it.
  - **Claim it would support:** “The product must decide when to clarify.”
  - **Notes file:** [computational-pragmatics-notes.md](../research/computational-pragmatics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Lorenz Kuhn, Yarin Gal, and Sebastian Farquhar (2022). *CLAM: Selective Clarification for Ambiguous Questions with Generative Language Models*. arXiv. <https://arxiv.org/abs/2212.07769>
  - **Why it matters here:** provides a primary ambiguity-and-clarification evaluation across dialogue tasks.
  - **Claim it would support:** “A safer harness asks or retrieves evidence before granting write authority.”
  - **Notes file:** [computational-pragmatics-notes.md](../research/computational-pragmatics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Kaustubh D. Dhole (2020). *Resolving Intent Ambiguities by Retrieving Discriminative Clarifying Questions*. arXiv. <https://arxiv.org/abs/2008.07559>
  - **Why it matters here:** grounds clarification as discriminating among competing intents rather than merely asking for more text.
  - **Claim it would support:** “Resolve the intended action, referent, and scope before retrieval or execution.”
  - **Notes file:** [computational-pragmatics-notes.md](../research/computational-pragmatics-notes.md)
  - **Miriah's notes:**

### ch04.03-examples-instructions-structured-outputs

- [ ] **unreviewed** — Saibo Geng et al. (2023). *Grammar-Constrained Decoding for Structured NLP Tasks without Finetuning*. arXiv. <https://arxiv.org/abs/2305.13971>
  - **Why it matters here:** grounds the distinction between guaranteed structural membership and semantic correctness.
  - **Claim it would support:** “Grammar-constrained decoding can guarantee syntactic membership, but neither syntax nor schema guarantees semantic correctness or safe values.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Saibo Geng et al. (2025). *Generating Structured Outputs from Language Models: Benchmark and Studies*. arXiv. <https://arxiv.org/abs/2501.10868>
  - **Why it matters here:** evaluates structured-output compliance, coverage, efficiency, and quality across real JSON schemas.
  - **Claim it would support:** “A schema-valid call is the beginning of execution validation, not the end.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — *On Evaluating the Integration of Reasoning and Action in LLM Agents* (2023). arXiv. <https://arxiv.org/abs/2311.09721>
  - **Why it matters here:** evaluates tool-interaction strategies and host-side execution boundaries.
  - **Claim it would support:** “The model proposes; the host validates, authorizes, and executes.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**

### ch05.01-toolformer-and-react

- [ ] **unreviewed** — Timo Schick et al. (2023). *Toolformer: Language Models Can Teach Themselves to Use Tools*. arXiv. <https://arxiv.org/abs/2302.04761>
  - **Why it matters here:** grounds learned tool selection and argument generation while leaving execution outside the model.
  - **Claim it would support:** “The important shift was moving part of tool selection into the model while leaving execution outside it.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shunyu Yao et al. (2023). *ReAct: Synergizing Reasoning and Acting in Language Models*. ICLR. <https://arxiv.org/abs/2210.03629>
  - **Why it matters here:** grounds the action/observation loop used by agentic retrieval and tool workflows.
  - **Claim it would support:** “Tool outputs are structured observations that inform subsequent actions.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yujia Qin et al. (2023). *ToolLLM: Facilitating Large Language Models to Master Tool Usage*. arXiv. <https://arxiv.org/abs/2307.16789>
  - **Why it matters here:** provides a tool-use dataset and benchmark for API selection and invocation at scale.
  - **Claim it would support:** “ToolBench's scale is evidence about that benchmark, not proof that a production tool catalog should be large.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**

### ch05.02-tool-schemas-and-function-calling

- [ ] **unreviewed** — Saibo Geng et al. (2023). *Grammar-Constrained Decoding for Structured NLP Tasks without Finetuning*. arXiv. <https://arxiv.org/abs/2305.13971>
  - **Why it matters here:** supports the distinction between syntax guarantees and semantic/authorization checks.
  - **Claim it would support:** “The schema constrains representation; it does not grant permission.”
  - **Notes file:** [semantic-contracts.md](../research/semantic-contracts.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Saibo Geng et al. (2025). *Generating Structured Outputs from Language Models: Benchmark and Studies*. arXiv. <https://arxiv.org/abs/2501.10868>
  - **Why it matters here:** benchmarks structured-output compliance and quality across JSON schemas.
  - **Claim it would support:** “A too-strict schema prevents valid use cases; test the exact schema behavior your model and SDK enforce.”
  - **Notes file:** [semantic-contracts.md](../research/semantic-contracts.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yujia Qin et al. (2023). *ToolLLM: Facilitating Large Language Models to Master Tool Usage*. arXiv. <https://arxiv.org/abs/2307.16789>
  - **Why it matters here:** provides a tool-use benchmark for API selection and invocation at scale.
  - **Claim it would support:** “Evaluate selection on ambiguous, adversarial, and no-tool cases.”
  - **Notes file:** [semantic-contracts.md](../research/semantic-contracts.md)
  - **Miriah's notes:**

### ch05.03-tool-selection-routing-validation

- [ ] **unreviewed** — Hongfei Xia et al. (2025). *SafeToolBench: Pioneering a Prospective Benchmark to Evaluating Tool Utilization Safety in LLMs*. arXiv. <https://arxiv.org/abs/2509.07315>
  - **Why it matters here:** evaluates safety before irreversible tool execution rather than after the side effect.
  - **Claim it would support:** “The host decides what it may do for this principal, task, and workflow state.”
  - **Notes file:** [semantic-contracts.md](../research/semantic-contracts.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Wenrui Liu et al. (2025). *MCPAgentBench: A Real-world Task Benchmark for Evaluating LLM Agent MCP Tool Use*. arXiv. <https://arxiv.org/abs/2512.24565>
  - **Why it matters here:** measures tool selection with distractors, task completion, and execution efficiency across multi-step tasks.
  - **Claim it would support:** “Multi-step workflows require explicit routing and completion criteria.”
  - **Notes file:** [semantic-contracts.md](../research/semantic-contracts.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Harsh Soni (2026). *ToolFailBench: Diagnosing Tool-Use Failures in LLM Agents*. arXiv. <https://arxiv.org/abs/2607.04686>
  - **Why it matters here:** distinguishes tool-selection, result-use, fabrication, and unnecessary-call failures.
  - **Claim it would support:** “Reliable systems make unsafe proposals rejectable, repeated side effects idempotent, failures observable, and recovery deterministic.”
  - **Notes file:** [semantic-contracts.md](../research/semantic-contracts.md)
  - **Miriah's notes:**

### ch05.04-tool-usage-pattern-detection

- [ ] **unreviewed** — Yibing Liu et al. (2026). *TrajAD: Trajectory Anomaly Detection for Trustworthy LLM Agents*. arXiv. <https://arxiv.org/abs/2602.06443>
  - **Why it matters here:** supports runtime inspection and localization of anomalous intermediate steps, relevant to “Model the Sequence, Not Just the Count.”
  - **Claim it would support:** “A trace can explain what happened and help locate an anomaly; it cannot by itself explain why the user acted.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**

### ch06.01-the-myth-of-model-memory

- [ ] **unreviewed** — Charles Packer et al. (2023). *MemGPT: Towards LLMs as Operating Systems*. arXiv. <https://arxiv.org/abs/2310.08560>
  - **Why it matters here:** describes virtual context management across memory tiers and interrupts, relevant to distinguishing bounded working context from durable state.
  - **Claim it would support:** “Invocation context ... is a temporary working set, not durable application state.”
  - **Notes file:** [letta-notes.md](../research/letta-notes.md)
  - **Miriah's notes:**

### ch06.02-persistent-state-and-retrieval

- [ ] **unreviewed** — Ming Wu and Pengyuan Zhu (2026). *Agent Zero Memory: Provenance-Aware Long-Term Memory for LLM Agents*. arXiv. <https://arxiv.org/abs/2608.29606>
  - **Why it matters here:** presents parallel episodic, graph, and documentary memory with provenance-locked retrieval, relevant to distinguishing source authority and derived copies.
  - **Claim it would support:** “A state platform should be able to say which source version produced a derived result.”
  - **Notes file:** [episodic-periodic-memory.md](../research/episodic-periodic-memory.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yuanyi Song et al. (2026). *Retrieval-Driven Memory Reconsolidation for Long-Term LLM Agents*. arXiv. <https://arxiv.org/abs/2609.16053>
  - **Why it matters here:** treats retrieval feedback as part of a continuing memory lifecycle, relevant to correction and reorganization.
  - **Claim it would support:** “The dotted maintenance path matters as much as the happy read path.”
  - **Notes file:** [episodic-periodic-memory.md](../research/episodic-periodic-memory.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Zicheng Zhao et al. (2026). *Accurate and Efficient Long-Term Memory for LLM Agents*. arXiv. <https://arxiv.org/abs/2607.16211>
  - **Why it matters here:** studies structured storage, conflict detection, updates, and deletions, relevant to write-path and consistency-gap controls.
  - **Claim it would support:** “The write path should distinguish source events from derived state.”
  - **Notes file:** [episodic-periodic-memory.md](../research/episodic-periodic-memory.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Genglin Liu and Saadia Gabriel (2026). *PM-Bench: Evaluating Prospective Memory in LLM Agents*. arXiv. <https://arxiv.org/abs/2607.12385>
  - **Why it matters here:** provides a controlled test of delayed intentions and future cues, relevant to treating apparent memory as a measurable state-and-retrieval capability.
  - **Claim it would support:** “When an agent appears to forget, retrieval is only one possible fault domain.”
  - **Notes file:** [letta-notes.md](../research/letta-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Hanqi Jiang et al. (2026). *SYNAPSE: Empowering LLM Agents with Episodic-Semantic Memory via Spreading Activation*. arXiv. <https://arxiv.org/abs/2601.02744>
  - **Why it matters here:** evaluates episodic and semantic memory for temporal and multi-hop recall, relevant to separating source evidence from derived meaning.
  - **Claim it would support:** “Raw interaction history is evidence, not automatically knowledge.”
  - **Notes file:** [letta-notes.md](../research/letta-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Kaixiang Wang et al. (2026). *E-mem: Multi-agent based Episodic Context Reconstruction for LLM Agent Memory*. arXiv. <https://arxiv.org/abs/2601.21714>
  - **Why it matters here:** examines episodic context reconstruction and the cost of memory preprocessing, relevant to the warning that summaries are lossy derived state.
  - **Claim it would support:** “Summarization and extraction are lossy transformations.”
  - **Notes file:** [letta-notes.md](../research/letta-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Wonjoong Kim et al. (2025). *Beyond the Final Answer: Evaluating the Reasoning Trajectories of Tool-Augmented Agents*. arXiv. <https://arxiv.org/abs/2510.02837>
  - **Why it matters here:** provides a trajectory-level evaluation frame for efficiency, hallucination, and adaptivity, relevant to the proposed change loop.
  - **Claim it would support:** “Evaluate candidate recall, false identity selection, attempts, latency, cost, and downstream effects.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yiqi Wang et al. (2026). *From Agent Traces to Trust: Evidence Tracing and Execution Provenance in LLM Agents*. arXiv. <https://arxiv.org/abs/2606.04990>
  - **Why it matters here:** surveys trace schemas and provenance links across tools, evidence, actions, and outcomes, relevant to the required event fields.
  - **Claim it would support:** “The trace must make that comparison possible.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Bhaskar Gurram (2026). *Auditing Automated Evaluation, Error Propagation, and Runtime Mitigation in Tool-Using Language Agents*. arXiv. <https://arxiv.org/abs/2604.16706>
  - **Why it matters here:** audits automated judging and error propagation over execution traces, relevant to validating pattern metrics before using them to change contracts.
  - **Claim it would support:** “A trace can explain what happened ... but it cannot by itself ... turn a common sequence into policy.”
  - **Notes file:** [toolformer-notes.md](../research/toolformer-notes.md)
  - **Miriah's notes:**

## Deferred / out of scope

- No chapter or module prose is edited during bootstrap; the loop requires the first iteration to create only this ledger.
- Slide edits are deferred to slide-specific passes; the Chicago deck currently has day-level decks (`day1`–`day5`) rather than one deck per book module.
- The standalone `ch00` chapter has no `.outline.md`; it was reviewed as the bootstrap exception rather than receiving an invented outline.

## Questions for Miriah

- `blog posts/` and `blog-posts/` both exist (space versus hyphen). Are both live, or is one legacy/archive material?
- The flat research files `chapter-3.md`, `chapter-5.md`, `chapter-6.md`, `chapter-7.md`, `chapter-12.md`, and `notes.md` sit alongside the structured `book/chapters/` tree. Are they legacy drafts, source material, or live manuscript inputs?
- Should the task ledger remain at `book/RESEARCH_LEDGER.md` despite `AGENTS.md` saying new research files belong under `research/`? The bootstrap used the task-specified location and recorded the conflict above.
- The curriculum-to-book mappings above are provisional; please confirm whether “mapping: unsure” rows should be narrowed to direct numbered modules or may cite adjacent modules.
- The source curriculum has modules `00_Setup` and `01_Dev_Environment`, while the book has no corresponding numbered chapter. Should those remain only as cross-cutting context, or receive explicit mappings in later passes?
- The source curriculum’s Chicago slides are organized by training day, not book module. Should each module ledger row name a likely day deck in later passes?
