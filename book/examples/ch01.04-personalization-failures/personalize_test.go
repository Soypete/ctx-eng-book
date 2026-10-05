package personalize

import (
	"errors"
	"testing"
	"time"
)

var (
	now      = time.Date(2026, 10, 16, 9, 0, 0, 0, time.UTC)
	sessions = Sessions{"tok-ana": "ana", "tok-ben": "ben"}
	policy   = TaskPolicy{
		"code-review": {"language", "experience", "current_project"},
		"billing":     {"invoice_email"},
	}
	maxAge = map[string]time.Duration{
		"language":        30 * 24 * time.Hour,
		"current_project": time.Hour,
	}
)

func store() *Store {
	return &Store{Rows: []Preference{
		{Owner: "ana", Field: "language", Value: "Python", Origin: Inferred, UpdatedAt: now.Add(-48 * time.Hour)},
		{Owner: "ana", Field: "language", Value: "Go", Origin: Declared, UpdatedAt: now.Add(-72 * time.Hour)},
		{Owner: "ana", Field: "experience", Value: "senior", Origin: Declared, UpdatedAt: now.Add(-90 * 24 * time.Hour)},
		{Owner: "ana", Field: "current_project", Value: "billing-api", Origin: Inferred, UpdatedAt: now.Add(-3 * time.Hour)},
		{Owner: "ana", Field: "invoice_email", Value: "ana@example.com", Origin: Declared, UpdatedAt: now},
		{Owner: "ben", Field: "language", Value: "Rust", Origin: Declared, UpdatedAt: now},
	}}
}

func TestCrossUserReadIsImpossible(t *testing.T) {
	ctx, err := ForRequest(sessions, policy, store(), maxAge, "tok-ben", "code-review", now)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range ctx.Prefs {
		if p.Owner != "ben" {
			t.Fatalf("ben's context contains %s's %s", p.Owner, p.Field)
		}
	}
	// A zero Capability (what a caller gets if it tries to build one by hand)
	// reads nothing.
	if _, err := store().Read(Capability{}, now); !errors.Is(err, ErrExpired) {
		t.Fatalf("forged capability: err = %v", err)
	}
}

func TestUnknownSessionGetsNoContext(t *testing.T) {
	if _, err := ForRequest(sessions, policy, store(), maxAge, "tok-guess", "code-review", now); !errors.Is(err, ErrUnauthenticated) {
		t.Fatalf("err = %v, want ErrUnauthenticated", err)
	}
}

func TestTaskScopeLimitsFields(t *testing.T) {
	ctx, err := ForRequest(sessions, policy, store(), maxAge, "tok-ana", "code-review", now)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := ctx.Prefs["invoice_email"]; ok {
		t.Fatal("code review context includes the billing email")
	}
}

func TestDeclaredBeatsInferred(t *testing.T) {
	ctx, _ := ForRequest(sessions, policy, store(), maxAge, "tok-ana", "code-review", now)
	if got := ctx.Prefs["language"].Value; got != "Go" {
		t.Fatalf("language = %q; the user said Go", got)
	}
}

func TestStalePreferenceIsExcludedNotUsed(t *testing.T) {
	ctx, _ := ForRequest(sessions, policy, store(), maxAge, "tok-ana", "code-review", now)
	if _, ok := ctx.Prefs["current_project"]; ok {
		t.Fatal("3h-old project context used under a 1h limit")
	}
	found := false
	for _, e := range ctx.Excluded {
		found = found || (e.Field == "current_project" && e.Reason == "stale: refresh before use")
	}
	if !found {
		t.Fatalf("stale exclusion not recorded: %+v", ctx.Excluded)
	}
	// Fields with no max age (experience) are kept regardless of age.
	if ctx.Prefs["experience"].Value != "senior" {
		t.Fatal("experience dropped")
	}
}

func TestExpiredCapabilityReadsNothing(t *testing.T) {
	who, _ := sessions.Authenticate("tok-ana")
	c, _ := policy.Issue(who, "code-review", now, time.Minute)
	if _, err := store().Read(c, now.Add(time.Hour)); !errors.Is(err, ErrExpired) {
		t.Fatalf("err = %v, want ErrExpired", err)
	}
}
