// SPDX-License-Identifier: Apache-2.0

// Package enforcement compares what a project declares it requires of its host against
// what a host adapter reports. It holds no host: the requirement names are section 13's
// and the words and states belong to whichever adapter under internal/host answered.
//
// It is deliberately not reachable from the gate path: a verdict has to be reproducible
// from the repository alone, and a gate that called out would make it depend on whether a
// service answered. Nothing under internal/gates imports this or an adapter.
package enforcement

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// State is what could be learned about one requirement. The three are different and
// collapsing them is what makes a report worthless: a setting nobody made is somebody's
// oversight, a setting the host does not have is nobody's, and a waived one is a
// decision already taken.
type State string

const (
	Met          State = "met"
	Unmet        State = "unmet"
	NotAvailable State = "not-available"
	Waived       State = "waived"
	Unknown      State = "unknown"
)

// Requirement is one line of the report.
type Requirement struct {
	Name     string `yaml:"name" json:"name"`
	Declared string `yaml:"declared" json:"declared"`
	Actual   string `yaml:"actual" json:"actual"`
	State    State  `yaml:"state" json:"state"`
	Note     string `yaml:"note,omitempty" json:"note,omitempty"`
}

// Report is written as a pipeline artifact and committed by nothing. CI verifies
// artifacts that exist and produces none, and a second place carrying a verdict beside
// gate.yaml is what the schema forbids.
type Report struct {
	Repository   string        `yaml:"repository"`
	Branch       string        `yaml:"branch"`
	CheckedAt    string        `yaml:"checked_at"`
	Requirements []Requirement `yaml:"requirements"`
}

// Unmet counts the requirements that are neither met nor excused. A waived one is
// excused by a decision and a not-available one that is not waived is not: that is the
// case a project has to look at once and then record.
func (r Report) Unmet() int {
	n := 0
	for _, q := range r.Requirements {
		if q.State == Unmet || q.State == NotAvailable {
			n++
		}
	}
	return n
}

// The requirement names are section 13's and the second standing rule makes the list a
// budget rather than a starting point. They are constants so that an adapter builds its
// answers from them: a host may choose its own words for what it found, never its own name
// for what was asked, and a name spelled wrong in an adapter does not compile.
const (
	NameRequiredPipeline     = "required_pipeline"
	NameAllowBypass          = "allow_bypass"
	NameApprovalsRequired    = "approvals.required"
	NameApprovalsNotByAuthor = "approvals.not_by_author"
	NameMergeMethod          = "merge_method"
)

// Names returns the requirement names in report order. An adapter's test asserts its own
// names are among these, which is where the budget is held: Compare drops a name it does
// not know, and there is nowhere in Report to say so that would not itself be a field the
// specification does not have.
func Names() []string {
	return []string{NameRequiredPipeline, NameAllowBypass, NameApprovalsRequired,
		NameApprovalsNotByAuthor, NameMergeMethod}
}

// Declared is the enforcement block of project.yaml.
type Declared struct {
	RequiredPipeline bool   `yaml:"required_pipeline"`
	AllowBypass      bool   `yaml:"allow_bypass"`
	MergeMethod      string `yaml:"merge_method,omitempty"`
	Approvals        struct {
		Required    int    `yaml:"required"`
		NotByAuthor bool   `yaml:"not_by_author"`
		Waived      string `yaml:"waived,omitempty"`
	} `yaml:"approvals"`
	// Waived excuses the block as a whole where a host cannot express any of it.
	Waived string `yaml:"waived,omitempty"`
}

// requirement is what the domain still knows about one line of the report: its name,
// whether the declaration asks for it, and the reason it may carry of its own. What was
// once here as well, how to read a host's answer, belongs to the adapter that has the
// answer: see A65, and #97 for the question it closes.
type requirement struct {
	name     string
	declared func(Declared) (string, bool)
	// waived is the reason this requirement may carry of its own, beyond the block's.
	waived func(Declared) string
}

var requirements = []requirement{
	{
		name:     NameRequiredPipeline,
		declared: func(d Declared) (string, bool) { return "true", d.RequiredPipeline },
	},
	{
		name:     NameAllowBypass,
		declared: func(d Declared) (string, bool) { return "false", !d.AllowBypass },
	},
	{
		name: NameApprovalsRequired,
		declared: func(d Declared) (string, bool) {
			return fmt.Sprintf("%d", d.Approvals.Required), d.Approvals.Required > 0
		},
		waived: func(d Declared) string { return d.Approvals.Waived },
	},
	{
		name:     NameApprovalsNotByAuthor,
		declared: func(d Declared) (string, bool) { return "true", d.Approvals.NotByAuthor },
		waived:   func(d Declared) string { return d.Approvals.Waived },
	},
}

// Compare turns a declaration and an adapter's answer into the report.
//
// The adapter has already decided met from unmet and what the host cannot express at all,
// because only the host knows how to read its own answer. What is left here is what the
// project decided rather than what the host reports: which requirements were asked for, and
// which unmet or unavailable ones a waiver excuses.
//
// A declared requirement the adapter did not answer for is not available. That is the
// honest reading of an absence: the host was asked and had nothing to say about it, which
// is the same state as a host that has no such setting.
func Compare(repo, branch string, d Declared, answered []Requirement, now time.Time) Report {
	rep := Report{Repository: repo, Branch: branch, CheckedAt: now.UTC().Format(time.RFC3339)}
	byName := make(map[string]Requirement, len(answered))
	for _, q := range answered {
		byName[q.Name] = q
	}
	for _, q := range requirements {
		declared, asked := q.declared(d)
		if !asked {
			continue
		}
		w := d.Waived
		if q.waived != nil && q.waived(d) != "" {
			w = q.waived(d)
		}
		got, ok := byName[q.name]
		if !ok {
			rep.add(q.name, declared, "not available", waiveOr(w, NotAvailable), noteOf(w, ""))
			continue
		}
		if got.State == Met {
			rep.add(q.name, declared, got.Actual, Met, got.Note)
			continue
		}
		rep.add(q.name, declared, got.Actual, waiveOr(w, got.State), noteOf(w, got.Note))
	}

	// merge_method is declared and deliberately not compared. Section 13 says "not
	// checked" for it, and the plan says it is meaningful once a commit predicate is
	// active, which no gate in this runner has yet: a squash replaces the commits a
	// predicate judged, so the requirement exists for a consumer that does not. Comparing
	// it would be a check the specification says is not made, which is a spec change
	// first. The field is reported as unchecked rather than left silent, so that a project
	// declaring it learns that nothing reads it.
	if d.MergeMethod != "" {
		rep.add(NameMergeMethod, d.MergeMethod, "not compared", Unknown,
			"section 13 leaves it unchecked; it is meaningful once a commit predicate is active")
	}
	return rep
}

func (r *Report) add(name, declared, actual string, s State, note string) {
	r.Requirements = append(r.Requirements,
		Requirement{Name: name, Declared: declared, Actual: actual, State: s, Note: note})
}

// waiveOr is where waived does its work. Without it a requirement the host cannot
// express is reported as unmet on every run forever, and a report that always says the
// same thing is ignored within a fortnight.
func waiveOr(waived string, otherwise State) State {
	if strings.TrimSpace(waived) != "" {
		return Waived
	}
	return otherwise
}

func noteOf(waived, reason string) string {
	if strings.TrimSpace(waived) != "" {
		return waived
	}
	return reason
}

// Token reads the enforcement token from the environment. It is never in the
// repository, which is the same rule the tracker credential follows.
func Token() string {
	for _, k := range []string{"XENO_ENFORCEMENT_TOKEN", "GITHUB_TOKEN"} {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return ""
}
