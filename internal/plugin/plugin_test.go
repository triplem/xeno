// SPDX-License-Identifier: Apache-2.0

// Package plugin holds no code. Its tests check the shipped plugin — the manifests, the skills and
// the hook wiring — against the tree the skills describe.
//
// A skill is prose an agent loads, and nothing here can check whether it produces better artifacts
// than no skill at all. Three things in it can be wrong mechanically, and those are checked: a
// command it names, a section it asks a phase to write, and a gate it says will refuse something.
// The rest is read by a person, which is the same limit the shipped rule set has.
package plugin

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/template"
)

const repoRoot = "../.."

// pluginDir is where the distribution tree lives in this repository: the same directory that
// already carries templates/, rules/given/builtin/ and secrets.yaml, which is what section 13
// lists at the distribution root.
const pluginDir = ".xeno/plugin"

// Section 13 names the skills. Six are the phases and xeno-learning is cross cutting, because
// learning happens at the end of each of the six and once more when an intent closes.
var phaseSkills = map[string]string{
	"xeno-intake":         "00-intake",
	"xeno-requirements":   "01-requirements",
	"xeno-design":         "02-design",
	"xeno-implementation": "03-implementation",
	"xeno-verification":   "04-verification",
	"xeno-review":         "05-review",
}

const learningSkill = "xeno-learning"

func skillPath(name string) string {
	return filepath.Join(repoRoot, pluginDir, "skills", name, "SKILL.md")
}

func read(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

// The seven of section 13, no more and no fewer: a client resolves a skill by name, and a set that
// drifted from the document would be a set nothing in the document describes.
func TestTheSkillsAreTheSevenSectionThirteenNames(t *testing.T) {
	entries, err := os.ReadDir(filepath.Join(repoRoot, pluginDir, "skills"))
	if err != nil {
		t.Fatal(err)
	}
	var found []string
	for _, e := range entries {
		if e.IsDir() {
			found = append(found, e.Name())
		}
	}
	sort.Strings(found)

	want := []string{learningSkill}
	for name := range phaseSkills {
		want = append(want, name)
	}
	sort.Strings(want)

	if strings.Join(found, ",") != strings.Join(want, ",") {
		t.Fatalf("skills are %v, want %v", found, want)
	}
	for _, name := range found {
		if _, err := os.Stat(skillPath(name)); err != nil {
			t.Errorf("%s has no SKILL.md", name)
		}
	}
}

// A skill is loaded by its frontmatter: a name that matches its directory and a description that
// says when to use it, which is what a client selects on.
func TestEverySkillCarriesANameAndADescription(t *testing.T) {
	for _, name := range allSkills() {
		body := read(t, skillPath(name))
		if !strings.HasPrefix(body, "---\n") {
			t.Errorf("%s has no frontmatter", name)
			continue
		}
		front := body[4 : strings.Index(body[4:], "---\n")+4]
		if !strings.Contains(front, "name: "+name+"\n") {
			t.Errorf("%s does not name itself in its frontmatter", name)
		}
		desc := ""
		for _, line := range strings.Split(front, "\n") {
			if strings.HasPrefix(line, "description: ") {
				desc = strings.TrimPrefix(line, "description: ")
			}
		}
		switch {
		case desc == "":
			t.Errorf("%s has no description", name)
		case len(strings.Fields(desc)) < 12:
			t.Errorf("%s's description is %d words, which is too short to select on",
				name, len(strings.Fields(desc)))
		case !strings.Contains(desc, "Use "):
			t.Errorf("%s's description does not say when to use it: %q", name, desc)
		}
	}
}

func allSkills() []string {
	out := []string{learningSkill}
	for name := range phaseSkills {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// commandPattern finds what a skill tells an agent to run: the indented command blocks, where the
// first two words after xeno are the command.
var commandPattern = regexp.MustCompile(`(?m)^ {4}xeno ([a-z-]+)(?: ([a-z-]+))?`)

// A skill that names a command the binary does not have is worse than no skill, and it is the
// likeliest defect in prose nothing compiles. The dispatch table is the authority.
func TestEveryCommandASkillNamesExists(t *testing.T) {
	surface := commandSurface(t)
	for _, name := range allSkills() {
		for _, m := range commandPattern.FindAllStringSubmatch(read(t, skillPath(name)), -1) {
			one, two := m[1], strings.TrimSpace(m[2])
			if surface[one] {
				continue // a one word command, init or version
			}
			if two != "" && surface[one+" "+two] {
				continue
			}
			t.Errorf("%s names `xeno %s %s`, which the dispatch table does not have", name, one, two)
		}
	}
}

// commandSurface is every command the binary answers, read out of main.go's dispatch table and the
// two it answers before dispatch.
func commandSurface(t *testing.T) map[string]bool {
	t.Helper()
	body := read(t, filepath.Join(repoRoot, "cmd/xeno/main.go"))
	out := map[string]bool{"init": true, "version": true}
	for _, m := range regexp.MustCompile(`(?m)^\t"([a-z]+(?: [a-z-]+)?)":\s+\{`).FindAllStringSubmatch(body, -1) {
		out[m[1]] = true
	}
	if len(out) < 10 {
		t.Fatalf("read %d commands out of the dispatch table, which cannot be right", len(out))
	}
	return out
}

// sectionPattern finds the sections a skill tells a phase to write, in its command block.
var sectionPattern = regexp.MustCompile(`(?m)^ {4}xeno section set\s+([a-z-]+)`)

// A phase skill may only ask for sections that phase's template has, for the same reason: a
// section the template does not carry cannot be written, and the agent would find out at the
// write rather than here.
func TestEverySectionAPhaseSkillAsksForIsInItsTemplate(t *testing.T) {
	for name, phase := range phaseSkills {
		tmpl, err := template.Load(repoRoot, model.TemplateID(phase), "en")
		if err != nil {
			t.Fatalf("%s: %v", phase, err)
		}
		for _, m := range sectionPattern.FindAllStringSubmatch(read(t, skillPath(name)), -1) {
			if !tmpl.Has(m[1]) {
				t.Errorf("%s asks for section %q, which %s does not have: it has %s",
					name, m[1], tmpl.Ref(), strings.Join(tmpl.Known(), ", "))
			}
		}
	}
}

// gatePattern finds the gates a skill says will refuse something.
var gatePattern = regexp.MustCompile(`G-[A-Z][a-z]+`)

// A skill may name a gate that exists, and may not promise a check from a gate that reports
// not-implemented: an agent told to expect a verdict it will never see learns to ignore the skill.
func TestEveryGateASkillNamesExistsAndIsImplemented(t *testing.T) {
	implemented := map[string]bool{}
	for _, phase := range model.Phases {
		for _, id := range gates.Applicable(phase) {
			implemented[id] = true
		}
	}
	// The three that report not-implemented. A skill may name one, because the gate list is a
	// declaration the runner grows into, and it has to say so in the same breath: an agent told
	// to expect a verdict it will never see learns to ignore the skill.
	unrun := map[string]bool{"G-Supply": true, "G-Secret": true, "G-Test": true}

	for _, name := range allSkills() {
		body := read(t, skillPath(name))
		for _, id := range gatePattern.FindAllString(body, -1) {
			if !implemented[id] {
				t.Errorf("%s names %s, which is not in the gate table", name, id)
				continue
			}
			if unrun[id] && !strings.Contains(body, "not-implemented") {
				t.Errorf("%s names %s, which reports not-implemented, and does not say so", name, id)
			}
		}
	}
}

// The skills travel to every project that installs the plugin, which is the bar given/builtin/ has:
// nothing in them may only make sense inside this repository.
func TestNoSkillNamesThisProjectsOwnWorld(t *testing.T) {
	forbidden := []string{"gofmt", "golang", "go.mod", "internal/", "cmd/xeno", "triplem", "detekt"}
	for _, name := range allSkills() {
		body := read(t, skillPath(name))
		for _, f := range forbidden {
			if strings.Contains(body, f) {
				t.Errorf("%s names %q, which is this repository rather than the process", name, f)
			}
		}
	}
}

// There is no MCP server, so a skill describing an operation describes something absent. WP11's
// done-when requires the command path to work alone, and these are the skills for that path.
func TestNoSkillPromisesTheServerThatDoesNotExist(t *testing.T) {
	for _, name := range allSkills() {
		body := strings.ToLower(read(t, skillPath(name)))
		for _, f := range []string{"mcp", "tool call", "operation"} {
			if strings.Contains(body, f) {
				t.Errorf("%s mentions %q, and no server exists", name, f)
			}
		}
	}
}

// The hook wiring comes from the plugin so that a project which installs it records a turn's cost
// without writing a settings file of its own.
func TestThePluginCarriesTheHookWiring(t *testing.T) {
	var hooks struct {
		Hooks map[string][]struct {
			Hooks []struct {
				Type    string `json:"type"`
				Command string `json:"command"`
			} `json:"hooks"`
		} `json:"hooks"`
	}
	body := read(t, filepath.Join(repoRoot, pluginDir, "hooks/hooks.json"))
	if err := json.Unmarshal([]byte(body), &hooks); err != nil {
		t.Fatal(err)
	}
	stop, ok := hooks.Hooks["Stop"]
	if !ok || len(stop) == 0 || len(stop[0].Hooks) == 0 {
		t.Fatal("the plugin wires no Stop hook, so an installing project records no cost")
	}
	cmd := stop[0].Hooks[0].Command
	if cmd != "xeno cost turn" {
		t.Errorf("the hook runs %q, want `xeno cost turn` on the path: a project receives the "+
			"runner as a binary, not as this repository's own build", cmd)
	}
}

// Both manifests are read rather than described, because the client loads them and a field this
// test invented would pass while the client refused the plugin.
func TestTheManifestsSayWhatTheClientReads(t *testing.T) {
	var plugin struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Version     string `json:"version"`
	}
	body := read(t, filepath.Join(repoRoot, pluginDir, ".claude-plugin/plugin.json"))
	if err := json.Unmarshal([]byte(body), &plugin); err != nil {
		t.Fatal(err)
	}
	if plugin.Name != "xeno" || plugin.Version == "" || len(strings.Fields(plugin.Description)) < 10 {
		t.Errorf("the plugin manifest is %+v", plugin)
	}

	var market struct {
		Name    string `json:"name"`
		Plugins []struct {
			Name   string `json:"name"`
			Source string `json:"source"`
		} `json:"plugins"`
	}
	body = read(t, filepath.Join(repoRoot, ".claude-plugin/marketplace.json"))
	if err := json.Unmarshal([]byte(body), &market); err != nil {
		t.Fatal(err)
	}
	if len(market.Plugins) != 1 || market.Plugins[0].Name != "xeno" {
		t.Fatalf("the marketplace declares %+v", market.Plugins)
	}
	// The source is the distribution tree inside this repository, which is where the manifest the
	// client loads actually sits (A77).
	if market.Plugins[0].Source != "./"+pluginDir {
		t.Errorf("the marketplace points at %q, want ./%s", market.Plugins[0].Source, pluginDir)
	}
}
