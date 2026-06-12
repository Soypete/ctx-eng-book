-- +goose Up
CREATE EXTENSION IF NOT EXISTS vector;

CREATE SCHEMA IF NOT EXISTS authorpedro;

CREATE TABLE IF NOT EXISTS authorpedro.book_embeddings (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    chapter_slug VARCHAR(255) NOT NULL,
    module_slug VARCHAR(255) NOT NULL,
    path VARCHAR(512) NOT NULL,
    embedding vector(4096),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(chapter_slug, module_slug)
);

CREATE INDEX IF NOT EXISTS idx_book_embeddings_chapter ON authorpedro.book_embeddings(chapter_slug);
CREATE INDEX IF NOT EXISTS idx_book_embeddings_module ON authorpedro.book_embeddings(module_slug);
CREATE INDEX IF NOT EXISTS idx_book_embeddings_updated ON authorpedro.book_embeddings(updated_at DESC);

ALTER TABLE authorpedro.book_embeddings OWNER TO postgres;
ALTER SCHEMA authorpedro OWNER TO postgres;

-- +goose Down
DROP TABLE IF EXISTS authorpedro.book_embeddings;
DROP SCHEMA IF NOT EXISTS authorpedro;