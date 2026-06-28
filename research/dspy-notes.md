# DSPy: Compiling Declarative Language Model Calls into Self-Improving Pipelines — Research Notes

## Primary Source

- **DSPy: Compiling Declarative Language Model Calls into Self-Improving Pipelines**
  - Khattab, Singhvi, Maheshwari, et al.
  - Stanford NLP Group
  - October 2023
  - https://arxiv.org/abs/2310.03714

## Related Sources

- Demonstrate-Search-Predict (DSP) — Khattab et al. (2022)
  - https://arxiv.org/abs/2212.14024
- Toolformer — Meta AI (2023)

*Reading Context: Context Engineering Book Research — Chapter 5/6: Pipeline Architecture*
*Reading List Reference: Section 3 — LLM Foundations, Section 4 — Agent Architectures*

---

# Key Insight

DSPy is often described as a prompt optimization framework.

A better interpretation:

> DSPy is a context composition compiler.

The paper demonstrates that:
1. LM pipelines can be expressed as declarative computational graphs
2. Modules are parameterized (learn how to apply prompting, finetuning, augmentation)
3. A compiler can optimize any pipeline to maximize a given metric
4. Short DSPy programs can express and optimize sophisticated pipelines

---

# The DSPy Abstraction

## Text Transformation Graphs

DSPy abstracts LM pipelines as imperative computational graphs where:
- LMs are invoked through declarative modules
- Modules are parameterized (can learn)
- The compiler optimizes prompt + weight combinations

## Modules as Parameters

DSPy modules are not hard-coded prompts. They are **learnable components**:

```python
# Instead of:
prompt = "Answer this question: {question}"

# DSPy:
class ChainOfThought(dspy.Module):
    def forward(self, question):
        reasoning = dspy.ChainOfThought()(question)
        answer = dspy.Predict()(reasoning)
        return answer
```

The module **learns** how to apply reasoning techniques through demonstration collection.

---

# Historical Context: Pre-Tool Calling Agents

Chris Brousseau demonstrated DSPy multi-step pipelines **before** tool calling became mainstream in LLMs.

This is significant:

> DSPy pioneered the idea of composing LMs into multi-step pipelines when most people were still writing single prompts.

Key observations:
- 2022-2023: DSPy showed retrieval + reasoning compositions
- Early 2023: Toolformer demonstrated tool use
- Mid-2023: GPT-4 function calling + OpenAI tool schema
- Late 2023+: Agent frameworks exploded

DSPy was doing multi-step composition **before** native tool calling existed.

---

# What DSPy Got Right

## 1. Declarative Pipeline Composition

> Instead of writing prompts, write Python code that describes the pipeline.

This is now the dominant paradigm:
- LangChain expressions
- LlamaIndex query engines
- Agent frameworks

## 2. Compilation as Optimization

The compiler idea was prescient:
- Automatically optimize prompts
- Select demonstrations
- Tune parameters

Modern equivalents:
- Prompt caching
- Few-shot selection
- Automatic few-shot curation

## 3. Modules as First-Class Citizens

Modules that wrap LM calls + logic became standard:
- Retrievers
- Rewriters
- Rerankers
- Validators

---

# What DSPy Introduced

## Prompt Optimization

DSPy compilers can:
- Generate candidate prompts
- Evaluate on metric
- Select best performing

This predates:
- Anthropic's prompt caching
- OpenAI's GPT fine-tuning
- Automatic few-shot selection

## Demonstration Collection

> Modules learn by creating and collecting demonstrations.

This is now called "example curation" or "few-shot selection" - a critical part of reliable LLM systems.

---

# The DSPy Compiler / Optimizers

The compiler (called "optimizers" in DSPy) automates prompt engineering:

## What the Compiler Does

1. **Prompt Optimization** — Generates candidate prompts, evaluates on metric, selects best
2. **Demonstration Selection** — Chooses which examples to include as few-shot
3. **Module Tuning** — Adjusts how modules apply reasoning/retrieval
4. **Metric-Driven** — You provide a scoring function, compiler optimizes toward it

## Available Optimizers

| Optimizer | Strategy |
|-----------|----------|
| **BootstrapFewShot** | Collects successful demonstrations, uses as few-shot |
| **GEPA** | Reflective prompt evolution (latest, most powerful) |
| **MIPROv2** | Multi-prompt instruction optimization |
| **COPRO** | Prompt co-traning |
| **BootstrapFinetune** | Fine-tunes model on generated data |

## Compiler Workflow

```python
# 1. Define your program
extract = dspy.Predict(ExtractEvent)

# 2. Define a metric (scoring function)
def semantic_f1(prediction, example):
    return f1_score(prediction.event_name, example.event_name)

# 3. Compile against the metric
optimizer = dspy.GEPA(metric=semantic_f1, auto="medium")
optimized = optimizer.compile(extract, trainset=examples)

# Result: optimized prompts + demonstrations
# Before: 62% accuracy (zero-shot)
# After: 89% accuracy (compiled)
```

## Why This Matters for Context Engineering

> The compiler automates what context engineers spend hours doing manually:
> - Choosing the right few-shot examples
> - Tuning prompt instructions
> - Optimizing for cost/quality tradeoffs

This is early automated context engineering.

---

# The Cost of Composition

Every composition has a context cost:

| Component | Context Cost |
|-----------|--------------|
| Retrieval | Tokens for context |
| Reasoning | Chain length |
| Validation | Constraint tokens |
| Multiple LMs | Sum of all |

> More sophisticated pipelines require larger context windows and more careful cost management.

# LLM Abstraction (Swapping Models)

**Yes** — DSPy abstracts LLM calls so you can swap models without changing your code.

## How It Works

DSPy uses **LiteLLM** under the hood to normalize different providers:

```python
# Swap models by changing the string - rest of code stays the same

# OpenAI
lm = dspy.LM("openai/gpt-4o-mini")

# Anthropic
lm = dspy.LM("anthropic/claude-3-5-sonnet-20241022")

# Google
lm = dspy.LM("gemini/gemini-1.5-pro")

# Local (Ollama)
lm = dspy.LM("ollama/llama3.1")

# OpenRouter (any model)
lm = dspy.LM("openrouter/anthropic/claude-3-opus")
```

Same code, different model string. The **signatures** and **modules** stay the same.

## Supported Providers

DSPy/LiteLLM supports 100+ models across:
- OpenAI, Anthropic, Google, Meta, Mistral
- Local: Ollama, llama.cpp, vLLM
- Routers: OpenRouter, Azure, AWS Bedrock

## Why This Matters for Context Engineering

> You write your pipeline once, then swap models based on cost/performance needs.

This is exactly what the user was trying to do with templates and tool calls. DSPy handles:
- API normalization
- Response parsing
- Tool calling format conversion
- Retry logic

The compiler also **re-optimizes** for each model, so prompts that work for GPT-4 may be different for Claude.

---

# Built-in Prompting Patterns (Modules)

DSPy provides pre-built modules that implement common prompting strategies:

## 1. ChainOfThought (CoT)

> "Think step by step before answering"

```python
cot = dspy.ChainOfThought("question -> answer")
```

- Injects reasoning into context before final answer
- Built-in prompting strategy: "Let's think step by step..."
- Context engineering: Adds intermediate reasoning tokens

## 2. MultiChainComparison

> Generate multiple reasoning chains, compare them

```python
multi = dspy.MultiChainComparison("question -> answer", num_chains=3)
```

- Generates several CoT responses
- Compares/selects best one
- Reduces reasoning errors
- Context engineering: Multiple reasoning paths + selection

## 3. ReAct (Reasoning + Acting)

> Interleave reasoning with tool use

```python
react = dspy.ReAct(
    "question -> answer",
    tools=[search, calc]
)
```

- Reasoning loop: think → act → observe → repeat
- Combines CoT with tool calling
- Context engineering: Maintains state across multiple steps

## 4. ProgramOfThought (PoT)

> Generate code, execute it, use result

```python
pot = dspy.ProgramOfThought("question -> answer")
```

- Model writes Python code
- Interpreter executes
- Result becomes context for final answer

## 5. CodeAct

> Execute code in REPL, use output

```python
code = dspy.CodeAct("question -> answer")
```

---

# How These Relate to Context Engineering

These are **prompting strategies** that DSPy abstracts into reusable modules:

| Pattern | What it does | Context Engineering Equivalent |
|---------|--------------|--------------------------------|
| ChainOfThought | Adds reasoning to prompt | "Include step-by-step reasoning instructions" |
| MultiChainComparison | Multiple attempts, select best | "Generate N candidates, rank them" |
| ReAct | Reasoning + tools loop | "Maintain stateful tool-use loop" |
| ProgramOfThought | Code generation + execution | "Offload computation to interpreter" |

## Why This Matters for the Book

These patterns are now **infrastructure**, not prompts. DSPy (and similar frameworks) absorb them:

1. **You don't write the prompt** — you choose the module
2. **The compiler optimizes** — which pattern works best for your task
3. **Swappable** — swap CoT for ReAct without rewriting prompts

**Key insight**: Context engineering is moving from "write prompts" to "compose modules."

---

# Retrieve Modules & RAG

DSPy includes built-in retrieval modules for **RAG** (Retrieval-Augmented Generation):

## Built-in Retrievers

- **ColBERTv2** — dense retrieval model
- **Embeddings** — standard embedding-based retrieval
- Custom retrievers via `dspy.Retrieve`

## How DSPy Abstracts RAG

```python
# Simple RAG in DSPy
class RAG(dspy.Module):
    def __init__(self):
        self.retrieve = dspy.Retrieve(k=3)
        self.generate = dspy.ChainOfThought("context, question -> answer")

    def forward(self, question):
        context = self.retrieve(question).passages
        return self.generate(context=context, question=question)
```

- Retrieval is just another **module** in the pipeline
- You compose retrieval + generation declaratively
- The compiler optimizes the whole RAG pipeline together

## Why People Use TypeScript Instead

This is a common question. Reasons people use TS/JS frameworks (LangChain, LlamaIndex.ts) instead of DSPy:

| Reason | Explanation |
|--------|-------------|
| Existing codebase | TS/JS project, don't want Python |
| Team skills | Python isn't an option |
| Frontend integration | Browser-based, Node.js environments |
| Deployment | Serverless functions, edge, etc. |
| Ecosystem | JS/TS has huge ecosystem |

**DSPy is Python-only** (Stanford NLP is a Python shop). For TS equivalents:
- LangChain.js / LangGraph.js
- LlamaIndex.ts
- Vercel AI SDK

## DSPy vs LlamaIndex.ts — What's Replicated?

LlamaIndex.ts is the TypeScript port of **LlamaIndex** (data indexing), NOT DSPy. Different projects.

| Feature | DSPy | LlamaIndex.ts |
|---------|------|---------------|
| RAG pipelines | Yes | Yes |
| Query engines | Yes | Yes |
| Agent loops | Yes (ReAct) | Yes |
| Tool calling | Yes | Yes |
| **Prompt optimization** | **Yes (compiler)** | ❌ No |
| **Demonstration learning** | Yes | ❌ No |
| Multi-model abstraction | Yes (LiteLLM) | Partial |

**No true DSPy TypeScript equivalent exists.** The unique DSPy features (compiler, automatic few-shot selection, prompt optimization) aren't replicated in any TS framework.

---

# Claims for Evidence Ledger

---

# Alternatives in Other Languages

## Known Ports/Alternatives

| Framework | Language | Description |
|-----------|----------|-------------|
| **incode-agentware/forge** | Go | Port of the Forge tool-calling framework |
| **LogiLLM** | Python | Enterprise alternative to DSPy (zero-dependency) |
| **DsPy** | Jupyter | Alternative framework (appears to be work in progress) |

Note: DSPy is primarily Python. Most alternatives are also Python or Jupyter Notebook.

---

# DSPy vs In-Context Learning vs Tool Calling

These are related but distinct concepts:

## In-Context Learning (ICL)

> Learning from examples provided in the context window.

DSPy **automates ICL** by:
- Collecting successful demonstrations (BootstrapFewShot)
- Selecting best examples (KNNFewShot)
- Curating few-shot sets automatically

## Tool Calling

> Model invokes external functions to get information or take actions.

DSPy **supports tool calling** via:
- ReAct module for reasoning + tool use
- Tool definitions as Python functions
- MCP (Model Context Protocol) integration

## Relationship

```
ICL (foundation) → DSPy modules (composition) → Tool calling (capability)
      ↓                    ↓                      ↓
  Few-shot examples    Declarative pipelines   External actions
```

**Key insight**: DSPy predates native tool calling but its modules now support tools. The compiler optimizes both ICL (demonstrations) and tool use.

---

# Claims for Evidence Ledger

## Claim: Declarative Pipeline Composition

DSPy pioneered the idea of expressing LM pipelines as declarative Python code rather than hard-coded prompts.

**Source:** dspy (Stanford NLP, 2023)

**Supports:** Chapter 5/6 — Pipeline Architecture; Chapter 7 — Modular LLM Systems

**Strength:** strong

---

## Claim: Compilation as Optimization

The DSPy compiler optimizes prompts and parameters automatically, predating modern prompt caching and few-shot selection.

**Source:** dspy (Stanford NLP, 2023)

**Supports:** Chapter 14 — The Cost of Context; Chapter 9 — Optimization Strategies

**Strength:** strong

---

## Claim: Multi-Step Before Tool Calling

DSPy demonstrated multi-step agent-like pipelines before native tool calling existed in LLMs.

**Source:** dspy (Stanford NLP, 2023); Chris Brousseau demo (2022-2023)

**Supports:** Chapter 6 — Agent Architectures

**Strength:** strong

---

## Claim: Modules Learn from Demonstrations

DSPy modules collect and learn from demonstrations, which is now standard practice for reliable LLM systems.

**Source:** dspy (Stanford NLP, 2023)

**Supports:** Chapter 5 — In-Context Learning; Chapter 10 — Demonstration Engineering

**Strength:** strong

---

## Claim: Context Cost of Composition

Every pipeline composition has measurable context cost that grows with complexity.

**Source:** dspy (Stanford NLP, 2023)

**Supports:** Chapter 14 — The Cost of Context

**Strength:** strong

---

# Gaps Identified

This source does **not** support:
- Specific authorization patterns for multi-tenant systems
- Tool calling schemas (covered by Toolformer)
- Evaluation frameworks (covered by agent eval research)
- Knowledge graph architecture

These require additional sources.

---

# Modules + Scaffolding = Programmable Context

## The Scaffolding Concept

**Scaffolding** in LLM systems = structured guidance that shapes model behavior:

| Scaffolding Type | What it does | DSPy Equivalent |
|------------------|--------------|-----------------|
| ChainOfThought | "Think step by step" | `dspy.ChainOfThought` module |
| Tool schemas | Structure tool calls | `dspy.ReAct` with tools |
| Output schemas | Typed JSON/structs | `dspy.Predict` with signature |
| Few-shot demos | Example-based guidance | BootstrapFewShot optimizer |
| Validation | Reject bad outputs | Custom module + metric |

## Modules as Composable Scaffolding

```
User Task → [Module 1: Retrieve] → [Module 2: Rewrite] → [Module 3: Answer]
                  ↓                      ↓                       ↓
            Scaffold: RAG         Scaffold: query        Scaffold: CoT
                                   transformation
```

**Key insight**: DSPy modules are **scaffolding as code**.

- Each module encapsulates a prompting pattern (scaffolding)
- Modules can be **composed** (pipelines of scaffolding)
- The **compiler** selects/optimizes which scaffolding works

## Why Scaffolding Matters

Think of it like physics:

> Scaffolding = boundary conditions for the LLM
> 
> Without scaffolding: LLM roams freely (high variance, unpredictable)
> 
> With scaffolding: LLM operates within structured bounds (lower variance, more reliable)

The compiler finds the **optimal boundary conditions** for your task — just like finding the right constraints to solve a differential equation.

## Scaffolding in LangChain/LangGraph

LangChain calls this the **harness**:

```
Agent = Model + Harness
Harness = prompt + tools + middleware
```

DSPy is similar but adds **automatic optimization**:

| Feature | LangChain | DSPy |
|---------|-----------|------|
| Module composition | Yes | Yes |
| Tool calling | Yes | Yes |
| **Prompt optimization** | No | Yes (compiler) |
| **Auto few-shot selection** | No | Yes |

## The Book Argument

**Context engineering = building the right scaffolding.**

- Manual scaffolding = write prompts by hand (trial & error)
- Automated scaffolding = let compiler optimize (DSPy approach)
- Hybrid = templates + tools + validation rules

This connects to your optimization analogy:
- Prompt space = function space
- Metric = loss function  
- Compiler = gradient descent
- Best scaffold = global minimum

---

# Teleprompters vs MCP/ReAct Loops

## Two Different Context Problems

| Aspect | Teleprompters | MCP/ReAct Loops |
|--------|---------------|-----------------|
| **Problem** | What instructions/examples to put in context? | How to get information INTO context dynamically? |
| **Approach** | Offline optimization (compile time) | Online execution (runtime) |
| **What changes** | Prompt text, demonstrations | Tool calls, retrieved data |
| **When it runs** | Before deployment | During inference |

## Teleprompters: Static Context Optimization

Teleprompters (like DSPy optimizers) solve:

```
Given: task T, metric M, training examples D
Find: best prompt P that maximizes M on D
```

**Analogy**: Like finding the right initial conditions for a physics simulation.

```python
# Teleprompter = "compile-time" optimization
optimizer = dspy.GEPA(metric=f1_score)
best_prompt = optimizer.compile(program, trainset=examples)
```

**What it optimizes**:
- Instruction text ("You are a helpful assistant that...")
- Few-shot demonstrations (input → output pairs)
- Module configuration (which modules to use)

## MCP/ReAct: Dynamic Context Injection

MCP + ReAct loops solve:

```
Given: user query Q, available tools T
Find: sequence of tool calls that produce answer A
```

**Analogy**: Like a physics simulation that queries sensors during execution.

```python
# ReAct loop = "runtime" context construction
while not done:
    thought = llm.think(context)    # Reasoning
    action = llm.act(thought)       # Tool call
    observation = execute(action)   # Get data
    context += observation           # Inject into context
```

**What it does**:
- Calls external APIs (search, database, calculators)
- Retrieves documents/context on-demand
- Iteratively builds answer through tool use

## Key Distinction

```
Teleprompters: "What should I TELL the model before it starts?"
                ↓
           Static context (pre-computed)

MCP/ReAct:     "What should I FETCH during execution?"
                ↓
           Dynamic context (computed at runtime)
```

## Why Both Matter

| Teleprompter alone | ReAct alone |
|--------------------|-------------|
| Limited to knowledge in prompt | Can access external data |
| No tool use | Tool use enabled |
| Fast (no tool calls) | Slower (depends on tools |
| Predictable | Can be unpredictable |

**Best of both worlds**: Use teleprompters to optimize the base prompt, then wrap in ReAct for tool use.

```
┌─────────────────────────────────────────┐
│  Teleprompter Output (optimized prompt) │
│  - Instructions                          │
│  - Few-shot examples                     │
│  - Output schema                         │
└────────────────┬────────────────────────┘
                 ↓
┌─────────────────────────────────────────┐
│  ReAct Loop (at runtime)                │
│  - Think                                 │
│  - Call MCP tools                        │
│  - Observe                               │
│  - Repeat                                │
└─────────────────────────────────────────┘
```

## For the Book: The Optimization Spectrum

Think of this as a spectrum from **static** to **dynamic** context:

```
Pure Prompt    →    Teleprompter    →    ReAct    →    Agent
(static)            (optimized)          (dynamic)    (full autonomy)

↑                                       ↑
Compile-time                          Runtime
```

Each step adds:
- More context flexibility
- More execution complexity
- Less predictability
- Higher cost

This is the core trade-off for context engineers to understand.

---

# Math Problems as Benchmark

## Why Math?

The DSPy paper uses **math word problems** as a key case study. This is significant because:

1. **LLMs are bad at math** — this makes it a hard problem
2. **Clear ground truth** — unambiguous right/wrong answers
3. **Composable** — can test different context strategies:
   - Tool calling (calculator, code execution)
   - APIs (Wolfram Alpha, math APIs)
   - Few-shot prompting (examples in context)
   - RAG (retrieve similar problems + solutions)

## Optimization Framework

This connects to your physics background:

```
Prompt Optimization ≈ Finding Global Minimum
                    ↓
          Loss = evaluation metric (accuracy, F1, etc.)
                    ↓
          Gradient-like search over prompt space
                    ↓
          DSPy compiler = optimizer algorithm
```

| Physics Concept | DSPy Equivalent |
|-----------------|-----------------|
| Loss function | Metric (accuracy, F1, etc.) |
| Gradient descent | BootstrapFewShot, MIPROv2 |
| Local minima | Suboptimal prompt configurations |
| Global optimum | Best prompt + demo combination |
| Hyperparameters | Temperature, k-shot, etc. |

## Why This Matters for the Book

**Key argument**: Context engineering is an optimization problem.

Different methods = different optimization landscapes:

| Method | Landscape |
|--------|-----------|
| Zero-shot | Flat, high variance |
| Few-shot | Can find minima with good examples |
| RAG | Adds structure to search space |
| Tool calling | External computation (bypasses LLM weakness) |
| Compilation | Automated search for global optimum |

> "Just as differential equations describe how systems evolve toward equilibrium, prompt optimization describes how context evolves toward correct answers."

## Book Example Proposal

A chapter section could walk through solving the same math problem with:

1. **Baseline** — zero-shot prompt
2. **Few-shot** — add 3 worked examples
3. **RAG** — retrieve similar problems from a dataset
4. **Tool/API** — use calculator or Wolfram API
5. **DSPy compiled** — let compiler optimize the whole pipeline

Measure: accuracy, token cost, latency, reliability.

This gives readers a concrete, reproducible benchmark for understanding context engineering tradeoffs.

---

# Potential Book Quote

> DSPy revealed that the future of LLM programming was not better prompts—it was better composition.
>
> Before tool calling existed, DSPy showed how to chain LMs together into pipelines that could reason, retrieve, and validate.
>
> This was the first step toward treating LMs as components in a larger system, not just as endpoints.