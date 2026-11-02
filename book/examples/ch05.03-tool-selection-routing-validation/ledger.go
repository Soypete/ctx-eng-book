package main

import (
	"errors"
	"fmt"
)

// Errors a payment service can return. A timeout says nothing about whether
// the transfer committed.
var (
	ErrTimeout           = errors.New("timeout waiting for ledger")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

// Ledger is a fake payment service. It honors idempotency keys: the same key
// returns the original transaction instead of moving money again.
type Ledger struct {
	Balance        map[string]int64  // account -> cents
	byKey          map[string]string // idempotency key -> transaction ID
	Transfers      int               // money movements actually committed
	TimeoutsBefore int               // commit, then time out, this many times
	Unreachable    bool              // time out without committing anything
}

// NewLedger returns a ledger with the given balances.
func NewLedger(balances map[string]int64) *Ledger {
	return &Ledger{Balance: balances, byKey: map[string]string{}}
}

// Transfer moves money once per key. With an empty key, every call moves money.
func (l *Ledger) Transfer(key, from, to string, cents int64) (string, error) {
	if l.Unreachable {
		return "", ErrTimeout
	}
	if id, seen := l.byKey[key]; key != "" && seen {
		return id, nil // replay: same answer, no second movement
	}
	if l.Balance[from] < cents {
		return "", ErrInsufficientFunds
	}
	l.Balance[from] -= cents
	l.Balance[to] += cents
	l.Transfers++
	id := fmt.Sprintf("txn_%03d", l.Transfers)
	if key != "" {
		l.byKey[key] = id
	}
	if l.TimeoutsBefore > 0 {
		l.TimeoutsBefore--
		return "", ErrTimeout // committed, but the caller never hears about it
	}
	return id, nil
}
