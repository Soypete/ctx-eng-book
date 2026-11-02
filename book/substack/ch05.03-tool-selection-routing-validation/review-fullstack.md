# Full-stack engineer review — A Timeout Doesn't Mean the Transfer Failed

**Would I finish it?** yes. The Dana opener is a bug I could ship next sprint, and the Go snippet showed up before I got bored.
**Would I share it?** yes, with the teammate wiring tool calls into our checkout assistant, once the retry-ownership gap below is closed.

## What lands
- "The first transfer went through; only the reply got lost." That's every flaky-webhook postmortem I've sat through, and it's a great hook.
- The five terminal states map straight onto an enum and a `status` column. That's an actual Monday change.
- `TestRetryWithoutKeyChargesTwice`: "Same code, same model behavior, one missing string." That's the line I'd paste in Slack.
- The diagram caption, "If it fails, the transfer stays done," sums up the whole section in one sentence.

## Where I got lost or rolled my eyes
- "the model sees an error and does the reasonable thing: it tries again." In the opener it's the *model* that retries. The fix is a retry loop inside the *harness*. After you return `FailedRetryable`, what stops the model from proposing `transfer_funds` again on the next turn? And if the user types "try again," that's a new message. Doesn't "a request ID assigned when the user hits send" give it a new key and charge Dana twice anyway? That's the hole I'd actually fall into.
- "Every step from one state to the next needs explicit state, explicit scope… and exit criteria," plus the next two paragraphs. This is the abstract stretch. Provenance, staleness, and credentials all go by with no code, and I skimmed it.
- `failed_retryable` is defined as "unknown whether money moved." Calling it "retryable" made me think it's safe to retry blindly. What does my UI show, "pending"?
- The loop has no backoff or context deadline. Copy that into a request handler and you've blown your latency budget.

## Missing for me
- Where the key lives: a Postgres table with a unique constraint on `(user_id, request_id)` and a stored result, so a repeat call (from the harness *or* the model) returns the first outcome.
- What to do when the payment API *doesn't* accept idempotency keys. Do I look up status by reference and reconcile?
- Three lines on how `failed_retryable` gets resolved: a background job, a status poll, or a user prompt.

## Top 3 edits (ranked)
1. Close the opener's loop. Show the harness deduping a *model-proposed* retry (and a "try again" from the user) against the same key, not just its own internal loop.
2. Add a short storage sketch for the idempotency key (table + unique constraint + cached result), and say what to do when the downstream service has no key support.
3. Cut or compress the provenance and scope paragraphs in "Routing is a pipeline problem," or tie them to a field in the Go `Result`. Add backoff and a context deadline to the snippet.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
