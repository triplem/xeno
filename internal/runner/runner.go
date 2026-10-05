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

	"github.com/triplem/xeno/internal/cost"
	"github.com/triplem/xeno/internal/evidence"
	"github.com/triplem/xeno/internal/external"
	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/git"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/index"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/plugin"
	"github.com/triplem/xeno/internal/rules"
	"github.com/triplem/xeno/internal/secrets"
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
	// Base and Head are the commit range under review. An input of the run and
	// deliberately not recorded: after a squash a recorded range would point at commits
	// that no longer exist. Section 9's commit predicates read them through gates.Ctx.
	Base, Head string
	// ToolVersion is the harness reporting its own version, for the one field of section 5
	// the runner cannot know (A35, #181). Reported rather than worked out, because working
	// it out means asking which harness this is, which section 7 forbids: the moment the
	// runner behaves differently per harness, the tools stop being interchangeable.
	//
	// Empty means absent, as before. A35's reason is that a plausible value in a field
	// nobody produced is worse than an absent one, and nothing here produces one.
	//
	// Set from --tool-version, or from XENO_HARNESS_VERSION where the flag is not given.
	// Section 7 lists the variable and calls it recorded only, which is what this is: the
	// value reaches two frontmatter fields and nothing reads it back.
	ToolVersion string
}

func New(root string) *Runner {
	return &Runner{Root: root, Now: func() time.Time { return time.Now().UTC() },
		ToolVersion: os.Getenv(HarnessVersionEnv)}
}

// HarnessVersionEnv is the variable section 7 lists for the harness's own version. The runner
// reads it here and nowhere else, and reads no client specific variable at all: normalising
// one into XENO_* is the entry point's work, which is #183 and is not built.
//
// It is the only XENO_* variable anything in this binary reads today, which is the finding that
// issue records. The three beside it in that list — the plugin root, the data location and the
// harness name — are specified and read by nothing.
const HarnessVersionEnv = "XENO_HARNESS_VERSION"

// HarnessEnv is the variable section 7 lists for which harness is running, recorded only. It
// reaches the `tool` field of section 5 and nothing else reads it back, which is what
// "recorded" means and what keeps the tools interchangeable.
const HarnessEnv = "XENO_HARNESS"

func (r *Runner) abs(rel string) string { return filepath.Join(r.Root, rel) }
func (r *Runner) stamp() string         { return r.Now().Format(time.RFC3339) }

func (r *Runner) marker(key, phase string) string {
	return model.LocalPath(r.Root, "runs", key, phase+".lock")
}

// phaseEnv is where the running phase is written for whatever makes model requests. It
// lives beside the run marker, under the gitignored local directory (A9), because it
// describes a machine's current state and not the trail.
// PhaseEnvFile is the name of the file naming the open phase, inside the local data location.
// Exported because `xeno cost turn` resolves the same file from the command and the two used to
// write the path as separate literals, which is half of what #205 was about.
const PhaseEnvFile = "phase.env"

func (r *Runner) phaseEnv() string { return model.LocalPath(r.Root, PhaseEnvFile) }

// PhaseEnv is what a harness wrapper or a hook sources so that a model request can carry
// the intent and the phase it belongs to. The plan asks `phase start` to export them; a
// child process cannot set its parent's environment, so it writes them where a later
// process can read them, and prints the same on request (A56).
//
// The names are the runner's own and say nothing about a harness or a gateway. Turning
// them into request headers is the plugin's work, per WP11: which variable a harness reads
// and which header a gateway keeps are both agent specific, and the runner holds no agent
// specific logic. The CI wrapper already reads these two names for `gate run`.
func (r *Runner) PhaseEnv(key, phase string) (string, error) {
	id, err := r.qualified(key)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("export XENO_INTENT=%q\nexport XENO_PHASE=%q\n", id, phase), nil
}

func (r *Runner) writePhaseEnv(key, phase string) error {
	content, err := r.PhaseEnv(key, phase)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(r.phaseEnv()), 0o755); err != nil {
		return err
	}
	return os.WriteFile(r.phaseEnv(), []byte(content), 0o644)
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
		RunnerVersion: model.RunnerVersion, PluginVersion: plugin.Version(r.Root)}, err
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

// ChangedSince is the files of the preceding phase's information base whose hashes no longer match
// the tree: what a repeated phase has to read again, and nothing else.
//
// Section 5's sentence is the whole design: "context.lock.yaml already carries paths with hashes,
// so a repeated phase knows which files changed and reads only those." The inputs are the
// predecessor's lock and the tree, so this is derived rather than recorded — the same section says
// the lock states what was declared rather than what was read, and its field list has no entry for
// a changed set. A derivation that is printed cannot drift from its inputs; one that is recorded
// can.
//
// Nothing here reads a file for the agent. The runner says which files moved; what to do about it
// is the agent's to decide, which is also why an empty answer and no predecessor are the same
// answer: nothing to say.
func (r *Runner) ChangedSince(key, phase string) []string {
	idx := model.PhaseIndex(phase)
	if idx <= 0 {
		return nil
	}
	var prev model.ContextLock
	rel := model.PhaseDir(key, model.Phases[idx-1]) + "/context.lock.yaml"
	if err := fm.ReadYAML(r.abs(rel), &prev); err != nil {
		return nil
	}
	var changed []string
	for _, f := range prev.Files {
		h, err := hashing.FileHash(r.abs(f.Path))
		switch {
		case err != nil:
			changed = append(changed, f.Path+" (gone)")
		case h != f.SHA256:
			changed = append(changed, f.Path)
		}
	}
	return changed
}

// headCommit is the commit the phase is started against, read from the clone that is already
// there. Everything that can go wrong — no repository, no commit yet, no git — is the same answer:
// the lock says nothing rather than something made up.
func (r *Runner) headCommit() string {
	return git.Head(r.Root)
}

// rulesApplied is the effective rule set as section 5 records it: a path and a version counter per
// rule. The same Load and Effective that produce rules_hash produce this, so a lock listing three
// rules and a frontmatter hash over four is not a state this code can reach.
func (r *Runner) rulesApplied() []model.AppliedRule {
	read, _ := rules.Load(r.Root)
	effective, _ := rules.Effective(read)
	if len(effective) == 0 {
		return nil
	}
	out := make([]model.AppliedRule, 0, len(effective))
	for _, rl := range effective {
		out = append(out, model.AppliedRule{Path: rl.Path, Version: rl.Version})
	}
	return out
}

// externalGates is the producer gates.Run calls for section 14's declared commands, or nil
// where a project declares none — which is the default, and what keeps the chain of trust closed
// for a project that wants it closed. The declaration is read per run rather than cached,
// because the hash it carries is checked before every run and a stale declaration would check a
// stale hash.
func (r *Runner) externalGates(key, artifactsHash, qualified string) func(string) []model.Check {
	var p model.Project
	if err := fm.ReadYAML(r.abs(model.ProjectFile), &p); err != nil || len(p.ExternalGates) == 0 {
		return nil
	}
	return func(phase string) []model.Check {
		return external.Run(r.Root, phase, artifactsHash, qualified,
			model.PhaseDir(key, phase), p.ExternalGates)
	}
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
	checks := gates.Run(gates.Ctx{Root: r.Root, Key: key, Phase: phase, ArtifactsHash: h,
		QualifiedID: common.Intent, Base: r.Base, Head: r.Head,
		External: r.externalGates(key, h, common.Intent)}, prev)
	// Checked where the verdict is produced rather than where the findings are, so that a
	// second path into it, an external gate above all, meets the same rule as the first.
	if err := gates.Invariants(checks); err != nil {
		return nil, refuse("%v", err)
	}
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
			for _, d := range gateSet(committed, p) {
				res.Divergences = append(res.Divergences, Divergence{k, p, d})
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

// gateSet compares the gates a committed verdict carries against the ones that apply at
// its phase. Without it an absence says two things: that a gate does not apply yet, and
// that it did not run. Recomputing cannot tell them apart, because the recomputation uses
// the same table and leaves out the same gate, so the comparison has to be against the
// table rather than against another run.
//
// It reports in both directions. A gate that vanished is the case worth catching, and a
// gate that appears before its phase is the same defect pointing the other way: a verdict
// claiming a check that could not have happened.
func gateSet(committed *model.Gate, phase string) []string {
	present := map[string]bool{}
	for _, ch := range committed.Checks {
		present[ch.Gate] = true
	}
	var out []string
	for _, id := range gates.Applicable(phase) {
		if !present[id] {
			out = append(out, id+" applies at this phase and the verdict does not carry it")
		}
		delete(present, id)
	}
	ids := make([]string, 0, len(present))
	for id := range present {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		out = append(out, id+" is in the verdict and does not apply at this phase")
	}
	return out
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
	// A phase that has been judged is not started again. Start writes context.lock.yaml
	// unconditionally, and the lock is inside artifacts_hash, so a second start rewrites a
	// sealed artifact: section 11's "what is sealed is never rewritten" caught that as a
	// divergence afterwards, and this refuses it at the point of the act. The lock is also
	// the only record of what the phase was given, which is what Changed compares against,
	// so the overwrite destroyed the answer to "what changed" as well (#215).
	//
	// Redoing the work does not need a second start: write the sections again and run
	// phase finish, which recomputes the hash and carries the decisions forward. Starting
	// over from nothing is deliberate and stays possible, by removing the phase directory —
	// the verdict alone is no longer enough, because from #225 the artifact left behind is
	// itself a phase under way.
	if fm.Exists(r.abs(model.PhaseDir(key, phase) + "/gate.yaml")) {
		return refuse("%s has a verdict; redo the work with section set and phase finish, "+
			"or remove %s to start it over", phase, model.PhaseDir(key, phase))
	}
	// And the state between the two guards above: sections written, no verdict yet, and the
	// marker gone. That is a phase under way, which section 6 already says is not started again;
	// the marker is only the machine-local evidence of it, and its absence is ordinary — A9 puts
	// it under .xeno/local/ because it describes a machine and not the trail, so a phase begun
	// elsewhere has none here and retention expires it at thirty days.
	//
	// Starting again would rewrite context.lock.yaml, which is inside artifacts_hash and is the
	// only record of what the phase was given: the two reasons #215 gives for refusing a judged
	// phase, both true one step earlier. The rewrite moves `created`, so the artifact's
	// context_hash stops matching the lock and phase finish reports red on G-Schema two commands
	// later, which is not where the mistake was made (#225).
	//
	// The condition is the artifact and not a comparison against the lock. Before a second start
	// the two still agree — the staleness is what the start causes, not what it finds — so
	// comparing them here would pass in exactly the case the damage is about to be done.
	if fm.Exists(r.abs(model.PhaseDir(key, phase) + "/output.md")) {
		return refuse("%s is already under way; carry on with section set and phase finish, "+
			"which need no second start, or remove %s to start the phase over",
			phase, model.PhaseDir(key, phase))
	}
	common, err := r.common(key, phase)
	if err != nil {
		return err
	}
	lock := model.ContextLock{Common: common, EvidenceSource: r.evidenceSource()}
	files, err := r.informationBase(key)
	if err != nil {
		return err
	}
	lock.Files = files
	// Section 5 writes both of these into the lock. repo_commit is absent where the repository
	// has none or is not one, because a lock naming the commit a phase ran against is a claim and
	// an invented value would be a false one. rules_applied answers what rules_hash cannot —
	// which rules, and which revision of each — and is absent rather than empty where no rule is
	// in force, since an empty list says a set was resolved and came out empty (A74).
	lock.RepoCommit = r.headCommit()
	lock.RulesApplied = r.rulesApplied()
	// The plugin block section 5 enumerates, and the other half of the sentence that explains
	// plugin_version: the frontmatter names what was used and the lock proves it with a hash.
	// Absent together where there is no vendored plugin, because a block naming a version with
	// no hash, or a hash with no version, would be half a claim (#177).
	if v, h := plugin.Version(r.Root), plugin.Hash(r.Root); v != "" || h != "" {
		lock.Plugin = &model.LockPlugin{Version: v, SHA256: h}
	}
	// Which template the phase will be rendered from, recorded because otherwise two
	// projects on the same template version are indistinguishable although one of them
	// overrode it. A repository without a vendored plugin records nothing here and
	// finds out at the first section write, which is where it matters.
	if t, err := template.Load(r.Root, model.TemplateID(phase), r.language()); err == nil {
		lock.TemplateSource = string(t.Source)
	}

	if idx > 0 {
		h, err := r.predecessorAllowsStart(key, model.Phases[idx-1])
		if err != nil {
			return err
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
	if err := r.writePhaseEnv(key, phase); err != nil {
		return err
	}
	return fm.WriteYAML(r.marker(key, phase), map[string]string{"phase": phase, "started": r.stamp()})
}

// predecessorAllowsStart is A12: the next phase starts only on a decided predecessor. It
// returns the hash the lock records, which is what G-Freshness compares against later.
//
// Building on a red verdict moves the failure downstream instead of resolving it, and
// building on a provisional one builds on evidence that has not arrived. So a provisional
// predecessor is attached to and judged again first: the pipeline may have finished since,
// and only evidence/ and gate.yaml of that one phase are written.
func (r *Runner) predecessorAllowsStart(key, pred string) (string, error) {
	g, err := r.readGate(key, pred)
	if err != nil {
		return "", refuse("%s has no verdict yet; run xeno phase finish for it first", pred)
	}
	h, err := r.hash(key, pred)
	if err != nil {
		return "", err
	}
	if g.ArtifactsHash != h {
		return "", refuse("%s changed after its verdict; run xeno phase finish for it again", pred)
	}
	if g.Status == "provisional" {
		att, err := evidence.Attach(r.Root, key, pred, r.EvidenceFrom)
		if err != nil {
			return "", err
		}
		if g, err = r.evaluate(key, pred); err != nil {
			return "", err
		}
		// An entry the pipeline published wrong keeps the phase provisional exactly as a
		// missing one does, and waiting is the wrong advice for it: nothing arrives by
		// waiting for a job that has already run. So the refusal names it instead, and each
		// sentence carries its own remedy, because the three reasons are fixed in three
		// different places (#221).
		if g.Status == "provisional" && len(att.Declined) > 0 {
			return "", refuse("%s waits for evidence the pipeline published in a form nothing can record:\n  %s\nWaiting will not help; republish it as each line says.",
				pred, strings.Join(att.Declined, "\n  "))
		}
		if g.Status == "provisional" {
			return "", refuse("%s still waits for evidence from the pipeline; start again once it has run", pred)
		}
	}
	if !Decided(g.Status) {
		return "", refuse("%s is red; decide every failing finding first: fix it, approve it or override it", pred)
	}
	return h, nil
}

// informationBase is the files a phase is given: the intent's scope, resolved.
//
// The scope is P0's artifact and applies to every phase of the intent, which is what section
// 12 means by a phase reading what it names; a phase does not get a scope of its own, so
// there is one budget per intent rather than six.
//
// A repository without a scope gets an empty list, and the staleness half of G-Freshness then
// has nothing to compare. That is a smaller claim than an empty scope would be: nothing was
// declared, rather than nothing was read. From #217 a P0 cannot be finished without one, so
// the case is a phase started before that rule rather than a project opting out.
func (r *Runner) informationBase(key string) ([]model.ContextFile, error) {
	s, err := r.readScope(key)
	if err != nil || s == nil {
		return nil, err
	}
	return r.resolveScope(s)
}

// firstMatch is the index of the first pattern that claims the path, or -1.
func firstMatch(patterns []string, path string) int {
	for i, pattern := range patterns {
		if model.MatchPath(pattern, path) {
			return i
		}
	}
	return -1
}

func matchesAny(patterns []string, path string) bool {
	for _, pattern := range patterns {
		if model.MatchPath(pattern, path) {
			return true
		}
	}
	return false
}

// Finish seals the phase: it computes artifacts_hash over what is there now, runs the
// gates and writes the verdict. It does not wait for the pipeline.
//
// A summary writes digest.md first. Section 5 divides the labour — the agent supplies the summary
// text, the runner writes the file — and section 6's working sequence puts the digest here, which
// is also the earliest honest moment to summarise a phase. Before the gates, because the digest is
// inside artifacts_hash: written afterwards it would seal a hash over a tree that lacked the file
// and no verdict would ever match the directory again.
//
// An empty summary writes nothing. Section 5 says the supported path is not an enforced one, so a
// phase finished without a summary is judged exactly as before, with G-Schema reporting the digest
// missing where it is missing.
func (r *Runner) Finish(key, phase, summary string) (*model.Gate, error) {
	if model.PhaseIndex(phase) < 0 {
		return nil, fmt.Errorf("unknown phase %q", phase)
	}
	// Section 5 puts the context scope among P0's artifacts and calls it a budget P0 produces,
	// and for a hundred intents nothing produced one, because no command wrote it and no check
	// asked for it (#217). The requirement is here and not in a gate for the reason D-2
	// records: Verify recomputes a verdict and compares it against the committed one, so a
	// check in a gate re-judges every phase already sealed, which is what A74 was written
	// about. A command is not re-run, so this binds every new P0 and reaches no old one.
	//
	// First statement, before the digest and the cost are written, so a refused finish writes
	// nothing at all. Not in phase start of P1 either: a P0 can only be left without a scope if
	// it was allowed to finish.
	if phase == model.Phases[0] &&
		!fm.Exists(r.abs(model.PhaseDir(key, phase)+"/"+model.ContextScope)) {
		return nil, refuse("%s has no %s, which section 5 has P0 produce as an artifact of its "+
			"own and the phases after it read; write it with xeno scope set --intent %s",
			phase, model.ContextScope, key)
	}
	if strings.TrimSpace(summary) != "" {
		if err := r.writeDigest(key, phase, summary); err != nil {
			return nil, err
		}
	}
	// Before the verdict, and it could be anywhere: cost.yaml lies outside artifacts_hash, by
	// section 11, so that a figure arriving after a verdict cannot invalidate it. A phase with
	// nothing attributed writes no file, because section 11 says a phase without one is complete.
	if err := r.writeCost(key, phase); err != nil {
		return nil, err
	}
	g, err := r.evaluate(key, phase)
	if err != nil {
		return nil, err
	}
	_ = os.Remove(r.marker(key, phase))
	// A phase that has ended attributes nothing. Left behind, the file would put the next
	// session's requests on a phase that is sealed, which is worse than attributing none.
	_ = os.Remove(r.phaseEnv())
	return g, nil
}

// writeDigest writes the summary of a phase with the frontmatter section 5 requires of a file
// produced in a session. Three groups, not two: a digest is neither rendered nor covered by a
// rule set, so it carries no template, no strings_hash and no rules_hash.
//
// It carries no secrets_hash either, and that absence is the honest half of this writer. Section
// 5 gives the runner two jobs, filtering and writing, and there is no filter to do the first: the
// effective filter is a shipped pattern file plus a project's additions, and no such file exists
// in this tree. A hash over nothing would assert that a digest passed through a filter, which is
// the one claim in a digest a reader would take on trust. G-Schema reports the field missing,
// which is what is true. Whoever ships the filter adds the filtering here.
func (r *Runner) writeDigest(key, phase, summary string) error {
	common, err := r.common(key, phase)
	if err != nil {
		return err
	}
	front := map[string]any{
		"intent": common.Intent, "phase": common.Phase, "created": common.Created,
		"schema_version": common.SchemaVersion,
		"runner_version": common.RunnerVersion, "plugin_version": common.PluginVersion,
		"language": r.language(),
	}
	if h, err := hashing.FileHash(r.abs(model.PhaseDir(key, phase) + "/context.lock.yaml")); err == nil {
		front["context_hash"] = h
	}
	if tool, mdl := r.agent(); tool != "" || mdl != "" {
		if mdl != "" {
			front["model"] = mdl
		}
		if tool != "" {
			front["tool"] = tool
		}
	}
	// Carried over from the artifact beside it rather than taken as an input of its own.
	// The digest is the same phase produced in the same session, `output.md` already records
	// which harness wrote it, and this function is rerun by every `phase finish` — an input
	// would have to be supplied again on each one, which is the hand edit this replaces
	// (#181). A reported version wins, for a digest written before any section was.
	if v := r.ToolVersion; v != "" {
		front["tool_version"] = v
	} else if v := r.recordedToolVersion(key, phase); v != "" {
		front["tool_version"] = v
	}
	// Section 16: the agent writes the summary text, the runner filters it against the
	// effective filter, hashes it and writes the file, so the filtering is deterministic and
	// outside the model's reach. An empty filter, which is a repository before its plugin is
	// vendored, redacts nothing and writes no field rather than a hash of nothing.
	filter, err := secrets.Load(r.Root)
	if err != nil {
		return err
	}
	summary = filter.Redact(summary)
	if h := filter.Hash(); h != "" {
		front["secrets_hash"] = h
	}
	body := strings.TrimRight(summary, "\n") + "\n"
	out := "---\n" + frontmatter(front) + "---\n" + body
	path := r.abs(model.PhaseDir(key, phase) + "/digest.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(out), 0o644)
}

// recordedToolVersion reads what the phase's output.md says wrote it. A phase with no artifact,
// no frontmatter or no such field answers empty, which is the field absent and is the state of
// every phase written before there was a way to report it.
func (r *Runner) recordedToolVersion(key, phase string) string {
	var front struct {
		ToolVersion string `yaml:"tool_version"`
	}
	if _, err := fm.ReadFront(r.abs(model.PhaseDir(key, phase)+"/output.md"), &front); err != nil {
		return ""
	}
	return front.ToolVersion
}

// writeCost writes the cost record of section 11 from what the hook attributed to this phase.
//
// The counts come from a ledger a hook appends to, one line per turn, naming the phase the runner
// said was open. Nothing is inferred: a turn spent with no phase open is in the ledger as such,
// and this sums only what names the phase. Measured against this repository, attributing by each
// phase's own window instead captures about seven per cent of a session, because the work is done
// before `phase start` is called, so the gap is left visible rather than distributed (A63).
//
// evidence: self-reported is mandatory and accurate here: a count read from a file the harness
// wrote, attributed by a marker the runner wrote.
func (r *Runner) writeCost(key, phase string) error {
	common, err := r.common(key, phase)
	if err != nil {
		return err
	}
	totals, sessions, err := cost.ForPhase(r.Root, common.Intent, phase)
	if err != nil || totals.Zero() {
		return nil // no ledger, or nothing attributed: a phase without a cost record is complete
	}
	return fm.WriteYAML(r.abs(model.PhaseDir(key, phase)+"/cost.yaml"), model.Cost{
		Common: common, Evidence: "self-reported",
		TokensIn: totals.In, TokensOut: totals.Out, TokensCached: totals.Cached,
		Sessions: sessions,
	})
}

// GateRun recomputes a verdict. It attaches first, because P5 has no successor whose
// start could do it.
func (r *Runner) GateRun(key, phase string) (*model.Gate, error) {
	if _, err := evidence.Attach(r.Root, key, phase, r.EvidenceFrom); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	return r.evaluate(key, phase)
}

// agent returns the tool and the model a project records, and empty strings where the file,
// the block or a field is absent. Section 12 puts both in project.yaml, so they are not among
// the fields that come from the harness: A35 listed five and the reason covers three, which is
// what this reads (#120).
//
// Absent rather than defaulted, because the row being amended says why: a plausible value in a
// field nobody produced is worse than an absent one, and the only tool this repository has ever
// used would be exactly such a plausible value.
func (r *Runner) agent() (tool, mdl string) {
	var p model.Project
	if err := fm.ReadYAML(r.abs(projectConfig), &p); err != nil {
		return os.Getenv(HarnessEnv), ""
	}
	tool = p.Agent.Tool
	// Section 7 lists XENO_HARNESS and calls it recorded only, which is this field. It beats
	// the project's declaration where both are there: `agent.tool` is what a project says it
	// uses and the variable is what is actually running, and a session under a harness the
	// project did not name is the case the pair exists to make visible rather than hide.
	//
	// Read and never branched on, which section 7 says in as many words and the `verify`
	// workflow now checks: the moment the runner behaves differently per harness, the tools
	// stop being interchangeable. This is A35's second amendment amended (#183).
	if h := os.Getenv(HarnessEnv); h != "" {
		tool = h
	}
	return tool, p.Agent.Model.Default
}

// language is the language artifacts are written in. English unless a project says
// otherwise; the process layer is English regardless.
func (r *Runner) language() string {
	var p model.Project
	if err := fm.ReadYAML(r.abs(projectConfig), &p); err == nil && p.Language.Artifacts != "" {
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
//
// context_hash is written on every render, not only on the first one, because the
// frontmatter of an existing file is carried over from whatever wrote it. The value does
// not drift within a phase: phase start writes context.lock.yaml once and nothing
// refreshes it, so every section write of one phase hashes the same bytes. Where the lock
// is absent the field is left out, which is the case of a file written outside a started
// phase; its absence there is G-Freshness's finding about the lock.
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
	// The filter output.md was written under. The prose itself is not redacted: section 16
	// puts the filtering on the digest, which is the file that leaves a session, and a filter
	// over the agent's own prose would redact a discussion of a pattern by that pattern.
	if filter, ferr := secrets.Load(r.Root); ferr == nil {
		if h := filter.Hash(); h != "" {
			front["secrets_hash"] = h
		}
	}
	// The rule set this phase is judged against. Written here rather than by the harness,
	// which is what wrote by-hand into every artifact before internal/rules existed; the
	// placeholder stays readable in those, because the field sits inside artifacts_hash and
	// rewriting one would change every verdict in its intent (A66).
	//
	// A tree with a problem in it still produces a hash, over the rules that were usable.
	// G-Rules reports the problem and the phase is red, and the field says which set
	// actually resolved rather than going absent and saying nothing.
	read, _ := rules.Load(r.Root)
	effective, _ := rules.Effective(read)
	front["rules_hash"] = rules.Hash(effective)
	// Section 12 records both in project.yaml, so they are not among the fields that come
	// from the harness; absent where the project does not say, never defaulted (A35, #120).
	if tool, mdl := r.agent(); tool != "" || mdl != "" {
		if mdl != "" {
			front["model"] = mdl
		}
		if tool != "" {
			front["tool"] = tool
		}
	}
	// The one field that does come from the harness, where the harness said. Written only
	// when reported, and a report never erases what an earlier write carried over: a phase
	// whose first sections named the version and whose last did not is one session, and the
	// field describes the session rather than the invocation.
	if r.ToolVersion != "" {
		front["tool_version"] = r.ToolVersion
	}
	lockPath := r.abs(model.PhaseDir(key, phase) + "/context.lock.yaml")
	if h, err := hashing.FileHash(lockPath); err == nil {
		front["context_hash"] = h
	}
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
	"open_questions", "decisions", "evidence", "review_checklist",
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

// symbolIndex reads the index the project produced, or says why there is none.
//
// The reason is returned rather than dropped, because it is what the tools entry of
// context.lock.yaml records and because somebody asking why a phase ran without an index
// deserves the sentence rather than silence. Nothing here fails: an absent index is ordinary.
func (r *Runner) symbolIndex(now time.Time) (*index.Index, string) {
	var p model.Project
	if err := fm.ReadYAML(r.abs(projectConfig), &p); err != nil {
		return nil, "project.yaml cannot be read, so no index is configured"
	}
	path := p.Index.Path
	if path != "" && !filepath.IsAbs(path) {
		path = r.abs(path)
	}
	maxAge := index.DefaultMaxAge
	if p.Index.MaxAgeHours > 0 {
		maxAge = time.Duration(p.Index.MaxAgeHours) * time.Hour
	}
	return index.Load(path, maxAge, now)
}

func (r *Runner) evidenceSource() string {
	var p model.Project
	if err := fm.ReadYAML(r.abs(projectConfig), &p); err == nil && p.Evidence.Source != "" {
		return p.Evidence.Source
	}
	return "ci"
}

// Decided reports whether a verdict settles its phase. Red is undecided by definition and a
// provisional one waits for evidence, so neither lets the next phase start and neither means the
// work is finished. Everything else does: approved and overridden are decisions a person took, and
// green needed none.
//
// It is one predicate because two parts of the tool ask the question. predecessorAllowsStart asks
// it before it lets a phase begin, and the intent listing asks it of P5 to say whether an intent is
// complete; a second list of accepted statuses could drift from the first and the two would then
// disagree about the same intent (#132).
func Decided(status string) bool { return status != "red" && status != "provisional" }

// IntentsRoot is where intent directories lie, relative to the repository root.
const IntentsRoot = ".xeno/intents"

// projectConfig is the file section 12 configures the project in. Named here because the
// readers of it in this file each want it twice, once to read and once to say which file
// an error is about.
const projectConfig = ".xeno/config/project.yaml"

// IntentSummary is one row of the listing: what an intent asserts about itself, plus how far
// it got. Nothing here is stored; it is read from intent.yaml and the phase verdicts.
type IntentSummary struct {
	Key     string
	Created string // intent.yaml's created, whole, or empty where it cannot be read
	// State is computed, not read: abandoned, complete, or the phase the work has reached.
	//
	// intent.yaml's own status carries two values, in-progress and abandoned, because section 5
	// gives it no third and section 8 says why: a merged intent's record is its P5 phase, so
	// completion is that verdict rather than a flag. Printing the field made every intent that
	// was not dropped read as unfinished, which is what #132 was about.
	State   string
	Phase   string // the furthest phase carrying a verdict, and that verdict
	Verdict string
	Problem string // why this row could not be read, where that happened
	// What the intent asked and what it settled, over every phase. Counted rather than
	// judged: a phase with nothing to ask is the ordinary case and most are, so these are
	// a figure and nothing reads them as a condition.
	Questions int
	Decisions int
}

// AskedNothing reports an intent that has reached the review phase without one question and
// without one decision anywhere in it.
//
// It is the figure #188 asks for and it is deliberately not a gate. G-Questions verifies
// that a question which was raised is resolved; it cannot see one that was never asked,
// because nothing declares how many there should have been, and a gate that demanded one
// would be answered with invented questions. What is being reported is not a phase without
// a question but an intent reaching the merge having asked nothing at all.
//
// From the review phase, for the reason G-Questions runs from there: an open question in P1
// is normal and should not hold up the work, and before a merge it is not normal any more.
// Not for an abandoned intent, for the reason that gate leaves one alone as well: abandoning
// means giving up, and reproaching it would only produce the invented question.
func (s IntentSummary) AskedNothing() bool {
	if s.Questions > 0 || s.Decisions > 0 || s.State == "abandoned" {
		return false
	}
	return s.State == "complete" || s.State == model.Phases[len(model.Phases)-1]
}

// Intents lists every intent in the order it was created.
//
// The order is read from the created field rather than from the key, which is the point: a key
// sorts one way and a recorded field can be asked in any order. Intent keys carried the issue
// number for exactly this purpose and sorted by when an issue was filed, which is not when the
// work happened (#118).
//
// An intent whose intent.yaml cannot be read, or which records no created, is listed with the
// reason in place of the date rather than left out. A record missing from a listing is worse
// than one that looks wrong in it, and the empty date sorts it to the front of an ascending
// order, which is deliberate: it belongs where somebody will see it.
//
// Ties break on the key, so two runs over one tree print the same thing. That is the only
// decision the key is still allowed to make.
func (r *Runner) Intents() ([]IntentSummary, error) {
	entries, err := os.ReadDir(r.abs(IntentsRoot))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // a repository that has started no intent is not in error
		}
		return nil, err
	}
	var out []IntentSummary
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		out = append(out, r.Summarise(e.Name()))
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Created != out[j].Created {
			return out[i].Created < out[j].Created
		}
		return out[i].Key < out[j].Key
	})
	return out, nil
}

// Summarise reads one intent. It reports rather than fails: the listing's job is to show what
// is there, and an unreadable intent is something to show.
//
// Exported because the one-intent form of the status command prints the same figure the
// listing marks, and a second walk computing it would be a second definition of it.
func (r *Runner) Summarise(key string) IntentSummary {
	s := IntentSummary{Key: key}
	var in model.Intent
	if err := fm.ReadYAML(r.abs(model.IntentDir(key)+"/intent.yaml"), &in); err != nil {
		s.Problem = "intent.yaml cannot be read"
		return s
	}
	// Kept whole, because the sort is over it: two intents of one day are ordered by their
	// time, and truncating to the date here would fall through to the key tie break and
	// reproduce the order this listing exists to correct. The command prints the date.
	s.Created = in.Created
	if s.Created == "" {
		s.Problem = "intent.yaml records no created"
	}
	if states, err := r.Status(key); err == nil {
		for _, st := range states {
			if st.Status != "" {
				s.Phase, s.Verdict = st.Phase, st.Status
			}
		}
		s.State = state(in.Status, states)
	}
	raised, taken := r.exchange(key)
	s.Questions, s.Decisions = len(raised), len(taken)
	return s
}

// state is the three answers in order. Abandoned wins and is read rather than computed, because an
// intent dropped in P1 has no P5 to derive anything from, which is the case intent close exists
// for. Otherwise a decided P5 is complete, and otherwise the work is where it has got to.
//
// The word is complete rather than closed: intent close writes abandoned in this tool, so a row
// reading closed would invite the command that falsifies it. It is a rendering of the verdict and
// not a value of the schema, which is why it is defined here and not in section 5.
func state(stored string, states []PhaseState) string {
	if stored == "abandoned" {
		return "abandoned"
	}
	last := model.Phases[len(model.Phases)-1]
	reached := ""
	for _, st := range states {
		if st.Status != "" || st.State == "running" {
			reached = st.Phase
		}
		if st.Phase == last && st.Status != "" && Decided(st.Status) {
			return "complete"
		}
	}
	if reached == "" {
		return "no phases"
	}
	return reached
}

// CompletenessResult is what the merge check reads: which intents a change touches, and which
// of them have neither finished nor been closed.
type CompletenessResult struct {
	Touched    []string        // the keys the range touches, by key
	Unfinished []IntentSummary // those among them that are neither complete nor abandoned
}

// Completeness answers, for one commit range, whether every intent the change touches has
// reached one of the two endings section 8 gives an intent: a decided P5, or `xeno intent
// close`. An intent that simply stopped is neither, and before this nothing said so —
// G-Complete runs at P5 and at intent close, and an intent that reaches neither has no phase
// for the gate to run in (#206).
//
// The touched intents and not the whole trail. A branch answers for what it changes; a check
// over every intent in the repository would turn one unclosed intent from last month into a
// red pull request for everybody, and the way to pass it would be to close intents nobody had
// finished.
//
// The state is read from summarise, which is the one definition of it in this tree. It is the
// same three answers `xeno intent status` prints, which is the point: the answer existed and
// nothing read it as a condition.
//
// An intent whose record cannot be read counts as unfinished. The listing reports such a row
// rather than leaving it out, because a record missing from a listing is worse than one that
// looks wrong in it; here the answer is a verdict, so the same fact has to refuse.
func (r *Runner) Completeness(base, head string) (*CompletenessResult, error) {
	paths, err := git.Paths(r.Root, base, head, IntentsRoot)
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	res := &CompletenessResult{}
	for _, p := range paths {
		rest, ok := strings.CutPrefix(filepath.ToSlash(p), IntentsRoot+"/")
		if !ok {
			continue
		}
		key, _, _ := strings.Cut(rest, "/")
		if key == "" || seen[key] {
			continue
		}
		seen[key] = true
		res.Touched = append(res.Touched, key)
	}
	sort.Strings(res.Touched)
	for _, key := range res.Touched {
		s := r.Summarise(key)
		if s.State == "complete" || s.State == "abandoned" {
			continue
		}
		if s.State == "" && s.Problem == "" {
			s.Problem = "its phases cannot be read"
		}
		res.Unfinished = append(res.Unfinished, s)
	}
	return res, nil
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
		// An artifact with no verdict is a phase under way, whether or not this machine holds
		// the marker. The marker is machine-local (A9) and expires at thirty days, so keying the
		// state on it alone reported a phase begun elsewhere as not-started — the first of the
		// three gaps #225 names, and what let `phase start` be suggested for a phase that had
		// already been started.
		case fm.Exists(filepath.Join(dir, "output.md")):
			s.State = "running"
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

// IntentStart creates an intent, which until now was the one artifact of this process
// that no command wrote. Three of its seven fields were guessed where the runner knows
// them: a typed timestamp, and two version strings copied from memory, so that every
// intent.yaml here records a version the artifacts beside it do not (#177, #179).
//
// It takes the one thing the runner cannot know, which is the issue. Everything else is
// derived: the key continues the sequence the intents directory holds, the host and the
// repository come from the tracker block, `created` is the runner's clock, the version
// fields are the runner's own, and `status` is the only value a person ever typed that
// the runner would have defaulted to anyway — `intent close` writes the other one and
// there is no third (A20).
//
// It refuses where the directory exists, as `phase start` refuses a running phase. A
// hand written intent.yaml is not touched by this and goes on working: seventy-nine of
// them exist, they sit inside `artifacts_hash`, and rewriting one changes every verdict
// that sealed it.
func (r *Runner) IntentStart(key, issue string) (*model.Intent, error) {
	var p model.Project
	// Read before the key is derived, so that a repository with no configuration is told
	// about the field it is missing rather than about the sequence it has.
	if err := fm.ReadYAML(r.abs(projectConfig), &p); err != nil && !os.IsNotExist(err) {
		return nil, refuse("%s cannot be read: %v", projectConfig, err)
	}
	id, err := p.Tracker.Qualified(issue)
	if err != nil {
		return nil, refuse("%v", err)
	}
	if key == "" {
		if key, err = r.nextKey(); err != nil {
			return nil, refuse("%v", err)
		}
	}
	rel := model.IntentDir(key)
	if fm.Exists(r.abs(rel)) {
		return nil, refuse("%s exists already; an intent is created once, and what it is "+
			"for is in %s", rel, rel+"/intent.yaml")
	}
	in := &model.Intent{
		Intent: id, Key: key, Status: "in-progress", Created: r.stamp(),
		SchemaVersion: model.SchemaVersion,
		RunnerVersion: model.RunnerVersion, PluginVersion: plugin.Version(r.Root),
	}
	return in, fm.WriteYAML(r.abs(rel+"/intent.yaml"), in)
}

// nextKey reads the sequence off the directory. A missing intents directory is the same
// case as an empty one: there is no sequence, so the first key is named rather than
// derived, which is what model.NextKey says.
func (r *Runner) nextKey() (string, error) {
	entries, err := os.ReadDir(r.abs(IntentsRoot))
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}
	var names []string
	for _, e := range entries {
		if e.IsDir() {
			names = append(names, e.Name())
		}
	}
	return model.NextKey(names)
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
	// As evaluate does for a phase, and for the reason Invariants gives about itself: what breaks
	// A25's rule is a second path into a verdict rather than a change to carryForward, and this is
	// that second path (#138).
	if err := gates.Invariants([]model.Check{check}); err != nil {
		return nil, refuse("%v", err)
	}
	status, err := gates.Status([]model.Check{check})
	if err != nil {
		return nil, err
	}
	g := &model.Gate{
		Common: model.Common{
			Intent: in.Intent, Created: r.stamp(),
			SchemaVersion: model.SchemaVersion,
			RunnerVersion: model.RunnerVersion, PluginVersion: plugin.Version(r.Root),
		},
		Status: status, RunAt: r.stamp(), ArtifactsHash: h,
		Checks: []model.Check{check},
	}
	return g, fm.WriteYAML(r.abs(model.IntentDir(key)+"/gate.yaml"), g)
}

// ---- The learning record

// learningPath is the record of a phase, or of the intent where no phase is given. Both
// exist: section 10 says a learning is owed at the end of every phase and once more when
// an intent closes, and G-Complete reads the intent level one from `intent close`.
func (r *Runner) learningPath(key, phase string) string {
	if phase == "" {
		return model.IntentDir(key) + "/learning.yaml"
	}
	return model.PhaseDir(key, phase) + "/learning.yaml"
}

// RecordLearning writes what a phase learned, which was the last artifact of this process
// that no command wrote.
//
// The split is `section set`'s: the content is the agent's, because an observation about a
// phase can come from nothing but whoever did it, and the header is the runner's — the same
// six fields it already writes into the three artifacts beside this one. A record typed by
// hand carries `0.1.0-dev` where they carry the commit, which is two strings for one build
// in one directory and is what #179 said about intent.yaml (#195).
//
// It does not refuse a sealed phase, which is `SectionSet`'s behaviour and the reason is the
// same: the file is inside artifacts_hash, so a write after the verdict makes `gate verify`
// report a divergence and `phase finish` is what judges it again. A command that refused
// would be refusing the first half of a re-judgement.
//
// Entries accumulate. A phase learns more than one thing often enough, and a second call
// appending is what keeps the first one from having to be remembered and retyped.
func (r *Runner) RecordLearning(key, phase string, noFinding bool, e model.LearningEntry) (*model.Learning, error) {
	if err := r.checkLearningArgs(noFinding, e); err != nil {
		return nil, err
	}
	rel := r.learningPath(key, phase)
	rec := model.Learning{}
	// A record that is not there is the ordinary case: this is usually what creates it.
	if err := fm.ReadYAML(r.abs(rel), &rec); err != nil && !os.IsNotExist(err) {
		return nil, refuse("%s cannot be read: %v", rel, err)
	}
	if rec.NoFinding && !noFinding {
		return nil, refuse("%s states no_finding, so an observation contradicts it; "+
			"withdraw the claim by hand before recording one", rel)
	}
	if noFinding && len(rec.Learnings) > 0 {
		return nil, refuse("%s already records %d learning(s), so no_finding contradicts it",
			rel, len(rec.Learnings))
	}
	common, err := r.common(key, phase)
	if err != nil {
		return nil, err
	}
	// The header is rewritten rather than carried over, because it describes the binary that
	// wrote the file and this write is the one that is happening.
	rec.Common = common
	if noFinding {
		rec.NoFinding = true
	} else {
		rec.Learnings = append(rec.Learnings, e)
	}
	return &rec, fm.WriteYAML(r.abs(rel), rec)
}

// checkLearningArgs refuses before anything is read, so a wrong category never reaches the
// file and G-Learning never has to report what a writer could have.
func (r *Runner) checkLearningArgs(noFinding bool, e model.LearningEntry) error {
	empty := e == model.LearningEntry{}
	switch {
	case noFinding && !empty:
		return refuse("--no-finding states there was nothing to record, so it takes none of " +
			"--category, --observation, --proposal or --target")
	case noFinding:
		return nil
	case empty:
		return refuse("a learning needs --category, --observation, --proposal and --target, " +
			"or --no-finding where there is nothing to record")
	}
	if !model.OneOf(e.Category, model.LearningCategories) {
		return refuse("--category is one of %s", strings.Join(model.LearningCategories, ", "))
	}
	for _, f := range []struct{ flag, v string }{
		{"--observation", e.Observation}, {"--proposal", e.Proposal}, {"--target", e.Target},
	} {
		if strings.TrimSpace(f.v) == "" {
			return refuse("%s is required: section 10 defines all four, and a record missing "+
				"one is not a proposal anybody can act on", f.flag)
		}
	}
	return nil
}

// ---- The assumption register

// assumptionsPath is at intent level: the register is carried forward across all phases
// and is therefore not part of any phase's artifacts_hash.
func (r *Runner) assumptionsPath(key string) string {
	return r.abs(model.IntentDir(key) + "/assumptions.yaml")
}

func (r *Runner) readRegister(key string) (*model.Assumptions, error) {
	var reg model.Assumptions
	if err := fm.ReadYAML(r.assumptionsPath(key), &reg); err != nil {
		return nil, refuse("%s has no assumption register; %s is missing", key, model.IntentDir(key)+"/assumptions.yaml")
	}
	return &reg, nil
}

func (r *Runner) writeRegister(key string, reg *model.Assumptions) error {
	reg.Updated = r.stamp()
	return fm.WriteYAML(r.assumptionsPath(key), reg)
}

// nextAssumptionID continues the register's own numbering rather than counting its
// entries, so that a removed record does not hand its id to the next one.
func nextAssumptionID(reg *model.Assumptions) string {
	high := 0
	for _, a := range reg.Assumptions {
		var n int
		if _, err := fmt.Sscanf(a.ID, "A-%d", &n); err == nil && n > high {
			high = n
		}
	}
	return fmt.Sprintf("A-%03d", high+1)
}

// RecordAssumption adds an open assumption to the register. Origin and confidence are
// what make it readable by somebody who did not write it, so both are required here
// although the schema leaves them out of a record that predates them.
func (r *Runner) RecordAssumption(key, phase, text, origin, confidence, resolves string) (*model.Assumption, error) {
	if text == "" {
		return nil, refuse("an assumption needs --text: the statement being assumed")
	}
	if !model.OneOf(origin, model.AssumptionOrigins) {
		return nil, refuse("--origin is one of %s", strings.Join(model.AssumptionOrigins, ", "))
	}
	if !model.OneOf(confidence, model.AssumptionConfidences) {
		return nil, refuse("--confidence is one of %s", strings.Join(model.AssumptionConfidences, ", "))
	}
	reg, err := r.readRegister(key)
	if err != nil {
		return nil, err
	}
	a := model.Assumption{
		ID: nextAssumptionID(reg), Phase: phase, Assumption: text,
		Origin: origin, Confidence: confidence, Status: "open", Resolves: resolves,
	}
	reg.Assumptions = append(reg.Assumptions, a)
	if err := r.writeRegister(key, reg); err != nil {
		return nil, err
	}
	return &a, nil
}

// DecideAssumption confirms or rejects one. Both are a statement by a person and section
// 8 gives each state its own field, so the person goes into the one that belongs to the
// status and the other stays absent.
func (r *Runner) DecideAssumption(key, id, status, by string) (*model.Assumption, error) {
	if status != "confirmed" && status != "rejected" {
		return nil, refuse("an assumption is confirmed or rejected, not %q", status)
	}
	if by == "" {
		return nil, refuse("--by is required: confirming or rejecting an assumption is a statement by a person")
	}
	reg, err := r.readRegister(key)
	if err != nil {
		return nil, err
	}
	for i := range reg.Assumptions {
		a := &reg.Assumptions[i]
		if a.ID != id {
			continue
		}
		if !a.Open() {
			return nil, refuse("%s is already %s; a decision is not replaced, it is made once", id, a.Status)
		}
		a.Status = status
		if status == "confirmed" {
			a.ConfirmedBy = by
		} else {
			a.RejectedBy = by
		}
		if err := r.writeRegister(key, reg); err != nil {
			return nil, err
		}
		return a, nil
	}
	return nil, refuse("%s has no assumption %s", key, id)
}
