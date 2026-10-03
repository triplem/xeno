// SPDX-License-Identifier: Apache-2.0

// Package git reads the commit range under review. It is the only place in this repository
// that starts a subprocess, and it starts two: `git log` over the range the run was given,
// and `git diff` over the two trees at its ends.
//
// It exists as a package rather than as a few lines inside a gate because four of section 9's
// predicate types read a commit, and process handling in the middle of a verdict would make
// the prohibition in internal/gates untestable. The arguments are fixed here; the only values
// that come from outside are the two refs, which the run supplies and never infers.
//
// Nothing here reaches the network. A clone has its history locally, which is what makes a
// commit predicate available to a gate that may not have a network at all.
package git

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// Commit is what the predicates of section 9 read: the subject for commit-message, the
// trailers for commit-trailer, the signature status for commit-signature, the author for
// approver-not-author, and the parent count so that a merge commit can be exempted.
type Commit struct {
	Hash      string
	Subject   string
	Author    string
	Email     string
	Signature string // git's %G? in one letter, see Signed
	Trailers  []string
	Parents   int
}

// Signed reports whether the commit carries a signature that verifies.
//
// Git answers with one of several letters and the rule needs a boolean. G is a good signature
// and U is a good signature from a key the local configuration does not trust; both are
// accepted, because trust lives in the verifier's keyring rather than in the commit, and a
// verdict that differed between two clones of one tree is the failure the whole gate path
// avoids. B, X, Y, R, E and N — bad, expired, expired key, revoked key, missing key and
// unsigned — are not signed for this purpose (A70).
func (c Commit) Signed() bool { return c.Signature == "G" || c.Signature == "U" }

// IsMerge reports whether the commit has more than one parent, which is what `exempt:
// [merge-commits]` leaves out of a range.
func (c Commit) IsMerge() bool { return c.Parents > 1 }

// HasTrailer reports whether the commit carries a trailer with this key, compared case
// insensitively as git itself reads trailers. The value is not examined: section 9 asks for a
// required trailer and nothing more, and checking what it says would need an expression, which
// is the line the whole section draws.
func (c Commit) HasTrailer(key string) bool {
	want := strings.ToLower(strings.TrimSuffix(strings.TrimSpace(key), ":"))
	for _, t := range c.Trailers {
		k, _, ok := strings.Cut(t, ":")
		if ok && strings.ToLower(strings.TrimSpace(k)) == want {
			return true
		}
	}
	return false
}

// The field and record separators: the ASCII unit and record separators, which no commit
// message carries in practice and which an editor cannot type by accident.
//
// Not NUL, although NUL is the separator git itself offers for machine reading. The format
// string is an argument to the process, and an argument containing NUL is rejected by exec
// before git ever runs — "fork/exec: invalid argument", which says nothing about the cause.
const (
	fieldSep  = "\x1f"
	recordSep = "\x1e"
)

// format asks for one record per commit. %(trailers) comes last because it is the only field
// that spans lines.
var format = strings.Join([]string{"%H", "%G?", "%P", "%an", "%ae", "%s",
	"%(trailers:only=true,unfold=true)"}, fieldSep) + recordSep

// Commits reads the range base..head, oldest first.
//
// An empty base or head is a caller error rather than a range: section 12 says the range is
// passed in and never inferred, so there is no default to fall back to and the caller words the
// finding. Everything else that goes wrong — a ref that does not resolve, a directory that is
// not a repository, git not being there at all — comes back as an error carrying what git said,
// because a gate reports it as a finding rather than failing the run.
func Commits(root, base, head string) ([]Commit, error) {
	if base == "" || head == "" {
		return nil, fmt.Errorf("no commit range: base %q, head %q", base, head)
	}
	cmd := exec.Command("git", "-C", root, "log", "--reverse", "--format="+format, base+".."+head)
	out, err := cmd.Output()
	if err != nil {
		msg := strings.TrimSpace(string(exitText(err)))
		if msg == "" {
			msg = err.Error()
		}
		return nil, fmt.Errorf("git log %s..%s: %s", base, head, msg)
	}
	var commits []Commit
	for _, record := range strings.Split(string(out), recordSep) {
		record = strings.TrimLeft(record, "\n")
		if strings.TrimSpace(record) == "" {
			continue
		}
		f := strings.Split(record, fieldSep)
		if len(f) < 7 {
			return nil, fmt.Errorf("git log %s..%s: a record carried %d fields, want 7", base, head, len(f))
		}
		c := Commit{Hash: f[0], Signature: f[1], Author: f[3], Email: f[4], Subject: f[5]}
		if p := strings.Fields(f[2]); len(p) > 0 {
			c.Parents = len(p)
		}
		for _, line := range strings.Split(f[6], "\n") {
			if strings.TrimSpace(line) != "" {
				c.Trailers = append(c.Trailers, strings.TrimSpace(line))
			}
		}
		commits = append(commits, c)
	}
	return commits, nil
}

// exitText is what git wrote to standard error, which says which ref did not resolve where the
// exit status alone says only that something did not.
func exitText(err error) []byte {
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		return ee.Stderr
	}
	return nil
}

// Head is the commit a repository is on, or an empty string where it has none: a directory that is
// not a repository, a repository with no commit yet, or no git at all are the same answer, because
// the caller records a fact about the repository and an invented one would be a false claim.
func Head(root string) string {
	out, err := exec.Command("git", "-C", root, "rev-parse", "HEAD").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// Paths lists the files that differ between the trees at base and head, under the directory
// given or over the whole tree where it is empty. Oldest first is meaningless here, so the
// order is git's, which is sorted by path.
//
// A tree comparison rather than a commit range, for the reason the verify job's trail guard
// gives: a directory created and dropped again inside one branch never existed on the base
// and is nobody's business, and the two-dot log form would still report it.
//
// Rename detection is off. A caller asking which intents a change touches wants both paths of
// a move, and this way it cannot be the step that lets a move through unseen.
//
// An empty base or head is a caller error rather than a range, as in Commits above: section 12
// says the range is passed in and never inferred, so there is nothing to fall back to.
func Paths(root, base, head, under string) ([]string, error) {
	if base == "" || head == "" {
		return nil, fmt.Errorf("no commit range: base %q, head %q", base, head)
	}
	args := []string{"-C", root, "diff", "--name-only", "--no-renames", base, head}
	if under != "" {
		args = append(args, "--", under)
	}
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		msg := strings.TrimSpace(string(exitText(err)))
		if msg == "" {
			msg = err.Error()
		}
		return nil, fmt.Errorf("git diff %s %s: %s", base, head, msg)
	}
	var paths []string
	for _, line := range strings.Split(string(out), "\n") {
		if line = strings.TrimSpace(line); line != "" {
			paths = append(paths, line)
		}
	}
	return paths, nil
}
