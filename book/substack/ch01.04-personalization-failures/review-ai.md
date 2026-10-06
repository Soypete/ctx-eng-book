# AI engineer review — Your Model Doesn't Know Your Users. Your Retrieval Does.

**Would I finish it?** skim. The Dana opener and the "`user_id` never gets a vote" line hooked me, but the middle is authz basics I already know.
**Would I share it?** yes. I'd send it to product engineers who are new to agents and think "memory" is magic. I wouldn't send it to my agent team.

## What lands
- "A framework might call your preference store 'memory.' It's still a table." This is the right deflation, and it's said plainly.
- Flagging a model-filled tool argument as an identity source is the real agent-era bug. I've seen MCP tools ship with `user_id` in the schema.
- Provenance (declared vs. inferred vs. imposed) is the most useful idea here. I don't see it written down often.
- The `Read` signature with no `user_id` parameter is a good teaching artifact.

## Where I got lost or rolled my eyes
- "Personalization fails in four ways, and every one of them is a context-assembly bug." Three of the four are authz and cache invalidation with new labels. Least privilege, scoped capabilities (these are OAuth scopes), and RLS are decades old. Say that out loud, then show what's *new* about doing it when an LLM sits downstream.
- "its fields are private to the package, so other code can't hand-build a working one." Unexported Go fields are an API boundary, not a security one. You do add "sign it" later, but this sentence oversells.
- "The model's learned behavior is ordinarily the same between requests." That hedge is clumsy. Just say "weights don't change per request."
- "Correction rate" and "Task relevance" are listed as metrics with no detail on how to measure them. How do I detect a correction? Who labels which fields a task "needed"? This is the eval hand-waving I watch for.

## Missing for me
- **The failure modes I'd actually hit.** Vector search that runs top-k first and filters by tenant afterward, so it leaks or returns nothing. Prompt or KV cache prefixes shared across users. Agent-written memories, where an injected instruction becomes an "inferred preference" that persists. That last one is a fifth failure, and it's the scary one.
- Where inference happens. Who writes inferred preferences, and with what confidence or decay?
- One eval harness snippet: a cross-tenant red-team test that drives the *model* to ask for Dana's data through a tool call, not just a unit test on `Store.Read`.

## Top 3 edits (ranked)
1. Add a fifth failure, memory poisoning through agent-written inferred preferences. Or at least a paragraph on it. It's the part that's genuinely LLM-specific, and it ties provenance to security.
2. Admit the lineage (least privilege, OAuth scopes, RLS) in one sentence, and pivot to what the LLM changes: the model fills tool arguments, and injected text reaches retrieval.
3. Define how to measure correction rate and task relevance, or cut them. Also fix the "private to the package" claim.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
