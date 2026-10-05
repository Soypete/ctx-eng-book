# AI engineer review — Before You Blame the Model, Check Three Gates

**Would I finish it?** yes. It's short, the gate framing is useful, and the Go snippet makes it concrete.
**Would I share it?** yes, with product engineers on my team who say "the model hallucinated" in every incident channel. I wouldn't send it to other eval people yet.

## What lands
- "a bug report about the system around the model, filed under the wrong name" is the right thesis, and it's well put.
- The table is good. "A bigger prompt does nothing" for source failures is advice I'd give.
- Treating tool calls with the same gates, including the "fifty tools, calls fifteen" overload case.

## Where I got lost or rolled my eyes
- "For every answer, keep four things together: ... the evidence a correct answer needs (from your eval set)." Production answers don't have an eval set. `Expected` only exists for labeled questions, so `Gate()` is an offline eval classifier, not a debugger for live traffic. Say so, or show how a production failure becomes a labeled case.
- `Unsupported []string // claims in the answer with no retrieved evidence behind them`. This field is doing all the hard work and gets no explanation. Claim extraction plus attribution is an LLM-judge or NLI step with its own error rate. That's the hand-waving about evals I worry about most.
- The code doesn't match the prose. The post says stale, contradictory, or "badly ranked evidence" cause hallucinations, but if retrieval admits both `policy-v3` and `policy-v1` and the model cites v1, `Unsupported` is empty and `Gate` returns `"ok"`. The overload case ("too much of it") has no gate either.
- Citing Vaswani to argue that attention "doesn't verify truth," and then saying it's "my inference, not a claim the paper makes," uses a citation as decoration. Cut it or replace it with an empirical source on unsupported generation.
- "megabytes of output" is hyperbole. Measure in tokens.
- "Lexicon / Semantics / Pragmatics" shows up with "In the language of this series" and is never defined for a cold reader.

## Missing for me
- An acknowledgment that this lines up with existing RAG eval metrics (context recall for retrieval, faithfulness/groundedness for generation). Without that, it reads as relabeled old ideas. Then say what the post adds: the source gate and the identity/scope check.
- One line on how `InSource` and `Unsupported` actually get populated.
- A fourth outcome for "evidence present but conflicting or buried," or say outright that `Gate` doesn't catch it.

## Top 3 edits (ranked)
1. Fix the `Gate` gap: handle contradictory or stale retrieved evidence (or distractor overload), or state that it's out of scope. Right now the code says "ok" for the post's own refund example if v1 and v3 are both retrieved.
2. Call it an eval-time tool, and explain how `Unsupported` is computed (a judge model, and that the judge needs its own validation).
3. Name the overlap with RAGAS-style recall and faithfulness metrics, and drop or replace the Vaswani paragraph.

## Miriah's notes

- **Stale copy = Source problem (provenance).** When retrieval returns an outdated copy (policy-v1 instead of v3), the failure is at the Source gate: the source layer didn't carry provenance/version/authority, so the stale copy looked legitimate. Make the table and the policy-v1/v3 example agree on that, and replace `InSource bool` with provenance fields (version, as-of, authoritative) so the code can express it.

- **Code:** the fixed example (provenance fields; stale copy = source failure; runs on labeled traces) lives at https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.01-missing-information with tests. Show only the key lines (≤15) and link to it; don't paste the program.

<!-- Add your notes for the rewrite here. Astra reads this section. -->
