# In-Context Learning: How Models Learn from Context

## Beat 1: Speech Act Interpretation

Jurafsky's computational pragmatics identifies four inferential problems [computational-pragmatics-notes, lines 15-19]:
1. Reference resolution — what does "it" refer to?
2. Speech act interpretation — what did the user mean?
3. Discourse coherence — how does this relate to what was said?
4. Abduction — what explains this utterance?

The most common failure mode isn't missing information — it's **speech act interpretation**. The model doesn't understand what the user wants DONE with the information provided.

Of the four inferential problems, speech act interpretation has the highest impact in production AI systems — getting this wrong leads to wrong tool calls, incorrect actions, and unreliable outcomes.

This is why frameworks like ReAct [arXiv:2210.03629], chain-of-thought prompting [arXiv:2201.11903], and planning modes (like Claude Code's plan mode [code.claude.com/docs]) exist. They're not adding intelligence — they're adding structure that helps the model extract the user's actual intent when the user hasn't efficiently communicated it.

A user says "fix the API" but means "find the bug in the auth handler." The model sees "fix the API" and guesses. ReAct forces the model to first reason about intent before taking action.

```python
# Example: ReAct loop — see examples/chapter-04-in-context-learning/01-react-loop.py
# Shows: Think → Act → Observe → Repeat forces intent extraction
```

## Beat 2: Why More Inference Means Less Reliability

In machine learning, inference is the act of reaching a conclusion. The problem: when humans communicate, there's a **shared background context** — from prior conversations, collaborative thinking, life experience. We build understanding together.

Models don't have shared context. They have a transformer — a mathematical mechanism for pattern matching, not shared understanding.

When harnesses introduce `/goal` or explicit goal definition, they're doing something profound: they're removing the burden of figuring out *how* to do something and instead focusing the model on *what* to do. The model no longer needs to infer the approach — it's given a target. This produces outcomes more aligned with user expectations.

This is the core insight: **more inference by the model means less reliability**. Context engineering flips this — we do the inference work so the model can just extract.

```python
# Example: Goal-directed vs inference-heavy — see examples/chapter-04-in-context-learning/02-goal-vs-inference.py
# Shows: /goal harness feature reduces model inference burden
```

## Beat 3: What "Doing the Inference Work" Actually Looks Like

In the past, inference gone wrong was just a hallucination. Now with agents in automation, it can be a **data breach**, **privilege escalation**, or simply **cost explosion** from extra tokens and tool calls.

This is where DSPy makes sense. DSPy helps you inject pragmatics into prompts for tools and workflows [dspy-notes, lines 37-43]. It's a formal way of saying: "do this, and I want this result."

People fail to be direct and pragmatic about specifying the results they want. We expect the model to infer the goal from a vague prompt when we should be explicit about the desired outcome and the constraints around it.

## Beat 4: Scope and Data Access — The Indexicality Problem

**What is indexicality?** [computational-pragmatics-notes, lines 28-38]

Indexicality is the linguistic property where meaning depends on context. Words like "I," "you," "here," "now," "this" — their meaning changes based on who's speaking, where, when, and what they're pointing to.

In context engineering: when a user says "the data," they mean different data than another user. When they say "my files," it's different from your files. The same phrase has different referents in different contexts.

```python
# Example: Indexicality definition — see examples/chapter-04-in-context-learning/05-indexicality.py
# Shows: Same phrase ("my data") means different things for different users
```

The user's example: "analyze the past month's sales data and compare to other monthly reports" — but if you pull in marketing data instead of sales data, you get the wrong result. Not a hallucination — just **wrong scope**. This is unreliable behavior.

The deeper question: when you say "the data," which data? The same query means different things to different users in different contexts.

Indexicality in semantic modeling [computational-pragmatics-notes, lines 28-50] means: the same data has different meanings in different contexts. "User's data" is different for user A vs user B. Without scoping, you're not just risking hallucinations — you're risking **wrong data access** that produces plausible but incorrect answers.

```python
# Example: Wrong scope failure — see examples/chapter-04-in-context-learning/03-wrong-scope.py
# Shows: Agent querying marketing data instead of sales data (not hallucination, wrong scope)
```

## Beat 5: Few-Shot as Teaching

This is how tool calls work. When you tell the model to do something and give it a few examples of the expected behavior, you're programming the model's actions. 

**A good few-shot example:** Shows the exact input format, the expected output format, and demonstrates the reasoning pattern you want.

**A bad few-shot example:** Ambiguous, too complex, or shows wrong behavior that the model learns anyway.

The key insight: few-shot isn't just "showing examples" — it's **teaching the model how to reason** about the problem. You're doing the inference work by demonstrating the correct inference path, so the model doesn't have to guess.

## Beat 6: Static vs Dynamic Context (Teleprompter Pattern)

The teleprompter pattern (from the DSPy paper [dspy-notes, lines 47-53]) is the distinction between:

- **Static context (teleprompter):** What you tell the model BEFORE it starts — system prompts, instructions, few-shot examples. Fixed at compile time.
- **Dynamic context (MCP):** What you FETCH during execution — tool results, retrieved documents, real-time data. Built at runtime.

MCP (Model Context Protocol) was created for Claude Code to enable dynamic context fetching. The model can now ask for more information mid-execution, rather than relying on a fixed context window.

**Which is more reliable?** Static context is more predictable — you know exactly what the model sees. Dynamic context is more powerful but introduces variability. The reliable approach: static context for invariants (instructions, constraints), dynamic context for fresh information (tool results, retrieved data).

```python
# Example: Static vs dynamic context — see examples/chapter-04-in-context-learning/04-mcp-dynamic-context.py
# Shows: Teleprompter (static) vs MCP (dynamic) context patterns
```

## A Note on Harnesses

This framework is fairly reliable **assuming the harness handles the tools well** [llms-in-production/chapter-7, p.358]. If the harness can't parse the model's output, context engineering fails at execution time. For example: some models output XML by default, while the harness expects JSON. The model did everything right — but the tool call fails.

Reliable context engineering requires matching your prompt structure to what your harness can actually parse.