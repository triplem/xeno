// Package evidence attaches pipeline results to declarations that were left pending.
// CI never writes into the repository; the runner pulls. Everything written here lies
// under evidence/, which is outside artifacts_hash.
package evidence

import (
	"os"
	"path/filepath"

	"example.com/xeno/internal/fm"
	"example.com/xeno/internal/hashing"
	"example.com/xeno/internal/model"
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

// Attach fills pending declarations of one phase from source. It returns how many
// items were attached in this call and how many remain pending.
func Attach(root, key, phase, source string) (attached, pending int, err error) {
	dir := filepath.Join(root, model.PhaseDir(key, phase))
	var o model.Output
	if _, err := fm.ReadFront(filepath.Join(dir, "output.md"), &o); err != nil {
		return 0, 0, err
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
			return 0, 0, err
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
			pending++
			continue
		}
		a := model.Attached{Kind: d.Kind, Job: d.Job, State: "attached", Result: m.Result,
			Pipeline: m.Pipeline, Commit: m.Commit}
		if m.File != "" {
			b, err := os.ReadFile(filepath.Join(source, m.File))
			if err != nil {
				return attached, pending, err
			}
			name := filepath.Base(m.File)
			if err := os.MkdirAll(filepath.Join(dir, "evidence"), 0o755); err != nil {
				return attached, pending, err
			}
			if err := os.WriteFile(filepath.Join(dir, "evidence", name), b, 0o644); err != nil {
				return attached, pending, err
			}
			a.Path = "evidence/" + name
			a.SHA256 = hashing.Hex(hashing.Normalise(b))
		} else {
			a.URI, a.SHA256 = m.URI, m.SHA256
		}
		att = append(att, a)
		attached++
	}
	if attached > 0 {
		if err := fm.WriteYAML(attPath, att); err != nil {
			return attached, pending, err
		}
	}
	return attached, pending, nil
}
