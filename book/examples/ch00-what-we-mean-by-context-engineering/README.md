# Context compiler (chapter 0)

Runnable companion to the Substack post *Context Engineering Is Not Prompt Engineering* and
[book chapter 0](../../chapters/ch00-what-we-mean-by-context-engineering.md).

`Compile` admits already-scored search candidates into a model's working set: it filters by
authorization, ranks with deterministic tie-breaks, applies a token budget, and records every
decision in a versioned `Manifest`. It is a post-filter. It keeps unauthorized text out of model
input, but it cannot recover eligible documents a top-k search never returned, so production
systems should also push access scope into the retrieval query.

```bash
go test ./...
```
