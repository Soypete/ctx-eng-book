# Data engineer review — Before You Blame the Model, Check Three Gates

**Would I finish it?** yes. "Source" is a gate, so this one is talking to me and not just the prompt people.
**Would I share it?** yes, with the analytics engineers getting pulled into "the RAG bot is wrong" tickets.

## What lands
- "A bigger prompt does nothing" on a source failure. Thank you. That's a data quality issue, and the post names it as one.
- The Trace struct. Keeping question, expected evidence, retrieved IDs, and answer together is basically lineage for an answer. I can build that.
- "Each 'no' has a different owner." Ownership is how I decide whether a ticket is mine.

## Where I got lost or rolled my eyes
- "the source holds `policy-v3` but retrieval admitted `policy-v1`, and it says `retrieval`." Why was v1 still in the index at all? The table files "only a stale copy" under **Source**, and then the worked example calls the same symptom a retrieval bug. In my world that's a broken deprecation or backfill, and it belongs to ingestion, not ranking. The post contradicts itself on the case I care most about.
- `InSource map[string]bool`. A bool can't say "authoritative, usable form." It doesn't tell me which version, as of when, or whether the OCR output was garbage. Gate 1 promises freshness and authority, but the code only checks whether the thing exists.
- "evidence IDs a correct answer needs." Whose IDs? Chunk IDs change every time someone re-chunks or re-embeds. If the IDs aren't stable across reindexes, your eval set quietly goes bad.

## Missing for me
- A definition of an evidence ID that survives schema evolution: document ID plus version plus effective date, not chunk hash.
- Something on what the traces cost. Storing retrieved context for every answer means retention, PII, and access policy decisions, since the trace now holds whatever the user was allowed to see.
- The ten-file example brings up scans, German, and spreadsheets, and then never ties them back to a gate. Is a failed OCR a source failure or a retrieval failure?

## Top 3 edits (ranked)
1. Settle the stale-copy case. Either reclassify the v1/v3 example as a source/ingestion failure, or explain why serving a superseded version counts as retrieval. Add `Version`/`AsOf` to the trace so the code can tell the difference.
2. Replace `InSource bool` with something that can say "present but stale/unusable," so the code matches the gate definition.
3. Add one line on stable evidence IDs and one on trace retention/governance.

## Miriah's notes

- **Stale copy = Source problem (provenance).** When retrieval returns an outdated copy (policy-v1 instead of v3), the failure is at the Source gate: the source layer didn't carry provenance/version/authority, so the stale copy looked legitimate. Make the table and the policy-v1/v3 example agree on that, and replace `InSource bool` with provenance fields (version, as-of, authoritative) so the code can express it.

- **Code:** the fixed example (provenance fields; stale copy = source failure; runs on labeled traces) lives at https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.01-missing-information with tests. Show only the key lines (≤15) and link to it; don't paste the program.

<!-- Add your notes for the rewrite here. Astra reads this section. -->
