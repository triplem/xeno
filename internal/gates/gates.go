// SPDX-License-Identifier: Apache-2.0

// Package gates evaluates the gates of one phase. Everything here is deterministic and
// free of network access: gates read, they never run anything and never ask a model.
package gates

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
)

// Ctx is what a gate run knows about.
type Ctx struct {
	Root, Key, Phase string
	ArtifactsHash    string
	QualifiedID      string // from intent.yaml
}

func (c Ctx) phaseRel(p string) string { return model.PhaseDir(c.Key, p) }
func (c Ctx) abs(rel string) string    { return filepath.Join(c.Root, rel) }

type gateFn func(Ctx) model.Check

type spec struct {
	id   string
	from int // phase index from which the gate applies
	fn   gateFn
}

// The order is the table in section 7: integrity first, then cost. Gates this runner
// does not implement yet are written as such rather than skipped.
var table = []spec{
	{"G-Supply", 0, notImplemented},
	{"G-Schema", 0, schema},
	{"G-Trace", 0, trace},
	{"G-Secret", 0, notImplemented},
	{"G-Assumptions", 0, assumptions},
	{"G-Questions", 5, questions},
	{"G-Learning", 0, learning},
	{"G-Freshness", 0, freshness},
	{"G-Evidence", 3, evidence},
	{"G-Build", 3, build},
	{"G-Test", 4, notImplemented},
	{"G-Rules", 0, notImplemented},
	{"G-Policy", 0, notImplemented},
	{"G-Complete", 5, notImplemented},
}

// ExternalProvenance marks a check whose findings were produced by foreign code.
const ExternalProvenance = "external"

// carryForward gives every finding its id and attaches the decision the previous run
// held for it. It is the only place a decision is attached, and it refuses to attach
// one to a finding from an external gate.
//
// The id is a hash over gate, rule, file and cause. For a gate of this runner the cause
// is written in this repository, so the id is stable by a promise we keep and any change
// to it is visible in a diff. An external gate's cause comes from foreign code, and the
// same id across two runs then means only that somebody else's wording did not change: a
// tool reporting a different problem in the same words would carry an old release onto it
// unnoticed, and one rewording cosmetically would drop every decision at once. Xeno
// cannot tell those apart from the outside, so it does neither, and the decision is taken
// again. A project wanting a lasting release for an external finding writes a rule of its
// own, whose cause Xeno words and can therefore keep stable.
//
// The provenance read is this run's, not the previous one's: the question is whether the
// finding is foreign now, not whether it once was.
func carryForward(ch *model.Check, decided map[string]*model.DecisionOnFinding) {
	for i := range ch.Findings {
		f := &ch.Findings[i]
		f.ID = hashing.FindingID(ch.Gate, "", f.File, f.Cause)
		if ch.Provenance == ExternalProvenance {
			continue
		}
		f.Decision = decided[f.ID]
	}
}

// Run evaluates every gate that applies to the phase, carrying decisions forward from
// the previous gate.yaml by finding id.
func Run(c Ctx, previous *model.Gate) []model.Check {
	idx := model.PhaseIndex(c.Phase)
	decided := map[string]*model.DecisionOnFinding{}
	if previous != nil {
		for _, ch := range previous.Checks {
			for _, f := range ch.Findings {
				if f.Decision != nil {
					decided[f.ID] = f.Decision
				}
			}
		}
	}
	var out []model.Check
	for _, s := range table {
		if idx < s.from {
			continue
		}
		ch := s.fn(c)
		ch.Gate = s.id
		carryForward(&ch, decided)
		out = append(out, ch)
	}
	return out
}

// Status derives the phase status from the checks. Nothing writes it directly.
// Order, loudest first: red, overridden, provisional, approved, green.
func Status(checks []model.Check) (string, error) {
	seen := map[string]bool{}
	undecided, overridden, pending, approved := false, false, false, false
	for _, ch := range checks {
		if ch.Result == "pending" {
			pending = true
		}
		for _, f := range ch.Findings {
			if seen[f.ID] {
				return "", fmt.Errorf("finding id collision %s: refusing to let one decision cover two findings", f.ID)
			}
			seen[f.ID] = true
			switch {
			case f.Decision == nil:
				undecided = true
			case f.Decision.Type == "overridden":
				overridden = true
			default:
				approved = true
			}
		}
	}
	switch {
	case undecided:
		return "red", nil
	case overridden:
		return "overridden", nil
	case pending:
		return "provisional", nil
	case approved:
		return "approved", nil
	}
	return "green", nil
}

func notImplemented(Ctx) model.Check {
	return model.Check{Result: "not-implemented", Provenance: "xeno"}
}

func result(fs []model.Finding) model.Check {
	r := "pass"
	if len(fs) > 0 {
		r = "fail"
	}
	return model.Check{Result: r, Provenance: "xeno", Findings: fs}
}

func finding(file, cause, next string) model.Finding {
	return model.Finding{File: file, Cause: cause, Next: next}
}

// ---- G-Schema

var (
	commonFields   = []string{"intent", "created", "runner_version", "plugin_version", "phase"}
	sessionFields  = []string{"language", "secrets_hash", "context_hash", "model", "tool", "tool_version"}
	renderedFields = []string{"template", "strings_hash", "rules_hash"}
)

// schema_version is deliberately absent from commonFields. An artifact written before
// the field existed does not carry it, and the process definition reads that absence as
// the schema that predates versioning rather than as a field somebody forgot: a gate
// that failed on it would invalidate every artifact behind the change that introduced
// it. What is checked is the shape of the value where there is one.
var schemaVersionShape = regexp.MustCompile(`^[0-9]+\.[0-9]+$`)

func schemaVersion(file string, raw map[string]any) []model.Finding {
	v, ok := raw["schema_version"]
	if !ok || v == nil {
		return nil // predates versioning, and that is readable rather than wrong
	}
	if s, isString := v.(string); !isString || !schemaVersionShape.MatchString(s) {
		return []model.Finding{finding(file, fmt.Sprintf("schema_version %v is not major.minor", v),
			"write it as two numbers separated by a dot, or leave it out where the artifact predates the field")}
	}
	return nil
}

func missing(raw map[string]any, fields ...[]string) []string {
	var m []string
	for _, group := range fields {
		for _, f := range group {
			if v, ok := raw[f]; !ok || v == nil || v == "" {
				m = append(m, f)
			}
		}
	}
	return m
}

func schema(c Ctx) model.Check {
	dir := c.phaseRel(c.Phase)
	var fs []model.Finding

	out := dir + "/output.md"
	var o model.Output
	raw, err := fm.ReadFront(c.abs(out), &o)
	switch {
	case os.IsNotExist(err):
		fs = append(fs, finding(out, "output.md is missing", "write the phase result through the MCP operation or its command"))
	case err != nil:
		fs = append(fs, finding(out, "frontmatter unreadable: "+err.Error(), "repair the frontmatter"))
	default:
		for _, f := range missing(raw, commonFields, sessionFields, renderedFields) {
			fs = append(fs, finding(out, "required field missing: "+f, "add "+f+" to the frontmatter"))
		}
		fs = append(fs, schemaVersion(out, raw)...)
		fs = append(fs, questionShape(out, o)...)
		fs = append(fs, decisionShape(out, o)...)
	}

	// Section 4 lists digest.md among the files every phase directory holds, and no
	// other gate looks for it: G-Learning guards learning.yaml, G-Freshness guards
	// context.lock.yaml from P1, and cost.yaml is deliberately unguarded because it may
	// arrive after the gate ran. Absent this check a phase finished without its digest
	// was green, which is a verdict on a phase that is not complete.
	dig := dir + "/digest.md"
	if !fm.Exists(c.abs(dig)) {
		fs = append(fs, finding(dig, "digest.md is missing", "write the digest of the session; a phase without one is incomplete"))
	} else {
		r, _ := fm.ReadFront(c.abs(dig), nil)
		for _, f := range missing(r, commonFields, sessionFields) {
			fs = append(fs, finding(dig, "required field missing: "+f, "add "+f+" to the frontmatter"))
		}
		fs = append(fs, schemaVersion(dig, r)...)
	}

	// Section 5 states its field sets for frontmatter and for the equivalent top level
	// keys in YAML artifacts alike, so the YAML ones are checked here rather than left
	// to G-Trace, which looks at two fields and judges binding rather than shape.
	for _, y := range []string{"context.lock.yaml", "learning.yaml"} {
		rel := dir + "/" + y
		raw := map[string]any{}
		if err := fm.ReadYAML(c.abs(rel), &raw); err != nil {
			continue // absence is G-Learning's finding, or G-Freshness's
		}
		for _, f := range missing(raw, commonFields) {
			fs = append(fs, finding(rel, "required field missing: "+f, "add "+f+" to the file"))
		}
		fs = append(fs, schemaVersion(rel, raw)...)
	}

	entries, _ := os.ReadDir(c.abs(dir))
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			if name != "evidence" {
				fs = append(fs, finding(dir+"/"+name, "unknown directory in the phase directory", "remove it or move it outside .xeno/"))
			}
			continue
		}
		if model.KnownPhaseFiles[name] || (name == model.ContextProfile && c.Phase == model.Phases[0]) {
			continue
		}
		fs = append(fs, finding(dir+"/"+name, "unknown file in the phase directory", "remove it; only files Xeno writes belong here"))
	}
	fs = append(fs, undeclaredEvidence(c, dir, o)...)
	return result(fs)
}

// A question carries two to four options plus the free entry, or states that none
// were found. It never carries two invented ones to satisfy this check.
func questionShape(file string, o model.Output) []model.Finding {
	var fs []model.Finding
	keys := map[string]bool{}
	for _, q := range o.OpenQuestions {
		if q.Key == "" {
			fs = append(fs, finding(file, "open question without a key", "give the question a stable key"))
			continue
		}
		if keys[q.Key] {
			fs = append(fs, finding(file, "duplicate question key "+q.Key, "keys are unique within the intent"))
		}
		keys[q.Key] = true
		if q.NoOptions {
			continue
		}
		proper, free := 0, 0
		for _, op := range q.Options {
			if op.Free {
				free++
			} else {
				proper++
			}
		}
		if proper < 2 || proper > 4 || free != 1 {
			fs = append(fs, finding(file, fmt.Sprintf("question %s needs two to four options and one free entry, has %d and %d", q.Key, proper, free),
				"offer the options found, or state no_options: true where there are none"))
		}
	}
	return fs
}

func decisionShape(file string, o model.Output) []model.Finding {
	var fs []model.Finding
	for _, d := range o.Decisions {
		if d.ID == "" || d.DecidedBy == "" || d.Rationale == "" || (!d.Withdrawn && d.Chosen == "") {
			fs = append(fs, finding(file, "decision "+d.ID+" incomplete: id, chosen, rationale and decided_by are required",
				"complete the decision; a withdrawal needs a reason and a person as well"))
		}
	}
	return fs
}

// Nothing may sit in evidence/ that no declaration and no attachment points at.
func undeclaredEvidence(c Ctx, dir string, o model.Output) []model.Finding {
	evDir := dir + "/evidence"
	entries, err := os.ReadDir(c.abs(evDir))
	if err != nil {
		return nil
	}
	known := map[string]bool{"attached.yaml": true}
	for _, e := range o.Evidence {
		known[filepath.Base(e.Path)] = e.Path != ""
	}
	var att []model.Attached
	_ = fm.ReadYAML(c.abs(evDir+"/attached.yaml"), &att)
	for _, a := range att {
		if a.Path != "" {
			known[filepath.Base(a.Path)] = true
		}
	}
	var fs []model.Finding
	for _, e := range entries {
		if !known[e.Name()] {
			fs = append(fs, finding(evDir+"/"+e.Name(), "undeclared file in evidence/", "declare it in output.md or remove it"))
		}
	}
	return fs
}

// ---- G-Trace

func trace(c Ctx) model.Check {
	dir := c.phaseRel(c.Phase)
	var fs []model.Finding
	check := func(rel string, raw map[string]any) {
		if raw["intent"] != c.QualifiedID {
			fs = append(fs, finding(rel, fmt.Sprintf("bound to intent %v, lies in %s", raw["intent"], c.QualifiedID), "move the artifact or correct its intent field"))
		}
		if raw["phase"] != c.Phase {
			fs = append(fs, finding(rel, fmt.Sprintf("bound to phase %v, lies in %s", raw["phase"], c.Phase), "move the artifact or correct its phase field"))
		}
	}
	for _, md := range []string{"output.md", "digest.md"} {
		if raw, err := fm.ReadFront(c.abs(dir+"/"+md), nil); err == nil {
			check(dir+"/"+md, raw)
		}
	}
	for _, y := range []string{"context.lock.yaml", "learning.yaml"} {
		raw := map[string]any{}
		if err := fm.ReadYAML(c.abs(dir+"/"+y), &raw); err == nil {
			check(dir+"/"+y, raw)
		}
	}
	return result(fs)
}

// ---- G-Assumptions

func readAssumptions(c Ctx) model.Assumptions {
	var a model.Assumptions
	_ = fm.ReadYAML(c.abs(model.IntentDir(c.Key)+"/assumptions.yaml"), &a)
	return a
}

func assumptions(c Ctx) model.Check {
	var fs []model.Finding
	for _, a := range readAssumptions(c).Assumptions {
		if a.ConfirmedBy == "" {
			fs = append(fs, finding(model.IntentDir(c.Key)+"/assumptions.yaml", "assumption "+a.ID+" is not confirmed", "confirm it, or replace it with a decision"))
		}
	}
	return result(fs)
}

// ---- G-Questions, from P5

func questions(c Ctx) model.Check {
	type raised struct{ key, file string }
	var qs []raised
	resolved := map[string]bool{}
	for _, p := range model.Phases {
		rel := c.phaseRel(p) + "/output.md"
		var o model.Output
		if _, err := fm.ReadFront(c.abs(rel), &o); err != nil {
			continue
		}
		for _, q := range o.OpenQuestions {
			qs = append(qs, raised{q.Key, rel})
		}
		for _, d := range o.Decisions {
			if d.Resolves != "" && d.DecidedBy != "" {
				resolved[d.Resolves] = true // decision or withdrawal
			}
		}
	}
	for _, a := range readAssumptions(c).Assumptions {
		if a.Resolves != "" && a.ConfirmedBy != "" {
			resolved[a.Resolves] = true
		}
	}
	var fs []model.Finding
	for _, q := range qs {
		if !resolved[q.key] {
			fs = append(fs, finding(q.file, "open question "+q.key+" is unresolved", "resolve it as a decision, a confirmed assumption or a withdrawal"))
		}
	}
	sort.Slice(fs, func(i, j int) bool { return fs[i].Cause < fs[j].Cause })
	return result(fs)
}

// ---- G-Learning

func learning(c Ctx) model.Check {
	rel := c.phaseRel(c.Phase) + "/learning.yaml"
	raw := map[string]any{}
	err := fm.ReadYAML(c.abs(rel), &raw)
	if os.IsNotExist(err) {
		return result([]model.Finding{finding(rel, "learning record missing", "write one, even when the result is no finding")})
	}
	for k := range raw {
		switch k {
		case "intent", "phase", "created", "runner_version", "plugin_version":
		default:
			return result(nil)
		}
	}
	return result([]model.Finding{finding(rel, "learning record is empty", "record an observation or state no_finding: true")})
}

// ---- G-Freshness: the first half, the context hash against the predecessor.

func freshness(c Ctx) model.Check {
	idx := model.PhaseIndex(c.Phase)
	if idx == 0 {
		return result(nil)
	}
	rel := c.phaseRel(c.Phase) + "/context.lock.yaml"
	var lock model.ContextLock
	if err := fm.ReadYAML(c.abs(rel), &lock); err != nil {
		return result([]model.Finding{finding(rel, "context lock missing", "start the phase with xeno phase start")})
	}
	cur, err := hashing.DirHash(c.Root, c.phaseRel(model.Phases[idx-1]), hashing.PhaseExcluded)
	if err != nil || cur != lock.PredecessorHash {
		return result([]model.Finding{finding(rel, "predecessor changed after this phase started", "rerun this phase against the current predecessor")})
	}
	return result(nil)
}

// ---- G-Evidence and G-Build

// Resolved evidence of a phase: declarations merged with what was attached.
type Resolved struct {
	Decl     model.EvidenceItem
	Attached *model.Attached
}

func Collect(c Ctx) ([]Resolved, error) {
	dir := c.phaseRel(c.Phase)
	var o model.Output
	if _, err := fm.ReadFront(c.abs(dir+"/output.md"), &o); err != nil {
		return nil, err
	}
	var att []model.Attached
	_ = fm.ReadYAML(c.abs(dir+"/evidence/attached.yaml"), &att)
	var out []Resolved
	for _, d := range o.Evidence {
		r := Resolved{Decl: d}
		if d.Pending() {
			for i := range att {
				if att[i].Kind == d.Kind && att[i].Job == d.Job && att[i].State == "attached" {
					r.Attached = &att[i]
				}
			}
		}
		out = append(out, r)
	}
	return out, nil
}

func evidence(c Ctx) model.Check {
	dir := c.phaseRel(c.Phase)
	items, err := Collect(c)
	if err != nil {
		return result(nil) // G-Schema already reports an unreadable output.md
	}
	var fs []model.Finding
	pending := false
	verify := func(label, sum, path, uri string) {
		switch {
		case path != "":
			got, err := hashing.FileHash(c.abs(dir + "/" + path))
			if err != nil {
				fs = append(fs, finding(dir+"/"+path, label+" does not resolve", "restore the file or correct the declaration"))
			} else if got != sum {
				fs = append(fs, finding(dir+"/"+path, label+" content does not match its hash", "the file changed after it was declared"))
			}
		case uri == "":
			fs = append(fs, finding(dir+"/output.md", label+" has neither path nor uri", "declare where it lives"))
		}
		// A uri is bound through its hash only; resolving it needs a network, which the
		// gate path never has.
	}
	for _, r := range items {
		label := "evidence " + r.Decl.Kind + "/" + r.Decl.Job
		switch {
		case !r.Decl.Pending():
			verify(label, r.Decl.SHA256, r.Decl.Path, r.Decl.URI)
		case r.Attached == nil:
			pending = true
		default:
			verify(label, r.Attached.SHA256, r.Attached.Path, r.Attached.URI)
		}
	}
	ch := result(fs)
	if pending && len(fs) == 0 {
		ch.Result = "pending"
	}
	return ch
}

func build(c Ctx) model.Check {
	items, err := Collect(c)
	if err != nil {
		return result(nil)
	}
	var fs []model.Finding
	pending := false
	for _, r := range items {
		if r.Decl.Kind != "build" {
			continue
		}
		res := r.Decl.Result
		if r.Decl.Pending() {
			if r.Attached == nil {
				pending = true
				continue
			}
			res = r.Attached.Result
		}
		if res != "pass" {
			fs = append(fs, finding(c.phaseRel(c.Phase)+"/output.md", "build "+r.Decl.Job+" did not succeed", "fix the build and let the pipeline produce a new result"))
		}
	}
	ch := result(fs)
	if pending && len(fs) == 0 {
		ch.Result = "pending"
	}
	return ch
}
