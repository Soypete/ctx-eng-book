---
title: Every Context Failure Has an Address
subtitle: Find the boundary that let bad context through, enforce the check, and keep evidence of the decision.
module: ch01.03-context-failure-case-studies
mood: broken
scheduled: 2026-10-30
---

A sales assistant recommends a discount to close a deal. The deal is already closed. The CRM knows; the search index that syncs daily still says “negotiating.”

You can rewrite the prompt all afternoon. Yesterday's data will still be yesterday's data.

**Context** is everything the model receives for a request: instructions, retrieved data, conversation, and tool results. An **agent** is a model that calls tools in a loop. The **harness** is the application code that builds its context, runs those tools, and decides whether another step happens.

This continues my [context engineering series](https://substack.com/@soypetetech). Last time I argued that application code has to decide when an agent stops. Here I apply that habit to five failures: find where the system let the failure happen, fix it there, and record whether the fix held.

## Give the failure an address

A **boundary** is a point where code checks information, authority, time, or cost before letting work proceed. A **decision record** stores what that check allowed or rejected and why.

“The model got it wrong” is a symptom. I want an address I can file a ticket against.

| Failure | Boundary that failed | Missing control |
|---|---|---|
| Endless debugging | Harness stop decision | Definition of done |
| Over-fetched customer data | Retrieval and authorization | Task-specific rows and fields |
| Unbounded bill | Retrieval and assembly | Result and token limits |
| Stale data | Freshness | Maximum age per question type |
| Authorization drift | Calls between components | Verification of the caller's authority |

![Request passes through harness, retrieval, freshness, and assembly checks before reaching the model; component calls also check authority, and each boundary writes a decision record.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch01.03-context-failure-case-studies/diagram-1.png)

*Follow the request toward the model: each check has a place to enforce its rule and record its decision.*

These are constructed scenarios, not incident reports. The customer-data case is a recurring pattern I've watched engineers build at work. They are debugging hypotheses until tested, not measured proof of reliability.

## Endless debugging: define enough evidence

A debugging agent gets “why did service X fail?” Every log points to another error. It keeps digging until a limit kills the run.

Define supporting signals before the run. For example, require an error naming the failed dependency plus a deployment event in the same window before presenting a deployment-related hypothesis. That supports a hypothesis; it doesn't prove causation. Missing or conflicting evidence should produce an unresolved result.

Have the harness enforce tool-call, time, and output limits regardless of the model's confidence. Record the termination reason, attempts, evidence references, and checks passed. “Budget exhausted” and “evidence requirement met” must be different outcomes.

## Over-fetched customer data: authorization is not relevance

I've seen teams retrieve every row a user may read, paste it into context, and ask the model to find what matters. That's a missing design decision wearing a prompt.

Authorized data can still be unnecessary. Extra fields consume **tokens**, the units of text models process, and widen exposure through logs, answers, or **prompt injection**: instructions hidden in material the model reads.

A **capability** is a short-lived, server-issued grant specifying whose data, which fields, which action, and until when. One concrete implementation could pass an opaque capability identifier that the protected service resolves against a trusted server-side record. For customer X, that record could permit reading name, email, and tickets until its expiration. The service checks the authenticated caller, customer, requested fields, action, and expiry before querying. A model-written scope string is not authority.

Apply existing row- and column-level policies too; a capability doesn't replace them. Trusted assembly code then selects only what this task needs. A ticket-summary request can omit email even when access permits it. Its decision record says which caller requested customer X, which capability was checked, which fields were permitted and selected, and which source supplied them.

Keep field meanings and authoritative sources in a catalog, with a **data contract** specifying expected fields, types, versions, and freshness. If `stage` becomes `deal_stage`, validate and explicitly map the supported version or reject the mismatch. Don't ask the model to guess what a schema change meant.

## Unbounded bill: budget the path

A documentation assistant using **retrieval-augmented generation (RAG)** searches documents and adds results to context. Without limits, query expansion—turning a question into several searches—multiplies work, and assembly concatenates full documents.

Cap search expansion and candidate results. Budget reranking, which re-scores results for relevance, against latency. Limit assembled passages and model output separately. Record retrieved versus used tokens, latency, and retrieval and model cost against the request.

Choose limits using representative tasks and retain their success results beside cost. A cheaper path that loses the needed passage isn't a fix. The module supplies no measured budget settings; I wouldn't invent defaults and call them evidence.

## Stale data: check before using it

Back to the closed deal. Give each question type a maximum acceptable age; the module's pricing example requires data less than an hour old. Compare that requirement with a timestamp backed by the source, not the time an indexer copied a row.

In the Go example below, `ObservedAt` means “when the source last vouched for the value.” It is not automatically a source commit timestamp or an index run timestamp. An old, unchanged record can still be current; a recently copied replica can still lag. The source adapter must establish that guarantee. Keep source version, observation time, and retrieval time distinct.

The [companion Go program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.03-context-failure-case-studies) implement only freshness; the other boundaries follow the same pattern: check at the boundary, record the decision.

`Resolve` receives a cached fact, question-type policy, and source-reading function `fetch`. After looking up the maximum age, `maxAge`, it fills decision record `d`:

```go
d.MaxAge, d.Age = maxAge, now.Sub(cached.ObservedAt)
if d.Age <= maxAge {
	d.Action, d.Reason = UsedCache, "within max age"
	return cached, d, nil
}
fresh, err := fetch(ctx, cached.Subject)
if err != nil { // don't fall back to the stale value: an old answer looks just like a current one
	d.Action, d.Reason = Rejected, err.Error()
	return Fact{}, d, fmt.Errorf("%w: %s is %s old (max %s): %v", ErrStale, cached.Subject, d.Age, maxAge, err)
}
d.Action, d.Reason = Refetched, fmt.Sprintf("cached copy %s old exceeds %s", d.Age, maxAge)
d.SourceVersion, d.ObservedAt = fresh.SourceVersion, fresh.ObservedAt
return fresh, d, nil
```

For a failed refresh, the record contains query type `deal-stage`, source `crm`, cached source version, observation and check times, age, maximum age, action `rejected`, and the source error. The caller receives no value. Show the rep “Current deal status unavailable; retry” and withhold the recommendation. Correctness costs availability here.

The example trusts a successful `fetch`; it doesn't revalidate the returned timestamp. Production adapters must validate that guarantee, bound refresh concurrency and retries, and honor a request deadline. On success, this record combines the old cache age with the refreshed version and observation time; retain separate before-and-after metadata if you need full lineage.

Upstream, capture source changes to update or invalidate the index, and agree on a producer freshness guarantee. That reduces stale reads; it doesn't remove the read-time check or justify unbounded load on the CRM.

## Authorization drift: check the caller

Suppose research can read raw data, analysis aggregates, and reporting only summaries. If reporting calls analysis and analysis uses its own access, reporting can receive aggregates it wasn't allowed to see. That's a **confused deputy**: a component using its authority for a caller who lacks it.

Pass the caller's capability through each call. Analysis and the protected data service both check it. A summary-only capability must reject an aggregate request; record caller, callee, capability reference, requested operation, denial reason, and whether data was returned.

This isn't unique to agents. Their model-selected tool calls make it especially important that authority comes from trusted code. A tool description that instructs the model to request broader access must not be able to mint that access.

## Make the repair observable

Store decision records as structured events linked to the request in an access-controlled audit store. Define retention and deletion rules. Record field names and protected references rather than customer values or reusable capability credentials; sanitize source errors before storing them.

Evaluate every step, not just the answer. For repeated runs of each fixture, report denied-data returns, stale-value returns after failed refresh, budget overruns, and task success. Any forbidden return or overrun fails that fixture. Choose and disclose the repeat count; don't imply a deterministic unit test establishes model reliability.

Deliberately bypass a check and confirm the evaluation fails. Turn successful injections into regression cases, distinguishing user input, retrieved content, and tool descriptions as entry points.

## Guidelines

1. **Name the failed boundary** before changing the prompt.
2. **Define evidence and stop conditions** and record which ended the run.
3. **Authorize and minimize retrieval** before assembling context; validate its data contract.
4. **Budget each stage** and compare task success alongside cost.
5. **Verify source-backed freshness** and reject stale values when refresh fails.
6. **Check caller capabilities at every handoff** and retain governed decision records.

I want a repair I can test: a named boundary, an enforced check, and evidence of its decision. Next in the [series](https://substack.com/@soypetetech): personalization, where context must fit one user without exposing another user's state.
