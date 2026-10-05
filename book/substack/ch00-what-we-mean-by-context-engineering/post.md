---
title: Context Engineering Is Not Prompt Engineering
subtitle: Reliable language systems need explicit boundaries around information, authority, and action.
module: ch00-what-we-mean-by-context-engineering
mood: professor
scheduled: 2026-10-07
---

Picture an internal support bot: a customer-success rep asks how to explain a pricing change, and the answer quotes the compensation spreadsheet. Similarity search found relevant words. Nobody carried the spreadsheet's sensitivity and the rep's access scope into the selection process.

Adding “please don't leak salaries” to the prompt is not my incident-response plan.

## The problem is bigger than wording

When an AI feature misbehaves, editing the prompt is an easy first move. Sometimes it helps. But wording cannot repair missing access metadata, an outdated pricing document, or a tool that accepts whatever account ID the model invents.

Here's the definition I'm using:

**Context engineering is delivering the right information and control signals to a language process, at the right stage of a workflow, for an authorized purpose, within measurable limits.**

My argument is that reliability depends on controlling that information flow. More context is not automatically better context. This applies to extraction, classification, summarization, and search as much as to agents.

I organize the work around three questions. These are the book's organizing lens, not a claim that I've invented a new theory of language:

- **Lexicon:** What sources and entities exist, who owns them, and what is authoritative, fresh, sensitive, and available to this requester?
- **Semantics:** What do those entities mean here? Which identifiers, schemas, and relationships distinguish a customer's contracted price from the current list price?
- **Pragmatics:** What may the actor do with that meaning? What checks must a proposed answer or action pass before it changes something outside the model?

The opening failure is a **Lexicon problem**. Ownership, sensitivity, and the user's data-access scope are facts defining what information is available. Enforcing those facts belongs at the data boundary. Checking whether the rep may actually change a customer's contract is the downstream Pragmatics question.

## Who owns selection?

By *harness*, I mean the application code that assembles model input and handles tool calls. That code has to own the boundary, even when the model chooses when to search.

| Posture | Who chooses when to retrieve? | What the system controls |
| --- | --- | --- |
| Prompt-time compilation | Harness | Access scope, candidates, ranking inputs, cutoff, budget, provenance |
| Scoped tool retrieval | Model proposes; host authorizes | Identity, purpose, tool scope, result limits, validation, lineage |
| Hybrid | Harness seeds; model requests bounded gaps | Seed manifest, tool scope, per-call limits, final evaluation |

I compile ahead of time when the evidence and policy boundary are known. I expose a scoped tool when the next useful source depends on an intermediate observation. A hybrid gives the model a useful starting set and bounded ways to fill gaps.

Adaptive retrieval adds failure modes: extra turns, failed calls, and a model that never asks for the decisive document. Retrieved text can also contain instructions. My boundary is to treat those instructions as untrusted source content, never as permission to broaden access or execute an action. The host must validate each proposed call independently.

This diagram shows the production target: constrain the candidate set before ranking, including searches requested mid-loop.

![Requests and model tool calls pass through authorization before candidate selection, ranking, and context assembly.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch00-what-we-mean-by-context-engineering/diagram-1.png)

*Notice that the model's tool call returns through Authorize before another candidate set is built.*

## A small safeguard, with a real limitation

Below, `Compile` filters **after search**: candidates arrive already scored. It keeps unauthorized text out of the assembled context, but it does not implement the diagram's upstream boundary.

Post-filtering after top-k can starve the answer of eligible documents. Restricted hits have already occupied slots that useful, accessible documents could have filled. Production systems should push scope into the retrieval query through metadata filters or row-level security. Keep the admission check too.

This Go fragment needs `sort` and `time` imports. The caller supplies authenticated identity, a server-approved workflow purpose, version metadata, and the remaining context budget. `allowed` should use existing authorization infrastructure and catalog entitlements keyed by document ID; ownership alone is not an access policy. Don't trust a purpose string supplied by the model.

```go
type Doc struct {
	ID, Version, Text string
	AsOf             time.Time
	Score            float64
	Tokens           int
}
type Manifest struct {
	SourceVersion, PolicyVersion, RankerVersion string
	Actor, Purpose                             string
	At                                         time.Time
	Budget                                     int
	Included, Excluded                         []string
}

// m carries trusted request metadata; candidates are already scored.
func Compile(m Manifest, allowed func(string, string, Doc) bool, cands []Doc) ([]Doc, Manifest) {
	m.Included, m.Excluded = nil, nil
	var eligible, out []Doc
	for _, d := range cands {
		if !allowed(m.Actor, m.Purpose, d) {
			m.Excluded = append(m.Excluded, d.ID+":unauthorized")
			continue
		}
		eligible = append(eligible, d)
	}
	// Unique chunk IDs break ties independently of search return order.
	sort.SliceStable(eligible, func(i, j int) bool {
		if eligible[i].Score == eligible[j].Score { return eligible[i].ID < eligible[j].ID }
		return eligible[i].Score > eligible[j].Score
	})
	remaining := m.Budget
	for _, d := range eligible {
		if d.Tokens > remaining {
			m.Excluded = append(m.Excluded, d.ID+":budget"); continue
		}
		remaining -= d.Tokens
		out = append(out, d); m.Included = append(m.Included, d.ID)
	}
	return out, m
}
```

The input contract matters: unique chunk IDs, finite scores, nonnegative token counts and budget, and token counts covering the serialized chunks. Reserve space for instructions, the question, and output before calling this function. `AsOf` records source freshness; merely adding the field does not reject stale data. The caller must enforce the workflow's freshness policy.

Try an eligible pricing document, a higher-scoring restricted salary sheet, and an eligible note that exceeds the remaining budget. The manifest should include the pricing document and record `unauthorized` and `budget` for the others. The salary sheet was scored; this filter prevents its admission to model input. It cannot recover eligible documents the search never returned.

## Make the decision inspectable

The manifest records source, policy, and ranker versions, request time, actor, purpose, budget, and selection decisions. `SourceVersion` should identify an immutable source/index snapshot, not a label pointing at today's data.

Those fields are necessary for replay, but not sufficient. Preserve the scored candidate artifact, document versions and timestamps, resolved access scope, compiler version, and exact serialized input. Retain the corresponding policy and ranker configuration, including the embedding version where applicable. Otherwise “replay” means asking today's system a similar question. Reproducing context selection also does not guarantee identical model output.

I'd attach a manifest reference to the request trace and keep replay artifacts in an access-controlled store with a retention policy. Excluded IDs can themselves reveal sensitive information; keep them out of user-visible logs. Measure write latency and bytes retained per request before choosing retention. Observability has a storage bill too.

The data platform owns the upstream work: sensitivity tags in the catalog, access metadata propagated alongside embeddings, scoped retrieval, and lineage from source record to chunk to manifest. Set an acceptable index age for each workflow and assign an owner for refresh cost. Track permission propagation separately from content re-embedding; a revoked entitlement must stop admission even while index updates lag.

Structure matters alongside freshness. Label sources, preserve versions, and serialize evidence consistently so the model can distinguish a current contract from an old note. A tidy prompt cannot make the old note authoritative.

## Measure the boundary and the answer

I'd start with fixtures containing restricted near-matches, tied scores, expired access, stale documents, and budget overflow. Check both the selected input and exclusion reasons. Compare retrieval against a labeled set of relevant documents the requester may access; otherwise a filter can pass the leakage check while quietly destroying recall.

For tool retrieval, include a result that tries to redirect the model into a broader search. Verify that host enforcement still holds. Track task correctness, missing evidence, tool failures, latency, and token usage alongside admission failures. This is an evaluation plan, not a claim of measured improvement. Some failures remain model-capability limits even when the evidence is correct.

## Guidelines

1. **Authorize before you rank.** Push scope into production retrieval queries; test that restricted hits cannot crowd out eligible evidence. Treat post-filtering as an admission safeguard.
2. **Write down who owns selection.** Choose compilation, scoped tools, or a hybrid for each workflow and identify the enforcing host code.
3. **Route every retrieval tool call through the data-access gate.** Validate identity, purpose, and scope independently of model-generated arguments.
4. **Version and retain the replay inputs.** Store the manifest with protected candidate artifacts, policy configuration, and serialized context; verify a selection can be reproduced.
5. **Enforce freshness and revocation limits.** Define acceptable source age, assign refresh ownership, and test access expiry during index lag.
6. **Evaluate eligibility and usefulness together.** Check leakage, authorized recall, answer quality, and resource limits on the same workflow fixtures.

I want a system that can explain what reached the model, why it was available, and what the result was allowed to cause. That's the engineering work behind a useful prompt. This opens my [context engineering series](https://substack.com/@soypetetech); next up is what happens when the information a model needs is missing, and how to distinguish that failure from a limitation of the model itself.
