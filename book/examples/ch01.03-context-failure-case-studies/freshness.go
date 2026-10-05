// Package freshness is the runnable companion to the post
// "Every Context Failure Has an Address" (book module 1.3).
//
// It implements one of the post's five boundaries, freshness, and shows the
// post's second claim: a fix is only useful if it leaves evidence. Resolve
// decides whether a cached fact is fresh enough for the question being asked,
// refetches it from the authoritative source when it is not, and returns a
// Decision record that says which path ran and why.
package freshness

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	// ErrNoPolicy means nobody decided how old this kind of fact may be.
	// Unknown query types fail closed instead of trusting the cache.
	ErrNoPolicy = errors.New("no freshness policy for query type")
	// ErrStale means the cached fact is too old and the source could not
	// confirm a current value. The caller gets no value rather than an old one.
	ErrStale = errors.New("fact is stale and source refetch failed")
)

// Fact is one retrieved value plus the metadata needed to judge whether it is
// still current.
type Fact struct {
	Subject       string    // what the value describes, e.g. "deal:4411/stage"
	Value         string    // e.g. "negotiating"
	Source        string    // the authoritative system, e.g. "crm"
	SourceVersion int64     // the source's version for this record
	ObservedAt    time.Time // when the source last vouched for the value
}

// Policy is the maximum acceptable age per query type. It belongs to the
// application, not the model: "deal-stage" questions might tolerate an hour,
// "product-docs" questions a week.
type Policy map[string]time.Duration

// Action records which path Resolve took.
type Action string

const (
	UsedCache Action = "used-cache"
	Refetched Action = "refetched"
	Rejected  Action = "rejected"
)

// Decision is the evidence row for one freshness check. Log it next to the
// request so you can later prove whether the boundary held.
type Decision struct {
	QueryType     string
	Subject       string
	Source        string
	SourceVersion int64
	ObservedAt    time.Time
	CheckedAt     time.Time
	Age           time.Duration
	MaxAge        time.Duration
	Action        Action
	Reason        string
}

// Fetcher reads the current value from the authoritative source.
type Fetcher func(ctx context.Context, subject string) (Fact, error)

// Resolve returns a fact that is fresh enough for queryType, together with the
// Decision that explains how it got there.
func Resolve(ctx context.Context, queryType string, cached Fact, pol Policy, fetch Fetcher, now time.Time) (Fact, Decision, error) {
	d := Decision{QueryType: queryType, Subject: cached.Subject, Source: cached.Source,
		SourceVersion: cached.SourceVersion, ObservedAt: cached.ObservedAt, CheckedAt: now}
	maxAge, ok := pol[queryType]
	if !ok {
		d.Action, d.Reason = Rejected, "no policy"
		return Fact{}, d, fmt.Errorf("%w: %q", ErrNoPolicy, queryType)
	}
	d.MaxAge, d.Age = maxAge, now.Sub(cached.ObservedAt)
	if d.Age <= maxAge {
		d.Action, d.Reason = UsedCache, "within max age"
		return cached, d, nil
	}
	fresh, err := fetch(ctx, cached.Subject)
	if err != nil { // don't fall back to the stale value: an old answer looks just like a current one
		d.Action, d.Reason = Rejected, err.Error()
		return Fact{}, d, fmt.Errorf("%w: %s is %s old (max %s): %v", ErrStale, cached.Subject, d.Age, maxAge, err)
	}
	d.Action, d.Reason = Refetched, fmt.Sprintf("cached copy %s old exceeds %s", d.Age, maxAge)
	d.SourceVersion, d.ObservedAt = fresh.SourceVersion, fresh.ObservedAt
	return fresh, d, nil
}
