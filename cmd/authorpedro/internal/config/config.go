package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	// Supabase
	DatabaseURL string `json:"database_url"`

	// Book paths
	BookPath string `json:"book_path"`

	// LLM settings
	LLMBaseURL string `json:"llm_base_url"`
	Model      string `json:"model"`

	// Embeddings
	EmbedModel string `json:"embed_model"`

	// Local state
	StateDir string `json:"state_dir"`
}

func DefaultConfig() Config {
	homeDir, _ := os.UserHomeDir()

	return Config{
		DatabaseURL: "",
		BookPath:    filepath.Join(homeDir, "books", "ctx-eng-book"),
		LLMBaseURL:  "http://100.121.229.114:8080/v1",
		Model:       "qwen3.6-27b",
		EmbedModel:  "qwen3.6-27b",
		StateDir:    filepath.Join(homeDir, ".authorpedro"),
	}
}

func Load() (Config, error) {
	cfg := DefaultConfig()

	envPath := ".env"
	if _, err := os.Stat(envPath); err == nil {
		if err := loadEnvFile(envPath, &cfg); err != nil {
			return cfg, err
		}
	}

	return cfg, nil
}

func loadEnvFile(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	for _, line := range parseEnvLines(string(data)) {
		if line == "" || line[0] == '#' {
			continue
		}
		parts := splitEnvLine(line)
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]
		switch key {
		case "DATABASE_URL":
			cfg.DatabaseURL = value
		case "BOOK_PATH":
			cfg.BookPath = value
		case "LLM_BASE_URL":
			cfg.LLMBaseURL = value
		case "MODEL":
			cfg.Model = value
		case "EMBED_MODEL":
			cfg.EmbedModel = value
		case "STATE_DIR":
			cfg.StateDir = value
		}
	}
	return nil
}

func parseEnvLines(content string) []string {
	var lines []string
	var current string
	inQuote := false

	for _, c := range content {
		if c == '"' || c == '\'' {
			inQuote = !inQuote
		}
		if c == '\n' && !inQuote {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(c)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

func splitEnvLine(line string) [2]string {
	var key, value string
	firstEquals := true

	for _, c := range line {
		if c == '=' && firstEquals {
			firstEquals = false
			continue
		}
		if firstEquals {
			key += string(c)
		} else {
			value += string(c)
		}
	}
	return [2]string{key, value}
}

func (c Config) Save(path string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
