## Editor Review — book/substack/ch00-what-we-mean-by-context-engineering/post.md (Post mode)

Scheduled: Wednesday 2026-10-07 (SCHEDULE.md status: `rewriting`). This is a re-review of the Astra rewrite 3, checked against the previous `editor-review.md` and the "Miriah's notes" sections in `review-ai.md`, `review-data.md`, and `review-fullstack.md`. The file is about 1,340 words including front matter, inside the 1,100–1,300 body target.

### Verdict
**revise**, but it's close. The rewrite fixed almost every finding from the last pass: agents and the harness are now set up first, "host" and "companion" are gone, the example is introduced before `Compile`, Authorize is the one name for the gate, and the broken book URL and typos are fixed. What's left is a smaller set of cohesion gaps:
1. Where Authorize actually sits reads as contradictory (L43, L49, L59, L67, L97).
2. "Authorize" gets a name before any pipeline exists (L29).
3. The "data contract" section still leans on terms it never explains (L85–89).

All three can be fixed in place in one more Astra pass, or by the author directly. The structure doesn't need to change.

### Status of earlier findings

**Previous editor-review.md (13 findings, 12 needing action): 11 resolved, 1 partly resolved.**

| # | Earlier finding | Status | Where |
|---|---|---|---|
| 1 | Typos, broken sentence, and broken book URL in L15/L21/L23 | **resolved** | L17 is clean and ends on "So what is context engineering?"; the L21 URL resolves; L23 says "lens", not "theory" |
| 2 | Agents and the harness showed up with no setup | **resolved** | L31 defines both before "Who owns selection?" (the bridging sentence is weak, see finding 7) |
| 3 | "host" and "harness" both in use | **resolved** | "harness" is used everywhere and "host" no longer appears |
| 4 | "The companion's `Compile`" was undefined | **resolved** | L57 introduces the runnable example with its link before `Compile` |
| 5 | Terms used before they're explained (seed manifest, adaptive retrieval, top-k, compile, language process) | **resolved** | "seed" and "adaptive" are gone; top-k is defined at L59, compiling at L35, and the manifest at L49 (before the diagram); "language process" is gone |
| 6 | No bridge from the lens to the rest of the post | **resolved** | L29: "This post works the Lexicon boundary in code; Semantics and Pragmatics get their own posts." |
| 7 | Hedges closing three sections | **resolved** | Two light hedges remain (L83 "not the whole retrieval system", L87 "version labels alone cannot do it"). Neither closes a section defensively, so they're fine |
| 8 | "Keep the data contract attached" was too dense | **partly resolved** | L87 now opens with why the manifest exists (the incident). L89 is still dense and uses undefined terms (finding 3) |
| 9 | The lens and the postures didn't tie back to the incident | **resolved** | L27 uses a discount the support bot could issue; L43 names the bot's posture |
| 10 | Four names for the gate; prompt injection not named | **resolved** | "Authorize" in the body, diagram, and guidelines; prompt injection is named and defined at L47 |
| 11 | No early series framing | **resolved** | L21 says this is the first post in a series drawn from the book |
| 12 | Code excerpt was 16 lines | **resolved** | L66–80 is 15 lines, with no packed statements |
| 13 | Evidence check | no action needed | The claims still match `book/examples/ch00-what-we-mean-by-context-engineering/compile.go` (post-filter, ID tie-break, budget exclusion, versioned `Manifest`) and its tests |

**Miriah's notes (review-*.md): all resolved.**
- **Lexicon framing of the opening failure:** resolved at L29, and authorization of data is called Lexicon throughout.
- **`Compile` as a post-filter:** resolved. The top-k starvation tradeoff is named at L59, and scope goes into the retrieval query for production. "Never reached the ranker" is gone. "Authorize before you rank" stays as Guideline 1.
- **Manifest version fields:** resolved. Source, policy, and ranker versions appear at L87 and in the example's `Manifest` struct.
- **Rewrite 2, voice:** resolved. The post is shorter and more direct, and the replay and measurement sections are each a single paragraph.
- **Rewrite 2, code:** resolved. The excerpt is 15 lines of the authorize-then-budget loop, with a link to the repo.
- **Rewrite 3:**
  - URL, typos, the question ending, and "lens": all resolved.
  - Agents and harness setup, "harness" everywhere, and "Authorize" as the only name: resolved.
  - The example is introduced first, compiling is defined, and "companion" is gone: resolved.
  - Findings 6 and 9 from the last review: resolved.
  - The early series sentence: resolved.

### Concept trace
Merged with a cold-read subagent that saw only the post text.

| term | first use | explained at | status |
|---|---|---|---|
| context engineering | title, L17 | L19 | ok |
| "authority" (subtitle) vs. "control signals" | L3, L19 | L19 (control signals only) | two names for one thing: the subtitle's "authority" never comes back, so connect them |
| control signals | L19 | L19 | ok (resolved since the last pass) |
| access scope / access metadata / data-access scope / scope | L9, L17, L29, L59 | never stated as one concept | **two names for one thing** (actually four). The cold reader couldn't tell whether they're one idea or three |
| Lexicon / Semantics / Pragmatics | L25–27 | L25–27 | ok |
| actor | L27 | implied at L61 | thin: say "actor (the authenticated requester)" the first time |
| Authorize | L29 ("the **Authorize** step") | L29 ("enforces those facts"); pipeline only at L51 | **used before explained**: it's called a "step" before any pipeline or diagram exists |
| Lexicon boundary / policy boundary | L29, L45 | never | borderline; readable from context |
| agent | L31 | L31 | ok (resolved) |
| harness | L31 | L31 | ok (resolved) |
| compiling context | L35 | L35 | ok (resolved) |
| candidates | L35 | L35 | ok |
| budget | L35 ("how much fits"), L39 | L63 ("context budget") | minor: "how much fits" covers it until then |
| prompt injection | L47 | L47 | ok (resolved) |
| manifest | L49 | L49 | ok (resolved: it's now defined before the diagram) |
| production target | L49 | never set against the example | feeds finding 1 |
| runnable example / `Compile` | L57 | L57 | ok (resolved) |
| working set | L57 | never | **two names for one thing**: candidates → eligible → working set → `out` → "Compiled context" (diagram) |
| top-k | L59 | L59 | ok (resolved) |
| row-level security | L59 | L59 | ok |
| `allowed`, `m`, `cands`, `remaining` | L61, L63 | L61, L63 | ok |
| "Recheck" (code comment) | L67 | never: the example has no earlier check | feeds finding 1 |
| "an oversized note" | L83 | never | minor: no note appears in the scenario |
| data contract (heading) | L85 | never | **never explained**: the heading names a concept the section never defines |
| replay | L87 | L87 (tied to the incident) | ok (improved) |
| ranker, policy version | L87 | never | minor: "ranking" appeared in the table; say "ranker version" plainly or accept it |
| data platform (team), lineage, indexed text fragments, index-age limits, index lag | L89 | never | **never explained**: a full-stack reader stalls here |
| `Doc.Version`, `Doc.AsOf` | L89 | never | **never explained**: the excerpt uses `d` and never shows a `Doc` type |
| embeddings | L89 | L89 | ok |
| authorized recall | L93 | L93 | ok |
| leakage | L93 | implied | ok for this audience |

Introduction order this piece needs (it's nearly there):
opening failure → why wording can't fix it → definition (including how "control signals" relates to the subtitle's "authority") → the lens, with the opening failure mapped to Lexicon and access scope named once → agents → harness → **Authorize introduced as the harness check that decides what text may be selected** → who owns selection, with compiling defined → manifest → diagram (the production target) → **the example as the recheck that runs after search inside the harness** → the post-filter tradeoff → **the data contract defined as what the data platform owes the harness** → evaluation → guidelines → next post.

### Findings (most severe first)
1. **[cohesion] L43, L49, L59, L67, L97: where Authorize sits reads as contradictory.** L43 puts Authorize "before candidate selection", the diagram (L49–53) shows it before the candidate set, Guideline 1 says "before you rank", and L59 says the example `Compile` filters "after search". The code comment at L67 says "Recheck access" even though the example has no earlier check. The cold reader couldn't work out where Authorize actually runs. Fix: one or two sentences at L57–59 that say it plainly. The diagram is the production target, where scope goes into the retrieval query. The example shows the second line of defense, the recheck inside `Compile` after search, which still matters when the query filter is missing or stale. Then the comment's "Recheck" makes sense. → **Astra**

2. **[cohesion] L29: "the Authorize step" is named before any pipeline exists.** The reader meets "step" with nothing to place it in until the diagram at L51. Fix: introduce it at L29 as a check rather than a step. Something like: "a check the harness runs before any text is selected; I'll call it Authorize." Or move the sentence after L31, so that the harness exists before it gets a check. → **Astra**

3. **[cohesion] L85–89: "data contract" is never explained, and L89 leans on undefined terms.** This is the remaining part of the earlier finding 8. The heading names a concept the body never defines. L89 assumes several things the reader hasn't been given:
   - a "data platform" team exists
   - lineage
   - "indexed text fragments" (chunks)
   - index-age limits
   - index lag
   - a `Doc` type that the excerpt never shows

   Fix: open L89 with one sentence saying what the contract is: what the data platform owes the harness (sensitivity tags, access metadata, versions, freshness), alongside the embeddings. Explain index lag in a clause (the gap between a change at the source and the change showing up in the index). Change `Doc.Version` and `Doc.AsOf` to "each document's `Version` and `AsOf` fields". Optionally move "measure storage and write latency under a retention policy" (L87) to the example README. → **Astra**

4. **[cohesion] L9, L17, L29, L59: four names for the requester's access scope.** The post says "access scope" (L9), "access metadata" (L17), "data-access scope" (L29), and "scope" (L59, L97). Fix: use "access scope" for the requester's permissions and "access metadata" only for tags on documents, and say that distinction once. → **Astra** (mechanical)

5. **[cohesion] L57: "working set" is a third name for the selected output.** The post goes candidates → eligible → working set → `out`, and the diagram ends at "Compiled context". Fix: say "selects which scored search results go into the model's input", or use "compiled context" to match the diagram. → **Astra** (mechanical)

6. **[argument] L3, L19: the subtitle promises "authority" and the body says "control signals".** The definition never connects the two, so the subtitle's claim reads as dropped. Fix: one clause at L19 tying control signals to authority (what the model is allowed to cause). The close at L104 already says "what the result was allowed to cause", so this would bookend the post. → **author**

7. **[flow] L31: "Agents bring these questions into a loop" is a vague bridge.** The cold reader asked which questions and what loop. The definitions that follow are good, but the opener isn't. Fix: open with the plain claim, for example that a lot of context now reaches models through agents. → **author / Astra**

8. **[flow] L21: the series sentence is followed by three disconnected thesis lines.** "Reliability depends on…", "More context is not…", and "The same discipline applies to…" each stand alone and read choppy, and the last one floats. Fix: fold them into one or two sentences, or move "the same discipline applies…" next to the definition at L19. → **author**

9. **[cohesion] L27, L83: small dangling references.** "Actor" (L27) is never tied to the rep or the authenticated requester, so add a few words. "An oversized note" (L83) refers to a note that isn't in the scenario, so say "an oversized document". → **Astra** (mechanical)

10. **[voice] L87: one long sentence of operational to-dos.** "Store a manifest reference…, protect the records…, and measure storage and write latency under a retention policy" is a checklist inside prose. It's minor and acceptable, but it's the last place the post sounds like a spec instead of like Miriah. → **Astra** (optional)

11. **[evidence]** No new claims. The opening scene is still framed as hypothetical ("Picture…", L43 "Our hypothetical support bot"), and the example claims match the code and tests. No action needed.

### What's working
- **The hook (L9–11)** is intact. "Not my incident-response plan" still lands.
- **L17's last line, "So what is context engineering?",** now leads cleanly into the bold definition.
- **L29** maps the incident to Lexicon and then says outright that this post works only the Lexicon boundary. That one line resolved the biggest argument gap from the last review.
- **L31** gives the agents and harness setup in two sentences, in the right place.
- **L43** places the incident in the posture table. The reader now knows which posture failed and why.
- **L47:** prompt injection is named, defined, and tied to a rule the harness enforces.
- **L59** names the post-filter tradeoff honestly, once, in plain words.
- **L87** opens with "After the opening incident, I'd want to know why the salary sheet reached the model". The manifest now has a reason to exist.
- **L93:** "An empty context can pass a leakage check and still fail the user." This is still the sharpest line in the post.
- **The close (L104)** bookends the thesis and bridges to ch01.01 without depending on it.

### Next step
Decide findings 6 and 8 yourself (subtitle tie-in, L21 thesis lines), then send findings 1–5, 7, and 9 to Astra as one short rewrite prompt. Finding 1 matters most: one sentence that separates the production target from the example's recheck. Re-run this editor pass and set SCHEDULE.md to `ready` only after no cohesion findings remain.
