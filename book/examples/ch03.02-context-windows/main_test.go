package main

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func TestRequiredEvidenceSurvivesOptionalOverflow(t *testing.T) {
	sources := fixture()
	sources[0].Text = strings.Repeat("optional investigation detail ", 100)
	task := "Describe the approved restart procedure."
	required, _ := json.Marshal(Input{Task: task, Evidence: []Source{sources[1]}})
	payload, manifest, err := Assemble(task, sources, len(required), "bytes", byteCount)
	if err != nil {
		t.Fatal(err)
	}
	if string(payload) != string(required) {
		t.Fatalf("required payload changed: %s", payload)
	}
	if manifest.Used != len(payload) || manifest.Used > manifest.Limit {
		t.Fatal("budget violated")
	}
	omitted, included := manifest.Decisions[0], manifest.Decisions[1]
	if omitted.Action != "omitted" || omitted.Position != 0 || omitted.Reason == "" {
		t.Fatalf("missing omission: %+v", omitted)
	}
	if included.Action != "included" || included.Position != 1 || included.Version != sources[1].Version || included.ConfirmedAt != sources[1].ConfirmedAt || included.Scope != sources[1].Scope {
		t.Fatalf("lost provenance: %+v", included)
	}
}

func TestRequiredOverflowEmitsNoRequest(t *testing.T) {
	sources := fixture()
	required, _ := json.Marshal(Input{Task: "restart", Evidence: []Source{sources[1]}})
	payload, manifest, err := Assemble("restart", sources, len(required)-1, "bytes", byteCount)
	if !errors.Is(err, ErrRequired) || payload != nil {
		t.Fatalf("expected no request, got %s, %v", payload, err)
	}
	for _, d := range manifest.Decisions {
		if d.Action != "rejected" || d.Position != 0 {
			t.Fatalf("false inclusion: %+v", d)
		}
	}
}

func TestBudgetIncludesTaskAndMetadata(t *testing.T) {
	s := fixture()[1]
	payload, _, err := Assemble("restart", []Source{s}, len(s.Text), "bytes", byteCount)
	if !errors.Is(err, ErrRequired) || payload != nil {
		t.Fatal("counted only source text")
	}
}

func TestOptionalEvidenceIncludedWhenItFits(t *testing.T) {
	sources := fixture()
	all, _ := json.Marshal(Input{Task: "restart", Evidence: []Source{sources[1], sources[0]}})
	payload, m, err := Assemble("restart", sources, len(all), "bytes", byteCount)
	if err != nil || string(payload) != string(all) {
		t.Fatalf("wrong payload: %s, %v", payload, err)
	}
	if m.Decisions[0].Position != 2 || m.Decisions[1].Position != 1 {
		t.Fatal("positions do not match serialized evidence")
	}
}

func TestCounterFailureEmitsNoRequest(t *testing.T) {
	failure := errors.New("counter unavailable")
	payload, _, err := Assemble("restart", fixture(), 1000, "tokens", func([]byte) (int, error) { return 0, failure })
	if !errors.Is(err, failure) || payload != nil {
		t.Fatal("counter failure must stop assembly")
	}
}
