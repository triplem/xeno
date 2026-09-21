// Package runner implements the commands on top of gates and evidence. It is the only
// writer of gate.yaml and context.lock.yaml, and it never writes inside the
// artifacts_hash of a phase whose verdict exists.
package runner

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"example.com/xeno/internal/evidence"
	"example.com/xeno/internal/fm"
	"example.com/xeno/internal/gates"
	"example.com/xeno/internal/hashing"
	"example.com/xeno/internal/model"
)

// Refusal is a command declining to proceed for a reason a person can act on.
type Refusal struct{ Reason string }

func (r *Refusal) Error() string { return r.Reason }

func refuse(format string, a ...any) error { return &Refusal{fmt.Sprintf(format, a...)} }

type Runner struct {
	Root         string
	Now          func() time.Time
	EvidenceFrom string // stand-in for the CI artifact fetch, see evidence.ManifestEntry
}

func New(root string) *Runner {
	return &Runner{Root: root, Now: func() time.Time { return time.Now().UTC() }}
}

func (r *Runner) abs(rel string) string { return filepath.Join(r.Root, rel) }
func (r *Runner) stamp() string         { return r.Now().Format(time.RFC3339) }

func (r *Runner) marker(key, phase string) string {
	return r.abs(filepath.Join(".xeno/local/runs", key, phase+".lock"))
}

func (r *Runner) qualified(key string) string {
	var in struct {
		Intent string `yaml:"intent"`
	}
	if err := fm.ReadYAML(r.abs(model.IntentDir(key)+"/intent.yaml"), &in); err == nil && in.Intent != "" {
		return in.Intent
	}
	return key
}

func (r *Runner) common(key, phase string) model.Common {
	return model.Common{Intent: r.qualified(key), Phase: phase, Created: r.stamp(),
		RunnerVersion: model.RunnerVersion, PluginVersion: model.PluginVersion}
}

func (r *Runner) readGate(key, phase string) (*model.Gate, error) {
	var g model.Gate
	if err := fm.ReadYAML(r.abs(model.PhaseDir(key, phase)+"/gate.yaml"), &g); err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *Runner) hash(key, phase string) (string, error) {
	return hashing.DirHash(r.Root, model.PhaseDir(key, phase), hashing.PhaseExcluded)
}

// compute runs the gates of a phase without writing anything.
func (r *Runner) compute(key, phase string) (*model.Gate, error) {
	h, err := r.hash(key, phase)
	if err != nil {
		return nil, err
	}
	prev, _ := r.readGate(key, phase)
	checks := gates.Run(gates.Ctx{Root: r.Root, Key: key, Phase: phase, ArtifactsHash: h, QualifiedID: r.qualified(key)}, prev)
	status, err := gates.Status(checks)
	if err != nil {
		return nil, err
	}
	g := &model.Gate{Common: r.common(key, phase), Status: status, RunAt: r.stamp(), ArtifactsHash: h, Checks: checks}
	if prev != nil && prev.Created != "" {
		g.Created = prev.Created
	}
	return g, nil
}

// evaluate computes and writes gate.yaml. gate.yaml lies outside artifacts_hash, so
// rewriting it never changes what a verdict is about.
func (r *Runner) evaluate(key, phase string) (*model.Gate, error) {
	g, err := r.compute(key, phase)
	if err != nil {
		return nil, err
	}
	return g, fm.WriteYAML(r.abs(model.PhaseDir(key, phase)+"/gate.yaml"), g)
}

// Divergence is one difference between a committed verdict and a recomputed one.
type Divergence struct {
	Key, Phase, What string
}

// VerifyResult is what CI reports. It writes nothing: CI recomputes and compares.
type VerifyResult struct {
	Checked     int
	Divergences []Divergence
	Red         []string
	Provisional []string
}

// Verify recomputes every committed verdict and compares it with what is in the
// repository. It never attaches evidence and never writes, so a pending item stays
// pending here and is reported as provisional rather than as a divergence.
func (r *Runner) Verify(key string) (*VerifyResult, error) {
	keys := []string{key}
	if key == "" {
		entries, err := os.ReadDir(r.abs(".xeno/intents"))
		if os.IsNotExist(err) {
			return &VerifyResult{}, nil // no intent yet is not a failure
		}
		if err != nil {
			return nil, err
		}
		keys = nil
		for _, e := range entries {
			if e.IsDir() {
				keys = append(keys, e.Name())
			}
		}
	}
	res := &VerifyResult{}
	for _, k := range keys {
		for _, p := range model.Phases {
			committed, err := r.readGate(k, p)
			if err != nil {
				continue
			}
			res.Checked++
			label := k + " " + p
			got, err := r.compute(k, p)
			if err != nil {
				return nil, err
			}
			if got.ArtifactsHash != committed.ArtifactsHash {
				res.Divergences = append(res.Divergences, Divergence{k, p, "artifacts changed after the verdict was written"})
				continue
			}
			if got.Status != committed.Status {
				res.Divergences = append(res.Divergences, Divergence{k, p,
					fmt.Sprintf("committed status %s, recomputed %s", committed.Status, got.Status)})
			}
			switch got.Status {
			case "red":
				res.Red = append(res.Red, label)
			case "provisional":
				res.Provisional = append(res.Provisional, label)
			}
		}
	}
	return res, nil
}

// Start begins a phase. The sequence is a property of the tool: a phase whose
// predecessor holds no verdict over its current content, a red one or a provisional
// one is refused, and so is a phase that is already running.
func (r *Runner) Start(key, phase string) error {
	idx := model.PhaseIndex(phase)
	if idx < 0 {
		return fmt.Errorf("unknown phase %q", phase)
	}
	if fm.Exists(r.marker(key, phase)) {
		return refuse("%s is already running; if that run died, remove %s", phase, r.marker(key, phase))
	}
	lock := model.ContextLock{Common: r.common(key, phase), EvidenceSource: r.evidenceSource()}

	if idx > 0 {
		pred := model.Phases[idx-1]
		g, err := r.readGate(key, pred)
		if err != nil {
			return refuse("%s has no verdict yet; run xeno phase finish for it first", pred)
		}
		h, err := r.hash(key, pred)
		if err != nil {
			return err
		}
		if g.ArtifactsHash != h {
			return refuse("%s changed after its verdict; run xeno phase finish for it again", pred)
		}
		// Pull what the pipeline has produced for the predecessor, then carry its
		// verdict forward. Only evidence/ and gate.yaml of that one phase are written.
		if g.Status == "provisional" {
			if _, _, err := evidence.Attach(r.Root, key, pred, r.EvidenceFrom); err != nil {
				return err
			}
			if g, err = r.evaluate(key, pred); err != nil {
				return err
			}
			if g.Status == "provisional" {
				return refuse("%s still waits for evidence from the pipeline; start again once it has run", pred)
			}
		}
		// The next phase starts only on a decided predecessor. Building on a red
		// verdict moves the failure downstream instead of resolving it.
		if g.Status == "red" {
			return refuse("%s is red; decide every failing finding first: fix it, approve it or override it", pred)
		}
		lock.PredecessorHash = h
	}

	dir := r.abs(model.PhaseDir(key, phase))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	if err := fm.WriteYAML(filepath.Join(dir, "context.lock.yaml"), lock); err != nil {
		return err
	}
	return fm.WriteYAML(r.marker(key, phase), map[string]string{"phase": phase, "started": r.stamp()})
}

// Finish seals the phase: it computes artifacts_hash over what is there now, runs the
// gates and writes the verdict. It does not wait for the pipeline.
func (r *Runner) Finish(key, phase string) (*model.Gate, error) {
	if model.PhaseIndex(phase) < 0 {
		return nil, fmt.Errorf("unknown phase %q", phase)
	}
	g, err := r.evaluate(key, phase)
	if err != nil {
		return nil, err
	}
	_ = os.Remove(r.marker(key, phase))
	return g, nil
}

// GateRun recomputes a verdict. It attaches first, because P5 has no successor whose
// start could do it.
func (r *Runner) GateRun(key, phase string) (*model.Gate, error) {
	if _, _, err := evidence.Attach(r.Root, key, phase, r.EvidenceFrom); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	return r.evaluate(key, phase)
}

func (r *Runner) evidenceSource() string {
	var p model.Project
	if err := fm.ReadYAML(r.abs(".xeno/config/project.yaml"), &p); err == nil && p.Evidence.Source != "" {
		return p.Evidence.Source
	}
	return "ci"
}

// PhaseState is computed, never stored: there is no position that could go stale.
type PhaseState struct {
	Phase  string
	State  string // not-started | running | finished | changed-after-verdict
	Status string // the verdict, where there is one
	Stale  bool   // its predecessor changed after it started
}

func (r *Runner) Status(key string) ([]PhaseState, error) {
	if !fm.Exists(r.abs(model.IntentDir(key))) {
		return nil, fmt.Errorf("no intent %s", key)
	}
	var out []PhaseState
	for i, p := range model.Phases {
		s := PhaseState{Phase: p, State: "not-started"}
		dir := r.abs(model.PhaseDir(key, p))
		switch {
		case fm.Exists(r.marker(key, p)):
			s.State = "running"
		case fm.Exists(filepath.Join(dir, "gate.yaml")):
			g, err := r.readGate(key, p)
			if err != nil {
				return nil, err
			}
			h, err := r.hash(key, p)
			if err != nil {
				return nil, err
			}
			s.Status = g.Status
			s.State = "finished"
			if g.ArtifactsHash != h {
				s.State = "changed-after-verdict"
			}
		}
		if i > 0 && fm.Exists(filepath.Join(dir, "context.lock.yaml")) {
			var lock model.ContextLock
			if err := fm.ReadYAML(filepath.Join(dir, "context.lock.yaml"), &lock); err == nil {
				if h, err := r.hash(key, model.Phases[i-1]); err == nil && h != lock.PredecessorHash {
					s.Stale = true
				}
			}
		}
		out = append(out, s)
	}
	return out, nil
}
