# Full-stack engineer review — Every Context Failure Has an Address

**Would I finish it?** skim. The sales-assistant opener and the failure table hooked me, but after the Go snippet the sections get shorter and vaguer, and I started skimming.
**Would I share it?** yes. I'd send it to the teammate who owns our RAG search box, with the freshness section highlighted.

## What lands
- The opener. "Nobody's prompt was wrong... What it was given was a day old." I've shipped that exact bug with a nightly sync job.
- The table mapping each failure to a boundary. It turns "the model got it wrong" into something I can file a ticket against.
- "A prompt can describe a rule. Only a boundary can enforce one." It's the same idea as auth middleware.
- The `Resolve` snippet, plus the refusal to fall back to the stale value. That's real code with a real decision in it.

## Where I got lost or rolled my eyes
- "The retrieval request carries a **capability**: a short-lived, server-issued grant..." Is that a JWT? A signed scope token? A row in Postgres? You named the most important fix in the post and never showed its shape.
- "a 2026 survey of execution-security work on coding agents... 69% to 98% against real denylists." Coding-agent denylists aren't my three-component reporting pipeline, and you admit as much in the next sentence. Either cut it or connect it.
- "Tell the model which signals count as supporting evidence." Which signals? For "why did service X fail," give me one concrete example, like "an error log line plus a deploy event in the same window."
- "return nothing rather than the old value." Fine for correctness, but what does my UI show the rep? The refetch also sits on the request's hot path, and you never mention its latency budget.

## Missing for me
- A code or JSON shape for a capability, and for a decision record. Is that a log line, an OTel span, or a table? I'd copy that schema on Monday.
- Real numbers for budgets, even illustrative ones: top-k 20, rerank to 5, 3k tokens, 800 ms.
- The diagram should show where these boundaries sit in a normal request path (API handler → retrieval → assembly → LLM call), not just list five boxes.

## Top 3 edits (ranked)
1. Add a capability snippet in about 10 lines of Go: the struct, how it's issued, and how the callee verifies it. It's the fix for two of the five failures, and right now it's only prose.
2. Show one decision record as JSON, and say where it lives (structured log or span attributes).
3. Cut or move the denylist statistic, and put a concrete "done" signal plus example budget numbers in its place.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
