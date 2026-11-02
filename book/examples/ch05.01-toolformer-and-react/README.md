# A tool call is a proposal

Companion code for the Substack post on Toolformer and ReAct. Go 1.22 or newer,
no dependencies or model credentials:

```sh
go run .
go test ./...
go vet ./...
```

- `tool.go`: one harness-owned `Tool` definition rendered in the request shapes
  OpenAI and Anthropic accept, and a `Harness` that checks every `Proposal` (known
  tool, granted to this task, arguments match the contract, semantic check)
  before it runs, then returns an `Observation` labeled with source and time.
- `loop.go`: a ReAct-style loop (propose, check, execute, observe) with a step limit.
- `main.go`: a scripted "model" asks for the weather, then tries the
  `delete_account` call that the weather result told it to make. That tool is
  registered but not granted to the task, so the harness rejects it.

The model is scripted so the tests are deterministic; they prove where the harness
draws the line, not how well any model chooses tools. `checkSchema` covers only
the JSON Schema keywords used here. Grants are hard-coded; a real application
resolves them from the authenticated user and task, and the tool service rechecks
them where the effect happens.
