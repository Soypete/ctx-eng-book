---
title: Your Agent Doesn't Need a Better Prompt. It Needs a Stop Condition.
subtitle: A model can propose another attempt, but application code must decide whether it may continue.
module: ch01.02-missing-state
mood: broken
scheduled: 2026-10-12
---

An agent—a model calling tools in a loop—needs to look up a user by email, fetch their orders, then fetch the order items. Simple. Except the column is `user_id` in one table and `userid` in another.

Say it hits a schema mismatch, guesses another column name, and gets another error. The messages change. The problem doesn't. Nothing tells it to stop, so now you've got a very expensive spelling bee.

Missing schema context—the table names, columns, and relationships needed to write the query—started the loop. Missing state and enforcement are why it never ends.

In my context-engineering series, I've asked whether information was available, understood, and usable. Now: what happens when the workflow still can't make progress?

The **harness** is the application code that sends context to the model, executes the tool calls it proposes, and decides whether there's another step. It keeps **attempt history**: a record of previous attempts and their results. It enforces limits before allowing another action. That check between what the model proposes and what actually runs is the **execution boundary**.

Its **trace** is the per-step record of what was proposed, what ran, and why the run stopped. That's where I want proof that the boundary worked.

## Give the loop somewhere to end

The model can suggest the next step. The harness decides whether to execute it or end the run, and records why.

I use these five terminal outcomes—the named ways a run ends—in the example:

- `answered`: the workflow produced its result.
- `needs-clarification`: an ambiguity requires input from the user.
- `denied`: the proposed action lacks a **capability**, a task-scoped grant to perform an action.
- `stuck`: attempts aren't adding information, or an error isn't safe to retry.
- `budget-exhausted`: another step would exceed the budget, a limit on steps or tokens—the units of text processed by the model.

The example represents capabilities with a list of allowed tools. Neither permission nor budget should be invented by the model.

An **error class** identifies a stable kind of failure even when its message changes. The harness counts consecutive failures in the same class and stops at a configured threshold.

![The model proposes a step; the harness checks state, budget, and capability, then executes or selects a terminal outcome.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch01.02-missing-state/diagram-1.png)

*In the full application, the harness enforces each decision; an injected instruction is untrusted text trying to redirect the workflow, covered below.*

A sentence in the system prompt is guidance, not enforcement. Tell the model the limits so it can plan. Then make the harness reject actions that cross them.

The accompanying Go example demonstrates this loop. `Run` takes a function that proposes the next step and one that executes it. These excerpts show its checks before execution and its stopping threshold after error classification:

```go
if !lim.Allowed[s.Tool] {
	return Denied, fmt.Sprintf("%s: %s not granted to this task", ErrDenied, s.Tool) // never retry around a permission decision
}
if spent+s.Tokens > lim.MaxTokens {
	return Exhausted, fmt.Sprintf("token budget %d would be exceeded at step %d", lim.MaxTokens, i)
}
// Later: after classifying the execution error and updating the counter.
if sameClass >= lim.MaxSameClass {
	return Stuck, fmt.Sprintf("%d attempts in a row failed with %s; another guess adds no information", sameClass, se.Class)
}
```

The [full program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.02-missing-state) include the omitted counter and all five outcomes. `ErrDenied` labels the harness's permission rejection; `ErrAmbiguous`, returned by the execution function, signals clarification. Successful execution produces `answered`; that function still has to validate its result.

This example demonstrates a tool allowlist, estimated step-token limits, and a returned outcome and reason. It leaves persistent traces, customer-level permissions and expiry checks, and actual model-usage accounting to the application around it.

The schema test varies the guessed column name but returns the same PostgreSQL error class, SQLSTATE `42703` for an undefined column. `Run` ends `stuck` at the configured threshold, regardless of message text.

That's a deliberate policy for blind schema guessing, not proof that every repeated error is hopeless. A corrected query could succeed. Stopping the loop doesn't repair the missing schema information. The team responsible for the database must supply the current catalog of tables and columns, and the harness must include the relevant schema before querying again.

## Remember what failed, and where facts came from

The harness can detect that loop because its attempt history retains the previous error class outside the model's current input. Trimming conversation history must not erase the retry counter.

For this workflow, I separate the attempt history from the **working state**: the instructions, facts, and messages assembled for the current model call. A schema fact reused later also needs provenance and a freshness rule. I'd store it as `subject, value, source, observed_at, expires_at, supersedes`. The subject identifies the table or column the fact describes; the source identifies its authoritative catalog. The observation and expiry times determine when to recheck it, and `supersedes` points to the older fact it replaces.

Keep the old observation in history while selecting the current fact for working state. Otherwise, yesterday's schema becomes today's confident mistake.

In the application around it, I'd record a stop event shaped like `run_id=… tool=sql.query outcome=stuck error_class=42703 reason=repeated_class`, linked to the preceding attempts. Give the team operating the harness ownership of that event stream, with access and retention rules. Don't dump customer rows or credentials into logs to prove a loop stopped.

## The same boundary owns permission

Now give that order-query agent a credential found in a query result. Possession doesn't authorize it to export unrelated accounts. 

**Trusted state** is the authority record: who the task acts for and what it may do, resolved from the authenticated session, policy store, and verified workflow decisions. Retrieved text cannot rewrite it. Retrieval simply means fetching data for the task; our SQL queries already do it.

If that query result also tells the agent to export accounts, it is attempting prompt injection ([Liu et al., 2024](https://arxiv.org/abs/2310.12815)). It must not widen the task's grants or limits. A related authorization failure is the **confused deputy**: a service is tricked into using its own authority for someone else's unauthorized action. Telling the model to behave doesn't constrain the service.

To choose a fix, I separate where permission handling breaks:

| Failure | What's wrong | Fix |
| --- | --- | --- |
| Missing authorization context | The system can't establish who it acts for, the task purpose, resources, actions, or expiry | Resolve these from trusted state; include delegation and required approvals |
| Over-scoped context | Access is allowed, but unnecessary data enters working state | Select only task-relevant rows and columns |
| Execution-boundary failure | The service performing the action doesn't recheck permission | Check the task's grant where the effect happens |

For the order lookup, enforce retrieval access through database row-level security or a policy-filtered view. The tool call lands in a **tool service**, the code that performs the requested operation. Before an export executes, that service must check who it acts for, the customer scope, permitted action, and expiry again. A grant to read orders must not quietly become a grant to export accounts.

Enforce those checks in the tool service even after the harness allows a call. Every execution path must pass through the execution boundary.

## A budget is another stop condition

That same order lookup can spend too much without ever hitting a schema error. Say it fetches every order and appends every item to each successive model call. Successful queries can still build an unbounded bill.

Cap rows returned during retrieval. During assembly—building the working state for the model call—select relevant results and summarize detail the task doesn't need. Before sending the prompt, remove redundant material and check the token limit. Keep trusted instructions and the messages needed to interpret the request explicitly marked as required; if those alone exceed the limit, stop instead of silently trimming them.

Production accounting must include the model calls that propose steps and reconcile estimates with reported usage. Otherwise the displayed budget and the bill describe different workflows.

I'd store request-level usage keyed by run, user, and query, linked to the attempt events. Measure retrieved tokens, model tokens, latency, cost, and task success together. For this order lookup, use fixed test requests with expected users, orders, and items; verify both the returned result and rejection of unauthorized requests. A smaller bill with worse answers is not a win.

## Guidelines

1. **End every run with one recorded outcome:** `answered`, `needs-clarification`, `denied`, `stuck`, or `budget-exhausted`.
2. **Classify failures and enforce a retry threshold.** Test that changing error text doesn't let repeated schema guesses escape it.
3. **Reject unauthorized actions before execution.** Recheck task scope at the tool service, and never retry around a denial.
4. **Attach source and freshness fields to reused facts.** Refresh the relevant schema before restarting a failed lookup.
5. **Keep attempt history outside working state.** Preserve required instructions and request context when trimming model input.
6. **Enforce budgets and record the rejected step.** Measure actual usage and expected task results alongside cost.

I want the model to propose useful work, and I want code to decide when that work ends. Remembering failed attempts, resolving task authority, and enforcing budgets are how the same SQL agent avoids a loop, a permission leak, and an unbounded bill. Next in this context-engineering series: real context failures as case studies.
