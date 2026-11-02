# AI engineer review — Five Tests Every Tool Needs Before an Agent Can Call It

**Would I finish it?** yes. It's short, the Ana hook is a bug I've actually seen, and the stage-labeled rejections are worth stealing.
**Would I share it?** yes, with the junior engineers on my team who think strict mode means the call is safe. Not with my peers until the eval section holds up better.

## What lands
- "If the only check in front of your tool is the schema, the only bugs you'll catch are typos." That's the right line.
- The eager-selector test. "If your eval has no no-tool cases, an agent that always reaches for a tool looks perfect" is the most useful sentence here.
- "p comes from the session, never from Args." Correct, and most demos get it wrong.

## Where I got lost or rolled my eyes
- "Do not use when the person already has a record' is the sentence that would have saved Ana." It wouldn't. The model can't know Ana exists unless it calls `get_customer` first. The real fix is "ask, or say there's no plan-change tool." The draft's own semantic check is what actually saves Ana.
- The ordering in the Go block: `checkMeaning` (duplicate lookup) runs **before** `hasRole`. An unauthorized caller can tell "already exists" apart from "unauthorized," so the check leaks whether a customer exists. Authorization should run before any check that reads data. That's the failure mode the post doesn't show.
- "To evaluate your own descriptions, swap in a real model call and keep the cases." This is where the eval is hand-waved. With replayed choices, the selection cases are unit tests of the harness, not evals of the model. A real selection eval needs repeated samples, pass rates per class, and a run per model version.
- The Kodathala citation says "the direction matches" but gives no result. What did it find? If it can't be stated in one clause, cut it.
- "Action contract" is preconditions/effects plus a policy engine, which is design-by-contract and OPA. Fine, but say so instead of presenting it as a new idea.

## Missing for me
- A **clarify** outcome next to abstention. For Ana, asking a follow-up question is the right behavior, and it needs its own eval class.
- Idempotency is listed and then never tested.
- MCP: when descriptions come from a third-party server, the description is an attack surface (tool poisoning), and the "model sees the description" model breaks.

## Top 3 edits (ranked)
1. Move authorization ahead of the semantic/duplicate check in the code, and add one sentence on why (it stops the existence leak).
2. Rework the Ana fix: the description can't encode state the model doesn't have, so make the right behavior ask or abstain, and add "clarify" as an eval class.
3. Say what a real selection eval needs (multiple samples, pass rate per class, rerun on model change), and state Kodathala's actual finding or drop it.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
