# The Context Engineering Manifesto

## Architectural Laws for LLM and NLP Systems Operating on Data

**Author:** Miriah Peterson

**Organization:** Haikei Labs

**Status:** Working draft

**Date:** September 2026

## Abstract

Agentic systems fail in production for reasons we too often attribute to the
model. We say the model hallucinated, retrieved the wrong document,
misunderstood the user, called the wrong tool, or ignored a policy. Sometimes
that is true. Increasingly, however, we blame the model for architectural
decisions we made around it.

Agents operate inside what I call the **Context Void**: an environment where
the information required to understand a task is incomplete, ambiguous, stale,
unauthorized, or missing. We then hand that incomplete picture to a
probabilistic system and ask it to fill in the blanks. It does. That is what we
built it to do.

Some blanks are safe for a model to infer. Some are not. A model can infer that
a frustrated customer wants help; it should not infer which customer record
belongs to them. It can reason about whether a refund seems appropriate; it
should not infer whether the requester is authorized to issue one. It can
propose an operation; it should not decide whether policy permits it.

Reliable agentic systems therefore need more than better prompts, longer
context windows, or elaborate memory systems. They need architecture around
context. This manifesto proposes three laws:

- **Lexicon:** Data must remain attributable, governed, and meaningful.
- **Syntax:** Meaning must be represented through explicit schemas, ontologies,
  and relationships.
- **Pragmatics:** Actions must be expressed as constrained, authorized, and
  auditable operations.

Together, these laws move us from treating context as prompt material toward
treating context as infrastructure. The model may remain probabilistic. The
environment around it cannot remain ambiguous everywhere.

## Context Is Infrastructure

An agent is not intelligent in isolation. Its behavior depends on a larger
system:

```text
Model + Data + Meaning + Identity + Tools + Policy + State
```

Change any one of these and behavior can change without changing a model
weight. Different documents produce different conclusions. A different
identity changes what information should be visible. A different tool surface
changes which actions are possible. A policy change can move the same proposed
operation from valid to forbidden.

That entire environment is context.

Yet many architectures assemble context immediately before inference: retrieve
documents, append history, inject a prompt, describe tools, and send the result
to the model. That approach is insufficient once models interact with
production systems, sensitive data, policy, and real users. Context needs
ownership, lifecycle management, versioning, access control, provenance,
observability, and explicit interfaces between information, meaning, and
action.

The transformer determines what information receives attention; it does not
establish truth, provenance, trust, authorization, or memory. [Our attention
research](../research/attention-is-all-you-need-notes.md) makes this boundary
explicit.

**Context is infrastructure.** Once an agent performs work rather than merely
generating text, that becomes an operational requirement.

## The Context Void

The Context Void exists when the model lacks information necessary to interpret
its environment. It does not mean the model has no data. Often it has too much
data and too little structure.

An agent might receive twenty documents about a customer without knowing which
one is authoritative. It might retrieve a policy without knowing that a newer
policy superseded it. It might see an identifier without knowing whether it
represents a customer, employee, patient, account, or service. It might know
which operation accomplishes a goal without knowing whether the current user
may perform it.

The failure chain is simple:

```text
Missing or ambiguous data
        ↓
Missing or ambiguous meaning
        ↓
Probabilistic inference
        ↓
Unvalidated action
```

Stale context, bad retrieval, entity confusion, prompt injection, and
unauthorized tool execution often begin here. Context assembly is a query over
multiple sources before inference, not merely text placed in a prompt. [Our
context-assembly research](../research/context-assembly-pipeline-patterns.md)
documents this distinction.

The goal is not to eliminate inference. Inference is why models are useful.
The goal is to decide where inference belongs and where deterministic systems
must take over.

## Law One: The Law of Lexicon

> **An agent may only reason from data whose source, ownership, authority, and
> access conditions are known.**

Before an agent can understand information, it needs to know what information
it is looking at. In an enterprise, the same customer may exist in five
systems; a policy may exist as both a PDF and an internal page; a name may
identify an employee in one dataset and a customer in another; and a valid
record may already be six months out of date.

Retrieval must answer more than “What is relevant?” It must also answer:

- Where did this information come from?
- Who owns it?
- Is it authoritative and current?
- What supersedes it?
- Who may access it, and for what purpose?
- Where may it legally be processed?

Those properties form the **Lexicon**: the governed information environment in
which an agent operates. It includes entities, definitions, relationships,
provenance, ownership, sensitivity, versions, retention, access attributes,
authority boundaries, and freshness or supersession information.

The Lexicon does not require copying every source into a central AI memory
store. That instinct recreates synchronization, permission, retention, and
data-residency problems under a new name. Customer records can remain in the
CRM, employee information in HR, telemetry in observability systems, and
intellectual property in its governed repository. The context layer need not
own everything. It must understand how sources relate, under whose authority
they exist, and under what conditions they may be used.

The objective is **reliable access to governed context**, not consolidation for
its own sake. Memory asks, “What has this agent seen before?” Governed context
asks, “What information is authoritative, current, relevant, and authorized
for this task?” Those are different questions. [Our research on memory and
knowledge graphs](../research/knowledge-graphs/05-context-engineering-connections.md)
supports treating memory as structured retrieval over governed state.

## Law Two: The Law of Syntax

> **Data without semantics is information without the narrative. Syntax is how
> an organization operationally conveys why data matters.**

Once we know where information came from, we still need to know what it means.
The string `123-45-6789` says little about how an organization should treat
it. Its operational meaning is richer:

```text
SSN
 ├── entity: Person
 ├── class: RegulatedIdentifier
 ├── display: MaskedOnly
 ├── purpose: IdentityVerification
 └── access: ABACRequired
```

I use **Syntax** here to mean structure, not punctuation. It is the mechanism
through which organizational meaning becomes legible to a model. It includes
ontologies, schemas, entity resolution, standard dataset mappings, relationship
graphs, typed tool parameters, domain vocabulary, and context-assembly rules.

Relationships matter because entities do not have one universal meaning. A
customer means something different to finance than to sales. A server means
something different to security than to platform engineering. A patient
identifier carries different implications in clinical care, billing, identity
verification, and analytics.

Ontologies make those distinctions explicit instead of forcing a model to
infer them from whichever documents retrieval returned. [Our semantic-contract
research](../research/semantic-contracts.md) treats ontologies, identifiers,
relationships, and constraints as components of an enforceable context
contract.

### Similarity Is Not Semantics

Vector search answers a useful question: “What looks similar to this request?”
It does not establish authority, identity, organizational definition, or
whether a relationship is causal, hierarchical, contractual, temporal, or
merely correlated.

Embedding proximity is not organizational truth. Vector search is an access
mechanism, not a complete semantic architecture. [Our hybrid-retrieval
research](../research/hybrid-retrieval-architectures.md) shows why lexical,
structured, graph, and vector retrieval can complement one another.

## Law Three: The Law of Pragmatics

> **Agents hallucinate constantly; sometimes those hallucinations happen to
> accomplish what we asked of them.** — Nick Humrich

A model continuously generates predictions. Once an agent calls tools, those
predictions become operational.

```text
Send this message.
Approve this request.
Update this record.
Deploy this version.
Delete this account.
```

These are operational speech acts. A call such as
`update_patient_record(patient_id, field, value)` is not merely text. It is a
proposed change to the world.

Pragmatics defines the architecture around action: typed tool definitions,
explicit parameters, purpose binding, user and workload identity, ABAC,
business-rule validation, pre-execution interception, audit records, and safe
failure behavior. Computational pragmatics studies the relationship between
utterances and context; [our research notes](../research/computational-pragmatics-notes.md)
connect that relationship to beliefs, goals, authorization, and action.

The model is allowed to propose. Infrastructure decides whether the proposal
becomes reality.

**The model can infer the work. It cannot infer permission.**

## The Determinism Boundary

Large language models are probabilistic systems. Their flexibility makes them
useful for ambiguous human tasks. It also marks the boundary of what they
should not decide.

Authorization cannot mean that the model thinks a user probably has permission.
Policy cannot mean that the model read a guideline and seems to remember it.
Compliance cannot mean that a requirement was placed in a system prompt.

Security and business policy must be executable:

```text
Given:
  identity, purpose, resource, action, policy

Return:
  allow or deny
```

This is the **Determinism Boundary**. On one side is probabilistic reasoning;
on the other is deterministic enforcement.

> **Never ask a probabilistic model to make a deterministic authorization
> decision.**

The model can determine what action may accomplish a task. Infrastructure must
determine whether that action is allowed.

## The Semantic Background Architecture

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
meaning and boundaries. It establishes identity, resolves relevant context,
applies organizational meaning, exposes permitted resources, intercepts
proposed actions, evaluates policy, and records what happened.

Intelligence happens in the foreground. Context makes that intelligence useful.

## How the Three Laws Change Agent Architecture

### Replace memory with governed context

Persistence without authority creates confident stale context. A conversation
from three months ago should not defeat a policy updated yesterday, and
information seen under one authorization context should not automatically carry
into another.

The question is not how much the agent remembers. It is whether the information
being used **applies now**.

### Replace similarity with meaning

Similarity is useful because language is messy. But it is an access mechanism,
not an architectural definition of meaning. Semantic infrastructure tells us
what an entity is, what it relates to, which definition applies, and why it
matters to the task.

### Replace intent inference with action contracts

Prompt instructions ask whether the model will remember the rule. Pragmatic
enforcement asks whether the proposed operation satisfies the rule before
execution. For consequential operations, production systems should prefer
architectural enforcement whenever a rule can be made deterministic.

## Failure Modes Through the Three Laws

| Failure | Broken law |
|---|---|
| Stale context overrides a current decision | Lexicon |
| Similar but incorrect document is retrieved | Lexicon / Syntax |
| Sensitive field is misunderstood | Syntax |
| Prompt injection changes behavior | Syntax / Pragmatics |
| Agent invokes an overly powerful tool | Pragmatics |
| Agent acts with a shared service identity | Lexicon / Pragmatics |
| Policy exists only in unexecutable prose | Pragmatics |
| Audit log cannot identify the human | Pragmatics |

Different failures require different fixes. If the problem is Lexicon, a
better tool description will not solve it. If the problem is Syntax, increasing
`top_k` probably will not solve it. If the problem is Pragmatics, another
paragraph in the system prompt definitely will not solve it.

The laws tell us where the missing architecture belongs.

## Evaluation and Operations

If agents operate through context infrastructure, the system—not only the
model—becomes the unit of evaluation. Retrieval authority matters because a
correct fact can come from the wrong source. Freshness matters because
yesterday’s correct answer may be wrong today. Entity resolution matters
because good reasoning fails when attached to the wrong entity. Policy and
rejection metrics matter because a correct model output can still produce an
unacceptable outcome.

Measure:

- retrieval authority;
- context freshness;
- entity-resolution accuracy;
- ontology coverage;
- policy-decision accuracy;
- unauthorized-action rejection rate;
- tool-call validity;
- stale-context recurrence;
- audit completeness;
- model-swap consistency; and
- time from request to validated action.

Evaluation must cover retrieval, tools, workflows, and outcomes, not merely
model text. [Our evaluation research](../research/llms-in-production/chapter-7.md)
supports using a harness to evaluate the system around the model.

The production unit is:

```text
Data → Meaning → Context → Model → Tool → Policy → Outcome
```

The model occupies one position in that chain. An important position, but one
position nevertheless.

## Sovereignty Means Controlling Context

Sovereignty means more than running an open-weight model on hardware you
control. You can self-host a model and still send sensitive information into
systems you do not control, operate every agent under a shared service account,
or lose the ability to explain why an action was allowed.

A sovereign AI system controls:

- where data lives;
- which model can see it;
- which identity accesses it;
- which meaning is applied;
- which actions are possible;
- which policies are enforced; and
- who can inspect the result.

Sovereignty is control over the full context-to-action path. Context
infrastructure moves enforcement outside the model, making model choice more
replaceable while policy, identity, data, and organizational meaning remain
under organizational control.

## Implementation Path

Do not rebuild the entire data stack before deploying an agent. Start where the
agent touches reality:

1. Inventory the agent’s tools and data sources.
2. Identify stale-context and ambiguous-context failures.
3. Define the core entities and vocabulary for the task domain.
4. Add provenance and freshness metadata.
5. Introduce identity-aware retrieval.
6. Replace broad tools with typed operations.
7. Intercept every consequential tool call before execution.
8. Apply ABAC using user, resource, action, and purpose.
9. Record traces and policy decisions.
10. Evaluate the full context-to-action pipeline.

You do not need an ontology for the universe. You need enough semantics for the
part of the organization the agent is expected to understand.

## Conclusion

We do not need agents that merely remember more. We need agents that understand
which information is authoritative, receive the organizational meaning around
that information, and operate inside systems that know who is asking, why they
are asking, what resources are involved, and which actions are permitted.

We need to stop expecting models to infer architecture we never built.

The Lexicon gives the agent governed knowledge. Syntax gives that knowledge
structure and meaning. Pragmatics gives action boundaries. The Determinism
Boundary separates what the model may reason about from what infrastructure
must enforce. Together, they create the **Semantic Background** required for
intelligent action.

Context engineering is not bigger prompts, longer memory, or better tricks for
stuffing documents into a context window. It is the discipline of building the
data, semantic, identity, and policy infrastructure that makes probabilistic
reasoning reliable enough to operate against real systems.

The model can remain probabilistic.

The world around it cannot be ambiguous everywhere.

That is the job of context engineering.

## Applying the Architecture: Haikei

The architecture can grow incrementally. Identity and attribution can come
first, followed by policy enforcement, then increasingly sophisticated context
resolution and semantic representation.

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

The principal, ABAC, proxy, and audit layers represent concrete portions of
this model. Ontology and semantic-injection capabilities represent the next
layer and should be described that way until a specific deployment implements
them directly.

A manifesto should tell us where we are going. It should not require us to
pretend we have already arrived.

## Works Cited

The manifesto draws on the following books, papers, and standards. The linked
repository notes record how each source informed the argument.

### Books and chapters

- Brousseau, Christopher, and Matthew Sharp. *LLMs in Production: From
  Language Models to Successful Products*. Manning, 2025. [Publisher
  page](https://www.manning.com/books/llms-in-production). See the
  [repository notes](../research/llms-in-production/notes.md), especially the
  material on evaluation and production reliability.
- Jurafsky, Dan. “Computational Pragmatics.” In *The Handbook of Pragmatics*.
  Blackwell. [Chapter PDF](https://web.stanford.edu/~jurafsky/prag.pdf). See
  the [repository notes](../research/computational-pragmatics-notes.md).
- Kejriwal, Mayank, Craig A. Knoblock, and Pedro Szekely. *Knowledge Graphs:
  Fundamentals, Techniques, and Applications*. MIT Press, 2021.
  [Publisher page](https://mitpress.mit.edu/9780262045094/knowledge-graphs/).
  See the [knowledge-graph bibliography](../research/knowledge-graphs/bibliography.md).
- Kleppmann, Martin. *Designing Data-Intensive Applications*. O’Reilly Media,
  2017. [Publisher page](https://www.oreilly.com/library/view/designing-data-intensive-applications/9781491903063/).

### Papers

- Vaswani, Ashish, et al. “Attention Is All You Need.” 2017.
  [arXiv:1706.03762](https://arxiv.org/abs/1706.03762). See the
  [repository notes](../research/attention-is-all-you-need-notes.md).
- Brown, Tom B., et al. “Language Models are Few-Shot Learners.” 2020.
  [arXiv:2005.14165](https://arxiv.org/abs/2005.14165).
- Schick, Timo, et al. “Toolformer: Language Models Can Teach Themselves to
  Use Tools.” 2023. [arXiv:2302.04761](https://arxiv.org/abs/2302.04761).
  See the [repository notes](../research/toolformer-notes.md).
- Park, Joon Sung, et al. “Generative Agents: Interactive Simulacra of Human
  Behavior.” 2023. [arXiv:2304.03442](https://arxiv.org/abs/2304.03442).
  See the [memory research notes](../research/episodic-periodic-memory.md).
- Allemang, Dean, and Juan F. Sequeda. “Ontologies to the Rescue?” 2024.
  [arXiv:2405.11706](https://arxiv.org/abs/2405.11706). See the
  [knowledge-graph evaluation notes](../research/reading-list.md#11-knowledge-graph-quality-metrics).
- Zaveri, Amrapali, et al. “Quality Assessment for Linked Data: A Survey.”
  *Semantic Web*, 2016. [DOI:10.3233/SW-150175](https://doi.org/10.3233/sw-150175).
- Hu, Edward J., et al. “LoRA: Low-Rank Adaptation of Large Language Models.”
  2021. [arXiv:2106.09685](https://arxiv.org/abs/2106.09685).

### Standards and technical specifications

- W3C. “RDF 1.1 Concepts and Abstract Syntax.”
  [W3C Recommendation](https://www.w3.org/TR/rdf11-concepts/).
- W3C. “OWL 2 Web Ontology Language Document Overview.”
  [W3C Recommendation](https://www.w3.org/TR/owl2-overview/).
- W3C. “SPARQL 1.1 Query Language.”
  [W3C Recommendation](https://www.w3.org/TR/sparql11-query/).
- NIST. *Attribute Considerations for Access Control Systems* (SP 800-162).
  [NIST publication](https://csrc.nist.gov/pubs/sp/800/162/upd2/final).

The repository’s [evidence ledger](../research/_evidence-ledger.md) maps these
sources to the claims and chapters they support.
