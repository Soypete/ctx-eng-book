# Schema-valid is only a proposal

Run with Go 1.22 or newer, with no dependencies or model credentials:

```sh
go run .
go test ./...
go vet ./...
```

The invented export requests demonstrate three outcomes: reversed dates pass
`schema.json` but fail semantic validation; payroll passes syntax and semantics
but fails authorization; an allowed sales export reaches a fake service once.
The tests assert that rejected proposals never reach that service and that service
failure remains distinct from rejection and success.

`Parse` implements only the fixed contract in `schema.json`, plus duplicate-key
rejection. It is not a general JSON Schema validator. Dates use YYYY-MM-DD with
an inclusive start and exclusive end. Policy and identity are trusted inputs,
kept outside model JSON. Keep the schema and parser in sync if changing fields.

This deterministic example proves check ordering, not model accuracy. It does not
implement authentication, source retrieval/freshness, persistent audit storage,
retries, or actual exports. A production service must enforce permissions at use,
verify evidence against authoritative sources, and handle partial execution and
idempotency before retrying. Decision records use generic failure reasons to avoid
copying service error details. Store them under appropriate access and retention rules.
