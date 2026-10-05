---
intent: github.com/triplem/xeno#225
phase: 01-requirements
created: "2026-10-05T17:15:07Z"
schema_version: "1.0"
runner_version: dev+9eec4ed.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3b60ffc5a053d3c98b74c30aee4667eb628723f1e41245120241ba59a66c0189
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `phase start` refuses a phase whose `output.md` exists and records a `context_hash` that the
   phase's current `context.lock.yaml` does not match, with the marker absent and no verdict
   written. That is the lost-marker state and it is the only new refusal.

2. The refusal names both ways on: carry on with `section set` and `phase finish` without
   starting again, or remove the artifact to start the phase over. It mirrors the verdict
   refusal #215 wrote, which offers the same two.

3. A first start is unaffected. No artifact, no lock, no refusal — this is every normal start
   and the criterion that keeps the change from being a regression for everybody.

4. A phase whose artifact exists and whose lock still matches is not refused, so a marker lost
   with nothing else changed lets the phase be re-entered. This is the case the refusal must not
   catch: if the lock still describes what the phase was given, starting again costs nothing.

5. The existing two refusals are unchanged in condition and in wording: a present marker still
   says "already running", a verdict still says what #215 made it say.

6. The refusal and the printed next step agree. For a phase in this state the suggestion must
   not offer `phase start`, or a person follows the printed advice into the refusal.

7. A test reproduces the defect with a moving clock and asserts the refusal, and a second test
   asserts the same scenario under the pinned fixture clock behaves consistently. The first is
   the one that would fail against today's code.

8. `./xeno gate verify` exits 0 with the 393 verdicts that exist now intact, plus this intent's
   own judged phases.

9. `go build`, `go test ./...`, `go vet ./...` pass and `gofmt -l` outside `vendor/` prints
   nothing.

10. One commit, `Closes #225`, and the issue carries `wp7`.

<!-- xeno:section:non-goals -->
## Non goals

Not `xeno intent resume`. #225 names the objection: a second entry point into the computation
`intent status` and the printed next step already perform. The maintainer chose the refusal.

Not the age figure. A figure nobody acts on is decoration, which the issue says of it.

Not telling a resumed phase what changed under it. Half of it exists — a phase whose predecessor
moved is reported `Stale` — and the other half is #215's, about re-reading a predecessor rather
than re-entering a phase.

Not a durable run marker. A9 puts it under `.xeno/local/` because it describes a machine's
current state and not the trail, and this intent's whole point is that its absence is normal.

Not a change to G-Schema. The red verdict it produces is correct; what was missing is anything
stopping the act that earns it.

Not a repair of an artifact already in this state. The refusal prevents it going forward; a
phase whose lock was already rewritten is a person's to resolve, and the refusal's second branch
says how.

Not a new field, a new gate or a new command. Section 6 already says a second start of a running
phase is refused; this makes the tool able to tell that a phase is running when the marker
cannot say so.

<!-- xeno:section:constraints -->
## Constraints

The marker's absence cannot be the condition. A9 makes it machine-local and gitignored and
`retention` expires it at thirty days, so refusing on a missing marker would refuse the ordinary
first start on every machine. The artifact is what distinguishes the states, because only a
phase already under way has an `output.md` recording a `context_hash`.

The pinned clock hides the defect. `newFixture` fixes `Now`, so a rewritten lock is
byte-identical and the artifact stays consistent with it; a test written in the package's usual
style passes against the broken code. Criterion 7 exists because of that and the moving-clock
test is the one that carries it.

The lock is inside `artifacts_hash` and is the only record of what the phase was given. That is
why the rewrite is a loss rather than an inconvenience, and why the refusal belongs before the
write rather than as a warning after it.

The refusal has to agree with `next.go`. The suggestion for a running phase names the sections
still wanted and offers `xeno section set`, which is the right advice in this state too, so
nothing in `next.go` needs changing — but criterion 6 asserts the agreement rather than assuming
it, because a refusal that contradicts the printed next step is worse than no refusal.

Backwards compatibility is total for the normal path. Every first start has no artifact, so the
condition cannot fire; criterion 3 asserts it rather than reasoning about it.

One intent, one branch, `Closes #225`, and the issue carries `wp7`.
