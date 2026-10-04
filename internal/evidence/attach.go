// SPDX-License-Identifier: Apache-2.0

// Package evidence attaches pipeline results to declarations that were left pending.
// CI never writes into the repository; the runner pulls. Everything written here lies
// under evidence/, which is outside artifacts_hash.
package evidence

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
)

// ManifestEntry describes one result a pipeline published. In a real setup the fetch
// adapter reads the CI artifact store; offline, a directory with manifest.yaml stands
// in for it, which keeps the attach logic testable without a host.
type ManifestEntry struct {
	Kind     string `yaml:"kind"`
	Job      string `yaml:"job"`
	Result   string `yaml:"result"`
	File     string `yaml:"file,omitempty"`   // small textual result, copied into evidence/
	URI      string `yaml:"uri,omitempty"`    // large or binary result, referenced
	SHA256   string `yaml:"sha256,omitempty"` // required with uri
	Pipeline string `yaml:"pipeline,omitempty"`
	Commit   string `yaml:"commit,omitempty"`
}

// Result is what one attach did. It carries what was declined as well as what was
// counted, because Attach writes no output and is called from three places: `phase
// start`, `gate run` and `evidence attach`. A refusal it swallowed would leave the item
// pending with no explanation anywhere, which is a worse failure than the one this
// refusal prevents, so the reason is returned and each caller says it in its own voice.
type Result struct {
	Attached int
	Pending  int
	// Declined names the entries the pipeline published that were not recorded, one sentence
	// each, and each sentence says what to republish. They are counted in Pending, because
	// pending is what they were and declining to record one changes nothing about the phase.
	//
	// The sentence carries its own remedy rather than the callers carrying one, because there
	// are three reasons now and they are fixed in different places: a hash and a uri are
	// republished by the step that uploads, and a result outside section 4's set is a word the
	// job writes into its own manifest (#221).
	Declined []string
}

// Attach fills pending declarations of one phase from source.
//
// An entry it cannot bind, or whose result is a word section 4 does not define, is declined
// rather than recorded. A8 says an item with a `uri` is bound by its declared hash alone,
// since resolving a uri needs a network the gate path never has, so an entry arriving
// without one is bound by nothing: section 4 pays for attaching outside the
// `artifacts_hash` with the hash `gate.yaml` records, and there is nothing to record.
// Recorded anyway, it would reach G-Evidence as a record that cannot be judged and no
// longer says whether the pipeline published it wrong or somebody edited the file, which is
// why this is refused here and not there.
//
// The result is declined for the same reason one step earlier. The set is closed and
// G-Evidence judges a recorded value against it since #221, so an entry carrying `success`
// would be written here and reported there, with nothing in between saying which of the two
// to believe. Fifteen records in this repository got in exactly that way.
func Attach(root, key, phase, source string) (Result, error) {
	var res Result
	dir := filepath.Join(root, model.PhaseDir(key, phase))
	var o model.Output
	if _, err := fm.ReadFront(filepath.Join(dir, "output.md"), &o); err != nil {
		return res, err
	}
	attPath := filepath.Join(dir, "evidence", "attached.yaml")
	var att []model.Attached
	_ = fm.ReadYAML(attPath, &att)
	has := func(kind, job string) bool {
		for _, a := range att {
			if a.Kind == kind && a.Job == job && a.State == "attached" {
				return true
			}
		}
		return false
	}

	var manifest []ManifestEntry
	if source != "" {
		if err := fm.ReadYAML(filepath.Join(source, "manifest.yaml"), &manifest); err != nil && !os.IsNotExist(err) {
			return res, err
		}
	}
	find := func(kind, job string) *ManifestEntry {
		for i := range manifest {
			if manifest[i].Kind == kind && manifest[i].Job == job {
				return &manifest[i]
			}
		}
		return nil
	}

	for _, d := range o.Evidence {
		if !d.Pending() || has(d.Kind, d.Job) {
			continue
		}
		m := find(d.Kind, d.Job)
		if m == nil {
			res.Pending++
			continue
		}
		// Before the binding reasons below, because a result is wrong whichever way the
		// entry is bound and because it is the one a publisher fixes in the step that
		// wrote it.
		if why := unrecordable(*m); why != "" {
			res.Declined = append(res.Declined, d.Kind+"/"+d.Job+" "+why)
			res.Pending++
			continue
		}
		a := model.Attached{Kind: d.Kind, Job: d.Job, State: "attached", Result: m.Result,
			Pipeline: m.Pipeline, Commit: m.Commit}
		if m.File != "" {
			b, err := os.ReadFile(filepath.Join(source, m.File))
			if err != nil {
				return res, err
			}
			name := filepath.Base(m.File)
			if err := os.MkdirAll(filepath.Join(dir, "evidence"), 0o755); err != nil {
				return res, err
			}
			if err := os.WriteFile(filepath.Join(dir, "evidence", name), b, 0o644); err != nil {
				return res, err
			}
			a.Path = "evidence/" + name
			a.SHA256 = hashing.Hex(hashing.Normalise(b))
		} else if why := unbindable(*m); why != "" {
			res.Declined = append(res.Declined, d.Kind+"/"+d.Job+" "+why)
			res.Pending++
			continue
		} else {
			a.URI, a.SHA256 = m.URI, m.SHA256
		}
		att = append(att, a)
		res.Attached++
	}
	if res.Attached > 0 {
		if err := fm.WriteYAML(attPath, att); err != nil {
			return res, err
		}
	}
	return res, nil
}

// unbindable says why an entry with no file cannot be bound, or "" where it can. A file
// backed entry never reaches this: its hash is computed from the bytes that were copied.
//
// The two cases are one defect with different amounts to look at. An entry with a uri and
// no hash names where something lives and nothing about what it is; one with neither names
// a result and nothing it belongs to.
func unbindable(m ManifestEntry) string {
	switch {
	case m.URI == "":
		return "was published with no file, no uri and no hash, so it names a result and " +
			"nothing it belongs to; republish it with the file, or with a uri and its sha256"
	case m.SHA256 == "":
		return "was published with a uri and no sha256, and a uri is bound by its hash " +
			"alone, because resolving one needs a network the gate path never has; " +
			"republish it with the sha256"
	}
	return ""
}

// unrecordable says why an entry cannot be written down at all, as against the two above,
// which are about what would bind it. One reason so far: a result section 4 does not define.
//
// It judges a value and not an absence, as the gate does. Section 4 has the field absent
// where a producer reports nothing — a bill of materials, a database metadata file — and this
// repository's vulnerability scan publishes one of those beside its report.
func unrecordable(m ManifestEntry) string {
	if m.Result == "" || model.OneOf(m.Result, model.EvidenceResults) {
		return ""
	}
	return "was published with result " + m.Result + ", which section 4 does not define; the " +
		"set is " + strings.Join(model.EvidenceResults, " and ") + ", and the value is what the " +
		"run reported against its own threshold. Republish it with one of the two"
}
