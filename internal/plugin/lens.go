// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/triplem/xeno/internal/fm"
)

// LensPrefix is how section 13's tree spells a lens: the four skills are named
// xeno-lens-security, xeno-lens-privacy, xeno-lens-operations and
// xeno-lens-architecture, beside the seven whose names are the phases and the learning
// record. The prefix is what tells the two kinds apart, and it is read from the directory
// rather than from a list here, so that a project which overrides or adds one is found by
// the same walk.
const LensPrefix = "xeno-lens-"

// skillsSubdir is the directory section 13 puts every skill in, inside the vendored plugin.
const skillsSubdir = "skills"

// Lens is one lens of section 13: a skill a phase pulls in, with the phases it declares it
// applies to. ID is the name without the prefix, which is what `lenses.enabled` in
// project.yaml names — Appendix A writes `enabled: [security, privacy]` and not the skill
// names, and the pair is resolved here rather than in each reader.
//
// Phases is the declaration, not a default. Section 13 says a lens "declares which phases
// it applies to", so a lens whose frontmatter names none applies nowhere: a lens that
// silently applied everywhere would be the four-lenses-in-every-request cost the same
// section warns about, arrived at by omission.
type Lens struct {
	ID     string
	Skill  string
	Phases []string
}

// AppliesTo reports whether the lens declared this phase.
func (l Lens) AppliesTo(phase string) bool {
	for _, p := range l.Phases {
		if p == phase {
			return true
		}
	}
	return false
}

// Lenses is every lens the vendored plugin carries, sorted by id, and empty where there is
// no plugin or no lens in it.
//
// Empty rather than an error. A repository with no vendored plugin runs its phases without
// lenses exactly as a project that enabled none does, and the two cases are the same thing
// from here: nothing is pulled in. What differs is visible one level up, where the project's
// `lenses.enabled` is read and a name that resolves to nothing is reported.
func Lenses(root string) []Lens {
	dir := filepath.Join(root, Dir, skillsSubdir)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var out []Lens
	for _, e := range entries {
		if !e.IsDir() || !strings.HasPrefix(e.Name(), LensPrefix) {
			continue
		}
		var front struct {
			Name   string   `yaml:"name"`
			Phases []string `yaml:"phases"`
		}
		if _, err := fm.ReadFront(filepath.Join(dir, e.Name(), "SKILL.md"), &front); err != nil {
			continue
		}
		skill := front.Name
		if skill == "" {
			skill = e.Name()
		}
		out = append(out, Lens{
			ID:     strings.TrimPrefix(e.Name(), LensPrefix),
			Skill:  skill,
			Phases: front.Phases,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// LensesFor is the lenses a phase actually works under: enabled by the project and
// declaring this phase. Sorted by id, so that two runs over one tree report the same order
// whatever order the project listed them in.
//
// It is a selection and nothing else. Section 13 fixes what a lens may produce — an open
// question, an assumption, a review checklist entry — and all three are written by the
// phase's own commands, so enabling a lens changes what the phase finds and no field, hash
// or verdict anywhere. That is why nothing records the selection: there is no field for it
// in section 5's list, and the one thing a record here could do is make a cost decision
// look like part of the trail.
func LensesFor(root string, enabled []string, phase string) []Lens {
	if len(enabled) == 0 {
		return nil
	}
	want := make(map[string]bool, len(enabled))
	for _, name := range enabled {
		want[strings.TrimSpace(name)] = true
	}
	var out []Lens
	for _, l := range Lenses(root) {
		if want[l.ID] && l.AppliesTo(phase) {
			out = append(out, l)
		}
	}
	return out
}

// UnknownLenses are the names a project enabled that no lens in the vendored plugin
// answers to, sorted.
//
// Reported rather than ignored, because the failure mode of a silent miss is the one this
// whole block is hard to notice in: a project that misspells a name gets exactly what a
// project that enabled nothing gets — every phase running, every gate green, and the lens
// it asked for never loaded.
func UnknownLenses(root string, enabled []string) []string {
	have := make(map[string]bool)
	for _, l := range Lenses(root) {
		have[l.ID] = true
	}
	var out []string
	for _, name := range enabled {
		if n := strings.TrimSpace(name); n != "" && !have[n] {
			out = append(out, n)
		}
	}
	sort.Strings(out)
	return out
}
