# Personalization as governed context assembly (module 1.4)

Runnable companion to the Substack post *Your Model Doesn't Know Your Users. Your Retrieval
Does.* and [book module 1.4](../../chapters/ch01-every-failure-is-a-context-failure/modules/ch01.04-personalization-failures.md).

Personalization is built at request time, outside the model:

1. `Sessions.Authenticate` derives the user from the session token. Nothing in the request
   body (and nothing the model says) chooses whose data is read.
2. `TaskPolicy.Issue` grants a short-lived capability for only the fields the task needs.
3. `Store.Read` enforces the capability: rows for other users and fields outside the grant
   never leave storage.
4. `Assemble` excludes preferences older than their field's max age and lets a declared
   preference override an inferred one, recording every exclusion and why.

The tests cover the post's four failures: cross-user reads fail, a forged or expired
capability reads nothing, a code-review task can't see the billing email, a stale project
preference is excluded (and recorded), and "I write Go" beats an inferred "Python."

This is an in-memory illustration. In production, enforce step 3 in the database
(row-level security or a policy-filtered view) and sign the capability.

```bash
go test ./...
```
