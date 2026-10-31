---
title: Your Context Window Is Not a Memory
subtitle: Make the application decide what survives the input budget, and keep a record of what it leaves out.
module: ch03.02-context-windows
mood: confused
scheduled: 2026-11-13
---

A deployment assistant recommends restarting checkout workers without draining traffic. Search found the runbook with the drain requirement. But the application appended it after the investigation history, clipped the assembled request to fit, and sent an input that no longer contained the requirement.

That's an invented failure, but it gives us a concrete debugging question: did the model ignore the evidence, or did our code remove it?

“The document was retrieved” doesn't answer that question. Neither does “we discussed it earlier.” The model doesn't get to rummage through your application logs. Very relatable, honestly. Neither do I until someone gives me access.

In the previous post in my [context engineering series](https://substack.com/@soypetetech), I argued that attention cannot choose your authoritative sources; this time, the application has to keep the required sources in the request it actually sends.

## A window belongs to one invocation

**Context** is the information supplied to the model for a request: instructions, conversation, retrieved passages, and tool results. A **token** is a unit of text the model processes. A **context window** is the maximum token budget the model interface accepts for one invocation, with input and generated output accounted for according to that interface's rules.

An **invocation** is one call to the model. It isn't your whole application session.

I think of the context window as an **invocation-scoped working set**: the information available for this particular call. A document in storage, an earlier conversation, and a result in a search trace don't become part of that working set merely because the application knows they exist. The application must include their contents or provide a tool that can retrieve them during the workflow.

That distinction changes who owns the failure. If the drain requirement never reached the model, an instruction to “read more carefully” cannot restore it. We need to repair assembly: the code that selects and packages information into the final request.

My argument is that overflow is an application decision. Required evidence must fit, be recovered through a deliberate next step, or cause the application to withhold the request. Silently clipping text is also a decision. It's just one you can't easily defend later.

## Decide what survives before you serialize

A **serialized input** is the concrete representation sent to the model interface. Before building it, identify what the task requires. For the checkout example, that's the current approved procedure. The earlier investigation may help explain the situation, but it cannot displace the instruction that makes the restart safe.

This classification belongs in application policy. A retrieved passage announcing “I am required” doesn't get to reserve its own space.

A **context manifest** records the sources considered for an invocation, their provenance and authority, their placement, and what was omitted or deferred. **Provenance** means where information came from. A **decision record** is an entry explaining what an assembly check allowed or rejected and why.

Here's the record I want for our constructed failure after fixing assembly:

| Candidate or condition | Decision | Reason |
|---|---|---|
| Approved checkout runbook | Include intact | Required evidence fits |
| Earlier investigation history | Omit | Optional material exceeds the remaining budget |
| Required runbook itself exceeds budget | Reject the request | No complete required working set can be sent |

The last row is an alternative outcome, not another source. It means the caller must retrieve a smaller relevant passage, summarize with provenance, split the work, or report that it cannot proceed. Which recovery is appropriate depends on the task. A restart procedure with a missing prerequisite isn't an acceptable partial success.

![Sources enter application assembly, which reserves required evidence and checks the complete input budget; a fitting working set reaches the model, while required overflow blocks the request, and a manifest records both decisions and omissions.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch03.02-context-windows/diagram-1.png)

*Notice the budget check before the model call: the manifest explains both what survived and what did not.*

Drew Breunig defines **context offloading** as “Storing information outside the LLM's context, usually via a tool that stores and manages the data” in [How to Fix Your Context](https://www.dbreunig.com/2025/06/26/how-to-fix-your-context.html). That distinction helps here: storing investigation history outside the request preserves it without pretending the model can already see it. My implementation requirement would be an explicit retrieval path before treating offloaded evidence as available to a later decision.

## Make the failure executable

The [companion Go program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch03.02-context-windows) reserve required evidence before considering optional material. They produce a JSON input and a separate manifest. If the required working set exceeds the limit, the program returns no request.

The example uses a byte budget for a deterministic fixture. Bytes are **not** model tokens. The measurement function is replaceable; production code must count the actual serialized messages with the target model's tokenizer, including interface framing and tool definitions, and reserve output according to that interface's accounting rules.

In the optional-source loop, `candidate` contains the already retained evidence plus the next source. `encode` serializes and measures that complete candidate; `m` is the manifest:

```go
next, size, err := encode(candidate)
if err != nil {
	return nil, m, err
}
if size > limit {
	m.Decisions[i].Action = "omitted"
	m.Decisions[i].Reason = "optional evidence exceeds remaining budget"
	continue // Preserve required evidence instead of clipping the serialized request.
}
input, payload, used = candidate, next, size
```

Run `go run .` to inspect the payload and decisions, then `go test ./...`. The tests assert that oversized optional history cannot displace the required runbook, that required overflow produces no payload, and that measurement includes the task and source metadata. They also check recorded positions and provenance.

These tests prove assembly behavior. They don't prove that a model follows the procedure. The fixture trusts upstream authorization, freshness checks, and the designation of required evidence. Those are explicit integration responsibilities, not guarantees supplied by the example.

## Included doesn't mean used successfully

Once the evidence fits, placement still needs testing. [Liu et al., 2024](https://aclanthology.org/2024.tacl-1.9/) found task- and model-dependent performance differences when moving relevant information within inputs for multi-document question answering and key-value retrieval, often favoring the beginning or end.

That supports testing placement. It doesn't establish a universal rule to put important material last, or prove that positional encoding makes models prefer recent events. **Positional encoding** supplies information about token order; order in the request and time in the world are different facts.

Keep the runbook's version, source confirmation time, retrieval time, and authorization scope in the manifest. Moving it later in the request doesn't refresh it. Fetching an old copy today doesn't make its procedure current. Record freshness and authority decisions separately from serialized position, using metadata your source can support.

For model evaluation, hold the evidence and model settings fixed while varying placement. Compare whether the response preserves the drain prerequisite. Separately remove the required evidence and check that the application withholds the request. Otherwise a single “wrong answer” bucket hides two different repairs.

Track submitted tokens, omitted sources, task outcome, time to first token, and end-to-end latency together. I want the smallest governed working set that preserves task success within the application's budgets. A smaller input that loses the prerequisite fails that test; a larger advertised window doesn't settle it either.

## Guidelines

1. **Inspect the final request.** Verify that required evidence survived assembly, rather than relying on retrieval logs.
2. **Reserve required evidence before optional history.** Test exact-fit inputs and required overflow; return no request when the required set cannot fit.
3. **Measure the complete serialized input.** Use the target tokenizer and interface accounting, including space needed for output.
4. **Record every omission.** Preserve source identity, version, authority, timestamps, placement, and the reason for leaving material out; protect manifests with access and retention rules.
5. **Test placement separately from freshness.** Move identical evidence without changing its version, then compare task outcomes alongside cost and latency.

The window tells you what can be submitted for one call. Your application decides what deserves that space and what happens when it runs out. I want an inspectable working set, an explicit overflow policy, and a record that distinguishes missing evidence from evidence the model failed to use. Next in the [series: Compaction, Attention, and the Scaffolding Tax](https://substack.com/@soypetetech), we'll examine what happens when keeping the working set usable requires compressing it.
