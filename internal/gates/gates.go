// SPDX-License-Identifier: Apache-2.0

// Package gates evaluates the gates of one phase. Everything here is deterministic and free of
// network access, and what that forbids is worth naming rather than generalising: no gate runs
// a build, a test suite or a scanner, and none asks a model. G-Build and G-Test read declared
// results for exactly that reason.
//
// One subprocess is started, by internal/git, and only where a rule asks for it: `git log` over
// the commit range the run was given, which reads the clone that is already there. Section 9's
// commit predicates read a subject, a trailer, a signature status and an author, and a history
// is not something a repository carries in a file a gate could parse instead.
package gates

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/git"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/rules"
	"github.com/triplem/xeno/internal/template"
)

// Ctx is what a gate run knows about.
type Ctx struct {
	Root, Key, Phase string
	ArtifactsHash    string
	QualifiedID      string // from intent.yaml
	// Base and Head are the commit range under review, as the run was given it. Section 12
	// forbids working it out here: a guessed range means different verdicts locally and in CI
	// from the same repository state, so an absent range is a finding and never a default.
	Base, Head string
	// External produces the checks of section 14's external gates, which are commands a project
	// declares. It is supplied rather than called from here, because this package reads and
	// that runs: the one subprocess this package's comment admits is a git log over a local
	// clone, and foreign code with repository access is a different kind of exception. Absent
	// where a project declares none, which is the default and every project today.
	External func(phase string) []model.Check
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
	{"G-Rules", 0, rulesGate},
	{"G-Policy", 0, policy},
	{"G-Complete", 5, completeInReview},
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
	// After the shipped gates, in the order the project declared them. They go through the same
	// carryForward as every other check, which is where a finding gets its id and where an
	// external one is refused a decision: a second path into the verdict that skipped it would
	// make the invariant a comment.
	if c.External != nil {
		for _, ch := range c.External(c.Phase) {
			carryForward(&ch, decided)
			out = append(out, ch)
		}
	}
	return out
}

// Status derives the phase status from the checks. Nothing writes it directly.
// Order, loudest first: red, overridden, provisional, approved, green.
//
// A check reporting fail without carrying a finding is refused rather than turned into a
// status. Red is a promise that something can be decided: a red phase carries a failure
// somebody approves or overrides by a finding's id, and there is no id here, so the verdict
// would name no cause and clear only by editing the file it came from. No gate of this runner
// reaches the state, because every one returns through result(), which writes fail only where
// there are findings; what reaches it is a second writer. Section 14 makes external gates the
// extension point, and a scanner wrapper reporting a failure it cannot attribute to a file is
// ordinary rather than malformed. So is a hand edited gate.yaml, which lies outside
// artifacts_hash by design and comes back through rewriteStatus.
//
// The rule is here and not in Invariants for that last path: Invariants runs from evaluate
// alone, on checks just produced, while every derivation of a status passes through here. The
// condition reads both fields, because pass, pending and not-implemented carry no finding as
// their ordinary state.
func Status(checks []model.Check) (string, error) {
	seen := map[string]bool{}
	undecided, overridden, pending, approved := false, false, false, false
	for _, ch := range checks {
		if ch.Result == "fail" && len(ch.Findings) == 0 {
			return "", fmt.Errorf("%s reports fail and carries no finding: "+
				"a check that failed without saying what failed is not a verdict", ch.Gate)
		}
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

// A hash field carries sixty four lowercase hex characters or the placeholder, per
// Appendix B, and where it carries a hash the hash is recomputed and compared. Both
// belong here rather than in G-Freshness: what is checked is a field against the file it
// names, inside one phase, where G-Freshness compares a phase against its predecessor.
// The placeholder is honest in two cases and wrong in a third, which is why
// this cannot be one regexp: `secrets_hash` and `rules_hash` have no writer yet, and an
// artifact with `tool: manual` was produced by nothing at all, so both may say `by-hand`.
// In `context_hash` or `strings_hash` of an artifact a session produced, a writer exists
// and the value was skipped.
var hashShape = regexp.MustCompile(`^[0-9a-f]{64}$`)

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
	for _, f := range model.HashFields {
		v, ok := raw[f].(string)
		if !ok || v == "" {
			continue // absence is the missing field finding, not this one
		}
		honest := byHand || model.OneOf(f, model.WriterlessHash) || (f == "strings_hash" && goneBundle)
		switch {
		case hashShape.MatchString(v):
			fs = append(fs, recomputed(c, file, f, v, goneBundle, raw)...)
		case v == model.HashPlaceholder && honest:
		case v == model.HashPlaceholder:
			fs = append(fs, finding(file, f+" says "+model.HashPlaceholder+" where a writer exists",
				"write the hash Appendix B defines; the placeholder is for a field nothing writes yet"))
		default:
			fs = append(fs, finding(file, f+" is neither a sha256 nor "+model.HashPlaceholder,
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

// links reports a declared link whose document is not in the tree.
//
// It is the profile's one unambiguous error. Everything else in the file is a pattern, and a
// pattern that matches nothing is a state rather than a mistake — the files may not be written
// yet. A link is a specific path, declared because section 5 forbids inferring one: "an inferred
// mapping is an assumption, and assumptions in this process are either registered or absent". A
// declaration whose target does not exist is neither.
//
// The finding names the profile rather than the lock, because the profile is the claim and the
// base is the consequence: a finding on the lock would point a reader at a file that is correct
// about what it was given. The runner, which writes that lock, skips such a link in silence and is
// right to — it records what the phase was given, and a file that is not there was not given.
//
// Checked wherever the profile is read, which is every phase, as the budget check established. A
// configuration error should keep being reported until it is corrected, where a check at P0 alone
// would report it once into a verdict nobody re-reads. #171 asked for this and did not do it; #172
// is it.
func links(c Ctx) []model.Finding {
	var p model.Profile
	profile := c.phaseRel(model.Phases[0]) + "/" + model.ContextProfile
	if err := fm.ReadYAML(c.abs(profile), &p); err != nil {
		return nil // no profile is no link
	}
	var fs []model.Finding
	for _, l := range p.Links {
		if l.Docs == "" {
			continue // a link with no document declares nothing to find
		}
		if _, err := os.Stat(c.abs(l.Docs)); err != nil {
			at := l.Component
			if at == "" {
				at = "a link"
			}
			fs = append(fs, finding(profile,
				"the link for "+at+" names "+quoted(l.Docs)+", which is not in the tree",
				"correct the path, or take the link out; a declared link is a claim about a file"))
		}
	}
	return fs
}

// budget reports a recorded context that exceeded the budget its profile declared.
//
// Section 5 puts this check here and says what it is for in the same breath: "deliberately a
// finding and not a red gate in the sense of stopping work", because "blocking against a number
// nobody has experience with yet would be the wrong way round", and because what it prevents is
// the budget quietly becoming decoration. So it is a finding like any other — visible, decidable,
// recorded — and it names both numbers, since a finding that says only "over budget" is one
// nobody can act on.
//
// The comparison is between two files: the profile declares the budget and the lock records what
// the phase was given. Bytes come from the tree when the check runs rather than from the lock,
// which carries paths and hashes and no sizes; a file the lock names and the tree has lost is
// skipped, because its absence is G-Freshness's finding and one cause reported by two gates is
// what #160 avoided.
func budget(c Ctx) []model.Finding {
	var p model.Profile
	profile := c.phaseRel(model.Phases[0]) + "/" + model.ContextProfile
	if err := fm.ReadYAML(c.abs(profile), &p); err != nil {
		return nil // no profile is no budget
	}
	if p.Budget.Files == 0 && p.Budget.Bytes == 0 {
		return nil
	}
	rel := c.phaseRel(c.Phase) + "/context.lock.yaml"
	var lock model.ContextLock
	if err := fm.ReadYAML(c.abs(rel), &lock); err != nil {
		return nil // the lock's absence is G-Freshness's finding
	}
	var fs []model.Finding
	if p.Budget.Files > 0 && len(lock.Files) > p.Budget.Files {
		fs = append(fs, finding(rel,
			fmt.Sprintf("the recorded context is %d files and the budget is %d",
				len(lock.Files), p.Budget.Files),
			"narrow the profile's include, or raise the budget in "+profile+" and say why"))
	}
	if p.Budget.Bytes > 0 {
		var total int
		for _, f := range lock.Files {
			if info, err := os.Stat(c.abs(f.Path)); err == nil {
				total += int(info.Size())
			}
		}
		if total > p.Budget.Bytes {
			fs = append(fs, finding(rel,
				fmt.Sprintf("the recorded context is %d bytes and the budget is %d", total, p.Budget.Bytes),
				"narrow the profile's include, or raise the budget in "+profile+" and say why"))
		}
	}
	return fs
}

func schema(c Ctx) model.Check {
	dir := c.phaseRel(c.Phase)
	o, fs := phaseResult(c, dir)
	fs = append(fs, digestFindings(c, dir)...)
	fs = append(fs, yamlFindings(c, dir)...)
	fs = append(fs, directoryFindings(c, dir)...)
	fs = append(fs, undeclaredEvidence(c, dir, o)...)
	fs = append(fs, budget(c)...)
	fs = append(fs, links(c)...)
	return result(fs)
}

// phaseResult judges output.md and returns it, because three of the four checks below need
// nothing from it and the fourth, the undeclared evidence, needs all of it.
func phaseResult(c Ctx, dir string) (model.Output, []model.Finding) {
	out := dir + "/output.md"
	var o model.Output
	raw, err := fm.ReadFront(c.abs(out), &o)
	switch {
	case os.IsNotExist(err):
		return o, []model.Finding{finding(out, "output.md is missing",
			"write the phase result through the MCP operation or its command")}
	case err != nil:
		return o, []model.Finding{finding(out, "frontmatter unreadable: "+err.Error(), "repair the frontmatter")}
	}
	var fs []model.Finding
	for _, f := range missing(raw, commonFields, sessionFields, renderedFields) {
		fs = append(fs, finding(out, "required field missing: "+f, "add "+f+" to the frontmatter"))
	}
	fs = append(fs, schemaVersion(out, raw)...)
	fs = append(fs, hashes(c, out, raw)...)
	fs = append(fs, questionShape(out, o)...)
	fs = append(fs, decisionShape(out, o)...)
	fs = append(fs, evidenceShape(out, o)...)
	return o, fs
}

// evidenceShape judges a declaration against the two closed sets of section 4. It is here
// and not in G-Evidence because this is shape, read off one file, and G-Evidence resolves
// content: a gate that reads a result should not also be deciding whether the word is one
// the document allows.
//
// `result` is required on `test-report` and `build-log` and written elsewhere only where
// the producer reports against a threshold, so a missing one is a finding on those two
// kinds alone. Demanding it everywhere would contradict the sentence that lets a bill of
// materials report nothing, and limitation 8 of section 16 rests on exactly that
// asymmetry: the omission shows on the two kinds that carry it and nowhere else.
//
// A pending item is exempt from that requirement and not from the set. Section 4 says an
// item a pipeline has yet to produce "declares only its kind and its job", so it has no
// result to carry: the run that would report one has not happened, and the value arrives
// in `evidence/attached.yaml` as the job's own verdict. Requiring it here would make every
// declaration of future evidence a finding, which is the state P4 is designed to be in
// between its finish and its pipeline. The kind is known at declaration time either way.
func evidenceShape(file string, o model.Output) []model.Finding {
	var fs []model.Finding
	for _, e := range o.Evidence {
		at := "evidence item " + e.Kind + "/" + e.Job
		if !model.OneOf(e.Kind, model.EvidenceKinds) {
			fs = append(fs, finding(file, at+" has a kind section 4 does not define",
				"section 4 fixes the set: "+strings.Join(model.EvidenceKinds, ", ")))
		}
		switch {
		case e.Result == "":
			if model.OneOf(e.Kind, model.ResultRequiredKinds) && !e.Pending() {
				fs = append(fs, finding(file, at+" carries no result",
					"section 4 requires it on "+strings.Join(model.ResultRequiredKinds, " and ")+
						", because G-Test and G-Build read it"))
			}
		case !model.OneOf(e.Result, model.EvidenceResults):
			fs = append(fs, finding(file, at+" has result "+e.Result,
				"section 4 fixes the set: "+strings.Join(model.EvidenceResults, ", ")+
					"; it is what the run reported against its own threshold"))
		}
	}
	return fs
}

// digestFindings guards digest.md. Section 4 lists it among the files every phase holds and
// no other gate looks for it: G-Learning guards learning.yaml, G-Freshness guards
// context.lock.yaml from P1, and cost.yaml is deliberately unguarded because it may arrive
// after the gate ran. Without this a phase finished without its digest is green, which is a
// verdict on a phase that is not complete.
func digestFindings(c Ctx, dir string) []model.Finding {
	dig := dir + "/digest.md"
	if !fm.Exists(c.abs(dig)) {
		return []model.Finding{finding(dig, "digest.md is missing",
			"write the digest of the session; a phase without one is incomplete")}
	}
	r, _ := fm.ReadFront(c.abs(dig), nil)
	var fs []model.Finding
	for _, f := range missing(r, commonFields, sessionFields) {
		fs = append(fs, finding(dig, "required field missing: "+f, "add "+f+" to the frontmatter"))
	}
	fs = append(fs, schemaVersion(dig, r)...)
	return append(fs, hashes(c, dig, r)...)
}

// yamlFindings checks the phase's YAML artifacts. Section 5 states its field sets for
// frontmatter and for the equivalent top level keys alike, so the YAML ones are checked
// here rather than left to G-Trace, which looks at two fields and judges binding rather
// than shape.
func yamlFindings(c Ctx, dir string) []model.Finding {
	var fs []model.Finding
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
	return fs
}

// directoryFindings reports what is in the phase directory and should not be. Every file a
// hash covers is one Xeno wrote, which is the property Appendix B's normalisation rests on,
// so anything else is a finding before it reaches a hash.
func directoryFindings(c Ctx, dir string) []model.Finding {
	entries, _ := os.ReadDir(c.abs(dir))
	var fs []model.Finding
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			if name != "evidence" {
				fs = append(fs, finding(dir+"/"+name, "unknown directory in the phase directory",
					"remove it or move it outside .xeno/"))
			}
			continue
		}
		if model.KnownPhaseFiles[name] || (name == model.ContextProfile && c.Phase == model.Phases[0]) {
			continue
		}
		fs = append(fs, finding(dir+"/"+name, "unknown file in the phase directory",
			"remove it; only files Xeno writes belong here"))
	}
	return fs
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
//
// Both sides of the comparison are paths relative to evidence/: a declaration carries one
// relative to the phase directory, which the prefix below removes, and the walk produces
// one per file. Section 4 forbids no subdirectory under evidence/, and a set keyed on
// basenames agrees with the paths only while the directory is flat, which is the shape
// Attach happens to write and not a property of the check. Two items of the same name in
// different directories would then stand in for each other.
//
// The walk skips directories rather than reporting them, because a directory is not
// something a declaration can name; the files in it are judged one by one.
func undeclaredEvidence(c Ctx, dir string, o model.Output) []model.Finding {
	evDir := dir + "/evidence"
	root := c.abs(evDir)
	if s, err := os.Stat(root); err != nil || !s.IsDir() {
		return nil
	}
	known := map[string]bool{"attached.yaml": true}
	declare := func(p string) {
		if rel, ok := insideEvidence(p); ok {
			known[rel] = true
		}
	}
	for _, e := range o.Evidence {
		declare(e.Path)
	}
	var att []model.Attached
	_ = fm.ReadYAML(c.abs(evDir+"/attached.yaml"), &att)
	for _, a := range att {
		declare(a.Path)
	}
	var fs []model.Finding
	_ = filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, rerr := filepath.Rel(root, p)
		if rerr != nil {
			return nil
		}
		rel = filepath.ToSlash(rel)
		if !known[rel] {
			fs = append(fs, finding(evDir+"/"+rel, "undeclared file in evidence/", "declare it in output.md or remove it"))
		}
		return nil
	})
	return fs
}

// insideEvidence turns a declared path into the key this check compares, the path relative
// to evidence/. A path that names nothing inside evidence/ yields no key rather than a
// meaningless one: an empty path gave filepath.Base the value ".", which matched no entry
// and sat in the set saying nothing, and a path pointing elsewhere in the phase directory
// declares nothing here either.
func insideEvidence(p string) (string, bool) {
	rel := strings.TrimPrefix(filepath.ToSlash(p), "evidence/")
	if rel == "" || rel == p || strings.HasPrefix(rel, "../") {
		return "", false
	}
	return rel, true
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
// The header fields every process file carries, which is what a learning record has to
// say something beyond. The rest of section 10's shape lives in internal/model.
var learningHeader = []string{"intent", "phase", "created", "schema_version", "runner_version", "plugin_version"}

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
		if model.OneOf(k, learningHeader) {
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
		for _, k := range model.LearningKeys {
			if s, isString := entry[k].(string); !isString || strings.TrimSpace(s) == "" {
				fs = append(fs, finding(rel, at+" has no "+k,
					"section 10 defines all four; a record missing one is not a proposal anybody can act on"))
			}
		}
		for k := range entry {
			if !model.OneOf(k, model.LearningKeys) {
				fs = append(fs, finding(rel, at+" carries unknown key "+k,
					"section 10 defines category, observation, proposal and target"))
			}
		}
		if cat, isString := entry["category"].(string); isString && cat != "" && !model.OneOf(cat, model.LearningCategories) {
			fs = append(fs, finding(rel, at+" has category "+cat,
				"section 10 fixes the set: "+strings.Join(model.LearningCategories, ", ")))
		}
	}
	return fs
}

// completeInReview is G-Complete in the mode a merging intent meets, as part of P5.
// Section 7: it checks that the artifacts of all preceding phases are present and green,
// approved or overridden, and nothing about the merge itself, which has not happened when
// the gate runs.
//
// A verdict is what "present" means here. An artifact without one was never judged, and a
// phase judged provisional is waiting for evidence rather than decided, so it is reported
// with the word the verdict carries: a red phase, a provisional one and a phase nobody ran
// are three different repairs.
//
// It reads the committed verdicts rather than recomputing them. Recomputation is
// `gate verify`'s work and it covers every phase anyway, so doing it again here would
// report the same divergence twice under a different name, and a gate that recomputed its
// predecessors would also be judging what another gate already judged.
func completeInReview(c Ctx) model.Check {
	var fs []model.Finding
	for _, phase := range model.Phases[:model.PhaseIndex(c.Phase)] {
		rel := c.phaseRel(phase) + "/gate.yaml"
		var g model.Gate
		if err := fm.ReadYAML(c.abs(rel), &g); err != nil {
			fs = append(fs, finding(rel, phase+" holds no verdict",
				"run xeno phase finish for it; an intent reaches review with every phase judged"))
			continue
		}
		switch g.Status {
		case "green", "approved", "overridden":
		case "provisional":
			fs = append(fs, finding(rel, phase+" is provisional, waiting for evidence",
				"attach what the pipeline produced and judge it again"))
		default:
			fs = append(fs, finding(rel, phase+" is "+g.Status,
				"decide every failing finding in it: fix it, approve it or override it"))
		}
	}
	return result(fs)
}

// intentDirectoryFindings reports what is in the intent directory and should not be. It is the
// twin of directoryFindings one level up, and it exists for the same reason: Appendix B rests the
// normalisation of a hash on every file it covers being one Xeno wrote, and calls that a checked
// property rather than an assumption. The phase level had the check and the intent level hash,
// added later, had none (#109).
//
// Directories are reported although DirHash does not descend into one, so a stray directory cannot
// change the value. A directory nobody wrote is where files appear next, and reporting it is the
// cheaper half of the same rule.
//
// An unreadable directory reports nothing here, as at the phase level: CompleteOnClose is about to
// read intent.yaml out of it and will say so.
func intentDirectoryFindings(root, key string) []model.Finding {
	dir := model.IntentDir(key)
	entries, _ := os.ReadDir(filepath.Join(root, dir))
	var fs []model.Finding
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			if name != model.PhasesDir {
				fs = append(fs, finding(dir+"/"+name, "unknown directory in the intent directory",
					"remove it or move it outside .xeno/"))
			}
			continue
		}
		if model.KnownIntentFiles[name] {
			continue
		}
		fs = append(fs, finding(dir+"/"+name, "unknown file in the intent directory",
			"remove it; only files Xeno writes belong here"))
	}
	return fs
}

// CompleteOnClose is G-Complete in its second mode, the one an abandoned intent meets.
// It has two invocation points because an intent has two ways of ending, and the mode
// follows from where the gate was invoked rather than from a field.
//
// Run from `xeno intent close` it checks what the process definition names for this mode:
// that the intent carries a reason, and that the closing learning record exists.
//
// Why a reason has to be checked at all, when the command requires one: the command is
// not the only way a file gets written, and a gate that trusts the writer checks
// nothing.
// Its findings are routed through carryForward before they are returned, as Run does for every
// phase check, because A25's promise is that every finding reaches a verdict through that one
// function and Invariants exists to check the promise rather than state it. This path was written
// after the invariant and met neither for as long as it produced at most one finding, whose empty
// id collided with nothing (#138).
func CompleteOnClose(root, key string) model.Check {
	// First, so that a verdict lists what the directory contained before what the intent
	// asserted: the hash this verdict carries covers that directory.
	fs := intentDirectoryFindings(root, key)
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
	// A decision taken on an intent level finding survives a re-close, the way a decision on a
	// phase finding survives a re-run, and for the same reason: the id is a hash over what the
	// finding reports, so the same finding is the same id.
	carryForward(&ch, intentDecisions(root, key))
	return ch
}

// intentDecisions reads the decisions the previous intent level verdict held, which is what Run
// does from the previous phase verdict.
func intentDecisions(root, key string) map[string]*model.DecisionOnFinding {
	decided := map[string]*model.DecisionOnFinding{}
	var prev model.Gate
	if err := fm.ReadYAML(filepath.Join(root, model.IntentDir(key), "gate.yaml"), &prev); err != nil {
		return decided
	}
	for _, ch := range prev.Checks {
		for _, f := range ch.Findings {
			if f.Decision != nil {
				decided[f.ID] = f.Decision
			}
		}
	}
	return decided
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
	// where names the file a finding belongs in. A declaration's failure is output.md's,
	// and an attachment's is attached.yaml's, which is the file somebody would have had to
	// edit for it to be in that state.
	attachedRel := dir + "/evidence/attached.yaml"
	verify := func(label, where, sum, path, uri string) {
		switch {
		case path != "":
			got, err := hashing.FileHash(c.abs(dir + "/" + path))
			if err != nil {
				fs = append(fs, finding(dir+"/"+path, label+" does not resolve", "restore the file or correct the declaration"))
			} else if got != sum {
				fs = append(fs, finding(dir+"/"+path, label+" content does not match its hash", "the file changed after it was declared"))
			}
		case uri == "":
			fs = append(fs, finding(where, label+" has neither path nor uri", "declare where it lives"))
		case sum == "":
			// A uri is bound through its hash and nothing else, so one without a hash is
			// bound by nothing: the verdict would record that evidence arrived without
			// being able to say what arrived. The attach declines to write this, and it is
			// checked here because attached.yaml lies outside the artifacts_hash by
			// design, so it is the one file in a judged phase that can be edited without
			// making any verdict stale. What the missing seal costs is paid by this check.
			fs = append(fs, finding(where, label+" has a uri and no hash, so nothing binds it",
				"record the sha256 the pipeline published, or let the attach decline it and republish"))
		}
		// Whether the bytes behind a uri still match is not asked: resolving one needs a
		// network, which the gate path never has.
	}
	for _, r := range items {
		label := "evidence " + r.Decl.Kind + "/" + r.Decl.Job
		switch {
		case !r.Decl.Pending():
			verify(label, dir+"/output.md", r.Decl.SHA256, r.Decl.Path, r.Decl.URI)
		case r.Attached == nil:
			pending = true
		default:
			verify(label, attachedRel, r.Attached.SHA256, r.Attached.Path, r.Attached.URI)
		}
	}
	ch := result(fs)
	if pending && len(fs) == 0 {
		ch.Result = "pending"
	}
	return ch
}

// BuildKind is the kind G-Build reads, as section 4 spells it. Named rather than written
// into the comparison below because the gate read `build` for as long as nothing checked
// the closed set, and matched no conformant declaration the whole time: a declared build
// with `result: fail` passed. The set now lives in one place and G-Schema judges a value
// against it, so a spelling nobody defined cannot reach this gate again.
const BuildKind = "build-log"

func build(c Ctx) model.Check {
	items, err := Collect(c)
	if err != nil {
		return result(nil)
	}
	var fs []model.Finding
	pending := false
	for _, r := range items {
		if r.Decl.Kind != BuildKind {
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

// ---- G-Rules

// rulesGate is section 7's row: rule collisions resolved, binding respected, scope matches the
// path. It reads the tree, resolves it, and reports what neither step could use.
//
// The gate words what internal/rules returns and judges nothing of its own, so that G-Policy
// resolving the same tree cannot disagree with it about what the set is. Section 7 puts the
// two together at the end of the table for that reason.
//
// A repository with no rule tree passes. That is every repository before a plugin is vendored
// and every project that maintains no rules, and a project without rules does not have a
// broken rule set.
func rulesGate(c Ctx) model.Check {
	read, problems := rules.Load(c.Root)
	_, collisions := rules.Effective(read)
	var fs []model.Finding
	for _, p := range append(problems, collisions...) {
		fs = append(fs, finding(p.Path, p.Cause, p.Next))
	}
	return result(fs)
}

// ---- G-Policy

// predicateFn is what a check of a given type is evaluated by. The registry is the five names
// section 9 fixes and it is a budget rather than an extension point: a project naming a type of
// its own gets the finding below, because project-defined predicates are outside v1 and an
// external gate covers the same ground.
type predicateFn func(*eval, rules.Rule) []model.Finding

var predicates = map[string]predicateFn{
	"section-implies-section": sectionImpliesSection,
	"section-non-empty":       sectionNonEmpty,
	"commit-message":          commitMessage,
	"commit-trailer":          commitTrailer,
	"commit-signature":        commitSignature,
	"approver-not-author":     approverNotAuthor,
}

// eval is one gate run's evaluation of the checked rules. It holds the range so that five rules
// naming a commit type read one `git log` rather than five, and it holds the error with it, so a
// range that does not resolve is reported once per rule that needed it and read once.
type eval struct {
	Ctx
	commits []git.Commit
	err     error
	loaded  bool
}

// commitsOf loads the range on the first predicate that asks.
func (e *eval) commitsOf() ([]git.Commit, error) {
	if !e.loaded {
		e.commits, e.err = git.Commits(e.Root, e.Base, e.Head)
		e.loaded = true
	}
	return e.commits, e.err
}

// inRange is the commits a rule judges: the range, less the merge commits where the rule
// exempts them. A rule that does not exempt them judges a merge subject like any other, which
// is what section 9's example makes explicit by carrying the field.
func (e *eval) inRange(r rules.Rule) ([]git.Commit, []model.Finding) {
	cs, err := e.commitsOf()
	if err != nil {
		return nil, []model.Finding{finding(r.Path,
			"rule "+r.ID+" reads the commit range and it could not be read: "+err.Error(),
			"pass --base and --head as refs this repository resolves; the range is an input of the run and is never inferred")}
	}
	if !exempts(r, "merge-commits") {
		return cs, nil
	}
	var out []git.Commit
	for _, c := range cs {
		if !c.IsMerge() {
			out = append(out, c)
		}
	}
	return out, nil
}

// exempts reads a check's exempt list, which section 9 writes as a sequence of names.
func exempts(r rules.Rule, what string) bool {
	if r.Check == nil {
		return false
	}
	list, ok := r.Check.Params["exempt"].([]any)
	if !ok {
		return false
	}
	for _, e := range list {
		if s, ok := e.(string); ok && s == what {
			return true
		}
	}
	return false
}

// param reads one string parameter of a check.
func param(r rules.Rule, key string) string {
	if r.Check == nil {
		return ""
	}
	s, _ := r.Check.Params[key].(string)
	return s
}

// shortHash is how a finding names a commit: the first eight characters, as git itself prints
// them, plus the subject so that a reader recognises the commit without looking it up.
func shortHash(c git.Commit) string {
	h := c.Hash
	if len(h) > 8 {
		h = h[:8]
	}
	return h + " " + quoted(c.Subject)
}

// ---- the five predicate types of section 9

// sectionImpliesSection is the one type that reads an artifact: where the when section is
// non-empty, the then section has to be.
//
// An empty or absent antecedent leaves the rule green whatever the consequent says, because an
// implication with a false antecedent is true. The alternative reading would turn every such
// rule into a requirement that every section of every template be filled, which is not what
// "a change to a published interface must come with a migration note" says.
func sectionImpliesSection(e *eval, r rules.Rule) []model.Finding {
	when, then := sectionParam(r, "when"), sectionParam(r, "then")
	rel := e.phaseRel(e.Phase) + "/output.md"
	if when.section == "" || then.section == "" {
		return []model.Finding{finding(r.Path,
			"rule "+r.ID+" names no section in when or then",
			"write when and then as a section and non_empty, as section 9 does")}
	}
	if fs := unknownSection(e, r, when.section, then.section); fs != nil {
		return fs
	}
	sections, ok := renderedSections(e, rel)
	if !ok {
		return nil // an absent or unreadable artifact is G-Schema's finding
	}
	filled := func(id string) bool { return strings.TrimSpace(sections[id]) != "" }
	if filled(when.section) && !filled(then.section) {
		return []model.Finding{finding(rel,
			"rule "+r.ID+": "+when.section+" is filled and "+then.section+" is empty",
			"write "+then.section+", or record a deviation against the rule")}
	}
	return nil
}

// sectionNonEmpty is the same reader with no antecedent: the section a rule names has to carry
// text. It is the second type in this registry that section 9's table does not list, admitted by
// the same sentence as the first — this work package delivers "the section predicates the shipped
// set needs" — and the shipped set needs it for the filled release notes section at P5. Nothing
// else covers that: a required section carrying nothing is read only by the next-step suggestion
// and by no gate (A73).
func sectionNonEmpty(e *eval, r rules.Rule) []model.Finding {
	name := param(r, "section")
	if name == "" {
		return []model.Finding{finding(r.Path, "rule "+r.ID+" names no section",
			"write section as the id of the section the rule requires")}
	}
	if fs := unknownSection(e, r, name); fs != nil {
		return fs
	}
	rel := e.phaseRel(e.Phase) + "/output.md"
	sections, ok := renderedSections(e, rel)
	if !ok {
		return nil // an absent or unreadable artifact is G-Schema's finding
	}
	if strings.TrimSpace(sections[name]) == "" {
		return []model.Finding{finding(rel,
			"rule "+r.ID+": "+name+" is empty",
			"write "+name+", or record a deviation against the rule")}
	}
	return nil
}

// unknownSection reports a rule naming a section the phase's template does not have, which is a
// configuration error in the rule rather than an empty section: an empty one would leave the rule
// vacuously green for ever. Where the template cannot be loaded at all nothing is reported, which
// is a repository with no vendored plugin and no way to answer the question.
func unknownSection(e *eval, r rules.Rule, names ...string) []model.Finding {
	t, err := template.Load(e.Root, model.TemplateID(e.Phase), "en")
	if err != nil {
		return nil
	}
	for _, name := range names {
		if !t.Has(name) {
			return []model.Finding{finding(r.Path,
				"rule "+r.ID+" names section "+quoted(name)+", which "+t.Ref()+" does not have",
				"name a section of the template the phase renders from: "+strings.Join(t.Known(), ", "))}
		}
	}
	return nil
}

// renderedSections parses the phase's artifact back into sections, which is the form a reader
// sees and the form a section predicate asks about.
func renderedSections(e *eval, rel string) (map[string]string, bool) {
	body, err := os.ReadFile(e.abs(rel))
	if err != nil {
		return nil, false
	}
	_, content, ferr := fm.Split(body)
	if ferr != nil {
		return nil, false
	}
	return template.Parse(string(content)), true
}

// sectionParam reads one half of a section-implies-section check, which section 9 writes as a
// mapping of a section and a non_empty flag. non_empty is read and not acted on: it is the only
// form the specification gives, and a predicate that behaved differently without it would be
// inventing a second form.
type sectionRef struct {
	section  string
	nonEmpty bool
}

func sectionParam(r rules.Rule, key string) sectionRef {
	if r.Check == nil {
		return sectionRef{}
	}
	m, ok := r.Check.Params[key].(map[string]any)
	if !ok {
		return sectionRef{}
	}
	s, _ := m["section"].(string)
	ne, _ := m["non_empty"].(bool)
	return sectionRef{section: s, nonEmpty: ne}
}

// commitMessage judges every subject in the range against a named shipped pattern. The pattern
// is named and never carried as an expression, and the same map answers here and in
// `xeno check commit-message`, so a hook cannot start rejecting what this accepts.
func commitMessage(e *eval, r rules.Rule) []model.Finding {
	name := param(r, "pattern")
	if name == "" {
		return []model.Finding{finding(r.Path, "rule "+r.ID+" names no pattern",
			"name a shipped pattern: "+strings.Join(PatternNames(), ", "))}
	}
	if _, ok := patterns[name]; !ok {
		return []model.Finding{finding(r.Path,
			"rule "+r.ID+" names pattern "+quoted(name)+", which is not shipped",
			"name a shipped pattern: "+strings.Join(PatternNames(), ", "))}
	}
	cs, fs := e.inRange(r)
	if fs != nil {
		return fs
	}
	for _, c := range cs {
		if err := CheckMessage(name, c.Subject); err != nil {
			fs = append(fs, finding(r.Path,
				"rule "+r.ID+": commit "+shortHash(c)+" does not match pattern "+quoted(name),
				"write the subject as "+name+" describes, or exempt the commit if the rule allows it"))
		}
	}
	return fs
}

// commitTrailer requires a trailer on every commit in the range, Xeno-Intent: being the shipped
// case. The key is compared and the value is not read.
func commitTrailer(e *eval, r rules.Rule) []model.Finding {
	key := param(r, "trailer")
	if key == "" {
		return []model.Finding{finding(r.Path, "rule "+r.ID+" names no trailer",
			"name the trailer the rule requires, for instance Xeno-Intent")}
	}
	cs, fs := e.inRange(r)
	if fs != nil {
		return fs
	}
	for _, c := range cs {
		if !c.HasTrailer(key) {
			fs = append(fs, finding(r.Path,
				"rule "+r.ID+": commit "+shortHash(c)+" carries no "+key+" trailer",
				"add the trailer to the commit message; a trailer is a key and a value on their own line at the end"))
		}
	}
	return fs
}

// commitSignature requires a signature that verifies on every commit in the range. What counts
// as verifying is Commit.Signed, and A70 records why an untrusted key is accepted.
func commitSignature(e *eval, r rules.Rule) []model.Finding {
	cs, fs := e.inRange(r)
	if fs != nil {
		return fs
	}
	for _, c := range cs {
		if !c.Signed() {
			fs = append(fs, finding(r.Path,
				"rule "+r.ID+": commit "+shortHash(c)+" carries no signature that verifies",
				"sign the commit, or take the rule out of the tree; what is checked is git's own verification"))
		}
	}
	return fs
}

// approverNotAuthor compares the by of every decision in this phase's gate.yaml against the
// authors of the range. Section 9 says both things worth knowing about it: it is evaluated on
// the run after the decision was written, which that decision's own commit triggers, and it
// compares two self-asserted strings, so it enforces the discipline of a team that means it.
//
// No decision is green. A rule about who may release a finding is not a rule requiring that one
// be released.
func approverNotAuthor(e *eval, r rules.Rule) []model.Finding {
	cs, fs := e.inRange(r)
	if fs != nil {
		return fs
	}
	authors := map[string]git.Commit{}
	for _, c := range cs {
		if c.Email != "" {
			authors[strings.ToLower(c.Email)] = c
		}
		if c.Author != "" {
			authors[strings.ToLower(c.Author)] = c
		}
	}
	rel := e.phaseRel(e.Phase) + "/gate.yaml"
	var g model.Gate
	if err := fm.ReadYAML(e.abs(rel), &g); err != nil {
		return nil // no verdict yet, so no decision to judge
	}
	for _, ch := range g.Checks {
		for _, f := range ch.Findings {
			if f.Decision == nil || f.Decision.By == "" {
				continue
			}
			if c, ok := authors[strings.ToLower(f.Decision.By)]; ok {
				fs = append(fs, finding(rel,
					"rule "+r.ID+": "+f.ID+" was "+f.Decision.Type+" by "+quoted(f.Decision.By)+
						", who authored commit "+shortHash(c),
					"let somebody who did not write the change decide it; that is what the rule is for"))
			}
		}
	}
	return fs
}

// policy is section 7's row: conformance against the effective rule set, every review checklist
// entry answered, from P0.
//
// The count runs from the rules to the entries and never the other way round. Section 12 says a
// lens entry carries no rule id and cannot add to or subtract from the set G-Policy counts,
// which only holds if the set is the rules; and section 9 pairs "every entry carries a result,
// and nothing more" with "passing over one is now a recorded deviation rather than an omission
// nobody sees", which only holds if a missing entry is a finding. A gate that walked the entries
// would pass an empty checklist.
//
// The tree is resolved through internal/rules and not again here, because section 7 says this
// gate and G-Rules resolve the same tree and should do it once: two resolutions that could
// disagree would leave rules_hash recording a set that nothing judged.
func policy(c Ctx) model.Check {
	read, _ := rules.Load(c.Root)
	effective, _ := rules.Effective(read)
	// A tree that does not resolve is G-Rules's finding, and repeating it here would report one
	// broken file twice in one verdict.
	//
	// A phase written under a different rule set is not judged against this one. The artifact
	// records which set was in force, which is what rules_hash is for, and a rule adopted today
	// cannot make a judgement taken last month wrong: section 10 routes a rule change through a
	// merge so that it takes effect after review, which is forward. Where the recorded hash is
	// absent or differs, this gate has nothing to say about the phase (A74).
	if !judgedUnder(c, rules.Hash(effective)) {
		return result(nil)
	}
	var fs []model.Finding
	fs = append(fs, checkedRules(&eval{Ctx: c}, effective)...)
	if c.Phase == model.Phases[len(model.Phases)-1] {
		fs = append(fs, reviewChecklist(c, effective)...)
	}
	sort.Slice(fs, func(i, j int) bool { return fs[i].Cause < fs[j].Cause })
	return result(fs)
}

// judgedUnder reports whether the phase's artifact says it was written under this rule set.
//
// An artifact that records no rules_hash predates the writer and cannot say; one that records the
// placeholder was written when no writer existed; one that records another hash was written under
// another set. In all three cases the set in hand is not the set that applied, and judging the
// phase against it would rewrite a verdict rather than recompute one.
func judgedUnder(c Ctx, want string) bool {
	var o model.Output
	if _, err := fm.ReadFront(c.abs(c.phaseRel(c.Phase)+"/output.md"), &o); err != nil {
		return false
	}
	return o.RulesHash != "" && o.RulesHash == want
}

// checkedRules reports a checked rule that applies to this phase and whose predicate type
// nothing implements. The alternative was to pass it, which is the silently green verdict
// section 16 catalogues: a rule in force that nothing evaluated. Where the type is implemented
// the predicate runs, which is what the next piece turns on.
func checkedRules(e *eval, effective []rules.Rule) []model.Finding {
	var fs []model.Finding
	for _, r := range effective {
		if r.Kind != rules.Checked || !appliesTo(r, e.Phase) {
			continue
		}
		t := ""
		if r.Check != nil {
			t = r.Check.Type
		}
		fn, ok := predicates[t]
		if !ok {
			fs = append(fs, finding(r.Path,
				"rule "+r.ID+" is checked and no implementation exists for predicate type "+quoted(t),
				"write the rule as kind review until the type ships, or take it out of the tree"))
			continue
		}
		fs = append(fs, fn(e, r)...)
	}
	return fs
}

// reviewChecklist counts the review rules of the effective set against the answers in the P5
// artifact. Whether an answer is a good one is not a question a deterministic gate may ask, and
// section 9 says so; what it asks is whether the rule was answered at all.
//
// Every review rule is answered here whatever its applies_to, because the checklist exists once
// and section 9 renders it from the effective rule set rather than from a phase's slice of it. A
// review rule about design is answered when the change is reviewed (A69).
func reviewChecklist(c Ctx, effective []rules.Rule) []model.Finding {
	rel := c.phaseRel(c.Phase) + "/output.md"
	var o model.Output
	if _, err := fm.ReadFront(c.abs(rel), &o); err != nil {
		return nil // the artifact's absence is G-Schema's finding, not this gate's
	}

	answered := map[string]bool{}
	known := map[string]bool{}
	for _, r := range effective {
		if r.Kind == rules.Review {
			known[r.ID] = true
		}
	}
	var fs []model.Finding
	for _, e := range o.ReviewChecklist {
		switch {
		case e.Result == "":
			fs = append(fs, finding(rel, "checklist entry "+entryName(e)+" carries no result",
				"answer it with met, deviation or not-applicable"))
		case !model.OneOf(e.Result, model.ChecklistResults):
			fs = append(fs, finding(rel, "checklist entry "+entryName(e)+" has result "+quoted(e.Result),
				"section 9 fixes the three: met, deviation, not-applicable"))
		case model.OneOf(e.Result, model.ChecklistNeedsNote) && e.Note == "":
			fs = append(fs, finding(rel, "checklist entry "+entryName(e)+" is "+e.Result+" and carries no note",
				"write why the rule was passed over; met is the only result that needs none"))
		}
		if e.Rule == "" {
			continue // a lens entry answers no rule, by section 12, whatever its source says
		}
		if !known[e.Rule] {
			fs = append(fs, finding(rel, "checklist entry answers "+quoted(e.Rule)+", which is no review rule of the effective set",
				"correct the id, or take the entry out; a rule that does not apply is not answered here"))
			continue
		}
		answered[e.Rule] = true
	}
	for _, r := range effective {
		if r.Kind == rules.Review && !answered[r.ID] {
			fs = append(fs, finding(rel, "review rule "+r.ID+" has no checklist entry",
				"answer it in review_checklist with met, deviation or not-applicable"))
		}
	}
	return fs
}

// appliesTo reads a rule's applies_to against a phase id.
func appliesTo(r rules.Rule, phase string) bool {
	for _, p := range r.AppliesTo {
		if p == phase {
			return true
		}
	}
	return false
}

// entryName is how a finding names an entry that may have no rule id, which is every lens entry.
func entryName(e model.ChecklistEntry) string {
	if e.Rule == "" {
		if e.Source != "" {
			return "from " + e.Source
		}
		return "with no rule"
	}
	return quoted(e.Rule)
}

func quoted(s string) string { return "\"" + s + "\"" }
