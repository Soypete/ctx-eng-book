# Full-stack engineer review — Stop Saying "The AI Failed"

**Would I finish it?** yes. It's short, and once the ticket classifier and the Go snippet show up it reads like a real PR and stops reading like a lecture.
**Would I share it?** yes, with my team lead and whoever writes our postmortems, plus the teammate who keeps proposing "just upgrade to the newer model."

## What lands
- The classifier example. "emitting a label the routing rules don't recognize, like 'P0,' so the ticket lands nowhere" is a bug I've shipped, with strings instead of labels.
- The Go snippet. A model behind an interface, a staleness check before the call, and an allowlist check after it is what I can copy on Monday.
- "Swapping the model changed the labels it proposed. It didn't change the checks they had to pass." I'd quote that one.
- The Postgres analogy works for me. It's exactly how I think about it.

## Where I got lost or rolled my eyes
- "each named with a term borrowed from linguistics" — Lexicon / Semantics / Pragmatics is the part I'd skim past. The plain-English names in parentheses ("missing information / unclear meaning / wrong use") are what I'll actually remember. As written, the jargon comes first and the useful words come second.
- "These aren't layers. They're three lenses you can point at any part of the system." The next paragraph says "Put together, the lenses describe a pipeline." So which is it, lenses or a pipeline? I lost the thread there.
- "Authorization, operational feedback, and evaluation hold who may do what" — permissions get mentioned here and in the Pragmatics bullet, but the example never touches auth. The earlier post's cross-user data leak is the case I care about most, and this post doesn't come back to it.
- "When that happens, you lose a little." Lose a little what? This sentence is vague.

## Missing for me
- **Where this sits in a request.** A sequence diagram (HTTP handler → load catalog → model call → validate → route or triage) with latency marked would help. Does the staleness check hit the DB on every request, or is it cached?
- **What the prompt actually contains.** `s.prompt(ticket)` hides the part the "send definitions with labels" guideline is about. Show 5 lines of the rendered prompt with the taxonomy inlined.
- **Structured output.** Should I use JSON schema / tool-call output mode, and then still validate? Say so.
- **Triage operations.** Who watches TriageQueue, and what metric tells me the rate is climbing?

## Top 3 edits (ranked)
1. Show the rendered prompt next to the Go snippet so guideline #3 has concrete code. Right now the most actionable guideline has no code behind it.
2. Lead with the plain names (missing info / unclear meaning / wrong use) and move the linguistics terms into parentheses. Fix the "not layers… a pipeline" contradiction.
3. Add one Pragmatics line about auth (for example, a per-tenant catalog keyed on the session's org ID) so the permissions claim lands in code.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
