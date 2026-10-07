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
// Three ports and one function, which is the whole of section 12's contract: the branch
// rules `xeno enforcement check` asks about, the issue a phase start reads, the comment a
// pipeline writes, and the credential all of them need. The first is the one that needed a
// decision about vocabulary; the other two answer in the words the host used, because a
// title and a body are the same two things on every host and there is nothing in them for
// the domain to judge.
//
// Nothing under internal/gates imports this. A verdict has to be reproducible from the
// repository alone, and A42 holds that property with a check in CI rather than a test.
package host

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/triplem/xeno/internal/enforcement"
	"github.com/triplem/xeno/internal/host/github"
	"github.com/triplem/xeno/internal/host/gitlab"
	"github.com/triplem/xeno/internal/model"
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

// Issues is section 12's second operation: read an issue. The key is the part after the
// `#` of a qualified intent id, which is what the host numbers its issues by; splitting a
// qualified id is the caller's work, because the adapter "knows no relationship between
// issue key and repository" and the project is therefore a parameter beside it.
//
// The answer is model.Issue rather than a type of this package's, and the same holds for
// the comment below. The reason is the import graph: this package selects the adapters, so
// an adapter cannot import it, and the vocabulary the two share therefore belongs where
// section 12's tracker block already is. The branch rules port answers in
// enforcement.Requirement for exactly the same reason.
type Issues interface {
	Issue(project, key, token string) (model.Issue, error)
}

// Comments is section 12's third operation: write a comment. One call and no read before
// it, which is what section 12 means by "it costs one call from a job that already runs";
// finding and editing a previous comment would be two, and the trail in the repository is
// the record either way.
type Comments interface {
	Comment(project, key, token, body string) (model.Written, error)
}

// Host is the three operations an adapter implements, as one interface, so that the claim
// this package makes holds for the whole of section 12's contract and not only for the
// branch rules port: a second host is an entry in the table below and a package beside
// github, and neither the ports nor the domain changes for it.
//
// The fourth operation, resolving credentials, is Credentials below. It is not a method
// because it reads the environment and no host: the variable is named by
// `tracker.auth.secret_env` and what it holds is the same string whichever host it is for.
type Host interface {
	BranchRules
	Issues
	Comments
}

// adapters is the adapters that exist, by the name project.yaml selects them with.
var adapters = map[string]func(*http.Client, string) Host{
	"github": func(c *http.Client, baseURL string) Host {
		return github.New(c, baseURL)
	},
	"gitlab": func(c *http.Client, baseURL string) Host {
		return gitlab.New(c, baseURL)
	},
}

// Adapters lists what can be selected, for an error that helps.
func Adapters() []string {
	names := make([]string, 0, len(adapters))
	for n := range adapters {
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
	return For(model.Tracker{Adapter: adapter, BaseURL: baseURL}, c)
}

// For selects the adapter a project configured, for all three operations.
//
// The two refusals are the ones BranchRulesFor has always made, and they are here rather
// than in each caller because Appendix A gives the tracker block one rule for all of them:
// a block that is there and incomplete makes `xeno phase start` refuse rather than guess,
// and the refusal names the field that is missing.
func For(t model.Tracker, c *http.Client) (Host, error) {
	if strings.TrimSpace(t.Adapter) == "" {
		return nil, fmt.Errorf("no tracker.adapter in .xeno/config/project.yaml; there is %s",
			strings.Join(Adapters(), ", "))
	}
	newAdapter, ok := adapters[t.Adapter]
	if !ok {
		return nil, fmt.Errorf("no adapter %q; there is %s", t.Adapter,
			strings.Join(Adapters(), ", "))
	}
	if strings.TrimSpace(t.BaseURL) == "" {
		return nil, fmt.Errorf(
			"no tracker.base_url in .xeno/config/project.yaml for adapter %q; "+
				"a self managed deployment has no canonical address, so the address is a "+
				"setting rather than a constant", t.Adapter)
	}
	return newAdapter(c, t.BaseURL), nil
}

// Credentials is section 12's fourth operation: resolve the tracker token. It comes from
// the environment and never from project.yaml, which Appendix A requires of the file and
// the implementation plan's credential table requires of this token in particular.
//
// The variable is the one `tracker.auth.secret_env` names and no other. There is no
// fallback to a variable a pipeline happens to provide: the enforcement token has one
// because A40 decided it, and the reason given there — that a scheduled job needs no
// secret of its own — is about a read of repository settings and not about a write that
// appears under somebody's name on an issue.
//
// An empty variable is an error and not an absence. The plan's table says what expiry
// costs, "P0 and every comment fail with exit code 2", which is the staircase separating a
// credential nobody renewed from a tracker nobody configured.
// ErrNoToken is what Credentials wraps where the variable naming the token is empty. The
// configuration is right and the machine simply holds no token, which is a different thing
// from a block that names no variable at all, and the two have different answers.
var ErrNoToken = errors.New("no token in the environment")

func Credentials(a model.Auth) (string, error) {
	name := strings.TrimSpace(a.SecretEnv)
	if name == "" {
		return "", fmt.Errorf("no tracker.auth.secret_env in .xeno/config/project.yaml, so " +
			"there is no variable to read the token from; the file never holds the token itself")
	}
	if s := strings.TrimSpace(a.Scheme); s != "" && s != "token" {
		return "", fmt.Errorf("tracker.auth.scheme is %q, and token is the only scheme either "+
			"adapter implements; a request sent under the wrong scheme fails at the host", s)
	}
	v := os.Getenv(name)
	if v == "" {
		// Wrapped in ErrNoToken, because an empty variable is not a misconfiguration: a
		// machine with no token is the ordinary state of a clone somebody is reading, and
		// what a caller does about it depends on what it wanted the token for. Reading an
		// issue carries on without one; writing a comment cannot.
		return "", fmt.Errorf("%w: %s is empty, and it is the variable "+
			"tracker.auth.secret_env names; set it to a token that may read an issue and "+
			"write a note", ErrNoToken, name)
	}
	return v, nil
}
