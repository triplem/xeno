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
	"runtime"
	"sort"
	"strings"
	"sync"

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
	var raw []byte
	var f Filter
	for _, rel := range []string{Shipped, Project} {
		b, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return Filter{}, err
		}
		raw = append(append(raw, b...), '\n')
		var one file
		if err := yaml.Unmarshal(b, &one); err != nil {
			return Filter{}, err
		}
		f.Patterns = append(f.Patterns, one.Patterns...)
		f.Paths = append(f.Paths, one.PathsNeverDigested...)
	}
	f.sort()
	return f.withCompiled(raw), nil
}

// compiled sets, keyed by the bytes of the files they came from. `xeno gate verify` loads the
// filter once per verdict and this repository's trail holds five hundred of them, and compiling
// two hundred regexes costs about sixty milliseconds, so the same two files were being compiled
// for half a minute of a verification.
//
// The key is the content and not the path or a timestamp, which is what makes the cache a
// memo rather than a second source of truth: the same bytes can only compile to the same set,
// so a file edited between two loads of one process misses the entry and is compiled again,
// and a test that rewrites a filter under one root sees what it wrote.
var (
	compiledMu    sync.Mutex
	compiledCache = map[string][]*regexp.Regexp{}
)

func (f Filter) withCompiled(raw []byte) Filter {
	sum := sha256.Sum256(raw)
	key := hex.EncodeToString(sum[:])
	compiledMu.Lock()
	defer compiledMu.Unlock()
	if set, ok := compiledCache[key]; ok {
		f.compiled = set
		return f
	}
	f.compile()
	compiledCache[key] = f.compiled
	return f
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

// Hit is one pattern firing in a text: which pattern, and the line it fired on. The matched
// text is deliberately not here. A gate's finding is written into `gate.yaml`, which is
// committed and sits in the trail for as long as the intent does, so a value copied out of an
// artifact into a finding is the secret published a second time, in a file nobody thinks of as
// the one holding it. The pattern id and the line are what a person needs to find it, and they
// are everything this carries.
type Hit struct {
	PatternID string
	Line      int // 1 based, counted over the text as given
}

// Matches reports every pattern that fires in text, at most once per pattern and line, sorted
// by line and then by pattern id so that two runs over one file produce the same list in the
// same order. A finding id is a hash over the cause, so an order that varied would move ids
// between runs and drop the decisions attached to them.
//
// It reads f.compiled, which is the set Redact walks, and that is what makes G-Secret and the
// digest writer share rules rather than claim to: section 4 asks for one file read by both, and
// one compiled set read by both is the same sentence one level further in. A pattern that will
// not compile is absent from both, for the same reason.
func (f Filter) Matches(text string) []Hit {
	// The offsets of every line start, computed once. A regex reports a byte offset, and
	// counting newlines from the beginning for each match would be quadratic over a long
	// artifact with a pattern that fires often.
	starts := []int{0}
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			starts = append(starts, i+1)
		}
	}
	lineAt := func(off int) int {
		return sort.Search(len(starts), func(i int) bool { return starts[i] > off })
	}
	// One worker per core over the patterns, each writing into its own slot. Two hundred
	// patterns with no keyword prefilter — section 4's file has no room for one — cost about
	// seventy microseconds per kilobyte of artifact, and `gate verify` recomputes every verdict
	// in the trail, so the serial form put this repository's own verification into the minutes.
	// Determinism is not at stake: a pattern's matches are collected under its own index and
	// the result is sorted below, so the order is a property of the filter and the text and
	// never of the scheduling.
	per := make([][]Hit, len(f.compiled))
	workers := runtime.NumCPU()
	if workers > len(f.compiled) {
		workers = len(f.compiled)
	}
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		lo := w * len(f.compiled) / workers
		hi := (w + 1) * len(f.compiled) / workers
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := lo; i < hi; i++ {
				re := f.compiled[i]
				if re == nil {
					continue
				}
				seen := map[int]bool{}
				for _, m := range re.FindAllStringIndex(text, -1) {
					line := lineAt(m[0])
					if seen[line] {
						continue
					}
					seen[line] = true
					per[i] = append(per[i], Hit{PatternID: f.Patterns[i].ID, Line: line})
				}
			}
		}()
	}
	wg.Wait()
	var hits []Hit
	for _, hs := range per {
		hits = append(hits, hs...)
	}
	sort.Slice(hits, func(i, j int) bool {
		if hits[i].Line != hits[j].Line {
			return hits[i].Line < hits[j].Line
		}
		return hits[i].PatternID < hits[j].PatternID
	})
	return hits
}
