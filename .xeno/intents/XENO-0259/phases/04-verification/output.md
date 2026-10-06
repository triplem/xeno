---
intent: github.com/triplem/xeno#260
phase: 04-verification
created: "2026-10-06T09:23:40Z"
schema_version: "1.0"
runner_version: dev+081da51.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5a80b80f1d7d7d5756e9354ca5524670f444c715584e1d5cd527efc88f12c7fe
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - format: other
      job: go-test
      kind: test-report
      path: evidence/go-test.txt
      produced_by: go test ./...
      result: pass
      sha256: 32bee7ff67eebb781e060cd859defe2c4fdcbdd1c6c3eb569c4f4d74b2cfc62d
    - format: other
      job: npm-audit
      kind: scan
      path: evidence/npm-audit.json
      produced_by: npm audit --json in the tree the action installs
      result: pass
      sha256: 1765c77d76fef280c5d141780a388db22e143b3c4b9a7f3daafe1f46e80dbe08
    - format: other
      job: audit-reproduction
      kind: other
      path: evidence/audit-reproduction.txt
      produced_by: the audit job's three steps, run by hand
      sha256: 223fffac1e15e2ad9dd60a177d96e1cf87ae16d420acbb7a93187df0978b7d8c
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

P1's criteria are numbered, so this maps by number. Ten criteria, ten rows, and the column that
matters is the last: four of the ten are read by a person, because what they assert is prose.

| # | what it asserts | how it is checked | result |
|---|---|---|---|
| 1 | the audit installs the action's three steps | the three steps run by hand from `release.yml`'s own two identifiers, then `npm audit`; `evidence/audit-reproduction.txt` | pass |
| 2 | the sha is read out of `release.yml`, not copied | both `sed` patterns run against the file: version `25.0.9`, sha `b12c8f6…`; `grep -c` for a second copy of either in `audit.yml` | pass |
| 3 | the baseline records that tree's figures | `npm audit --json` of the reproduced tree against `.github/npm-audit-baseline.json`, by the job's own comparison script | pass |
| 4 | the commit message says the raise is a correction | a person reads the commit message before it is made | pass |
| 5 | the check still fails only on a rise | the comparison script is unchanged; `git diff` over the `audit` step's Python | pass |
| 6 | every false claim is corrected where it is made | `grep -rn "exactly as the release"` over the tree | pass |
| 7 | `docs/supply-chain.md` says 25.0.9 | `grep -n "24.2.9"` over the document | pass |
| 8 | A44 records what it now covers | a person reads the row | pass |
| 9 | `gate verify`, `go test`, `go vet`, `gofmt` | run; `evidence/go-test.txt` | pass |
| 10 | the audit job passes on this intent's pull request | the run itself, which has not happened yet | **open** |

Criterion 10 cannot be closed here and saying otherwise is what the gaps section is for. It is the
one claim in this intent whose evidence is a CI run and not a local one.

<!-- xeno:section:results -->
## Results

## The reproduction

The job's three steps were run by hand on 2026-10-06 under node 26.9.0 and npm 12.0.2, starting
from the same two `sed` patterns the job uses against the real `release.yml`:

    resolved version=25.0.9
    resolved action=b12c8f6015dc215fe37bc154d4ad456dd3833c90
    npm --loglevel error ci --omit=dev
    npm install --no-audit --no-fund semantic-release@25.0.9
    npm audit --json

    counts: {"critical": 2, "high": 23, "moderate": 2, "low": 1}
    risen: none
    comparison exit=0

The comparison is the job's own script, unmodified, reading the committed baseline. It exits 0,
which is criterion 3. `evidence/npm-audit.json` is the report it read and
`evidence/audit-reproduction.txt` the transcript.

## The flag that differs

`--only=prod`, which the action passes, and `--omit=dev`, which this job passes in its place, were
each run against the same checkout and produced the same tree: 2 critical, 23 high, 2 moderate, 1
low in both. The deviation is measured and not argued.

## The two patterns

Both were run against `.github/workflows/release.yml` rather than reasoned about. The version
pattern yields `25.0.9`; the sha pattern yields `b12c8f6015dc215fe37bc154d4ad456dd3833c90`; the
`- uses:` variant the pattern was first written as yields nothing, which is how the shape of the
line in `release.yml` was settled. `grep -c` for either identifier in `audit.yml` returns 0, so
neither is copied and criterion 2 holds.

## The claim, where it was and where it is now

`grep -rn "exactly as the release"` over the tree returns eleven lines. One is
`.github/workflows/audit.yml:7`, which quotes the old claim in order to say it was false. The other
ten are sealed artifacts — XENO-0230's three phases, which were written when it was believed, and
this intent's own five, which are about it. **No file asserts it any more**, which is what criterion
6 asks; the sealed ones are not rewritten, under section 11.

`grep -c "24.2.9"` over `docs/supply-chain.md` returns 0 and the table's row reads 25.0.9, which is
criterion 7.

## The suite

`go test ./...` exits 0 across every package, in `evidence/go-test.txt`. `go vet ./...` is clean,
`gofmt -l .` outside `vendor/` prints nothing, and `xeno gate verify` reports **432 verdicts
verified** at exit 0. No Go source was touched, so these say nothing was touched by accident —
which is the whole of what criterion 9 claims for them.

## What a person read

Criteria 4, 8 and the prose half of 6 are answered by somebody reading, and are marked pass on that
basis rather than on a command. The commit message names all four severity classes and says the
raise is a correction; A44's row carries the measurement, its date and the two criticals it now
covers. `docs/clause-readers.md` is the project's own record that a clause read by a person is a
real answer, and these are three of them.

<!-- xeno:section:gaps -->
## Gaps

**Criterion 10 is open and only CI can close it.** Every figure here was measured under node 26.9.0
and npm 12.0.2; CI pins node 24 and the action declares `using: node24`. 531 of the 532 packages
come from the action's lockfile, which no node version re-resolves, so the exposure is the single
`npm install semantic-release@25.0.9` — and that install added no path the lockfile did not already
hold. The risk is therefore small and it is not zero, and the honest statement is that the baseline
is confirmed by the first run of the job and not by this phase. If the run disagrees, the figures
and not the design are wrong.

**The second checkout has not run under `actions/checkout`.** The reproduction used `git fetch
--depth 1` of the same sha, which produces the same tree and is not the same command. What is
untested is the action's own behaviour with `repository:`, `ref:` and `path:` together — a public
repository, a 40-character sha and a relative path, all three of which it documents. It fails loudly
if it fails at all, and that is the whole mitigation.

**The baseline still cannot tell which advisory is which.** It is four counts, so a critical
replacing a critical reads as no change. A44 accepted this at ten and two and it is accepted again
at twenty-eight; what is new is that the comment names all twenty-eight, so a reader can diff by
hand what the check cannot. Nothing checks that the comment and the counts agree, and the next
person to raise a number has to update both.

**Nothing watches whether the pins are still the ones to be on.** Unchanged by this intent and
explicitly out of scope, but worth stating beside a baseline that now carries two criticals: the
numbers are correct and stable, and a correct stable number is not evidence of currency. #44 holds
that question and the deferred sha bump is where it will next be asked.

**Two criticals are now permanently visible on every run.** That is the intended effect and it is
also the standing hazard A44 names from the other side — a number nobody can act on is a number
people stop reading. It is bounded here by the counts only being able to move on a commit in this
repository, so the next move is attributable rather than ordinary. It is not eliminated.

**The `release-toolchain/` tree lives in the workspace for the length of the job.** Nothing reads it
but the two steps that build and audit it and nothing uploads it, but the workspace is no longer
only this repository while the job runs. No check asserts that, and a later step added to this job
would meet it.
