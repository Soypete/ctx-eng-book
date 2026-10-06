# Freshness as an enforceable boundary (module 1.3)

Runnable companion to the Substack post *Every Context Failure Has an Address* and
[book module 1.3](../../chapters/ch01-every-failure-is-a-context-failure/modules/ch01.03-context-failure-case-studies.md).

The post locates five context failures at five boundaries. This example implements one of
them, freshness, using the post's stale-CRM scenario: a retrieval index synced yesterday
still says a deal is `negotiating`, but the CRM closed it this morning.

`Resolve` takes the cached fact, the query type, and a per-query-type maximum age:

- within the limit → use the cached fact;
- too old → refetch from the authoritative source;
- too old and the source can't answer → return no value (`ErrStale`), never the old one;
- no policy for this query type → fail closed (`ErrNoPolicy`).

Every call also returns a `Decision`: source, version, observed time, check time, age,
limit, the path taken, and why. That record is the evidence that the fix worked.

```bash
go test ./...
```
