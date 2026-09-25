// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
)

func initFixture(t *testing.T) *Runner {
	t.Helper()
	r := New(t.TempDir())
	r.PluginSource = filepath.Join("..", "..", ".xeno", "plugin")
	return r
}

// WP9's first acceptance. init is the command somebody runs again after reading the
// output of the first one, so a version of it that overwrote would eat whatever they
// changed in between.
func TestRunningInitTwiceChangesNothingTheSecondTime(t *testing.T) {
	r := initFixture(t)
	first, err := r.Init(InitOptions{Vendor: true, TrackerKey: "owner/repo"})
	if err != nil {
		t.Fatal(err)
	}
	if len(first.Created) == 0 {
		t.Fatal("the first run created nothing")
	}
	before := snapshot(t, r.Root)

	second, err := r.Init(InitOptions{Vendor: true, TrackerKey: "somebody/else"})
	if err != nil {
		t.Fatal(err)
	}
	if len(second.Created) != 0 {
		t.Fatalf("the second run created %v", second.Created)
	}
	if after := snapshot(t, r.Root); after != before {
		t.Fatal("the second run changed the tree")
	}
}

// It never edits a file it did not create. The one exception is .gitignore, where it
// appends a line that is not there and leaves the rest alone, because that file belongs
// to the project.
func TestInitAppendsToGitignoreAndOverwritesNothing(t *testing.T) {
	r := initFixture(t)
	existing := "build/\nnode_modules/\n"
	if err := os.WriteFile(filepath.Join(r.Root, ".gitignore"), []byte(existing), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := r.Init(InitOptions{}); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(r.Root, ".gitignore"))
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)
	if !strings.HasPrefix(got, existing) {
		t.Fatalf("what was there did not survive:\n%s", got)
	}
	if !strings.Contains(got, ".xeno/local/") {
		t.Fatalf("the entry was not added:\n%s", got)
	}
	// And again: the line is there, so nothing is appended a second time.
	if _, err := r.Init(InitOptions{}); err != nil {
		t.Fatal(err)
	}
	b2, _ := os.ReadFile(filepath.Join(r.Root, ".gitignore"))
	if string(b2) != got {
		t.Fatalf("a second run appended again:\n%s", b2)
	}
}

// A version mismatch stops it, on any difference and not only a major one. A different
// runner produces a different hash or a different verdict than CI.
func TestAVersionMismatchStopsInit(t *testing.T) {
	r := initFixture(t)
	if _, err := r.Init(InitOptions{}); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(r.Root, ".xeno/config/project.yaml")
	b, _ := os.ReadFile(path)
	edited := strings.Replace(string(b), "runner_version: "+model.RunnerVersion, "runner_version: 9.9.9", 1)
	if edited == string(b) {
		t.Fatal("the pin is not written the way this test expects")
	}
	if err := os.WriteFile(path, []byte(edited), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := r.Init(InitOptions{})
	if err == nil {
		t.Fatal("init proceeded against a repository pinned to another runner")
	}
	if !strings.Contains(err.Error(), "9.9.9") {
		t.Fatalf("the refusal does not name the pin: %v", err)
	}
}

// The settings a person still has to make are the point of the output. A first contact
// that leaves the project believing the gate is binding when it is not is worse than no
// first contact.
func TestInitNamesWhatItCannotDo(t *testing.T) {
	r := initFixture(t)
	res, err := r.Init(InitOptions{})
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(res.Manual, "\n")
	for _, want := range []string{"required", "review", "token", "schedule"} {
		if !strings.Contains(joined, want) {
			t.Errorf("the manual settings do not mention %q:\n%s", want, joined)
		}
	}
	if len(res.Outstand) == 0 {
		t.Fatal("nothing was reported as undetermined, although no host was asked")
	}
}

func TestInitWritesAConfigurationTheRunnerCanReadBack(t *testing.T) {
	r := initFixture(t)
	if _, err := r.Init(InitOptions{TrackerKey: "owner/repo", Model: "a-model", Language: "de"}); err != nil {
		t.Fatal(err)
	}
	var p model.Project
	if err := fm.ReadYAML(filepath.Join(r.Root, ".xeno/config/project.yaml"), &p); err != nil {
		t.Fatal(err)
	}
	if p.Language.Artifacts != "de" {
		t.Errorf("language is %q, want de", p.Language.Artifacts)
	}
	if p.RunnerVersion != model.RunnerVersion {
		t.Errorf("the pin is %q, want the running version %q", p.RunnerVersion, model.RunnerVersion)
	}
	if p.Evidence.Source != "ci" {
		t.Errorf("evidence source is %q, want the ci default", p.Evidence.Source)
	}
	// And the language it wrote is the one a phase renders in.
	if r.language() != "de" {
		t.Errorf("the runner reads %q back", r.language())
	}
}

func TestVendorPutsTheShippedSetInTheRepository(t *testing.T) {
	r := initFixture(t)
	if _, err := r.Init(InitOptions{Vendor: true}); err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"intake", "design", "review"} {
		for _, f := range []string{"template.yaml", "strings.en.yaml", "strings.de.yaml"} {
			p := filepath.Join(r.Root, ".xeno/plugin/templates", id, f)
			if !fm.Exists(p) {
				t.Errorf("%s was not vendored", p)
			}
		}
	}
	// Without --vendor nothing is copied, since a repository may carry its own.
	r2 := initFixture(t)
	if _, err := r2.Init(InitOptions{}); err != nil {
		t.Fatal(err)
	}
	if fm.Exists(filepath.Join(r2.Root, ".xeno/plugin/templates/intake/template.yaml")) {
		t.Error("the plugin was vendored without being asked for")
	}
}

// snapshot is every path with the size of its content, which is enough to see a write.
func snapshot(t *testing.T, root string) string {
	t.Helper()
	var b strings.Builder
	err := filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(root, p)
		b.WriteString(rel)
		b.WriteString(":")
		b.WriteString(string(rune(info.Size())))
		b.WriteString("\n")
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return b.String()
}
