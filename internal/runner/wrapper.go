// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"fmt"
	"sort"
	"strings"

	"github.com/triplem/xeno/internal/scaffold"
)

// WrapperHost is what a host contributes to the generated wrapper: where the file goes,
// and how that host names the two ends of the commit range.
//
// Those two expressions are the parameters WP10 asks the generator to take from the
// start although there is one wrapper to generate. They are the only host specific
// thing in it, so another host is an entry in this table rather than a second generator.
type WrapperHost struct {
	ID      string
	Path    string
	BaseRef string
	HeadRef string
	// Adapter and APIBase are what project.yaml's tracker block is written with. They are
	// here rather than in a table of their own because the range expressions and the
	// address are the same fact about one host arriving at two generated files, and two
	// tables holding one fact is how the two files come to disagree.
	Adapter string
	APIBase string
	// Squash is the setting the commit convention rests on, in this host's own words. A
	// squash rewrites the message, so a Refs or Closes footer survives only where the
	// template carries the description through; CONTRIBUTING.md records the dependency.
	// It is a row rather than a branch on the host's name, so a third host adds no code.
	Squash string
}

var wrapperHosts = map[string]WrapperHost{
	"github": {
		ID:      "github",
		Path:    ".github/workflows/xeno-gate.yml",
		BaseRef: "${{ github.event.pull_request.base.sha }}",
		HeadRef: "${{ github.event.pull_request.head.sha }}",
		Adapter: "github",
		APIBase: "https://api.github.com",
		Squash: "set the pull request squash commit message to the description, " +
			"or a Refs or Closes footer does not survive a squash merge",
	},
	"gitlab": {
		ID:   "gitlab",
		Path: ".gitlab-ci.yml",
		// CI_MERGE_REQUEST_DIFF_BASE_SHA exists only in a merge request pipeline, and only
		// while the merge request is open, which is why the generated job is scoped to
		// merge_request_event rather than matching the other host's trigger for symmetry.
		BaseRef: "$CI_MERGE_REQUEST_DIFF_BASE_SHA",
		HeadRef: "$CI_COMMIT_SHA",
		Adapter: "gitlab",
		APIBase: "https://gitlab.com/api/v4",
		Squash: "set the squash commit message template to contain %{description}, " +
			"or a Refs or Closes footer does not survive a squash merge",
	},
}

// WrapperHosts lists what can be generated, for an error that helps. It is derived from the
// table and sorted: a literal list beside it was correct for as long as there was one host,
// and correct is how a duplicate survives being noticed.
func WrapperHosts() []string {
	out := make([]string, 0, len(wrapperHosts))
	for h := range wrapperHosts {
		out = append(out, h)
	}
	sort.Strings(out)
	return out
}

// Wrapper generates the CI wrapper for a host, from the scaffold of that host: the body
// is a file rather than a string literal, so a repository can replace it without forking
// the runner, and the generated file says which copy it came from.
//
// It calls one command and does nothing else. A wrapper that built or tested would be
// this project's own pipeline, which is a different file with the opposite purpose, and
// mixing them up puts build steps into a template or a gate call into a pipeline.
func Wrapper(root, host, runnerVersion string) (WrapperHost, string, error) {
	if strings.TrimSpace(host) == "" {
		var h WrapperHost
		return h, "", fmt.Errorf("no --host; there is %s.\n"+
			"It is not defaulted: the host decides the wrapper and the tracker block both, and a "+
			"tool that picks one is a tool with one host", strings.Join(WrapperHosts(), ", "))
	}
	h, ok := wrapperHosts[host]
	if !ok {
		return h, "", fmt.Errorf("no wrapper for host %q; there is %s", host, strings.Join(WrapperHosts(), ", "))
	}
	body, _, err := scaffold.RenderWrapper(root, host, scaffold.Wrapper{
		RunnerVersion: runnerVersion, BaseRef: h.BaseRef, HeadRef: h.HeadRef,
	})
	return h, body, err
}
