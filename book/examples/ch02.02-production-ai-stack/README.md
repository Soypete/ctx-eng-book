# The model is one layer

A standard-library Go lookup program with invented deployment documentation and deterministic model substitutes. No API key, network, or dependencies.

From this directory:

```sh
go run .
go test ./...
go vet ./...
```

The unchecked first-result selector picks the v1 guide for a v2 request. The same selector behind the source contract returns the v2 passage. A last-result selector produces the same checked outcome. Tests assert that wrong-release, wrong-product, stale, future-dated, and incomplete evidence never reaches the model; an excluded citation cannot return; model errors produce no text; and an adapter cannot mutate the evidence used for rendering.

`Lookup` receives already-authorized documents from a trusted source adapter. It checks product/release identity, required provenance fields, source time, and the proposed citation, returning decision records. The one-hour policy is an illustrative fixture, not a recommended production threshold. `AsOf` is when the source vouched for the version, not when it was copied. Source metadata is trusted; this program cannot verify its truth.

This is a passage lookup, not a general deployment advisor: application code renders the selected source text and executes nothing. The tests prove deterministic boundaries, not real-model accuracy or document correctness. A production integration must supply authentication and source-side authorization, schema validation, publication/invalidation guarantees, deadlines, context budgets, protected decision-record storage, and evaluations of complete requests. Model adapters implement `Select(context.Context, Request, []Document) (string, error)`; prompt formatting and quality still need evaluation after a model swap.
