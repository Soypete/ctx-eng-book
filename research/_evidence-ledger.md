# Evidence Ledger

Claims extracted from research sources, mapped to book pillars.

---

## Pragmatics — Declarative Pipeline Composition

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "DSPy is a context composition compiler... LM pipelines can be expressed as declarative computational graphs"
- **Locator:** DSPy paper, abstract
- **Supports:** Chapter 5/6 — Pipeline Architecture; Chapter 7 — Modular LLM Systems
- **Strength:** strong

---

## Pragmatics — Compilation as Optimization

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "The compiler optimizes prompts and parameters automatically, predating modern prompt caching and few-shot selection"
- **Locator:** DSPy paper, section on optimizers
- **Supports:** Chapter 14 — The Cost of Context; Chapter 9 — Optimization Strategies
- **Strength:** strong

---

## Pragmatics — Multi-Step Before Tool Calling

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "DSPy pioneered the idea of composing LMs into multi-step pipelines when most people were still writing single prompts"
- **Locator:** Historical context section
- **Supports:** Chapter 6 — Agent Architectures
- **Strength:** strong

---

## Pragmatics — Modules as Composable Scaffolding

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "DSPy modules are parameterized, meaning they can learn (by creating and collecting demonstrations) how to apply compositions of prompting, finetuning, augmentation, and reasoning techniques"
- **Locator:** DSPy paper, abstract
- **Supports:** Chapter 5 — In-Context Learning; Chapter 7 — Modular LLM Systems
- **Strength:** strong

---

## Pragmatics — Teleprompters vs Runtime Context

- **Source:** dspy-notes (synthesized from DSPy + LangChain)
- **Quote:** "Teleprompters: 'What should I TELL the model before it starts?' (static context, compile-time) vs MCP/ReAct: 'What should I FETCH during execution?' (dynamic context, runtime)"
- **Locator:** Teleprompters vs MCP/ReAct Loops section
- **Supports:** Chapter 9 — Optimization Strategies; Chapter 6 — Agent Architectures
- **Strength:** strong

---

## Pragmatics — Math Problems as Optimization Benchmark

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "Prompt Optimization ≈ Finding Global Minimum... Gradient-like search over prompt space"
- **Locator:** Math Problems as Benchmark section
- **Supports:** Chapter 9 — Optimization Strategies; Chapter 14 — The Cost of Context
- **Strength:** strong
- **Counterpoint:** LLMs are not differentiable — this is a useful analogy but not mathematically exact

---

## Pragmatics — Context Cost of Composition

- **Source:** dspy-notes (Khattab et al., 2023)
- **Quote:** "Every pipeline composition has measurable context cost that grows with complexity"
- **Locator:** The Cost of Composition section
- **Supports:** Chapter 14 — The Cost of Context
- **Strength:** strong

---

## Evaluation — Agent Evals Need Harnesses

- **Source:** llms-in-production/chapter-7
- **Quote:** "A harness is a system for [evaluation]"
- **Locator:** chapter-7, line 358
- **Supports:** Chapter 12 — Evaluation; Harness Design
- **Strength:** strong