# LLMs in Production — Chapter 7 Notes
## Evaluation, Hallucination, Grounding, and Context Engineering

**Source:** *LLMs in Production* (Chapter 7)
**Date:** June 2026
**Research Theme:** Evaluation, Hallucination, Grounding, Reliability

---

# Key Insight

Chapter 7 focuses on evaluation and the hallucination problem.

This is where Context Engineering connects most directly to measurable outcomes.

Key themes:

- evaluation types
- hallucination sources
- grounding techniques
- reliability metrics
- benchmarking

All require context engineering to solve.

---

# Hallucination as Symptom

## Not a Model Problem

Industry often treats hallucination as:

> A property of the model.

This is misleading.

Hallucination is a symptom of:

- missing context
- incorrect context
- insufficient grounding
- retrieval failures
- state failures

---

## Hallucination Categories

### Factual Hallucination

Model generates false facts.

Context engineering cause:

- No retrieval of facts
- No grounding in verified sources

---

### Logical Hallucination

Model generates internally inconsistent statements.

Context engineering cause:

- No consistency constraints
- No state verification

---

### Instructional Hallucination

Model ignores instructions in context.

Context engineering cause:

- Conflicting instructions
- Ambiguous constraints
- Poor instruction formatting

---

### Referential Hallucination

Model misattributes information to wrong source.

Context engineering cause:

- No provenance tracking
- No citation context

---

# Grounding

## What Is Grounding?

Grounding means connecting model outputs to:

- verified facts
- authoritative sources
- real-world state

---

## Grounding Techniques

### Retrieval Grounding

Inject relevant facts into context.

```text
Query → Retrieve Facts → Inject → Generate
```

This is the RAG pattern.

---

### Constraint Grounding

Limit output space through semantic constraints.

```text
Allowed Values: [A, B, C]
Output must be in allowed values.
```

---

### State Grounding

Verify outputs against real system state.

```text
Model Output: "User has 3 orders"
→ Verify against database
→ Confirm or correct
```

---

### Source Grounding

Provide citations or references.

```text
Generate → Add citations → Output
```

---

# The Context Engineering Position on Hallucination

> Hallucination is a retrieval failure, not a generation failure.

If the model had the right information, it would use it.

The problem is not the model.

The problem is the system's failure to provide the right information.

---

# Evaluation Types

## Generation Evaluation

Did the model produce a good output?

Metrics:

- BLEU
- ROUGE
- BERTScore
- LLM-as-judge

Problem:

These measure output quality.

They do not measure context quality.

---

## Retrieval Evaluation

Did the system retrieve the right information?

Metrics:

- Precision
- Recall
- F1
- MRR
- NDCG

Problem:

These measure retrieval quality.

They do not measure downstream impact.

---

## Agent Trajectory Evaluation

Did the agent take the right path?

Metrics:

- Task completion rate
- Tool call accuracy
- Step-by-step correctness
- Error recovery rate

Problem:

Expensive to evaluate.

Requires annotated trajectories.

---

## Context Quality Evaluation

Does the context enable reliable generation?

This is the missing evaluation layer.

---

# Context Quality Metrics

## Relevance

Is the retrieved information relevant to the task?

```text
Task: Answer question about X
Context: Contains information about X?
```

---

## Completeness

Does the context contain all required information?

```text
Task: Generate report
Context: Contains all required data fields?
```

---

## Accuracy

Is the information correct?

```text
Retrieved Fact: "Company founded in 2020"
Ground Truth: "Company founded in 2019"
→ Context inaccuracy
```

---

## Freshness

Is the information up-to-date?

```text
Task: Current stock price
Context: Stock price from yesterday
→ Stale context
```

---

## Authorization

Does the context respect permissions?

```text
User: Regular user
Context: Includes admin-only data
→ Authorization failure
```

---

## Provenance

Can the source of information be traced?

```text
Output: "X is true"
Context: Which source proved X?
→ Provenance tracking
```

---

# Benchmarking

## Common Benchmarks

### MMLU

Massive Multitask Language Understanding.

Tests general knowledge.

---

### HumanEval

Code generation benchmark.

---

### TruthfulQA

Tests tendency to repeat falsehoods.

---

### HotpotQA

Multi-hop reasoning benchmark.

---

## Benchmark Limitations

Benchmarks evaluate:

- Model capabilities
- Prompting strategies

They do not evaluate:

- Retrieval quality
- Context assembly
- Authorization enforcement
- State management

Context engineering requires different benchmarks.

---

# Building Evaluation Harnesses

## What Is a Harness?

A harness is a system for:

- Running tests
- Measuring outcomes
- Comparing approaches

---

## Evaluation Dimensions

### Functional

Does it work?

```text
Task completion
Output correctness
```

### Reliability

Does it work consistently?

```text
Failure rate
Error handling
Recovery
```

### Efficiency

Does it work efficiently?

```text
Latency
Token usage
Cost
```

### Safety

Does it work safely?

```text
Authorization
Guardrails
PII handling
```

---

## Continuous Evaluation

### Build-Time Evaluation

- Unit tests
- Integration tests
- Retrieval tests

### Deployment-Time Evaluation

- A/B testing
- Canary releases
- Shadow mode

### Production-Time Evaluation

- Monitoring
- Alerts
- Logging

---

# Context Engineering and Evaluation

## The Feedback Loop

```text
Context → Generate → Evaluate → Improve Context → ...
```

Evaluation tells us if context engineering is working.

---

## Evaluation-Driven Development

1. Define success criteria
2. Build evaluation harness
3. Measure baseline
4. Improve context
5. Re-evaluate
6. Repeat

---

# Key Insight: Context Engineering Is Testable

Unlike model behavior, context engineering is testable.

We can test:

- Retrieval quality
- State management
- Authorization enforcement
- Constraint application
- Context assembly

This makes context engineering an engineering discipline.

Model engineering remains partially empirical.

---

# Open Questions for the Book

1. What should context quality benchmarks measure?
2. How do we correlate context quality with output quality?
3. What is the minimum evaluation for a production context engineering system?
4. How do we evaluate context engineering across different task types?
5. What are the key metrics for each pillar?
6. How do we build continuous evaluation for context engineering?

---

# Connection to Other Chapters

- Chapter 3 (retrieval): Retrieval evaluation
- Chapter 5 (prompting): Prompt evaluation
- Chapter 6 (agents): Agent evaluation
- Future chapters: Reliability engineering

---

# Summary

Hallucination is a context failure.

Evaluation is the mechanism for measuring context engineering success.

Context engineering makes AI systems testable, measurable, and improvable.

This is what distinguishes it from prompt engineering.