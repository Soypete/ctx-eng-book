---
title: Stop Saying "The AI Failed"
subtitle: A support-ticket classifier shows why reliability needs current data, explicit definitions, and checks beyond the model.
module: ch02.01-ai-marketing
mood: mean
scheduled: 2026-11-04
---

A checkout outage lands in the support queue as a cosmetic issue. The classifier returned perfectly valid labels. “The AI failed” tells me almost nothing about what to fix.

That's an invented incident, but it's the distinction I want in a postmortem: did we supply bad information, leave the task ambiguous, accept a bad decision, or get the wrong answer from the model despite doing the surrounding work?

In the previous post in my [context engineering series](https://substack.com/@soypetetech), the wrong user's greeting traced back to which records the system retrieved.

**Context** is everything a model receives for a request. **Context engineering** is deciding what goes into it, from where, and under what limits. My argument here is narrower than “every failure is context”: locate the failure before prescribing a model upgrade.

## Give the failure an address

Take a support-ticket classifier. It proposes a product and severity; application code uses the severity to select a queue. No autonomous workflow required.

I ask three questions. The names in parentheses come from linguistics, but the engineering work is familiar: data quality, shared definitions, and enforced rules.

- **Missing information (Lexicon):** Is the product catalog current, complete, and authorized for this request? Who owns its publication and refresh deadline?
- **Unclear meaning (Semantics):** Does “sev2” mean a degraded feature with a workaround, or something else? Send the definition with the label.
- **Wrong use (Pragmatics):** Can an unrecognized label trigger routing? Who is allowed to route this ticket, and what happens when validation fails?

These are diagnostic questions you can ask at any step, not a sequence of processing stages. The classifier's actual sequence is simpler: load approved data, assemble the context, call the model, validate its proposal, then route or send it to a person.

A stronger model may follow definitions better or decline to guess. It cannot guarantee that your catalog is current or your routing permissions are enforced. Conversely, a correct catalog cannot guarantee a correct classification. Both deserve tests. Neither gets a participation trophy.

## Put the contract before the call

For this classifier, I would publish a versioned catalog from the authoritative product registry through a transformation and snapshot job. Assign that job an owner, a refresh deadline, and an access policy. Publish the severity **taxonomy**—the labels and their definitions—as a versioned contract too.

The catalog's `AsOf` should describe how current the source data is, not when a loader copied it. Keep load time separately. Reloading yesterday's snapshot today must not reset its freshness clock. Preserve provenance: the source and job run behind each published version.

In a request handler, authenticate the caller, then select the catalog allowed for their organization. If caching it, include that verified organization and catalog version in the cache key. The model doesn't get to choose another organization's catalog. This is the same identity boundary as the previous post, applied to classification.

The Go example below starts after that boundary, with a catalog already in memory. Its freshness check is a local comparison, not a database query per ticket. In production I would measure snapshot loading, model-call time, and routing separately before deciding where latency needs work.

![A versioned catalog supplies definitions to the orange model box; application validation then routes recognized labels or sends unknown labels to human triage.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch02.01-ai-marketing/diagram-1.png)

*The model proposes labels; the surrounding system supplies their meaning and decides whether they may be routed.*

Freshness has a catch: a retired product can still pass membership validation if the loaded snapshot says it's active. The age gate rejects an expired snapshot, but cannot detect a retirement inside the allowed age. I would invalidate affected snapshots on retirement and keep the maximum age as a backstop.

Rejecting stale data also has a cost. This example sends every affected ticket to human triage. Before shipping that policy, I would assign support ownership of that queue, alert the catalog owner on stale-data failures, and monitor triage arrival rate and oldest waiting ticket. Failing closed without staffing the fallback is just an outage with better intentions.

## Make the boundary executable

The [companion Go example](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch02.01-ai-marketing) keeps the existing classifier behind a `Model` interface: anything implementing `Classify(ctx, prompt) (Label, error)` can supply a proposal. `s` is the system holding the catalog, severity definitions, and routing rules; `m` is the selected model. `TriageQueue` names the human triage queue.

This excerpt is from `System.Route`:

```go
if now.Sub(s.Catalog.AsOf) > s.Catalog.MaxAge { // a great model can't fix a stale list of products
	return Result{Queue: TriageQueue, CatalogVersion: s.Catalog.Version, Reason: ErrStaleCatalog.Error()}, ErrStaleCatalog
}
l, err := m.Classify(ctx, s.prompt(ticket))
if err != nil {
	return Result{Queue: TriageQueue, CatalogVersion: s.Catalog.Version, Reason: err.Error()}, err
}
res := Result{Label: l, CatalogVersion: s.Catalog.Version}
if !s.Catalog.Products[l.Product] {
	res.Queue, res.Reason = TriageQueue, fmt.Sprintf("%v: %q", ErrUnknownProduct, l.Product)
	return res, ErrUnknownProduct
}
```

The full function also checks that the severity has both a definition and a route. The prompt is a typed Go value, not a hidden prose trick. For the checkout fixture its contents are:

> Ticket: checkout returns 500 for every card  
> Products: payments, search  
> sev1: customer-facing outage or data loss  
> sev2: degraded feature with a workaround  
> sev3: question or cosmetic issue

A provider adapter would serialize those fields. I would use structured output—a schema specifying the response fields and allowed labels—where supported, then validate against the request's catalog and routes in application code. A correctly shaped response still needs permission to become an action.

The example sends the active product list and definitions on every call. For a larger catalog, I would select an authorized, task-relevant subset and evaluate whether that selection omits the right product. There is no measured size threshold here.

Catalog versions now survive stale-data and model-error returns as well as successful calls. Production decision records should also carry taxonomy and routing-policy versions. Keep historical labels tied to their original definitions when severity meanings change; don't silently reinterpret them. Give those records access and retention rules without copying raw ticket text into general logs.

## Valid is not correct

Suppose the ticket describes checkout failing for every card, but the model returns `search / sev3`. Both labels exist. The route exists. Validation passes, and the ticket goes to the wrong queue.

The example's added evaluation test makes that failure explicit. A **golden set** is a collection of tickets with expected labels. This tiny, invented set runs through the same routing code with both a keyword-based fake model and a fake that always answers `search / sev3`. The test reports exact matches by expected label pair and a **confusion matrix**, a count of expected-versus-predicted pairs. It asserts that the wrong but valid answer passes routing while failing the expected classification.

That proves the distinction, not real-model accuracy. Run the same evaluation structure on representative, reviewed tickets for each candidate model, including ambiguous cases and severe incidents. Compare missed urgent tickets separately from overall accuracy before a swap.

Drew Breunig argues, “Your eval is the most valuable AI asset you own, not your model or your prompts,” in [Your Eval is More Important Than the Model](https://www.dbreunig.com/2025/01/08/evaluating-llms-as-knowledge-banks.html). Here, that distinction is concrete: the routing checks survive a swap, while the evaluation tells us whether the replacement makes better decisions.

I think about this like a data platform. Choosing a database matters; choosing Postgres doesn't make the whole application reliable. Likewise, I would budget for prompt retuning when changing models, while preserving the source contracts, definitions, access boundaries, and evaluation cases around them.

## Guidelines

1. **Name the failing boundary in the incident review.** Record the supplied context, proposed labels, validation result, and actual routing outcome before choosing a fix.
2. **Test freshness against source time.** Reload an old snapshot and verify that it still fails; test retirement invalidation inside the maximum age too.
3. **Version definitions with their consumers.** Check that every severity has a definition and route, and preserve the versions needed to interpret old decisions.
4. **Enforce scope before context assembly.** Test that another organization's catalog cannot reach the model through the request or cache path.
5. **Exercise the fallback.** Inject an expired catalog and verify human triage, a recorded reason, and an alert to the responsible owner.
6. **Evaluate valid but wrong answers on every model swap.** Report expected-versus-predicted labels alongside rejection counts; passing validation is not an accuracy score.

I want “the AI failed” to start an investigation, not finish one. The model matters, and so does the code that supplies its facts and acts on its answer. Next in the series, [The Production Context Stack](https://substack.com/@soypetetech) maps those responsibilities so each failure has somewhere concrete to land.
