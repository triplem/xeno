---
intent: github.com/triplem/xeno#110
phase: 04-verification
created: "2026-09-29T19:13:35Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 227b7fabe08be017cc21a999e4a05a7977fd86702beaed737e882296fe8e9e7d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
evidence:
  - kind: test-report
    job: go-test
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it |
|---|---|
| AC1 only `main` names the real files | the transcript greps the package: one line, line 51 |
| AC2 the staircase | `TestTwoIsForWhatCouldNotRunAtAll` over five cases, `TestOneIsARefusalAndZeroIsSuccess`, `TestARedVerdictExitsOneAndPrintsToStandardOutput` |
| AC3 provisional exits 0, divergence and red exit 1 | `TestAProvisionalVerdictVerifiesAsZero` over seven results, including provisional beside red; `TestGateVerifyOverThisRepositoryExitsZero` for the whole command |
| AC4 the table and the usage agree | `TestTheDispatchTableAndTheUsageAgree`, which reads the names out of the usage rather than listing them |
| AC5 `init` is one word, the rest two | `TestInitIsOneWordAndTheRestAreTwo`, and `one word that is not init` in the first table |
| AC6 a missing `--intent` or an unresolvable phase exits 2 | the first table, two of its five cases |
| AC7 a missing positional argument is not silent | `TestAMissingPositionalArgumentIsNotSilent` |
| AC8 `--export` carries no suggestion | `TestExportPrintsTheEnvironmentAndNothingElse` |
| AC9 every assertion names its stream | read: each case checks `out` or `errw` and most check both, and every exit-2 case asserts standard output is empty |
| AC10 nothing changed | `gate verify` over 139, the whole suite green, and no word, format or code in the diff |

Ten rows. AC9 is the one the refactor needed and it earned itself: two defects were found by the change
and a third by the tests.

<!-- xeno:section:results -->
## Results

Thirteen tests in `cmd/xeno`, none skipped. The suite passes, `gofmt` and `go vet` are
clean, and `gate verify` matches 139 verdicts.

Coverage over the module goes from 72.0% to 81.5%, and `cmd/xeno` from 28 of 30
functions at zero to 19 of 31. The number is a consequence rather than the aim, and what
it does not say is that the five things the issue named are now asserted while twenty
functions of printing are not.

**Three defects, which is the answer to whether the package needed this.** The refactor
found two: a loop in `suggest` binding `o` over the owed obligations, so `o.out` inside
it resolved to a field of a string; and `parse` building its `opts` without the writers,
which surfaced as a nil pointer the first time the binary ran. The tests found the
third: `version` is named in the usage and absent from the dispatch table.

The third is not a bug. `run` answers `version` before the split and before the usage
check, so it works in a directory holding nothing, which is when somebody asks for a
version. The test exempts it by name and says that, rather than relaxing so that any
name could be missing.

Read rather than executed: that no word, format string or exit code changed. Every hunk
of the diff is a writer or a signature, which is what makes fifty-five mechanical edits
reviewable.

<!-- xeno:section:gaps -->
## Gaps

**Twenty functions are still at zero, and they are the printers.** `printInit`,
`printEnforcement`, `printGate` and the command wrappers around the runner. What they
print has no test, which XENO-0207 already recorded about a sentence in output and is
now true of twenty of them. The issue named five things and those are asserted; the rest
of the package is exercised only where a test happened to pass through it.

**`enforcement check` reaches a host and is untested**, deliberately. It is the one path
a test cannot take without a network, and the coverage it is missing is the coverage the
issue excluded.

**The provisional branch is asserted away from the command.** `verifyCode` takes a
result and returns a code, and seven constructed results cover it exactly. What is not
covered is the path from a real provisional phase to that result, which is `Verify`'s
and the runner's, and no test in this repository builds a provisional phase and verifies
it end to end. The deviation says why and it remains a gap.

**A test reads this repository.** `TestGateVerifyOverThisRepositoryExitsZero` runs `gate
verify` against `../..`, so it passes because this repository is clean and would fail if
somebody committed a divergence — which is the verify workflow's job, from a test. It is
a useful assertion in the wrong place and it is the cheapest end-to-end cover of the
command that exists.

**`main` is still untested**, being one line, and it is now the one line that would
break if the writers were passed in the wrong order.

**The evidence is a local run**, fifteenth in a row.
