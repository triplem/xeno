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
	if len(res.Unbindable) != 1 || !strings.Contains(res.Unbindable[0], "scan/trivy") {
		t.Errorf("the reason does not name the entry: %q", res.Unbindable)
	}
	if !strings.Contains(strings.Join(res.Unbindable, " "), "sha256") {
		t.Errorf("the reason does not say what is missing: %q", res.Unbindable)
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
	if res.Attached != 0 || res.Pending != 1 || len(res.Unbindable) != 1 {
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
	if res.Attached != 1 || res.Pending != 0 || len(res.Unbindable) != 0 {
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
	if res.Attached != 1 || len(res.Unbindable) != 0 {
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
	if res.Attached != 1 || res.Pending != 1 || len(res.Unbindable) != 1 {
		t.Fatalf("%+v", res)
	}
	got := attachedRecords(t, root, dir)
	if len(got) != 1 || got[0].Job != "cyclonedx" {
		t.Fatalf("the bindable entry was not the one recorded: %+v", got)
	}
}
