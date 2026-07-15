# Web Crawlers for Knowledge Graphs

## Core Concepts

### Focused Crawling

Focused crawling (topic-specific crawling) is a more common term than "domain discovery" for web crawlers that target specific topics or domains.

- **Definition**: Crawlers that prioritize pages relevant to a specific topic rather than crawling the entire web
- **vs. General Crawling**: Focused crawlers use classifiers to determine relevance and follow links strategically

### Classifiers in Web Crawlers

Classifiers are important to data engineering when it comes to knowledge graphs. They are the way to process chats and unstructured texts effectively.

- **Use case**: Whether to use LLMs to create labeled data or train classifiers
- **Training**: Train classifier on pages that are known relevant (seed set), then assign new pages to layers based on their predicted relevance

### Context Graph Crawlers

> Crawlers in context graphs ask for immediate parents/children to navigate graphs and gain information. This is a non-reasoning approach to graph navigation.

- **Key paper**: Diligenti et al. (2000) — "Focused Crawling Using Context Graphs"
- **Approach**: Builds classifiers for sets of pages mainly at distance 1 or 2 from relevant pages
- **Method**: Uses Hidden Markov Models to browse using sequence labeling and context-focused crawlers

### Domain Discovery Tool (DDT)

**Krishnamurthy et al. (2016)** — "Interactive Exploration for Domain Discovery on the Web" (IDEA Workshop at KDD 2016)

- **Definition**: Visual analytics framework for interactive domain discovery
- **Approach**: Augments ordinary search engine functionality, supports exploratory data analysis of webpages
- **ML Methods**: Uses naive Bayes classifiers and sequences — crawlers are NLP-based
- **Key insight**: Translates analyst interactions with web data into a computational model of the domain of interest
- **Relation to focused crawling**: DDT relates to focused crawling but adds interactive visual analytics

### Deterministic Procedural Approaches

> Procedural approach where we don't have ontologies. DDT (focused crawling) might be a more deterministic way to explore without LLM sprawl.

**Advantages**:
- More deterministic behavior
- Avoids LLM "sprawl" (uncontrolled graph expansion)
- No ontology maintenance overhead

**Limitations**:
- May lack semantic richness of ontology-based approaches
- Requires analyst interaction (DDT)

### LLMs Augmenting Focused Crawling

> LLMs have a lot of the information in their training set, so in theory can augment DDT if it is providing queries.

**Potential**:
- LLMs can generate queries for focused crawlers
- Can provide domain knowledge from training data

**Risks**:
- Training data may be outdated
- Hallucination risk

### Legacy Algorithms

**PageRank and HITS**:
- Common and old web search algorithms
- LLMs do not leverage these
- Modern web search does not prominently use these either

**Note**: Modern search engines may use link signals implicitly; graph-based retrieval could reintroduce these algorithms.

---

## Related Research

- Context graphs (Diligenti et al., 2000)
- Focused crawling literature
- DDT (Krishnamurthy et al., 2016)
- Modern LLM-augmented retrieval

## Evidence Ledger Entries

See `research/_evidence-ledger.md`:
- Classifier Training for Unstructured Text in Knowledge Graphs
- Context Graph Crawlers — Non-Reasoning Graph Navigation
- DDT — Deterministic Procedural Focused Crawling
- Procedural Domain Discovery — Deterministic Without Ontologies
- LLMs Augmenting DDT — Query Generation from Training Knowledge
- PageRank and HITS — Legacy Algorithms Not Used by Modern LLMs