# Data engineer review — Five Tests Every Tool Needs Before an Agent Can Call It

**Would I finish it?** yes — the duplicate-Ana opener is a data-quality incident I've cleaned up, so I'm in from line one.
**Would I share it?** yes, with caveats — with the platform team that's getting "can the agent write to the CRM?" tickets, but I'd add a comment about the freshness gap.

## What lands
- "every type is right... The call passes validation. Now there are two Anas." A data contract that checks shape and never checks meaning. I live with this every day.
- The contract fields include "where its data comes from and how fresh it must be". It's good to see the author put provenance next to authorization.
- "in a real system it belongs in your policy data". Yes. Rules belong in governed tables, not in `if` statements.

## Where I got lost or rolled my eyes
- "which source is authoritative for customer data, and how fresh it has to be" shows up in the can't-say list and in the contract, then none of the five tests checks it. If `checkMeaning` reads a replica that's 20 minutes behind, Ana is still missing there and the duplicate check passes. Freshness is named as a check and then never enforced.
- "whether a repeated request should create a second record or return the first" also gets dropped. Two concurrent proposals can both pass `checkMeaning` before either one writes. That's a check-then-act race. The real guard is a unique constraint or an idempotency key at the store, not a lookup in the harness.
- "duplicates, unparseable email": duplicate by what key? `Ana@Example.com` vs `ana@example.com` is how most CRMs end up with two Anas. Normalization is part of meaning.
- The diagram says the harness "records the call". The code says `h.Executed = append(...)`. What gets recorded: who, which contract version, which source snapshot? That's my lineage, and the post skips it.

## Missing for me
- A sixth test class, or a sub-case of #4: **stale or wrong-source read**. The authoritative store has Ana, the harness's read path doesn't, and the call must be rejected.
- One line on schema evolution. When `tier` gets a new enum value, which version of the description did the model see, and does the eval pin it?
- An audit-record example: one JSON row per call with stage reached, principal, contract version, and source/as-of timestamp.

## Top 3 edits (ranked)
1. Make freshness and source authority testable. Add a stale-read case to test #4 (or make it a sixth test), or stop listing it as part of the contract.
2. Say that the duplicate check needs a store-level guarantee (unique key or idempotency key), and name the normalized key.
3. Show what "records the call" writes, so audit and lineage aren't left as an implication.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
