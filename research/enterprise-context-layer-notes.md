# Source Note: "What an Enterprise Context Layer Actually Is"

Author: Prukalpa Sankar  
Publication: Context & Chaos  
Date: June 1, 2026  
URL: https://open.substack.com/pub/contextandchaos/p/what-an-enterprise-context-layer

## Relevance to Context Engineering Book

This article provides a useful enterprise architecture framing for context engineering. It defines an enterprise context layer as the system that turns knowledge, expertise, and norms into machine-usable context for AI agents.

This maps closely to the book thesis:

> Reliable AI systems require engineered context, structured state, semantic constraints, governed retrieval, and evaluation.

## Key Mapping

- Knowledge → data context, canonical facts, knowledge graphs
- Expertise → procedural context, workflows, SOPs, reusable skills
- Norms → authorization, governance, approval paths, policy constraints
- Context layer → governed context substrate for AI systems
- Skills → versioned, testable procedural knowledge
- Learning loops → evaluation, traces, corrections, and promotion into durable state

## Useful Distinction

The article separates the context layer into:

1. Core substrate:
   - AI-ready data and knowledge graph
   - semantics and ontology
   - skills

2. Operating capabilities:
   - context mining
   - context development lifecycle
   - compounding learning loops
   - activation and retrieval
   - governance and observability

This is useful because it distinguishes the material of context from the lifecycle that maintains it.

## How This Supports the Book

The article reinforces that context engineering is not prompt engineering. A context layer is not a vector database, semantic layer, data catalog, or memory store by itself. It is the governed system that makes those components usable by AI systems.

The book should build on this but stay more engineering-specific:

- What are the storage patterns?
- How are permissions enforced?
- How is context retrieved?
- How are changes versioned?
- How are skills evaluated?
- How does context promotion happen?
- How do we prevent context drift?
- How do we measure reliability improvement?

## Skeptical Notes

The "shared enterprise brain" metaphor is useful for market explanation but risky. It can make the architecture sound more magical than it is. A better engineering framing is:

> A context layer is a governed substrate of data, semantics, procedures, permissions, retrieval interfaces, and evaluation loops.

The claim that AI can build and curate ontologies at enterprise speed should be treated as a hypothesis. AI can assist with extraction, conflict detection, and candidate generation, but ownership and approval still need humans.

## Failure Modes This Source Connects To

- Agents retrieving facts without understanding definitions
- Different agents using different versions of business logic
- Prompts becoming ungoverned procedural knowledge
- Memory turning into unverified institutional folklore
- Semantic layers being mistaken for full context infrastructure
- Vector databases being treated as enterprise memory
- Context drift without owners, tests, or rollback
- AI systems acting without permission-aware retrieval

## Book Pattern

Context engineering should treat reusable procedures as governed artifacts, not prompt snippets.

A "skill" should be:

- named
- owned
- versioned
- tested
- permission-scoped
- observable
- replaceable
- callable by agents and workflows

This connects tool calling, memory, retrieval, and evaluation into one systems pattern.