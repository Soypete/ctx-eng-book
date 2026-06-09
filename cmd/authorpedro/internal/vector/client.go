package vector

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type Client struct {
	db *sql.DB
}

type EmbeddingRecord struct {
	ID          int
	Content     string
	ChapterSlug string
	ModuleSlug  string
	Path        string
	Embedding   []float32
	UpdatedAt   string
}

func NewClient(connStr string) (*Client, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Client{db: db}, nil
}

func (c *Client) Close() error {
	return c.db.Close()
}

func (c *Client) UpsertEmbedding(chapterSlug, moduleSlug, path, content string, embedding []float32) error {
	embeddingStr := vectorToString(embedding)

	query := `
		INSERT INTO authorpedro.book_embeddings (content, chapter_slug, module_slug, path, embedding, updated_at)
		VALUES ($1, $2, $3, $4, $5::vector, NOW())
		ON CONFLICT (chapter_slug, module_slug) DO UPDATE SET
			content = EXCLUDED.content,
			path = EXCLUDED.path,
			embedding = EXCLUDED.embedding,
			updated_at = NOW()
	`

	_, err := c.db.Exec(query, content, chapterSlug, moduleSlug, path, embeddingStr)
	return err
}

func (c *Client) Search(queryEmbedding []float32, topK int) ([]EmbeddingRecord, error) {
	if topK <= 0 {
		topK = 5
	}

	embeddingStr := vectorToString(queryEmbedding)

	query := `
		SELECT id, content, chapter_slug, module_slug, path, embedding, updated_at
		FROM authorpedro.book_embeddings
		ORDER BY embedding <=> $1::vector
		LIMIT $2
	`

	rows, err := c.db.Query(query, embeddingStr, topK)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}
	defer rows.Close()

	var results []EmbeddingRecord
	for rows.Next() {
		var r EmbeddingRecord
		var embStr string
		if err := rows.Scan(&r.ID, &r.Content, &r.ChapterSlug, &r.ModuleSlug, &r.Path, &embStr, &r.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		r.Embedding = stringToVector(embStr)
		results = append(results, r)
	}

	return results, rows.Err()
}

func (c *Client) GetEmbedding(chapterSlug, moduleSlug string) (*EmbeddingRecord, error) {
	query := `
		SELECT id, content, chapter_slug, module_slug, path, embedding, updated_at
		FROM authorpedro.book_embeddings
		WHERE chapter_slug = $1 AND module_slug = $2
	`

	var r EmbeddingRecord
	var embStr string
	err := c.db.QueryRow(query, chapterSlug, moduleSlug).Scan(
		&r.ID, &r.Content, &r.ChapterSlug, &r.ModuleSlug, &r.Path, &embStr, &r.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	r.Embedding = stringToVector(embStr)
	return &r, nil
}

func (c *Client) GetRecentEmbeddings(limit int) ([]EmbeddingRecord, error) {
	if limit <= 0 {
		limit = 5
	}
	query := `
		SELECT id, content, chapter_slug, module_slug, path, embedding, updated_at
		FROM authorpedro.book_embeddings
		ORDER BY updated_at DESC
		LIMIT $1
	`

	rows, err := c.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []EmbeddingRecord
	for rows.Next() {
		var r EmbeddingRecord
		var embStr string
		if err := rows.Scan(&r.ID, &r.Content, &r.ChapterSlug, &r.ModuleSlug, &r.Path, &embStr, &r.UpdatedAt); err != nil {
			return nil, err
		}
		r.Embedding = stringToVector(embStr)
		results = append(results, r)
	}

	return results, rows.Err()
}

func (c *Client) GetAllEmbeddings() ([]EmbeddingRecord, error) {
	query := `
		SELECT id, content, chapter_slug, module_slug, path, embedding, updated_at
		FROM authorpedro.book_embeddings
		ORDER BY chapter_slug, module_slug
	`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []EmbeddingRecord
	for rows.Next() {
		var r EmbeddingRecord
		var embStr string
		if err := rows.Scan(&r.ID, &r.Content, &r.ChapterSlug, &r.ModuleSlug, &r.Path, &embStr, &r.UpdatedAt); err != nil {
			return nil, err
		}
		r.Embedding = stringToVector(embStr)
		results = append(results, r)
	}

	return results, rows.Err()
}

func vectorToString(v []float32) string {
	parts := make([]string, len(v))
	for i, f := range v {
		parts[i] = fmt.Sprintf("%f", f)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func stringToVector(s string) []float32 {
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	if s == "" {
		return nil
	}

	parts := strings.Split(s, ",")
	result := make([]float32, 0, len(parts))
	for _, p := range parts {
		var f float32
		fmt.Sscanf(p, "%f", &f)
		result = append(result, f)
	}
	return result
}
