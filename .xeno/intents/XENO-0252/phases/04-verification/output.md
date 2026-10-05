---
intent: github.com/triplem/xeno#205
phase: 04-verification
created: "2026-10-05T16:51:55Z"
schema_version: "1.0"
runner_version: dev+9eec4ed.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1fe9913977a28a53a9e6d30cd71b11695456aca8ee8ccb8a018b74dcbf5cbb08
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
      sha256: 423d90a94681412807a3290c6215b60a38390978dd762f6bd1e3470c82cb907a
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Eight tests across two packages, plus one run of the real command and one declaration so G-Test
judges this intent.

| criterion | what answers it |
|---|---|
| 1, the resolver | `TestUnsetResolvesToWhatTheLiteralsMeant`, `TestAnAbsoluteValueIsUsedAsItStands`, `TestARelativeValueIsRelativeToTheRoot`, `TestAnEmptyOrBlankValueReadsAsUnset` |
| 2, the helper and the constants | the four tests above all go through `LocalPath` or `LocalDir`, and the constants are what they name rather than strings repeated in the test |
| 3, the literals are gone | one grep over Go sources outside tests: the only hits are the `.gitignore` site, the constant itself and comments |
| 4, no duplicated `phase.env` | read, plus `PhaseEnvFile` being the one definition both the runner and the command use |
| 5, the gitignore entry | read, and `init`'s own tests still pass unchanged |
| 6, setting it moves the files | `TestTheLocalDataLocationMovesWithTheVariable` for the marker and `phase.env`, `TestTheReportAndTheLedgerResolveToTheSameDirectory` for the other two |
| 7, unset moves nothing | asserted twice: over the strings per path, and over the disk after a real `phase start` |
| 8, the trail | `./xeno gate verify` at exit 0 over 390 verdicts |
| 9, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 10, the register row | A97 |
| 11, one commit | the commit itself |

Criterion 7 is asserted twice on purpose and the second assertion is the one that matters.
Checking the resolver's strings catches a wrong `filepath.Join`; running `phase start` and
looking at the disk catches a writer that resolves correctly and then writes somewhere else,
which is a different mistake and the one a string test cannot see.

The absolute case is tested because the entry point produces it. `xeno-env.sh` exports
`${XENO_PLUGIN_DATA:=$root/.xeno/local}`, so the value in practice is absolute and a resolver
that joined it to the root would nest one path inside another — the only behaviour here that
could have been wrong in a way nothing in this repository would ever have shown.

Beyond the tests, it was run. In a scratch repository with the variable set, `intent start` and
`phase start` wrote `phase.env` and `runs/DEMO-0001/00-intake.lock` under the redirected
directory and left `.xeno/local` empty. That is the promise the export has made since it was
written, exercised through the binary rather than through a fixture.

This phase declares `test-report/go-test`, the output of `go test ./...` bound by hash, so G-Test
judges it — the gate implemented in the previous intent reading this one's evidence.

<!-- xeno:section:results -->
## Results

Ten criteria met, one pending until the commit.

1. Met. `LocalDir` resolves the variable where set, the default where not, an absolute value as
   it stands and a relative one against the root.

2. Met. `LocalPath` joins inside it; `LocalDataEnv` and `LocalDefault` are the only places either
   string appears.

3. Met. The six literals are gone. One grep over Go sources outside tests finds the `.gitignore`
   site, the constant and three comments.

4. Met. `PhaseEnvFile` is defined once and used by the runner and the command.

5. Met. `init` writes `model.LocalDefault + "/"` with the comment saying why it does not resolve.

6. Met. Setting the variable moves the run marker, `phase.env`, the enforcement report and the
   ledger; the first two asserted on disk after a real `phase start`, the last two over their
   resolvers.

7. Met, twice. Unset resolves to the four strings the literals produced, and a real `phase start`
   with the variable unset leaves the marker and `phase.env` exactly where every existing
   repository has them.

8. Met. `./xeno gate verify` exits 0 over 390 verdicts: the 387 that existed are intact and the
   three are this intent's own judged phases. Nothing under the local data location is hashed,
   which is why the expectation was that this change is invisible to the trail, and the command
   is what confirms it.

9. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
   nothing, and `go test ./...` is `ok` across all eighteen packages that have tests — the run
   this phase declares.

10. Met. A97, with A89's measurement as the contrast and the state-splitting cost stated.

11. Pending. The commit comes after this phase is judged.

**One defect in this intent's own record, which no criterion covers.** P3's learning carries a
dropped word: its proposal reads "list its readers as well as the places that duplicate its
value. for the identifier and not only for the string", and the missing word is the name of the
tool. The shell ate it — the proposal was passed as a double-quoted argument containing a
backquoted word, so the shell ran that word as a command, the command failed, and the empty
output was substituted in. The learning is recorded, the sentence is readable, and the phase is
judged, so repairing it would mean rewriting a sealed artifact or re-judging a verdict over a
typo. It is left, said here, and its cause is this phase's learning, because the hazard will
recur and is worth more written down than the typo is harmful.

<!-- xeno:section:gaps -->
## Gaps

A person can split their own state and nothing can tell them. The variable is read at each use,
so setting it between `phase start` and `phase finish` leaves the marker in one directory and the
lookup in another, and the second command reports the phase as not running — which is
indistinguishable from a phase that genuinely is not. A97 states it. Detecting it would mean
recording the location inside the phase, which would put an environment value into the trail and
is the thing A89 refused for a different variable.

Nothing validates the value. A path that cannot be created fails at the write, as `.xeno/local`
does today, so `XENO_PLUGIN_DATA=/nonexistent/x` is a confusing error at `phase start` rather
than a clear refusal at the point of setting. A resolver that checked would be checking at a
moment when nothing is being written, which is why it does not.

The entry point is unchanged and untested here. `xeno-env.sh` exports the variable and this
intent reads it, but nothing in this repository runs the script: the absolute-value case is
tested against a value shaped like what the script produces, not against the script. That is the
same gap #201 recorded for the release workflow, one level smaller.

A relocated directory outside the repository is not gitignored and cannot be. Criterion 5 keeps
the entry at the default deliberately; a person who redirects state to a path inside the tree but
outside `.xeno/local/` gets untracked files and no warning. The runner cannot know which case
they are in.

The two constants that changed meaning are a compatibility edge nothing guards. `cost.Ledger`
and `runner.ReportPath` were exported paths and are now a file name and a function; anything
outside this repository reading them as paths breaks at compile time, which is the good case, and
nothing does. It is worth knowing rather than checking.
