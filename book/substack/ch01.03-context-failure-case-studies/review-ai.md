# AI engineer review — Every Context Failure Has an Address

**Would I finish it?** skim. The table and the freshness code are worth my time, but I skimmed the definitions of tokens, RAG, and reranking.
**Would I share it?** yes, with the platform/backend engineers on my team who still think "fix the prompt" is the default incident response. I wouldn't send it to other agent builders.

## What lands
- "A prompt can describe a rule. Only a boundary can enforce one." That's the thesis, and it's correct.
- The failure → boundary table. It's the most reusable thing in the post, and I'd steal it for a postmortem template.
- "If the source can't answer, return nothing rather than the old value, because a stale answer looks exactly like a fresh one." Good line, and the Go snippet backs it up.
- "A cheaper path that loses the right answer isn't a fix." Few cost posts bother to say this.

## Where I got lost or rolled my eyes
- "a 2026 survey ... found failure rates from 69% to 98% against real denylists." Those studies are about coding agents and command denylists. The section is about a confused deputy between services. You hedge it ("not a measurement of this scenario"), but if it needs that hedge, it isn't supporting evidence. Cut it or find a source on delegated authorization.
- "the caller passes its capability with every call, and the callee ... verify it." This is the hand-wavy part. How? OAuth token exchange (RFC 8693)? Macaroons? Downscoped JWTs? Without a mechanism it reads as least privilege under a new name. The confused deputy dates to 1988, so tell me what's specific to agents.
- "Tell the model which signals count as supporting evidence." Which signals, for a debugging agent? This is the one failure with no concrete fix.
- "Run the same task repeatedly and record every step." That's not an eval. How many runs, what pass criterion, and what metric tells me a boundary regressed?

## Missing for me
- An example decision record, shown as JSON, for one boundary. You list the fields but never show one.
- The MCP and tool-description injection path gets half a clause in the last section. For me that's the failure mode I haven't seen handled, and it deserves its own row in the table.
- The cost of failing closed on freshness. When does "return nothing" hurt availability, and what do you show the user?
- The snippet ends after `Refetched` with no return, so it reads as truncated.

## Top 3 edits (ranked)
1. Make "capability" concrete: name a mechanism and show the token or struct that gets passed and checked. Right now it carries two sections on a definition alone.
2. Replace the Rashidi citation with evidence about delegated authorization, or drop it.
3. Show one real decision record, and give the "test it like any regression" paragraph a pass/fail criterion (N runs, a metric, and what counts as a regression).

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
