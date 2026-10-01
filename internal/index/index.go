// SPDX-License-Identifier: Apache-2.0

// Package index reads a symbol index a project produced, so that an agent can ask where a
// name is defined instead of reading files to find out.
//
// Xeno ships no indexer and names none. What it fixes is the format of the answer, because
// a project's own toolchain knows its languages better than a tool built around a process
// can, and because every indexer carries its own licence terms and its own idea of what a
// symbol is. Section 5 of the process definition is the authority on both the format and
// the configuration.
//
// Nothing here may be reached from internal/gates. The index is used while a phase runs and
// never by a gate, which is why it is allowed to be missing, stale or wrong: a verdict never
// depends on it. A42 holds the equivalent property for the network with a check in CI, and
// this package is checked by the same step for the same reason.
package index

import (
	"fmt"
	"os"
	"time"

	"go.yaml.in/yaml/v3"
)

// DefaultMaxAge is how old an index may be before it is treated as absent. Section 5 gives
// index.max_age_hours this default.
const DefaultMaxAge = 24 * time.Hour

// Symbol is one entry: where a name is defined, and what encloses it.
//
// Kind and Container are free strings and deliberately not enumerations. Xeno cannot
// enumerate what every project's indexer thinks a symbol is — func, method, class, trait,
// object, record, impl — and a closed set would be Xeno's idea of that imposed on a tool the
// project chose. What makes the open set safe is that nothing in Xeno compares either value:
// no gate reads them, Lookup matches on Name, and the consumer is an agent.
type Symbol struct {
	Name string `yaml:"name"`
	Kind string `yaml:"kind"`
	File string `yaml:"file"`
	Line int    `yaml:"line"`
	// Container is what encloses the symbol, empty where it is at the top level. Empty is a
	// value here and not a missing field.
	Container string `yaml:"container,omitempty"`
}

// Index is a symbol index with the provenance section 5 requires of it: the tool, that
// tool's version, and when it was produced. The provenance is not decoration. An index that
// cannot say how old it is cannot be judged, and a stale index is worse than none.
type Index struct {
	Tool        string    `yaml:"tool"`
	ToolVersion string    `yaml:"tool_version"`
	ProducedAt  time.Time `yaml:"produced_at"`
	Symbols     []Symbol  `yaml:"symbols"`
}

// Load reads the index at path, or says why there is none.
//
// It cannot fail. Absent, unreadable, malformed and stale are four causes with one outcome,
// because a phase has to run without an index and an absent one is the state every
// repository is in until somebody produces one. An error return would invite a caller to
// treat that ordinary state as a problem, so the signature says it is not one: a nil index
// and a reason, which is what the tools entry of context.lock.yaml will record.
//
// now is a parameter rather than read from the clock, because staleness is the one boundary
// in this package worth testing from both sides.
func Load(path string, maxAge time.Duration, now time.Time) (*Index, string) {
	if path == "" {
		return nil, "no index.path in .xeno/config/project.yaml"
	}
	if maxAge <= 0 {
		maxAge = DefaultMaxAge
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Sprintf("no index at %s; the project produces it", path)
		}
		return nil, fmt.Sprintf("%s cannot be read: %v", path, err)
	}
	var i Index
	if err := yaml.Unmarshal(raw, &i); err != nil {
		return nil, fmt.Sprintf("%s is not a symbol index: %v", path, err)
	}
	if i.Tool == "" || i.ProducedAt.IsZero() {
		return nil, fmt.Sprintf(
			"%s names no tool or no time it was produced, so its age cannot be judged", path)
	}
	if age := now.Sub(i.ProducedAt); age > maxAge {
		return nil, fmt.Sprintf("the index at %s is %s old, past %s; a stale index is worse than none",
			path, age.Round(time.Minute), maxAge)
	}
	return &i, ""
}

// Lookup answers where a name is defined, with every location the index holds for it.
//
// The match is exact. "Where is X" is the question the index exists for and X is a name the
// agent already has, from a call site or an error; anything looser would be a search over
// the index, and a search is what the index replaces.
//
// A nil receiver returns nothing, so a caller that has no index needs no branch for it. So
// does a name the index does not hold: an index is allowed to be incomplete, and an absent
// name is an answer rather than a fault.
func (i *Index) Lookup(name string) []Symbol {
	if i == nil {
		return nil
	}
	var out []Symbol
	for _, s := range i.Symbols {
		if s.Name == name {
			out = append(out, s)
		}
	}
	return out
}

// Age is how old the index was when it was read, which the record names alongside the tool.
func (i *Index) Age(now time.Time) time.Duration {
	if i == nil {
		return 0
	}
	return now.Sub(i.ProducedAt)
}
