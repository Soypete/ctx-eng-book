# AI engineer review — Stop Saying "The AI Failed"

**Would I finish it?** yes. It's short, the Go snippet is real, and "make the model a parameter" is how I already build things.
**Would I share it?** yes, with PMs and incident-review leads who still write "the model hallucinated" as a root cause. I wouldn't send it to my ML team as-is, because it skips evals.

## What lands
- Model `m` as an interface, with checks owned by system `s`. Two fake models run through the same tests is the right demo.
- The catalog carries an owner, a version, and an as-of date, and `CatalogVersion` is stamped on each result. That gives you provenance you can debug from.
- Unknown labels go to triage with a reason instead of being silently dropped. Good.

## Where I got lost or rolled my eyes
- "each named with a term borrowed from linguistics": I'd call these data quality, schema/ontology, and output validation plus authz. The renaming makes it look like a new idea when it isn't one. "Pragmatics = wrong use" also stretches what the linguistics term means.
- "Nothing stops the classifier from emitting a label the routing rules don't recognize, like 'P0'": structured outputs and enum-constrained decoding stop this at generation time. Not mentioning them makes the post look like it predates 2024.
- "A better model fixes none of those." That overclaims. A stronger model follows sent definitions better and abstains more often. Say "a better model can't *guarantee*" any of those.
- "and check whether the outcome was right": this line is the whole problem, and it's never shown. The code only catches *out-of-vocabulary* labels.
- "Every one was fixed outside the model": I'd want a link or a number backing that.

## Missing for me
- **The failure mode I actually see:** a valid, in-catalog, wrong label. "Billing / sev3" when it's "Auth / sev1" passes every check in this snippet. The fix is a labeled eval set and per-model confusion matrices, run on every model swap. Show that.
- The stale-catalog guard sends *every* ticket to triage. What's the blast radius when the catalog job breaks on a Friday? Add alerting, or a degraded mode.
- One line on structured-output schemas as the first line of defense, with code validation as the backstop.

## Top 3 edits (ranked)
1. Add the "valid but wrong" case and a minimal eval: a golden set, accuracy per label, and a test run against both fake models. Without it, "the system decides whether the feature works" is hand-waving.
2. Soften "A better model fixes none of those" to "can't guarantee," and name constrained decoding or enum schemas next to the routing validation.
3. Own the relabeling. Add one sentence saying the lenses map onto data quality, schema, and guardrails that engineers already know, and explain what the new names buy you.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
