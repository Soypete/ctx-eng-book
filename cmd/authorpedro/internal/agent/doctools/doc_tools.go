package doctools

import (
	"context"
	"os"
	"path/filepath"

	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/outline"
	"github.com/soypete/pedro-agentware/go/tools"
)

type DocTools struct {
	cfg  config.Config
	book outline.Book
}

func NewDocTools(cfg config.Config, book outline.Book) *DocTools {
	return &DocTools{cfg: cfg, book: book}
}

func (dt *DocTools) ReadModule() tools.ExtendedTool  { return &ReadModuleTool{dt: dt} }
func (dt *DocTools) WriteModule() tools.ExtendedTool { return &WriteModuleTool{dt: dt} }
func (dt *DocTools) ListOutline() tools.ExtendedTool { return &ListOutlineTool{book: dt.book} }

type ReadModuleTool struct {
	dt *DocTools
}

func (t *ReadModuleTool) Name() string { return "read_module" }
func (t *ReadModuleTool) Description() string {
	return "Read content from a book module. Input: chapter_slug, module_slug"
}

func (t *ReadModuleTool) Execute(ctx context.Context, args map[string]any) (*tools.Result, error) {
	chapterSlug, _ := args["chapter_slug"].(string)
	moduleSlug, _ := args["module_slug"].(string)

	if chapterSlug == "" || moduleSlug == "" {
		return &tools.Result{Output: "Error: chapter_slug and module_slug required"}, nil
	}

	relPath, err := t.dt.book.ResolvePath(chapterSlug, moduleSlug)
	if err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	fullPath := filepath.Join(t.dt.cfg.BookPath, relPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return &tools.Result{Output: "Module not found or empty"}, nil
	}

	return &tools.Result{Output: string(content)}, nil
}

func (t *ReadModuleTool) InputSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"chapter_slug": map[string]any{"type": "string", "description": "Chapter slug (e.g., 'introduction')"},
			"module_slug":  map[string]any{"type": "string", "description": "Module slug (e.g., 'why-context-engineering')"},
		},
		"required": []string{"chapter_slug", "module_slug"},
	}
}

func (t *ReadModuleTool) Examples() []tools.ToolExample {
	return []tools.ToolExample{
		{
			Input:       map[string]any{"chapter_slug": "introduction", "module_slug": "why-context-engineering"},
			Output:      "Content of the module...",
			Explanation: "Reads the introduction/why-context-engineering module",
		},
	}
}

type WriteModuleTool struct {
	dt *DocTools
}

func (t *WriteModuleTool) Name() string { return "write_module" }
func (t *WriteModuleTool) Description() string {
	return "Write content to a book module. Input: chapter_slug, module_slug, content"
}

func (t *WriteModuleTool) Execute(ctx context.Context, args map[string]any) (*tools.Result, error) {
	chapterSlug, _ := args["chapter_slug"].(string)
	moduleSlug, _ := args["module_slug"].(string)
	content, _ := args["content"].(string)

	if chapterSlug == "" || moduleSlug == "" || content == "" {
		return &tools.Result{Output: "Error: chapter_slug, module_slug, and content required"}, nil
	}

	relPath, err := t.dt.book.ResolvePath(chapterSlug, moduleSlug)
	if err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	fullPath := filepath.Join(t.dt.cfg.BookPath, relPath)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	existing, _ := os.ReadFile(fullPath)
	if len(existing) > 0 {
		existing = append(existing, '\n')
		content = string(existing) + content
	}

	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	return &tools.Result{Output: "Saved to " + relPath}, nil
}

func (t *WriteModuleTool) InputSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"chapter_slug": map[string]any{"type": "string", "description": "Chapter slug"},
			"module_slug":  map[string]any{"type": "string", "description": "Module slug"},
			"content":      map[string]any{"type": "string", "description": "Content to write"},
		},
		"required": []string{"chapter_slug", "module_slug", "content"},
	}
}

func (t *WriteModuleTool) Examples() []tools.ToolExample {
	return nil
}

type ListOutlineTool struct {
	book outline.Book
}

func (t *ListOutlineTool) Name() string { return "list_outline" }
func (t *ListOutlineTool) Description() string {
	return "List all chapters and modules in the book outline"
}

func (t *ListOutlineTool) Execute(ctx context.Context, args map[string]any) (*tools.Result, error) {
	var output string
	for _, ch := range t.book.Chapters {
		output += "## " + ch.Title + " (" + ch.Slug + ")\n"
		for _, mod := range ch.Modules {
			output += "  - " + mod.Title + " (" + mod.Slug + ")\n"
		}
	}
	return &tools.Result{Output: output}, nil
}

func (t *ListOutlineTool) InputSchema() map[string]any {
	return map[string]any{
		"type":       "object",
		"properties": map[string]any{},
	}
}

func (t *ListOutlineTool) Examples() []tools.ToolExample {
	return nil
}
