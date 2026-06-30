package main

import (
	"context"
	"fmt"
)

// ReAct demonstrates speech act interpretation via structured reasoning.
// The model doesn't just act — it first REASONS about what the user meant.
// This example shows the ReAct loop: Think → Act → Observe → Repeat.

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
			// BEAT 1: Speech act interpretation — reason BEFORE acting
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

	// This is what happens WITHOUT ReAct (pure inference):
	// The model would guess and call a random tool, likely wrong.
	// With ReAct, the model explicitly reasons about intent first.

	fmt.Println("User request:", userRequest)
	fmt.Println("\n--- ReAct Reasoning Trace ---")
	for i, step := range reactLoop {
		fmt.Printf("\nStep %d:\n", i+1)
		fmt.Printf("  Thought: %s\n", step.Thought)
		fmt.Printf("  Action:  %s\n", step.Action)
		fmt.Printf("  Observe: %s\n", step.Observation)
	}

	// Key insight: ReAct adds STRUCTURE to force intent extraction.
	// Without it, the model infers incorrectly.
	_ = context.Background() // suppress unused warning
}
