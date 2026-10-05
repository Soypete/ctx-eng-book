// Package contextcompile is the runnable companion to the post
// "Context Engineering Is Not Prompt Engineering" (book chapter 0).
//
// Compile admits already-scored search candidates into a model's working set.
// It is a post-filter: it keeps unauthorized text out of model input, but it
// cannot recover eligible documents that a top-k search never returned.
// Production systems should also push access scope into the retrieval query.
package contextcompile

import (
	"sort"
	"time"
)

// Doc is one scored search candidate.
type Doc struct {
	ID, Version, Text string
	AsOf              time.Time // source freshness; enforcing a freshness policy is the caller's job
	Score             float64
	Tokens            int
}

// Manifest records the trusted request metadata and every admission decision,
// so a selection can be inspected and, with retained artifacts, replayed.
type Manifest struct {
	SourceVersion, PolicyVersion, RankerVersion string
	Actor, Purpose                              string // from authenticated state, never from the model
	At                                          time.Time
	Budget                                      int
	Included, Excluded                          []string
}

// Allowed decides whether actor may use d for purpose. Back it with your
// existing authorization infrastructure, not ownership alone.
type Allowed func(actor, purpose string, d Doc) bool

// Compile filters, ranks, and budgets candidates, recording why each one was
// included or excluded.
func Compile(m Manifest, allowed Allowed, cands []Doc) ([]Doc, Manifest) {
	m.Included, m.Excluded = nil, nil

	var eligible []Doc
	for _, d := range cands {
		if !allowed(m.Actor, m.Purpose, d) {
			m.Excluded = append(m.Excluded, d.ID+":unauthorized")
			continue
		}
		eligible = append(eligible, d)
	}

	// Unique IDs break score ties so the order never depends on search return order.
	sort.SliceStable(eligible, func(i, j int) bool {
		if eligible[i].Score == eligible[j].Score {
			return eligible[i].ID < eligible[j].ID
		}
		return eligible[i].Score > eligible[j].Score
	})

	var out []Doc
	remaining := m.Budget
	for _, d := range eligible {
		if d.Tokens > remaining {
			m.Excluded = append(m.Excluded, d.ID+":budget")
			continue
		}
		remaining -= d.Tokens
		out = append(out, d)
		m.Included = append(m.Included, d.ID)
	}
	return out, m
}
