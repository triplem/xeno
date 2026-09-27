// SPDX-License-Identifier: Apache-2.0

// Package scaffold holds the files xeno init writes into a repository: the CI wrapper
// and the initial project.yaml. They are starting points somebody edits afterwards,
// which is what separates them from the templates in internal/template: a template
// renders an artifact a gate then judges, and carries a version and a strings hash for
// that reason. A scaffold carries neither, and sharing the resolution code would invite
// somebody to give the wrapper a version it has no use for.
//
// The defaults are embedded rather than read from the vendored plugin. xeno init is
// what creates .xeno/ and what vendors the plugin, so a file it needs in order to run
// cannot live in what it produces; G-Supply makes the same point from the other side,
// since the digest it checks belongs to a plugin that is not there yet.
package scaffold

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed files
var files embed.FS

// OverrideDir is where a repository puts its own version of a scaffold. It is read
// after the first init, which is too late for the first wrapper and right for every
// one after, and it is how an organisation applies its defaults without forking the
// runner.
const OverrideDir = ".xeno/config/scaffold"

// Source says where a rendered scaffold came from. The generated file names it, because
// the first question anybody asks of a generated file that does not look like they
// expect is which copy produced it. `context.lock.yaml` records `template_source` for
// the same reason.
type Source string

const (
	FromRunner  Source = "the runner's own default"
	FromProject Source = "the project's override under " + OverrideDir
)

// Wrapper is what the CI wrapper needs. The two refs are the host's expressions for the
// ends of the commit range, which WP10 asks the generator to take as parameters.
type Wrapper struct {
	RunnerVersion string
	BaseRef       string
	HeadRef       string
	Source        Source
}

// Project is what the initial project.yaml needs, which is the three answers xeno init
// asks for plus the version it pins.
type Project struct {
	RunnerVersion  string
	Language       string
	Model          string
	TrackerProject string
	Source         Source
}

// RenderWrapper renders the CI wrapper of one host.
func RenderWrapper(root, host string, d Wrapper) (string, Source, error) {
	return render(root, "ci-"+host+".yml", &d, func(s Source) { d.Source = s })
}

// RenderProject renders the initial project.yaml.
func RenderProject(root string, d Project) (string, Source, error) {
	return render(root, "project.yaml", &d, func(s Source) { d.Source = s })
}

// render resolves a scaffold, project override first, and executes it. The source is
// set on the data before execution so a template can name it: the one field every
// scaffold has in common.
func render(root, name string, data any, setSource func(Source)) (string, Source, error) {
	body, source, err := resolve(root, name)
	if err != nil {
		return "", source, err
	}
	setSource(source)
	t, err := template.New(name).Parse(string(body))
	if err != nil {
		return "", source, fmt.Errorf("scaffold %s: %w", name, err)
	}
	var out bytes.Buffer
	if err := t.Execute(&out, data); err != nil {
		return "", source, fmt.Errorf("scaffold %s: %w", name, err)
	}
	return out.String(), source, nil
}

func resolve(root, name string) ([]byte, Source, error) {
	if root != "" {
		b, err := os.ReadFile(filepath.Join(root, OverrideDir, name))
		if err == nil {
			return b, FromProject, nil
		}
	}
	b, err := files.ReadFile("files/" + name)
	if err != nil {
		return nil, FromRunner, fmt.Errorf("no scaffold %q", name)
	}
	return b, FromRunner, nil
}

// Names lists what can be rendered, for an error that helps.
func Names() []string {
	entries, err := files.ReadDir("files")
	if err != nil {
		return nil
	}
	var out []string
	for _, e := range entries {
		out = append(out, e.Name())
	}
	return out
}
