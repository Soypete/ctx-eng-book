# A timeout is not a failed transfer

Companion code for the Substack post on tool routing and validation. Go 1.22
or newer, no dependencies or model credentials:

```sh
go run .
go test ./...
go vet ./...
```

- `workflow.go`: `Workflow.Run` takes the model's `Proposal` (recipient name,
  amount) and routes it through resolve payee, authorize amount, transfer, and
  notify. Every run ends in one of five outcomes: `completed`,
  `needs_clarification`, `denied`, `failed_retryable`, `failed_terminal`.
- `ledger.go`: a fake payment service that honors idempotency keys and can be
  told to commit a transfer and then time out.

The tests show the claims in the post: a timeout followed by a retry with the
same key charges once; the same retry without a key charges twice; a failed
notification leaves the committed transfer alone; ambiguous or unknown payees
stop before any money moves; and the summary returned to the model has no
account numbers.

Not shown: real authentication, persistence of the idempotency key and run
state across process restarts, compensation for multi-ledger transfers, or a
notification queue. People, accounts, and amounts are invented.
