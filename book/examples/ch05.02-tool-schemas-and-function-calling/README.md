# The schema is one field of the contract

Companion code for the Substack post on tool schemas and function calling.
Go 1.22 or newer, no dependencies or model credentials:

```sh
go run .
go test ./...
go vet ./...
```

- `contract.go`: a `Contract` that carries purpose, when to use, when not to
  use, effects, the argument shape, and the role required. `Description()` is
  the text the model sees.
- `harness.go`: `Check` runs three stages in order (syntax, semantic,
  authorization) and records only calls that pass all of them.
- `eval.go` and `main.go`: an eval with cases in five classes: correct
  selection, no-tool abstention, malformed arguments, schema-valid but
  semantically invalid arguments, and unauthorized requests.

The `Selector` is a stand-in for the model. `recorded` replays fixed choices so
the tests are deterministic; `TestEagerSelectorFailsAbstention` shows that a
selector that always calls a tool passes the happy path and still fails the
abstention case. To evaluate your own tool descriptions, replace the selector
with a model call and keep the cases. Customers and requests are invented.
`checkShape` covers only the schema features used here.
