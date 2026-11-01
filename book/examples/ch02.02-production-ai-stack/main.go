// This lookup demo uses invented data and deterministic model substitutes.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Request struct{ Product, Release string }

// Documents arrive from an already-authorized, trusted source adapter.
// AsOf means the source vouched for this version then, not ingestion time.
type Document struct {
	ID, Product, Release, Source, Version, Text string
	AsOf                                        time.Time
}

type Policy struct {
	Version string
	MaxAge  time.Duration
}

type Decision struct {
	DocumentID, Source, SourceVersion     string
	AsOf                                  time.Time
	Stage, Outcome, Reason, PolicyVersion string
}

type Result struct {
	Text      string
	Decisions []Decision
}

// A provider adapter may serialize these values, but cannot change admission rules.
type Model interface {
	Select(context.Context, Request, []Document) (string, error)
}

func rejectReason(r Request, d Document, now time.Time, p Policy) string {
	switch {
	case d.ID == "" || d.Source == "" || d.Version == "" || d.Text == "":
		return "missing source fields"
	case d.Product != r.Product || d.Release != r.Release:
		return "product or release mismatch"
	case d.AsOf.IsZero() || d.AsOf.After(now) || now.Sub(d.AsOf) > p.MaxAge:
		return "source time outside policy"
	default:
		return ""
	}
}

func Lookup(ctx context.Context, r Request, docs []Document, now time.Time, p Policy, m Model) (Result, error) {
	out := Result{}
	if r.Product == "" || r.Release == "" || p.Version == "" || p.MaxAge <= 0 {
		return out, errors.New("invalid request or policy")
	}
	eligible := []Document{}
	seen := map[string]bool{}
	for _, d := range docs {
		reason := rejectReason(r, d, now, p)
		outcome := "rejected"
		if reason == "" {
			if seen[d.ID] {
				return out, errors.New("ambiguous document ID")
			}
			seen[d.ID] = true
			eligible = append(eligible, d)
			outcome, reason = "admitted", "source contract passed"
		}
		out.Decisions = append(out.Decisions, Decision{d.ID, d.Source, d.Version, d.AsOf, "assembly", outcome, reason, p.Version})
	}
	if len(eligible) == 0 {
		return out, errors.New("no eligible evidence; ask documentation owner")
	}
	// Keep a private copy so even an adapter cannot mutate validation evidence.
	id, err := m.Select(ctx, r, append([]Document(nil), eligible...))
	if err != nil {
		out.Decisions = append(out.Decisions, Decision{Stage: "model", Outcome: "rejected", Reason: "model error", PolicyVersion: p.Version})
		return out, err
	}
	// A proposed citation cannot restore a document rejected during assembly.
	for _, d := range eligible {
		if d.ID == id {
			out.Text = d.Text // render the approved passage, not invented deployment steps
			out.Decisions = append(out.Decisions, Decision{d.ID, d.Source, d.Version, d.AsOf, "validation", "returned", "citation admitted", p.Version})
			return out, nil
		}
	}
	out.Decisions = append(out.Decisions, Decision{Stage: "validation", Outcome: "rejected", Reason: "citation outside admitted evidence", PolicyVersion: p.Version})
	return out, errors.New("citation outside admitted evidence")
}

type firstModel struct{}

func (firstModel) Select(_ context.Context, _ Request, docs []Document) (string, error) {
	return docs[0].ID, nil
}

type lastModel struct{}

func (lastModel) Select(_ context.Context, _ Request, docs []Document) (string, error) {
	return docs[len(docs)-1].ID, nil
}

func fixture() (Request, []Document, time.Time, Policy) {
	now := time.Date(2026, 11, 6, 12, 0, 0, 0, time.UTC)
	return Request{"checkout", "v2"}, []Document{
		{"guide-v1", "checkout", "v1", "release-docs", "rev-a", "Restart the worker directly.", now},
		{"guide-v2", "checkout", "v2", "release-docs", "rev-b", "Drain the worker before restarting it.", now},
	}, now, Policy{"deployment-lookup-v1", time.Hour}
}

func main() {
	r, docs, now, p := fixture()
	fmt.Println("Unchecked first result:", docs[0].Text)
	for _, m := range []Model{firstModel{}, lastModel{}} {
		out, err := Lookup(context.Background(), r, docs, now, p, m)
		if err != nil {
			panic(err)
		}
		data, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Checked %T:\n%s\n", m, data)
	}
}
