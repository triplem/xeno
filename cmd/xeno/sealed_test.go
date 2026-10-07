// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
)

// ---- WP1: what is sealed is never rewritten, asserted over the whole surface

// walked is the phase every command of the walk is aimed at. P0 is the one phase a fixture
// can reach through the commands themselves, because the sequence refuses a later one until
// its predecessor is green, and a phase is sealed by having a verdict rather than by that
// verdict being green. A red one also leaves findings behind, which is what lets `gate
// approve`, `gate override` and `obligation close` act on something instead of refusing for
// want of an id: the four commands that write into `gate.yaml` are the interesting half of
// this walk, since `hashing.PhaseExcluded` keeps that file outside `artifacts_hash` and the
// property worth asserting is that writing there still leaves the hash alone.
const walked = "00-intake"

// walk is one command's place in the surface: what it is given beyond the flags the dispatch
// table itself says it needs, and what it is allowed to change.
type walk struct {
	// args are the command's own arguments, a positional one first as `parse` reads it.
	// A value in braces is filled per fixture, because a finding id and a file path cannot
	// be written down in advance.
	args []string
	// needsGit makes the fixture a repository, for the one command that asks git what a
	// commit range touches. Only that command pays the skip when git is absent.
	needsGit bool
	// moves, where set, is why this command is expected to change the sealed phase's
	// artifacts_hash. It is the point of the whole table: the exceptions are visible in one
	// place, and adding one is a decision somebody wrote down rather than a silent pass.
	moves string
	// verdict is what the command may do to the judged phase's gate.yaml, which is the file
	// hashing.PhaseExcluded keeps outside artifacts_hash.
	verdict verdictWrite
	// stdin, for the one command that reads a hook's JSON instead of taking flags. Without it
	// that command reads the empty standard input a test has, does nothing, and passes while
	// saying nothing about where it writes.
	stdin bool
	// cannotRun, where set, is why the walk cannot get this command as far as running: the
	// exit code is then 2 and the walk says nothing about sealing for it. It is asserted the
	// same way moves is, so a command that becomes runnable loses its entry rather than
	// keeping a note that stopped being true.
	cannotRun string
}

// verdictWrite is what a command is allowed to do to a judged phase's gate.yaml. Section 7
// gives a person three commands that write a decision into a phase that is already sealed,
// and `gate run` a fourth that writes the whole verdict again; what makes all four safe is
// the exclusion rather than a rule about when they may be used, so the walk holds them to
// writing there and to nothing else.
type verdictWrite int

const (
	// leavesTheVerdict is the default and the case for most of the surface: the command has
	// no business in gate.yaml and must leave its bytes alone.
	leavesTheVerdict verdictWrite = iota
	// writesTheVerdict is a person's decision written into it, which is asserted in both
	// directions: a command named here that wrote nothing would leave the walk passing while
	// proving nothing about the exclusion it rests on.
	writesTheVerdict
	// rejudges is the verdict computed and written again. Over a phase the walk has not
	// changed, the bytes are the ones already there apart from run_at, so whether the file
	// changed is a question about which second the test ran in and the walk does not ask it.
	rejudges
)

// aimedAtTheSealedPhase carries one entry per command. The names are not listed here as a
// second copy of the surface: the test reads them out of the dispatch table and fails on an
// entry the table does not have and on a command this map does not cover, which is how a
// command added tomorrow arrives in the walk whether or not anybody remembered it.
var aimedAtTheSealedPhase = map[string]walk{
	// --host because init refuses without one and names the two it knows; the flag is not in
	// the usage string, which is a finding about the usage and not about this walk.
	"init": {args: []string{"--host", "github"}},
	"enforcement check": {cannotRun: "it asks the host what it enforces, which takes a token " +
		"and a network, and a test has neither; it is also the one command that never opens " +
		"an intent directory, so there is nothing in a phase for it to reach"},
	"check commit-message": {args: []string{"--file", "{message}"}},
	"gate verify":          {},
	"intent verify":        {args: []string{"--base", "{base}", "--head", "HEAD"}, needsGit: true},
	// The ledger cost turn appends to lies under .xeno/local/, outside every intent directory,
	// which is why a turn recorded against a judged phase leaves that phase alone.
	"cost turn":     {stdin: true},
	"intent start":  {args: []string{"--for", "9"}},
	"intent status": {args: []string{"--all"}},
	"intent close":  {args: []string{"--reason", "the walk is over"}},
	// --phase by hand, because the dispatch table has needsPhase off here: the record is
	// owed per phase and once more when an intent closes, and without the flag this would
	// write at intent level and never reach the phase the walk is about.
	"learning record": {args: []string{"--phase", walked, "--no-finding"},
		moves: "the record lies in the phase directory and inside artifacts_hash, and " +
			"RecordLearning deliberately does not refuse a sealed phase: a command that " +
			"refused would be refusing the first half of a re-judgement"},
	"scope set": {args: []string{"--file", "{scope}"},
		moves: "the context scope is P0's own artifact, so writing it after P0 is judged " +
			"changes P0; cmdScopeSet prints that caveat where somebody types the command"},
	"assumption confirm": {args: []string{"A-001", "--by", "a.person"}},
	"assumption reject":  {args: []string{"A-001", "--by", "a.person"}},
	"section set": {args: []string{"problem", "--file", "{text}"},
		moves: "the section is written into output.md, which is inside artifacts_hash; " +
			"SectionSet does not refuse a sealed phase because rewriting the sections and " +
			"running phase finish again is how a phase is redone"},
	"phase start": {},
	"phase finish": {args: []string{"--summary", "{text}"}, verdict: rejudges,
		moves: "this is the command that computes artifacts_hash, and the digest it writes " +
			"is inside it; WP1 names it as the only command expected to move one"},
	// --evidence-from because `gate run` attaches before it judges, so the walk gives it the
	// pipeline the fixture's pending declaration is waiting for: what the attach writes lands
	// under evidence/, which DirHash does not descend into, and that is a second way into the
	// phase's directory that has to leave the hash alone.
	"gate run": {args: []string{"--evidence-from", "{pipeline}"}, verdict: rejudges},
	"gate approve": {args: []string{"{finding}", "--by", "a.person", "--reason", "assessed"},
		verdict: writesTheVerdict},
	"gate override": {args: []string{"{finding}", "--by", "a.person", "--reason", "needed now"},
		verdict: writesTheVerdict},
	"assumption record": {args: []string{"--text", "the host reflows to 72",
		"--origin", "repo-convention", "--confidence", "low"}},
	"question record": {args: []string{"--file", "{question}"},
		moves: "the entry is a frontmatter block of output.md, which is inside artifacts_hash"},
	"decision record": {args: []string{"--chosen", "fail fast", "--reason", "the caller retries",
		"--by", "a.person"},
		moves: "the decision is a frontmatter block of output.md, which is inside artifacts_hash"},
	// review answer resolves the review phase itself rather than taking one, so it cannot be
	// aimed at an earlier sealed phase at all. What the walk asserts about it is that an answer
	// given while P0 is sealed does not reach P0; #242's three refusals are what keep it off a
	// sealed P5, which this fixture does not reach.
	"review answer":    {args: []string{"deviations-are-traceable", "--result", "met"}},
	"obligation close": {args: []string{"{overridden}"}, verdict: writesTheVerdict},
	"evidence attach":  {args: []string{"--from", "{pipeline}"}},
	// The route writes the rule files, the patch and the description into --out and nothing
	// anywhere else: it reads the learning records of a sealed phase and leaves that phase
	// alone, which is the property worth asserting for a command that walks the whole trail.
	"learning propose": {args: []string{"--out", "{proposals}"}},
	// --dry-run, because the command's whole purpose is a call to the code host and the walk
	// has neither a token nor a network. What it still proves is the half that matters here:
	// composing a comment from a sealed phase's verdict reads that phase and writes nothing
	// into it.
	"report verdict": {args: []string{"--dry-run"}},
	"evidence declare": {args: []string{"--kind", "scan", "--job", "semgrep"},
		moves: "the declaration is a frontmatter block of output.md, so a declaration after " +
			"the verdict leaves the phase stale rather than rewriting it (#216)"},
}

// seal is a repository whose P0 has been judged, with the values a command needs to be aimed
// at it. Each command gets its own, because a command that wrote something would otherwise be
// deciding what the next one found.
type seal struct {
	root string
	// hash is artifacts_hash as the directory computes it, and recorded is the value the
	// verdict sealed. Both are asserted: the first is what `gate verify` recomputes, the
	// second is what it compares against, and a command could move either one alone.
	hash, recorded string
	subst          map[string]string
	// hook is a session hook's JSON, for the command that reads one on standard input.
	hook string
}

// sealedP0 drives P0 to a verdict through the commands, so that the fixture is a state the
// tool can actually produce rather than files written beside it. The verdict is red — the
// shipped gates want a plugin, a secrets hash and a learning record that this fixture has no
// reason to carry — and red is what leaves findings to decide about.
func sealedP0(t *testing.T) *seal {
	t.Helper()
	root := repo(t)
	at := []string{"--root", root, "--intent", "PROJ-1", "--no-next"}
	aimed := append(at, "--phase", walked)

	// A project configuration, so that `intent start` has a tracker to qualify an issue
	// against and refuses for nothing else.
	writeUnder(t, root, ".xeno/config/project.yaml",
		"tracker:\n  adapter: github\n  project: triplem/xeno\n"+
			"  base_url: https://api.github.com\n")
	// One review rule, so that `review answer` is refused for where the checklist belongs
	// rather than for an empty rule set, which is a refusal about the fixture.
	writeUnder(t, root, ".xeno/plugin/rules/given/builtin/deviations-are-traceable.yaml",
		"id: deviations-are-traceable\nversion: 1\nscope: builtin\nbinding: false\n"+
			"kind: review\napplies_to: [05-review]\nstatement: >\n  A deviation is traceable.\n")

	step(t, append([]string{"phase", "start"}, aimed...))
	for _, section := range []string{"problem", "scope", "context-rationale"} {
		step(t, append([]string{"section", "set", section, "--file",
			inputFile(t, section+".md", "what the "+section+" section says\n")}, aimed...))
	}
	// A pending declaration, so that `evidence attach` has an entry to bind and writes
	// something rather than finding nothing to do.
	step(t, append([]string{"evidence", "declare", "--kind", "test-report", "--job", "unit",
		"--format", "junit"}, aimed...))
	// An assumption, so that confirm and reject have a record to decide on. The register is
	// an intent level file, which is why neither command can reach a phase.
	step(t, append([]string{"assumption", "record", "--text", "the gates are the shipped set",
		"--origin", "repo-convention", "--confidence", "low"}, aimed...))
	step(t, append([]string{"phase", "finish", "--summary",
		inputFile(t, "summary.md", "a summary of the intake\n")}, aimed...))

	ids := findingIDs(t, root)
	if len(ids) < 2 {
		t.Fatalf("the fixture's verdict carries %d findings, so there is nothing to decide about", len(ids))
	}
	// One finding carries an override already, because `obligation close` acts on an
	// obligation and an override is what owes one.
	overridden := ids[len(ids)-1]
	step(t, append([]string{"gate", "override", overridden, "--by", "a.person",
		"--reason", "the merge cannot wait for it"}, aimed...))

	s := &seal{root: root, subst: map[string]string{
		"{finding}":    ids[0],
		"{overridden}": overridden,
		"{text}":       inputFile(t, "content.md", "a line of content\n"),
		"{message}":    inputFile(t, "message.txt", "fix(gates): a subject\n"),
		"{question}": inputFile(t, "question.yaml", "text: which error behaviour?\noptions:\n"+
			"  - text: fail fast\n    consequence: the caller retries\n    recommended: true\n"+
			"    reason: a retry the caller can see is one it can decide about\n"+
			"  - text: retry internally\n    consequence: the caller never sees it\n"+
			"  - text: something else\n    free: true\n"),
		"{scope}":    inputFile(t, "scope.yaml", "include:\n  - src/**\n"),
		"{pipeline}": pipelineDir(t),
		// A destination outside the root entirely, which is the point of the command: it has
		// no default and refuses a path inside .xeno, because the trail it reads is the one
		// place a learning may not take effect.
		"{proposals}": t.TempDir(),
	}}
	// A hook's JSON and the transcript it points at, the two inputs `cost turn` takes. The
	// usage figures are what make the turn worth recording: a transcript summing to nothing
	// is dropped, which would be the same silence as an empty standard input.
	transcript := inputFile(t, "transcript.jsonl",
		"{\"message\":{\"usage\":{\"input_tokens\":1200,\"output_tokens\":340,"+
			"\"cache_read_input_tokens\":900}}}\n")
	s.hook = "{\"transcript_path\":\"" + transcript + "\",\"session_id\":\"s-0001\"}"
	s.hash, s.recorded = phaseHash(t, root), recordedHash(t, root)
	if s.hash != s.recorded {
		t.Fatalf("the fixture is stale before the walk begins: directory %s, verdict %s", s.hash, s.recorded)
	}
	return s
}

// TestEveryCommandLeavesASealedPhaseWhereItWas is WP1's done-when in the enumerated form it
// asks for: every command the dispatch table holds, each against a phase that has been judged.
// The per command assertions elsewhere catch the command somebody was thinking about; this
// catches the one nobody was, which is the one that gets added next (#242 is the issue that
// took).
func TestEveryCommandLeavesASealedPhaseWhereItWas(t *testing.T) {
	for name := range commands {
		if _, ok := aimedAtTheSealedPhase[name]; !ok {
			t.Errorf("the dispatch table holds %q and the walk does not aim it anywhere; "+
				"give it arguments, and a reason if it is allowed to move the hash", name)
		}
	}
	for name := range aimedAtTheSealedPhase {
		if _, ok := commands[name]; !ok {
			t.Errorf("the walk aims %q and the dispatch table has no such command", name)
		}
	}

	for name, c := range commands {
		w, ok := aimedAtTheSealedPhase[name]
		if !ok {
			continue
		}
		t.Run(name, func(t *testing.T) {
			s := sealedP0(t)
			args := append(strings.Fields(name), s.fill(t, w)...)
			args = append(args, "--root", s.root, "--no-next")
			if c.needsKey {
				args = append(args, "--intent", "PROJ-1")
			}
			if c.needsPhase {
				args = append(args, "--phase", walked)
			}

			if w.stdin {
				s.feed(t)
			}
			verdictBefore := verdictBytes(t, s.root)
			code, out, errw := invoke(t, args...)
			// 2 is "could not run at all", which here means the walk handed the command
			// arguments it could not use. The command did nothing, so the hash below would
			// be unchanged for a reason that has nothing to do with sealing.
			switch {
			case code == 2 && w.cannotRun == "":
				t.Fatalf("the walk could not run %q, so it proves nothing about sealing:\n%s%s",
					name, out, errw)
			case code != 2 && w.cannotRun != "":
				t.Fatalf("%s ran after all, exit %d, and the walk says it cannot: %s",
					name, code, w.cannotRun)
			case code == 2:
				return
			}

			after, recorded := phaseHash(t, s.root), recordedHash(t, s.root)
			moved := after != s.hash
			switch {
			case moved && w.moves == "":
				t.Errorf("%s moved %s's artifacts_hash from %s to %s.\n"+
					"Either it must not write inside a judged phase, or the walk needs an "+
					"entry saying why it may:\n  exit %d\n%s%s",
					name, walked, s.hash, after, code, out, errw)
			case !moved && w.moves != "":
				t.Errorf("%s is named in the walk as moving %s and it did not: %s.\n"+
					"An exception nothing needs is one the next reader will trust:\n  exit %d\n%s%s",
					name, walked, w.moves, code, out, errw)
			}
			// The hash the verdict recorded is a second thing a command could move, and the
			// one `gate verify` compares the directory against. Only a command that computes
			// the verdict again may write it, and over an unchanged phase even that one
			// writes back the value already there.
			if recorded != s.recorded && w.verdict != rejudges {
				t.Errorf("%s rewrote the artifacts_hash %s's verdict recorded, from %s to %s",
					name, walked, s.recorded, recorded)
			}
			wrote := verdictBytes(t, s.root) != verdictBefore
			switch {
			case wrote && w.verdict == leavesTheVerdict:
				t.Errorf("%s wrote %s's gate.yaml and the walk does not say it may; "+
					"say what it writes there, or stop writing into a judged phase:\n%s%s",
					name, walked, out, errw)
			case !wrote && w.verdict == writesTheVerdict:
				t.Errorf("%s is named in the walk as writing a decision into %s's gate.yaml "+
					"and it wrote nothing, so the walk no longer holds that file outside "+
					"artifacts_hash:\n  exit %d\n%s%s", name, walked, code, out, errw)
			}
		})
	}
}

// feed puts the fixture's hook JSON on standard input for the rest of the subtest. The
// commands take their writers as arguments and `run` can be driven in process because of
// it, which is #110's point; standard input is the one channel still read from the process,
// so a test that wants to drive it has to lend the process its own.
func (s *seal) feed(t *testing.T) {
	t.Helper()
	p := filepath.Join(t.TempDir(), "hook.json")
	if err := os.WriteFile(p, []byte(s.hook), 0o644); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}
	saved := os.Stdin
	os.Stdin = f
	t.Cleanup(func() { os.Stdin = saved; f.Close() })
}

// fill substitutes the per fixture values into one command's arguments and fails on a
// placeholder nothing fills, which would otherwise reach the command as a literal.
func (s *seal) fill(t *testing.T, w walk) []string {
	t.Helper()
	if w.needsGit {
		// The fixture becomes a repository here rather than in sealedP0, so that the skip
		// when git is absent costs this command alone.
		s.subst["{base}"] = gitRepo(t, s.root)
	}
	args := make([]string, 0, len(w.args))
	for _, a := range w.args {
		if strings.HasPrefix(a, "{") {
			v, ok := s.subst[a]
			if !ok {
				t.Fatalf("nothing fills the placeholder %s", a)
			}
			a = v
		}
		args = append(args, a)
	}
	return args
}

// phaseHash is Appendix B's artifacts_hash over the walked phase as it lies on disk, which is
// what `gate verify` recomputes.
func phaseHash(t *testing.T, root string) string {
	t.Helper()
	h, err := hashing.DirHash(root, model.PhaseDir("PROJ-1", walked), hashing.PhaseExcluded)
	if err != nil {
		t.Fatal(err)
	}
	return h
}

// recordedHash is the value the verdict sealed, read off gate.yaml.
func recordedHash(t *testing.T, root string) string {
	t.Helper()
	var g model.Gate
	if err := fm.ReadYAML(filepath.Join(root, model.PhaseDir("PROJ-1", walked), "gate.yaml"), &g); err != nil {
		t.Fatal(err)
	}
	return g.ArtifactsHash
}

// verdictBytes is the walked phase's gate.yaml as it lies on disk. Read as bytes rather
// than parsed, because what is being compared is whether anything was written at all.
func verdictBytes(t *testing.T, root string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, model.PhaseDir("PROJ-1", walked), "gate.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

// findingIDs lists the findings of the walked phase's verdict in the order it carries them,
// so that the walk takes its ids from the fixture rather than naming one that a change to a
// gate's cause would silently retire.
func findingIDs(t *testing.T, root string) []string {
	t.Helper()
	var g model.Gate
	if err := fm.ReadYAML(filepath.Join(root, model.PhaseDir("PROJ-1", walked), "gate.yaml"), &g); err != nil {
		t.Fatal(err)
	}
	var ids []string
	for _, c := range g.Checks {
		for _, f := range c.Findings {
			ids = append(ids, f.ID)
		}
	}
	return ids
}

// pipelineDir is a directory standing in for the artifact store, with a manifest naming the
// item the fixture declared as pending.
func pipelineDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	write := func(name, body string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("junit.xml", "<testsuite failures=\"0\"/>\n")
	write("manifest.yaml", "- kind: test-report\n  job: unit\n  result: pass\n  file: junit.xml\n"+
		"  pipeline: \"4711\"\n  commit: abc123\n")
	return dir
}

// writeUnder writes one file inside the fixture, making the directories it needs.
func writeUnder(t *testing.T, root, rel, body string) {
	t.Helper()
	p := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

// step runs one command of the fixture's own setup. A red verdict is 1 and expected here, so
// only 2 is fatal: that is the command not having run, which would leave the fixture in a
// state nothing else in the test is about.
func step(t *testing.T, args []string) {
	t.Helper()
	if code, out, errw := invoke(t, args...); code == 2 {
		t.Fatalf("the fixture could not run %q: %s%s", strings.Join(args, " "), out, errw)
	}
}
