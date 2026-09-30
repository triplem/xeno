// SPDX-License-Identifier: Apache-2.0

// Package host holds the ports a code host is reached through, and selects the adapter a
// project configured. It is the boundary #97 decided and A65 fixed the vocabulary of.
//
// An adapter answers with requirements rather than with a struct of host neutral facts.
// The reason is in A65 and is worth restating where the interface is: GitHub answers
// allow_bypass with one boolean and GitLab with three lists of grants, and reducing grants
// to "somebody may bypass" is a judgement about one host's access levels. A neutral struct
// has to hold that judgement in the domain, which is the one package meant to know no
// host, or grow a field shaped like one host's answer.
//
// Nothing under internal/gates imports this. A verdict has to be reproducible from the
// repository alone, and A42 holds that property with a check in CI rather than a test.
package host

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/triplem/xeno/internal/enforcement"
	"github.com/triplem/xeno/internal/host/github"
	"github.com/triplem/xeno/internal/host/gitlab"
)

// BranchRules is what xeno enforcement check asks of a host: what the branch is configured
// to require, expressed as the requirements a project declares.
//
// Declared is a parameter because the adapter decides met from unmet, and some of those
// decisions need the declared value: approvals.required is met against a count. The
// alternative, an adapter that reports what it found for the domain to judge, is the
// reading A65 rejects.
//
// An implementation returns Met, Unmet or NotAvailable and never Waived. A waiver is a
// decision the project took and recorded in project.yaml, so it is applied where the
// declaration is read and not where the host is asked.
//
// The names in the returned requirements are enforcement's constants. A host chooses its
// own words for what it found, never its own name for what was asked.
type BranchRules interface {
	Requirements(repo, branch, token string, d enforcement.Declared) (
		[]enforcement.Requirement, error)
}

// branchRules is the adapters that exist, by the name project.yaml selects them with.
// A second host is an entry here and a package beside github, and neither the port nor
// the domain changes for it: that is the claim this package makes, and the one that is
// tested the day the second entry arrives.
var branchRules = map[string]func(*http.Client, string) BranchRules{
	"github": func(c *http.Client, baseURL string) BranchRules {
		return github.New(c, baseURL)
	},
	"gitlab": func(c *http.Client, baseURL string) BranchRules {
		return gitlab.New(c, baseURL)
	},
}

// Adapters lists what can be selected, for an error that helps.
func Adapters() []string {
	names := make([]string, 0, len(branchRules))
	for n := range branchRules {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// BranchRulesFor selects the adapter a project configured.
//
// There is no default. A missing adapter name is a refusal naming the field rather than an
// assumption about which host this is: a tool that guesses the host is a tool that only ever
// had one, and the guess is invisible for exactly as long as there is only one.
func BranchRulesFor(adapter, baseURL string, c *http.Client) (BranchRules, error) {
	if strings.TrimSpace(adapter) == "" {
		return nil, fmt.Errorf("no tracker.adapter in .xeno/config/project.yaml; there is %s",
			strings.Join(Adapters(), ", "))
	}
	newAdapter, ok := branchRules[adapter]
	if !ok {
		return nil, fmt.Errorf("no adapter %q; there is %s", adapter,
			strings.Join(Adapters(), ", "))
	}
	if strings.TrimSpace(baseURL) == "" {
		return nil, fmt.Errorf(
			"no tracker.base_url in .xeno/config/project.yaml for adapter %q; "+
				"a self managed deployment has no canonical address, so the address is a "+
				"setting rather than a constant", adapter)
	}
	return newAdapter(c, baseURL), nil
}
