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

// This deliberately tiny golden set demonstrates evaluation, not model quality.
// Both models return allowed labels; only expected-vs-predicted checks find the miss.
func TestGoldenSetSeparatesValidityFromCorrectness(t *testing.T) {
	golden := []struct {
		ticket string
		want   Label
	}{
		{"checkout returns 500 for every card", Label{"payments", "sev1"}},
		{"where is the search help page?", Label{"search", "sev3"}},
	}
	models := []struct {
		name        string
		model       Model
		wantCorrect int
	}{
		{"keyword", keywordModel{}, 2},
		{"confident-but-wrong", confidentModel{Label{"search", "sev3"}}, 1},
	}
	for _, candidate := range models {
		t.Run(candidate.name, func(t *testing.T) {
			matrix := make(map[Label]map[Label]int)
			correct := 0
			for _, sample := range golden {
				r, err := system().Route(context.Background(), candidate.model, sample.ticket, now)
				if err != nil || r.Queue == TriageQueue {
					t.Fatalf("allowed labels should pass routing: %+v, %v", r, err)
				}
				if matrix[sample.want] == nil {
					matrix[sample.want] = make(map[Label]int)
				}
				matrix[sample.want][r.Label]++
				if r.Label == sample.want {
					correct++
				} else if r.Queue != "support-queue" {
					t.Fatalf("expected valid-but-wrong support route, got %+v", r)
				}
			}
			for _, sample := range golden {
				t.Logf("expected=%+v exact=%d/1 predicted=%v", sample.want, matrix[sample.want][sample.want], matrix[sample.want])
			}
			if correct != candidate.wantCorrect {
				t.Fatalf("exact matches=%d, want %d", correct, candidate.wantCorrect)
			}
		})
	}
}

type failingModel struct{ calls *int }

func (m failingModel) Classify(context.Context, Prompt) (Label, error) {
	*m.calls++
	return Label{}, errors.New("model unavailable")
}

func TestFailuresKeepCatalogVersion(t *testing.T) {
	for _, stale := range []bool{false, true} {
		s := system()
		if stale {
			s.Catalog.AsOf = now.Add(-72 * time.Hour)
		}
		calls := 0
		r, err := s.Route(context.Background(), failingModel{&calls}, "checkout failed", now)
		if err == nil || r.Queue != TriageQueue || r.CatalogVersion != s.Catalog.Version || r.Reason == "" {
			t.Fatalf("failure lost provenance: %+v, %v", r, err)
		}
		if stale && calls != 0 {
			t.Fatal("stale catalog reached the model")
		}
		if !stale && calls != 1 {
			t.Fatal("fresh catalog did not reach the model")
		}
	}
}
