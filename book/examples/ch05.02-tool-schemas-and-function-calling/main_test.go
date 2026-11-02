package main

import (
	"strings"
	"testing"
)

func TestSuiteCoversFiveClasses(t *testing.T) {
	seen := map[Class]bool{}
	for _, c := range suite() {
		seen[c.Class] = true
	}
	for _, want := range []Class{Selection, Abstention, Malformed, Semantic, Unauthorized} {
		if !seen[want] {
			t.Errorf("no case for %s", want)
		}
	}
}

func TestGoodRunPassesEveryCase(t *testing.T) {
	for _, r := range Evaluate(suite(), recorded(goodRun), newHarness()) {
		if !r.Passed {
			t.Errorf("%s/%s: %s", r.Case.Class, r.Case.Name, r.Detail)
		}
	}
}

// A model that reaches for a tool on every request passes the happy path and
// fails abstention. Only a no-tool case catches it.
func TestEagerSelectorFailsAbstention(t *testing.T) {
	eager := func(string, []string) string { return "create_customer" }
	failed := map[Class]int{}
	for _, r := range Evaluate(suite(), eager, newHarness()) {
		if !r.Passed {
			failed[r.Case.Class]++
		}
	}
	if failed[Abstention] == 0 {
		t.Fatal("abstention case did not catch an eager selector")
	}
}

func TestOnlyFullyCheckedCallsExecute(t *testing.T) {
	h := newHarness()
	Evaluate(suite(), recorded(goodRun), h)
	if len(h.Executed) != 1 || h.Executed[0].Args["email"] != "ben@example.com" {
		t.Fatalf("executed %v", h.Executed)
	}
}

func TestSchemaValidDuplicateIsSemanticNotSyntax(t *testing.T) {
	err := newHarness().Check(agent, Call{"create_customer", map[string]any{"email": "ana@example.com", "name": "Ana"}})
	if err == nil || !strings.HasPrefix(err.Error(), "semantic") {
		t.Fatalf("got %v", err)
	}
}

func TestDescriptionSaysWhenNotToUse(t *testing.T) {
	if !strings.Contains(createCustomer.Description(), "Do not use when") {
		t.Fatal("description has no exclusion criteria")
	}
}
