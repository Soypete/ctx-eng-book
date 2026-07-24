# Episodic and Periodic Memory in AI/LLM Systems

## Definitions in AI Context

| Term | Human Psychology | AI/LLM Implementation |
|------|------------------|----------------------|
| **Episodic Memory** | Experiences/events stored with temporal/contextual markers | Stored conversation segments, interaction histories, event logs—structured records of what happened |
| **Periodic Memory** | Not a standard psych term—likely refers to periodic/cyclical refresh or refresh cycles | Context compaction, summarization cycles, periodic memory refresh/dump operations |

**Key distinction**: There's no actual "memory" in transformers. There's only retrieval from external systems. This is the foundational thesis of ch06.

---

## How They're Used in AI/LLM Systems

### Episodic Memory Implementation Patterns

| System | Episodic Memory Implementation |
|--------|-------------------------------|
| **Mem0** | Stores independent fact records with temporal markers; ADD-only extraction preserves history of state changes |
| **LangMem** | Captures structured episodes: `observation`, `thoughts`, `action`, `result`—used as few-shot examples |
| **Letta/MemGPT** | Virtual memory management; conversation history tiered across storage layers |
| **Pedro** | Conversation state persisted to document store with semantic search |

### Key Algorithms

1. **Single-pass hierarchical distillation** (Mem0): Extract facts in one LLM call, ADD-only—no overwrites

2. **Memory tiers** (MemGPT): OS-inspired virtual memory with paging between context "RAM" and archival storage

3. **Multi-signal retrieval** (Mem0): Parallel scoring—semantic similarity + keyword matching + entity matching

4. **Git-based memory versioning** (Letta Context Repositories): Memory as filesystem with git-backed versioning for conflict resolution

### Periodic Memory / Context Compaction

**Definition**: Periodic processes that compress or summarize context to fit within window constraints

**Techniques**:
- Summarization: Compress conversation history into compact form
- Hierarchical contexts: Nested summary structures
- Memory dumping: Periodic flush of working memory to persistent storage
- Sleep-time compute: Asynchronous memory consolidation during idle periods (Letta)

**The core problem**:
> "Summarization requires predicting the future. At the moment a model creates a summary, it must decide which information future tasks will require before those tasks exist. Every compaction step becomes an irreversible information loss."

---

## How It Relates to Context Engineering

### Direct Connections to Book Chapters

| Book Chapter | Connection |
|--------------|-----------|
| **ch04** (In-Context Learning) | Episodic memories serve as few-shot examples; episodic structures provide demonstrations for ICL |
| **ch06** (Memory is a Database) | **Core chapter**—all memory systems are database engineering; episodic = stored records; periodic = compaction/summarization |
| **ch07** (Context is a Query) | Memory retrieval IS context query; episodic memory is query result from storage; periodic memory operations are scheduled queries |
| **ch14** (Cost of Context) | Memory compaction reduces tokens; episodic storage has retrieval latency costs; periodic summarization trades detail for efficiency |

### Foundational Principles from the Book

1. **Memory is Retrieval, Not Retention**: There's no memory—only structured retrieval over persistent state (ch06.01)

2. **Context = Query Results**: Every piece of context arrives via a query. Memory is just a query against stored state (ch07.01)

3. **Relationships Given to Agent**: Semantic relationships should be engineered upfront, not inferred

---

## Key Research Papers and Approaches

### Foundational Papers

| Paper | Contribution | Relevance |
|-------|-------------|-----------|
| **MemGPT** (arXiv:2310.08560) | Virtual context management; OS-inspired memory tiers | Foundation for Letta's approach; memory paging concept |
| **Mem0 Research** | Single-pass hierarchical distillation; multi-signal retrieval | Production memory system; benchmarks (LoCoMo, LongMemEval, BEAM) |
| **LangMem** | Memory as tools (manage_memory, search_memory); hot-path vs background processing | Agent-centric memory control |
| **Sleep-time Compute** (Letta, arXiv:2504.13171) | Asynchronous memory consolidation during idle periods | Shifts memory work from user-blocking to background |

### Storage Layer Mapping

| Memory Type | Storage System |
|-------------|----------------|
| Short-term (working context) | In-memory, Redis |
| Episodic (conversation history) | Document store (MongoDB), vector DB |
| Semantic (user profiles, preferences) | Relational (PostgreSQL), knowledge graphs |
| Procedural (agent behavior) | System prompts, prompt rules |

---

## Open Problems

1. **Temporal reasoning at scale**: BEAM benchmark shows 16.3% at 10M tokens
2. **Contradiction resolution**: 35.7% accuracy (1M tokens)
3. **Cross-session structure**: Connecting scattered interactions into timelines
4. **Periodic compaction loss**: Irreversible information loss from summarization

These are **engineering problems**, not model problems—solved by context engineering: retrieval pipelines, storage systems, and scheduled processing.

---

## Summary: Connection to Book Themes

| Theme | Episodic Memory | Periodic Memory |
|-------|-----------------|-----------------|
| **Reliability** | Stored records must be governed, versioned | Compaction must preserve relevance |
| **Context as Engineering** | Structured retrieval | Scheduled context pipelines |
| **Cost** | Retrieval latency, storage | Summarization compute, information loss |
| **Inference vs. Given** | Entity links = given relationships | Trigger conditions = given logic |

The episodic/periodic distinction maps to:
- **Episodic** = what we store (database records, event logs)
- **Periodic** = how we maintain it (compaction, refresh cycles)

Both are engineering problems solved by context engineering—not by expecting models to "remember."

---

## Related Research

- [ch06-memory-is-a-database-problem](/cmd/authorpedro/books/ctx-eng-book/chapters/ch06-memory-is-a-database-problem/modules/ch06.01-the-myth-of-model-memory.md)
- [ch07-context-is-a-query](/cmd/authorpedro/books/ctx-eng-book/chapters/ch07-context-is-a-query/modules/ch07.01-sources-of-context.md)
- [research/mem0-notes.md](/research/mem0-notes.md) — Mem0 implementation details
- [research/letta-notes.md](/research/letta-notes.md) — Letta/MemGPT approach