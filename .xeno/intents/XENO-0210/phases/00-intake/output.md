---
intent: github.com/triplem/xeno#110
phase: 00-intake
created: "2026-09-29T19:01:51Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a39abc00d216eef899527c10ff906983ef0a1c506836483669a15fb520da73a5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

`cmd/xeno` states a contract in its package comment:

> Exit codes follow one staircase throughout: 0 where the command did what was asked and
the > verdict is not red, 1 on a red verdict or a refusal with a reason, 2 where it
could not run at > all.

Nothing asserts it. Of the thirty functions in the package, twenty-eight are at zero
coverage; the two that are not are `tail` and `sprintRow`, which XENO-0207 extracted so
that a table's arithmetic could be tested. Cross package coverage is 72.0% and this
package is what is left of the gap that is not a network path.

The staircase is not glue. `report` distinguishes a `runner.Refusal` from an ordinary
error to choose between 1 and 2, and returns 1 again on a red verdict. `cmdGateVerify`
returns 1 on a divergence or a red phase and 0 on a provisional one, which is the
distinction section 6 rests the evidence arrangement on: the wrapper reports provisional
rather than failing, because a P4 that failed verification for waiting on a pipeline is
what section 6 names as the thing that would make people stop taking verification
seriously. Nothing checks that a provisional verdict exits 0.

Three more pieces with nothing behind them. The dispatch, where `needsKey` and
`needsPhase` are fields precisely because they used to be switches listing their
exceptions, and nothing checks that the table and the usage string agree. The positional
argument, which `parse` takes off the front because Go's flag package stops at the first
non-flag, and which is silently empty when absent. `--no-next` and `--export`, which are
exclusive because a shell eval would otherwise swallow a sentence meant for a person.

`run([]string) int` is already shaped for a test: it takes its arguments and returns the
code. What stops one is where it writes. Fifty-five calls go to the package level
`os.Stdout` and `os.Stderr`, so nothing can read what a command printed without either
replacing those globally or running the built binary.

<!-- xeno:section:scope -->
## Scope

The writers become explicit. `run` takes them, `opts` carries them, and `report` and
`suggest` become methods on `opts` so that every command writes through something a test
can supply. `main` passes `os.Stdout` and `os.Stderr` and is the only place they appear.

Tests over `run`, in process, asserting the staircase and the four pieces the issue
names: the exit codes for a refusal, an error, a red verdict and a green one; that a
provisional verdict exits 0; that the dispatch resolves every command the usage names
and rejects what it does not; that a missing positional argument and a missing
`--intent` exit 2 with a reason; and that `--export` writes only the environment.

Not the commands' behaviour, which the runner's own tests cover. This is about what the
command layer returns and where it writes, not about what the runner does when asked.

Not a network path. `enforcement check` reaches a host and stays untested.

Not coverage as a number. The issue names five things and those are the subject; a
percentage is what follows rather than what is aimed at.

<!-- xeno:section:context-rationale -->
## Why this context

**The writers are made explicit rather than swapped globally.** A test could reassign
`os.Stdout` for the length of a case, which would need no production change at all, and
it would make every test in the package share one mutable global and forbid running them
in parallel. The issue weighed the same two options and chose this one: what writes goes
behind something a caller supplies, and `main` is the only place the real files appear.

**`report` and `suggest` become methods on `opts`, which is where the writers already
have to live.** Every command function already takes `*opts`, so once it carries the
writers nothing else has to be threaded. `run` takes them as parameters because it
writes the usage before an `opts` exists.

**The subject is the staircase and the four pieces, not the coverage.** A percentage
would rise fastest by testing whichever function is longest, and the issue names what
matters: the three exit codes, the provisional case the CI wrapper depends on, the
dispatch table against the usage string, the positional argument, and the two flags that
must not both print.

**The provisional case is the one worth the whole intent.** Section 6 says the wrapper
reports provisional rather than failing, and gives the reason: a P4 that failed
verification for waiting on a pipeline would teach everybody to stop reading
verification. `cmdGateVerify` returns 0 there, which is one `if` away from being wrong,
and every push out of a P4 depends on it.

**Testing `run` rather than the built binary.** `run` is already shaped for it, a test
in process can read what a command wrote, and a test of the binary would depend on a
build step and on the machine's shell. The cost is the refactor, which is fifty-five
call sites of mechanical change and is the reason this was not done when the package was
written.

**The refactor's own risk, named before it is taken.** Fifty-five edits can send
something to the wrong stream, and a test of the exit code would not notice. So the
tests assert which stream each thing arrives on: a usage message and a refusal on
standard error, a verdict and a listing on standard output.
