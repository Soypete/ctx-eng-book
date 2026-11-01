package main

import (
	"errors"
	"testing"
)

func TestValidationBeforeExecution(t *testing.T) {
	good := `{"dataset":"sales","start":"2026-10-01","end":"2026-11-01"}`
	for _, tc := range []struct {
		name, raw, principal, version, stage, outcome string
		calls                                         int
		serviceErr                                    bool
	}{
		{"malformed", `{`, "analyst", "v1", "schema", "rejected", 0, false},
		{"wrong type", `{"dataset":3,"start":"x","end":"y"}`, "analyst", "v1", "schema", "rejected", 0, false},
		{"missing", `{"dataset":"sales"}`, "analyst", "v1", "schema", "rejected", 0, false},
		{"reversed", `{"dataset":"sales","start":"2026-11-01","end":"2026-10-01"}`, "analyst", "v1", "semantic", "rejected", 0, false},
		{"equal", `{"dataset":"sales","start":"2026-10-01","end":"2026-10-01"}`, "analyst", "v1", "semantic", "rejected", 0, false},
		{"invalid date", `{"dataset":"sales","start":"2026-02-30","end":"2026-11-01"}`, "analyst", "v1", "semantic", "rejected", 0, false},
		{"forbidden", `{"dataset":"payroll","start":"2026-10-01","end":"2026-11-01"}`, "analyst", "v1", "authorization", "rejected", 0, false},
		{"no identity", good, "", "v1", "authorization", "rejected", 0, false},
		{"no policy version", good, "analyst", "", "authorization", "rejected", 0, false},
		{"valid", good, "analyst", "v1", "execution", "executed", 1, false},
		{"service failure", good, "analyst", "v1", "execution", "failed", 1, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Explicitly establish that semantic/authority failures passed the schema.
			if tc.stage != "schema" {
				if _, err := Parse(tc.raw); err != nil {
					t.Fatal(err)
				}
			}
			calls := 0
			policy := Policy{tc.version, map[string]map[string]bool{"analyst": {"sales": true}}}
			d := Run(tc.raw, tc.principal, policy, func(p Proposal) error {
				calls++
				if p != (Proposal{"sales", "2026-10-01", "2026-11-01"}) {
					t.Fatalf("unexpected export: %+v", p)
				}
				if tc.serviceErr {
					return errors.New("private service detail")
				}
				return nil
			})
			if d.Stage != tc.stage || d.Outcome != tc.outcome || calls != tc.calls {
				t.Fatalf("decision=%+v calls=%d", d, calls)
			}
			if d.Reason == "" || d.ContractVersion != "v1" || d.PolicyVersion != tc.version || d.Principal != tc.principal {
				t.Fatalf("incomplete decision: %+v", d)
			}
		})
	}
}

func TestRejectAmbiguousJSON(t *testing.T) {
	for _, raw := range []string{
		`null`, `[]`, `{"dataset":null,"start":"x","end":"y"}`,
		`{"dataset":"","start":"x","end":"y"}`,
		`{"dataset":"sales","start":"x","end":"y","principal":"admin"}`,
		`{"dataset":"sales","dataset":"payroll","start":"x","end":"y"}`,
		`{"dataset":"sales","start":"x","end":"y"} {}`,
	} {
		if _, err := Parse(raw); err == nil {
			t.Errorf("accepted %s", raw)
		}
	}
}
