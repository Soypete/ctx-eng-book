# Every Failure Is a Context Failure

## Argument (one sentence)
Every AI system failure—hallucinations, wrong tool calls, loops, permission errors—is ultimately a failure of context: missing, incorrect, inaccessible, or unconstrained information.

## Why it matters (reliability goal)
The Pragmatics pillar's goal is aligning model behavior with user intent. Without proper context, intent cannot be inferred. The Governance pillar requires controlled information access. This module establishes the foundational thesis that AI reliability is context engineering.

## Beats (the path to the conclusion)

1. **The Demo-to-Production Gap** — Most AI demos work because they run in controlled, context-rich environments; they fail in production where context is partial, stale, or absent. — evidence: GAP (needs case study)

2. **Hallucinations Are Missing Context** — When models lack sufficient grounding information, they generate plausible but false outputs. The failure is not the model's creativity but the system's failure to provide necessary reference material. — evidence: GAP

3. **Wrong Tool Calls Are Missing Context** — Agents call wrong tools when they don't have accurate descriptions, state, or user intent. Tool selection requires contextual understanding of what's available and what the user wants. — evidence: GAP

4. **Agent Loops Are Missing Exit Criteria** — Loops occur when the model lacks terminal context—what success looks like, when to stop, or what information would break the cycle. — evidence: GAP (could cite infinite loops case study)

5. **Permission Failures Are Unconstrained Context** — Giving agents excessive access leads to failures: they attempt operations they shouldn't, retrieve data they shouldn't see, or act beyond authorization. — evidence: GAP

6. **Cost Overruns Are Unmeasured Context** — Without context budget controls, systems consume unlimited tokens, make redundant calls, or retrieve unnecessary data. — evidence: GAP

7. **The Pattern: All Fail Trace to Context** — Synthesize: every failure mode maps to one of four context problems: missing, incorrect, inaccessible, or unconstrained. — evidence: GAP (could link to five-pillars framework)

## Conclusion to reach
The reader should believe that AI failures are not mysterious or inherent to the technology—they are engineering problems with engineering solutions: proper context construction, validation, and governance.

## Examples needed
- Side-by-side comparison: working demo context vs. failing production context
- Code: simple context audit that flags missing pieces before a call

## Readiness
needs-research ({list: case studies for hallucinations, tool calls, loops, permissions, cost; evidence for context-as-root-cause})