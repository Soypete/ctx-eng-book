package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Source metadata comes from trusted adapters, after authorization and freshness checks.
type Source struct {
	ID          string
	Version     string
	ConfirmedAt string
	RetrievedAt string
	Scope       string
	Required    bool
	Text        string
}

type Input struct {
	Task     string
	Evidence []Source
}

type Decision struct {
	ID          string
	Version     string
	ConfirmedAt string
	RetrievedAt string
	Scope       string
	Action      string
	Reason      string
	Position    int // One-based evidence position; zero means absent from the payload.
}

type Manifest struct {
	Unit      string
	Limit     int
	Used      int
	Decisions []Decision
}

// Measure must count the complete serialized request in the budget's units.
// A production adapter must include provider framing and use its target tokenizer.
type Measure func([]byte) (int, error)

var ErrRequired = errors.New("required evidence exceeds input budget")

func Assemble(task string, sources []Source, limit int, unit string, measure Measure) ([]byte, Manifest, error) {
	m := Manifest{Unit: unit, Limit: limit}
	if limit <= 0 || measure == nil || unit == "" {
		return nil, m, errors.New("positive limit, unit and measure are required")
	}
	seen := make(map[string]bool)
	for _, s := range sources {
		if s.ID == "" || seen[s.ID] {
			return nil, m, errors.New("source IDs must be nonempty and unique")
		}
		seen[s.ID] = true
		m.Decisions = append(m.Decisions, Decision{
			ID: s.ID, Version: s.Version, ConfirmedAt: s.ConfirmedAt,
			RetrievedAt: s.RetrievedAt, Scope: s.Scope,
			Action: "withheld", Reason: "request not assembled",
		})
	}
	input := Input{Task: task, Evidence: []Source{}}
	for _, s := range sources {
		if s.Required {
			input.Evidence = append(input.Evidence, s)
		}
	}
	encode := func(in Input) ([]byte, int, error) {
		payload, err := json.Marshal(in)
		if err != nil {
			return nil, 0, err
		}
		used, err := measure(payload)
		if err != nil {
			return nil, 0, err
		}
		if used < 0 {
			return nil, 0, errors.New("negative measurement")
		}
		return payload, used, nil
	}
	payload, used, err := encode(input)
	if err != nil {
		return nil, m, err
	}
	if used > limit {
		for i := range m.Decisions {
			m.Decisions[i].Action = "rejected"
			m.Decisions[i].Reason = "required working set does not fit; no request emitted"
		}
		return nil, m, ErrRequired
	}
	for i, s := range sources {
		if s.Required {
			continue
		}
		candidate := Input{Task: task, Evidence: append(append([]Source{}, input.Evidence...), s)}
		next, size, err := encode(candidate)
		if err != nil {
			return nil, m, err
		}
		if size > limit {
			m.Decisions[i].Action = "omitted"
			m.Decisions[i].Reason = "optional evidence exceeds remaining budget"
			continue // Preserve required evidence instead of clipping the serialized request.
		}
		input, payload, used = candidate, next, size
	}
	for position, s := range input.Evidence {
		for i := range m.Decisions {
			if m.Decisions[i].ID == s.ID {
				m.Decisions[i].Action = "included"
				m.Decisions[i].Reason = "fits after reserving required evidence"
				m.Decisions[i].Position = position + 1
			}
		}
	}
	m.Used = used
	return payload, m, nil
}

// Bytes are an honest deterministic fixture unit, NOT an estimate of model tokens.
func byteCount(payload []byte) (int, error) { return len(payload), nil }

func fixture() []Source {
	return []Source{
		{ID: "incident-history", Version: "history-v1", ConfirmedAt: "2026-11-12T09:00:00Z", RetrievedAt: "2026-11-12T09:01:00Z", Scope: "checkout-read", Text: "Earlier investigation notes; optional background."},
		{ID: "checkout-runbook", Version: "runbook-v3", ConfirmedAt: "2026-11-12T09:00:00Z", RetrievedAt: "2026-11-12T09:01:00Z", Scope: "checkout-read", Required: true, Text: "Drain checkout traffic before restarting workers."},
	}
}

func main() {
	const task = "Describe the approved checkout restart procedure."
	sources := fixture()
	required, _ := json.Marshal(Input{Task: task, Evidence: []Source{sources[1]}})
	// Set the fixture limit to exactly the required payload, including its metadata.
	payload, manifest, err := Assemble(task, sources, len(required), "bytes (fixture only)", byteCount)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("INPUT:", string(payload))
	record, _ := json.MarshalIndent(manifest, "", "  ")
	fmt.Println("MANIFEST:", string(record))
}
