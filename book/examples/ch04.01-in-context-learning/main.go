package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Example struct{ Input, Output string }
type Invocation struct {
	ContractVersion                            string
	Rule                                       string
	ExamplesVersion                            string
	Examples                                   []Example
	SourceID, SourceVersion, SourceConfirmedAt string
	Input                                      string
}
type Record struct{ SnapshotID, Action, Reason string }

// Freeze retains the bytes as well as their identity: a hash cannot replay input.
func Freeze(in Invocation) ([]byte, string, error) {
	if in.ContractVersion == "" || in.Rule == "" || in.ExamplesVersion == "" || in.SourceID == "" || in.SourceVersion == "" || in.SourceConfirmedAt == "" {
		return nil, "", fmt.Errorf("missing contract or provenance")
	}
	b, err := json.Marshal(in)
	if err != nil {
		return nil, "", err
	}
	sum := sha256.Sum256(b)
	return b, hex.EncodeToString(sum[:]), nil
}

// Check compares a candidate with a reviewed fixture, not with another model's guess.
func Check(snapshotID, candidate, expected string) Record {
	d := Record{SnapshotID: snapshotID, Action: "rejected"}
	var out struct {
		Label string `json:"label"`
	}
	if err := json.Unmarshal([]byte(candidate), &out); err != nil {
		d.Reason = "invalid JSON"
		return d
	}
	// Valid JSON alone does not establish the intended meaning of a label.
	if out.Label != "buyer" && out.Label != "unknown" {
		d.Reason = "label outside contract"
		return d
	}
	if out.Label != expected {
		d.Reason = "label disagrees with reviewed fixture"
		return d
	}
	d.Action, d.Reason = "accepted", "label matches reviewed fixture"
	return d
}

func fixture() Invocation {
	return Invocation{
		ContractVersion: "buyer-v1",
		Rule:            "buyer requires an explicit purchase commitment; otherwise unknown",
		ExamplesVersion: "examples-v1",
		Examples:        []Example{{"We will purchase the annual plan.", "buyer"}, {"We are evaluating the annual plan.", "unknown"}, {"Please send pricing.", "unknown"}},
		SourceID:        "synthetic-ticket", SourceVersion: "v1", SourceConfirmedAt: "2026-11-01T00:00:00Z",
		Input: "Please send pricing.",
	}
}

func main() {
	b, id, err := Freeze(fixture())
	if err != nil {
		panic(err)
	}
	fmt.Println("Retained invocation:", string(b))
	// Candidates are scripted so the semantic failure is reproducible without an API.
	for _, candidate := range []string{`{"label":"buyer"}`, `{"label":"unknown"}`} {
		d, _ := json.Marshal(Check(id, candidate, "unknown"))
		fmt.Println(string(d))
	}
}
