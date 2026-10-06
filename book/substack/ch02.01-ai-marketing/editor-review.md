## Editor Review — book/substack/ch02.01-ai-marketing/draft.md (Post mode)

### Verdict
ship (for persona review). The first-pass cohesion findings are fixed. One open evidence decision remains, and one internal-reference leak was removed.

Sources: `term_trace.py`, a top-to-bottom read, and a cold-read subagent that saw only the draft (first version). Line numbers refer to the revised draft.

### First-pass findings, now fixed in the draft
| # | Finding (first draft) | Fix |
|---|---|---|
| 1 | "context" and "context engineering" were never defined, and the bridge listed four failures under "the last three posts" | L13 defines both and says "the last few posts", listing only failures posts 3–5 covered |
| 2 | Lexicon/Semantics/Pragmatics appeared with no reason for the names | L17 says they're borrowed from linguistics. Each lens is paired with its plain question (missing information / unclear meaning / wrong use) |
| 3 | the three-part frame had four different wordings | one wording throughout: missing information, unclear meaning, wrong use (L19–21, L76, Guideline 1) |
| 4 | the caption depended on the color orange, and the alt text didn't mention it | the alt text now says "the orange model box". The diagram re-rendered with the model in brand orange |
| 5 | "sev2" vs "P0" mixed two scales | sev2 is now the undefined-meaning case and P0 the unrecognized-label case (L46–47) |
| 6 | "governed catalog", `s`, `m`, the `Model` interface, `TriageQueue`, and "stand-in models" were unexplained | L49 defines the catalog by its fields. L51 states the interface and what `s` holds. L68 defines `TriageQueue` and says "fake models" |
| 7 | "This isn't only about agents" pushed back on a premise the post never set up | L41 now sets it up from the series' earlier agent posts |
| 8 | the prompt section came out of nowhere, and the list of five investments wasn't mapped | L72 opens from where teams spend effort. L74 maps the investments onto the three lenses |
| 9 | the close said "layers" after the post said the lenses aren't layers | the close now says "the parts of a production context stack" |
| 10 | catalog / reference data, and label / category | "label" throughout. "Reference data" is used only in Guideline 2, which names the catalog |

**Internal reference removed.** The module cites the curriculum's `02_Prompt_Patterns` and `03_Agents_101` READMEs. Readers can't see those, so they're left out of the post.

### Concept trace (revised draft)
| term | first use | explained at | status |
|---|---|---|---|
| context / context engineering | L13 | L13 | fine |
| Lexicon / Semantics / Pragmatics | L19–21 | L17–21 | fine |
| agent | L13 ("agent loops") | L41 | used before explained. It appears only as a modifier in the bridge, so this is minor |
| postmortem / on-call | L9, L46 | — | common engineering vocabulary. Fine |
| taxonomy | L49 | L49 (by example) | fine |
| `Model` / `s` / `m` / `TriageQueue` | L51, L68 | L51, L68 | fine |

Introduction order: "the AI failed" → context → three lenses → pipeline diagram → models matter but can't carry the system → classifier (no agent) → code → prompt as one interface → close.

### Findings (most severe first)
1. [evidence] the module's readiness is "ready-with-open-evidence": its primary systems sources are queued and not yet reviewed. The post makes no cited claims. The model-quality caveat (L33) and the database analogy (L37) are worded as the author's argument, which the outline asks for ("do not make an unsupported commoditization claim"). → none now; research later
2. [evidence] L72: "prompts that were tuned for one model often need rework for the next" is the author's experience, not a cited finding. Keep it in first-person framing, or cite a source if one is in the queue. → author
3. [argument] the module is short (~650 words), so this post adds its own scaffolding: the three questions, the database comparison, and the classifier code. All of it follows the module's argument, and nothing new is claimed as evidence. Miriah should confirm the "Ban 'the AI failed'" stance (Guideline 1) is one she'll stand behind. → author
4. [altitude] the excerpt shows the product check, not the severity check. L68 says severity works "the same way" and the link has the full code. Minor. → none

### What's working
- The hook is a line every engineer has heard in a postmortem, and the post answers it with a method.
- The ticket classifier demonstrates "no agent required" in code. Two fake models run through the same checks.
- The close points to post 7 (the production context stack) without depending on it.

### Next step
Persona reviews. Then Miriah decides on Finding 2 and on the strength of Guideline 1.
