package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

// Proposal contains model-selected values, never the caller's identity.
type Proposal struct {
	Dataset string
	Start   string
	End     string
}

type Policy struct {
	Version string
	Exports map[string]map[string]bool
}

type Decision struct {
	Principal, ContractVersion, PolicyVersion string
	Stage, Outcome, Reason                    string
}

// Parse implements this example's fixed schema.json contract, not JSON Schema
// generally. Reject duplicate keys too, so consumers cannot disagree on values.
func Parse(raw string) (Proposal, error) {
	d := json.NewDecoder(strings.NewReader(raw))
	tok, err := d.Token()
	if err != nil || tok != json.Delim('{') {
		return Proposal{}, fmt.Errorf("expected JSON object")
	}
	fields := map[string]string{}
	for d.More() {
		key, err := d.Token()
		if err != nil {
			return Proposal{}, err
		}
		name, ok := key.(string)
		if !ok || (name != "dataset" && name != "start" && name != "end") {
			return Proposal{}, fmt.Errorf("unknown field")
		}
		if _, exists := fields[name]; exists {
			return Proposal{}, fmt.Errorf("duplicate field")
		}
		var value any
		if err := d.Decode(&value); err != nil {
			return Proposal{}, err
		}
		s, ok := value.(string)
		if !ok || s == "" {
			return Proposal{}, fmt.Errorf("expected nonempty string")
		}
		fields[name] = s
	}
	if _, err := d.Token(); err != nil {
		return Proposal{}, err
	}
	if _, err := d.Token(); err != io.EOF {
		return Proposal{}, fmt.Errorf("trailing content")
	}
	if len(fields) != 3 {
		return Proposal{}, fmt.Errorf("missing required field")
	}
	return Proposal{fields["dataset"], fields["start"], fields["end"]}, nil
}

// Run takes a trusted session principal and policy separately from model text.
// execute stands in for the export service; no external writes occur in the demo.
func Run(raw, principal string, policy Policy, execute func(Proposal) error) Decision {
	d := Decision{Principal: principal, ContractVersion: "v1", PolicyVersion: policy.Version,
		Stage: "schema", Outcome: "rejected"}
	p, err := Parse(raw)
	if err != nil {
		d.Reason = "output does not match contract"
		return d
	}
	start, startErr := time.Parse("2006-01-02", p.Start)
	end, endErr := time.Parse("2006-01-02", p.End)
	// Well-typed strings can still describe an impossible reporting interval.
	if startErr != nil || endErr != nil || !start.Before(end) {
		d.Stage, d.Reason = "semantic", "require valid dates and start before end"
		return d
	}
	// Permission comes from the trusted session and policy, never model JSON.
	if principal == "" || policy.Version == "" || !policy.Exports[principal][p.Dataset] {
		d.Stage, d.Reason = "authorization", "export not permitted"
		return d
	}
	d.Stage = "execution"
	if err := execute(p); err != nil {
		d.Outcome, d.Reason = "failed", "export service failed"
		return d
	}
	d.Outcome, d.Reason = "executed", "export service completed"
	return d
}

func main() {
	policy := Policy{"demo-v1", map[string]map[string]bool{"analyst": {"sales": true}}}
	for _, raw := range []string{
		`{"dataset":"sales","start":"2026-11-01","end":"2026-10-01"}`,
		`{"dataset":"payroll","start":"2026-10-01","end":"2026-11-01"}`,
		`{"dataset":"sales","start":"2026-10-01","end":"2026-11-01"}`,
	} {
		calls := 0
		d := Run(raw, "analyst", policy, func(Proposal) error { calls++; return nil })
		fmt.Printf("%s: %s; exports=%d; reason=%s\n", d.Stage, d.Outcome, calls, d.Reason)
	}
}
