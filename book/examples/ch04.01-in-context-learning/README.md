# Inspect the assumption

Run with Go 1.22 or newer, no dependencies or API credentials:

```sh
go run .
go test ./...
go vet ./...
```

This synthetic classification fixture defines `buyer` as an explicit purchase
commitment. A pricing inquiry must yield `unknown`. The program prints the exact
serialized invocation, its SHA-256 identity, and decisions for two scripted
candidate answers. Valid JSON with the wrong meaning is rejected.

Tests verify semantic rejection, accepted abstention, required metadata, snapshot
round trips, and trace changes when example content/order, source version, or the
rule changes. These are application checks, not measured model performance.
`Check` is an offline evaluation function with a reviewed expected label; it is
not a general production oracle. The decoder is not a strict JSON schema validator.

The timestamp is synthetic provenance, not an implemented freshness check.
Production code must enforce access/freshness before assembly, retain actual
provider request bytes (including message roles and tool definitions), model
identity/settings and responses, and protect snapshots under retention rules.
A digest identifies retained bytes; it cannot reconstruct deleted content or
establish source authority. Identical inputs do not promise identical model outputs.
