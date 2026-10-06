## Editor Review — book/substack/ch01.03-context-failure-case-studies/draft.md (Post mode)

### Verdict
ship (for persona review). The cohesion findings from the first pass are fixed in the draft. What's left is minor, plus one open decision for Miriah.

Sources: `term_trace.py`, a top-to-bottom read, and a cold-read subagent that saw only the draft. The cold read ran on the first version. The table below describes the revised draft. Line numbers refer to the revised draft.

### First-pass findings, now fixed in the draft
| # | Finding (first draft) | Fix |
|---|---|---|
| 1 | "boundary" was defined as a crossing point (retrieval, assembly, tool call, handoff), but the table called "workflow state" and "freshness" boundaries | L19 now defines a boundary as any check where data, authority, time, or cost crosses, and lists all five |
| 2 | "Overhydrated", "terminal state", "query expansion", "candidates", "rerank", "tokens", "prompt injection", "trajectory", "fail closed", "denylist" were never explained | renamed to "Over-fetched" and "definition of done". Glossed inline at L47, L53, L55, L83. Dropped "trajectory" and "fail closed" |
| 3 | the scoped grant had three names: "server-issued, task-scoped access", "capability", "grant" | **capability** is defined at L49 and used from there on. "Grant" appears only inside the definition and in Guideline 5 |
| 4 | the diagram sat inside the stale-data section but covered all five failures | moved to sit right after the table (L31). Node labels renamed to match the table |
| 5 | "from the example" came before the link, and `d`/`maxAge` were unexplained | L61 names the companion Go example and explains `maxAge`, `fetch`, and `d` before the excerpt |
| 6 | confused deputy was defined as "tricked", and the components weren't set up | L81 says the components can be agents, services, or steps, and the definition no longer says "tricked" |
| 7 | the Rashidi citation came with no setup | L83 frames it as related research on coding agents, defines denylist, and keeps the module's caveat |
| 8 | failure names drifted between the table, the headings, and the evidence list | table, headings, and evidence-list labels now match |

### Concept trace (revised draft)
| term | first use | explained at | status |
|---|---|---|---|
| context | L17 | L17 | fine |
| agent / harness | L13 ("an agent stops") | L17 | used four lines early, in the bridge sentence. Minor |
| boundary | L19 | L19 | fine |
| tokens | L47 | L47 | fine |
| prompt injection | L47 | L47 | fine |
| capability | L49 | L49 | fine |
| RAG / query expansion / rerank | L53–55 | L53–55 | fine |
| decision record | L31 (alt text) | L61, L87 | the diagram caption uses it before the prose explains it, but it reads clearly in context. Minor |
| confused deputy | L81 | L81 | fine |
| denylist | L83 | L83 | fine |

Introduction order: context → agent/harness → boundary → five failures (table + diagram) → capability → RAG → freshness → decision record → confused deputy → evidence per boundary.

### Findings (most severe first)
1. [argument] L59–75: the code covers only the freshness boundary, but the post covers five, so the example backs one fifth of the claims. This is intentional because a five-boundary program would bury the point, but Miriah should confirm it. → author
2. [evidence] L83: the Rashidi (2026) range comes from the module and the evidence ledger (`research/_evidence-ledger.md`, the Rashidi entry). The post keeps the module's caveat. The source is about coding-agent execution security, not inter-service delegation, so it is only analogous. A source on confused-deputy or delegation failures would fit better. → research
3. [altitude] L3 and L41: the module's two scenarios are both constructed, and the post says so at L35. Nothing quantitative is invented, and "a pile of searches" replaces the module's "hundreds". → none
4. [flow] L13: the bridge says "an agent stops" before L17 defines agent. Readers who skipped post 3 still follow it, but the definitions could move above the bridge. → author/Astra
5. [voice] length is 1,561 words, near the 1,600 cap. If Astra adds anything, cut the debugging-loop section (L37–41) first, because post 3 covered loops. → Astra

### What's working
- The hook (closed deal, stale index) is concrete and matches the code test exactly.
- The table plus the diagram give the post one organizing idea: a failure has an address.
- The "A fix you can't observe is a hypothesis" section ties the five cases to the module's conclusion, which is to name the telemetry.

### Next step
Persona reviews. Then Miriah decides on Finding 1, keeping a single-boundary example or asking for a second example.
