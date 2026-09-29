---
intent: github.com/triplem/xeno#110
phase: 05-review
created: "2026-09-29T19:14:14Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cf4b09894211fd0a1b0d591484d3c32beda49ef4a917dfff4eeb6b0fab824a93
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed.** The staircase is the package's own comment and the
provisional case answers to section 6; both already said what this asserts.

**No behaviour changed.** No word, no format string and no exit code is in the diff;
every hunk is a writer or a signature. That is what makes fifty-five mechanical edits
reviewable, and it is the property AC10 states and `gate verify` over 139 verdicts
supports.

**Only `main` names the real files**, which is the point of the change: everything else
writes through what it was given.

**The staircase the package promises is asserted**, all three steps, and each assertion
says which stream its output arrived on — which is the criterion the refactor needed and
the issue did not ask for.

**The case section 6 depends on is asserted as a branch**, with seven results including
provisional beside red, and the whole command is covered separately over this
repository. The deviation says why it is not one test instead of two, and the gaps say
what that leaves uncovered.

**No test skips.** The first attempt at the provisional case did, which reports success
while asserting nothing, and that is what sent the design back.

**The dispatch test derives its list rather than carrying one.** A third copy of the
command names would be a third place to be wrong, which is the defect `needsKey` being a
field rather than a switch was meant to avoid.

**`version` is exempted by name with its reason**, rather than the test being relaxed so
that any name could be missing from the table.

**The phases were written in order**, and the refactor was built and run before the
tests were written, which is how its two defects were found.

<!-- xeno:section:release-notes -->
## Release notes

`cmd/xeno` has tests behind its exit codes. The staircase the package promises — 0 where
the command did what was asked and the verdict is not red, 1 on a red verdict or a
refusal, 2 where it could not run at all — is asserted through `run`, in process, with
each case naming the stream its output arrived on.

The case a CI wrapper depends on has a test of its own: a provisional verdict verifies
as 0, a divergence or a red phase as 1. Section 6 rests the evidence arrangement on the
first, and the reason is now beside the branch.

Four more things the issue named are covered: the dispatch table against the usage
string, `init` being one word where the rest are two, a missing positional argument, and
`--export` carrying no suggestion that a shell would eval.

What made it possible: `run` takes the two writers, `opts` carries them, `report` and
`suggest` are methods on it, and `main` supplies the real files and is the only line in
the package that names them. No word, format or exit code changed.

Coverage over the module rises from 72.0% to 81.5%.

<!-- xeno:section:residual-risk -->
## Residual risk

**Twenty functions are still at zero and they are the printers.** What the tool prints
has almost no test, which XENO-0207 recorded about one sentence and is now true of
twenty functions. The five things the issue named are asserted; the printing around them
is exercised only where a test passed through it.

**No test builds a provisional phase and verifies it end to end.** The branch is
asserted directly and the command over this repository, and the path between them — a
real declaration, a real pending item, a real recompute — is covered by the runner's own
tests and by nothing that reaches an exit code. The deviation explains the choice and
does not close the gap.

**A test reads this repository**, running `gate verify` against `../..`. It passes
because the repository is clean and would fail on a committed divergence, which is the
verify workflow's job. It is the cheapest end to end cover of the command and it is an
assertion in the wrong place.

**`enforcement check` stays untested**, being the one path that needs a host.

**`main` is untested and is now the line that matters most**, being where the two
writers are supplied in an order nothing checks. Passing them the wrong way round would
send verdicts to standard error and refusals to standard output, and every test in the
package would pass.

**Accepted with the five named.** The state it replaces is a package that promised a
staircase in a comment and asserted none of it, where a shadowed receiver and a struct
missing its fields both went unnoticed until something was compiled.
