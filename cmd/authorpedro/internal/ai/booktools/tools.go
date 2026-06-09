package booktools

import (
	"context"
	"os"
	"path/filepath"

	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/outline"
	"github.com/soypete/authorpedro/internal/vector"
	"github.com/soypete/pedro-agentware/go/tools"
)

type DocTools struct {
	cfg       config.Config
	book      outline.Book
	vectorCli *vector.Client
	embedder  *vector.Embedder
}

func NewDocTools(cfg config.Config, book outline.Book) *DocTools {
	dt := &DocTools{cfg: cfg, book: book}

	if cfg.DatabaseURL != "" {
		cli, err := vector.NewClient(cfg.DatabaseURL)
		if err == nil {
			dt.vectorCli = cli
			dt.embedder = vector.NewEmbedder(cfg.LLMBaseURL, cfg.EmbedModel)
		}
	}

	return dt
}

func (dt *DocTools) ReadModule() tools.ExtendedTool       { return &ReadModuleTool{dt: dt} }
func (dt *DocTools) WriteModule() tools.ExtendedTool      { return &WriteModuleTool{dt: dt} }
func (dt *DocTools) ListOutline() tools.ExtendedTool      { return &ListOutlineTool{book: dt.book} }
func (dt *DocTools) SearchContent() tools.ExtendedTool    { return &SearchContentTool{dt: dt} }
func (dt *DocTools) ResearchMaterial() tools.ExtendedTool { return &ResearchTool{dt: dt} }
func (dt *DocTools) WriteSection() tools.ExtendedTool     { return &WriteSectionTool{dt: dt} }
func (dt *DocTools) AppendCode() tools.ExtendedTool       { return &AppendCodeTool{dt: dt} }
func (dt *DocTools) EditModule() tools.ExtendedTool       { return &EditModuleTool{dt: dt} }

func (dt *DocTools) GetRecentModules(limit int) []vector.EmbeddingRecord {
	if dt.vectorCli == nil {
		return nil
	}
	records, _ := dt.vectorCli.GetRecentEmbeddings(limit)
	return records
}

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

	if t.dt.vectorCli != nil && t.dt.embedder != nil {
		emb, err := t.dt.embedder.Embed(content)
		if err == nil {
			t.dt.vectorCli.UpsertEmbedding(chapterSlug, moduleSlug, relPath, content, emb)
		}
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

type SearchContentTool struct {
	dt *DocTools
}

func (t *SearchContentTool) Name() string { return "search_prior_content" }
func (t *SearchContentTool) Description() string {
	return "Search previously written content using semantic similarity. Input: query, top_k (optional, default 5)"
}

func (t *SearchContentTool) Execute(ctx context.Context, args map[string]any) (*tools.Result, error) {
	if t.dt.vectorCli == nil || t.dt.embedder == nil {
		return &tools.Result{Output: "Error: vector search not configured"}, nil
	}

	query, _ := args["query"].(string)
	if query == "" {
		return &tools.Result{Output: "Error: query required"}, nil
	}

	topK := 5
	if tk, ok := args["top_k"].(float64); ok {
		topK = int(tk)
	}

	emb, err := t.dt.embedder.Embed(query)
	if err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	results, err := t.dt.vectorCli.Search(emb, topK)
	if err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	if len(results) == 0 {
		return &tools.Result{Output: "No similar content found"}, nil
	}

	var output string
	for i, r := range results {
		output += "--- Result " + itoa(i+1) + " ---\n"
		output += "Chapter: " + r.ChapterSlug + " | Module: " + r.ModuleSlug + "\n"
		if len(r.Content) > 500 {
			output += r.Content[:500] + "...\n\n"
		} else {
			output += r.Content + "\n\n"
		}
	}
	return &tools.Result{Output: output}, nil
}

func (t *SearchContentTool) InputSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"query": map[string]any{"type": "string", "description": "Search query text"},
			"top_k": map[string]any{"type": "number", "description": "Number of results (default 5)"},
		},
		"required": []string{"query"},
	}
}

func (t *SearchContentTool) Examples() []tools.ToolExample {
	return nil
}

type ResearchTool struct {
	dt *DocTools
}

func (t *ResearchTool) Name() string { return "research_material" }
func (t *ResearchTool) Description() string {
	return "Research content: search prior writing, read modules, list outline. Use this to gather context before writing."
}

func (t *ResearchTool) Execute(ctx context.Context, args map[string]any) (*tools.Result, error) {
	query, _ := args["query"].(string)
	searchType, _ := args["type"].(string)

	var output string

	if searchType == "outline" || query == "" {
		output += "## Book Outline\n"
		for _, ch := range t.dt.book.Chapters {
			output += "### " + ch.Title + " (" + ch.Slug + ")\n"
			for _, mod := range ch.Modules {
				output += "- " + mod.Title + " (" + mod.Slug + ")\n"
			}
		}
	}

	if query != "" && t.dt.vectorCli != nil && t.dt.embedder != nil {
		emb, err := t.dt.embedder.Embed(query)
		if err == nil {
			results, err := t.dt.vectorCli.Search(emb, 5)
			if err == nil && len(results) > 0 {
				output += "\n## Relevant Prior Content\n"
				for _, r := range results {
					output += "**" + r.ChapterSlug + "/" + r.ModuleSlug + "**\n"
					preview := r.Content
					if len(preview) > 300 {
						preview = preview[:300] + "..."
					}
					output += preview + "\n\n"
				}
			}
		}
	}

	if output == "" {
		output = "Use: research_material(query='topic', type='outline'|'search')"
	}

	return &tools.Result{Output: output}, nil
}

func (t *ResearchTool) InputSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"query": map[string]any{"type": "string", "description": "Search query for prior content"},
			"type":  map[string]any{"type": "string", "description": "Type: 'outline' for structure, 'search' for semantic search"},
		},
	}
}

func (t *ResearchTool) Examples() []tools.ToolExample {
	return nil
}

type WriteSectionTool struct {
	dt *DocTools
}

func (t *WriteSectionTool) Name() string { return "write_section" }
func (t *WriteSectionTool) Description() string {
	return "Write a new section/module. Provide chapter_slug, module_slug, content. Auto-embeds for search."
}

func (t *WriteSectionTool) Execute(ctx context.Context, args map[string]any) (*tools.Result, error) {
	chapterSlug, _ := args["chapter_slug"].(string)
	moduleSlug, _ := args["module_slug"].(string)
	content, _ := args["content"].(string)

	if chapterSlug == "" || moduleSlug == "" || content == "" {
		return &tools.Result{Output: "Error: chapter_slug, module_slug, content required"}, nil
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

	if t.dt.vectorCli != nil && t.dt.embedder != nil {
		emb, err := t.dt.embedder.Embed(content)
		if err == nil {
			t.dt.vectorCli.UpsertEmbedding(chapterSlug, moduleSlug, relPath, content, emb)
		}
	}

	return &tools.Result{Output: "Written to " + relPath}, nil
}

func (t *WriteSectionTool) InputSchema() map[string]any {
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

func (t *WriteSectionTool) Examples() []tools.ToolExample {
	return nil
}

type AppendCodeTool struct {
	dt *DocTools
}

func (t *AppendCodeTool) Name() string { return "append_code" }
func (t *AppendCodeTool) Description() string {
	return "Append code example to a module. Adds to code/ directory under the module path."
}

func (t *AppendCodeTool) Execute(ctx context.Context, args map[string]any) (*tools.Result, error) {
	chapterSlug, _ := args["chapter_slug"].(string)
	moduleSlug, _ := args["module_slug"].(string)
	filename, _ := args["filename"].(string)
	code, _ := args["code"].(string)
	language, _ := args["language"].(string)

	if chapterSlug == "" || moduleSlug == "" || code == "" {
		return &tools.Result{Output: "Error: chapter_slug, module_slug, code required"}, nil
	}

	if filename == "" {
		filename = "example.go"
	}

	relPath, err := t.dt.book.ResolvePath(chapterSlug, moduleSlug)
	if err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	codeDir := filepath.Join(t.dt.cfg.BookPath, "code", relPath)
	if err := os.MkdirAll(codeDir, 0755); err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	codePath := filepath.Join(codeDir, filename)
	ext := ""
	switch language {
	case "go":
		ext = ".go"
	case "python":
		ext = ".py"
	case "javascript", "js":
		ext = ".js"
	case "typescript", "ts":
		ext = ".ts"
	case "shell", "bash":
		ext = ".sh"
	}
	if ext != "" && filepath.Ext(filename) == "" {
		codePath += ext
	}

	existing, _ := os.ReadFile(codePath)
	if len(existing) > 0 {
		existing = append(existing, '\n')
		code = string(existing) + code
	}

	if err := os.WriteFile(codePath, []byte(code), 0644); err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	relCodePath := filepath.Join("code", relPath, filepath.Base(codePath))
	return &tools.Result{Output: "Code written to " + relCodePath}, nil
}

func (t *AppendCodeTool) InputSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"chapter_slug": map[string]any{"type": "string", "description": "Chapter slug"},
			"module_slug":  map[string]any{"type": "string", "description": "Module slug"},
			"filename":     map[string]any{"type": "string", "description": "Code filename (default: example.go)"},
			"code":         map[string]any{"type": "string", "description": "Code content"},
			"language":     map[string]any{"type": "string", "description": "Language: go, python, javascript, etc."},
		},
		"required": []string{"chapter_slug", "module_slug", "code"},
	}
}

func (t *AppendCodeTool) Examples() []tools.ToolExample {
	return nil
}

type EditModuleTool struct {
	dt *DocTools
}

func (t *EditModuleTool) Name() string { return "edit_module" }
func (t *EditModuleTool) Description() string {
	return "Edit/replace content in an existing module. Use read_module first to see current content."
}

func (t *EditModuleTool) Execute(ctx context.Context, args map[string]any) (*tools.Result, error) {
	chapterSlug, _ := args["chapter_slug"].(string)
	moduleSlug, _ := args["module_slug"].(string)
	content, _ := args["content"].(string)
	mode, _ := args["mode"].(string)

	if chapterSlug == "" || moduleSlug == "" || content == "" {
		return &tools.Result{Output: "Error: chapter_slug, module_slug, content required"}, nil
	}

	relPath, err := t.dt.book.ResolvePath(chapterSlug, moduleSlug)
	if err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	fullPath := filepath.Join(t.dt.cfg.BookPath, relPath)

	if mode == "append" {
		existing, _ := os.ReadFile(fullPath)
		if len(existing) > 0 {
			existing = append(existing, '\n')
			content = string(existing) + content
		}
	}

	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return &tools.Result{Output: "Error: " + err.Error()}, nil
	}

	if t.dt.vectorCli != nil && t.dt.embedder != nil {
		emb, err := t.dt.embedder.Embed(content)
		if err == nil {
			t.dt.vectorCli.UpsertEmbedding(chapterSlug, moduleSlug, relPath, content, emb)
		}
	}

	return &tools.Result{Output: "Updated " + relPath}, nil
}

func (t *EditModuleTool) InputSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"chapter_slug": map[string]any{"type": "string", "description": "Chapter slug"},
			"module_slug":  map[string]any{"type": "string", "description": "Module slug"},
			"content":      map[string]any{"type": "string", "description": "New content (or appended if mode=append)"},
			"mode":         map[string]any{"type": "string", "description": "Mode: 'replace' or 'append' (default replace)"},
		},
		"required": []string{"chapter_slug", "module_slug", "content"},
	}
}

func (t *EditModuleTool) Examples() []tools.ToolExample {
	return nil
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	var s string
	for i > 0 {
		s = string(rune('0'+i%10)) + s
		i /= 10
	}
	return s
}
