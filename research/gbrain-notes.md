# GBrain Agent Brain Framework - Research Notes

**Source:** https://github.com/garrytan/gbrain  
**Date Analyzed:** July 2026  
**Focus:** Agent memory architecture, knowledge retrieval, and relationship handling

---

## 1. What Problem GBrain Solves

### The Core Problem: Amnesiac Agents

GBrain tackles the fundamental limitation of AI agents: **memorylessness**. The README explicitly frames this as:

> "your coding agent stops being amnesiac about everything that isn't code"

Most AI agents only see what's in their immediate context window. They can't:
- Remember previous conversations
- Connect new information to existing knowledge
- Synthesize answers across multiple sources
- Identify gaps in their own knowledge

### Scale and Ambition

GBrain is built for **large-scale personal knowledge management**:
- 146,646 pages in Garry Tan's personal deployment
- 24,585 people, 5,339 companies tracked
- 66 autonomous cron jobs running 24/7 for enrichment
- Multi-user team ("company brain") deployment with OAuth-scoped access

This is not a toy memory system. It's production-grade infrastructure for managing institutional knowledge at scale.

### Synthesis as the Differentiator

GBrain's key innovation is the **synthesis layer** - not just retrieval, but actual answer generation:

**Most Tools Return:**
```
1. people/alice — Alice runs engineering at Acme...
2. meetings/2026-03-15-alice-q1 — Q1 product review with Alice...
```

**GBrain Returns:**
```
Alice runs engineering at Acme (a series-B fintech). You last spoke
on April 22 in a quick pricing chat. Three things are still open
from that conversation:
1. She owes you the security review...
2. You committed to pricing...
3. She mentioned they're hiring a CISO...

Heads up: nothing's been added to the brain about Alice or Acme
since April 22, six weeks ago.
```

This synthesis includes:
- Explicit citations to source pages
- Gap analysis (what the brain doesn't know)
- Temporal awareness (staleness detection)

---

## 2. How It Handles Memory/Context

### The Signal-Response Loop

GBrain implements a continuous memory loop:

```
signal → search → respond → write → auto-link → sync
```

Each component:

| Component | Description |
|-----------|-------------|
| **Signal detector** | Runs on every message; captures ideas, entity mentions, time-sensitive todos |
| **Brain-first lookup** | Queries the brain BEFORE any external API call |
| **Auto-link** | Extracts entity references on every write, creates typed edges |
| **Cron-driven enrichment** | Overnight: dedup, fix citations, score salience, find contradictions |

### Two-Tier Context Architecture

1. **Short-term context:** Signal detection and immediate retrieval
2. **Long-term context:** The synthesized knowledge graph

This mirrors the book's "State" pillar (Part IV in five-pillars-outline.md) around persistent memory and context lifecycles.

### Schema Packs: Defining Brain Shape

GBrain doesn't impose a fixed layout. It uses **schema packs** to define:

- What page types exist (`person`, `company`, `meeting`, `deal`, `email`, etc.)
- What relationships are valid (`attended`, `works_at`, `invested_in`, `founded`, `advises`)
- What facts get extracted automatically

The default `gbrain-base-v2` ships with 15 canonical types, but users can author their own schemas. This is directly aligned with the book's "Semantics" pillar - reducing ambiguity through explicit structure.

---

## 3. Knowledge Storage and Retrieval

### Storage Layer

**Dual engine support:**
- **PGLite** (Postgres 17 via WASM): Zero-config, default for personal brains up to ~50K pages
- **Postgres + pgvector**: For large/multi-machine deployments (Supabase or self-hosted)

The brain repo (markdown files on disk) is the system of record. DB stores:
- Page embeddings
- Knowledge graph edges
- Timeline entries
- Citation metadata

### Retrieval Architecture

GBrain uses a **four-strategy hybrid approach:**

1. **Vector search (HNSW on pgvector)** - Semantic similarity
2. **BM25 keyword** - Lexical matching
3. **Reciprocal-rank fusion (RRF)** - Merges rankings without weighting one globally
4. **Knowledge graph traversal** - Typed-edge following

### The Benchmark Results

| Strategy | P@5 | R@5 |
|----------|-----|-----|
| ripgrep BM25 only | ~18 | ~75 |
| vector-only RAG | ~18 | ~80 |
| hybrid + RRF (no graph) | ~18 | ~85 |
| **GBrain full stack** | **49.1** | **97.9** |

The graph adds **+31 P@5 points** - this is the critical finding. Vector search alone underdelivers on relational queries.

### Zero-LLM Edge Extraction

Every `put_page` runs `extractEntityRefs` - pattern matching on:
- Standard markdown links: `[Garry Tan](wiki/people/garry-tan)`
- Obsidian wikilinks: `[[wiki/people/garry-tan|Garry Tan]]`
- Typed-link blockquotes

This creates typed edges (`attended`, `works_at`, `invested_in`, etc.) with **zero LLM calls**. The graph grows on every write at near-zero cost.

### Advanced Retrieval Features

- **Title-phrase boost:** Query matching a page's title gets priority
- **Alias hop:** Free-text aliases bridge synonyms ("Hall of Light" → Mingtang page)
- **Intent-aware rewriting:** Classifies queries as `entity`, `temporal`, `event`, or `general`
- **Multi-query expansion:** For high-detail queries, generates 2-3 variants and merges via RRF

### Named-Thing Retrieval

A specific focus on retrieving pages by their canonical names, not just semantic similarity:
- Per-page max-pool (best chunk per page, not N scattered chunks)
- Evidence tags on every result (`alias_hit`, `exact_title_match`, `high_vector_match`, etc.)
- `create_safety` hints (`exists`, `probable`, `unknown`) for agents deciding whether to write duplicates

---

## 4. Is GBrain a Knowledge Graph?

GBrain has **retrieval structure** but not **classical graph reasoning**:

### What GBrain Has (KG-Shaped Retrieval)
- Schema packs: explicit types + relationships
- Auto-linked edges: pattern-matching on wikilinks (zero LLM calls)
- Graph traversal as retrieval strategy (+31 P@5 over vector-only)
- Named-thing retrieval with typed edges

### What GBrain Does NOT Have
- No OWL reasoning / transitive closure
- No rule-based inference in the graph
- No formal ontology enforcement

### The Distinction
Reasoning happens in the **LLM synthesis layer**, not the graph. This aligns with Ch8's thesis: *"Modern knowledge graphs with LLMs perform reasoning in the model. The graph provides structure and retrieval. The model provides interpretation and inference."*

GBrain is a **KG-shaped retrieval index** — graph structure for better retrieval, LLM for synthesis/reasoning. Not a classical knowledge graph with built-in reasoning.

---

## 5. Comparison to Book Thesis

### Thesis: "reliable AI = engineered context (pragmatics, data, semantics) — relationships given to agent, no inference needed"

Let me map GBrain's approach against each component:

### Semantics: What does the information mean?

GBrain aligns strongly with this pillar:

- **Schema packs** define explicit page types and relationships
- **Knowledge graph** stores typed edges (`works_at`, `invested_in`, etc.) - relationships are explicitly given, not inferred
- **Entity resolution** through auto-linking with zero LLM calls
- **Gap analysis** in synthesis explicitly states what the brain doesn't know

**Alignment: Strong** - GBrain treats semantics as engineering, not inference.

### Pragmatics: What should be done with the information?

GBrain's synthesis layer directly addresses this:

- **Signals** capture intent from messages
- **Skills** (43 curated) encode behavioral patterns
- **Intent classification** routes queries to appropriate retrieval strategies
- **Tool routing** through MCP with scoped permissions

**Alignment: Strong** - GBrain encodes what to do with retrieved information through skills and the synthesis layer.

### Governance: What is allowed?

GBrain has explicit governance:

- **OAuth 2.1 + PKCE** for team brain access
- **Scope-gated access** (`read` / `write` / `admin`)
- **Per-user, per-source permissions** - team members only see what they're allowed to see
- **Rate limiting** on the HTTP server

**Alignment: Strong** - Governance is built into the architecture.

### State: What persists over time?

GBrain directly addresses state management:

- **Persistent memory** via markdown repo + Postgres
- **Timeline entries** track page history
- **Dream cycle** (cron jobs) keeps memory fresh overnight
- **Event-sourcing style** - deletes in git become soft-deletes in DB

**Alignment: Strong** - State persistence is the core value proposition.

### Where GBrain Goes Further

1. **Synthesis over retrieval:** GBrain doesn't just find pages; it generates answers with citations. This is a pragmatic layer on top of semantic storage.

2. **Self-improving:** The dream cycle automatically:
   - Fixes citations
   - Finds contradictions
   - Consolidates memory
   - Scores salience

3. **Multi-agent orchestration:** The Minions queue handles durable subagents that survive crashes via two-phase pending→done persistence.

4. **Company brain as first-class concern:** Multi-user, OAuth-scoped, federated access is not an afterthought.

### Where It Potentially Differs

The book emphasizes "relationships given to agent, no inference needed." GBrain delivers this through:
- Auto-linking (pattern-based, not inferred)
- Schema packs (explicit relationships)
- Graph traversal (explicit paths)

However, some aspects still use inference:
- **Intent classification** is deterministic but still classification
- **Query expansion** uses Haiku for variant generation
- **Synthesis** uses LLMs to compose answers

The "no inference" principle is approximated but not absolute - which is realistic for practical deployment.

---

## 6. Key Takeaways for the Book

### What GBrain Demonstrates

1. **Memory is infrastructure, not feature.** GBrain treats the knowledge base as critical infrastructure requiring:
   - Health checks (`gbrain doctor`)
   - Repair workflows
   - Migration paths
   - Backup strategies

2. **Hybrid retrieval wins.** Pure vector search is insufficient. The +31 P@5 improvement from graph traversal proves that semantic similarity alone cannot capture relational queries.

3. **Synthesis is the killer app.** Returning pages is okay; returning answers with citations and gap analysis changes how users work. This maps directly to the book's "Pragmatics" pillar.

4. **Schema as contract.** Defining what types exist and how they relate is essential for reliable behavior. GBrain's schema packs are a practical implementation of the book's "Semantics" pillar.

5. **Automation keeps the brain fresh.** The dream cycle (overnight enrichment) addresses the staleness problem that plagues most personal knowledge systems.

### Architecture Patterns to Highlight

| Pattern | Book Pillar | GBrain Implementation |
|---------|-------------|----------------------|
| Explicit relationships | Semantics | Schema packs, typed edges |
| Intent-driven retrieval | Pragmatics | Intent classifier, query routing |
| Scope-gated access | Governance | OAuth scopes, per-user permissions |
| Persistent memory | State | Markdown repo + Postgres, dream cycle |
| Gap awareness | Pragmatics | Synthesis includes "what we don't know" |

### Questions for Further Research

1. How does GBrain handle conflicting information across sources?
2. What are the failure modes of the auto-linker?
3. How does the schema evolution work in practice?
4. What embedding/dimensionality choices work best for personal knowledge?

---

## 7. References

- Main repo: https://github.com/garrytan/gbrain
- Design doc: https://github.com/garrytan/gbrain/blob/master/DESIGN.md
- Retrieval architecture: https://github.com/garrytan/gbrain/blob/master/docs/architecture/RETRIEVAL.md
- Deployment topologies: https://github.com/garrytan/gbrain/blob/master/docs/architecture/topologies.md
- Eval methodology: https://github.com/garrytan/gbrain/blob/master/docs/eval/SEARCH_MODE_METHODOLOGY.md
- Benchmark suite: https://github.com/garrytan/gbrain-evals
