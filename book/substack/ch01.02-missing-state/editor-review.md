## Editor Review — book/substack/ch01.02-missing-state/post.md (Post mode, after Astra rewrite)

### Verdict
revise

The rewrite fixed what mattered most in the last round. The opener and the code now describe the same failure, there is one set of five outcomes, harness/agent/trace/capability are defined before they're used, the middle sections tie back to the order-lookup agent, and the excerpt is short, accurate, and linked. What's left is smaller but still affects cohesion. The word **state** names three different things. **Boundary** carries the thesis but is never defined. The close points to "that diagnostic," which the post never set up. All of it can be fixed in place, and none of it needs restructuring.

### Comparison with the previous review (draft.md)

12 of 13 findings are resolved. One is partly resolved.

| # | Earlier finding | Status | Evidence in post.md |
|---|---|---|---|
| 1 | harness never defined; host/harness split across series | **resolved** | L17 defines it before first reliance (L23). ch01.01's post now also says "harness" (ch01.01 post L19) |
| 2 | code compared raw strings, so the opener ran to budget-exhausted | **resolved** | L53–55 stop on a repeated error class. L60 explains SQLSTATE 42703 and says message text isn't compared. `TestSchemaGuessingEndsStuck` passes |
| 3 | memory taxonomy interrupted the setup | **resolved** | moved after the code (L64–72) and tied to the retry counter (L66) |
| 4 | permissions and cost sections had no bridges | **resolved** | L76 ("that order-query agent") and L96 ("That same order lookup"). Retrieval is tied to the SQL queries at L78 |
| 5 | three outcome lists | **resolved** | one list at L27–31, matching the code constants, the diagram, and Guideline 1 |
| 6 | diagram promised pre-execution checks the code didn't have | **resolved** | capability defined at L33, before the diagram (L37). Excerpt L46–51 shows both checks before execution |
| 7 | working set / working state; Who / principal | **resolved** | only "working state" is used, and "who it acts for" (L86, L90) |
| 8 | trace never defined | **resolved** | L19 |
| 9 | excerpt was 30 lines, had no link, and didn't say who raises the sentinels | **resolved** | 11-line excerpt, link at L58, and L58 says who returns `ErrDenied` and `ErrAmbiguous` |
| 10 | no series marker, bridge, or close | **partly resolved** | series bridge at L15 and closing paragraph at L113. But L113's "that diagnostic" has no antecedent (new Finding 3) |
| 11 | Liu et al. cited for the impersonation history | **resolved** | the citation is attached only to the prompt-injection clause (L80), which the evidence ledger supports (`research/_evidence-ledger.md` L311–317) |
| 12 | parenthetical hedge; abstract root set | **resolved** | "Say it fetches..." (L96). The root set became "explicitly marked as required" (L98) |
| 13 | cost section not tied to the stop condition | **resolved** | heading L94 "A budget is another stop condition" |

**Miriah's notes** (same note in all three review-*.md files): **all honored.** The run stops on the error class, not the text. Capability and token budget are checked before execution. The same five outcomes appear in the prose, the code, the diagram, and the Guidelines. The excerpt is 15 lines or fewer and links to the repo.

**Code excerpt vs `book/examples/ch01.02-missing-state/run.go`: matches.** Post L46–51 are run.go L60–65 word for word, and post L53–55 are run.go L81–83. L52's "Later:" comment marks the skipped lines honestly: `spent += s.Tokens`, the exec call, and the class counter at L66–80. The surrounding claims also check out. `ErrAmbiguous` comes from the exec callback (run.go L73). Success returns `answered` (L71–72). Unclassified errors end `stuck` (L84–85). The tests cover all five outcomes, and `go test ./...` passes. The diagram PNG was regenerated after the .mmd change (16:42 vs 16:41) and matches the source.

### Concept trace
Sources: `term_trace.py`, a top-to-bottom read, and a cold-read subagent given only post.md.

| term | first use | explained at | status |
|---|---|---|---|
| agent | L9 | L9 | fine |
| schema context | L11 ("schema mismatch") | L13 | fine. The gloss follows within two lines |
| harness | L17 | L17 | fine |
| state | L17 ("a record of previous attempts") | L17 | **one name for three things**: L17 state = attempt record; L68 **working state** = what's assembled for the model call, kept *separate* from attempt history; L78 **trusted state** = session/policy store. The cold reader called this the biggest clash |
| constraints / enforcement | L17 | L17 | fine. "those limits" has no antecedent until "constraints" in the same sentence. Minor |
| trace | L19 | L19 | fine. The diagram's Trace node is a step in the loop, but L72 says the example doesn't persist a trace (minor) |
| boundary | L19 ("proof that the boundary worked") | never | **never explained.** Reused at L76 ("owns the boundary") and L92 ("pass through the boundary") as the thesis noun |
| transition | L23 | never | minor. Implies a state machine the post never shows |
| terminal outcomes | L25 | L25 | fine |
| budget / tokens | L31, L33 | L33 | fine |
| capability / permission / grant | L29, L33 | L33 | fine. L33 defines capability as a grant. Permission is used as the ordinary word |
| injected instruction | L35 | L35 | defined, but **premature**: nothing in the story has read untrusted text yet. The scenario that needs it arrives at L76–80 |
| error class | L39 (caption) | L60 | **used before explained**, 21 lines early |
| system prompt | L41 | — | common vocabulary. Fine |
| `Run`, callbacks | L43, L58, L100 | L58, partly | "execution callback" (L58) and "proposal callback" (L100) refer to `Run`'s two function parameters, which the excerpt never shows. Minor altitude |
| catalog / schema contract / data owner | L62 | never | minor. Readable as "the schema source," but "data owner" is a new role |
| working state | L68 | L68 | defined. See the **state** row |
| provenance | L68 | L68, by its fields | fine |
| platform team | L72 | never | minor. A new party with no setup |
| trusted state | L78 | L78 | defined, but see the **state** row |
| retrieval | L78 | L78 | fine. Tied to the SQL queries |
| prompt injection / confused deputy | L80 | L80 | fine, cited |
| tool service / protected service | L90, L92 | never | **two names for one thing**, and the first sign that execution happens in a service separate from the harness |
| capability token / downstream enforcement | L92 | never | minor. Named only to say the example lacks them |
| assembly | L98 | never | minor. L98 assumes retrieval → assembly → send stages that were never laid out |
| "that diagnostic" | L113 | never | **never explained.** The post never calls anything a diagnostic |

Introduction order this piece needs:
agent → the schema loop (opener) → series bridge → harness, with state as *attempt history* → boundary (the line the harness enforces between proposal and execution) → trace → five terminal outcomes → budget / capability → diagram → code (error class defined here or before the caption) → working state vs attempt history → untrusted text and injected instructions (with the credential scenario) → trusted state → tool service as the place the effect happens → budget as a stop condition across retrieval → assembly → send → Guidelines → close.

### Findings (most severe first)
1. [cohesion] L17, L68, L78: "State" names three things. At L17 it's "a record of previous attempts." At L68 attempt history is kept *separate from* working state. At L78 "trusted state" is the session or policy store. Guideline 5 (L110) only works if the reader follows the L68 split. → Give L17's concept a distinct name, attempt history, and use it again at L66–68. Keep "working state" for the assembled model input. Either rename "trusted state" (for example, "trusted sources") or add one clause saying it's a different kind. → Astra / author
2. [cohesion] L19, L76, L92: "The boundary" carries the thesis ("owns the boundary," "every execution path must pass through the boundary") but is never defined. L19 asks for "proof that the boundary worked" before any boundary exists. → At or before L19, one plain clause: the boundary is the check between what the model proposes and what actually runs. → Astra
3. [cohesion] L113: "applying that diagnostic to concrete failures, with documented incidents clearly separated from teaching examples." The post never calls anything a diagnostic, so the antecedent is either L15's previous-post questions or nothing. The second half reads like the author's editorial policy leaking to readers. → Name what the next post does in reader terms ("Next: real incidents where these stop conditions were missing"), and drop the process note. → author / Astra
4. [flow] L35: The injected-instruction definition arrives before the story has any untrusted text. It belongs with the credential scenario at L76–80, where the post already defines prompt injection. Defining it twice is also repetition. → Move L35 into the permission section, or cut it and keep only L80. If the diagram's "Injected instruction" node stays, say in the caption that it's covered below. → Astra
5. [cohesion] L39: The caption relies on "error class," which isn't explained until L60. → Gloss it in the caption ("an error class, such as a Postgres SQLSTATE code") or move the L60 explanation ahead of the excerpt. → Astra
6. [cohesion] L90, L92: "The tool service" and "the protected service" are two names for one thing, and they're the first sign that the tool runs in a separate service the harness calls. → One sentence at L90 saying the tool call lands in a service that performs the effect, then use one name. → Astra
7. [voice] L17–35: Ten terms are defined in about 18 lines: harness, state, constraints, enforcement, trace, terminal outcomes, budget, tokens, capability, injected instruction. It reads like a spec, and the cold reader found it dense. → Keep the L17 harness sentence. Let capability and budget arrive in the outcome list where they're used. Move injected instruction out (Finding 4). → author
8. [voice] L62, L72, L92, L100: Four "the example doesn't/isn't..." caveats in four sections. Each one is honest, but together they make a hedging stack. → Keep one consolidated scope note right after the code link (L58), and cut the rest down to what the reader must do in production. → author
9. [altitude] L58, L100: "execution callback" and "proposal callback" refer to `Run`'s parameters, which the excerpt doesn't show. → Add one clause at L43 or L58: `Run` takes a function that proposes the next step and one that executes it. → Astra
10. [flow] L80–88: The text sets up prompt injection and confused deputy, then says "These failures need different fixes," but the table lists three other failures. → Bridge with one clause saying the table sorts *where* the permission failure happens, or relabel the transition. → Astra
11. [cohesion, minor] L37 / diagram, L72: The diagram shows the trace as part of the loop, but L72 says the example doesn't persist one. "Answer ready" leaves the harness, but in the code `answered` comes from a successful execution. → Acceptable as a picture of the full application. Optionally caption it "in the application around the example." → author
12. [cohesion, minor] L62, L72, L98: "data owner," "platform team," and "assembly" arrive without setup. → One clause each, or cut "platform team." → Astra

### What's working
- The hook (L9–11) is intact and now honored by the code. "A very expensive spelling bee" is Miriah's voice.
- L23 and L41: "The harness owns the transition" and "A sentence in the system prompt is guidance, not enforcement." This is still the thesis, and it's in the right place.
- The excerpt (L46–55) is accurate, short, and commented for *why* ("never retry around a permission decision").
- L60–62: The honest note that a repeated-class threshold is a policy, not proof, and that the stop condition "doesn't repair the schema contract." It's good argument hygiene.
- L70: "Otherwise, yesterday's schema becomes today's confident mistake."
- L76 and L96: The bridges keep one running example across all three failures, and the close (L113) restates that.
- L102: "A smaller bill with worse answers is not a win."
- The Guidelines (L106–111) are imperative and testable, and each one traces to a section.

### Next step
Send Findings 1–6 to the next Astra pass, or have the author fix them: one name per kind of state, a one-clause definition of the boundary at L19, a reader-facing next-post pointer at L113, and the injected-instruction definition moved into the permission section. Findings 7–8 are author voice edits.
