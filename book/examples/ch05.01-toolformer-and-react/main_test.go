package main

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestSameDefinitionBothProviders(t *testing.T) {
	w := weatherTool()
	oa := OpenAIFormat(w)["function"].(map[string]any)["parameters"]
	an := AnthropicFormat(w)["input_schema"]
	if !reflect.DeepEqual(oa, an) {
		t.Fatalf("provider formats diverged:\n%v\n%v", oa, an)
	}
}

func TestHarnessRejectsBeforeExecution(t *testing.T) {
	cases := []struct {
		name string
		p    Proposal
		want error
	}{
		{"unknown tool", Proposal{"drop_tables", json.RawMessage(`{}`)}, ErrUnknownTool},
		{"registered, not granted", Proposal{"delete_account", json.RawMessage(`{}`)}, ErrNotGranted},
		{"missing field", Proposal{"get_weather", json.RawMessage(`{"unit":"celsius"}`)}, ErrBadArgs},
		{"bad enum", Proposal{"get_weather", json.RawMessage(`{"location":"Provo","unit":"kelvin"}`)}, ErrBadArgs},
		{"extra field", Proposal{"get_weather", json.RawMessage(`{"location":"Provo","admin":"true"}`)}, ErrBadArgs},
		{"schema-valid but empty", Proposal{"get_weather", json.RawMessage(`{"location":"  "}`)}, ErrBadArgs},
		{"not JSON", Proposal{"get_weather", json.RawMessage(`location=Provo`)}, ErrBadArgs},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			h := newHarness()
			if _, err := h.Handle(c.p); !errors.Is(err, c.want) {
				t.Fatalf("got %v, want %v", err, c.want)
			}
			if len(h.Ran) != 0 {
				t.Fatalf("rejected proposal executed: %v", h.Ran)
			}
		})
	}
}

func TestRegisteredButNotGranted(t *testing.T) {
	h := newHarness()
	h.Granted = map[string]bool{}
	_, err := h.Handle(Proposal{"get_weather", json.RawMessage(`{"location":"Provo"}`)})
	if !errors.Is(err, ErrNotGranted) || len(h.Ran) != 0 {
		t.Fatalf("got %v ran=%v", err, h.Ran)
	}
}

func TestObservationCarriesSourceAndTime(t *testing.T) {
	h := newHarness()
	obs, err := h.Handle(Proposal{"get_weather", json.RawMessage(`{"location":"Provo"}`)})
	if err != nil {
		t.Fatal(err)
	}
	if obs.Source == "" || obs.ObservedAt.IsZero() {
		t.Fatalf("observation missing provenance: %+v", obs)
	}
}

func TestToolResultCannotWidenGrants(t *testing.T) {
	h := newHarness()
	script := []Step{
		{Proposal: &Proposal{"get_weather", json.RawMessage(`{"location":"Provo"}`)}},
		{Proposal: &Proposal{"delete_account", json.RawMessage(`{}`)}},
		{Answer: "done"},
	}
	_, history, err := Loop(h, func(hist []Turn) Step { return script[len(hist)] }, 5)
	if err != nil {
		t.Fatal(err)
	}
	if !errors.Is(history[1].Rejected, ErrNotGranted) {
		t.Fatalf("injected tool request not rejected: %v", history[1].Rejected)
	}
	if !reflect.DeepEqual(h.Ran, []string{"get_weather"}) {
		t.Fatalf("executed %v", h.Ran)
	}
}

func TestLoopStopsWithoutAnswer(t *testing.T) {
	h := newHarness()
	p := &Proposal{"get_weather", json.RawMessage(`{"location":"Provo"}`)}
	_, history, err := Loop(h, func([]Turn) Step { return Step{Proposal: p} }, 3)
	if !errors.Is(err, ErrStepLimit) || len(history) != 3 {
		t.Fatalf("got %v after %d turns", err, len(history))
	}
}
