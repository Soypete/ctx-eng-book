package harness

import (
	"fmt"
	"strings"
	"testing"
)

var limits = Limits{MaxSteps: 10, MaxTokens: 5000, MaxSameClass: 3, Allowed: map[string]bool{"sql.query": true}}

func sqlStep(int) Step { return Step{Tool: "sql.query", Tokens: 200} }

// The post's opener: every guess at the column name produces different error
// text, but the same SQLSTATE class. The run must end as Stuck, not loop to the
// step limit.
func TestSchemaGuessingEndsStuck(t *testing.T) {
	guesses := []string{"user_id", "userid", "userId", "uid"}
	exec := func(s Step) (string, error) {
		col := guesses[0]
		guesses = guesses[1:]
		return "", &StepError{Class: "42703", Msg: fmt.Sprintf("column %q does not exist", col)}
	}
	o, why := Run(sqlStep, exec, limits)
	if o != Stuck || !strings.Contains(why, "42703") {
		t.Fatalf("got %s (%s), want stuck on 42703", o, why)
	}
}

func TestPermissionDeniedIsTerminal(t *testing.T) {
	propose := func(int) Step { return Step{Tool: "billing.export", Tokens: 10} }
	o, _ := Run(propose, func(Step) (string, error) { return "ok", nil }, limits)
	if o != Denied {
		t.Fatalf("got %s, want denied", o)
	}
}

func TestTokenBudgetStopsTheRun(t *testing.T) {
	big := func(int) Step { return Step{Tool: "sql.query", Tokens: 3000} }
	fail := func(Step) (string, error) { return "", &StepError{Class: "57014", Msg: fmt.Sprint("timeout")} }
	o, _ := Run(big, fail, Limits{MaxSteps: 10, MaxTokens: 5000, MaxSameClass: 5, Allowed: limits.Allowed})
	if o != Exhausted {
		t.Fatalf("got %s, want budget-exhausted", o)
	}
}

func TestAnswerAndClarify(t *testing.T) {
	if o, _ := Run(sqlStep, func(Step) (string, error) { return "42 orders", nil }, limits); o != Answered {
		t.Fatalf("got %s, want answered", o)
	}
	if o, _ := Run(sqlStep, func(Step) (string, error) { return "", ErrAmbiguous }, limits); o != Clarify {
		t.Fatalf("got %s, want needs-clarification", o)
	}
}
