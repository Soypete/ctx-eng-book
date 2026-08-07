package main

import (
	"context"
	"fmt"
)

// This toy trace demonstrates ReAct: reason, act, observe, repeat.
// ReAct structures tool use; it does not prove that inferred intent is correct.

type Tool func(args map[string]interface{}) (string, error)

func main() {
	// Simulated user request with ambiguous intent
	userRequest := "fix the API"

	// In a real harness, this would be the model's reasoning trace
	reactLoop := []struct {
		Thought     string
		Action      string
		Observation string
	}{
		{
			// In production, reading logs also requires an authorization check.
			Thought: "The user said 'fix the API' but didn't specify what kind of fix. " +
				"I need to interpret what they actually want. Possibilities: " +
				"1) Fix a bug, 2) Update endpoint, 3) Improve performance. " +
				"I should ask for clarification or check recent error logs first.",
			Action:      "get_recent_errors",
			Observation: "Found 3 auth-related errors in the last hour",
		},
		{
			Thought: "The errors are in the auth handler. The user likely means " +
				"'find the bug in the auth handler' rather than a general API fix.",
			Action:      "search_code(pattern=auth, location=api)",
			Observation: "Found 2 files: api/auth.go, api/middleware.go",
		},
	}

	// A safer harness would require clarification or a read capability before
	// executing the first action. A visible reasoning trace is not enforcement.

	fmt.Println("User request:", userRequest)
	fmt.Println("\n--- ReAct Reasoning Trace ---")
	for i, step := range reactLoop {
		fmt.Printf("\nStep %d:\n", i+1)
		fmt.Printf("  Thought: %s\n", step.Thought)
		fmt.Printf("  Action:  %s\n", step.Action)
		fmt.Printf("  Observe: %s\n", step.Observation)
	}

	// Key insight: ReAct structures iteration; the harness still owns policy,
	// validation, budgets, and termination.
	_ = context.Background() // suppress unused warning
}
