---
intent: github.com/triplem/xeno#167
phase: 01-requirements
created: "2026-10-01T16:47:15Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 227b859d9a650b2b25f777b7fd4f84096c17990de217a4b8efff2a1e08bb8a81
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**A project that declares nothing is unaffected.** No declaration, no check, no subprocess, and
the verdict carries the same fourteen rows it carried before. That is every project today and the
default Appendix A names: "none run".

**A declared gate whose hash matches runs, and its verdict is the exit status.** Exit zero is a
pass, anything else is a fail, and the check carries the gate's `id` and
`provenance: external`.

**Its findings appear in `gate.yaml` with ids derived the same way as every other finding.** The
file, the cause and the next step come from the command's answer; the id is computed by the
runner, through the same `carryForward` every other check goes through, so an external finding is
indistinguishable from a Xeno one in how it is identified and distinguishable in where it came
from.

**A declared gate whose hash does not match does not run and is red.** The finding names the
declared hash and the one the file has. Nothing is executed before the comparison, which is what
"checked before every run" means and what WP4's done-when asks for in its own words: a modified
external gate refuses to run rather than running unnoticed.

**A declaration that cannot be run is red, naming what happened.** A path that does not exist, a
file that is not executable, a command that cannot be started: each is its own cause, and none is
a panic or a refused run.

**A command that answers something other than the agreed JSON is red.** Not valid JSON, the wrong
shape, a finding with no cause: the check fails and the finding says what came back, truncated to
something a verdict can carry. A foreign tool's broken answer must not be able to produce a
green.

**A non-zero exit with no findings still produces a finding.** `Status` refuses a check that fails
without one, so the runner synthesises one naming the exit status. A project's tool cannot make
the verdict unrepresentable.

**A command that does not finish in time is red.** The limit is stated in the register and in the
project configuration's own comment, the process is killed, and the finding names the limit.

**A decision on an external finding does not survive the next run.** `gate approve` against an
external finding is accepted, the release is recorded, and the next `gate run` produces the same
finding without the decision, because section 4 says a release on foreign wording is taken again
rather than kept. This is the first time that rule has had anything to act on.

**A gate declared for other phases does not run at this one.** The `phases` list decides, and a
declaration naming a phase that does not exist is red rather than silently inert.

**`G-Complete` accepts an external check like any other**, which it already does, and the status
of a phase with a failing external gate is red for the same reason any other fail is.

**Nothing else moves.** No rule set changes, so no `rules_hash` changes. No sealed artifact is
rewritten. `./xeno gate verify` stays at exit 0 over the trail, and the whole suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No sandbox, no environment scrubbing, no working directory games.** The command runs with what
the pipeline gives it. Section 14 says the marking is the point, and a half-measure would read as
containment while providing none.

**No example declaration anywhere.** `xeno init` writes none, the scaffold's `project.yaml` carries
the block as a comment and nothing else. An example declaration is a command somebody has to
delete before their first run.

**No discovery, no catalogue, no installer.** A project declares a path and a hash.

**No opinion on what an external gate should check.** The one in the tests exists to be run.

**No second way to produce an external finding.** The rule engine does not gain a project-defined
predicate type, which section 14 puts outside v1 and for which this is the alternative.

**No change to how a decision is recorded.** `gate approve` and `gate override` work on an external
finding exactly as they do on any other; what differs is that the next run does not carry it
forward, and that is already implemented.

**No retry.** A command that fails is a fail. A gate that needs retrying is a gate reading
something that is not ready, which is what the evidence mechanism is for.

**No parallelism.** Declared gates run in the order they are declared. Two of them are not worth a
goroutine and the ordering is one less thing to explain in a verdict.

<!-- xeno:section:constraints -->
## Constraints

**The hash is checked before every run**, which Appendix A says in those words. Not cached, not
checked at declaration time, not trusted from a previous run: read the file, hash it, compare, and
only then consider running.

**Execution does not belong in `internal/gates`.** That package reads. Its comment was rewritten
once already, in #162, for one `git log` over a local clone, and foreign code with repository
access is a different kind of exception. The gates package may hold the rules an external check
has to satisfy — it already does — and not the code that starts the process.

**Every external check meets the same verdict path as every other.** `carryForward` for the ids,
`Invariants` for the rules, `Status` for the result. A second path into the verdict that skipped
any of them would make the invariant a comment.

**The contract is written down where a project will meet it**: in the register as an assumption
and in the configuration's own comments, because a project's tool has to agree with it and
nothing else specifies it.

**A foreign answer is data, not instruction.** Whatever the command returns is a finding's text
and nothing else: no path is resolved against the repository, no field is executed, and what is
printed into a verdict is bounded in length.

**Deterministic in the only sense available.** The same tree and the same command produce the same
check, and what the command itself does is its own business. That is the honest limit of a gate
whose body is foreign.

**One dependency, 88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0.**

**One intent, one branch, one issue.** `167-a-projects-own-tool`, #167, labelled wp4, off a `main`
that carries the other four pieces.
