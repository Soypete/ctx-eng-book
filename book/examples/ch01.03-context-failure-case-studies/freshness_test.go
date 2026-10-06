package freshness

import (
	"context"
	"errors"
	"testing"
	"time"
)

var now = time.Date(2026, 10, 14, 9, 0, 0, 0, time.UTC)

// The vector index synced yesterday and still says "negotiating".
var indexed = Fact{Subject: "deal:4411/stage", Value: "negotiating", Source: "crm",
	SourceVersion: 17, ObservedAt: now.Add(-24 * time.Hour)}

var policy = Policy{"deal-stage": time.Hour, "product-docs": 7 * 24 * time.Hour}

// The CRM closed the deal this morning.
func crm(_ context.Context, subject string) (Fact, error) {
	return Fact{Subject: subject, Value: "closed-won", Source: "crm",
		SourceVersion: 18, ObservedAt: now.Add(-5 * time.Minute)}, nil
}

func TestStaleDealStageIsRefetched(t *testing.T) {
	got, d, err := Resolve(context.Background(), "deal-stage", indexed, policy, crm, now)
	if err != nil {
		t.Fatal(err)
	}
	if got.Value != "closed-won" {
		t.Fatalf("assistant would act on %q; want the CRM's current stage", got.Value)
	}
	if d.Action != Refetched || d.SourceVersion != 18 || d.Age != 24*time.Hour {
		t.Fatalf("decision does not prove the refetch: %+v", d)
	}
}

func TestSameFactIsFreshEnoughForADifferentQuestion(t *testing.T) {
	called := false
	fetch := func(ctx context.Context, s string) (Fact, error) { called = true; return crm(ctx, s) }
	got, d, err := Resolve(context.Background(), "product-docs", indexed, policy, fetch, now)
	if err != nil {
		t.Fatal(err)
	}
	if called || got.Value != "negotiating" || d.Action != UsedCache {
		t.Fatalf("a 24h-old fact is within a 7-day policy; got %+v, fetched=%v", d, called)
	}
}

func TestFailedRefetchReturnsNoValue(t *testing.T) {
	down := func(context.Context, string) (Fact, error) { return Fact{}, errors.New("crm timeout") }
	got, d, err := Resolve(context.Background(), "deal-stage", indexed, policy, down, now)
	if !errors.Is(err, ErrStale) {
		t.Fatalf("err = %v, want ErrStale", err)
	}
	if got.Value != "" || d.Action != Rejected {
		t.Fatalf("stale value leaked: %q, %+v", got.Value, d)
	}
}

func TestUnknownQueryTypeFailsClosed(t *testing.T) {
	_, d, err := Resolve(context.Background(), "pricing", indexed, policy, crm, now)
	if !errors.Is(err, ErrNoPolicy) || d.Action != Rejected {
		t.Fatalf("err = %v, decision = %+v; want ErrNoPolicy and rejected", err, d)
	}
}
