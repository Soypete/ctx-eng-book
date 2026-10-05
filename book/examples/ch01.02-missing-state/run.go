// Package harness is the runnable companion to the post
// "Your Agent Doesn't Need a Better Prompt. It Needs a Stop Condition."
// (book module 1.2).
//
// Run owns an agent loop. The model proposes each step; the harness checks
// capability and budget before executing it, and every run ends in exactly
// one named Outcome with a recorded reason.
package harness

import (
	"errors"
	"fmt"
)

// Outcome is how a run ended. The prose, the code, and the post's Guidelines
// use the same five.
type Outcome string

const (
	Answered  Outcome = "answered"
	Clarify   Outcome = "needs-clarification"
	Denied    Outcome = "denied"
	Stuck     Outcome = "stuck"
	Exhausted Outcome = "budget-exhausted"
)

var (
	ErrDenied    = errors.New("capability denied")
	ErrAmbiguous = errors.New("ambiguous request")
)

// StepError carries a stable error class, such as a Postgres SQLSTATE, so the
// harness can tell "the same problem again" from "new information" even when
// the message text changes on every guess.
type StepError struct {
	Class string // e.g. "42703" undefined_column
	Msg   string
}

func (e *StepError) Error() string { return e.Class + ": " + e.Msg }

// Step is one model-proposed action.
type Step struct {
	Tool   string
	Tokens int // estimated cost of executing this step
}

// Limits are enforced by the harness, not described in the prompt.
type Limits struct {
	MaxSteps, MaxTokens, MaxSameClass int
	Allowed                           map[string]bool // capabilities granted to this task
}

// Run asks propose for the next step, enforces limits, executes it, and stops
// with a named outcome.
func Run(propose func(attempt int) Step, exec func(Step) (string, error), lim Limits) (Outcome, string) {
	spent, sameClass, lastClass := 0, 0, ""
	for i := 1; i <= lim.MaxSteps; i++ {
		s := propose(i)
		if !lim.Allowed[s.Tool] {
			return Denied, fmt.Sprintf("%s: %s not granted to this task", ErrDenied, s.Tool) // never retry around a permission decision
		}
		if spent+s.Tokens > lim.MaxTokens {
			return Exhausted, fmt.Sprintf("token budget %d would be exceeded at step %d", lim.MaxTokens, i)
		}
		spent += s.Tokens

		out, err := exec(s)
		var se *StepError
		switch {
		case err == nil:
			return Answered, out
		case errors.Is(err, ErrAmbiguous):
			return Clarify, err.Error()
		case errors.As(err, &se):
			if se.Class == lastClass {
				sameClass++
			} else {
				lastClass, sameClass = se.Class, 1
			}
			if sameClass >= lim.MaxSameClass {
				return Stuck, fmt.Sprintf("%d attempts in a row failed with %s; another guess adds no information", sameClass, se.Class)
			}
		default:
			return Stuck, err.Error() // unclassified errors are not safe to retry blindly
		}
	}
	return Exhausted, fmt.Sprintf("stopped after %d steps", lim.MaxSteps)
}
