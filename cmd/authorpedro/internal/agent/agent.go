package agent

import (
	"context"
	"time"

	"github.com/soypete/authorpedro/internal/agent/doctools"
	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/outline"
	"github.com/soypete/pedro-agentware/go/executor"
	"github.com/soypete/pedro-agentware/go/llm"
	"github.com/soypete/pedro-agentware/go/middleware"
	"github.com/soypete/pedro-agentware/go/tools"
)

type Agent struct {
	executor executor.Executor
	registry *tools.ToolRegistry
}

func New(cfg config.Config, book outline.Book) (*Agent, error) {
	dt := doctools.NewDocTools(cfg, book)

	registry := tools.NewToolRegistry()
	registry.Register(dt.ReadModule())
	registry.Register(dt.WriteModule())
	registry.Register(dt.ListOutline())

	backend := llm.NewServerBackend(llm.ServerConfig{
		BaseURL:       cfg.LLMBaseURL,
		Model:         cfg.Model,
		ContextWindow: 128000,
		Timeout:       120 * time.Second,
	})

	policy := middleware.Policy{
		Rules: []middleware.Rule{
			{
				Name:   "allow-all",
				Tools:  []string{"read_module", "write_module", "list_outline"},
				Action: middleware.ActionAllow,
			},
		},
		DefaultDeny: false,
	}

	auditor := middleware.NewInMemoryAuditor()

	exec := executor.NewDispatchExecutor(
		backend,
		registry,
		&policy,
		auditor,
		cfg.Model,
	)

	return &Agent{
		executor: exec,
		registry: registry,
	}, nil
}

func (a *Agent) Execute(ctx context.Context, task string) (*executor.ExecuteResult, error) {
	req := executor.ExecuteRequest{
		SystemPrompt: `You are a book-writing assistant helping write "Context Engineering: Building Reliable AI Systems".
You have access to tools to read and write book modules.
Use write_module to save content, read_module to reference existing content.
When finished, output "TASK_COMPLETE" to signal completion.`,
		UserMessage:   task,
		MaxIterations: 10,
		CallerCtx:     middleware.CallerContext{},
	}

	return a.executor.Execute(ctx, req)
}

func (a *Agent) Registry() *tools.ToolRegistry {
	return a.registry
}
