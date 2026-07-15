# WordNet, Anchor Text, and Semantic Web Crawlers

Research notes on lexical resources for web search agents.

> **Convention:** This note follows the principles-first paradigm.
> - Core principles (timeless) are marked with "## Core Principle"
> - Current technology snapshots are marked with "## Current Implementation"

---

## Core Principle: WordNet as Guardrails for Semantic Expansion (Timeless)

The other thing WordNet can do is let web crawlers expand a search to related concepts while keeping our guardrails.

WordNet provides a structured, curated lexical database that constrains semantic expansion to known, verified relationships — unlike arbitrary web content which may drift or hallucinate connections.

---

## Core Principle: Anchor Text as Semantic Validation (Timeless)

### Definition

Anchor text is the clickable, visible text in a hyperlink. It serves as a descriptor of the linked content.

Example:
```html
<a href="https://example.com">click here</a>
                         ^^^^^^^^^^^ anchor text
```

### Relevance to Web Search Agents

- Anchor text historically signals page relevance (SEO signal)
- For AI agents: anchor text can be used as ground truth labels for linked content
- Semantic crawlers can use anchor text patterns to validate and constrain retrieval

### In Context Engineering

When building web search agents, anchor text provides:
- A human-curated signal about page content
- A constraint layer: "only follow links where anchor text matches the query domain"
- A validation mechanism: compare retrieved content against what anchor text promised

---

## Core Principle: Semantic Crawlers with Ontologies and Taxonomies (Timeless)

### Concept

Semantic crawlers go beyond keyword matching by:

1. **Using ontologies** — formal definitions of concepts and their relationships
2. **Using taxonomies** — hierarchical classifications of topics
3. **Mapping content** — matching page content against ontology classes
4. **Constrained traversal** — only following links that maintain semantic relevance

### Architecture

```
Query → Ontology Lookup → Related Concepts → Taxonomy Traversal → Filtered Crawl
```

### Why This Matters for Web Search Agents

- **Precision**: Only retrieve pages that map to relevant ontology nodes
- **Recall**: Expand to related concepts without brute-force crawling
- **Guardrails**: The ontology acts as a bounds — you can't "drift" into irrelevant content

### Metadata / Topic Web Pages

Semantic crawlers can also parse:
- Metadata tags (Dublin Core, schema.org)
- Topic web page descriptions
- Structured data annotations

These provide additional guardrails for relevance.

---

## Core Principle: WordNet for Constrained Semantic Expansion (Timeless)

### What is WordNet?

A large lexical database of English words, organized into synsets (sets of synonymous words). Provides:
- Synonym sets (synsets)
- Hypernyms (broader concepts)
- Hyponyms (narrower concepts)
- Meronymy (part-whole)
- Antonyms (opposites)

### Using WordNet for Search Expansion

Given a query term, WordNet can:

1. **Find synonyms** — expand "car" → "automobile", "vehicle"
2. **Find hypernyms** — expand "car" → "vehicle", "transportation"
3. **Find hyponyms** — expand "vehicle" → "car", "truck", "bus"
4. **Navigate the lexical graph** — traverse relationships programmatically

### Guardrail Properties

Unlike vector similarity (which can return anything semantically "close"), WordNet:

- Is human-curated
- Has verified relationships
- Has defined sense (disambiguation) for polysemous words
- Provides provenance: each relationship has a source

### Practical Application for Web Search Agents

```python
# Pseudocode for WordNet-guided crawl
query = "medical research"
synsets = wordnet.synsets(query)

for synset in synsets:
    # Expand to related concepts
    hypernyms = synset.hypernyms()  # broader
    hyponyms = synset.hyponyms()    # narrower
    related = synset.similar_tos()  # similar topics

    # Build constrained crawl list
    crawl_targets = [s.name() for s in hypernyms + hyponyms + related]
```

This gives you a bounded, explainable expansion — not a blind vector search.

---

## Connecting to Context Engineering

### Why This Matters

1. **Web search agents are the easiest to build** — they're testable, have clear signals (links), and bounded scope
2. **Semantic crawlers demonstrate context engineering via tool calls** — the agent must decide what to fetch, when, and how to validate
3. **WordNet provides guardrails without hallucination** — the relationships are explicit, not inferred

### The Pattern

```
User Query → Disambiguate (WordNet senses) → Expand (WordNet relations) →
  Filter (ontology mapping) → Crawl (semantic constraints) →
    Validate (anchor text, metadata) → Context (retrieved content)
```

This is a complete context engineering pipeline: structured retrieval over curated state.

---

## Harvest Rate — Crawler Performance Metric

### Definition

**Harvest rate** = the percentage of downloaded pages that are relevant to a topic.

Industry standard metric for focused crawler performance.

### Connecting to Agent Tool Calls

> Any tool call that fetches data is just a crawler with a little spice.

This reframes agent tool calls as a generalization of web crawlers:

| Crawler (Web) | Agent Tool Call (Generalized) |
|---------------|-------------------------------|
| Downloaded pages | Fetched data from tool |
| Relevant pages | Relevant to current task |
| Harvest rate | Precision of tool calls |
| Focused crawling | Selective tool invocation |

### What This Enables

If we measure tool calls by harvest rate:
- **Database queries** — what % of returned rows were actually needed?
- **Graph queries** — what % of retrieved nodes were used?
- **Document retrieval** — what % of chunks were referenced in the response?
- **Image retrieval** (multimodal) — what % of images were relevant?

This connects directly to the context precision/recall formulation from earlier in the book:

```
Harvest Rate = Useful Retrieved / Total Retrieved
```

### The Generalized Context

It's not just web sources anymore. The "crawl" now includes:
- Web pages
- Databases
- Graph stores
- Text documents
- Images (for multimodal models)
- APIs
- File systems

All of these can be measured by harvest rate to evaluate tool call efficiency.

---

## Proposed Context Engineering Metrics Framework

Starting with harvest rate as the foundation:

### 1. Harvest Rate (Core)
```
Harvest Rate = Useful Retrieved / Total Retrieved
```
- Measures the precision of any data-fetching tool call
- Analogous to IR precision but for tool calls

### 2. Context Recall (Bounded)
```
Context Recall = Relevant Retrieved / Relevant Available
```
- Requires bounding the domain (like the KG approach)
- Harder to measure but essential for completeness

### 3. Latency Impact
```
Latency Cost = f(token_count, model_size, cache_state)
```
- Measures the cost trade-off of retrieval decisions

### 4. Agent Path Efficiency
```
Path Efficiency = Goal Achieved / Tool Calls Made
```
- Measures how efficiently the agent reached its objective
- Analogous to focused crawler's "number of pages visited to find relevant content"

### 5. Staleness Score
```
Staleness = Time Since Last Retrieval / Allowed Staleness Threshold
```
- Measures how fresh the context is
- Critical for agents that rely on dynamic data

### 6. Hallucination Rate (via Citation)
```
Citation Match = References That Actually Support Claim / Total References
```
- Measures whether retrieved context was used accurately
- Connects to the precision at the "usage" layer

---

## Expanding the Framework

This framework can expand to:
- **Cross-source harvest rate** — aggregate across multiple tool types
- **Domain-specific precision** — calibrated per use case
- **Cumulative context value** — measuring the value of context over a session
- **Agent learning rate** — how quickly does the agent improve its tool call precision?

The key insight: **harvest rate is the foundation** because it directly measures the efficiency of the retrieval mechanism, which is what tool calls fundamentally are.

---

## Learning Crawlers: User Preference Learning

### Concept

Learning crawlers can learn user preferences from a topic — they observe what the user engages with and build a model of preferences within that domain.

### Relevance to LLMs and "Memoroes"

This is directly relevant to:
- **Memory systems**: If you have a personalized set of patterns recognized into a harness, this can help tool calls behave in specific ways
- **User modeling**: The crawler learns what the user finds valuable within a topic area
- **Adaptive retrieval**: Instead of static prompts, the agent learns from interaction patterns

### The Current Problem

This preference information is typically:
- Hard-coded in the system prompt
- Static and not learned from behavior
- Not adaptable to individual users

### The Context Engineering Opportunity

Move preference learning from system prompt to tool call:

```
Current: System prompt contains "user prefers concise answers"

Opportunity: Tool call returns preference model learned from past interactions
  → Tool: get_user_preferences(topic="software engineering")
  → Returns: {response_length: "short", format: "code", detail_level: "high"}
```

### The Harness Pattern

A "harness" is a structured recognition system that:
1. Observes user behavior within a topic
2. Extracts preference patterns
3. Encodes them in a machine-readable format
4. Provides them to agents via tool calls (not system prompts)

This is more maintainable than:
- Giant system prompts
- Manual preference documentation
- Hard-coded behavior assumptions

---

## Generating the Context Graph (Section 3.3.1.1)

### The Key Idea

The crawler uses the context graph to track relationships and pull data for all the 1 to 2 node hops. This works if you have an ontology for a domain.

Even if the ontology is just **hydrated** (populated with instances), you can query that directly instead of crawling. The graph becomes the index.

**Does NOT work if the agent is building the graph itself.**

---

## The Sprawl Problem — AI is Additive

Via Kendall Clark (linked above, has been working on graphs for years):

> AI is additive. It will always increase the size of the graph if it is in charge of building it.

This is the fundamental problem:
- Agents build graphs by adding nodes/edges
- There's no inherent constraint on what gets added
- Graph grows unboundedly = sprawl
- Query performance degrades
- Precision drops

---

## The Solution: Predefined Domain Ontologies

### Scaffolding vs. Structured Ontologies

| Approach | Agent Builds Graph | Problem |
|----------|-------------------|---------|
| **Scaffolding** | Agent builds its own steps | Not predictive, workflow loops are costly |
| **Predefined Domain Ontology** | Agent queries predefined structure | Constrained, efficient, bounded |

### How It Works

1. **Define the domain ontology first** — what entities and relationships matter
2. **Make the context graph a predefined problem** — the ontology defines where the agent should go and find data
3. **New data from the agent fits into those ontologies** — agent extracts and maps, not creates
4. **If we need to increase the ontology** — that's a systems engineering process for data engineers, not the agent

### The Key Insight

> Predefined domain ontologies make the context graph a **bounded problem** rather than an **emergent one**.

The agent doesn't decide what to add to the graph. The ontology already defines the structure. The agent just:
- Queries the ontology
- Extracts data that fits existing schema
- Returns results in expected format

This is the opposite of the "agent builds its own knowledge graph" approach, which leads to sprawl.

---

## The Alternative: Tool Calls as Encapsulated Relationships

Instead of building a full ontology, you can:

1. **Limit tool calls** — constrain the number of data sources the agent can query
2. **Finite goals** — each tool call has a bounded objective
3. **Provided pragmatics** — embed relationships in the tool call itself

### The Pattern

Instead of:
```
Agent builds graph → Agent decides what to add → Sprawl
```

Do this:
```
Tool call has pragmatics: "fetch X, relationship is Y"
Tool call has limit: max 5 sources
Tool call has goal: "get user profile data"
Graph emerges from tool call results, not from agent decisions
```

### Why This Beats GBrain and Similar Tools

| Approach | Problem |
|----------|---------|
| **GBrain and graph-defined tools** | Always sprawl because they let the agent decide what to add |
| **Tool-call encapsulated relationships** | Bounded by tool definition, finite by design |

The key insight: **if you don't let the agent build the graph, it can't sprawl.**

### Why This Matters

This gives agents **real incremental value** without:
- Full knowledge graph maintenance
- Ontology engineering overhead
- Graph storage costs
- Sprawl management

The relationships are **encapsulated in the tool call** rather than built by the agent.

**This is why I don't like GBrain and other graph-defined tools** — they will always sprawl because they're architecturally designed to let the agent build relationships, which is additive by nature.

---

## Context Graphs — The Foundation

### The Quote

From **Knowledge Graphs: Fundamentals, Techniques, and Applications** (Kerjriwal, Knoblock, & Szekely, 2021):

> "The structure of paths leading to relevant pages can be an important factor in focused crawling, as first shown with context graphs by Diligenti et al. (2000)."

Also relevant: Foundation Capital's "Context Graphs: AI's Trillion-Dollar Opportunity" (2020) — published before GPT, anticipating structured context needs.

This was published in 2020, before GPT became mainstream. It anticipates the need for structured context in AI systems.

### Diligenti et al. (2000) — Context Graphs

The original context graph method:

- Builds classifiers for sets of pages at distance 1 or 2 from relevant pages
- Uses Hidden Markov Models (HMMs) for browsing
- Applies **sequence labeling** and context-focused crawling
- Works backward: collects backlinks to relevant pages to discover pages from the context graph

The key insight: instead of crawling forward from seed pages, work backward through the link graph to find pages that connect to relevant content.

### Sequence Labeling — Connecting NER and Learning Crawlers

**Named Entity Recognition (NER)** uses sequence labeling to identify and classify entities in text (PERSON, ORG, DATE, etc.).

**Learning crawlers** also use sequence labeling to classify pages and predict navigation paths.

**This is the key technique for distinguishing context graphs from agent runs:**

> Use sequence labeling to label and differentiate nodes in the context graph vs. the agent's execution trace.

The pattern:
- **Context graph nodes** = pages, entities, relationships (labeled via NER-style sequence labeling)
- **Agent run nodes** = tool calls, decisions, state changes (labeled via execution trace)
- **Edges** = link structure (context graph) vs. causal flow (agent run)

This creates a unified representation where both web structure and agent behavior can be analyzed together using the same sequence labeling framework.

### Relevance to LLMs and Tool Calls

Context graphs as described by the VBQ (presumably the book they're reading) are supposed to capture data from actions of a graph and make them usable for agents via some kind of graph and crawling.

**The key insight for context engineering:**

> If you map the data to a domain, it is much easier to store in a subgraph via ontologies, to reason about, and use the reasoner/inference to reduce crawling steps.

This connects directly to:
- **Ontologies as domain constraints** — mapping data to a domain makes it manageable
- **Subgraph storage** — storing in domain-specific subgraphs rather than the full web graph
- **Inference for reduction** — using reasoners to infer relationships and reduce the need for exhaustive crawling

### The Pattern for LLM Web Agents

```
1. Start with relevant pages (known good sources)
2. Collect backlinks (who links to these pages?)
3. Build context graph from link structure
4. Train classifier on pages at distance 1-2 from relevant
5. Use HMM to predict likely relevant paths
6. Crawl strategically, not exhaustively
```

For LLM tool calls: instead of "search the web broadly", the agent could:
- Use known relevant URLs as seeds
- Fetch their backlinks via tool calls
- Build a bounded context graph
- Traverse strategically based on learned paths

### Connection to the Book's Thesis

This reinforces: **context = structured retrieval over persistent state**. The context graph is the persistent state (link graph), and the crawler is the retrieval mechanism. Adding ontology mapping makes it domain-specific and reason-able.

---

## Research Gap

Need to find:
- Existing semantic crawler implementations
- WordNet APIs for programmatic access (NLTK, Princeton's official API)
- Ontology mapping approaches (schema.org integration)
- Evaluation metrics for semantic crawl quality
- **Diligenti et al. (2000) — Context Graphs paper (CRITICAL)**
- Foundation Capital "Context Graphs: AI's Trillion-Dollar Opportunity" article
- **Is "merged context graphs" an alternative to Semantic Web?** (Open question)

---

## Open Question: Context Graphs vs Semantic Web — Are They in Conflict?

The 2000 paper calls them "merged context graphs." This was published around the same time as the Semantic Web (early 2000s).

**Are they alternatives or complementary?**

### The Short Answer

**Not exactly in conflict**, but they represent fundamentally different philosophies about how to make web data usable.

### Detailed Comparison

| Aspect | Semantic Web | Context Graphs (Merged) |
|--------|-------------|------------------------|
| **Philosophy** | Make everything machine-readable upfront | Find relevant data dynamically |
| **Mechanism** | Annotate with RDF/OWL | Use link structure + classifiers |
| **Assumption** | We CAN annotate everything | Annotation is impractical |
| **Scope** | Universal (all web data) | Focused (relevant pages only) |
| **Era** | 1999-2001 (Berners-Lee) | 2000 (Diligenti et al.) |

### The Key Tension

Semantic Web asked: "How do we represent all knowledge so machines can understand it?"

Context Graphs asked: "How do we efficiently FIND the knowledge we need?"

These aren't mutually exclusive, BUT:

### Why Context Graphs Won (Practically)

The Semantic Web **didn't achieve its vision**. Most web content was never marked up in RDF/OWL. It was too much work for too little payoff.

Context Graphs offered a **more practical approach**:
- Don't require universal annotation
- Work with existing link structure
- Focus on finding relevant pages, not representing everything

### The Conflict (If Any)

> Semantic Web assumes a **declarative** solution (mark it up, then machines can reason).

> Context Graphs represent an **procedural** solution (use the graph to find what you need).

For LLM agents: **Context Graphs win** because:
1. We can't annotate the whole internet
2. We can't predict what the agent will need
3. Dynamic discovery via graph traversal is more practical

### Conclusion

Not in direct conflict, but Context Graphs represent a more pragmatic philosophy that actually worked, while Semantic Web's vision remained largely unrealized. For agents, Context Graphs are the better foundation.

---

## Classifier Training in Context Graphs

### From the Paper

The original context graph paper uses ML classifiers to assign pages to layers in the Merged Context Graph (MCG). This is key:

1. **Train classifier** on pages that are known relevant (seed set)
2. **Assign new pages to layers** based on their predicted relevance
3. **Use layer assignments** to prioritize crawling

### The Same Thing for Agents + Knowledge Graphs

This is exactly what we need for agent knowledge graphs:

> Model domains in a way that data pipelines can do automatic hydration of links, URLs, paths, or database data based on classification.

The pattern:
```
Domain model → Classification → Automatic hydration
                  ↓
          "This entity belongs
           to class X"
                  ↓
          Pipeline hydrates:
          - Links to related entities
          - URLs for external data
          - Paths to internal data
          - Database queries
```

This is the **automatic data pipeline** approach vs. the "agent builds it manually" approach.

### Do People Know Context Graphs Exist?

**Question:** Like most people do scraping for their crawlers right? Do people actually use context graphs this way?

**Probably not.** Here's why:
- Context graphs are academic (2000 paper)
- Most practical crawlers use simpler approaches:
  - BFS/DFS with depth limits
  - Relevance scoring via keywords
  - Sitemap parsing
  - Link equity (PageRank proxies)

The "merged context graph" approach with classifier-based layer assignment is sophisticated and not widely adopted in mainstream crawler tooling.

### What Does Google Do?

**Unknown publicly**, but we can infer:

- Uses link structure extensively (the original context graph insight)
- Has massive computing to do exhaustive crawling
- Uses ML for classification (likely)
- PageRank = link structure as a ranking signal
- Probably doesn't use "context graphs" as a named concept, but uses the underlying principles

**Fun fact:** Google's original PageRank (1998) came from the same academic tradition as context graphs — using link structure to find relevant pages. The context graph paper (2000) explicitly built on this idea.

---

## Research Questions

1. Are any modern crawler tools explicitly based on context graphs?
2. What do open-source crawlers (Scrapy, etc.) use for focused crawling?
3. Is there commercial tooling that does classifier-based layer assignment?
4. How do modern search engines (Google, Bing) approach "focused crawling" at scale?

---

## Declarative vs Procedural in Context Engineering

Based on our reading:

### Declarative Approaches

| Approach | Definition | Example |
|----------|------------|---------|
| **Predefined Ontologies** | Define domain structure upfront, agent queries it | "Here is the schema, query against it" |
| **Semantic Web** | Annotate data with RDF/OWL, then reason | Mark up documents, use OWL推理 |
| **Fixed Tool Schemas** | Tool calls have fixed interfaces | MCP tools with defined parameters |
| **System Prompts** | Tell the agent the rules upfront | "Always do X before Y" |

**Philosophy:** "State the structure/ rules first, then execute"

### Procedural Approaches

| Approach | Definition | Example |
|----------|------------|---------|
| **Context Graphs** | Use link structure to find relevant data dynamically | Crawl 1-2 hops from relevant pages |
| **Agent Builds Graph** | Agent discovers and adds relationships | GBrain, Mem0, let agents add edges |
| **Scaffolding** | Agent builds its own steps during execution | ReAct-style reasoning traces |
| **In-Context Learning** | Agent learns from examples in conversation | Few-shot prompts |

**Philosophy:** "Discover and adapt as you go"

### The Key Trade-off

| Aspect | Declarative | Procedural |
|--------|-------------|------------|
| **Upfront work** | High (define schema) | Low (let agent discover) |
| **Sprawl risk** | Low (bounded) | High (additive) |
| **Precision** | High (structured) | Variable (depends on agent) |
| **Maintenance** | Engineering process | Emergent (no maintenance) |
| **For LLMs** | Query predefined KG | Agent builds from scratch |

### The Book's Position (Synthesized)

Context engineering should lean **declarative** because:
1. Agent-built graphs always sprawl
2. Procedural approaches are less predictable
3. Predefined ontologies make the problem bounded
4. "Context = structured retrieval over persistent state"

But with a twist: **use tool calls as the interface** rather than full KG construction. The agent doesn't build the graph — it queries structured tools that return bounded data.

### Refined Position: Procedural Agents + Declarative Data

**User's refinement:**

> Procedural processes make the agents work better, but data injection is best done declaratively.

This is a crucial distinction:

| Layer | Approach | Why |
|-------|----------|-----|
| **Agent behavior** | Procedural | Agent learns, adapts, builds steps dynamically |
| **Data injection** | Declarative | Fixed schemas, predefined ontologies, bounded tool returns |
| **Pragmatics** | **The bridge** | How agent interprets declarative data into procedural action |

### Where Pragmatics Fits

**Pragmatics = the interface between declarative data and procedural agent behavior.**

```
Declarative Data → [Pragmatics] → Procedural Agent Action
       ↓                    ↓                ↓
   Schemas,            "How does the      Agent reasoning,
   ontologies,        agent interpret    tool calls,
   tool definitions   this context?"     learned steps
```

Pragmatics is the translation layer:
- **What** the declarative data says (semantics)
- **How** the agent uses it (pragmatics)
- The gap between "what we provided" and "what the agent does with it"

In the book, pragmatics covers:
- Tool call pragmatics: how the agent decides which tool to call
- Context pragmatics: how the agent interprets system prompts
- Retrieval pragmatics: how the agent decides what's relevant

**The insight:** You can give the agent declarative data (schemas, tools, prompts), but what the agent actually **does** with that data is a procedural question. Pragmatics is the study of that gap.

---

## The Three Layers of Context Management

### 1. Declarative = Data Pipeline (Contract-Based)

- Schemas define contracts
- Tool definitions are contracts
- Ontology structures are contracts
- Data pipelines hydrate based on classification

**Focus:** What data is available and how it's structured

### 2. Pragmatics = Prompting and "How"

- How the agent interprets declarative data
- Prompt design that translates contracts into action
- Tool call pragmatics (when to call, how to interpret results)
- The "translation" layer between data and behavior

**Focus:** How declarative data becomes actionable

### 3. Procedural = Agenting System and Infrastructure

- Agent reasoning loops
- Tool execution flow
- Scaffolding and execution traces
- What the agent actually does at runtime

**Focus:** What the agent does with the interpreted data

---

### The Three Are Interdependent

```
Data Pipeline (Declarative) → Prompt Design (Pragmatics) → Agent System (Procedural)
        ↓                            ↓                           ↓
    Contracts                    Translation                  Execution
    Schemas                      Tool interpretation         Reasoning loops
    Ontology                     Context injection           Action selection
```

All three are needed:
- Without **declarative**: no bounded data to work with
- Without **pragmatics**: agent can't translate data into action
- Without **procedural**: nothing executes

The three layers form a complete context management system.

---

## Next Steps

1. Find working WordNet API / Python library
2. Prototype a simple semantic expansion loop
3. Test with a bounded domain (e.g., "software engineering")
4. Compare to vector-based expansion
5. Research learning crawlers and preference extraction
6. Investigate "memoroes" concept