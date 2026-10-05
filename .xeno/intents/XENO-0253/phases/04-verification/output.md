---
intent: github.com/triplem/xeno#225
phase: 04-verification
created: "2026-10-05T17:28:35Z"
schema_version: "1.0"
runner_version: dev+6adc0f9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2f7c6312fc52c1f37d989f482d56537ca53f61004fd54bb760a30b21e7ca7ed5
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
      sha256: ccae2c0f116e32bcdfb29c577718108f60769af18eac38cec1fe1e4277a9d9c2
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Six tests in `internal/runner/runner_test.go`, two existing ones rewritten, and one declaration
so G-Test judges this phase.

| criterion | what answers it |
|---|---|
| 1, the refusal | `TestAStartThatWouldStaleTheArtifactIsRefused`, which also asserts the lock was not rewritten by the refused start |
| 2, both ways out | the same test, on the message: "already under way", "section set and phase finish", and the phase directory |
| 3, a first start is unaffected | `TestAFirstStartIsUnaffected`, under the moving clock |
| 4, the narrow case | **falsified, not met.** See the results section: there is no harmless second start, and `TestThereIsNoHarmlessSecondStartOfAPhaseUnderWay` asserts the opposite of what the criterion asked for |
| 5, the older refusals | `TestTheOlderTwoRefusalsAreUnchanged` for the conditions; their wording changed with the escape hatch, which `TestStartIsRefusedWhereThePhaseHasAVerdict` now asserts |
| 6, refusal and suggestion agree | `TestTheSuggestionAgreesWithTheRefusal`, which failed against the first implementation and is the reason the state computation changed |
| 7, the moving clock | `f.moving()` and the reproduction; the pinned-clock half of the criterion became `TestAPhaseWithAnArtifactIsRunningWithoutItsMarker` |
| 8, the trail | `./xeno gate verify` at exit 0 over 396 verdicts |
| 9, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 10, one commit | the commit itself |

**The moving clock is the whole of why this was findable.** `newFixture` pins `Now`, so a
rewritten lock comes out byte-identical, the artifact stays consistent with it, and a test
written in the package's usual style passes against the broken code. The defect is the `created`
stamp moving. `f.moving()` is four lines and it is the only reason the reproduction fails against
`Start` as it was.

Two tests assert absences and both earned their place. The refusal test checks the lock's hash is
unchanged after the refusal, because a guard that refuses and writes anyway would pass a test
that only read the error; and the agreement test checks the suggestion does not contain
`phase start`, which is what caught the state computation being wrong.

`TestAPhaseWithAnArtifactIsRunningWithoutItsMarker` is the one test here about `intent status`
rather than about `Start`. It covers the gap #225 names first — a phase begun on one machine
reading as not-started on another — which this intent closes as a consequence of making the
suggestion correct rather than as a goal.

<!-- xeno:section:results -->
## Results

Eight criteria met, one falsified, one pending until the commit.

1. Met. A phase whose artifact exists and whose verdict does not is refused, and the refused
   start leaves the lock's hash unchanged.

2. Met. The message names carrying on with `section set` and `phase finish`, and removing the
   phase directory to start over.

3. Met. A first start has no artifact and is unaffected, asserted under the moving clock.

4. **Falsified.** The criterion asked that a phase whose artifact exists and whose lock still
   matches be allowed to start, on the reasoning that a marker lost with nothing else changed
   costs nothing. That state does not survive a real clock: `created` moves, the rewrite is never
   byte-identical, and the artifact goes stale every time. There is no harmless second start of a
   phase under way, and the test now asserts that rather than its opposite. The criterion was
   written from P2's alternatives, which were wrong for the same reason.

5. Met in condition, changed in wording. A present marker still says "already running" and a
   verdict still says "has a verdict"; both now name the phase directory as the way to start
   over, because removing the verdict alone leaves the artifact and lands on the new refusal.

6. Met, and it failed first. The suggestion offered `phase start` for the state being refused,
   because the phase state read a phase under way as `not-started` when this machine had no
   marker. Fixed in the state computation rather than in `next.go`.

7. Met. `f.moving()` carries the reproduction; the pinned-clock case became the state test.

8. Met. `./xeno gate verify` exits 0 over 396 verdicts: the 393 that existed are intact and the
   three are this intent's own judged phases.

9. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
   nothing, and `go test ./...` is `ok` across all eighteen packages that have tests — the run
   this phase declares.

10. Pending. The commit comes after this phase is judged.

**One behaviour changed beyond what the criteria named**, and it is the thing a reviewer should
look at hardest: starting a phase over now takes the phase directory rather than the verdict.
#215 documented removing the verdict, and after this that leaves the artifact and lands on the
second refusal. Both messages and #215's comment and test were corrected together, so nothing
sends a person from one refusal into another — but it is a change to a documented escape hatch,
made because the alternative was leaving the instruction wrong.

<!-- xeno:section:gaps -->
## Gaps

Two of #225's three gaps remain untouched and one is closed as a side effect. Nothing says how
long an intent has been sitting, and nothing tells a resumed phase what changed under it; both
shapes the issue offered for them, a figure and a command, were weighed and declined. What
closed, unplanned, is the first: a phase begun on one machine no longer reads as `not-started`
on another, because the state now keys on the artifact rather than on the marker.

A phase can still be entered in a state nothing warns about. The refusal stops a second start;
it does not help somebody who has been away for a week and wants to know what the phase already
contains. `xeno intent status` and the printed next step remain the whole answer, which is
#225's own conclusion, and this intent does not improve it.

An artifact whose lock was already rewritten is not repaired. Anyone who hit this before the
refusal has a phase that will go red on G-Schema, and the message's second branch — remove the
phase directory — is the only route back. Nothing detects the state and says so; `intent status`
reports the phase as running and the finding arrives at `phase finish`.

The guard cannot tell a deliberate restart from an accident, and now neither can it be told.
Before, removing the verdict expressed "start over"; now the only expression is removing the
phase directory, which is more destructive and therefore a clearer statement of intent — but
somebody who wanted to keep their sections and re-run the start has no way to say so. That is
deliberate, because keeping the sections is exactly what `section set` and `phase finish` do
without a start.

The marker is now almost vestigial. It still gives the "already running" refusal on the machine
holding it, but the artifact gives a stronger one everywhere; what the marker adds is the window
between `phase start` and the first `section set`, when no artifact exists yet. Worth knowing
before anybody prunes it, and not this intent's to decide.
