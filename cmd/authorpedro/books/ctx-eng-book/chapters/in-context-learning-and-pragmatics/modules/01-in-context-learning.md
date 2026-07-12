# In-Context Learning: How Models Learn from Context

## Beat 1: Speech Act Interpretation

Jurafsky's computational pragmatics identifies four inferential problems:
1. Reference resolution — what does "it" refer to?
2. Speech act interpretation — what did the user mean?
3. Discourse coherence — how does this relate to what was said?
4. Abduction — what explains this utterance?

The most common failure mode isn't missing information — it's **speech act interpretation**. The model doesn't understand what the user wants DONE with the information provided.

Of the four inferential problems, speech act interpretation has the highest impact in production AI systems — getting this wrong leads to wrong tool calls, incorrect actions, and unreliable outcomes.

This is why frameworks like ReAct [arXiv:2210.03629], chain-of-thought prompting [arXiv:2201.11903], and planning modes (like Claude Code's plan mode [code.claude.com/docs]) exist. They're not adding intelligence — they're adding structure that helps the model extract the user's actual intent when the user hasn't efficiently communicated it.

A user says "fix the API" but means "find the bug in the auth handler." The model sees "fix the API" and guesses. ReAct forces the model to first reason about intent before taking action.

The ReAct pattern forces the model to first reason about intent before taking action.

## Beat 2: Why More Inference Means Less Reliability

In machine learning, inference is the act of reaching a conclusion. The problem: when humans communicate, there's a **shared background context** — from prior conversations, collaborative thinking, life experience. We build understanding together.

Models don't have shared context. They have a transformer — a mathematical mechanism for pattern matching, not shared understanding.

When harnesses introduce `/goal` or explicit goal definition, they're doing something profound: they're removing the burden of figuring out *how* to do something and instead focusing the model on *what* to do. The model no longer needs to infer the approach — it's given a target. This produces outcomes more aligned with user expectations.

This is the core insight: **more inference by the model means less reliability**. Context engineering flips this — we do the inference work so the model can just extract.

The /goal directive removes the burden of figuring out *how* to do something and instead focuses the model on *what* to do.

## Beat 3: What "Doing the Inference Work" Actually Looks Like

In the past, inference gone wrong was just a hallucination. Now with agents in automation, it can be a **data breach**, **privilege escalation**, or simply **cost explosion** from extra tokens and tool calls.

This is where DSPy makes sense. DSPy helps you inject pragmatics into prompts for tools and workflows. It's a formal way of saying: "do this, and I want this result."

People fail to be direct and pragmatic about specifying the results they want. We expect the model to infer the goal from a vague prompt when we should be explicit about the desired outcome and the constraints around it.

## Beat 4: Scope and Data Access — The Indexicality Problem

**What is indexicality?**

Indexicality is the linguistic property where meaning depends on context. Words like "I," "you," "here," "now," "this" — their meaning changes based on who's speaking, where, when, and what they're pointing to.

In context engineering: when a user says "the data," they mean different data than another user. When they say "my files," it's different from your files. The same phrase has different referents in different contexts.

When a user says "the data," they mean different data than another user. Without explicit scoping, the same phrase has different referents.

The user's example: "analyze the past month's sales data and compare to other monthly reports" — but if you pull in marketing data instead of sales data, you get the wrong result. Not a hallucination — just **wrong scope**. This is unreliable behavior.

The deeper question: when you say "the data," which data? The same query means different things to different users in different contexts.

Indexicality in semantic modeling means: the same data has different meanings in different contexts. "User's data" is different for user A vs user B. Without scoping, you're not just risking hallucinations — you're risking **wrong data access** that produces plausible but incorrect answers.

The agent might query marketing data instead of sales data—not a hallucination, but wrong scope. This is unreliable behavior.

## Beat 5: Few-Shot as Teaching

This is how tool calls work. When you tell the model to do something and give it a few examples of the expected behavior, you're programming the model's actions. 

**A good few-shot example:** Shows the exact input format, the expected output format, and demonstrates the reasoning pattern you want.

**A bad few-shot example:** Ambiguous, too complex, or shows wrong behavior that the model learns anyway.

The key insight: few-shot isn't just "showing examples" — it's **teaching the model how to reason** about the problem. You're doing the inference work by demonstrating the correct inference path, so the model doesn't have to guess.

## Beat 6: Static vs Dynamic Context (Teleprompter Pattern)

The teleprompter pattern (from the DSPy paper) is the distinction between:

- **Static context (teleprompter):** What you tell the model BEFORE it starts — system prompts, instructions, few-shot examples. Fixed at compile time.
- **Dynamic context (MCP):** What you FETCH during execution — tool results, retrieved documents, real-time data. Built at runtime.

MCP (Model Context Protocol) is a protocol that enables the model to fetch additional context during execution, rather than relying on a fixed context window. It was created for Claude Code to enable dynamic context fetching. The model can now ask for more information mid-execution, rather than relying on a fixed context window.

**Which is more reliable?** Static context is more predictable — you know exactly what the model sees. Dynamic context is more powerful but introduces variability. The reliable approach: static context for invariants (instructions, constraints), dynamic context for fresh information (tool results, retrieved data).

Static context is what you tell the model before it starts (system prompts, instructions, few-shot examples). Dynamic context is what you fetch during execution (tool results, retrieved documents, real-time data).

## Beat 7: Tool Formatting — How Models Choose Actions

Tool formatting is the bridge between intent and execution. When a model decides to call a function, it must output a correctly formatted call — name, arguments, structure. This is where most agent systems fail.

The problem: models are trained on text, not tool schemas. They need explicit instruction on:
- What tools exist and when to use each one
- The exact argument format each tool expects
- What a valid response looks like vs. what to skip

Tool formatting requires explicit instruction on what tools exist, when to use each one, the exact argument format, and what a valid response looks like.

**Tool formatting best practices:**
1. **Schema-first**: Define tools with JSON Schema or equivalent. The schema becomes part of the prompt.
2. **One-shot examples**: Show exactly one valid call for each tool in the system prompt.
3. **Failure handling**: Tell the model what to do when a tool call fails — retry, fall back, or escalate.

When tool formatting is ambiguous, the model guesses. The result: wrong function names, malformed arguments, or calls to non-existent tools. This is not a model intelligence problem — it's a context engineering problem.

## Beat 8: Structured Outputs — Reliability Through Schema

Structured outputs constrain model behavior to known shapes. Instead of free-form text, the model outputs JSON, XML, or typed objects that downstream systems can parse reliably.

**Why this matters for reliability:**
- **Parseable**: No regex extraction, no fragile string matching
- **Validated**: Schema violations fail fast, before they cause downstream errors
- **Auditable**: Every output has a known structure to inspect

Structured output with validation (using Pydantic or similar) guarantees parseable output that downstream systems can handle reliably.

The pattern: constrain the output space. The more you specify what the output should look like, the more reliable the model behaves. This is why function calling was one of the earliest reliable agent patterns — it constrains output to a known schema.

**Structured output patterns:**
1. **JSON Schema forcing**: Tell the model "output valid JSON matching this schema"
2. **Grammar-based generation**: Use constrained decoding to only allow valid tokens
3. **Two-stage: extract then validate**: Let the model output freely, then validate and reject non-conforming results

## Beat 9: Prompt Tuning — Optimizing the Signal

Few-shot examples and instructions are not static — they can be tuned. Prompt tuning (including methods like A/B testing, evolutionary prompt optimization, and learned prompts) searches the space of possible instructions to find what works.

The key insight: **prompt performance is measurable**. Given a task and evaluation criteria, you can systematically improve the prompt. This is no different from tuning any other system parameter.

**What to tune:**
- Instruction wording — "Analyze the following" vs "Extract key insights from"
- Example selection — which few-shot examples produce the best results
- Order effects — do high-quality examples work better at the start or end
- Constraint framing — "never do X" vs "avoid doing X"

**What not to tune on:**
- Single examples — prompt performance has variance; test across diverse inputs
- One-off evaluations — prompts that work on one query may fail on others

## Beat 10: The Reliability Question — How Do We Communicate Intent Clearly?

This chapter started with a question: **how do we communicate intent clearly to the model?**

The answer, through the lens of in-context learning and pragmatics:

1. **Do the inference work**: Don't make the model guess what you want. State the goal explicitly. Use goal directives (`/goal`) rather than hoping the model infers from vague instructions.

2. **Scope the data**: Address indexicality directly. "The data" means nothing — "the sales data from the past 30 days for account A" means something precise.

3. **Show, don't just tell**: Few-shot examples teach the model the reasoning pattern you want, not just the format you expect.

4. **Separate static from dynamic**: Use static context (teleprompter) for invariants — instructions, constraints, examples. Use dynamic context (MCP) for runtime information.

5. **Constrain outputs**: Tool formatting and structured outputs reduce the model's output space to what your system can reliably parse.

6. **Measure and tune**: Prompt performance is observable and optimizable. Treat prompts as parameters to tune, not magic incantations.

Context engineering is the practice of removing ambiguity from the model's input so it can focus on extraction rather than inference. The more precisely you communicate intent, the more reliably the model acts on it.

## A Note on Harnesses

This framework is fairly reliable assuming the harness handles the tools well. If the harness can't parse the model's output, context engineering fails at execution time. For example: some models output XML by default, while the harness expects JSON. The model did everything right — but the tool call fails.

Reliable context engineering requires matching your prompt structure to what your harness can actually parse.