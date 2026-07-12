# Retrieval Beyond Vector Databases

## Beat 1: The Retrieval Problem

Every production AI system faces the same fundamental question: how do we retrieve the right information? Not the most similar information. Not the most semantically related. The *right* information—the exact context needed to produce a reliable response.

Vector databases have become the default answer. Embed your data, store the vectors, query by similarity, return the top-k results. This works surprisingly well for semantic similarity tasks—finding passages that mean something similar to your query. But reliability demands more. It demands precision, completeness, and the ability to retrieve specific facts, not just semantically related ones.

The vector database alone cannot answer "what is the user's billing tier?" or "which tools does this agent have access to?" or "what were the three actions taken in the previous session?" These are precise lookups, not semantic searches. They're the queries that determine whether your AI system works correctly or fails in ways that seem inexplicable.

This chapter argues that reliable retrieval requires a hybrid approach—combining semantic search with keyword search, graph traversal, structured queries, and metadata filtering. The vector database is one tool in the retrieval toolbox, not the whole toolbox.

## Beat 2: The Limits of Semantic Similarity

Vector search answers a specific question: which documents are semantically similar to this query? This is useful. It's not sufficient.

Consider a legal document system. A query for "breach of contract" should return contract breach clauses—not articles about data breaches, security incidents, or ethical violations. Semantic similarity might surface any of these. The vector model learned that "breach" relates to security incidents, data leaks, and contract violations. It doesn't know which one you need.

Consider a support ticket system. A query for "login failed" should return login failure incidents—not all incidents containing the word "login" or semantically related to authentication. Vector search might return password reset requests, account lockout issues, or two-factor authentication problems. The retrieval is "close" but not correct.

Consider an agent that needs to check its own permissions. A query for "what tools can I use?" should return exactly the tool definitions the agent was granted—not similar tools, not tools other agents use, not tools mentioned in documentation. This is a precise lookup, not a similarity search.

In each case, the failure isn't that the vector database returns bad results. It's that semantic similarity is the wrong retrieval primitive for the task. The system needs keyword matching, exact lookup, graph traversal, or structured filtering—and vector search alone cannot provide these.

## Beat 3: BM25 and Keyword Search

BM25 is a ranking function used by full-text search engines like Elasticsearch and OpenSearch. It scores documents based on term frequency, inverse document frequency, and document length normalization. It's older than vector embeddings. It's still essential.

Keyword search excels at exact matching. When you need documents containing "Order-2024-001" or "GDPR compliance" or "emergency shutdown procedure," BM25 finds them. The query terms appear in the documents. There's no ambiguity about whether semantic similarity qualifies.

A practical retrieval pipeline often starts with keyword search. The user query is parsed for exact-match terms—IDs, proper nouns, technical phrases—and used to filter or boost results. Only then does semantic similarity rank what remains.

This is particularly important for domain-specific terminology. A vector model trained on general text may not distinguish between "ticket" as a support request and "ticket" as a travel authorization. BM25 treats them as different terms. It respects the exact vocabulary your system uses.

The hybrid pattern is straightforward: run keyword queries and vector queries in parallel, combine the results, rerank by a learned model or simple weighted scoring. This gives you exact-match precision with semantic generalization.

## Beat 4: Graph Traversal

When information has relationships, graph traversal retrieves what keyword and vector search cannot: connected context.

Consider a knowledge graph representing a company's organizational structure, asset hierarchy, and causal dependencies. A query about "server outages affecting payment processing" requires traversing from the payment service to its dependencies, then to the servers hosting those dependencies, then to the outages affecting those servers. This is a graph query—starting from payment-processing, following dependency edges, collecting related incidents.

Vector search cannot express this. There's no semantic similarity between "payment processing" and "server X" that captures the dependency relationship. The connection is structural, not semantic.

Graph traversal also handles temporal reasoning. "What happened after the deployment?" requires ordering—following edges that represent temporal sequence. A knowledge graph with time-indexed edges supports queries like "find all downstream effects of event X."

Practical graph retrieval in AI systems works as follows: identify the starting nodes from structured queries or vector search, then traverse along relationship edges to collect connected context. The depth and direction of traversal are parameters you control based on how much context you need.

This is essential for agentic systems. Agents act on structured state—task graphs, dependency chains, authorization hierarchies. Retrieving "what can I do?" requires traversing the permission graph. Retrieving "what's the current task state?" requires traversing the task dependency graph. These are not semantic queries. They're graph queries.

The performance impact is substantial. On the GBrain benchmark, a four-strategy hybrid (vector + BM25 + RRF + graph) achieves 49.1 P@5. Each component alone—ripgrep BM25 only, vector-only RAG, or hybrid without graph—sits around 18 P@5. Graph traversal contributes +31 points. It's the difference between random performance and production-ready retrieval.

## Beat 5: Hybrid Retrieval Architectures

The reliable pattern is hybrid retrieval: multiple retrieval primitives combined into a single pipeline that produces the right context for each request.

A typical architecture:

1. Query parsing: Extract structured filters (date ranges, user IDs, document types), keyword terms, and semantic query components from the user's request.

2. Structured query: Run exact matches against structured stores—user preferences from Redis, permissions from a policy database, IDs from a relational store. These return precise values, not documents.

3. Keyword search: Run BM25 against full-text indexes. Filter by the structured constraints from step 1. Return the top-k keyword-matched documents.

4. Vector search: Run the semantic component against vector indexes. Filter by the structured constraints. Return the top-k semantically similar documents.

5. Graph traversal: If the query requires relationship context, traverse from the retrieved nodes along relationship edges. Collect connected context.

6. Reranking: Combine results from multiple sources, score by relevance, filter duplicates, return the final ranked list.

Each stage is optional. A simple system might use only structured queries. A complex system uses all five. The point is that you choose the retrieval primitive based on what you're trying to retrieve, not because one primitive is universally best.

Evaluating hybrid retrieval requires benchmarks that measure retrieval quality, not just semantic similarity. Context-Bench by Letta tests how well LLMs chain file operations, trace relationships across sessions, and manage long-horizon retrieval—the core skills context engineering demands. The best model (Claude Sonnet 4.5) scores 74% on the benchmark. This gap signals that context engineering is a distinct competency, not something that emerges automatically from larger models.

## Beat 6: Semantic Indexing and Knowledge Stores

Semantic indexing goes beyond simple vector embedding. It structures the indexed representation to support precise retrieval.

Consider a support knowledge base. Instead of embedding every article as a single vector, you embed each article section, each FAQ answer, each procedure step as separate chunks. You index them with metadata: product version, issue category, user role, severity level. The retrieval query includes these filters. The vector search operates on a filtered subset, not the entire corpus.

This is semantic indexing—embedding with structure. The structure supports filtering that pure vector search cannot express. You don't just find semantically similar content. You find semantically similar content that applies to the right product version, the right user role, the right issue category.

Knowledge stores formalize this pattern. A knowledge store is a retrieval system that maintains structured metadata alongside embeddings, supports structured filtering alongside semantic search, and exposes query interfaces that combine both. Think of it as a document store with a semantic search layer on top.

The practical implication: design your indexing schema before you design your embedding strategy. Decide what metadata you need to filter on—user type, time period, document type, authorization scope—then build the vector index to support those filters. The retrieval query is only as precise as the index allows.

Mem0 approaches this through hierarchical distillation—storing user interactions at multiple granularity levels (user, session, message) and querying across all levels for context. This is semantic indexing as a retrieval architecture: maintaining structured relationships between context chunks at indexing time so that traversal at query time finds the right level of detail. Mem0's LoCoMo benchmark score (92.5) and LongMemEval (94.4) demonstrate the approach works for personal memory; BEAM 1M (64.1) and BEAM 10M (48.6) show temporal reasoning at scale remains unsolved—a gap for systems that need multi-year context spans.

## Beat 7: The Reliability Question Answered

How do we retrieve the right information? By building a retrieval system that combines the right primitives for each query type.

Vector search alone cannot answer precise questions. BM25 cannot understand context. Graph traversal cannot handle unstructured content. Each primitive has a domain. The reliable system uses all of them.

The architecture is straightforward: parse the query into components, route each component to the appropriate retrieval primitive, combine results, return context. This is database engineering—the same discipline that powers traditional applications—but with a language model as the consumer.

The reliability question has a practical answer: retrieve the right information by building retrieval that knows what "right" means for each query. Semantic similarity is not enough. You need keyword matching, graph traversal, structured filtering, and semantic search—combined into a pipeline that produces precise, complete, reliable context.

Your AI system is only as reliable as its retrieval. Build retrieval that matches.

(End of file)