# Full-stack engineer review — Your Agent Doesn't Need a Better Prompt. It Needs a Stop Condition.

**Would I finish it?** yes. The `user_id`/`userid` opener is a bug I've actually shipped, and the Go loop is the part I'd copy on Monday.
**Would I share it?** yes, with the teammate who owns our support-bot service and keeps adding "IMPORTANT: stop after 5 tries" to the system prompt.

## What lands
- "A sentence in the system prompt is guidance, not enforcement." That's the thesis, and I'm on board.
- The terminal-outcomes list plus `type Outcome string` gives me something concrete to put in a PR.
- The failure table (missing authz / over-scoped / execution-boundary) maps onto auth middleware vs. row-level filtering vs. downstream checks. I'd screenshot it.

## Where I got lost or rolled my eyes
- "it returns `stuck-same-error` on the second identical error." Your own opener says the agent "guesses a new column name, retries, gets a different error." Different guesses give different error strings (`column "userid" does not exist` vs `column "user_ID" does not exist`), so `err.Error() == last` never fires and the example just runs to `Exhausted`. Comparing raw error strings is also brittle. Real drivers put positions and request IDs in the message.
- The diagram says the harness "checks state, budget, and capability," but `Run` only counts steps. There's no token budget and no capability check, so the code doesn't do what the picture shows.
- Your outcome list includes "stops for missing evidence or permission," but the enum swaps that for `Stuck`, and the Guidelines list a third set again. Pick one set.
- "Some of that working state is a root set that must never be trimmed." What's a root set? It sounds like GC jargon, and no example follows.

## Missing for me
- Code for the permissions section: a `Principal{ActingFor, Resources, Actions, ExpiresAt}` struct checked in the tool executor, not in the prompt.
- One line of the token-budget check in the loop (`if tokensUsed > budget { return Exhausted }`).
- What the trace record looks like. You say "the trace shows the request that was rejected," so show me the log line.

## Top 3 edits (ranked)
1. Fix the stuck detector so it matches the story: normalize errors to a class (e.g. Postgres SQLSTATE `42703` undefined_column) and stop on a repeated *class*, or rewrite the opener so the errors really are identical.
2. Make the code match the diagram by adding a token/cost budget and a capability check to `Run`, or redraw the diagram.
3. Reconcile the three versions of the terminal-outcome list, and add a small `Principal` check snippet to the permissions section.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
