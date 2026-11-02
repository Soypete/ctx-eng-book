package main

import "fmt"

func newHarness() *Harness {
	return &Harness{
		Contracts: map[string]Contract{createCustomer.Name: createCustomer, getCustomer.Name: getCustomer},
		Customers: map[string]string{"ana@example.com": "cus_001"},
	}
}

var (
	agent = Principal{ID: "support-agent-7", Roles: []string{"support"}}
	guest = Principal{ID: "anonymous", Roles: nil}
)

// suite is the eval: at least one case per class. Requests and people are invented.
func suite() []Case {
	call := func(tool string, kv ...string) *Call {
		args := map[string]any{}
		for i := 0; i+1 < len(kv); i += 2 {
			args[kv[i]] = kv[i+1]
		}
		return &Call{Tool: tool, Args: args}
	}
	return []Case{
		{Name: "new signup", Class: Selection, Request: "Sign up Ben Ortiz, ben@example.com, on the free plan", WantTool: "create_customer"},
		{Name: "existing account", Class: Selection, Request: "What plan is ana@example.com on?", WantTool: "get_customer"},
		{Name: "tier change, no tool for it", Class: Abstention, Request: "Move ana@example.com to the pro plan", WantTool: ""},
		{Name: "policy question", Class: Abstention, Request: "What's the difference between pro and enterprise?", WantTool: ""},
		{Name: "missing name", Class: Malformed, Principal: agent, Call: call("create_customer", "email", "ben@example.com"), WantErr: ErrSyntax},
		{Name: "unknown tier", Class: Malformed, Principal: agent, Call: call("create_customer", "email", "ben@example.com", "name", "Ben", "tier", "platinum"), WantErr: ErrSyntax},
		{Name: "duplicate customer", Class: Semantic, Principal: agent, Call: call("create_customer", "email", "ANA@example.com", "name", "Ana"), WantErr: ErrSemantic},
		{Name: "not an address", Class: Semantic, Principal: agent, Call: call("create_customer", "email", "ben at example", "name", "Ben"), WantErr: ErrSemantic},
		{Name: "anonymous caller", Class: Unauthorized, Principal: guest, Call: call("create_customer", "email", "ben@example.com", "name", "Ben"), WantErr: ErrUnauthorized},
		{Name: "enterprise without sales", Class: Unauthorized, Principal: agent, Call: call("create_customer", "email", "ben@example.com", "name", "Ben", "tier", "enterprise"), WantErr: ErrUnauthorized},
		{Name: "valid create", Class: Control, Principal: agent, Call: call("create_customer", "email", "ben@example.com", "name", "Ben", "tier", "free"), WantErr: nil},
	}
}

// recorded replays what a model chose for each request in a past run. Swap in
// a real model call to evaluate your own descriptions.
func recorded(answers map[string]string) Selector {
	return func(request string, _ []string) string { return answers[request] }
}

var goodRun = map[string]string{
	"Sign up Ben Ortiz, ben@example.com, on the free plan": "create_customer",
	"What plan is ana@example.com on?":                     "get_customer",
}

func main() {
	fmt.Println("model sees:", createCustomer.Description())
	for _, r := range Evaluate(suite(), recorded(goodRun), newHarness()) {
		mark := "PASS"
		if !r.Passed {
			mark = "FAIL"
		}
		fmt.Printf("%s %-22s %-20s %s\n", mark, r.Case.Class, r.Case.Name, r.Detail)
	}
}
