package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Policy struct {
	ID, Version, Scope string
	ConfirmedAt        time.Time
	General            string
	Exceptions         map[string]string
}

type Manifest struct {
	SourceID, SourceVersion, Scope, Method string
	ConfirmedAt                            time.Time
	Deferred                               []string
}

type Summary struct {
	Text     string
	Manifest Manifest
}

type Decision struct {
	SourceID, SourceVersion, Method, Action, Reason string
}

// Compact deliberately drops exceptions to model a lossy summary, not an LLM.
func Compact(p Policy) Summary {
	return Summary{p.General, Manifest{p.ID, p.Version, p.Scope, "general-only/v1", p.ConfirmedAt, []string{"exceptions"}}}
}

// Fetch must enforce current caller authorization and source freshness.
type Fetch func(Manifest) (Policy, error)

func Resolve(s Summary, plan string, fetch Fetch) (string, Decision) {
	d := Decision{SourceID: s.Manifest.SourceID, SourceVersion: s.Manifest.SourceVersion, Method: s.Manifest.Method}
	// Eligibility depends on exceptions; never promote the summary to authority.
	p, err := fetch(s.Manifest)
	if err != nil {
		d.Action, d.Reason = "withheld", err.Error()
		return "", d
	}
	if p.ID != s.Manifest.SourceID || p.Version != s.Manifest.SourceVersion {
		d.Action, d.Reason = "withheld", "source identity or version changed; rebuild summary"
		return "", d
	}
	d.Action, d.Reason = "rehydrated", "eligibility requires original exceptions"
	answer := p.General
	if exception, ok := p.Exceptions[plan]; ok {
		answer = exception
	}
	return answer, d
}

// Reader is an in-memory stand-in for a governed source adapter. Its inputs
// come from trusted application policy, never from a generated summary.
func Reader(store map[string]Policy, callerScope string, now time.Time, maxAge time.Duration) Fetch {
	return func(m Manifest) (Policy, error) {
		p, ok := store[m.SourceID]
		if !ok {
			return Policy{}, fmt.Errorf("source unavailable")
		}
		if p.Scope != callerScope {
			return Policy{}, fmt.Errorf("source access denied")
		}
		age := now.Sub(p.ConfirmedAt)
		if maxAge <= 0 || age < 0 || age > maxAge {
			return Policy{}, fmt.Errorf("source freshness rejected")
		}
		return p, nil
	}
}

func fixture(now time.Time) Policy {
	return Policy{ID: "refund-policy", Version: "v7", Scope: "support", ConfirmedAt: now,
		General:    "Refunds are available within 30 days.",
		Exceptions: map[string]string{"annual": "Annual plans are nonrefundable after activation."}}
}

func main() {
	now := time.Date(2026, 11, 16, 12, 0, 0, 0, time.UTC)
	p := fixture(now)
	s := Compact(p)
	store := map[string]Policy{p.ID: p}
	fetch := Reader(store, "support", now, time.Hour)
	recovered, decision := Resolve(s, "annual", fetch)
	delete(store, p.ID)
	missing, blocked := Resolve(s, "annual", fetch)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(struct {
		Summary             Summary
		Recovered           string
		Recovery            Decision
		MissingSourceAnswer string
		MissingSource       Decision
	}{s, recovered, decision, missing, blocked}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
