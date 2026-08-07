# Toolformer, Tool Calling, and the Economics of Context — Research Notes

## Primary Source

- **Toolformer: Language Models Can Teach Themselves to Use Tools**
  - Meta AI Research
  - 2023
  - https://arxiv.org/abs/2302.04761

## Related Sources

- Language Models are Few-Shot Learners — Brown et al. (2020)
  - https://arxiv.org/abs/2005.14165
- *LLMs in Production: From Language Models to Successful Products* — Christopher Brousseau and Matthew Sharp (Manning, 2024)
  - https://www.manning.com/books/llms-in-production

*Reading Context: Context Engineering Book Research — Chapter 5: Tool Use Is Structured Context*
*Reading List Reference: Section 3 — LLM Foundations*

---

# Key Insight

Toolformer is often described as a tool-use paper.

A better interpretation is:

> Toolformer is a context acquisition paper.

The paper demonstrates that language models can learn:

1. when they lack information
2. when they lack capability
3. which external system can provide that capability
4. how to incorporate the result back into generation

The tool is not the important part.

The retrieval of missing context is.

---

# Tools as Capability Extension

LLMs are not general-purpose computers.

They are probabilistic token predictors.

There are many tasks they perform poorly:

- arithmetic
- counting
- current information retrieval
- time calculations
- deterministic business logic
- stateful operations
- external actions

Toolformer proposes a simple solution:

> Delegate work to systems that already solve the problem.

Examples:

| Capability Gap | Tool |
|---------------|------|
| Arithmetic | Calculator |
| Current Information | Search |
| Time | Clock |
| Data Access | Database |
| Files | Filesystem |
| Actions | APIs |
| Business Logic | External Services |

---

# Tool Calls Are Retrieval

One of the most important observations from this study session:

> Tool calls are retrieval.

This broadens retrieval beyond vector databases.

Retrieval includes:

- SQL queries
- graph traversals
- filesystem access
- search engines
- web search
- APIs
- calculators
- clocks

Any external system that provides information to the model is functioning as retrieval.

This supports the broader thesis:

> RAG is not a vector database.
>
> RAG is any process that retrieves external information and injects it into context.

---

# Context Engineering and Tool Design

Tool design is part of context engineering.

When designing an AI system, one of the primary questions becomes:

> What capabilities should exist as tools?

The model itself will never reliably solve:

- current state problems
- deterministic calculations
- authorization decisions
- large-scale retrieval

Those capabilities should exist as external systems.

The model becomes an orchestrator.

The system performs the work.

---

# Context Is Not Free

A major realization from this study session:

> Every piece of context has a cost.

This remains true regardless of hosting model.

---

## API Models

Cost is easy to measure:

```text
$/token
```

---

## Self-Hosted Models

The cost changes form:

- latency
- GPU time
- RAM
- KV cache growth
- network transfer
- energy consumption

Tokens may appear free.

Context is not.

---

# Context Cost Model

Potential framework:

```text
Context Cost =
    Retrieval Cost +
    Injection Cost +
    Attention Cost +
    Latency Cost +
    Governance Cost
```

---

## Retrieval Cost

Cost of obtaining information.

Examples:

- SQL query
- graph traversal
- search request
- API call

---

## Injection Cost

How many tokens are added.

Examples:

- current date
- customer record
- retrieved document

---

## Attention Cost

The model must process the information.

More context means:

- more compute
- larger KV cache
- longer inference

---

## Latency Cost

Time added to user response.

Includes:

- retrieval
- serialization
- prompt construction
- inference

---

## Governance Cost

Cost of:

- authorization
- provenance
- lineage
- auditing

---

# Context ROI

Potential metric:

```text
Context ROI =
    Improvement Produced
    --------------------
      Context Cost
```

Examples:

| Context | Cost | Benefit |
|----------|----------|----------|
| Current Date | Low | High |
| User Profile | Medium | High |
| Entire Wiki Page | High | Low |
| Full Database Dump | Extreme | Low |

---

# Tool Calls Have Context Cost

Every tool call adds:

- tokens
- latency
- retrieval work
- context window usage

This means:

> More tools is not always better.

The important question becomes:

> Which tools create enough value to justify their context cost?

---

# Tool Calling and Training

Modern models are explicitly trained for tool use.

Tool calling is no longer purely emergent behavior.

However:

> Models are trained on specific tool call formats.

Not all tool formats.

---

# Zero-Shot Tool Calling

A model can perform zero-shot tool calling when the schema matches its training distribution.

Examples:

- OpenAI-style JSON
- Anthropic XML variants
- Gemini function calling
- Qwen schemas
- MiniMax schemas

Benefits:

- lower context cost
- fewer examples
- fewer errors
- higher reliability

---

# Few-Shot Tool Calling

If a harness uses a different schema than the model was trained on:

> Tool use becomes a few-shot learning problem.

The model must learn the interface from context.

Additional context is required:

- instructions
- examples
- schemas
- demonstrations

This increases:

- prompt size
- latency
- context consumption

---

# Native Format vs Standardized Format

## Native Format

Pros:

- lower token cost
- fewer examples
- better performance on small models

Cons:

- harder model switching
- harness complexity

---

## Standardized Format

Pros:

- portability
- orchestration simplicity
- model independence

Cons:

- more prompting
- larger context windows
- more formatting failures

---

# Quantized Models Reveal This First

Large frontier models often recover from formatting mistakes.

Smaller models frequently do not.

Observed example:

MiniMax (trained on XML-like formats) occasionally attempts tool calls using XML when the harness expects JSON.

Example:

```xml
<tool_call>
  <name>bash</name>
  <arguments>ls</arguments>
</tool_call>
```

instead of:

```json
{
  "tool": "bash",
  "arguments": {
    "command": "ls"
  }
}
```

This appears to be a reasoning failure.

In reality:

> It is often a schema mismatch.

---

# Pedro Agentware Pattern

Tool formatting should be handled by the harness.

The harness should support:

- adapters
- templates
- schema validation
- deterministic repair

The model should not be forced to relearn interface conventions.

String formatting is already a solved problem.

---

# Tool Calls vs LoRA

A major discussion point from Christopher Schauer.

Question:

> When should we use context engineering and tool calls versus refinement training?

---

# Context Engineering as Research

Context engineering is an exploration process.

It helps discover:

- required context
- retrieval patterns
- workflows
- prompts
- tool usage
- evaluation criteria

The goal is to identify:

> What actually works?

---

# LoRA as Specialization

LoRA becomes valuable when:

- workflows stabilize
- outputs stabilize
- evaluations stabilize
- retrieval patterns stabilize

At that point:

> Runtime context engineering stops producing significant gains.

---

# Reliability Ladder

## Stage 1

Prompting

```text
User → Prompt → Model
```

---

## Stage 2

Context Engineering

```text
User
↓
Retrieval
↓
Tools
↓
Scaffolding
↓
Model
```

---

## Stage 3

Refinement Training

```text
User
↓
Specialized Model
↓
Output
```

---

# The Context Engineering Limit

Important principle:

> There is a point where context engineering stops winning.

Signs:

- repeated prompts
- repeated retrieval
- repeated instructions
- stable workflows
- stable outputs

At that point:

> The behavior should be moved into the model.

---

# Storage, Retrieval, and Memory

A major realization:

> There is no memory system to rule them all.

Different workloads require different retrieval architectures.

---

# Markdown as the Current Hammer

Historically:

```text
Postgres = hammer
```

Increasingly:

```text
Markdown = hammer
```

Modern agent systems frequently use:

- markdown files
- JSON files
- filesystem traversal

because:

- easy
- human-readable
- git-friendly
- model-friendly

---

# Why Markdown Works

Benefits:

- zero infrastructure
- easy inspection
- agent compatibility
- low setup cost

These are legitimate advantages.

---

# When Markdown Stops Working

Failure modes:

## Retrieval

Too many files.

## Relationships

You accidentally build a graph.

## Permissions

Authorization becomes difficult.

## Analytics

You start rebuilding a warehouse with grep.

## Freshness

Dynamic state becomes painful.

---

# Discovery vs Access

An important distinction.

## Discovery

Question:

> What should I look at?

Solutions:

- vector search
- keyword search
- graph traversal

---

## Access

Question:

> Give me the thing.

Solutions:

- databases
- caches
- key-value stores

---

# Retrieval Architecture Matters More Than Storage

A better framing:

> There is no retrieval architecture to rule them all.

Different access patterns require different systems.

Examples:

| Pattern | System |
|----------|----------|
| Point Lookup | Database |
| Similarity Search | Vector Store |
| Traversal | Graph |
| Exploration | Filesystem |
| Massive Scan | S3 / Iceberg |
| Stateful Execution | Tool |

---

# S3 Is Not Slow

Common misconception.

S3 is slow for:

- transactional lookups
- row-level access

S3 is excellent for:

- large scans
- parquet streaming
- analytical workloads

The question is not:

> Is S3 fast?

The question is:

> Fast for what workload?

---

# Hammer vs Scalpel

The key storage question for the book:

> When should you use the hammer?
>
> When should you use the scalpel?

Start with:

- markdown
- filesystem
- postgres
- duckdb

Introduce specialized systems only when:

- access patterns justify them
- context cost justifies them
- operational complexity is worth it

---

# Emerging Thesis

Reliable AI systems are not built by maximizing context.

Reliable AI systems are built by:

- optimizing context value
- minimizing context cost
- choosing appropriate retrieval architectures
- introducing specialization only when justified

Context engineering is ultimately:

> The discipline of delivering the right information, at the right time, for the lowest total system cost.

---

# Claims for Evidence Ledger

## Claim: Tool Calls Are Retrieval

Tool calls extend the definition of retrieval beyond vector databases to any external system that provides information to the model.

**Source:** toolformer (Meta AI, 2023)

**Supports:** Chapter 5 — Tool Use Is Structured Context; Chapter 7 — Context Is a Query Over Distributed State

**Strength:** strong

---

## Claim: Context Has Cost

Every piece of context has a measurable cost in tokens, latency, and compute.

**Source:** toolformer (Meta AI, 2023)

**Supports:** Chapter 14 — The Cost of Context

**Strength:** strong

---

## Claim: Zero-Shot Tool Calling

Models can perform zero-shot tool calling when the schema matches their training distribution, reducing context requirements.

**Source:** toolformer (Meta AI, 2023)

**Supports:** Chapter 5 — Tool Use Is Structured Context

**Strength:** strong

---

## Claim: Context Engineering Precedes Model Training

Context engineering is an exploration process that identifies what works before moving behavior into the model via LoRA or fine-tuning.

**Source:** toolformer (Meta AI, 2023)

**Supports:** Chapter 15 — When Context Engineering Stops Working

**Strength:** strong

---

# Gaps Identified

This source does **not** support:

- Specific authorization patterns
- Knowledge graph architecture
- Observability frameworks
- DDSO implementation

These require additional sources from:

- Unix research
- Semantic web research
- OpenTelemetry research

---

# Potential Book Quote

> Toolformer reveals that tool calls are not a feature of the model—they are a mechanism for context acquisition.
>
> The model learns to recognize when it lacks capability and delegates to external systems.
>
> This is retrieval, extended beyond databases to any system that can provide missing information.
