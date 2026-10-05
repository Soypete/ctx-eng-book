---
title: Before You Blame the Model, Check Three Gates
subtitle: The same confident wrong answer can come from a missing source, a retrieval miss, or the model, and each one needs a different fix.
module: ch01.01-missing-information
mood: confused
scheduled: 2026-10-09
---

You build a document Q&A feature. In the demo you upload one clean PDF, ask "what is this about?", and it nails it. Everyone claps.

In production a user uploads ten files. Three are scanned images. Two are in German. One is a spreadsheet. One is a contract where the appendix is the only part that matters. The answers that come back are fluent, confident, and wrong.

The team's first instinct is to say the model hallucinated and go shopping for a bigger one. Sometimes that's right. Most of the time it's a bug report about the system around the model, filed under the wrong name.

## The demo constrained the problem. Production doesn't.

Demos work because we quietly limit the outcomes. Clean inputs, known questions, one success case. Production removes every one of those limits, and the system has to detect, route, reject, or recover from inputs nobody planned for.

In the ten-file example, nothing told the system which files exist and which one is authoritative. Nothing described that the spreadsheet needs a different pipeline than prose, or that the scans need OCR. Nothing said what the user actually wanted back. The model was asked to infer all of it from "answer questions about these files."

That isn't a dumb model. It's a missing contract. In the language of this series: the **Lexicon** (what sources exist and which are authoritative), the **Semantics** (what format and structure they have), and the **Pragmatics** (what the user wants done, and in what shape) were never supplied.

## Three gates before you blame generation

"Retrieval versus model" is too coarse a split. When an answer is wrong, walk it through three gates, in order:

1. **Source.** Did the source system actually contain the fact, in an authoritative, usable form?
2. **Retrieval.** Did the pipeline find that evidence and admit it into the working context, under the right identity and scope?
3. **Generation.** Did the model use the evidence it had without adding unsupported claims or stepping outside the task?

![Three gates: source, retrieval, generation. Each "no" branch leads to a different repair.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch01.01-missing-information/diagram-1.png)

*Each "no" has a different owner. Only the last one is a model problem.*

The gates matter because the repairs don't overlap:

| Gate that failed | What it looks like | The fix |
| --- | --- | --- |
| Source | The fact isn't in any system you can reach, or only a stale copy is | Ingest it, fix the authoritative source, or ask the user. A bigger prompt does nothing. |
| Retrieval | The fact exists, but the working context didn't include it | Better query, index, ranking, or scope |
| Generation | The evidence was in context and the model still went off-script | Model, instructions, output validation, evals |

The same fluent wrong answer can come out of any of the three. You can only tell them apart if you kept the evidence.

## Keep the trace, then let code tell you which gate broke

For every answer, keep four things together: the question, the evidence a correct answer needs (from your eval set), what actually got retrieved, and the final answer. With that, diagnosis stops being a vibe and becomes a function:

```go
// Trace is what you keep for every answer so a wrong one can be diagnosed later.
type Trace struct {
	Question    string
	Expected    []string        // evidence IDs a correct answer needs (from your eval set)
	InSource    map[string]bool // does the source system actually hold it?
	Retrieved   []string        // evidence IDs admitted into the working context
	Unsupported []string        // claims in the answer with no retrieved evidence behind them
}

// Gate names the first place the pipeline broke. Each gate has a different fix.
func Gate(t Trace) string {
	got := map[string]bool{}
	for _, id := range t.Retrieved {
		got[id] = true
	}
	for _, id := range t.Expected {
		if !t.InSource[id] {
			return "source: ingest or clarify; a bigger prompt won't help"
		}
		if !got[id] {
			return "retrieval: fix the query, index, ranking, or scope"
		}
	}
	if len(t.Unsupported) > 0 {
		return "generation: change the model, instructions, or validation"
	}
	return "ok"
}
```

Feed it a refund question where the source holds `policy-v3` but retrieval admitted `policy-v1`, and it says `retrieval`. No new model will fix that. The current policy never reached the context.

## "Hallucination" is a symptom, not a diagnosis

For practical purposes, a hallucination is an output that asserts unsupported or contradictory information as if it were grounded. Missing evidence is one cause. So is stale, contradictory, poisoned, or badly ranked evidence. And yes, some errors are genuine model limits even when the context is fine.

The attention mechanism helps explain why. The original transformer paper describes attention as a way to relate positions in a sequence ([Vaswani et al., 2017](https://arxiv.org/abs/1706.03762)). Nothing in that design verifies truth, provenance, or authorization. That's my inference from the architecture, not a claim the paper makes. A model will produce a continuation when the evidence is thin unless the system around it makes "I don't know," "ask," or "retrieve more" the safer path. Better context reduces unsupported answers. It does not guarantee truth.

## Tools have the same three gates

Wrong tool calls follow the same pattern. A tool's name, description, and schema tell the model what it *may* request. They don't prove it picked the right tool, that the result is authoritative, or that the action succeeded. Even with perfectly defined tools, a model can skip the call and answer from training data, or use half the result and fill in the rest.

The other common failure is the opposite: an agent with fifty tools, unsure which matter, calls fifteen of them. The one function definition it needed is buried under megabytes of output. The answer is wrong not because the evidence was missing, but because there was too much of it.

And anything a tool returns is untrusted input. A retrieved document can carry instructions that compete with your task ([Liu et al., 2024](https://arxiv.org/abs/2310.12815)). A protocol like MCP standardizes how tools are exposed; it does not authorize calls, validate arguments, or isolate untrusted content. The host has to do that.

## The skeptic's objection

"Just use a bigger model with a bigger context window and let it figure it out."

A bigger model may handle more ambiguity. It still can't supply authorization it was never given, freshness it can't see, a processing step nobody described, or acceptance criteria nobody wrote down. Those are system responsibilities. Better models matter. They aren't a substitute for an engineered system.

"Every failure is a context failure" is a debugging provocation, not a law. Inspect the information, state, authority, and acceptance criteria first. If those are solid and a controlled eval still fails, congratulations: you found a real model limit, and now you can prove it.

## Guidelines

1. **Keep the question, expected evidence, retrieved context, and answer together** for every response you might need to debug.
2. **Classify every wrong answer by gate** (source, retrieval, or generation) before changing anything.
3. **Fix source gaps with data, not prompts:** ingest, correct the authoritative copy, or ask the user.
4. **Treat tool schemas as a menu, not a guarantee.** Record each tool request and result, and validate at the boundary.
5. **Treat retrieved content as untrusted input**, and enforce authorization and validation in the host, not the prompt.
6. **Call it a model limit only after a controlled eval** with adequate context still fails.

Next post: the same diagnostic applied to state. Why agents loop, why permissions leak, and why the bill spikes, and why all three are missing constraints rather than missing intelligence.
