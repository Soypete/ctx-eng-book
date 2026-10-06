## Editor Review — book/substack/ch01.04-personalization-failures/draft.md (Post mode)

### Verdict
ship (for persona review). First-pass cohesion findings are fixed. Open evidence gaps carried over from the module are listed below.

Sources: `term_trace.py`, a top-to-bottom read, and a cold-read subagent that saw only the draft (first version). Line numbers refer to the revised draft.

### First-pass findings, now fixed in the draft
| # | Finding (first draft) | Fix |
|---|---|---|
| 1 | the bridge used "boundary", which this post never defines | L13 now says "the exact place in the code where each one got through" |
| 2 | "provenance" appeared in the diagram alt text before it was defined | named at L34, in Failure 1. Diagram moved to the enforcement section (L56) |
| 3 | "tool argument", "tenant", "task policy", "sign it", "row-level security" were unexplained | glossed at L40 (tool argument, tenant), L60 (server-side task policy), L79 (signing, Postgres RLS) |
| 4 | "billing email" first appeared in the test summary | introduced at L50 (a code review doesn't need your billing email) |
| 5 | "can't be forged" was followed by "a forged capability reads nothing" | now "can't hand-build a working one" and "a hand-built or expired capability reads nothing" |
| 6 | "preference version" and "policy decision" were never set up | replaced with "source and timestamp" and "capability issued" |
| 7 | the four questions didn't map to the four failures | the questions now include "as of now" (staleness) and "for that task" (scope). L28 sets up four common failures instead of a one-to-one mapping |
| 8 | "policy" meant both a preference source and access rules | preference source renamed "organization rule". "Policy" now means only the server's task-to-fields mapping |
| 9 | the store had six names | "preference store" throughout. "Store" appears only in code |

### Concept trace (revised draft)
| term | first use | explained at | status |
|---|---|---|---|
| context / retrieval | L17 | L17 | fine |
| personalization | L19 | L19 | fine |
| provenance / declared / inferred | L34 | L34 | fine |
| tool argument | L40 | L40 | fine |
| tenant isolation | L40 | L40 | fine |
| prompt injection | L50 | L50 | fine |
| least privilege / capability | L52 | L52 | fine |
| task policy | L60 | L60 | fine |
| row-level security | L79 | L79 | fine |
| Dana | L9 | invented name | fine. Invented scenario, no real data |

Introduction order: context → retrieval → personalization (not memory) → four questions → provenance → tenant isolation → freshness → least privilege/capability → enforcement code → measurement.

### Findings (most severe first)
1. [evidence] the module has an open `[RESEARCH NEEDED]` for a primary source on preference elicitation and personalization failure. The post makes no research claim about it and frames all four failures as engineering patterns, so nothing in it is unsupported. But it carries no citation at all. If Miriah wants one, the ledger's Shift (2024) RAG-security entry could support "retrieval widens the data-leakage surface". Other sources are still queued per the module outline. → research / author
2. [evidence] L46 and L50: "a month" vs "an hour", and "a thousand documents", are illustrative values from the module, worded as hypotheticals. They are not measurements. Keep the hedged wording ("probably", "Say"). → none
3. [argument] L81: the example is in-memory. The post says production enforcement belongs in the database (L79), which matches the module's caveat. A reviewer may still ask for an RLS/SQL version. That is a decision for Miriah, not a fix. → author
4. [flow] L83–93: the "Measure it like a system" list repeats some of Guideline 6. Minor. Astra can merge them if length is needed. → Astra

### What's working
- The Dana hook makes the most serious failure land in two sentences.
- "A framework might call your preference store 'memory.' It's still a table." This is the post's thesis in Miriah's voice.
- The code shows the missing `user_id` parameter, and the post points that out. It closes the module's `[CODE EXAMPLE NEEDED]` with tests that prove cross-user reads fail.

### Next step
Persona reviews. Then Miriah decides whether to add a citation (Finding 1).
