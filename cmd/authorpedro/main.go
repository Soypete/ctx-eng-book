package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/metrics"
	"github.com/soypete/authorpedro/internal/tui"
)

//go:embed migrations/*.sql
var migrations embed.FS

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("Prometheus metrics on :9090/metrics")
		if err := http.ListenAndServe(":9090", nil); err != nil {
			log.Printf("Metrics server error: %v", err)
		}
	}()

	if err := runMigrations(cfg.DatabaseURL); err != nil {
		log.Printf("Migration warning: %v", err)
	}

	model, err := tui.NewModel(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing TUI: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
		os.Exit(1)
	}
}

func RecordAgentMetrics(ttft time.Duration, duration time.Duration, iterations int, tokens float64, isError bool) {
	metrics.RecordRequest(ttft.Seconds(), duration.Seconds(), iterations, tokens, isError)
}

func IncInFlight(delta float64) {
	metrics.IncInFlight(delta)
}

func runMigrations(dbURL string) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open DB: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}

	goose.SetBaseFS(migrations)
	goose.SetDialect("postgres")

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Migrations up to date")
	return nil
}
