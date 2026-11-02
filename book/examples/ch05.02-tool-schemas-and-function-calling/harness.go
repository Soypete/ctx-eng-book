package main

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

// Call is a model's proposed tool call after JSON decoding.
type Call struct {
	Tool string
	Args map[string]any
}

// Principal is who the call acts for, resolved from the session, never from
// model output.
type Principal struct {
	ID    string
	Roles []string
}

// Each stage has its own error so tests and logs can say where a call stopped.
var (
	ErrSyntax       = errors.New("syntax")
	ErrSemantic     = errors.New("semantic")
	ErrUnauthorized = errors.New("unauthorized")
)

// Harness holds the contracts and the state the semantic checks need.
type Harness struct {
	Contracts map[string]Contract
	Customers map[string]string // email -> customer ID
	Executed  []Call
}

// Check runs the stages in order. A call that passes every stage executes.
func (h *Harness) Check(p Principal, c Call) error {
	k, ok := h.Contracts[c.Tool]
	if !ok {
		return fmt.Errorf("%w: unknown tool %q", ErrSyntax, c.Tool)
	}
	if err := checkShape(k, c.Args); err != nil {
		return fmt.Errorf("%w: %v", ErrSyntax, err)
	}
	if err := h.checkMeaning(k, c.Args); err != nil { // duplicates, unparseable email
		return fmt.Errorf("%w: %v", ErrSemantic, err)
	}
	if !hasRole(p, k.RequiredRole) { // p comes from the session, never from Args
		return fmt.Errorf("%w: %s lacks role %q for %s", ErrUnauthorized, p.ID, k.RequiredRole, k.Name)
	}
	if t, _ := c.Args["tier"].(string); t == "enterprise" && !hasRole(p, "sales") {
		return fmt.Errorf("%w: enterprise tier requires sales", ErrUnauthorized)
	}
	h.Executed = append(h.Executed, c) // stands in for running the tool
	return nil
}

// checkShape is what the schema can express: names, types, enums.
func checkShape(k Contract, args map[string]any) error {
	allowed := map[string]bool{}
	for _, f := range append(append([]string{}, k.Required...), k.Optional...) {
		allowed[f] = true
	}
	for _, f := range k.Required {
		if _, ok := args[f]; !ok {
			return fmt.Errorf("missing %q", f)
		}
	}
	for f, v := range args {
		if !allowed[f] {
			return fmt.Errorf("unexpected field %q", f)
		}
		s, ok := v.(string)
		if !ok {
			return fmt.Errorf("%q must be a string", f)
		}
		if enum, ok := k.Enums[f]; ok && !contains(enum, s) {
			return fmt.Errorf("%q must be one of %v", f, enum)
		}
	}
	return nil
}

// checkMeaning is what the schema can't express: a parseable address, a
// nonblank name, and a customer who doesn't already exist.
func (h *Harness) checkMeaning(k Contract, args map[string]any) error {
	email, _ := args["email"].(string)
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("email %q is not an address", email)
	}
	if k.Name != "create_customer" {
		return nil
	}
	if name, _ := args["name"].(string); strings.TrimSpace(name) == "" {
		return errors.New("name is blank")
	}
	if id, exists := h.Customers[strings.ToLower(email)]; exists {
		return fmt.Errorf("customer %s already exists for this email", id)
	}
	return nil
}

func hasRole(p Principal, role string) bool { return contains(p.Roles, role) }

func contains(xs []string, s string) bool {
	for _, x := range xs {
		if x == s {
			return true
		}
	}
	return false
}
