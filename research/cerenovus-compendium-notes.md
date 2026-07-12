# Cerenovus Compendium — Research Notes

## Primary Source

- **Compendium: Shared Claude Code Sessions and Team Memory for AI Agents**
  - Cerenovus (YC S26)
  - https://cerenovus.ai/compendium-info
  - Published: 2026

## Related Context

- Cerenovus is building "the AI company brain" — AI that reads everything a company produces
- Y Combinator S26 batch
- Based in San Francisco

---

# 1. What Problem It Solves

## The Core Problem

**AI agents and new teammates start from zero.**

Every time a new Claude Code session begins, or a new engineer joins, or a new agent is spun up, the system has no knowledge of:

- The current state of the codebase
- Decisions already made
- Customer context and history
- Team norms and conventions

This creates massive inefficiency:

- Repeated context reconstruction
- Inconsistent understanding across team members and agents
- Slow onboarding
- Agents acting without full context

## Compendium's Solution

> Compendium keeps one current, shared picture of the codebase, decisions, and customers, so onboarding an engineer or spinning up an agent starts from everything the team already knows.

Key capabilities:

1. **Shared Claude Code sessions** — Team members can share Claude Code sessions, not just their own isolated sessions
2. **Shared context for AI agents** — Agents read from and write to a shared knowledge base
3. **Durable team memory** — Context survives beyond a single session or a single machine
4. **Real-time collaborative knowledge base** — A "live" knowledge base that everyone (humans and agents) can read and write
5. **Onboarding from collective knowledge** — New engineers start with everything the team already knows

---

# 2. How It Handles Memory/Context

## Session-Level Memory

Traditional AI assistants (including Claude Code) operate in isolated sessions:

- Each session starts with minimal context
- Memory is ephemeral — it dies when the session ends
- Context does not transfer between machines or team members

## Compendium's Approach

Compendium provides a **durable team memory layer** that sits between individual sessions and the team's collective knowledge:

- **Persistent across sessions**: Memory outlives any single Claude Code session
- **Persistent across machines**: Not tied to a specific local environment
- **Shared across team**: Everyone reads from and writes to the same memory
- **Agent-accessible**: AI agents can both read (to get context) and write (to record findings)

This is essentially a **context substrate** — a governed layer of structured memory that persists and is shared.

---

# 3. Knowledge Storage and Retrieval

## Storage Model

Compendium appears to be a **live knowledge base** — a real-time, collaborative system that:

- Stores team knowledge, decisions, and context
- Is continuously updated as team members and agents work
- Provides both read and write access to humans and agents

This is different from:

- **Static documentation** — which becomes stale and is often out of date
- **Vector databases** — which store embeddings but lack write semantics and governance
- **Simple memory plugins** — which are typically individual, not team-based

## Retrieval Model

The system supports:

- **Context retrieval** — Agents and team members retrieve relevant context when needed
- **Real-time updates** — The knowledge base is "live," not a periodic snapshot
- **Shared access** — Everyone works from "the same current picture"

## What Gets Stored

Based on the product description, Compendium stores:

- **Codebase context** — Current state of the code, architecture decisions
- **Decisions** — Team decisions, reasoning, outcomes
- **Customer context** — Customer information, interactions, history
- **Team knowledge** — Norms, conventions, institutional knowledge

---

# 4. Comparison to Book Thesis

## Book Thesis Recap

> Reliable AI = engineered context (pragmatics, data, semantics) — relationships given to agent, no inference needed

The book argues that:

1. **Pragmatics** — The intended meaning, goals, and constraints must be explicitly provided
2. **Data** — Raw data and facts must be accessible in structured form
3. **Semantics** — Meaning and relationships must be explicitly represented

The key insight: **relationships should be given to the agent, not inferred by the agent.**

## How Compendium Maps to the Thesis

| Thesis Component | Compendium Implementation |
|-----------------|--------------------------|
| **Engineered context** | Provides a durable, governed knowledge base as the context substrate |
| **Pragmatics** | Stores team decisions, goals, and reasoning — the "why" behind code |
| **Data** | Stores codebase state, customer data, factual information |
| **Semantics** | Provides structured relationships (decisions → outcomes, customers → history) |
| **Relationships given to agent** | Agents read from shared memory instead of inferring context |
| **No inference needed** | Context is explicitly provided; agents don't need to reconstruct it |

## Alignment Analysis

Compendium is a **direct practical implementation** of the book's thesis:

1. **Explicit Context Over Implicit Inference**: Instead of hoping agents can infer team context, Compendium explicitly provides it
2. **Governed Memory**: The knowledge base is not unconstrained memory — it's a structured, team-maintained system
3. **Persistent State**: Context is engineered to persist, not rely on fragile session state
4. **Shared Understanding**: Multiple agents and humans work from the same context (the "single source of truth" for AI)

## Key Insight

The book argues that **context engineering** is the discipline of making AI reliable. Compendium is a product that **implements context engineering** for the specific domain of team-based AI development.

This validates the book's thesis: there is real demand for systems that provide engineered context rather than relying on inference.

---

# 5. Distinctive Positioning

## What Makes Compendium Different

1. **Team memory, not individual memory**: Most memory systems focus on a single user's session. Compendium is explicitly team-based.

2. **Human + Agent sharing**: The knowledge base is designed for both humans and agents to read and write. This is a "collaborative" model, not a "tool" model.

3. **Live, not static**: Unlike documentation or wikis that become stale, Compendium is described as "live" — continuously updated as work happens.

4. **Session sharing**: The ability to share Claude Code sessions directly is a unique capability — not just sharing the memory, but sharing the actual session context.

## Relation to Enterprise Context Layers

Compendium appears to be a specialized version of the "enterprise context layer" concept:

- **Enterprise context layer** (general): The system that turns knowledge, expertise, and norms into machine-usable context for AI agents
- **Compendium** (specific): Focuses on developer teams using Claude Code, with emphasis on code context, decisions, and customer data

This maps to the three layers from the enterprise context layer research:

- **Knowledge** → codebase, decisions, customer data (data context)
- **Expertise** → team norms, conventions, workflows (procedural context)
- **Norms** → governance of what gets stored, who can write (authorization)

---

# 6. Implications for the Book

## Validation of Thesis

Compendium validates the book's core argument:

- There is real market demand for systems that provide engineered context
- The problem of "agents starting from zero" is recognized as significant
- The solution is explicitly about **giving context** rather than hoping inference works

## Patterns to Draw From

1. **Session sharing as a feature**: The ability to share Claude Code sessions directly is an innovative approach to context transfer
2. **Write access for agents**: Agents shouldn't just read context — they should write findings, decisions, and state
3. **Team memory over individual memory**: The unit of context is the team, not the individual
4. **Live updating**: Context systems should be continuously updated, not periodically synchronized

## Gaps in Information

The public page doesn't provide details on:

- How context is stored (vector? graph? hybrid?)
- How retrieval is governed (permissions, freshness, relevance scoring)
- How conflicts are resolved (two agents write conflicting info)
- How context quality is ensured (governance, verification)
- How context promotion works (ephemeral → durable)

These would be valuable areas to explore in a follow-up conversation with Cerenovus or through deeper research.

---

# 7. Conclusion

Cerenovus Compendium is a practical implementation of the book's thesis. It demonstrates that:

1. **The problem is real**: AI agents and new team members starting from zero is a significant pain point
2. **The solution works**: Providing engineered, shared, durable context addresses this problem
3. **The market agrees**: Y Combinator backed this idea, indicating commercial viability

The key insight from both the book and Compendium is the same:

> Don't ask AI to infer what you can explicitly provide.

Compendium makes this operational for team-based AI development.

---

*Research Date: July 11, 2026*
*Context: Context Engineering Book Research*
