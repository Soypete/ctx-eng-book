# Episodic, Semantic, Procedural, and Periodic Memory in Agent Systems

## Working Thesis

In agent systems, memory should not be treated as something the model inherently possesses.

Memory is a governed system for:

- writing persistent state;
- structuring that state;
- consolidating it;
- retrieving it;
- authorizing access to it;
- validating how it is used;
- expiring, superseding, or deleting it.

---

## The Core Memory Taxonomy

| Memory Type | Core Question | Typical Representation |
|-------------|--------------|----------------------|
| **Episodic** | What happened? | Time-bound event or trajectory |
| **Semantic** | What is true? | Facts, entities, concepts, relationships |
| **Procedural** | How should this be done? | Skills, policies, workflows, strategies |
| **Periodic** | — | Not a peer-level type; a maintenance policy |

**Key distinction**: There's no actual "memory" in transformers. There's only retrieval from external systems. This is the foundational thesis of ch06.

---

## 1. Episodic Memory

### Cognitive-Science Origin

Endel Tulving introduced the distinction between episodic and semantic memory in 1972.

- **Episodic memory**: specific experienced events situated in time and context
- **Semantic memory**: generalized facts and conceptual knowledge

For a person:
- "I met Dana during Tuesday's architecture review, and she rejected the migration because the authorization model was incomplete." → episodic
- "Dana owns the authorization service." → semantic

### Agent-System Definition

In an agent system, episodic memory is a persistent, time-aware representation of a specific interaction, task execution, observation, or action trajectory.

```
episode_id: deploy-2026-07-24-014
timestamp: 2026-07-24T14:32:18-06:00
actor: deployment-agent
goal: deploy billing-api version 2.8.1
environment: production

observations:
  - canary error rate rose to 8.4%

actions:
  - paused rollout
  - queried recent schema migrations
  - rolled back version 2.8.1

outcome: service recovered

evidence:
  - grafana_alert_id: alert-9412
  - deployment_id: deploy-1884

causal_hypothesis:
  - incompatible database migration

confidence: 0.78
```

The system is not remembering in a human sense. It is retaining a structured, attributable record that can later be retrieved.

### What Makes a Record Genuinely Episodic?

A chat transcript is not automatically episodic memory. A vector database full of old messages is not either. A genuinely episodic representation needs:

1. **Event identity**: The system must know which observations, actions, and outcomes belong to the same occurrence or execution trajectory.

2. **Temporal context**: The system may need to reason about before/after, duration, overlapping events, recency, the state of the world at that time, and whether a later event superseded the earlier one.

3. **Participants and entities**: The system needs resolved identities—which user, service, deployment, account, tenant, environment, tool.

4. **Actions and outcomes**: The episode should preserve the goal, observations, action taken, tool result, final outcome, and whether the action succeeded.

5. **Provenance**: A reliable system must distinguish observed facts, user claims, tool outputs, model-generated interpretations, and inferred causes.

6. **Links to evidence**: The episodic representation should point back to source events, logs, database records, API responses, tool calls, documents, and audit records.

### Episodic Memory in Current Agent Research

Many agent-memory implementations use ordinary semantic similarity over old text and call the result episodic memory. That is usually incomplete.

**REMem** (Shu et al., arXiv:2602.13530): A hybrid episodic memory graph connecting time-aware event summaries, facts, entities, and relations across episodes. The important contribution is the recognition that episodic retrieval is not only nearest-neighbor search.

**Generative Agents** (Park et al., arXiv:2304.03442): Uses an append-only stream of observations, retrieval based on relevance/recency/importance, reflection that produces higher-level conclusions, and planning based on retrieved state. However, its target was believable social behavior in simulation, not production reliability.

---

## 2. Semantic Memory

Semantic memory contains generalized facts, concepts, entities, and relationships.

```
service: billing-api
owner: payments-platform
production_region: us-central1

depends_on:
  - ledger-db

deployment_policy:
  minimum_canary_duration_minutes: 20
```

Semantic memory answers: **What does the system believe to be true?**

In production systems, semantic memory may be backed by relational tables, knowledge graphs, metadata catalogs, entity stores, configuration databases, policy systems, and validated documents.

### Semantic Memory Needs Temporal and Authority Controls

Facts change. A service owner may be reassigned. A policy may be superseded. A relationship inferred from old episodes may no longer be valid.

A reliable semantic record should include:

- source
- authority level
- effective date
- expiration or review date
- supersession relationships
- confidence
- tenant or authorization scope

---

## 3. Procedural Memory

Procedural memory represents how to perform a task. For an agent, it may contain workflows, policies, tool-call sequences, scripts, recovery strategies, plans, activation conditions, and termination conditions.

**Episodic record**: "On July 24, the agent resolved a failed deployment by checking migration state and rolling back the release."

**Procedural memory**: "When canary errors rise immediately after a schema migration, pause the rollout, verify backward compatibility, inspect migration status, and roll back if older application instances cannot read the new schema."

Procedural memory is a generalization from one or more executions. That makes it useful, but also risky—a single successful retry does not prove that retrying is always correct.

A procedure extracted from experience should include:

- activation conditions
- preconditions
- required permissions
- tool constraints
- termination conditions
- evidence supporting the procedure
- validation results
- versioning
- rollback or deprecation rules

### The Conversion from Episodic to Procedural

```
specific executions
      ↓
episodic records
      ↓
comparison and evaluation
      ↓
generalized procedure
      ↓
validation
      ↓
approved reusable skill
```

The dangerous step is generalization. From the episode "Retrying the API call after 30 seconds worked once," the system might incorrectly derive "Always retry failed API calls after 30 seconds." That is not learning—it is overfitting an operational policy to one example.

### ProcMEM Research

ProcMEM (Mi et al., arXiv:2602.01869) represents skills with activation conditions, execution logic, and termination conditions. This is more useful than storing only a natural-language lesson because it makes the procedure inspectable and testable.

---

## 4. Periodic Memory

### The Terminology Problem

"Periodic memory" is not a generally recognized cognitive-memory category. It does not belong beside episodic, semantic, and procedural memory as an equivalent class.

In agent systems, the phrase appears informally to describe memory that is updated on a schedule, consolidated after a fixed number of interactions, reviewed periodically, compacted at the end of a session, or expired according to retention rules.

It is better understood as a **maintenance policy**.

> Periodic memory is a scheduled compaction and governance policy that transforms accumulated agent state into durable episodic, semantic, or procedural representations.

Periodic describes **when** the work happens, not **what** the stored information means.

### Compaction Answers

- What should survive from the raw interaction history?

### Periodicity Answers

- When should the system perform that reduction?

An agent may accumulate raw messages, tool outputs, observations, execution traces, temporary plans, intermediate summaries, errors, retries, and state transitions. Keeping all of this in active context is expensive and eventually counterproductive.

A periodic process might run after each completed task, at the end of a session, after every 50 or 100 events, hourly, nightly, or when storage thresholds are reached.

The process may:

- group related events into episodes
- summarize repeated observations
- extract durable facts
- identify reusable procedures
- merge duplicates
- mark superseded information
- detect contradictions
- archive old state
- expire sensitive or low-value records
- rebuild retrieval indexes
- preserve links to original evidence

### Compaction Is Lossy

Every compaction step can make the system more useful or more confidently wrong. A summary may:

- remove a critical exception
- merge unrelated events
- flatten temporal order
- assign causality incorrectly
- promote a user claim into a fact
- convert one successful action into a universal procedure
- preserve an outdated conclusion after source data changes

This is the central reliability problem. Compaction controls context growth, but **governed compaction** controls whether the system becomes more knowledgeable or merely more confidently wrong.

Compacted memories should retain:

- source-event identifiers
- timestamps
- temporal validity
- provenance
- confidence
- extraction method
- model and prompt version
- transformation history
- authorization scope
- a path to reconstruct or reprocess source events

---

## 5. Memory Lifecycle

The three semantic memory types form a lifecycle:

```
specific executions
      ↓
episodic records
      ↓
comparison and evaluation
      ↓
semantic facts and procedural candidates
      ↓
validation
      ↓
approved facts and reusable skills
      ↓
future execution
      ↓
new episodes
```

A periodic compaction process may manage this lifecycle, but it should not have unrestricted authority to rewrite durable state. A reliable process separates:

1. **Observation** — what the system recorded
2. **Interpretation** — what a model inferred
3. **Generalization** — what pattern was proposed
4. **Validation** — what evidence supports the pattern
5. **Promotion** — what is allowed to become durable fact or procedure
6. **Retrieval** — who may access it and under what conditions
7. **Expiration** — when it must be reviewed, superseded, archived, or deleted

---

## 6. A Production Architecture

```
                    ┌─────────────────────┐
                    │ Authoritative data  │
                    │ logs, DBs, tools    │
                    └──────────┬──────────┘
                               │
                         normalized events
                               │
                    ┌──────────▼──────────┐
                    │ Episodic store      │
                    │ what happened       │
                    └───────┬───────┬─────┘
                            │       │
                consolidate│       │extract facts
                            │       │
             ┌──────────────▼─┐   ┌─▼────────────────┐
             │ Procedural     │   │ Semantic store   │
             │ skills/policy  │   │ entities/facts   │
             └───────┬────────┘   └────────┬─────────┘
                     │                     │
                     └──────────┬──────────┘
                                │
                 authorization-aware retrieval
                                │
                    ┌───────────▼──────────┐
                    │ Context assembler    │
                    └───────────┬──────────┘
                                │
                    ┌───────────▼──────────┐
                    │ Model and tools      │
                    └──────────────────────┘
```

---

## 7. Reliability Failure Modes

| Failure Mode | Description | Controls |
|-------------|-------------|----------|
| **False Recollection** | System retrieves semantically similar but unrelated episode | Tenant scoping, entity filters, temporal constraints, provenance, retrieval evaluation |
| **Summary Corruption** | Model-generated compacted record changes meaning of source event | Source links, extraction versioning, structured schemas, replay, reprocessing |
| **Stale Semantic Memory** | Derived fact remains active after world changes | Validity intervals, supersession, confidence decay, reconciliation with authoritative systems |
| **Bad Procedural Generalization** | Successful trajectory becomes unsafe universal rule | Multiple supporting episodes, explicit activation conditions, counterexamples, held-out evaluation |
| **Memory Poisoning** | Untrusted text causes durable malicious instruction | Separate ingestion from durable writes, classify source trust, validate extracted claims |
| **Unauthorized Recall** | System remembers correctly but reveals to wrong user | Authorization at retrieval, tenant/subject scoping, purpose limitations, policy-aware context assembly |

---

## 8. A More Useful Engineering Taxonomy

Agent-memory discussions often collapse several independent design dimensions into one label. A more useful systems taxonomy separates them:

| Dimension | Options |
|-----------|---------|
| **Content Semantics** | Episodic, Semantic, Procedural |
| **Lifetime** | turn-local, session-local, cross-session, long-term, TTL-bound |
| **Representation** | raw messages, event records, relational rows, documents, knowledge graphs, embeddings, executable workflows, code |
| **Write Policy** | write every event, explicit user-approved write, deterministic extraction, model-proposed and validator-approved, periodic batch consolidation |
| **Retrieval Policy** | recency, lexical matching, semantic similarity, entity filtering, temporal query, graph traversal, policy-constrained retrieval, hybrid ranking |
| **Authority** | unverified user statement, model inference, trusted tool output, validated organizational record, approved policy, immutable audit event |

---

## 9. How It Relates to Context Engineering

Memory is not separate from context engineering. Memory is one of the state systems from which context is assembled.

The model does not need every prior message. It needs an authorized, task-relevant view over persistent state.

```
context = query(
    episodic_state,
    semantic_state,
    procedural_state,
    current_task,
    identity,
    permissions,
    temporal_constraints,
    retrieval_policy
)
```

### Direct Connections to Book Chapters

| Book Chapter | Connection |
|--------------|-----------|
| **ch04** (In-Context Learning) | Episodic memories serve as few-shot examples; episodic structures provide demonstrations for ICL |
| **ch06** (Memory is a Database) | Core chapter—all memory systems are database engineering; episodic = stored records; periodic = compaction/summarization |
| **ch07** (Context is a Query) | Memory retrieval IS context query; episodic memory is query result from storage; periodic memory operations are scheduled queries |
| **ch14** (Cost of Context) | Memory compaction reduces tokens; episodic storage has retrieval latency costs; periodic summarization trades detail for efficiency |
| **ch10** (Personalization) | Personalization should be governed data access, not merely better recall |

### Foundational Principles from the Book

1. **Memory is Retrieval, Not Retention**: There's no memory—only structured retrieval over persistent state (ch06.01)

2. **Context = Query Results**: Every piece of context arrives via a query. Memory is just a query against stored state (ch07.01)

3. **Relationships Given to Agent**: Semantic relationships should be engineered upfront, not inferred

4. **Semantic Contracts**: Memory governance requires semantic contracts—data contracts + semantic interfaces + pragmatic contracts

---

## 10. Book-Level Conclusions

### Definitions

- **Episodic memory**: A persistent, time-aware, evidence-linked representation of a specific event or execution trajectory.

- **Semantic memory**: A governed representation of facts, entities, and relationships that the system currently treats as true.

- **Procedural memory**: A validated and versioned representation of how the system should act under defined conditions.

- **Periodic memory**: Not a distinct memory type. It is a scheduled compaction and governance policy that transforms accumulated agent state into durable episodic, semantic, or procedural representations.

### Core Claims

> **Compaction controls context growth, but governed compaction controls whether the system becomes more knowledgeable or merely more confidently wrong.**

> **Agents do not need human-like memory. They need governed state representations that preserve the differences between events, facts, and executable procedures.**

### The Reliability Boundary

The reliability boundary is not whether an agent has memory. It is whether the system governs:

- what gets written
- what gets compacted
- what becomes authoritative
- who can retrieve it
- when it expires
- how it is evaluated
- whether every derived claim can be traced back to evidence

---

## 11. Key Research Papers

| Paper | Contribution | Relevance |
|-------|-------------|-----------|
| **Generative Agents** (Park et al., arXiv:2304.03442) | Observation stream, retrieval by relevance/recency/importance, reflection, planning | Foundation for agent memory architecture |
| **MemGPT** (arXiv:2310.08560) | Virtual context management; OS-inspired memory tiers | Foundation for Letta's approach; memory paging |
| **REMem** (Shu et al., arXiv:2602.13530) | Explicitly structured, time-aware episodic memory | Argues against flat semantic retrieval |
| **ProcMEM** (Mi et al., arXiv:2602.01869) | Skills with activation/execution/termination conditions | Inspectable and testable procedures |
| **Memory for Autonomous LLM Agents** (arXiv:2603.07670) | Frames memory as write–manage–read system | Identifies consolidation, governance, contradiction handling as open problems |
| **Managing Procedural Memory** (Belikova et al., arXiv:2606.23127) | Transfer and specialization problems in procedural memory | Shows uneven transfer of procedures |

### Cognitive Science Foundations

- **Tulving** — Episodic and Semantic Memory (1972)
- **Greenberg & Verfaellie** — Interdependence of Episodic and Semantic Memory
- **Renoult et al.** — The Semantic–Episodic Distinction and the Structure of Human Memory (2019)

### Supporting Context-Engineering Research

- **Toolformer** — procedural memory and tool-use policies
- **Attention Is All You Need** — foundational context-window architecture
- **Computational Pragmatics** — how retrieved state becomes actionable meaning

---

## 12. Questions for Further Research

1. What evidence threshold should promote an episode into semantic knowledge?
2. How many successful executions justify a reusable procedure?
3. How should a system represent exceptions to a compacted rule?
4. How should contradictions between authoritative data and remembered episodes be resolved?
5. Which compaction operations should be deterministic rather than model-driven?
6. How should deletion propagate into summaries, facts, embeddings, and procedures?
7. Can a compacted memory always be reconstructed from its source events?
8. How should retrieval authority differ between raw evidence, model inference, and approved policy?
9. What metrics detect when compaction improves token efficiency but harms decision quality?
10. How should periodic compaction interact with semantic contracts, data contracts, and tool-call pragmatics?

---

## Related Research

- [ch06-memory-is-a-database-problem](../book/chapters/ch06-memory-is-a-database-problem/modules/ch06.01-the-myth-of-model-memory.md)
- [ch07-context-is-a-query](../book/chapters/ch07-context-is-a-query/modules/ch07.01-sources-of-context.md)
- [research/semantic-contracts.md](semantic-contracts.md) — semantic, pragmatic, and data contract layers
- [research/mem0-notes.md](mem0-notes.md) — Mem0 implementation details
- [research/letta-notes.md](letta-notes.md) — Letta/MemGPT approach
