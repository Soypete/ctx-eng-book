---
title: Engineering the Context Boundaries
subtitle: Give every failure boundary an owner, an executable check, and a record that makes repair possible.
module: ch02.03-future-ai-engineering
mood: professor
scheduled: 2026-11-09
---

A deployment assistant serves an outdated runbook. The data team says the pipeline ran, the application team says search returned a result, and the AI team says the model followed its instructions. Everybody's component worked. The user still got the wrong procedure.

This is an invented scenario, but it's a useful test of your team structure: who owns getting that procedure corrected, and who owns preventing it from reaching another user?

In the previous post in my [context engineering series](https://substack.com/@soypetetech), I argued that production reliability depends on contracts across the stack; now I want to assign the work of enforcing and maintaining them.

**Context** is everything a model receives for a request, including instructions, retrieved information, conversation, and tool results. **Context engineering** is the work of building and maintaining that information, its meaning, its state, and the rules governing its use.

My argument is simple: organize this work around failure boundaries. Each boundary needs an owner who can change its behavior and demonstrate that the repair holds. A new job title doesn't make that happen. Neither does putting “AI” in the team name. I've checked; the letters don't run tests.

## Own the handoff

A **boundary** is a point where application code checks information or authority before letting work proceed. A **contract** states what that boundary accepts, guarantees, and rejects.

For the runbook, “the pipeline ran” isn't the contract we need. We need to know which source supplied the procedure, which version it supplied, and when the source last confirmed it was current. **Provenance** is that record of where information came from. Copying yesterday's snapshot into today's index doesn't renew its authority.

The source owner must make those fields meaningful. The application owner must check them before assembling context. Both sides have work; assigning everything to “the AI team” hides the handoff.

Here's the ownership record I would start with for this scenario:

| Boundary | Contract to maintain | Repair responsibility |
|---|---|---|
| Source publication | Preserve source identity, version, and source-backed confirmation time. | Source owner corrects publication and metadata. |
| Context assembly | Reject evidence outside the request's freshness requirement. | Application owner maintains the check and unresolved response. |
| Source access | Verify the caller's permission before returning material. | Policy owner maintains access rules and denial tests. |
| Evaluation | Detect outdated procedures reaching the user after a change. | Evaluation owner maintains the failure case and release decision. |

These are responsibilities, not four required hires. One person can own several in a small team. What matters is that a rejected document has somewhere specific to go, and that the person receiving it has authority to repair the relevant component.

![Source evidence enters an owned contract check; allowed evidence reaches the model, while rejected evidence stops, and decision records route failures to an owner whose repairs feed regression tests.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch02.03-future-ai-engineering/diagram-1.png)

*Follow the rejected path: the check stops the request, and the record connects that failure to repair and a regression test.*

## Make ownership executable

A **decision record** stores what a check allowed or rejected and why. For this runbook, I want the boundary name, owner, policy version, source reference, source version, source confirmation time, and check time alongside the outcome.

That gives an engineer a starting point more useful than “the answer looked weird.” If the source confirmation is too old, the publishing owner can investigate. If the application accepts it anyway, the assembly check needs repair. If the confirmation time is recent but false, the producer's guarantee failed. The timestamp comparison cannot discover that lie by itself.

Drew Breunig describes the enduring engineering choice in [The State of Context Engineering](https://www.dbreunig.com/2026/10/24/the-state-of-context-engineering.html): deciding what belongs in deterministic code and what should be left to the probabilistic model. My extension is that somebody must own each of those decisions after launch. The freshness rule belongs in code, and its owner belongs beside it.

The [companion Go program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch02.03-future-ai-engineering) implement a small source-admission pipeline. A **pipeline** here is an ordered sequence of checks before downstream work runs. Each check declares its name, owner, policy version, and validation function. Construction fails if any of those are absent.

The program accepts evidence already authorized by the source service. It checks provenance and freshness, then calls a model substitute that simply returns the admitted text. No API key is required. Inside the loop, `b` is the boundary, `e` is the evidence, and `d` is its decision record, already populated with ownership and source metadata:

```go
err := b.Check(e, now)
if err != nil {
	d.Outcome, d.Reason = "rejected", err.Error()
}
decisions = append(decisions, d) // Keep the repair owner beside the decision.
if err != nil {
	return "", decisions, err // Rejected evidence never reaches the model.
}
```

The demo supplies an old source confirmation and prints a rejection assigned to `source-publishing`. It returns no answer, and the model substitute never runs. The tests assert that behavior, verify the decision metadata, and reject an ownerless configuration. They also admit current evidence, so a system that rejects everything cannot pass.

The fixture uses an illustrative one-hour maximum age. Choose your own policy from the question's needs and the producer's guarantee. For a rejected runbook, I would show “Current procedure unavailable” and route the protected source reference to its owner. Silently using the old procedure would undo the check.

These tests prove the fixture's control flow. They don't prove model accuracy, source truth, or that `source-publishing` is staffed. Production configuration needs a maintained team directory and an escalation path. An owner string is a useful requirement; it is not an on-call rotation.

## Assign the lifecycle, too

Ownership can't end at initial ingestion. **State** is information retained across requests or workflow attempts. If an assistant saves a runbook summary, correcting the source won't necessarily correct that saved summary. Someone must own revision, expiry, and recovery of those retained copies.

[Orogat and Mansour, 2026](https://arxiv.org/abs/2605.26252), frame long-term agent memory as a data-management workload and propose ingestion, revision, forgetting, and retrieval as state-level operations. That supports treating saved information as something with a lifecycle. It doesn't establish a new job title or prove that a particular team structure improves reliability.

For our scenario, I would add a failure case that corrects the source and then resumes a workflow holding the old summary. The state owner must specify whether that summary is invalidated, refreshed, or rejected. “The source was fixed” is insufficient evidence that the next answer will be fixed.

Store decision records under access and retention rules. Prefer protected references over copying runbook contents into general logs. The record should help the responsible engineer investigate without creating another uncontrolled data store.

Then evaluate the complete path after the repair: did the right procedure become available, did stale evidence remain blocked, and did the change introduce excessive delays or unnecessary refusals? A **regression test** preserves a failure case so later changes can be checked against it. Keep the owner attached to that case, including when the owner changes teams.

## Guidelines

1. **Assign an owner to each enforced contract.** Reject configuration missing the owner, rule version, or executable check; verify owner references against your team directory.
2. **Split producer guarantees from consumer checks.** Test both incorrect source metadata and correct metadata that the application mishandles.
3. **Keep repair metadata beside the decision.** Verify rejected requests retain the boundary, owner, reason, and source version without returning rejected content.
4. **Test the unresolved path.** Assert that rejected evidence never reaches downstream generation, and that current evidence can still produce a result.
5. **Carry repairs through retained state.** Resume a workflow with an obsolete saved summary and verify its documented invalidation or refresh behavior.
6. **Give every repaired failure a regression case.** Recheck task success, refusals, and latency after changes, and transfer test ownership when responsibilities move.

I want context engineering to produce work a team can inspect: a maintained source guarantee, an enforced boundary, a useful decision record, and a failure case somebody owns. My bet is that these responsibilities will outlast whatever titles we give them. Next in the [series: Tokens, Embeddings, and Attention](https://substack.com/@soypetetech), we'll look at how models represent and use the information those boundaries admit.
