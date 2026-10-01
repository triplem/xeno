// SPDX-License-Identifier: Apache-2.0

// Package external runs the external gates of section 14: a command a project declares in
// project.yaml with its path and hash, receiving and returning JSON, with the exit status
// deciding, running only when the hash matches.
//
// It is a package of its own because internal/gates reads and this runs. That package holds
// every rule an external check has to satisfy — the provenance mark, the id derivation, the
// refusal to carry a decision on foreign wording — and it holds none of the code that starts a
// process. The import edge runs from the runner to both.
//
// What makes foreign code in the gate path acceptable is not containment: there is no sandbox,
// and the command runs with whatever the pipeline gives it. It is the mark. Section 14: without
// external gates the chain of trust is closed, and an external gate brings foreign code with
// repository access into the pipeline, so the trail has to show which statement came from Xeno
// and which did not.
package external

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/triplem/xeno/internal/model"
)

// Timeout bounds a foreign command. Nothing in the documents bounds one, and an unbounded
// subprocess in the gate path is a property nobody chose: the first external gate somebody
// writes will have a bug in it, and a gate that hangs is not deterministic in the only sense
// this project can claim. A project that needs longer argues with a number that is written
// down rather than with a discovery (A75).
const Timeout = 60 * time.Second

// timeoutForTest is the limit actually used, so that the test for the mechanism does not wait a
// minute. Nothing outside this package sets it and the shipped value is Timeout.
var timeoutForTest = Timeout

// Provenance is what every check from here carries, per section 5's vocabulary. It is the same
// constant internal/gates checks for; the string is the contract between them, as it is in
// gate.yaml.
const Provenance = "external"

// hashShape is sixty-four lowercase hex characters, which is what Appendix B fixes a sha256 to
// and what a declaration has to carry.
const hashLength = 64

// Input is what the command is given on stdin. It is what a gate needs in order to find the
// artifacts and to name them back in a finding, and nothing else: the repository is already
// there, with the pipeline's own access to it (A76).
type Input struct {
	Intent        string `json:"intent"`
	Phase         string `json:"phase"`
	PhaseDir      string `json:"phase_dir"`
	ArtifactsHash string `json:"artifacts_hash"`
	Root          string `json:"root"`
}

// Output is what the command may answer. Anything else in the object is ignored, so that a tool
// may answer a superset of this; what is read is the findings, and of a finding only the three
// fields a verdict carries.
type Output struct {
	Findings []Finding `json:"findings"`
}

// Finding is one statement from foreign code. Cause is required, because a finding with nothing
// to say cannot be acted on and the verdict refuses a fail without one; file and next are
// optional, and both are text rather than anything the runner resolves.
type Finding struct {
	File  string `json:"file"`
	Cause string `json:"cause"`
	Next  string `json:"next"`
}

// Run evaluates every declared gate that applies to the phase, in the order it was declared.
//
// One check per declaration, each carrying the declaration's id and this provenance. The caller
// routes them into the verdict with every other check, which is where they get their finding
// ids and where the rule that they carry no decision is enforced.
func Run(root, phase, artifactsHash, qualifiedIntent, phaseDir string, declared []model.ExternalGate) []model.Check {
	var out []model.Check
	for _, d := range declared {
		if !applies(d, phase) {
			continue
		}
		in := Input{Intent: qualifiedIntent, Phase: phase, PhaseDir: phaseDir,
			ArtifactsHash: artifactsHash, Root: root}
		out = append(out, one(root, d, in))
	}
	return out
}

// applies reads the declaration's phases. A declaration naming no phase runs nowhere rather than
// everywhere: section 14 has a project name the phases, and the generous reading would run
// foreign code at six phases because a field was forgotten.
func applies(d model.ExternalGate, phase string) bool {
	for _, p := range d.Phases {
		if p == phase {
			return true
		}
	}
	return false
}

// one is the whole life of a declaration: validated, hashed, run, read.
func one(root string, d model.ExternalGate, in Input) model.Check {
	if fs := declaration(d); fs != nil {
		return check(d.ID, "fail", fs)
	}
	path := filepath.Join(root, d.Path)
	sum, err := fileHash(path)
	if err != nil {
		return check(d.ID, "fail", []model.Finding{{File: d.Path,
			Cause: "external gate " + d.ID + " cannot be read: " + err.Error(),
			Next:  "declare a path to a file in this repository, relative to its root"}})
	}
	// Before every run, per Appendix A, and before anything is started: a modified gate refuses
	// to run rather than running unnoticed. Reported as a failed check and not as a skipped one,
	// because a modified tool must not be able to make a verdict quieter.
	if sum != d.SHA256 {
		return check(d.ID, "fail", []model.Finding{{File: d.Path,
			Cause: "external gate " + d.ID + " is declared with sha256 " + short(d.SHA256) +
				" and the file hashes to " + short(sum),
			Next: "re-declare the hash if the change was intended; the gate runs only when it matches"}})
	}
	return run(path, d, in)
}

// declaration reports what is wrong with the declaration itself, so that a project's
// configuration mistake reads as one rather than as an attempt to run something.
func declaration(d model.ExternalGate) []model.Finding {
	var fs []model.Finding
	add := func(cause, next string) {
		fs = append(fs, model.Finding{File: model.ProjectFile, Cause: cause, Next: next})
	}
	name := d.ID
	if name == "" {
		name = "an external gate"
		add("an external gate is declared with no id",
			"give it an id; it is what names its check in gate.yaml")
	}
	if d.Path == "" {
		add("external gate "+name+" is declared with no path",
			"declare the path to the command, relative to the repository root")
	}
	switch {
	case d.SHA256 == "":
		add("external gate "+name+" is declared with no sha256",
			"declare the hash of the command; the gate runs only when it matches")
	case len(d.SHA256) != hashLength || strings.ToLower(d.SHA256) != d.SHA256 || !isHex(d.SHA256):
		add("external gate "+name+" is declared with sha256 "+short(d.SHA256)+", which is not a sha256",
			"write sixty four lowercase hex characters, as Appendix B defines it")
	}
	if len(d.Phases) == 0 {
		add("external gate "+name+" is declared for no phase",
			"list the phases it applies to; a gate declared for none runs nowhere")
	}
	for _, p := range d.Phases {
		if model.PhaseIndex(p) < 0 {
			add("external gate "+name+" is declared for phase "+quoted(p)+", which is not a phase",
				"name one of the six phases: "+strings.Join(model.Phases, ", "))
		}
	}
	return fs
}

// run starts the command, gives it the input on stdin, and reads the answer.
func run(path string, d model.ExternalGate, in Input) model.Check {
	body, err := json.Marshal(in)
	if err != nil {
		return check(d.ID, "fail", []model.Finding{{File: d.Path,
			Cause: "external gate " + d.ID + ": its input could not be written: " + err.Error(),
			Next:  "report this; the runner writes the input and the gate did not run"}})
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeoutForTest)
	defer cancel()

	cmd := exec.CommandContext(ctx, path)
	cmd.Stdin = strings.NewReader(string(body))
	// Killing the command is not enough to return: a grandchild that inherited the pipe keeps
	// stdout open, and Output waits for it. WaitDelay closes the pipes shortly after the kill, so
	// a gate that starts something long-running cannot hold the verdict open behind it.
	cmd.WaitDelay = time.Second
	out, runErr := cmd.Output()

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return check(d.ID, "fail", []model.Finding{{File: d.Path,
			Cause: fmt.Sprintf("external gate %s did not finish within %s", d.ID, timeoutForTest),
			Next:  "make the gate answer inside the limit, or take it out of the declaration"}})
	}

	var exit *exec.ExitError
	switch {
	case runErr == nil:
	case errors.As(runErr, &exit):
		// A non-zero exit is the verdict section 14 asks for, so the answer is still read: a
		// tool that fails and says why is the ordinary case.
	default:
		return check(d.ID, "fail", []model.Finding{{File: d.Path,
			Cause: "external gate " + d.ID + " could not be run: " + runErr.Error(),
			Next:  "declare a path to an executable file; the runner starts it directly, not through a shell"}})
	}

	fs, parseErr := findings(d, out)
	if parseErr != nil {
		return check(d.ID, "fail", []model.Finding{{File: d.Path,
			Cause: "external gate " + d.ID + " answered something other than the agreed JSON: " + parseErr.Error(),
			Next:  `answer {"findings":[{"file":"","cause":"","next":""}]} on standard output; cause is required`}})
	}
	if runErr == nil {
		// Exit zero with findings is a pass carrying them, which is how a tool reports what it
		// does not consider fatal. The exit status decides, not the findings.
		return check(d.ID, "pass", fs)
	}
	if len(fs) == 0 {
		// Status refuses a fail with no finding, so the absence is filled here rather than
		// allowed to make a verdict unrepresentable.
		fs = []model.Finding{{File: d.Path,
			Cause: fmt.Sprintf("external gate %s exited %d and reported no finding", d.ID, exit.ExitCode()),
			Next:  "have the gate say what is wrong, or exit zero where nothing is"}}
	}
	return check(d.ID, "fail", fs)
}

// findings reads the answer. Everything from foreign code is text: trimmed to a length a verdict
// can carry, with the file recorded as the command wrote it and never resolved against the
// repository.
func findings(d model.ExternalGate, out []byte) ([]model.Finding, error) {
	trimmed := strings.TrimSpace(string(out))
	if trimmed == "" {
		return nil, nil // nothing to say, which an exit status still decides
	}
	var answer Output
	if err := json.Unmarshal([]byte(trimmed), &answer); err != nil {
		return nil, fmt.Errorf("%v (%s)", err, short(trimmed))
	}
	var fs []model.Finding
	for i, f := range answer.Findings {
		if strings.TrimSpace(f.Cause) == "" {
			return nil, fmt.Errorf("finding %d carries no cause", i+1)
		}
		file := strings.TrimSpace(f.File)
		if file == "" {
			file = d.Path
		}
		fs = append(fs, model.Finding{File: clip(file), Cause: clip(f.Cause), Next: clip(f.Next)})
	}
	return fs, nil
}

func check(id, result string, fs []model.Finding) model.Check {
	return model.Check{Gate: id, Result: result, Provenance: Provenance, Findings: fs}
}

func fileHash(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:]), nil
}

// clip bounds what reaches a verdict. A gate.yaml is read by people and by a diff, and a foreign
// command's output is not something this project controls the length of.
func clip(s string) string {
	const max = 400
	s = strings.Join(strings.Fields(s), " ")
	if len(s) <= max {
		return s
	}
	return s[:max] + "…"
}

func short(s string) string {
	if len(s) > 12 {
		return s[:12] + "…"
	}
	return s
}

func isHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

func quoted(s string) string { return `"` + s + `"` }
