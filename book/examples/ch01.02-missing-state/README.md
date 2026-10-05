# Harness-owned stop conditions (module 1.2)

Runnable companion to the Substack post *Your Agent Doesn't Need a Better Prompt. It Needs a Stop
Condition.* and [book module 1.2](../../chapters/ch01-every-failure-is-a-context-failure/modules/ch01.02-missing-state.md).

`Run` owns the agent loop. Before executing each model-proposed step it checks the task's granted
capabilities and the token budget; after each failure it compares the **error class** (for example
Postgres SQLSTATE `42703`, undefined column) rather than the message text, so a model guessing
`user_id`, `userid`, `userId`... ends as `stuck` instead of looping to the step limit. Every run
ends in one of five outcomes: `answered`, `needs-clarification`, `denied`, `stuck`,
`budget-exhausted`.

```bash
go test ./...
```
