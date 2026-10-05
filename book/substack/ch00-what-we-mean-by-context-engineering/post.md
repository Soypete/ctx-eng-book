---
title: Context Engineering Is Not Prompt Engineering
subtitle: Reliability starts with controlling what information and authority reach the model.
module: ch00-what-we-mean-by-context-engineering
mood: professor
scheduled: 2026-10-07
---

Picture an internal support bot: a customer-success rep asks how to explain a pricing change, and the answer quotes the compensation spreadsheet. Similarity search found relevant words. Nobody carried the spreadsheet's sensitivity and the rep's access scope into the selection process.

Adding “please don't leak salaries” to the prompt is not my incident-response plan.

![A support bot answers a pricing question with rows from the salary spreadsheet.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch00-what-we-mean-by-context-engineering/scene-1.png)

## The problem is bigger than wording

When an AI feature misbehaves, I can edit the prompt. But wording cannot repair missing permissions, stale pricing, or a tool accepting account IDs the model invents. So what is context engineering?

**Context engineering is delivering the right information and control signals to a model, at the right stage of a workflow, for an authorized purpose, within measurable limits.** Control signals include instructions and tool limits: the task and the authority to cause an action.

This is the first post in a series drawn from my [book on context engineering](https://github.com/Soypete/ctx-eng-book/blob/main/book/chapters/ch00-what-we-mean-by-context-engineering.md). For support bots, extraction, classification, summarization, and search, reliability depends on controlling information flow. More context is not automatically better context.

I use a lens with three questions:

- **Lexicon:** What sources and entities exist? Who owns them? What's authoritative, current, sensitive, and available to this requester?
- **Semantics:** What do those entities mean here? Which identifiers and relationships distinguish a customer's contracted price from the current list price?
- **Pragmatics:** What may the actor—the authenticated requester, here the rep—do with that meaning? What must be checked before the support bot issues a discount?

The opening failure is a **Lexicon problem**: ownership, sensitivity, and the user's access scope define what information is available. Access scope describes the requester's permissions; access metadata records the restrictions attached to documents. This post works the Lexicon boundary in code; Semantics and Pragmatics get their own posts.

Let the support bot search based on earlier results, and you have an agent: a model calling tools in a loop to decide what to do next. The **harness** is the application code around the model that builds its input and runs its tools. That's where context engineering happens. I'll call its check of what text the requester may access **Authorize**.

## Who owns selection?

Compiling context means assembling the model's input from selected information. Search returns **candidates**, documents considered for inclusion; the harness decides what fits in the compiled context.

| Posture | Who chooses when to retrieve? | What the harness controls |
| --- | --- | --- |
| Prompt-time compilation | Harness, before generation | Access scope, candidates, ranking, budget, source records |
| Scoped tool retrieval | Model proposes; harness authorizes | Identity, purpose, access scope, tool limits, validation |
| Hybrid | Harness supplies a starting set; model requests gaps | Starting set, access scope, tool limits, final evaluation |

Our hypothetical support bot used prompt-time compilation with no access scope in the query. Authorize should exclude the salary sheet before candidate selection.

I compile ahead of time when the evidence and access scope are known. I expose a scoped tool when the next useful source depends on an earlier result. A hybrid combines both.

Scoped tool retrieval adds turns, failed calls, and the risk that the model never asks for the decisive document. It also exposes **prompt injection**: retrieved text that tries to redirect the model with instructions. The harness must validate each proposed call without letting untrusted source text grant broader access.

The manifest is the record of what was selected or excluded and why.

![Requests and model tool calls pass through Authorize before candidate selection, ranking, and context assembly.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch00-what-we-mean-by-context-engineering/diagram-1.png)

*Notice that the model's tool call returns through Authorize before another candidate set is built.*

## Put Authorize in code

I put a [runnable Go example and tests in the book repo](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch00-what-we-mean-by-context-engineering). `Compile` selects scored search results for the compiled context. Tests use a pricing document, a higher-scoring restricted salary sheet, and an oversized note.

The diagram shows the production target, where Authorize scopes the search itself; the Go example shows the recheck inside `Compile` **after search**, needed when query filters are missing or stale. Post-filtering after **top-k**, the k highest-scoring results search returns, can starve the answer: restricted hits already occupied slots that eligible documents needed. Production systems should push access scope into the retrieval query using metadata filters or row-level security, database rules that restrict which rows a requester can read. Keep Authorize in `Compile` too.

Here, `allowed` is the permission check; back it with existing authorization infrastructure. Ownership alone is not permission. Take actor identity from authenticated session state and purpose from the server's approved workflow, never from model-generated arguments.

In this excerpt, `m` is the manifest, `cands` holds search results, and `remaining` is the context budget. Setup, sorting, and inclusion logging are omitted. The full code sorts eligible documents by score and breaks ties by unique document ID before applying the budget.

```go
for _, d := range cands {
    if !allowed(m.Actor, m.Purpose, d) { // Recheck access before using text.
        m.Excluded = append(m.Excluded, d.ID+":unauthorized")
        continue
    }
    eligible = append(eligible, d)
}
for _, d := range eligible {
    if d.Tokens > remaining {
        m.Excluded = append(m.Excluded, d.ID+":budget")
        continue
    }
    remaining -= d.Tokens
    out = append(out, d)
}
```

Reserve room for instructions, the question, and output before budgeting. Run `go test ./...` in the example directory: tests check exclusion of the higher-scoring salary sheet, rejection of an oversized note, and ordering by ID when scores tie. These tests cover selection only.

## Keep source history and permissions attached

After that incident, I’d use the manifest to investigate why the salary sheet reached the model. The example’s `Manifest` records source, policy, and ranker versions, actor, purpose, timestamp, budget, and selection decisions. I link that record from the request trace and protect it, including excluded IDs. Replay needs retained source versions, scored candidates, and configuration; version labels alone cannot reconstruct decisions. Set a retention limit and measure the storage and write cost.

Whoever maintains search owes the harness sensitivity tags, access metadata, source versions, and freshness. Carry those alongside embeddings—the numeric representations used for similarity search—and preserve links from source records through searchable text to manifests. Assign refresh ownership and cost, set a maximum age for searchable copies, and check each document's `Version` and `AsOf` fields. Update permissions separately from rebuilding embeddings, and reject revoked access while search is catching up with source changes.

## Measure the boundary and the answer

I'd test restricted near-matches, stale sources, expired access, score ties, budget overflow, and tool results containing prompt injection. Compare selected inputs and exclusion reasons with expected results, then measure authorized recall: how much relevant, permitted evidence the system retrieves. Track answer correctness, tool failures, latency, and token use alongside leakage. An empty context can pass a leakage check and still fail the user.

## Guidelines

1. **Authorize before you rank.** Push access scope into production retrieval; test that restricted hits cannot crowd out eligible evidence.
2. **Write down who owns selection.** Choose compilation, scoped tools, or a hybrid and identify the harness code enforcing it.
3. **Route every retrieval tool call through Authorize.** Validate identity, purpose, and access scope independently of model arguments.
4. **Version and retain replay inputs.** Verify that protected records reproduce selection and budget decisions.
5. **Enforce freshness and revocation limits.** Assign refresh ownership and test expired access while search catches up.
6. **Evaluate eligibility and usefulness together.** Check leakage, authorized recall, answer quality, and resource limits on the same test cases.

I want a system that can explain what reached the model, why it was available, and what the result was allowed to cause. That's the engineering work behind a useful prompt. Next in my [context engineering series](https://substack.com/@soypetetech): distinguishing missing information from limitations of the model itself.
