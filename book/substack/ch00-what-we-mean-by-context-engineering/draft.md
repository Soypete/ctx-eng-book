---
title: Context Engineering Is Not Prompt Engineering
subtitle: Reliability comes from controlling what information and authority reach the model, not from finding better words.
module: ch00-what-we-mean-by-context-engineering
mood: professor
scheduled: 2026-10-07
---

Picture an internal support bot. A customer-success rep asks it how to explain a pricing change, and the answer comes back with a line from the compensation spreadsheet. The prompt was fine. The vector search was fine; that spreadsheet really was the most similar document in the index. The model did exactly what it was asked.

So what broke? Nobody decided whether that spreadsheet was allowed to compete for a spot in the answer.

That is the problem this series is about, and it is not a prompt problem.

## What we keep calling the wrong thing

When an AI feature misbehaves, the first move is usually to edit the prompt. Add a sentence. Add a rule in capital letters. Add three more examples. Sometimes that helps. But a lot of the failures we see in production (hallucinated answers, wrong tool calls, data leaking across users, bills that spike) do not start in the wording. They start in the information system around the model.

Here is the definition I'm using for the whole series:

**Context engineering is delivering the right information and control signals to a language process, at the right stage of a workflow, for an authorized purpose, within measurable limits.**

The goal is not to hand the model everything that might be useful someday. It is to give it what the current decision needs, in a form it can use, bounded by who is asking and why. Reliable systems come from controlling information flow, not from maximizing context size.

I organize that work around three questions:

- **Lexicon:** what data, entities, and sources exist? Who owns them, what's authoritative, what's stale, what's sensitive?
- **Semantics:** what do those things mean here, in this domain, at this time? Schemas, identifiers, relationships, definitions.
- **Pragmatics:** what is this actor allowed to say or do with that meaning, and what gets checked before an output becomes an action?

The support bot failed at the Lexicon layer. It never knew the spreadsheet was out of scope for that user, so no amount of Semantics or Pragmatics downstream could save it.

## Who owns the selection?

"Retrieval" hides a decision that matters more than which embedding model you pick: who decides what the model gets to see?

There are three common postures.

**Prompt-time compilation.** The application retrieves candidates, ranks them, trims to a budget, and builds the model input before generation. The harness owns selection, so you can inspect the exact working set, replay it, and compare it across ranker versions.

**Scoped tool retrieval.** The model gets a search or fetch tool and decides what to ask for mid-loop. This is adaptive, and you need it when the next question depends on what the model just learned. But each call is a new opportunity to fail, and the model may never ask for the one document that would have settled things.

**Hybrid.** The harness seeds the prompt with high-confidence context and exposes a bounded tool for the gaps.

None of these is universally right. What matters is that in every one of them, **the system, not the model, owns the boundary.** The model can decide *when* it needs more context. The system decides what that request is allowed to expose.

| Posture | Who decides when to retrieve | What the system must make deterministic |
| --- | --- | --- |
| Prompt-time compilation | Harness | authorization, candidate set, ranking inputs, cutoff, budget, provenance |
| Scoped tool retrieval | Model proposes, host authorizes | tool scope, identity and purpose, result limits, validation, lineage |
| Hybrid | Harness seeds, model fills bounded gaps | seed manifest, tool scope, per-call limits, final evaluation |

The rule I use: compile context ahead of time when you already know what evidence and policy apply. Expose a scoped tool when the useful context depends on execution. Use both when a good seed shrinks the uncertainty and the remaining questions are genuinely dynamic. And don't make `search_all_company_data(query)` your default just because the model is capable of writing queries.

![Authorization happens before ranking, and the model's tool calls loop back through the same gate.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch00-what-we-mean-by-context-engineering/diagram-1.png)

*Notice where the model's tool call goes: back through Authorize, not around it.*

## What this looks like in code

Here's the smallest version of that pipeline I could write. The point isn't the sorting. It's the order of operations, and the fact that every rule lives in code you control:

```go
type Doc struct {
	ID, Owner, Text string
	Score           float64 // relevance from any ranker: BM25, vector, graph
	Tokens          int
}

type Manifest struct{ Included, Excluded []string }

// Compile builds the model's working set. The system owns every rule here;
// the model only ever sees what survives.
func Compile(user string, allowed func(user, owner string) bool, cands []Doc, budget int) ([]Doc, Manifest) {
	var m Manifest
	var eligible []Doc
	for _, d := range cands { // 1. authorize BEFORE ranking: unauthorized docs never compete
		if allowed(user, d.Owner) {
			eligible = append(eligible, d)
		} else {
			m.Excluded = append(m.Excluded, d.ID+":unauthorized")
		}
	}
	sort.Slice(eligible, func(i, j int) bool { return eligible[i].Score > eligible[j].Score }) // 2. rank
	var out []Doc
	for _, d := range eligible { // 3. deterministic token budget
		if d.Tokens > budget {
			m.Excluded = append(m.Excluded, d.ID+":budget")
			continue
		}
		budget -= d.Tokens
		out = append(out, d)
		m.Included = append(m.Included, d.ID)
	}
	return out, m // 4. the manifest makes the selection replayable
}
```

Run it with a public pricing doc, a finance-owned salary sheet that scores *higher*, and an oversized personal note, and the manifest reads `Included:[a] Excluded:[b:unauthorized c:budget]`. The salary sheet had the best relevance score and still never reached the ranker. That's the bug from the top of this post, fixed in four lines.

Ranking can be as probabilistic as you like: lexical, vector, graph, learned. Scope, cutoff, budget, and admission should not be.

## Context is broader than retrieval, and broader than agents

Agents get most of the attention right now, but they aren't the definition of the field. The same discipline applies to search, extraction, classification, summarization, and plain old workflows. In every one, the model is one participant in a bigger pipeline. Something has to assemble the relevant context, make the meaning explicit, and constrain what the result is allowed to cause.

Two more dimensions cut across everything above:

- **Time.** Sources change, definitions get versioned, permissions expire. Ask not just what a fact means, but when it was true and whether the authorization is still valid.
- **Structure.** The same information is more or less usable depending on ordering, labels, and serialization. Syntax makes meaning available. It does not supply authority.

When something misbehaves, walk it in that order. Which Lexicon, Semantics, or Pragmatics assumption failed? Did the information arrive at the wrong time or in the wrong shape? Then look at enforcement: retrieval filters, provenance, schemas, authorization checks, workflow state, evals. The prompt is the last place to look, not the first.

## Guidelines

1. **Authorize before you rank.** An unauthorized record should never compete for a slot, even if it's later dropped from the prompt.
2. **Decide who owns selection** for each feature: harness, model, or a bounded mix. Write it down.
3. **Route every tool call through the same authorization gate** as prompt-time context.
4. **Make budgets and cutoffs deterministic** and record them in code, not in prompt prose.
5. **Emit a manifest** of what was included, what was excluded, and why, so any answer can be replayed.
6. **Debug the information system before the prompt:** sources, meaning, permissions, timing, structure, then wording.

This is the first post in a series on context engineering, one post per section of the book I'm writing. Next up: what happens when the information a model needs simply isn't there, and why "the model hallucinated" is usually a retrieval bug report in disguise.
