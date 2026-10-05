---
title: Context Engineering Is Not Prompt Engineering
subtitle: Reliability starts with controlling what information and authority reach the model.
module: ch00-what-we-mean-by-context-engineering
mood: professor
scheduled: 2026-10-07
---

Picture an internal support bot: a customer-success rep asks how to explain a pricing change, and the answer quotes the compensation spreadsheet. Similarity search found relevant words. Nobody carried the spreadsheet's sensitivity and the rep's access scope into the selection process.

Adding “please don't leak salaries” to the prompt is not my incident-response plan.

## The problem is bigger than wording

When an AI feature misbehaves, editing the prompt is an easy first move. Sometimes it helps. But wording cannot repair the fact that we were  missing access metadata; we ha dan outdated pricing document, or our tool that accepts whatever account ID the model invents. This is all show we shape the context of a model. it is modling and configureing the information it needs to itneract and return a rsponse. so What is conext Engineering? 

Here's the definition I'm using:

**Context engineering is delivering the right information and control signals to a language process, at the right stage of a workflow, for an authorized purpose, within measurable limits.**

My argument is that reliability depends on controlling that information flow. More context is not automatically better context. This applies to the "classical"  AI Large Language Model activities extraction, classification, summarization, and search as much as to agents.

 my theory for context engineering is that there are 3 major compnents. I organize the work around these questions. They're the organizing lens for my [book](htt AI Large Language Model activities ps://github.com/Soypete/ctx-eng-book/blob/main/book/chapters/ch00-what-we-mean-by-context-engineering.md), not a new theory of language:

- **Lexicon:** What sources and entities exist? Who owns them? What's authoritative, current, sensitive, and available to this requester?
- **Semantics:** What do those entities mean here? Which identifiers and relationships distinguish a customer's contracted price from the current list price?
- **Pragmatics:** What may the actor do with that meaning? Which checks must an answer or action pass before it changes something outside the model?

The opening failure is a **Lexicon problem**. Ownership, sensitivity, and the user's data-access scope define what information is available. Enforce those facts at the data boundary. Checking whether the rep may actually change a customer's contract is the downstream Pragmatics question.

## Who owns selection?

By *harness*, I mean the application code that assembles model input and handles tool calls. That code owns the boundary, even when the model chooses when to search.

| Posture | Who chooses when to retrieve? | What the system controls |
| --- | --- | --- |
| Prompt-time compilation | Harness | Access scope, candidates, ranking inputs, cutoff, budget, provenance |
| Scoped tool retrieval | Model proposes; host authorizes | Identity, purpose, tool scope, result limits, validation, lineage |
| Hybrid | Harness seeds; model requests bounded gaps | Seed manifest, tool scope, per-call limits, final evaluation |

I compile ahead of time when the evidence and policy boundary are known. I expose a scoped tool when the next useful source depends on an intermediate observation. A hybrid supplies a useful starting set and bounded ways to fill gaps.

Adaptive retrieval adds extra turns, failed calls, and the possibility that the model never asks for the decisive document. Retrieved text can also contain instructions. Treat those as untrusted source content; the host must validate each proposed call without letting a document grant broader access.

This diagram shows the production target, including searches requested mid-loop:

![Requests and model tool calls pass through authorization before candidate selection, ranking, and context assembly.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch00-what-we-mean-by-context-engineering/diagram-1.png)

*Notice that the model's tool call returns through Authorize before another candidate set is built.*

## Put the admission check in code

The companion's `Compile` function filters **after search**: candidates arrive already scored. Post-filtering after top-k can starve the answer of eligible documents because restricted hits already occupied search-result slots. Production systems should push scope into the retrieval query through metadata filters or row-level security, then keep this admission check too.

Use your existing authorization infrastructure for `allowed`, backed by catalog entitlements. Ownership alone is not an access policy. Take the actor from authenticated session state and purpose from the server's approved workflow or route, never from model-generated arguments.

Here are the authorization and budget loops from the [full Go example and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch00-what-we-mean-by-context-engineering), with formatting condensed. Variable setup and the intervening sort are omitted; that sort breaks score ties by unique document ID so search return order cannot change selection.

```go
for _, d := range cands {
    if !allowed(m.Actor, m.Purpose, d) {
        m.Excluded = append(m.Excluded, d.ID+":unauthorized")
        continue
    }
    eligible = append(eligible, d)
}
// Sort eligible candidates here; budget only what the actor may use.
for _, d := range eligible {
    if d.Tokens > remaining {
        m.Excluded = append(m.Excluded, d.ID+":budget")
        continue
    }
    remaining -= d.Tokens; out = append(out, d); m.Included = append(m.Included, d.ID)
}
```

Reserve space for instructions, the question, and output before supplying the context budget. Run `go test ./...` in the example directory: the tests check that a higher-scoring restricted salary sheet is excluded, an oversized note is dropped, and tied scores resolve by ID. That's evidence for the admission behavior, not proof that the whole retrieval system works.

## Keep the data contract attached

The example's `Manifest` struct includes `SourceVersion`, `PolicyVersion`, and `RankerVersion`, plus actor, purpose, timestamp, budget, and inclusion/exclusion decisions. Replay also needs retained scored candidates, immutable source snapshots, policy and ranker configuration, compiler version, and the exact serialized input; version labels alone cannot reconstruct them. Attach a manifest reference to the request trace and protect the artifacts with access controls and a retention policy, including excluded IDs. Measure storage and write latency; reproducing selection does not promise identical model output.

The data platform owns the upstream Lexicon work: catalog sensitivity tags, access metadata carried alongside embeddings, scoped retrieval, and lineage from source record to chunk to manifest. The example's `Doc` has `Version` and `AsOf`; the caller must enforce freshness. Set an acceptable index age and assign an owner for refresh costs. Propagate permission changes separately from content re-embedding, and reject revoked access while the index catches up.

Structure matters too. Label sources and serialize evidence consistently so the model can distinguish a current contract from an old note. A tidy prompt cannot make that old note authoritative.

## Measure the boundary and the answer

I'd evaluate restricted near-matches, stale sources, expired access, tied scores, budget overflow, and tool results that try to redirect the model. Check selected inputs and exclusion reasons against expected results, then measure recall against relevant documents the requester may access. Track answer correctness, tool failures, latency, and token use alongside leakage: an empty context can pass a leakage check and still fail the user. These are checks to run, not measured improvements; some failures remain model-capability limits.

## Guidelines

1. **Authorize before you rank.** Make scoped retrieval the production target; test that restricted hits cannot crowd out eligible evidence.
2. **Write down who owns selection.** Choose compilation, scoped tools, or a hybrid and identify the host code enforcing it.
3. **Route every retrieval tool call through the data-access gate.** Validate identity, purpose, and scope independently of model arguments.
4. **Version and retain replay inputs.** Verify that protected artifacts reproduce selection and its budget decisions.
5. **Enforce freshness and revocation limits.** Assign refresh ownership and test access expiry during index lag.
6. **Evaluate eligibility and usefulness together.** Check leakage, authorized recall, answer quality, and resource limits on the same fixtures.

I want a system that can explain what reached the model, why it was available, and what the result was allowed to cause. That's the engineering work behind a useful prompt. This opens my [context engineering series](https://substack.com/@soypetetech); next up is what happens when the information a model needs is missing, and how to distinguish that from a limitation of the model itself.
