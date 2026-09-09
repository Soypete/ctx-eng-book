# Context Engineering: Building Reliable AI Systems Through Data, State, and Constraints

**Subtitle:** Building Reliable AI Systems from Data to Action

**Author:** Miriah Peterson
**Author title and affiliation:** CEO, Haikei Labs
**Author pronouns:** To be supplied

**Format:** Practical teaching book with durable reference value
**Estimated length:** 350–400 published pages
**Estimated visuals:** 40–50 diagrams and figures
**Estimated code listings:** 120–150

> Private submission fields—including mailing address, phone number, preferred
> email, nationality disclosure, contracting entity, and pronouns—should be
> completed by the author in the final O’Reilly submission document.

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

## About the author

### Author biography

Miriah Peterson is CEO of Haikei Labs, where she is building infrastructure
for reliable, governed AI systems. She is a data engineer and educator focused
on AI infrastructure, context engineering, and distributed systems, with
experience building production systems across startups and enterprise
technology companies.

She created SoyPete Tech, teaches for Boot.dev and O’Reilly, and hosts the
Domesticating AI podcast, which explores homelabs, local AI, open models, and
the infrastructure required to run AI outside of hosted platforms.

Peterson is the right person to write this book because her work connects the
data and infrastructure foundations of AI with the practical realities of
teaching engineers how to build and operate these systems. Her public work is
available through [SoyPete Tech](https://soypetetech.substack.com/),
[LinkedIn](https://www.linkedin.com/in/miriah-peterson-tech), and the
[Domesticating AI podcast](https://www.listennotes.com/podcasts/domesticating-ai-soypete-tech-hOU7ncLIV42/).

### What are your qualifications for writing this book?

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

## Marketing description

AI systems rarely fail because a model cannot produce fluent text. They fail
because the system supplied the wrong information, lost important state,
exposed the wrong tool, ignored authorization, or had no way to tell whether
the task actually succeeded. *Context Engineering* shows engineers how to
build the reliable layer around a model: governed data, retrieval, memory,
structured tools, bounded workflows, validation, observability, and evaluation.

Unlike books focused primarily on prompting, model selection, or a single
application framework, this book follows the complete path from a user request
to a trustworthy outcome. It gives readers a durable, vendor-independent
vocabulary—Lexicon, Semantics, and Pragmatics—for diagnosing failures and
designing systems that remain useful as models and frameworks change.

## About the topic

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

## Audience

### Level

Beginner to intermediate. The book assumes practical programming experience,
but no prior LLM or advanced mathematics background.

### Who is the book for?

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

### How will readers use the book?

Readers can work through the book sequentially to learn the context-to-outcome
method, then return to individual chapters as an architecture and debugging
reference. A data engineer may use the retrieval and provenance chapters while
designing an ingestion pipeline; an application engineer may use the tool,
authorization, and workflow chapters before shipping an agent; and an
engineering leader may use the evaluation, cost, and platform chapters to set
team standards. The material is also suited to company training and advanced
courses in data engineering, software architecture, or applied AI.

### Estimated market and adoption

The addressable audience includes software and data engineers building AI
features, platform and SRE teams operating AI workloads, technical founders,
and engineering leaders standardizing AI architecture. Adoption is being
driven by the rapid incorporation of foundation models, retrieval systems,
copilots, and agents into products and internal workflows. O’Reilly can add
current analyst and platform-adoption statistics during acquisition; the book
itself will avoid tying its durable thesis to a single market forecast.

### Keywords

Context engineering, AI engineering, LLM applications, generative AI,
retrieval-augmented generation, RAG, prompt engineering, agents, agentic
workflows, tool calling, function calling, structured outputs, memory,
knowledge graphs, ontologies, semantic search, hybrid retrieval, data
engineering, data lineage, provenance, authorization, least privilege,
observability, evaluation, reliability, inference cost, local AI, and open
models.

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

## What the reader will learn—and how to apply it

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

## Other book features

### Code repository

Yes. Code samples, Mermaid diagrams, datasets, fixtures, context contracts,
retrieval examples, evaluation harnesses, and observability templates will be
maintained in public GitHub repositories. Examples will be versioned and
annotated when provider APIs or model behavior changes.

### O’Reilly sandbox

To be discussed with the editor. The examples are designed to run locally and
could be adapted to an O’Reilly sandbox. The manuscript’s primary examples are
expected to use Python and interoperable services; .NET is not the proposed
primary environment.

### Software dependencies

Provider APIs, open-weight models, embedding and reranking models, vector and
graph stores, Python libraries, and agent frameworks may change during
production. The book will minimize framework-specific coupling, pin examples
where practical, provide local alternatives, and document tested versions in
the companion repositories. The core designs rely on portable interfaces,
schemas, HTTP, SQL, and ordinary command-line tools.

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

## Competing titles

### Print and adjacent competition

- *AI Engineering* by Chip Huyen (O’Reilly, 2025, ISBN 978-1-098-16630-4): a
  broad guide to building applications with foundation models. *Context
  Engineering* is narrower and more opinionated about the data, context,
  authorization, workflow, and evaluation boundaries around those models.
- *LLMs in Production* by Christopher Brousseau and Matt Sharp (Manning, 2025,
  ISBN 978-1-63343-720-3): a closely
  related production and operational title. *Context Engineering* focuses on
  the context-to-outcome layer—retrieval, state, semantics, tools, and governed
  execution—that determines what production systems can safely do.
- *Building Reliable AI Systems* by Rush Shahani (Manning, ISBN
  978-1-63343-673-2): a related architecture and reliability title.
  *Context Engineering* centers the data, semantic context, retrieval,
  authorization, and workflow state that make reliability possible.
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

### Related O’Reilly titles

- *AI Engineering* by Chip Huyen: broad foundation-model application
  engineering; this proposal concentrates on context and control boundaries.
- *Designing Data-Intensive Applications* by Martin Kleppmann: durable data
  systems foundations; this proposal applies those ideas to runtime AI context,
  probabilistic proposals, and governed effects.
- *Hands-On Large Language Models* by Jay Alammar and Maayan Salameh:
  model and application foundations; this proposal emphasizes the surrounding
  system’s reliability and authorization responsibilities.

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

## Book outline

The detailed chapter-by-chapter outline follows. Each chapter will end with
three to five review questions, and the book will conclude with a 20-question
final quiz unless otherwise agreed with the acquisitions editor.

### Tech reviewers

Potential technical and target-reader reviewers will be supplied with the
submission. Suggested categories include data-engineering leaders, AI
platform engineers, security and reliability practitioners, and engineers who
build retrieval-augmented and agentic applications.

## Specs and schedule

The manuscript is expected to be approximately 350–400 published pages. It will
include approximately 40–50 figures, including context pipelines, retrieval
flows, memory boundaries, tool contracts, authorization paths, workflow state
machines, evaluation loops, and observability dashboards. Figures will be
maintained as Mermaid source and exported to PNG/SVG for publication and web
use. It will include approximately 120–150 code listings, with runnable
examples maintained in the companion repositories.

### Delivery schedule

- **Two draft chapters:** Month 3
- **Half draft manuscript:** Month 9
- **Full draft manuscript ready for technical review:** Month 15
- **Final manuscript ready for production:** Month 18

## Contact information

**Name:** Miriah Peterson
**Online presence:** [SoyPete Tech](https://soypetetech.substack.com/),
[LinkedIn](https://www.linkedin.com/in/miriah-peterson-tech), and
[Domesticating AI](https://www.listennotes.com/podcasts/domesticating-ai-soypete-tech-hOU7ncLIV42/)
**Website:** [SoyPete Tech Linktree](https://linktr.ee/soypete_tech)

### Schedule notes

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

## Table of contents

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
