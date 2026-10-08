// SPDX-License-Identifier: Apache-2.0

package model

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Tracker is section 12's tracker block, which section 3's qualified id is built from.
// The whole block is optional, so every field here can be empty and the error that
// follows names the field rather than filling it in.
type Tracker struct {
	Adapter string `yaml:"adapter"`
	Project string `yaml:"project"`
	BaseURL string `yaml:"base_url"`
	Auth    Auth   `yaml:"auth"`
}

// Auth is the `tracker.auth` of Appendix A: the scheme, and the name of the environment
// variable the token lives in. The token itself is never here, which Appendix A says in
// as many words — "Credentials never appear in this file. secret_env names the variable,
// the environment holds the value" — and which is why this struct has no field for one.
//
// The scheme is read and not acted on yet, because `token` is the only one either adapter
// implements and a value nothing reads is better carried than silently dropped: a project
// that writes something else learns it from the refusal rather than from a request that
// went out with the wrong header.
type Auth struct {
	Scheme    string `yaml:"scheme"`
	SecretEnv string `yaml:"secret_env"`
}

// Issue is an issue's content as the person who raised it wrote it, and what the host
// says around it. Title and Body are the host's own two fields and what P0 quotes, as the
// problem as stated and not a summary of it. The rest is what section 12's "Starting an
// intent" reads: the labels, the milestone, the comments, and the project's open
// milestones where the issue carries one. Reading an issue stays one operation of the
// adapter contract and is up to three calls, which is why they arrive together.
//
// The fields past the body are the words both hosts share — a label is a name on either, a
// milestone a title and a due date, a comment an author, a time and a body — and nothing
// in them is judged by the adapter. What approval means is Approval below, in the domain,
// so that two adapters cannot answer it differently (A65).
//
// Milestone is a pointer because an issue with none and an answer that carried none are
// different things, and OpenMilestones is read only where the issue carries a milestone,
// which is the one case the ordering is asked about.
//
// Reason says why there is no content, where the host refused or there is no such issue.
// Section 12 makes the whole tracker block optional and Appendix A says an absent one
// means "no issue is read", so neither case is a failure of the command that asked: it is
// an answer, the same shape the branch rules port already gives a requirement it could not
// read, and the reason is what reaches the person who has to go looking.
//
// It lives here, beside the tracker block, because it is the vocabulary an adapter and the
// port it sits behind both need, and the port package selects the adapters: anything they
// share has to be somewhere neither of them is.
type Issue struct {
	Title          string
	Body           string
	Reason         string
	Labels         []string
	Milestone      *Milestone
	Comments       []Comment
	OpenMilestones []Milestone
}

// Milestone is what both hosts call one: a title, a due date as the host wrote it, and
// the number the host orders them by where two share a date. Due is a string and not a
// time, because one host writes `2026-11-30T00:00:00Z` and the other `2026-11-30`, and
// both sort as text within one host, which is the only comparison made; empty means the
// milestone has no date. Closed is true for a milestone whose turn has passed.
type Milestone struct {
	Title  string
	Due    string
	Number int
	Closed bool
}

// Comment is one comment on the issue: who wrote it, when, in the host's own stamp, and
// what. Nothing about the author's standing is carried, by the decision on #330: the
// label's right is the check of standing, and the comment is quoted rather than vetted.
type Comment struct {
	Author string
	At     string
	Body   string
}

// ApprovedLabel and ApprovedWord are the two halves of approval section 12 fixes: the
// label an issue carries, and the first line of the comment that carries the reason. They
// are constants and not configuration, so that a reader of any trail knows what the
// intake's sentence meant without a project file that may have changed since.
const (
	ApprovedLabel = "approved"
	ApprovedWord  = "approved"
)

// Approval is who approved an issue, when, and why, read off the last comment whose first
// line is ApprovedWord. By and At are the host's own author and stamp.
type Approval struct {
	By     string
	At     string
	Reason string
}

// Approval answers whether the issue carries both halves. The second return names what is
// missing, in the words a refusal prints, and is empty where nothing is.
//
// The last qualifying comment wins and not the first, because an approval is withdrawn
// and given again by writing another, and the latest is the one that stands. The first
// line is compared trimmed, with case ignored and trailing punctuation dropped, so that
// `Approved.` and `approved:` are not two different words; the reason is everything after
// that line, trimmed, and may be empty.
func (i Issue) Approval() (Approval, []string) {
	var missing []string
	labelled := false
	for _, l := range i.Labels {
		if strings.EqualFold(strings.TrimSpace(l), ApprovedLabel) {
			labelled = true
		}
	}
	if !labelled {
		missing = append(missing, "the label "+ApprovedLabel)
	}
	var found *Approval
	for _, c := range i.Comments {
		first, rest, _ := strings.Cut(strings.TrimSpace(c.Body), "\n")
		if strings.EqualFold(strings.TrimRight(strings.TrimSpace(first), ".:-"), ApprovedWord) {
			found = &Approval{By: c.Author, At: c.At, Reason: strings.TrimSpace(rest)}
		}
	}
	if found == nil {
		missing = append(missing, "a comment whose first line is "+ApprovedWord)
		return Approval{}, missing
	}
	return *found, missing
}

// Ahead is the open milestone whose turn comes before the issue's, or nil where the issue
// carries no milestone, its milestone is closed, or its milestone is the earliest open one.
// The order is section 12's: by due date, a milestone without one last, and by number
// between equals, so that the answer does not depend on the order the host listed them.
func (i Issue) Ahead() *Milestone {
	if i.Milestone == nil || i.Milestone.Closed {
		return nil
	}
	var first *Milestone
	for k := range i.OpenMilestones {
		m := &i.OpenMilestones[k]
		if m.Closed {
			continue
		}
		if first == nil || m.before(*first) {
			first = m
		}
	}
	if first == nil || first.Title == i.Milestone.Title {
		return nil
	}
	return first
}

// before orders two milestones: a dated one before an undated one, an earlier date before
// a later one, and the lower number first between equals.
func (m Milestone) before(o Milestone) bool {
	switch {
	case (m.Due == "") != (o.Due == ""):
		return m.Due != ""
	case m.Due != o.Due:
		return m.Due < o.Due
	}
	return m.Number < o.Number
}

// Found says whether there is content to carry. A title with no body counts, since an
// issue may legitimately have none; the zero value does not, which is what a caller that
// never asked a host holds, and an empty section written from it would say that the problem
// was stated nowhere rather than that nobody was asked.
func (i Issue) Found() bool { return i.Reason == "" && (i.Title != "" || i.Body != "") }

// Written is where a comment landed, or why it did not. Same shape and same reason as
// Issue above: a host that refuses the write and an issue that is not there are answers a
// pipeline step reports, not errors that fail the job that reported the verdict.
//
// URL is empty where the host's answer carries none. GitLab's note payload has no web
// address in it, so the field is one host's answer and not the contract's promise.
type Written struct {
	URL    string
	Reason string
}

// Qualified builds the canonical intent id of section 3, which is the tracker host, the
// project and the key, from as much of it as a person named and the configuration for
// the rest.
//
// The argument is the issue: a bare key, `176`, or any longer prefix of the id ending
// in one, up to the whole `github.com/triplem/xeno#176`. One rule covers all of them,
// because the parts a person leaves out are the parts the configuration already holds,
// and a command that accepted only the two ends would make the middle unsayable.
//
// The key is not required to be a number. Section 3 calls it a key and names Jira as
// the case where it is not one, so what is checked is that it carries no space: YAML
// reads " #" as the start of a comment, and `qualified` in the runner exists because an
// id cut off by one is read back shorter without anybody noticing.
func (t Tracker) Qualified(issue string) (string, error) {
	project, key := t.Project, strings.TrimSpace(issue)
	if i := strings.LastIndex(key, "#"); i >= 0 {
		project, key = strings.TrimSpace(key[:i]), strings.TrimSpace(key[i+1:])
	}
	if key == "" {
		return "", fmt.Errorf("--for is the issue the intent belongs to, as a key, 176, " +
			"or as the whole qualified id, github.com/triplem/xeno#176")
	}
	if project == "" {
		return "", fmt.Errorf("no tracker.project in .xeno/config/project.yaml, so the " +
			"repository the key belongs to is not in the configuration; give --for the " +
			"whole qualified id, host, repository and key")
	}
	if !strings.Contains(firstSegment(project), ".") {
		host, err := t.host()
		if err != nil {
			return "", err
		}
		project = host + "/" + project
	}
	id := project + "#" + key
	if strings.ContainsAny(id, " \t") {
		return "", fmt.Errorf("%q holds a space, and an intent id with one in it is read "+
			"back cut off where it sits beside a comment", id)
	}
	return id, nil
}

// Locate is Qualified read backwards: the project and the key an adapter has to be given
// in order to reach the issue a qualified intent id names. Section 12's first operation is
// resolving an intent, and this is the half of it the other three need, because the adapter
// "knows no relationship between issue key and repository" and takes both as parameters.
//
// The project comes from the id and not from `tracker.project`, because `--for` takes a
// whole qualified id and an intent may therefore belong to an issue in another repository.
// Reading the configured project instead would send the read, or worse the comment, to
// whichever repository this one is configured for.
//
// The id's host is checked against the configured address rather than ignored. An adapter
// pointed at one deployment cannot answer about an issue on another, and the id it would
// otherwise ask about reads as a repository path with a hostname as its first segment,
// which that host answers 404 to: a wrong question asked of the wrong host, reported as a
// missing issue. The error names both hosts instead, and what the caller does with it is
// the caller's: the runner reports it as an issue that was not read, because an intent
// legitimately may belong to one.
func (t Tracker) Locate(qualified string) (project, key string, err error) {
	i := strings.LastIndex(qualified, "#")
	if i <= 0 || i == len(qualified)-1 {
		return "", "", fmt.Errorf("%q is not a qualified intent id, which is the tracker "+
			"host, the project and the key, as github.com/triplem/xeno#176", qualified)
	}
	path, key := qualified[:i], qualified[i+1:]
	h, err := t.host()
	if err != nil {
		return "", "", err
	}
	rest, ok := strings.CutPrefix(path, h+"/")
	if !ok || rest == "" {
		return "", "", fmt.Errorf("the intent belongs to an issue on %s and "+
			"tracker.base_url names %s, so this adapter cannot reach it", path, h)
	}
	return rest, key, nil
}

// host is the tracker host of the id, taken from the API address the project configured
// and not from the adapter's name: a self managed deployment of either host answers at
// its own address, and a name mapped to an address is a table that is right until the
// first such deployment.
//
// The leading `api.` goes, which is the one place the two differ: GitHub's API for
// github.com is api.github.com and its id host is github.com, while every other
// deployment of either host serves the API under the host itself.
func (t Tracker) host() (string, error) {
	u, err := url.Parse(t.BaseURL)
	if err != nil || u.Host == "" {
		return "", fmt.Errorf("tracker.base_url in .xeno/config/project.yaml is %q, which "+
			"names no host, so the first part of the intent id cannot be read off it; give "+
			"--for the whole qualified id", t.BaseURL)
	}
	return strings.TrimPrefix(u.Host, "api."), nil
}

func firstSegment(s string) string { return strings.SplitN(s, "/", 2)[0] }

// NextKey continues the sequence of keys an intents directory already holds: the prefix
// they share and the number after the highest, padded as that highest one is.
//
// The sequence is read off the directory rather than configured, because that is where
// it is: a key is a prefix and a number, nothing in project.yaml names the prefix, and
// a new field for it would be a second place the answer lives. The padding comes from
// the key the number was taken from, so a project that pads to three keeps three and
// this one keeps four (D-7).
//
// Names that are not a key are skipped rather than refused. A directory is only in the
// sequence if it reads as one, and a repository holding none has no sequence to
// continue, which is reported so that the caller can ask for the first key by name.
func NextKey(names []string) (string, error) {
	prefix, high, width := "", 0, 0
	for _, name := range names {
		p, n, w, ok := splitKey(name)
		if !ok {
			continue
		}
		if prefix == "" {
			prefix = p
		}
		if p != prefix {
			return "", fmt.Errorf("the keys here carry two prefixes, %s and %s, so there is "+
				"no one sequence to continue; name the key with --intent", prefix, p)
		}
		if n > high {
			high, width = n, w
		}
	}
	if prefix == "" {
		return "", fmt.Errorf("there is no key here to continue a sequence from, so the " +
			"first one is named rather than derived; give --intent")
	}
	return fmt.Sprintf("%s-%0*d", prefix, width, high+1), nil
}

// splitKey reads a key as a prefix and a number. The split is the last hyphen, so a
// prefix may hold one; what has to be a number is the part after it, because that is
// the part the sequence counts.
func splitKey(name string) (prefix string, number, width int, ok bool) {
	i := strings.LastIndex(name, "-")
	if i <= 0 || i == len(name)-1 {
		return "", 0, 0, false
	}
	n, err := strconv.Atoi(name[i+1:])
	if err != nil || n < 0 {
		return "", 0, 0, false
	}
	return name[:i], n, len(name) - i - 1, true
}
