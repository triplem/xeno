// SPDX-License-Identifier: Apache-2.0

// Package secrets implements the secret filter of section 4: the shipped pattern file plus a
// project's additions, the hash over the effective set, and the redaction a digest passes
// through.
//
// It is a package of its own because section 4 says G-Secret and the digest writer read the
// same file. A gate written later reads this rather than copying it, which is what makes the
// claim that the two share rules verifiable rather than stated.
package secrets

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"go.yaml.in/yaml/v3"
)

// Shipped and Project are where section 4 puts the two files, relative to a repository root.
const (
	Shipped = ".xeno/plugin/secrets.yaml"
	Project = ".xeno/config/secrets.yaml"
)

// Pattern is one entry of the filter: an id a redaction can name, and the regex it matches.
type Pattern struct {
	ID    string `yaml:"id"`
	Regex string `yaml:"regex"`
}

type file struct {
	Patterns           []Pattern `yaml:"patterns"`
	PathsNeverDigested []string  `yaml:"paths_never_digested"`
}

// Filter is the effective set: every shipped pattern and every one a project adds. A project
// removes nothing, because section 4 says a filter a project can switch off is not a filter,
// so this is a union and matching is a disjunction. A project pattern reusing a shipped id
// therefore leaves the shipped one in force rather than replacing it.
type Filter struct {
	Patterns []Pattern
	Paths    []string
	compiled []*regexp.Regexp // in the order of Patterns; a pattern that will not compile is dropped
}

// Load assembles the effective filter. A missing shipped file yields an empty filter and no
// error: a repository before its plugin is vendored has no filter, which is a state the runner
// reports by writing no secrets_hash rather than one it refuses a phase for.
func Load(root string) (Filter, error) {
	var f Filter
	for _, rel := range []string{Shipped, Project} {
		b, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return Filter{}, err
		}
		var one file
		if err := yaml.Unmarshal(b, &one); err != nil {
			return Filter{}, err
		}
		f.Patterns = append(f.Patterns, one.Patterns...)
		f.Paths = append(f.Paths, one.PathsNeverDigested...)
	}
	f.sort()
	f.compile()
	return f, nil
}

// Empty reports whether any pattern or path is in force. An empty filter is the state before
// a plugin is vendored, and it is what makes secrets_hash absent rather than a hash of nothing.
func (f Filter) Empty() bool { return len(f.Patterns) == 0 && len(f.Paths) == 0 }

// sort makes the set a set rather than two concatenated files, which is what lets two projects
// with the same effective filter agree on its hash whatever order they wrote it in.
func (f *Filter) sort() {
	sort.SliceStable(f.Patterns, func(i, j int) bool {
		if f.Patterns[i].ID != f.Patterns[j].ID {
			return f.Patterns[i].ID < f.Patterns[j].ID
		}
		return f.Patterns[i].Regex < f.Patterns[j].Regex
	})
	sort.Strings(f.Paths)
}

// compile drops a pattern whose regex will not compile rather than failing the run. A filter
// that refuses to load stops every phase of a repository over one bad line in a project file,
// and the pattern still enters the hash, so what was in force stays reconstructable.
func (f *Filter) compile() {
	f.compiled = make([]*regexp.Regexp, len(f.Patterns))
	for i, p := range f.Patterns {
		if re, err := regexp.Compile(p.Regex); err == nil {
			f.compiled[i] = re
		}
	}
}

// Hash is secrets_hash: a sha256 over the effective set in a canonical rendering, one line per
// entry, sorted, tab separated.
//
// Appendix B fixes every other hash to the byte and delegates this one to "the package that
// first writes them", because how a set of several files reduces to one value belongs with the
// code that assembles the set. So this covers the set's content and not the bytes of the files
// it came from: a comment edited in either file leaves the value alone, and two projects whose
// effective filters agree agree here, which is what makes the field say which filter a digest
// passed through rather than which files happened to exist. A21 records the decision.
//
// An empty filter has no hash. The caller writes no field, rather than the hash of an empty
// string, which would assert a filter that is not there.
func (f Filter) Hash() string {
	if f.Empty() {
		return ""
	}
	var b strings.Builder
	for _, p := range f.Patterns {
		b.WriteString("pattern\t" + p.ID + "\t" + p.Regex + "\n")
	}
	for _, path := range f.Paths {
		b.WriteString("path\t" + path + "\n")
	}
	sum := sha256.Sum256([]byte(b.String()))
	return hex.EncodeToString(sum[:])
}

// Redact replaces every match with the id of the pattern that fired. The line around it is
// left, because a digest is read by a person and a silent deletion would make the sentence
// around it a lie; naming the pattern says what was removed and by which rule.
//
// Section 16 is why this is here rather than in the agent: the runner holds the text between
// the summary and the file, so the filtering is deterministic and outside the model's reach.
func (f Filter) Redact(text string) string {
	for i, re := range f.compiled {
		if re == nil {
			continue
		}
		text = re.ReplaceAllString(text, "[redacted: "+f.Patterns[i].ID+"]")
	}
	return text
}
