# LLMs in Production — Chapter 6 Notes
## Multi-LLM Architectures, Agent Systems, Tool Use, and Context Engineering

**Source:** *LLMs in Production* (Chapter 6)
**Date:** June 2026
**Research Theme:** Agent Orchestration, Multi-Model Systems, Tool Calling

---

# Key Insight

Chapter 6 introduces multi-LLM architectures and agent systems.

This is where the book's content begins to intersect heavily with Context Engineering.

Key themes:

- routing
- agents
- tools
- orchestration
- state management
- evaluation

All of these are context engineering problems.

---

# Multi-LLM Architectures

## Types

### Specialist Models

Different models for different tasks.

Example:

- Model A: Classification
- Model B: Generation
- Model C: Reasoning

### Router Models

A model decides which model handles a request.

### Ensemble Models

Multiple models contribute to a single response.

---

## Context Engineering Connection

Multi-LLM architectures create new context engineering challenges:

- How does each model receive context?
- How is state shared between models?
- How are authorization policies applied?
- How is retrieval scoped per model?
- How is evaluation performed?

The answer is always: **context engineering**.

---

# Agent Systems

## What Makes Something an Agent?

Chapter 6 defines an agent as a system that:

1. Perceives environment
2. Decides action
3. Executes action
4. Observes result
5. Repeats

This is the classic agent loop.

---

## Agent Loop in LLM Terms

```text
User Input
→ Model
→ Action (Tool Call / Text)
→ Execution
→ Observation
→ Model
→ ...
```

Each iteration requires context assembly.

---

# Context Engineering in Agent Systems

## The Context Assembly Problem

Every agent loop iteration requires:

1. **State Update** — What has happened so far?
2. **Retrieval Update** — What new information is relevant?
3. **Tool Scope** — What tools are available now?
4. **Memory Integration** — What should be remembered?
5. **Authorization Check** — Has permission changed?
6. **Constraint Update** — Have constraints been modified?

This is the core context engineering challenge in agents.

---

# Tool Calling

## Tool Calling as Context

Tools are not just actions.

Tools are context sources.

Example:

```json
{
  "tool": "get_user_orders",
  "result": [...]
}
```

The tool result becomes context for the next inference.

---

## Tool Selection as Context Engineering

Choosing which tools to make available is a context engineering decision.

Providing too many tools:

- Increases ambiguity
- Increases token consumption
- Decreases reliability

Providing too few tools:

- Limits capability
- May prevent task completion

The optimal tool set is task-dependent.

---

## Tool Definition as Context

Tool definitions include:

- name
- description
- parameters
- return types
- examples

All of this becomes context.

Poor tool definitions → poor context → unreliable behavior.

---

# Tool Calling Failure Modes

## 1. Wrong Tool Selected

The model selects an inappropriate tool.

Context engineering solution:

- Scope available tools
- Provide semantic constraints
- Improve tool descriptions

---

## 2. Wrong Parameters

The model calls the right tool with wrong arguments.

Context engineering solution:

- Schema validation
- Parameter constraints
- Type checking

---

## 3. Missing Tool

The model does not call a required tool.

Context engineering solution:

- Explicit tool availability
- Capability declarations
- Tool necessity hints

---

## 4. Tool Not Available

The tool exists but is not in scope.

Context engineering solution:

- Authorization-aware tool availability
- Permission filtering
- Feature flags

---

## 5. Tool Failure

The tool executes but returns an error.

Context engineering solution:

- Error handling context
- Retry context
- Fallback context

---

# Orchestration Patterns

## Sequential Processing

```text
Step 1 → Step 2 → Step 3 → Output
```

Context flows sequentially.

Each step adds to context.

---

## Parallel Processing

```text
Step 1
Step 2 → Output
Step 3
```

Context may be assembled from multiple sources.

Requires context merging logic.

---

## Conditional Routing

```text
Input → Decision → Branch A / Branch B
```

Context determines routing.

Routing determines context.

---

## Iterative / Agent Loop

```text
Input → Model → Action → Observation → Model → ...
```

Context is rebuilt each iteration.

State management becomes critical.

---

# State Management in Agents

## Conversation State

What the user said.
What the model said.
What tools were called.

---

## Execution State

What actions have been taken.
What results were obtained.
What errors occurred.

---

## Business State

What task is being performed.
What entities are involved.
What constraints apply.

---

## State Failure Modes

### Lost State

Context is lost between iterations.

Result: Agent forgets what it was doing.

---

### Corrupted State

Context becomes inconsistent.

Result: Agent makes decisions based on wrong information.

---

### Divergent State

Agent loses track of original intent.

Result: Agent pursues wrong goal.

---

# Evaluation Challenges in Agent Systems

## Trajectory Evaluation

Did the agent take the right path?

Not just:

> Did it produce the right output?

But:

> Did it use the right tools?
> Did it call them in the right order?
> Did it handle errors correctly?

---

## Multi-Step Correctness

A single wrong step can cascade.

Example:

1. Wrong tool selected
2. Wrong data retrieved
3. Wrong decision made
4. Wrong output generated

The failure originated in step 1.

The symptom appeared in step 4.

---

## Context Quality vs Output Quality

Important distinction:

- Output quality: Did the agent succeed?
- Context quality: Did the agent have the right information?

A system can have perfect context and still fail.
A system with poor context will usually fail.

Evaluating context quality is harder but more actionable.

---

# Emergent Pattern: Context as Infrastructure

This chapter reinforces a key theme:

> Context engineering treats context as infrastructure.

In agent systems:

- Context is assembled, not just written
- Context flows between components
- Context has state, not just content
- Context has failures, not just errors
- Context requires evaluation, not just testing

This is the infrastructure view.

---

# Open Questions for the Book

1. How do we evaluate context quality in agent systems?
2. What is the minimum context required for reliable tool calling?
3. How should context be partitioned between agent iterations?
4. How do authorization policies apply in multi-agent systems?
5. What is the relationship between context complexity and agent reliability?
6. How do we debug context failures in agent systems?
7. What patterns emerge from successful agent architectures?

---

# Connection to Other Chapters

This chapter connects to:

- Chapter 3 (retrieval): Tools as retrieval sources
- Chapter 5 (prompting): Tool definitions as pragmatic context
- Chapter 7 (evaluation): Agent trajectory evaluation
- Future chapters on state, memory, and governance

---

# Summary

Agent systems are context engineering systems.

The complexity of agents is the complexity of context assembly, state management, retrieval, authorization, and evaluation.

Building reliable agents requires building reliable context infrastructure.

This is the core thesis of Context Engineering.