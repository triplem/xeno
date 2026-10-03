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
