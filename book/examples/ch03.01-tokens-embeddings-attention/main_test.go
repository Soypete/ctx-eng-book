package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestAdmissionAndAblation(t *testing.T) {
	now := time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)
	sources := fixture(now)
	payload, decisions, err := Assemble("Restart checkout?", "checkout", sources, now, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	var input Input
	if err := json.Unmarshal(payload, &input); err != nil {
		t.Fatal(err)
	}
	want := []Evidence{{"runbook", "v4", sources[3].Text}}
	if !reflect.DeepEqual(input.Evidence, want) {
		t.Fatalf("evidence = %+v", input.Evidence)
	}
	for i, reason := range []string{"stale", "wrong service", "unauthorized", "included"} {
		if decisions[i].Reason != reason || decisions[i].Version != sources[i].Version || !decisions[i].ConfirmedAt.Equal(sources[i].ConfirmedAt) {
			t.Fatalf("decision %d = %+v", i, decisions[i])
		}
	}
	for _, s := range sources[:3] {
		if strings.Contains(string(payload), s.Text) {
			t.Fatalf("rejected text leaked: %s", s.ID)
		}
	}
	// Removing a rejected source must not change the serialized model input.
	withoutRejected, _, err := Assemble("Restart checkout?", "checkout", sources[1:], now, time.Hour)
	if err != nil || string(withoutRejected) != string(payload) {
		t.Fatal("rejected source affected input")
	}
	// Removing the only admitted source must stop generation, not use a fallback.
	empty, records, err := Assemble("Restart checkout?", "checkout", sources[:3], now, time.Hour)
	if err == nil || empty != nil || len(records) != 3 {
		t.Fatal("missing evidence did not stop assembly")
	}
}

func TestMetadataAndFreshness(t *testing.T) {
	now := time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)
	for _, tc := range []struct {
		name     string
		change   func(*Source)
		admitted bool
	}{
		{"at limit", func(s *Source) { s.ConfirmedAt = now.Add(-time.Hour) }, true},
		{"past limit", func(s *Source) { s.ConfirmedAt = now.Add(-time.Hour - time.Nanosecond) }, false},
		{"future", func(s *Source) { s.ConfirmedAt = now.Add(time.Second) }, false},
		{"missing time", func(s *Source) { s.ConfirmedAt = time.Time{} }, false},
		{"missing version", func(s *Source) { s.Version = "" }, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := fixture(now)[3]
			tc.change(&s)
			payload, _, err := Assemble("Restart checkout?", "checkout", []Source{s}, now, time.Hour)
			if (err == nil) != tc.admitted || (payload != nil) != tc.admitted {
				t.Fatalf("payload=%s err=%v", payload, err)
			}
		})
	}
}
