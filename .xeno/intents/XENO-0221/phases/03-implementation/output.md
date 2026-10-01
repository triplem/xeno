---
intent: github.com/triplem/xeno#167
phase: 03-implementation
created: "2026-10-01T16:56:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3791e2004efb4cf75f1d8afcbf47632a72b6683790965ecd705623e7e853a54b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One new package of 283 lines plus 330 of test, one type and one constant in the model, one field
on the gate context, one producer on the runner, one documented block in the scaffold, two
register rows, and two tests through the verdict.

**`internal/external`.** `Run(root, phase, artifactsHash, qualifiedIntent, phaseDir, declared)`
returns a check per declared gate that applies to the phase, in the order declared. `one` is the
life of a declaration: validated, hashed, run, read. `declaration` reports what is wrong with the
declaration itself — no id, no path, no hash, a hash that is not sixty-four lowercase hex
characters, no phase, a phase that is not one of the six — each naming
`.xeno/config/project.yaml`, so a configuration mistake reads as one. `fileHash` is read from
disk on every call, because Appendix A's "checked before every run" is a frequency.

**The refusal.** A hash that does not match produces a failed check naming the declared hash and
the one the file has, and nothing is started before the comparison. A test modifies the command
after the declaration was written and asserts both the finding and that the command left no
trace of having run.

**The contract, A76.** `Input` on stdin with `intent`, `phase`, `phase_dir`, `artifacts_hash`
and `root`; `Output` on stdout with `findings`, each carrying a required `cause` and an optional
`file` and `next`. Unknown fields are ignored so a tool may answer a superset. `findings` rejects
an answer that is not the agreed JSON, naming what came back, and fills an absent `file` with the
gate's own path so that a finding always names something.

**The exit status decides, A75 bounds the wait.** Zero is a pass carrying whatever findings came
with it; non-zero is a fail, and a non-zero exit with nothing to say gets a synthesised finding
naming the status, because `Status` refuses a fail without one. Sixty seconds, in a context
deadline, with `WaitDelay` a second behind it.

**`clip` bounds what reaches a verdict**: whitespace collapsed, four hundred characters, then an
ellipsis. A `gate.yaml` is read by people and by a diff, and the length of foreign output is not
this project's to control.

**`internal/model`.** `ExternalGate` with the four fields of Appendix A, `Project.ExternalGates`,
and `ProjectFile` for the path a declaration finding names. The comment on the field says what
makes it different from every other field in that file: the others change what the runner reads,
this one changes what it runs.

**`internal/gates`.** `Ctx` gains `External func(phase string) []model.Check`, which `Run` calls
after the table and routes through the same `carryForward` as every shipped gate. That is where
an external finding gets its id and where it is refused a decision, so the verdict has one
assembly point. The package starts no process and its comment still holds.

**`internal/runner`.** `externalGates` reads `project.yaml` per run and returns nil where nothing
is declared, so the common case adds no closure and no check. `compute` passes it into the
context, which puts external checks in front of `Invariants` and `Status` with everything else —
the second path the comment there has been describing since #66.

**`internal/scaffold/files/project.yaml`.** The block as a comment, with what it declares, that
the hash is checked before every run, that the command is started directly rather than through a
shell, that it has sixty seconds, that its findings are marked, that a decision on one does not
survive, and that there is no sandbox. Nothing is declared, because an example declaration is a
command somebody has to delete.

**Tests, `internal/external`.** Nothing declared runs nothing; a declared gate passes; exit zero
with findings is a pass carrying them; a non-zero exit is a fail; a non-zero exit with nothing to
say still produces a finding; a modified gate refuses to run and leaves no trace; a missing gate
and an unexecutable one each name what happened; three malformed answers are each a fail naming
the contract; a hung gate is killed and named; six wrong declarations each name the configuration
file; a declaration for no phase runs nowhere; a gate runs only at the phases it names; the
command is given the five fields of the contract, read back out of the process; and a two
thousand character cause is bounded before it reaches the verdict.

**Tests, `internal/runner`.** A declared gate's check reaches `gate.yaml` with
`provenance: external`, its finding carries an id, the phase is red, a release is taken on it, and
the next `gate run` produces the same finding with the same id and no decision. And a project that
declares none produces no external check at all.

<!-- xeno:section:deviations -->
## Deviations from the design

**Killing the command does not bound the wait, and the first version did not bound it at all.**
The timeout was a context deadline, which kills the process, and the test for it took thirty
seconds — the full `sleep` the fake gate ran. `Output` waits for standard output to close, and the
grandchild that inherited the pipe keeps it open after its parent is killed. So a gate that starts
something long-running could hold a verdict open for as long as that thing lived, with a timeout
that looked implemented and was not. `WaitDelay` closes the pipes a second after the kill; the
test now takes 1.2 seconds, and A75 records the whole of it rather than only the limit.

**The limit is a package variable so that a test can shorten it.** `Timeout` is the constant and
`timeoutForTest` is what the code reads. That is a seam for a test inside production code, which
this project has not needed before, and the alternative was a test that waits a minute to assert a
mechanism that behaves identically at any limit. The variable is unexported, unset anywhere but
the test, and the comment says so.

**The first version of the malformed-answer test built a gate it threw away.** Two declarations
per case, one unused, left over from writing the answer into a `printf` before switching to a
heredoc. It passed and asserted the right thing by accident of the second declaration. Removed.
That is twice in three pieces that a test of mine was green while carrying something that was not
doing anything — the pattern is a variable or a value created and then not used, and the learning
in #160 said to read a new test back rather than trust the green.

**`run1` in the tests fixes the phase, so the applicability cases needed their own calls.** A
small thing recorded because it shaped the tests: two of them call `Run` directly rather than
through the helper, which looks inconsistent and is the only way to assert what happens at a phase
a gate does not name.

**The runner test could not use the fixture the way I first wrote it.** `f.run` starts the phase,
writes the artifact and finishes it, and my version called `Start` and `SectionSet` first, so the
phase was already running when `f.run` started it. The second judgement also has to go through
`GateRun` rather than `f.run`, because a finished phase cannot be started again. Both are the
fixture working as intended; the deviation is that I wrote against what I assumed it did.

**No end-to-end run against this repository.** Every other piece in WP4 was exercised against the
real tree on a throwaway copy. This one is not: declaring an external gate here would mean writing
a tool for this project to run on itself, which is content rather than verification, and the
verification is in the runner test instead. It is the first piece of the five with no copy-of-the-
tree demonstration, and the gaps say what that costs.
