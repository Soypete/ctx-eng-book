package main

import (
	"errors"
	"strings"
	"testing"
)

func okNotify(string, string) error { return nil }

func TestTimeoutThenRetryChargesOnce(t *testing.T) {
	l := NewLedger(map[string]int64{"acct_user": 500_000})
	l.TimeoutsBefore = 1
	w := Workflow{Ledger: l, Notify: okNotify, MaxAttempts: 3}
	r := w.Run(session(), Proposal{"Dana", 120_000}, "req_1")
	if r.Outcome != Completed || r.Attempts != 2 || r.TransactionID != "txn_001" {
		t.Fatalf("got %+v", r)
	}
	if l.Transfers != 1 || l.Balance["acct_user"] != 380_000 {
		t.Fatalf("charged %d times, balance %d", l.Transfers, l.Balance["acct_user"])
	}
}

// The same retry loop without a key is the double charge.
func TestRetryWithoutKeyChargesTwice(t *testing.T) {
	l := NewLedger(map[string]int64{"acct_user": 500_000})
	l.TimeoutsBefore = 1
	w := Workflow{Ledger: l, Notify: okNotify, MaxAttempts: 3}
	w.Run(session(), Proposal{"Dana", 120_000}, "")
	if l.Transfers != 2 {
		t.Fatalf("expected the keyless retry to move money twice, got %d", l.Transfers)
	}
}

func TestReplayedRequestReturnsSameTransaction(t *testing.T) {
	l := NewLedger(map[string]int64{"acct_user": 500_000})
	w := Workflow{Ledger: l, Notify: okNotify, MaxAttempts: 3}
	a := w.Run(session(), Proposal{"Dana", 50_000}, "req_9")
	b := w.Run(session(), Proposal{"Dana", 50_000}, "req_9")
	if a.TransactionID != b.TransactionID || l.Transfers != 1 {
		t.Fatalf("replay moved money again: %+v %+v transfers=%d", a, b, l.Transfers)
	}
}

func TestNotificationFailureKeepsTransfer(t *testing.T) {
	l := NewLedger(map[string]int64{"acct_user": 500_000})
	w := Workflow{Ledger: l, Notify: func(string, string) error { return errors.New("down") }, MaxAttempts: 3}
	r := w.Run(session(), Proposal{"Dana", 10_000}, "req_3")
	if r.Outcome != Completed || r.Notified || l.Transfers != 1 {
		t.Fatalf("got %+v transfers=%d", r, l.Transfers)
	}
}

func TestTerminalStates(t *testing.T) {
	cases := []struct {
		name    string
		p       Proposal
		down    bool
		balance int64
		want    Outcome
	}{
		{"ambiguous payee", Proposal{"Sam", 1_000}, false, 500_000, NeedsClarification},
		{"unknown payee", Proposal{"Mallory", 1_000}, false, 500_000, Denied},
		{"over limit", Proposal{"Dana", 900_000}, false, 2_000_000, Denied},
		{"insufficient funds", Proposal{"Dana", 100_000}, false, 5_000, FailedTerminal},
		{"ledger keeps timing out", Proposal{"Dana", 1_000}, true, 500_000, FailedRetryable},
		{"happy path", Proposal{"Dana", 1_000}, false, 500_000, Completed},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			l := NewLedger(map[string]int64{"acct_user": c.balance})
			l.Unreachable = c.down
			r := Workflow{Ledger: l, Notify: okNotify, MaxAttempts: 3}.Run(session(), c.p, "req_"+c.name)
			if r.Outcome != c.want {
				t.Fatalf("got %s (%s), want %s", r.Outcome, r.Reason, c.want)
			}
			if c.want != Completed {
				if l.Transfers != 0 {
					t.Fatal("money moved on a run that did not complete")
				}
			}
		})
	}
}

func TestSummaryHidesAccounts(t *testing.T) {
	l := NewLedger(map[string]int64{"acct_user": 500_000})
	r := Workflow{Ledger: l, Notify: okNotify, MaxAttempts: 3}.Run(session(), Proposal{"Dana", 1_000}, "k")
	for _, acct := range []string{"acct_user", "acct_dana"} {
		if strings.Contains(r.Summary(), acct) {
			t.Fatalf("summary leaks %s: %s", acct, r.Summary())
		}
	}
}
