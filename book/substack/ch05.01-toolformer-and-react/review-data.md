# Data engineer review — Your Model Can Ask for a Tool. Your Code Decides.

**Would I finish it?** yes — short, and "registered vs. granted" is the same fight I have over catalog entries vs. access policies.
**Would I share it?** yes — with the platform team getting "can the agent query the warehouse?" tickets.

## What lands
- "The observations are new data entering the model's context, and their source and freshness have to stay visible." That's provenance and freshness, said plainly. Good.
- "*Registered* and *granted* are different questions." Discoverable in the catalog isn't the same as having SELECT. I'll steal this framing.
- "the service behind the tool checks again where the effect happens." Enforce at the source, not only in the client. Correct.

## Where I got lost or rolled my eyes
- "labeled with its source and time." Which time? The time the harness fetched it, or the time the weather data was observed? A cached reading from 40 minutes ago, fetched now, gets a fresh-looking timestamp. That's event time vs. processing time, and it's exactly how stale data gets passed off as current. The post says freshness matters and then never shows what the label contains.
- "terrible when you need it to query a database exactly once." Good line, but the post never comes back to it. Nothing about retries, idempotency, or what happens when the loop proposes the same call twice. Calling `get_weather` twice is harmless. Calling a query billed by bytes scanned, or a write, is not.
- "Grants are hard-coded; a real application resolves them from the signed-in user." Grants here work per tool. My policies work per row and per column. A granted `query_orders` tool that returns every customer's PII is still a leak. Tool-level grants are necessary, not enough.

## Missing for me
- A concrete observation struct: `source`, `fetched_at`, `as_of`, maybe a request ID I can join to the tool's own logs. Show the label, don't just name it.
- Even one sentence on logging proposals, rejections, and observations as an audit trail. When security asks "what did the agent read and why," that's lineage.
- Cost: who pays when the model proposes an expensive call that passes schema and grant checks?

## Top 3 edits (ranked)
1. Show the labeled observation in the code or a short snippet, with separate fetch time and data as-of time.
2. Add one line that tool grants sit on top of data-level access policy (row/column scoping); they don't replace it.
3. Close the "exactly once" loop: a sentence on idempotency or budgets for repeated or costly calls, or cut the phrase.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
