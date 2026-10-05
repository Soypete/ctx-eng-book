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
