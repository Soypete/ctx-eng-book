## Editor Review — book/substack/ch05.01-toolformer-and-react/draft.md (Post mode)

### Verdict
revise. The cohesion findings below are fixed in this draft. What's left is evidence and linking, and those are the author's call.

Sources: `term_trace.py`, a top-to-bottom read, and a cold-read subagent that saw only the draft. Line numbers refer to the draft after fixes.

### Concept trace
| term | first use | explained at | status |
|---|---|---|---|
| tool | L9 ("weather tool") | L19 | fine. Ordinary use in the hook; defined before any reliance |
| `delete_account` | L9 | L9 | **fixed**. The cold reader asked why a weather app has this tool; L9 now says it's for the settings page |
| schema / JSON Schema | L19 | L19 | fine |
| tool call / proposal | L19, L21 | L19, L21 | **fixed**. The bridge (L13) used "proposal" before L21; it now says "a request that code has to check" |
| harness | L23 | L23 | fine. Same term as earlier posts in the series |
| observation | L23 | L23 | fine |
| grants | L27 caption | L23 | **fixed**. It was used in the caption before any definition; now defined in the harness paragraph |
| Toolformer / fine-tuned | L33 | L33 | fine |
| ReAct | L43 | L43 | fine |
| agent | L45 | L45 | fine |
| read/write tools | L47 | — | **fixed**. The split was never set up; rephrased as "every tool result your agent reads ... every other tool it can call" |
| scripted model | L61 | L61 | **fixed**. It was explained only at L81 |
| `h`, `p`, `args` in the excerpt | L63 | L63 comment | **fixed**. A signature comment was added; the excerpt is 12 lines |
| registered vs granted | L79 | L79 | fine |

Introduction order: tool → schema → tool call/proposal → harness → observation → grants → Toolformer → ReAct → agent → the boundary in code.

### Findings (most severe first)
1. [evidence] L43: ReAct gets one sentence ("interleaves reasoning with actions; results come back as observations"). The module cites Yao et al. (2022) but never develops ReAct beyond the citation, and there is no evidence-ledger row for it. The sentence restates the paper's framing and adds no number. → Keep it as is, or add a ledger row so the description is anchored. → author / research
2. [evidence] L33: "kept the calls that improved the model's predictions." This matches the module's wording. The cold reader asked "improved how?" The paper's criterion is whether the call's result lowered the loss on the following tokens, but the module doesn't say that. → Leave it vague, or have research confirm the criterion before the post adds detail. → author / research
3. [cohesion] L13, L95: the bridge and the next-post pointer link to the Substack profile, not to the previous post (ch04.03, "Valid JSON Is Still Only a Proposal"). → Link that post's URL once it's published. → author
4. [cohesion] L84 (now fixed): the prose said a rejection "goes back as an observation," but the excerpt returns an error. It now says the loop records the rejection in the history the model sees next, which matches `loop.go`.
5. [altitude] L51: Google is named, but only OpenAI's and Anthropic's request shapes are described, and the example renders only those two. This is minor. → Say "the example shows two of them," or drop Google from the shape comparison. → author
6. [argument] Overlap with the series: the step limit and "the harness decides" repeat ch01.02 and ch04.03. What's new here is Toolformer's split between selection and authority, and "registered vs granted." Keep the emphasis there. → Astra

### What's working
- The hook is concrete and the code test reproduces it: the injected `delete_account` gets `ErrNotGranted`, and only `get_weather` runs.
- "A model that learned to ask well has still only learned to ask" carries the argument in one line.
- The excerpt matches `book/examples/ch05.01-toolformer-and-react/tool.go` L79–99 (trimmed), and `go test` passes.

### Next step
Miriah decides on Findings 1–3 in her notes. No further cohesion work is needed before the persona reviews.
