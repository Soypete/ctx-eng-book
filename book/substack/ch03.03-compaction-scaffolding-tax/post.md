---
title: Your Summary Is Not the Source of Truth
subtitle: Compaction earns its cost when the workflow can recover the detail it left out.
module: ch03.03-compaction-scaffolding-tax
mood: broken
scheduled: 2026-11-16
---

A support assistant tells a customer their annual plan qualifies for a refund. The original policy says annual plans are nonrefundable after activation. The conversation summary says “refunds are available within 30 days,” and the application deleted the original evidence.

That's an invented failure. The summary sounds reasonable. It also removed the exact detail the next decision needed. Asking the model to read more carefully is going to be a very expensive way to discover that the sentence is gone.

In the previous post in my [context engineering series](https://substack.com/@soypetetech), I argued that the application must decide what survives the context window; now I want to examine what survives summarization.

My argument: **a summary should be a derived artifact with a recovery path, never an automatic replacement for its source.** If you measure only how much smaller it got, you're missing the failure that matters.

## The next question wasn't in the summary's job description

**Context** is the information supplied to a model for a request. The **context window** is the token budget the model interface accepts for one call; **tokens** are units of text the model processes. **Compaction** reduces the material carried forward, through filtering, selected excerpts, or a generated summary.

Compaction has a job. A long support conversation might become a short handoff note. That note can preserve the customer's issue and the next action while leaving out policy exceptions that didn't matter to the handoff.

Then the next request asks about an exception.

A summary is produced for an anticipated use. It cannot preserve every possible future distinction while also promising to discard information. When the application deletes the source, omitted information is unavailable to later calls unless another source can supply it. That loss belongs to the system's retained information, not some mysterious failure of model memory.

Drew Breunig defines context summarization as “Boiling down an accrued context into a condensed summary” in [How to Fix Your Context](https://www.dbreunig.com/2025/06/26/how-to-fix-your-context.html). He separately describes context offloading as storing information outside the model's context. That separation is useful here: reducing what the model receives and retaining what the application can recover are different decisions.

I want the summary to help find the evidence. I don't want a handoff note quietly promoted to company policy.

## Keep the source and the transformation inspectable

A **derived artifact** is information produced from another record, such as a summary of a policy. **Provenance** records where that information came from. **Rehydration** means retrieving original evidence when the compact representation lacks the detail a task needs.

For our refund example, the application retains the policy and records which version produced the summary. Before answering an eligibility question, it retrieves the policy's exceptions. The trigger comes from the task's requirements. Waiting for the model to announce that something is missing assumes it knows what disappeared.

A **context manifest** records the working set and its transformations. For compaction, I want source IDs and versions, source authority, freshness information, access scope, compaction method and version, and the material deferred for recovery. The manifest tells us what happened to the evidence. It cannot replace that evidence.

![A retained authoritative source produces a compact summary and manifest; a later eligibility request recovers the source, checks it, and either supplies original evidence or withholds the answer.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch03.03-compaction-scaffolding-tax/diagram-1.png)

*Notice that the recovery path reaches the retained source; the summary alone cannot restore its missing exception.*

A **decision record** explains what a check allowed or rejected and why. For this workflow, I would expect records like these:

| Condition | Decision | Reason |
|---|---|---|
| Summary created | Defer exceptions; retain source reference | Handoff text carries only the general rule |
| Eligibility requested | Recover original exceptions | This task depends on omitted detail |
| Source missing | Withhold answer | Summary cannot establish eligibility |
| Source version changed | Require summary rebuild | Recorded transformation describes an older version |
| Current access or freshness check fails | Withhold answer | Retention does not establish permission or currency |

These are application policies I'm proposing, not measured findings about every summarizer. You might support a different version-change workflow. Make that choice explicit enough to test.

Freshness needs particular care. A newly generated summary of an old policy is still derived from an old policy. Keep source confirmation time distinct from summary creation time. Likewise, a source ID in a manifest is a retrieval reference, not an access grant. Recovery must enforce the caller's current permissions. Retaining evidence means retaining it under the source's access and retention rules.

## Make the missing exception executable

The [companion Go program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch03.03-compaction-scaffolding-tax) are a small runnable demonstration of this failure. They use an invented refund policy, with a general rule and an annual-plan exception. A deterministic compaction function keeps the general rule and deliberately drops the exception.

No model call is necessary to prove that dropping the source removes this recovery path.

The summary carries a manifest. `Resolve` handles an eligibility request; `fetch` reads the referenced source while checking current access and freshness. `d` is its decision record. This excerpt comes directly from that function:

```go
// Eligibility depends on exceptions; never promote the summary to authority.
p, err := fetch(s.Manifest)
if err != nil {
	d.Action, d.Reason = "withheld", err.Error()
	return "", d
}
if p.ID != s.Manifest.SourceID || p.Version != s.Manifest.SourceVersion {
	d.Action, d.Reason = "withheld", "source identity or version changed; rebuild summary"
	return "", d
}
d.Action, d.Reason = "rehydrated", "eligibility requires original exceptions"
```

The rest of the function selects the applicable exception from the recovered policy. Run `go run .` to see the summary, recovered answer, and decision records. Then run `go test ./...`.

The tests compare the recovered answer with the uncompacted policy's exception. They delete the source and require an empty answer instead of a fallback to the general rule. They also reject changed source identity or version, denied access, and invalid source freshness.

This proves deterministic recovery and refusal behavior. It doesn't prove that a model writes faithful summaries or uses recovered evidence correctly. The example's in-memory storage and scope check also aren't a production authorization system. Its refund window and freshness limit are fixture choices, not recommended defaults.

## Pay for recovery, then measure what it buys

I first heard **scaffolding tax** from [Saul Ramirez](https://www.linkedin.com/in/saulramirez/) during a Utah meetup talk. That's my recollection, not a claim that he coined it. I use the phrase for the latency, tokens, storage, code, and extra failure modes introduced by orchestration around a model.

Keeping sources, recording transformations, and retrieving exceptions all charge that tax. My argument is that these layers earn their place when later tasks need details compaction removes. A successful recovery matters more than a pretty compression ratio. It still has a cost.

Compare each compaction policy against an uncompacted baseline on the tasks you actually support. Include questions that depend on an exception, event order, a disputed attribution, a revoked permission, an unresolved question, or a failed attempt that must not be repeated. Deliberately place the needed detail before compaction, then ask for it afterward.

Track task success alongside token reduction, latency, storage, tool calls, source-recovery rate, contradictions, and whether output claims retain provenance. The Go fixture checks one exception and its recovery path. A model evaluation must also check whether the final answer respects that exception across repeated runs. Record what each layer costs and which failures it prevents before deciding to keep it.

## Guidelines

1. **Retain authoritative sources under an explicit retention policy.** Delete one in a test and verify that dependent decisions stop.
2. **Version every compaction transformation.** Record source identity, version, authority, freshness, access scope, method, and deferred material.
3. **Trigger recovery from task requirements.** Make exception-dependent tasks retrieve exceptions before answering.
4. **Recheck access and freshness during recovery.** Test revoked access and stale sources; neither should return a usable answer.
5. **Evaluate lost details against the original evidence.** Include exceptions and chronology, and compare task success alongside operational cost.

I want compaction that makes the working set smaller while keeping necessary evidence recoverable. That requires a retained source, an inspectable transformation, and a tested decision about when to retrieve more or stop. Next in the [series: In-Context Learning](https://substack.com/@soypetetech), I'll look at how examples change what the model infers from the working set we give it.
