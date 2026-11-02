# AI engineer review — A Timeout Doesn't Mean the Transfer Failed

**Would I finish it?** yes. The opener is a failure I've watched happen, and the post is short.
**Would I share it?** yes, with the junior folks on my team who wire up payment or ticketing tools. I wouldn't send it to peers, because there's nothing here they haven't seen.

## What lands
- "A timeout is 'unknown,' not 'failed'" is the right framing, and the key-vs-no-key test pair makes it concrete.
- "The key has to come from the user's request, not from the attempt" is the one detail people actually get wrong.
- The notification is kept out of the transfer's retry path, and the summary sent back to the model leaves out account numbers. Both are good.

## Where I got lost or rolled my eyes
- "The model sees an error and does the reasonable thing: it tries again." In the fix, the *harness* retries, not the model. The case that bites me in production is the user saying "did it go through? try again." That's a new turn, which means a new request ID, which means a new key, and Dana gets paid twice. The post never closes that hole.
- "Same code, same model behavior, one missing string." No model runs in these tests. It's a fake ledger and a deterministic loop. Calling that "model behavior" is hand-waving about evals.
- `failed_retryable`: "it's unknown whether money moved." Guideline 3 says unknown, but the state's name says retryable. When the outcome is unknown, the right move is to look up the transfer by its key and reconcile, not to retry blind.
- Idempotency keys and compensation are standard payments practice (Stripe has shipped idempotency keys for years). That's fine, but the post never says what makes this an *agent* problem as opposed to a distributed-systems problem.
- The module is "tool selection," but I didn't see any selection. Routing gets a definition and is then dropped.

## Missing for me
- A status-check step (`get_transfer(key)`) before any retry, ideally shown in the diagram.
- The agent-specific failure: the model proposes a second `transfer_funds` with a slightly different amount or payee. How does the harness detect a near-duplicate proposal?
- Backoff, and what "bound attempts per kind of error" looks like in code. The loop shown retries immediately.

## Top 3 edits (ranked)
1. Cover the cross-turn duplicate ("try again" from the user). Either derive the key from the intent (payee, amount, period) or check for a pending or unknown transfer before accepting a new one.
2. Rename `failed_retryable` to something like `unknown_outcome`, and show a reconcile-by-key step before any retry.
3. Drop "same model behavior," or add a real eval: run a model against the timeout fixture and count duplicate proposals.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
