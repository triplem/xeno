---
intent: github.com/triplem/xeno#237
phase: 04-verification
created: "2026-10-06T12:10:41Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 335c05490988ca9e77dcec89a9bd28af55ebd353f4bd47087d7d9cadeca372c6
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
      sha256: 869b496a82ca275dcf6e31d5e316ef7d76403a33cc1c9036f5cfe4dedcb3bf2c
    - format: other
      job: remedy-followed
      kind: other
      path: evidence/remedy-followed.txt
      produced_by: the remedy followed end to end on a scratch copy
      sha256: 2ba6e9d911f2a5b91d0ded7fceae597a1f9b1ca8b305c21147241ee97591f60b
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Ten criteria, by number. Criterion 1's row is the issue's own "Done when" and it was checked by
following the remedy rather than by reading it.

| # | what it asserts | how it is checked | result |
|---|---|---|---|
| 1 | the two remedies name section 7's two routes | both routes followed on a scratch copy; `evidence/remedy-followed.txt`, and `TestAStalePhasesRemedyNamesBothRoutesSectionSevenNames` | pass |
| 2 | each says why `phase finish` will not clear it | the `assertSealedLockRemedy` helper, in both tests that share the remedy | pass |
| 3 | each names which phase to act on | the same helper, asserting the earlier phase and the gated phase separately | pass |
| 4 | the re-run route says what it costs | the same helper, "every phase after it"; confirmed able to fail by removing the clause | pass |
| 5 | the approval is a second person's act, not a command | a person reads the string; `grep` for `gate approve` in `gates.go` returns nothing | pass |
| 6 | the unreadable remedy stays and says why it differs | `TestAnUnreadableInputIsTheTreesFaultAndNeedsNoRelease`, which asserts the absence of the other two routes | pass |
| 7 | a test asserts what each remedy names | three tests added; each confirmed able to fail by mutating the remedy | pass |
| 8 | no condition of the check changes | `git diff` over `staleReads` | pass |
| 9 | `go test`, `go vet`, `gofmt`, `gate verify` | run; `evidence/go-test.txt` | pass |
| 10 | no normative document is touched | `git diff --stat`, which names two files in `internal/gates` | pass |

<!-- xeno:section:results -->
## Results

## The remedy was followed, not read

On a scratch copy of this repository, an intent was run to P2 with a lock recording one file, the
file was changed and committed, and P2 was gated over that range. The finding appeared with the new
remedy:

    F-2c54fb .../01-requirements/context.lock.yaml: 01-requirements was given probe.txt
      and it has changed since
      next: start 01-requirements over, which discards it and every phase after it, or a
      second person approves this finding on 02-design as still valid; section set and
      phase finish will not clear it, because the lock keeps what the phase was given and
      only phase start writes one

Then the second route was taken — `gate approve F-2c54fb --phase 02 --by "a second person" --reason
…` — and three things followed:

    the finding carries decision: {type: approved, by: a second person, against: e3d9e063…}
    02-design status: approved
    03-implementation started

So following the remedy clears the finding and the work proceeds, which is the issue's "Done when".
`evidence/remedy-followed.txt` carries the verdict and the decision as the runner wrote them.

## The first attempt passed, and that is why there is a second

The check was first built by moving the ground under P0 and gating P1. G-Freshness passed.
**P0's `context.lock.yaml` records no files**: `phase start` writes the lock and `scope set` writes
P0's scope afterwards, so it is born empty. Across the trail that is 0 of 117 P0 locks, and 17 of 66
P1 locks carry files. A green gate was about to be written down as evidence that the remedy worked.

That is #263's failure shape exactly — a tool asked about something absent answers as it does about
something that does not match — and it is why the proof rests on P1's lock and P2's gate.

## The tests can fail

Each assertion was checked by mutation and not by being seen to pass:

| mutation | what failed |
|---|---|
| the release route removed from the remedy | both shared-remedy tests, on the gated phase's name |
| "which discards it and every phase after it" removed | the changed-file test, on what the re-run costs |
| the unreadable case given the shared remedy | the unreadable test, on asking for an approval it does not need |

Restoring the file returns all three to passing. A90's objection applies to a test as much as to a
gate, and three assertions over a string are the kind that pass on something incidental.

## What did not change

`git diff` over `staleReads` touches the three `finding` calls and nothing else: the loop bound `i <
idx`, the `touched[f.Path]` guard, the hashing, the three causes and the sort are as #236 left them.
The three tests #236 brought are untouched and pass. `grep` for `gate approve` in `gates.go` returns
nothing, so the command is still only in `next.go`, where its comment explains why.

## The suite

`go test ./...` exits 0 across every package, in `evidence/go-test.txt`. `go vet` is clean, `gofmt
-l` outside `vendor/` prints nothing, and `xeno gate verify` exits 0.

<!-- xeno:section:gaps -->
## Gaps

**The check this repairs has almost no input, and that is not fixed here.** 0 of 117 P0 locks and 17
of 66 P1 locks record a `files` list. So for most of this trail the staleness half could not fire at
all, which is both why the broken remedy survived and the reason to doubt that a repaired one will be
met often. It is a change to when `phase start` writes the lock relative to `scope set`, which is
#215's territory and section 5's, and it is filed rather than absorbed under the third standing rule.

**The composed suggestion names the release twice.** `Runner.red` adds "Or a second person releases
it, with gate approve or gate override" after the remedy, so a reader of `xeno gate run` meets the
approval in both. Recorded in P3's deviations and visible in the scratch run. The alternative left
the reader of `gate.yaml` with one route where section 7 gives two.

**"Then judge it again" sits between the two routes in that composed sentence**, which reads right
for the re-run and wrong for the approval. It is `next.go`'s sentence and a non-goal here, and it is
only visible now that the remedy carries two routes.

**The re-run route has not been followed end to end.** Only the approval was. Starting the earlier
phase over means removing its directory and every phase after it, which the scratch copy would have
had to be rebuilt to show; the refusal that names the act was reproduced instead, and the remedy's
wording was written against it. So one of the two routes is proved and the other is quoted from the
runner's own refusal.

**Nothing asserts that the remedy and the refusal agree.** The remedy says `section set` and `phase
finish` will not clear the finding because only `phase start` writes a lock; `Start` refuses a second
start for the reasons #215 gives. Two statements in two packages that have to stay true together,
with no test over the pair. A reader is what connects them.

**The unreadable case's test depends on not running as root.** It chmods a file to 0o000 and expects
the read to fail. As root it would succeed and the test would report no finding, failing for a reason
that has nothing to do with the remedy. It is not guarded, and it is the kind of thing that passes
everywhere until it does not.
