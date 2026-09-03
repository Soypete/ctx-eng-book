# Context Engineering: Building Reliable AI Systems Through Data, State, and Constraints

## Thesis

> Reliable AI systems require engineered context, structured state, semantic constraints, governed retrieval, authorization boundaries, and continuous evaluation.

The book's organizing spine is **Lexicon → Semantics → Pragmatics**:

- **Lexicon:** What data, entities, definitions, sources, and authorities are
  available, and under whose control?
- **Semantics:** What do those entities and relationships mean in this domain,
  task, and time?
- **Pragmatics:** What may a language system produce or do with that meaning,
  for which actor and purpose, under which constraints?

These are not agent-only concerns. They apply to search, extraction,
classification, summarization, recommendation, workflow automation, and
agents. The model is one participant in a larger context-to-outcome pipeline.
The existing temporal, syntactic, semantic, and pragmatic dimensions describe
how context changes and is represented; Lexicon, Semantics, and Pragmatics
provide the recurring editorial questions that connect the chapters.

Context engineering is not prompt engineering.

Prompt engineering is a technique.

Context engineering is a systems discipline that spans:

- data engineering
- retrieval
- memory/state
- authorization
- orchestration
- tool use
- evaluation
- cost management
- reliability engineering

The discipline can be summarized as:

```text
Lexicon       →  Semantics       →  Pragmatics
data/authority   meaning/relations   purpose/action
     ↓                 ↓                  ↓
assembled context → interpreted context → constrained outcome
```

The goal is not to make models guess better. It is to engineer the background
that makes useful interpretation and reliable outcomes possible.

---

# What We Mean by Context Engineering

A chapter establishing precise definitions before diving into the technical content. This clarifies that context engineering is a systems discipline (not prompt engineering), defines the four dimensions (temporal, syntactic, semantic, pragmatic), and distinguishes context graphs from memory/chat history.

---

# Part I — Why AI Systems Fail

## Chapter 1: Every Failure Is a Context Failure

Topics:

- Hallucinations as missing context
- Wrong tool calls as missing context
- Agent loops as missing context
- Permission failures as missing context
- Personalization failures as missing context
- Cost overruns as context failures

Case Studies:

- AI Psychosis
- Infinite agent loops
- Lost conversational state
- Over-scoped assistants

Reliability Question: What information was missing, incorrect, inaccessible, or unconstrained?

---

## Chapter 2: AI Is a Systems Problem

Topics:

- AI as a marketing category
- Why models are becoming commodities
- The infrastructure stack underneath AI
- The rise of the AI generalist
- The future specialization of AI engineering

Reliability Question: Which layers actually determine production behavior?

---

# Part II — How Models Actually Use Context

## Chapter 3: Attention Is All You Need (But We Stopped Paying Attention)

Topics:

- Tokens
- Embeddings
- Attention
- Context windows
- Position encoding
- Why context is expensive

Reliability Question: What information can the model actually see?

---

## Chapter 4: In-Context Learning and Pragmatics

Topics:

- Few-shot prompting
- In-context learning
- Computational pragmatics
- Prompt tuning
- Tool formatting
- Structured outputs

Reliability Question: How do we communicate intent clearly to the model?

---

## Chapter 5: Tool Use Is Structured Context

Topics:

- Toolformer
- Modern tool calling
- JSON/XML schemas
- Function calling
- Tool selection
- Tool routing
- Tool usage pattern detection (TwiCal-style event + time extraction)

Reliability Question: How do we constrain model actions?

---

# Part III — Context Is Data

## Chapter 6: Memory Is a Database Problem

Topics:

- Why models do not remember
- Memory vs retrieval
- User state
- Session state
- Long-term state

Case Studies:

- Pedro
- Memstore research
- Agent memory systems

Reliability Question: Where does state actually live?

---

## Chapter 7: Context Is a Query Over Distributed State

Topics:

- Databases
- APIs
- Documents
- Search indexes
- Event streams
- Lakehouses
- Information extraction pipelines
- Context assembly from multiple sources

Key Thesis: Context is not stored. Context is assembled.

Reliability Question: Which systems contribute information?

---

## Chapter 8: Knowledge Graphs and Semantic Context

Topics:

- RDF
- OWL
- SPARQL
- Ontologies
- Entity resolution
- Relationship traversal
- Ontology-guided information extraction
- Knowledge extraction methods (NER, OpenIE, rule-based, LLM-based)
- Guardrails for extraction validation
- Multilingual extraction with LLMs

Reliability Question: How do we represent meaning instead of text?

---

## Chapter 9: Retrieval Beyond Vector Databases

Topics:

- BM25
- Keyword search
- Graph traversal
- Hybrid retrieval
- Semantic indexing
- Knowledge stores

Reliability Question: How do we retrieve the right information?

---

# Part IV — Context Must Be Governed

## Chapter 10: Guardrails and Ontology-Based Validation

Topics:

- Guardrails reduce hallucinations
- Ontology as guardrails
- Knowledge graph validation
- Entity constraints
- Relation constraints
- Output schema validation
- Agent ABAC (Attribute-Based Access Control)
- Scoped hydration across stores
- Wiki permissions
- Database row-level security
- Knowledge graph access control
- Provenance
- Derived ontologies
- User context

Key Thesis: Guardrails and context work in tandem — context reduces hallucination, guardrails validate output against structured knowledge.

Reliability Question: How do we ensure the model stays within bounds?

---

## Chapter 11: Stop Giving Agents Permissions

Topics:

- Least privilege
- RBAC
- OAuth
- OpenID Connect
- Capability-based access

Reliability Question: How do we prevent dangerous retrieval?

---

## Chapter 12: The UNIX Philosophy of AI Systems

Topics:

- Small composable systems
- Pipes
- Files
- Mounts
- Namespaces
- Process isolation

New Material:

- Context as mounted state
- Agent workspaces
- Scoped filesystem access
- Secret management

Reliability Question: How do we reduce system complexity?

---

# Part V — Orchestration and Cost

## Chapter 13: Agents Are Workflows

Topics:

- ReAct
- Planning
- Harnesses
- State machines
- Durable execution
- Event-driven systems

Key Thesis: Most agents are workflows wearing a trench coat.

Reliability Question: When should the model decide and when should software decide?

---

## Chapter 14: The Cost of Context

Topics:

- Token economics
- Context windows
- Latency
- Retrieval cost
- Tool cost
- Local models
- NER vs LLM extraction cost tradeoff
- Extraction method selection (rule-based vs LLM)

New Material:

- One-shot vs loops
- Subagent costs
- Memory consumption
- Compute tradeoffs
- Cost-aware extraction pipeline design

Reliability Question: What is the cheapest reliable solution?

---

## Chapter 15: When Context Engineering Stops Working

Topics:

- Fine-tuning
- LoRA
- Distillation
- Specialized models
- SLMs

Key Thesis: Context engineering is the research phase. Model modification begins when:
- retrieval cannot solve the problem
- prompting cannot solve the problem
- orchestration cannot solve the problem

Reliability Question: Is this a model problem or a context problem?

---

# Part VI — Reliability Engineering for AI

## Chapter 16: Observability for Context Systems

Topics:

- OpenTelemetry
- Traces
- Prompt lineage
- Retrieval lineage
- Tool lineage

Reliability Question: Why did the model do that?

---

## Chapter 17: Evaluating AI Systems

Topics:

- Evals
- Benchmarks
- Regression testing
- Reliability metrics
- QA-driven SRL benchmarks (QAMR, QA-SRL, LSOIE)
- OpenIE evaluation (RelVis)

Reliability Question: How do we prove improvement?

---

## Chapter 18: Building a Context Engineering Platform

Topics:

- End-to-end architecture
- Knowledge stores
- Retrieval
- Permissions
- Memory
- Evaluation

Reliability Question: What does a production-ready system look like?

---

# Conclusion: The Context Engineer

The final chapter is no longer about prompts.

It is about a new role emerging from the collision of:

- data engineering
- platform engineering
- distributed systems
- information retrieval
- security
- AI infrastructure

The context engineer is responsible for:

- getting the right information
- to the right model
- at the right time
- with the right permissions
- at the lowest possible cost
- while producing reliable outcomes
