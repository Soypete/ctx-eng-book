# LangMem Research Notes

**Source:** langchain-ai/langmem (GitHub repository)  
**Documentation:** https://langchain-ai.github.io/langmem/  
**Stars:** 1.5k  
**License:** MIT

---

## Overview

LangMem is LangChain's library for memory management in AI agents. It provides tooling for agents to "learn and adapt from their interactions over time" through extracting important information from conversations, optimizing agent behavior through prompt refinement, and maintaining long-term memory.

---

## 1. Problem Solved

### The Core Problem

LLM applications suffer from **context loss** — agents forget important details across conversations. Unlike human memory which naturally consolidates and retrieves relevant information, LLM context windows are finite and ephemeral.

LangMem addresses:

1. **Extracting meaningful details** from conversations that warrant remembering
2. **Persistent storage** of learned information across sessions
3. **Optimizing agent behavior** through prompt refinement based on experience
4. **Balancing latency vs. thoroughness** in memory operations

### Why It Matters

Without memory management, every conversation starts from scratch. Users must重复 (repeat) preferences, context, and background. This creates poor user experience and limits agent utility for long-term tasks.

---

## 2. Memory/Context Handling

### Core Memory API

LangMem provides a **functional core** — functions that transform memory state without side effects. These primitives work with any storage system:

- **`create_memory_manager`**: Extract new memories, update/remove outdated memories, consolidate and generalize from existing memories based on new conversation information
- **`create_prompt_optimizer`**: Update prompt rules and core behavior based on conversation information (with optional feedback)

The core pattern for every memory operation:
1. Accept conversation(s) and current memory state
2. Prompt an LLM to determine how to expand or consolidate
3. Respond with updated memory state

### Memory Types

LangMem recognizes distinct memory types (inspired by human memory psychology):

| Type | Purpose | Agent Example | Typical Storage |
|------|---------|---------------|-----------------|
| **Semantic** | Facts & Knowledge | User preferences; knowledge triplets | Profile or Collection |
| **Episodic** | Past Experiences | Few-shot examples; conversation summaries | Collection |
| **Procedural** | System Behavior | Core personality and response patterns | Prompt rules |

#### Semantic Memory: Collections vs Profiles

- **Collections**: Unbounded knowledge to be searched at runtime. Individual documents/records. Requires reconciliation with existing beliefs (insert, update, delete, consolidate).
- **Profiles**: Task-specific, single document representing current state. Updates the existing document rather than creating new ones. Good for user preferences, latest state.

#### Episodic Memory

Captures successful interactions as learning examples — the situation, thought process, and why it worked. Structured as:
- `observation`: The situation and relevant context
- `thoughts`: Key considerations and reasoning process
- `action`: What was done in response
- `result`: What happened and why it worked

#### Procedural Memory

Encodes how an agent should behave. Starts with system prompts, evolves through feedback and experience. Uses `create_prompt_optimizer` to refine instructions.

### Memory Management Tools

Two primary tools give agents direct memory access:

- **`create_manage_memory_tool`**: Store, update, and delete memories during conversations
- **`create_search_memory_tool`**: Semantic search across stored memories

Example:
```python
from langmem import create_manage_memory_tool, create_search_memory_tool

agent = create_react_agent(
    "anthropic:claude-3-5-sonnet-latest",
    tools=[
        create_manage_memory_tool(namespace=("memories", "{user_id}")),
        create_search_memory_tool(namespace=("memories", "{user_id}")),
    ],
    store=store,
)
```

The agent decides *what* and *when* to store — no special commands needed.

### Hot Path vs Background Memory Management

#### Hot Path ("Conscious Formation")

Memory operations happen *during* the conversation — the agent actively invokes memory tools. 

**Pros:**
- Easy to implement
- Agent chooses what to store
- Immediate context updates

**Cons:**
- Adds perceptible latency to user interactions
- Adds complexity to agent tool decisions

#### Background ("Subconscious Formation")

Memory extracted *after* conversations complete (or after inactivity). Prompt an LLM to reflect, find patterns, extract insights without impacting response time.

**Pros:**
- No latency impact on user interactions
- Deeper pattern analysis possible
- Agent tool decisions remain simple

**Cons:**
- Delayed updates
- Requires post-conversation processing

LangMem supports both patterns via:
- Hot path: direct tool invocation
- Background: `create_memory_store_manager` with deferred processing

---

## 3. Knowledge Storage and Retrieval

### LangGraph Integration

LangMem's stateful integration works with LangGraph's storage layer:

- **`create_memory_store_manager`**: Automatically persists extracted memories to LangGraph store
- Works with any `BaseStore` implementation (InMemoryStore, AsyncPostgresStore, etc.)

### Storage System

The storage system is built on LangGraph's `BaseStore` interface:

#### Memory Namespaces

Hierarchical organization of memories:
```python
namespace = ("acme_corp", "{user_id}", "code_assistant")
```

Template variables (like `{user_id}`) are populated at runtime from `configurable` fields in `RunnableConfig`.

#### Retrieval Methods

LangGraph's store supports:
- **Direct Access** (`store.get`): Get specific memory by key
- **Semantic Search** (`store.search`): Find memories by vector similarity
- **Metadata Filtering**: Filter by attributes

### InMemoryStore Example

```python
from langgraph.store.memory import InMemoryStore

store = InMemoryStore(
    index={
        "dims": 1536,
        "embed": "openai:text-embedding-3-small",
    }
)
```

**Note:** For production, use persistent stores like `AsyncPostgresStore`.

---

## 4. Comparison to Book Thesis

### Thesis Reminder

**"reliable AI = engineered context (pragmatics, data, semantics) — relationships given to agent, no inference needed"**

### Alignment Analysis

| Thesis Component | LangMem Approach | Alignment |
|------------------|------------------|-----------|
| **Engineered Context** | Explicit memory extraction, storage, retrieval | Strong alignment — memory *is* engineered context |
| **Pragmatics** | Memory tools that agent uses during conversations | Partial — gives agent capability, but doesn't hardcode relationships |
| **Data** | Semantic search, vector storage | Strong alignment — structured memory storage |
| **Semantics** | Profile schemas, episodic structures | Partial — memory has semantic structure but retrieval requires inference |
| **Relationships Given to Agent** | Namespaces, profiles with schemas | Partial — namespace structure is given, but content requires inference |
| **No Inference Needed** | Direct lookup by key vs semantic search | Mixed — key lookup = no inference, semantic search = similarity inference |

### Key Insights

1. **Memory as Engineered Context**: LangMem treats memory as explicit context to be engineered, not inferred. The agent doesn't guess what to remember — explicit extraction mechanisms decide.

2. **The Inference Trade-off**: 
   - **Key-based lookup** (`store.get`): No inference needed — direct retrieval
   - **Semantic search** (`store.search`): Requires similarity inference
   - This aligns with thesis: "relationships given to agent" = use key-based lookup for guaranteed reliability

3. **Namespace as Given Relationship**: The namespace structure (`("memories", "{user_id}")`) is *engineered* — it's given to the agent. This is exactly what the thesis advocates: explicit relationships, not inferred.

4. **Hot Path Trade-off**: Hot path memory management requires the agent to *decide* when to use memory tools — inference. Background processing removes this but adds latency. The thesis would favor engineered retrieval (background) over inferred usage (hot path).

5. **Missing from LangMem**: The thesis emphasizes *pragmatics* — the relationships between context elements. LangMem stores memories but doesn't enforce structural relationships (like knowledge graph edges) between them.

---

## Key Features Summary

| Feature | Description |
|---------|-------------|
| `create_memory_manager` | Core API for memory extraction/consolidation |
| `create_manage_memory_tool` | Agent tool for storing/updating memories |
| `create_search_memory_tool` | Agent tool for semantic memory retrieval |
| `create_memory_store_manager` | Background memory processing with auto-persist |
| `create_prompt_optimizer` | Procedural memory — refine system prompts |
| InMemoryStore | Development storage (not persistent) |
| Namespaces | Hierarchical memory organization |
| Hot path | Agent actively manages memory during conversation |
| Background | Memory extracted after conversations |

---

## References

- GitHub: https://github.com/langchain-ai/langmem
- Docs: https://langchain-ai.github.io/langmem/
- Core Concepts: https://langchain-ai.github.io/langmem/concepts/conceptual_guide
- Memory Tools Guide: https://langchain-ai.github.io/langmem/guides/memory_tools
- Background Quickstart: https://langchain-ai.github.io/langmem/background_quickstart

---

*Research note created: July 2026*
