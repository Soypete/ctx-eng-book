package triage

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

var now = time.Date(2026, 10, 19, 9, 0, 0, 0, time.UTC)

func system() System {
	return System{
		Catalog: Catalog{Version: "2026-10-19.1", AsOf: now.Add(-2 * time.Hour), MaxAge: 24 * time.Hour,
			Products: map[string]bool{"payments": true, "search": true, "legacy-reports": false}},
		Severities: map[string]string{
			"sev1": "customer-facing outage or data loss",
			"sev2": "degraded feature with a workaround",
			"sev3": "question or cosmetic issue",
		},
		Routes: map[string]string{"sev1": "oncall", "sev2": "product-queue", "sev3": "support-queue"},
	}
}

// keywordModel stands in for a careful model.
type keywordModel struct{}

func (keywordModel) Classify(_ context.Context, p Prompt) (Label, error) {
	if strings.Contains(p.Ticket, "checkout") {
		return Label{"payments", "sev1"}, nil
	}
	return Label{"search", "sev3"}, nil
}

// confidentModel stands in for a different model that answers fluently with
// labels this system never defined.
type confidentModel struct{ label Label }

func (m confidentModel) Classify(context.Context, Prompt) (Label, error) { return m.label, nil }

func TestVerifiedLabelIsRouted(t *testing.T) {
	r, err := system().Route(context.Background(), keywordModel{}, "checkout returns 500 for every card", now)
	if err != nil || r.Queue != "oncall" || r.CatalogVersion != "2026-10-19.1" {
		t.Fatalf("got %+v, %v", r, err)
	}
}

func TestSwappingModelsKeepsTheSameChecks(t *testing.T) {
	cases := []struct {
		name  string
		model Model
		want  error
	}{
		{"made-up severity", confidentModel{Label{"payments", "P0"}}, ErrUnknownSeverity},
		{"retired product", confidentModel{Label{"legacy-reports", "sev2"}}, ErrUnknownProduct},
		{"invented product", confidentModel{Label{"payments-v2", "sev2"}}, ErrUnknownProduct},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r, err := system().Route(context.Background(), c.model, "checkout is slow", now)
			if !errors.Is(err, c.want) || r.Queue != TriageQueue || r.Reason == "" {
				t.Fatalf("got %+v, %v; want %v routed to triage with a reason", r, err, c.want)
			}
		})
	}
}

func TestStaleCatalogStopsBeforeTheModelRuns(t *testing.T) {
	s := system()
	s.Catalog.AsOf = now.Add(-72 * time.Hour)
	r, err := s.Route(context.Background(), keywordModel{}, "checkout returns 500", now)
	if !errors.Is(err, ErrStaleCatalog) || r.Queue != TriageQueue {
		t.Fatalf("got %+v, %v", r, err)
	}
}

func TestPromptCarriesDefinitionsNotJustLabels(t *testing.T) {
	p := system().prompt("x")
	if p.Severities["sev2"] == "" {
		t.Fatal("model would have to guess what sev2 means")
	}
	for _, name := range p.Products {
		if name == "legacy-reports" {
			t.Fatal("retired product offered to the model")
		}
	}
}
