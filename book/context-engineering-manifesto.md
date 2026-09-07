# The Context Engineering Manifesto

## Architectural Laws for LLM and NLP Systems Operating on Data

## Abstract

AI systems fail in production not because models lack intelligence, but because
they operate inside a context void. They receive incomplete data, ambiguous
meanings, and weak action boundaries, then use probabilistic reasoning to fill
in the gaps.

This manifesto defines three architectural laws for reliable AI systems:

- **Lexicon:** Data must remain attributable, governed, and meaningful.
- **Syntax:** Meaning must be represented through explicit schemas, ontologies,
  and relationships.
- **Pragmatics:** Actions must be expressed as constrained, authorized, and
  auditable operations.

Together, these laws move AI systems beyond prompt engineering toward semantic
engineering: building the background that makes model behavior useful,
inspectable, and safe.

## 1. The First Principle: Context Is Infrastructure

An agent is not intelligent in isolation.

Its behavior depends on:

```text
Model
+ Data
+ Meaning
+ Identity
+ Tools
+ Policy
+ State
```

Most current systems treat context as prompt material. Context must instead be
treated as infrastructure with ownership, lifecycle, versioning, access
control, and observability.

## 2. The Context Void

The context void appears when an AI system does not know:

- which data is authoritative;
- whether information is current;
- what an entity means;
- which user and purpose apply;
- whether an action is allowed; or
- which tool is appropriate.

The failure compounds across the agent loop:

```text
Missing data
    ↓
Ambiguous meaning
    ↓
Probabilistic inference
    ↓
Unvalidated action
```

This is the root of stale-context failures, retrieval errors, hallucinations,
prompt injection, and unauthorized tool calls.

## 3. Law One: The Law of Lexicon

> An agent may only reason from data whose source, ownership, authority, and
> access conditions are known.

The Lexicon is the governed information environment that gives business and
operational meaning to context. It includes:

- entities;
- definitions;
- relationships;
- provenance;
- ownership;
- sensitivity;
- versions;
- retention; and
- access attributes.

Relationships are not decoration. They define how an entity connects to the
business and operational meaning of the data. The Lexicon spans everything
needed to operate a company, a product, or an industrial system: intellectual
property, operations, employee data, strategy documents, sales data, and the
governed context each domain requires.

The Lexicon law rejects the idea that all useful context must be copied into a
central memory store. The objective is not data consolidation for its own
sake. The objective is reliable context with clear ownership and scope.

### Lexicon requires

- Federated data access
- Identity-aware retrieval
- Authorization and access control
- Source provenance
- Freshness and supersession rules
- Privacy controls
- Compliance and GDPR requirements
- Data sovereignty
- Explicit authority boundaries

## 4. Law Two: The Law of Syntax

> Data without semantics is information without the narrative. There is a why
> behind every piece of data that is relayed; syntax is how we operationally
> convey that reason.

Syntax defines how meaning is represented and supplied to a model’s working
context. Ontologies express domain-specific relationship meaning. An entity
means something different to a business, operations, a product, a sales team,
or an industrial system, so the ontology must make explicit which relationships
and definitions apply to the domain at hand.

For example:

```text
SSN
 ├── entity: Person
 ├── class: RegulatedIdentifier
 ├── display: MaskedOnly
 ├── purpose: IdentityVerification
 └── access: ABACRequired
```

Syntax includes:

- Ontologies
- Schemas
- Entity resolution
- Standard dataset mappings
- Relationship graphs
- Typed tool parameters
- Domain vocabulary
- Context assembly rules

Syntax prevents the model from inferring organizational meaning from vague
examples or semantically nearby documents.

## 5. Law Three: The Law of Pragmatics

> Agents hallucinate constantly; sometimes those hallucinations happen to
> accomplish what we asked of them.

Pragmatics concerns what an agent is doing through a tool call:

```text
“Send this message.”
“Approve this request.”
“Update this record.”
“Deploy this version.”
```

These are not merely strings. They are operational speech acts.

Pragmatics requires:

- Typed tool definitions
- Explicit parameters
- Purpose binding
- User and workload identity
- ABAC evaluation
- Business-rule validation
- Pre-execution interception
- Audit records
- Safe failure behavior

The model can propose:

```text
update_patient_record(patient_id, field, value)
```

The middleware must decide whether the action is valid, authorized, and safe.

## 6. The Determinism Boundary

Models are probabilistic:

```text
Given context, predict the most likely next action.
```

Security and business policy must be deterministic:

```text
Given identity, purpose, resource, and action, allow or deny.
```

Therefore:

> Never ask a probabilistic model to make a deterministic authorization
> decision.

The model can reason about what it wants to do. The Neural Proxy decides
whether it may do it.

## 7. The Semantic Background Architecture

```text
Human Principal
       ↓
Task and Purpose
       ↓
Neural Proxy
  ├── Lexicon resolution
  ├── Syntax / ontology injection
  ├── Context retrieval
  ├── Identity and ABAC evaluation
  ├── Tool-call validation
  ├── Model routing
  └── Audit and trace
       ↓
Agent / Model
       ↓
Validated Tool Call
       ↓
Enterprise Data and Services
```

The Neural Proxy is not the agent. It is the background that gives the agent
meaning and boundaries.

## 8. Mapping the Laws to Haikei

| Manifesto concept | Haikei implementation or direction |
|---|---|
| Principal identity | OIDC bridge and user identity |
| Lexicon / data access | Connected data sources and governed context |
| Syntax / meaning | Tool definitions, mappings, and ontology layer |
| Pragmatics / action | ABAC authorization before tool execution |
| Neural Proxy | Kei middleware and ABAC engine |
| Sovereignty | VPC or self-hosted control plane |
| Accountability | Audit logs with user attribution |
| Agent interoperability | MCP, BYOA, and Agentware SDK |
| Policy enforcement | Policy decision and enforcement points |
| Operational evidence | Traces, evaluations, and action outcomes |

The existing implementation strongly supports the principal, ABAC, audit, and
proxy claims. Ontology and semantic-injection capabilities should be described
as the next layer being built unless they are already implemented in a specific
deployment.

## 9. What the Three Laws Replace

### Replace memory with governed context

Memory asks:

> What has the agent seen before?

Governed context asks:

> What information is authoritative, relevant, current, and authorized for
> this task?

### Replace similarity with meaning

Vector search asks:

> What is nearby in embedding space?

Semantic infrastructure asks:

> What does this entity mean, what is it related to, and which definition
> applies?

### Replace intent inference with action contracts

Prompt instructions ask:

> Will the model remember the rule?

Pragmatic enforcement asks:

> Does the tool call satisfy the rule before execution?

## 10. Failure Modes

| Failure | Broken law |
|---|---|
| Stale context overrides a current decision | Lexicon |
| Similar but incorrect document is retrieved | Lexicon / Syntax |
| Sensitive field is misunderstood | Syntax |
| Prompt injection changes behavior | Syntax / Pragmatics |
| Agent invokes an overly powerful tool | Pragmatics |
| Agent acts with a shared service identity | Lexicon / Pragmatics |
| Policy exists only in unexecutable text | Pragmatics |
| Audit log cannot identify the human | Pragmatics |

## 11. Evaluation and Operations

Semantic systems need more than answer-quality benchmarks.

Measure:

- Retrieval authority
- Context freshness
- Entity-resolution accuracy
- Ontology coverage
- Policy decision accuracy
- Unauthorized-action rejection rate
- Tool-call validity
- Stale-context recurrence
- Audit completeness
- Model-swap consistency
- Time from request to validated action

The production unit is:

```text
Data → Meaning → Context → Model → Tool → Policy → Outcome
```

It is not the model alone.

## 12. The Sovereignty Claim

Sovereignty means more than where the model runs.

A sovereign AI system controls:

- where data lives;
- which model sees it;
- which identity accesses it;
- which meaning is applied;
- which actions are possible;
- which policies are enforced; and
- who can inspect the result.

This distinguishes infrastructure sovereignty from merely hosting an
open-weight model.

## 13. An Implementation Path

1. Inventory agent tools and data sources.
2. Identify stale-context and ambiguous-context failures.
3. Define core entities and organizational vocabulary.
4. Add provenance and freshness metadata.
5. Introduce identity-aware retrieval.
6. Replace broad tools with typed operations.
7. Intercept every tool call before execution.
8. Apply ABAC using user, resource, action, and purpose.
9. Record traces and policy decisions.
10. Evaluate the full context-to-action pipeline.

## 14. Closing Declaration

> We do not need agents that merely remember more.
>
> We need agents that know what information means, why it applies, who may use
> it, and what actions are authorized.
>
> The Lexicon gives the agent governed knowledge.
>
> Syntax gives that knowledge structure.
>
> Pragmatics gives action boundaries.
>
> Together, they form the semantic background for intelligent action.
