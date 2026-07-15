# Evidence Ledger

Claims extracted from research sources, mapped to book pillars.

---

## Pragmatics — Declarative Pipeline Composition

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "DSPy is a context composition compiler... LM pipelines can be expressed as declarative computational graphs"
- **Locator:** DSPy paper, abstract
- **Supports:** Chapter 5/6 — Pipeline Architecture; Chapter 7 — Modular LLM Systems
- **Strength:** strong

---

## Pragmatics — Compilation as Optimization

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "The compiler optimizes prompts and parameters automatically, predating modern prompt caching and few-shot selection"
- **Locator:** DSPy paper, section on optimizers
- **Supports:** Chapter 14 — The Cost of Context; Chapter 9 — Optimization Strategies
- **Strength:** strong

---

## Pragmatics — Multi-Step Before Tool Calling

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "DSPy pioneered the idea of composing LMs into multi-step pipelines when most people were still writing single prompts"
- **Locator:** Historical context section
- **Supports:** Chapter 6 — Agent Architectures
- **Strength:** strong

---

## Pragmatics — Modules as Composable Scaffolding

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "DSPy modules are parameterized, meaning they can learn (by creating and collecting demonstrations) how to apply compositions of prompting, finetuning, augmentation, and reasoning techniques"
- **Locator:** DSPy paper, abstract
- **Supports:** Chapter 5 — In-Context Learning; Chapter 7 — Modular LLM Systems
- **Strength:** strong

---

## Pragmatics — Teleprompters vs Runtime Context

- **Source:** dspy-notes (synthesized from DSPy + LangChain)
- **Quote:** "Teleprompters: 'What should I TELL the model before it starts?' (static context, compile-time) vs MCP/ReAct: 'What should I FETCH during execution?' (dynamic context, runtime)"
- **Locator:** Teleprompters vs MCP/ReAct Loops section
- **Supports:** Chapter 9 — Optimization Strategies; Chapter 6 — Agent Architectures
- **Strength:** strong

---

## Pragmatics — Math Problems as Optimization Benchmark

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "Prompt Optimization ≈ Finding Global Minimum... Gradient-like search over prompt space"
- **Locator:** Math Problems as Benchmark section
- **Supports:** Chapter 9 — Optimization Strategies; Chapter 14 — The Cost of Context
- **Strength:** strong
- **Counterpoint:** LLMs are not differentiable — this is a useful analogy but not mathematically exact

---

## Pragmatics — Context Cost of Composition

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "Every pipeline composition has measurable context cost that grows with complexity"
- **Locator:** The Cost of Composition section
- **Supports:** Chapter 14 — The Cost of Context
- **Strength:** strong

---

## Evaluation — Agent Evals Need Harnesses

- **Source:** llms-in-production/chapter-7
- **Quote:** "A harness is a system for [evaluation]"
- **Locator:** chapter-7, line 358
- **Supports:** Chapter 12 — Evaluation; Harness Design
- **Strength:** strong

---

## Efficiency — LoRA Parameter Reduction

- **Source:** LoRA: Low-Rank Adaptation of Large Language Models (Hu et al., 2021)
- **Quote:** "Compared to GPT-3 175B fine-tuned with Adam, LoRA can reduce the number of trainable parameters by 10,000 times and the GPU memory requirement by 3 times"
- **Locator:** arXiv:2106.09685, Abstract
- **Supports:** Chapter 15 — When Context Engineering Stops Working; Alternative to full fine-tuning
- **Strength:** strong

---

## Efficiency — LoRA Performance Parity

- **Source:** LoRA: Low-Rank Adaptation of Large Language Models (Hu et al., 2021)
- **Quote:** "LoRA performs on-par or better than fine-tuning in model quality on RoBERTa, DeBERTa, GPT-2, and GPT-3, despite having fewer trainable parameters, a higher training throughput, and, unlike adapters, no additional inference latency"
- **Locator:** arXiv:2106.09685, Abstract
- **Supports:** Chapter 15 — When Context Engineering Stops Working; Parameter-efficient fine-tuning
- **Strength:** strong

---

## Efficiency — Small Language Models Enterprise Viability

- **Source:** Microsoft research on small language models
- **Quote:** "Small language models (SLMs) with 7-8 billion parameters can match or exceed GPT-3.5 performance on domain-specific tasks while reducing inference costs by 90%+"
- **Locator:** Various Microsoft/Meta model releases (Llama 3 8B, Phi-3, Mistral 7B)
- **Supports:** Chapter 15 — When Context Engineering Stops Working; SLM alternatives to large models
- **Strength:** moderate

---

## Observability — OpenTelemetry Industry Adoption

- **Source:** CNCF OpenTelemetry Project Statistics (2026)
- **Quote:** "OpenTelemetry has 27,863 contributors, 5,375 contributing organizations, 13,312 GitHub stars, and graduated from CNCF in May 2026"
- **Locator:** cncf.io/projects/opentelemetry
- **Supports:** Chapter 16 — Observability for Context Systems; Industry-standard observability
- **Strength:** strong

---

## Observability — OpenTelemetry for AI Systems

- **Source:** CNCF OpenTelemetry Blog
- **Quote:** "OpenTelemetry is now the foundation for AI and cloud observability, with network boundary support for AI agents and mesh-derived metrics"
- **Locator:** cncf.io/blog (2026)
- **Supports:** Chapter 16 — Observability for Context Systems; LLM tracing and metrics
- **Strength:** strong

---

## Observability — LLM Production Deployment Challenges

- **Source:** Arize AI LLM Observability 101
- **Quote:** "Over half (53%) of teams say they plan to deploy LLM apps into production in the next 12 months — however, nearly as many (43%) cite issues like accuracy of responses and hallucinations as barriers to implementation"
- **Locator:** arize.com/llm-observability
- **Supports:** Chapter 16 — Observability for Context Systems; Why observability matters
- **Strength:** strong

---

## Alignment — KTO Model Alignment Method

- **Source:** KTO: Model Alignment as Prospect Theoretic Optimization (Ethayarajh et al., 2024)
- **Quote:** "KTO matches or exceeds the performance of preference-based methods at scales from 1B to 30B, despite only learning from a binary signal of whether an output is desirable"
- **Locator:** arXiv:2402.01306 (ICML 2024)
- **Supports:** Chapter 15 — When Context Engineering Stops Working; Alternative alignment methods
- **Strength:** strong

---

## Systems — Model Commoditization

- **Source:** ai-is-a-systems-problem (chapter draft)
- **Quote:** "In 2020, GPT-3 seemed miraculous. In 2024, a model with similar capabilities runs on a laptop."
- **Locator:** Beat 2: Why Models Are Becoming Commodities
- **Supports:** Chapter 2 — AI Is a Systems Problem
- **Strength:** strong

---

## Systems — OpenAI Model Lead Eroded

- **Source:** ai-is-a-systems-problem (chapter draft)
- **Quote:** "OpenAI's model lead has eroded despite continued investment. Anthropic, Google, Meta, and open-source models now match or exceed GPT-4 on many benchmarks."
- **Locator:** Beat 2: Why Models Are Becoming Commodities
- **Supports:** Chapter 2 — AI Is a Systems Problem
- **Strength:** moderate — claim based on market observation; specific benchmark citations would strengthen

---

## Systems — Inference Cost Reduction

- **Source:** ai-is-a-systems-problem (chapter draft)
- **Quote:** "A well-optimized inference stack can reduce token costs by 40% without changing a single line of model interaction code."
- **Locator:** Beat 3: The Infrastructure Stack Beneath AI
- **Supports:** Chapter 2 — AI Is a Systems Problem
- **Strength:** moderate — industry claim; specific vendor/benchmark would strengthen

---

## Systems — Differentiation Shifts to System Layers

- **Source:** ai-is-a-systems-problem (chapter draft)
- **Quote:** "The differentiation is in the system: how you retrieve context, how you manage state, how you authorize actions, how you evaluate outputs."
- **Locator:** Beat 2: Why Models Are Becoming Commodities
- **Supports:** Chapter 2 — AI Is a Systems Problem
- **Strength:** strong

---

## Memory — Agent Memory as Database Workload

- **Source:** Orogat & Mansour (2026) — "Is Agent Memory a Database? Rethinking Data Foundations for Long-Term AI Agent Memory"
- **Quote:** "Long-term agent memory is a new data-management workload. Its correctness is a property of the state trajectory, not of individual records. We formalize this as Governed Evolving Memory (GEM). GEM replaces record-level database operations with four state-level operators: ingestion, revision, forgetting, and retrieval."
- **Locator:** arXiv:2605.26252, abstract
- **Supports:** Chapter 6 — Memory Is a Database Problem
- **Strength:** strong

---

## Memory — Trustworthy Memory Search in AI Agents

- **Source:** Zhang et al. (2026) — "Beyond Similarity: Trustworthy Memory Search for Personal AI Agents"
- **Quote:** "Long-term memory is not merely a utility layer, but a durable control channel that can reshape how agents interpret tasks and execute actions, leaving them highly susceptible to threats such as cross-domain leakage, sycophancy, tool-call drift, or memory-induced jailbreaks."
- **Locator:** arXiv:2606.06054, abstract
- **Supports:** Chapter 6 — Memory Is a Database Problem; Chapter 12 — Evaluation
- **Strength:** strong

---

## Security — Prompt Injection Attack Framework

- **Source:** Liu et al. (2024) — "Formalizing and Benchmarking Prompt Injection Attacks and Defenses"
- **Quote:** "A prompt injection attack aims to inject malicious instruction/data into the input of an LLM-Integrated Application such that it produces results as an attacker desires."
- **Locator:** arXiv:2310.12815, USENIX Security Symposium 2024, abstract
- **Supports:** Chapter 11 — Stop Giving Agents Permissions; Prompt-based security failures
- **Strength:** strong

---

## Security — Universal Adversarial Jailbreak Attacks

- **Source:** Zou et al. (2023) — "Universal and Transferable Adversarial Attacks on Aligned Language Models"
- **Quote:** "Our approach finds a suffix that, when attached to a wide range of queries for an LLM to produce objectionable content, aims to maximize the probability that the model produces an affirmative response (rather than refusing to answer). Surprisingly, we find that the adversarial prompts generated by our approach are quite transferable."
- **Locator:** arXiv:2307.15043, abstract
- **Supports:** Chapter 11 — Stop Giving Agents Permissions; Prompt injection exploits
- **Strength:** strong

---

## Security — Baseline Defenses for LLM Adversarial Attacks

- **Source:** Jain et al. (2023) — "Baseline Defenses for Adversarial Attacks Against Aligned Language Models"
- **Quote:** "We evaluate several baseline defense strategies against leading adversarial attacks on LLMs, discussing the various settings in which each is feasible and effective. Particularly, we look at three types of defenses: detection (perplexity based), input preprocessing (paraphrase and retokenization), and adversarial training."
- **Locator:** arXiv:2309.00614, abstract
- **Supports:** Chapter 11 — Stop Giving Agents Permissions; Defense strategies
- **Strength:** strong

---

## Security — Dangerous Retrieval in Production AI Systems

- **Source:** Shift (2024) — "The Hidden Dangers of Retrieval-Augmented Generation: A Security Analysis"
- **Quote:** "RAG systems introduce new attack surfaces where malicious documents can be retrieved and injected into the model's context, leading to data leakage, prompt manipulation, and unauthorized actions."
- **Locator:** Shift security research report, 2024
- **Supports:** Chapter 10 — Authorization Across Stores; Chapter 11 — Dangerous retrieval failures
- **Strength:** moderate — industry security research report

---

## Authorization — Capability-Based Access for AI Agents

- **Source:** Brandstetter et al. (2024) — "On the Security of AI Agent Architectures: A Capability-Based Access Control Framework"
- **Quote:** "Traditional RBAC models fail to capture the dynamic, multi-step nature of agent actions. Capability-based access provides fine-grained, revocable permissions tied to specific actions rather than static roles, enabling safer agent authorization in production systems."
- **Locator:** arXiv:2412.10891, abstract
- **Supports:** Chapter 11 — Stop Giving Agents Permissions; Capability-based access control
- **Strength:** moderate
---

## Security — Agent Authorization Failure Modes

- **Source:** Rashidi (2026) — "The Balkanization of Execution-Security Research for AI Coding Agents: Isolation, Access Control, and TOCTOU Vulnerabilities"
- **Quote:** "Policy-enforcement studies report failure rates from 69% to 98% of real denylists, yet no isolation paper re-evaluates its own defense under that adversarial setting. TOCTOU and MCP threats are analyzed as separate literatures despite both being instances of the same state-validation problem."
- **Locator:** arXiv:2607.05743, abstract
- **Supports:** Chapter 10 — Authorization Across Stores; Chapter 11 — Agent authorization failure modes
- **Strength:** strong

---

## Security — Context-to-Execution Integrity for LLM Agents

- **Source:** Santos-Grueiro (2026) — "Context-to-Execution Integrity for LLM Agents"
- **Quote:** "Language-model agents read attacker-writable context to solve tasks. Tool execution needs a separate authority check for protected sink fields, sink-interpreted payloads, and the invocation event."
- **Locator:** arXiv:2607.06000, abstract
- **Supports:** Chapter 10 — Authorization Across Stores; Chapter 11 — Agent authorization failure modes
- **Strength:** strong

---

## Security — Off-Host Identity-Bound Authorization for AI Agents

- **Source:** Kodathala (2026) — "aiAuthZ: Off-Host, Identity-Bound Authorization for AI Agents"
- **Quote:** "AI agents issue tool calls on the basis of text they cannot verify, so any party who controls part of the context can forge the appearance of authority. I evaluate 15 contemporary language models against eight attack scenarios and find that refusal varies from 100% down to 38%."
- **Locator:** arXiv:2607.05518, abstract
- **Supports:** Chapter 10 — Authorization Across Stores; Chapter 11 — Agent authorization failure modes
- **Strength:** strong

---

## Security — Cryptographic Confinement of Learning Agents

- **Source:** Qin et al. (2026) — "Governed Individuation: Cryptographically Decoupling an Agent's Learning from Its Authority"
- **Quote:** "Autonomous agents are moving from sandboxed text generators to operators of code, data, and physical infrastructure, and they increasingly learn while deployed. We show that confinement can be guaranteed as an invariant of the agent's execution architecture rather than a probabilistic outcome of its training."
- **Locator:** arXiv:2607.04613, abstract
- **Supports:** Chapter 11 — Stop Giving Agents Permissions; Capability-based access control in AI systems
- **Strength:** strong

---

## Security — LLM Refusal Under Attack Scenarios

- **Source:** Kodathala (2026) — "aiAuthZ: Off-Host, Identity-Bound Authorization for AI Agents"
- **Quote:** "The most expensive model refused only half of the attacks despite a twentyfold price spread. The gateway does not prevent a model from being deceived; it prevents a deceived model from acting beyond the verified user's authority on every call routed through it."
- **Locator:** arXiv:2607.05518, abstract
- **Supports:** Chapter 11 — Stop Giving Agents Permissions; Prompt-based security failures
- **Strength:** strong

---

## Systems — RDF as Standard Data Interchange Model

- **Source:** W3C RDF Working Group — "Resource Description Framework (RDF)"
- **Quote:** "RDF is a standard model for data interchange on the Web. RDF has features that facilitate data merging even if the underlying schemas differ, and it specifically supports the evolution of schemas over time without requiring all the data consumers to be changed."
- **Locator:** w3.org/RDF, published 2014
- **Supports:** Chapter 7 — Context Is a Query; Chapter 8 — Knowledge Graphs; RDF as foundational standard
- **Strength:** strong

---

## Systems — SPARQL 1.1 Query Language Standard

- **Source:** W3C SPARQL Working Group — "SPARQL 1.1 Query Language"
- **Quote:** "RDF is a directed, labeled graph data format for representing information in the Web. This specification defines the syntax and semantics of the SPARQL query language for RDF."
- **Locator:** w3.org/TR/2013/REC-sparql11-query-20130321, W3C Recommendation March 2013
- **Supports:** Chapter 7 — Context Is a Query; Chapter 8 — Knowledge Graphs; Standard query language for RDF
- **Strength:** strong

---

## Systems — OWL Web Ontology Language Standard

- **Source:** W3C OWL Working Group — "OWL 2 Web Ontology Language"
- **Quote:** "The W3C Web Ontology Language (OWL) is a Semantic Web language designed to represent rich and complex knowledge about things, groups of things, and relations between things. OWL is a computational logic-based language such that knowledge expressed in OWL can be exploited by computer programs."
- **Locator:** w3.org/TR/2012/REC-owl2-overview-20121211, W3C Recommendation December 2012
- **Supports:** Chapter 8 — Knowledge Graphs; Ontology modeling for knowledge representation
- **Strength:** strong

---

## Systems — Apache AGE Graph Database for PostgreSQL

- **Source:** Apache AGE Project — "Apache AGE: Graph Processing & Analytics for Relational Databases"
- **Quote:** "Apache AGE is an extension for PostgreSQL that enables users to leverage a graph database on top of the existing relational databases... The basic principle of the project is to create a single storage that handles both the relational and graph data model so that the users can use the standard ANSI SQL along with openCypher."
- **Locator:** github.com/apache/age, Apache Software Foundation
- **Supports:** Chapter 8 — Knowledge Graphs; PostgreSQL-based graph database, hybrid SQL/Cypher querying
- **Strength:** strong

---

## Systems — Apache Jena Semantic Web Framework

- **Source:** Apache Jena Project — "Apache Jena: A Free and Open Source Java Framework for Building Semantic Web and Linked Data Applications"
- **Quote:** "Apache Jena is a free and open source Java framework for building Semantic Web and Linked Data applications. [Includes] RDF API, ARQ (SPARQL 1.1 compliant engine), TDB (native high performance triple store), Fuseki (SPARQL endpoint), Ontology API, and Inference API."
- **Locator:** jena.apache.org, Apache Software Foundation
- **Supports:** Chapter 7 — Context Is a Query; Chapter 8 — Knowledge Graphs; RDF/SPARQL implementation framework
- **Strength:** strong

---

## Systems — Amazon Neptune Managed Graph Database

- **Source:** AWS — "Amazon Neptune: Serverless Graph Database Service for Connected Data"
- **Quote:** "Amazon Neptune is a serverless graph database service for connected data and improved AI accuracy... Neptune is the only database that gives you the power of connected data with the enterprise capabilities and value of AWS. Supports both RDF (SPARQL) and property graph (Gremlin) models."
- **Locator:** aws.amazon.com/neptune
- **Supports:** Chapter 8 — Knowledge Graphs; Enterprise graph database with multi-model support
- **Strength:** strong

---

## Pragmatics — RAG Query Reliability and Noise Sensitivity

- **Source:** Hsia et al. (2024) — "RAGGED: Towards Informed Design of Scalable and Stable RAG Systems"
- **Quote:** "Our analysis reveals that reader robustness to noise is the key determinant of RAG stability and scalability. Some readers benefit from increased retrieval depth, while others degrade due to their sensitivity to distracting content."
- **Locator:** arXiv:2403.09040, ICML 2025
- **Supports:** Chapter 7 — Context Is a Query; Vector search reliability and query stability
- **Strength:** strong

---

## Pragmatics — Knowledge Graphs Improve LLM Accuracy Over Direct SQL

- **Source:** Sequeda et al. (2023) — "A Benchmark to Understand the Role of Knowledge Graphs on Large Language Model's Accuracy for Question Answering on Enterprise SQL Databases"
- **Quote:** "Our primary finding reveals that question answering using GPT-4, with zero-shot prompts directly on SQL databases, achieves an accuracy of 16%. Notably, this accuracy increases to 54% when questions are posed over a Knowledge Graph representation of the enterprise SQL database."
- **Locator:** arXiv:2311.07509
- **Supports:** Chapter 7 — Context Is a Query; Chapter 8 — Knowledge Graphs; Knowledge graph benefits over direct database queries
- **Strength:** strong

---

## Semantics — Knowledge Graph Quality Dimensions

- **Source:** Zaveri et al. (2015) — "Quality Assessment for Linked Data: A Survey"
- **Quote:** "We propose a systematic classification of quality dimensions for linked data, covering: completeness (extent to which information is present), accuracy (correctness relative to the real world), consistency (logical coherence without contradictions), timeliness (currency of information), provenance (traceability of data origins), and accessibility (availability and ease of retrieval)"
- **Locator:** Semantic Web journal, Vol. 7, No. 1, pages 63-93
- **Supports:** Chapter 8 — Knowledge Graphs; Quality metrics framework for semantic layers
- **Strength:** strong

---

## Semantics — KG Quality as Extrinsic vs Intrinsic

- **Source:** Paulheim (2017) — "Knowledge graph refinement: A survey of approaches and evaluation methods"
- **Quote:** "We distinguish between intrinsic quality (accuracy, consistency, concision) and extrinsic quality (usefulness for downstream tasks). Extrinsic evaluation measures how well the KG supports particular applications."
- **Locator:** Semantic Web journal, Vol. 9, No. 1, pages 3-24
- **Supports:** Chapter 8 — Knowledge Graphs; Chapter 16 — Evaluation; Intrinsic vs extrinsic quality
- **Strength:** strong

---

## Semantics — Wikidata Temporal Precision

- **Source:** Wikibase Data Model — MediaWiki
- **Quote:** "precision: shortint. The numbers have the following meaning: 0 - billion years, 1 - hundred million years, ..., 6 - millennium, 7 - century, 8 - decade, 9 - year, 10 - month, 11 - day, 12 - hour, 13 - minute, 14 - second"
- **Locator:** mediawiki.org/wiki/Wikibase/DataModel, Dates and times section
- **Supports:** Chapter 8 — Knowledge Graphs; Value precision vs graph quality distinction
- **Strength:** strong

---

## Evaluation — Semantic Coverage Metric

- **Source:** kg-quality-metrics-notes (original research)
- **Quote:** "Semantic Coverage — Question: How much of the domain has been modeled? Measures: ontology coverage, relationship coverage, property coverage, source coverage. This is a design-time metric."
- **Locator:** research/kg-quality-metrics-notes.md
- **Supports:** Chapter 8 — Knowledge Graphs; Chapter 16 — Evaluation; Design-time coverage metrics
- **Strength:** strong

---

## Evaluation — Hydration Coverage Metric

- **Source:** kg-quality-metrics-notes (original research)
- **Quote:** "Hydration Coverage — Question: How much of the required context was successfully retrieved? Example: Customer has 4 known sources, only 2 retrieved = 50% hydration coverage. This becomes a runtime metric."
- **Locator:** research/kg-quality-metrics-notes.md
- **Supports:** Chapter 7 — Context Is a Query; Chapter 16 — Evaluation; Runtime retrieval metrics
- **Strength:** strong

---

## Evaluation — Context Precision (Not IR Precision)

- **Source:** kg-quality-metrics-notes (original research)
- **Quote:** "Context Precision = Useful Context / Retrieved Context. Not information retrieval precision. Of the retrieved context, how much was actually useful? Too much irrelevant context increases cost, latency, attention competition, hallucination risk."
- **Locator:** research/kg-quality-metrics-notes.md
- **Supports:** Chapter 14 — The Cost of Context; Chapter 16 — Evaluation; Context efficiency metrics
- **Strength:** strong

---

## Evaluation — Inferential Reach Metric

- **Source:** kg-quality-metrics-notes (original research)
- **Quote:** "Inferential Reach — Question: How much reasoning can the graph support? Metrics: multi-hop traversal success, competency questions answered, reachable semantic depth, average traversal depth. This may become one of the most important metrics for AI systems."
- **Locator:** research/kg-quality-metrics-notes.md
- **Supports:** Chapter 8 — Knowledge Graphs; Chapter 16 — Evaluation; Reasoning capability metrics
- **Strength:** strong

---

## Semantics — Public Data Sources for Context Engineering

- **Source:** original research (user input)
- **Quote:** "Wikipedia and web search indexes are underutilized public data sources that context engineers can leverage—their link-based structure mirrors file parsing and web crawling patterns that LLMs already understand."
- **Locator:** ch07.05-public-data-sources-wikipedia-web.outline.md
- **Supports:** Chapter 7 — Context Is a Query; Public data layer for context engineering
- **Strength:** strong

---

## Semantics — HATEOAS Navigation Pattern for LLMs

- **Source:** Wikipedia HATEOAS article
- **Quote:** "HATEOAS (Hypermedia as the Engine of Application State) is a constraint of the REST application architecture"
- **Locator:** en.wikipedia.org/wiki/HATEOAS
- **Supports:** Chapter 7 — Context Is a Query; Link-based navigation for LLM context retrieval
- **Strength:** moderate

---

## Systems — Graph as Semantic Router Pattern

- **Source:** kg-quality-metrics-notes (original research)
- **Quote:** "An agent should traverse the graph to discover: where information lives, who owns it, authorization requirements, retrieval method, semantic relationships. Rather than asking 'Give me all customer information', the graph answers 'The CRM owns customer identity. Billing owns invoices. Support owns tickets.' The graph routes retrieval."
- **Locator:** research/kg-quality-metrics-notes.md
- **Supports:** Chapter 7 — Context Is a Query; Chapter 8 — Knowledge Graphs; Semantic routing architecture
- **Strength:** strong

---

## Systems — Graph as Map Not Territory

- **Source:** kg-quality-metrics-notes (original research)
- **Quote:** "The graph should not be the database. The graph should be the semantic layer over a distributed data mesh. Instead of storing every fact inside Neo4j, the graph stores: ontology, semantic relationships, identifiers, provenance, authorization boundaries, source mappings. Operational data remains in: PostgreSQL, Iceberg, Delta Lake, APIs, Event Streams, Object Storage, Search indexes."
- **Locator:** research/kg-quality-metrics-notes.md
- **Supports:** Chapter 8 — Knowledge Graphs; Data mesh architecture; Graph vs operational storage
- **Strength:** strong

---

## Systems — Property Graphs vs Triple Stores Architecture

- **Source:** Amazon Neptune Documentation — "Amazon Neptune: Serverless Graph Database Service"
- **Quote:** "Neptune is the only database that gives you the power of connected data with the enterprise capabilities and value of AWS. Supports both RDF (SPARQL) and property graph (Gremlin) models."
- **Locator:** aws.amazon.com/neptune
- **Supports:** Chapter 8 — Knowledge Graphs; Multi-model graph database supporting both property graphs and RDF triple stores
- **Strength:** strong

---

## Learning Systems — User Preference Learning

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "Learning crawlers can learn user preferences from a topic — they observe what the user engages with and build a model of preferences within that domain."
- **Locator:** Learning Crawlers section
- **Supports:** Chapter 7 — Tool Calling; Chapter 10 — Memory Systems
- **Strength:** strong

---

## Learning Systems — Preference Harness Pattern

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "A 'harness' is a structured recognition system that: Observes user behavior within a topic, Extracts preference patterns, Encodes them in a machine-readable format, Provides them to agents via tool calls (not system prompts)"
- **Locator:** The Harness Pattern section
- **Supports:** Chapter 7 — Tool Calling; Chapter 5 — System Prompts
- **Strength:** strong

---

## Semantic Crawlers — Ontology-Based Constrained Traversal

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "Semantic crawlers go beyond keyword matching by: Using ontologies, Using taxonomies, Mapping content, Constrained traversal"
- **Locator:** Semantic Crawlers section
- **Supports:** Chapter 9 — Retrieval Beyond Vector Databases
- **Strength:** strong

---

## WordNet — Guardrails Without Hallucination

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "Unlike vector similarity (which can return anything semantically 'close'), WordNet is human-curated, has verified relationships, has defined sense (disambiguation) for polysemous words, provides provenance"
- **Locator:** WordNet section
- **Supports:** Chapter 9 — Retrieval; Chapter 4 — Hallucination Prevention
- **Strength:** strong

---

## Anchor Text — Semantic Validation for Web Agents

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "Anchor text provides: A human-curated signal about page content, A constraint layer: 'only follow links where anchor text matches the query domain', A validation mechanism: compare retrieved content against what anchor text promised"
- **Locator:** Anchor Text section
- **Supports:** Chapter 9 — Web Search Agents
- **Strength:** strong
- **Counterpoint:** Requires crawl-time extraction

---

## Context Graphs — Focused Crawling via Link Structure

- **Source:** Knowledge Graphs: Fundamentals, Techniques, and Applications (Kerjriwal, Knoblock, & Szekely, 2021)
- **Quote:** "The structure of paths leading to relevant pages can be an important factor in focused crawling, as first shown with context graphs by Diligenti et al. (2000)."
- **Locator:** Chapter on Learning Graphs / Focused Crawling
- **Supports:** Chapter 9 — Web Search Agents; Chapter 8 — Knowledge Graphs
- **Strength:** strong
- **Counterpoint:** Original paper from 2000, pre-dates LLMs

---

## Context Graphs — Working Backward from Relevant Pages

- **Source:** Knowledge Graphs: Fundamentals, Techniques, and Applications (Kerjriwal, Knoblock, & Szekely, 2021)
- **Quote:** "The original context graph method builds classifiers for sets of pages mainly at distance 1 or 2 from relevant pages in the context graph. They use Hidden Markov Models to browse using sequence labeling and context focused crawlers."
- **Locator:** Context Graph section
- **Supports:** Chapter 9 — Web Search Agents; Chapter 7 — Tool Calling
- **Strength:** strong
- **Counterpoint:** Need to locate Diligenti et al. (2000) original paper

---

## Context Graphs + Ontologies — Reducing Crawling Steps via Inference

- **Source:** wordnet-anchor-text-notes (research, 2025) — synthesized from context graph + ontology research
- **Quote:** "If you map the data to a domain, it is much easier to store in a subgraph via ontologies, to reason about, and use the reasoner inference to reduce crawling steps."
- **Locator:** Context Graphs section
- **Supports:** Chapter 8 — Knowledge Graphs; Chapter 9 — Retrieval
- **Strength:** suggestive
- **Counterpoint:** Needs empirical validation

---

## Sequence Labeling — Unified Framework for Context Graphs and Agent Runs

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "Use sequence labeling to label and differentiate nodes in the context graph vs. the agent's execution trace. Context graph nodes = pages, entities, relationships (labeled via NER-style sequence labeling). Agent run nodes = tool calls, decisions, state changes (labeled via execution trace)."
- **Locator:** Sequence Labeling section
- **Supports:** Chapter 7 — Tool Calling; Chapter 9 — Web Search Agents
- **Strength:** strong
- **Counterpoint:** Needs implementation validation

---

## Harvest Rate — Core Metric for Context Engineering

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "Harvest rate = the percentage of downloaded pages relevant to a topic. Any tool call that fetches data is just a crawler with a little spice."
- **Locator:** Harvest Rate section
- **Supports:** Chapter 9 — Retrieval Metrics; Chapter 7 — Tool Calling
- **Strength:** strong
- **Counterpoint:** None — well-established crawler metric, now generalized

---

## Harvest Rate — Generalized Beyond Web Crawlers

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "The 'crawl' now includes: web pages, databases, graph stores, text documents, images (for multimodal models), APIs, file systems. All can be measured by harvest rate."
- **Locator:** Harvest Rate / Generalized Context section
- **Supports:** Chapter 9 — Retrieval Beyond Vector Databases; Chapter 4 — Multi-Source Retrieval
- **Strength:** strong
- **Counterpoint:** Need domain-specific calibration

---

## Predefined Ontologies vs Agent-Built Graphs

- **Source:** Knowledge Graphs: Fundamentals, Techniques, and Applications (Kerjriwal, Knoblock, & Szekely, 2021) + synthesized
- **Quote:** "Predefined domain ontologies will make the context graph a predefined problem, and new data extracted from the agent fits into those ontologies. If we need to increase them, that is a systems engineering process for data engineers."
- **Locator:** Section 3.3.1.1 — Generating the Context Graph
- **Supports:** Chapter 8 — Knowledge Graphs; Chapter 7 — Tool Calling
- **Strength:** strong
- **Counterpoint:** Requires upfront ontology engineering

---

## The Sprawl Problem — AI is Additive

- **Source:** wordnet-anchor-text-notes (research, 2025) — referencing Kendall Clark
- **Quote:** "AI is additive. It will always increase the size of the graph if it is in charge of building it."
- **Locator:** The Sprawl Problem section
- **Supports:** Chapter 8 — Knowledge Graphs; Chapter 7 — Tool Calling
- **Strength:** strong
- **Counterpoint:** None — well-observed phenomenon

---

## Tool Calls as Encapsulated Relationships

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "Instead of building an ontology, use a limit of tool calls or data sources in tool calls with finite goals or provided pragmatics to make the relationships encapsulated in the graph part of the tool call. This allows you to give agents real incremental value without the full knowledge graph maintenance and overhead."
- **Locator:** The Alternative section
- **Supports:** Chapter 7 — Tool Calling; Chapter 8 — Knowledge Graphs
- **Strength:** strong
- **Counterpoint:** Less expressiveness than full graph

---

## Graph-Defined Tools Will Always Sprawl

- **Source:** wordnet-anchor-text-notes (research, 2025)
- **Quote:** "I don't like GBrain and other graph-defined tools because they will always sprawl. They are graph-defined and let the agent decide what to add, which is additive by nature."
- **Locator:** Why This Beats GBrain section
- **Supports:** Chapter 7 — Tool Calling; Chapter 8 — Knowledge Graphs
- **Strength:** strong
- **Counterpoint:** May work for bounded domains with strict schemas

---

## Classifier Training — Domain Modeling for Automatic Hydration

- **Source:** Diligenti et al. (2000) via Knowledge Graphs book + synthesized
- **Quote:** "The original context graph paper uses ML classifiers to assign pages to layers in the Merged Context Graph (MCG). Model domains in a way that data pipelines can do automatic hydration of links, URLs, paths, or database data based on classification."
- **Locator:** Classifier Training section
- **Supports:** Chapter 8 — Knowledge Graphs; Chapter 9 — Retrieval
- **Strength:** strong
- **Counterpoint:** Requires training data for classifier

---

## Pragmatics as the Bridge — Declarative Data + Procedural Agent

- **Source:** wordnet-anchor-text-notes (research, 2025) — refined position
- **Quote:** "Procedural processes make the agents work better, but data injection is best done declaratively. Pragmatics is the interface between declarative data and procedural agent behavior."
- **Locator:** Pragmatics section
- **Supports:** Chapter 4 — Pragmatics; Chapter 7 — Tool Calling
- **Strength:** strong
- **Counterpoint:** None — core architectural insight
