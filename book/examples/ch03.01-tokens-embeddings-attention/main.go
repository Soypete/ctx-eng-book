package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Source metadata comes from a trusted source adapter, never document prose.
type Source struct {
	ID, Version, Service, Text string
	Authorized                 bool
	ConfirmedAt                time.Time
}

type Decision struct {
	ID, Version, Reason string
	ConfirmedAt         time.Time
	Included            bool
}

type Evidence struct {
	ID, Version, Text string
}

type Input struct {
	Question string
	Evidence []Evidence
}

// Assemble creates the complete input for this toy interface and its manifest.
// The caller supplies a trusted freshness policy and exact service identifier.
func Assemble(question, service string, sources []Source, now time.Time, maxAge time.Duration) ([]byte, []Decision, error) {
	if question == "" || service == "" || now.IsZero() || maxAge <= 0 {
		return nil, nil, fmt.Errorf("invalid request or policy")
	}
	input := Input{Question: question, Evidence: []Evidence{}}
	decisions := []Decision{}
	for _, s := range sources {
		reason := "included"
		switch {
		case !s.Authorized:
			reason = "unauthorized"
		case s.ID == "" || s.Version == "" || s.Text == "":
			reason = "missing provenance or text"
		case s.Service != service:
			reason = "wrong service"
		case s.ConfirmedAt.IsZero() || s.ConfirmedAt.After(now):
			reason = "invalid confirmation time"
		case now.Sub(s.ConfirmedAt) > maxAge:
			reason = "stale"
		}
		d := Decision{ID: s.ID, Version: s.Version, ConfirmedAt: s.ConfirmedAt, Reason: reason, Included: reason == "included"}
		decisions = append(decisions, d) // Keep the reason even when text is withheld.
		if !d.Included {
			continue // Attention cannot enforce our source-admission policy for us.
		}
		input.Evidence = append(input.Evidence, Evidence{s.ID, s.Version, s.Text})
	}
	if len(input.Evidence) == 0 {
		return nil, decisions, fmt.Errorf("no admissible evidence")
	}
	payload, err := json.Marshal(input)
	return payload, decisions, err
}

func fixture(now time.Time) []Source {
	return []Source{
		{"archive", "v1", "checkout", "Restart every checkout worker.", true, now.Add(-24 * time.Hour)},
		{"catalog", "v3", "catalog", "Restart the catalog queue.", true, now},
		{"restricted", "v2", "checkout", "Private recovery procedure.", false, now},
		{"runbook", "v4", "checkout", "Drain checkout traffic before restarting a worker.", true, now},
	}
}

func main() {
	now := time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)
	payload, decisions, err := Assemble("How do I restart checkout?", "checkout", fixture(now), now, time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Model input: %s\n", payload)
	manifest, err := json.MarshalIndent(decisions, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decision manifest: %s\n", manifest)
}
