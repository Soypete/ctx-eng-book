# Compaction must preserve a recovery path

Run with Go 1.22 or newer, from this directory:

```sh
go run .
go test ./...
go vet ./...
```

No dependencies, network calls, or model credentials. The invented refund policy
has a general rule and an annual-plan exception. `Compact` deterministically
omits the exception while retaining source lineage. `Resolve` recovers the
original rule through a source adapter; deleting that source produces no answer.
The JSON output includes both decisions and the summary manifest.

Tests compare recovery with the uncompacted exception, demonstrate missing-source
failure, and reject changed identity/version, denied scope, and invalid freshness.
The 30-day refund window and one-hour freshness limit are fixture choices, not
research results or recommended defaults. Text reduction is not token accounting.

This proves recovery and refusal behavior, not model summarization quality.
The in-memory source adapter uses a trusted caller scope and source confirmation
time. Production needs authenticated identity, real access and retention policies,
source-backed timestamps, persistent storage, bounded reads, and task evaluation.
A changed version deliberately requires rebuilding the summary; this demo stops
there. A manifest is lineage, not authorization and not a copy of the source.
