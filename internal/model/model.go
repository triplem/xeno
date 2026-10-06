// SPDX-License-Identifier: Apache-2.0

// Package model holds the directory layout and the file formats of the process
// definition. It has no behaviour beyond naming things.
package model

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

// The version the runner reports and writes into every artifact. It is var rather than
// const so that the release build sets it through -ldflags from the tag the pipeline
// derived, which is what keeps the number out of the source.
//
// A build that was not made by the release carries devVersion, and init below appends
// the commit it came from. Without that the field names every unreleased build ever
// made and therefore identifies none of them, which is a poor showing for a field whose
// only job is to say which binary wrote an artifact.
//
// devVersion claims no number. It read `0.1.0-dev` for twenty-nine minor releases, because
// a literal in the source cannot learn the newest tag — `debug.ReadBuildInfo` reports
// `(devel)` for the main module and the tag is known only to git, at build time. So every
// artifact in this trail said `0.1.0-dev` while the project shipped v0.29.2, which is not a
// stale number but a wrong one. `dev` says what the build is and leaves the number to the
// release, which sets it through ldflags and is the only thing that knows it (#177).
//
// There is no PluginVersion here any more. It was a constant beside this one, with a comment
// saying there was no plugin yet to have a version, and a release set it to the runner's own
// through ldflags — so a 0.28.0 runner working in a project whose vendored plugin was 0.26.0
// recorded 0.28.0. A number that cannot disagree with the runner can never be proved wrong,
// which is the one thing section 5 wants it for. It is read from the vendored plugin's own
// manifest now, by internal/plugin (#177).
var RunnerVersion = devVersion

const devVersion = "dev"

func init() {
	if RunnerVersion != devVersion {
		return // a release build already said what it is
	}
	revision, modified := vcsStamp()
	RunnerVersion = stamp(devVersion, revision, modified)
}

// stamp appends the commit to a development version as semver build metadata: "+", then
// dot separated alphanumerics. An empty revision leaves the version alone, which is what
// happens in a test binary, where the toolchain stamps no vcs settings at all.
//
// "dirty" says the tree was not the commit. It does not say which tree it was, so two
// builds from one commit with different uncommitted changes report the same string.
func stamp(base, revision string, modified bool) string {
	if revision == "" {
		return base
	}
	if len(revision) > 7 {
		revision = revision[:7]
	}
	if modified {
		return base + "+" + revision + ".dirty"
	}
	return base + "+" + revision
}

func vcsStamp() (revision string, modified bool) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", false
	}
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			revision = s.Value
		case "vcs.modified":
			modified = s.Value == "true"
		}
	}
	return revision, modified
}

// SchemaVersion is the shape of the artifacts this runner writes, and it moves for a
// different reason than the two above: the minor when a field is added that an older
// reader can ignore, the major when one changes meaning or goes away. It is const,
// because a build cannot be told to write a shape other than the one it implements.
//
// Artifacts written before the field existed do not carry it, and that absence is the
// schema that predates versioning rather than a field somebody forgot.
const SchemaVersion = "1.0"

// Phases in order. The order is the process; nothing else defines it.
var Phases = []string{
	"00-intake", "01-requirements", "02-design",
	"03-implementation", "04-verification", "05-review",
}

// PhaseIndex returns the position of a phase id, or -1.
func PhaseIndex(id string) int {
	for i, p := range Phases {
		if p == id {
			return i
		}
	}
	return -1
}

// ResolvePhase accepts "04" or "04-verification".
func ResolvePhase(s string) (string, error) {
	if len(s) == 1 {
		s = "0" + s
	}
	for _, p := range Phases {
		if p == s || p[:2] == s {
			return p, nil
		}
	}
	return "", fmt.Errorf("unknown phase %q", s)
}

func IntentDir(key string) string { return filepath.ToSlash(filepath.Join(".xeno/intents", key)) }
func PhaseDir(key, phase string) string {
	return filepath.ToSlash(filepath.Join(IntentDir(key), "phases", phase))
}

// Files a phase directory may hold, per section 4. Anything else is a G-Schema finding.
var KnownPhaseFiles = map[string]bool{
	"output.md": true, "context.lock.yaml": true, "digest.md": true,
	"gate.yaml": true, "learning.yaml": true, "cost.yaml": true,
}

// Files an intent directory may hold, per section 4, which lists these four and the phases
// directory beside them. Anything else is a G-Complete finding when the intent is closed.
//
// Two maps rather than one, because the two levels hold different files and a single map would
// accept output.md in an intent directory and assumptions.yaml in a phase, so neither check would
// distinguish what section 4 distinguishes (#109).
var KnownIntentFiles = map[string]bool{
	"intent.yaml": true, "assumptions.yaml": true,
	"learning.yaml": true, "gate.yaml": true,
}

// PhasesDir is the one directory an intent directory may hold, as evidence is the one a phase may.
const PhasesDir = "phases"

const ContextScope = "context-scope.yaml" // P0 only

// TemplateID is the phase without its ordering prefix: 02-design is rendered from the
// template design. The prefix orders the phases and says nothing a template needs.
func TemplateID(phase string) string {
	if i := strings.IndexByte(phase, '-'); i >= 0 {
		return phase[i+1:]
	}
	return phase
}

// Common is carried by every process file.
type Common struct {
	Intent        string `yaml:"intent"`
	Phase         string `yaml:"phase,omitempty"`
	Created       string `yaml:"created"`
	SchemaVersion string `yaml:"schema_version,omitempty"`
	RunnerVersion string `yaml:"runner_version"`
	// PluginVersion is omitempty because absent is not a default: a repository with no
	// vendored plugin rendered from none, and section 5 requiring the field in every process
	// file is what makes that a G-Schema finding rather than something to paper over (A35).
	PluginVersion string `yaml:"plugin_version,omitempty"`
}

// Question is an entry of the open-questions section, carried structurally.
type Question struct {
	Key     string   `yaml:"key"`
	Text    string   `yaml:"text"`
	Options []Option `yaml:"options,omitempty"`
	// NoOptions states that the agent found none, rather than inventing two.
	NoOptions bool `yaml:"no_options,omitempty"`
}

type Option struct {
	Text        string `yaml:"text"`
	Consequence string `yaml:"consequence,omitempty"`
	Recommended bool   `yaml:"recommended,omitempty"`
	Free        bool   `yaml:"free,omitempty"` // the free entry, an option like the others
	// Section 8: "The reason belongs to the option the recommendation names and is carried
	// there rather than on the question." Empty on every option but one, which the clause
	// names as the cost: a recommendation that moves to another option takes its reason with
	// it or QuestionAsked refuses, where a reason beside the question would go on describing
	// the option it used to be about.
	Reason string `yaml:"reason,omitempty"`
}

// Decision is an entry of the decisions section.
type Decision struct {
	ID         string `yaml:"id"`
	Resolves   string `yaml:"resolves,omitempty"`
	Chosen     string `yaml:"chosen"`
	Rationale  string `yaml:"rationale"`
	DecidedBy  string `yaml:"decided_by"`
	ProposedBy string `yaml:"proposed_by,omitempty"`
	Withdrawn  bool   `yaml:"withdrawn,omitempty"`
}

// EvidenceItem is a declaration in the frontmatter of output.md. An item that exists
// at finish time carries sha256 and path or uri; a pending one carries kind and job only.
//
// The fields stand in the order section 4 prints them, because a reader compares an
// artifact against the document and not against this struct, and the order a written item
// carries is this one. `id` is the one field of the section with no field here: an item is
// identified by its kind and its job, which is what Collect and evidence.Attach key an
// attachment to a declaration by, and a second identity both of them ignore would be a
// name for the pair that no reader uses (D-2).
type EvidenceItem struct {
	Kind string `yaml:"kind"`
	// Result is what the run reported against its own threshold, never a statement by Xeno.
	Result string `yaml:"result,omitempty"`
	// ProducedBy is the command as run and Format is the shape of the report. Both are
	// section 4 fields, both are provenance, and nothing reads either: the section says so
	// of format in as many words, and A-002 records what producedBy means on an item a
	// pipeline has yet to produce.
	ProducedBy string `yaml:"produced_by,omitempty"`
	Format     string `yaml:"format,omitempty"`
	SHA256     string `yaml:"sha256,omitempty"`
	Path       string `yaml:"path,omitempty"`
	URI        string `yaml:"uri,omitempty"`
	Job        string `yaml:"job,omitempty"`
}

// Pending reports whether the item still waits for a pipeline.
func (e EvidenceItem) Pending() bool { return e.SHA256 == "" }

// Output is the frontmatter of output.md as far as the gates read it.
type Output struct {
	Common        `yaml:",inline"`
	Language      string         `yaml:"language"`
	SecretsHash   string         `yaml:"secrets_hash"`
	ContextHash   string         `yaml:"context_hash"`
	Model         string         `yaml:"model"`
	Tool          string         `yaml:"tool"`
	ToolVersion   string         `yaml:"tool_version"`
	Template      string         `yaml:"template"`
	StringsHash   string         `yaml:"strings_hash"`
	RulesHash     string         `yaml:"rules_hash"`
	OpenQuestions []Question     `yaml:"open_questions,omitempty"`
	Decisions     []Decision     `yaml:"decisions,omitempty"`
	Evidence      []EvidenceItem `yaml:"evidence,omitempty"`
	// The P5 checklist of section 9, one entry per review rule of the effective set. It is a
	// frontmatter list beside the three above, which is A3's shape and A68's spelling: the
	// section id is hyphenated and the key is not, as open-questions and open_questions
	// already are.
	ReviewChecklist []ChecklistEntry `yaml:"review_checklist,omitempty"`
}

// LocalDataEnv is the variable section 7 lists in the normalised environment, and
// LocalDefault the location it names: "XENO_PLUGIN_DATA  local data location, always
// .xeno/local/".
//
// The entry point exports it and, until #205, nothing read it: six places wrote the path as a
// literal instead, one of them twice. So the export promised something the runner did not keep.
const (
	LocalDataEnv = "XENO_PLUGIN_DATA"
	LocalDefault = ".xeno/local"
)

// LocalDir is where a repository's local data goes: the run marker, phase.env, the cost ledger
// and the enforcement report.
//
// Reading the variable is safe in a way that reading XENO_PLUGIN_ROOT was not, which is the
// distinction A89 and this one's row turn on: nothing under here is covered by artifacts_hash,
// so no verdict can be made to depend on the environment. The plugin root resolves rules_hash
// and a rendered artifact, which is why section 7 lost that one rather than gaining a reader.
//
// An absolute value is returned as it stands, because the entry point exports
// `${XENO_PLUGIN_DATA:=$root/.xeno/local}` and joining that to the root would nest one path
// inside another. A relative value is relative to the repository root, as every other path the
// runner handles is, and not to the process's working directory: a value meaning different
// things depending on where the command was invoked is a worse promise than none.
//
// Empty reads as unset. os.Getenv cannot tell them apart, an exported-but-empty variable is
// what `export XENO_PLUGIN_DATA=` produces, and resolving it to the root would put the ledger
// and the run marker at the top of the tree.
//
// Nothing is created here. The writers already make what they need, and a resolver that made
// directories would do it on every read, including the reads that only report a path.
func LocalDir(root string) string {
	if v := strings.TrimSpace(os.Getenv(LocalDataEnv)); v != "" {
		if filepath.IsAbs(v) {
			return v
		}
		return filepath.Join(root, v)
	}
	return filepath.Join(root, LocalDefault)
}

// LocalPath is a file or directory inside LocalDir.
func LocalPath(root string, parts ...string) string {
	return filepath.Join(append([]string{LocalDir(root)}, parts...)...)
}

// ChecklistEntry is one answer in the P5 checklist. Rule is the anchor section 9 gives each
// entry, and its absence is what keeps a lens entry out of the counted set: section 12 says a
// lens entry carries source: lens and no rule id, and the gate keys on the missing rule rather
// than on the source, so a lens that omitted its own label still cannot answer a rule.
type ChecklistEntry struct {
	Rule   string `yaml:"rule,omitempty"`
	Result string `yaml:"result"`
	Note   string `yaml:"note,omitempty"`
	Source string `yaml:"source,omitempty"` // lens, where a lens wrote it
}

// Cost is the record of section 11, one per phase, outside artifacts_hash so that a figure
// arriving after a verdict cannot invalidate it. cost_usd is deliberately absent: section 11
// makes it optional and puts the authoritative money view in v2, and a price is not a fact this
// repository holds.
type Cost struct {
	Common       `yaml:",inline"`
	Evidence     string   `yaml:"evidence"`
	TokensIn     int      `yaml:"tokens_in"`
	TokensOut    int      `yaml:"tokens_out"`
	TokensCached int      `yaml:"tokens_cached"`
	Sessions     []string `yaml:"sessions"`
}

// Attached is one entry of evidence/attached.yaml, outside artifacts_hash.
type Attached struct {
	Kind     string `yaml:"kind"`
	Job      string `yaml:"job"`
	State    string `yaml:"state"`
	Result   string `yaml:"result"`
	SHA256   string `yaml:"sha256"`
	Path     string `yaml:"path,omitempty"`
	URI      string `yaml:"uri,omitempty"`
	Pipeline string `yaml:"pipeline,omitempty"`
	Commit   string `yaml:"commit,omitempty"`
}

// Assumption is the record section 8 defines. Status carries the state, and an empty
// status reads as open: absent means nothing has decided it yet, which is the same
// answer a register written by hand before the field existed gives.
//
// Resolves is not in section 8's schema. The same section resolves an open question as
// an assumption somebody confirms, which needs the question's key on the record to be
// checkable, and G-Questions reads it for exactly that.
type Assumption struct {
	ID          string `yaml:"id"`
	Phase       string `yaml:"phase,omitempty"`
	Assumption  string `yaml:"assumption"`
	Origin      string `yaml:"origin,omitempty"`
	Confidence  string `yaml:"confidence,omitempty"`
	Status      string `yaml:"status"`
	ConfirmedBy string `yaml:"confirmed_by,omitempty"`
	RejectedBy  string `yaml:"rejected_by,omitempty"`
	Resolves    string `yaml:"resolves,omitempty"`
}

// OneOf is the closed set test every enumeration in the documents needs. It lived twice,
// once in the gates and once in the runner, which is how a helper of four lines becomes
// two helpers of four lines.
func OneOf(value string, set []string) bool {
	for _, s := range set {
		if s == value {
			return true
		}
	}
	return false
}

// The closed sets the documents fix. They are here rather than in the packages that read
// them because they are the specification's enumerations, not one gate's detail, and a
// reader looking for what a field may carry looks for the type it belongs to.
var (
	// Section 10: what a learning record's category may be, and the four keys of an entry.
	LearningCategories = []string{"template", "prompt", "context-rule", "project-convention"}
	LearningKeys       = []string{"category", "observation", "proposal", "target"}
	// Appendix B: the fields that carry a hash, and the two that have no writer yet.
	HashFields      = []string{"context_hash", "secrets_hash", "strings_hash", "rules_hash"}
	WriterlessHash  = []string{"secrets_hash", "rules_hash"}
	HashPlaceholder = "by-hand"
	// ProjectFile is where a project's configuration lives, which is the file a finding about a
	// declaration has to name.
	ProjectFile = ".xeno/config/project.yaml"
	// Section 4: what an evidence declaration's kind may be, and what its result may say.
	// Section 9: what an answer to a review rule may say, and which of them owe a note. A met
	// needs none; the note exists to record why a rule was passed over.
	ChecklistResults   = []string{"met", "deviation", "not-applicable"}
	ChecklistNeedsNote = []string{"deviation", "not-applicable"}
	EvidenceKinds      = []string{"test-report", "coverage", "build-log", "scan", "sbom", "other"}
	EvidenceResults    = []string{"pass", "fail"}
	// Section 4's five shapes of a report. The set is closed there and judged here for the
	// reason BuildKind records: a value nobody compares against is a spelling that survives
	// for as long as nothing reads it, which `kind: build` did until #198.
	EvidenceFormats = []string{"junit", "trx", "tap", "go-test-json", "other"}
	// The two kinds that must carry a result, because G-Test and G-Build read it.
	// Elsewhere the field follows the producer: a run with a threshold reports against it
	// and one without reports nothing, so an absent result there is a producer that had
	// nothing to say rather than a writer who forgot.
	ResultRequiredKinds = []string{"test-report", "build-log"}
)

// The closed sets of section 8. A value outside one of them is refused where a record is
// written, so that the register cannot fill up with spellings the reader has to guess at.
var (
	AssumptionOrigins     = []string{"template-default", "repo-convention", "rules", "user-input"}
	AssumptionConfidences = []string{"high", "medium", "low"}
	AssumptionStatuses    = []string{"open", "confirmed", "rejected"}
)

// Open says whether the gate of the assumption's phase goes red for it. Rejected is
// decided and not confirmed: the assumption was examined and dropped.
func (a Assumption) Open() bool { return a.Status != "confirmed" && a.Status != "rejected" }

// DecidedBy is the person behind whichever decided state the record is in. Section 8
// gives each state its own field and says exactly one of them is present, so this reads
// the one that belongs to the status rather than whichever is filled in.
func (a Assumption) DecidedBy() string {
	if a.Status == "rejected" {
		return a.RejectedBy
	}
	return a.ConfirmedBy
}

type Assumptions struct {
	Common      `yaml:",inline"`
	Updated     string       `yaml:"updated,omitempty"`
	Assumptions []Assumption `yaml:"assumptions"`
}

// ContextLock records the input state. PredecessorHash is what G-Freshness and the
// computed staleness in `intent status` compare against.
type ContextLock struct {
	Common          `yaml:",inline"`
	PredecessorHash string `yaml:"predecessor_artifacts_hash,omitempty"`
	EvidenceSource  string `yaml:"evidence_source"`
	// TemplateSource is plugin or project, and it is recorded because otherwise two
	// projects on the same template version are indistinguishable although one of them
	// overrode it. Empty where no template could be resolved, which is what a
	// repository without a vendored plugin looks like until xeno init puts one there.
	TemplateSource string `yaml:"template_source,omitempty"`
	// RepoCommit is the commit the phase was started against, as section 5 writes it. Absent
	// where the repository has none or is not one at all, because a lock saying which commit a
	// phase ran against is a claim and an invented value would be a false one.
	RepoCommit string `yaml:"repo_commit,omitempty"`
	// Files is the information base the phase was given, resolved from the context
	// scope when the phase started and never refreshed: the lock describes the input
	// state, and one rewritten at the end would describe nothing. G-Freshness compares
	// these hashes against the tree for every preceding phase.
	//
	// The order is the scope's include order, with paths sorted inside each pattern. Section 5
	// asks for an order of volatility and says the lock records the assembly order rather than
	// only the set, so the project writes its patterns from stable to volatile and this follows;
	// the tie-break is what keeps two runs over one tree byte-identical, which the hash needs.
	Files []ContextFile `yaml:"files,omitempty"`
	// RulesApplied is the effective rule set the phase was judged against, one entry per rule.
	// It answers what rules_hash cannot: which rules, and which revision of each. Absent where
	// no rule is in force, because an empty list says a set was resolved and came out empty,
	// where absence says there was nothing to resolve (A74's distinction).
	RulesApplied []AppliedRule `yaml:"rules_applied,omitempty"`
	// Plugin is what section 5 enumerates beside the frontmatter's plugin_version, and the
	// sentence that explains the pair: "the frontmatter names what was used,
	// context.lock.yaml proves it with a hash. Where the two disagree, the hash wins and
	// G-Supply fails." That gate is not implemented, so this is written and read back by
	// nothing yet — a field with a writer and no reader, which is the way round that leaves
	// the trail able to answer the question later.
	Plugin *LockPlugin `yaml:"plugin,omitempty"`
}

// LockPlugin is the lock's plugin block: the version the vendored plugin declares and the
// hash over its tree, which internal/plugin defines.
type LockPlugin struct {
	Version string `yaml:"version,omitempty"`
	SHA256  string `yaml:"sha256,omitempty"`
}

// AppliedRule is one entry of rules_applied, as section 5 writes it: the path the rule was read
// from and its own version counter, which claims nothing about compatibility and says how often
// the file has moved.
type AppliedRule struct {
	Path    string `yaml:"path"`
	Version int    `yaml:"version"`
}

// ContextFile is one entry of the information base, as section 5 writes it.
//
// Bytes is the size of the file as the walk saw it, taken with the hash so the two describe one
// read of one file. Section 5: "the size is recorded because the budget is judged against what the
// phase was given and not against what the tree holds now". Absent in every lock written before the
// field existed, which is why the budget asks whether any entry carries a size rather than whether
// the sum is zero.
type ContextFile struct {
	Path   string `yaml:"path"`
	SHA256 string `yaml:"sha256"`
	Bytes  int64  `yaml:"bytes,omitempty"`
}

// Scope is the context scope section 12 defines: a budget, produced by P0 and read by
// every phase. Links are declared and never inferred, and the runner records them without
// resolving them: what they are for is the agent's reading, not the gate's.
type Scope struct {
	Common  `yaml:",inline"`
	Include []string `yaml:"include"`
	Exclude []string `yaml:"exclude,omitempty"`
	Links   []struct {
		Component string `yaml:"component"`
		Docs      string `yaml:"docs"`
	} `yaml:"links,omitempty"`
	Budget struct {
		Files int `yaml:"files,omitempty"`
		Bytes int `yaml:"bytes,omitempty"`
	} `yaml:"budget,omitempty"`
}

type DecisionOnFinding struct {
	Type       string `yaml:"type"`
	By         string `yaml:"by"`
	At         string `yaml:"at"`
	Against    string `yaml:"against"`
	Reason     string `yaml:"reason"`
	Obligation string `yaml:"obligation,omitempty"`
}

type Finding struct {
	ID       string             `yaml:"id"`
	File     string             `yaml:"file"`
	Cause    string             `yaml:"cause"`
	Next     string             `yaml:"next"`
	Decision *DecisionOnFinding `yaml:"decision,omitempty"`
}

type Check struct {
	Gate       string    `yaml:"gate"`
	Result     string    `yaml:"result"` // pass | fail | pending | not-implemented
	Provenance string    `yaml:"provenance"`
	Findings   []Finding `yaml:"findings,omitempty"`
}

type Gate struct {
	Common        `yaml:",inline"`
	Status        string  `yaml:"status"`
	RunAt         string  `yaml:"run_at"`
	ArtifactsHash string  `yaml:"artifacts_hash"`
	Checks        []Check `yaml:"checks"`
}

// Intent is intent.yaml. The field order is the one the file already has, so that
// rewriting it at close does not reorder a file that sits inside the intent level hash.
type Intent struct {
	Intent        string `yaml:"intent"`
	Key           string `yaml:"key"`
	Status        string `yaml:"status"`
	Reason        string `yaml:"reason,omitempty"`
	Created       string `yaml:"created"`
	SchemaVersion string `yaml:"schema_version,omitempty"`
	RunnerVersion string `yaml:"runner_version"`
	PluginVersion string `yaml:"plugin_version,omitempty"`
}

// Learning is learning.yaml, at phase level and at intent level both. Section 10 fixes the
// shape: the common header, then either entries or the statement that there were none.
//
// NoFinding and Learnings are both omitempty and exactly one of them is written. A record
// carrying neither is what G-Learning calls empty, and one carrying both would claim there
// was nothing to say beside something said.
type Learning struct {
	Common    `yaml:",inline"`
	NoFinding bool            `yaml:"no_finding,omitempty"`
	Learnings []LearningEntry `yaml:"learnings,omitempty"`
}

// LearningEntry is one entry, with the four keys section 10 defines and no fifth. The
// order is the order LearningKeys lists them in, which is the order the section reads.
type LearningEntry struct {
	Category    string `yaml:"category"`
	Observation string `yaml:"observation"`
	Proposal    string `yaml:"proposal"`
	Target      string `yaml:"target"`
}

// Project is the part of project.yaml the core reads. Appendix A has the whole file;
// what is here is what the runner acts on today.
type Language struct {
	// Artifacts is the language artifact content is written in. The process layer,
	// meaning keys, ids, gate names and section ids, is English regardless.
	Artifacts string `yaml:"artifacts"`
}

// Project is the part of project.yaml the core reads.
type Project struct {
	RunnerVersion string `yaml:"runner_version"`
	Evidence      struct {
		Source string `yaml:"source"`
	} `yaml:"evidence"`
	Language Language `yaml:"language"`
	// Section 12's tracker block, read for the qualified intent id and by the branch rules
	// port. One shape rather than one per reader: the two read the same three fields, and
	// two declarations of them is how a field comes to be spelled differently in each.
	Tracker Tracker `yaml:"tracker"`
	// Section 12's agent block. The runner reads it for the two session fields that have a
	// source in the repository: a plausible value in a field nobody produced is worse than an
	// absent one (A35), and these two are produced here.
	Agent struct {
		Tool  string `yaml:"tool"`
		Model struct {
			Default string `yaml:"default"`
		} `yaml:"model"`
	} `yaml:"agent"`
	// Section 5's index block, context economy. Both fields are optional and an absent block
	// means no index, which is what every repository that has not produced one is in.
	Index struct {
		Path        string `yaml:"path"`
		MaxAgeHours int    `yaml:"max_age_hours"`
	} `yaml:"index"`
	// Section 14's external gates. Every other field in this file changes what the runner
	// reads; this one changes what it runs, and the default is that none is declared and
	// none runs, which is what keeps the chain of trust closed for a project that wants it.
	ExternalGates []ExternalGate `yaml:"external_gates"`
}

// ExternalGate is one declaration of section 14: a command, the hash it has to have, and the
// phases it applies to. The hash is checked before every run, which Appendix A says in those
// words, so a gate that was modified refuses to run rather than running unnoticed.
type ExternalGate struct {
	ID     string   `yaml:"id"`
	Path   string   `yaml:"path"`
	SHA256 string   `yaml:"sha256"`
	Phases []string `yaml:"phases"`
}
