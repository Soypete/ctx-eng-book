package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Request is resolved by a trusted UI or clarification flow, not a model guess.
type Request struct {
	Action, Dataset, ResolutionEvidence string
	Start, End                          time.Time // Inclusive start, exclusive end; UTC in this example.
}

// Session and Policy come from authenticated application state.
type Session struct{ Principal string }
type Policy struct {
	Version string
	Readers map[string]map[string]bool
}
type Decision struct {
	Principal, Action, Dataset, ResolutionEvidence, PolicyVersion string
	Start, End                                                    time.Time
	Outcome, Reason                                               string
}

type Reader func(dataset string, start, end time.Time) (string, error)

func Execute(s Session, r Request, p Policy, read Reader) (string, Decision) {
	d := Decision{Principal: s.Principal, Action: r.Action, Dataset: r.Dataset,
		ResolutionEvidence: r.ResolutionEvidence, PolicyVersion: p.Version, Start: r.Start, End: r.End}
	if s.Principal == "" {
		d.Outcome, d.Reason = "denied", "authentication required"
		return "", d
	}
	// Missing scope must not become a best-guess query.
	if r.Action == "" || r.Dataset == "" || r.ResolutionEvidence == "" ||
		r.Start.IsZero() || r.End.IsZero() || !r.Start.Before(r.End) {
		d.Outcome, d.Reason = "clarify", "resolve action, dataset, interval, and evidence"
		return "", d
	}
	// A complete interpretation still does not grant permission.
	if r.Action != "analyze" || p.Version == "" || !p.Readers[s.Principal][r.Dataset] {
		d.Outcome, d.Reason = "denied", "operation or dataset outside policy"
		return "", d
	}
	value, err := read(r.Dataset, r.Start, r.End)
	if err != nil {
		d.Outcome, d.Reason = "failed", "source read failed"
		return "", d
	}
	d.Outcome, d.Reason = "read", "resolved scope allowed by policy"
	return value, d
}

func fixture() (Session, Request, Policy) {
	return Session{"miriah"}, Request{
		Action: "analyze", Dataset: "sales", ResolutionEvidence: "selection:event-demo",
		Start: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC),
	}, Policy{"demo-v1", map[string]map[string]bool{"miriah": {"sales": true}}}
}

func main() {
	s, resolved, policy := fixture()
	ambiguous := resolved
	ambiguous.Dataset, ambiguous.ResolutionEvidence = "", ""
	denied := resolved
	denied.Dataset = "marketing"
	for _, request := range []Request{ambiguous, denied, resolved} {
		reads := 0
		_, decision := Execute(s, request, policy, func(dataset string, start, end time.Time) (string, error) {
			reads++
			return "invented sales rows", nil
		})
		encoded, err := json.Marshal(decision)
		if err != nil {
			panic(err)
		}
		fmt.Printf("reads=%d decision=%s\n", reads, encoded)
	}
}
