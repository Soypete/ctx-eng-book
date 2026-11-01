---
title: Resolve the Meaning Before You Grant the Action
subtitle: Make purpose, identity, and scope explicit before a plausible guess becomes a database read.
module: ch04.02-computational-pragmatics
mood: confused
scheduled: 2026-11-20
---

“Analyze the past month's sales data and compare it to other monthly reports.” The assistant produces a convincing comparison using marketing data. The arithmetic works; the answer belongs to a different question.

This is a constructed failure, not an incident report. Nothing needs to be fabricated for it to happen. The application can retrieve real records, calculate correctly, and still deliver the wrong result because it never established what “the data” referred to.

In the previous post in my [context engineering series](https://substack.com/@soypetetech), I argued that examples hide specifications; here I want to expose the assumptions between a user's words and the action we permit.

My argument: **resolve the intended action and its scope before retrieval or execution.** A plausible interpretation is a proposal. It shouldn't quietly become permission.

## The missing step isn't better grammar

**Pragmatics** concerns what someone is trying to accomplish with an utterance and what may appropriately follow from it. “Can you check this?” looks like a question about ability, but usually functions as a request. **Speech-act interpretation** is figuring out that intended act.

**Reference resolution** identifies what an expression points to: which report is “this,” or which records count as “the data.” **Indexicality** is the dependence of expressions such as “my,” “here,” and “now” on the situation in which they are used.

These distinctions come from [Jurafsky's *Pragmatics and Computational Linguistics*](https://web.stanford.edu/~jurafsky/prag.pdf), which identifies reference resolution, speech-act interpretation, discourse structure and coherence, and abduction as core inferential problems. The latter two concern how statements fit together and what explains an utterance. That taxonomy doesn't tell us which failure dominates production systems.

For our sales request, the missing decisions are concrete. Does “past month” mean the previous calendar month or a rolling interval? Which sales dataset? Which comparison reports? Is the user asking for a read-only analysis or an update to a published report?

You can turn the sentence into immaculate JSON and leave every one of those questions unanswered. Congratulations, the ambiguity has a schema.

Direct wording helps expose missing decisions. It doesn't establish that the selected dataset is the intended one, or that the caller may read it. I wouldn't make users learn a miniature query language just to keep the product from guessing.

## Give interpretation an execution boundary

An **agent** is a model that calls tools in a loop. The **harness** is the application code that supplies its context and controls those calls. The same scope problem exists in a search box or summarizer without an agent loop.

I would put a **resolution gate** in that application code: a check that the intended operation, authenticated caller, target, and permitted scope are established before a read or write proceeds.

Some answers already exist. If the user selected the sales report in the interface, preserve that selection as evidence. If the product defines reporting periods in a particular business timezone, resolve the dates using that definition. When the available evidence leaves a consequential choice open, ask a focused question: “Previous calendar month or trailing month?”

Don't ask again when trusted state already answers it. Don't treat a model's confident completion of a missing field as trusted state either.

Drew Breunig describes the engineering choice in [The State of Context Engineering](https://www.dbreunig.com/2026/10/24/the-state-of-context-engineering.html) as choosing when to be deterministic—what goes in code—and when to be probabilistic—when the model drives. My application of that distinction is to let a model propose an interpretation while code controls whether its consequences are allowed.

![A request and trusted state feed scope resolution; missing fields lead to clarification, complete scope proceeds through policy checks, and only allowed requests reach retrieval. All outcomes produce a decision record.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch04.02-computational-pragmatics/diagram-1.png)

*Notice that resolving the meaning and authorizing the read are separate checks, both before retrieval.*

## “My files” needs an identity

An **authenticated principal** is the identity established by the application's authentication system. The word “my” does not supply one. Nor should a model invent one from a name mentioned in a conversation.

Resolve references against resources that principal is allowed to use. Then verify that the resolved operation and resource remain within policy. A user might legitimately access both marketing and sales; that makes both available, not interchangeable. Authorization alone cannot choose the right report.

For data engineers, this is where the catalog earns its keep. Bind the user's selection to a stable dataset identifier and a defined reporting interval. **Provenance** is the record of where information or a decision came from. Preserve the selection or clarification that resolved the reference, rather than logging only “sales query.”

Freshness is another contextual requirement I would make explicit. A request for the current report needs a source-backed account of when its contents were confirmed. Retrieval time alone doesn't supply that. If the source cannot meet the task's freshness requirement, return an unresolved outcome rather than silently using an older report. This is my proposed control, not a measured result from the pragmatics taxonomy.

## Make the boundary testable

The [companion Go program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch04.02-computational-pragmatics) implement a small version of this gate. They use invented inputs and a fake reader, a function standing in for a database read. There is no model call.

The request `r` contains the resolved action, dataset, time interval, and a reference to the evidence resolving the request. Session `s` supplies the caller; policy `p` supplies allowed datasets. A **decision record**, `d`, explains what the check did and why:

```go
// Missing scope must not become a best-guess query.
if r.Action == "" || r.Dataset == "" || r.ResolutionEvidence == "" ||
	r.Start.IsZero() || r.End.IsZero() || !r.Start.Before(r.End) {
	d.Outcome, d.Reason = "clarify", "resolve action, dataset, interval, and evidence"
	return "", d
}
// A complete interpretation still does not grant permission.
if r.Action != "analyze" || p.Version == "" || !p.Readers[s.Principal][r.Dataset] {
	d.Outcome, d.Reason = "denied", "operation or dataset outside policy"
	return "", d
}
```

Run `go run .` to see these outcomes:

| Request state | Decision | Database reads |
|---|---|---|
| Dataset and resolution evidence missing | Clarify | None |
| Marketing selected but policy allows only sales | Denied | None |
| Sales selected, interval resolved, policy allows analysis | Read | One scoped call |

Those are fixture results, not model performance statistics. `go test ./...` asserts the read counts, preserves the selected interval, and checks missing identity, unsupported actions, invalid dates, and source failure. It also checks that rejected requests return no data.

The example proves the ordering of checks. It cannot prove that a filled-in request reflects human intent. Its evidence reference is metadata; production code must obtain and verify that reference through a trusted selection or clarification flow. Letting model output populate every field would merely formalize the guess.

The example also stops at retrieval. It does not implement analysis, comparison-report selection, freshness validation, or authentication. Those belong in the real integration, with authorization enforced by the data service too.

## Keep the reason, not just the result

Retain the caller, resolved action, dataset, interval, resolution evidence reference, policy version, outcome, and reason. For a real read, add source version and source-confirmed freshness. Protect those records with access and retention rules; references can reveal sensitive activity even without copied rows.

Then test ambiguous requests as product behavior. The expected result might be a specific clarification, an existing selection reused, or a denied action. Include a case where both datasets are authorized but the intended one remains unresolved. That's the case that catches an application confusing access with meaning.

A deterministic gate test establishes a code property. Evaluating whether a model resolves language correctly requires reviewed requests and repeated model runs. Keep those claims separate.

## Guidelines

1. **Resolve action, caller, target, and scope before reading.** Assert that incomplete requests make no data calls.
2. **Reuse trusted selections and ask about missing choices.** Test that an ambiguous reference stays unresolved when no evidence selects it.
3. **Authorize the resolved operation independently.** Test a complete request for a forbidden dataset.
4. **Preserve resolution provenance and source freshness.** Verify that retrieval time cannot stand in for source confirmation.
5. **Record clarification and denial as distinct outcomes.** Check that each retains a reason and returns no data.

I want users to describe what they need without every unstated assumption turning into an invisible product decision. Resolve what matters, keep the evidence, and enforce the permitted consequence before acting. Next in the [series: Examples, Instructions, and Structured Outputs](https://substack.com/@soypetetech), I'll turn those requirements into interfaces that can be validated.
