---
title: The Model Is Only One Layer
subtitle: Production reliability depends on the contracts that carry a request from source data to a usable answer.
module: ch02.02-production-ai-stack
mood: happy
scheduled: 2026-11-06
---

A deployment assistant tells you to restart a worker without draining it. The instructions are correct for the previous release. Your search index has both guides, and the application never checked which release the answer belonged to.

That's an invented failure, but the repair is concrete. Put the release check in the path that assembles the model's information. A more articulate explanation of the wrong procedure is still the wrong procedure.

In the previous post in my [context engineering series](https://substack.com/@soypetetech), I argued that “the AI failed” should start an investigation; here I want to show where the repair belongs in the production stack.

**Context** is everything a model receives for a request: instructions, retrieved information, conversation, and tool results. A **contract** is an explicit agreement about what a component accepts, guarantees, and rejects. My argument is that production reliability depends on enforcing those contracts across the complete path from source data to outcome.

The model participates in that path. It doesn't own all of it.

## Follow one request through the stack

Suppose the user asks for the restart procedure for checkout release v2. The application already knows the product and release from the deployment page. Those identifiers should travel with the request. Asking the model to reconstruct them from a conversation adds a guessing step we don't need.

The **data layer** stores and retrieves possible evidence. The **semantic layer** makes its meaning explicit: which product an identifier names, which release a guide applies to, and which source owns that claim. These can be tables and typed service responses. You don't need to buy a graph database to agree on what `release` means.

**Orchestration** is the application code coordinating retrieval, context assembly, the model call, and validation. **Governance** supplies rules about access, permitted actions, retention, and resource use across that path.

Here is how I would divide the work for this request:

| Responsibility | Contract for the deployment lookup |
|---|---|
| Data and meaning | Return authorized guides with exact product and release identifiers, source version, and source-backed time. |
| Context assembly | Admit only evidence matching the request and freshness policy; stop if none qualifies. |
| Model | Propose a passage identifier from the admitted evidence. |
| Output validation | Accept only an admitted identifier, then render its source passage. |
| Governance and measurement | Enforce access at the source and retain a protected record of the decisions. |

Hardware, storage, and model serving sit underneath this work. They constrain capacity, latency, throughput, and availability. A serving timeout deserves its own diagnosis. It doesn't explain why a successful response cited the wrong release.

This is the useful part of drawing the stack: each failure gets an owner and an executable check. The boxes aren't a shopping list.

![Authorized source guides pass through source-contract checks, then an interchangeable model and citation validation; rejected evidence stops, while decisions feed a protected record.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch02.02-production-ai-stack/diagram-1.png)

*Notice that source checks happen before the model, and output validation cannot restore evidence those checks rejected.*

## Retrieval isn't finished when search returns

**Vector search** finds material by similarity between numerical representations of content. Similarity is useful for finding deployment instructions. It cannot substitute for an exact release constraint or permission to read the document.

**Retrieval-augmented generation**, usually shortened to RAG, adds retrieved material to a model's context. In [RAGGED, Hsia et al., 2024](https://arxiv.org/abs/2403.09040), reader robustness to noise changes whether greater retrieval depth helps or hurts across the evaluated configurations. Adding more results is not a universal repair.

For our deployment request, I'd first enforce source access, then use exact product and release fields to restrict eligible evidence. Search can rank within that scope. If the question instead asks how restart behavior changed between releases, the scope changes. Evaluate those question types separately; don't silently reuse the single-release contract.

**Provenance** records where information came from and which source version supplied it. A fresh copy of an old guide is still an old guide. Keep source version, the time the source vouched for it, and ingestion time distinct. Someone has to own the publication process that makes those fields trustworthy.

This is where data engineering becomes part of answer quality. If a producer republishes an incorrect release identifier, the consumer's equality check will happily accept it. A contract needs producer validation and correction procedures as well as a reader that rejects bad input.

## Put the invariant in code

An **invariant** is a rule that must remain true through the operation. For this lookup, the returned passage must belong to evidence admitted for this product, release, and freshness policy.

Drew Breunig describes the engineering choice in [The State of Context Engineering](https://www.dbreunig.com/2026/10/24/the-state-of-context-engineering.html) as deciding what belongs in deterministic code and what should be left to the probabilistic model. That is the distinction I want here: let the model select a useful passage; let application code enforce which passages are eligible.

The [companion Go program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch02.02-production-ai-stack) implement a deliberately narrow passage lookup. Documents arrive already authorized from a trusted source adapter. The program rejects mismatched product or release fields, missing provenance, and source timestamps outside its policy before calling the model.

The `Model` interface accepts the request and eligible documents and returns a proposed document ID. The example uses deterministic substitutes that select the first or last document. They need no API key. A real provider adapter would translate the same inputs into its request format.

After the call, `id` is the proposed citation, `eligible` contains the admitted documents, and `out` holds the result and decision records:

```go
// A proposed citation cannot restore a document rejected during assembly.
for _, d := range eligible {
	if d.ID == id {
		out.Text = d.Text // render the approved passage, not invented deployment steps
		out.Decisions = append(out.Decisions, Decision{d.ID, d.Source, d.Version, d.AsOf, "validation", "returned", "citation admitted", p.Version})
		return out, nil
	}
}
out.Decisions = append(out.Decisions, Decision{Stage: "validation", Outcome: "rejected", Reason: "citation outside admitted evidence", PolicyVersion: p.Version})
return out, errors.New("citation outside admitted evidence")
```

With no admission check, the first-result substitute selects the previous release. With the check, that same substitute returns the current release's passage. Swapping in the last-result substitute preserves the checked outcome. A separate test proposes the rejected guide's ID anyway and verifies that no text returns.

The tests also verify that stale or incomplete evidence never reaches the model. The fixture's one-hour freshness policy is illustrative, not a production recommendation. When nothing qualifies, the lookup returns an error directing the caller to the documentation owner. I would show that unresolved state instead of silently falling back to an older guide.

This proves the checks on these fixtures. It doesn't prove real-model accuracy, source correctness, or the safety of a deployment. The program renders an existing passage and executes nothing. Free-form instructions would need additional evaluation; a valid citation alone cannot establish that generated prose follows its source.

## Keep the decision with the outcome

A **decision record** stores what a check accepted or rejected and why. Here it carries the document reference, source version, source time, stage, outcome, reason, and policy version. You can distinguish “no eligible guide” from “model proposed an excluded citation” without reading a paragraph of generated explanation.

In production, I would attach those records to the request and measure source retrieval, assembly, model execution, and validation separately. The [OpenTelemetry specification](https://opentelemetry.io/docs/specs/otel/) supplies interoperable telemetry primitives. Domain lineage and evaluated reliability objectives are still our responsibility.

Keep records under access and retention rules; don't turn general logs into a second documentation store. Evaluate the complete request after a change, including whether it returned the right procedure, rejected usable evidence, or exceeded its deadline. A faster model call is only useful if the resulting workflow still succeeds.

## Guidelines

1. **Carry exact identifiers through the request.** Test that a guide for another release cannot reach the model in a single-release lookup.
2. **Validate source metadata before assembly.** Reject missing provenance and invalid source times; verify that copying a snapshot doesn't renew its freshness.
3. **Keep admission rules outside the model adapter.** Run the same boundary tests against every replacement adapter.
4. **Validate proposed citations against admitted evidence.** Deliberately return an excluded ID and assert that no passage reaches the user.
5. **Record rejection reasons and source versions.** Make empty evidence, model failure, and failed output validation distinguishable in protected records.
6. **Evaluate the whole request after changes.** Check answer correctness and unresolved outcomes alongside latency, including the fallback when evidence is unavailable.

I want a stack whose contracts survive a model change and whose failures leave enough evidence to repair them. The model still needs evaluation; the surrounding system needs owners, checks, and a usable failure path. Next in the series, [Engineering the Context Boundaries](https://substack.com/@soypetetech) turns those responsibilities into engineering work a team can actually own.
