# Full-stack engineer review — Five Tests Every Tool Needs Before an Agent Can Call It

**Would I finish it?** yes — the duplicate-Ana opener is a bug I've basically shipped, and the Go snippet looks like middleware I already write.
**Would I share it?** yes — with the teammate who owns our support-bot tool handlers.

## What lands
- "If the only check in front of your tool is the schema, the only bugs you'll catch are typos." That's the line for the PR description.
- The check chain reads like a request pipeline: validate, business rules, authz, execute. I know where that goes in my service.
- "p comes from the session, never from Args" — exactly how I treat a user ID in a request body. Good.
- The eager-selector test is the most useful idea in the post. I'd steal it today.

## Where I got lost or rolled my eyes
- The code runs `h.checkMeaning` (a duplicate lookup, so a DB hit) *before* `hasRole`. In my API, authz comes before anything that touches data. Otherwise an unauthorized caller can learn "this email already exists" from the error, and I pay for a query on calls I'm going to reject anyway. Either swap the order or say why meaning comes first.
- "whether a repeated request should create a second record or return the first." Idempotency gets listed and then dropped. Retries are what actually cause duplicates in prod, and the example never shows an idempotency key.
- The Kodathala paragraph. I skimmed it. It doesn't change what I build.
- "Three words before we start" then defines five terms.

## Missing for me
- The real fix for Ana. There's no `change_plan` tool, so the right answer is abstain *and then add the tool*. The post never says that, so it reads like the model is the only one at fault.
- What the harness sends back to the model when a call is rejected. "Return errors in one consistent format" — show me the payload. Does `ErrSemantic` go back so the model can retry, or to the user?
- Where this lives: one wrapper around every tool handler, or per tool? And what does it cost in latency per turn?

## Top 3 edits (ranked)
1. Put authorization before the semantic check in the snippet and the diagram, or explain the order. As written it's a pattern I'd flag in code review.
2. Add a short "what the model gets back" example for each rejection stage.
3. Close the Ana loop: abstention plus "add the missing tool," and one line on idempotency keys for retries.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
