# How Attention Works

## Argument
The transformer attention mechanism is fundamentally an information routing system—it decides what to focus on, not what is true.

## Why it matters
Context engineering's first job is deciding what information to put in context. Understanding attention explains why some information gets used and other information gets ignored.

## Beats
1. What attention actually does — weighted information retrieval within the context window (lines 23-27 of research)
2. The critical assumption — all relevant information must already exist in the context (lines 34-38)
3. What attention does NOT solve: truth, provenance, authorization, memory (lines 16-27)
4. Position encoding — why order matters and how it affects retrieval
5. Context windows as fixed memory — the hard limit every context engineer must work within
6. The cost of context — why more tokens = more computation

## Conclusion to reach
Attention is a routing mechanism, not a reasoning mechanism. Context engineering must ensure the right information is present BEFORE attention runs.

## Examples needed
- Visualization of attention weights on a sample prompt
- Diagram showing context window as fixed buffer

## Readiness
needs-research (position encoding details, context window limits per model)