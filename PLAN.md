# AuthorPedro — Implementation Plan

## Project Overview

**AuthorPedro** is a Go TUI agent harness for writing a book incrementally each day. It uses pedro-agentware for model-agnostic tool calling, Supabase pgvector for semantic search over prior content, and integrates with the OpenCode skill (adapted to Go) to reference prior work.

**Book being written**: "Context Engineering: Building Reliable AI Systems" (this repo)

---

## Architecture

### Stack

- **TUI**: Bubble Tea (`charmbracelet/bubbletea`) + Lipgloss — matches professor_pedro
- **Agent Loop**: pedro-agentware `Executor` → Nemotron via `pedrogpt:8000`
- **State**: SQLite (modernc.org/sqlite) — session position, outline cache
- **Vector Search**: Supabase pgvector — connection via `.env`
- **Tool Calling**: GenericFormatter (OpenAI-style JSON) — works with Nemotron
- **Embeddings**: Same model via `/v1/embeddings` endpoint

### Data Flow

```
User Input → TUI → Agent Loop (pedro-agentware)
           → Parse Tool Calls → Execute Tools
           → Vector Search (Supabase)
           → OpenCode Skill Tool
           → Stream Response → TUI
```

---

## Directory Structure

This repo contains both the book content and the AuthorPedro codebase.

```
ctx-eng-book/
├── .env                          # Supabase connection string (gitignored)
├── .gitignore                    # Ignore .env, binary, SQLite
├── PLAN.md                       # This file
├── README.md
├── chapters.md                   # Book outline
├── research/                     # Research materials
│
├── cmd/authorpedro/              # Main CLI/TUI entrypoint
│   ├── main.go
│   ├── go.mod
│   └── go.sum
│
├── internal/
│   ├── config/
│   │   └── config.go             # Load .env, set defaults
│   ├── outline/
│   │   ├── outline.go            # Parse chapters.md → tree
│   │   └── resolver.go           # Resolve outline nodes → file paths
│   ├── tui/
│   │   ├── app.go                # Bubble Tea model
│   │   ├── views/
│   │   │   ├── writing.go        # Writing pane
│   │   │   ├── agent.go          # Agent output pane
│   │   │   └── status.go         # Current chapter/module
│   │   └── styles/
│   │       └── theme.go
│   ├── agent/
│   │   ├── executor.go           # Wire pedro-agentware Executor
│   │   └── tools/
│   │       ├── doc_tool.go       # read_module, write_module
│   │       ├── outline_tool.go   # list_outline
│   │       ├── search_tool.go    # search_prior_content (vector)
│   │       └── opencode_tool.go  # Reference prior work
│   ├── vector/
│   │   ├── client.go             # Supabase pgvector client
│   │   └── embedder.go           # Embedding generation (via pedrogpt)
│   └── storage/
│       └── db.go                 # SQLite for session state
│
└── books/
    └── ctx-eng-book/             # Actual book content (scaffolded)
        ├── introduction/
        │   └── modules/
        │       └── 01-intro.md
        └── chapters/
            ├── 01-reliability/
            │   └── modules/
            │       └── 01-why-ai-fails.md
            └── ...
```

---

## Outline Format

Use **YAML** for the book outline (diffable, deterministic, easy to generate paths).

Location: `book.yaml` in repo root (generated from `chapters.md` on first run).

```yaml
book:
  title: "Context Engineering: Building Reliable AI Systems"
  slug: "ctx-eng-book"

chapters:
  - number: 1
    title: "The Reliability Problem"
    slug: "reliability"
    modules:
      - number: 1
        title: "Why AI Feels Magical in Demos"
        slug: "why-ai-feels-magical"

  - number: 2
    title: "What Is Context?"
    slug: "what-is-context"
    modules:
      - number: 1
        title: "Beyond the Prompt"
        slug: "beyond-prompt"
      # ...

  # ... (all chapters from chapters.md)
```

**Path resolution**:
- Module → `books/{book-slug}/chapters/{ch:02d}-{slug}/modules/{mn:02d}-{slug}.md`
- Introduction → `books/{book-slug}/introduction/modules/{mn:02d}-{slug}.md`

---

## Tool Definitions (v1)

### 1. `read_module`
- **Input**: `{ "chapter": "01-reliability", "module": "01-why-ai-feels-magical" }`
- **Output**: Markdown content of the module file
- **Error**: File not found, invalid path

### 2. `write_module`
- **Input**: `{ "chapter": "01-reliability", "module": "01-why-ai-feels-magical", "content": "# New content..." }`
- **Output**: `"Written to books/ctx-eng-book/chapters/01-reliability/modules/01-why-ai-feels-magical.md"`
- **Side effects**: Re-embed content to vector store

### 3. `list_outline`
- **Input**: `{}` (no params)
- **Output**: JSON array of all chapters/modules with paths

### 4. `search_prior_content`
- **Input**: `{ "query": "semantic search query", "limit": 5 }`
- **Output**: Array of `{ "content": "...", "path": "...", "score": 0.95 }`
- **Backend**: Supabase pgvector

### 5. `opencode_skill`
- **Input**: `{ "query": "what did I write about retrieval last week?" }`
- **Output**: Relevant content from OpenCode skill memory
- **Implementation**: Go adaptation of experiments repo's MemPalace adapter

---

## Implementation Phases

### Phase 1: Foundation (Week 1)

**Goal**: Minimal TUI shell that can read/write modules manually

1. Initialize `cmd/authorpedro/` with Go module
2. Create `.env` template + `.gitignore`
3. Implement `config.go` — load Supabase from `.env`
4. Implement `outline.go` — parse `chapters.md` to outline tree
5. Implement `resolver.go` — map outline nodes to file paths; scaffold missing dirs
6. Build Bubble Tea TUI:
   - Writing pane (multiline input)
   - Status line (current chapter/module)
   - Agent pane (output only)
7. Wire read/write — user types, saves to file

**Deliverable**: TUI where you can navigate chapters and write content manually

---

### Phase 2: Agent Integration (Week 2)

**Goal**: Agent assists with writing via tool calls

1. Wire pedro-agentware `Executor`:
   - Backend: HTTP client to `http://pedrogpt:8000/v1/chat/completions`
   - Model: `nemotron-3-super-120b`
   - Formatter: `GenericFormatter` (OpenAI-style)
2. Create `ToolRegistry` with v1 tools
3. Implement `doc_tool.go`: read_module, write_module
4. Implement `outline_tool.go`: list_outline
5. Add agent pane streaming (receive tokens, append to view)
6. System prompt: "You are a book-writing assistant. Help the user write 'Context Engineering'..."

**Deliverable**: User types, agent can read/write files

---

### Phase 3: Vector Search (Week 3)

**Goal**: Semantic search over written content

1. Implement `vector/client.go`:
   - Connect to Supabase via `lib/pq`
   - Initialize `book_embeddings` table if not exists
   - Schema: `id, content, chapter, module, path, embedding vector(1536), updated_at`
2. Implement `vector/embedder.go`:
   - Use Nemotron via `http://pedrogpt:8000/v1/embeddings`
3. On `write_module`: embed new content, upsert to Supabase
4. Implement `search_tool.go`: `search_prior_content` — query Supabase, return top-k results
5. Add to tool registry

**Deliverable**: Agent can search "what did I say about retrieval?" and get relevant passages

---

### Phase 4: Resume + OpenCode Tool (Week 4)

**Goal**: "Where did we leave off" on launch + OpenCode skill integration

1. Resume pass on launch:
   - Query Supabase for most recently updated modules
   - Sort by `updated_at` DESC, take top 5
   - Agent summarizes: "Last worked on Chapter X, Module Y. Unfinished: ..."
2. Implement `opencode_tool.go`:
   - Adapt experiments repo's MemPalace adapter to Go
   - Connect to OpenCode skill data (stored in `~/.opencode/memory/` or similar)
   - Tool: "Given query, search OpenCode memory for relevant prior work"
3. Add to tool registry

**Deliverable**: Launch → agent reports where we left off; agent can query OpenCode skill

---

### Phase 5: Polish (Week 5)

1. Prometheus metrics for agent loop (iterations, tool calls, latency)
2. Error handling + retry logic
3. Offline mode (disable vector search, work from cached embeddings)
4. Session persistence — resume exact cursor position in module
5. Commands: `/search <query>`, `/outline`, `/next`, `/prev`

---

## Configuration (.env)

```bash
# Supabase
DATABASE_URL=postgres://user:pass@host:5432/db?sslmode=disable

# AuthorPedro
BOOK_PATH=./books/ctx-eng-book
MODEL=nemotron-3-super-120b
LLM_BASE_URL=http://pedrogpt:8000/v1
```

---

## Key Design Decisions

| Decision | Rationale |
|----------|-----------|
| YAML outline | Diffable, deterministic path generation, matches existing patterns |
| Supabase pgvector | Already running, Go driver (`lib/pq`) available |
| GenericFormatter | Works with Nemotron's OpenAI-style tool calls |
| Nemotron for embeddings | Same model for writing and embedding (simpler) |
| Go adaptation of OpenCode tool | Good exercise, keeps deps minimal |
| SQLite for state | Matches professor_pedro, sufficient for session metadata |
| Bubble Tea | Matches professor_pedro TUI |

---

## Upstream Contributions

As we build, if pedro-agentware needs changes:

1. **Nemotron formatter**: If GenericFormatter doesn't work, add explicit case to `toolformat/selector.go`
2. **Auditor hooks**: Already present in middleware — use for observability

Contribute back via PRs to respective repos.

---

## Fallback Plan

If Nemotron embeddings don't work:
- Use `qwen3.6-27b` for embeddings (loaded on pedrogpt)
- Update `EMBED_MODEL` in config

---

## Next Steps (after plan review)

1. Create `cmd/authorpedro/` skeleton
2. Add `.env` template + update `.gitignore`
3. Implement Phase 1 — Foundation

---

## Blocked by SSH Access

The following items require SSH access to resolve:

1. **Push upstream changes** — Any PRs to pedro-agentware, professor_pedro, or other repos need SSH to push
2. **pedrogpt connectivity** — Need to verify pedrogpt is reachable from local machine (may need VPN/Tailscale)
3. **Supabase connectivity** — Need to verify Supabase instance is accessible

### To resolve next session:

```bash
# Test connectivity to pedrogpt (when on same network/VPN)
curl http://pedrogpt:8000/v1/models

# Test Supabase connection
psql "$DATABASE_URL" -c "SELECT 1"

# Push upstream changes (once SSH is working)
git push origin main
```

### Currently Blocked
- **Git commits**: 1Password integration blocking commit (not SSH)
- **SSH access**: Can't push to pedro repos
- **pedrogpt**: Can't test connectivity until on VPN/network