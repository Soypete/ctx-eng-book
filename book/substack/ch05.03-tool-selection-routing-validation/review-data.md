# Data engineer review — A Timeout Doesn't Mean the Transfer Failed

**Would I finish it?** yes. It's short, and "a timeout is unknown, not failed" is the same lesson I learned from at-least-once delivery.
**Would I share it?** yes, with the platform team that keeps getting asked to "just expose the warehouse as a tool."

## What lands
- "Routing is a pipeline problem." Typed stages, explicit state and exit criteria are how I already think about a DAG.
- "The key has to come from the user's request, not from the attempt." That's the distinction between a dedupe key and a run ID, and most people get it wrong.
- The two tests, `TestTimeoutThenRetryChargesOnce` and `TestRetryWithoutKeyChargesTwice`. "One missing string" is a good line for a postmortem.

## Where I got lost or rolled my eyes
- "record where each value came from and when (a payee list loaded yesterday may miss today's change)". This is the most important data sentence in the post, and it's a parenthetical. What happens to it? Nothing later checks freshness. Is Dana resolved from a cache, a replica, or the source of record? A stale payee is a real way to pay the wrong person, and it never comes up again.
- "A later run with the same key can't double-charge." Only if the key and the run state survive. Where is the key stored? For how long? Who replays `failed_retryable` runs: a person, a scheduler, the model? The post makes durability sound free, and it's the actual design work.
- "executed once" in the diagram caption. The service dedupes. Your side is still at-least-once. Say that.

## Missing for me
- The record schema. The diagram has a "recorded" step, and the next post promises to mine "these outcomes and attempt counts," but I never see the columns: request_id, idempotency key, outcome, attempts, txn_id, notified, timestamps, source version of the payee. If that's the future evidence, show the table now.
- A name for the notification pattern. A separately retried side effect after a commit is a transactional outbox. Naming it gives readers like me something to search for.
- Retention and access for that record. It holds payees and amounts, so who can read it?

## Top 3 edits (ranked)
1. Answer where the idempotency key and run state are persisted, how long the dedupe window is, and who retries `failed_retryable`. Without that, "a later run can't double-charge" is a promise and not a design.
2. Turn the provenance parenthetical into a step: resolve the payee from a source with a known freshness, record that version, and treat stale data as `needs_clarification`.
3. Show the outcome record's schema and call the notification step an outbox.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
