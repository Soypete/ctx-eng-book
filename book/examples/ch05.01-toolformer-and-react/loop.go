package main

import "errors"

// Step is one model turn in a ReAct-style loop: either a tool proposal or a
// final answer. Reasoning text, if any, is the model's business, not the harness's.
type Step struct {
	Proposal *Proposal
	Answer   string
}

// Turn records what the model proposed and what the harness did about it.
type Turn struct {
	Proposal    Proposal
	Observation Observation
	Rejected    error
}

// ErrStepLimit ends a loop that never produces an answer.
var ErrStepLimit = errors.New("step limit reached without an answer")

// Loop alternates model proposals with harness decisions. A rejection is fed back
// as an observation so the model can try another route; it never executes.
func Loop(h *Harness, propose func(history []Turn) Step, maxSteps int) (string, []Turn, error) {
	var history []Turn
	for i := 0; i < maxSteps; i++ {
		s := propose(history)
		if s.Proposal == nil {
			return s.Answer, history, nil
		}
		obs, err := h.Handle(*s.Proposal)
		history = append(history, Turn{Proposal: *s.Proposal, Observation: obs, Rejected: err})
	}
	return "", history, ErrStepLimit
}
