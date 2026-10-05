---
title: Your Agent Doesn't Need a Better Prompt. It Needs a Stop Condition.
subtitle: Loops, permission leaks, and runaway bills are missing state and constraints, and they belong in the harness, not the system prompt.
module: ch01.02-missing-state
scheduled: 2026-10-12
---

An agent needs three queries: look up a user's ID by email, fetch that user's orders, then fetch the order items. Simple. Except the column is `user_id` in one table and `userid` in another.

So the agent queries, fails, guesses a new column name, retries, gets a different error, guesses again. Nothing tells it to stop. It isn't failing loudly. It's quietly burning tokens in a loop, confidently trying the same broken idea with small variations.

Missing schema context started the loop. Missing state and enforcement are why it never ends.

## "Memory" is three different jobs

We say "memory" when we mean at least three things, and mixing them up is how workflows lose the plot:

- **Episodic state** is what happened: the request, the tool calls, the results, the outcome, the trace.
- **Semantic state** is the current facts you may recall later, each with a subject and a rule for replacing stale values.
- **Working state** is the bounded context assembled for *this* turn: the instructions, selected facts, summaries, and recent messages the model can actually see.

Some of that working state is a root set that must never be trimmed just because the budget is tight: the trusted instructions and the turns needed to interpret the current request. Everything else can be summarized, archived, superseded, or dropped, but those operations should be explicit and visible in the trace.

Durable state isn't "the model remembering." It's the harness deciding what to record, keep, retrieve, and put back in front of the model.

## Loops are missing exit criteria

A workflow loops when it can't tell that it has succeeded, failed, or run out of budget. The fix is to give it named ways to stop. A reliable loop ends in one of a handful of terminal outcomes:

- it **answers**,
- it **asks for clarification**,
- it **declines** an out-of-scope or unauthorized request,
- it **stops for missing evidence or permission**,
- or it **reports that the budget is exhausted**.

The model can suggest the next step. The harness owns the transition, and it records why.

You can put exit criteria in the prompt ("stop after 10 tool calls"), and you should, because it helps the model plan. But a sentence in the system prompt is guidance, not enforcement. The model can keep retrying for "higher confidence." A budget is real only when the trace shows the request that was rejected or the run that was terminated.

![The model proposes; the harness checks state, budget, and capability, then either executes or ends in a named outcome.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch01.02-missing-state/diagram-1.png)

*Every arrow out of the harness is a decision the model doesn't get to make.*

Here's what owning the loop looks like:

```go
type Outcome string

const (
	Answered  Outcome = "answered"
	Clarify   Outcome = "needs-clarification"
	Denied    Outcome = "denied"
	Exhausted Outcome = "budget-exhausted"
	Stuck     Outcome = "stuck-same-error"
)

// Run owns the loop. The model proposes each step; the harness decides whether
// there is another one, and records why it stopped.
func Run(step func(attempt int) (string, error), maxSteps int) (Outcome, string) {
	var last string
	for i := 1; i <= maxSteps; i++ {
		out, err := step(i)
		switch {
		case err == nil:
			return Answered, out
		case errors.Is(err, ErrDenied):
			return Denied, err.Error() // never retry around a permission decision
		case errors.Is(err, ErrAmbiguous):
			return Clarify, err.Error()
		case err.Error() == last:
			return Stuck, err.Error() // another attempt can't add information
		}
		last = err.Error()
	}
	return Exhausted, fmt.Sprintf("stopped after %d steps", maxSteps)
}
```

Point it at the schema-mismatch query and it returns `stuck-same-error` on the second identical error instead of spinning until someone notices the bill. The repeated error became a typed failure that ends the run, not permission to keep improvising.

## Permissions are context, not prompt text

Early agent automation leaned on impersonation: the agent acted as the user, inherited the user's tokens, and was told in the prompt what it was allowed to do. That's how you get prompt injection with real consequences. Someone convinces the bot to act for a person it doesn't represent ([Liu et al., 2024](https://arxiv.org/abs/2310.12815)).

An agent can have valid credentials, a working tool connection, and a legitimate goal, and still do the wrong thing, because the missing context is about authority:

- **Who** is it acting for: the user, itself, a service, another agent?
- **Which resources** belong to this task?
- **Which actions** are allowed: read, write, delete, purchase, send?
- **When** does that authority expire, and can a sub-agent inherit it?
- **Where** did this instruction come from: trusted task state, or a document it just read?

The biggest trap is confusing possession with permission. A workflow that finds a valid credential while processing a document has the *capability* to use it. That doesn't mean the current task is *allowed* to. Summarizing a customer account shouldn't let the agent export unrelated accounts just because the credential technically could.

Three failures look alike and need different fixes:

| Failure | What's wrong | Fix |
| --- | --- | --- |
| Missing authorization context | The system can't establish who, what, why, or until when | Resolve it from authenticated state and policy |
| Over-scoped context | Access check passes, but far more data enters the working set than the task needs | Retrieve the minimum task-relevant set |
| Execution-boundary failure | Context was scoped, but the downstream service accepts the action without rechecking | Enforce capabilities where the effect happens |

A prompt that says "ask before deleting" does not create an approval boundary. Code that resolves the principal, resource, action, and expiry from trusted state, and enforces that decision at retrieval *and* at execution, does.

## Cost overruns are unmeasured context

AI is billed by the token, and unbounded context is how bills spike. Picture a request that fans out into hundreds of retrieval calls and concatenates every result into the prompt. (That's a constructed example, not an incident report.) The missing invariant is obvious once you say it: nothing put an upper bound on fan-out or on assembled context.

Put a budget at every stage:

1. **Retrieval:** cap results and rerank to a small top set.
2. **Assembly:** cap tokens; prefer summaries where detail isn't needed.
3. **Prompt:** strip redundant boilerplate.
4. **Tracking:** record cost per request and attribute it to users and queries.

Then measure whether it worked: task success per request alongside retrieved tokens, model tokens, latency, and cost. A smaller bill with worse answers is not a win.

## Guidelines

1. **Name your terminal outcomes** (answered, clarify, denied, stuck, budget-exhausted) and make the harness end every run in exactly one.
2. **Treat a repeated identical error as a stop signal**, not a retry.
3. **Never retry around a permission denial.**
4. **Resolve who, what, which action, and until when from trusted state**, and enforce it at retrieval and again at execution.
5. **Separate episodic, semantic, and working state**, and never trim the root set to save tokens.
6. **Put a budget at every stage and log every rejection**, so the trace proves the limit was real.

Next post: real context failures, the scenarios and case studies, and how to tell an incident from a teaching example.
