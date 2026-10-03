---
intent: github.com/triplem/xeno#184
phase: 04-verification
created: "2026-10-03T12:57:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2c336b3beb1e260ab3badb389c37adcf0d4ac499fda1c94ea837f73e773250f5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Each acceptance criterion of P1 against what holds it.

**The specification change is its own commit, before any code** — `git show --stat e8f68b1`: two
documents, nineteen insertions, no `.go` file. Held by the history rather than by a test, which is
the only place an ordering can be held.

**The variable carries a whole session** — `TestTheHarnessVersionIsReadFromTheEnvironment` at the
runner level, and measured on a throwaway tree where the variable was exported, no flag was passed
anywhere, and both artifacts of a phase came out carrying the field. This intent's own six phases
were written the same way and are the larger sample.

**The flag beats the variable** — `TestTheFlagBeatsTheHarnessVersionVariable`, at the surface
because `parse` is where the precedence lives. It sets the variable to `9.9.9` so that a passing
test cannot be a coincidence of both channels carrying the same string.

**Neither said leaves the field absent in both files** — `TestWithoutAReportTheFieldStaysAbsent`,
unchanged from #181 and now meaningful in a second way: it only tests what it claims because the
fixture clears the variable.

**The variable is read once, where the runner is constructed** — `TestTheHarnessVersionIsReadFromTheEnvironment`
sets it after `newFixture` has cleared it and calls `reopen`, which fails if the read is anywhere
but `New`. A per-command read would pass without `reopen` and that is the distinction the helper
exists to draw.

**The tests are hermetic** — `newFixture` calls `t.Setenv(HarnessVersionEnv, "")` before
constructing anything, so every test in the package is independent of the machine. Checked by
running the suite with the variable exported, which is how this session runs: the absence tests
still pass.

**The runner reads this variable and no other** — `grep -rn os.Getenv internal cmd` excluding tests
returns two hits: the enforcement token and this one. That is the assertion, and it is a grep
because what is being claimed is an absence of readers.

**The flag is not removed** — `TestTheToolVersionFlagReachesBothArtifacts` from #181, unchanged and
passing.

**Nothing else changes** — the existing suite, and `./xeno gate verify` over the whole trail.

<!-- xeno:section:results -->
## Results

**`go build ./cmd/xeno`** — builds.

**`go test ./...`** — eighteen packages, all `ok`, none skipped. Eight cases across the two
`tool_version` groups now pass, two of them new here.

**Run with `XENO_HARNESS_VERSION` exported**, which is how this session runs, so the hermeticity
claim is tested rather than reasoned: the absence tests pass with the variable set in the
environment the suite inherits.

**`gofmt -l .` outside `vendor/`** — nothing. **`go vet ./...`** — nothing.

**`./xeno gate verify`** — `verified 253 verdicts`, exit 0, no divergence, nothing red, nothing
provisional.

**`grep -rn os.Getenv internal cmd`, excluding tests** — two hits: `internal/enforcement` for the
token, and `internal/runner` for this variable. The three other names in section 7's list appear in
no Go file at all.

**The specification commit** — `e8f68b1`, two documents, 19 insertions, 8 deletions, no code. The
first commit in this repository to change the process definition at the agent's hand, and it exists
because the maintainer approved the change in as many words.

**The session channel, measured end to end.** On a throwaway tree: `export
XENO_HARNESS_VERSION=2.1.276`, then `intent start`, `phase start`, three `section set` and a
`phase finish` with no flag on any of them. Both artifacts carry `tool_version: 2.1.276`.

**This intent's own trail.** Six phases, twelve artifacts, every one carrying the field, written
with the variable alone — no `--tool-version` was passed anywhere in XENO-0229.

**Evidence is self-reported and local.** The gate path makes no network call; the pipeline runs the
same commands on a push.

<!-- xeno:section:gaps -->
## Gaps

**Nothing sets the variable for a session yet.** This intent moved the remembering from six flags
to one export, and an export somebody has to type is still something somebody has to type. The
entry point that would set it from the client's own environment is #183 and it is unbuilt, so the
claim here is deliberately the smaller one and P0 said so before any code was written.

**The runner reads one of section 7's four variables.** `XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA` and
`XENO_HARNESS` are still read by nothing, and the `--plugin-root` argument and the resolution order
are still unbuilt. After this change #183's title is less true and its substance is untouched,
which is worth saying because a partially answered issue is the kind that gets read as answered.

**The value is still unverifiable.** `export XENO_HARNESS_VERSION=9.9.9` is recorded as readily as
the truth — the surface test relies on exactly that — and so is a wrong `--tool-version`. The same
holds for `model` and `tool` from `project.yaml`. The process records what a session reports about
itself, which is A35's territory and not a defect of this change.

**Hermeticity is asserted by one line in one fixture.** A test in this package that builds a runner
without `newFixture` would inherit the environment again, and nothing stops one being written. A
lint for it would be a rule, and rules go through section 10.

**Three provenances for one field now.** Typed by hand before #181, flag-reported in XENO-0228,
variable-reported from XENO-0229 on, with nothing in any artifact saying which. A reader has the
commits and `runner_version` to date them by. Backfilling is not available and should not be: the
hashes are sealed.
