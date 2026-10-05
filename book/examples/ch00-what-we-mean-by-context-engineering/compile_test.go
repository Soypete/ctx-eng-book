package contextcompile

import (
	"reflect"
	"testing"
	"time"
)

// The post's example: a higher-scoring restricted salary sheet is scored by
// search but never admitted, and an oversized note is cut by the budget.
func TestCompileExcludesUnauthorizedAndOverBudget(t *testing.T) {
	allowed := func(actor, purpose string, d Doc) bool { return d.ID != "salary" }
	cands := []Doc{
		{ID: "pricing", Score: 0.90, Tokens: 300},
		{ID: "salary", Score: 0.95, Tokens: 200},
		{ID: "note", Score: 0.40, Tokens: 900},
	}
	m := Manifest{Actor: "rep", Purpose: "pricing-help", Budget: 1000, At: time.Now(),
		SourceVersion: "index@2026-10-01", PolicyVersion: "acl-v12", RankerVersion: "bm25+e5-v3"}

	out, got := Compile(m, allowed, cands)

	if len(out) != 1 || out[0].ID != "pricing" {
		t.Fatalf("admitted %v, want only pricing", out)
	}
	if want := []string{"salary:unauthorized", "note:budget"}; !reflect.DeepEqual(got.Excluded, want) {
		t.Fatalf("excluded %v, want %v", got.Excluded, want)
	}
	if got.PolicyVersion != "acl-v12" {
		t.Fatalf("manifest lost policy version")
	}
}

// Ties break by ID, so replaying the same candidates gives the same order.
func TestCompileTieBreakIsDeterministic(t *testing.T) {
	all := func(string, string, Doc) bool { return true }
	cands := []Doc{{ID: "b", Score: 0.5, Tokens: 1}, {ID: "a", Score: 0.5, Tokens: 1}}
	_, m := Compile(Manifest{Budget: 10}, all, cands)
	if want := []string{"a", "b"}; !reflect.DeepEqual(m.Included, want) {
		t.Fatalf("included %v, want %v", m.Included, want)
	}
}
