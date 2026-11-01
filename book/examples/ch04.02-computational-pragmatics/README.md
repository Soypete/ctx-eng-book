# Resolve scope before reading

Run with Go 1.22 or later, using only the standard library:

```sh
go run .
go test ./...
go vet ./...
```

The invented fixtures produce `clarify` with zero reads, `denied` with zero reads,
and `read` with one call restricted to the resolved sales dataset and interval.
Tests also cover identity, missing fields, invalid intervals, unsupported actions,
policy metadata, decision records, and source failure without returning partial data.

This is a deterministic boundary demonstration, not a language model evaluation.
`Request` must come from a trusted resolution flow: the evidence reference is
metadata, not a verified credential. A model could fill every field incorrectly.
Session and policy are server-owned; never deserialize them from model output.
The production data service must enforce access too. No real authentication,
retrieval, analysis, freshness validation, evidence lookup, or audit storage is
implemented. The UTC interval is illustrative, not a definition of “past month.”
