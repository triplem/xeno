// SPDX-License-Identifier: Apache-2.0

// Package model holds the directory layout and the file formats of the process
// definition. It has no behaviour beyond naming things.
package model

import (
	"fmt"
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
// PluginVersion takes no stamp. It describes the vendored plugin and not the binary
// that wrote the artifact, and there is no plugin yet to have a commit.
var (
	RunnerVersion = devVersion
	PluginVersion = devVersion
)

const devVersion = "0.1.0-dev"

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

const ContextProfile = "context-profile.yaml" // P0 only

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
	PluginVersion string `yaml:"plugin_version"`
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
type EvidenceItem struct {
	Kind   string `yaml:"kind"`
	Job    string `yaml:"job,omitempty"`
	SHA256 string `yaml:"sha256,omitempty"`
	Path   string `yaml:"path,omitempty"`
	URI    string `yaml:"uri,omitempty"`
	Result string `yaml:"result,omitempty"`
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
	// Files is the information base the phase was given, resolved from the context
	// profile when the phase started and never refreshed: the lock describes the input
	// state, and one rewritten at the end would describe nothing. G-Freshness compares
	// these hashes against the tree for every preceding phase.
	Files []ContextFile `yaml:"files,omitempty"`
}

// ContextFile is one entry of the information base, as section 5 writes it.
type ContextFile struct {
	Path   string `yaml:"path"`
	SHA256 string `yaml:"sha256"`
}

// Profile is the context profile section 12 defines: a budget, produced by P0 and read by
// every phase. Links are declared and never inferred, and the runner records them without
// resolving them: what they are for is the agent's reading, not the gate's.
type Profile struct {
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
	PluginVersion string `yaml:"plugin_version"`
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
}
