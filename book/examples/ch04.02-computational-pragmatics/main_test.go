package main

import (
	"errors"
	"testing"
	"time"
)

func TestScopeBeforeRead(t *testing.T) {
	tests := []struct {
		name, outcome string
		mutate        func(*Session, *Request, *Policy)
	}{
		{"resolved", "read", func(*Session, *Request, *Policy) {}},
		{"missing dataset", "clarify", func(_ *Session, r *Request, _ *Policy) { r.Dataset = "" }},
		{"missing action", "clarify", func(_ *Session, r *Request, _ *Policy) { r.Action = "" }},
		{"missing evidence", "clarify", func(_ *Session, r *Request, _ *Policy) { r.ResolutionEvidence = "" }},
		{"missing time", "clarify", func(_ *Session, r *Request, _ *Policy) { r.Start = time.Time{} }},
		{"reversed time", "clarify", func(_ *Session, r *Request, _ *Policy) { r.Start = r.End.Add(time.Hour) }},
		{"empty interval", "clarify", func(_ *Session, r *Request, _ *Policy) { r.Start = r.End }},
		{"wrong dataset", "denied", func(_ *Session, r *Request, _ *Policy) { r.Dataset = "marketing" }},
		{"write request", "denied", func(_ *Session, r *Request, _ *Policy) { r.Action = "delete" }},
		{"anonymous", "denied", func(s *Session, _ *Request, _ *Policy) { s.Principal = "" }},
		{"other caller", "denied", func(s *Session, _ *Request, _ *Policy) { s.Principal = "other" }},
		{"unversioned policy", "denied", func(_ *Session, _ *Request, p *Policy) { p.Version = "" }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, r, p := fixture()
			tt.mutate(&s, &r, &p)
			calls := 0
			value, d := Execute(s, r, p, func(dataset string, start, end time.Time) (string, error) {
				calls++
				if dataset != r.Dataset || start != r.Start || end != r.End {
					t.Fatal("read scope changed")
				}
				return "rows", nil
			})
			if d.Outcome != tt.outcome {
				t.Fatalf("outcome=%s, want %s", d.Outcome, tt.outcome)
			}
			wantCalls := 0
			if tt.outcome == "read" {
				wantCalls = 1
			}
			if calls != wantCalls {
				t.Fatalf("reads=%d, want %d", calls, wantCalls)
			}
			if wantCalls == 0 && value != "" {
				t.Fatal("rejected request returned data")
			}
			if d.Principal != s.Principal || d.Dataset != r.Dataset || d.ResolutionEvidence != r.ResolutionEvidence || d.PolicyVersion != p.Version || d.Reason == "" {
				t.Fatal("incomplete decision record")
			}
		})
	}
}

func TestReadFailureReturnsNoData(t *testing.T) {
	s, r, p := fixture()
	value, d := Execute(s, r, p, func(string, time.Time, time.Time) (string, error) {
		return "partial data", errors.New("sensitive source details")
	})
	if value != "" || d.Outcome != "failed" || d.Reason != "source read failed" {
		t.Fatalf("value=%q decision=%+v", value, d)
	}
}
