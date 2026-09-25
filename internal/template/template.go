// SPDX-License-Identifier: Apache-2.0

// Package template renders a phase result from a structure and a strings bundle.
//
// The two are separate so that adding a language is a bundle and never a change to the
// structure: template.yaml holds section ids, their order and which are required, and
// the bundle holds the headings. Section ids are part of the process layer and are
// always English, whatever the artifacts are written in.
package template

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/hashing"
)

// AnchorPrefix opens the HTML comment that marks a section. An anchor stands on its own
// line before the heading and there are no end markers: a section runs from its anchor
// to the next one or to the end of the file. The agent never writes one, which is what
// makes them impossible to get wrong.
const AnchorPrefix = "<!-- xeno:section:"

// Template is template.yaml. Order is the order of the slice.
type Template struct {
	ID       string    `yaml:"id"`
	Version  string    `yaml:"version"`
	Phase    string    `yaml:"phase"`
	Sections []Section `yaml:"sections"`
}

type Section struct {
	ID string `yaml:"id"`
	// Required sections are rendered whether or not they carry content, because a
	// missing one is the finding. An optional section is rendered only where there is
	// something in it: an absent one is fine and a present empty one is not, and empty
	// headings accumulate.
	Required bool `yaml:"required"`
}

// Bundle is strings.<lang>.yaml. It carries what a reader sees and nothing a gate reads.
type Bundle struct {
	Language string            `yaml:"language"`
	Title    string            `yaml:"title"`
	Headings map[string]string `yaml:"headings"`
}

// Source says which of the two resolution paths a template came from. It is recorded in
// context.lock.yaml, because otherwise two projects on the same template version are
// indistinguishable although one of them overrode it.
type Source string

const (
	FromPlugin  Source = "plugin"
	FromProject Source = "project"
)

const (
	pluginDir  = ".xeno/plugin/templates"
	projectDir = ".xeno/config/templates"
)

// Resolved is a template together with the bundle and where they were found.
type Resolved struct {
	Template    Template
	Bundle      Bundle
	Source      Source
	StringsHash string
}

// Ref is what an artifact records in its `template` field: id@version.
func (r Resolved) Ref() string { return r.Template.ID + "@" + r.Template.Version }

// Load resolves one template id in one language. Project beats plugin, per id, so a
// project replaces a single template without forking the set.
//
// A missing bundle is an error and never a fall back to another language. Falling back
// would render an artifact in a language nobody asked for and record that it was
// rendered in the one they did.
func Load(root, id, language string) (Resolved, error) {
	var r Resolved
	dir, src := "", FromPlugin
	for _, c := range []struct {
		path string
		src  Source
	}{{filepath.Join(root, projectDir, id), FromProject}, {filepath.Join(root, pluginDir, id), FromPlugin}} {
		if fm.Exists(filepath.Join(c.path, "template.yaml")) {
			dir, src = c.path, c.src
			break
		}
	}
	if dir == "" {
		return r, fmt.Errorf("no template %q under %s or %s", id, projectDir, pluginDir)
	}
	if err := fm.ReadYAML(filepath.Join(dir, "template.yaml"), &r.Template); err != nil {
		return r, fmt.Errorf("%s/template.yaml: %w", dir, err)
	}
	bundlePath := filepath.Join(dir, "strings."+language+".yaml")
	if err := fm.ReadYAML(bundlePath, &r.Bundle); err != nil {
		return r, fmt.Errorf("no strings bundle for %q in template %q: %w", language, id, err)
	}
	b, err := os.ReadFile(bundlePath)
	if err != nil {
		return r, err
	}
	r.StringsHash = hashing.Hex(hashing.Normalise(b))
	r.Source = src
	for _, s := range r.Template.Sections {
		if _, ok := r.Bundle.Headings[s.ID]; !ok {
			return r, fmt.Errorf("bundle %s has no heading for section %q", bundlePath, s.ID)
		}
	}
	return r, nil
}

// Render writes the body of a phase result: the title, then every section the template
// asks for, in the order it defines, each preceded by its anchor.
//
// What is rendered does not depend on the language: the same content in another bundle
// produces the same sections with the same ids in the same order, and only the headings
// move. That is what lets an artifact be re-rendered without changing what it asserts.
func (r Resolved) Render(content map[string]string) string {
	var b strings.Builder
	if r.Bundle.Title != "" {
		fmt.Fprintf(&b, "# %s\n", r.Bundle.Title)
	}
	for _, s := range r.Template.Sections {
		text := strings.TrimSpace(content[s.ID])
		if text == "" && !s.Required {
			continue
		}
		fmt.Fprintf(&b, "\n%s%s -->\n## %s\n", AnchorPrefix, s.ID, r.Bundle.Headings[s.ID])
		if text != "" {
			fmt.Fprintf(&b, "\n%s\n", text)
		}
	}
	return b.String()
}

// Parse reads a rendered body back into its sections, so that setting one section
// re-renders the rest unchanged. Anything before the first anchor is the title and is
// dropped, since the bundle owns it.
func Parse(body string) map[string]string {
	out := map[string]string{}
	id := ""
	var buf []string
	flush := func() {
		if id != "" {
			out[id] = strings.TrimSpace(strings.Join(buf, "\n"))
		}
		buf = nil
	}
	for _, line := range strings.Split(body, "\n") {
		if t := strings.TrimSpace(line); strings.HasPrefix(t, AnchorPrefix) && strings.HasSuffix(t, "-->") {
			flush()
			id = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(t, AnchorPrefix), "-->"))
			continue
		}
		if id == "" {
			continue // the title, which the bundle owns
		}
		if strings.HasPrefix(strings.TrimSpace(line), "## ") && len(buf) == 0 {
			continue // the heading, which the bundle owns
		}
		buf = append(buf, line)
	}
	flush()
	return out
}

// Missing lists the required sections that carry nothing, in template order.
func (r Resolved) Missing(content map[string]string) []string {
	var out []string
	for _, s := range r.Template.Sections {
		if s.Required && strings.TrimSpace(content[s.ID]) == "" {
			out = append(out, s.ID)
		}
	}
	return out
}

// Known lists every section id the template defines, sorted, for an error that helps.
func (r Resolved) Known() []string {
	out := make([]string, 0, len(r.Template.Sections))
	for _, s := range r.Template.Sections {
		out = append(out, s.ID)
	}
	sort.Strings(out)
	return out
}

// Has says whether the template defines a section.
func (r Resolved) Has(id string) bool {
	for _, s := range r.Template.Sections {
		if s.ID == id {
			return true
		}
	}
	return false
}
