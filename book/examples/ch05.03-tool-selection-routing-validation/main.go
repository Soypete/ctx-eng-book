package main

import (
	"errors"
	"fmt"
)

// All names, accounts, and amounts are invented.
func session() Session {
	return Session{
		UserID:  "u_42",
		Account: "acct_user",
		Payees: map[string]string{
			"Sam Rivera (landlord)": "acct_sam_r",
			"Sam Chen":              "acct_sam_c",
			"Dana Okafor":           "acct_dana",
		},
		LimitCents: 200_000,
	}
}

func main() {
	l := NewLedger(map[string]int64{"acct_user": 500_000})
	l.TimeoutsBefore = 1 // the first attempt commits and then times out
	w := Workflow{Ledger: l, Notify: func(string, string) error { return errors.New("sms provider down") }, MaxAttempts: 3}

	fmt.Println(w.Run(session(), Proposal{RecipientName: "Sam", AmountCents: 120_000}, "req_1").Summary())
	fmt.Println(w.Run(session(), Proposal{RecipientName: "Dana", AmountCents: 120_000}, "req_2").Summary())
	fmt.Println("money movements:", l.Transfers, "user balance:", l.Balance["acct_user"])
}
