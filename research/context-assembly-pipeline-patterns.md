# Context Assembly Pipeline Patterns

## Core Concept

Context assembly is the query pipeline that assembles context from multiple sources before inference.

From Chapter 7: "Context is a Query"

### Pipeline Architecture

```
User Request → Query Understanding → Source Selection → Retrieval → Ranking → Assembly → LLM
```

**Stages:**
1. **Query Understanding**: Extract intent, entities, constraints
2. **Source Selection**: Choose which stores to query (SQL, API, vector, graph, stream)
3. **Retrieval**: Execute queries against selected sources
4. **Ranking**: Score and filter retrieved context
5. **Assembly**: Combine into prompt/instruction
6. **LLM**: Final inference

### Connection to Web IE

Web IE is the **data ingestion layer** for context pipelines:

```
Web Page → Extraction → Structured Entities → Knowledge Graph → Context Pipeline → LLM
```

IE output becomes context pipeline input.

### Key Patterns

**Hydration**: Populating context slots with retrieved data

**Coverage Metrics**: What % of required context was retrieved

**Freshness Handling**: Context assembly must handle timeouts, failures across queries

**Partial Failure**: Not all sources succeed — pipeline must degrade gracefully

---

## Evidence Ledger

See `research/_evidence-ledger.md`:
- Context Assembly Pipelines — Query pipeline for context
- Hydration Coverage — Measures retrieval success

---

## Cross-References

- **Chapter 7**: Context is a Query — core pipeline concepts
- **Chapter 8**: Knowledge Graphs as context sources
- **Chapter 9**: Hybrid retrieval in pipeline
- **webinformationextraction.md**: IE as ingestion for pipelines