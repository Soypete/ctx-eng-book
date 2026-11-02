// Package main shows a multi-step tool workflow with software-defined terminal
// states, an idempotency key on the side effect, and a notification step
// that cannot undo a committed transfer.
package main

import (
	"errors"
	"fmt"
	"strings"
)

// Outcome is the closed set of ways a transfer run can end.
type Outcome string

const (
	Completed          Outcome = "completed"
	NeedsClarification Outcome = "needs_clarification"
	Denied             Outcome = "denied"
	FailedRetryable    Outcome = "failed_retryable"
	FailedTerminal     Outcome = "failed_terminal"
)

// Proposal is what the model extracted from the user's message. It is a
// request for the workflow to consider, not an instruction to move money.
type Proposal struct {
	RecipientName string
	AmountCents   int64
}

// Session is trusted state from authentication, never from the model.
type Session struct {
	UserID     string
	Account    string
	Payees     map[string]string // payee display name -> account
	LimitCents int64
}

// Result is the typed record of the run. Only Summary goes back to the model.
type Result struct {
	Outcome       Outcome
	TransactionID string
	Attempts      int
	Notified      bool
	Reason        string
}

// Summary is what the model may see: outcome and IDs, no account numbers.
func (r Result) Summary() string {
	return fmt.Sprintf("outcome=%s txn=%s notified=%t reason=%q", r.Outcome, r.TransactionID, r.Notified, r.Reason)
}

// Workflow wires the steps to their services and limits.
type Workflow struct {
	Ledger      *Ledger
	Notify      func(userID, txn string) error
	MaxAttempts int
}

// Run routes one proposal through resolve, authorize, transfer, and notify.
// key must be derived from the user's request (for example a request ID),
// so a retried request reuses it.
func (w Workflow) Run(s Session, p Proposal, key string) Result {
	to, r, ok := resolvePayee(s, p.RecipientName)
	if !ok {
		return r
	}
	if p.AmountCents <= 0 || p.AmountCents > s.LimitCents {
		return Result{Outcome: Denied, Reason: fmt.Sprintf("amount %d outside 1..%d", p.AmountCents, s.LimitCents)}
	}

	var txn string
	var err error
	attempts := 0
	for attempts < w.MaxAttempts {
		attempts++
		txn, err = w.Ledger.Transfer(key, s.Account, to, p.AmountCents) // same key every attempt
		if !errors.Is(err, ErrTimeout) {
			break
		}
	}
	switch {
	case errors.Is(err, ErrTimeout):
		// Unknown whether it committed. A later run with the same key can't double-charge.
		return Result{Outcome: FailedRetryable, Attempts: attempts, Reason: "ledger timed out; retry with the same key"}
	case err != nil:
		return Result{Outcome: FailedTerminal, Attempts: attempts, Reason: err.Error()}
	}

	res := Result{Outcome: Completed, TransactionID: txn, Attempts: attempts}
	if nerr := w.Notify(s.UserID, txn); nerr != nil {
		res.Reason = "transfer committed; notification failed, retry it on its own" // never roll back or redo the transfer
		return res
	}
	res.Notified = true
	return res
}

func resolvePayee(s Session, name string) (string, Result, bool) {
	var matches []string
	for display, acct := range s.Payees {
		if strings.Contains(strings.ToLower(display), strings.ToLower(strings.TrimSpace(name))) {
			matches = append(matches, acct)
		}
	}
	switch len(matches) {
	case 1:
		return matches[0], Result{}, true
	case 0:
		return "", Result{Outcome: Denied, Reason: "recipient is not a saved payee"}, false
	default:
		return "", Result{Outcome: NeedsClarification, Reason: fmt.Sprintf("%d payees match %q", len(matches), name)}, false
	}
}
