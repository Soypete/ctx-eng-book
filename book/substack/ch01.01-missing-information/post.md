---
title: Before You Blame the Model, Check Three Gates
subtitle: A missing source, a retrieval miss, and an unsupported answer need different repairs, even when they look like the same model failure.
module: ch01.01-missing-information
mood: confused
scheduled: 2026-10-09
---

You build a document Q&A feature, demo it with one clean PDF, and it nails the answer. Everyone claps. In production, a user uploads ten files and gets fluent, confident nonsense.

Three files are scans. Two are in German. One is a spreadsheet. One is a contract where the appendix is the only part that matters. The rest are fine.

The team's first instinct is to say the model hallucinated and go shopping for a bigger one. I want to inspect the input first. This may be a bug report about the system around the model, filed under the wrong name.

This is post 2 of my [context engineering series](https://substack.com/@soypetetech): post 1 covered controlling what reaches the model; now I want to diagnose what went wrong along that path.

## The demo constrained the problem. Production doesn't.

The **harness** is the code around the model that builds its input and runs its tools. In the demo, we handed it clean inputs and a known question. Production requires it to detect, route, reject, or recover from inputs nobody planned for.

The harness needs to be told which files exist, how to process them, and what the user wants back. I call those **Lexicon** (sources, identities, and authority), **Semantics** (formats, meaning, and relationships), and **Pragmatics** (the task, permitted actions, and required output). They're questions the harness needs answers to before “answer questions about these files” becomes an executable task.

An **authoritative** copy is one the source owner vouches for as current for the task. **Provenance** records where it came from, its version, its as-of date, and whether a newer copy supersedes it. For user uploads, ask the user which copy governs the task. Existence alone doesn't make it usable evidence.

## Three gates before you blame generation

The **working context** is what actually gets put in front of the model for this request. I check three gates in order:

1. **Source:** Is the needed fact available, authoritative, current, and usable, and does every admitted copy carry provenance showing it is current? This checks Lexicon and the Semantics needed to interpret the data.
2. **Retrieval:** Did the search and selection steps admit the relevant evidence into working context? Check Lexicon’s identity and scope: whose permissions the search used and which collections it could access.
3. **Generation:** Did the model use that evidence faithfully and stay within the task? This checks Pragmatics as well as support for the answer's claims.

![A wrong answer passes through Source, Retrieval, and Generation checks; the Source gate asks whether the evidence is current and versioned.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch01.01-missing-information/diagram-1.png)

*Check freshness before diagnosing retrieval; only the final failure branch calls for investigating generation.*

| Failed gate | What it looks like | Repair |
| --- | --- | --- |
| Source | Missing fact, unusable file, or stale copy presented as legitimate | Ingest, process, version, correct authority, or ask the user. A bigger prompt does nothing. |
| Retrieval | Usable evidence exists but is missing or buried in context | Fix the query, index, ranking, scope, or amount admitted. |
| Generation | Adequate evidence reached the model, but the answer is unsupported or violates the task | Change instructions, validation, or the model; evaluate again. |

Back to our ten files: scans without optical character recognition (OCR), German documents without language handling, and a spreadsheet without suitable processing fail **Source**. The data isn't in a usable, described form. A usable contract appendix that never reaches context fails **Retrieval**. If the appendix reaches the model and it invents a conflicting contract term, investigate **Generation**.

For a refund question, admitting `policy-v1` when `policy-v3` is current is a **Source** failure. The source let an obsolete copy look legitimate. Ranking cannot reliably choose the current policy when the records don't say which version supersedes which.

Retrieval-augmented generation supplies retrieved evidence to a model before it answers. These gates overlap with its evaluation: context recall asks whether needed evidence arrived, and faithfulness asks whether the answer follows it. [RAGAS](https://arxiv.org/abs/2309.15217) likewise evaluates retrieval and faithful use separately. My emphasis here is the earlier Source gate, plus the identity and scope checks. Faithfulness to an obsolete policy still produces the wrong answer.

## Keep the trace, then label the failure

You can only tell the gates apart if you kept the evidence. A **trace** is the record connecting a request to its inputs and answer. An **eval set** is a collection of labeled questions with expected evidence and behavior.

The runnable Go example classifies **labeled traces**: eval cases or production failures someone has triaged. Its `Trace` record holds the question and the evidence fields below; `Gate(t Trace) string` takes that record and returns a failure message or `"ok"`.

For production, I'd store a request ID, question, requester identity and scope, source versions, admitted evidence references, and final answer together in a protected trace table. When a ticket arrives, a reviewer identifies the required facts (`Needs`), checks source history for the authoritative records applicable to the request (`Current`), and reconstructs what reached the model (`Retrieved`). Then run `Gate`.

`Unsupported` holds answer claims a checker found unsupported by the evidence. Start with human review: split the answer into factual claims and match each to supporting evidence. That matching is **attribution**. An automated checker can assist; validate it against reviewed cases.

Here are the evidence fields and the stale-copy branch from `Gate`; `r` is an admitted record and `cur` is the source's current record for the same fact. The branch belongs inside the function's evidence loop.

```go
type Evidence struct {
	ID, Fact, Version string
	AsOf              time.Time
	Authoritative     bool // the source vouches for this copy
	Superseded        bool // the source has marked a newer version
}

			// A stale copy is a provenance failure at the source: if the source
			// does not version and supersede its records, an old copy looks
			// exactly as legitimate as the current one.
			if !r.Authoritative || r.Superseded || r.Version != cur.Version {
				return fmt.Sprintf("source: stale copy %s@%s competes with %s@%s; mark authority and supersession at the source",
					r.ID, r.Version, cur.ID, cur.Version)
			}
```

The [full program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.01-missing-information) also check missing provenance and missing evidence. Run `go test ./...` in that directory. The refund tests return `source: stale copy` for both the obsolete copy alone and obsolete-plus-current copies together.

Use stable evidence references: document ID, version, and effective date, with a mapping to the fact being checked. A hash of a document fragment—a chunk—changes when you split the document differently. Retain immutable evidence snapshots when replay requires them; references alone cannot reconstruct overwritten content. Set retention and access rules for traces, redact personal data where possible, and measure storage cost before duplicating every retrieved passage.

`Gate` trusts its labels and source records; with empty `Needs` and `Unsupported`, it returns `"ok"` even for an unlabeled trace. It doesn't measure OCR quality, enforce permissions, detect arbitrary contradictions, or catch evidence buried among distractions. Its `"ok"` means only these checks passed.

## Tools pass through the same gates

An **agent** is a model calling tools in a loop. Tool descriptions and schemas—the declared shapes of their inputs—tell it what it may request. They don't prove it chose correctly or that an action succeeded.

A missing or stale tool result fails Source. Burying the useful result among too many tool outputs fails Retrieval/selection. Skipping a required call or inventing claims beyond a usable result violates the Generation task boundary. Record each request and result, then validate what happened.

Tool output is untrusted input. A retrieved document can contain instructions intended to redirect the task, an application of the prompt-injection threat model studied by [Liu et al., 2024](https://arxiv.org/abs/2310.12815). Model Context Protocol (MCP) standardizes tool discovery and invocation. The harness still must authorize calls, validate arguments, and isolate untrusted content.

## A bigger model still needs evidence

I use “hallucination” for an answer that asserts unsupported or contradictory information as though it were grounded. Missing, stale, poisoned, or poorly selected evidence can produce that symptom. Better context helps; it doesn't guarantee truth.

A bigger model may handle ambiguity better. It cannot supply missing authorization, freshness, processing requirements, or acceptance criteria. Give the harness paths to ask, retrieve more, or stop when evidence is inadequate.

“Every failure is a context failure” is a debugging provocation, not a law. If adequate evidence and enforced boundaries still fail a controlled evaluation, congratulations: you found a real model limit, and now you can prove it.

## Guidelines

1. **Version and supersede source records.** Require provenance and classify stale copies in context as Source failures.
2. **Record live traces; diagnose labeled cases.** Fill `Needs` and verify `Current` before interpreting `Gate`.
3. **Classify failures before changing the model.** Check Source, Retrieval, then Generation against retained evidence.
4. **Preserve replayable evidence under access and retention rules.** Verify that references survive document reprocessing.
5. **Check claims and task boundaries separately.** Review attribution labels and test contradictions and overload.
6. **Enforce tool boundaries in the harness.** Authorize calls, validate arguments and results, and test that retrieved instructions cannot grant permissions.

I want the incident ticket to name a repair: fix the source, fix selection, or fix generation. A stale policy needs a data repair before another model comparison. Next in the series, I'll extend these gates to state: what the workflow has done, what it may do next, and when it must stop.
