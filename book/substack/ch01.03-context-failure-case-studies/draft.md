---
title: Every Context Failure Has an Address
subtitle: Five production failures, the boundary each one crossed, and the record that proves the fix worked.
module: ch01.03-context-failure-case-studies
mood: broken
scheduled: 2026-10-14
---

A sales assistant tells a rep to offer a discount to close a deal. The deal closed yesterday. The CRM knows that. The assistant doesn't, because it reads from a search index that syncs once a day, and the index still says "negotiating."

Nobody's prompt was wrong. The model reasoned perfectly from what it was given. What it was given was a day old.

This is post 4 in my [context engineering series](https://substack.com/@soypetetech). Last time I argued that application code, not the prompt, has to decide when an agent stops. This time I apply the same habit to five failures I see over and over: find the place in the system where the failure happened, fix it there, and keep a record that proves the fix held.

## Give the failure an address

Some vocabulary first. **Context** is everything the model receives for one request: instructions, retrieved data, conversation, tool results. An **agent** is a model that calls tools in a loop, and the **harness** is the ordinary application code around it that builds the context, runs the tools, and decides whether there's another step.

A **boundary** is a check in that code, a point where something crosses from one part of the system to the next and code can inspect it: data, authority, time, or cost. The harness deciding whether to take another step is a boundary. So is a database query, a freshness check on cached data, the step that assembles the context, and a call from one service to another. A prompt can describe a rule. Only a boundary can enforce one.

"The model got it wrong" is a symptom, not an address. The address is the boundary that let bad context through:

| Failure | Boundary that failed | What was missing or unbounded |
|---|---|---|
| Endless debugging | The harness's stop decision | No definition of done |
| Over-fetched customer data | Retrieval and authorization | Every readable record entered the context |
| Unbounded bill | Retrieval and assembly | No limit on results or tokens |
| Stale data | Freshness | Cached data outlived its source |
| Authorization drift | Calls between components | A callee used broader authority than its caller |

![Five boundaries between the request and the model, each writing a decision record.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch01.03-context-failure-case-studies/diagram-1.png)

*Every fix sits at a boundary, and every boundary writes down what it decided.*

These are constructed examples built on common production architectures, not reports of specific incidents. The second is a pattern I've watched engineers build at work.

## Endless debugging: no definition of done

A production-debugging agent gets "why did service X fail?" and starts reading logs. Every log points to another error. Without a definition of done, it keeps digging until a time or tool-call limit kills it, having spent the budget without an answer.

"Find the root cause" isn't a definition of done. Tell the model which signals count as supporting evidence, then have the harness enforce tool-call, time, and output limits. Don't stop on the model's self-reported confidence; that isn't calibrated evidence. Require named supporting signals and record which checks passed.

## Over-fetched customer data: authorization is not relevance

I've seen teams query every row the user is allowed to read, paste it into the context, and tell the model to figure out what matters. That hands a missing design decision to the model.

"The user may read this" and "the model needs this for this task" are different questions. Raw rows don't say what their fields mean, which source is authoritative, or what the task requires. Every extra field costs tokens (the units of text a model reads and bills by) and widens the damage from a logging mistake, an overly chatty answer, or a **prompt injection**: instructions hidden in data the model reads. If the query also lacks a customer filter, it's no longer a relevance problem. It's an authorization failure.

The repair has three parts. The retrieval request carries a **capability**: a short-lived, server-issued grant naming whose data, which fields, which action, and until when. Trusted code then selects the task-specific rows and fields before anything reaches the model. And the context says what the data means: field definitions, source, freshness.

## Unbounded bill: no budget at any stage

A documentation assistant built on **retrieval-augmented generation (RAG)**, meaning it searches documents and adds the results to the context, has no limits. Query expansion (rewriting one question into several searches) fans out. Every result gets concatenated. Full documents go in where a paragraph would do. As usage grows, the bill grows faster than the useful answers.

Budget each stage separately: cap the search results you consider, rerank them (re-score the top results with a slower, more accurate model) within a latency limit, assemble passages up to a task-specific token limit, and attribute cost to the request. Then check that task success held. A cheaper path that loses the right answer isn't a fix.

## Stale data: back to the sales assistant

The CRM updated in real time; the index refreshed daily; nothing compared the two. The repair is a freshness boundary. Every retrieved fact carries when the source last vouched for it, and each type of question gets a maximum age. A deal-stage question might tolerate an hour; a product-docs question, a week. Too old means fetch it again. If the source can't answer, return nothing rather than the old value, because a stale answer looks exactly like a fresh one.

Here's that check from the companion Go example. `Resolve` gets the cached fact, the maximum age for this question type (`maxAge`), and a `fetch` function that reads the source. It fills in a decision record, `d`, as it goes:

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
```

The [full program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.03-context-failure-case-studies) replay the closed-deal scenario. The returned record also carries the source and its version, set earlier in the function. A question type with no policy gets an error instead of the cached value; when in doubt, the check refuses.

## Authorization drift: the confused deputy

Picture three components (agents, services, or workflow steps; it doesn't matter): research can read raw data, analysis can read aggregates, reporting can read only summaries. Reporting asks analysis for help. If analysis answers with its own broader access instead of what reporting is allowed to see, reporting just got raw data. That's a **confused deputy**: a component using its own authority on behalf of a caller who doesn't have it.

Telling each component "only share what the caller may see" makes the boundary advisory. Instead, the caller passes its capability with every call, and the callee and the service that touches the data verify it. Advisory enforcement doesn't hold up well in related research: a 2026 survey of execution-security work on coding agents reports that the policy-enforcement studies it reviewed found failure rates from 69% to 98% against real denylists, meaning lists of forbidden commands ([Rashidi, 2026](https://arxiv.org/abs/2607.05743)). That's a range across those studies, not a measurement of this scenario.

## A fix you can't observe is a hypothesis

Each repair is a guess until you can show it held. So each boundary writes a decision record:

- **Stop decision:** why the run ended, how many attempts, which signals supported the answer.
- **Retrieval and authorization:** who the request acted for, the capability's scope, which fields were selected, and where they came from.
- **Budgets:** tokens retrieved versus used, latency, cost, and whether the task succeeded.
- **Freshness:** source version, observed time, retrieval time, and the decision.
- **Calls between components:** caller, callee, the capability passed, the decision, and the effect.

Then test it like any regression. Run the same task repeatedly and record every step, not just the final answer. Deliberately break one boundary and confirm your tests notice. When a prompt injection gets through, turn it into a test case and record how it arrived: user input, retrieved content, or a tool's description.

## Guidelines

1. **Name the boundary before you change the prompt.** For every incident, write down which boundary let the bad context through.
2. **Separate "may read" from "needs for this task."** Select task-specific rows and fields in code before assembling the context.
3. **Budget every stage.** Cap results and assembled tokens, and track task success next to cost.
4. **Give every cached fact an age and every question type a maximum.** Refetch when it's too old; return nothing if the source can't confirm.
5. **Pass capabilities on every call.** The callee checks the caller's grant, not its own.
6. **Write a decision record at every boundary.** If you can't show what the boundary decided, you can't show the fix worked.

The model is rarely the place to start. Find the boundary, fix it in code, and keep the record. Next in the series: personalization, where the context has to fit one user without leaking another user's data.
