# Metadata for Context Engineering

## Core Concept

Metadata is first-class context — often overlooked but critical for reliable AI systems.

### Types of Metadata

| Category | Examples |
|----------|----------|
| **Source metadata** | timestamps, authors, version, origin URL |
| **Content metadata** | content-type, language, classification, schema |
| **Retrieval metadata** | relevance scores, freshness, authority, trust |
| **System metadata** | permissions, provenance, chain-of-custody, extraction method |

### Why Metadata Matters

Without metadata, the model cannot assess:

- **How fresh** is this information?
- **Who** provided this?
- **Is the user authorized** to see this?
- **Where** did this come from?
- **How reliable** is this source?

### Metadata as Operational Parameters

From Chapter 5: Context Engineering is the process of supplying operational parameters to the model.

**Metadata IS operational parameters:**

```json
{
  "content": { ... },
  "metadata": {
    "timestamp": "2024-01-15",
    "author": "Jane Doe",
    "permissions": ["read", "write"],
    "provenance": "extracted-from-web-page-x",
    "confidence": 0.92,
    "source_type": "ie_extraction"
  }
}
```

### Connection to Web IE

Web IE pipelines produce rich metadata:

- Extraction method (wrapper, grammar, LLM)
- Source URL and timestamp
- Confidence scores
- Schema/taxonomy mappings

This metadata flows through context pipelines to inform the model.

---

## Evidence Ledger

See `research/_evidence-ledger.md`:
- Metadata as Operational Parameters — First-class context

---

## Cross-References

- **Chapter 5**: Context Engineering as Parameter Supply
- **Chapter 7**: Context assembly pipelines carry metadata
- **Chapter 8**: Knowledge graphs can store metadata on entities/edges
- **webinformationextraction.md**: Metadata importance in extraction