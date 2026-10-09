// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"go.yaml.in/yaml/v3"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/plugin"
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
	for _, want := range []string{"required", "review", "token", "schedule", model.ApprovedLabel, "xeno-labels.sh"} {
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

// Section 13 gives CI the image and the developer machine the binaries, and the wrapper is
// where that reaches a project: a job that names the image declares its runner version,
// where one that downloads a release binary resolves whichever version the download
// returns. The reference is asserted whole, repository and version together, because a
// wrapper naming the right repository at the wrong version compares the phase against a
// runner that computed it differently and reports that as a divergence in the trail.
func TestEveryWrapperNamesTheRunnerImage(t *testing.T) {
	const version = "1.2.3"
	for _, id := range WrapperHosts() {
		_, body, err := Wrapper("", id, version)
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		if want := model.RunnerImage(version); !strings.Contains(body, want) {
			t.Errorf("%s: the wrapper does not name the image %q:\n%s", id, want, body)
		}
		// The image is what the runner comes from, so an install step is not a step that
		// is merely redundant: two routes to a runner is two versions it could be.
		for _, unwanted := range []string{"install xeno", "alpine"} {
			if strings.Contains(body, unwanted) {
				t.Errorf("%s: the wrapper still carries %q, beside the image it names", id, unwanted)
			}
		}
	}
}

// The uid a container job runs as is the adopter's to set and nobody else's: the host mounts
// its work directory in and chowns nothing, actions/checkout runs inside the container, and a
// hosted runner happens to match the image at 1001 while a self-hosted one may not. So the
// GitHub wrapper carries the knob commented, with its reason, and the GitLab wrapper carries
// nothing: that executor leaves the build directory world writable, so a line about a uid
// there would describe a problem the host does not have (#324).
func TestOnlyTheGitHubWrapperOffersTheContainerUser(t *testing.T) {
	const line = "# options: --user 1001"
	for _, id := range WrapperHosts() {
		_, body, err := Wrapper("", id, "1.2.3")
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		switch id {
		case "github":
			if !strings.Contains(body, line) {
				t.Errorf("the github wrapper does not offer %q:\n%s", line, body)
			}
			// Commented, and the assertion says so: an uncommented --user would carry a uid
			// this repository cannot know, right on one runner and wrong on every other.
			if strings.Contains(body, "\n      options:") {
				t.Error("the github wrapper sets a container user rather than offering one")
			}
		default:
			if strings.Contains(body, "--user") {
				t.Errorf("%s: the wrapper names a container user and that host needs none", id)
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

// Section 13 says --vendor copies plugin.json and skills/ beside the templates and the rules. The
// skills are what a project reads to know what a phase owes, so a vendored project without them
// has the runner and no description of the process (#169).
func TestVendorPutsThePluginsOwnArtifactsInTheRepository(t *testing.T) {
	r := initFixture(t)
	if _, err := r.Init(InitOptions{Host: "github", Vendor: true}); err != nil {
		t.Fatal(err)
	}
	for _, rel := range []string{
		".xeno/plugin/.claude-plugin/plugin.json",
		".xeno/plugin/hooks/hooks.json",
		// Section 13 puts mcp.json at the distribution root and in the --vendor set. A
		// project that vendored everything but this one file has a server on the machine
		// and no client that knows to start it (#285).
		".xeno/plugin/mcp.json",
		".xeno/plugin/skills/xeno-intake/SKILL.md",
		".xeno/plugin/skills/xeno-learning/SKILL.md",
		".xeno/plugin/skills/xeno-review/SKILL.md",
	} {
		if !fm.Exists(filepath.Join(r.Root, rel)) {
			t.Errorf("%s was not vendored", rel)
		}
	}
	// Eleven skills, as section 13 names them: the six phases, the cross cutting learning record
	// and the four lenses. The lenses travel whether or not a project enables any, because
	// enablement is a line in project.yaml and a lens that is not vendored could not be enabled
	// at all.
	entries, err := os.ReadDir(filepath.Join(r.Root, ".xeno/plugin/skills"))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 11 {
		t.Errorf("%d skills vendored, want the eleven of section 13", len(entries))
	}

	// And a second run changes nothing, which is WP9's own criterion and now has a third tree and
	// two files to not break.
	before := snapshot(t, r.Root)
	if _, err := r.Init(InitOptions{Host: "github", Vendor: true}); err != nil {
		t.Fatal(err)
	}
	if after := snapshot(t, r.Root); after != before {
		t.Error("a second init changed the vendored plugin")
	}
}

// ---- the vendored tree is the plugin, not a list of its parts (#199)

// TestTheVendoredTreeHasTheDigestTheSourceHas is the property G-Supply needs and the one that
// was impossible before this walk: a released runner carries the digest of the plugin it
// shipped, and `init --vendor` has to produce exactly those bytes at exactly those paths or
// the gate fails for every adopter on every phase.
//
// It compares digests rather than listing files, because the digest is what the gate compares
// and a list would pass while a byte differed.
func TestTheVendoredTreeHasTheDigestTheSourceHas(t *testing.T) {
	r := initFixture(t)
	if _, err := r.Init(InitOptions{TrackerKey: "o/r", Vendor: true, Host: "github"}); err != nil {
		t.Fatal(err)
	}
	want := plugin.Hash(filepath.Join("..", ".."))
	if want == "" {
		t.Fatal("no digest for this repository's own plugin")
	}
	if got := plugin.Hash(r.Root); got != want {
		t.Errorf("the vendored tree hashes %s and the source %s: an adopter's G-Supply "+
			"would fail on every phase", got, want)
	}
}

// Every file, at any depth, with nothing named. The walk replaced three functions that each
// knew a tree and a depth, and twice something was added to the plugin that none of them
// copied — `secrets.yaml`, which the code's own comment listed, and `bin/`, added with the
// entry point. A list of parts goes stale; this asserts the whole.
func TestVendorCopiesEveryFileOfThePlugin(t *testing.T) {
	r := initFixture(t)
	if _, err := r.Init(InitOptions{TrackerKey: "o/r", Vendor: true, Host: "github"}); err != nil {
		t.Fatal(err)
	}
	var want, got []string
	collect := func(root string, into *[]string) {
		base := filepath.Join(root, plugin.Dir)
		if err := filepath.WalkDir(base, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}
			rel, rerr := filepath.Rel(base, p)
			*into = append(*into, filepath.ToSlash(rel))
			return rerr
		}); err != nil {
			t.Fatal(err)
		}
	}
	collect(filepath.Join("..", ".."), &want)
	collect(r.Root, &got)
	sort.Strings(want)
	sort.Strings(got)
	if strings.Join(want, "\n") != strings.Join(got, "\n") {
		t.Errorf("vendored set differs from the plugin:\n  want %v\n  got  %v", want, got)
	}
}

// The entry point the plugin's hook invokes has to be executable, and the mode cannot travel:
// an embed.FS reports every file read-only whatever was committed, so init sets it.
func TestTheVendoredEntryPointIsExecutable(t *testing.T) {
	r := initFixture(t)
	if _, err := r.Init(InitOptions{TrackerKey: "o/r", Vendor: true, Host: "github"}); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(filepath.Join(r.Root, plugin.Dir, "bin", "xeno-env.sh"))
	if err != nil {
		t.Fatalf("the entry point was not vendored: %v", err)
	}
	if info.Mode()&0o111 == 0 {
		t.Errorf("mode %v, want executable: the hook runs this file", info.Mode())
	}
}

// A build that carries no plugin and is pointed at no plugin refuses, naming which build it
// is rather than which file was missing. That is a development build with a wrong
// --plugin-from, and the message has to say that the release would have carried one.
func TestWithoutAPluginToVendorInitRefuses(t *testing.T) {
	r := initFixture(t)
	r.PluginSource = filepath.Join(t.TempDir(), "not-a-plugin")
	_, err := r.Init(InitOptions{TrackerKey: "o/r", Vendor: true, Host: "github"})
	var ref *Refusal
	if !errors.As(err, &ref) {
		t.Fatalf("got %v, want a refusal", err)
	}
	for _, want := range []string{"carries no plugin", "--plugin-from"} {
		if !strings.Contains(ref.Reason, want) {
			t.Errorf("the refusal does not say %q: %q", want, ref.Reason)
		}
	}
}

// Where it vendored from is reported, because the answer decides whether G-Supply can pass:
// a release carries the bytes its digest was taken over and a directory is whatever is in it.
func TestInitSaysWhereThePluginCameFrom(t *testing.T) {
	r := initFixture(t)
	res, err := r.Init(InitOptions{TrackerKey: "o/r", Vendor: true, Host: "github"})
	if err != nil {
		t.Fatal(err)
	}
	if res.PluginFrom == "" {
		t.Error("init does not say where the plugin came from")
	}
}

// A pipeline's results reach the gate as a directory, and the generated wrapper is where
// an adopter meets that. The runner does not fetch them: `xeno gate ...` never touches
// the network, so the fetch is a step of the job and the directory it leaves behind is
// what the gate reads. A wrapper that named no directory, or named one it then did not
// pass, would leave that step to be guessed at.
func TestEveryWrapperFillsTheEvidenceDirectoryAndPassesIt(t *testing.T) {
	for _, id := range WrapperHosts() {
		_, body, err := Wrapper("", id, "1.2.3")
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		if !strings.Contains(body, scaffold.EvidenceDir) {
			t.Errorf("%s: the wrapper names no evidence directory:\n%s", id, body)
		}
		if want := "--evidence-from " + scaffold.EvidenceDir; !strings.Contains(body, want) {
			t.Errorf("%s: the wrapper does not pass %q to the runner:\n%s", id, want, body)
		}
	}
}

// Both wrappers are generated into somebody else's repository, where a syntax error is
// found by their CI and not here. Each template is prose and YAML in one file, so an
// added comment or step is exactly the edit that can indent a key wrongly.
func TestEveryRenderedWrapperParsesAsYAML(t *testing.T) {
	for _, id := range WrapperHosts() {
		_, body, err := Wrapper("", id, "1.2.3")
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		var doc any
		if err := yaml.Unmarshal([]byte(body), &doc); err != nil {
			t.Errorf("%s: the rendered wrapper is not YAML: %v\n%s", id, err, body)
		}
	}
}
