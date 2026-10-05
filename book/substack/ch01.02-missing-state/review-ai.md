# AI engineer review: Your Agent Doesn't Need a Better Prompt. It Needs a Stop Condition.

**Would I finish it?** yes. It's short and the harness-owns-the-transition point is right, but the code doesn't catch the failure the post opens with.
**Would I share it?** yes, with teammates who still put "stop after 10 tool calls" in the system prompt and treat that as enforcement.

## What lands
- "The model can suggest the next step. The harness owns the transition, and it records why." This is the thesis, and it's correct.
- "Never retry around a permission denial" is a rule I'd put in a code review checklist.
- The three-row failure table (missing authz / over-scoped / execution-boundary) is the most useful thing here. Most people lump those together.

## Where I got lost or rolled my eyes
- "it returns `stuck-same-error` on the second identical error." But your own opening says the agent "guesses a new column name, retries, gets a different error." `column "userid" does not exist` and `column "user_id" does not exist` are different strings, so `err.Error() == last` never fires and you fall through to `Exhausted`. The real failure is *near*-duplicate errors, and the post skips it. You need error fingerprinting (normalize identifiers, compare error class) or a no-progress check, not string equality.
- The prose lists five outcomes, including "stops for missing evidence or permission." The code has `Stuck` instead and no missing-evidence case. Pick one list.
- "Episodic / semantic / working state" is the CoALA split (Sumers et al., 2023) with new labels, and there's no citation. Credit it, or say what's different. "Root set" also needs one sentence on how you mark it.
- The Liu et al. citation is a prompt-injection benchmark. "Act for a person it doesn't represent" is the confused-deputy problem, so call it that.

## Missing for me
- A token/cost budget inside `Run`. Right now it only counts steps, but the cost section says "budget at every stage."
- What "task success per request" means here: a labeled eval set, an LLM judge, or user outcomes? "Measure whether it worked" is hand-waving.
- A short example of enforcing at the execution boundary, e.g. a capability token checked by the tool server.

## Top 3 edits (ranked)
1. Fix stuck detection so it catches the opening example: fingerprint errors or track no progress, and show `userid`/`user_id` actually ending in `stuck`.
2. Make the outcome list match between prose, code, and Guidelines, and add a token budget to `Run`.
3. Cite CoALA for the memory split and say "confused deputy" for the impersonation point.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
