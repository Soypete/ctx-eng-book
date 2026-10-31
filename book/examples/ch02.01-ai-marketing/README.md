# The model is a parameter; the system is the work (module 2.1)

Runnable companion to the Substack post *Stop Saying "The AI Failed"* and
[book module 2.1](../../chapters/ch02-ai-is-a-systems-problem/modules/ch02.01-ai-marketing.md).

A support-ticket classifier assigns a product and a severity. No agent, no tool loop.
`System.Route` takes the model as an argument and checks its answer against three things
the model doesn't own:

- **Lexicon:** a versioned product catalog with an as-of time and max age. A stale catalog
  stops the request before the model runs.
- **Semantics:** a severity taxonomy with written definitions, sent to the model with the
  labels.
- **Pragmatics:** routing rules. A label with no queue, or a product that's retired or
  invented, goes to `human-triage` with a reason instead of being routed on the model's word.

The tests run two stand-in models through the same system. Swapping the model changes the
labels it proposes, not the checks they must pass.

```bash
go test ./...
```

## Validity is not accuracy

`TestGoldenSetSeparatesValidityFromCorrectness` runs a tiny invented golden set
through a keyword fake and a constant-answer fake. Both pass output validation;
the constant-answer fake misroutes a checkout outage as `search / sev3`. Run
`go test -v ./...` to see exact matches by expected label pair and the confusion
matrix (expected -> predicted counts). These fixtures demonstrate the evaluation
boundary, not real-model performance. Evaluate actual model adapters against a
representative, reviewed dataset before replacing a production model.

Failure returns preserve `CatalogVersion`, including stale snapshots and model
errors. The tests also assert that an expired snapshot never reaches the model.

## Deliberate limits

This example loads data in memory and supplies typed prompt fields on each call.
It has no live provider, output-schema adapter, authentication, organization
filter, snapshot publisher, cache invalidation, alerting, or staffed triage queue.
`AsOf` must be set from source freshness, not refreshed by copying an old artifact.
The maximum age is a fixture choice, not a production recommendation. A retirement
inside that age requires snapshot invalidation; membership alone cannot detect it.
Production records also need catalog ownership and source lineage, taxonomy and
routing-policy versions, and access and retention rules. Keep historical labels
attached to their original definitions when those definitions change.
