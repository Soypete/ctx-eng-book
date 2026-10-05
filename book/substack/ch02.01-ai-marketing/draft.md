---
title: Stop Saying "The AI Failed"
subtitle: "AI" names a model, a product, a workflow, even a company. Reliability lives in the system around the model, and that system outlasts whichever model you pick.
module: ch02.01-ai-marketing
mood: professor
scheduled: 2026-10-19
---

"The AI failed" is my least favorite line in a postmortem. Which AI? The model? The search step that fed it? The code that turned its answer into an action? The whole product?

"AI" gets used for a model, a product feature, a workflow, and sometimes an entire company. That's fine in marketing. It's a problem in an incident review, because it lumps very different failures into one bucket, and the bucket's default fix is "try a better model."

This is post 6 in my [context engineering series](https://substack.com/@soypetetech). **Context** is everything a model receives for one request, and context engineering is the work of deciding what goes into it, from where, and with what limits. The last few posts walked through failures that looked like model mistakes: runaway agent loops, stale data, and one user's data showing up in another user's answer. Every one was fixed outside the model. This post zooms out to say why.

## A better question than "did the AI fail?"

When something goes wrong, I ask three narrower questions, each named with a term borrowed from linguistics:

- **Lexicon (missing information):** Did the system have the right sources, entities, and definitions? Does someone accountable vouch for them, and are they current and authorized for this request?
- **Semantics (unclear meaning):** Did the system make clear what those things mean in this domain, for this task, at this time? Or did the model have to infer identities, schemas, and categories from incidental text?
- **Pragmatics (wrong use):** Did the system allow the wrong use of an output that was otherwise reasonable? Are the output format, purpose, and permissions enforced in code outside the model?

These aren't layers. They're three lenses you can point at any part of the system. Time cuts across all three: sources change, definitions get versioned, permissions expire. You'll see that in the example below, where the reference data carries an as-of date.

Put together, the lenses describe a pipeline: authoritative sources become structured meaning, the model works from that meaning, and code constrains what happens with the result.

![Sources become structured meaning, the orange model box works from it, and validation decides whether the result is routed or sent to human triage.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch02.01-ai-marketing/diagram-1.png)

*Swap the orange model box and everything else is still there, still doing the work that makes the answer trustworthy.*

## Models matter. They just can't carry the system.

I'm not going to tell you models are interchangeable commodities. They aren't. They differ in quality, cost, latency, what input they accept, and how you can deploy them, and the rankings shift with every release and benchmark.

My claim is narrower: you can't hand production behavior to the model alone. Whatever model you use, something still has to retrieve the right data, track state, define what things mean, enforce who may do what, and check whether the outcome was right.

Databases are a useful comparison. Engines differ, and picking one matters. But no application became reliable just by choosing Postgres. The pipelines, transformations, access rules, and operations around it decide whether the platform works. A model isn't a database, but the point holds: neither component carries the reliability of the whole system.

## No agent required

Most of this series so far has featured agents, models that call tools in a loop. The lenses don't need one. Take something boring: a support-ticket classifier that gives each ticket a product label and a severity label so it lands in the right queue.

The model can be excellent and the system can still fail three ways:

- **Lexicon:** The product catalog is stale. The model labels a ticket with a product that was retired last quarter, or one that never existed.
- **Semantics:** "Severity" has no written definition. The model guesses what "sev2" means, and its guess doesn't match what the on-call engineer means.
- **Pragmatics:** Nothing stops the classifier from emitting a label the routing rules don't recognize, like "P0," so the ticket lands nowhere.

A better model fixes none of those. A catalog with an owner, a version, and an as-of date does. So does a severity taxonomy with definitions sent to the model alongside the labels, and routing code that accepts only labels it can verify.

The [example](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch02.01-ai-marketing) builds exactly that in Go. The model `m` is a parameter: anything with a `Classify(ctx, prompt) (Label, error)` method. The checks belong to the system `s`, which holds the catalog, the taxonomy, and the routing rules:

```go
if now.Sub(s.Catalog.AsOf) > s.Catalog.MaxAge { // a great model can't fix a stale list of products
	return Result{Queue: TriageQueue, Reason: ErrStaleCatalog.Error()}, ErrStaleCatalog
}
l, err := m.Classify(ctx, s.prompt(ticket))
if err != nil {
	return Result{Queue: TriageQueue, Reason: err.Error()}, err
}
res := Result{Label: l, CatalogVersion: s.Catalog.Version}
if !s.Catalog.Products[l.Product] {
	res.Queue, res.Reason = TriageQueue, fmt.Sprintf("%v: %q", ErrUnknownProduct, l.Product)
	return res, ErrUnknownProduct
}
```

`TriageQueue` is the human triage queue. The full program checks severity against the taxonomy the same way. The tests run two fake models through the same system: a careful one, and a confident one that invents "P0" and labels tickets with a retired product. The second model's answers go to human triage with a reason attached instead of being routed on its word. Swapping the model changed the labels it proposed. It didn't change the checks they had to pass.

## The prompt is one interface

Prompts are where most teams spend their tuning effort, and they do matter: wording, examples, and output format all shape what the model does. But the prompt is one replaceable interface inside a larger system, and prompts that were tuned for one model often need rework for the next. When that happens, you lose a little.

The knowledge that's hard to rebuild lives in the rest of the system, and it maps onto the same three lenses. Retrieval and durable state hold which sources are authoritative and what's current (Lexicon). Schemas and taxonomies hold how your domain defines its terms (Semantics). Authorization, operational feedback, and evaluation hold who may do what and which outcomes count as correct (Pragmatics). That work keeps paying off when you switch models next quarter.

So when the postmortem says "the AI failed," push back with the three questions. Missing information? Unclear meaning? Wrong use? You'll usually find the answer, and the fix, outside the model.

## Guidelines

1. **Ban "the AI failed" from incident reviews.** Name the gap: missing information, unclear meaning, or wrong use.
2. **Give reference data an owner, a version, and an as-of time.** Refuse to run on a catalog past its max age.
3. **Send definitions with labels.** If the model must pick a label, put what each label means in the context.
4. **Validate model output against what the system recognizes.** Unknown labels go to the human triage queue with a reason, never silently routed or dropped.
5. **Make the model a parameter.** Keep data, meaning, and permission checks in code that doesn't change when the model does, and test them with more than one model.

Models will keep changing. The system around them is the part you own, and the part that decides whether the feature works. Next in the series, I'll turn that system into a map: the parts of a production context stack and which failures show up in each.
