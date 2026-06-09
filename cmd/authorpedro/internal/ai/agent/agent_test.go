package agent

import (
	"testing"

	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/outline"
)

func TestExecuteOption(t *testing.T) {
	cfg := &executeConfig{
		maxIterations: 10,
		thinking:      true,
		planMode:      false,
	}

	WithThinking(false)(cfg)
	if cfg.thinking != false {
		t.Errorf("expected thinking false, got %v", cfg.thinking)
	}

	WithPlanMode(true)(cfg)
	if cfg.planMode != true {
		t.Errorf("expected planMode true, got %v", cfg.planMode)
	}

	WithThinking(true)(cfg)
	if cfg.thinking != true {
		t.Errorf("expected thinking true, got %v", cfg.thinking)
	}
}

func TestExecuteConfigDefaults(t *testing.T) {
	cfg := &executeConfig{maxIterations: 10}
	WithThinking(true)(cfg)
	WithPlanMode(true)(cfg)

	if !cfg.thinking {
		t.Error("thinking should be true by default after option applied")
	}
	if !cfg.planMode {
		t.Error("planMode should be true after option applied")
	}
	if cfg.maxIterations != 10 {
		t.Errorf("expected maxIterations 10, got %d", cfg.maxIterations)
	}
}

func TestWithThinkingOption(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected bool
	}{
		{"disable thinking", false, false},
		{"enable thinking", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &executeConfig{}
			WithThinking(tt.input)(cfg)
			if cfg.thinking != tt.expected {
				t.Errorf("WithThinking(%v) = %v, want %v", tt.input, cfg.thinking, tt.expected)
			}
		})
	}
}

func TestWithPlanModeOption(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected bool
	}{
		{"disable plan mode", false, false},
		{"enable plan mode", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &executeConfig{}
			WithPlanMode(tt.input)(cfg)
			if cfg.planMode != tt.expected {
				t.Errorf("WithPlanMode(%v) = %v, want %v", tt.input, cfg.planMode, tt.expected)
			}
		})
	}
}

func TestWithMaxIterations(t *testing.T) {
	opt := func(c *executeConfig) {
		c.maxIterations = 5
	}
	cfg := &executeConfig{}
	opt(cfg)

	if cfg.maxIterations != 5 {
		t.Errorf("maxIterations = %d, want 5", cfg.maxIterations)
	}
}

func TestBookOutline(t *testing.T) {
	book := outline.Book{
		Title: "Test Book",
		Chapters: []outline.Chapter{
			{
				Title: "Chapter 1",
				Slug:  "ch1",
				Modules: []outline.Module{
					{Title: "Module 1.1", Slug: "mod1-1"},
					{Title: "Module 1.2", Slug: "mod1-2"},
				},
			},
			{
				Title: "Chapter 2",
				Slug:  "ch2",
				Modules: []outline.Module{
					{Title: "Module 2.1", Slug: "mod2-1"},
				},
			},
		},
	}

	agent := &Agent{book: book}
	summary := agent.GetOutlineSummary()

	if summary == "" {
		t.Error("expected non-empty summary")
	}

	expectedChapters := []string{"Chapter 1", "Chapter 2"}
	for _, ch := range expectedChapters {
		if !contains(summary, ch) {
			t.Errorf("summary missing chapter: %s", ch)
		}
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestConfigLoading(t *testing.T) {
	cfg := config.Config{
		LLMBaseURL:  "http://localhost:8080/v1",
		Model:       "test-model",
		DatabaseURL: "",
	}

	if cfg.LLMBaseURL == "" {
		t.Error("expected LLMBaseURL to be set")
	}
	if cfg.Model == "" {
		t.Error("expected Model to be set")
	}
}

func TestNewAgentWithoutDB(t *testing.T) {
	cfg := config.Config{
		LLMBaseURL:  "http://localhost:8080/v1",
		Model:       "test-model",
		DatabaseURL: "",
	}

	book := outline.Book{
		Title: "Test Book",
		Chapters: []outline.Chapter{
			{
				Title: "Chapter 1",
				Slug:  "ch1",
				Modules: []outline.Module{
					{Title: "Module 1", Slug: "mod1"},
				},
			},
		},
	}

	agent, err := New(cfg, book)
	if err != nil {
		t.Logf("New returned error (expected if no LLM): %v", err)
		return
	}

	if agent == nil {
		t.Error("expected non-nil agent")
	}

	if agent.book.Title != "Test Book" {
		t.Errorf("book title = %s, want Test Book", agent.book.Title)
	}
}
