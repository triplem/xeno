// SPDX-License-Identifier: Apache-2.0

// Package model holds the directory layout and the file formats of the process
// definition. It has no behaviour beyond naming things.
package model

import (
	"fmt"
	"path/filepath"
)

const (
	RunnerVersion = "0.1.0-dev"
	PluginVersion = "0.1.0-dev"
)

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

// Common is carried by every process file.
type Common struct {
	Intent        string `yaml:"intent"`
	Phase         string `yaml:"phase,omitempty"`
	Created       string `yaml:"created"`
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

type Assumption struct {
	ID          string `yaml:"id"`
	Phase       string `yaml:"phase,omitempty"`
	Text        string `yaml:"text"`
	Resolves    string `yaml:"resolves,omitempty"`
	ConfirmedBy string `yaml:"confirmed_by,omitempty"`
}

type Assumptions struct {
	Common      `yaml:",inline"`
	Assumptions []Assumption `yaml:"assumptions"`
}

// ContextLock records the input state. PredecessorHash is what G-Freshness and the
// computed staleness in `intent status` compare against.
type ContextLock struct {
	Common          `yaml:",inline"`
	PredecessorHash string `yaml:"predecessor_artifacts_hash,omitempty"`
	EvidenceSource  string `yaml:"evidence_source"`
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

// Project is the part of project.yaml the core reads.
type Project struct {
	Evidence struct {
		Source string `yaml:"source"`
	} `yaml:"evidence"`
}
