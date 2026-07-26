# Additional Papers to Add to Priority Reading List

## Recommended Additions

---

### 17. BM25 and the Probabilistic Relevance Framework

**Authors:** Stephen Robertson, Hugo Zaragoza
**Year:** 2009
**Venue:** Foundations and Trends in Information Retrieval

**Paper:** [The Probabilistic Relevance Framework: BM25 and Beyond](https://cs.uwaterloo.ca/~jdeprez/publications/robertson09foundations.pdf)

### Why it belongs in the book

BM25 remains the baseline retrieval function against which all vector and hybrid approaches should be measured. The probabilistic relevance framework provides a principled derivation of term frequency and document length normalization, grounding retrieval in formal IR theory rather than embeddings alone.

### Book connection

- Retrieval baselines
- Lexical vs semantic search
- Hybrid retrieval foundations
- TF-IDF → BM25 evolution

### Failure mode explained

A system uses dense embeddings without recognizing that sparse lexical matching often outperforms semantic similarity for exact-term queries, entity lookups, and structured filters.

---

### 18. Event Sourcing

**Author:** Martin Fowler
**Year:** 2005 (ongoing)
**Venue:** Martin Fowler's Patterns of Enterprise Application Architecture

**Paper:** https://martinfowler.com/eaaDev/EventSourcing.html

### Why it belongs in the book

Event sourcing is the foundational pattern for treating memory as structured state rather than append-only log. It provides:

- Complete audit trail
- Temporal queries (state at any point in time)
- Deterministic replay
- Separation of event storage from current-state projection

This is directly relevant to episodic memory architecture.

### Book connection

- Structured state
- Memory as query over event log
- Temporal reasoning
- Replay and recovery

### Failure mode explained

The agent treats conversation history as a simple context window rather than as a queryable event sequence that can be projected, summarized, or filtered by time, actor, or type.

---

### 19. Attribute-Based Access Control (ABAC) for Distributed Systems

**Authors:** David Ferraiolo, Ravi Sandhu, Serban Gavrila, Richard Kuhn, Ramaswamy Chandramouli
**Year:** 2001
**Venue:** ACM Transactions on Information and System Security

**Paper:** https://csrc.nist.gov/projects/role-based-access-control

### Why it belongs in the book

RBAC is well-covered in your existing list. ABAC extends this with attribute-based policies that can express context-dependent authorization:

- Subject attributes (role, department, clearance)
- Resource attributes (classification, owner, sensitivity)
- Environment attributes (time, location, threat level)
- Action attributes (read, write, execute)

This maps directly to your semantic/pragmatic contract arguments.

### Book connection

- Policy-driven authorization
- Context-aware access control
- Governance infrastructure

### Failure mode explained

Authorization is hard-coded into tool visibility rather than expressed as declarative policies that can be audited, versioned, and evaluated at runtime.

---

### 20. Open Policy Agent (OPA): The Policy Decoupling Principle

**Authors:** Torin Sandall, Max Smythe, etc.
**Year:** 2018-present

**Paper:** https://www.openpolicyagent.org/

### Why it belongs in the book

OPA demonstrates production-grade policy decoupling:

- Policies written in Rego (declarative)
- Decision separated from enforcement
- Used by Kubernetes, Terraform, API gateways
- Supports audit trails and policy testing

This is the closest real-world implementation of your "deterministic external governance" argument (from Deontic Policies paper).

### Book connection

- Runtime governance
- Policy as code
- External enforcement
- Policy testing

### Failure mode explained

Policy logic lives inside prompts or application code rather than in a separate, testable, auditable policy engine.

---

## Integrated Reading Order

### After Phase 3 (Governed Agent Systems), add:

17. BM25 and Probabilistic Relevance Framework
18. Event Sourcing

### After Phase 1 (Establish the Thesis), optionally add:

19. ABAC / NIST RBAC Model (already in your list — ensure it's read)
20. Open Policy Agent (if available as paper, otherwise documentation)

---

## Summary

| # | Paper | Book Pillar |
|---|-------|-------------|
| 17 | BM25 and Probabilistic Relevance | Retrieval |
| 18 | Event Sourcing | State/Memory |
| 19 | ABAC | Governance |
| 20 | Open Policy Agent | Governance |