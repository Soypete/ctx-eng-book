// Package main shows the split between a model's tool proposal and the harness's
// decision to execute it. One tool definition is rendered for two provider
// request formats; a Harness checks every proposal before anything runs.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"
)

// Tool is the harness-owned definition of one operation. The model sees the
// name, description, and parameter schema; it never sees the handler.
type Tool struct {
	Name        string
	Description string
	Parameters  map[string]any // JSON Schema for the arguments
	Source      string         // where results come from, attached to observations
	Handler     func(args map[string]any) (string, error)
	Validate    func(args map[string]any) error // semantic checks the schema can't express
	SideEffect  bool                            // changes something; needs the user's confirmation
}

// OpenAIFormat renders the definition in the shape OpenAI's chat completions
// API accepts in its tools field.
func OpenAIFormat(t Tool) map[string]any {
	return map[string]any{
		"type": "function",
		"function": map[string]any{
			"name":        t.Name,
			"description": t.Description,
			"parameters":  t.Parameters,
		},
	}
}

// AnthropicFormat renders the same definition in the shape Anthropic's
// Messages API accepts in its tools field.
func AnthropicFormat(t Tool) map[string]any {
	return map[string]any{
		"name":         t.Name,
		"description":  t.Description,
		"input_schema": t.Parameters,
	}
}

// Proposal is what a model returns when it asks for a tool: a name and raw
// JSON arguments. It carries no authority.
type Proposal struct {
	Tool      string
	Arguments json.RawMessage
}

// Observation is a tool result returned to the model, labeled with where it
// came from and when. It is external data, not verified truth.
type Observation struct {
	Tool       string
	Source     string
	ObservedAt time.Time
	Content    string
}

// Rejection reasons. Each one stops a proposal before execution.
var (
	ErrUnknownTool = errors.New("unknown tool")
	ErrNotGranted  = errors.New("tool not granted to this task")
	ErrBadArgs     = errors.New("arguments do not match the tool contract")
	ErrNeedsUser   = errors.New("side-effecting call needs the user's confirmation")
)

// Harness owns the action boundary: the registry, the task's grants, and the clock.
type Harness struct {
	Tools   map[string]Tool
	Granted map[string]bool // set from the authenticated task, never from model output
	Now     func() time.Time
	// Confirm asks the signed-in user to approve one specific side-effecting call,
	// outside the model's context. Nil means nobody can confirm, so those calls stop.
	Confirm func(p Proposal) bool
	Ran     []string // tools actually executed, for tests and audit
}

// Definitions returns only the tools granted to this task, so the model is never
// offered a tool it may not use. Handle still checks grants as the backstop.
func (h *Harness) Definitions() []Tool {
	var out []Tool
	for name, t := range h.Tools {
		if h.Granted[name] {
			out = append(out, t)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// Handle checks a proposal and executes it only if every check passes.
func (h *Harness) Handle(p Proposal) (Observation, error) {
	t, ok := h.Tools[p.Tool]
	if !ok {
		return Observation{}, fmt.Errorf("%w: %q", ErrUnknownTool, p.Tool)
	}
	if !h.Granted[p.Tool] { // grants come from the authenticated task, not the model
		return Observation{}, fmt.Errorf("%w: %s", ErrNotGranted, p.Tool)
	}
	var args map[string]any
	if err := json.Unmarshal(p.Arguments, &args); err != nil {
		return Observation{}, fmt.Errorf("%w: %v", ErrBadArgs, err)
	}
	if err := checkSchema(t.Parameters, args); err != nil {
		return Observation{}, fmt.Errorf("%w: %v", ErrBadArgs, err)
	}
	if t.Validate != nil {
		if err := t.Validate(args); err != nil {
			return Observation{}, fmt.Errorf("%w: %v", ErrBadArgs, err)
		}
	}
	if t.SideEffect && (h.Confirm == nil || !h.Confirm(p)) {
		return Observation{}, fmt.Errorf("%w: %s", ErrNeedsUser, t.Name)
	}
	out, err := t.Handler(args) // only now does anything happen
	if err != nil {
		return Observation{}, fmt.Errorf("%s failed: %w", t.Name, err)
	}
	h.Ran = append(h.Ran, t.Name)
	return Observation{Tool: t.Name, Source: t.Source, ObservedAt: h.Now(), Content: out}, nil
}

// checkSchema enforces the subset of JSON Schema this example uses:
// required fields, string types, enums, and no extra fields.
// It is not a general validator.
func checkSchema(schema map[string]any, args map[string]any) error {
	props, _ := schema["properties"].(map[string]any)
	for _, r := range asStrings(schema["required"]) {
		if _, ok := args[r]; !ok {
			return fmt.Errorf("missing required field %q", r)
		}
	}
	for k, v := range args {
		spec, ok := props[k].(map[string]any)
		if !ok {
			return fmt.Errorf("unexpected field %q", k)
		}
		s, isString := v.(string)
		if spec["type"] == "string" && !isString {
			return fmt.Errorf("field %q must be a string", k)
		}
		if enum := asStrings(spec["enum"]); len(enum) > 0 && !contains(enum, s) {
			return fmt.Errorf("field %q must be one of %v", k, enum)
		}
	}
	return nil
}

func asStrings(v any) []string {
	switch x := v.(type) {
	case []string:
		return x
	case []any:
		out := make([]string, 0, len(x))
		for _, e := range x {
			if s, ok := e.(string); ok {
				out = append(out, s)
			}
		}
		return out
	}
	return nil
}

func contains(xs []string, s string) bool {
	for _, x := range xs {
		if x == s {
			return true
		}
	}
	return false
}
