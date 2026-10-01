// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Shipped patterns, named. A rule names one and carries no expression of its own,
// because a regex in a rule file would be code without provenance, which is the line
// external gates exist to draw.
//
// The same map answers at the gate and in `xeno check commit-message`, which is what a
// hook calls. One implementation decides in both places, so a push cannot start
// rejecting what a gate accepts.
var patterns = map[string]*regexp.Regexp{
	// Conventional Commits, subject only: type, optional scope, optional breaking
	// marker, colon, space, description. The types are the Angular set.
	"conventional-commits": regexp.MustCompile(
		`^(feat|fix|docs|style|refactor|perf|test|build|ci|chore|revert)(\([^()]+\))?!?: .+$`),
	// The same form carrying an issue reference in the subject, which is what a project
	// needs when the reference has to survive a squash merge: the squashed message is built
	// from the merge request title, and a footer written on a branch commit does not reach
	// it. Both hosts' spellings are accepted, GitHub's (#123) and GitLab's (!123), because a
	// shipped pattern that held for one host would be a pattern per host (A71).
	"conventional-commits-with-issue": regexp.MustCompile(
		`^(feat|fix|docs|style|refactor|perf|test|build|ci|chore|revert)(\([^()]+\))?!?: .+ \([#!][0-9]+\)$`),
}

// PatternNames lists what is shipped, for an error that helps rather than refuses.
func PatternNames() []string {
	names := make([]string, 0, len(patterns))
	for n := range patterns {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// CheckMessage evaluates one message against a named pattern. It reads the subject and
// nothing else: a body and its trailers are not what any of these patterns describe.
func CheckMessage(pattern, message string) error {
	re, ok := patterns[pattern]
	if !ok {
		return fmt.Errorf("no shipped pattern %q; there is %s", pattern, strings.Join(PatternNames(), ", "))
	}
	subject := message
	if i := strings.IndexByte(subject, '\n'); i >= 0 {
		subject = subject[:i]
	}
	subject = strings.TrimRight(subject, "\r")
	if subject == "" {
		return fmt.Errorf("the message has no subject")
	}
	if !re.MatchString(subject) {
		return fmt.Errorf("subject does not match %s:\n  %s", pattern, subject)
	}
	return nil
}
