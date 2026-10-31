package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestSemanticFailureDespiteValidJSON(t *testing.T) {
	_, id, err := Freeze(fixture())
	if err != nil {
		t.Fatal(err)
	}
	wrong := `{"label":"buyer"}`
	if !json.Valid([]byte(wrong)) {
		t.Fatal("fixture must be valid JSON")
	}
	d := Check(id, wrong, "unknown")
	if d.Action != "rejected" || d.Reason != "label disagrees with reviewed fixture" || d.SnapshotID != id {
		t.Fatalf("lost semantic failure: %+v", d)
	}
	if Check(id, `{"label":"unknown"}`, "unknown").Action != "accepted" {
		t.Fatal("expected abstention rejected")
	}
}

func TestSnapshotPreservesOrderAndProvenance(t *testing.T) {
	original := fixture()
	b, id, err := Freeze(original)
	if err != nil {
		t.Fatal(err)
	}
	var restored Invocation
	if err := json.Unmarshal(b, &restored); err != nil {
		t.Fatal(err)
	}
	again, replayID, _ := Freeze(restored)
	if !bytes.Equal(b, again) || id != replayID {
		t.Fatal("snapshot did not round trip")
	}
	for name, mutate := range map[string]func(*Invocation){
		"order":   func(in *Invocation) { in.Examples[0], in.Examples[1] = in.Examples[1], in.Examples[0] },
		"rule":    func(in *Invocation) { in.Rule = "interest counts as buyer" },
		"source":  func(in *Invocation) { in.SourceVersion = "v2" },
		"example": func(in *Invocation) { in.Examples[2].Output = "buyer" },
	} {
		t.Run(name, func(t *testing.T) {
			changed := fixture()
			mutate(&changed)
			changedBytes, changedID, err := Freeze(changed)
			if err != nil {
				t.Fatal(err)
			}
			if bytes.Equal(b, changedBytes) || id == changedID {
				t.Fatal("change disappeared from trace")
			}
		})
	}
}

func TestMissingContractAndInvalidCandidates(t *testing.T) {
	in := fixture()
	in.Rule = ""
	if _, _, err := Freeze(in); err == nil {
		t.Fatal("missing rule accepted")
	}
	for _, candidate := range []string{`broken`, `{"label":"lead"}`, `{}`} {
		if Check("snapshot", candidate, "unknown").Action != "rejected" {
			t.Fatalf("accepted %s", candidate)
		}
	}
}
