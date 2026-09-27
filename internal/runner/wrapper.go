// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"fmt"
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
}

var wrapperHosts = map[string]WrapperHost{
	"github": {
		ID:      "github",
		Path:    ".github/workflows/xeno-gate.yml",
		BaseRef: "${{ github.event.pull_request.base.sha }}",
		HeadRef: "${{ github.event.pull_request.head.sha }}",
	},
}

// WrapperHosts lists what can be generated, for an error that helps.
func WrapperHosts() []string { return []string{"github"} }

// Wrapper generates the CI wrapper for a host, from the scaffold of that host: the body
// is a file rather than a string literal, so a repository can replace it without forking
// the runner, and the generated file says which copy it came from.
//
// It calls one command and does nothing else. A wrapper that built or tested would be
// this project's own pipeline, which is a different file with the opposite purpose, and
// mixing them up puts build steps into a template or a gate call into a pipeline.
func Wrapper(root, host, runnerVersion string) (WrapperHost, string, error) {
	h, ok := wrapperHosts[host]
	if !ok {
		return h, "", fmt.Errorf("no wrapper for host %q; there is %s", host, strings.Join(WrapperHosts(), ", "))
	}
	body, _, err := scaffold.RenderWrapper(root, host, scaffold.Wrapper{
		RunnerVersion: runnerVersion, BaseRef: h.BaseRef, HeadRef: h.HeadRef,
	})
	return h, body, err
}
