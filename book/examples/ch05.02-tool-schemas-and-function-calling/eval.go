package main

import (
	"errors"
	"fmt"
	"sort"
)

// Class names the five failure classes every tool should be tested against.
type Class string

const (
	Selection    Class = "correct-selection"
	Abstention   Class = "no-tool-abstention"
	Malformed    Class = "malformed-arguments"
	Semantic     Class = "semantically-invalid"
	Unauthorized Class = "unauthorized"
	Control      Class = "valid-control" // a call that should execute, so the checks aren't just rejecting everything
)

// Selector stands in for the model: given a request and the tool
// descriptions, it returns the tool it would call, or "" for none.
type Selector func(request string, descriptions []string) string

// Case is one eval row. Selection and abstention cases test the Selector;
// the other three hand the harness a fixed call and expect a specific stage to stop it.
type Case struct {
	Name      string
	Class     Class
	Request   string
	WantTool  string // "" means the right answer is no tool
	Principal Principal
	Call      *Call
	WantErr   error // nil means the call should execute
}

// Result says whether a case passed and why.
type Result struct {
	Case   Case
	Passed bool
	Detail string
}

// Evaluate runs every case. Selection is scored against WantTool; harness cases
// are scored by which stage, if any, rejected the call.
func Evaluate(cases []Case, sel Selector, h *Harness) []Result {
	var descs []string
	for _, k := range h.Contracts {
		descs = append(descs, k.Description())
	}
	sort.Strings(descs)
	var out []Result
	for _, c := range cases {
		if c.Call == nil {
			got := sel(c.Request, descs)
			out = append(out, Result{c, got == c.WantTool, fmt.Sprintf("selected %q, want %q", got, c.WantTool)})
			continue
		}
		err := h.Check(c.Principal, *c.Call)
		pass := (c.WantErr == nil && err == nil) || (c.WantErr != nil && errors.Is(err, c.WantErr))
		out = append(out, Result{c, pass, fmt.Sprintf("got %v, want %v", err, c.WantErr)})
	}
	return out
}
