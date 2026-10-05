# Three gates (module 1.1)

Runnable companion to the Substack post *Before You Blame the Model, Check Three Gates* and
[book module 1.1](../../chapters/ch01-every-failure-is-a-context-failure/modules/ch01.01-missing-information.md).

`Gate` classifies a wrong answer by the first gate it failed: **source**, **retrieval**, or
**generation**. A stale copy (policy v1 reaching the context while v3 is current) is a **source**
failure: the source did not carry provenance (version, as-of, authority, supersession), so the old
copy looked legitimate.

`Gate` runs on labeled traces, such as eval cases or production failures you have triaged, because
it needs to know which facts a correct answer requires. `Unsupported` comes from a claim checker
(human review or an attribution eval); building that checker is out of scope here.

```bash
go test ./...
```
