// Package main shows a tool contract that is wider than its schema, and an
// eval that tests the five ways a tool call can go wrong.
package main

import (
	"fmt"
	"strings"
)

// Contract is everything the harness knows about one tool. The schema is one
// field of it. The model sees Description(); the harness enforces the rest.
type Contract struct {
	Name         string
	Purpose      string
	UseWhen      []string // selection context: when this tool is the right one
	DoNotUseWhen []string // selection context: near-misses that belong elsewhere
	Effects      string   // what changes in the world if it runs
	Required     []string // schema: required argument names
	Optional     []string // schema: other allowed argument names
	Enums        map[string][]string
	RequiredRole string // authorization: who may cause the effect
}

// Description is the model-visible text. "When not to use" is what separates
// two tools that sound alike.
func (c Contract) Description() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s: %s Effects: %s.", c.Name, c.Purpose, c.Effects)
	if len(c.UseWhen) > 0 {
		fmt.Fprintf(&b, " Use when: %s.", strings.Join(c.UseWhen, "; "))
	}
	if len(c.DoNotUseWhen) > 0 {
		fmt.Fprintf(&b, " Do not use when: %s.", strings.Join(c.DoNotUseWhen, "; "))
	}
	return b.String()
}

// createCustomer and getCustomer are invented for the example.
var createCustomer = Contract{
	Name:         "create_customer",
	Purpose:      "Create a new customer record.",
	UseWhen:      []string{"a person who is not yet a customer is signing up", "sales is onboarding a new account"},
	DoNotUseWhen: []string{"the person already has a record (use get_customer)", "the request is to change an existing customer's tier"},
	Effects:      "writes one row to customers; sends no email",
	Required:     []string{"email", "name"},
	Optional:     []string{"tier"},
	Enums:        map[string][]string{"tier": {"free", "pro", "enterprise"}},
	RequiredRole: "support",
}

var getCustomer = Contract{
	Name:         "get_customer",
	Purpose:      "Look up an existing customer by email.",
	UseWhen:      []string{"the user asks about an existing account"},
	DoNotUseWhen: []string{"the person is not a customer yet"},
	Effects:      "none; read-only",
	Required:     []string{"email"},
	RequiredRole: "support",
}
