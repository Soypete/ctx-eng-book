# Attention Cannot Choose Your Sources

Run with Go 1.22 or newer, from this directory:

```sh
go run .
go test ./...
go vet ./...
```

Standard library only; no model, API key, or network call. The invented runbook
fixture includes stale, wrong-service, unauthorized, and current evidence. The
program prints the exact JSON input for a toy model interface plus an ordered
decision manifest. Only the current, authorized checkout runbook enters the input.

Tests assert admitted content, exclusion of rejected text, provenance records,
freshness boundaries, and source removal (ablation). Removing a rejected source
leaves the input unchanged; removing the only admitted source returns no payload.
This proves application control flow, not model accuracy or attention behavior.

The one-hour freshness limit is an illustrative fixture policy, not a recommendation.
`ConfirmedAt` means when the source vouched for the content, not index refresh time.
`Authorized` is a trusted adapter result for the caller, not a field supplied by a
user or document. Production source services must enforce access before fetching
text. This example assumes truthful adapter metadata and does not resolve conflicts
between multiple admissible records.

For a real model, adapt the payload to its message format, capture that final input
in protected storage, and count it with the target tokenizer including message
framing. This example intentionally does not substitute word counts for tokens.
Keep the manifest outside model input; govern its access and retention too.
