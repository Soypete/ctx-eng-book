# Book Edit Audit

## Purpose

This is the resumable editorial checkpoint for *Context Engineering: Building Reliable AI Systems*. `chapters.md` is authoritative because `README.md` and the local `draft-plan` skill both identify it as the canonical outline.

This file preserves the sequence of the review. The ordered inventory and per-module entries record the state at each checkpoint; they are not all current-state summaries. Later reconciliation sections supersede earlier research-gap counts and citation notes. The `Final Summary`, subsequent source audits, and the manuscript's searchable review markers are the current publication state.

## Discovery Baseline

- Manuscript: `book/preface.md`, the chapter-zero definition essay, and Chapters 1–18 under `book/chapters/`.
- Research: 52 Markdown files under `research/`, including `research/_evidence-ledger.md` and `research/knowledge-graphs/bibliography.md`.
- Local skills: `.opencode/skills/`; the editor workflow is selected by invoking the `editor` skill by name (for example, asking OpenCode to use `editor` on a chapter). No separate editor executable or package script exists.
- Examples: `book/examples/chapter-04-in-context-learning/` plus fenced examples in manuscript modules.
- Validation: no book build, Markdown lint, or link checker is configured. Applicable checks are repository searches, Python `py_compile`, Go `gofmt`/`go build` for standalone examples, and the local `code-audit` workflow.
- Baseline size: 85 manuscript Markdown files, 85 module outlines, approximately 60,000 manuscript words, and 38 fenced code blocks.

## Global Structural Issues

- Chapter 5.4 is now drafted narrowly as structured tool-trace evidence for contract design; Chapters 16 and 17 retain observability and evaluation ownership.
- Chapter 10's canonical title/path mismatch remains recorded in its completion entry.
- Cross-chapter ownership is now explicit: Chapter 8 owns semantic mechanisms, Chapter 10 guardrails and governed context, and Chapter 11 authorization mechanics.
- Definitions for memory, state, retrieval, context graph, and knowledge graph were reconciled across Chapters 0, 6, 7, and 8.
- The promised conclusion, “The Context Engineer,” now exists at `book/conclusion.md`.

## Status Legend

`not reviewed` · `reviewed` · `edited` · `needs author review` · `blocked by missing research` · `blocked by structural decision`

## Ordered Review Inventory

The `Research` and `Issues` columns below are the baseline captured before the chapter-by-chapter edits. Resolved items remain visible as audit history; current unresolved items are enumerated in the final summary.

### Front Matter and Definitions

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/preface.md` | edited | “Agent memory” is usually a data, indexing, semantics, and governance problem. | Start with concrete workflows and engineer retrieval before adopting a memory metaphor. | Preface; not explicitly listed as an outline chapter. | Audited; comparative framework claims still need evidence | None expected | Check overlap with Chapters 6–8. |
| `book/chapters/ch00-what-we-mean-by-context-engineering.md` | edited | Context engineering is the systems discipline of assembling temporal, syntactic, semantic, and pragmatic information around inference. | Treat prompts as one interface within a larger context system. | Matches “What We Mean by Context Engineering.” | Audited; focused-crawling definition needs page citation | Architecture example deferred to Chapter 7 | Establishes working definitions; cross-chapter pass remains. |

### Chapter 1: Every Failure Is a Context Failure

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.01-missing-information.md` | edited | Missing or unusable evidence is a common, diagnosable cause of unsupported output and wrong tool behavior, but not the only cause. | Control evidence, intent, tool scope, and validation before attributing a failure to model capability. | Missing and Incorrect Information | 3 primary-source gaps marked | Complete-context diagram still useful | Recommend Substack split before tool calling. |
| `book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.02-missing-state.md` | edited | Loops, permission failures, and overruns expose missing workflow state and enforcement. | Describe boundaries in context and enforce them in the harness and data layer. | Missing State and Constraints | Security counts and cost figures marked | Existing SQL and policy snippets; executable loop example belongs in Chapter 13 | Permissions/decoding may move during structural pass. |
| `book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.03-context-failure-case-studies.md` | edited | Information, authority, time, and cost failures can be localized to enforceable system boundaries. | Instrument termination, authorization, freshness, context volume, latency, cost, and outcomes. | Context Failure Case Studies | Quantitative claims removed or marked | Summary table remains useful | Observed customer-data pattern distinguished from constructed examples. |
| `book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.04-personalization-failures.md` | edited | Personalization is governed context assembly over authenticated identity and relevant, fresh state. | Enforce subject scope in storage and measure freshness, denial behavior, and correction rates. | Personalization Failures as Missing Context | 3 research gaps marked | Executable cross-tenant denial example needed | Prepares Chapter 10; avoid repeating it there. |

### Chapter 2: AI Is a Systems Problem

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch02-ai-is-a-systems-problem/modules/ch02.01-ai-marketing.md` | edited | “AI” obscures component boundaries; reliability emerges from a capable model and its surrounding system. | Treat prompt wording as a replaceable interface and invest in retrieval, state, authorization, operations, and evaluation. | AI as a Marketing Category | Market-comparison evidence marked | Stack diagram still useful | Database analogy narrowed; Chapter 2 neighbor review in progress. |
| `book/chapters/ch02-ai-is-a-systems-problem/modules/ch02.02-production-ai-stack.md` | edited | Capability comes from the model; reliability emerges from contracts and behavior across the full production stack. | Assign each failure to a component boundary, enforce its contract, and measure end-to-end success. | The Production AI Stack | 2 source gaps marked; LoRA miscitation corrected | Stack diagram still needed | Establishes system map for later chapters. |
| `book/chapters/ch02-ai-is-a-systems-problem/modules/ch02.03-future-ai-engineering.md` | edited | Context engineering is an emerging systems skill set for software engineers who build and maintain agents. | Name owners for retrieval, state recovery, authorization, platform budgets, and evaluation regressions. | The Future of AI Engineering | Organizational prediction explicitly framed as the book's thesis | Team ownership diagram useful | Not presented as a required new job title. |

### Chapter 3: Attention Is All You Need (But We Stopped Paying Attention)

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch03-attention-is-all-you-need/modules/ch03.01-tokens-embeddings-attention.md` | edited | Tokenization, internal representations, retrieval embeddings, and attention are distinct mechanisms with different failure boundaries. | Record serialized input, use the target tokenizer, trace retrieval, and test context ablations. | Tokens, Embeddings, and Attention | Long-context empirical source marked | Tokenization/attention diagrams useful | Corrected token/retrieval embedding conflation and equal-attention claim. |
| `book/chapters/ch03-attention-is-all-you-need/modules/ch03.02-context-windows.md` | edited | A context window is an invocation-scoped token budget, not persistent memory or a linearly read tape. | Build the minimum working set and measure utilization, truncation, latency, recall, and task success. | Context Windows and Positional Limits | 3 primary-source gaps marked | Dated cost example still useful | Removed stale vendor table and universal recency explanation. |
| `book/chapters/ch03-attention-is-all-you-need/modules/ch03.03-compaction-scaffolding-tax.md` | edited | Compaction is a versioned state policy; generated summaries are derived artifacts, not automatic sources of truth. | Preserve source and provenance, test what compaction loses, and retain only scaffolding with measured value. | Compaction, Attention, and the Scaffolding Tax | Talk provenance identified; company architecture claims scoped | Failure-injection example specified | “Scaffolding tax” attributed to Miriah's recollection of Saul Ramirez's Utah meetup talk. |

### Chapter 4: In-Context Learning and Pragmatics

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch04-in-context-learning-and-pragmatics/modules/ch04.01-in-context-learning.md` | edited | Hidden, unconstrained assumptions make failures hard to reproduce; context engineering makes them inspectable and evaluable. | Version invocation inputs, trace workflow evidence, and evaluate residual inference. | In-Context Learning | 2 evidence gaps marked | Invocation/workflow diagram useful | Corrected DSPy/MCP categories and prompt-tuning terminology. |
| `book/chapters/ch04-in-context-learning-and-pragmatics/modules/ch04.02-computational-pragmatics.md` | edited | Intended action and indexical references must be resolved before retrieval or execution. | Use authenticated state, clarification, and policy gates instead of structured guessing. | Computational Pragmatics | 2 complete citations needed | Existing scenario useful; ReAct example corrected | Removed unsupported “most common/highest impact” ranking. |
| `book/chapters/ch04-in-context-learning-and-pragmatics/modules/ch04.03-examples-instructions-structured-outputs.md` | edited | Examples and schemas define an interface; validation, authorization, and evaluation make it reliable. | Generate, parse, validate syntax and semantics, enforce authority, and measure held-out outcomes. | Examples, Instructions, and Structured Outputs | Metric sources retained; citation pass remains | Schema-first executable example still needed | Corrected validation guarantees, example policy, recall, and SER claims. |

### Chapter 5: Tool Use Is Structured Context

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch05-tool-use-is-structured-context/modules/ch05.01-toolformer-and-react.md` | edited | Tools give the model a proposal space; the host creates the action boundary through validation, authorization, execution, and result handling. | Never let a structured proposal acquire authority before host checks. | Toolformer and ReAct | ToolBench source marked; Toolformer corrected | Existing provider-specific snippet remains illustrative | Historical and vendor claims corrected. |
| `book/chapters/ch05-tool-use-is-structured-context/modules/ch05.02-tool-schemas-and-function-calling.md` | edited | Schemas constrain representation; semantic rules and policy create enforceable action boundaries. | Test selection, abstention, malformed, semantically invalid, and unauthorized cases. | Tool Schemas and Function Calling | 2026 citations require verification | Existing schema examples retained | Corrected email-format, selection, and authorization claims. |
| `book/chapters/ch05-tool-use-is-structured-context/modules/ch05.03-tool-selection-routing-validation.md` | edited | Multi-step tool workflows require typed state, terminal states, idempotency, and side-effect-aware recovery. | Make unsafe proposals rejectable, repeated effects idempotent, and recovery deterministic. | Tool Selection, Routing, and Validation | 2026 citation requires verification | Money-transfer routing example materially strengthened | Prepares Chapter 6 state discussion. |
| `book/chapters/ch05-tool-use-is-structured-context/modules/ch05.04-tool-usage-pattern-detection.md` | edited | Structured tool traces reveal hypotheses for safer tool contracts and workflows. | Compare sequences by task and outcome before making a bounded, evaluated change. | Tool Usage Pattern Detection | TwiCal retained only as representation analogy | Customer-search disambiguation example | Chapter 16 owns collection; Chapter 17 owns proof. |

### Chapter 6: Memory Is a Database Problem

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch06-memory-is-a-database-problem/modules/ch06.01-the-myth-of-model-memory.md` | edited | Cross-invocation continuity comes from a governed state lifecycle around the model, not an internal per-user faculty. | Trace apparent memory through capture, authority, storage, indexing, retrieval, context assembly, use, and deletion. | The Myth of Model Memory | Orogat/Mansour plus repository taxonomy retained with narrower claims | User-preference lifecycle example added | Deprecated project/vendor survey removed; retrieval-only and universal-failure claims corrected. |
| `book/chapters/ch06-memory-is-a-database-problem/modules/ch06.02-persistent-state-and-retrieval.md` | edited | State authority, access patterns, consistency, governance, and trajectory evaluation determine the platform—not memory labels or product categories. | Design and test both read and write paths, including supersession, revocation, deletion, rebuild, and replay. | Persistent State and Retrieval | Orogat/Mansour trajectory framing retained | Current-policy versus similar-case architecture example added | Product-prescriptive mappings and disposable-session assumption removed. |
| `book/chapters/ch06-memory-is-a-database-problem/modules/ch06.03-user-session-workflow-state.md` | edited | User, session, and workflow state require different ownership, lifetime, authority, and recovery boundaries. | Place consequential execution facts in durable workflow state and make every cross-category promotion explicit. | User, Session, and Workflow State | Uses chapter research synthesis; no new factual citation required | State table and expense-report recovery example added | Title/body mismatch resolved; long-term recast as retention horizon. |

### Chapter 7: Context Is a Query Over Distributed State

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch07-context-is-a-query/modules/ch07.01-sources-of-context.md` | edited | Invocation context is a task-specific materialized view whose contributors require explicit source contracts. | Inventory authority, scope, selection, time, failure, budget, and provenance for every contributor. | Sources of Context | Durable systems claims; official standards inventory removed | SQL and API-envelope examples | Absolute query, ACID, search, and stream claims corrected. |
| `book/chapters/ch07-context-is-a-query/modules/ch07.02-context-assembly-pipelines.md` | edited | A typed, authorized pipeline emits bounded context, an assembly manifest, and a proceed/degrade/stop decision. | Test source planning, outcomes, budgeting, injection boundaries, and final inclusion independently of the model. | Context Assembly Pipelines | Repository pipeline research synthesized | Executable standard-library Python plus manifest | Concatenation-only design replaced. |
| `book/chapters/ch07-context-is-a-query/modules/ch07.03-freshness-consistency-and-partial-failure.md` | edited | Task invariants determine acceptable time, consistency, retry, conflict, and degradation behavior. | Stop when required evidence cannot meet its authority or consistency contract. | Freshness, Consistency, and Partial Failure | Distributed-systems claims narrowed | Outcome and response table | Universal eventual-consistency and graceful-degradation claims removed. |
| `book/chapters/ch07-context-is-a-query/modules/ch07.04-hydration-coverage-and-retrieval-success.md` | edited | Hydration coverage is a proposed comparison between versioned context requirements and usable results. | Report required and optional `(hydrated, expected)` pairs with outcomes and companion quality metrics. | Hydration Coverage and Retrieval Success | Explicitly labeled original book metric | Executable coverage calculation | Raw source-count ratio and unsupported anecdotes removed. |
| `book/chapters/ch07-context-is-a-query/modules/ch07.05-public-data-sources-wikipedia-web.md` | edited | Public retrieval supplies attributed evidence, but rank, snippets, links, and citations do not establish authority. | Use search for discovery, verify underlying sources, preserve revisions, and isolate retrieved instructions. | Public Data Sources: Wikipedia and Web Indexing | Official Wikibase model retained | Public regulation/private plan and completeness examples | HATEOAS, prevalence, and completeness overclaims corrected. |
| `book/chapters/ch07-context-is-a-query/modules/ch07.06-information-extraction-pipelines.md` | edited | Extraction creates derived proposals that require identity, source-entailment, schema, and policy validation before promotion. | Preserve modality, attribution, time, identity uncertainty, source spans, and extractor versions. | Information Extraction Pipelines | Repository IE synthesis used cautiously | Rich acquisition-event record and pipeline | Noise, necessity, KG benchmark, and model-knowledge claims corrected. |

### Chapter 8: Knowledge Graphs and Semantic Context

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.01-schemas-taxonomies-and-ontologies.md` | edited | Semantic models make decision-relevant meaning engineerable, but graph structure does not guarantee truth or deterministic model reasoning. | Separate schema, taxonomy, ontology, inference, and validation; model from competency questions and invariants. | Schemas, Taxonomies, and Ontologies | Standards claims handed to 8.2; unsupported benchmark removed | Supply-agreement representation | Corrected identity, constraint, model-reasoning, and graph-guarantee overclaims. |
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.02-rdf-owl-and-sparql.md` | edited | RDF data, ontology entailment, graph querying, and operational validation enforce different properties. | Choose RDF, OWL, SPARQL, SHACL, relational constraints, or application rules according to the invariant. | RDF, OWL, SPARQL, and Shapes | Primary W3C standards plus Allemang/Sequeda benchmark | Turtle, SHACL, and SPARQL examples | Public “Data as an AI Guardrail” essay linked and identified as author inference. |
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.03-entity-resolution-and-relationship-traversal.md` | edited | Entity equivalence must be reversible and evidence-bearing; traversal must be scoped and authorized during expansion. | Preserve source identity, tier evidence, and test resolution plus traversal end to end. | Entity Resolution and Relationship Traversal | Claims narrowed to engineering practice | Versioned resolution assertion | Exact-match, canonical-authority, and final-filter assumptions corrected. |
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.04-knowledge-graph-tradeoffs.md` | edited | Graphs earn their cost only for measured relationship-centric workloads that simpler interfaces cannot serve as reliably. | Run a source-backed reversible pilot and count lifecycle, identity, authorization, and evaluation costs. | Knowledge Graph Tradeoffs | Allemang/Sequeda scoped; memstore synthesis | Adoption decision framework | Inevitable-graph, reasoner, algorithm, vendor, and wiki claims removed. |
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.05-instance-coverage-and-ontology-population.md` | edited | Graph population must be measured against versioned source eligibility and competency questions, not class diagrams or raw node counts. | Measure source-backed required paths and keep population failure separate from runtime hydration failure. | Instance Coverage and Ontology Population | Proposed metrics explicitly qualified; adjacent linked-data literature noted | Education required-path example | Dimensionally invalid ratio and universal storage rules corrected. |
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.06-property-completeness-and-schema-quality.md` | edited | Completeness applies only to versioned, applicable task requirements and must preserve distinct absence states. | Measure `(present, applicable)` by task without rewarding invented or overexposed values. | Property Completeness and Schema Quality | Proposed metric qualified | Missingness taxonomy | “More properties is better” removed. |
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.07-ontology-guided-information-extraction.md` | edited | Ontologies should compile into bounded extraction contracts while preserving ambiguity and novelty. | Separate guidance, validation, identity resolution, source entailment, and promotion. | Ontology-Guided Information Extraction | Repository synthesis narrowed | Ingestion pipeline | Prompt compliance no longer treated as validation. |
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.08-knowledge-extraction-methods.md` | edited | Extraction methods are selected per field using representative errors and downstream consequences. | Cascade parsers, classifiers, open extraction, and models behind one contract. | Knowledge Extraction Methods | Universal rankings removed | Method table and cascade | “Hybrid is best” removed. |
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.09-guardrails-for-extraction-validation.md` | edited | Validation proves conformance to evidence and policy, not universal truth. | Layer structure, semantics, entailment, identity, authority, conflict, and promotion checks. | Guardrails for Extraction Validation | Benchmark duplication removed | Validation sequence | Confidence-based auto-accept corrected. |
| `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.10-multilingual-extraction-with-llms.md` | edited | Multilingual extraction reliability requires per-language evaluation and fallback, not one-model claims. | Preserve source language and evaluate every consequential slice and route. | Multilingual Extraction with LLMs | Unsupported quantitative claims removed | Routing/evaluation architecture | Broad transfer, code-switch, and cost claims corrected. |

### Chapter 9: Retrieval Beyond Vector Databases

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch09-retrieval-beyond-vector-databases/modules/ch09.01-lexical-and-relational-retrieval.md` | edited | Query class determines whether keyed, relational, lexical, dense, graph, or hybrid retrieval is appropriate. | Never substitute relevance ranking for scoped authoritative state. | Lexical and Relational Retrieval | Stable IR/database principles | Versioned entitlement SQL | BM25/exact and lexical/semantic absolutes corrected. |
| `book/chapters/ch09-retrieval-beyond-vector-databases/modules/ch09.02-vector-and-semantic-retrieval.md` | edited | Dense retrieval is a governed, versioned candidate generator over derived indexes. | Design chunks, filters, lifecycle, and evaluation from task evidence backward. | Vector and Semantic Retrieval | Volatile vendor evidence removed | Index lifecycle architecture | Similarity no longer implies precision or knowledge. |
| `book/chapters/ch09-retrieval-beyond-vector-databases/modules/ch09.03-graph-and-hybrid-retrieval.md` | edited | Graph queries retrieve explicit paths; hybrid retrieval plans and fuses only required candidate generators. | Preserve path/result lineage and justify combinations through ablation. | Graph and Hybrid Retrieval | Misattributed benchmarks removed | RRF and planned flow | Causality, authorization, and fan-out corrected. |
| `book/chapters/ch09-retrieval-beyond-vector-databases/modules/ch09.04-ranking-reranking-and-query-planning.md` | edited | Query plans define task-relative rightness, routes, reranking, budgets, and stop conditions. | Keep generation, reranking, and final selection observable. | Ranking, Reranking, and Query Planning | Chapter synthesis | Planning boundary | Expanded from repetitive conclusion. |
| `book/chapters/ch09-retrieval-beyond-vector-databases/modules/ch09.05-context-precision-and-context-recall.md` | edited | Metrics require explicit judgment units and bounded denominators; context precision is a proposed view. | Evaluate retrieval, assembly, evidence use, and outcomes separately. | Context Precision and Context Recall | Proposed metric qualified | End-to-end harness | Invented equations and invalid proxy removed. |

### Chapter 10: Guardrails and Ontology-Based Validation

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch10-personalization-is-governed-data-access/modules/ch10.00-guardrails-and-ontology-based-validation.md` | edited | Guardrails are enforceable checks at named system boundaries. | Match schemas, shapes, policy, and postconditions to the invariant they enforce. | Guardrails and Ontology-Based Validation | Chapter 8 semantics retained | Expense validation chain | Directory title mismatches canonical chapter. |
| `book/chapters/ch10-personalization-is-governed-data-access/modules/ch10.01-personalization-as-retrieval.md` | edited | Personalization binds trusted identity and purpose to retrieval decisions. | Never treat prompts or model attributes as authorization. | Personalization as Governed Retrieval | Semantic-contract synthesis | Authorization envelope | Storage-only and ABAC superiority claims narrowed. |
| `book/chapters/ch10-personalization-is-governed-data-access/modules/ch10.02-scoped-hydration.md` | edited | Each store must enforce scope without weaker index, cache, path, or credential bypasses. | Push authorization into candidate generation and preserve denial semantics. | Scoped Hydration | Durable security practice | Adapter enforcement matrix | Universal store assumptions removed. |
| `book/chapters/ch10-personalization-is-governed-data-access/modules/ch10.03-provenance-and-derived-context.md` | edited | Provenance chains source, authorization, transformation, inclusion, and descendants. | Inherit restrictions and design lineage for replay and deletion. | Provenance and Derived Context | Provenance literature synthesized | Derived-item record | Lineage explicitly not proof. |
| `book/chapters/ch10-personalization-is-governed-data-access/modules/ch10.04-policy-aware-user-context.md` | edited | Personal context is governed by purpose, inference, retention, and disclosure beyond access. | Compile common decisions into store-specific enforcement and test transitions. | Policy-Aware User Context | Semantic-contract framework | Policy transition suite | Canonical-ABAC prescription removed. |
| `book/chapters/ch10-personalization-is-governed-data-access/modules/ch10.05-provenance-coverage-metrics.md` | edited | Coverage measures conformance to item-type evidence contracts. | Report item/field coverage and reviewed claim support separately. | Provenance Coverage Metrics | Proposed metric qualified | Coverage formulas | Byte ratio and universal fields removed. |

### Chapter 11: Stop Giving Agents Permissions

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch11-stop-giving-agents-permissions/modules/ch11.01-least-privilege.md` | edited | Constrain principal, resource, action, time, and quantity for every model-proposed operation. | Use least privilege to contain, not eliminate, model and pipeline failures. | Least Privilege | Stable security principle | Four-scope model | Root-cause and capability-only claims removed. |
| `book/chapters/ch11-stop-giving-agents-permissions/modules/ch11.02-rbac-abac-capability-based-access.md` | edited | Policy models compose; protocols and token formats are not authorization guarantees. | Keep OAuth/OIDC credentials outside the model and validate exact delegated authority. | RBAC, ABAC, Delegation, and Capabilities | RFC 6749 and OIDC Core | Trust flow | OAuth, OIDC, JWT, capability claims corrected. |
| `book/chapters/ch11-stop-giving-agents-permissions/modules/ch11.03-scoped-credentials-knowledge-stores.md` | edited | Trusted brokers attach narrow credentials to protected calls outside model context. | Prefer handles, explicit audience/task binding, attenuation, and tested revocation. | Scoped Credentials and Trusted Brokers | Security architecture | Broker flow | “Knowledge store” and self-enforcing claims corrected. |
| `book/chapters/ch11-stop-giving-agents-permissions/modules/ch11.04-retrieval-execution-boundaries.md` | edited | Retrieval and execution need separate grants and commit-time re-authorization. | Bind protected fields and current versions to the execution decision. | Retrieval and Execution Boundaries | Chapter synthesis | Commit boundary | Impossibility and primary-failure claims removed. |
| `book/chapters/ch11-stop-giving-agents-permissions/modules/ch11.05-authorization-coverage-and-necessary-access.md` | edited | Allowed access and task necessity are distinct and must be evaluated together. | Pair minimization with required-slot coverage and task success. | Necessary Access and Authorization Coverage | Proposed metric qualified | Necessary-access formula | ReAct/graph and user-only claims removed. |

### Chapter 12: The UNIX Philosophy of AI Systems

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch12-the-unix-philosophy-of-ai-systems/modules/ch12.01-small-composable-systems.md` | edited | Composition improves reliability only at coherent, enforceable boundaries with explicit contracts. | Split by reason to change and test replacement through shadowing and rollback. | Small Composable Systems | UNIX attribution corrected | Boundary decision model | Microservice and compaction overclaims removed. |
| `book/chapters/ch12-the-unix-philosophy-of-ai-systems/modules/ch12.02-pipes-files-explicit-interfaces.md` | edited | Context stages require typed, provenance-bearing envelopes; file APIs remain an analogy. | Preserve backend semantics, errors, deadlines, and metadata through composition. | Pipes, Files, and Explicit Interfaces | Repository pipeline research | JSON envelope | Text/file universality removed. |
| `book/chapters/ch12-the-unix-philosophy-of-ai-systems/modules/ch12.03-mounts-namespaces-isolation.md` | edited | Logical namespaces organize context, while real isolation requires OS/runtime and backend enforcement. | Name and test process, filesystem, network, identity, and resource boundaries. | Mounts, Namespaces, and Isolation | Volatile research removed | Safe namespace tree | Fork/exec and failure-containment equivalence corrected. |
| `book/chapters/ch12-the-unix-philosophy-of-ai-systems/modules/ch12.04-agent-workspaces-secret-management.md` | edited | Workspaces hold scoped task state and handles; secrets remain in trusted executors. | Govern creation, concurrency, quotas, promotion, recovery, and cleanup. | Agent Workspaces and Secret Management | Chapter 11 boundary retained | Secret-free workspace tree | Environment and permission/capability claims corrected. |

### Chapter 13: Agents Are Workflows

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch13-agents-are-workflows/modules/ch13.01-planning-and-react.md` | edited | ReAct and planning expose model proposals to software control; they do not themselves enforce safe action. | Persist plans as versioned proposed state and validate every action at the harness boundary. | Planning and ReAct | ReAct primary paper retained; volatile framework survey removed | Typed plan record and control-flow architecture | None material. |
| `book/chapters/ch13-agents-are-workflows/modules/ch13.02-harnesses-and-state-machines.md` | edited | The harness owns invariants and wraps probabilistic decisions in explicit operational states. | Let models propose; let trusted reducers validate transitions, versions, approvals, and terminal outcomes. | Harnesses and State Machines | No external empirical claim required | Executable Python transition reducer | None material. |
| `book/chapters/ch13-agents-are-workflows/modules/ch13.03-durable-and-event-driven-execution.md` | edited | Durable execution is safe recovery from known states, not blind restart from a transcript. | Classify ambiguous effects, use idempotency and reconciliation, and define what replay means. | Durable and Event-Driven Execution | Unsupported future-dated memory citation removed | Checkpoint and crash-window architecture | None material. |
| `book/chapters/ch13-agents-are-workflows/modules/ch13.04-loops-retries-and-bounded-autonomy.md` | edited | Bounded loops require observable progress, classified failures, and software-owned terminal policy. | Retry only safe failure classes and expand autonomy one enforced boundary at a time. | Loops, Retries, and Bounded Autonomy | No unsupported quantitative claim retained | Decision matrix and retry taxonomy | None material. |

### Chapter 14: The Cost of Context

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch14-the-cost-of-context/modules/ch14.01-token-and-context-economics.md` | edited | Optimize lifecycle cost per successful outcome, not token count alone. | Enforce context admission and evaluate the marginal value of each context policy. | Token and Context Economics | Volatile prices/specs removed | Cost equations and admission policy | None material. |
| `book/chapters/ch14-the-cost-of-context/modules/ch14.02-retrieval-tool-and-latency-costs.md` | edited | Retrieval and tools consume lifecycle cost and the task's end-to-end deadline. | Trace the critical path and price failures, retries, and side effects. | Retrieval, Tool, and Latency Costs | Fabricated case study removed | Lifecycle and expected-tool-cost equations | None material. |
| `book/chapters/ch14-the-cost-of-context/modules/ch14.03-one-shot-execution-loops-and-subagents.md` | edited | Calls and delegation earn their cost only through measurable outcome improvement. | Budget full workflow trees and load-test context length with concurrency. | One-Shot Execution, Loops, and Subagents | Unsupported hardware rule removed | Workflow cost equation | None material. |
| `book/chapters/ch14-the-cost-of-context/modules/ch14.04-local-models-and-model-routing.md` | edited | Deployment and routing decisions require total cost and route-specific reliability measurement. | Use observable cascades and enforce endpoint eligibility outside the router. | Local Models and Model Routing | Volatile model comparisons removed | Routing/cascade architecture | None material. |
| `book/chapters/ch14-the-cost-of-context/modules/ch14.05-context-efficiency-metrics.md` | edited | Context efficiency is constrained multi-objective comparison, not a universal scalar. | Establish hard outcomes, trace a resource vector, and compare nondominated policies. | Context Efficiency Beyond Token Cost | DSPy primary paper linked with narrow claim | Resource vector and frontier method | None material. |
| `book/chapters/ch14-the-cost-of-context/modules/ch14.06-ner-vs-llm-extraction-costs.md` | edited | Extraction methods are comparable only against the same assertion contract and lifecycle horizon. | Benchmark development, operation, review, error, and change costs on production slices. | NER vs. LLM Extraction Costs | Broad local-note claims narrowed | Lifecycle cost equation | None material. |
| `book/chapters/ch14-the-cost-of-context/modules/ch14.07-extraction-method-selection.md` | edited | Route assertion types to the least expensive method that meets their consequence-aware contract. | Evaluate rules, models, LLMs, review, shared validation, and the router itself. | Extraction Method Selection | Unsupported quotation removed | Hybrid routing architecture | None material. |
| `book/chapters/ch14-the-cost-of-context/modules/ch14.08-cost-aware-extraction-pipeline-design.md` | edited | Extraction economics span ingestion through correction and deletion. | Optimize cost per accepted assertion and downstream success with versioned stage ledgers. | Cost-Aware Extraction Pipeline Design | Unsupported scaffolding-tax appeal removed | Stage ledger and ablation plan | None material. |

### Chapter 15: When Context Engineering Stops Working

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch15-when-context-engineering-stops-working/modules/ch15.01-diagnosing-model-problems.md` | edited | A model problem is a stable failure that survives controlled system-boundary ablations. | Preserve run manifests, change one variable, and require evidence before training. | Diagnosing Model Problems | No unsupported empirical claim retained | Diagnostic intervention matrix | None material. |
| `book/chapters/ch15-when-context-engineering-stops-working/modules/ch15.02-fine-tuning-and-lora.md` | edited | Adaptation is justified by stable behavior, representative data, and held-out system improvement. | Compare full, parameter-efficient, and no-training routes with rollback. | Fine-Tuning and LoRA | LoRA primary paper scoped to evaluated settings | Adaptation tradeoff table | None material. |
| `book/chapters/ch15-when-context-engineering-stops-working/modules/ch15.03-distillation-and-specialized-models.md` | edited | Distillation, size, and specialization reshape rather than universally improve capability and cost. | Benchmark full workflows and specialize the narrowest necessary component. | Distillation and Specialized Models | Foundational distillation paper added | Candidate decision framework | None material. |
| `book/chapters/ch15-when-context-engineering-stops-working/modules/ch15.04-context-engineering-as-the-research-phase.md` | edited | Context engineering produces the observable contracts that remain after weights change. | Keep governed facts and enforcement outside weights; order experiments by evidence and cost. | Context Engineering as the Research Phase | Unsupported anecdotes removed | Reversible seven-step decision process | None material. |

### Chapter 16: Observability for Context Systems

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch16-observability-for-context-systems/modules/ch16.01-tracing-context-assembly.md` | edited | Traces reconstruct operational boundaries but do not reveal faithful model causality. | Trace planning through terminal outcome, govern content capture, and test hypotheses. | Tracing Context Assembly | OTel and W3C primary specifications added | Context trace and diagnostic sequence | None material. |
| `book/chapters/ch16-observability-for-context-systems/modules/ch16.02-prompt-retrieval-tool-lineage.md` | edited | Lineage connects governed, versioned inputs and decisions to outputs and effects. | Separate candidates from admission and proposals from committed effects. | Prompt, Retrieval, and Tool Lineage | No unsupported empirical claim retained | Tool-effect chain and incident drill | None material. |
| `book/chapters/ch16-observability-for-context-systems/modules/ch16.03-state-cost-latency-observability.md` | edited | Aggregate state, cost, and tail-latency signals reveal systemic regressions that individual traces cannot. | Use low-cardinality metrics, outcome attribution, critical-path tracing, and controlled evaluation. | State, Cost, and Latency Observability | Vendor survey removed | State/cost/latency measurement plan | None material. |

### Chapter 17: Evaluating AI Systems

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch17-evaluating-ai-systems/modules/ch17.01-evals-and-benchmarks.md` | edited | Improvement is a bounded claim supported by validated judgment and controlled comparison. | Predefine claims, thresholds, slices, and uncertainty. | Evals and Benchmarks | Unsupported judge generalizations removed | Evaluation decision sequence | None material. |
| `book/chapters/ch17-evaluating-ai-systems/modules/ch17.02-retrieval-and-tool-evaluation.md` | edited | Boundary evaluation localizes retrieval, admission, tool, and effect failures. | Use ablations, fixtures, and failure injection before end-to-end judgment. | Retrieval and Tool Evaluation | No unsupported claim retained | Retrieval/tool evaluation stack | None material. |
| `book/chapters/ch17-evaluating-ai-systems/modules/ch17.03-regression-and-scenario-testing.md` | edited | Regression compares invariants and outcome distributions, not prose strings. | Encode incidents as versioned scenarios and nearby mutations. | Regression and Scenario Testing | No external claim required | Scenario contract | None material. |
| `book/chapters/ch17-evaluating-ai-systems/modules/ch17.04-reliability-metrics-and-failure-budgets.md` | edited | Objectives and non-fungible budgets make acceptable shortfall explicit. | Define denominators, slices, windows, and release responses. | Reliability Metrics and Failure Budgets | Incorrect budget arithmetic removed | SLI/SLO/budget framework | None material. |
| `book/chapters/ch17-evaluating-ai-systems/modules/ch17.05-qa-driven-srl-benchmarks.md` | edited | QA benchmarks expose linguistic structure but do not define production truth. | Evaluate representation stages and domain transfer separately. | QA-Driven SRL Benchmarks | Primary papers retained | QA-SRL example | None material. |
| `book/chapters/ch17-evaluating-ai-systems/modules/ch17.06-openie-evaluation-relvis.md` | edited | RelVis demonstrates systematic error analysis, not production validity. | Version matching policy and evaluate assertion promotion downstream. | OpenIE Evaluation (RelVis) | RelVis primary paper retained | Error-analysis framework | None material. |

### Chapter 18: Building a Context Engineering Platform

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/chapters/ch18-building-a-context-engineering-platform/modules/ch18.01-source-and-ingestion-architecture.md` | edited | Platform contracts connect authoritative sources, rebuildable derived state, decisions, and outcomes. | Version ingestion, provenance, correction, deletion, and failure outcomes. | Source and Ingestion Architecture | Vendor claims removed | End-to-end evidence flow | None material. |
| `book/chapters/ch18-building-a-context-engineering-platform/modules/ch18.02-semantic-and-retrieval-infrastructure.md` | edited | Retrieval produces an authorized evidence manifest with explicit uncertainty. | Plan slots and routes; preserve scores, failures, provenance, and admission decisions. | Semantic and Retrieval Infrastructure | Unsupported universal retrieval claims removed | Retrieval contract | None material. |
| `book/chapters/ch18-building-a-context-engineering-platform/modules/ch18.03-authorization-state-and-tooling.md` | edited | Trusted software binds authority, durable state, and effects around model proposals. | Enforce at query and commit, keep credentials external, and reconcile before retry. | Authorization, State, and Tooling | Future-dated and universal security claims removed | Recovery architecture | None material. |
| `book/chapters/ch18-building-a-context-engineering-platform/modules/ch18.04-observability-evaluation-cost-control.md` | edited | Production readiness is the ability to detect, diagnose, contain, evaluate, and reverse failure. | Govern run manifests, tests, rollouts, SLOs, cost, ownership, and rollback. | Observability, Evaluation, and Cost Control | No unsupported empirical claim retained | Feedback loop and checklist | None material. |

### Conclusion

| File | Status | Central assertion | Practical takeaway | Outline relationship | Research | Code/example | Issues |
|---|---|---|---|---|---|---|---|
| `book/conclusion.md` | edited | Context engineering makes evidence, state, authority, action, and evaluation explicit around uncertain models. | Build systems that can constrain, explain operationally, recover, evaluate, and be corrected. | Conclusion: The Context Engineer | Synthesizes established book arguments | Reliability progression | None material. |

## Cross-Chapter Duplication and Prerequisite Baseline

- Chapters 1 and 7 both argue that missing or incorrectly assembled information causes failures; Chapter 1 should diagnose, while Chapter 7 should implement.
- Chapters 5, 10, and 11 all discuss constraints. Chapter 5 should own typed action interfaces, Chapter 10 semantic/output validation and governed hydration, and Chapter 11 execution authority.
- Chapters 6 and 13 both discuss state. Chapter 6 should own storage semantics and retrieval horizons; Chapter 13 should own workflow transitions, recovery, and durable execution.
- Chapters 8 and 9 both discuss graphs. Chapter 8 should own representation and data quality; Chapter 9 should own retrieval strategy and ranking.
- Chapter 3 is prerequisite to Chapters 4–5; Chapter 6 is prerequisite to Chapters 7–10; Chapters 10–11 are prerequisite to Chapter 12; Chapters 16–17 are required before the platform synthesis in Chapter 18.

## Per-Module Completed Reviews

### Preface — Everyone Is Talking About AI Memory

- **File:** `book/preface.md`
- **Outline section:** Front matter; establishes the data-engineering thesis used by Chapters 6–9.
- **Dominant thesis:** “Memory” requests usually expose retrieval, semantic-modeling, governance, and data-quality needs.
- **Edits made:** Corrected grammar and headings; reduced fragmentation; narrowed claims about graphs, scale, and model interchangeability; aligned the thesis with durable state and authorization.
- **Assertions verified:** No new factual claim added; graph claims narrowed to what the manuscript supports.
- **Citations added or corrected:** None.
- **Unsupported claims remaining:** Comparative framework scalability needs evidence if retained as more than opinion.
- **Counterarguments:** Memory tools can help bounded systems; models are not interchangeable.
- **Code examples:** None needed.
- **Practical agent-building takeaway:** Start from a workflow and its retrieval requirements before selecting a memory product or graph.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** none.
- **Validation commands run:** chapter validation pending.
- **Status:** edited.

### Definitions — What We Mean by Context Engineering

- **File:** `book/chapters/ch00-what-we-mean-by-context-engineering.md`
- **Outline section:** What We Mean by Context Engineering.
- **Dominant thesis:** Context engineering controls information flow across temporal, syntactic, semantic, and pragmatic dimensions; prompting is one interface.
- **Edits made:** Recognized genuine model limits; declared the book's usage of overloaded “context graph”; added enforcement-layer diagnostics and measurement.
- **Assertions verified:** Focused crawling is presented as one context-graph usage, not a universal definition.
- **Citations added or corrected:** Added Hugging Face's primary July 2026 agent-intrusion timeline for the documented escape, credential acquisition, lateral movement, trust-boundary weaknesses, and remediations; labeled the authorization-context mapping as the book's inference.
- **Unsupported claims remaining:** Focused-crawling definition needs a page-level citation.
- **Counterarguments:** Explicitly recognizes model-capability failures and term overloading.
- **Code examples:** Architecture implementation deferred to Chapter 7.
- **Practical agent-building takeaway:** Inspect timing, structure, meaning, intended use, and enforcement when behavior fails.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** none.
- **Validation commands run:** chapter validation pending.
- **Status:** edited.

### Chapter 1.1 — Missing and Incorrect Information

- **File:** `book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.01-missing-information.md`
- **Outline section:** Hallucinations, wrong tool calls, and the demo-to-production gap.
- **Dominant thesis:** Missing or unusable evidence is a common cause of unsupported output and incorrect tool behavior, but not the only cause.
- **Edits made:** Added an operational hallucination definition; corrected the false Vaswani quotation; distinguished ReAct from decomposition; corrected decoding versus attention; clarified MCP's boundary; added model-capability counterexamples and a transition.
- **Assertions verified:** Vaswani et al. support the architecture description, not the former truth/provenance quotation.
- **Citations added or corrected:** Canonical Vaswani arXiv URL added.
- **Unsupported claims remaining:** Hallucination definition, pragmatics transfer, and retrieved prompt-injection threat model need primary sources.
- **Counterarguments:** Better models can reduce ambiguity; adequate context cannot guarantee correctness.
- **Code examples:** Complete-context architecture diagram still useful.
- **Practical agent-building takeaway:** Make intent and preconditions explicit, scope tools and retrieval, validate outputs, then evaluate residual failures.
- **Substack recommendation:** multi-part post; split before “Wrong Tool Calls.”
- **Manual-review markers added:** 3 `RESEARCH NEEDED`.
- **Validation commands run:** chapter validation pending.
- **Status:** edited.

### Chapter 1.2 — Missing State and Constraints

- **File:** `book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.02-missing-state.md`
- **Outline section:** Loops, permission failures, and cost overruns.
- **Dominant thesis:** State, authorization, budgets, and termination must be enforced by the harness rather than left as prompt advice.
- **Edits made:** Separated schema context from workflow enforcement; reframed authorization as governed principal, resource, credential, purpose, action, delegation, temporal, approval, and provenance context; narrowed guarantee language; reframed unverified security counts; separated manifests from runtime enforcement; corrected decoding terminology; labeled the 200-call/500,000-token scenario as a constructed example with illustrative values; added metrics.
- **Assertions verified:** Existing arXiv links retained but require source-level checking later.
- **Citations added or corrected:** Added Hugging Face's primary July 2026 agent-intrusion timeline for the documented escape, credential acquisition, lateral movement, trust-boundary weaknesses, and remediations; labeled the authorization-context mapping as the book's inference.
- **Unsupported claims remaining at this checkpoint:** OWASP taxonomy/counts and the 3% result were later removed. The 200-call/500,000-token scenario is now explicitly constructed and its values explicitly illustrative.
- **Counterarguments:** Declarations are not enforcement; budget reduction may damage recall and task success.
- **Code examples:** Existing snippets retained; executable loop enforcement belongs in Chapter 13.
- **Practical agent-building takeaway:** Track attempts and errors, enforce terminal conditions and capability scope, and measure success with cost.
- **Substack recommendation:** multi-part post; state/termination and permission/decoding have separate outcomes.
- **Manual-review markers added:** none remaining for this module.
- **Validation commands run:** chapter validation pending.
- **Status:** edited.

### Chapter 1.3 — Context Failure Case Studies

- **File:** `book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.03-context-failure-case-studies.md`
- **Outline section:** Chapter 1 case studies.
- **Dominant thesis:** Information, authority, time, and cost failures can be localized to enforceable system boundaries.
- **Edits made:** Distinguished Miriah's observed overhydrated SQL/context pattern from the constructed examples; made explicit that raw rows do not supply meaning, authority, or task relevance and that asking the model to “figure it out” delegates a missing data contract to inference; separated coarse authorization from task-specific minimization and semantic context; removed fabricated-looking precise outcomes; replaced model confidence with evidence and harness budgets; named confused-deputy behavior; added telemetry and transitions.
- **Assertions verified:** No quantitative incident claim is treated as verified.
- **Citations added or corrected:** None.
- **Unsupported claims remaining at this checkpoint:** The original checkpoint left real-versus-constructed status, Rashidi evidence, and a quantitative range unresolved. The later research reconciliation resolved the Rashidi evidence issue and removed the unsupported range. Author review then identified the customer-data design as an observed workplace pattern; all other scenarios are now labeled constructed.
- **Counterarguments:** Lower cost is not reliability if recall or task success falls.
- **Code examples:** Comparative architecture table remains useful.
- **Practical agent-building takeaway:** Capture termination, authorization, freshness, volume, latency, cost, and outcome telemetry.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** none remaining for this module.
- **Validation commands run:** chapter validation pending.
- **Status:** edited.

### Chapter 1.4 — Personalization Failures as Missing Context

- **File:** `book/chapters/ch01-every-failure-is-a-context-failure/modules/ch01.04-personalization-failures.md`
- **Outline section:** Personalization failures; prepares Chapter 10.
- **Dominant thesis:** Personalization is governed context assembly over authenticated identity, relevant state, permission, provenance, and freshness.
- **Edits made:** Separated parameters, state, and context; distinguished declared, inferred, and policy preferences; replaced unverified incident lists with tenant isolation; fixed stale-preference code; tightened identity and least-privilege boundaries; added metrics and a transition.
- **Assertions verified:** Park et al. no longer receives attribution for the repository's synthesized failure taxonomy.
- **Citations added or corrected:** None.
- **Unsupported claims remaining:** Preference-learning evidence and proposed OWASP taxonomy/incidents.
- **Counterarguments:** More personalization increases privacy and stale-state risk; assertions cannot replace storage enforcement.
- **Code examples:** Authenticated retrieval with cross-tenant denial tests is needed.
- **Practical agent-building takeaway:** Authenticate, issue task-scoped authority, enforce it in storage, attach freshness/provenance, and measure correction and denial behavior.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 3 `RESEARCH NEEDED`, 1 `CODE EXAMPLE NEEDED`.
- **Validation commands run:** chapter validation pending.
- **Status:** edited.

### Chapter 2.1 — AI as a Marketing Category

- **File:** `book/chapters/ch02-ai-is-a-systems-problem/modules/ch02.01-ai-marketing.md`
- **Outline section:** AI as a marketing category; model commoditization; database analogy.
- **Dominant thesis:** “AI” obscures component boundaries, while production reliability emerges from a capable model and its surrounding system.
- **Edits made:** Removed intent attribution to vendors; aligned the Chapter 1 recap with its narrowed thesis; qualified falling inference cost and model interchangeability; rejected benchmark cherry-picking; narrowed the database analogy; corrected the claim that prompts cannot be versioned or tested; added a transition.
- **Assertions verified:** Sequeda et al. do not establish model commoditization; the outline readiness was corrected accordingly.
- **Citations added or corrected:** Removed reliance on the secondary market article and misapplied Sequeda citation for the comparative claim.
- **Unsupported claims remaining:** Dated primary model evaluations across specified tasks, costs, and deployment constraints.
- **Counterarguments:** Models remain meaningfully different; inference is not cheap for every workload; database engines are not literal commodities.
- **Code examples:** A production-stack diagram remains useful and belongs primarily in Chapter 2.2.
- **Practical agent-building takeaway:** Assign failures to model, retrieval, state, authorization, orchestration, or evaluation boundaries before choosing a fix.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 1 `RESEARCH NEEDED`.
- **Validation commands run:** balanced-fence check; stale absolute-claim search.
- **Status:** edited.

### Chapter 2.2 — The Production AI Stack

- **File:** `book/chapters/ch02-ai-is-a-systems-problem/modules/ch02.02-production-ai-stack.md`
- **Outline section:** Production stack, inference, retrieval, generalists, and production behavior.
- **Dominant thesis:** Model capability is necessary, while reliability emerges from contracts across hardware, serving, data, orchestration, governance, and evaluation.
- **Edits made:** Corrected the LoRA miscitation; distinguished stored data from assembled context; narrowed hybrid-retrieval claims; separated model decisions from workflow cascades; removed dismissive generalist language; added component-boundary diagnosis and a transition.
- **Assertions verified:** Hu et al. do not support a 40 percent serving-cost reduction.
- **Citations added or corrected:** LoRA citation retained only for its actual subject; vendor observability citation removed from evidentiary use.
- **Unsupported claims remaining:** Primary hybrid-retrieval evaluation and standards-based observability/reliability support.
- **Counterarguments:** Models still cause failures; hybrid retrieval is not required for every workload; prompts remain part of the stack.
- **Code examples:** Stack diagram needed; no weak quota-driven code recommended.
- **Practical agent-building takeaway:** Map every request through component contracts and attribute failures before changing prompts or models.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 2 `RESEARCH NEEDED`.
- **Validation commands run:** Chapter 2 balanced-fence and stale-claim checks pending.
- **Status:** edited.

### Chapter 2.3 — The Future of AI Engineering

- **File:** `book/chapters/ch02-ai-is-a-systems-problem/modules/ch02.03-future-ai-engineering.md`
- **Outline section:** AI generalism and future specialization.
- **Dominant thesis:** Context engineering is the emerging systems skill set software engineers need to build and maintain production agents, not necessarily a separate job title.
- **Edits made:** Resolved the role framing; reframed forecasting as the book's hypothesis about software-engineering skills; clarified that data remains data until assembled for a task; made techniques optional; added lifecycle, recovery, revocation, and evaluation responsibilities; qualified OpenTelemetry; recognized combined roles and specialist collaboration; added an ownership-based reliability conclusion.
- **Assertions verified:** No organizational forecast treated as established evidence.
- **Citations added or corrected:** Removed broken local-link citations from prose; retained explicit markers for two paper checks.
- **Unsupported claims remaining:** None material. The skill-set forecast is explicitly the book's thesis rather than established labor-market evidence. Orogat/Mansour is verified and scoped; the invalid Brandstetter attribution was removed.
- **Counterarguments:** Small teams can combine roles; specialization has coordination cost and is not automatically reliable.
- **Code examples:** Team ownership diagram useful; code not appropriate.
- **Practical agent-building takeaway:** Name an accountable owner and regression signal for each failure boundary.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** none remaining for this module.
- **Validation commands run:** Chapter 2 balanced-fence and stale-claim checks pending.
- **Status:** edited.

### Chapter 3.1 — Tokens, Embeddings, and Attention

- **File:** `book/chapters/ch03-attention-is-all-you-need/modules/ch03.01-tokens-embeddings-attention.md`
- **Outline section:** Tokens, embeddings, attention, and what a model can use.
- **Dominant thesis:** Tokenization, internal representations, retrieval embeddings, and attention are separate mechanisms whose boundaries must be observable.
- **Edits made:** Distinguished token IDs, input embeddings, contextual representations, and retrieval embeddings; added causal masks, heads, and layers; removed equal-attention and deterministic-dilution claims; added an ablation-based debugging path and transition.
- **Assertions verified:** Vaswani et al. supports transformer mechanics, not the stronger production distractor claims.
- **Citations added or corrected:** Vaswani citation retained with narrowed attribution.
- **Unsupported claims remaining:** Primary retrieval-embedding and long-context distractor evidence.
- **Counterarguments:** Attention weights are not faithful explanations; more context is task-dependent rather than always harmful.
- **Code examples:** Diagrams are more appropriate than code.
- **Practical agent-building takeaway:** Serialize, count, trace, ablate, and compare task outcomes.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 1 `RESEARCH NEEDED`.
- **Validation commands run:** Chapter 3 validation pending.
- **Status:** edited.

### Chapter 3.2 — Context Windows and Positional Limits

- **File:** `book/chapters/ch03-attention-is-all-you-need/modules/ch03.02-context-windows.md`
- **Outline section:** Position, context limits, and context cost.
- **Dominant thesis:** The context window is an invocation working-set limit; effective utilization and cost depend on model, task, content, and placement.
- **Edits made:** Distinguished positional methods; removed causal/recency claims; reframed lost-in-the-middle as empirical; removed stale vendor capacities and prices; replaced ticket tape with working set; corrected the DSPy miscitation; added operational metrics and transition.
- **Assertions verified:** Original transformer used sinusoidal positions; later-family and serving claims require targeted sources.
- **Citations added or corrected:** Misapplied Khattab evidentiary claim removed.
- **Unsupported claims remaining:** Position-utilization study, official dated window table if retained, and serving benchmarks.
- **Counterarguments:** Advertised capacity is not effective utilization; larger context can help recall while hurting cost or selection.
- **Code examples:** A dated cost calculation is useful; no timeless vendor values.
- **Practical agent-building takeaway:** Measure minimum working-set success with truncation, latency, retrieval, and outcome telemetry.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 3 `RESEARCH NEEDED`.
- **Validation commands run:** Chapter 3 validation pending.
- **Status:** edited.

### Chapter 3.3 — Compaction, Attention, and the Scaffolding Tax

- **File:** `book/chapters/ch03-attention-is-all-you-need/modules/ch03.03-compaction-scaffolding-tax.md`
- **Outline section:** Compaction strategies and scaffolding cost.
- **Dominant thesis:** Compaction is a versioned state-lifecycle policy; summaries should remain derived, attributable, and recoverable artifacts.
- **Edits made:** Removed equal-attention and absolute-output claims; attributed “scaffolding tax” to Miriah's recollection of Saul Ramirez's Utah meetup talk; connected it cautiously to Subquadratic's published description of long-context architecture; distinguished essential from model-compensatory scaffolding; qualified irreversible loss; filled the TBD with five compaction policies, an iterative agent pattern, failure-injection cases, metrics, and a Chapter 4 transition.
- **Assertions verified:** Saul Ramirez publicly identifies himself as leading language-modeling research at Subquadratic; Subquadratic publicly describes retrieval, chunking, and agentic scaffolding as workarounds for dense-attention context limits. The talk attribution remains an author recollection, not a published record.
- **Citations added or corrected:** Removed the unsupported “AI infrastructure research, 2024” authority claim; added Saul Ramirez's public profile and Subquadratic's technical report with explicit source limitations.
- **Unsupported claims remaining:** Comparative effectiveness among compaction policies remains workload-dependent and must be evaluated rather than asserted.
- **Counterarguments:** Some scaffolding is required for security and durability; retaining source makes summary loss recoverable.
- **Code examples:** A compaction failure-injection harness would add value later.
- **Practical agent-building takeaway:** Keep authoritative source, attach policy/version/provenance, rehydrate details, and test known exceptions after compaction.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** none remaining for this module.
- **Validation commands run:** Chapter 3 validation pending.
- **Status:** edited.

### Chapter 4.1 — In-Context Learning

- **File:** `book/chapters/ch04-in-context-learning-and-pragmatics/modules/ch04.01-in-context-learning.md`
- **Outline section:** Few-shot learning, prompt optimization, and runtime context.
- **Dominant thesis:** Reliability improves when consequential assumptions and context inputs are observable and evaluated, not when model inference is rhetorically eliminated.
- **Edits made:** Disambiguated inference; narrowed Brown attribution; reframed goals; replaced static/dynamic teleprompter taxonomy with invocation/workflow context; clarified MCP; separated prompt optimization from learned prompt tuning; qualified novel-label and multilingual extraction.
- **Assertions verified:** Brown and DSPy do not support the former universal reliability law.
- **Citations added or corrected:** Existing paper links retained with narrower claims.
- **Unsupported claims remaining:** Explicit task-specification reliability and multilingual/novel-label comparisons.
- **Counterarguments:** Planning and inference remain necessary; invocation context can be dynamically assembled.
- **Code examples:** Invocation/workflow diagram useful.
- **Practical agent-building takeaway:** Record exact inputs, trace runtime evidence, clarify missing goal fields, and evaluate held-out cases.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 2 `RESEARCH NEEDED`.
- **Validation commands run:** Chapter 4 fence checks; all example runs described below.
- **Status:** edited.

### Chapter 4.2 — Computational Pragmatics

- **File:** `book/chapters/ch04-in-context-learning-and-pragmatics/modules/ch04.02-computational-pragmatics.md`
- **Outline section:** Speech acts, inference, scope, and indexicality.
- **Dominant thesis:** Resolve intended action and contextual referents before granting retrieval or execution authority.
- **Edits made:** Removed unsupported failure ranking; corrected ReAct/chain-of-thought/planning claims; made authority the execution-risk boundary; narrowed DSPy interpretation; corrected indexicality versus semantic meaning; added enforcement and transition.
- **Assertions verified:** ReAct structures reasoning/actions/observations but does not establish user intent.
- **Citations added or corrected:** Yao and Wei paper links retained for their actual mechanisms.
- **Unsupported claims remaining:** Complete Jurafsky taxonomy and authoritative indexicality citations.
- **Counterarguments:** Reference resolution and discourse coherence may matter as much as speech acts.
- **Code examples:** Existing ReAct examples were corrected to distinguish trace from enforcement.
- **Practical agent-building takeaway:** Authenticate the principal, resolve ambiguous references, clarify missing scope, and gate execution.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 2 `RESEARCH NEEDED`.
- **Validation commands run:** Go and Python example execution; fence checks.
- **Status:** edited.

### Chapter 4.3 — Examples, Instructions, and Structured Outputs

- **File:** `book/chapters/ch04-in-context-learning-and-pragmatics/modules/ch04.03-examples-instructions-structured-outputs.md`
- **Outline section:** Few-shot examples, tool formatting, structured outputs, and extraction metrics.
- **Dominant thesis:** Examples and schemas define a probabilistic interface; parsing, semantic validation, authorization, and held-out evaluation create reliability.
- **Edits made:** Removed deterministic programming language; replaced one-example-per-tool guidance; distinguished validation from constrained decoding; corrected schema prompting, recovery stages, recall, and SER; added semantic and authorization failure tests and Chapter 5 transition.
- **Assertions verified:** Syntax conformance does not imply semantic correctness or safety.
- **Citations added or corrected:** No new citation; existing evaluation source remains for later source-level audit.
- **Unsupported claims remaining:** Tool-format prevalence and function-calling history need evidence if stated comparatively.
- **Counterarguments:** Some failures remain model capability failures even with clear schemas.
- **Code examples:** Executable schema-first tool with semantic validation remains needed.
- **Practical agent-building takeaway:** Test malformed, invalid, unauthorized, and tool-error cases—not only happy-path JSON.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** none.
- **Validation commands run:** `python3 -m py_compile`, execution of all five Python examples, `gofmt -l`, `go run`, balanced-fence and stale-claim searches.
- **Status:** edited.

### Chapter 5.1 — Toolformer and ReAct

- **File:** `book/chapters/ch05-tool-use-is-structured-context/modules/ch05.01-toolformer-and-react.md`
- **Outline section:** Toolformer origins, modern calls, and tool ecosystem.
- **Dominant thesis:** A tool call is a model proposal; the host owns validation, authority, execution, and observation.
- **Edits made:** Corrected Toolformer training history and author name; separated runtime definitions from training; added omission/misuse failure modes; corrected provider request flow, data verification, schema support, ToolBench attribution, and duplicate heading.
- **Assertions verified:** Toolformer uses generated/filtered calls and fine-tuning, not runtime in-context examples alone.
- **Citations added or corrected:** Schick et al. corrected; ToolBench source marked.
- **Unsupported claims remaining:** Primary ToolBench attribution and current provider-feature comparisons.
- **Counterarguments:** A bounded proposal is not bounded authority; tools increase side-effect risk.
- **Code examples:** Provider-specific snippet is illustrative and will age; cross-provider code is not required.
- **Practical agent-building takeaway:** Parse, validate, authorize, execute with limits, and return scoped provenance-bearing observations.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 1 `RESEARCH NEEDED`.
- **Validation commands run:** balanced-fence and stale-claim checks.
- **Status:** edited.

### Chapter 5.2 — Tool Schemas and Function Calling

- **File:** `book/chapters/ch05-tool-use-is-structured-context/modules/ch05.02-tool-schemas-and-function-calling.md`
- **Outline section:** Schemas, selection, design, and semantic contracts.
- **Dominant thesis:** Schemas constrain recognized representation; the host combines semantic validation, policy, idempotency, and audit into an action contract.
- **Edits made:** Qualified JSON Schema `format`; separated field shape from authority; narrowed selection-as-retrieval analogy; added deterministic routing and held-out selection evaluation; reframed narrow scope as narrow authority; added five negative test classes.
- **Assertions verified:** Schema validity cannot establish authorization or business correctness.
- **Citations added or corrected:** Unverified 2026 claims consolidated into a research marker.
- **Unsupported claims remaining:** Two 2026 paper attributions and quantitative results.
- **Counterarguments:** Clear schemas cannot compensate for weak tool-use capability.
- **Code examples:** Existing simple/complex schemas useful; executable host policy tests still desirable.
- **Practical agent-building takeaway:** Test no-tool, malformed, semantic-invalid, and unauthorized proposals alongside success.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 1 `RESEARCH NEEDED`.
- **Validation commands run:** balanced-fence and stale-claim checks.
- **Status:** edited.

### Chapter 5.3 — Tool Selection, Routing, and Validation

- **File:** `book/chapters/ch05-tool-use-is-structured-context/modules/ch05.03-tool-selection-routing-validation.md`
- **Outline section:** Multi-step routing, exit criteria, and pragmatic acts.
- **Dominant thesis:** Reliable routing requires typed workflow state, enforced terminals, idempotency, and side-effect-aware retries.
- **Edits made:** Rebuilt the money-transfer example with authenticated resolution, clarification, confirmation, capability, idempotency, commit state, and separate messaging retry; prevented credential/context leakage; replaced prompt completion with terminal states; added timeout ambiguity and compensation; removed inevitable-behavior claims.
- **Assertions verified:** No quantitative paper claim treated as verified.
- **Citations added or corrected:** Santos-Grueiro claim marked for verification.
- **Unsupported claims remaining:** 2026 paper provenance and findings.
- **Counterarguments:** Smaller action spaces remain unsafe when tools carry broad authority.
- **Code examples:** Routing diagram useful; Chapter 13 should own executable state-machine code.
- **Practical agent-building takeaway:** Persist commit IDs, use idempotency keys, classify failure states, and never retry ambiguous side effects blindly.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** 1 `RESEARCH NEEDED`.
- **Validation commands run:** balanced-fence and stale-claim checks.
- **Status:** edited.

### Chapter 5.4 — Tool Usage Pattern Detection

- **File:** `book/chapters/ch05-tool-use-is-structured-context/modules/ch05.04-tool-usage-pattern-detection.md`.
- **Outline section:** Tool-usage pattern detection.
- **Dominant thesis:** Tool traces may be modeled as temporal events and queried for repeated sequences and failures.
- **Edits made:** Drafted the module around structured trace sequences, outcome comparison, bounded contract redesign, and cross-chapter ownership.
- **Assertions verified:** TwiCal does not establish that agent trace mining optimizes behavior.
- **Citations added or corrected:** Existing Ritter source retained as inspiration only.
- **Unsupported claims remaining:** None material; TwiCal is explicitly an analogy rather than transfer evidence.
- **Counterarguments:** Structured traces may not require NLP extraction at all.
- **Code examples:** Customer-search disambiguation illustrates the design loop; no executable example required.
- **Practical agent-building takeaway:** Start with a failure or optimization decision, then collect only telemetry that informs it.
- **Substack recommendation:** standalone post with links to Chapters 16 and 17.
- **Manual-review markers added:** none.
- **Validation commands run:** outline review, neighbor cohesion, link, fence, and global marker checks.
- **Status:** edited.

### Chapter 6.1 — The Myth of Model Memory

- **File:** `book/chapters/ch06-memory-is-a-database-problem/modules/ch06.01-the-myth-of-model-memory.md`
- **Outline section:** Why models do not remember; memory versus retrieval.
- **Dominant thesis:** Cross-invocation continuity is produced by a governed application-state lifecycle, not an internal per-user model faculty.
- **Edits made:** Distinguished parameters, invocation context, caches, and durable state; expanded memory beyond retrieval; added lifecycle and fault domains; clarified the taxonomy and graph boundary; removed vendor inventory and deprecated project references.
- **Assertions verified:** The repository research supports episodic/semantic/procedural distinctions and treats periodic work as maintenance; the Orogat/Mansour citation supports the state-trajectory framing.
- **Citations added or corrected:** Retained Orogat/Mansour and replaced repeated internal references with one repository taxonomy link.
- **Unsupported claims remaining:** None material.
- **Counterarguments:** Correct authorization can withhold state; the model can misuse correct evidence; caches and provider history do not create an internal application record.
- **Code examples:** No code needed; a typed preference lifecycle supplies the production example and evaluation measures.
- **Practical agent-building takeaway:** Trace writes, identity, authority, indexing, retrieval, context inclusion, model use, revision, and deletion separately.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** none.
- **Validation commands run:** Chapter 6 internal-link, deprecated-name, heading, fence, whitespace, and diff checks.
- **Status:** edited.

### Chapter 6.2 — Persistent State and Retrieval

- **File:** `book/chapters/ch06-memory-is-a-database-problem/modules/ch06.02-persistent-state-and-retrieval.md`
- **Outline section:** Persistent state, storage, indexing, retrieval, authorization, and governance.
- **Dominant thesis:** Reliable agent memory requires explicit authority and consistency boundaries across read, write, transformation, and recovery paths.
- **Edits made:** Replaced state-to-product prescriptions with requirements; added consistency gaps, pre-ranking authorization, governed writes, compaction/deletion behavior, and trajectory evaluation.
- **Assertions verified:** The governed evolving memory source supports evaluating state evolution rather than isolated records.
- **Citations added or corrected:** Orogat/Mansour used only for the trajectory-level claim; repository taxonomy linked.
- **Unsupported claims remaining:** None material.
- **Counterarguments:** Semantic retrieval remains useful for evidence discovery but is not the authority or consistency model for binding policy.
- **Code examples:** Current-policy versus similar-case architecture example; executable code is unnecessary at this boundary.
- **Practical agent-building takeaway:** Choose stores from access/lifecycle requirements and test supersession, revocation, deletion, index rebuild, and replay.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** none.
- **Validation commands run:** Chapter 6 internal-link, deprecated-name, heading, fence, whitespace, and diff checks.
- **Status:** edited.

### Chapter 6.3 — User, Session, and Workflow State

- **File:** `book/chapters/ch06-memory-is-a-database-problem/modules/ch06.03-user-session-workflow-state.md`
- **Outline section:** User, session, and long-term state, reconciled with the module's workflow-state title.
- **Dominant thesis:** State placement follows owner, lifetime, authority, and recovery requirements; long-term is a retention horizon rather than an owner.
- **Edits made:** Replaced the mismatched three-type definition; added a comparison table, authoritative versus inferred user state, consequence-based session durability, durable workflow recovery, explicit promotions, and Chapter 13 boundary.
- **Assertions verified:** No external comparative or quantitative assertion introduced.
- **Citations added or corrected:** None required.
- **Unsupported claims remaining:** None material.
- **Counterarguments:** Some session data is safely disposable; only consequential or unreconstructable values require durable workflow storage.
- **Code examples:** Expense-report execution demonstrates idempotency and crash recovery; Chapter 13 should own executable orchestration code.
- **Practical agent-building takeaway:** Name the state owner, lifetime-ending event, authoritative copy, and enforcement boundary before persisting it.
- **Substack recommendation:** standalone post.
- **Manual-review markers added:** none.
- **Validation commands run:** Chapter 6 internal-link, deprecated-name, heading, fence, whitespace, and diff checks.
- **Status:** edited.

### Chapter 7.1 — Sources of Context

- **File / outline:** `book/chapters/ch07-context-is-a-query/modules/ch07.01-sources-of-context.md`; Sources of Context.
- **Dominant thesis:** Context is a task-specific materialized view assembled under explicit source contracts.
- **Edits made:** Added source contract; corrected absolute query, ACID, lexical/dense index, stream, and analytical claims; removed vendor survey.
- **Assertions verified / citations:** Durable database/distributed-system principles used; no new source-dependent quantitative claim.
- **Unsupported claims / counterarguments:** None material; direct inputs need not be queried and source correctness does not guarantee model correctness.
- **Examples:** Tenant-scoped SQL, API envelope, source inventory.
- **Practical takeaway:** Trace eligible sources, transformations, exclusions, and final inclusion.
- **Substack recommendation:** standalone post.
- **Markers:** none. **Validation:** Chapter 7 checks below. **Status:** edited.

### Chapter 7.2 — Context Assembly Pipelines

- **File / outline:** `book/chapters/ch07-context-is-a-query/modules/ch07.02-context-assembly-pipelines.md`; Context Assembly Pipelines.
- **Dominant thesis:** Typed planning and outcomes turn distributed retrieval into an enforceable, replayable boundary.
- **Edits made:** Replaced concatenation example with planning, criticality, outcome taxonomy, budgets, injection treatment, manifest, decisions, and failure tests.
- **Assertions verified / citations:** No external comparative claims; repository pipeline pattern supplies chapter terminology.
- **Unsupported claims / counterarguments:** None material; not every failure should degrade to a prompt.
- **Code:** Executable standard-library Python for source outcomes and decisions; manifest JSON.
- **Practical takeaway:** Assert the assembly result before asking a model to use it.
- **Substack recommendation:** standalone post.
- **Markers:** none. **Validation:** Python execution and Chapter 7 checks below. **Status:** edited.

### Chapter 7.3 — Freshness, Consistency, and Partial Failure

- **File / outline:** `book/chapters/ch07-context-is-a-query/modules/ch07.03-freshness-consistency-and-partial-failure.md`; matching outline section.
- **Dominant thesis:** Task invariants, not a universal consistency preference, determine whether distributed context is safe to use.
- **Edits made:** Added temporal dimensions, authoritative/action-gating reads, outcome table, deadline-bounded retry, deterministic conflict rules, and degradation measures.
- **Assertions verified / citations:** Removed claims needing universal distributed-system support rather than adding citations.
- **Unsupported claims / counterarguments:** None material; attribution is useful but insufficient for governed conflict resolution.
- **Examples:** Billing/profile conflict and witness-statement counterexample.
- **Practical takeaway:** Measure incorrect decisions to proceed with insufficient evidence.
- **Substack recommendation:** standalone post.
- **Markers:** none. **Validation:** Chapter 7 checks below. **Status:** edited.

### Chapter 7.4 — Hydration Coverage and Retrieval Success

- **File / outline:** `book/chapters/ch07-context-is-a-query/modules/ch07.04-hydration-coverage-and-retrieval-success.md`; matching outline section.
- **Dominant thesis:** Coverage compares versioned task slots with results that satisfy their full contracts.
- **Edits made:** Labeled original terminology; split required/optional coverage; added outcomes, limitations, consequence-aware alerting, and failure tests; removed anecdotes and unverified benchmark claims.
- **Assertions verified / citations:** Presented as book hypothesis rather than established evidence.
- **Unsupported claims / counterarguments:** Metric needs empirical validation; it does not measure relevance, correct use, or safety.
- **Code:** Executable standard-library coverage example.
- **Practical takeaway:** Keep numerator, denominator, plan version, and outcome semantics visible.
- **Substack recommendation:** standalone post with proposed-metric qualification.
- **Markers:** none. **Validation:** Python execution and Chapter 7 checks below. **Status:** edited.

### Chapter 7.5 — Public Data Sources: Wikipedia and the Web

- **File / outline:** `book/chapters/ch07-context-is-a-query/modules/ch07.05-public-data-sources-wikipedia-web.md`; Public Data Sources.
- **Dominant thesis:** Public retrieval is an untrusted evidence path whose ranking and links must not be confused with authority.
- **Edits made:** Separated discovery from support; added revision provenance, public/private authority, injection boundary, and evaluation; corrected HATEOAS and completeness claims.
- **Assertions verified / citations:** Official Wikibase data-model link retained for structured statements and metadata.
- **Unsupported claims / counterarguments:** None material; navigability does not imply exhaustive coverage.
- **Examples:** Public regulation/private plan; all-cities completeness counterexample.
- **Practical takeaway:** Open and verify underlying sources, preserve passage-level provenance, and test citation entailment.
- **Substack recommendation:** standalone post.
- **Markers:** none. **Validation:** Chapter 7 checks below. **Status:** edited.

### Chapter 7.6 — Information Extraction Pipelines

- **File / outline:** `book/chapters/ch07-context-is-a-query/modules/ch07.06-information-extraction-pipelines.md`; Information Extraction Pipelines.
- **Dominant thesis:** Extracted structure remains proposed derived state until grounded validation and policy promote it.
- **Edits made:** Distinguished mentions/identity/relations/events/OpenIE; preserved modality and attribution; separated extraction from promotion; added versioned evaluation and lifecycle.
- **Assertions verified / citations:** Removed task-specific KG benchmark and unsupported knowledge-distillation generalization from the argument.
- **Unsupported claims / counterarguments:** None material; documents can be useful without IE and schema precision has costs.
- **Examples:** Acquisition statement and provenance-bearing JSON record.
- **Practical takeaway:** Retain sources and rebuild derived state; never treat generation absent from the passage as extraction.
- **Substack recommendation:** standalone post or direct Chapter 8 bridge.
- **Markers:** none. **Validation:** Chapter 7 checks below. **Status:** edited.

### Chapter 8.1 — Schemas, Taxonomies, and Ontologies

- **File / outline:** `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.01-schemas-taxonomies-and-ontologies.md`; matching outline section.
- **Dominant thesis:** Semantic modeling creates an engineering interface for meaning but does not itself establish truth, identity, or deterministic reasoning.
- **Edits made:** Separated schema/taxonomy/ontology and description/inference/validation; added identity policy, n-ary assertion, competency questions, lifecycle, and limits; removed vendor and benchmark material.
- **Assertions verified / citations:** Standards-specific claims deferred to primary-source module 8.2.
- **Unsupported claims / counterarguments:** None material; explicit graphs can still be stale, wrong, incomplete, or unauthorized.
- **Examples:** Provenance- and time-bearing supply agreement.
- **Practical takeaway:** Add concepts only when they change retrieval, validation, authorization, or a decision.
- **Substack recommendation:** standalone post.
- **Markers:** none. **Validation:** Chapter 8 partial checks below. **Status:** edited.

### Chapter 8.2 — RDF, OWL, SPARQL, and Shapes

- **File / outline:** `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.02-rdf-owl-and-sparql.md`; RDF, OWL, and SPARQL.
- **Dominant thesis:** Data representation, entailment, query, and validation have distinct semantics and must not be treated as interchangeable guarantees.
- **Edits made:** Corrected RDF terms, open-world reasoning, domain/range and cardinality behavior, property-graph equivalence, and relational storage claims; added SHACL and bounded-selection guidance.
- **Assertions verified / citations:** W3C RDF 1.1, OWL 2, SPARQL 1.1, and SHACL; Allemang/Sequeda supports ontology-based query checking and repair only within its reported benchmark.
- **Citations added or corrected:** Added all four W3C primary specifications and arXiv:2405.11706. Miriah's Substack draft 191132102 recorded as accompanying author interpretation.
- **Unsupported claims / counterarguments:** Substack publish link is private; benchmark does not prove universal ontology inference accuracy.
- **Examples:** Executable-looking Turtle data, SHACL shape, and SPARQL query; no repository runtime is configured for them.
- **Practical takeaway:** Assign each invariant to the mechanism whose semantics actually enforce it.
- **Substack recommendation:** standalone standards explainer.
- **Markers:** none. Replaced the private Substack editor URL with the public “Data as an AI Guardrail” post and distinguished its architectural argument from benchmark evidence. **Validation:** Chapter 8 partial checks below. **Status:** edited.

### Chapter 8.3 — Entity Resolution and Relationship Traversal

- **File / outline:** `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.03-entity-resolution-and-relationship-traversal.md`; matching outline.
- **Dominant thesis:** Reliable traversal depends on reversible identity assertions and authorization during bounded path expansion.
- **Edits:** Preserved source records; tiered match signals; added policy versions, temporal scope, per-hop authorization, discovery budgets, and joint failure tests.
- **Evidence / unsupported claims:** No quantitative claim added; removed unsupported Sequeda attribution and product survey.
- **Counterarguments:** Exact identifiers may be strong evidence but shared/recycled attributes are not proof; canonical entities do not own all attributes.
- **Example:** Versioned, evidence-bearing resolution assertion.
- **Practical takeaway:** Measure false merges, missed matches, unauthorized expansion, path quality, and downstream outcomes together.
- **Substack:** standalone post. **Markers:** none. **Validation:** Chapter 8 partial checks. **Status:** edited.

### Chapter 8.4 — Knowledge Graph Tradeoffs

- **File / outline:** `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.04-knowledge-graph-tradeoffs.md`; matching outline.
- **Dominant thesis:** Graph adoption is justified by a measured relationship-centric reliability gain, not by architectural inevitability.
- **Edits:** Replaced universal/topical prescriptions with competency questions, adoption signals, cross-domain contracts, capability separation, full costs, and reversible pilot.
- **Evidence / citations:** Allemang/Sequeda retained only for its ontology-aware text-to-SPARQL benchmark; repository memstore synthesis supports domain-first learning.
- **Unsupported claims:** None material; deleted vendor, wiki, multilingual, and graph-algorithm generalizations.
- **Counterarguments:** Recursive SQL, join tables, tools, and policy engines may remain the correct long-term design.
- **Example:** Source-backed pilot and decision criteria.
- **Practical takeaway:** Benchmark the actual query and failure mode before duplicating state into a graph.
- **Substack:** standalone decision post. **Markers:** none. **Validation:** Chapter 8 partial checks. **Status:** edited.

### Chapter 8.5 — Instance Coverage and Ontology Population

- **File / outline:** `book/chapters/ch08-knowledge-graphs-and-semantic-context/modules/ch08.05-instance-coverage-and-ontology-population.md`; matching outline.
- **Dominant thesis:** Population is task-ready only relative to explicit source eligibility, required paths, and competency questions.
- **Edits made:** Preserved and focused the education example; replaced invalid ratio with class, source-record, relationship, and question coverage; added ingestion contracts, materialization tradeoffs, metric-boundary trace, and replay tests.
- **Assertions verified / citations:** Metrics are labeled as book proposals; linked-data completeness is adjacent evidence rather than claimed validation.
- **Unsupported claims / counterarguments:** Denominators may be unavailable and should not be fabricated; node/edge density does not establish usefulness.
- **Code / architecture examples:** Student → CourseOffering → Assignment → Timestamp path; no code needed.
- **Practical takeaway:** Reconcile eligible source records to authorized, temporally valid paths and evaluate question coverage.
- **Substack recommendation:** standalone post retaining the education example.
- **Manual-review markers:** none.
- **Validation commands:** fence, stale-formula, terminology, and `git diff --check` checks.
- **Status:** edited.

### Chapter 8.6–8.10 — Completion Entries

- **Files / outline:** `ch08.06-property-completeness-and-schema-quality.md` through `ch08.10-multilingual-extraction-with-llms.md`, in canonical outline order.
- **Dominant theses:** Applicable requirements define completeness; ontology guidance narrows proposals; method selection is empirical; validation establishes policy conformance; multilingual quality must be evaluated by slice.
- **Edits made:** Rebuilt the five modules as a continuous ingestion pipeline; removed repeated vendor and benchmark material; distinguished missingness, guidance, extraction, validation, promotion, and multilingual routing.
- **Assertions verified / citations:** No unsupported benchmark result retained. Allemang/Sequeda remains solely in 8.2 and 8.4 with its scoped claim; W3C standards remain in 8.2.
- **Unsupported claims remaining:** The property-completeness formulation is explicitly a proposed metric whose downstream predictive value needs evaluation. No material unsupported factual assertion remains in 8.7–8.10.
- **Counterarguments:** More schema is not necessarily better; ontology guidance can suppress novelty; no extraction method dominates; accepted data may be false; a multilingual endpoint does not provide uniform quality.
- **Code / architecture:** Applicable-property formula, missingness taxonomy, extraction-contract pipeline, method-selection table, layered validation sequence, and multilingual routing/evaluation design. No additional executable code is necessary.
- **Practical agent takeaway:** Keep source evidence and policy versions through extraction, validate before promotion, and restrict capability where a language or assertion type misses its threshold.
- **Substack:** 8.6, 8.8, and 8.10 standalone; 8.7 and 8.9 can stand alone or form a two-part extraction-contract/validation series.
- **Markers:** none. The former Chapter 8.2 public-URL marker is resolved.
- **Validation:** full Chapter 8 fence, stale-claim, marker, link-presence, and `git diff --check`; no book build or Markdown linter exists.
- **Status:** all five edited.

### Chapter 9.1 — Lexical and Relational Retrieval
- **File / outline / thesis:** `ch09.01-lexical-and-relational-retrieval.md`; classify information needs before selecting retrieval primitives.
- **Edits / evidence:** Corrected BM25 versus exact matching and added scoped authoritative lookup; no volatile claim retained.
- **Remaining claims / counterargument:** None material; analyzers can extend lexical behavior and ranked search remains non-authoritative.
- **Example / takeaway / Substack:** Versioned entitlement SQL; use authoritative queries for state; standalone post.
- **Citations / markers / validation / status:** No citation added; none; Chapter 9 checks; edited.

### Chapter 9.2 — Vector and Semantic Retrieval
- **File / outline / thesis:** `ch09.02-vector-and-semantic-retrieval.md`; dense indexes are versioned derived candidate systems.
- **Edits / evidence:** Added evidence-boundary chunking, filtering, canonical fetch, deletion/rebuild, ANN tradeoff, and baseline comparison; removed Mem0 scores.
- **Remaining claims / counterargument:** None material; dense retrieval earns inclusion only on representative queries.
- **Example / takeaway / Substack:** Semantic index lifecycle; design from queries backward; standalone post.
- **Citations / markers / validation / status:** None; none; Chapter 9 checks; edited.

### Chapter 9.3 — Graph and Hybrid Retrieval
- **File / outline / thesis:** `ch09.03-graph-and-hybrid-retrieval.md`; retrieve explicit paths and fuse selected candidate generators.
- **Edits / evidence:** Separated time/causality/authorization, added planned routing, RRF, lineage, and ablation; removed GBrain and Context-Bench claims.
- **Remaining claims / counterargument:** None material; hybrid fan-out can reduce reliability.
- **Example / takeaway / Substack:** RRF and flow; preserve lineage and ablate components; standalone post.
- **Citations / markers / validation / status:** None; none; Chapter 9 checks; edited.

### Chapter 9.4 — Ranking, Reranking, and Query Planning
- **File / outline / thesis:** `ch09.04-ranking-reranking-and-query-planning.md`; task policy defines candidate routes and final selection.
- **Edits / evidence:** Added generation/reranking separation, score incompatibility, authority/freshness/diversity, deadline, budget, and failure tests.
- **Remaining claims / counterargument:** None material; not every task benefits from diversity or every retriever.
- **Example / takeaway / Substack:** Query-plan architecture; inspect pre/post-rerank lists; standalone post.
- **Citations / markers / validation / status:** None; none; Chapter 9 checks; edited.

### Chapter 9.5 — Context Precision and Context Recall
- **File / outline / thesis:** `ch09.05-context-precision-and-context-recall.md`; judged metrics need explicit units and bounded universes.
- **Edits / evidence:** Replaced invented CE/ODE with standard IR framing, qualified evidence-unit context precision, bounded recall, and an end-to-end harness.
- **Remaining claims / counterargument:** Context precision is a book-proposed view; token-level usefulness is often ambiguous.
- **Example / takeaway / Substack:** Evaluation manifest and failure injection; keep boundary metrics separate; standalone post.
- **Citations / markers / validation / status:** No unsupported citation retained; none; Chapter 9 checks; edited.

### Chapter 10 — Per-Module Completion Entries

- **10.00 Guardrails:** Replaced metaphor and vendor survey with input, retrieval, output, execution, and postcondition enforcement boundaries. Ontology, graph, SHACL, schema, and policy roles are distinct. Expense validation chain added. Standalone post; edited; no marker.
- **10.01 Governed personalization:** Added trusted authorization envelope, decision/enforcement separation, ABAC tradeoffs, purpose and personal-data lifecycle. Standalone post; edited; no marker.
- **10.02 Scoped hydration:** Added adapter capability contracts, authorization-aware candidate generation, pooling/cache/index/graph risks, outcome semantics, and alternate-path tests. Standalone post; edited; no marker.
- **10.03 Provenance:** Added multi-parent lineage, inherited restrictions, protected replay, correction/deletion propagation, and a derived-item record. Standalone post; edited; no marker.
- **10.04 Policy-aware user context:** Separated access from use/inference/retention, removed universal canonical ABAC, added a purchasing-agent delegated-authority object with deterministic enforcement boundaries, and added policy-transition evaluation. Standalone post; edited; no marker.
- **10.05 Provenance coverage:** Replaced byte coverage with item-type contracts, item/field formulas, and reviewed claim-support coverage. Proposed metric is qualified. Standalone post; edited; no marker.
- **Assertions/citations:** No volatile quantitative assertion retained. Chapter 8 owns W3C ontology/SHACL semantics; Chapter 11 owns capability mechanisms. Unsupported claims remaining: none material in prose.
- **Practical agent takeaway:** Bind trusted identity and purpose, enforce at every retrieval/action boundary, propagate restrictions through derived state, and retain auditable evidence.
- **Validation:** Full Chapter 10 fence, stale-claim, deprecated-name, marker, and `git diff --check`; no build/linter configured.
- **Structural issue:** **[OUTLINE MISMATCH: canonical Chapter 10 is “Guardrails and Ontology-Based Validation,” but its directory and uploader mapping remain `ch10-personalization-is-governed-data-access`; rename requires a coordinated path migration.]**

### Chapter 11 — Per-Module Completion Entries

- **11.01 Least privilege:** Four scoped dimensions, multiple enforcement mechanisms, separate read/action risk, and misuse tests. Broad permission is a blast-radius issue, not the universal root cause. Standalone; edited; no markers.
- **11.02 Authorization models:** Corrected RBAC/ABAC/capability comparisons, explained these mechanisms as ways to evaluate governed authorization context without making the prompt an enforcement point, corrected OAuth delegation, OIDC ID-token purpose, bearer/revocation tradeoffs, and JWT misconception. Added RFC 6749 and OIDC Core primary citations. Standalone; edited; no markers.
- **11.03 Credential broker:** Replaced credential “knowledge store” with vault/broker/proxy, model-invisible handles, attenuation, revocation/availability choices, and security operations. Standalone; edited; no markers.
- **11.04 Retrieval/execution:** Added distinct grants, exact protected-field binding, commit-time re-authorization, TOCTOU and bypass coverage. Standalone; edited; no markers.
- **11.05 Necessary access:** Qualified necessary-access precision as proposed terminology; separated unauthorized inclusion; paired minimization with sufficiency and task outcomes. Standalone; edited; no markers.
- **Assertions/counterarguments:** No quantitative research claim retained. Capabilities are not the only least-privilege mechanism; OAuth scopes may be broad; self-contained credentials weaken immediate revocation; aggressive minimization can starve a task.
- **Practical takeaway:** The trusted host derives narrow authority, keeps credentials out of context, and re-authorizes exact effects at commit.
- **Validation:** Full Chapter 11 fence, stale security-claim, primary-link, marker, and `git diff --check`; no build/linter configured.

### Chapter 12 — Per-Module Completion Entries

- **12.01 Small composable systems:** Corrected UNIX maxim attribution, made decomposition conditional, added contract/replacement tests and distributed-cost counterargument. Primary 1974 Bell Labs paper linked. Standalone; edited; no marker.
- **12.02 Pipes/files:** Replaced prose/text universal interface with typed envelopes, cancellation, backpressure, and provenance; preserved typed backend operations and bounded mounted-state metaphor. Standalone; edited; no marker.
- **12.03 Namespaces/isolation:** Separated logical organization from real OS/container enforcement; added network, resource, path, symlink, cache, cleanup, and external-effect tests; removed volatile benchmark claims and fork/exec equivalence. Standalone; edited; no marker.
- **12.04 Workspaces/secrets:** Removed model-readable secrets mount, added lifecycle, ownership, concurrency, safe cleanup, trusted handles, credential leakage risks, recovery, and output promotion. Standalone; edited; no marker.
- **Assertions/counterarguments:** UNIX is a design analogy, not proof of reliability; smaller components can increase distributed complexity; files are leaky abstractions for query systems; process isolation contains only isolated resources.
- **Practical takeaway:** Use explicit contracts and real containment mechanisms where the analogy maps; retain typed semantics, backend policy, and durable workflows where it ends.
- **Validation:** Full Chapter 12 fence, historical-attribution, stale-claim, secret-path, marker, link, and `git diff --check`; no build/linter configured.

### Chapter 13 — Per-Module Completion Entries

- **13.01 Planning and ReAct:** Reframed ReAct as an interaction pattern rather than access to faithful internal reasoning; added typed proposal/validation boundaries, versioned plan state, search prerequisites, and framework-independent adoption tests. ReAct paper retained with a narrower claim. Standalone post; edited; no marker.
- **13.02 Harnesses and state machines:** Assigned invariants to trusted code; clarified that operational state can be finite without model behavior being deterministic; added an executable reducer with valid transitions, approval enforcement, terminal states, and optimistic concurrency. Standalone post; edited; no marker.
- **13.03 Durable and event-driven execution:** Replaced seamless-resume and exactly-once implications with checkpoint contents, ambiguous-outcome reconciliation, outbox/idempotency patterns, ordering and schema obligations, fallible compensation, and three meanings of replay. Standalone post; edited; no marker.
- **13.04 Loops, retries, and bounded autonomy:** Added an invariant-based decision matrix, task-specific progress, context compaction policy, failure taxonomy, circuit-breaker limitations, and an incremental autonomy progression. Standalone post; edited; no marker.
- **Assertions/counterarguments:** Planning search is useful only with credible evaluators and bounded state; state machines can create transition sprawl; persistence mechanisms have different semantics; replay cannot recreate the past external world; loops can amplify rather than correct errors.
- **Practical takeaway:** Treat model output as a proposal, preserve workflow state outside the model, classify recovery and retry outcomes, and measure containment before adding autonomy.
- **Validation:** Full Chapter 13 fence, stale-claim, citation, marker, embedded-Python, and `git diff --check`; no build/linter configured.

### Chapter 14 — Per-Module Completion Entries

- **14.01–14.04:** Replaced volatile pricing/specifications, fabricated case studies, universal one-shot advice, free-local-inference framing, and capability stereotypes with outcome cost, admission policy, critical-path accounting, full workflow budgets, capacity tests, and observable model routing. All four are standalone posts; edited; no markers.
- **14.05:** Defined context efficiency as constrained multi-objective comparison, added resource-vector tracing and Pareto-frontier selection, and retained the DSPy primary paper only for its compilation/optimization claim. Standalone post; edited; no marker.
- **14.06–14.08:** Rebuilt extraction economics around comparable assertion contracts, lifecycle horizons, consequence-aware routes, shared validation, router evaluation, stage ledgers, accepted assertions, downstream outcomes, and rebuildable derived state. 14.06 and 14.07 form a useful two-part series; 14.08 stands alone; edited; no markers.
- **Assertions/counterarguments:** More context is not always better; fewer calls are not always cheaper per success; caching and parallelism add consistency/load costs; local and API economics depend on utilization; rules can fail systematically; there is no universal extraction break-even volume.
- **Practical takeaway:** Define hard reliability constraints, attach all resource and recovery costs to terminal outcomes, and select only among designs that meet those constraints.
- **Validation:** Full Chapter 14 fence, stale-price/case-study/overclaim, citation, marker, internal-link, and `git diff --check`; no build/linter configured.

### Chapter 15 — Per-Module Completion Entries

- **15.01 Diagnosis:** Replaced the rigid reliability ladder and intuition-based model/context split with run manifests, controlled ablations, knowledge/behavior/capability/enforcement distinctions, and an escalation matrix. Standalone post; edited; no marker.
- **15.02 Fine-tuning and LoRA:** Scoped LoRA results to the paper's evaluated configurations; removed universal example counts, quality hierarchy, forgetting claims, and speculative J-space path; added data provenance, leakage-resistant splits, system evaluation, canarying, and rollback. Standalone post; edited; no marker.
- **15.03 Distillation and specialized models:** Added the foundational distillation paper; removed invented SLM comparisons and universal degradation; added teacher-data risk, deployed-artifact tests, operational local-serving costs, and narrow-component specialization. Standalone post; edited; no marker.
- **15.04 Research phase:** Made experiment order conditional on reversibility, information gain, consequence, and lifecycle cost; clarified which contracts persist and which facts should remain outside weights. Standalone post; edited; no marker.
- **Assertions/counterarguments:** Correct context does not alone prove a model failure; switching models may beat training; LoRA results do not generalize universally; distilled students can outperform teachers narrowly; model size is not a task-quality metric; context infrastructure also evolves.
- **Practical takeaway:** Prove a stable capability or behavior gap with controlled evidence, choose the smallest intervention, and keep system enforcement and governed state outside the model.
- **Validation:** Full Chapter 15 fence, stale-claim, primary-citation, marker, internal-link, and `git diff --check`; no build/linter configured.

### Chapter 16 — Per-Module Completion Entries

- **16.01 Tracing:** Replaced “trace explains why” with operational reconstructability; added spans and links, proposal/effect separation, OpenTelemetry’s proper scope, W3C propagation, sensitive-content controls, sampling limitations, and hypothesis-driven diagnosis. Standalone post; edited; no marker.
- **16.02 Lineage:** Expanded prompt lineage into rendered context manifests; added retrieval route and exclusion outcomes, end-to-end tool-effect states, protected references, counterfactual replay, and incident drills. Standalone post; edited; no marker.
- **16.03 State/cost/latency:** Removed vendor survey and permanence/low-overhead claims; added transition metrics, cardinality controls, cost-to-outcome attribution, critical-path tails, telemetry failure handling, schema/retention lifecycle, and controlled experiments. Standalone post; edited; no marker.
- **Assertions/counterarguments:** Traces do not expose model reasoning or establish causality; complete payload capture can violate security and privacy; sampling limits diagnosis; telemetry is not the domain source of truth; instrumentation and retention have operational costs.
- **Practical takeaway:** Instrument trusted decisions, preserve governed lineage references, connect aggregate signals to protected traces, and confirm hypotheses through evaluation.
- **Validation:** Full Chapter 16 fence, stale-adoption/survey/causality claim, primary-link, marker, internal-link, and `git diff --check`; no build/linter configured.

### Chapter 17 — Per-Module Completion Entries

- **17.01–17.04:** Added evaluator validation, paired repeated comparison, boundary ablations, scenario contracts, consequence-aware indicators, and correct failure-budget semantics. Removed vendor surveys and deterministic/probabilistic oversimplifications. Standalone posts; edited; no markers.
- **17.05–17.06:** Scoped QA-SRL, QAMR, LSOIE, and RelVis to their representations and datasets; added stage-specific matching, independent verification, domain transfer, promotion, and downstream evaluation. Recommended two-part extraction-evaluation series; edited; no markers.
- **Practical takeaway:** Define the improvement claim, validate judgments, compare against a baseline with uncertainty, protect hard invariants, and confirm offline evidence through guarded production measurement.
- **Validation:** Full Chapter 17 fence, stale-claim, primary-link, marker, internal-link, and `git diff --check`; no build/linter configured.

### Chapter 18 — Per-Module Completion Entries

- **18.01 Sources/ingestion:** Replaced vendor-oriented store layers with end-to-end evidence flow, ingestion contracts, authoritative/derived separation, rebuildability, and store selection by invariant. Standalone synthesis post; edited; no marker.
- **18.02 Semantics/retrieval:** Added task/identity/purpose planning, route-specific scores, explicit failure outcomes, context manifests, semantic limits, ablations, and boundary metrics. Standalone synthesis post; edited; no marker.
- **18.03 Authorization/state/tools:** Removed claims that defense in depth makes violations impossible; added trusted-host authority, protected caches/indexes, durable workflow state, brokered credentials, commit-time authorization, and ambiguous-effect recovery. Standalone synthesis post; edited; no marker.
- **18.04 Observability/evaluation/cost:** Added governed run manifests, multi-layer tests, guarded rollout, consequence-aware SLOs, lifecycle cost, component ownership, migration, and rollback. Standalone platform checklist; edited; no marker.
- **Conclusion:** Created `book/conclusion.md`; narrowed the opening slogan, defined the context engineer by accountable boundaries, preserved the full systems thesis, and closed on correctability rather than model certainty. Standalone concluding essay; edited; no marker.
- **Practical takeaway:** Build the smallest platform that supplies sufficient authorized evidence, constrains action, survives failure, and proves improvement.
- **Validation:** Full-manuscript structural, marker, link, fence, outline-coverage, stale-claim, and `git diff --check`; no build/linter configured.

## Resume Checkpoint

- Current phase: complete manuscript reviewed and edited in canonical order.
- Completed: preface, definition essay, Chapters 1–18, Chapter 5.4, and the conclusion.
- Exact next target: author review of remaining global markers and structural decisions listed in the final summary.
- Last validation: full-manuscript structural, link, fence, marker, outline-coverage, and whitespace checks completed. No book build, Markdown linter, or configured external link checker exists.

### Repository migration checkpoint

- The authoritative book moved from `cmd/authorpedro/books/ctx-eng-book/` to the repository-level `book/` directory.
- The deprecated `cmd/authorpedro/` Go application, its logs, generated Python bytecode, application-specific environment template, and obsolete implementation plan were removed.
- Repository instructions, local authoring-skill paths, README structure, ignore rules, and the Notion uploader now point to `book/`.
- The manuscript migration preserves the chapter/module structure while the editorial pass changes the manuscript content substantially; Git may display these as deletions and additions until rename similarity is calculated after staging.

## Final Summary

The complete manuscript review is finished.

- **Chapters processed:** 18 of 18, plus the definition chapter, preface, and conclusion.
- **Modules processed:** 85 manuscript modules in canonical outline order; 85 corresponding module outlines; no outline lacks a manuscript.
- **Editorial artifacts covered:** 173 Markdown files under `book/`, plus this audit file. This includes the newly drafted Chapter 5.4 and `book/conclusion.md`.
- **Validation run:** all-book outline/manuscript inventory, balanced-fence scan, relative-link resolution, global manual-marker scan, deprecated `authorpedro` reference scan, stale-claim checks recorded per chapter, embedded Python examples where present, and `git diff --check`. The repository defines no build, test, lint, or external-link-check command.

### Remaining manual-review markers

| Category | Count |
|---|---:|
| `AUTHOR REVIEW` | 0 |
| `RESEARCH NEEDED` | 5 |
| `CLAIM TOO BROAD` | 0 |
| `COUNTERARGUMENT NEEDED` | 0 |
| `CODE EXAMPLE NEEDED` | 1 |
| `ARCHITECTURE EXAMPLE NEEDED` | 0 |
| `OUTLINE MISMATCH` | 1 |
| `CROSS-CHAPTER ISSUE` | 0 |
| `SUBSTACK SPLIT` | 0 |

### Weakest remaining areas

- **Research support after repository-note reconciliation:** Chapter 1 (2 research markers), Chapter 3 (1), and Chapter 4 (2). Repository notes and verified primary records resolved 15 markers; four obsolete unsupported examples were removed rather than cited. Chapters 2 and 5 now have no remaining research markers.
- **Practical guidance:** Chapters 2 and 3 remain the most conceptual. Chapter 1.4 has the only explicit code-example marker. Chapters 13, 14, 17, and 18 now provide the strongest implementation, cost, evaluation, and platform progressions.

### Recommended Substack splits

- Chapter 1.1: split before “Wrong Tool Calls.”
- Chapter 1.2: separate state/termination from permission/decoding.
- Chapters 8.7 and 8.9: extraction contract and validation as a two-part series.
- Chapters 14.6 and 14.7: extraction economics and method selection as a two-part series.
- Chapters 17.5 and 17.6: QA-driven semantic extraction and OpenIE evaluation as a two-part series.

### Unresolved structure and author priorities

- The sole outline mismatch is Chapter 10: the canonical title is “Guardrails and Ontology-Based Validation,” while its directory and uploader mapping retain `ch10-personalization-is-governed-data-access`. The manuscript content is aligned; the path migration remains a coordinated repository decision.
- All author-review markers are resolved. The Chapter 8 essay is publicly linked as “Data as an AI Guardrail” and identified as the author's architectural inference rather than benchmark evidence. “Scaffolding tax” is attributed to Miriah's recollection of Saul Ramirez's Utah meetup talk and connected cautiously to Subquadratic's published architecture work. “Context engineer” is defined as the emerging systems skill set for software engineers who build and maintain production agents, not a required new job title. The Chapter 1 customer-data example is identified as an observed workplace design pattern, while the other case-study scenarios and the 200-call/500,000-token values are explicitly constructed.
- Highest-priority research work: add primary evidence for the Chapter 1 hallucination and preference-elicitation claims; then add long-context serving benchmarks for Chapter 3 and task-specific specification/example-selection and multilingual extraction comparisons for Chapter 4.

### Research-note reconciliation (2026-08-06)

- **Supported and added:** Lost in the Middle (long-context position sensitivity), ToolLLM/ToolBench, Orogat and Mansour (agent-memory lifecycle), RAGGED (retrieval depth and reader noise sensitivity), OpenTelemetry, Jurafsky (computational-pragmatics taxonomy and indexicality), Liu et al. (prompt injection), Rashidi (execution-security systematization), Santos-Grueiro (context-to-execution integrity), Kodathala (off-host identity-bound authorization), and Zhang et al. (memory admission as a trust boundary).
- **Unsupported and removed:** former OWASP/ClawHavoc names and counts, the untraceable 3-percent cost result, and the broad model-commoditization claim.
- **Ledger correction:** the former Brandstetter authorization entry was invalid. arXiv:2412.10891 is Bai et al.'s *Zigzag Diffusion Sampling*, not an agent-access-control paper. The evidence ledger now records that correction explicitly.
- **Notes still insufficient:** the local long-context cost notes provide an engineering synthesis but no primary serving benchmark; multilingual IE notes repeat broad advantages without comparative primary evaluations; personalization notes do not contain preference-elicitation evidence; hallucination notes contain book hypotheses rather than a source that separates context-induced from model-capability errors.

### Foundational-book source audit (2026-08-06)

- **Kejriwal, Knoblock, and Szekely, *Knowledge Graphs* (MIT Press, 2021):** directly supports Chapter 8's foundations, construction, refinement, querying, and instance-matching discussion, with secondary relevance to Chapters 7, 9, 14, and 17. Four evidence-ledger entries were corrected so the manuscript's synthesis is no longer presented as verbatim quotation. This textbook does not by itself establish comparative LLM retrieval or multilingual-extraction performance.
- **Kleppmann, *Designing Data-Intensive Applications* (O'Reilly, 2017):** supports the systems foundation used across Chapters 2, 6, 7, 11, 13, 16, and 18: reliability, data models, indexes, schema evolution, replication, transactions, partial failure, consistency, event logs, batch/stream processing, derived state, and end-to-end correctness. It supports architectural principles, not claims about model behavior.
- **Brousseau and Sharp, *LLMs in Production* (Manning, 2024):** supports broad production-lifecycle discussion involving datasets, platform scale, deployment, cost, security, load testing, and model evaluation. It is practitioner guidance rather than primary experimental evidence. Repository attribution was corrected from Chris Fregly and Antje Barth; unattributed references to statements by "Chris" remain provenance-limited until their actual source is identified.
- **Effect on remaining research markers:** none of these three books resolves the five remaining empirical gaps: the Chapter 1 hallucination taxonomy and preference-elicitation claims, Chapter 3 long-context serving benchmarks, or Chapter 4 task-specification/example-selection and comparative multilingual-extraction claims.
