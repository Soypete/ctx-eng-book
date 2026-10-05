// Package gates is the runnable companion to the post
// "Before You Blame the Model, Check Three Gates" (book module 1.1).
//
// Gate classifies a wrong answer by the first of three gates it failed:
// source, retrieval, or generation. It runs on labeled traces (eval cases, or
// production failures you have triaged and labeled), because it needs to know
// which facts a correct answer requires. Live traffic gets the trace recorded;
// diagnosis happens once someone labels it.
package gates

import (
	"fmt"
	"time"
)

// Evidence is one record as it reached the working context, with the
// provenance the source attached to it.
type Evidence struct {
	ID, Fact, Version string
	AsOf              time.Time
	Authoritative     bool // the source vouches for this copy
	Superseded        bool // the source has marked a newer version
}

// Trace is what you keep for every answer so a wrong one can be diagnosed later.
type Trace struct {
	Question    string
	Needs       []string            // facts a correct answer requires (from the label)
	Current     map[string]Evidence // the source's authoritative, current record per fact
	Retrieved   []Evidence          // what was admitted into the working context
	Unsupported []string            // answer claims a checker found no evidence for
}

// Gate names the first gate that failed. Each gate has a different owner and fix.
func Gate(t Trace) string {
	for _, fact := range t.Needs {
		cur, ok := t.Current[fact]
		if !ok {
			return fmt.Sprintf("source: no authoritative record for %q; ingest it or ask the user", fact)
		}
		found := false
		for _, r := range t.Retrieved {
			if r.Fact != fact {
				continue
			}
			// A stale copy is a provenance failure at the source: if the source
			// does not version and supersede its records, an old copy looks
			// exactly as legitimate as the current one.
			if r.Version == "" || r.AsOf.IsZero() {
				return fmt.Sprintf("source: %s has no provenance (version, as-of)", r.ID)
			}
			if !r.Authoritative || r.Superseded || r.Version != cur.Version {
				return fmt.Sprintf("source: stale copy %s@%s competes with %s@%s; mark authority and supersession at the source",
					r.ID, r.Version, cur.ID, cur.Version)
			}
			found = true
		}
		if !found {
			return fmt.Sprintf("retrieval: %q exists at the source but never reached the context; fix the query, index, ranking, or scope", fact)
		}
	}
	if len(t.Unsupported) > 0 {
		return "generation: the evidence was there; change the model, instructions, or validation"
	}
	return "ok"
}
