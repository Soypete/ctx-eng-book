package agent

import (
	"testing"

	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/outline"
)

func TestAgentWithMockServer(t *testing.T) {
	book := outline.Book{
		Chapters: []outline.Chapter{
			{
				Title: "Chapter 1",
				Slug:  "chapter-1",
				Modules: []outline.Module{
					{Title: "Module 1", Slug: "module-1"},
				},
			},
		},
	}

	cfg := config.Config{
		LLMBaseURL: "http://localhost:12345/v1",
		Model:      "test-model",
	}

	ag, err := New(cfg, book)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	var toolCalls []ToolCall
	ag.SetToolCallback(func(tc ToolCall) {
		toolCalls = append(toolCalls, tc)
	})

	_ = toolCalls
}
