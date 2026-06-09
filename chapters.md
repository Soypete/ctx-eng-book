# Context Engineering: Building Reliable AI Systems

## Working Thesis

Reliable AI systems do not emerge from larger models alone.

They are built through engineered context, structured state, semantic constraints, governed retrieval, authorization-aware data access, and continuous evaluation.

The primary challenge in production AI is not generating tokens. The challenge is ensuring the model receives the correct information, at the correct time, in the correct scope, for the correct task.

Context engineering is the discipline responsible for solving that problem.

---

## Definition of Context Engineering

Context engineering is the discipline of designing, governing, retrieving, constraining, and evaluating the information used by AI systems.

Context includes far more than prompts. It encompasses:

- Retrieved knowledge
- Structured state
- User and organizational data
- Authorization policies
- Tool outputs
- Memory systems
- Semantic relationships
- Workflow state
- External system interactions

Context engineering treats context as infrastructure rather than prompt text.

Its goal is to make AI systems more reliable, predictable, secure, observable, and cost-effective.

---

# Part I: Why AI Systems Fail

## Chapter 1: The Reliability Problem

### Why AI Feels Magical in Demos
- The narrow scope of demonstrations
- Curated examples versus production reality
- The illusion of capability

### Why Production Is Different
- Open-ended user behavior
- Data quality problems
- Organizational complexity

### Hallucinations Are Not the Root Problem
- Hallucination as a symptom
- Missing information
- Retrieval failures
- State failures

### The Reliability Gap
- Why larger models do not solve operational problems
- Failure modes observed in production systems
- Lessons from software engineering

### Context as Infrastructure
- The central argument of the book
- Why context determines reliability

---

## Chapter 2: What Is Context?

### Beyond the Prompt
- Why prompts are only one component
- The misconception of prompt engineering

### The Anatomy of a Context Packet
- Instructions
- Retrieved knowledge
- State
- Tool outputs
- User information

### Sources of Context
- Databases
- APIs
- Documents
- Tools
- Memory systems

### Context Lifecycles
- Creation
- Retrieval
- Transformation
- Expiration

### Context Failures
- Missing context
- Incorrect context
- Excessive context
- Unauthorized context

---

## Chapter 3: The Five Pillars of Context Engineering

### Retrieval
- Acquiring relevant information

### State
- Maintaining continuity

### Constraints
- Limiting ambiguity

### Governance
- Controlling access and provenance

### Evaluation
- Measuring effectiveness

### Interactions Between the Pillars
- Why reliability requires all five

---

# Part II: Retrieval

## Chapter 4: Search Before AI

### The History of Information Retrieval
- Libraries
- Search engines
- Enterprise search

### Ranking and Relevance
- Precision
- Recall
- Relevance scoring

### BM25 and Lexical Search
- How traditional search works
- Strengths and weaknesses

### Why Retrieval Still Matters
- Search as the foundation of AI systems

### Retrieval Failure Modes
- Missing information
- Poor ranking
- Stale results

---

## Chapter 5: Vector Search and Its Limits

### Embeddings Explained
- Semantic representations
- Similarity spaces

### Similarity Metrics
- Cosine similarity
- Dot product
- Euclidean distance

### Building a Vector Index
- Chunking
- Embedding generation
- Retrieval workflows

### Failure Modes of Vector Search
- Semantic drift
- False similarity
- Missing context

### Hybrid Retrieval
- Combining lexical and semantic search
- Why hybrid systems often perform better

---

## Chapter 6: Knowledge Graphs and Semantic Retrieval

### Why Relationships Matter
- Entities versus documents

### Semantic Web Foundations
- RDF
- Triples
- Linked data

### Ontologies
- Classes
- Relationships
- Constraints

### Knowledge Graph Retrieval
- Traversal
- Semantic querying
- Context expansion

### Semantic Constraints and Reliability
- Reducing ambiguity
- Improving retrieval precision

---

# Part III: State and Memory

## Chapter 7: Agents Need State

### Stateless Models
- What LLMs actually are

### Stateful Systems
- Why applications need memory

### Workflow State
- Tasks
- Processes
- Execution tracking

### Event Sourcing Concepts
- Immutable logs
- Replayability
- Auditing

### State Failure Modes
- Lost state
- Corrupted state
- Divergent state

---

## Chapter 8: Memory Is Retrieval

### The Myth of AI Memory
- What models remember
- What systems remember

### Short-Term Memory
- Session history
- Conversation context

### Long-Term Memory
- Persistent storage
- Retrieval strategies

### Personalization
- User profiles
- Preferences
- Behavioral context

### Memory Failure Modes
- Memory contamination
- Incorrect recall
- Privacy risks

---

## Chapter 9: Context Windows and Token Economics

### Understanding Context Windows
- Limits of transformer architectures

### Token Budgets
- Cost considerations
- Performance implications

### Context Compression
- Summarization
- Distillation
- Aggregation

### Context Prioritization
- Relevance ranking
- Information density

### The Economics of Context
- Token waste
- Retrieval costs
- Operational efficiency

---

# Part IV: Constraints and Governance

## Chapter 10: Authorization Is Context

### Security Failures in AI Systems
- Oversharing
- Over-permissioned retrieval

### Authentication and Authorization
- Identity
- Permissions
- Access control

### RBAC and ABAC
- Role-based access
- Attribute-based access

### Permission-Aware Retrieval
- Filtering before generation

### Least Privilege for AI
- Reducing risk through scoped context

---

## Chapter 11: Semantic Constraints

### Why Constraints Matter
- Reliability through structure

### Taxonomies and Controlled Vocabularies
- Standardizing meaning

### Ontologies in Production
- Semantic modeling
- Domain representation

### Derived Domain-Specific Ontologies
- DDSOs
- Knowledge hydration

### Constraints as Guardrails
- Preventing invalid outputs
- Improving retrieval accuracy

---

## Chapter 12: Tool Use and Context Boundaries

### Tools as Context Sources
- APIs
- Databases
- External systems

### Tool Calling Architectures
- Planning
- Execution
- Validation

### Model Context Protocol
- MCP concepts
- Tool interoperability

### Context Leakage
- Cross-tool contamination
- Data exposure risks

### Designing Safe Tool Systems
- Validation
- Sandboxing
- Auditing

---

# Part V: Reliability Engineering for AI

## Chapter 13: Observability

### Why AI Systems Need Observability
- Debugging non-deterministic systems

### Logging Context
- Inputs
- Retrievals
- Outputs

### Tracing Agent Behavior
- Tool calls
- Decision paths

### OpenTelemetry and AI
- Standardized instrumentation

### Diagnosing Context Failures
- Root cause analysis

---

## Chapter 14: Evaluation

### Why Evaluation Matters
- Reliability requires measurement

### Types of Evals
- Retrieval evals
- Generation evals
- Agent evals

### Building Evaluation Harnesses
- Test cases
- Benchmarks
- Reproducibility

### Measuring Context Quality
- Relevance
- Completeness
- Accuracy

### Continuous Evaluation
- Regression testing
- Production monitoring

---

## Chapter 15: Building Reliable AI Systems

### The Context Engineering Architecture
- Bringing the pillars together

### Designing a Context Pipeline
- Retrieval
- State
- Constraints
- Governance

### Reference System Design
- End-to-end architecture

### Reliability Patterns
- Proven approaches
- Anti-patterns

### The Future of Context Engineering
- Agents
- Knowledge systems
- Semantic infrastructure

### Final Thesis
- Reliability emerges from engineered context, not larger models