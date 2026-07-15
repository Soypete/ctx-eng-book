# Web Information Extraction

## Core Concepts

### "Unstructured" Text Is Actually Organized

Text is often defined as "unstructured" because it's not tabular, but it IS organized:

- **Process structure**: steps, workflows, procedures
- **Section structure**: headers, paragraphs, lists
- **Document structure**: title, abstract, body, references

Web pages add:

- **DOM structure**: HTML elements, nesting, attributes
- **Visual structure**: headings, paragraphs, tables, lists
- **Link structure**: internal and external links

This means web IE can leverage both:
1. **NLP organization** — sentences, paragraphs, sections
2. **DOM organization** — HTML tree structure, element types

This is why web IE is more feasible than "unstructured" text IE — the page has explicit structural signals that NLP and DOM parsing can exploit.

### Metadata Is Overlooked in Context Engineering

Metadata is crucial for context engineering but often overlooked:

- **Source metadata**: timestamps, authors, version, origin
- **Content metadata**: content-type, language, classification
- **Retrieval metadata**: relevance scores, freshness, authority
- **System metadata**: permissions, provenance, chain-of-custody

Context engineering should treat metadata as first-class context:

```
{
  "content": {...},
  "metadata": {
    "source": "web",
    "timestamp": "2024-01-15",
    "author": "...",
    "permissions": ["read", "write"],
    "provenance": "extracted-from-page-x"
  }
}
```

Without metadata, the model cannot assess:
- How fresh is this information?
- Who provided this?
- Is the user authorized to see this?
- Where did this come from?

**Metadata is operational parameters — exactly what Chapter 5 says context engineering should supply.**

### Regex for Normalization, Not Classification

A key distinction in extraction pipelines:

- **Regex for normalization**: Standardize formats (dates, phone numbers, emails) — extract and transform
- **Not for classification**: Don't use regex to decide categories or entity types

Example:
- Normalization: `"01/15/2024"` → `"2024-01-15"` (standardize date format)
- Classification: Don't use regex to decide if a ticket is "billing" vs "technical" — use NER/LLM for that

Regex is deterministic and fast. Classification requires understanding — use ML/LLMs for that.

### The Cost Question: Volume, Velocity, and Context Overhead

IE faces the same DE challenges as most LLM data processing:

**Volume:**
- The graph grows with each extraction
- Number of documents to iterate scales with source data
- More documents = more extraction passes = more cost

**Velocity:**
- Real-time extraction requires parallel workers
- Each extraction pass needs context (conversation history, agent steps)
- Processing becomes very expensive at scale

**The core tension:**

> To extract meaningful data (even with ontologies), you often need to provide the model's full operational context — agent steps, conversation history, prior retrievals.

This creates a paradox:
- Web pages and context aggregations promise value
- But the overhead to make them valuable may exceed their worth
- You need parallel workers, full context replay, expensive retrieval

**The fundamental insight:**

> Data is expensive. LLMs are NOT a shortcut for the cost of data processing.

Augmenting experience with data will always be expensive. The question is whether the extracted value justifies the extraction overhead:

| Approach | Overhead | Value |
|----------|----------|-------|
| Raw web pages | Low (just crawl) | Low (unstructured) |
| Web IE | Medium (extraction) | Medium (entities) |
| KG with context | High (full context replay) | High? |
| Structured DB | Low (already structured) | High |

The value proposition of IE must account for the full context engineering cost — not just extraction, but the infrastructure to make that extracted data usable by the model.

### Ethical Considerations: Scraping and Data Aggregation

Web scraping raises ethical concerns that have come up more with LLMs:

- **Terms of service**: Many sites prohibit scraping
- **Resource usage**: Automated crawling burdens target servers
- **Content ownership**: Extracted data may violate copyrights
- **Consent**: Data was often not intended for extraction

**Alternatives for ethical data aggregation:**

1. **Validate sources**: Use APIs, approved data feeds, licensed content
2. **Direct partnerships**: Work with content providers directly
3. **User-provided data**: Data users explicitly share
4. **Public datasets**: Curated, permissioned collections
5. **Synthetic data**: Generated for training/testing

The LLM era has made this more visible — training data concerns, attribution, and content rights are now mainstream discussions.

IE pipelines should consider: is this data source ethically obtainable?

### Wrappers: Exploiting Markup for Generalization

**Wrappers** are a classic IE technique that exploits HTML markup tags to generalize extraction:

- Use DOM structure (divs, tables, lists) as extraction signals
- Wrap around content to identify repeatable patterns
- Work well on structured/semi-structured pages (e.g., product listings, tables)

**How wrappers work:**

```html
<div class="product">
  <span class="name">Product A</span>
  <span class="price">$99</span>
</div>
```

Wrapper learns: "extract name from span.product span.name"

**Advantages:**
- Deterministic — no ML required
- Fast — DOM traversal is cheap
- Generalizable across similar page structures

**Limitations:**
- Brittle — page structure changes break wrappers
- Only works on structured/semi-structured pages
- Doesn't handle natural language well

Wrappers complement NER/LLM approaches — use wrappers for structured pages, NLP for unstructured content.

### Grammars for Extraction: Minerva Approach

Minerva uses grammars as a core tool for web IE — a structured extraction approach:

**Grammars define:**
- Expected extraction patterns (date formats, quantities, relationships)
- Can be composed and nested
- Provide deterministic parsing rules

**Why grammars matter for web IE:**

1. **Deterministic**: Predictable, reproducible extraction
2. **Composable**: Build complex grammars from simple rules
3. **Self-host friendly**: Can run grammars locally without external APIs
4. **Hybrid potential**: Use grammars for structure, LLMs for semantics

**Example pipeline:**

```
Web Page → Grammar Parser → Structured AST → LLM Enhancement → Knowledge Graph
```

Grammars handle the mechanical extraction; LLMs handle the semantic interpretation.

### Logits as Guardrails for Pragmatics

**Key insight from logits substack:**

Logits (model output probabilities) can serve as guardrails for pragmatics:

- **What logits capture**: Raw probability distributions over next tokens
- **Guardrail use**: Constrain outputs based on probability thresholds
- **Pragmatics connection**: Chapter 5 discusses pragmatics — context that informs meaning

**How logits work as guardrails:**

```
LLM Output → Logit Analysis → Threshold Check → Allow/Block/Flag
```

**Applications:**
- Filter improbable completions
- Detect out-of-distribution responses
- Enforce topic constraints
- Identify uncertainty signals

This is another form of context engineering — using model internals (not just prompts) to reduce uncertainty before downstream processing.

### Part-of-Speech Tagging for Extraction

POS tagging has been a traditional extraction technique:

- Identify noun phrases, verbs, adjectives
- Use grammatical structure to find entities
- Feed into pattern-based extraction

**The cost:**
- Required specialized NLP researchers
- Complex annotation pipelines
- Multiple model pipelines (tokenizer → POS → NER → relation extraction)

**LLMs reduce this cost:**
- Single model handles POS + NER + relation extraction
- No need for NLP research team to build pipeline
- In-context learning adapts to new extraction tasks

This is a significant change — what once required a team of NLP specialists can now be done with prompts and a single LLM call.

### Other Classic IE Systems: Rapier

**Rapier** is another IE system that uses classical NLP techniques:

- **POS tagging**: Identifies grammatical roles
- **WordNet**: Leverages lexical databases for semantic relationships
- **Pattern matching**: Combines linguistic features for extraction

**Comparison:**
| System | Key Technique | Modern Equivalent |
|--------|---------------|-------------------|
| Rapier | POS + WordNet + patterns | LLM in-context learning |
| Minerva | Grammars | Grammar + LLM hybrid |
| Wrappers | DOM structure | Wrapper + NER |

These systems represent the pre-LLM extraction landscape — each pioneered techniques now simplified by modern models.

### Stalker: Hierarchical Data → Embedded Catalog

**Stalker** uses hierarchical data structures to pull content into embedded catalogs:

- **Hierarchical approach**: Exploits site structure, category trees
- **Embedded catalog**: Extracts into structured, nested representations
- **Taxonomy mapping**: Maps extracted content to predefined categories

This is essentially **ontology/taxonomy-based extraction**:
- Define a category hierarchy (taxonomy)
- Map extracted content to appropriate nodes
- Build nested/embedded output structure

**LLMs supercharge this:**
- No need to build complex taxonomy mappers
- In-context learning maps to any ontology you describe
- Can handle implicit taxonomies (infer hierarchy from examples)

```
Taxonomy: Product > Electronics > Computers > Laptops
LLM prompt: "Categorize this product into: [taxonomy]"
→ Laptops (automatically maps through hierarchy)
```

The same principle applies to knowledge graphs — LLMs can navigate and populate hierarchical ontologies without building dedicated taxonomy pipelines.

### LLMs Enable Image/Video Processing

A key advantage LLMs have over classic IE/crawler systems:

- **Image embeddings**: Process visual content as dense vectors
- **Video frame analysis**: Extract information from video streams
- **OCR-like extraction**: Text extraction from images (screenshots, photos, documents)
- **Multimodal understanding**: Combine text, image, and video in extraction

**Dense models are really good at this** — vision-language models can:
- Identify entities in images
- Extract text from screenshots
- Understand diagrams and charts
- Process video frames for context

This is a massive expansion of what's extractable — classic crawlers only handled HTML/text; LLMs can process the full multimodal web.

### Before Building Your Own KG: Use Existing Cross-Domain Knowledge Graphs

**DBpedia** and **YAGO** are large cross-domain KGs worth leveraging:

**YAGO** (Yet Another Great Ontology):
- From Max Planck Institute
- Cross-domain like DBpedia
- High-quality ontological annotations
- Temporal knowledge (facts with time bounds)
- Strong academic foundation

**Schema.org**:
- Standardized vocabulary for structured data
- Used by Google, Microsoft, Yahoo for search
- Covers products, events, organizations, people, recipes, etc.
- Embedded in web pages as JSON-LD, Microdata

**Wikidata**:
- Structured Wikipedia data
- Free, open knowledge base
- Used as backend for Wikipedia
- Covers millions of entities with properties

**Why use existing KGs instead of building:**
1. **Already populated**: Millions of entities, relationships
2. **Well-structured**: Clean schemas, proven ontologies
3. **Maintained**: Updates, fixes handled by community
4. **Connected**: Links to other KGs (Wikipedia, WordNet)
5. **Free**: Publicly available

**Only build your own KG if:**
- Domain is not covered by existing KGs
- Need proprietary/private data
- Need real-time updates not available elsewhere

"Before building a knowledge graph, check what's already there."

### Web IE vs NER

**Two versions of IE for KG construction:**

1. **Web Information Extraction**
   - Operates over raw webpages
   - Attempts to extract a KG with entities, relations, or even events
   - Focuses on structure from web pages (tables, lists, templates)

2. **Named Entity Recognition (NER)**
   - Extracts instances of concepts such as PERSON or LOCATION
   - Concepts come from an ontology (domain-specific)
   - Provides the "nodes" of the knowledge graph

### Web IE as Context Engineering

Web IE connects directly to Chapter 5's Context Engineering theme:

> Context Engineering assumes inference is expensive and unreliable, so the system should remove ambiguity before inference begins.

Web IE is a form of **pre-inference uncertainty reduction**:

- Raw web content → Structured entities/relations
- The system does extraction work BEFORE the LLM needs to infer
- Reduces the model's inference burden

### Web IE Pipeline

```
Web Page → HTML Parsing → Structure Extraction → Entity Extraction → Relationship Extraction → Knowledge Graph
```

**Each stage reduces uncertainty:**
1. HTML parsing: Converts ambiguous markup to structured data
2. Structure extraction: Identifies tables, lists, templates
3. Entity extraction: NER or LLM-based extraction
4. Relationship extraction: Builds edges

### LLMs for Web IE

LLMs can perform web IE in-context:

**Advantages:**
- Can extract from unstructured text within pages
- Handles context-dependent entities
- No training required for new entity types
- Can leverage page context (headers, surrounding text)

**Limitations:**
- Computational cost (per-page inference)
- May miss structure that traditional IE exploits
- Context window limits for long pages

### Hybrid Approaches

Best practice: Combine traditional web IE with LLMs:

1. **Traditional IE** for structure (tables, lists, templates)
2. **LLM IE** for unstructured text sections
3. **Ensemble** for final extraction

This mirrors the Chapter 5 principle:
> The more parameters that can be supplied deterministically, the less inference the model must perform.

### Token Efficiency in Web IE

From Chapter 5: "Every token should have a job."

Web IE contributes to token efficiency by:
- Extracting only relevant entities (not full page content)
- Storing knowledge as structured data (not raw text)
- Enabling retrieval of specific facts rather than context replay

---

## Evidence Ledger Entries

See `research/_evidence-ledger.md`:
- Web IE vs NER — Two Approaches to KG Construction
- Web IE as Pre-Inference Uncertainty Reduction
- Hybrid Web IE — Traditional + LLM
- Token Efficiency Through Structured Extraction

---

## Cross-References to Context Engineering Book

### Chapter 5: Context Engineering

This research connects to Chapter 5:

- **ch05.01 Context Engineering as Parameter Supply**: Web IE supplies structured parameters (entities, relations) to the model instead of raw text
- **ch05.02 Token Efficiency**: Web IE reduces token waste by extracting structured data instead of storing raw网页 content
- **ch05.03 Reliability Through Uncertainty Reduction**: Web IE reduces uncertainty by pre-extracting entities before inference
- **ch05.04 Prompt Engineering vs Context Engineering**: Web IE is context engineering — extracting structure before inference, not prompting the model to infer

### Related Chapters

- **Chapter 7 (Context Assembly)**: Web IE is a core component of context assembly pipelines
- **Chapter 8 (Knowledge Graphs)**: Web IE is the first step in building KGs from web sources