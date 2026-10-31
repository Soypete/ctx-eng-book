package main

import (
	"testing"
	"time"
)

func TestBoundaryDecisions(t *testing.T) {
	now := time.Date(2026, 11, 9, 12, 0, 0, 0, time.UTC)
	for _, tc := range []struct {
		name     string
		change   func(*Evidence)
		rejected string
	}{
		{"current", func(e *Evidence) {}, ""},
		{"at limit", func(e *Evidence) { e.ObservedAt = now.Add(-time.Hour) }, ""},
		{"stale", func(e *Evidence) { e.ObservedAt = now.Add(-2 * time.Hour) }, "freshness"},
		{"missing version", func(e *Evidence) { e.Version = "" }, "provenance"},
		{"missing time", func(e *Evidence) { e.ObservedAt = time.Time{} }, "freshness"},
		{"future time", func(e *Evidence) { e.ObservedAt = now.Add(time.Second) }, "freshness"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p, err := NewPipeline(SourceBoundaries(time.Hour))
			if err != nil {
				t.Fatal(err)
			}
			e := Evidence{"guide-7", "runbook-store", "rev-3", "approved passage", now}
			tc.change(&e)
			called := false
			answer, ds, err := p.Run(e, now, func(text string) string { called = true; return text })
			if tc.rejected == "" {
				if err != nil || !called || answer != e.Text || len(ds) != 2 {
					t.Fatalf("accepted path: %q, %v, %v", answer, ds, err)
				}
			} else {
				if err == nil || called || answer != "" {
					t.Fatalf("rejection leaked downstream: %q, called=%v, err=%v", answer, called, err)
				}
				d := ds[len(ds)-1]
				if d.Boundary != tc.rejected || d.Outcome != "rejected" || d.Reason == "" {
					t.Fatalf("wrong rejection: %+v", d)
				}
			}
			for _, d := range ds {
				if d.Owner != "source-publishing" || d.PolicyVersion == "" || d.EvidenceID != e.ID || d.Source != e.Source || d.SourceVersion != e.Version || !d.ObservedAt.Equal(e.ObservedAt) || !d.CheckedAt.Equal(now) {
					t.Fatalf("lost repair metadata: %+v", d)
				}
			}
		})
	}
}

func TestOwnershipRequired(t *testing.T) {
	for _, owner := range []string{"", " "} {
		bs := SourceBoundaries(time.Hour)
		bs[1].Owner = owner
		if _, err := NewPipeline(bs); err == nil {
			t.Fatal("unowned boundary accepted")
		}
	}
}

func TestInvalidConfiguration(t *testing.T) {
	for _, mutate := range []func([]Boundary) []Boundary{
		func(bs []Boundary) []Boundary { return nil },
		func(bs []Boundary) []Boundary { bs[0].Name = ""; return bs },
		func(bs []Boundary) []Boundary { bs[0].PolicyVersion = ""; return bs },
		func(bs []Boundary) []Boundary { bs[0].Check = nil; return bs },
		func(bs []Boundary) []Boundary { return append(bs, bs[0]) },
	} {
		if _, err := NewPipeline(mutate(SourceBoundaries(time.Hour))); err == nil {
			t.Fatal("invalid configuration accepted")
		}
	}
	var p Pipeline
	called := false
	if _, _, err := p.Run(Evidence{}, time.Now(), func(s string) string { called = true; return s }); err == nil || called {
		t.Fatal("unconfigured pipeline ran")
	}
}
