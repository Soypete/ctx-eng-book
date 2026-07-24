# Semantic Contracts: The Missing Agent Interface for Data Mesh

## Research Thesis

Data contracts make decentralized data products dependable. Semantic contracts make those data products interpretable and actionable by AI agents. Together, they supply the missing interface that can make data mesh valuable for agentic systems.

---

## 1. The Original Promise of Data Mesh

Data mesh is a decentralized approach to analytical data management. Rather than treating data as a centralized exhaust stream collected by a specialized platform team, it assigns data-product ownership to the domains that understand and produce the data.

Data mesh is a socio-technical architecture:

- Domain-oriented ownership
- Data as a product
- Self-service infrastructure
- Federated computational governance

This means data mesh is already an agreement problem:

- domains retain control of their data;
- consumers must be able to discover and compose products;
- global governance must coexist with local autonomy;
- interfaces must survive independent domain evolution.

---

## 2. What Data Contracts Contribute

A data contract is an explicit agreement between a data producer and its consumers:

- schema and types;
- required and optional fields;
- compatibility rules;
- freshness and availability;
- ownership;
- data quality thresholds;
- lineage;
- usage limitations;
- versioning and deprecation.

Data contracts provide the stability required for autonomous consumption:

- Does the field exist?
- What type does it have?
- How fresh is it?
- Who owns it?
- Is the producer meeting the promised service level?
- Will a change break known consumers?

These guarantees make the distributed product dependable. They do not necessarily make it understandable.

---

## 3. The Unresolved Semantic Gap

An agent may receive perfectly valid data and still use it incorrectly.

Consider a field named `revenue`. A data contract may guarantee:

```
name: revenue
type: decimal
nullable: false
currency: USD
freshness: 24h
```

But an agent still may not know whether revenue means:

- invoiced revenue
- collected revenue
- recognized revenue
- gross revenue
- net revenue
- annual recurring revenue
- revenue attributed to the sale date
- revenue attributed to the recognition period

This is the **semantic gap**: the distance between the shape and reliability of information and the meaning a consumer is entitled to assign to it.

---

## 4. What a Semantic Contract Adds

A semantic contract is a versioned, machine-readable agreement about meaning, valid interpretation, policy, behavior, and observable guarantees.

### Meaning

- What real-world concept does this field or entity represent?
- Which ontology or vocabulary defines it?
- Which aliases are equivalent?
- Which superficially similar concepts are explicitly not equivalent?

### Context

- In which domain, workflow, jurisdiction, or time period is the definition valid?
- At what granularity may the data be interpreted?
- Does the value describe current state, historical state, an observation, or an estimate?

### Derivation and Evidence

- How was the value produced?
- Which source is authoritative?
- Which transformations and inference rules were applied?
- What provenance must accompany a claim?

### Permission and Purpose

- Who may access the data?
- For which purpose?
- Which uses, inferences, disclosures, or persisted outputs are prohibited?

### Behavior

- What preconditions must hold before a tool or data operation is used?
- What state changes after execution?
- Which invariants must remain true?
- Is the operation idempotent, reversible, or compensatable?

### Verification

- How can the interpretation or action be validated?
- What constitutes a contract violation?
- What evidence should be recorded for audit and repair?

---

## 5. The Three Contract Layers

Consider a user asking: *"Which customers should we reach out to before renewal?"*

### Data Contract

The data contract establishes that:

- `customer_id` is stable;
- renewal dates are current;
- account-health data refreshes every six hours;
- required fields are present;
- the owning domains meet their service guarantees.

It says the data is **dependable**.

### Semantic Contract

The semantic contract establishes that:

- `customer` means a contracted account, not an individual user;
- `renewal_date` means the contractual renewal date;
- `health_score` is predictive, not authoritative;
- churn risk cannot be interpreted as dissatisfaction;
- billing is authoritative for renewal status.

It says what the data **means** and what conclusions it **supports**.

### Pragmatic Contract

The pragmatic layer asks:

- Who is asking?
- In what role?
- What does "reach out" mean here?
- Is this analysis, recommendation, or authorization to contact customers?
- Which communication channels are appropriate?
- Should the agent merely identify accounts, draft outreach, or actually send it?
- What consequences will the action have?

It says what the user is trying to accomplish **with that meaning in this context**. That is **computational pragmatics**.

---

## 6. Why Agents Require Both Contracts

| Contract Layer | Primary Guarantee | Prevents |
|---------------|-------------------|----------|
| Data contract | The product is structurally and operationally dependable | Schema breaks, missing fields, stale data, undocumented ownership |
| Semantic contract | The product is interpreted and acted upon according to governed meaning | Concept confusion, invalid inference, misuse, unauthorized action |
| Combined contract | The agent can reliably consume, reason over, and act through the product | Technically valid but operationally wrong agent behavior |

> A data contract tells an agent what it can read. A semantic contract tells it what the data means, what it may conclude, and what it is allowed to do next.

---

## 7. Evidence from Current Agent Research

### 7.1 Semantic Contracts as Executable Intent

STEPS introduces a semantic contract as an executable intermediate interface between ambiguous natural-language requests and constrained system execution. An LLM interprets the request, but a structured contract captures service preferences, fulfillment bounds, confidence, and uncertainty before the scheduler acts.

### 7.2 Ontologies Compiled into Agent Tools

Zhou et al. propose ontology-to-tools compilation, in which formal ontological constraints are transformed into executable tool interfaces for LLM agents. Instead of asking the model to remember semantic rules from a prompt, the system constrains knowledge-graph creation and modification through ontology-aware tools.

### 7.3 Tool Schemas Are Not Behavioral Contracts

Contract2Tool identifies a central limitation of ordinary agent tool schemas: they describe how to call a tool, but not when the tool is causally appropriate or what state it produces. The proposed contracts include preconditions, effects, risk, and cost.

### 7.4 Semantic Grounding Reduces Tool Hallucination

Chethan formalizes a "semantic training gap" between an LLM's statistical familiarity with domain vocabulary and the operational relationships that define meaning inside an organization. The proposed ontology-grounded tool architecture uses a three-operation interface—resolve, contextualize, and annotate—and reports a reduction in domain-identifier hallucinations.

### 7.5 Semantic Metadata Improves Actionable Retrieval

Chen et al. compare open-web agent retrieval with retrieval over a large corpus enriched with structured semantic metadata. Their results show that semantic metadata substantially improves precision for finding machine-actionable, FAIR-compliant datasets.

### 7.6 Ontologies Constrain Both Context and Output

Tuan proposes a three-layer ontological framework for role, domain, and interaction semantics. The architecture uses symbolic knowledge to constrain context assembly, tool discovery, governance thresholds, and potentially output validation.

### 7.7 Semantic Contracts for Agent Communication

Research tracing the relationship between semantic-web systems and agentic AI argues that modern agent interoperability is revisiting longstanding problems from multi-agent systems and semantic-web research: shared vocabularies, capability descriptions, semantic agreement, and protocol interpretation.

---

## 8. Semantics Is Not Enough

Semantics can tell the agent:

> `send_customer_email` sends a message to the primary account contact.

Pragmatics determines:

> Is sending an email the appropriate interpretation of "reach out" in this conversation?

The user might have meant:

- prepare a list for the account managers;
- draft an outreach campaign;
- create CRM tasks;
- send the messages now;
- investigate why the accounts are at risk.

All are semantically related to "reach out." Only one may be pragmatically appropriate.

This is why an ontology or semantic layer cannot, by itself, make an agent reliable. It represents the available concepts and relationships, but it does not completely resolve:

- speaker intent;
- institutional role;
- conversational history;
- organizational norms;
- implied authority;
- acceptable risk;
- expected degree of initiative;
- the distinction between describing and doing.

That is the gap between **knowing what an action means** and **knowing whether this is the right moment to perform it**.

---

## 9. Where Computational Pragmatics Lives

Pragmatics is distributed across the agent architecture:

### 1. The User Interaction

- deixis: "this customer," "that report," "next quarter";
- implicature: what the user expects without stating it;
- role and relationship;
- prior commitments;
- level of urgency;
- whether a statement is a question, request, command, or exploration.

### 2. The Agent's Role

The same request means different things to different agents. "Fix the invoice" might mean:

- explain the discrepancy to an analyst agent;
- create a correction proposal for a finance agent;
- modify a pending invoice for a billing operator;
- refuse the request for a read-only reporting agent.

### 3. The Semantic Contract

- authoritative meanings;
- domain boundaries;
- purpose restrictions;
- valid inferences;
- obligations;
- prohibited uses;
- behavioral preconditions.

### 4. The Available Tools

The tool set defines the agent's actionable vocabulary. A tool is an **affordance**: a representation of something the environment permits the agent to do.

### 5. The Orchestration and Policy Layer

- which tools are exposed for the current task;
- which require confirmation;
- which can be chained;
- which data products can supply arguments;
- which postconditions must be checked;
- when the agent must stop and ask a person.

---

## 10. Tool Calling Is Computational Pragmatics Made Operational

A tool call has a structure remarkably similar to a speech act.

When someone says "I approve the refund," the utterance does more than describe approval—it **performs** approval under the right institutional conditions.

Likewise, when an agent calls:

```json
{
  "tool": "approve_refund",
  "arguments": { "refund_id": "R-1049" }
}
```

the output is not merely language about a refund. It changes institutional state.

The call is valid only if pragmatic conditions hold:

- the caller has the correct authority;
- the referenced refund exists;
- the statement is intended as approval rather than discussion;
- the request occurs within the proper workflow;
- prerequisite reviews are complete;
- the consequences are understood;
- the operation is allowed at this point in the conversation.

These are essentially the **felicity conditions** of the tool call.

> **Tool calls are where computational pragmatics becomes action.**

---

## 11. Tool Curation Is Action-Space Governance

Tool curation is often framed as:

- reducing tokens;
- helping the model select accurately;
- avoiding confusion;
- improving tool-routing benchmarks.

But the deeper reason is:

> **The tools you expose define the set of pragmatic actions the agent can consider.**

Suppose the user says: "There's a duplicate customer record. Take care of it."

You might expose:

- `inspect_customer_records`
- `compare_customer_entities`
- `propose_entity_merge`

but withhold:

- `merge_customer_entities`
- `delete_customer`
- `rewrite_billing_history`

The agent can investigate and prepare a repair. But the available affordances prevent it from interpreting "take care of it" as permission to perform an irreversible merge.

That is not merely guardrailing after reasoning. It shapes the reasoning itself.

---

## 12. A Tool Description Should Be a Pragmatic Contract

Today, tools are often described like this:

```yaml
name: update_customer
description: Updates a customer.
```

That is nearly useless.

A pragmatic tool contract would include:

```yaml
name: propose_customer_merge

purpose:
  description: >
    Prepare a non-destructive proposal to merge records believed
    to represent the same legal customer.

use_when:
  - two records have been resolved as probable duplicates
  - the user requested investigation or remediation
  - conflicting identifiers have been surfaced

do_not_use_when:
  - records represent separate subsidiaries
  - either customer has an unresolved legal hold
  - confidence is below 0.95

caller_roles:
  - customer_data_steward_agent
  - support_operations_agent

preconditions:
  - both records exist
  - identity evidence has been retrieved
  - authoritative domains have been checked

effects:
  - creates a review proposal
  - does not modify customer records

follow_up_capabilities:
  - submit_merge_for_review

evidence_required:
  - matching identifiers
  - conflicting attributes
  - source provenance
  - confidence score

failure_modes:
  unresolved_identity:
    next_action: request_more_evidence
  policy_conflict:
    next_action: escalate_to_data_steward
```

This is a machine-readable account of the tool's **pragmatic appropriateness**.

---

## 13. Semantic Contracts Constrain Claims; Pragmatic Contracts Constrain Acts

### Semantic Contract

Controls what the agent can safely say:

- this metric means recognized revenue;
- this source is authoritative;
- these entities are equivalent;
- this inference is supported;
- this confidence level must be disclosed.

### Pragmatic Contract

Controls what the agent can appropriately do:

- retrieve this information for this purpose;
- propose this change but do not execute it;
- request clarification because intent is ambiguous;
- obtain confirmation before an irreversible operation;
- choose this tool because it matches the user's institutional goal;
- stop because no available action satisfies the request safely.

> **Semantic contracts govern warranted interpretation. Pragmatic contracts govern warranted action.**

---

## 14. Why the Semantic Layer Cannot Simply Be Placed in the Prompt

Current agent architecture treats business meaning as context text:

```
Here are our definitions.
Here are 300 tool descriptions.
Here are the governance policies.
Now decide what to do.
```

That asks the model to reproduce the entire institution's semantics and pragmatics probabilistically on every invocation.

Instead, the contract layer should compile the situation into a bounded interface:

```
User intent
      +
Role and authority
      +
Relevant semantic contracts
      +
Current workflow state
      ↓
Pragmatic capability resolver
      ↓
Curated tools and scoped data
      ↓
Agent decision
      ↓
Precondition and postcondition validation
```

The agent should not see every tool and then be told not to misuse most of them. It should receive the tools that are pragmatically valid in the current context.

---

## 15. The Agent-Ready Data Mesh

An agent-ready data mesh needs four cooperating layers:

### Layer 1: Domain-Owned Data Products

- billing owns invoices and recognized financial events;
- identity owns users, credentials, and account membership;
- product owns feature use and interaction events;
- support owns cases and service interactions.

### Layer 2: Data Contracts

- structure;
- service levels;
- quality;
- ownership;
- versioning;
- access endpoints.

### Layer 3: Semantic Contracts

- domain concepts;
- identity and equivalence;
- temporal and aggregation semantics;
- authority and provenance;
- purpose-bound access;
- permissible inference;
- interaction preconditions and effects;
- validation and audit.

### Layer 4: Agent Execution Infrastructure

- context assembly;
- product and tool discovery;
- scoped credentials;
- query planning;
- ontology-aware tools;
- policy checks;
- output validation;
- traces and contract-violation events.

---

## 16. Semantic Contracts as Context Engineering

Context engineering has two connected responsibilities:

1. **Data engineering**: retrieve the correct information with identity, provenance, permissions, and temporal integrity.
2. **Pragmatics**: determine how that information should be interpreted and used in the current interaction.

A semantic contract joins them:

```
Data engineering
  supplies authoritative information
              ↓
Semantic contract
  states valid interpretation and action
              ↓
Pragmatics
  applies the contract to current intent
```

The contract should not be dumped wholesale into the prompt. It should govern how context and tools are assembled:

- exclude products whose contracts do not match the purpose;
- resolve concepts before querying;
- expose only permitted operations;
- carry definitions and provenance into the answer;
- validate conclusions against declared semantics;
- prevent actions whose preconditions are unsatisfied.

The semantic contract is therefore a **context boundary**, not merely additional context.

---

## 17. Core Assertions

### Assertion 1

Data mesh is not inherently agent-ready. Decentralized ownership creates valuable domain boundaries but also multiplies vocabularies, identities, policies, and local assumptions.

### Assertion 2

Data contracts are necessary but primarily guarantee product reliability. They make decentralized products testable, stable, and discoverable.

### Assertion 3

Semantic contracts guarantee safe interpretation and action. They encode domain meaning, authority, provenance, purpose, behavioral conditions, and verifiable outcomes.

### Assertion 4

The combination is the missing agent interface for data mesh. Agents require a machine-executable agreement that spans both the technical product interface and the meaning of its use.

### Assertion 5

Semantic contracts should be compiled into context and tools, not pasted into prompts. Their purpose is to constrain retrieval, capability exposure, authorization, reasoning, and validation.

### Assertion 6

Knowledge graphs and ontologies are components, not substitutes. They provide identifiers, vocabulary, relationships, and constraints. The contract adds obligations, guarantees, purpose, and behavior.

### Assertion 7

Federated semantics preserve data-mesh autonomy. The goal is not one universal ontology. The goal is explicit, enforceable agreements at domain boundaries.

### Assertion 8

Contract violations provide the error model for self-healing agents. An agent can only repair itself reliably when it knows which invariant failed.

---

## 18. Proposed Concise Definition

A **semantic contract** is a versioned, machine-executable agreement that defines:

- what information means;
- which interpretations and inferences are valid;
- who may use it for which purposes;
- what behavioral guarantees must hold when an agent acts upon it.

An **agent-ready data product** pairs:

- a **data contract** governing structure and service reliability;
- a **semantic contract** governing meaning, policy, provenance, and behavior.

---

## 19. The Missing Piece: Open Pragmatic Interface

There are two closely related missing artifacts:

### Open Semantic Interface

Exposes what the organization knows:

- concepts;
- entities;
- relationships;
- metrics;
- definitions;
- mappings;
- provenance;
- authority;
- data-product locations.

> What does the organization mean by this?

### Open Pragmatic Interface

Exposes what an agent may appropriately do with that meaning:

- available capabilities;
- role requirements;
- allowed purposes;
- action preconditions;
- side effects;
- confirmation requirements;
- risk and cost;
- expected postconditions;
- evidence requirements;
- escalation rules;
- repair options.

> Given who is asking, what they are trying to accomplish, and the current state, what action is appropriate?

---

## 20. Architecture Summary

```
Data mesh
  distributes domain-owned information

Data contracts
  make products reliable

Semantic contracts
  make their meaning explicit

Open semantic interchange
  makes that meaning portable

Computational pragmatics
  applies meaning to situated intent

Tool curation
  turns that interpretation into a bounded action space

Tool calls
  perform the selected institutional action
```

---

## 21. Strongest Assertions

> **Data contracts are the interfaces of a data mesh. Semantic contracts are the interfaces of an organizational knowledge system. Curated tools are the pragmatic interface through which agents are permitted to act on that knowledge.**

> **A semantic layer tells an agent what the business means. A curated tool layer tells it what the business permits this agent to do right now.**

> **Agent tool curation is computational pragmatics implemented as infrastructure.**

---

## 22. Product Data vs Personal Data: The Contract Boundary

A critical dimension that semantic contracts must address: **product data** and **personal data** require fundamentally different contract structures.

### Product Data (B2B / Operational)

Product data describes business entities, transactions, and operations:

- invoices, orders, inventory, pricing
- company hierarchies, organizational structure
- internal processes, workflows, policies

For product data, semantic contracts govern:

- **What it means**: invoice status = recognized vs pending vs void
- **What conclusions it supports**: can be used for financial reporting
- **What actions it enables**: create_credit_memo requires approved invoice
- **Who owns it**: finance domain
- **Authority**: billing system is authoritative for invoice status

### Personal Data (PII / User Data)

Personal data describes individuals:

- names, emails, addresses, phone numbers
- behavioral data, preferences, history
- identity credentials, authentication state

For personal data, semantic contracts additionally govern:

- **Consent boundaries**: what purposes is this data authorized for?
- **Retention limits**: how long may it be persisted?
- **Access restrictions**: which roles can see which fields?
- **Processing lawful basis**: consent, legitimate interest, contract performance
- **Cross-border rules**: which jurisdictions apply?
- **Individual rights**: access, deletion, portability requests
- **Purpose binding**: analytics use ≠ marketing use ≠ support use

### The Privacy Contract

Personal data requires an additional **privacy contract** layer:

```yaml
name: customer_contact_info
data_type: personal

consent_model:
  - marketing_communications: explicit_consent
  - product_updates: legitimate_interest
  - support_interactions: contract_performance

retention:
  - marketing_list: 2_years_from_last_interaction
  - support_history: 7_years_from_last_case
  - authentication_logs: 90_days

access_control:
  - support_agent: read_name_email_phone
  - marketing_system: read_email_preferences_only
  - analytics: read_aggregated_only

purpose_binding:
  - do_not_use_for: employee_performance_review
  - do_not_use_for: credit_decisions
  - requires_minimization: true

individual_rights:
  - deletion_supported: true
  - portability_supported: true
  - objection_mechanism: available
```

---

## 23. The What vs How Architecture

This is the core architectural insight:

### The Agent Workflow = THE WHAT

The user or agent expresses **intent**:

> "Which customers should we reach out to before renewal?"
> "There's a duplicate customer record. Take care of it."
> "Generate a report on Q3 revenue by region."

This is the **workflow** — what the user wants to accomplish.

### The Contract + Tools = THE HOW

The system responds with **capability**:

- What information is available and reliable? (data contract)
- What does the information mean? (semantic contract)
- When, why, and by whom may it be used? (pragmatic contract)
- Which executable actions realize that use? (tool interface)

### The Contract Becomes the Automation Framework

The semantic contract is not metadata. It is **executable infrastructure**:

```
User Intent (WHAT)
     ↓
Contract Resolution (HOW)
     ↓
Curated Tool Selection (HOW)
     ↓
Validated Execution (HOW)
     ↓
Contract-Aware Response (HOW)
```

The contract tells you:

1. **What you can do** — available actions filtered by purpose, role, consent
2. **How you can do it** — preconditions, required arguments, sequencing
3. **What you cannot do** — prohibited purposes, missing preconditions, policy conflicts
4. **What evidence you must produce** — provenance, audit trails, verification

### Example: "Take care of the duplicate"

| Layer | Answer |
|-------|--------|
| **What** (workflow) | Resolve duplicate customer record |
| **Data contract** | customer records exist, have IDs, have matching evidence |
| **Semantic contract** | "customer" means contracted account, "duplicate" means probable match at 0.95+ confidence |
| **Pragmatic contract** | merge requires steward role, confirmed authority, no legal holds |
| **Tool (how)** | `propose_customer_merge` — non-destructive, requires approval |

The agent doesn't guess. The contract answers what is possible and how to achieve it safely.

---

## 24. Contracts as Governance Infrastructure

This transforms how we think about automation:

| Traditional Automation | Contract-Guided Automation |
|------------------------|---------------------------|
| IFTTT / workflow rules | Semantic + pragmatic constraints |
| Hard-coded permissions | Role-based contract resolution |
| Static tool descriptions | Dynamic pragmatic contracts |
| Batch process governance | Real-time contract validation |
| Audit after the fact | Contract-aware trace generation |

The semantic contract **is** the automation framework. It encodes:

- **Authority**: who can do what
- **Preconditions**: what must be true before action
- **Effects**: what changes after action
- **Invariants**: what must remain true
- **Evidence**: what must be recorded
- **Repair**: what to do when contracts fail

This is the infrastructure layer that makes agents safe to deploy in enterprise environments.

---

## References

### Primary Sources

- Goedegebuure et al. — Data mesh academic review (decentralized, domain-driven data architecture)
- Wider et al. — Self-service infrastructure and federated governance
- Chen et al. — Semantic metadata improves agentic retrieval precision
- STEPS — Semantic contracts as executable intermediate interface
- Zhou et al. — Ontology-to-tools compilation
- Contract2Tool — Tool schemas are not behavioral contracts
- Chethan — Semantic training gap and ontology-grounded tools
- Tuan — Three-layer ontological framework (role, domain, interaction)
- VASO — Semantic contracts for physical-agent skills

### Related Research

- "A Survey of LLM-Driven AI Agent Communication" — arXiv:2506.19676
- "Agent-OM: Leveraging LLM Agents for Ontology Matching" — arXiv:2312.00326
- "The Ontology for Agents, Systems and Integration of Services" — arXiv:2306.10061
- "Deontic Policies for Runtime Governance of Agentic AI" — arXiv:2606.19464
- "Scalable LLM Agent Tool Access in the Cloud" — arXiv:2607.15593

### Standards and Specifications

- Open Semantic Interchange (OSI) — now Apache Ossie (incubating)
- MCP (Model Context Protocol) — data-product access
- SKOS — Simple Knowledge Organization System
- OWL — Web Ontology Language
- RDF — Resource Description Framework

---

## Cross-References

- **Data Mesh**: research/knowledge-graphs/05-context-engineering-connections.md
- **Context Engineering Thesis**: research/five-pillars-outline.md
- **Entity Resolution**: research/knowledge-graphs/8-instancematching.md
- **Evaluation Metrics**: research/kg-quality-metrics-notes.md
- **Evidence Ledger**: research/_evidence-ledger.md
- **Knowledge Graphs**: research/knowledge-graphs/01-overview-and-big-ideas.md