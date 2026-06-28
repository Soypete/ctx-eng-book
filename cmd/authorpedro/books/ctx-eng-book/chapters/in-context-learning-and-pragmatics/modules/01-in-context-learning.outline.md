# In-Context Learning: How Models Learn from Context

## Argument
Models don't just read context—they learn from it. In-context learning is the mechanism that makes context engineering possible, and understanding inference is the key to doing it well.

## Why it matters
The entire discipline of context engineering rests on the fact that models can learn from provided context. Without in-context learning, there would be no point to retrieval, memory, or prompt construction.

## Beats
1. The four inferential problems (Jurafsky) — reference resolution, speech act, discourse coherence, abduction (lines 15-19)
2. What pragmatics teaches us: inference is the core problem (line 22)
3. Context engineering as inference reduction — solving inference BEFORE the model sees input (lines 56-66)
4. Indexicality — why "user's data" means different things to different users (lines 28-50)
5. Few-shot as teaching — giving examples is programming the model's behavior
6. The teleprompter pattern — static context vs dynamic context (from dspy-notes)

## Conclusion to reach
In-context learning is the mechanism. Pragmatics is the theory. Together they tell us: context engineering is the practice of doing the model's inference work for it, so the model can focus on extraction rather than guessing.

## Examples needed
- Before/after: ambiguous prompt vs disambiguated with context
- Few-shot example showing how demonstrations shape model behavior

## Readiness
ready