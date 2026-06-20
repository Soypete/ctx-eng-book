package agent

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/soypete/authorpedro/internal/ai/booktools"
	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/outline"
	"github.com/soypete/pedro-agentware/go/executor"
	"github.com/soypete/pedro-agentware/go/llm"
	"github.com/soypete/pedro-agentware/go/middleware"
	agenttools "github.com/soypete/pedro-agentware/go/tools"
)

var (
	ErrMaxRetriesExceeded = errors.New("max retries exceeded")
	ErrToolTimeout        = errors.New("tool execution timeout")
)

const (
	maxRetries     = 3
	baseRetryDelay = 1 * time.Second
	maxRetryDelay  = 30 * time.Second
	toolTimeout    = 60 * time.Second
)

type ToolCall struct {
	Name string
	Args map[string]any
}

type StreamCallback func(content string)

type Agent struct {
	executor       executor.Executor
	registry       *agenttools.ToolRegistry
	docTools       *booktools.DocTools
	book           outline.Book
	toolCalls      []ToolCall
	ToolCallback   func(ToolCall)
	StreamCallback StreamCallback
	backend        llm.Backend
}

func New(cfg config.Config, book outline.Book) (*Agent, error) {
	dt := booktools.NewDocTools(cfg, book)

	registry := agenttools.NewToolRegistry()
	registry.Register(dt.ReadModule())
	registry.Register(dt.WriteModule())
	registry.Register(dt.ListOutline())
	registry.Register(dt.SearchContent())
	registry.Register(dt.ResearchMaterial())
	registry.Register(dt.WriteSection())
	registry.Register(dt.AppendCode())
	registry.Register(dt.EditModule())

	backend := llm.NewServerBackend(llm.ServerConfig{
		BaseURL:       cfg.LLMBaseURL,
		Model:         cfg.Model,
		ContextWindow: 128000,
		Timeout:       600 * time.Second,
	})

	policy := middleware.Policy{
		Rules: []middleware.Rule{
			{
				Name:   "allow-all",
				Tools:  []string{"read_module", "write_module", "list_outline", "search_prior_content", "research_material", "write_section", "append_code", "edit_module"},
				Action: middleware.ActionAllow,
			},
		},
		DefaultDeny: false,
	}

	auditor := middleware.NewInMemoryAuditor()

	agent := &Agent{
		executor: nil,
		registry: registry,
		docTools: dt,
		book:     book,
		backend:  backend,
	}

	exec := executor.NewDispatchExecutor(
		backend,
		registry,
		&policy,
		auditor,
		cfg.Model,
		agent.toolCallbackAdapter(),
	)

	agent.executor = exec
	return agent, nil
}

func (a *Agent) toolCallbackAdapter() executor.ToolCallCallback {
	return func(toolName string, args map[string]any) {
		if a.ToolCallback != nil {
			a.ToolCallback(ToolCall{
				Name: toolName,
				Args: args,
			})
		}
	}
}

func (a *Agent) SetStreamCallback(cb StreamCallback) {
	a.StreamCallback = cb
}

func (a *Agent) Execute(ctx context.Context, task string, opts ...ExecuteOption) (*executor.ExecuteResult, error) {
	cfg := &executeConfig{
		maxIterations: 10,
		thinking:      false,
		planMode:      false,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	thinkingStr := ""
	if cfg.thinking {
		thinkingStr = "\n## Thinking\nEnable extended thinking to plan your approach before executing."
	}

	planPrompt := ""
	if cfg.planMode {
		planPrompt = "\n## Plan Mode\nFirst, create a detailed plan for this task. Output your plan, then wait for user confirmation before executing. Start your response with 'PLAN:'"
	}

	execToUse := a.executor
	if a.StreamCallback != nil {
		streamingBackend := WrapBackendForStreaming(a.backend, a.StreamCallback)
		policy := middleware.Policy{
			Rules: []middleware.Rule{
				{
					Name:   "allow-all",
					Tools:  []string{"read_module", "write_module", "list_outline", "search_prior_content", "research_material", "write_section", "append_code", "edit_module"},
					Action: middleware.ActionAllow,
				},
			},
			DefaultDeny: false,
		}
		auditor := middleware.NewInMemoryAuditor()
		modelName := a.backend.ModelName()
		log.Printf("[Agent] Using streaming with model: %s", modelName)
		execToUse = executor.NewDispatchExecutor(
			streamingBackend,
			a.registry,
			&policy,
			auditor,
			modelName,
			a.toolCallbackAdapter(),
		)
	}

	req := executor.ExecuteRequest{
		SystemPrompt: `You are a book-writing assistant helping write "Context Engineering: Building Reliable AI Systems", a technical book with code examples.` + thinkingStr + planPrompt + `

## Your Workflow
1. When asked "what should I work on?" or on launch: check the outline and recent work to suggest next module
2. Research: use research_material to search prior content and list outline
3. Write: use write_section for new content, append_code for code examples
4. Edit: use edit_module to improve existing content
5. On error: retry up to 3 times before giving up

## Available Tools
- list_outline: See full book structure
- research_material(query, type): Search prior writing or get outline
- search_prior_content(query, top_k): Semantic search over written content
- write_section(chapter_slug, module_slug, content): Write new section
- append_code(chapter_slug, module_slug, code, filename, language): Add code example
- edit_module(chapter_slug, module_slug, content, mode): Replace or append content
- read_module(chapter_slug, module_slug): Read existing content

## Tool Calling
When you need to use a tool, you MUST output it in this exact XML format:
<tool_call>
<tool name="TOOL_NAME">
{"arg1": "value1", "arg2": "value2"}
</tool>
</tool_call>

Do NOT output tool descriptions - actually call the tool when needed.

## Error Handling
If a tool fails, retry with exponential backoff. If it fails 3 times, explain the error to the user and suggest alternatives.

## Output
When finished, output "TASK_COMPLETE" to signal completion.`,
		UserMessage:   task,
		MaxIterations: cfg.maxIterations,
		CallerCtx:     middleware.CallerContext{},
	}

	return a.executeWithRetry(ctx, execToUse, req)
}

func (a *Agent) executeWithRetry(ctx context.Context, exec executor.Executor, req executor.ExecuteRequest) (*executor.ExecuteResult, error) {
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			delay := baseRetryDelay * time.Duration(1<<uint(attempt-1))
			if delay > maxRetryDelay {
				delay = maxRetryDelay
			}
			log.Printf("[Agent] Retry %d/%d after %v", attempt, maxRetries-1, delay)
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(delay):
			}
		}

		result, err := exec.Execute(ctx, req)
		if err == nil {
			return result, nil
		}

		lastErr = err

		if isTransientError(err) {
			log.Printf("[Agent] Transient error on attempt %d: %v", attempt+1, err)
			continue
		}

		return nil, err
	}

	return nil, lastErr
}

func isTransientError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	transientPatterns := []string{
		"connection refused",
		"connection reset",
		"timeout",
		"temporary failure",
		"service unavailable",
		"too many requests",
		"rate limit",
		"429",
		"500",
		"502",
		"503",
		"504",
		"i/o timeout",
		"network",
	}
	for _, p := range transientPatterns {
		if containsFold(errStr, p) {
			return true
		}
	}
	return false
}

func containsFold(s, substr string) bool {
	s = toLower(s)
	substr = toLower(substr)
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func toLower(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		result[i] = c
	}
	return string(result)
}

type executeConfig struct {
	maxIterations int
	thinking      bool
	planMode      bool
}

type ExecuteOption func(*executeConfig)

func WithThinking(enabled bool) ExecuteOption {
	return func(c *executeConfig) {
		c.thinking = enabled
	}
}

func WithPlanMode(enabled bool) ExecuteOption {
	return func(c *executeConfig) {
		c.planMode = enabled
	}
}

func (a *Agent) ExecutePlanMode(ctx context.Context, task string) (*executor.ExecuteResult, error) {
	return a.Execute(ctx, task, WithThinking(true), WithPlanMode(true))
}

func (a *Agent) ExecuteWithThinking(ctx context.Context, task string) (*executor.ExecuteResult, error) {
	return a.Execute(ctx, task, WithThinking(true))
}

func (a *Agent) GetRecentModules(limit int) string {
	records := a.docTools.GetRecentModules(limit)
	if len(records) == 0 {
		return "No recent work found. Start fresh with the outline."
	}
	var output string
	output += "## Recent Work\n"
	for _, r := range records {
		output += "- " + r.ChapterSlug + "/" + r.ModuleSlug + " (" + r.UpdatedAt + ")\n"
	}
	return output
}

func (a *Agent) GetOutlineSummary() string {
	var output string
	for _, ch := range a.book.Chapters {
		output += "## " + ch.Title + "\n"
		for _, mod := range ch.Modules {
			output += "- " + ch.Slug + "/" + mod.Slug + ": " + mod.Title + "\n"
		}
	}
	return output
}

func (a *Agent) Registry() *agenttools.ToolRegistry {
	return a.registry
}

func (a *Agent) SetToolCallback(cb func(ToolCall)) {
	a.ToolCallback = cb
}

func (a *Agent) SetOfflineMode(offline bool) {
	a.docTools.SetOffline(offline)
}
