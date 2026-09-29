---
intent: github.com/triplem/xeno#110
phase: 01-requirements
created: "2026-09-29T19:02:49Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 59a69283634bd0c19cd5369128f8d166e998ebf75e9cfea2accc8a9893e9659d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**AC1.** `run` takes the two writers and nothing in the package refers to `os.Stdout` or
`os.Stderr` except `main`.

**AC2.** The staircase, asserted through `run`: 0 where the command did what was asked
and the verdict is not red, 1 on a red verdict and on a refusal with a reason, 2 where
it could not run at all.

**AC3.** A provisional verdict exits 0 from `gate verify`, and a divergence and a red
phase each exit 1. The provisional case carries a comment naming section 6 and the
wrapper that depends on it.

**AC4.** Every command the usage string names resolves in the dispatch table, and every
entry in the table appears in the usage string. A name that is in neither exits 2 with
the usage on standard error.

**AC5.** `init` is reachable as one word and every other command as two, and a single
word that is not `init` exits 2.

**AC6.** A command needing `--intent` without one exits 2 and says which flag. A command
needing a phase with an unresolvable one exits 2 and says so.

**AC7.** A missing positional argument does not become a silent empty string: `gate
approve` without a finding id exits 2 or refuses with a reason, and the test records
which.

**AC8.** `--export` writes the phase environment to standard output and no suggestion
with it. `--no-next` suppresses the suggestion, and exit code 2 suppresses it whether or
not the flag is given.

**AC9.** Each assertion says which stream its output arrived on, so that the refactor
cannot quietly move something from one to the other.

**AC10.** Everything green stays green. `gate verify` matches every verdict and no
behaviour changes: the same words on the same streams with the same codes.

<!-- xeno:section:non-goals -->
## Non goals

The commands' behaviour. What the runner does when asked is the runner's own tests; this
is what the command layer returns and where it writes.

`enforcement check`, which reaches a host. It stays untested and the gap is named rather
than closed.

A coverage number as the goal. The issue names five things and those are the subject.

The output's wording. XENO-0207 recorded that a sentence in output has no test and that
this is the right trade; the tests here assert a stream and a substring that identifies
the case, not a phrasing.

Parallel tests. The writers make them possible and nothing here marks a test parallel,
since the package's tests use a temporary directory each and are fast.

`main`, which stays one line.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes. The staircase is the package's own comment and section 6
is what the provisional case answers to; both already say what this asserts.

No behaviour may change. The same words on the same streams with the same exit codes,
which is what makes a refactor of fifty-five call sites safe to review: every difference
in the diff is a writer.

The writers are supplied, never global. That is the point of the change and a test that
reassigned `os.Stdout` would defeat it.

Each test names the stream it read. A refactor this shape fails by sending something to
the wrong one, and an exit code would not notice.

`run` keeps its signature except for the writers, so its callers stay `main` and the
tests.

One intent, one issue. The commits reference #110.
