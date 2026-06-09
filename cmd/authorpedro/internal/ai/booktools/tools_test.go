package booktools

import (
	"testing"

	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/outline"
)

func TestDocToolsNew(t *testing.T) {
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

	dt := NewDocTools(cfg, book)
	if dt == nil {
		t.Error("expected non-nil DocTools")
	}
	if dt.cfg.LLMBaseURL != cfg.LLMBaseURL {
		t.Errorf("expected LLMBaseURL %s, got %s", cfg.LLMBaseURL, dt.cfg.LLMBaseURL)
	}
	if dt.book.Title != book.Title {
		t.Errorf("expected book title %s, got %s", book.Title, dt.book.Title)
	}
}

func TestDocToolsWithoutDB(t *testing.T) {
	cfg := config.Config{
		LLMBaseURL:  "http://localhost:8080/v1",
		Model:       "test-model",
		DatabaseURL: "",
	}

	book := outline.Book{
		Title:    "Test Book",
		Chapters: []outline.Chapter{},
	}

	dt := NewDocTools(cfg, book)
	if dt.vectorCli != nil {
		t.Error("expected nil vectorCli when no DatabaseURL")
	}
	if dt.embedder != nil {
		t.Error("expected nil embedder when no DatabaseURL")
	}
}

func TestDocToolsToolsReturnNonNil(t *testing.T) {
	cfg := config.Config{
		LLMBaseURL:  "http://localhost:8080/v1",
		Model:       "test-model",
		DatabaseURL: "",
	}

	book := outline.Book{
		Title:    "Test Book",
		Chapters: []outline.Chapter{},
	}

	dt := NewDocTools(cfg, book)

	tools := []struct {
		name string
		tool interface{}
	}{
		{"ReadModule", dt.ReadModule()},
		{"WriteModule", dt.WriteModule()},
		{"ListOutline", dt.ListOutline()},
		{"SearchContent", dt.SearchContent()},
		{"ResearchMaterial", dt.ResearchMaterial()},
		{"WriteSection", dt.WriteSection()},
		{"AppendCode", dt.AppendCode()},
		{"EditModule", dt.EditModule()},
	}

	for _, tt := range tools {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tool == nil {
				t.Errorf("%s returned nil", tt.name)
			}
		})
	}
}

func TestListOutlineTool(t *testing.T) {
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

	tool := &ListOutlineTool{book: book}

	if tool.Name() != "list_outline" {
		t.Errorf("Name() = %s, want list_outline", tool.Name())
	}

	if tool.Description() == "" {
		t.Error("Description() should not be empty")
	}
}

func TestGetRecentModulesEmpty(t *testing.T) {
	cfg := config.Config{
		LLMBaseURL:  "http://localhost:8080/v1",
		Model:       "test-model",
		DatabaseURL: "",
	}

	book := outline.Book{
		Title:    "Test Book",
		Chapters: []outline.Chapter{},
	}

	dt := NewDocTools(cfg, book)
	records := dt.GetRecentModules(5)

	if records != nil {
		t.Errorf("expected nil records when no DB, got %v", records)
	}
}
