## Editor Review — book/substack/ch05.02-tool-schemas-and-function-calling/draft.md (Post mode)

### Verdict
revise. The cohesion findings below are fixed in this draft. One argument question and one linking item remain for the author.

Sources: `term_trace.py`, a top-to-bottom read, and a cold-read subagent that saw only the draft. Line numbers refer to the draft after fixes.

### Concept trace
| term | first use | explained at | status |
|---|---|---|---|
| `create_customer`, `get_customer` | L9 | L9, by name and context | fine |
| allowed list / enum | L11 | L25 | **fixed**. It now reads "one of the allowed plan names" |
| harness | L15 | L21 | **fixed**. The bridge said "the harness" before L21; it now says "your application code" |
| tool / schema / tool call | L19 | L19 | fine |
| agent | L21 | L21 | fine |
| selection | L36 | L38 | fine |
| action contract | L52 | L52 | fine |
| scope / evidence the call needs | L52 | L52 | **fixed**. Both were undefined; each now has an inline gloss |
| Kodathala, gateway | L60 | L60 | **fixed**. "Identity-bound ... gateway" read as jargon; it now says "a separate service in front of the tools" |
| eval | L64 | L64 | **fixed**. It was used as a noun with no definition |
| semantic | L69 | L33 ("questions about meaning") | fine |
| `k`, `c`, `p` in the excerpt | L76 | L74 | **fixed** |
| selector | L94 | L94 | **fixed**. It was introduced only implicitly |
| "different owners" | L105 | — | **fixed**. It now names where each kind of failure gets fixed |

Introduction order: tool → schema → tool call → harness → agent → what a schema can't say → description as selection context → action contract → eval → five classes → code.

### Findings (most severe first)
1. [argument] L25 vs L69: the schema can declare `format: email`, but the post files "not an address" under semantic failures. The cold reader asked which stage owns it. **Fixed in place** with a parenthetical: many validators skip `format`, so the example checks it in code. Miriah may prefer a cleaner semantic example, such as a customer who already exists, and drop the email case from the list.
2. [argument] L92–96: the hook's request ("Move Ana to pro") is blocked twice, by the semantic duplicate check and by the abstention eval. The cold reader saw this as a contradiction. **Fixed in place**: the draft now says the semantic check blocks the call but can't reveal an agent that reaches for tools it shouldn't. This is the post's subtlest point, so watch it in the rewrite.
3. [cohesion] L15, L109: the bridge links the Substack profile, not the previous post (ch05.01). → Use the post URL once it's published. → author
4. [evidence] L60: the Kodathala numbers (15 models, eight attack scenarios) appear in the module and in `research/_evidence-ledger.md` L385–388. They're supported. Scope is limited to that paper's gateway, and the draft says so.
5. [altitude] L78–88: the enterprise-tier rule is hard-coded after the generic role check. The note after the excerpt now says it belongs in policy data in a real system.
6. [voice] L42–44: the `get_user` poor/good description pair comes from the book module, but the hook's tools are `create_customer` and `get_customer`. Consider rewriting the pair around `get_customer` so the post uses one running example. → Astra

### What's working
- The hook is a schema-valid failure that every persona will recognize.
- "Exclusions, not purposes, separate look-alike tools" is the post's most quotable idea, and it's new relative to the earlier posts in the series.
- The eager-selector test is a strong, concrete demonstration of why abstention cases matter. The excerpt matches `harness.go` L43–55, and `go test` passes.

### Next step
Miriah decides on Finding 1 (keep or drop the email example) and Finding 6 in her notes.
