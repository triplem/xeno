// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/rules"
	"github.com/triplem/xeno/internal/scaffold"
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
	first, err := r.Init(InitOptions{Host: "github", Vendor: true, TrackerKey: "owner/repo"})
	if err != nil {
		t.Fatal(err)
	}
	if len(first.Created) == 0 {
		t.Fatal("the first run created nothing")
	}
	before := snapshot(t, r.Root)

	second, err := r.Init(InitOptions{Host: "github", Vendor: true, TrackerKey: "somebody/else"})
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
	if _, err := r.Init(InitOptions{Host: "github"}); err != nil {
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
	if _, err := r.Init(InitOptions{Host: "github"}); err != nil {
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
	if _, err := r.Init(InitOptions{Host: "github"}); err != nil {
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
	_, err := r.Init(InitOptions{Host: "github"})
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
	res, err := r.Init(InitOptions{Host: "github"})
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
	if _, err := r.Init(InitOptions{Host: "github", TrackerKey: "owner/repo", Model: "a-model", Language: "de"}); err != nil {
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
	if _, err := r.Init(InitOptions{Host: "github", Vendor: true}); err != nil {
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
	if _, err := r2.Init(InitOptions{Host: "github"}); err != nil {
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

// WP10: the two host specific things in the wrapper are the expressions for the ends of
// the commit range, which is what makes another host an entry in a table rather than a
// second generator.
func TestTheWrapperPassesBothEndsOfTheRange(t *testing.T) {
	h, body, err := Wrapper("", "github", "1.2.3")
	if err != nil {
		t.Fatal(err)
	}
	if h.BaseRef == "" || h.HeadRef == "" || h.BaseRef == h.HeadRef {
		t.Fatalf("the host contributes no distinct range expressions: %+v", h)
	}
	for _, want := range []string{"--base " + h.BaseRef, "--head " + h.HeadRef, "1.2.3", "xeno gate run"} {
		if !strings.Contains(body, want) {
			t.Errorf("the wrapper does not carry %q", want)
		}
	}
	// It calls one command and does nothing else: no build, no test.
	for _, unwanted := range []string{"go build", "go test", "semantic-release"} {
		if strings.Contains(body, unwanted) {
			t.Errorf("the wrapper carries %q, which belongs in the project's own pipeline", unwanted)
		}
	}
	if _, _, err := Wrapper("", "nowhere", "1.2.3"); err == nil {
		t.Fatal("a wrapper was generated for an unknown host")
	}
}

func TestInitGeneratesTheWrapper(t *testing.T) {
	r := initFixture(t)
	res, err := r.Init(InitOptions{Host: "github"})
	if err != nil {
		t.Fatal(err)
	}
	h := wrapperHosts["github"]
	if !fm.Exists(filepath.Join(r.Root, h.Path)) {
		t.Fatalf("%s was not generated", h.Path)
	}
	found := false
	for _, p := range res.Created {
		if p == h.Path {
			found = true
		}
	}
	if !found {
		t.Errorf("the wrapper is not in what init reports it created: %v", res.Created)
	}
}

// The scaffolds are files, so a repository replaces one without forking the runner, and
// the generated file says which copy it came from: the first question anybody asks of a
// generated file that does not look like they expect.
func TestAProjectOverridesAScaffold(t *testing.T) {
	r := initFixture(t)
	dir := filepath.Join(r.Root, scaffold.OverrideDir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := "# ours, from {{.Source}}\nname: xeno gate\nsteps: [{run: xeno gate run --base {{.BaseRef}}}]\n"
	if err := os.WriteFile(filepath.Join(dir, "ci-github.yml"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	_, generated, err := Wrapper(r.Root, "github", "1.2.3")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(generated, "# ours") {
		t.Fatalf("the override was not used:\n%s", generated)
	}
	if !strings.Contains(generated, string(scaffold.FromProject)) {
		t.Fatalf("the generated file does not name its source:\n%s", generated)
	}
	// The host still contributes the range expressions, so an override inherits them
	// rather than hard coding one host's syntax.
	if !strings.Contains(generated, "${{ github.event.pull_request.base.sha }}") {
		t.Fatalf("the override did not receive the host's base ref:\n%s", generated)
	}

	// Without an override the runner's own default is used, and says so.
	r2 := initFixture(t)
	_, fallback, err := Wrapper(r2.Root, "github", "1.2.3")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(fallback, string(scaffold.FromRunner)) {
		t.Fatalf("the default does not name itself:\n%s", fallback)
	}
}

// The same for project.yaml, which is the other file init writes into somebody else's
// repository, and the one where a positional argument in the wrong place used to produce
// a file that looked right.
func TestTheProjectConfigurationComesFromTheScaffold(t *testing.T) {
	r := initFixture(t)
	res, err := r.Init(InitOptions{Host: "github", TrackerKey: "o/r", Model: "m", Language: "de"})
	if err != nil {
		t.Fatal(err)
	}
	_ = res
	b, err := os.ReadFile(filepath.Join(r.Root, ".xeno/config/project.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)
	for _, want := range []string{"project: o/r", "default: m", "artifacts: de", string(scaffold.FromRunner)} {
		if !strings.Contains(got, want) {
			t.Errorf("the configuration does not carry %q:\n%s", want, got)
		}
	}
}

// #104's first done-when, the half #143 did not reach: no default host name appears in
// code. The flag's default was the visible one and Init filled an empty host in as well,
// so the absence is asserted rather than left to a reading of two files.
func TestInitRefusesWithoutAHost(t *testing.T) {
	r := initFixture(t)
	_, err := r.Init(InitOptions{})
	if err == nil {
		t.Fatal("init picked a host instead of refusing")
	}
	if !strings.Contains(err.Error(), "github") || !strings.Contains(err.Error(), "gitlab") {
		t.Errorf("the refusal does not say what there is: %v", err)
	}
	var refusal *Refusal
	if !errors.As(err, &refusal) {
		t.Errorf("a missing host is %T, so the exit code is 2 rather than 1 by A11", err)
	}
}

// A wrapper whose range is wrong produces verdicts about the wrong commits, and nothing
// downstream says so: gate.yaml would be internally consistent and about another change.
// So the two expressions are asserted in the generated text of every host.
func TestEveryWrapperPassesBothEndsOfTheRange(t *testing.T) {
	for _, id := range WrapperHosts() {
		r := initFixture(t)
		if _, err := r.Init(InitOptions{Host: id}); err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		h := wrapperHosts[id]
		body, err := os.ReadFile(filepath.Join(r.Root, h.Path))
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		for _, want := range []string{h.BaseRef, h.HeadRef, "xeno gate run", "xeno gate verify"} {
			if !strings.Contains(string(body), want) {
				t.Errorf("%s: the wrapper does not carry %q", id, want)
			}
		}
	}
}

// CI_MERGE_REQUEST_DIFF_BASE_SHA is populated only in a merge request pipeline, so a
// GitLab wrapper that ran on a branch pipeline would pass an empty base. The scoping is
// what makes the range reachable, which is why it is asserted and not left to the reader.
func TestTheGitLabWrapperRunsOnlyOnMergeRequests(t *testing.T) {
	r := initFixture(t)
	if _, err := r.Init(InitOptions{Host: "gitlab"}); err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(filepath.Join(r.Root, wrapperHosts["gitlab"].Path))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `$CI_PIPELINE_SOURCE == "merge_request_event"`) {
		t.Error("the job is not scoped to merge request pipelines, so the base SHA is empty")
	}
}

// The wrapper and the tracker block are the same choice, and they used to be made twice:
// once from the flag and once from a literal in the project scaffold. A repository whose
// pipeline is one host's and whose tracker configuration is another's works as neither.
func TestTheTrackerBlockNamesTheHostTheWrapperWasGeneratedFor(t *testing.T) {
	for _, id := range WrapperHosts() {
		r := initFixture(t)
		if _, err := r.Init(InitOptions{Host: id, TrackerKey: "o/r"}); err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		var p struct {
			Tracker struct {
				Adapter string `yaml:"adapter"`
				BaseURL string `yaml:"base_url"`
			} `yaml:"tracker"`
		}
		if err := fm.ReadYAML(filepath.Join(r.Root, ".xeno/config/project.yaml"), &p); err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		h := wrapperHosts[id]
		if p.Tracker.Adapter != h.Adapter {
			t.Errorf("%s: tracker.adapter is %q, want %q", id, p.Tracker.Adapter, h.Adapter)
		}
		if p.Tracker.BaseURL != h.APIBase {
			t.Errorf("%s: tracker.base_url is %q, want %q", id, p.Tracker.BaseURL, h.APIBase)
		}
	}
}

// WrapperHosts returned a literal list beside the table that already held the names, which
// was correct for as long as there was one host. A third is one row, and this is what says
// so.
func TestWrapperHostsIsDerivedFromTheTableAndSorted(t *testing.T) {
	got := WrapperHosts()
	if len(got) != len(wrapperHosts) {
		t.Fatalf("WrapperHosts has %d entries and the table has %d", len(got), len(wrapperHosts))
	}
	for _, id := range got {
		if _, ok := wrapperHosts[id]; !ok {
			t.Errorf("%q is listed and not in the table", id)
		}
	}
	for i := 1; i < len(got); i++ {
		if got[i-1] > got[i] {
			t.Fatalf("not sorted, so an error message depends on map order: %v", got)
		}
	}
}

// The commit convention rests on a host setting, and the symptom of skipping it is a
// footer that vanished from history. Every host names its own, because the settings differ
// and the consequence does not.
func TestTheSquashSettingIsNamedForEveryHost(t *testing.T) {
	for _, id := range WrapperHosts() {
		r := initFixture(t)
		res, err := r.Init(InitOptions{Host: id})
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		found := false
		for _, m := range res.Manual {
			if strings.Contains(m, "squash") {
				found = true
			}
		}
		if !found {
			t.Errorf("%s: nothing in the manual list names the squash setting: %v", id, res.Manual)
		}
	}
}

// The rule tree is the second thing the plugin carries, and it was not vendored until #165: a
// shipped set that does not arrive is not shipped.
func TestVendorPutsTheShippedRuleSetInTheRepository(t *testing.T) {
	r := initFixture(t)
	if _, err := r.Init(InitOptions{Host: "github", Vendor: true}); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(r.Root, ".xeno/plugin/rules/given/builtin")
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("the rule tree was not vendored: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("the rule tree was vendored empty")
	}

	// And what arrived resolves, with no finding from the gate that reads it, in a repository
	// that has nothing else of its own.
	read, problems := rules.Load(r.Root)
	if len(problems) != 0 {
		t.Fatalf("the vendored set does not resolve: %v", problems)
	}
	effective, collisions := rules.Effective(read)
	if len(collisions) != 0 {
		t.Fatalf("the vendored set collides with itself: %v", collisions)
	}
	if len(effective) != len(entries) {
		t.Errorf("%d files vendored and %d rules in force", len(entries), len(effective))
	}

	// Without --vendor the rules are not copied either.
	r2 := initFixture(t)
	if _, err := r2.Init(InitOptions{Host: "github"}); err != nil {
		t.Fatal(err)
	}
	if fm.Exists(filepath.Join(r2.Root, ".xeno/plugin/rules/given/builtin")) {
		t.Error("the rule set was vendored without being asked for")
	}
}
