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
	"strings"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/template"
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

// Applicable is the set of gates that apply at a phase, in the order of the table above,
// which is the table in section 7 of the process definition. A verdict carries exactly
// these, and the one caller that does not run gates reads it from here rather than from a
// list of its own: two lists of applicability would answer differently the first time a
// gate's From column moved.
func Applicable(phase string) []string {
	idx := model.PhaseIndex(phase)
	if idx < 0 {
		return nil
	}
	var ids []string
	for _, s := range table {
		if idx >= s.from {
			ids = append(ids, s.id)
		}
	}
	return ids
}

// Invariants is what the routing through carryForward is supposed to guarantee, checked
// rather than promised. A25 recorded the promise: every finding reaches a verdict through
// that one function, so every id is the hash the appendix defines and no external finding
// carries a decision. Nothing enforced it, and a test of carryForward cannot, because what
// would break the rule is a second path into the verdict rather than a change to the
// function.
//
// So the verdict is checked instead of the path. A finding with no id, an id that is not
// the hash of what it reports, or a decision on foreign code fails the run where it is
// produced, whichever code produced it.
//
// The rule id is empty until WP4 brings rules; when a finding carries one, it enters the
// hash here exactly as it does in carryForward, which is why both read it from the same
// place.
func Invariants(checks []model.Check) error {
	for _, ch := range checks {
		for _, f := range ch.Findings {
			switch {
			case f.ID == "":
				return fmt.Errorf("%s produced a finding with no id (%s): every finding is routed through carryForward",
					ch.Gate, f.Cause)
			case f.ID != hashing.FindingID(ch.Gate, "", f.File, f.Cause):
				return fmt.Errorf("%s finding %s does not carry the id its content hashes to: an id is derived and never written",
					ch.Gate, f.ID)
			case ch.Provenance == ExternalProvenance && f.Decision != nil:
				return fmt.Errorf("%s finding %s is external and carries a decision: a release on foreign wording is taken again, not kept",
					ch.Gate, f.ID)
			}
		}
	}
	return nil
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

// oneOf is the closed set test the specification's enumerations need.
func oneOf(v string, set []string) bool {
	for _, s := range set {
		if s == v {
			return true
		}
	}
	return false
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

// A hash field carries sixty four lowercase hex characters or the placeholder, per
// Appendix B, and where it carries a hash the hash is recomputed and compared. Both
// belong here rather than in G-Freshness: what is checked is a field against the file it
// names, inside one phase, where G-Freshness compares a phase against its predecessor.
// The placeholder is honest in two cases and wrong in a third, which is why
// this cannot be one regexp: `secrets_hash` and `rules_hash` have no writer yet, and an
// artifact with `tool: manual` was produced by nothing at all, so both may say `by-hand`.
// In `context_hash` or `strings_hash` of an artifact a session produced, a writer exists
// and the value was skipped.
var (
	hashShape   = regexp.MustCompile(`^[0-9a-f]{64}$`)
	hashFields  = []string{"context_hash", "secrets_hash", "strings_hash", "rules_hash"}
	writerless  = map[string]bool{"secrets_hash": true, "rules_hash": true}
	placeholder = "by-hand"
)

// language is the artifact's own, because the bundle a phase rendered from is the one in
// that language and no other.
func language(raw map[string]any) string {
	if l, ok := raw["language"].(string); ok && l != "" {
		return l
	}
	return "en"
}

func hashes(c Ctx, file string, raw map[string]any) []model.Finding {
	// Nothing produced a manual artifact, so none of its hashes had a writer.
	byHand := raw["tool"] == "manual"
	// A bundle the repository no longer carries cannot be hashed by anybody. The
	// artifact names the version it rendered from, and where that is not the version
	// here, the value is unrecoverable rather than skipped.
	goneBundle := false
	if ref, ok := raw["template"].(string); ok {
		if t, err := template.Load(c.Root, model.TemplateID(c.Phase), language(raw)); err == nil {
			goneBundle = ref != t.Ref()
		}
	}
	var fs []model.Finding
	for _, f := range hashFields {
		v, ok := raw[f].(string)
		if !ok || v == "" {
			continue // absence is the missing field finding, not this one
		}
		honest := byHand || writerless[f] || (f == "strings_hash" && goneBundle)
		switch {
		case hashShape.MatchString(v):
			fs = append(fs, recomputed(c, file, f, v, goneBundle, raw)...)
		case v == placeholder && honest:
		case v == placeholder:
			fs = append(fs, finding(file, f+" says "+placeholder+" where a writer exists",
				"write the hash Appendix B defines; the placeholder is for a field nothing writes yet"))
		default:
			fs = append(fs, finding(file, f+" is neither a sha256 nor "+placeholder,
				"write sixty four lowercase hex characters, as Appendix B defines it"))
		}
	}
	return fs
}

// recomputed compares a value against the file it covers. Two of the four can be
// recomputed: `context_hash` against the lock file beside the artifact, and
// `strings_hash` against the bundle the phase rendered from, where that bundle is still
// the one the repository carries. The other two have no writer and therefore nothing to
// recompute against, which Appendix B says outright.
func recomputed(c Ctx, file, field, value string, goneBundle bool, raw map[string]any) []model.Finding {
	var want, covers string
	switch {
	case field == "context_hash":
		covers = c.phaseRel(c.Phase) + "/context.lock.yaml"
		h, err := hashing.FileHash(c.abs(covers))
		if err != nil {
			return nil // the file's absence is G-Freshness's finding, not a hash mismatch
		}
		want = h
	case field == "strings_hash" && !goneBundle:
		t, err := template.Load(c.Root, model.TemplateID(c.Phase), language(raw))
		if err != nil {
			return nil // an unresolvable template is the section writer's finding
		}
		covers = "the strings bundle of " + t.Ref()
		want = t.StringsHash
	default:
		return nil
	}
	if value == want {
		return nil
	}
	return []model.Finding{finding(file, field+" does not match "+covers,
		"recompute it as Appendix B defines, or write the artifact from the file it names")}
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
		fs = append(fs, hashes(c, out, raw)...)
		fs = append(fs, questionShape(out, o)...)
		fs = append(fs, decisionShape(out, o)...)
	}

	// Section 4 lists digest.md among the files every phase directory holds, and no
	// other gate looks for it: G-Learning guards learning.yaml, G-Freshness guards
	// context.lock.yaml from P1, and cost.yaml is deliberately unguarded because it may
	// arrive after the gate ran. Without this check a phase finished without its digest
	// is green, which is a verdict on a phase that is not complete.
	dig := dir + "/digest.md"
	if !fm.Exists(c.abs(dig)) {
		fs = append(fs, finding(dig, "digest.md is missing", "write the digest of the session; a phase without one is incomplete"))
	} else {
		r, _ := fm.ReadFront(c.abs(dig), nil)
		for _, f := range missing(r, commonFields, sessionFields) {
			fs = append(fs, finding(dig, "required field missing: "+f, "add "+f+" to the frontmatter"))
		}
		fs = append(fs, schemaVersion(dig, r)...)
		fs = append(fs, hashes(c, dig, r)...)
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

// Section 8: every open assumption turns the gate of its phase red. The state is in
// status, so a rejected assumption is decided rather than unconfirmed, which is the
// difference an empty confirmed_by cannot express.
func assumptions(c Ctx) model.Check {
	rel := model.IntentDir(c.Key) + "/assumptions.yaml"
	var fs []model.Finding
	for _, a := range readAssumptions(c).Assumptions {
		switch {
		case a.Open():
			fs = append(fs, finding(rel, "assumption "+a.ID+" is open", "confirm it, reject it, or replace it with a decision"))
		case a.DecidedBy() == "":
			fs = append(fs, finding(rel, "assumption "+a.ID+" is "+a.Status+" by nobody", "name who decided it: section 8 gives each decided state its person"))
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
		if a.Resolves != "" && a.Status == "confirmed" {
			resolved[a.Resolves] = true // section 8's second exit: somebody accepts a placeholder
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
	return result(learningFindings(c.abs(c.phaseRel(c.Phase)+"/learning.yaml"), c.phaseRel(c.Phase)+"/learning.yaml"))
}

// The shape section 10 defines: `category` from a closed set, and an entry that carries
// an observation, a proposal and a target. A record that says nothing says `no_finding`,
// which is the honest empty case and not a missing one.
var (
	learningCategories = []string{"template", "prompt", "context-rule", "project-convention"}
	learningKeys       = []string{"category", "observation", "proposal", "target"}
	learningHeader     = []string{"intent", "phase", "created", "schema_version", "runner_version", "plugin_version"}
)

// learningFindings is the same judgement wherever a learning record sits: present, saying
// something beyond the header fields, and shaped as section 10 defines. Counting keys
// was A7's gap: a record could invent `observations:` with a category of its own and pass,
// which is how the gap was found rather than how it was predicted.
func learningFindings(abs, rel string) []model.Finding {
	raw := map[string]any{}
	err := fm.ReadYAML(abs, &raw)
	if os.IsNotExist(err) {
		return []model.Finding{finding(rel, "learning record missing", "write one, even when the result is no finding")}
	}
	var fs []model.Finding
	body := false
	for k := range raw {
		if oneOf(k, learningHeader) {
			continue
		}
		body = true
		switch k {
		case "no_finding", "learnings":
		default:
			fs = append(fs, finding(rel, "unknown key "+k+" in the learning record",
				"section 10 defines learnings and no_finding; anything else is not read by anybody"))
		}
	}
	if !body {
		return []model.Finding{finding(rel, "learning record is empty", "record an observation or state no_finding: true")}
	}
	fs = append(fs, learningEntries(rel, raw["learnings"])...)
	sort.Slice(fs, func(i, j int) bool { return fs[i].Cause < fs[j].Cause })
	return fs
}

// learningEntries judges the list itself. An entry is a mapping with four keys, and the
// category is one of four words, because a free category is a word the rule set cannot
// route and the merge request section 10 asks for would have nowhere to go.
func learningEntries(rel string, v any) []model.Finding {
	if v == nil {
		return nil
	}
	list, ok := v.([]any)
	if !ok {
		return []model.Finding{finding(rel, "learnings is not a list",
			"section 10 defines it as a list of entries, each with category, observation, proposal and target")}
	}
	var fs []model.Finding
	for i, e := range list {
		at := fmt.Sprintf("entry %d", i+1)
		entry, ok := e.(map[string]any)
		if !ok {
			fs = append(fs, finding(rel, at+" is not a mapping",
				"write category, observation, proposal and target under it"))
			continue
		}
		for _, k := range learningKeys {
			if s, isString := entry[k].(string); !isString || strings.TrimSpace(s) == "" {
				fs = append(fs, finding(rel, at+" has no "+k,
					"section 10 defines all four; a record missing one is not a proposal anybody can act on"))
			}
		}
		for k := range entry {
			if !oneOf(k, learningKeys) {
				fs = append(fs, finding(rel, at+" carries unknown key "+k,
					"section 10 defines category, observation, proposal and target"))
			}
		}
		if cat, isString := entry["category"].(string); isString && cat != "" && !oneOf(cat, learningCategories) {
			fs = append(fs, finding(rel, at+" has category "+cat,
				"section 10 fixes the set: "+strings.Join(learningCategories, ", ")))
		}
	}
	return fs
}

// CompleteOnClose is G-Complete in its second mode, the one an abandoned intent meets.
// It has two invocation points because an intent has two ways of ending, and the mode
// follows from where the gate was invoked rather than from a field.
//
// Run as part of P5 it checks the preceding phases; that mode is in the table above and
// is not implemented, since no intent in this repository reaches P5 yet. Run from
// `xeno intent close` it checks what the process definition names for this mode: that
// the intent carries a reason, and that the closing learning record exists.
//
// Why a reason has to be checked at all, when the command requires one: the command is
// not the only way a file gets written, and a gate that trusts the writer checks
// nothing.
func CompleteOnClose(root, key string) model.Check {
	var fs []model.Finding
	rel := model.IntentDir(key) + "/intent.yaml"
	var in model.Intent
	if err := fm.ReadYAML(filepath.Join(root, rel), &in); err != nil {
		fs = append(fs, finding(rel, "intent.yaml cannot be read: "+err.Error(), "repair it"))
	} else {
		if in.Status != "abandoned" {
			fs = append(fs, finding(rel, "status is "+in.Status+", not abandoned", "an intent that closes this way is abandoned"))
		}
		if strings.TrimSpace(in.Reason) == "" {
			fs = append(fs, finding(rel, "abandoned without a reason", "why something was dropped is usually worth more than why it was built"))
		}
	}
	lrel := model.IntentDir(key) + "/learning.yaml"
	fs = append(fs, learningFindings(filepath.Join(root, lrel), lrel)...)
	ch := result(fs)
	ch.Gate = "G-Complete"
	return ch
}

// ---- G-Freshness, both halves: the context hash against the predecessor, and the files a
// preceding phase was given against the tree.

func freshness(c Ctx) model.Check {
	idx := model.PhaseIndex(c.Phase)
	var fs []model.Finding
	if idx > 0 {
		rel := c.phaseRel(c.Phase) + "/context.lock.yaml"
		var lock model.ContextLock
		if err := fm.ReadYAML(c.abs(rel), &lock); err != nil {
			return result([]model.Finding{finding(rel, "context lock missing", "start the phase with xeno phase start")})
		}
		cur, err := hashing.DirHash(c.Root, c.phaseRel(model.Phases[idx-1]), hashing.PhaseExcluded)
		if err != nil || cur != lock.PredecessorHash {
			return result([]model.Finding{finding(rel, "predecessor changed after this phase started", "rerun this phase against the current predecessor")})
		}
	}
	return result(append(fs, staleReads(c, idx)...))
}

// staleReads is the second half: no file a preceding phase was given has changed since it
// was given. The lock records the information base with a hash each and is never
// refreshed, so a file whose hash no longer matches the tree is one this phase, or a later
// one, moved out from under an earlier phase's reading.
//
// It looks at this phase and every phase before it, because the question is whether the
// work already done still rests on what it was given, and a phase does not stop being
// stale by having a successor.
//
// Where no profile was written the lists are empty and there is nothing to compare, which
// is every intent in this repository so far: the check is not weaker for it, it has simply
// been told nothing.
func staleReads(c Ctx, idx int) []model.Finding {
	var fs []model.Finding
	for i := 0; i <= idx; i++ {
		phase := model.Phases[i]
		rel := c.phaseRel(phase) + "/context.lock.yaml"
		var lock model.ContextLock
		if err := fm.ReadYAML(c.abs(rel), &lock); err != nil {
			continue // absence is the first half's finding, or a phase that has not run
		}
		for _, f := range lock.Files {
			h, err := hashing.FileHash(c.abs(f.Path))
			switch {
			case os.IsNotExist(err):
				fs = append(fs, finding(rel, phase+" was given "+f.Path+" and it is gone",
					"read the phase again against what is there, or record why the file left"))
			case err != nil:
				fs = append(fs, finding(rel, phase+" was given "+f.Path+" and it cannot be read: "+err.Error(),
					"make it readable, or read the phase again against what is there"))
			case h != f.SHA256:
				fs = append(fs, finding(rel, phase+" was given "+f.Path+" and it has changed since",
					"read the phase again for what changed; the lock records what it was given, not what is there now"))
			}
		}
	}
	sort.Slice(fs, func(i, j int) bool { return fs[i].Cause < fs[j].Cause })
	return fs
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
