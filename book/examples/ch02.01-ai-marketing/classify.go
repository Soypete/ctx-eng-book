// Package triage is the runnable companion to the post
// "Stop Saying \"The AI Failed\"" (book module 2.1).
//
// A support-ticket classifier assigns a product and a severity. The model is
// a parameter: anything that satisfies Model. Everything that decides whether
// the result can be trusted lives around it, in the System:
//
//   - Lexicon:    a governed product catalog with a version and an as-of time
//   - Semantics:  a severity taxonomy whose labels have written definitions
//   - Pragmatics: routing rules that accept only labels a queue exists for
//
// Swap the model and those checks stay the same. That is the post's point.
package triage

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrStaleCatalog    = errors.New("product catalog older than its max age")
	ErrUnknownProduct  = errors.New("product not in current catalog")
	ErrUnknownSeverity = errors.New("severity not in taxonomy")
)

// Label is what a model proposes.
type Label struct{ Product, Severity string }

// Prompt is everything the model is given. The definitions travel with the
// labels so the model doesn't have to guess what "sev2" means here.
type Prompt struct {
	Ticket     string
	Products   []string
	Severities map[string]string // label -> definition
}

// Model is any classifier: an LLM call, a fine-tuned model, or a rules engine.
type Model interface {
	Classify(ctx context.Context, p Prompt) (Label, error)
}

// Catalog is the Lexicon: which products exist, according to whom, as of when.
type Catalog struct {
	Version  string
	AsOf     time.Time
	MaxAge   time.Duration
	Products map[string]bool // name -> active
}

// System is the infrastructure around the model.
type System struct {
	Catalog    Catalog
	Severities map[string]string // Semantics: label -> definition
	Routes     map[string]string // Pragmatics: severity -> queue
}

// Result is the routed ticket plus the evidence of what it was checked against.
type Result struct {
	Label          Label
	Queue          string
	CatalogVersion string
	Reason         string
}

const TriageQueue = "human-triage"

// Route classifies a ticket with m and decides where it may go. A label the
// system can't verify goes to human triage with a reason; it is never dropped
// and never routed on the model's word alone.
func (s System) Route(ctx context.Context, m Model, ticket string, now time.Time) (Result, error) {
	if now.Sub(s.Catalog.AsOf) > s.Catalog.MaxAge { // a great model can't fix a stale list of products
		return Result{Queue: TriageQueue, Reason: ErrStaleCatalog.Error()}, ErrStaleCatalog
	}
	l, err := m.Classify(ctx, s.prompt(ticket))
	if err != nil {
		return Result{Queue: TriageQueue, Reason: err.Error()}, err
	}
	res := Result{Label: l, CatalogVersion: s.Catalog.Version}
	if !s.Catalog.Products[l.Product] {
		res.Queue, res.Reason = TriageQueue, fmt.Sprintf("%v: %q", ErrUnknownProduct, l.Product)
		return res, ErrUnknownProduct
	}
	queue, ok := s.Routes[l.Severity]
	if _, defined := s.Severities[l.Severity]; !ok || !defined {
		res.Queue, res.Reason = TriageQueue, fmt.Sprintf("%v: %q", ErrUnknownSeverity, l.Severity)
		return res, ErrUnknownSeverity
	}
	res.Queue, res.Reason = queue, "label verified against catalog and taxonomy"
	return res, nil
}

func (s System) prompt(ticket string) Prompt {
	p := Prompt{Ticket: ticket, Severities: s.Severities}
	for name, active := range s.Catalog.Products {
		if active {
			p.Products = append(p.Products, name)
		}
	}
	return p
}
