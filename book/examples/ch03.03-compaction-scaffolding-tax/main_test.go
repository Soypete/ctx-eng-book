package main

import (
	"strings"
	"testing"
	"time"
)

func TestCompactionRequiresSourceToRecoverException(t *testing.T) {
	now := time.Date(2026, 11, 16, 12, 0, 0, 0, time.UTC)
	p := fixture(now)
	s := Compact(p)
	if strings.Contains(s.Text, "Annual") || len(s.Text) >= len(p.General)+len(p.Exceptions["annual"]) {
		t.Fatal("fixture must omit the exception and reduce text")
	}
	if s.Manifest.SourceVersion != p.Version || s.Manifest.Method != "general-only/v1" || s.Manifest.Deferred[0] != "exceptions" {
		t.Fatal("missing lineage or deferred detail")
	}
	store := map[string]Policy{p.ID: p}
	fetch := Reader(store, "support", now, time.Hour)
	answer, d := Resolve(s, "annual", fetch)
	if answer != p.Exceptions["annual"] || d.Action != "rehydrated" || d.SourceID != p.ID {
		t.Fatalf("must match uncompacted exception: %q %+v", answer, d)
	}
	general, _ := Resolve(s, "monthly", fetch)
	if general != p.General {
		t.Fatalf("unexpected general rule: %q", general)
	}
	delete(store, p.ID)
	answer, d = Resolve(s, "annual", fetch)
	if answer != "" || d.Action != "withheld" || d.Reason != "source unavailable" {
		t.Fatalf("must not fall back to summary: %q %+v", answer, d)
	}
}

func TestRecoveryRejectsInvalidSource(t *testing.T) {
	now := time.Date(2026, 11, 16, 12, 0, 0, 0, time.UTC)
	for _, tc := range []struct {
		name   string
		change func(*Policy)
	}{
		{"changed version", func(p *Policy) { p.Version = "v8" }},
		{"changed identity", func(p *Policy) { p.ID = "other" }},
		{"revoked scope", func(p *Policy) { p.Scope = "restricted" }},
		{"stale source", func(p *Policy) { p.ConfirmedAt = now.Add(-2 * time.Hour) }},
		{"future timestamp", func(p *Policy) { p.ConfirmedAt = now.Add(time.Hour) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := fixture(now)
			s := Compact(p)
			tc.change(&p)
			answer, d := Resolve(s, "annual", Reader(map[string]Policy{s.Manifest.SourceID: p}, "support", now, time.Hour))
			if answer != "" || d.Action != "withheld" || d.Reason == "" {
				t.Fatalf("invalid recovery returned %q, %+v", answer, d)
			}
		})
	}
}
