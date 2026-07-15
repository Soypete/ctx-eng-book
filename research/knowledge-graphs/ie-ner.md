# Information Extraction and Named Entity Recognition (NER)

## Core Concepts

### Named Entity Recognition (NER)

NER is a sequence labeling task that identifies and classifies entities in text into predefined categories such as:
- PERSON (people)
- ORGANIZATION
- LOCATION
- DATE/TIME
- MONEY, PERCENT, etc.

**Key insight**: NER uses sequence labeling to identify and classify entities in text — same underlying technique as learning crawlers.

### Information Extraction (IE)

Broader than NER — IE encompasses:
- Named Entity Recognition
- Relationship Extraction
- Event Extraction
- Coreference Resolution

### NER in Knowledge Graphs

NER is a critical step in building knowledge graphs from unstructured text:

1. **Entity Extraction**: Identify entities mentioned in text
2. **Entity Linking**: Connect extracted entities to canonical knowledge base entries
3. **Relationship Extraction**: Identify relationships between entities

### In-Context Learning for NER

LLMs can perform NER in-context with few-shot examples:

```
Input: "John works at Google in Mountain View."
Output: PERSON: John; ORG: Google; LOCATION: Mountain View
```

**Advantages**:
- No training required
- Can adapt to new entity types with examples
- Handles context-dependent entity mentions

**Limitations**:
- May miss entities or misclassify
- Context window limits number of examples
- Consistency can vary across inputs

### Training Custom NER Classifiers

For production systems, training domain-specific NER classifiers is recommended:

> Classifiers are important to data engineering when it comes to graphs. It is the way to process chats and unstructured texts effectively. Whether to use LLMs to create labeled data or train classifiers.

**Approaches**:
1. Use LLMs to generate labeled training data (reduces cold-start problem)
2. Fine-tune existing NER models on domain-specific data
3. Use sequence labeling frameworks (CRF, HMM)

### Two Versions of IE for KG Construction

**1. Web Information Extraction**
- Operates over raw webpages
- Attempts to extract a KG with entities, relations, or even events
- Focuses on structure from web pages (tables, lists, templates)

**2. Named Entity Recognition (NER)**
- Extracts instances of concepts such as PERSON or LOCATION
- Concepts come from an ontology (domain-specific)
- Provides the "nodes" of the knowledge graph

**Relationship Extraction**
- To supplement the instances and interconnect them with relations
- Extracts edges/relationships between entities
- Builds the full graph structure

```
Web IE → Full KG (entities + relations + events)
NER → Entity instances + Ontology concepts
Relationship Extraction → Connects NER instances into graph
```

---

### LLMs Enabling Novel Entity Extraction and Classification into Ontology

Traditional NER relies on pre-defined entity types (PERSON, ORG, LOCATION). LLMs enable a more powerful approach:

**Unnamed Entity Extraction**:
- LLMs can identify entities that don't match predefined categories
- Can extract novel entity types based on context and domain knowledge
- Example: Extract "authentication failure" as an EVENT type even if not in training data

**Classification into Ontology**:
- Given an ontology, LLMs can classify extracted entities into the correct concept classes
- Uses in-context learning to understand the ontology structure
- Can handle hierarchical ontology relationships (e.g., PERSON > RESEARCHER > AI_RESEARCHER)

**Workflow**:
```
Raw Text → LLM extracts candidate entities → LLM classifies into ontology concepts → Knowledge Graph
```

**Advantages over traditional NER**:
- No retraining needed for new entity types
- Can leverage ontology structure for zero-shot classification
- Handles context-dependent entity classification

**Limitations**:
- Requires clear ontology definition in prompt
- May hallucinate entity types outside the ontology
- Consistency can vary across large document sets

NER-style sequence labeling connects to learning crawlers:

- **Context graph nodes** = pages, entities, relationships (labeled via NER-style sequence labeling)
- **Agent run nodes** = tool calls, decisions, state changes (labeled via execution trace)

---

### Tasks Redefined Through Ontologies

In knowledge graph construction, ontologies can redefine tasks by constraining what entities and relationships are valid:

- **Task definition via ontology**: The ontology defines what the agent should extract and how entities relate
- **Predefined problem space**: Instead of letting the agent discover entities freely, the ontology bounds the extraction task
- **Example**: A "customer" ontology defines that extraction should find PERSON, ORG, and their relationships — not arbitrary entities

> Make the context graph a predefined problem — the ontology defines where the agent should go and find data.

**Trade-off**:
- Ontologies provide structure and consistency
- But can miss entities outside the defined schema
- Alternative: Use procedural approaches without predefined ontologies (see web-crawlers.md)

---

### IE with LLMs as Knowledge Distillation

IE using LLMs can be viewed as **knowledge distillation** — extracting and structuring knowledge that exists in the LLM's training data into a explicit knowledge graph:

- **What was implicit**: Knowledge encoded in LLM weights (learned from web-scale text)
- **What becomes explicit**: Structured entities, relationships, and facts in a KG
- **Process**: LLM acts as a "knowledge extractor" — reading raw sources and outputting structured data

**Why "distillation" fits**:
- LLMs have seen massive amounts of unstructured text during training
- IE prompts the LLM to "distill" relevant knowledge into structured form
- The KG is a compressed, explicit representation of the LLM's knowledge

**Trade-off**:
- KG provides determinism and governability that raw LLM outputs lack

### LLMs Help with NER — Background Knowledge and "Poverty of the Stimulus"

Traditional NER struggles with implicit information that requires background knowledge:

> "Significant amounts of relevant information could be implicit and difficult to discern without a requisite (and enormous) amount of background knowledge that humans have managed to acquire despite a 'poverty of the stimulus.' In other words, we are able to use, learn, compose, and understand sentences in very creative ways, despite not having heard the vast majority of sentences we end up using and understanding."

**Why LLMs help**:
- LLMs have acquired background knowledge during training on web-scale text
- They can infer implicit entities and relationships that require world knowledge
- Can handle context-dependent entity mentions

**The challenge**:
- Traditional NER models trained on predefined ontologies cannot be easily transferred
- Genre changes (newswire vs. social media) or language changes break models
- LLMs with in-context learning can adapt without retraining

**Ontology Changes Require New NER Models**:
> If the ontology changes or expands, you need to train a new NER model.

This is a key limitation:
- Traditional NER is tied to fixed entity classes
- Adding new entity types requires retraining

**This is a pre-LLM statement. LLMs allow for domain expansion:**

- **In-context learning**: Give the LLM examples of the new entity type in the prompt, and it can immediately recognize them — no retraining needed
- **Zero-shot extraction**: LLMs can extract novel entity types without any examples, using only the ontology description
- **Cross-domain transfer**: An LLM trained on general text can apply knowledge to specialized domains (medical, legal, scientific) without fine-tuning
- **Genre adaptation**: LLMs handle social media, newswire, conversational text differently — same model, different prompts

The key difference: traditional NER uses **closed vocabulary** (learned during training), while LLMs use **open extraction** (can recognize anything they can describe in context).

### Supervised vs Unsupervised NER Performance

> NER is more computationally efficient and extraction difficult. Open IE and unsupervised NER, which require no ontology and/or training instances, still lag significantly in performance compared to supervised, ontologically mediated NER.

**Key insight**:
- **Supervised NER**: More accurate, but requires labeled training data and fixed ontology
- **Unsupervised / Open IE**: No ontology or training needed, but lower performance

**Modern update**:
- Performance of unsupervised NER has improved steadily over the past decade
- Improvements in language models and self-supervised learning have closed the gap
- LLMs blur the line — they can do zero-shot extraction without training while approaching supervised quality

### LLMs Enable Multi-lingual NER

> What LLMs enable is classifying and processing multi-lingual data.

- **Cross-lingual transfer**: Train on one language, apply to another without parallel data
- **Single model, many languages**: One LLM can do NER in dozens of languages
- **No language-specific models**: Traditional NER required separate models per language; LLMs handle all in one
- **Low-resource languages**: LLMs can extract entities from languages with limited labeled training data

This is a key advantage over traditional NER systems that required language-specific models and training data.

- Context graphs (Diligenti et al., 2000)
- NER with transformers
- Few-shot learning for NER
- LLM-based information extraction

### Traditional NER: Supervised Sequence Labeling

The current popular paradigm for NER is supervised learning with sequence-labeling techniques:

**Traditional Models**:
- Decision trees
- Maximum entropy models
- Hidden Markov Models (HMMs)
- Support Vector Machines (SVMs)
- Conditional Random Fields (CRFs)

**Why Sequence Labeling Matters**:
- Traditional classification treats each word independently
- Problem: "the" followed by "United States" — classifying "the" as "O" doesn't help predict "United States" as "LOC"
- Sequence labelers (HMMs, CRFs) assign output states to input terms **without making strong independence assumptions**
- They consider dependencies between adjacent tokens

> "We should not be classifying every token independently, but if possible, classifying the sequence as a whole."

**CRFs** are particularly important for NER — they model the entire sequence jointly rather than individual tokens.

### CRFs Applied to NER vs LLMs

**CRFs (Conditional Random Fields)**:
- Apply probabilistic modeling to sequence labeling
- Explicitly model dependencies between adjacent labels
- Well-established for NER tasks

**LLMs can do NER**, but:
- Need data on computational cost comparison
- LLMs are much more computationally expensive than CRFs
- For well-defined entity types, CRFs may be more efficient

**Comparison**:

| Aspect | CRF | LLM |
|--------|-----|-----|
| Inference Speed | Milliseconds | Seconds |
| Memory | Small matrices | Billions of parameters |
| Cost per token | Very low | Higher |
| Novel entities | No (fixed labels) | Yes (zero-shot) |

**When to use each**:
- **CRFs**: Well-defined entity types, low latency requirements, cost-sensitive
- **LLMs**: Novel entity types, context-dependent extraction, flexibility over efficiency

### Active Learning with NER — Use Case for Small Parameter Models

> Active learning with NER is a use case for small parameter models.

**What is Active Learning**:
- Start with small set of human-annotated examples
- Model actively decides which examples to present to human next for maximal gain
- Usually: samples where the learner's prediction has greatest uncertainty
- Human annotator gives the system data that benefits it most

**Why this favors small models**:
- **Efficient updates**: Small models can be retrained quickly with new labeled data
- **Human-in-the-loop**: Small periodic interventions from annotators
- **Annotation efficiency**: Preempts labeling of redundant samples the system can already label with high certainty

> "Active learning is one example of 'human-in-the-loop' learning that has recently witnessed resurgence, because it represents a hybrid situation where an appropriate balance of system design, data labeling, knowledge engineering, and human intervention together lead to effective performance."

**Use case**: When you have limited labeled data and want to build a NER model iteratively with human feedback — small parameter models (CRFs, smaller neural models) are more practical than LLMs.

### Neural NER — RNNs and Character Embeddings

> Several NLP problems like POS tagging and NER in a unified feature framework, neural NER systems with minimal feature engineering have gained in popularity. Such models are appealing because they do not normally require domain-specific resources like lexicons or ontologies and can scale more easily without significant manual tuning.

**Neural NER architectures**:
- Based on RNNs (Recurrent Neural Networks) over word, subword, and character embeddings
- **RNN definition**: A network with loops that allows information to persist. Can be unrolled over a sequence.
- Chain-like structure makes them suited for sequences: speech recognition, language modeling, translation, image captioning, NER

**Advantages over traditional NER**:
- Minimal feature engineering required
- No domain-specific resources (lexicons, ontologies) needed
- Scale more easily without manual tuning

**Similar to CRFs**: RNNs accept input vectors and model sequences, but with learned representations instead of explicit feature functions.

### Data Labeling from Agent Actions and Conversations

> Something specific is how data labeling should be justified and grappled from agent actions and tool conversations.

**Key insight**: Instead of manual annotation, use agent execution traces as labeled data:

- **Agent actions**: Tool calls, decisions, state changes become labeled examples
- **Conversations**: User interactions provide natural labeled data
- **Execution traces**: "This query led to this tool being called, which retrieved this data"

**Why this matters**:
- Reduces need for human annotation
- Labels come from actual system behavior
- Can bootstrap NER models from agent interaction data

> Context graph nodes = pages, entities, relationships (labeled via NER-style sequence labeling)
> Agent run nodes = tool calls, decisions, state changes (labeled via execution trace)

### NER Evaluation Metrics

> To quantify the success of NER in data pipelines, the knowledge graph book provides:

**Core Metrics (from IR)**:
- **Precision** = #correct / (#correct + #incorrect) — measures correctness
- **Recall** = #correct / #total — measures completeness
- **F1** = 2 × Precision × Recall / (Precision + Recall) — weighted harmonic mean

**Key distinctions**:
- **Spurious slot**: System produces a slot that doesn't align with gold standard
- **Invalid value**: Slot aligns but has wrong value
- **Missing slot**: Gold standard slot not produced by system

**IE-Specific Metric**:
- **Slot Error Rate (SER)** = (#incorrect + #missing) / #total

**Customization**:
- F-measure can weight precision vs recall (β parameter)
- Metrics can be measured per slot type
- Validation sets needed to tune hyperparameters

---

## Evidence Ledger Entries

See `research/_evidence-ledger.md`:
- Classifier Training for Unstructured Text in Knowledge Graphs

---

## Cross-References to Context Engineering Book

### Chapter 4: In-Context Learning and Pragmatics

This research connects to Chapter 4:

- **ch04.01 In-Context Learning**: "LLMs enabling novel entity extraction" and "multi-lingual NER" support the in-context learning discussion — LLMs can extract novel entities without retraining
- **ch04.02 Computational Pragmatics**: "IE as knowledge distillation" connects to "do the inference work" — extracting structured data from LLMs rather than having the model infer
- **ch04.03 Examples, Instructions, Structured Outputs**: "NER evaluation metrics" (precision, recall, F1) supports the "Measure and tune" section — prompt performance is measurable using the same metrics as NER

### Related Chapters

- **Chapter 7 (Context Assembly)**: IE/NER is a core component of context assembly pipelines
- **Chapter 8 (Knowledge Graphs)**: NER is the first step in building KGs from unstructured text