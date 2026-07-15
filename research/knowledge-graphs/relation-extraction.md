# Relation Extraction (KG Book Chapter 6)

## Core Concepts

### RE Is More Complex Than NER

**Why relation extraction is harder:**

- **NER**: Identify entities within a sentence (local)
- **RE**: Understand the relationship between entities, often requiring full sentence context

Example:
```
"John gave Mary a book"
```
- NER: John (PERSON), Mary (PERSON), book (OBJECT)
- RE: John → gave → Mary (giver-receiver relationship), John → has → book (ownership)

**The complexity:**
- Same entities can have multiple relationship types
- Relationship is often implicit (not explicitly stated)
- Requires understanding verb semantics, prepositions, clauses
- Context matters: "Apple" (fruit) vs "Apple" (company)

LLMs help by understanding context — but extraction still requires careful prompt design to capture the right relationship type.

### Ontologies Are FOUNDATIONAL for RE

**Correction from earlier deduction:**

The book clarifies: unlike NER, **ontologies are foundational for RE** because:

- **NER**: Entity types are often clear (person, org, location)
- **RE**: What constitutes a "relationship" or "event" is not always clear without ontology constraints

**Why ontologies matter more for RE:**

- Without an ontology, it's ambiguous what counts as a relation
- Same text can imply many relationship types
- Ontology provides the frame to interpret what's happening

**Example:**
```
"John and Mary met in 2020"
```
- With ontology: PERSON-PERSON met (EVENT relation)
- Without ontology: ambiguous — is it a meeting? A date? A collaboration?

This is the opposite of my earlier deduction — ontologies PREVENT sprawl BY providing the constraints that make RE tractable.

### ACE: Automatic Content Extraction

**ACE** (Automatic Content Extraction) is a foundational ontology and NIST standard for RE:

- **Developed by NIST**: National Institute of Standards and Technology
- **Defines relation types**: Standardized set of entity relationships
- **Used for evaluation**: Benchmark for extraction systems
- **Covers**: Entities, relations, events

**ACE relation types include:**
- Physical relations (located, near)
- Personal-social relations (spouse, sibling, parent)
- Employment/organizational relations (employer, member)
- Actor-action relations (agent, patient)

**Why ACE matters:**
- Provides a ready-made ontology for RE
- Enables comparison across systems
- Industry-accepted standard

### Traditional RE vs LLM RE: Cost Trade-off

**Traditional approach:**
- Trained supervised/semi-supervised models (CRF, SVM, neural)
- High upfront cost: labeled data, model training, pipeline development
- Lower long-term cost: inference is fast, cheap, deterministic

**LLM approach:**
- Zero/few-shot prompting, no training required
- Low upfront cost: just write prompts
- Higher long-term cost: per-inference API costs, latency

**The trade-off:**
- Traditional: High capex, low opex
- LLM: Low capex, high opex

This is a classic build vs buy tradeoff — applies to RE as it does to all ML tasks.

### WordNet: Semantic Resource for RE

WordNet is already noted in this project:

- **See**: `research/wordnet-anchor-text-notes.md`
- **What it is**: Large lexical database of English words (synsets, relationships)
- **RE connection**: Provides semantic relationships (synonym, hypernym, hyponym, meronym)
- **As guardrails**: WordNet provides verified relationships — unlike vector similarity which can hallucinate

WordNet can constrain RE by:
- Disambiguating word senses
- Expanding relation types to known semantic relationships
- Providing provenance for extracted relations

**This connects to Chapter 5's context engineering:**
> "Make the context graph a predefined problem — the ontology defines where the agent should go and find data."

### LLMs Don't "Do" RE — They Process and Suggest

**A clarifying point:**

LLMs don't autonomously perform relation extraction. They:

1. **Process data**: Read text, identify candidate relations
2. **Suggest options**: Output relationship suggestions
3. **Require human in the loop**: For validation, expansion, ontology updates

**The ontology-as-config approach:**

> If you can generate the ontology via configuration, you make it easier to:
> - Automate ontology creation
> - Audit ontology changes
> - Enable NER/RE pipelines consistently

**This enables a "shift left" motion:**
- Move semantics/knowledge work earlier (design-time)
- Define ontology in config, not at runtime
- Keep data contracts clean
- Human validates suggestions to expand domain

```
Config (ontology) → NER/RE Pipeline → Graph (governed)
                              ↓
                    Human-in-the-loop for expansion
```

**The benefit:** Prevents sprawl by keeping the ontology as the source of truth, while LLMs assist by processing data and suggesting controlled expansions.

---

## Evidence Ledger

See `research/_evidence-ledger.md`:
- (To be added)

---

## Cross-References

- **Chapter 5**: Context Engineering — relation extraction as data supply
- **Chapter 7**: Context assembly — relations as structured context
- **ie-ner.md**: NER provides entities, relation extraction provides edges
- **webinformationextraction.md**: IE pipeline includes relation extraction