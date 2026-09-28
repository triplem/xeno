---
intent: github.com/triplem/xeno#108
phase: 04-verification
created: "2026-09-28T17:00:16Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+023193e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 570ea184183d87caeae9a62435cc1b79532d2c7afcd06e85c68ed0c8d637dbe8
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
| AC1 the field is the hash of the lock beside the artifact | `TestSectionSetWritesTheContextHashOfTheLockBesideIt`, first assertion: the value read back from `output.md` equals `hashing.FileHash` of the phase's `context.lock.yaml` |
| AC2 the supported walk reports nothing about the field | `TestTheSupportedWalkIsNotRedOnTheContextHash`: start, three section writes, evaluate, fail on any finding whose cause or next step names `context_hash`. The other five fields are still reported and the test does not touch them |
| AC3 a second write does not move the value | `TestSectionSetWritesTheContextHashOfTheLockBesideIt`, second assertion: the value after a second `section set` equals the first |
| AC4 no lock means no field and no refusal | `TestSectionSetRendersAnchorsTheCallerNeverWrites` and every existing test that writes sections without starting a phase: they pass unchanged, which is only possible if the absent lock is skipped rather than refused |
| AC5 the amended rows keep what they said | read, not tested. `git show d39537a -- ASSUMPTIONS.md` is a two line diff and both lines are additions to the status column |

AC4 has no test of its own, and that is a gap rather than an oversight; it is named below.

<!-- xeno:section:results -->
## Results

`go test ./...` passes on every package. `gofmt -l .` outside `vendor/` prints nothing,
`go vet ./...` is silent, and `./xeno gate verify` recomputes and matches every verdict
the repository carries, which now includes the phases of this intent.

Both new tests were run against the tree with the writer removed and both fail there:
`context_hash is "", want the hash of the lock beside it`, and `G-Schema reports on
context_hash: required field missing: context_hash / add context_hash to the
frontmatter`. That is what makes them tests of this change rather than of the fixtures.

The strongest result is not in the report. Every phase of this intent was produced by
`phase start` and `section set` and carries a `context_hash` the runner wrote, and each
one is green on G-Schema. The record of the fix is the first artifact set the fix made
possible.

The declared evidence is the transcript of that run, attached by hash and copied into
`evidence/` of this phase.

<!-- xeno:section:gaps -->
## Gaps

**The pipeline publishes no artifact.** `.github/workflows/xeno.yml` runs the tests, the
format check, the vet and `gate verify`, and uploads nothing, so there is no CI artifact
for this phase's declaration to bind to. The evidence attached here is a local run,
bound by its hash and carried in the repository, through the `--evidence-from` directory
the tree provides as the stand-in for a fetch adapter. It proves what was run on this
machine; it does not prove what the pipeline saw. That is a finding about the pipeline
and belongs to whichever package owns evidence publication, not to this intent.

**AC4 is argued, not asserted.** That an absent lock leaves the field out is covered
only by the existing tests continuing to pass. A test that writes a section with no
phase started and reads the frontmatter back would state it directly, and it is one test
away.

**G-Test and G-Build read a result this phase cannot fully give.** G-Build passes here,
and G-Test is not in the implemented set at all, so the mapping above is checked by a
reader rather than by a gate. The mapping's correctness is therefore as good as the
reading.

**`digest.md` still carries a `context_hash` a hand wrote.** Section 5 requires the
field on both files a session produces. The runner writes one of them. The digests of
this intent were written by hand, this one included, so half of what the fix is about is
still done the unsupported way one file over.
