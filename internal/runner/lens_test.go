// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/plugin"
)

// WP11's done-when: "a disabled lens changes nothing but the findings". These are what makes that
// clause checkable, which it was not while the four lenses did not exist — enabling and disabling
// nothing changes nothing, and the clause was not false but untestable (#286).

// lensFixture is a repository whose vendored plugin carries one lens declaring the first phase,
// and whose project.yaml enables what the caller passes. The lens is vendored in both the enabled
// and the disabled fixture on purpose: the plugin tree, and therefore the hash G-Supply compares
// and the block the lock records, has to be the same in the two, or the comparison below would be
// about the plugin rather than about the enablement.
func lensFixture(t *testing.T, enabled string) *fixture {
	f := newFixture(t)
	f.write(plugin.Dir+"/skills/"+plugin.LensPrefix+"security/SKILL.md",
		"---\nname: "+plugin.LensPrefix+"security\ndescription: the security lens, for the fixture\n"+
			"phases: ["+model.Phases[0]+"]\n---\n\n# Security lens\n")
	f.write(model.ProjectFile, "runner_version: "+model.RunnerVersion+"\nlenses:\n  enabled: ["+enabled+"]\n")
	return f
}

// phaseBytes is every file of a phase directory with its content, which is what "changes nothing"
// has to mean: the lock, the artifact, the digest, the learning record and the verdict, compared by
// content rather than by the hash one of them carries. A comparison of artifacts_hash alone would
// miss anything written beside the hashed set.
func phaseBytes(t *testing.T, root, phase string) string {
	t.Helper()
	dir := filepath.Join(root, model.PhaseDir(key, phase))
	var b strings.Builder
	err := filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(dir, p)
		content, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		b.WriteString(rel + ":\n" + string(content) + "\n")
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return b.String()
}

func TestADisabledLensChangesNothingButTheFindings(t *testing.T) {
	phase := model.Phases[0]
	with, without := lensFixture(t, "security"), lensFixture(t, "")

	// The one thing enablement decides: which lenses the phase works under.
	applying, unknown := with.r.Lenses(phase)
	if len(applying) != 1 || applying[0].ID != "security" {
		t.Fatalf("with security enabled, %v applies at %s", applying, phase)
	}
	if len(unknown) != 0 {
		t.Errorf("the enabled lens was not resolved: %v", unknown)
	}
	if off, _ := without.r.Lenses(phase); len(off) != 0 {
		t.Fatalf("with none enabled, %v applies at %s", off, phase)
	}

	// And everything the phase writes is the same, because every finding a lens produces is
	// written by a command the phase already has. Nothing in section 5's field list records which
	// lenses ran, so no artifact, no hash and no verdict may move with the block.
	a, b := with.run(phase, ""), without.run(phase, "")
	if a.Status != b.Status {
		t.Errorf("the verdict is %s with the lens and %s without it", a.Status, b.Status)
	}
	if x, y := phaseBytes(t, with.root, phase), phaseBytes(t, without.root, phase); x != y {
		t.Errorf("enabling a lens changed what the phase wrote:\n--- with\n%s\n--- without\n%s", x, y)
	}
}

// A lens applies where it declared it applies, and enablement cannot widen that: section 13's cost
// is paid per phase the lens names, not per phase the project has.
func TestAnEnabledLensAppliesOnlyToThePhasesItDeclared(t *testing.T) {
	f := lensFixture(t, "security")
	if applying, _ := f.r.Lenses(model.Phases[0]); len(applying) != 1 {
		t.Fatalf("the lens does not apply to the phase it declared: %v", applying)
	}
	for _, phase := range model.Phases[1:] {
		if applying, _ := f.r.Lenses(phase); len(applying) != 0 {
			t.Errorf("the lens applies at %s, which it did not declare", phase)
		}
	}
}

// A name the plugin does not answer to is reported, because the alternative is the failure that
// looks like success: the project asked for a lens, every phase ran, every gate was green, and
// nothing was ever loaded.
func TestALensNameNothingAnswersToIsReported(t *testing.T) {
	f := lensFixture(t, "security, cryptography")
	applying, unknown := f.r.Lenses(model.Phases[0])
	if len(applying) != 1 || applying[0].ID != "security" {
		t.Errorf("the lens that exists did not apply: %v", applying)
	}
	if len(unknown) != 1 || unknown[0] != "cryptography" {
		t.Errorf("unknown lenses are %v, want the one the plugin has no skill for", unknown)
	}
}

// An absent block is Appendix A's default, and it is the state every repository is in until it
// decides otherwise. It has to be distinguishable from nothing at all only in the file, not here:
// both mean no lens runs.
func TestWithNoLensBlockNoLensRuns(t *testing.T) {
	f := newFixture(t)
	f.write(model.ProjectFile, "runner_version: "+model.RunnerVersion+"\n")
	if applying, unknown := f.r.Lenses(model.Phases[0]); len(applying) != 0 || len(unknown) != 0 {
		t.Errorf("a project.yaml with no lens block ran %v and could not resolve %v", applying, unknown)
	}
}
