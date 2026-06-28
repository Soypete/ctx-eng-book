# LLMs in Production — Chapter 12: Ethics

*Notes from listening to Chapter 12 (Ethics / Appendix)*

---

## Larger Token Windows

One of the largest innovations that relate to this book is larger token windows. The frontier models have 1M+ parameter context windows, which means a lot for context engineering — not only in the Fin-Ops context optimization size but in compute, optimization, speed, overload, etc.

The fact that we have such large context windows means we can create incredible models with lots of proprietary insight. Tool calls and multimodal inputs will gobble up tokens in context windows, so we need to keep that as a constraint. Having options and knowing tradeoffs is just engineering.

---

## Parallelism and Context

### Tensor Parallelism

Split model weights across multiple GPUs. Each GPU holds a slice of attention/FFN layers. Used in Megatron-LM (arxiv:1909.08053, arxiv:2104.04473).

**How it affects context:**

- **Memory** — KV cache scales with sequence length × batch size × heads. TP splits model weights but NOT the context itself. Each request's context still needs proportional memory.
- **Speed** — Attention is O(n²). TP doesn't fix this. You need FlashAttention, sliding windows, or prefix caching.
- **Batching** — TP enables larger batch sizes, but longer contexts still compete for memory per request.

### Pipeline Parallelism

Split layers across GPUs. Sequential processing of layers creates bubbles (idle time).

### Data Parallelism

Copy model to all GPUs, split training data. Doesn't help with inference context.

### Key Insight for Context Engineering

Parallelism solves **model size** (fitting larger models), not **context length** (what fits in context). Context engineering is about managing what's *in* that window given compute/memory constraints — that's a different problem than parallelism.

**Key papers:**
- Megatron-LM: https://arxiv.org/abs/1909.08053
- Efficient Large-Scale Training: https://arxiv.org/abs/2104.04473

---

## Data as the Moat

> "The exciting change we are seeing inside the industry due to the introduction of LLMs is that companies are finally starting to understand the importance of governing and managing their data. For some, it's the drive to finetune their own LLMs and get in on the exciting race to deliver AI products. For others, it's the fear of becoming obsolete, as the capabilities of these systems far surpass previous technologies; they are finding it's only their data that provides any type of moat or protection from competition. And for everyone, it's the worry they'll make the same mistakes they've seen other companies make.
>
> LLMs aren't just a driving factor; they are also helping teams label, tag, organize, and clean their data. Many companies had piles of data they didn't know what to do with, but with LLM models like CLIP, captioning images has become a breeze. Some companies have found that simply creating embedding spaces of their text, images, audio, and video has allowed them to create meaningful structures for datasets previously unstructured. Structured data is much easier to operate around, opening doors for search, recommendations, and other insights."

**Why this matters for the book:**
- **Your code is not your IP — the data is**
- Context engineering is about structuring, governing, and retrieving that data effectively
- This quote captures the core thesis: reliable AI systems require engineered context

---

## Appendix: LLMs for Data Processing

LLMs are transforming how organizations process and structure their data:

### Structuring Unstructured Data

- **Text → Structured formats** — LLMs can extract entities, relationships, and schemas from free-form text
- **CLIP and multimodal models** — captioning images, creating embeddings for unstructured media
- **Embedding spaces** — creating meaningful vector representations of text, images, audio, video

### Topic Modeling & Classification

- Zero-shot and few-shot classification without training dedicated models
- Automated tagging and categorization of documents
- Semantic clustering for discovery

### Data Labeling for Smaller Models

- LLMs can generate training labels for smaller, specialized models
- Cost-effective alternative to human annotation
- Enables dataset curation for domain-specific statistical models

### Why This Matters for Context Engineering

The data pipeline feeding your context window is as important as the window itself. Structured, labeled, well-organized data means:

- Better retrieval quality
- Cleaner context construction
- More reliable downstream AI systems

> This could be its own chapter — call it "Data Processing for AI Systems" or similar.

---

## Knowledge Graphs vs Documents

**Key distinction:**
- **Knowledge graphs are not for documents** — they're for **relationships**
- **Ontologies are the relationships** of linked data
- That data can live **anywhere** — databases, documents, APIs, etc.

The value isn't in the documents themselves — it's in the connections between things. Documents are just one possible source of entity data. A knowledge graph can pull from structured DBs, document extraction, API calls, and more.

This is why context engineering isn't just "put more docs in context" — it's about building and querying the relationships that connect distributed data.

---

---

## DSPy for Math with a Dataset

DSPy (Stanford NLP) is a framework for optimizing LLM prompts and pipelines. It uses gradient-like optimization to find the best prompt/inference configuration for a given task.

**Key paper:** https://arxiv.org/abs/2312.08382

**For math/datasets:**
- DSPy can optimize chains of thought for mathematical reasoning
- Uses bootstrapped examples to improve pipeline performance
- Relevant for context engineering: shows how to structure inference pipelines programmatically

---

## Grounding: Preventing Hallucinations

**Key insight:** Hallucination is a data validation problem — we know how to solve those.

### Why System Prompts Fail

- The more context you add, the less reliably the model follows any single rule
- Prompt length increases processing time and can be ignored entirely

### The Data as Guardrail Approach

Using ontologies to ground agentic AI:

1. **Ontology as Schema Constraint** — constrain outputs to typed predicate space (Apple ODKE+, Graph-Constrained Reasoning)

2. **Ontology as Context Engineering** — organize what context gets retrieved and injected (Palantir OAG)

3. **Ontology Behind Tool Calls** — KG schema drives what functions exist and how they're called (FNCTOD)

### Agent Validation Pattern

```
Agent generates → Ontology validates via query → Validated output
                → Reject + retry if invalid
```

### Key Takeaways

- **Data without semantics is just noise to an LLM**
- Deterministic validation beats prompt engineering
- Your data engineering expertise transfers to grounding AI

**See also:** Talk at `~/code/misc/talks/data-as-ai-guardrail/talk.md`

---

## The Right Tool for the Right Job

This connects back to the "Unix filesystem vs DB" argument:

| Use Case | Best Tool |
|----------|-----------|
| Known queries, structured data | SQL / MemPalace |
| Discovery, exploration | Unix filesystem / wiki models (markdown) |
| Text similarity, embedding comparison | Vector stores |
| Multi-line relationships with inference | Ontologies with reasoners |

**The key insight:** Context engineering is a data engineering problem. If you have a good model and know where you want to go, SQL and MemPalace work best. But if you need discovery, flat files with markdown are great. For parsing and comparing vectors, vector stores win. For complex multi-line relationships, ontologies with reasoners win.

This isn't a "one tool to rule them all" problem — it's about choosing the right storage/query system based on the access pattern.

---

## Knowledge Editing

**From LLMs in Production:**

> Knowledge editing is the process of efficiently adjusting specific behaviors. Optimally, this would look like surgery where we precisely go in and change the exact model weights that activate when we get incorrect responses.

**Use case:** Combat factual decay — facts that change over time (current Super Bowl winner, presidents, etc.)

**Why not just retrain/finetune?** Too heavy for updating a fact or two — may change model in unexpected ways.

### Ontology / Wikidata as Guardrail

This is another practical use case for ontologies:

- **Wikidata API** — live query for current factual information
- Before responding with time-sensitive facts, query Wikidata to verify
- Acts as an external knowledge guardrail — model权重不改变, just validate against current data

**Practical code example:**
```python
async def verify_fact(entity: str, claim: str) -> bool:
    # Query Wikidata for current facts
    result = await wikidata_query(entity)
    return claim_matches(result, claim)
```

This combines knowledge editing (external validation) with ontology-backed grounding — without modifying model weights.

---

## Appendix A.5: Mid-20th Century and Modern Linguistics

*From LLMs in Production — Appendix on linguistics history*

### Computers Started as Linguistic Tools

The emphasis on scientific method during early 20th-century linguistics helped set the stage for computational linguistics (CompLing) and NLP.

**Early computers were designed for explicitly linguistic purposes:**
- Alan Turing — information theory, AI
- Claude Shannon — information theory
- Mary Rosamund Haas — comparative historical linguistics

> "What I cannot create, I do not understand."
> — Richard Feynman (relates to file systems — if you can't build it, you don't understand it)

### Why This Matters for Context Engineering

The computer started as a word calculator — a tool for processing language. This history shows that:

- **Computation has always been about managing context** — information, meaning, symbols
- **Automation of context management** isn't new — it's the original purpose
- Modern LLMs continue this tradition — but we need proper data engineering to make them reliable

The linguistics roots of computing reinforce that context engineering is fundamental, not an afterthought.

---

## Aside: Grammars and Chomsky

Noam Chomsky's work on formal grammars is relevant to context engineering — especially for tool calls:

- **Context-free grammars** can constrain what tool calls look like
- Similar to guardrails: you define the valid structure, model must conform
- **Logit biases** and **grammar-based decoding** can enforce valid tool call tags
- Constrained decoding ensures the model outputs valid JSON, correct tool names, proper parameter types

This connects back to the idea that context engineering = applying data engineering principles to AI. Just as grammars formalize valid sentence structures, ontologies formalize valid domain knowledge.

---