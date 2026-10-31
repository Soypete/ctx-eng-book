---
title: Valid JSON Is Still Only a Proposal
subtitle: Examples teach the shape, but code must check meaning and permission before execution.
module: ch04.03-examples-instructions-structured-outputs
mood: happy
scheduled: 2026-11-23
---

The report-export assistant returns perfect JSON. All the required fields are present, every value has the right type, and the reporting period ends before it starts. The application sends it to the export service anyway.

This is an invented failure, but the mistake is concrete: the team checked whether software could read the request and called that validation complete. Congratulations, the nonsense has types.

In the previous post in my [context engineering series](https://substack.com/@soypetetech), I argued for resolving meaning and scope before granting an action; now I want that boundary expressed as an executable contract.

My argument is simple: **a schema-valid output is a proposal, and application code must decide whether it is fit to execute.** Better examples and clearer instructions help the model produce that proposal. They don't replace the decision.

## Give each check a job

**Structured output** is a response organized into a specified representation, such as a JSON object. A **schema** describes its permitted shape: fields, types, required values, and other declared constraints.

For the export request, our schema requires nonempty strings named `dataset`, `start`, and `end`, with no extra fields. An object containing `sales`, `2026-11-01`, and `2026-10-01` satisfies that contract. It still describes a backwards reporting interval.

**Syntax** is the form software can parse. **Semantics** is whether the values and their relationships mean the right thing. **Pragmatics** concerns whether the result is appropriate for the intended purpose and actor.

Those distinctions give us separate questions:

| Check | Question for this export |
|---|---|
| Parse and schema | Is this an object with the required string fields? |
| Semantic validation | Are these real dates, with the start before the end? |
| Authorization | May this authenticated caller export this dataset? |
| Execution result | Did the export service actually complete? |

A richer schema can express more constraints. That moves some checks earlier; it doesn't establish that the requested dataset matches the user's intent or that the caller has permission. I want each rule enforced where the application has the evidence needed to decide it.

## Examples teach the boundary; code enforces it

**Few-shot examples** are sample inputs and expected outputs supplied alongside the task. They can demonstrate which label to choose, how to represent a date, or when to abstain. They influence a probabilistic response; they don't program a deterministic action.

For this export interface, I would show a resolved reporting interval that becomes a proposal and an unresolved interval that leads to clarification before proposal generation. A useful instruction says: “Use the selected dataset and resolved dates. If either is missing, request clarification. Do not invent a reporting period.” The product needs an explicit clarification path to make that instruction actionable.

Examples should earn their place by distinguishing behavior the model otherwise confuses. Repeating a happy path with different dataset names doesn't teach what happens when “last month” is unresolved. Nor does a beautifully written explanation prove the model follows a faithful internal reasoning process.

Keep examples separate from the **held-out cases**, reviewed inputs reserved for evaluation rather than included in the prompt. Compare task outcomes on those cases when changing instructions or examples. If the model still selects the wrong operation despite a clear interface, retain model capability as a possible cause. Prompt editing isn't a universal repair.

Drew Breunig's [Let the Model Write the Prompt](https://www.dbreunig.com/2025/06/10/let-the-model-write-the-prompt.html) argues, “Decouple your task from the LLM. Write tasks, not prompts.” My application here is to put the acceptance contract in the program, where changing the model doesn't silently change what counts as an acceptable export.

## Asking for JSON isn't enforcing JSON

**Schema prompting** asks the model to follow a declared structure. The returned text remains untrusted until checked. Validation detects nonconformance; it cannot retroactively make malformed output valid.

**Constrained decoding** limits generation to tokens allowed by a supported grammar or schema. It can enforce membership in that supported syntax. It doesn't make the values true, current, or authorized.

The **harness** is the application code that supplies model context and controls execution. Its job continues after parsing. That applies to an ordinary extraction workflow as much as to an **agent**, a model that calls tools in a loop.

![Examples and instructions guide a model proposal through schema, semantic, and authorization checks before execution; rejected proposals stop and outcomes produce decision records.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch04.03-examples-instructions-structured-outputs/diagram-1.png)

*Notice that passing the schema opens another check, not a direct route to execution.*

## Prove that the bad request stops

The [companion Go program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch04.03-examples-instructions-structured-outputs) use invented requests and a fake export function. There is no model call. The folder includes the schema and a parser implementing that fixed contract, including rejection of unknown fields, duplicate keys, and trailing content.

After parsing, `p` contains the proposal, `principal` comes from the trusted session, and `policy` supplies export permissions. A **decision record**, `d`, stores the check's outcome and reason. These lines come directly from the program:

```go
start, startErr := time.Parse("2006-01-02", p.Start)
end, endErr := time.Parse("2006-01-02", p.End)
// Well-typed strings can still describe an impossible reporting interval.
if startErr != nil || endErr != nil || !start.Before(end) {
	d.Stage, d.Reason = "semantic", "require valid dates and start before end"
	return d
}
// Permission comes from the trusted session and policy, never model JSON.
if principal == "" || policy.Version == "" || !policy.Exports[principal][p.Dataset] {
	d.Stage, d.Reason = "authorization", "export not permitted"
	return d
}
```

Run `go run .` to see a reversed interval rejected, an otherwise valid payroll request denied, and an allowed sales request reach the fake export function. Run `go test ./...` to assert that rejected proposals make no export calls. The tests also cover invalid dates, missing identity, and service failure.

This proves a code property: those failing checks prevent execution. It does not measure how often a model produces a correct proposal. It also does not implement authentication or a real export service. The service must enforce permissions when it acts, and a reported execution error must not automatically trigger a retry that could duplicate a partially completed export.

## Keep evidence outside the proposal

For a data engineer, a typed source reference is only the beginning. **Provenance** records where information came from. A model-generated `source_version` string cannot establish which source the application actually read.

I would retain the contract and policy versions separately from runtime evidence: the selected dataset, resolved interval, actual source version, and source-backed confirmation of freshness. A recent retrieval timestamp doesn't establish that the underlying values are current. When freshness matters to acceptance, verify it through the source integration before using the result. The small Go example leaves that integration out.

Keep the invocation context and retrieval trace available under appropriate access and retention rules. The decision record should distinguish schema rejection, semantic rejection, authorization denial, execution failure, and success. Don't turn the audit trail into an unrestricted copy of customer data.

For extraction evaluation, the module uses the familiar [information-extraction metrics](https://github.com/Soypete/ctx-eng-book/blob/main/research/knowledge-graphs/ie-ner.md): **precision** measures correct extracted items among extracted items, and **recall** measures recovered relevant items against labeled relevant items. Those measure task quality. Track rejected proposals, unauthorized execution, and appropriate abstention separately; a parse-success rate cannot stand in for them.

## Guidelines

1. **Define the output contract before writing examples.** Test missing fields, wrong types, and unexpected fields against it.
2. **Demonstrate difficult boundaries and clarification behavior.** Evaluate changes on held-out inputs, including missing scope.
3. **Test schema-valid nonsense.** Assert that reversed or impossible dates never reach execution.
4. **Authorize using trusted identity and policy.** Test a well-formed, meaningful request for a forbidden dataset.
5. **Verify provenance and required freshness through source integrations.** Reject model-written metadata as proof of either.
6. **Record the failed stage and measure task outcomes separately.** Keep execution failures distinct from rejected proposals and compare model quality beyond parsing.

I want structured output to give us an interface we can inspect, test, and refuse when its contents fail the task. Examples communicate expectations; executable checks decide what may happen next. Next in the [series: Toolformer and ReAct](https://substack.com/@soypetetech), I'll follow the proposal across the tool boundary, where selection, execution, and observations become a loop.
