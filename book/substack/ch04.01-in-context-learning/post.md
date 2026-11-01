---
title: Your Examples Are Hiding a Specification
subtitle: Make the assumptions behind your examples inspectable before tuning the prompt.
module: ch04.01-in-context-learning
mood: professor
scheduled: 2026-11-18
---

A ticket classifier labels “Please send pricing” as a buyer. The JSON parses, the dashboard updates, and sales gets another person who never agreed to buy anything. Every example in the prompt showed an obvious purchase; none explained where interest ends and commitment begins.

That's an invented failure, but it's a useful debugging case. The application asked for a label without deciding what the label meant. Beautiful curly braces. Undefined business logic.

In the previous post in my [context engineering series](https://substack.com/@soypetetech), I argued that compaction needs a recovery path to original evidence; now I want to make the assumptions in the examples we retain just as inspectable.

My argument: **examples belong in a versioned, testable specification of the task.** If an important distinction exists only in whoever selected the examples, the model has to supply its own interpretation. You then have to discover that interpretation by debugging the output.

## An example teaches more than a format

**Context** is the information supplied to a model for a request. **In-context learning** is using instructions and demonstrations in that context to guide a task without updating the model's trained parameters. **Few-shot examples** are a small set of input/output demonstrations included with the request.

Examples are useful. “Here is a ticket; here is its label” makes an interface concrete. But examples can leave the reason for the label unstated.

In our classifier, does `buyer` mean someone interested in pricing, someone with purchasing authority, or someone explicitly committing to purchase? All are plausible business categories. They are different tasks.

**Syntax** is how information is arranged and serialized. **Semantics** is what that information means. A JSON object can satisfy the first and fail the second. A schema permitting `buyer` and `unknown` cannot, by itself, decide whether a pricing inquiry establishes purchase intent.

I would write a **task contract**: the intended outcome, label definitions, required evidence, and conditions for an acceptable answer. For this invented application, `buyer` requires an explicit purchase commitment. Otherwise return `unknown`. That is a product decision, not a universal definition of a buyer.

Then I would choose examples that expose the distinction:

| Input | Output | Reason |
|---|---|---|
| We will purchase the annual plan. | `buyer` | Explicit commitment |
| We are evaluating the annual plan. | `unknown` | Evaluation does not establish commitment |
| Please send pricing. | `unknown` | Pricing inquiry does not establish commitment |

An **abstention** is an explicit refusal to assign a substantive label when the required evidence is missing. Here, `unknown` is that outcome. It must be an acceptable result in the application, or somebody will eventually “fix” the prompt to stop producing it.

## Keep the evidence narrower than the promise

[Brown et al. (2020)](https://arxiv.org/abs/2005.14165) demonstrated strong few-shot behavior through text interaction at scale. That supports taking examples seriously as an interface. It does not establish that every additional example improves your classifier, or that less reasoning always means more reliability.

**Inference** has two meanings here: running a trained model, and drawing a conclusion from evidence. Neither is inherently the bug. The problem I want to expose is consequential interpretation that nobody specified or evaluated.

My engineering argument is that making those assumptions explicit makes failures easier to investigate. I am not claiming a measured reliability gain for this ticket classifier. That would require a defined workload and an evaluation comparing alternatives.

Drew Breunig makes a related point in [The Problem is Prompt Debt](https://www.dbreunig.com/2026/06/22/the-problem-is-prompt-debt.html): “specify your system's behavior with measurements, not prose.” That strengthens the practical requirement here: writing down a label definition is a start; checking whether outputs satisfy it is the next step.

A longer prompt is not a substitute for deciding what counts as correct.

## Retain what the model actually received

**Invocation context** is the information assembled before one model call. **Workflow context** is information obtained between calls, such as retrieved records or tool results. A dynamically retrieved ticket becomes fixed invocation context when the application sends that call.

Version the stable pieces: label definitions, instructions, and example sets. Retrieve changing evidence when needed. Record both in the same trace.

I want the exact serialized request, including example order, rather than a log saying “used buyer prompt.” Keep the source identity and version beside it. **Provenance** records where the evidence came from; it lets us distinguish a changed ticket from a changed interpretation of the same ticket.

![A task contract, ordered examples, and source evidence form a retained invocation snapshot; model output is checked against reviewed labels and produces an accepted or rejected decision record.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch04.01-in-context-learning/diagram-1.png)

*Notice that the check depends on a reviewed expected label, not merely on whether the output parses.*

Freshness and governance still apply. A timestamp for assembling a prompt doesn't establish when the source last confirmed its contents. An example copied from an old policy doesn't become current because you pasted it today. Check source permissions and freshness before assembly, and protect retained requests under the same access and retention rules as their contents.

These are controls I would put around the task. In-context learning does not supply them automatically.

## Make the disagreement executable

The [companion Go program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch04.01-in-context-learning) implement the small, deterministic part of this argument. They retain a serialized invocation with its rule, ordered examples, versions, source reference, and input. A content digest identifies those retained bytes. The digest alone cannot reconstruct the request.

The program then checks scripted candidate outputs against the reviewed answer for our pricing inquiry. `out.Label` is the decoded candidate label, `expected` is the fixture's reviewed label, and `d` is a **decision record** explaining what the check accepted or rejected:

```go
// Valid JSON alone does not establish the intended meaning of a label.
if out.Label != "buyer" && out.Label != "unknown" {
	d.Reason = "label outside contract"
	return d
}
if out.Label != expected {
	d.Reason = "label disagrees with reviewed fixture"
	return d
}
d.Action, d.Reason = "accepted", "label matches reviewed fixture"
return d
```

Run `go run .` to see a syntactically valid `buyer` answer rejected and an `unknown` answer accepted. Run `go test ./...` to verify that distinction and that changing example order, example content, the rule, or the source version changes the retained snapshot's identity.

| Candidate or assembly condition | Recorded outcome |
|---|---|
| Required rule or provenance missing | Assembly error; no snapshot emitted |
| Valid JSON, wrong reviewed label | Rejected; semantic disagreement |
| Unsupported label | Rejected; outside contract |
| Expected abstention | Accepted; reviewed label matched |

This proves that the application can preserve the inputs and distinguish formatting success from this semantic failure. It does not prove that examples improve a model's accuracy. There is no model call, and the expected label is available because this is an offline test. Production requests don't arrive with an answer key. The timestamp is recorded metadata, not an implemented freshness check.

For a real integration, retain the actual provider request, model identity and settings, response, and evaluation result. Replaying the same inputs helps investigate a failure; it does not guarantee identical model output.

## Evaluate the distinction you meant to teach

**Prompt optimization** here means changing discrete instructions or examples and measuring the result. It can include example selection, ordering, wording, and constraints. Changing all of them at once makes the result harder to attribute.

Build a labeled evaluation set that includes commitments, inquiries, ambiguous language, and abstentions. Keep evaluation cases separate from demonstrations used to tune the prompt. Compare candidate prompts across diverse inputs and repeated runs; a single satisfying answer is not a reliability result.

Measure errors for each label and inspect the boundary cases. If you add languages or domains, evaluate those separately. A model can attempt a newly described category without retraining, but that is a fast first experiment. Compare it with rules or trained alternatives before deciding it meets the requirement.

## Guidelines

1. **Define each label and its evidence requirement.** Include an explicit outcome for missing evidence.
2. **Add examples at the disputed boundary.** Test an inquiry beside a commitment, not just obvious purchases.
3. **Retain exact invocation inputs and artifact versions.** Verify that changing example order appears in the trace.
4. **Check source access and freshness before assembly.** Keep source confirmation separate from request creation time.
5. **Score meaning separately from formatting.** Include valid JSON with the wrong label in your tests.
6. **Evaluate prompt changes on held-out cases.** Compare label errors and abstentions across repeated runs before adopting a revision.

I want examples that make the task's assumptions visible, retained inputs that make failures investigable, and tests that tell me whether the intended distinction survived. The model still has to interpret language; our job is to specify and evaluate the consequential parts. Next in the [series: Computational Pragmatics](https://substack.com/@soypetetech), I'll examine how purpose and situation change what an instruction means.
