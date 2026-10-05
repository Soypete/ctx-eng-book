## Editor Review — book/substack/ch01.01-missing-information/post.md (Post mode, after the Astra rewrite)

Scheduled: Friday 2026-10-09 (SCHEDULE.md status: `rewriting`). The body is about 1,530 words including the code, inside the 900–1,600 range. Inputs for this pass: `post.md`, the previous `editor-review.md` (on `draft.md`), "Miriah's notes" in `review-fullstack.md`, `review-ai.md`, and `review-data.md`, `book/examples/ch01.01-missing-information/` (`gate.go`, `gate_test.go`, README; `go test ./...` passes), `diagram-1.mmd`, `research/_evidence-ledger.md`, and `term_trace.py` output. A **cold-read subagent** read the post text only, and its list is merged below.

### Verdict
**revise** (light)

The rewrite fixes the problem that blocked the last pass. A stale copy is now a Source/provenance failure in every place it comes up. The code excerpt is accurate, short, and linked. Nearly every earlier finding is resolved. What remains is small, but some of it is cohesion: one overloaded word, one gate definition that doesn't cover the case the post leans on, and two terms the reader needs that are never explained. You can fix all of it in place. Nothing has to be reordered.

### Checks the author asked for

**Code excerpt vs `gate.go`: matches.** Post L62–67 is identical to `gate.go` L18–23 (the `Evidence` struct). Post L69–75 is identical to `gate.go` L46–48 plus L52–55 (the "why" comment and the stale-copy branch). It leaves out the missing-provenance check at L49–51 on purpose, and L78 points to it ("also check missing provenance"). The excerpt is 14 lines, under the 15-line limit. L78's test claim is accurate: `gate_test.go` L33–34 return `source: stale copy` both for v1 alone and for v3 plus v1.

**Stale copy = Source/provenance failure, everywhere: yes.** It appears at L23 (authoritative and provenance defined first), L29 (the Source gate asks "authoritative, current"), L35 (the caption: freshness is checked before retrieval), L39 (the table: "stale copy presented as legitimate" falls under Source), L45 (`policy-v1`/`policy-v3` is a Source failure, and ranking can't fix it), L69–75 (the code comment and return), L78 (the tests), L88 (a stale tool result fails Source), L102 (Guideline 1), and L109 ("a stale policy needs a data repair"). `diagram-1.mmd` agrees: the `S` node asks "current and versioned?", `SF` says "add provenance," and `focus` is on `S`. Nothing contradicts it. One definitional gap remains (Finding 2).

### Earlier findings: status

| # | earlier finding | status |
|---|---|---|
| 1 | Code and worked example call stale a retrieval failure | **resolved** (L45, L69–78) |
| 2 | Code block 29 lines, stale version, no link | **resolved** (14 lines, linked at L78) |
| 3 | `Gate` only runs on labeled traces; `Unsupported` unexplained | **resolved** (L53, L57) |
| 4 | Series unnamed; Lexicon/Semantics/Pragmatics orphaned | **resolved** (series named and linked at L15; terms tied to gates at L29–31) |
| 5 | Ten-file opener never sorted into gates | **resolved** (L43, using Miriah's classification: German = Source) |
| 6 | Vaswani citation is decoration; ledger labels an inference as a Quote | **resolved in the post** (cut). The **ledger entry is still open**: `research/_evidence-ledger.md` L10 still records the inference as a "Quote." That's for research, not the post. |
| 7 | Tools section doesn't map to the gates; overload has no gate | **resolved** (L88; overload = Retrieval/selection; says `Gate` misses it) |
| 8 | "megabytes" unsupported number | **resolved** |
| 9 | working context / identity-scope / agent / host undefined | **resolved** (L27, L30, L86, L19; "harness" used throughout, "host" gone) |
| 10 | Admit the overlap with RAG metrics, add a ledger row with a URL | **partial**: the overlap is named at L47, but there's no citation in the post and no ledger row |
| 11 | Stable evidence IDs and trace governance | **resolved** (L80) |
| 12 | No closing takeaway; clunky next-post line | **resolved** (L109) |
| 13 | Guidelines lack the versioning rule and "labeled" | **resolved** (L102–103) |
| 14 | Diagram encodes the old model | **resolved** (`diagram-1.mmd`) |
| 15 | Hedge stacks at L86/L102 | **resolved** for those lines. A new caveat pattern appears (Finding 7). |

**13 of 15 fully resolved**, 2 partial (6: ledger only; 10: citation).

Miriah's notes are all applied: provenance fields replace `InSource`, the excerpt is ≤15 lines plus a link, L/S/P are kept and tied to the gates, the opener is sorted, overload counts as Retrieval/selection, agent and harness each get one sentence, "harness" is the only name, the post-2 bridge is in, labeled traces are stated, the versioning Guideline is in, "megabytes" is gone, and the diagram is updated.

### Concept trace
| term | first use | explained at | status |
|---|---|---|---|
| harness | L19 | L19 | OK |
| contract | L11 (the legal document), L21 (the system contract), L43 (the "contract appendix") | L21 | **two meanings for one word** (Finding 1) |
| Lexicon / Semantics / Pragmatics | L21 | L21, tied to gates at L29–31 | OK. Retrieval maps to none of them, which the cold reader found muddled because "identities" sits under Lexicon at L21 (minor) |
| authoritative / provenance / supersede | L23 | L23 | OK. "Source owner" has no obvious meaning for a user-uploaded PDF (minor, Finding 6) |
| working context | L27 | L27 | OK |
| Source / Retrieval / Generation | L29–31 | L29–31 | OK, but Source is defined over the source, not the admitted copy (Finding 2) |
| OCR | L43 | L43 | OK |
| "the source layer" | L45 | never | minor: a new name next to "Source gate" |
| context recall / faithfulness | L47 | L47, in one clause each | OK in wording, but uncited (Finding 4) |
| trace / eval set | L51 | L51 | OK. "eval cases" at L53 is close enough |
| `Gate`, labeled traces | L53 | L53 | OK. The link only comes at L78 (minor) |
| `Needs` / `Current` / `Retrieved` / `Unsupported` | L55–57 | glossed in parentheses | **fields of a type the reader never sees** (Finding 3) |
| attribution check / attribution labels | L57, L106 | never | **never explained** (Finding 5) |
| chunk / chunk hash | L80 | never | minor. Most of this audience knows chunking |
| agent | L86 | L86 | OK |
| task boundary | L88, L105 | implied by L31 ("stay within the task") | minor |
| MCP | L90 | L90 | OK |
| hallucination | L13 | L94 | OK. The everyday meaning carries L13, and L94 sharpens it on purpose |
| workflow / state | L109 | forward reference | OK |

Introduction order this piece needs (it now follows this order): demo vs. production → series bridge → harness → system contract (L/S/P) → authoritative and provenance → working context → three gates → sort the ten files → stale copy = Source → trace and eval set → labeled traces → `Trace` fields → `Unsupported` and claim checking → excerpt and link → limits → agent and tools → hallucination → skeptic → guidelines → takeaway and next post.

### Findings (most severe first)

1. **[cohesion] L21 (also L11, L43): "contract" means two things.** L11 and L43 use it for the legal document with the appendix. L21 uses it for the harness's input agreement ("lacks a contract: which files exist..."). The cold reader tripped on this, and it sits right where L/S/P are introduced. → Give the L21 sense a different noun (for example "spec" or "input agreement") or qualify it once ("a system contract, not the legal one"), so "contract" means the uploaded document everywhere else. → author / Astra.

2. **[cohesion] L29 vs L45: the Source gate is defined over the source, but the post's key case is about an admitted copy.** L29 asks whether "the needed fact is available, authoritative, current." L30 then defines Retrieval as admitting evidence. At L45, `policy-v1` *was admitted*. So the cold reader asked whether that makes it a Retrieval issue, and why Source, the first gate, inspects retrieved evidence. L45's "ranking cannot reliably choose" answers the why, but the gate definition doesn't make room for it. → Extend L29 by one clause: Source also asks whether every admitted copy carries provenance from the source showing it is current. Then L45 and the code (which checks `r` against `cur`) follow from the definition. Optionally change "the source layer" at L45 to "the source" to drop a third name. → Astra.

3. **[altitude] L55–57, L82, L103: `Needs`, `Current`, `Retrieved`, `Unsupported`, and the `ok` return belong to a `Trace` type and a `Gate` signature the reader never sees.** The excerpt shows only `Evidence`. The cold reader couldn't tell the field types, or whether `ok` is a bool or a string. → Add one sentence before the excerpt, without adding code lines: "`Gate(t Trace) string` takes a trace with four fields... and returns the first failed gate as a message, or `"ok"`." That keeps the excerpt at 14 lines. → Astra.

4. **[evidence] L47: context recall and faithfulness are named with no source.** This is the rest of earlier Finding 10. → Link the RAGAS paper (https://arxiv.org/abs/2309.15217) or the RAGAS metrics docs, and add a row to `research/_evidence-ledger.md`. → research, then Astra.

5. **[cohesion] L57, L106: "attribution" is never explained.** L57 says "an automated attribution check," and Guideline 5 says "review attribution labels." Neither tells the reader that attribution means tying each answer claim to the evidence record that supports it. L57 nearly says this already ("match each against the admitted records"). → Name it there ("that matching is attribution") so Guideline 5 has an antecedent. → Astra.

6. **[argument] L23: "the source owner vouches" doesn't cover the opener's case.** In the ten-file scenario, the user uploaded the files. Who vouches for authority or supersession? → One clause: for user uploads, the harness either asks the user or records the upload as the as-of version. Or leave L23 general and say which party plays owner in the upload case at L43. → author.

7. **[voice] L53, L57, L82, L88, L106: the same caveat repeated across sections.** The disclaimers about what `Gate` can't do now appear five times: "does not diagnose unlabeled live traffic... not proof of success," "validate that checker... outside this example," a full paragraph of limits (L82), "the small classifier above doesn't catch that overload," and "beyond the example's checks." L82 is the right home for them. → Keep L82, fold the overload limit into it, and cut the others to the minimum the sentence needs. Also, L13 "This *may* be a bug report" softens the thesis line the reviewers quoted, and L98 "now you can demonstrate it" is flatter than the earlier "prove it." The author should decide whether to restore both. → author / Astra.

8. **[flow] L88: "the useful function definition gets buried" is unclear.** The overload example used to be about an answer buried in tool output. "Function definition" reads as tool schemas, which is a different failure. → Say what got buried: the one useful result among many tool outputs. → Astra.

9. **[flow] L53: "An empty list of required facts is not proof of success" dangles.** Nothing tells the reader that `Gate` returns `ok` when `Needs` is empty. → Fold it into the Finding 3 sentence ("with no `Needs`, it returns `ok`, so an unlabeled trace proves nothing"). → Astra.

10. **[minor] L9–11: "ten files," but the list adds up to seven.** Three scans, two German files, one spreadsheet, and one contract. → Say "the rest are fine," or change the count. → author.

11. **[altitude, code only] `gate.go` L52–54 (post L72–74): misleading message wording.** For the v1-alone case and for a non-authoritative record, the message says "competes with," but nothing competes. This is optional. It's the example file, not the post, and changing it means re-checking the excerpt and test prefixes (the tests match on the `source: stale copy` prefix, so the tail can change). → author.

Not acted on from the cold read: the front matter (Substack strips it); "hallucination" defined at L94 (deliberate, as in the last pass); the series link points to the profile rather than post 1 (fine until post 1 has a URL, then link it).

Outside the post: `research/_evidence-ledger.md` L10 still labels the Vaswani inference as a **Quote** (earlier Finding 6). → research.

### What's working
- The provenance thread is fully consistent across prose, table, diagram, code, tests, guidelines, and takeaway. That was the main ask, and it landed.
- L45: "Ranking cannot reliably choose the current policy when the records don't say which version supersedes which." That one line justifies the decision.
- L43 sorts the opener's failures into gates and pays off the hook.
- The excerpt is accurate, 14 lines, has its "why" comment and an explicit loop-context sentence (L59), and links to tests that cover the post's own example.
- L47, "Faithfulness to an obsolete policy still produces the wrong answer," answers the AI reviewer's "relabeled RAGAS" objection in one sentence.
- L109's takeaway ("I want the incident ticket to name a repair") and the bridge to the ch01.02 stop-condition post.
- The skeptic section (L96–98) is still Miriah's voice.

### Next step
Make the small in-place fixes for Findings 1–3, 5, 8, and 9 (author edit or a narrow Astra pass). Add the RAGAS citation and ledger row (Finding 4). Miriah decides Findings 6, 7, and 10. Then mark the post ready for 2026-10-09. A full re-review isn't needed unless the gate definitions change.
