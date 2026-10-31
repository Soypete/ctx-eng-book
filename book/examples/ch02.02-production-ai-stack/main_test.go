package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

type modelFunc func(context.Context, Request, []Document) (string, error)

func (f modelFunc) Select(c context.Context, r Request, d []Document) (string, error) {
	return f(c, r, d)
}

func TestContractsSurviveModelSwap(t *testing.T) {
	r, docs, now, p := fixture()
	// The unchanged first-result model selects the wrong release without admission.
	raw, _ := (firstModel{}).Select(context.Background(), r, docs)
	if raw != "guide-v1" {
		t.Fatal("fixture no longer reproduces failure")
	}
	for _, m := range []Model{firstModel{}, lastModel{}} {
		out, err := Lookup(context.Background(), r, docs, now, p, m)
		if err != nil || out.Text != docs[1].Text {
			t.Fatalf("%T: %+v %v", m, out, err)
		}
		if len(out.Decisions) != 3 || out.Decisions[0].Outcome != "rejected" || out.Decisions[2].SourceVersion != "rev-b" {
			t.Fatalf("missing lineage or rejection: %+v", out.Decisions)
		}
	}
}

func TestInvalidEvidenceNeverReachesModel(t *testing.T) {
	for _, tc := range []struct {
		name   string
		change func(*Document)
	}{
		{"wrong release", func(d *Document) { d.Release = "v1" }},
		{"wrong product", func(d *Document) { d.Product = "search" }},
		{"stale", func(d *Document) { d.AsOf = d.AsOf.Add(-2 * time.Hour) }},
		{"future", func(d *Document) { d.AsOf = d.AsOf.Add(time.Second) }},
		{"missing time", func(d *Document) { d.AsOf = time.Time{} }},
		{"missing provenance", func(d *Document) { d.Version = "" }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r, docs, now, p := fixture()
			d := docs[1]
			tc.change(&d)
			called := false
			m := modelFunc(func(context.Context, Request, []Document) (string, error) { called = true; return d.ID, nil })
			out, err := Lookup(context.Background(), r, []Document{d}, now, p, m)
			if err == nil || called || out.Text != "" || len(out.Decisions) != 1 || out.Decisions[0].Outcome != "rejected" {
				t.Fatalf("failed boundary: %+v %v called=%v", out, err, called)
			}
		})
	}
}

func TestRejectedCitationCannotReturn(t *testing.T) {
	r, docs, now, p := fixture()
	m := modelFunc(func(_ context.Context, _ Request, admitted []Document) (string, error) {
		if len(admitted) != 1 || admitted[0].ID != "guide-v2" {
			t.Fatal("wrong model context")
		}
		return "guide-v1", nil
	})
	out, err := Lookup(context.Background(), r, docs, now, p, m)
	if err == nil || out.Text != "" || out.Decisions[len(out.Decisions)-1].Stage != "validation" {
		t.Fatalf("citation escaped: %+v %v", out, err)
	}
}

func TestModelErrorAndMutation(t *testing.T) {
	r, docs, now, p := fixture()
	fail := modelFunc(func(context.Context, Request, []Document) (string, error) { return "", errors.New("provider failed") })
	out, err := Lookup(context.Background(), r, docs, now, p, fail)
	if err == nil || out.Text != "" || out.Decisions[len(out.Decisions)-1].Stage != "model" {
		t.Fatalf("lost failure: %+v %v", out, err)
	}
	mutate := modelFunc(func(_ context.Context, _ Request, d []Document) (string, error) {
		d[0].Text = "invented steps"
		return d[0].ID, nil
	})
	out, err = Lookup(context.Background(), r, docs, now, p, mutate)
	if err != nil || out.Text != docs[1].Text {
		t.Fatalf("adapter mutated evidence: %+v %v", out, err)
	}
}
