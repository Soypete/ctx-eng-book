package gates

import (
	"strings"
	"testing"
	"time"
)

var (
	jan = time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)
	sep = time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC)
	v1  = Evidence{ID: "refund-policy", Fact: "refund-window", Version: "v1", AsOf: jan, Authoritative: true, Superseded: true}
	v3  = Evidence{ID: "refund-policy", Fact: "refund-window", Version: "v3", AsOf: sep, Authoritative: true}
)

func trace(retrieved []Evidence, unsupported ...string) Trace {
	return Trace{
		Question:    "What is the refund window?",
		Needs:       []string{"refund-window"},
		Current:     map[string]Evidence{"refund-window": v3},
		Retrieved:   retrieved,
		Unsupported: unsupported,
	}
}

func TestGate(t *testing.T) {
	cases := []struct {
		name string
		tr   Trace
		want string
	}{
		{"current copy retrieved", trace([]Evidence{v3}), "ok"},
		{"stale copy is a source/provenance failure", trace([]Evidence{v1}), "source: stale copy"},
		{"stale and current both retrieved still fails at source", trace([]Evidence{v3, v1}), "source: stale copy"},
		{"copy with no version is a provenance failure", trace([]Evidence{{ID: "wiki-copy", Fact: "refund-window"}}), "source: wiki-copy has no provenance"},
		{"fact never retrieved", trace(nil), "retrieval:"},
		{"evidence present but answer unsupported", trace([]Evidence{v3}, "refunds take 90 days"), "generation:"},
		{"no authoritative record anywhere", Trace{Needs: []string{"refund-window"}, Current: map[string]Evidence{}}, "source: no authoritative record"},
	}
	for _, c := range cases {
		if got := Gate(c.tr); !strings.HasPrefix(got, c.want) {
			t.Errorf("%s: got %q, want prefix %q", c.name, got, c.want)
		}
	}
}
