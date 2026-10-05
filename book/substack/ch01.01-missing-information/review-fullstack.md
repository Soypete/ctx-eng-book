# Full-stack engineer review — Before You Blame the Model, Check Three Gates

**Would I finish it?** yes — the opening is my actual Q&A feature, and the gate table is the triage I wish I'd had last quarter.
**Would I share it?** yes — with the team that owns our "ask your docs" endpoint, as the thing to read before the next "the model is dumb" Slack thread.

## What lands
- "Most of the time it's a bug report about the system around the model, filed under the wrong name." Great line.
- The gate table. Source / retrieval / generation, each with a different owner. I can put that in a runbook today.
- The `policy-v3` vs `policy-v1` example is concrete and I've seen it happen.

## Where I got lost or rolled my eyes
- "the evidence a correct answer needs (from your eval set)". This is where the Monday plan falls apart. `Gate()` needs `Expected`, and a live user question in prod has no expected evidence IDs. So this works on my eval set, not on the angry support ticket. The post never says that, so I'd ship the struct, log `Expected: nil`, and every answer would come back `"ok"`.
- `Unsupported []string // claims in the answer with no retrieved evidence behind them`. How do I fill that in? Claim extraction plus an entailment check plus an LLM judge? That's the hard part, and it's hidden in a field comment.
- "the **Lexicon**... the **Semantics**... the **Pragmatics**". New series jargon right when I'm trying to get to the code. I don't need three linguistics terms to know I'm missing a file manifest.
- The ten-file opening never comes back. Are the scanned PDFs a source failure or a retrieval failure? The German files? Run the gates on the story you started with.
- I skimmed the Vaswani paragraph. "That's my inference from the architecture" is honest, but it slows down a post that's otherwise practical.

## Missing for me
- What a trace row looks like in prod: which columns, where it's logged (one table? OpenTelemetry span attributes?), and how much it costs to store retrieved chunks for each request.
- A sentence on the prod path: no `Expected`, so you replay the failing question against the eval harness, or have a human label it, and *then* run `Gate`.
- One concrete way to compute `Unsupported`, even a rough one.

## Top 3 edits (ranked)
1. Say plainly that `Gate()` runs on eval or labeled traces, not live traffic, and show how a prod failure gets into that set.
2. Go back to the ten-file example and sort each failure (scans, German, spreadsheet, appendix) into a gate.
3. Show how `Unsupported` gets populated, or drop it from the struct and say generation checks are a separate eval.

## Miriah's notes

- **Stale copy = Source problem (provenance).** When retrieval returns an outdated copy (policy-v1 instead of v3), the failure is at the Source gate: the source layer didn't carry provenance/version/authority, so the stale copy looked legitimate. Make the table and the policy-v1/v3 example agree on that, and replace `InSource bool` with provenance fields (version, as-of, authoritative) so the code can express it.

- **Code:** the fixed example (provenance fields; stale copy = source failure; runs on labeled traces) lives at https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.01-missing-information with tests. Show only the key lines (≤15) and link to it; don't paste the program.

<!-- Add your notes for the rewrite here. Astra reads this section. -->
