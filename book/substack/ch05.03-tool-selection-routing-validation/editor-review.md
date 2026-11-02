## Editor Review — book/substack/ch05.03-tool-selection-routing-validation/draft.md (Post mode)

### Verdict
revise. The cohesion findings below are fixed in this draft. Overlap with an earlier post and two linking items remain.

Sources: `term_trace.py`, a top-to-bottom read, and a cold-read subagent that saw only the draft. Line numbers refer to the draft after fixes.

### Concept trace
| term | first use | explained at | status |
|---|---|---|---|
| transfer tool / payees | L9 | L9, L19 | fine |
| tool / tool call / harness / agent | L19 | L19 | fine. Four definitions in one paragraph is dense but readable |
| routing / side effect | L21 | L21 | fine |
| scope | L27 | L27 | **fixed**. It was used three times and never defined |
| freshness | L29 | L29 | **fixed**. It now has a one-clause example (a payee list loaded yesterday) |
| terminal states | L41 | L41 | fine |
| idempotency key | L55 | L55 | fine. How the key is derived is now explained (a request ID; next month's rent gets a new key) |
| `Ledger`, `s`, `p`, `key` in the excerpt | L59 | L57 | **fixed**. The text said "payment service" and the code said `Ledger`; they're now tied together |
| success path in the excerpt | L74 | L74 comment | **fixed**. A comment marks the fall-through to `completed` |
| `Notified` flag | L85 | L85 | **fixed** ("a flag" was vague) |
| compensation | L87 | L87 | fine |
| traces (close) | L100 | — | **fixed**. It now says "records like these outcomes and attempt counts" |

Introduction order: tool → tool call → harness → agent → routing → side effect → scope → terminal states → idempotency key → compensation.

### Findings (most severe first)
1. [argument] L39–49: terminal states and bounded attempts overlap ch01.02 ("Your Agent Doesn't Need a Better Prompt. It Needs a Stop Condition."). The new material here is idempotency, timeout-as-unknown, and per-side-effect retry. → Consider shortening the terminal-states section and linking ch01.02 for the stop-condition argument. → author / Astra
2. [evidence] The opening numbers ($1,200, Dana, two payees named Sam) are invented, and the draft says so at L13. The module cites Santos-Grueiro (2026) on context-to-execution integrity. The draft leaves it out on purpose because its manifest/sink vocabulary would need its own introduction. → Miriah decides whether to add one sentence and a link. → author
3. [cohesion] L15, L100: the bridge links the Substack profile, not the ch05.02 post. → Use the post URL once it's published. → author
4. [altitude] L59–74: the excerpt is exactly 15 lines, the cap. There's no backoff between retries; the cold reader noticed. The example omits it on purpose. → Add one clause ("add backoff in production"), or leave it. → Astra
5. [cohesion] L27: "extraction pipelines and ordinary workflow systems" is a passing mention with no example. → Cut it, or keep it as one clause. → Astra

### What's working
- The hook ("only the reply got lost") states the thesis in one sentence, and `TestRetryWithoutKeyChargesTwice` reproduces it.
- The diagram puts the notification after the record, with a crossed edge back to the transfer, so readers see the separation before they read about it.
- The excerpt matches `book/examples/ch05.03-tool-selection-routing-validation/workflow.go` L74–87, and `go test` passes.

### Next step
Miriah decides on Findings 1 and 2 in her notes.
