# Book Proposal: Context Engineering

## Building Reliable AI Systems from Data to Action

**Author:** Miriah Peterson
**Proposed for:** Joe Reis's media company
**Format:** Practical teaching book with durable reference value
**Estimated length:** 350–400 published pages
**Estimated visuals:** 40–50 diagrams and figures
**Estimated code listings:** 120–150

## Proposal summary

The central problem in AI is no longer access to a capable model. It is giving
that model the right information, state, tools, constraints, and feedback at
the right time—and making the resulting system reliable enough to use in the
real world.

*Context Engineering* is a practical guide to designing the systems around a
model. It explains how to turn raw data into usable context, retrieve the right
evidence, manage memory and state, define tool contracts, govern permissions,
evaluate outcomes, and operate AI workflows in production.

The book’s central argument is simple: model output is only one part of an AI
system. A model can propose an answer, action, or plan, but trusted software
must assemble the context, validate the proposal, enforce authorization, carry
out the effect, and record what happened. Context engineering is the discipline
that connects data engineering to dependable AI behavior.

This is a durable topic. Models, providers, and frameworks will change. The
need to manage information, provenance, freshness, identity, memory, tools,
workflow state, cost, and feedback will not.

## 1. Tell us about the author

### What are your qualifications for writing this book?

This proposal is grounded in the practical engineering problem between model
capability and production usefulness: how data gets selected, how context is
assembled, how agents interact with tools, and how systems are made observable
and safe. The author’s qualifications and publication history should be
expanded here with Miriah Peterson’s specific experience before submission.

The book’s approach is based on a simple observation: most AI failures are not
simply model failures. They are failures of missing information, missing state,
ambiguous instructions, weak retrieval, ungoverned tool use, poor evaluation,
or unreliable workflows. This book provides a unified way to diagnose and
engineer those failures.

### What makes you distinctive as an author?

I work across the boundaries between AI, data engineering, software
architecture, and operations. I am interested in the complete path from a
user’s request to a system outcome:

```text
request -> context assembly -> model proposal -> validation -> effect
       -> state, provenance, evaluation, and feedback
```

That perspective makes the book both conceptual and practical. It gives
readers a vocabulary for understanding why AI systems fail and concrete
patterns for building systems that can recover, improve, and be trusted.

## 2. Tell us about the book’s topic

### What is the technology or idea that you’re writing about?

Context engineering is the design of the information and execution environment
around a model. It includes:

- identifying the information a task actually requires;
- retrieving and ranking evidence from documents, databases, APIs, and graphs;
- preserving provenance, freshness, permissions, and uncertainty;
- managing conversation history, long-term memory, and workflow state;
- designing prompts, examples, structured outputs, and tool schemas;
- validating model proposals before they become actions;
- controlling identity, authority, side effects, retries, and recovery; and
- evaluating the entire system rather than judging model text in isolation.

The book treats context as an engineered system, not as a longer prompt. A
reliable AI application is a pipeline that transforms governed data and user
intent into a bounded, inspectable outcome.

### Why is it important now?

Organizations have moved rapidly from experimenting with chat interfaces to
building AI features, copilots, and agents into products and internal
operations. Many of those systems work in demonstrations and fail in
production because the surrounding engineering is underspecified.

They retrieve irrelevant or stale information. They confuse conversation
history with durable memory. They give models tools without clear contracts or
adequate authorization. They cannot explain which evidence produced an answer,
why an action was taken, or whether a workflow actually completed.

This creates a major opportunity for the data-engineering community. AI does
not make data engineering less important; it exposes how much AI depends on
it. Retrieval is governed data access. Grounding is a provenance problem.
Memory is a state-management problem. Tool use is an interface and
authorization problem. Agents are workflows with probabilistic proposals.

Context engineering gives practitioners a common discipline for solving these
problems.

### What makes it different from its alternatives?

Most AI resources focus on prompting, model selection, or application
frameworks. Those topics are useful, but they do not answer the harder
questions:

- What information should enter the model’s context?
- How do we know that information is authoritative and current?
- What should persist outside the context window?
- How do we prevent a plausible model proposal from becoming an unsafe action?
- How do we measure retrieval quality, context quality, and task success?
- How do we debug a failure across data, retrieval, prompts, tools, state, and
  workflow execution?

This book answers those questions as one connected engineering problem. It
teaches readers to build context pipelines and execution boundaries that remain
useful as models and frameworks change.

## 3. Tell us about the book you plan to write

This book will show readers how to build a reliable AI system from the outside
in. They will begin with a task and its required context, then construct the
data, retrieval, memory, tool, workflow, authorization, and evaluation layers
needed to make that task dependable.

After reading the book, readers will be able to:

- diagnose AI failures as problems of missing information, missing state,
  ambiguous intent, weak retrieval, invalid proposals, or broken workflows;
- define the lexical, semantic, and pragmatic context a task requires;
- build ingestion and transformation pipelines that preserve provenance and
  support reliable retrieval;
- choose among lexical, vector, graph, hybrid, and structured retrieval;
- measure context precision, context recall, retrieval success, freshness, and
  source coverage;
- design memory systems that distinguish conversation history, user facts,
  task state, and durable organizational knowledge;
- create prompts, examples, structured outputs, and tool contracts that make
  model behavior inspectable and rejectable;
- build agents as bounded workflows with explicit state, retries, deadlines,
  checkpoints, and completion criteria;
- enforce least privilege and authorization at the data and tool boundaries;
- trace context assembly, model proposals, validation, tool execution, and
  outcomes end to end;
- evaluate AI systems against real task distributions rather than relying only
  on generic benchmarks; and
- make informed decisions about when to improve context, change the workflow,
  fine-tune a model, or replace a model altogether.

### Is the book designed to teach a topic or to be used as a reference?

It is primarily a teaching book with strong reference value. The reader will
learn a general method for designing AI systems, then apply it through complete
examples and reusable patterns. The material is also well suited to company
training, internal AI-platform programs, and advanced college courses in data
engineering, software architecture, or applied AI.

### Does this book fall into a series?

The book fits a practical, systems-oriented “From Scratch” positioning. It
could also work as a companion to books about data engineering, AI engineering,
or LLMs in production because it addresses the layer that connects those
subjects.

### What are the unique characteristics of the proposed book?

The book will use a consistent visual language to show how context moves
through a system and where reliability boundaries exist. Chapters will include
architecture diagrams, failure traces, decision tables, exercises, and
end-to-end examples.

Diagrams will be created as Mermaid source in the existing example repository
and rendered into PNG and SVG assets for the manuscript, website, and talks.
Flowcharts, sequence diagrams, state machines, retrieval pipelines, and
authorization boundaries will remain editable as text so they can evolve with
the examples. The existing scripts and code examples will reference the same
architecture, making the visuals and implementations consistent rather than
producing one-off artwork disconnected from the book.

Supplementary materials will include:

- public GitHub repositories with runnable examples;
- context contracts, retrieval, evaluation, and observability templates;
- example datasets and reproducible test fixtures;
- worksheets for context budgets, cost, latency, and coverage; and
- optional video lectures covering difficult implementation details.

## 4. Frequently asked questions

### Q1: Isn’t context engineering just prompt engineering?

No. Prompt engineering shapes the instructions given to a model. Context
engineering includes the entire system that decides what the model should see,
what state should persist, what tools are available, what the model is allowed
to propose, and how the result is validated and used.

Prompt quality matters, but a perfect prompt cannot repair missing source data,
stale retrieval, incorrect permissions, lost workflow state, or an unsafe tool
boundary.

### Q2: How do I know what context to provide?

Start with the task contract: the user, purpose, required output, authoritative
sources, freshness requirements, constraints, and success criteria. Then build
a context pipeline that retrieves candidate information, applies policy and
quality checks, and records what was admitted.

The book gives readers practical methods for measuring context precision,
recall, coverage, freshness, and usefulness rather than relying on intuition
about prompt length.

### Q3: How do I give an agent access to tools safely?

Treat every model tool call as a proposal, not as authority. A trusted host
should validate the schema, resolve identity and scope, authorize the exact
operation, enforce limits and idempotency, execute the call, and record the
result.

The book covers tool contracts, least privilege, capability-based access,
sandboxing, approval boundaries, retries, ambiguous outcomes, and auditability.

### Q4: How do I know whether an AI system is improving?

Measure the complete task outcome. A system can produce fluent text while
retrieving the wrong evidence, violating a permission boundary, wasting its
context budget, or failing to complete the workflow.

The book develops an evaluation approach that combines task success, retrieval
quality, tool correctness, authorization decisions, latency, cost, failure
recovery, and human intervention.

## 5. Tell us about the readers

### Primary audience

The primary reader is a technical practitioner responsible for making AI useful
inside a real product or organization. Specifically, the book is for:

- software engineers building AI features and agentic applications;
- data engineers designing ingestion, retrieval, knowledge, and governance
  systems for AI;
- ML and ML-adjacent practitioners who want to improve systems beyond model
  selection and prompt tuning;
- platform, DevOps, and SRE professionals operating AI workloads;
- technical founders building products around proprietary data or workflows;
- product-minded engineers responsible for AI quality and user trust; and
- engineering leaders establishing standards for AI architecture and
  evaluation.

### Reader prerequisites

Readers should be comfortable with intermediate Python, functions and classes,
common libraries, and basic command-line work. They should understand at a
high level what a neural network, training, inference, and model parameters
are.

No prior LLM experience is required. The book will teach the relevant AI
concepts as needed and will not assume advanced mathematics,
metaprogramming, distributed-systems, or database expertise.

### What will motivate readers to learn this topic?

Readers are motivated by the gap between an AI demo and a dependable product:

- answers that sound plausible but are unsupported or stale;
- retrieval systems that return too much irrelevant information;
- agents that lose state or repeat work;
- tools that are difficult to constrain and audit;
- data and permissions that do not survive transformation into context;
- rising inference and operational costs; and
- the need to explain, evaluate, and improve AI behavior.

### Why the audience and topic will last

The audience is responsible for the engineering work that every useful AI
system requires, regardless of which model or provider is currently strongest.
Their problems are rooted in information architecture, data quality,
retrieval, state, software interfaces, security, observability, and workflow
design.

The book’s lasting importance comes from treating models as replaceable
components inside a larger system. It teaches readers how to build the layer
that adapts when models improve, APIs change, organizations add new data, or
workflows become more consequential.

## 6. Tell us about the competition

### What sources of information are available?

- *Building LLM Apps* by Haney (Manning): application development and common
  LLM patterns.
- *LLMs in Production* by Brousseau and Sharp (Manning): production deployment
  and operational concerns. It is a particularly relevant adjacent title,
  while this proposal focuses on the context, data, state, and control layer
  that production systems require.
- *AI Engineering* by Chip Huyen: a broad treatment of building applications
  with foundation models.
- Ontology, knowledge-graph, and semantic data-pipeline books and resources:
  useful foundations for representing meaning and relationships, but generally
  not a complete guide to turning those structures into runtime context,
  bounded tool use, and evaluated workflows.
- Resources on mixed-model and hybrid AI architectures: useful for routing
  different models to different tasks, but typically focused on model
  composition rather than the context and governance layer shared by those
  models.
- Books and documentation focused on LangChain, vector databases, prompt
  engineering, or agent frameworks.
- Research papers, technical blogs, vendor documentation, and conference talks
  covering individual pieces of the stack.

### How does the proposed book compare?

*Context Engineering* fills the gap between “write a better prompt” and “build
  a reliable AI system.” It is:

- **Systematic:** it connects data, semantics, pragmatics, retrieval, memory,
  tools, workflows, security, and operations.
- **Model- and vendor-independent:** the principles apply across providers,
  models, and frameworks.
- **Practical:** every chapter includes runnable code, measurements, or a
  concrete design exercise.
- **Opinionated:** it makes clear recommendations and explains tradeoffs.
- **Durable:** it focuses on boundaries and engineering responsibilities that
  remain after today’s tools are replaced.
- **Relevant to data engineering:** it shows how existing disciplines around
  data quality, lineage, access, schemas, and operations become the foundation
  of AI systems.

The current learning experience is fragmented across prompt examples, vector
database tutorials, ontology and knowledge-pipeline resources, mixed-model
architecture guides, agent frameworks, and production-infrastructure books.
There is no single resource that explains how these pieces should fit together
or how to debug the system when the answer is wrong.

## 7. Book size and illustrations

The manuscript is expected to be approximately 350–400 published pages. It will
include approximately 40–50 figures, including context pipelines, retrieval
flows, memory boundaries, tool contracts, authorization paths, workflow state
machines, evaluation loops, and observability dashboards. Figures will be
maintained as Mermaid source and exported to PNG/SVG for publication and web
use. It will include approximately 120–150 code listings, with runnable
examples maintained in the companion repositories.

## 8. Contact information

**Name:** Miriah Peterson
**Online presence:** To be supplied
**Website:** To be supplied

## 9. Schedule

The target schedule is approximately 18 months from contract to completed
manuscript.

- **Month 1:** Deliver and revise Chapter 1; finalize the example system,
  evaluation plan, and development environment.
- **Months 2–6:** Complete Part I and the first third of the manuscript.
- **Months 7–12:** Complete Parts II and III and the second third of the
  manuscript.
- **Months 13–18:** Complete Part IV, appendices, revisions, figures, and
  companion materials.

There is no single external deadline driving the project. The field is moving
quickly, so the manuscript will emphasize durable architectural principles,
reproducible evaluation methods, and clear decision frameworks. Tool-specific
examples will be maintained with version notes in the companion repositories.

## 10. Table of contents

### Part I: The Context Engineering Mindset

#### Chapter 1: Every AI Failure Is a Context Failure

- Why model quality is only one variable
- Missing information, missing state, and ambiguous intent
- The context-to-outcome pipeline
- A diagnostic framework for AI failures
- What the reader will build

#### Chapter 2: AI Is a Systems Problem

- The production AI stack
- Models, applications, data, and control planes
- Why demos fail at organizational boundaries
- Separating model proposals from trusted effects
- Designing for replacement and change

#### Chapter 3: Attention, Tokens, and Context Windows

- Tokens, embeddings, and attention
- Context windows and the economics of context
- What the model can and cannot “remember”
- Context compression and compaction
- Latency and cost implications

#### Chapter 4: Instructions, Examples, and Structured Meaning

- In-context learning and pragmatic instructions
- Examples as executable context
- Structured outputs and semantic contracts
- Clarifying intent and defining task boundaries
- When prompting stops being enough

### Part II: Data and Retrieval as Context

#### Chapter 5: Tool Use Is Structured Context

- Agents, tools, and action proposals
- Tool schemas and function calling
- Selection, routing, validation, and authorization
- Idempotency, side effects, and receipts
- Designing tools that are safe to reject

#### Chapter 6: Memory Is a Database Problem

- Conversation history versus durable memory
- User, session, task, and workflow state
- Memory write policies and correction
- Expiration, deletion, and provenance
- Persistent state outside the model context

#### Chapter 7: Context Is a Query

- Sources of context
- Ingestion and context assembly pipelines
- Freshness, consistency, and partial failure
- Hydration, coverage, and retrieval success
- Information extraction as derived state

#### Chapter 8: Knowledge Graphs and Semantic Context

- Schemas, taxonomies, and ontologies
- Entities, relationships, and identity resolution
- Graph retrieval and relationship traversal
- Ontology-guided extraction
- Validation, provenance, and knowledge quality

#### Chapter 9: Retrieval Beyond Vector Databases

- Lexical and relational retrieval
- Vector and semantic retrieval
- Graph and hybrid retrieval
- Ranking, reranking, and query planning
- Context precision and context recall

### Part III: Context, Personalization, and Agency

#### Chapter 10: Personalization Is Governed Data Access

- Personalization as retrieval
- Scoped hydration and policy-aware context
- Provenance and derived context
- Multi-tenant isolation
- Coverage, consent, and correction

#### Chapter 11: Stop Giving Agents Permissions

- Least privilege for AI systems
- RBAC, ABAC, and capability-based access
- Scoped credentials and knowledge stores
- Retrieval and execution boundaries
- Authorization coverage and necessary access

#### Chapter 12: The Unix Philosophy of AI Systems

- Small, composable systems
- Pipes, files, and explicit interfaces
- Mounts, namespaces, and isolation
- Task workspaces and secret management
- Replacing black boxes with inspectable boundaries

#### Chapter 13: Agents Are Workflows

- Planning and ReAct
- Harnesses and state machines
- Durable and event-driven execution
- Loops, retries, and bounded autonomy
- Completion criteria and recovery

### Part IV: Operating Context Systems

#### Chapter 14: The Cost of Context

- Token and context economics
- Retrieval, tool, and latency costs
- One-shot execution, loops, and subagents
- Model routing and cost-aware design
- Context-efficiency metrics

#### Chapter 15: When Context Engineering Stops Working

- Diagnosing model limitations
- Fine-tuning and adaptation
- Distillation and specialized models
- Choosing between better context and a different model
- Context engineering as the research phase

#### Chapter 16: Observability for Context Systems

- Tracing context assembly
- Prompt, retrieval, tool, and model lineage
- State, cost, latency, and failure observability
- Protecting sensitive content in telemetry
- Debugging from request to outcome

#### Chapter 17: Evaluating AI Systems

- Evals and benchmarks
- Retrieval and tool evaluation
- Task success and terminal states
- Failure injection and recovery testing
- Human review, calibration, and regression control

#### Chapter 18: Building a Context Engineering Platform

- Source and ingestion architecture
- Semantic and retrieval infrastructure
- Authorization, state, and tooling
- Observability, evaluation, and cost control
- A reference platform and implementation roadmap

### Appendices

- **A. Context Contract Templates**
- **B. Retrieval and Evaluation Checklists**
- **C. Tool, Authorization, and Workflow Schemas**
- **D. Cost, Latency, and Context-Budget Worksheets**

## Closing pitch

The AI industry has taught people how to call models. The next generation of
practitioners needs to learn how to build systems around them.

*Context Engineering* gives readers the practical framework to make AI useful
with real organizational data, explicit state, governed tools, and measurable
outcomes. Its immediate value is helping teams fix unreliable AI systems. Its
lasting value is defining the engineering layer that connects data to model
behavior and model behavior to trustworthy action.
