# Explicit context overflow

Run with Go 1.22 or newer; no dependencies or model credentials:

```sh
go run .
go test ./...
go vet ./...
```

The invented checkout fixture puts optional investigation history before a required
runbook in retrieval order. Assembly reserves the required working set, measures
complete JSON candidates, and includes optional sources only when they fit.
The program prints the exact input and a separate manifest with provenance,
source confirmation and retrieval times, scope, omission reasons, and positions.
If the task plus required evidence exceeds the limit, it emits no request.

The demo measures **bytes**, not tokens, and picks a limit equal to the required
JSON's length. This is a deterministic assembly test, not a model-capacity or
answer-quality benchmark. `Measure` is injected: a production integration must
serialize the actual provider messages, include framing and tools, count with
the target tokenizer, and derive the available input budget from the interface's
output accounting rules. Byte counts must never be substituted for token counts.

Authorization, freshness validation, and selection of required evidence happen
upstream and are trusted here. Required items retain their input order; optional
items follow them in input order. This placement is illustrative, not an optimal
model ordering. The example does not summarize or implement deferred retrieval.
A nil payload on any error means the caller must withhold the model request.
The manifest is separate from model input and must be stored under appropriate
access and retention rules in a real system.

Tests cover optional overflow without losing required evidence, exact-fit
inclusion, required overflow without a payload, full-request accounting,
provenance and positions, and measurement failure.
