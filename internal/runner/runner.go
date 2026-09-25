// SPDX-License-Identifier: Apache-2.0

// Package runner implements the commands on top of gates and evidence. It is the only
// writer of gate.yaml and context.lock.yaml, and it never writes inside the
// artifacts_hash of a phase whose verdict exists.
package runner

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"go.yaml.in/yaml/v3"
	"time"

	"github.com/triplem/xeno/internal/evidence"
	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/template"
)

// Refusal is a command declining to proceed for a reason a person can act on.
type Refusal struct{ Reason string }

func (r *Refusal) Error() string { return r.Reason }

func refuse(format string, a ...any) error { return &Refusal{fmt.Sprintf(format, a...)} }

type Runner struct {
	Root         string
	Now          func() time.Time
	EvidenceFrom string // stand-in for the CI artifact fetch, see evidence.ManifestEntry
	// PluginSource is where xeno init copies the shipped plugin from. A released
	// runner carries it; here it is the repository being developed.
	PluginSource string
}

func New(root string) *Runner {
	return &Runner{Root: root, Now: func() time.Time { return time.Now().UTC() }}
}

func (r *Runner) abs(rel string) string { return filepath.Join(r.Root, rel) }
func (r *Runner) stamp() string         { return r.Now().Format(time.RFC3339) }

func (r *Runner) marker(key, phase string) string {
	return r.abs(filepath.Join(".xeno/local/runs", key, phase+".lock"))
}

// qualified reads the qualified intent id. It never falls back to the directory name:
// a guessed id would make every trace comparison against it meaningless.
func (r *Runner) qualified(key string) (string, error) {
	var in struct {
		Intent string `yaml:"intent"`
	}
	path := model.IntentDir(key) + "/intent.yaml"
	if err := fm.ReadYAML(r.abs(path), &in); err != nil {
		return "", refuse("%s cannot be read: %v", path, err)
	}
	if in.Intent == "" {
		return "", refuse("%s has no intent field", path)
	}
	// YAML reads " #" as the start of a comment, so an unquoted id with a space before
	// its key loses the key silently. If every file is cut the same way they all agree,
	// and the verdict is green on the wrong id; hence the check on the raw line.
	raw, _ := os.ReadFile(r.abs(path))
	for _, line := range strings.Split(string(raw), "\n") {
		v, ok := strings.CutPrefix(strings.TrimSpace(line), "intent:")
		v = strings.TrimSpace(v)
		if ok && v != "" && v[0] != '"' && v[0] != '\'' && strings.Contains(v, " #") {
			return "", refuse("%s: the intent id is cut off by a YAML comment; write it in quotes, as intent: \"%s\"", path, strings.ReplaceAll(v, " #", "#"))
		}
	}
	var k struct {
		Key string `yaml:"key"`
	}
	if fm.ReadYAML(r.abs(path), &k) == nil && k.Key != "" && k.Key != key {
		return "", refuse("%s names key %s but lies in %s", path, k.Key, model.IntentDir(key))
	}
	return in.Intent, nil
}

func (r *Runner) common(key, phase string) (model.Common, error) {
	q, err := r.qualified(key)
	return model.Common{Intent: q, Phase: phase, Created: r.stamp(),
		SchemaVersion: model.SchemaVersion,
		RunnerVersion: model.RunnerVersion, PluginVersion: model.PluginVersion}, err
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
	common, err := r.common(key, phase)
	if err != nil {
		return nil, err
	}
	prev, _ := r.readGate(key, phase)
	checks := gates.Run(gates.Ctx{Root: r.Root, Key: key, Phase: phase, ArtifactsHash: h, QualifiedID: common.Intent}, prev)
	status, err := gates.Status(checks)
	if err != nil {
		return nil, err
	}
	g := &model.Gate{Common: common, Status: status, RunAt: r.stamp(), ArtifactsHash: h, Checks: checks}
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
	common, err := r.common(key, phase)
	if err != nil {
		return err
	}
	lock := model.ContextLock{Common: common, EvidenceSource: r.evidenceSource()}
	// Which template the phase will be rendered from, recorded because otherwise two
	// projects on the same template version are indistinguishable although one of them
	// overrode it. A repository without a vendored plugin records nothing here and
	// finds out at the first section write, which is where it matters.
	if t, err := template.Load(r.Root, model.TemplateID(phase), r.language()); err == nil {
		lock.TemplateSource = string(t.Source)
	}

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

// language is the language artifacts are written in. English unless a project says
// otherwise; the process layer is English regardless.
func (r *Runner) language() string {
	var p model.Project
	if err := fm.ReadYAML(r.abs(".xeno/config/project.yaml"), &p); err == nil && p.Language.Artifacts != "" {
		return p.Language.Artifacts
	}
	return "en"
}

// SectionSet writes one section of a phase result and renders the whole file again.
//
// Rendering on every write rather than at the end is what keeps the file on disk always
// in rendered form, and it is why the anchors can never be wrong: the agent supplies
// content per section and never writes an anchor.
//
// The frontmatter is carried over untouched where the file exists. Where it does not,
// what the runner knows is written and the rest is left out rather than filled with
// something plausible: model, tool and tool_version come from the harness, and
// secrets_hash from a filter that does not exist yet. G-Schema reports them missing,
// which is the honest state of a phase nothing has produced yet.
func (r *Runner) SectionSet(key, phase, section, content string) (*template.Resolved, error) {
	if model.PhaseIndex(phase) < 0 {
		return nil, fmt.Errorf("unknown phase %q", phase)
	}
	t, err := template.Load(r.Root, model.TemplateID(phase), r.language())
	if err != nil {
		return nil, refuse("%v", err)
	}
	if !t.Has(section) {
		return nil, refuse("template %s has no section %q; it has %s",
			t.Ref(), section, strings.Join(t.Known(), ", "))
	}
	path := r.abs(model.PhaseDir(key, phase) + "/output.md")
	front, sections := map[string]any{}, map[string]string{}
	if b, err := os.ReadFile(path); err == nil {
		f, body, ferr := fm.Split(b)
		if ferr == nil {
			_ = yaml.Unmarshal(f, &front)
		}
		sections = template.Parse(string(body))
	} else {
		common, err := r.common(key, phase)
		if err != nil {
			return nil, err
		}
		front["intent"], front["phase"] = common.Intent, common.Phase
		front["created"] = common.Created
		front["schema_version"] = common.SchemaVersion
		front["runner_version"], front["plugin_version"] = common.RunnerVersion, common.PluginVersion
	}
	front["language"] = t.Bundle.Language
	front["template"] = t.Ref()
	front["strings_hash"] = t.StringsHash
	sections[section] = content

	out := "---\n" + frontmatter(front) + "---\n\n" + t.Render(sections)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	return &t, os.WriteFile(path, []byte(out), 0o644)
}

// frontmatterOrder is the order section 5 lists the fields in, group by group: every
// process file, then what a phase file adds, then what a file produced in a session
// adds, then what a rendered file adds. A map would sort them alphabetically, and every
// artifact written by hand in this repository is in this order.
var frontmatterOrder = []string{
	"intent", "phase", "created", "schema_version", "runner_version", "plugin_version",
	"language", "secrets_hash", "context_hash", "model", "tool", "tool_version",
	"template", "strings_hash", "rules_hash",
	"open_questions", "decisions", "evidence",
}

func frontmatter(front map[string]any) string {
	var b strings.Builder
	written := map[string]bool{}
	emit := func(k string) {
		v, ok := front[k]
		if !ok || written[k] {
			return
		}
		written[k] = true
		out, err := yaml.Marshal(map[string]any{k: v})
		if err != nil {
			return
		}
		b.Write(out)
	}
	for _, k := range frontmatterOrder {
		emit(k)
	}
	rest := make([]string, 0, len(front))
	for k := range front {
		if !written[k] {
			rest = append(rest, k)
		}
	}
	sort.Strings(rest)
	for _, k := range rest {
		emit(k)
	}
	return b.String()
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

// Decide writes a decision onto one finding of a phase's verdict. With CloseObligation
// below it is the only writer of a decision block, which is what keeps that block from
// being something a person edits by hand.
//
// It refuses where the phase directory has moved since the verdict was written. A
// decision carries `against`, the artifacts_hash it was made on, and that is what makes
// it verifiable rather than a claim: written against a hash nobody judged, the field
// would assert something untrue. The finding may not even survive a recompute, since an
// id is a hash over gate, rule, file and cause, so content that moved can take the
// finding with it. The repair is `xeno phase finish`, after which either the finding
// comes back with the same id and the decision is about something that exists, or it
// does not and there was nothing to decide. Recorded as D-6.
func (r *Runner) Decide(key, phase, id, kind, by, reason string) (*model.Gate, error) {
	if kind != "approved" && kind != "overridden" {
		return nil, refuse("unknown decision %q", kind)
	}
	if by == "" || reason == "" {
		return nil, refuse("a decision needs --by and --reason: it is a statement by a person")
	}
	g, err := r.readGate(key, phase)
	if err != nil {
		return nil, refuse("%s %s has no verdict to decide on; run xeno phase finish first", key, phase)
	}
	h, err := r.hash(key, phase)
	if err != nil {
		return nil, err
	}
	if h != g.ArtifactsHash {
		return nil, refuse("%s %s changed after it was judged, so there is nothing to decide against.\n"+
			"  verdict:   %s\n  directory: %s\nRun xeno phase finish and decide again; the finding may not survive it.",
			key, phase, g.ArtifactsHash, h)
	}
	f := findFinding(g, id)
	if f == nil {
		return nil, refuse("%s %s has no finding %s", key, phase, id)
	}
	if f.Decision != nil {
		return nil, refuse("%s is already %s by %s at %s; a decision is not replaced, it is made once",
			id, f.Decision.Type, f.Decision.By, f.Decision.At)
	}
	f.Decision = &model.DecisionOnFinding{
		Type: kind, By: by, At: r.stamp(), Against: h, Reason: reason,
	}
	if kind == "overridden" {
		f.Decision.Obligation = "open"
	}
	return r.rewriteStatus(key, phase, g)
}

// CloseObligation marks the artifacts an override owed as delivered.
//
// Unlike Decide it does not require the directory to match the verdict, and the reason
// is the mechanism rather than an exception to it: an override is taken so that a merge
// can proceed while artifacts are still missing, so by the time they exist the phase has
// moved by definition. Requiring a matching hash here would make an obligation
// impossible to close, which is the opposite of the visibility it exists for.
//
// What it does require is that the finding is still there. An id that is gone took its
// override and its obligation with it.
func (r *Runner) CloseObligation(key, phase, id string) (*model.Gate, error) {
	g, err := r.readGate(key, phase)
	if err != nil {
		return nil, refuse("%s %s has no verdict", key, phase)
	}
	f := findFinding(g, id)
	if f == nil {
		return nil, refuse("%s %s has no finding %s; if the phase was redone, the obligation went with it", key, phase, id)
	}
	if f.Decision == nil || f.Decision.Type != "overridden" {
		return nil, refuse("%s carries no override, so it owes nothing", id)
	}
	if f.Decision.Obligation == "closed" {
		return nil, refuse("%s was closed already", id)
	}
	f.Decision.Obligation = "closed"
	return r.rewriteStatus(key, phase, g)
}

func findFinding(g *model.Gate, id string) *model.Finding {
	for i := range g.Checks {
		for j := range g.Checks[i].Findings {
			if g.Checks[i].Findings[j].ID == id {
				return &g.Checks[i].Findings[j]
			}
		}
	}
	return nil
}

// rewriteStatus derives the status again and writes the verdict back. run_at is left
// alone: nothing was re-run, a person decided about what the last run found.
func (r *Runner) rewriteStatus(key, phase string, g *model.Gate) (*model.Gate, error) {
	status, err := gates.Status(g.Checks)
	if err != nil {
		return nil, err
	}
	g.Status = status
	return g, fm.WriteYAML(r.abs(model.PhaseDir(key, phase)+"/gate.yaml"), g)
}

// IntentClose ends an intent that was dropped rather than merged.
//
// It exists because an intent abandoned in P1 never reaches P5 and would otherwise meet
// no gate at all. A merged intent needs no such command: G-Complete has already run as
// part of P5, where it has to run, since a gate that reports after the merge cannot gate
// it.
//
// The order is: record what the person asserted, then judge it. intent.yaml sits inside
// the intent level hash, so it is written before the hash is taken, and the gate reads
// the file rather than the argument, because a gate that trusts its caller checks
// nothing. A failing verdict is written rather than refused, the same as phase finish:
// the intent is abandoned either way and the record says what is missing.
func (r *Runner) IntentClose(key, reason string) (*model.Gate, error) {
	if strings.TrimSpace(reason) == "" {
		return nil, refuse("--reason is required: why something was dropped is usually worth more than why it was built")
	}
	rel := model.IntentDir(key) + "/intent.yaml"
	var in model.Intent
	if err := fm.ReadYAML(r.abs(rel), &in); err != nil {
		return nil, refuse("%s cannot be read: %v", rel, err)
	}
	if in.Status == "abandoned" {
		return nil, refuse("%s is already abandoned: %s", key, in.Reason)
	}
	in.Status = "abandoned"
	in.Reason = reason
	if err := fm.WriteYAML(r.abs(rel), in); err != nil {
		return nil, err
	}

	h, err := hashing.DirHash(r.Root, model.IntentDir(key), hashing.IntentExcluded)
	if err != nil {
		return nil, err
	}
	check := gates.CompleteOnClose(r.Root, key)
	status, err := gates.Status([]model.Check{check})
	if err != nil {
		return nil, err
	}
	g := &model.Gate{
		Common: model.Common{
			Intent: in.Intent, Created: r.stamp(),
			SchemaVersion: model.SchemaVersion,
			RunnerVersion: model.RunnerVersion, PluginVersion: model.PluginVersion,
		},
		Status: status, RunAt: r.stamp(), ArtifactsHash: h,
		Checks: []model.Check{check},
	}
	return g, fm.WriteYAML(r.abs(model.IntentDir(key)+"/gate.yaml"), g)
}
