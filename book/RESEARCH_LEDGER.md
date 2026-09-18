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
| ch06 | ch06.03-user-session-workflow-state | mapping: unsure — 10, 11, 23 | done | 2026-09-18 | open: review three queued workflow/isolation sources; confirm whether 23 is an adjacent mapping; align day3 state/recovery slides in a later slide pass |
| ch07 | ch07.01-sources-of-context | 05, 06, 10, 11 | done | 2026-09-18 | open: review three retrieval sources; align day2 retrieval/source-contract slides in a later slide pass |
| ch07 | ch07.02-context-assembly-pipelines | 05, 06, 07, 11 | done | 2026-09-18 | open: review three long-context sources; align day2 retrieval comparison with the manifest/decision boundary in a slide pass |
| ch07 | ch07.03-freshness-consistency-and-partial-failure | mapping: unsure — 05, 09, 22 | done | 2026-09-18 | open: review three conflict/freshness sources; confirm mapping; align day2/day3 freshness and drift material in a slide pass |
| ch07 | ch07.04-hydration-coverage-and-retrieval-success | 05, 06, 08, 21 | done | 2026-09-18 | open: review three retrieval-evaluation sources; align day2 evaluation material in a later slide pass |
| ch07 | ch07.05-public-data-sources-wikipedia-web | mapping: unsure — 05, 17 | done | 2026-09-18 | open: review three public-retrieval sources; confirm mapping; align day2/day4 web-research material in a slide pass |
| ch07 | ch07.06-information-extraction-pipelines | 05, 08, 16 | done | 2026-09-18 | open: review three queued IE/provenance sources; carry the proposal/validation boundary into the day4 graph-extractor slide |
| ch08 | ch08-knowledge-graphs-and-semantic-context | mapping: unsure — 16 | done | 2026-09-18 | open: review three queued GraphRAG/semantic sources; carry governed-semantic-surface framing into the day4 GraphRAG slide; standalone overview has no outline by repository structure |
| ch08 | ch08.01-schemas-taxonomies-and-ontologies | 16 | done | 2026-09-18 | open: review three queued ontology/validation sources; carry the representation-selection flow into the day4 GraphRAG slide |
| ch08 | ch08.02-rdf-owl-and-sparql | mapping: unsure — 16 | done | 2026-09-18 | open: review three queued RDF/SPARQL/query-check sources; carry the standards-boundary flow into the day4 GraphRAG slide |
| ch08 | ch08.03-entity-resolution-and-relationship-traversal | 16 | done | 2026-09-18 | open: review three queued entity-resolution/linking sources; carry entity-linking and bounded k-hop vocabulary into the day4 GraphRAG slide |
| ch08 | ch08.04-knowledge-graph-tradeoffs | mapping: unsure — 05, 16 | done | 2026-09-18 | open: review three queued graph-evaluation/refinement sources; carry the baseline-to-pilot loop into the day4 GraphRAG slide |
| ch08 | ch08.05-instance-coverage-and-ontology-population | 16, 21 | done | 2026-09-18 | open: review three queued completeness/refinement sources; carry the population-to-evaluation flow into the day4 GraphRAG and day5 eval slides |
| ch08 | ch08.06-property-completeness-and-schema-quality | 16, 21 | done | 2026-09-18 | open: review three queued completeness/evaluation sources; carry requirement-specific missingness into the day4 GraphRAG and day5 RAGAS slides |
| ch08 | ch08.07-ontology-guided-information-extraction | 16 | done | 2026-09-18 | open: review three queued ontology-guided IE sources; carry the contract and validation flows into the day4 GraphRAG slide |
| ch08 | ch08.08-knowledge-extraction-methods | 08, 16 | done | 2026-09-18 | pre-edit gaps: asset—method cascade lacked Mermaid; coverage—curriculum's three graph construction methods and schema comparison were not named; evidence—method-selection claims needed primary sources; research—queue extraction-method sources; open: review three queued sources and carry the method-routing flow into the day4 GraphRAG slide |
| ch08 | ch08.09-guardrails-for-extraction-validation | 13, 18, 21 | done | 2026-09-18 | open: review three queued validation/evaluation sources; carry the layered gates into the day5 RAGAS/DeepEval slides |
| ch08 | ch08.10-multilingual-extraction-with-llms | mapping: unsure — 05, 16 | done | 2026-09-18 | open: review three queued multilingual evaluation sources; carry the stratified route/evaluation loop into a future day5 slide pass; curriculum mapping remains unsure |
| ch09 | ch09.01-lexical-and-relational-retrieval | 05, 06 | done | 2026-09-18 | open: review three queued lexical/fusion sources; carry query routing into the day2 retrieval slide |
| ch09 | ch09.02-vector-and-semantic-retrieval | 05, 06 | done | 2026-09-18 | open: review four queued dense-retrieval/ANN sources; carry the dense flow and source-path example into the day2 retrieval slide |
| ch09 | ch09.03-graph-and-hybrid-retrieval | 06, 16 | done | 2026-09-18 | open: review three queued graph/hybrid sources; create a maintained source repository for the illustrative implementations |
| ch09 | ch09.04-ranking-reranking-and-query-planning | 06, 07 | done | 2026-09-18 | open: review three queued reranking/expansion/evaluation sources; create a maintained source repository for the illustrative implementations |
| ch09 | ch09.05-context-precision-and-context-recall | 05, 06, 08, 21 | done | 2026-09-18 | open: review three queued evaluation sources; create a maintained source repository for the illustrative implementations |
| ch10 | ch10.00-guardrails-and-ontology-based-validation | 13, 16, 18 | done | 2026-09-18 | open: review three queued guardrail sources; create a maintained source repository for the illustrative implementations |
| ch10 | ch10.01-personalization-as-retrieval | 10, 11 | done | 2026-09-18 | open: review three queued memory sources; create a maintained source repository for the illustrative implementations |
| ch10 | ch10.02-scoped-hydration | 10, 13 | done | 2026-09-18 | open: review three queued scoped-retrieval sources; create a maintained source repository for the illustrative implementations |
| ch10 | ch10.03-provenance-and-derived-context | 10, 22 | done | 2026-09-18 | open: review three queued provenance sources; create a maintained source repository for the illustrative implementations |
| ch10 | ch10.04-policy-aware-user-context | 13, 19, 20 | done | 2026-09-18 | open: review three queued policy-aware context sources; create a maintained source repository for the illustrative implementations |
| ch10 | ch10.05-provenance-coverage-metrics | 09, 21, 22 | done | 2026-09-18 | open: review three queued coverage/evaluation sources; create a maintained source repository for the illustrative implementations |
| ch11 | ch11.01-least-privilege | 13, 19, 20 | done | 2026-09-18 | open: review three queued least-authority sources; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch11 | ch11.02-rbac-abac-capability-based-access | mapping: unsure — 13, 19 | done | 2026-09-18 | open: review three queued capability/deputy sources; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch11 | ch11.03-scoped-credentials-knowledge-stores | 11, 13, 20 | done | 2026-09-18 | open: review four queued scoped-credential sources; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch11 | ch11.04-retrieval-execution-boundaries | 07, 11, 13 | done | 2026-09-18 | open: review three queued authorization/action-boundary sources; add direct TOCTOU evidence if needed; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch11 | ch11.05-authorization-coverage-and-necessary-access | 13, 19, 20 | done | 2026-09-18 | open: review three queued minimization/coverage sources; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch12 | ch12.01-small-composable-systems | 01, 11 | done | 2026-09-18 | open: review three queued UNIX/distributed-systems sources; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch12 | ch12.02-pipes-files-explicit-interfaces | 01, 11 | done | 2026-09-18 | open: review three queued interface/provenance sources; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch12 | ch12.03-mounts-namespaces-isolation | mapping: unsure — 01, 13, 20 | done | 2026-09-18 | open: review three queued namespace/isolation sources; confirm mapping; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch12 | ch12.04-task-workspaces-secret-management | 01, 11, 13 | done | 2026-09-18 | open: review three queued workspace/secret sources; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch13 | ch13.01-planning-and-react | 03, 11 | done | 2026-09-18 | open: review three queued planning/agent-evaluation sources; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch13 | ch13.02-harnesses-and-state-machines | 11, 23 | done | 2026-09-18 | open: review three queued harness/state sources; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch13 | ch13.03-durable-and-event-driven-execution | mapping: unsure — 11, 23 | done | 2026-09-18 | open: review four queued durable-workflow sources; confirm mapping; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch13 | ch13.04-loops-retries-and-bounded-autonomy | 03, 11, 13, 23 | done | 2026-09-18 | open: review five queued loop/retry/restoration sources; confirm mapping; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch14 | ch14.01-token-and-context-economics | 02, 05 | done | 2026-09-18 | open: review four queued context-economics sources; confirm mapping; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch14 | ch14.02-retrieval-tool-and-latency-costs | 05, 06, 07 | done | 2026-09-18 | open: review four queued retrieval/latency sources; confirm mapping; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch14 | ch14.03-one-shot-execution-loops-and-subagents | 03, 11, 12 | done | 2026-09-18 | open: review four queued one-shot/loop/delegation sources; confirm mapping; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch14 | ch14.04-local-models-and-model-routing | mapping: unsure — 03, 15 | done | 2026-09-18 | open: review four queued local-model/routing sources; confirm mapping; create a maintained source repository; clean curriculum-facing notes from earlier modules |
| ch14 | ch14.05-context-efficiency-metrics | 08, 09, 21, 22 | done | 2026-09-18 | open: review four queued context-efficiency/evaluation sources; confirm mapping; create a maintained source repository; clean curriculum-facing notes from earlier modules |
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

### ch11.03-scoped-credentials-knowledge-stores

- [ ] **unreviewed** — Michael Jones, Anthony Nadalin, Brian Campbell, John Bradley, and Chuck Mortimore (2020). *RFC 8693: OAuth 2.0 Token Exchange*. IETF Proposed Standard. <https://www.rfc-editor.org/rfc/rfc8693>
  - **Why it matters here:** defines token exchange for impersonation and delegation, including downstream tokens that can be more narrowly scoped; relevant to “Broker Narrow Authority”.
  - **Claim it would support:** “A trusted host should obtain or reference the credential and attach it only on the protected network call.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Laurent Chuat, AbdelRahman Abdou, Ralf Sasse, Christoph Sprenger, David Basin, and Adrian Perrig (2020). *SoK: Delegation and Revocation, the Missing Links in the Web's Chain of Trust*. IEEE European Symposium on Security and Privacy. <https://arxiv.org/abs/1906.10775>
  - **Why it matters here:** surveys delegation and revocation trade-offs, including short-lived credentials and revocation latency; relevant to “Choose Revocation and Validation Deliberately”.
  - **Claim it would support:** “Short lifetimes reduce exposure and increase renewal load.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Meng Sun, Junzuo Lai, Wei Wu, Ye Yang, Cheng Kang Chu, and Robert H. Deng (2024). *How to Securely Delegate and Revoke Partial Authorization Credentials*. IEEE Transactions on Dependable and Secure Computing. <https://doi.org/10.1109/TDSC.2024.3424520>
  - **Why it matters here:** formalizes partial delegation and revocation; relevant to “Credential chaining and delegation require explicit attenuation”.
  - **Claim it would support:** “A child grant must not exceed its parent.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Michael B. Jones, John Bradley, and Nat Sakimura (2023). *RFC 9449: OAuth 2.0 Demonstrating Proof-of-Possession at the Application Layer (DPoP)*. IETF Proposed Standard. <https://www.rfc-editor.org/rfc/rfc9449>
  - **Why it matters here:** describes sender-constrained tokens that reduce replay after token leakage; relevant to “Operate the Broker as Security-Critical Infrastructure”.
  - **Claim it would support:** “Test replay, wrong audience, stale policy, parent revocation, concurrent single-use, leaked handle, broker outage.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**

### ch11.04-retrieval-execution-boundaries

- [ ] **unreviewed** — Jon Howell and David Kotz (2000). *End-to-end authorization*. Proceedings of the 4th USENIX Symposium on Operating System Design and Implementation. <https://www.usenix.org/legacy/events/osdi2000/full_papers/howell/howell_html/>
  - **Why it matters here:** analyzes how gateways can lose the authority of the original caller and how resource servers can verify end-to-end authority; relevant to “Contain a Deceived Model”.
  - **Claim it would support:** “Alternate tools, service credentials, caches, and direct network access can bypass the intended boundary.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Dimitrios Stamatios Bouras, Yihan Dai, and Sergey Mechtaev (2026). *Authority Is Not a String: A Capability-Scoped Harness for Prompt-Injection-Resistant Coding Agents*. arXiv:2609.08371. <https://arxiv.org/abs/2609.08371>
  - **Why it matters here:** evaluates host-side pre-dispatch capability checks against injected tool proposals; relevant to “Contain a Deceived Model”.
  - **Claim it would support:** “Authorization cannot prevent a language process from being manipulated; it can prevent a manipulated proposal from exceeding verified authority.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Qiusi Zhan, Zhixiang Liang, Zifan Ying, and Daniel Kang (2024). *InjecAgent: Benchmarking Indirect Prompt Injections in Tool-Integrated Large Language Model Agents*. Findings of the Association for Computational Linguistics. <https://aclanthology.org/2024.findings-acl.624/>
  - **Why it matters here:** benchmarks indirect prompt injection against tool-integrated agents, relevant to testing retrieved instructions as untrusted inputs to action proposals.
  - **Claim it would support:** “Retrieval and execution need separate authorization because they protect different resources, actions, purposes, and consequences.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**

### ch11.05-authorization-coverage-and-necessary-access

- [ ] **unreviewed** — Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv:2309.15217. <https://arxiv.org/abs/2309.15217>
  - **Why it matters here:** separates context relevance, faithfulness, and answer quality; relevant to “Evaluate Minimization and Sufficiency Together”.
  - **Claim it would support:** “Pair necessary-access precision with required-slot coverage, retrieval recall, task success, and safe abstention.”
  - **Notes file:** [necessary-access-notes.md](../research/necessary-access-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Asia J. Biega, Peter Potash, Hal Daumé III, Fernando Diaz, and Michèle Finck (2020). *Operationalizing the Legal Principle of Data Minimization for Personalization*. Proceedings of SIGIR 2020. <https://doi.org/10.1145/3397271.3401034>
  - **Why it matters here:** studies operational definitions and performance effects of data minimization; relevant to “Bound Necessity Before Retrieval”.
  - **Claim it would support:** “A metric should not allow unnecessary-but-authorized data to hide an access-control incident.”
  - **Notes file:** [necessary-access-notes.md](../research/necessary-access-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Kaige Xie, Philippe Laban, Prafulla Kumar Choubey, Caiming Xiong, and Chien-Sheng Wu (2025). *Do RAG Systems Cover What Matters? Evaluating and Optimizing Responses with Sub-Question Coverage*. NAACL 2025. <https://aclanthology.org/2025.naacl-long.301/>
  - **Why it matters here:** evaluates coverage of core and secondary sub-questions; relevant to “Evaluate Minimization and Sufficiency Together”.
  - **Claim it would support:** “Required-slot coverage” should be reported beside precision and retrieval cost.
  - **Notes file:** [necessary-access-notes.md](../research/necessary-access-notes.md)
  - **Miriah's notes:**

### ch12.01-small-composable-systems

- [ ] **unreviewed** — Dennis M. Ritchie and Ken Thompson (1978). *The UNIX Time-Sharing System*. Bell System Technical Journal, 57, 1905–1929. <https://doi.org/10.1002/j.1538-7305.1978.tb02136.x>
  - **Why it matters here:** documents UNIX command and file-system interfaces; relevant to “Composition Moves Complexity to Contracts”.
  - **Claim it would support:** “The UNIX analogy is useful because it emphasizes focused programs and composition.”
  - **Notes file:** [unix-composition-notes.md](../research/unix-composition-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — W3C Provenance Working Group (2013). *PROV-DM: The PROV Data Model*. W3C Recommendation. <https://www.w3.org/TR/prov-dm/>
  - **Why it matters here:** defines a provenance model for entities, activities, and agents; relevant to “Interfaces Carry the Spine”.
  - **Claim it would support:** “A stage envelope can carry the boundary metadata alongside its payload.”
  - **Notes file:** [interface-envelope-notes.md](../research/interface-envelope-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Apostolos Destounis, Georgios S. Paschos, and Iordanis Koutsopoulos (2016). *Streaming Big Data meets Backpressure in Distributed Network Computation*. IEEE INFOCOM 2016. <https://arxiv.org/abs/1601.03876>
  - **Why it matters here:** studies query streams limited by communication and computation capacity; relevant to “Context Pipelines Need Typed Envelopes”.
  - **Claim it would support:** “They must preserve cancellation, deadlines, errors, provenance, and backpressure.”
  - **Notes file:** [interface-envelope-notes.md](../research/interface-envelope-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Ioannis Chatzigiannakis, Sotiris Nikoletseas, and Paul G. Spirakis (2004). *Distributed Computation and Communication in Wireless Sensor Networks*. Theoretical Computer Science, 323(1–3), 175–197. <https://doi.org/10.1016/j.tcs.2004.04.012>
  - **Why it matters here:** treats communication and computation as coupled distributed resources; relevant to “File-Like Is an Analogy, Not a Universal API”.
  - **Claim it would support:** “Consumers should know the guarantees they depend on even if implementations remain replaceable.”
  - **Notes file:** [interface-envelope-notes.md](../research/interface-envelope-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Jerome H. Saltzer, David P. Reed, and David D. Clark (1984). *End-to-End Arguments in System Design*. ACM Transactions on Computer Systems, 2(4), 277–288. <https://doi.org/10.1145/357401.357402>
  - **Why it matters here:** gives a principled basis for placing functions at system layers; relevant to “Split at Enforceable Boundaries”.
  - **Claim it would support:** “A component earns separation when it has a coherent responsibility, independently testable contract, different scaling or security needs, or a failure that should be contained.”
  - **Notes file:** [unix-composition-notes.md](../research/unix-composition-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Jeffrey Dean and Luiz André Barroso (2013). *The Tail at Scale*. Communications of the ACM, 56(2), 74–80. <https://doi.org/10.1145/2408776.2408794>
  - **Why it matters here:** analyzes latency variability in large distributed services; relevant to “Composition Moves Complexity to Contracts”.
  - **Claim it would support:** “Decomposition can also increase network hops, partial failure, version skew, and operational ownership.”
  - **Notes file:** [unix-composition-notes.md](../research/unix-composition-notes.md)
  - **Miriah's notes:**

### ch12.03-mounts-namespaces-isolation

- [ ] **unreviewed** — Yuqiong Sun, David Safford, Mimi Zohar, Dimitrios Pendarakis, Zhongshu Gu, and Trent Jaeger (2018). *Security Namespace: Making Linux Security Frameworks Available to Containers*. 27th USENIX Security Symposium. <https://www.usenix.org/conference/usenixsecurity18/presentation/sun>
  - **Why it matters here:** distinguishes virtualization from scoped security enforcement; relevant to “A Namespace Is Not Isolation by Itself”.
  - **Claim it would support:** “Operating-system mechanisms may help enforce [logical context boundaries], but they do not replace backend authorization.”
  - **Notes file:** [namespace-isolation-notes.md](../research/namespace-isolation-notes.md)
  - **Miriah's notes:**

### ch12.04-task-workspaces-secret-management

- [ ] **unreviewed** — Setu Kumar Basak, Lorenzo Neil, Bradley Reaves, and Laurie Williams (2022). *What are the Practices for Secret Management in Software Artifacts?* IEEE Secure Development Conference. <https://arxiv.org/abs/2208.11280>
  - **Why it matters here:** identifies practices for external secret storage, scanning, and short-lived secrets; relevant to “Give Executors Handles, Not Secrets”.
  - **Claim it would support:** “Rotate and revoke it independently of workspace deletion.”
  - **Notes file:** [workspace-secret-management-notes.md](../research/workspace-secret-management-notes.md)
  - **Miriah's notes:**

### ch13.01-planning-and-react

- [ ] **unreviewed** — Shunyu Yao et al. (2023). *ReAct: Synergizing Reasoning and Acting in Language Models*. ICLR. <https://arxiv.org/abs/2210.03629>
  - **Why it matters here:** describes interleaved actions and observations; relevant to “ReAct Is an Interaction Pattern”.
  - **Claim it would support:** “ReAct interleaves reasoning and acting so that observations from an environment can inform later actions.”
  - **Notes file:** [agent-planning-notes.md](../research/agent-planning-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Xiao Liu et al. (2023). *AgentBench: Evaluating LLMs as Agents*. ICLR 2024. <https://arxiv.org/abs/2308.03688>
  - **Why it matters here:** evaluates agents across environments and identifies long-horizon reasoning and decision failures; relevant to “Design From Failure Paths”.
  - **Claim it would support:** “Measure invalid-transition attempts, stale-write conflicts, time spent per state, terminal outcome rates, and manual interventions.”
  - **Notes file:** [harness-state-machine-notes.md](../research/harness-state-machine-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Edoardo Debenedetti et al. (2024). *AgentDojo: A Dynamic Environment to Evaluate Prompt Injection Attacks and Defenses for LLM Agents*. NeurIPS 2024. <https://arxiv.org/abs/2406.13352>
  - **Why it matters here:** evaluates tool-using agents under dynamic state and untrusted content; relevant to “Put Enforcement Outside the Model”.
  - **Claim it would support:** “The model may propose ... but those proposals are inputs to trusted code.”
  - **Notes file:** [harness-state-machine-notes.md](../research/harness-state-machine-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shunyu Yao et al. (2024). *Understanding the Planning of LLM Agents: A Survey*. arXiv:2402.02716. <https://arxiv.org/abs/2402.02716>
  - **Why it matters here:** surveys planning and state-handling patterns; relevant to “Make Workflow State Explicit”.
  - **Claim it would support:** “Model only states that change authority, recovery behavior, user-visible progress, resource ownership, or the meaning of later operations.”
  - **Notes file:** [harness-state-machine-notes.md](../research/harness-state-machine-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Edoardo Debenedetti, Jie Zhang, Mislav Balunovic, Luca Beurer-Kellner, Marc Fischer, and Florian Tramèr (2024). *AgentDojo: A Dynamic Environment to Evaluate Prompt Injection Attacks and Defenses for LLM Agents*. NeurIPS 2024. <https://arxiv.org/abs/2406.13352>
  - **Why it matters here:** evaluates realistic tool-using tasks and security properties under untrusted data; relevant to “Decide With an Evaluation, Not a Framework List”.
  - **Claim it would support:** “Test ... deterministic tests for invalid transitions and injected failures.”
  - **Notes file:** [agent-planning-notes.md](../research/agent-planning-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shunyu Yao et al. (2024). *Understanding the Planning of LLM Agents: A Survey*. arXiv:2402.02716. <https://arxiv.org/abs/2402.02716>
  - **Why it matters here:** surveys planning methods and their components; relevant to “A Plan Is Proposed State”.
  - **Claim it would support:** “Use more elaborate planning only when the system has a bounded action space ... an evaluator correlated with the production outcome ... and a budget.”
  - **Notes file:** [agent-planning-notes.md](../research/agent-planning-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Joel Reardon, Hubert Ritzdorf, David Basin, and Srdjan Čapkun (2013). *Secure Data Deletion from Persistent Media*. Proceedings of the 2013 ACM SIGSAC Conference on Computer and Communications Security. <https://doi.org/10.1145/2508859.2516699>
  - **Why it matters here:** analyzes secure deletion with encryption and key wrapping; relevant to “Define Lifecycle and Ownership”.
  - **Claim it would support:** “Deleting a directory is not enough if checkpoints, logs, caches, snapshots, or credential leases still retain the task's data or authority.”
  - **Notes file:** [workspace-secret-management-notes.md](../research/workspace-secret-management-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Zhihao Chen, Ying Zhang, Yi Liu, Gelei Deng, Yuekang Li, Yanjun Zhang, Jianting Ning, Leo Yu Zhang, Lei Ma, and Zhiqiang Li (2026). *How Your Credentials Are Leaked by LLM Agent Skills: An Empirical Study*. arXiv:2604.03070, version 2. <https://arxiv.org/abs/2604.03070>
  - **Why it matters here:** measures debug logging and other skill pathways that expose credentials to model context; relevant to “Give Executors Handles, Not Secrets”.
  - **Claim it would support:** “Environment variables and credential files ... can leak through child processes, crash dumps, logs, diagnostics, and accidental reads.”
  - **Notes file:** [workspace-secret-management-notes.md](../research/workspace-secret-management-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — William Findlay, David Barrera, and Anil Somayaji (2021). *BPFContain: Fixing the Soft Underbelly of Container Security*. arXiv:2102.06972. <https://arxiv.org/abs/2102.06972>
  - **Why it matters here:** analyzes the limited guarantees of namespaces and resource partitioning and the need for additional confinement policy; relevant to “Use Real Isolation for Untrusted Execution”.
  - **Claim it would support:** “Namespaces and cgroups on their own cannot enforce fine-grained access policies.”
  - **Notes file:** [namespace-isolation-notes.md](../research/namespace-isolation-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Maryam Rostamipoor, Seyedhamed Ghavamnia, and Michalis Polychronakis (2023). *Confine: Fine-grained System Call Filtering for Container Attack Surface Reduction*. Computers & Security, 132, 103325. <https://doi.org/10.1016/j.cose.2023.103325>
  - **Why it matters here:** evaluates syscall filtering as defense in depth for containers; relevant to “Contain Resources and Side Effects”.
  - **Claim it would support:** “Use mechanisms appropriate to the threat model: ... syscall controls, resource limits, and distinct credentials.”
  - **Notes file:** [namespace-isolation-notes.md](../research/namespace-isolation-notes.md)
  - **Miriah's notes:**

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

### ch06.03-user-session-workflow-state

- [ ] **unreviewed** — Riya Samanta et al. (2026). *AgentR: A Stateful and Recovery-Aware Software Architecture for LLM-based Auditable Workflows*. arXiv. <https://arxiv.org/abs/2608.15264>
  - **Why it matters here:** describes durable workflow artifacts, explicit transitions, retries, and orphan-job handling, relevant to workflow state as a recovery contract.
  - **Claim it would support:** “Workflow state records the pragmatic progress of a specific execution.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch07.01-sources-of-context

- [ ] **unreviewed** — Vladimir Karpukhin et al. (2020). *Dense Passage Retrieval for Open-Domain Question Answering*. EMNLP. <https://aclanthology.org/2020.emnlp-main.550/>
  - **Why it matters here:** provides a primary dense-retrieval formulation and evaluation, relevant to separating candidate selection from authority and answer generation.
  - **Claim it would support:** “Dense retrieval finds proximity in a learned representation. Neither establishes that a result is true, current, authorized, or sufficient.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch07.02-context-assembly-pipelines

- [ ] **unreviewed** — Nelson F. Liu et al. (2023). *Lost in the Middle: How Language Models Use Long Contexts*. arXiv. <https://arxiv.org/abs/2307.03172>
  - **Why it matters here:** measures position-sensitive use of long inputs, relevant to ranking, truncation, and non-uniform context budgets.
  - **Claim it would support:** “The assembler still has to ... reserve space ... truncate at semantic boundaries.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch07.03-freshness-consistency-and-partial-failure

- [ ] **unreviewed** — Arie Cattan et al. (2025). *DRAGged into Conflicts: Detecting and Addressing Conflicting Sources in Search-Augmented LLMs*. arXiv. <https://arxiv.org/abs/2506.08500>
  - **Why it matters here:** studies conflict types including temporal freshness conflicts, relevant to resolving disagreement before serialization.
  - **Claim it would support:** “Some contradictions are governed by an application invariant; others require investigation.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch07.04-hydration-coverage-and-retrieval-success

- [ ] **unreviewed** — Shahul Es et al. (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv. <https://arxiv.org/abs/2309.15217>
  - **Why it matters here:** separates retrieval relevance, faithfulness, and generation quality, relevant to keeping hydration coverage diagnostic.
  - **Claim it would support:** “Coverage also does not measure relevance within a slot.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**

### ch07.05-public-data-sources-wikipedia-web

- [ ] **unreviewed** — Tu Vu et al. (2023). *FreshLLMs: Refreshing Large Language Models with Search Engine Augmentation*. arXiv. <https://arxiv.org/abs/2310.03214>
  - **Why it matters here:** introduces dynamic QA and search-augmented prompting, relevant to freshness and retrieval-order effects in public evidence.
  - **Claim it would support:** “A search result is routing metadata, not final factual context.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yufang Hou et al. (2024). *WikiContradict: A Benchmark for Evaluating LLMs on Real-World Knowledge Conflicts from Wikipedia*. arXiv. <https://arxiv.org/abs/2406.13805>
  - **Why it matters here:** tests contradictory Wikipedia passages, relevant to treating a secondary source as attributed evidence rather than unqualified truth.
  - **Claim it would support:** “Wikipedia is a navigable secondary source.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Badrinath Ramakrishnan and Akshaya Balaji (2025). *Securing AI Agents Against Prompt Injection Attacks*. arXiv. <https://arxiv.org/abs/2511.15759>
  - **Why it matters here:** benchmarks injection risks in retrieval-augmented agents, relevant to public-page trust boundaries.
  - **Claim it would support:** “Public pages also create prompt-injection risk.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Nandan Thakur et al. (2021). *BEIR: A Heterogeneous Benchmark for Zero-shot Evaluation of Information Retrieval Models*. arXiv. <https://arxiv.org/abs/2104.08663>
  - **Why it matters here:** evaluates retrieval across diverse tasks and model families, relevant to within-slot retrieval success.
  - **Claim it would support:** “Pair coverage with retrieval precision and recall estimates.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Hao Yu et al. (2024). *Evaluation of Retrieval-Augmented Generation: A Survey*. arXiv. <https://arxiv.org/abs/2405.07437>
  - **Why it matters here:** organizes retrieval, generation, relevance, accuracy, and faithfulness measures, relevant to the module's limitation boundaries.
  - **Claim it would support:** “Hydration coverage answers one narrow question.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Authors not captured in this pass (2025). *Retrieval-Augmented Generation with Conflicting Evidence*. arXiv. <https://arxiv.org/abs/2504.13079>
  - **Why it matters here:** evaluates ambiguity, misinformation, and noise in retrieved evidence, relevant to deciding when to answer, qualify, or abstain.
  - **Claim it would support:** “Graceful degradation is not ‘answer with whatever arrived.’”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Ziyu Ge et al. (2025). *Resolving Conflicting Evidence in Automated Fact-Checking: A Study on Retrieval-Augmented LLMs*. arXiv. <https://arxiv.org/abs/2505.17762>
  - **Why it matters here:** evaluates source credibility when retrieved evidence conflicts, relevant to authority rules and incorrect-proceed measurements.
  - **Claim it would support:** “The model is not the authority boundary.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Junqing He et al. (2023). *Never Lost in the Middle: Mastering Long-Context Question Answering with Position-Agnostic Decompositional Training*. arXiv. <https://arxiv.org/abs/2311.09198>
  - **Why it matters here:** tests positional retrieval failures and mitigation, relevant to evaluating assembly order rather than assuming uniform model access.
  - **Claim it would support:** “Budget allocation is policy.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — George Arthur Baker et al. (2024). *Lost in the Middle, and In-Between: Enhancing Language Models' Ability to Reason Over Long Contexts in Multi-Hop QA*. arXiv. <https://arxiv.org/abs/2412.10079>
  - **Why it matters here:** extends positional failure analysis to multi-hop evidence, relevant to preserving related evidence through truncation.
  - **Claim it would support:** “Summarization may save tokens, but it is a lossy transformation.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Gordon V. Cormack et al. (2009). *Reciprocal Rank Fusion outperforms Condorcet and Individual Rank Learning Methods*. SIGIR. <https://doi.org/10.1145/1571941.1572114>
  - **Why it matters here:** supports rank fusion as a selection operation rather than a source-authority decision.
  - **Claim it would support:** “These mechanisms solve different problems.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Omar Khattab and Matei Zaharia (2020). *ColBERT: Efficient and Effective Passage Search via Contextualized Late Interaction over BERT*. SIGIR. <https://arxiv.org/abs/2004.12832>
  - **Why it matters here:** provides a primary late-interaction retrieval design, relevant to distinguishing selection operations in a source contract.
  - **Claim it would support:** “Documents can be retrieved through exact metadata filters, lexical search, dense similarity, or hybrid ranking.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yu Zhuang et al. (2026). *AgentRewind: Recoverable Execution for Long-Horizon LLM Agents*. arXiv. <https://arxiv.org/abs/2608.14380>
  - **Why it matters here:** uses aligned checkpoints of agent context and environment state for recovery, relevant to distinguishing session context from resumable workflow state.
  - **Claim it would support:** “If losing a value could repeat an effect ... that value belongs in durable workflow state.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Xinyu Gao et al. (2026). *Isolated but Exposed: Persistence-Based Memory Extraction Attack on LLM Agents*. arXiv. <https://arxiv.org/abs/2607.23444>
  - **Why it matters here:** shows that user-level memory isolation does not eliminate tool-side exfiltration risk, relevant to testing state boundaries and downstream capability use.
  - **Claim it would support:** “Session isolation should be tested explicitly.”
  - **Notes file:** [context-engineering-foundations-notes.md](../research/context-engineering-foundations-notes.md)
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

### ch07.06-information-extraction-pipelines

- [ ] **unreviewed** — Yuan Yao, Deming Ye, Peng Li, Xu Han, Yankai Lin, Zhenghao Liu, Zhiyuan Liu, Lixin Huang, Jie Zhou, and Maosong Sun (2019). *DocRED: A Large-Scale Document-Level Relation Extraction Dataset*. Proceedings of ACL. <https://aclanthology.org/P19-1074/>
  - **Why it matters here:** establishes that document-level relation extraction may require synthesizing evidence across multiple sentences and remains difficult even with a dedicated benchmark.
  - **Claim it would support:** “The choice is empirical. Evaluate it on representative documents” and the warning that aggregate extraction performance does not make the pipeline authoritative.
  - **Notes file:** [semantic-web-paper-notes.md](../research/semantic-web-paper-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Ying Lin, Heng Ji, Fei Huang, and Lingfei Wu (2020). *A Joint Neural Model for Information Extraction with Global Features*. Proceedings of ACL. <https://aclanthology.org/2020.acl-main.713/>
  - **Why it matters here:** presents OneIE as a joint graph extraction framework for entity mentions, event triggers, and links, giving a primary source for structured IE as a coordinated output rather than isolated labels.
  - **Claim it would support:** “Rules, statistical models, and language models can all participate” and the distinction between proposing structured state and promoting it.
  - **Notes file:** [semantic-web-paper-notes.md](../research/semantic-web-paper-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Gabriel Amaral, Odinaldo Rodrigues, and Elena Simperl (2022). *ProVe: A Pipeline for Automated Provenance Verification of Knowledge Graphs against Textual Sources*. arXiv:2210.14846. <https://arxiv.org/abs/2210.14846>
  - **Why it matters here:** directly studies automated verification that a graph triple is supported by the text documented as its provenance, matching the module's source-entailment and evidence-preservation boundary.
  - **Claim it would support:** “A model's parameters are not a source for a document extraction” and “require exact passage locators so unsupported additions can be detected.”
  - **Notes file:** [semantic-web-paper-notes.md](../research/semantic-web-paper-notes.md)
  - **Miriah's notes:**

### ch08.02-rdf-owl-and-sparql

- [ ] **unreviewed** — W3C RDF Working Group (2014). *RDF 1.1 Concepts and Abstract Syntax*. W3C Recommendation. <https://www.w3.org/TR/rdf11-concepts/>
  - **Why it matters here:** provides the primary data-model definition for RDF graphs, IRIs, blank nodes, literals, and merging, supporting the module's representation boundary.
  - **Claim it would support:** “An RDF triple has an IRI or blank node as its subject, an IRI as its predicate, and an IRI, blank node, or literal as its object.”
  - **Notes file:** [04-storage-indexing-and-querying.md](../research/knowledge-graphs/04-storage-indexing-and-querying.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — W3C SPARQL Working Group (2013). *SPARQL 1.1 Query Language*. W3C Recommendation. <https://www.w3.org/TR/sparql11-query/>
  - **Why it matters here:** provides the primary query semantics for graph patterns and result bindings, supporting the module's distinction between querying and reasoning.
  - **Claim it would support:** “SPARQL matches graph patterns and returns bindings, constructs graphs, asks whether a pattern exists, or describes resources.”
  - **Notes file:** [04-storage-indexing-and-querying.md](../research/knowledge-graphs/04-storage-indexing-and-querying.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Dean Allemang and Juan F. Sequeda (2024). *Ontologies to the Rescue? An Ontology-Based Query Check for LLMs*. arXiv:2405.11706. <https://arxiv.org/abs/2405.11706>
  - **Why it matters here:** evaluates ontology-aware checking and repair of generated SPARQL in one enterprise QA benchmark, supporting the module's bounded claim that semantic checking can improve a pipeline without guaranteeing general correctness.
  - **Claim it would support:** “This is evidence that ontology-aware query checking can improve one text-to-SPARQL pipeline, not that an ontology guarantees correct answers across domains.”
  - **Notes file:** [04-storage-indexing-and-querying.md](../research/knowledge-graphs/04-storage-indexing-and-querying.md)
  - **Miriah's notes:**

### ch08.03-entity-resolution-and-relationship-traversal

- [ ] **unreviewed** — Olivier Binette and Rebecca C. Steorts (2020). *(Almost) All of Entity Resolution*. arXiv:2008.04443. <https://arxiv.org/abs/2008.04443>
  - **Why it matters here:** reviews record linkage, deduplication, clustering, canonicalization, and the practical difficulty of integrating records without unique identifiers.
  - **Claim it would support:** “Entity resolution is the process of deciding which records refer to the same domain entity under a stated policy.”
  - **Notes file:** [03-entity-resolution-and-relation-extraction.md](../research/knowledge-graphs/03-entity-resolution-and-relation-extraction.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Matt Barnes (2015). *A Practitioner's Guide to Evaluating Entity Resolution Results*. arXiv:1509.04238. <https://arxiv.org/abs/1509.04238>
  - **Why it matters here:** surveys entity-resolution evaluation metrics and warns that rankings can conflict, supporting consequence-specific evaluation rather than a single match score.
  - **Claim it would support:** “Evaluate by entity type and consequence.”
  - **Notes file:** [03-entity-resolution-and-relation-extraction.md](../research/knowledge-graphs/03-entity-resolution-and-relation-extraction.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Rostislav Nedelchev, Debanjan Chaudhuri, Jens Lehmann, and Asja Fischer (2020). *End-to-End Entity Linking and Disambiguation leveraging Word and Knowledge Graph Embeddings*. arXiv:2002.11143. <https://arxiv.org/abs/2002.11143>
  - **Why it matters here:** studies entity linking as connecting mentions to graph entities and uses relational context for disambiguation, matching the curriculum's entity-linking and graph-neighborhood handoff.
  - **Claim it would support:** “The curriculum calls the question-side version of this process entity linking: find the nodes a question is about, then inspect a bounded k-hop subgraph.”
  - **Notes file:** [03-entity-resolution-and-relation-extraction.md](../research/knowledge-graphs/03-entity-resolution-and-relation-extraction.md)
  - **Miriah's notes:**

### ch08.04-knowledge-graph-tradeoffs

- [ ] **unreviewed** — Haoyu Han, Harry Shomer, Yu Wang, Yongjia Lei, Kai Guo, Zhigang Hua, Bo Long, Hui Liu, and Jiliang Tang (2025). *RAG vs. GraphRAG: A Systematic Evaluation and Key Insights*. arXiv:2502.11371. <https://arxiv.org/abs/2502.11371>
  - **Why it matters here:** compares RAG and GraphRAG across question answering and query-focused summarization, finding task-specific strengths and motivating selection or integration rather than a universal winner.
  - **Claim it would support:** “Benchmark the actual workload and failure modes” and the curriculum-aligned baseline comparison.
  - **Notes file:** [05-context-engineering-connections.md](../research/knowledge-graphs/05-context-engineering-connections.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — *Knowledge graph quality control: A survey* (2021). Future Internet. <https://doi.org/10.1016/j.fmre.2021.09.003>
  - **Why it matters here:** surveys quality dimensions and their differences, supporting the module's broader cost model beyond answer accuracy.
  - **Claim it would support:** “Graph adoption includes more than database licensing or query latency.”
  - **Notes file:** [05-context-engineering-connections.md](../research/knowledge-graphs/05-context-engineering-connections.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Philipp Cimiano and Heiko Paulheim (2016). *Knowledge Graph Refinement: A Survey of Approaches and Evaluation Methods*. Semantic Web. <https://doi.org/10.3233/SW-160218>
  - **Why it matters here:** frames refinement and evaluation as ongoing knowledge-graph operations, supporting reversible pilots, correction workflows, and rebuild requirements.
  - **Claim it would support:** “Inject source corrections and deletion, rebuild the projection, and verify replay.”
  - **Notes file:** [05-context-engineering-connections.md](../research/knowledge-graphs/05-context-engineering-connections.md)
  - **Miriah's notes:**

### ch08.05-instance-coverage-and-ontology-population

- [ ] **unreviewed** — Pascal Hitzler, Amrapali Zaveri, Anisa Rula, Andrea Maurino, Ricardo Pietrobon, Jens Lehmann, and Sören Auer (2016). *Quality Assessment for Linked Data: A Survey*. Semantic Web. <https://doi.org/10.3233/SW-150175>
  - **Why it matters here:** organizes quality dimensions including completeness, accuracy, consistency, timeliness, provenance, and accessibility, supporting the module's warning that population coverage is only one part of graph readiness.
  - **Claim it would support:** “Each denominator needs a version, tenant or domain scope, cutoff time, and eligibility rule.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Subhi Issa, Onaopepo Adekunle, Fayçal Hamdi, Samira Si-Said Cherfi, Michel Dumontier, and Amrapali Zaveri (2021). *Knowledge Graph Completeness: A Systematic Literature Review*. IEEE Access. <https://doi.org/10.1109/ACCESS.2021.3056622>
  - **Why it matters here:** surveys completeness as a distinct knowledge-graph quality dimension and supports treating denominators and completeness claims as explicit evaluation choices.
  - **Claim it would support:** “Coverage must be defined against an expected population or a task requirement.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Philipp Cimiano and Heiko Paulheim (2016). *Knowledge Graph Refinement: A Survey of Approaches and Evaluation Methods*. Semantic Web. <https://doi.org/10.3233/SW-160218>
  - **Why it matters here:** connects graph quality to refinement and evaluation workflows, supporting correction propagation, rebuilds, and competency-question testing.
  - **Claim it would support:** “Then test ... deletion and correction propagation; full rebuild equivalence.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**

### ch08.06-property-completeness-and-schema-quality

- [ ] **unreviewed** — Axel-Cyrille Ngonga Ngomo, Irini Fundulaki, Anastasia Krithara, Mohammad Rashid, Marco Torchiano, Giuseppe Rizzo, Nandana Mihindukulasooriya, and Oscar Corcho (2019). *A Quality Assessment Approach for Evolving Knowledge Bases*. Semantic Web. <https://doi.org/10.3233/SW-180324>
  - **Why it matters here:** distinguishes schema, property, population, and interlinking completeness and defines property completeness relative to a class and release, matching the module's requirement-specific denominator.
  - **Claim it would support:** “The denominator is ... the set required or expected for this entity under a versioned class, task, jurisdiction, lifecycle state, and source scope.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Subhi Issa, Onaopepo Adekunle, Fayçal Hamdi, Samira Si-Said Cherfi, Michel Dumontier, and Amrapali Zaveri (2021). *Knowledge Graph Completeness: A Systematic Literature Review*. IEEE Access. <https://doi.org/10.1109/ACCESS.2021.3056622>
  - **Why it matters here:** surveys completeness as a distinct quality dimension and supports keeping applicability, scope, and intended use explicit.
  - **Claim it would support:** “Property completeness is diagnostic, not a target to maximize.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv:2309.15217. <https://arxiv.org/abs/2309.15217>
  - **Why it matters here:** separates context relevance, answer relevance, and faithfulness, supporting the module's warning that graph property fill rates do not uniformly predict downstream retrieval or task quality.
  - **Claim it would support:** “Track downstream failures caused by missing or invalid properties rather than assuming completeness predicts performance uniformly.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**

### ch08.07-ontology-guided-information-extraction

- [ ] **unreviewed** — Raghu Anantharangachar, Srinivasan Ramani, and S. Rajagopalan (2013). *Ontology Guided Information Extraction from Unstructured Text*. arXiv:1302.1335. <https://arxiv.org/abs/1302.1335>
  - **Why it matters here:** describes populating an existing ontology from natural-language text and using ontology concepts to guide triple extraction, directly supporting the proposal-space argument.
  - **Claim it would support:** “An ontology can narrow that proposal space by defining the entities, relationships, roles, and temporal distinctions the application recognizes.”
  - **Notes file:** [07-informal-text-ie.md](../research/knowledge-graphs/07-informal-text-ie.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Daya C. Wimalasuriya and Dejing Dou (2010). *Ontology-based Information Extraction: An Introduction and a Survey of Current Approaches*. Journal of Information Science. <https://doi.org/10.1177/0165551509360123>
  - **Why it matters here:** surveys ontology-based IE architectures, implementation choices, and evaluation metrics, supporting the module's distinction between guidance and validation.
  - **Claim it would support:** “Ontology guidance can reduce invalid proposals, but model output must still cross deterministic checks.”
  - **Notes file:** [07-informal-text-ie.md](../research/knowledge-graphs/07-informal-text-ie.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Bowen Zhang and Harold Soh (2024). *Extract, Define, Canonicalize: An LLM-based Framework for Knowledge Graph Construction*. Proceedings of EMNLP. <https://aclanthology.org/2024.emnlp-main.548/>
  - **Why it matters here:** retrieves relevant schema elements and separates extraction, schema definition, and canonicalization, matching the module's task-specific contract and deliberate novelty handling.
  - **Claim it would support:** “The contract is smaller, testable, and versioned.”
  - **Notes file:** [07-informal-text-ie.md](../research/knowledge-graphs/07-informal-text-ie.md)
  - **Miriah's notes:**

### ch08.08-knowledge-extraction-methods

- [ ] **unreviewed** — Daya C. Wimalasuriya and Dejing Dou (2010). *Ontology-based Information Extraction: An Introduction and a Survey of Current Approaches*. Journal of Information Science. <https://doi.org/10.1177/0165551509360123>
  - **Why it matters here:** surveys IE architectures and evaluation choices, supporting method selection by task rather than a universal extraction ranking.
  - **Claim it would support:** “Knowledge-graph ingestion rarely needs one universal extractor.”
  - **Notes file:** [ie-ner.md](../research/knowledge-graphs/ie-ner.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Ying Lin, Heng Ji, Fei Huang, and Lingfei Wu (2020). *A Joint Neural Model for Information Extraction with Global Features*. Proceedings of ACL. <https://aclanthology.org/2020.acl-main.713/>
  - **Why it matters here:** provides a primary example of jointly modeling entities, triggers, and links, useful for comparing structured model outputs with simpler field-level methods.
  - **Claim it would support:** “These are tendencies, not universal precision/recall rankings.”
  - **Notes file:** [ie-ner.md](../research/knowledge-graphs/ie-ner.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Bowen Zhang and Harold Soh (2024). *Extract, Define, Canonicalize: An LLM-based Framework for Knowledge Graph Construction*. Proceedings of EMNLP. <https://aclanthology.org/2024.emnlp-main.548/>
  - **Why it matters here:** separates open extraction, schema definition, and canonicalization, and retrieves relevant schema elements for larger schemas.
  - **Claim it would support:** “The correct cascade depends on volume, latency, languages, review capacity, and the cost of false acceptance versus missed extraction.”
  - **Notes file:** [ie-ner.md](../research/knowledge-graphs/ie-ner.md)
  - **Miriah's notes:**

### ch08.09-guardrails-for-extraction-validation

- [ ] **unreviewed** — Sewon Min, Kalpesh Krishna, Xinxi Lyu, Mike Lewis, Wen-tau Yih, Pang Wei Koh, Mohit Iyyer, Luke Zettlemoyer, and Hannaneh Hajishirzi (2023). *FActScore: Fine-grained Atomic Evaluation of Factual Precision in Long Form Text Generation*. EMNLP. <https://arxiv.org/abs/2305.14251>
  - **Why it matters here:** decomposes generated content into atomic facts and evaluates support, supporting claim-level source entailment and the module's rejection of a single undifferentiated confidence score.
  - **Claim it would support:** “Source quality, extractor score, agreement, and validation results describe different uncertainties and should remain separate features.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Potsawee Manakul, Adian Liusie, and Mark J. F. Gales (2023). *SelfCheckGPT: Zero-Resource Black-Box Hallucination Detection for Generative Large Language Models*. EMNLP. <https://arxiv.org/abs/2303.08896>
  - **Why it matters here:** demonstrates a sampling-based consistency signal for black-box outputs, useful as one uncertainty feature but not a replacement for source and authority validation.
  - **Claim it would support:** “Do not trust a single confidence number.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv:2309.15217. <https://arxiv.org/abs/2309.15217>
  - **Why it matters here:** separates faithfulness, answer relevance, and context relevance, supporting the module's recommendation to measure validation and downstream escape as separate outcomes.
  - **Claim it would support:** “Measure false acceptance, false rejection, review yield, validator escape rate, correction latency, provenance coverage, and downstream use of unapproved state.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**

### ch08.10-multilingual-extraction-with-llms

- [ ] **unreviewed** — Junjie Hu, Sebastian Ruder, Aditya Siddhant, Graham Neubig, Orhan Firat, and Melvin Johnson (2020). *XTREME: A Massively Multilingual Multi-task Benchmark for Evaluating Cross-lingual Generalization*. arXiv:2003.11080. <https://arxiv.org/abs/2003.11080>
  - **Why it matters here:** evaluates cross-lingual generalization across typologically diverse languages and multiple tasks, supporting slice-based evaluation rather than aggregate multilingual claims.
  - **Claim it would support:** “Build reviewed sets for each supported language, script, and relevant language pair.”
  - **Notes file:** [language-models-few-shot-learners-notes.md](../research/language-models-few-shot-learners-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Jack FitzGerald, Christopher Hench, Charith Peris, Scott Mackie, Laurie Crist, Misha Britan, Wouter Leeuwis, Gokhan Tur, and Prem Natarajan et al. (2023). *MASSIVE: A 1M-Example Multilingual Natural Language Understanding Dataset with 51 Typologically-Diverse Languages*. Proceedings of ACL. <https://aclanthology.org/2023.acl-long.235/>
  - **Why it matters here:** provides a large multilingual slot-filling and intent benchmark with parallel labeled data, illustrating the reviewed cases and per-language metrics needed for language-aware routing.
  - **Claim it would support:** “Measure ... entity and relation performance by type and language” and preserve language-specific evaluation artifacts.
  - **Notes file:** [language-models-few-shot-learners-notes.md](../research/language-models-few-shot-learners-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Tyler A. Chang, Catherine Arnett, Zhuowen Tu, and Benjamin K. Bergen (2024). *When Is Multilinguality a Curse? Language Modeling for 250 High- and Low-Resource Languages*. Proceedings of EMNLP. <https://aclanthology.org/2024.emnlp-main.236/>
  - **Why it matters here:** measures how multilingual training affects languages differently and finds that adding languages can help some and hurt others, supporting bounded fallback and language-specific thresholds.
  - **Claim it would support:** “Aggregate accuracy can conceal ... language-specific failure modes” and “Languages below threshold need a safer fallback.”
  - **Notes file:** [language-models-few-shot-learners-notes.md](../research/language-models-few-shot-learners-notes.md)
  - **Miriah's notes:**

### ch08-knowledge-graphs-and-semantic-context

- [ ] **unreviewed** — Darren Edge, Ha Trinh, Newman Cheng, Joshua Bradley, Alex Chao, Apurva Mody, Steven Truitt, Dasha Metropolitansky, Robert Osazuwa Ness, and Jonathan Larson (2024, version 2 read 2026-09-18). *From Local to Global: A Graph RAG Approach to Query-Focused Summarization*. arXiv:2404.16130. <https://arxiv.org/abs/2404.16130>
  - **Why it matters here:** the curriculum's GraphRAG reading supplies primary evidence for testing relationship-oriented context against a strong retrieval baseline rather than assuming a graph helps.
  - **Claim it would support:** “Knowledge graphs are particularly useful when relationship-centric questions ... justify their operational cost.”
  - **Notes file:** [knowledge-graphs-km-thesis.md](../research/knowledge-graphs-km-thesis.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Lingfeng Zhong, Jia Wu, Qian Li, Hao Peng, and Xindong Wu (2023). *A Comprehensive Survey on Automatic Knowledge Graph Construction*. arXiv:2302.05019. <https://arxiv.org/abs/2302.05019>
  - **Why it matters here:** organizes graph construction as acquisition, refinement, and evolution, matching the chapter's semantic infrastructure and lifecycle framing.
  - **Claim it would support:** “The chapter moves from semantic design to operational use” and the distinction between semantic context and graph storage.
  - **Notes file:** [knowledge-graphs-km-thesis.md](../research/knowledge-graphs-km-thesis.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — W3C Provenance Working Group (2013). *Semantics of the PROV Data Model*. W3C Working Group Note. <https://www.w3.org/TR/prov-sem/>
  - **Why it matters here:** gives a primary standards reference for representing derivation and history, central to the chapter's downstream contract for semantic context.
  - **Claim it would support:** “A semantic representation should therefore expose ... the evidence, identity, time interval, authority, uncertainty, and scope needed to decide.”
  - **Notes file:** [knowledge-graphs-km-thesis.md](../research/knowledge-graphs-km-thesis.md)
  - **Miriah's notes:**

### ch08.01-schemas-taxonomies-and-ontologies

- [ ] **unreviewed** — W3C OWL Working Group (2012). *OWL 2 Web Ontology Language Primer (Second Edition)*. W3C Recommendation. <https://www.w3.org/TR/owl2-primer/>
  - **Why it matters here:** supplies the standards definition of ontology language, classes, properties, restrictions, and reasoning used by the module's terminology boundary.
  - **Claim it would support:** “An ontology states the concepts and relationships a domain recognizes, plus axioms that support shared interpretation or inference.”
  - **Notes file:** [02-semantics-and-ontologies.md](../research/knowledge-graphs/02-semantics-and-ontologies.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — W3C RDF Data Shapes Working Group (2017). *Shapes Constraint Language (SHACL)*. W3C Recommendation. <https://www.w3.org/TR/shacl/>
  - **Why it matters here:** provides the primary validation specification that separates RDF shape conformance from ontology entailment.
  - **Claim it would support:** “Closed-world application requirements ... need a validation boundary.”
  - **Notes file:** [02-semantics-and-ontologies.md](../research/knowledge-graphs/02-semantics-and-ontologies.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — C. Maria Keet and Paolo D. D. Ferrario (2013). *Evaluating Ontologies with Competency Questions*. IEEE/WIC/ACM International Joint Conferences on Web Intelligence and Intelligent Agent Technology. <https://doi.org/10.1109/WI-IAT.2013.199>
  - **Why it matters here:** grounds competency questions as a way to state ontology requirements and evaluate whether a model supports its intended use.
  - **Claim it would support:** “Reliable domain models grow from competency questions and invariants.”
  - **Notes file:** [02-semantics-and-ontologies.md](../research/knowledge-graphs/02-semantics-and-ontologies.md)
  - **Miriah's notes:**

### ch09.01-lexical-and-relational-retrieval

- [ ] **unreviewed** — Stephen E. Robertson and Hugo Zaragoza (2009). *The Probabilistic Relevance Framework: BM25 and Beyond*. Foundations and Trends in Information Retrieval. <https://doi.org/10.1561/1500000019>
  - **Why it matters here:** supplies a primary account of BM25's probabilistic assumptions for the module's bounded explanation of term frequency, document frequency, and length normalization.
  - **Claim it would support:** “BM25 ranks documents using term frequency, document frequency, and length normalization.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Gordon V. Cormack, Charles L. A. Clarke, and Stefan Büttcher (2009). *Reciprocal Rank Fusion Outperforms Condorcet and Individual Rank Learning Methods*. Proceedings of SIGIR. <https://doi.org/10.1145/1571941.1572114>
  - **Why it matters here:** provides primary evidence for rank fusion, supporting the distinction between candidate generation and combining result lists after retrieval.
  - **Claim it would support:** “Scores from different systems are not directly comparable, so fusion should be an explicit design step.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — *COIL: Revisit Exact Lexical Match in Information Retrieval with Contextualized Inverted List* (2021). Proceedings of NAACL. <https://aclanthology.org/2021.naacl-main.241/>
  - **Why it matters here:** revisits exact lexical matching with contextualized representations, supporting the module's warning that lexical behavior depends on analyzers, fields, and scoring design.
  - **Claim it would support:** “Lexical search behavior depends on analyzers, query expansion, fields, and corpus statistics.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**

### ch09.02-vector-and-semantic-retrieval

- [ ] **unreviewed** — Patrick Lewis, Ethan Perez, Aleksandara Piktus, Fabio Petroni, Vladimir Karpukhin, Naman Goyal, Heinrich Küttler, Mike Lewis, Wen-tau Yih, Tim Rocktäschel, Sebastian Riedel, and Douwe Kiela (2020). *Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks*. arXiv:2005.11401. <https://arxiv.org/abs/2005.11401>
  - **Why it matters here:** defines retrieval-augmented generation as a retriever-plus-generator system and supports treating retrieved passages as context rather than authority.
  - **Claim it would support:** “Retrieval is a candidate-generation stage ... not a truth-making stage.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Vladimir Karpukhin, Barlas Oğuz, Sewon Min, Patrick Lewis, Ledell Wu, Sergey Edunov, Danqi Chen, and Wen-tau Yih (2020). *Dense Passage Retrieval for Open-Domain Question Answering*. arXiv:2004.04906. <https://arxiv.org/abs/2004.04906>
  - **Why it matters here:** provides a primary dual-encoder dense-retrieval account and evaluates dense candidates against BM25.
  - **Claim it would support:** “Dense retrieval maps queries and content into a learned representation and ranks by proximity under a chosen similarity function.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Nils Reimers and Iryna Gurevych (2019). *Sentence-BERT: Sentence Embeddings using Siamese BERT-Networks*. Proceedings of EMNLP. <https://arxiv.org/abs/1908.10084>
  - **Why it matters here:** grounds reusable sentence embeddings and cosine-based semantic comparison.
  - **Claim it would support:** “For L2-normalized vectors, a dot product is cosine similarity.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yu. A. Malkov and D. A. Yashunin (2018). *Efficient and robust approximate nearest neighbor search using Hierarchical Navigable Small World graphs*. IEEE Transactions on Pattern Analysis and Machine Intelligence. <https://arxiv.org/abs/1603.09320>
  - **Why it matters here:** supplies a primary ANN index account for the latency/accuracy and parameter-tuning tradeoff.
  - **Claim it would support:** “Approximate nearest-neighbor search trades retrieval accuracy for latency and memory.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**

### ch09.03-graph-and-hybrid-retrieval

- [ ] **unreviewed** — Yuntong Hu, Zhihan Lei, Zheng Zhang, Bo Pan, Chen Ling, and Liang Zhao (2025). *GRAG: Graph Retrieval-Augmented Generation*. Findings of NAACL. <https://aclanthology.org/2025.findings-naacl.232/>
  - **Why it matters here:** studies textual subgraph retrieval and graph-aware context for networked documents, supporting relationship-oriented multi-hop retrieval.
  - **Claim it would support:** “Use graph paths when they are the answer or route to evidence.”
  - **Notes file:** [knowledge-graphs-km-thesis.md](../research/knowledge-graphs-km-thesis.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Bernal Jiménez Gutiérrez, Yiheng Shu, Yu Gu, Michihiro Yasunaga, and Yu Su (2024). *HippoRAG: Neurobiologically Inspired Long-Term Memory for Large Language Models*. NeurIPS. <https://arxiv.org/abs/2405.14831>
  - **Why it matters here:** combines a knowledge graph with Personalized PageRank and compares graph-supported retrieval with iterative retrieval, supporting the cost and multi-hop tradeoff framing.
  - **Claim it would support:** “A hybrid pipeline is justified only when its gain survives the operational cost and failure injection.”
  - **Notes file:** [knowledge-graphs-km-thesis.md](../research/knowledge-graphs-km-thesis.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Rishi Kalra, Zekun Wu, Ayesha Gulley, Airlie Hilliard, Xin Guan, Adriano Koshiyama, and Philip Colin Treleaven (2024). *HyPA-RAG: A Hybrid Parameter Adaptive Retrieval-Augmented Generation System for AI Legal and Policy Applications*. CustomNLP4U. <https://aclanthology.org/2024.customnlp4u-1.18/>
  - **Why it matters here:** evaluates adaptive combinations of dense, sparse, and knowledge-graph retrieval, supporting planned hybrid routing rather than unbounded fan-out.
  - **Claim it would support:** “Running relational, lexical, dense, and graph retrieval for every request increases latency and candidate noise.”
  - **Notes file:** [knowledge-graphs-km-thesis.md](../research/knowledge-graphs-km-thesis.md)
  - **Miriah's notes:**

### ch09.04-ranking-reranking-and-query-planning

- [ ] **unreviewed** — Rodrigo Nogueira and Kyunghyun Cho (2019). *Passage Re-ranking with BERT*. arXiv:1901.04085. <https://arxiv.org/abs/1901.04085>
  - **Why it matters here:** provides a primary passage-reranking account, supporting the separation of cheap candidate generation from richer query–document scoring.
  - **Claim it would support:** “Reranking spends more computation on a smaller set using richer query–document interaction.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Liang Wang, Nan Yang, and Furu Wei (2023). *Query2doc: Query Expansion with Large Language Models*. arXiv:2303.07678. <https://arxiv.org/abs/2303.07678>
  - **Why it matters here:** evaluates LLM-generated pseudo-document expansion for sparse and dense retrieval, supporting a measured multi-query rung.
  - **Claim it would support:** “Diverse query expansion can improve recall but can also introduce drift.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Alistair Moffat (2022). *Batch Evaluation Metrics in Information Retrieval: Measures, Scales, and Meaning*. arXiv:2207.03103. <https://arxiv.org/abs/2207.03103>
  - **Why it matters here:** examines reciprocal rank and other ranking metrics, supporting the warning that no single ranking metric captures system quality.
  - **Claim it would support:** “Measure selection accuracy, required-slot satisfaction, ranking metrics, latency, cost, and task outcome.”
  - **Notes file:** [hybrid-retrieval-architectures.md](../research/hybrid-retrieval-architectures.md)
  - **Miriah's notes:**

### ch09.05-context-precision-and-context-recall

- [ ] **unreviewed** — Nandan Thakur, Nils Reimers, Andreas Rücklé, Abhishek Srivastava, and Iryna Gurevych (2021, arXiv v4 read 2026-09-18). *BEIR: A Heterogenous Benchmark for Zero-shot Evaluation of Information Retrieval Models*. NeurIPS. <https://arxiv.org/abs/2104.08663>
  - **Why it matters here:** evaluates multiple retrieval families across heterogeneous tasks, supporting bounded judged corpora and cross-task caution.
  - **Claim it would support:** “Estimate recall against a bounded judged corpus ... and report it as such.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv:2309.15217. <https://arxiv.org/abs/2309.15217>
  - **Why it matters here:** separates context relevance, answer relevance, and faithfulness, supporting distinct retrieval and answer-support measurements.
  - **Claim it would support:** “Retrieval, authorization, freshness, source truth, tool selection, model use, and task success are different boundaries.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Sewon Min, Kalpesh Krishna, Xinxi Lyu, Mike Lewis, Wen-tau Yih, Pang Wei Koh, Mohit Iyyer, Luke Zettlemoyer, and Hannaneh Hajishirzi (2023). *FActScore: Fine-grained Atomic Evaluation of Factual Precision in Long Form Text Generation*. EMNLP. <https://arxiv.org/abs/2305.14251>
  - **Why it matters here:** evaluates atomic claim support after generation, supporting claim-to-source measurement.
  - **Claim it would support:** “Measure ... model answer with claim-to-source support.”
  - **Notes file:** [kg-quality-metrics-notes.md](../research/kg-quality-metrics-notes.md)
  - **Miriah's notes:**

### ch10.00-guardrails-and-ontology-based-validation

- [ ] **unreviewed** — Saibo Geng, Martin Josifoski, Maxime Peyrard, and Robert West (2023). *Grammar-Constrained Decoding for Structured NLP Tasks without Finetuning*. arXiv:2305.13971. <https://arxiv.org/abs/2305.13971>
  - **Why it matters here:** supports structural validity during decoding as distinct from semantic, authorization, and policy checks.
  - **Claim it would support:** “Engineered context can give a model better evidence ... Validation narrows recognized structure and policy permits or denies effects.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Eric Wallace, Kai Xiao, Reimar Leike, Lilian Weng, Johannes Heidecke, and Alex Beutel (2024). *The Instruction Hierarchy: Training LLMs to Prioritize Privileged Instructions*. arXiv:2404.13208. <https://arxiv.org/abs/2404.13208>
  - **Why it matters here:** evaluates prompt-injection robustness and over-refusal under instruction priorities.
  - **Claim it would support:** “Post-generation filtering alone is insufficient.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yen-Shan Chen, Sian-Yao Huang, Cheng-Lin Yang, and Yun-Nung Chen (2026). *TraceSafe: A Systematic Assessment of LLM Guardrails on Multi-Step Tool-Calling Trajectories*. arXiv:2604.07223. <https://arxiv.org/abs/2604.07223>
  - **Why it matters here:** evaluates guardrails over intermediate tool-use trajectories, supporting checks before side effects.
  - **Claim it would support:** “Test ... unauthorized paths, prompt-injected instructions, and validator outages.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**

### ch10.01-personalization-as-retrieval

- [ ] **unreviewed** — Charles Packer, Sarah Wooders, Kevin Lin, Vivian Fang, Shishir G. Patil, Ion Stoica, and Joseph E. Gonzalez (2023). *MemGPT: Towards LLMs as Operating Systems*. arXiv:2310.08560. <https://arxiv.org/abs/2310.08560>
  - **Why it matters here:** describes virtual context management across memory tiers, supporting the separation of working context from longer-lived state.
  - **Claim it would support:** “A model has no memory of its own. Everything it remembers is something a harness chose to put back in front of it.”
  - **Notes file:** [episodic-periodic-memory.md](../research/episodic-periodic-memory.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Joon Sung Park, Joseph C. O'Brien, Carrie J. Cai, Meredith Ringel Morris, Percy Liang, and Michael S. Bernstein (2023). *Generative Agents: Interactive Simulacra of Human Behavior*. arXiv:2304.03442. <https://arxiv.org/abs/2304.03442>
  - **Why it matters here:** combines experience records, reflection, and dynamic retrieval, supporting the distinction between episodic traces and semantic memory.
  - **Claim it would support:** “Semantic memory is distilled facts ... recalled by similarity.”
  - **Notes file:** [episodic-periodic-memory.md](../research/episodic-periodic-memory.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Weizhi Wang, Li Dong, Hao Cheng, Xiaodong Liu, Xifeng Yan, Jianfeng Gao, and Furu Wei (2023). *Augmenting Language Models with Long-Term Memory*. arXiv:2306.07174. <https://arxiv.org/abs/2306.07174>
  - **Why it matters here:** presents a long-term memory retriever and reader for context beyond the model window, supporting lifecycle and staleness considerations.
  - **Claim it would support:** “Production memory is per user, scoped, and forgets on purpose.”
  - **Notes file:** [episodic-periodic-memory.md](../research/episodic-periodic-memory.md)
  - **Miriah's notes:**

### ch10.02-scoped-hydration

- [ ] **unreviewed** — Pengcheng Zhou, Yinglun Feng, and Zhongliang Yang (2025). *Provably Secure Retrieval-Augmented Generation*. arXiv:2508.01084. <https://arxiv.org/abs/2508.01084>
  - **Why it matters here:** addresses authorization and confidentiality for retrieved content and vector embeddings.
  - **Claim it would support:** “Post-filtering is insufficient when unauthorized candidates influence ranking, counts, snippets, graph expansion, cache keys, or timing.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Yining Chen, Jihao Zhao, Bo Tang, Haofen Wang, Feiyu Xiong, and Zhiyu Li (2026). *MemPrivacy: Privacy-Preserving Personalized Memory Management for Edge-Cloud Agents*. arXiv:2605.09530. <https://arxiv.org/abs/2605.09530>
  - **Why it matters here:** evaluates privacy-preserving personalized memory management and utility-versus-disclosure tradeoffs.
  - **Claim it would support:** “Cache entries need principal- and policy-aware keys or content whose reuse is safe across scopes.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Jeff Z. Pan et al. (2023). *Large Language Models and Knowledge Graphs*. arXiv:2308.06374. <https://arxiv.org/abs/2308.06374>
  - **Why it matters here:** discusses privacy and policy concerns when knowledge and personal data are integrated with LLMs.
  - **Claim it would support:** “A visible endpoint does not imply a visible relationship.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**

### ch10.03-provenance-and-derived-context

- [ ] **unreviewed** — W3C Provenance Working Group (2013). *PROV-DM: The PROV Data Model*. W3C Recommendation. <https://www.w3.org/TR/prov-dm/>
  - **Why it matters here:** provides a standards vocabulary for entities, activities, agents, derivation, and attribution.
  - **Claim it would support:** “For a retrieved item, record a stable source locator ... transformations, and final context item ID.”
  - **Notes file:** [provenance-notes.md](../research/provenance-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Sewon Min, Kalpesh Krishna, Xinxi Lyu, Mike Lewis, Wen-tau Yih, Pang Wei Koh, Mohit Iyyer, Luke Zettlemoyer, and Hannaneh Hajishirzi (2023). *FActScore: Fine-grained Atomic Evaluation of Factual Precision in Long Form Text Generation*. EMNLP. <https://arxiv.org/abs/2305.14251>
  - **Why it matters here:** supports claim-level support checks after transformation rather than treating a provenance link as proof.
  - **Claim it would support:** “This is lineage, not proof.”
  - **Notes file:** [provenance-notes.md](../research/provenance-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv:2309.15217. <https://arxiv.org/abs/2309.15217>
  - **Why it matters here:** separates faithfulness and context relevance, supporting distinct evidence-path and answer evaluations.
  - **Claim it would support:** “Measure lineage completeness ... and unsupported claims whose evidence path is absent.”
  - **Notes file:** [provenance-notes.md](../research/provenance-notes.md)
  - **Miriah's notes:**

### ch10.04-policy-aware-user-context

- [ ] **unreviewed** — Yining Chen, Jihao Zhao, Bo Tang, Haofen Wang, Feiyu Xiong, and Zhiyu Li (2026). *MemPrivacy: Privacy-Preserving Personalized Memory Management for Edge-Cloud Agents*. arXiv:2605.09530. <https://arxiv.org/abs/2605.09530>
  - **Why it matters here:** evaluates privacy-preserving personalized memory and utility loss.
  - **Claim it would support:** “Minimize active context and durable retention separately.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Eric Wallace, Kai Xiao, Reimar Leike, Lilian Weng, Johannes Heidecke, and Alex Beutel (2024). *The Instruction Hierarchy: Training LLMs to Prioritize Privileged Instructions*. arXiv:2404.13208. <https://arxiv.org/abs/2404.13208>
  - **Why it matters here:** evaluates prompt-injection robustness and over-refusal.
  - **Claim it would support:** “Do not let the model supply trusted values for those fields.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Pengcheng Zhou, Yinglun Feng, and Zhongliang Yang (2025). *Provably Secure Retrieval-Augmented Generation*. arXiv:2508.01084. <https://arxiv.org/abs/2508.01084>
  - **Why it matters here:** addresses authorization and confidentiality for retrieved content.
  - **Claim it would support:** “A trusted policy decision point must derive or validate it.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**

### ch10.05-provenance-coverage-metrics

- [ ] **unreviewed** — W3C Provenance Working Group (2013). *PROV-DM: The PROV Data Model*. W3C Recommendation. <https://www.w3.org/TR/prov-dm/>
  - **Why it matters here:** defines provenance entities, activities, agents, derivation, and attribution for item-type contracts.
  - **Claim it would support:** “Define a provenance contract per context-item type.”
  - **Notes file:** [provenance-notes.md](../research/provenance-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Sewon Min, Kalpesh Krishna, Xinxi Lyu, Mike Lewis, Wen-tau Yih, Pang Wei Koh, Mohit Iyyer, Luke Zettlemoyer, and Hannaneh Hajishirzi (2023). *FActScore: Fine-grained Atomic Evaluation of Factual Precision in Long Form Text Generation*. EMNLP. <https://arxiv.org/abs/2305.14251>
  - **Why it matters here:** provides claim-level support measurement beyond citation presence.
  - **Claim it would support:** “Also measure claim-support coverage for generated outputs.”
  - **Notes file:** [provenance-notes.md](../research/provenance-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv:2309.15217. <https://arxiv.org/abs/2309.15217>
  - **Why it matters here:** separates faithfulness, answer relevance, and context relevance for distinct release metrics.
  - **Claim it would support:** “It does not measure source truth, semantic correctness, authorization correctness, relevance, or final-answer correctness.”
  - **Notes file:** [provenance-notes.md](../research/provenance-notes.md)
  - **Miriah's notes:**

### ch11.01-least-privilege

- [ ] **unreviewed** — Jerome H. Saltzer and Michael D. Schroeder (1975). *The Protection of Information in Computer Systems*. Proceedings of the IEEE. <https://doi.org/10.1109/PROC.1975.9939>
  - **Why it matters here:** provides the primary account of least privilege, fail-safe defaults, complete mediation, and separation of privilege.
  - **Claim it would support:** “A language-system component should receive only the authority required for the current task.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Norman Hardy (1988). *The Confused Deputy (or Why Capabilities Might Have Been Invented)*. ACM SIGOPS Operating Systems Review. <https://doi.org/10.1145/54289.871709>
  - **Why it matters here:** defines the confused-deputy failure that motivates binding authority to the intended principal and operation.
  - **Claim it would support:** “Test confused-deputy requests.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Paulius Rauba, Dominykas Seputis, Patrikas Vanagas, and Mihaela van der Schaar (2026). *No More, No Less: Least-Privilege Language Models*. arXiv:2601.23157. <https://arxiv.org/abs/2601.23157>
  - **Why it matters here:** applies least-privilege reasoning to language-model deployments while separating model controls from resource enforcement.
  - **Claim it would support:** “Prompt instructions can help the model choose correctly but do not enforce authority.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**

### ch11.02-rbac-abac-capability-based-access

- [ ] **unreviewed** — Norman Hardy (1988). *The Confused Deputy (or Why Capabilities Might Have Been Invented)*. ACM SIGOPS Operating Systems Review. <https://doi.org/10.1145/54289.871709>
  - **Why it matters here:** provides the primary confused-deputy account for separating a caller's authority from an intermediary's broader authority.
  - **Claim it would support:** “A trusted host, broker, or proxy performs the protocol and attaches credentials to outbound calls.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Mark S. Miller, Ka-Ping Yee, and Jonathan Shapiro (2003). *Capability Myths Demolished*. Technical report. <https://srl.cs.jhu.edu/pubs/SRL2003-02.pdf>
  - **Why it matters here:** examines capability semantics and common misconceptions.
  - **Claim it would support:** “JWT is only a token format, not proof that a token is least-privilege or capability-safe.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Zhiyuan Li, Jingzheng Wu, Yuhao Peng, Tianyue Luo, Xing Cui, and Xiang Ling (2026). *Confused Deputy Attack Against Model Context Protocol*. ACM Transactions on Software Engineering and Methodology. <https://doi.org/10.1145/3830467>
  - **Why it matters here:** applies confused-deputy analysis to a contemporary tool protocol.
  - **Claim it would support:** “Enforcement belongs at the resource server, tool host, database, or trusted proxy.”
  - **Notes file:** [guardrails-notes.md](../research/guardrails-notes.md)
  - **Miriah's notes:**

### ch13.03-durable-and-event-driven-execution

- [ ] **unreviewed** — ZenML (2026). *Kitaru: An open-source, durable execution platform for long-running Python agents*. GitHub repository. <https://github.com/zenml-io/kitaru>
  - **Why it matters here:** provides a current agent-specific implementation example for checkpointing, replay, waits, and artifacts; relevant to “Durable Execution” without serving as general distributed-systems evidence.
  - **Claim it would support:** “The durable boundary can sit around agent actions such as model calls, tool calls, human waits, and artifact saves.”
  - **Notes file:** [durable-workflow-notes.md](../research/durable-workflow-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Hector Garcia-Molina and Kenneth Salem (1987). *Sagas*. Princeton University technical report. <https://www.cs.princeton.edu/techreports/1987/070.pdf>
  - **Why it matters here:** introduces long-lived transactions and compensating actions; relevant to “Choose Persistence Semantics Deliberately”.
  - **Claim it would support:** “Compensation is another fallible business action, not a database rollback.”
  - **Notes file:** [durable-workflow-notes.md](../research/durable-workflow-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Zhenchao Zhuang et al. (2023). *ExoFlow: A Universal Workflow System for Exactly-Once DAGs*. 17th USENIX Symposium on Operating Systems Design and Implementation. <https://www.usenix.org/system/files/osdi23-zhuang.pdf>
  - **Why it matters here:** studies recovery and execution trade-offs for workflow DAGs; relevant to “Durable Execution”.
  - **Claim it would support:** “Exactly-once execution is rarely an end-to-end property a workflow can simply declare.”
  - **Notes file:** [durable-workflow-notes.md](../research/durable-workflow-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Akshat Verma et al. (2015). *Deterministic Replay: A Survey*. ACM Computing Surveys. <https://doi.org/10.1145/2790077>
  - **Why it matters here:** surveys deterministic replay scope and trade-offs; relevant to “Replay Is Evidence, Not Time Travel”.
  - **Claim it would support:** “Deterministic replay can reconstruct state transitions ... but cannot generally reproduce a nondeterministic model response or the past state of an external service byte for byte.”
  - **Notes file:** [durable-workflow-notes.md](../research/durable-workflow-notes.md)
  - **Miriah's notes:**

### ch13.04-loops-retries-and-bounded-autonomy

- [ ] **unreviewed** — Aman Madaan et al. (2023). *Self-Refine: Iterative Refinement with Self-Feedback*. arXiv. <https://arxiv.org/abs/2303.17651>
  - **Why it matters here:** evaluates iterative generation, feedback, and revision at test time; relevant to “Define Progress Before Repeating”.
  - **Claim it would support:** “A new paragraph of model output is not necessarily progress”; iterative refinement needs feedback and task-specific evaluation.
  - **Notes file:** [retry-and-conversation-restoration-notes.md](../research/retry-and-conversation-restoration-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Alejandro Forero Cuervo (ed.). *Handling Overload*. Google SRE Book. <https://sre.google/sre-book/handling-overload/>
  - **Why it matters here:** documents overload-aware retry budgets and stopping retries when broader service capacity is exhausted; relevant to “Retry Failures, Not Business Decisions”.
  - **Claim it would support:** “A retry policy needs a per-request budget and must sometimes let failure reach the caller.”
  - **Notes file:** [retry-and-conversation-restoration-notes.md](../research/retry-and-conversation-restoration-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Dan Sandler (ed.). *Addressing Cascading Failures*. Google SRE Book. <https://sre.google/sre-book/addressing-cascading-failures/>
  - **Why it matters here:** explains retry amplification and randomized exponential backoff; relevant to the retry taxonomy and dependency-outage boundary.
  - **Claim it would support:** “Retries can amplify the effects seen in server overload.”
  - **Notes file:** [retry-and-conversation-restoration-notes.md](../research/retry-and-conversation-restoration-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Kubernetes Authors. *Update a Deployment Without Downtime*. Kubernetes Documentation. <https://kubernetes.io/docs/tasks/run-application/update-deployment-rolling/>
  - **Why it matters here:** defines rolling replacement, readiness, rollout monitoring, and rollback behavior; relevant to “Restore Conversations Across Deploys”.
  - **Claim it would support:** “Old and new pods overlap during a rolling deployment, so durable state and version compatibility cannot live only in the worker process.”
  - **Notes file:** [retry-and-conversation-restoration-notes.md](../research/retry-and-conversation-restoration-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Vercel. *Execution Model and Durability*. eve documentation. <https://github.com/vercel/eve/blob/main/docs/concepts/execution-model-and-durability.mdx>
  - **Why it matters here:** gives a current implementation example of durable conversations transferring settled state across production deployments; relevant to restoring chats without treating worker memory as the source of truth.
  - **Claim it would support:** “A durable conversation can be handed to the deployment that accepts the next turn.”
  - **Notes file:** [retry-and-conversation-restoration-notes.md](../research/retry-and-conversation-restoration-notes.md)
  - **Miriah's notes:**

### ch14.01-token-and-context-economics

- [ ] **unreviewed** — Nelson F. Liu et al. (2023). *Lost in the Middle: How Language Models Use Long Contexts*. Transactions of the Association for Computational Linguistics. <https://arxiv.org/abs/2307.03172>
  - **Why it matters here:** measures how the position of relevant information affects long-context use; relevant to the distinction between fitting context and using it effectively.
  - **Claim it would support:** “A long context can fit within the advertised window while still increasing the chance that relevant evidence is overlooked or poorly positioned.”
  - **Notes file:** [token-and-context-economics-notes.md](../research/token-and-context-economics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — In Gim et al. (2024). *Prompt Cache: Modular Attention Reuse for Low-Latency Inference*. arXiv. <https://arxiv.org/abs/2311.04934>
  - **Why it matters here:** studies reusing attention states for repeated prompt segments; relevant to prefix stability, latency, and the economics of repeatedly sending stable context.
  - **Claim it would support:** “Keep stable prefixes stable when the serving stack can reuse them.”
  - **Notes file:** [token-and-context-economics-notes.md](../research/token-and-context-economics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Woosuk Kwon et al. (2023). *Efficient Memory Management for Large Language Model Serving with PagedAttention*. Proceedings of the 29th ACM Symposium on Operating Systems Principles. <https://doi.org/10.1145/3600006.3613165>
  - **Why it matters here:** connects sequence length and KV-cache management to serving memory and throughput; relevant to treating context length as an operational cost, not only a provider limit.
  - **Claim it would support:** “Context length is an operational cost involving serving memory and throughput.”
  - **Notes file:** [token-and-context-economics-notes.md](../research/token-and-context-economics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Patrick Lewis et al. (2020). *Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks*. Advances in Neural Information Processing Systems. <https://arxiv.org/abs/2005.11401>
  - **Why it matters here:** provides the primary retrieval-versus-parametric-memory framing; relevant to pricing retrieval and context assembly as part of the full pipeline.
  - **Claim it would support:** “Cost follows context from source acquisition and retrieval through model use and outcome.”
  - **Notes file:** [token-and-context-economics-notes.md](../research/token-and-context-economics-notes.md)
  - **Miriah's notes:**

### ch14.02-retrieval-tool-and-latency-costs

- [ ] **unreviewed** — Jeffrey Dean and Luiz André Barroso (2013). *The Tail at Scale*. Communications of the ACM. <https://doi.org/10.1145/2408776.2408794>
  - **Why it matters here:** analyzes tail latency in fan-out services and techniques for tolerating latency variability; relevant to “Allocate a Latency Budget”.
  - **Claim it would support:** “Averages hide the tail where timeouts and retries concentrate.”
  - **Notes file:** [retrieval-tool-latency-notes.md](../research/retrieval-tool-latency-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Vladimir Karpukhin et al. (2020). *Dense Passage Retrieval for Open-Domain Question Answering*. Proceedings of EMNLP 2020. <https://arxiv.org/abs/2004.04906>
  - **Why it matters here:** provides a primary dense-retrieval baseline and evaluates passage-retrieval accuracy; relevant to “Price Retrieval as a Lifecycle”.
  - **Claim it would support:** “A retrieval rung should be compared by evidence quality and operational cost, not assumed to be an unconditional improvement.”
  - **Notes file:** [retrieval-tool-latency-notes.md](../research/retrieval-tool-latency-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Gordon V. Cormack, Charles L. A. Clarke, and Stefan Büttcher (2009). *Reciprocal Rank Fusion Outperforms Condorcet and Individual Rank Learning Methods*. Proceedings of SIGIR 2009. <https://doi.org/10.1145/1571941.1572114>
  - **Why it matters here:** evaluates a simple rank-fusion method across information-retrieval systems; relevant to the cost/quality tradeoff of adding fusion before reranking.
  - **Claim it would support:** “Fusion is an additional retrieval stage whose quality and cost should be measured.”
  - **Notes file:** [retrieval-tool-latency-notes.md](../research/retrieval-tool-latency-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Omar Khattab and Matei Zaharia (2020). *ColBERT: Efficient and Effective Passage Search via Contextualized Late Interaction over BERT*. Proceedings of SIGIR 2020. <https://arxiv.org/abs/2004.12832>
  - **Why it matters here:** separates document encoding from query-time interaction to reduce reranking cost; relevant to “Price Retrieval as a Lifecycle”.
  - **Claim it would support:** “Reranking can improve candidate quality while adding a measurable query-time cost.”
  - **Notes file:** [retrieval-tool-latency-notes.md](../research/retrieval-tool-latency-notes.md)
  - **Miriah's notes:**

### ch14.03-one-shot-execution-loops-and-subagents

- [ ] **unreviewed** — Qingyun Wu et al. (2023). *AutoGen: Enabling Next-Gen LLM Applications via Multi-Agent Conversation*. arXiv. <https://arxiv.org/abs/2308.08155>
  - **Why it matters here:** describes a framework for agents that converse, use tools, and involve humans; relevant to distinguishing a coordination mechanism from a reliability guarantee.
  - **Claim it would support:** “A framework for multi-agent conversation does not itself guarantee bounded cost, valid handoffs, or safe effects.”
  - **Notes file:** [one-shot-loops-delegation-notes.md](../research/one-shot-loops-delegation-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Chen Qian et al. (2023). *ChatDev: Communicative Agents for Software Development*. arXiv. <https://arxiv.org/abs/2307.07924>
  - **Why it matters here:** presents role-based multi-agent workflows with staged communication; relevant to typed handoffs, phase boundaries, and the cost of coordinating workers.
  - **Claim it would support:** “Delegation turns context into an interface between specialized workers.”
  - **Notes file:** [one-shot-loops-delegation-notes.md](../research/one-shot-loops-delegation-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Junlin Wang et al. (2024). *Mixture-of-Agents Enhances Large Language Model Capabilities*. International Conference on Learning Representations. <https://arxiv.org/abs/2406.04692>
  - **Why it matters here:** evaluates layered agents whose outputs become context for later agents; relevant to measuring whether added calls improve outcomes enough to justify multiplied context and inference cost.
  - **Claim it would support:** “Extra calls are justified only by a measurable reduction in failure or recovery cost.”
  - **Notes file:** [one-shot-loops-delegation-notes.md](../research/one-shot-loops-delegation-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Anthropic. *How we built our multi-agent research system*. Anthropic Engineering. <https://www.anthropic.com/engineering/multi-agent-research-system>
  - **Why it matters here:** reports a production orchestrator-worker pattern with parallel subagents, retries, and checkpoints; relevant as engineering evidence for the benefits and operational controls of delegation.
  - **Claim it would support:** “Parallel delegation can reduce wall-clock latency while increasing total compute and coordination obligations.”
  - **Notes file:** [one-shot-loops-delegation-notes.md](../research/one-shot-loops-delegation-notes.md)
  - **Miriah's notes:**

### ch14.04-local-models-and-model-routing

- [ ] **unreviewed** — Tim Dettmers et al. (2022). *LLM.int8(): 8-bit Matrix Multiplication for Transformers at Scale*. arXiv. <https://arxiv.org/abs/2208.07339>
  - **Why it matters here:** demonstrates a way to reduce inference memory while retaining model performance under a specified quantization procedure; relevant to “Compare Total Cost Under Load”.
  - **Claim it would support:** “Quantization changes the resource curve; it does not remove the quality test.”
  - **Notes file:** [local-models-routing-notes.md](../research/local-models-routing-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Guangxuan Xiao et al. (2022). *SmoothQuant: Accurate and Efficient Post-Training Quantization for Large Language Models*. arXiv. <https://arxiv.org/abs/2211.10438>
  - **Why it matters here:** studies memory and speed effects of post-training quantization; relevant to the local-serving resource tradeoff.
  - **Claim it would support:** “A lower-memory checkpoint is useful only when its deployed quality and hardware behavior remain acceptable.”
  - **Notes file:** [local-models-routing-notes.md](../research/local-models-routing-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Lingjiao Chen, Matei Zaharia, and James Zou (2023). *FrugalGPT: How to Use Large Language Models While Reducing Cost and Improving Performance*. arXiv. <https://arxiv.org/abs/2305.05176>
  - **Why it matters here:** formalizes prompt adaptation, approximation, and cascades as cost/quality strategies; relevant to “Route With Observable Policy”.
  - **Claim it would support:** “A cascade can reduce cost only when its acceptance and escalation behavior are measured.”
  - **Notes file:** [local-models-routing-notes.md](../research/local-models-routing-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Isaac Ong et al. (2024). *RouteLLM: Learning to Route LLMs with Preference Data*. arXiv (version 4, 2025). <https://arxiv.org/abs/2406.18665>
  - **Why it matters here:** evaluates learned routers for choosing between stronger and weaker models under quality/cost tradeoffs; relevant to calibration, route-level evaluation, and routing regret.
  - **Claim it would support:** “A router must be evaluated on quality and cost across the task distribution, not trusted from a confidence score alone.”
  - **Notes file:** [local-models-routing-notes.md](../research/local-models-routing-notes.md)
  - **Miriah's notes:**

### ch14.05-context-efficiency-metrics

- [ ] **unreviewed** — Percy Liang et al. (2022). *Holistic Evaluation of Language Models*. arXiv. <https://arxiv.org/abs/2211.09110>
  - **Why it matters here:** defines scenario-based, multi-metric evaluation and exposes trade-offs beyond accuracy; relevant to “Compare Policies on a Frontier”.
  - **Claim it would support:** “Context efficiency should not be reported as a universal scalar.”
  - **Notes file:** [context-efficiency-metrics-notes.md](../research/context-efficiency-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shahul Es, Jithin James, Luis Espinosa-Anke, and Steven Schockaert (2023). *RAGAS: Automated Evaluation of Retrieval Augmented Generation*. arXiv. <https://arxiv.org/abs/2309.15217>
  - **Why it matters here:** separates retrieval context quality, faithfulness, and generation quality; relevant to “Define the Outcome First”.
  - **Claim it would support:** “A context policy can change retrieval quality and answer quality in different directions, so measure them separately.”
  - **Notes file:** [context-efficiency-metrics-notes.md](../research/context-efficiency-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Xiao Liu et al. (2023). *AgentBench: Evaluating LLMs as Agents*. arXiv. <https://arxiv.org/abs/2308.03688>
  - **Why it matters here:** evaluates agents across multiple interactive environments rather than a single answer score; relevant to task slices, trajectory outcomes, and representative evaluation sets.
  - **Claim it would support:** “Run experiments on versioned representative sets and important slices.”
  - **Notes file:** [context-efficiency-metrics-notes.md](../research/context-efficiency-metrics-notes.md)
  - **Miriah's notes:**
- [ ] **unreviewed** — Shunyu Yao, Noah Shinn, Pedram Razavi, and Karthik Narasimhan (2024). *τ-bench: A Benchmark for Tool-Agent-User Interaction in Real-World Domains*. arXiv. <https://arxiv.org/abs/2406.12045>
  - **Why it matters here:** evaluates policy-following tool interactions against terminal database state and introduces pass^k; relevant to reliability, repeated-run variation, and terminal-outcome measurement.
  - **Claim it would support:** “A lower bill can disguise a reliability regression unless repeated-run and terminal-outcome measures are retained.”
  - **Notes file:** [context-efficiency-metrics-notes.md](../research/context-efficiency-metrics-notes.md)
  - **Miriah's notes:**

## Deferred / out of scope

- No chapter or module prose is edited during bootstrap; the loop requires the first iteration to create only this ledger.
- Slide edits are deferred to slide-specific passes; the Chicago deck currently has day-level decks (`day1`–`day5`) rather than one deck per book module.
- The standalone `ch00` chapter has no `.outline.md`; it was reviewed as the bootstrap exception rather than receiving an invented outline.
- Create a maintained source repository for the illustrative retrieval and GraphRAG implementations before presenting the book excerpts as a book companion.

## Questions for Miriah

- `blog posts/` and `blog-posts/` both exist (space versus hyphen). Are both live, or is one legacy/archive material?
- The flat research files `chapter-3.md`, `chapter-5.md`, `chapter-6.md`, `chapter-7.md`, `chapter-12.md`, and `notes.md` sit alongside the structured `book/chapters/` tree. Are they legacy drafts, source material, or live manuscript inputs?
- Should the task ledger remain at `book/RESEARCH_LEDGER.md` despite `AGENTS.md` saying new research files belong under `research/`? The bootstrap used the task-specified location and recorded the conflict above.
- The curriculum-to-book mappings above are provisional; please confirm whether “mapping: unsure” rows should be narrowed to direct numbered modules or may cite adjacent modules.
- `ch14.04-local-models-and-model-routing` is mapped to modules 03 and 15, but those sources cover agent harnesses and prompt optimization rather than local inference or model routing. Is there a missing curriculum module, or should this row remain research-only?
- The source curriculum has modules `00_Setup` and `01_Dev_Environment`, while the book has no corresponding numbered chapter. Should those remain only as cross-cutting context, or receive explicit mappings in later passes?
- The source curriculum’s Chicago slides are organized by training day, not book module. Should each module ledger row name a likely day deck in later passes?
