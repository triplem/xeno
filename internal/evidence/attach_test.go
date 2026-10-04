// SPDX-License-Identifier: Apache-2.0

package evidence

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
)

const intent = "git.example/group/proj#1"

// phase writes a phase declaring one pending item, and a source directory holding the
// manifest the pipeline is standing in for. Both are what Attach reads and nothing else.
func phase(t *testing.T, declared, manifest string) (root, source, dir string) {
	t.Helper()
	root = t.TempDir()
	dir = model.PhaseDir("PROJ-1", "04-verification")
	if err := os.MkdirAll(filepath.Join(root, dir), 0o755); err != nil {
		t.Fatal(err)
	}
	out := "---\nintent: " + intent + "\nphase: 04-verification\nevidence:\n" + declared + "---\n\nbody\n"
	if err := os.WriteFile(filepath.Join(root, dir, "output.md"), []byte(out), 0o644); err != nil {
		t.Fatal(err)
	}
	source = filepath.Join(t.TempDir(), "artifacts")
	if err := os.MkdirAll(source, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "manifest.yaml"), []byte(manifest), 0o644); err != nil {
		t.Fatal(err)
	}
	return root, source, dir
}

func attachedRecords(t *testing.T, root, dir string) []model.Attached {
	t.Helper()
	var att []model.Attached
	err := fm.ReadYAML(filepath.Join(root, dir, "evidence", "attached.yaml"), &att)
	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
	return att
}

const pending = "  - kind: scan\n    job: trivy\n"

// The defect this fixes. A uri names where something lives and its hash is the only thing
// that says what it is, per A8, so an entry arriving without one is bound by nothing.
func TestAUriWithoutAHashIsDeclined(t *testing.T) {
	root, source, dir := phase(t, pending,
		"- kind: scan\n  job: trivy\n  result: pass\n  uri: https://ci.example/artifacts/7\n")
	res, err := Attach(root, "PROJ-1", "04-verification", source)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attached != 0 {
		t.Errorf("an entry nothing can bind was attached: %+v", res)
	}
	if res.Pending != 1 {
		t.Errorf("pending = %d, want 1: declining to record it leaves the phase as it was", res.Pending)
	}
	if len(res.Declined) != 1 || !strings.Contains(res.Declined[0], "scan/trivy") {
		t.Errorf("the reason does not name the entry: %q", res.Declined)
	}
	if !strings.Contains(strings.Join(res.Declined, " "), "sha256") {
		t.Errorf("the reason does not say what is missing: %q", res.Declined)
	}
	if got := attachedRecords(t, root, dir); len(got) != 0 {
		t.Errorf("attached.yaml holds %d records, want none: %+v", len(got), got)
	}
}

// The same defect with less to look at: a result and nothing it belongs to.
func TestAnEntryWithNoFileUriOrHashIsDeclined(t *testing.T) {
	root, source, dir := phase(t, pending, "- kind: scan\n  job: trivy\n  result: pass\n")
	res, err := Attach(root, "PROJ-1", "04-verification", source)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attached != 0 || res.Pending != 1 || len(res.Declined) != 1 {
		t.Fatalf("an entry naming nothing was attached: %+v", res)
	}
	if got := attachedRecords(t, root, dir); len(got) != 0 {
		t.Errorf("attached.yaml holds %d records, want none", len(got))
	}
}

// The guard is narrow: an entry that carries both is attached as before.
func TestAUriWithAHashIsAttached(t *testing.T) {
	const sum = "1111111111111111111111111111111111111111111111111111111111111111"
	root, source, dir := phase(t, pending,
		"- kind: scan\n  job: trivy\n  result: pass\n  uri: https://ci.example/a/7\n  sha256: "+sum+"\n")
	res, err := Attach(root, "PROJ-1", "04-verification", source)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attached != 1 || res.Pending != 0 || len(res.Declined) != 0 {
		t.Fatalf("a bindable entry was not attached: %+v", res)
	}
	got := attachedRecords(t, root, dir)
	if len(got) != 1 || got[0].SHA256 != sum || got[0].URI == "" {
		t.Fatalf("the record does not carry what the pipeline published: %+v", got)
	}
}

// A file backed entry never reaches the guard: its hash comes from the bytes copied in, so
// it cannot be in the state the guard is for. Kept so that the narrowing is checked and not
// assumed.
func TestAFileBackedEntryIsHashedFromItsBytes(t *testing.T) {
	root, source, dir := phase(t, pending,
		"- kind: scan\n  job: trivy\n  result: pass\n  file: trivy.json\n")
	if err := os.WriteFile(filepath.Join(source, "trivy.json"), []byte("{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	res, err := Attach(root, "PROJ-1", "04-verification", source)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attached != 1 || len(res.Declined) != 0 {
		t.Fatalf("a file backed entry was declined: %+v", res)
	}
	got := attachedRecords(t, root, dir)
	if len(got) != 1 || got[0].SHA256 == "" || got[0].Path != "evidence/trivy.json" {
		t.Fatalf("the record does not carry the hash of what was copied: %+v", got)
	}
}

// One entry declined does not take a good one with it. Attach is called once per phase and
// a phase declares several items, so a refusal that aborted would hold back evidence that
// arrived correctly.
func TestOneDeclinedEntryDoesNotHoldBackTheOthers(t *testing.T) {
	const sum = "2222222222222222222222222222222222222222222222222222222222222222"
	root, source, dir := phase(t,
		"  - kind: scan\n    job: trivy\n  - kind: sbom\n    job: cyclonedx\n",
		"- kind: scan\n  job: trivy\n  result: pass\n  uri: https://ci.example/a/7\n"+
			"- kind: sbom\n  job: cyclonedx\n  uri: https://ci.example/a/8\n  sha256: "+sum+"\n")
	res, err := Attach(root, "PROJ-1", "04-verification", source)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attached != 1 || res.Pending != 1 || len(res.Declined) != 1 {
		t.Fatalf("%+v", res)
	}
	got := attachedRecords(t, root, dir)
	if len(got) != 1 || got[0].Job != "cyclonedx" {
		t.Fatalf("the bindable entry was not the one recorded: %+v", got)
	}
}

// ---- #221: a result the document does not define is not recorded either

// The way all fifteen of this repository's hand written attachments got in. The set is
// closed in section 4, G-Evidence judges a recorded value against it, and an entry written
// here with `success` in it would be reported there with nothing in between saying which of
// the two to believe.
func TestAResultOutsideTheSetIsDeclined(t *testing.T) {
	root, source, dir := phase(t, pending,
		"- kind: scan\n  job: trivy\n  result: success\n  file: trivy.json\n")
	if err := os.WriteFile(filepath.Join(source, "trivy.json"), []byte("{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	res, err := Attach(root, "PROJ-1", "04-verification", source)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attached != 0 || res.Pending != 1 || len(res.Declined) != 1 {
		t.Fatalf("an entry carrying a word section 4 does not define was recorded: %+v", res)
	}
	if !strings.Contains(res.Declined[0], "scan/trivy") || !strings.Contains(res.Declined[0], "success") {
		t.Errorf("the reason does not name the entry and the value: %q", res.Declined)
	}
	// The remedy is in the sentence, because the three reasons are fixed in three places.
	if !strings.Contains(res.Declined[0], "Republish") {
		t.Errorf("the reason does not say what to do: %q", res.Declined)
	}
	// Nothing is left behind: no record, and no report copied beside the artifact, so a
	// corrected manifest binds the declaration that is still pending.
	if got := attachedRecords(t, root, dir); len(got) != 0 {
		t.Fatalf("the declined entry was recorded anyway: %+v", got)
	}
	if _, err := os.Stat(filepath.Join(root, dir, "evidence", "trivy.json")); err == nil {
		t.Error("the report of a declined entry was copied into evidence/")
	}
}

// It judges a value and not an absence. Section 4 has the field absent where a producer
// reports nothing, and this repository's vulnerability scan publishes exactly such an entry
// beside its report: `other/trivy-db`, the database metadata, with a file and no result.
func TestAnEntryWithNoResultIsStillRecorded(t *testing.T) {
	root, source, dir := phase(t, "  - kind: other\n    job: trivy-db\n",
		"- kind: other\n  job: trivy-db\n  file: db-metadata.json\n")
	if err := os.WriteFile(filepath.Join(source, "db-metadata.json"), []byte("{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	res, err := Attach(root, "PROJ-1", "04-verification", source)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attached != 1 || len(res.Declined) != 0 {
		t.Fatalf("an entry whose producer reports nothing was declined: %+v", res)
	}
	if got := attachedRecords(t, root, dir); len(got) != 1 || got[0].Result != "" {
		t.Fatalf("the record invented a result: %+v", got)
	}
}

// Every value the document defines is recorded as it was published, including fail: a run
// that reported failure is a successful run reporting fail, and section 4 says so.
func TestEveryDocumentedResultIsRecorded(t *testing.T) {
	for _, want := range model.EvidenceResults {
		root, source, dir := phase(t, pending,
			"- kind: scan\n  job: trivy\n  result: "+want+"\n  file: trivy.json\n")
		if err := os.WriteFile(filepath.Join(source, "trivy.json"), []byte("{}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		res, err := Attach(root, "PROJ-1", "04-verification", source)
		if err != nil {
			t.Fatal(err)
		}
		if res.Attached != 1 || len(res.Declined) != 0 {
			t.Fatalf("result %s was declined: %+v", want, res)
		}
		if got := attachedRecords(t, root, dir); len(got) != 1 || got[0].Result != want {
			t.Fatalf("result %s was not recorded as published: %+v", want, got)
		}
	}
}
