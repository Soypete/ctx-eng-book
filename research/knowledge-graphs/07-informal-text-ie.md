# Chapter 7: Information Extraction from Informal Text

**Source:** Knowledge Graphs Book (Chapter 7)
**Theme:** IE from social media, SMS, and informal text sources

---

## Key Themes

- Social media IE
- SMS / text message extraction
- Informal text challenges
- Noisy text processing

---

## Title: Non-Traditional Information Expressions

### Example: Tweets about California Wildfires

The chapter opens with Figure 1 showing 3 tweets about California wildfires:

- All share hashtag `#CaliforniaFires`
- Entities include: people, locations, groups, ways to help fight
- **Key insight**: Hashtags provide a clustering mechanism — entities can be categorized into domains by linking them through hashtags

### Your Observation: LLMs Struggle with Trending Topics

> "I've often asked for LLMs to give me trending topics to support the posts I write and it does a horrible job at finding them."

This connects to:
- **Ranking algorithms** from search engines are used in social media to suggest trends/recommend content
- LLMs don't have access to real-time social signals
- **Context engineering problem**: The model lacks the operational context (trending data, engagement metrics, temporal signals) to surface relevant topics

### Relevance to Context Engineering

The gap between what LLMs can do and what trending topic discovery requires:
- LLMs trained on static data → no access to real-time trends
- Need external context (APIs, retrieval) to provide trending signals
- This is exactly what context engineering solves — supply the missing operational context

## Temporal Expressions and Multi-Ontology IE

### Tweets as Temporal Expressions

Tweets (and similar short-form social text) are inherently temporal — they capture moments, events, and real-time information. This is similar to what chats from LLMs are.

### Fine-Grained IE Requires Multiple Ontologies

Unlike traditional IE with single-domain ontologies, informal text extraction requires multiple overlapping ontologies:

- **Location/Geo**: Physical places, coordinates, regions
- **FOAF** (Friend of a Friend): People, social relationships, groups
- **Events/Disasters**: Happenings, incidents, temporal events

> Fine-grained IE doesn't need a single domain — informal text naturally spans multiple domains simultaneously.

### Context Engineering Connection

This multi-ontology nature reflects what context engineering deals with:
- Real-world information isn't neatly separated into domains
- Context graphs must handle heterogeneous entity types
- The same piece of text can reference locations, people, events, and temporal markers all at once

## Public vs Custom Ontologies

### Two Approaches to Ontology Design

**Public Ontologies:**
- Standard, widely-used vocabularies (FOAF, Geo, schema.org)
- Already defined, well-documented, interoperable
- Good for open-world data (social media, web)

**Custom/Domain-Specific Ontologies:**
- Built for specific use cases or industries
- Can encode company-specific knowledge, best practices, guardrails

### Use Case: Product Chats

For chats within your product, a domain-specific ontology is powerful:
- Define entities relevant to your domain (products, features, users)
- Encode business rules, permitted actions, constraints
- Use as guardrails for what the system can and cannot do

### Use Case: Software Engineering Hardening

For software engineering interactions, a custom ontology can:
- Define software techniques, patterns, best practices
- Encode company-specific coding standards
- Serve as **global guardrails** — referenceable across all interactions
- Build organizational knowledge into the extraction/inference process

> A custom ontology for software engineering could capture: design patterns, security requirements, code review criteria, deployment practices — then use these as constraints on what the system recommends.

### Context Engineering Insight

This connects directly to context engineering's governance pillar:
- Ontologies are the "config" that constrains behavior
- Custom ontologies enable domain-specific reliability
- Public ontologies provide interoperability

The choice between public vs custom is a design-time decision that shapes runtime behavior.

## Social Media IE Requires Domain-Specific ID

### The Challenge

Social media is hard for ordinary IE because:
- Informal language, abbreviations, slang
- Domain-specific references that general models don't understand
- Fast-moving vocabulary (new terms, memes, trends)
- Requires domain-specific identification that general-purpose IE can't handle

### This Is the Same with Chats

Product chats within your system face the same challenge:
- Domain-specific terminology (your products, features, users)
- Company-specific acronyms and references
- Internal jargon that general models don't know
- Need domain-specific ID to be effective

### And Agent Artifacts

Agent interactions generate their own informal text:
- Tool calls and responses
- Execution traces
- Agent decisions and reasoning
- Error messages and states

All of these require domain-specific understanding:
- What does this tool do? What are its inputs/outputs?
- What does this error mean in this context?
- What decisions were made and why?

> "The same challenge that makes social media IE hard — domain-specific ID — is exactly what makes extracting from chats and agent artifacts hard. You need to understand YOUR domain to extract meaningfully."

### Context Engineering Connection

This is why context engineering matters:
- Generic models don't know your domain
- You must supply the domain context (ontology, terminology, constraints)
- Without domain-specific ID, extraction is noise, not signal

## Open-World Ontologies for Context Analysis

### The Power of Cross-Domain Public KGs

For analyzing context from agents and chats, open-world ontologies are powerful:

- **Wikidata**: Massive cross-domain knowledge base, millions of entities
- **DBpedia**: Structured Wikipedia data, diverse domains
- **YAGO**: High-quality ontological annotations with temporal knowledge
- **schema.org**: Standardized vocabulary used by major search engines

### Why These Work for Agent/Chat Context

These public KGs allow classification across lots of domains:
- Don't need to build a custom ontology from scratch
- Leverage existing entity definitions and relationships
- Handle the multi-domain nature of informal text
- Interoperable — connect to other systems and data

### The Hybrid Approach

For agent/chat context, consider combining:

1. **Open-world ontologies** (Wikidata, DBpedia, YAGO) for general entity classification
2. **Custom ontologies** for domain-specific knowledge, company rules, guardrails

> "Use Wikidata for 'what is this person/place/event?' — use your custom ontology for 'what can this user do in our system?'"

### Context Engineering Application

- Extract entities from chat/agent text → classify against open KGs
- Map to your custom ontology for domain-specific constraints
- The open KG provides the "what", your custom ontology provides the "so what"

This is the pattern: open-world ontologies handle broad classification, custom ontologies handle your specific reliability requirements.

## Bootstrapping NER Systems with Public KGs

### Using Public KGs as Training Data

Public knowledge graphs (Wikidata, DBpedia, YAGO, schema.org) can bootstrap NER systems:

- **Entity lists**: Extract all instances of a class (e.g., all "Person" entities from Wikidata)
- **Label generation**: Use KG labels as training examples
- **Relation patterns**: Learn patterns for relation extraction
- **Cross-lingual**: Many KGs have multilingual coverage

### The Bootstrap Process

```
Public KG → Extract entity lists → Generate training examples → Train NER classifier
```

This solves the cold-start problem:
- Don't need manually labeled data to start
- Leverage existing knowledge instead of building from zero
- Can refine with domain-specific data later

### For Agent/Chats Specifically

- Use Wikidata to recognize general entity types (person, org, location)
- Use your custom ontology to recognize domain-specific entities
- Combine for robust NER that handles both general and specific

> "Public KGs are a jumpstart for NER — they give you labeled data when you have none. Then you fine-tune for your specific domain."

## Ontologies and Guardrails Reduce Hallucinations

### Constraining LLM Outputs

Ontologies and guardrails for LLMs allow restricting hallucinations:

- **Valid entity types**: Only extract entities that exist in the ontology
- **Valid relations**: Only allow relationships defined in the schema
- **Output constraints**: Constrain generation to known values/types

### Your Talk/Blog Post

> You wrote a talk/blog post on this topic — ontologies as hallucination guardrails.

The mechanism:
- Define what is valid (ontology)
- Reject/flag what isn't (guardrails)
- Model output is checked against constraints before returning

### NER Is Computationally Cheaper

**NER** (BiLSTM+CRF or similar traditional models) is computationally cheaper than LLMs:

- Smaller model size → faster inference
- No massive parameter overhead
- Can run on CPU, not just GPU
- Lower cost per inference

For well-defined entity types, traditional NER (BER, CRF) may be preferable:
- Speed matters for high-volume extraction
- Cost matters at scale
- LLMs for complex/novel cases, NER for known domains

### The Trade-off

| Approach | Cost | Flexibility | Best For |
|----------|------|-------------|----------|
| NER/CRF | Low | Fixed entity types | High-volume, known domains |
| LLM | High | Open extraction | Novel domains, complex cases |

> "Use NER when you know what you're looking for. Use LLMs when you don't — then graduate to NER once you know."

## OpenIE Automation: What Agents Solve

### The Book Says

The book notes that automation is hard for OpenIE (Open Information Extraction):
- OpenIE extracts relations without predefined schemas
- Hard to automate at scale — requires human oversight
- Validation and quality control are challenging

### What Agents Actually Solve

Agents solve this automation problem:
- **Autonomous extraction**: Agents can decide what to extract, when to extract
- **Multi-step IE**: Chain tool calls for extraction → validation → storage
- **Self-correction**: Agents can retry failed extractions, verify outputs
- **Continuous operation**: Run extraction pipelines without constant human intervention

> "OpenIE automation was hard — then agents came along. Agents provide the 'automation glue' that OpenIE always needed."

### The Agent + IE Pattern

```
Text Input → Agent decides extraction strategy → Tool calls (NER, RE, validation) → Graph storage
```

This is exactly what context engineering does:
- Agents orchestrate the IE pipeline
- Tools do the extraction work
- Guardrails ensure quality
- The result feeds into the context graph

## LLMs Simplify the IE Pipeline

### Syntactic Checks and POS Tagging No Longer Needed

With LLMs, we don't need separate syntactic checks or part-of-speech tagging:
- LLMs inherently understand grammar and syntax
- These tasks can be solved by the smallest models (even 1B parameters)
- No need for separate NLP pipelines (tokenizer → POS → NER → relation extraction)

### The Old Pipeline (Pre-LLM)

```
Text → Tokenizer → POS Tagger → NER → Relation Extraction → Graph
```

Each step required specialized models, research teams, and complex pipelines.

### The New Pipeline (LLM)

```
Text → LLM → Extract entities + relations → Graph
```

One model handles what used to require multiple specialized components.

### What Still Needs Small Models

For well-defined tasks at scale, small models (CRF, BiLSTM) are still useful:
- POS tagging: Fast, accurate for standard text
- Syntax checking: Rule-based or lightweight models
- Cost-sensitive extraction: When inference cost matters

> "LLMs democratized IE — what required an NLP team now requires a prompt. But for high-volume production systems, the smallest models still win on cost."

## Rule-Based Systems: Kraken and Examplar

### The Book's Examples

The book mentions **Kraken** and **Examplar** as rule-based IE systems:
- Use hand-crafted rules and patterns
- Create n-ary extraction tuples (not just subject-verb-object, but multi-way relationships)
- Leverage lexical and syntactic patterns

### How Rule-Based Systems Work

```
Text → Hand-crafted rules/patterns → N-ary tuples → Graph
```

Rules define:
- Syntactic patterns to match (e.g., "NP VP NP" structures)
- Slot fillers (what goes in each position)
- Constraints on valid extractions

### Comparison: Rule-Based vs LLM + Ontology

| Aspect | Rule-Based (Kraken/Examplar) | LLM + Ontology |
|--------|------------------------------|-----------------|
| **Setup cost** | High (hand-craft rules) | Low (write prompts) |
| **Maintenance** | Brittle (rules break with new text) | Flexible (adapt via prompts) |
| **Coverage** | Limited to defined patterns | Open-world (any described type) |
| **Accuracy** | High for known patterns | Variable, can hallucinate |
| **Inference cost** | Very low (regex/parsing) | High (LLM API) |
| **N-ary relations** | Built-in (by design) | Requires prompt engineering |

### The Value of Rule-Based Systems

**Computational advantages:**
- Extremely fast (regex + parsing, no neural inference)
- Deterministic — same input = same output
- No API costs, runs locally
- Great for high-volume, well-understood patterns

**What we lose with LLMs:**
- Determinism (LLMs can vary)
- Cost control (per-token pricing)
- Transparency (rules are explicit, LLM behavior is opaque)

### What We Gain with LLMs + Ontology

- Handle edge cases rules miss
- Adapt to new domains without rewriting rules
- Open-world extraction (not limited to predefined patterns)
- Handle informal text (slang, abbreviations, new expressions)

### The Hybrid Approach

For production systems, consider:

1. **Rules for hot paths**: High-volume, well-defined extractions (e.g., email, phone, dates)
2. **LLMs for edge cases**: Unusual text, novel entity types, complex relations
3. **Ontologies as guardrails**: Validate both rule and LLM outputs

> "Rules are like handrails — they work great on the main path. LLMs are like a guide dog — they help you navigate the unexpected. Use both."

## Clause-Level OpenIE

### How OpenIE Works at Clause Level

The book describes OpenIE as working at the clause level:
- **Subject + Verb + Object**: Core extraction pattern
- **Prepositions**: Additional relationships (e.g., "with", "in", "at")
- **Sentence splitting**: Break text into clauses
- **Graph dependencies**: Use dependency parsing to understand grammatical relationships

### The Pipeline

```
Sentence → Clause splitting → Dependency parsing → Subject-Verb-Object + Propositions → Tuples
```

### Dependency Graph Example

For "John gave Mary a book in the library":
- Subject: John
- Verb: gave
- Objects: Mary, book
- Prepositional: in the library
- Dependency graph links these into structured tuples

### What This Means

OpenIE at clause level:
- Extracts relations between entities within each clause
- Handles complex sentences by breaking them down
- Uses grammatical structure (not just surface patterns)
- Creates n-ary tuples (subject, verb, object, prepositions)

### Connection to Context Engineering

This is similar to what a context graph does:
- Break down user messages into actionable components
- Extract entities, relations, and context
- Store as structured triples for reasoning

> "Clause-level extraction is essentially doing what a good context engineering system does — breaking complex text into structured, processable pieces."

## Inter-Proposition Relationship Modeling

### The Last OpenIE Class

The book describes **inter-proposition relationship modeling** as the final class of OpenIE:
- Models relationships **between** propositions/clauses
- Not just within a single clause, but across multiple clauses
- Captures discourse-level structure

### What This Captures

Examples of inter-proposition relationships:
- **Causal**: "Because X happened, Y resulted"
- **Temporal**: "First X, then Y"
- **Contrast**: "X but Y"
- **Elaboration**: "X. Specifically, Y"
- **Condition**: "If X, then Y"

### Why This Matters

Inter-proposition modeling captures:
- Argument structure (premise → conclusion)
- Narrative flow (event sequencing)
- Logical relationships (cause-effect, conditions)
- Context that single-clause extraction misses

### Connection to Context Engineering

This is crucial for context engineering:
- User messages often have implicit relationships between parts
- "Do X, but also handle Y" requires understanding the conjunction
- "I want A, and I previously asked about B" requires cross-turn memory

> "Inter-proposition relationships are what make context engineering necessary — the model needs to understand how the pieces fit together, not just what the pieces are."

## Graph Hydration and Creation via NLP Techniques

### Inter-Proposition Modeling as Graph Technique

Inter-proposition relationship modeling is a key technique for **graph hydration and creation**:
- Populates the graph with structured relationships
- Not just entities (nodes), but the connections between them (edges)
- Builds the context graph from raw text

### Why Populate Ontologies via NLP?

We want ontologies that we can populate via NLP techniques because:

1. **Scalability**: POS, NER, dependency parsing can process text at scale
2. **Cost efficiency**: Small models (1B params) can do this work cheaply
3. **Determinism**: Rule-based NLP gives consistent, auditable outputs
4. **Graph growth**: Each extraction adds nodes and edges to the context graph

### The Pipeline for Graph Creation

```
Text → POS tagging → NER → Clause extraction → Inter-proposition RE → Graph
```

Each NLP stage contributes:
- **POS**: Identifies grammatical roles (nouns, verbs, modifiers)
- **NER**: Finds entities (nodes for the graph)
- **Clause extraction**: Identifies propositions
- **Inter-proposition RE**: Creates edges between propositions
- **Graph**: The hydrated context structure

### This Is Context Graph Construction

This is exactly what building a context graph looks like:
- Start with raw text (user messages, agent outputs, documents)
- Apply NLP to extract structure
- Populate the graph with entities and relationships
- The graph becomes the context for future inference

> "Graph hydration is the process of turning text into structure — and NLP techniques (POS, NER, clause extraction, inter-proposition RE) are the tools that do it."

## Triples vs Tuples: Clarifying the Concept

### You've Been Thinking in Triples

Triples are the standard KG format:
- **Subject → Predicate → Object**
- Example: (John, gave, Mary)
- Each triple is one edge in the graph

### Tuples Are More General

**Tuples** are n-ary — they can have any number of elements:
- **Binary tuple**: (subject, object) — like a simple relationship
- **Ternary tuple**: (subject, verb, object) — more detail
- **N-ary tuple**: (subject, verb, object, preposition, time, manner, etc.)

> "A triple IS a tuple — it's just a tuple with 3 elements. Tuples are the general concept, triples are the specific case."

### Why N-ary Tuples Matter

Standard triples lose information:
- "John gave Mary a book yesterday" → (John, gave, Mary)
- What about the book? The time?

N-ary tuples preserve more:
- (John, gave, Mary, book, yesterday)
- Or separate tuples: (John, gave, Mary), (gave, time, yesterday), (gave, object, book)

### Graphene

**Graphene** is likely the tool you heard of — it's a system for:
- Creating tuples from text
- Managing n-ary relations
- Building knowledge graphs from extraction

This connects back to Kraken and Examplar — they also create n-ary tuples.

## Polarity, Modality, Attribution, Quantities

### What These Are

These are linguistic features that add nuance to extractions:

**Polarity:**
- Positive or negative sentiment/factuality
- "John likes coffee" vs "John doesn't like coffee"
- Ontology impact: Can be modeled as a boolean property on the relation

**Modality:**
- Possibility, necessity, obligation
- "John may come" vs "John must come" vs "John will come"
- Ontology impact: Requires modal logic or separate relation types

**Attribution:**
- Who said something, source of information
- "According to John, the meeting is at 3pm"
- Ontology impact: Attribution edge in the graph (fact → source)

**Quantities:**
- Numbers, measurements, counts
- "John bought 5 books" — the 5 is a quantity
- Ontology impact: Quantity as an attribute on the entity or relation

### How They Apply to Ontologies

These features can be encoded in ontologies:

```
Relation: likes
  - subject: Person
  - object: Thing
  - polarity: boolean (positive/negative)
  - modality: enum (possible, necessary, certain)
  - attributed_to: Person (who stated this)
  - quantity: number (if applicable)
```

### The Complexity Trade-off

Adding these to ontologies increases:
- **Expressiveness**: More detail captured
- **Complexity**: More schema elements to manage
- **Inference cost**: More to reason over

> "Polarity, modality, attribution, and quantities are what make simple triples insufficient — they require richer tuple representations or additional graph edges to capture fully."

## KnowItAll: POS-Only OpenIE

### What the Book Says

**KnowItAll** is an OpenIE system mentioned in the book:
- Uses **only POS tagging** — no NER, no linguistic parsing
- Extracts patterns from POS sequences
- Relies on surface patterns rather than deep linguistic analysis

### How KnowItAll Works

```
Text → POS tagging only → Pattern matching → Tuples
```

It looks for POS patterns like:
- "NNP VBZ NNP" → extract as (X, verb, Y)
- "NNP VBZ JJ NN" → extract with adjective modifier
- No entity recognition — just pattern matching on tags

### Comparison: KnowItAll vs LLMs

| Aspect | KnowItAll (POS-only) | LLMs |
|--------|---------------------|------|
| **Linguistic analysis** | Shallow (POS only) | Deep (internal representation) |
| **Entity recognition** | None (pattern-based) | Built-in (NER capabilities) |
| **Parsing** | None | Grammar understanding |
| **Setup** | Pre-built patterns | Just prompt |
| **Coverage** | Limited to POS patterns | Open-world |
| **Cost** | Very low (regex on tags) | High (LLM inference) |
| **Quality** | High for matching patterns | Variable, can hallucinate |

### What KnowItAll Shows

The book uses KnowItAll to demonstrate that you don't need complex NLP pipelines:
- POS tagging alone can extract useful tuples
- Surface patterns have surprising power
- Simpler approaches can work for known domains

### The Lesson for Context Engineering

This connects to the hybrid approach:
- **Simple POS patterns** for fast, cheap extraction of known patterns
- **LLMs** for complex, open-world extraction
- **Combine both**: Use KnowItAll-style patterns for hot paths, LLMs for edge cases

> "KnowItAll proves that you don't need fancy NLP — just POS tags and good patterns. LLMs give you more, but at higher cost. The best systems use both."

## KnowItAll Components: Early Agent Pattern

### The Components

KnowItAll has these components:
- **Extractor**: Extracts tuples from text using POS patterns
- **Search Engine**: Validates/expands extractions via web search
- **Assessor**: Scores/validates the extractions
- **Database**: Stores the extracted tuples

### This Is Like Agent Tool Calls

Yes — this is exactly the pattern of an agent:
- Extractor → tool for extraction
- Search Engine → tool for validation/expansion
- Assessor → tool for scoring/verification
- Database → tool for storage/retrieval

### The Parallel

```
KnowItAll:          Modern Agent:
Extractor    →      Extract tool
Search Engine →    Search/Verify tool
Assessor     →      Evaluate tool
Database     →      Store/Retrieve tool
```

### What This Means

The book is showing that **agent-like architectures existed before LLMs**:
- Break the problem into steps
- Each step is a tool/component
- Chain them together for the full pipeline

> "KnowItAll is essentially a rule-based agent — it has tools (extractor, search, assessor, database) and orchestrates them. Modern LLMs just make the tools smarter."

This connects directly to context engineering:
- Agents orchestrate IE pipelines
- Tools do the work
- The architecture has been stable since before LLMs

### Limitation: Not Scalable for New Relations

The book notes a key limitation:
- **KnowItAll is not amenable to scalable addition of new relations**
- To add a new relation type, you need to write new POS patterns
- Each new relation = new hand-crafted rules

### Why This Matters

This is the core trade-off:
- **Rule-based systems** (KnowItAll): Great for fixed relations, hard to extend
- **LLM-based systems**: Easy to add new relations (just update the prompt)

### The Scalability Problem

```
KnowItAll:
Add "CEO_of" relation → Write POS pattern for "CEO_of" → Test → Deploy
Add "founded_by" relation → Write POS pattern for "founded_by" → Test → Deploy
... (each new relation is a manual process)
```

vs.

```
LLM:
Add "CEO_of" relation → Update prompt: "Extract CEO_of relations" → Done
Add "founded_by" relation → Update prompt: "Extract founded_by relations" → Done
```

> "This is why LLMs won for IE — the scalability of adding new relations via prompts beats the brittleness of adding new rules."

For context engineering, this means:
- Use rule-based (KnowItAll-style) for stable, high-volume relations
- Use LLMs for new or changing domains
- The hybrid approach handles both

## TextRunner: Weak Supervision + Classifier Approach

### The Problem TextRunner Solves

- **Latency**: Real-time extraction is slow with traditional OpenIE
- **Too many ontologies**: Managing multiple ontologies is expensive

### TextRunner's Approach

TextRunner uses:
- **Weak supervision**: Not fully supervised (no manual labels), but not unsupervised either
- **Classifier (Naive Bayes)**: Learns to classify extractions as valid/invalid
- **Self-supervised**: Generates training data from the text itself

### How It Works

```
Text → Extract candidate tuples → Train classifier on patterns → Classify new extractions
```

The classifier learns:
- Which surface patterns produce valid relations
- Which don't — without human labeling

### Comparison: TextRunner vs KnowItAll

| Aspect | TextRunner | KnowItAll |
|--------|-----------|-----------|
| **Learning** | Classifier (Naive Bayes) | Pattern matching |
| **Supervision** | Weak (self-supervised) | None (POS-only) |
| **Adaptation** | Learns from data | Fixed patterns |
| **Latency** | Higher (training + inference) | Lower (rule-based) |

### Why This Matters

TextRunner shows a middle ground:
- Not fully manual (like KnowItAll's patterns)
- Not fully unsupervised (like early clustering)
- Uses weak supervision to bootstrap a classifier

This is relevant for context engineering:
- Can train on your own data (chat logs, agent traces)
- Classifier learns what "valid" looks like in YOUR domain
- More adaptable than pure rules, cheaper than full LLM

## Knowledge Extraction Options: The Practical Menu

### The Idea

The book should probably include a section that lists the options for knowledge extraction in context and data processing:

- **Not every method captures everything** — that's OK
- **Different trade-offs**: cost, accuracy, coverage, latency
- **Combined with guardrails**, these can produce usable knowledge graphs

### The Menu of Options

| Method | Coverage | Cost | Latency | When to Use |
|--------|----------|------|---------|-------------|
| **NER (BiLSTM-CRF, BER)** | Fixed entity types | Low | Fast | Known domains, high volume |
| **KnowItAll (POS patterns)** | Pattern-based | Very low | Fast | Stable relations, quick extraction |
| **TextRunner (Naive Bayes)** | Learns from data | Medium | Medium | Adaptable, self-supervised |
| **Rule-based (Kraken/Examplar)** | Defined patterns | Low | Fast | Deterministic, auditable |
| **LLM-based extraction** | Open-world | High | Slow | Novel domains, complex cases |
| **Clause-level OpenIE** | Relations within clauses | Medium | Medium | Complex sentences, multi-part relations |
| **Inter-proposition RE** | Cross-clause relations | Medium | Medium | Narrative understanding, discourse |

### The Point

These methods don't need to capture everything:
- They just need to capture **enough** for the use case
- **Guardrails** filter out noise, validate outputs
- The **graph** grows incrementally — more extraction = more context

### Guardrails as the Filter

Guardrails are essential:
- Validate extractions against the ontology
- Reject invalid entities/relations
- Flag uncertain extractions for human review
- Ensure the graph stays clean even with imperfect extraction

> "You don't need perfect extraction — you need good enough extraction with good guardrails. The graph grows over time as you add more extractors and refine your ontologies."

### For Context Engineering Specifically

Practical options for building context graphs:
1. **Start with NER** for your core entities (users, products, tools)
2. **Add rule-based extraction** for common patterns (dates, emails, IDs)
3. **Use LLM for edge cases** and novel relations
4. **Layer guardrails** to validate everything

This is the realistic path — not one perfect extractor, but a menu of options that together give you enough to work with.

## Event Detection vs Extraction in Social Media

### Event Detection First

The chapter notes that event detection is typically the first task in social media IE:
- Detect that something is happening (trending, notable)
- THEN extract details about the event

### Knowing vs Understanding

> "Knowing that something is trending is more important than what that trend implies"

This distinction matters:
- **Detection**: Is this noteworthy? (classification)
- **Extraction**: What are the details? (information extraction)
- **Understanding**: What does it mean? (reasoning)

### For Context Engineering

This maps to agent systems:
- **Detection**: Is this user message important? Does it mention a tool, a problem, a goal?
- **Extraction**: What tool? What problem? What goal?
- **Understanding**: How should the agent respond? (requires reasoning over the graph)

## Links as Augmentation

### The Book Notes

Links in social media can augment extraction:
- URLs point to more information
- But extracting from links requires fetching and parsing the linked content

### This Is a Search Problem

Using links effectively is a search/retrieval task:
- Fetch the linked page
- Extract relevant info
- Use as additional context

For chat-based systems:
- Links in user messages can trigger retrieval
- The linked content becomes part of the context graph
- This is different from extraction — it's augmentation via search

## TwiCal: Events + Time = Usage Patterns

### What Is TwiCal?

**TwiCal** (Twitter Calendar) is an open-domain event extraction and categorization system for Twitter, developed by **Alan Ritter** and colleagues. It's described as "the first open-domain event-extraction and categorization system for Twitter."

**Original Paper**: Ritter et al., EMNLP 2011
- Title: "TwiCal: Event Extraction and Categorization from Twitter"

### Technical Details

TwiCal uses:
- A tagger specifically trained with Twitter data
- The **TempEx time tagger** for temporal extraction
- **LinkLDA** (a topic model) for event categorization
- Extracts a **four-tuple** structure: (event, time, entity, context)

Output:
- **Event phrases** (event mentions)
- **Named entities**
- **Temporal information** (dates/times)
- Builds an "event calendar" from Twitter — organizing extracted events by date and topic

### Impact

- 1200+ citations in related survey papers
- Influenced subsequent work in social media event extraction and detection
- Enables open-domain event extraction without being limited to pre-defined event schemas

### What TwiCal Does

TwiCal extracts:
- **Event phrases**: What's happening
- **Temporal information**: When it happened
- **Named entities**: Who/what is involved
- Creates a four-tuple: (event, time, entity, context)

### Why This Matters for Agents/Chats

The interesting insight from this chapter:

**In chat-based systems, "events" are user actions and tool calls, and "time" is the sequence of the conversation.**

This means:
- Detect when a tool is mentioned/used → event detection
- Track when it happens → temporal extraction
- Build patterns over time → usage trends

### What This Enables

1. **Tool usage optimization**: Which tools are frequently used together?
2. **Problem pattern detection**: What problems recur in conversations?
3. **Contextual timing**: When in a conversation is something mentioned?
4. **Trend analysis for graph relationships**: If tool X is often mentioned with problem Y, that's a relationship that can influence graph structure

### The Connection to Context Engineering

> "Events in chats + time = tool usage patterns"

This is essentially doing what TwiCal does, but for agent interactions:
- Extract "events" (tool calls, user requests, agent decisions)
- Track "time" (conversation sequence, timing patterns)
- Build a calendar of agent activity
- Use patterns to inform graph relationships and agent behavior

The graph doesn't just store facts — it stores **temporal patterns of usage** that can be queried to optimize future responses.

## QA-Driven Semantic Role Labeling (SRL) Benchmarks

### The Paradigm

QA-driven SRL reformulates semantic role labeling as a question-answering task:
- Instead of formal labels (Agent, Patient, Theme), use natural language questions
- "Who took a walk?" → identifies the Agent role
- Reduces annotation cost — crowd workers need minimal training
- Provides interpretable outputs

### Key Benchmarks

**1. QA-SRL (Question-Answer Driven SRL)**
- Original paper: He, Lee, Zettlemoyer et al. (2015)
- Dataset: QA-SRL 2.0 — a larger version
- The foundational paradigm for QA-driven SRL

**2. QAMR (Question-Answer Meaning Representations)**
- Paper: "Crowdsourcing Question-Answer Meaning Representations" (https://aclanthology.org/N18-2089/)
- Dataset: ~5,000 sentences, ~100,000 questions
- Covers predicate-argument relationships from PropBank, NomBank, QA-SRL
- Includes previously under-resourced: implicit arguments and relations

**3. LSOIE (Large-Scale Open Information Extraction)**
- Paper: https://arxiv.org/abs/2101.11177
- Dataset: https://github.com/jacobsolawetz/LSOIE
- Created by converting QA-SRL 2.0 — 20x more data than next largest human-annotated OIE dataset

### Relevance to Context Engineering

- Frame extraction prompts as QA: "What tool was used?" instead of "Extract tool relation"
- More natural, interpretable outputs
- Can use the same QA pipeline for both SRL and verification

## RelVis: OIE System Benchmark

### What Is RelVis?

**Paper**: "Analysing Errors of Open Information Extraction Systems" (https://arxiv.org/abs/1707.07499)
- Authors: Rudolf Schneider, Tom Oberhauser, Tobias Klatt, Felix A. Gers, Alexander Löser
- Presented at Building Linguistically Generalizable NLP Systems at EMNLP 2017

### The Benchmark

RelVis is a **toolkit and benchmark for evaluating Open Information Extraction systems**:
- **4 datasets**: 3 from news domain + 1 from Wikipedia
- **4,522 labeled sentences**
- **11,243 binary or n-ary OIE relations**

Used to evaluate four OIE systems:
- ClausIE
- OpenIE 4.2
- Stanford OpenIE
- PredPatt

Also analyzed impact of five common error classes on 749 n-ary tuples.

### Relevance

- Provides a standard way to compare OIE systems
- Helps identify error patterns in extraction
- Useful for benchmarking your own extraction pipeline against known systems

## Key Takeaway: Providing an Ontology for IE Systems

### The Central Insight

The largest takeaway from this chapter is the importance of **providing an ontology for IE systems**:

> "The key is providing an ontology that guides what to extract and how to structure it."

### Why Ontologies Matter

1. **Guides Extraction**: The ontology tells the IE system what entities and relations matter
2. **Constrains Output**: Prevents over-extraction and noise
3. **Enables Guardrails**: Validates that extractions conform to expected patterns
4. **Provides Structure**: Transforms raw text into structured, queryable data

### Without Ontology = Noise

- OpenIE without ontology = lots of tuples, little usable structure
- No guidance on what's important
- Hard to query, reason over, or use for decisions

### With Ontology = Signal

- Extraction focused on meaningful entities/relations
- Graph becomes queryable and useful
- Enables reasoning, guardrails, and decision-making

### For Context Engineering

This is the core insight:

> "The ontology IS the context engineering — it defines what the system knows about, what it can extract, and what it should ignore."

- Build your ontology first
- Let the ontology guide extraction
- Use the graph for reasoning
- Apply guardrails to maintain quality

This connects to all pillars of context engineering:
- **Representation**: The ontology is the schema
- **Acquisition**: Extraction populates the graph per the ontology
- **Governance**: Ontologies as guardrails
- **Reasoning**: Graph queries over ontology-defined structure

## Multilingual IE: LLMs Show Massive Advantages

### The Chapter's Point

The chapter likely discusses IE systems for multilingual and cross-lingual settings (Spanish, Slavic languages). Traditional IE requires:
- Language-specific models
- Language-specific training data
- Separate pipelines per language

### LLMs Change This

**LLMs, even small ones, show huge advantages:**

1. **Single Model, Many Languages**
   - One model handles dozens of languages
   - No language-specific models needed

2. **Cross-lingual Transfer**
   - Train on high-resource (English)
   - Apply to low-resource languages
   - Zero-shot works surprisingly well

3. **Low-resource Languages**
   - LLMs can extract from languages with limited labeled data
   - Synthesis approaches (LLM-generated training data) address scarcity

### Research Findings

| Paper | Languages | Key Finding |
|-------|-----------|-------------|
| **Universal NER v2** | 100+ | Gold-standard benchmark for massively multilingual |
| **Otter** | 100+ | Outperforms GLiNER-x-base by 5.3pp F1 |
| **Cross-lingual RE (Romanian)** | Romanian | Zero-shot: 3-5pp drop vs English; QLoRA reduces to 1.4pp |
| **DetIE** | Portuguese, Spanish | Zero-shot OpenIE, 75% F1, 3.35x faster |
| **milIE** | Arabic, Galician, Chinese | First OpenIE dataset for some languages |

### Relevance to Context Engineering

For chat-based systems with global users:
- **One extraction pipeline** for all languages
- **Ontology guides extraction** regardless of language
- **Cross-lingual transfer** from English to other languages works well
- **Fine-tuning** can close the gap if needed

> "LLMs democratize multilingual IE — what required a team of language-specific models now requires one model and an ontology."