# LLMs in Production — Chapter 5 Notes
## Prompt Tuning, Fine-Tuning, Pragmatics, and Context Engineering

**Source:** *LLMs in Production* (Chapter 5)
**Date:** June 2026
**Research Theme:** Context Engineering vs Prompt Engineering vs Model Engineering

---

# Key Quote

> "Language modeling techniques all operate with the underlying assumption that the language model, given inputs and expected outputs, can divine the task to be done and do it in the best way within the number of parameters specified."

## Why This Matters

This quote captures the fundamental assumption behind early prompt engineering:

> If we provide the right examples and instructions, the model can infer the task and perform it correctly.

The production AI experience has shown that this assumption breaks down as systems become more complex.

The model is not simply inferring a task.

It is simultaneously inferring:

- what task is being performed
- what information matters
- what constraints apply
- what permissions exist
- what tools are available
- what output format is expected
- what success looks like

Every one of these is another inference burden placed on the model.

This observation leads directly to the concept of Context Engineering.

---

# Prompt Engineering vs Context Engineering

## Prompt Engineering

Prompt engineering assumes:

> The model can infer the correct behavior from instructions and examples.

Examples:

- role prompts
- system prompts
- few-shot examples
- chain-of-thought prompts
- output instructions

The burden remains on the model to infer the task.

---

## Context Engineering

Context engineering assumes:

> Inference is expensive and unreliable.

Instead of asking the model to infer everything, we explicitly provide operational parameters.

Examples:

- retrieval
- memory
- authorization
- semantic constraints
- tool availability
- state management
- evaluation

The goal is to reduce uncertainty before inference begins.

---

# Context Engineering as Parameter Supply

One useful framing:

> Context Engineering is the process of supplying operational parameters to the model.

Not neural-network parameters.

Operational parameters.

Examples:

| Parameter Type | Example |
|---------------|----------|
| State | User profile |
| Retrieval | Relevant documents |
| Authorization | Accessible resources |
| Tools | Available capabilities |
| Constraints | Valid actions |
| Output Requirements | Schema definitions |
| Memory | Retrieved state |

The more parameters that can be supplied deterministically, the less inference the model must perform.

---

# Computational Pragmatics

Chapter 5 discusses prompt tuning as a mechanism for providing pragmatic context.

This connects directly to the Computational Pragmatics paper.

## Pragmatics

Pragmatics is:

> The study of how meaning is derived from context rather than literal language.

Humans constantly infer intent from context.

LLMs do the same thing.

Examples, instructions, demonstrations, and previous interactions all act as pragmatic context.

---

## Few-Shot Learning as Pragmatic Context

Example:

```text
Ticket: Cannot login
Category: Authentication

Ticket: Credit card declined
Category: Billing

Ticket: User reports missing invoices
Category:
```

The examples teach:

- the task
- valid categories
- expected output structure

Few-shot learning works because it supplies pragmatic context.

---

# Important Distinction

Prompt engineering supplies pragmatic context through text.

Context engineering supplies pragmatic context through systems.

Instead of:

```text
You are an expert support agent...
```

The system supplies:

```json
{
  "customer_tier": "enterprise",
  "valid_categories": [...],
  "product": "SchoolAI"
}
```

Prompt engineering treats pragmatics as language.

Context engineering treats pragmatics as data.

---

# Prompt Engineering vs Prompt Tuning vs Fine-Tuning

One major source of confusion is terminology.

---

## Prompt Engineering

Changes:

- instructions
- examples
- prompts

Does NOT change model weights.

Occurs at inference time.

---

## Few-Shot Prompting

Changes:

- examples in context

Does NOT change model weights.

Occurs at inference time.

---

## Prompt Tuning (ML Research Meaning)

Changes:

- learned prompt embeddings

Does NOT change model weights.

Occurs during training.

Prompt embeddings are optimized mathematically.

Humans do not write them.

---

## Fine-Tuning

Changes:

- model weights

Occurs during training.

Knowledge becomes part of the model.

Examples:

- SFT
- LoRA
- QLoRA
- instruction tuning

---

# Why This Creates Confusion

Many practitioners use "prompt tuning" to mean:

> Improving prompts through experimentation.

Researchers use prompt tuning to mean:

> Learning prompt embeddings while freezing model weights.

These are not the same thing.

---

# Critical Boundary for the Book

This discussion revealed an important scope boundary.

## Context Engineering Ends Where Model Engineering Begins

Context engineering operates on inputs.

Model engineering operates on weights.

---

## Context Engineering

Includes:

- prompt construction
- retrieval
- memory
- authorization
- semantic constraints
- tool selection
- state management
- orchestration
- evaluation harnesses
- guardrails

No model changes required.

---

## Model Engineering

Includes:

- fine-tuning
- instruction tuning
- LoRA
- QLoRA
- RLHF
- DPO
- prompt tuning
- knowledge distillation

Model behavior is modified through training.

---

# Reliability Hierarchy

## Layer 1 — Context Engineering

Modify the environment around the model.

Examples:

- retrieval
- memory
- tools
- state
- semantic constraints

---

## Layer 2 — Model Engineering

Modify model behavior.

Examples:

- fine-tuning
- distillation
- prompt tuning

---

## Layer 3 — Model Architecture

Modify the model itself.

Examples:

- transformers
- MoE
- recurrent architectures
- multimodal architectures

---

# Core Thesis

Most production teams should spend most of their effort in Layer 1.

Many failures blamed on the model are actually failures of:

- retrieval
- memory
- state management
- authorization
- tool routing

These are context failures.

Not model failures.

---

# Guardrails vs Context Engineering

Another important distinction.

---

## Guardrails Restrict Actions

Guardrails are deterministic software systems.

Examples:

- RBAC
- OAuth
- schema validation
- approval workflows
- rate limiting

Guardrails prevent classes of behavior.

---

## Context Engineering Restricts Possibilities

Context engineering changes probabilities.

Examples:

- scoped retrieval
- ontology hydration
- semantic constraints
- memory selection
- context assembly

Context engineering cannot guarantee correctness.

It can only make incorrect behavior less likely.

---

# Reliability Through Uncertainty Reduction

A useful framing for the entire book:

> Reliability is uncertainty reduction.

Each pillar reduces uncertainty before inference.

---

## State

What information is available?

---

## Retrieval

Which information is selected?

---

## Semantic Constraints

How is information structured?

---

## Authorization & Governance

What information is allowed?

---

## Evaluation & Reliability

Did the system improve outcomes?

---

# Context Engineering and Token Efficiency

A major insight from this session:

Token efficiency is not a pillar.

Token efficiency is a consequence of good context engineering.

---

## Principle

> Every token should have a job.

If a token does not reduce uncertainty, why is it present?

---

# Sources of Token Waste

## 1. Redundant Pragmatic Context

Examples:

- repeated instructions
- repeated examples
- repeated tool descriptions

---

## 2. Historical Baggage

Examples:

- transcript replay
- failure replay
- reflection loops
- unbounded scratchpads

---

## 3. Unscoped Retrieval

Examples:

- irrelevant tools
- irrelevant documents
- excessive memory
- excessive schema injection

---

# Observation

Most token waste is actually a retrieval problem.

Not a model problem.

The system is forcing the model to perform:

- retrieval
- authorization
- routing
- memory selection

inside the context window.

Those tasks should be performed before inference.

---

# Practical Token Efficiency Principles

1. Every token should reduce uncertainty.
2. Retrieve before you prompt.
3. Scope tools before inference.
4. Summarize state instead of replaying transcripts.
5. Store memory as data, not conversations.
6. Remove instructions the model already understands.
7. Compress semantics, not meaning.
8. Eliminate retrieval work from the context window whenever possible.

---

# Potential Book Thesis Statement

> Prompt engineering assumes the model can infer the correct task from instructions.
>
> Context engineering assumes inference is expensive and unreliable, so the system should remove ambiguity before inference begins.

---

# Connection to DDSOs

DDSOs are not primarily a token optimization strategy.

They are a relevance optimization strategy.

Token savings emerge as a side effect.

Instead of injecting:

- entire databases
- large documents
- massive context windows

DDSOs provide:

- relevant entities
- relationships
- constraints
- authorized state

This improves:

- reliability
- security
- explainability
- cost

simultaneously.

---

# Final Conclusion

The central insight from this session is that Context Engineering should be treated as a systems discipline focused on reducing uncertainty before inference occurs.

The goal is not to make the model deterministic.

The goal is to construct an information environment in which incorrect behavior becomes less likely.

Prompt engineering modifies instructions.

Model engineering modifies weights.

Context engineering modifies the environment around the model.

That distinction should remain a hard boundary throughout the book.