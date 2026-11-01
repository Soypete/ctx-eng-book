// This example makes boundary ownership part of an executable source contract.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Evidence struct {
	ID, Source, Version, Text string
	ObservedAt                time.Time // When the source vouched for the value, not ingestion time.
}

type Boundary struct {
	Name, Owner, PolicyVersion string
	Check                      func(Evidence, time.Time) error
}

type Decision struct {
	EvidenceID, Source, SourceVersion               string
	ObservedAt, CheckedAt                           time.Time
	Boundary, Owner, PolicyVersion, Outcome, Reason string
}

type Pipeline struct{ boundaries []Boundary }

// NewPipeline rejects incomplete ownership before any request can run.
func NewPipeline(boundaries []Boundary) (Pipeline, error) {
	if len(boundaries) == 0 {
		return Pipeline{}, errors.New("no boundaries configured")
	}
	seen := map[string]bool{}
	for _, b := range boundaries {
		if strings.TrimSpace(b.Name) == "" || strings.TrimSpace(b.Owner) == "" || strings.TrimSpace(b.PolicyVersion) == "" || b.Check == nil {
			return Pipeline{}, errors.New("boundary needs name, owner, policy version, and check")
		}
		if seen[b.Name] {
			return Pipeline{}, fmt.Errorf("duplicate boundary: %s", b.Name)
		}
		seen[b.Name] = true
	}
	return Pipeline{append([]Boundary(nil), boundaries...)}, nil
}

// Run accepts evidence already authorized by the source service.
// model is a stand-in for downstream work; rejection must prevent that call.
func (p Pipeline) Run(e Evidence, now time.Time, model func(string) string) (string, []Decision, error) {
	var decisions []Decision
	if len(p.boundaries) == 0 {
		return "", decisions, errors.New("pipeline not configured")
	}
	for _, b := range p.boundaries {
		d := Decision{EvidenceID: e.ID, Source: e.Source, SourceVersion: e.Version, ObservedAt: e.ObservedAt, CheckedAt: now,
			Boundary: b.Name, Owner: b.Owner, PolicyVersion: b.PolicyVersion, Outcome: "allowed", Reason: "contract satisfied"}
		err := b.Check(e, now)
		if err != nil {
			d.Outcome, d.Reason = "rejected", err.Error()
		}
		decisions = append(decisions, d) // Keep the repair owner beside the decision.
		if err != nil {
			return "", decisions, err // Rejected evidence never reaches the model.
		}
	}
	return model(e.Text), decisions, nil
}

func SourceBoundaries(maxAge time.Duration) []Boundary {
	return []Boundary{
		{Name: "provenance", Owner: "source-publishing", PolicyVersion: "source-v1", Check: func(e Evidence, _ time.Time) error {
			if e.ID == "" || e.Source == "" || e.Version == "" {
				return errors.New("missing source reference")
			}
			return nil
		}},
		{Name: "freshness", Owner: "source-publishing", PolicyVersion: "freshness-v1", Check: func(e Evidence, now time.Time) error {
			if maxAge <= 0 {
				return errors.New("invalid freshness policy")
			}
			if e.ObservedAt.IsZero() || e.ObservedAt.After(now) {
				return errors.New("invalid source time")
			}
			if now.Sub(e.ObservedAt) > maxAge {
				return errors.New("source confirmation too old")
			}
			return nil
		}},
	}
}

func main() {
	now := time.Date(2026, 11, 9, 12, 0, 0, 0, time.UTC)
	p, err := NewPipeline(SourceBoundaries(time.Hour)) // Illustrative policy, not a recommended default.
	if err != nil {
		panic(err)
	}
	evidence := Evidence{"guide-7", "runbook-store", "rev-3", "Restart instructions", now.Add(-2 * time.Hour)}
	called := false
	answer, decisions, runErr := p.Run(evidence, now, func(text string) string { called = true; return text })
	result := struct {
		Answer      string
		ModelCalled bool
		Decisions   []Decision
		Error       string
	}{answer, called, decisions, ""}
	if runErr != nil {
		result.Error = runErr.Error()
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(result); err != nil {
		panic(err)
	}
}
