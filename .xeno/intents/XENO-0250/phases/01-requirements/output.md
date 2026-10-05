---
intent: github.com/triplem/xeno#229
phase: 01-requirements
created: "2026-10-05T15:21:10Z"
schema_version: "1.0"
runner_version: dev+325b27a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cfae9e2e7076083310f9b153d44ec6320c4e5d599892d00900410d0b43823116
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

1. `QuestionShape` reports a question whose proper options do not all carry a consequence,
   naming how many of how many, and the free entry is exempt.

2. The writer additionally refuses a question that recommends no option, and one that
   recommends more than one, naming the count. `xeno question record` is the path that
   refuses; the gate is not.

3. The split is one function pair rather than two checkers: the writer's check calls the
   gate's and adds to it, so the two cannot disagree about what the gate already asks.

4. Both halves say in the code why the division exists, naming XENO-3's Q-2 and the
   divergence, so that the next reader does not close the gap by moving the check and
   breaking `gate verify`.

5. `./xeno gate verify` exits 0 over the 375 verdicts that exist now, plus this intent's own
   judged phases and nothing else. This is the criterion the whole design turns on and it is
   a command, not a judgement.

6. A question missing a consequence is refused by the writer too, which follows from criterion
   3 without extra code and is asserted rather than assumed.

7. The `askable` fixture in `exchange_test.go` still passes unchanged, since it is this
   project's own example of a well-formed question and a change that refused it would be
   wrong about the clause rather than strict.

8. Tests cover: a bare option refused by the gate and by the writer, a free entry without a
   consequence accepted, no recommendation accepted by the gate and refused by the writer,
   two recommendations refused by the writer, and `no_options: true` exempt from both.

9. The two gaps that need a specification commit are recorded with their evidence and not
   worked around: no field for the recommendation's reason, and no rule that questions are put
   in sequence. Each gets an issue rather than a hole in this intent.

10. `go build`, `go test ./...`, `go vet ./...` pass and `gofmt -l` outside `vendor/` prints
    nothing.

11. One commit, `Closes #229`, and the issue carries `wp5`.

<!-- xeno:section:non-goals -->
## Non goals

Not a field for the recommendation's reason. Section 8 asks for it and `model.Option` has
nowhere to put it, and adding one is an addition to what the section enumerates: a
specification change first by the second standing rule, and a person's commit by the first.

Not a rule that questions are put in sequence. It is in neither normative document, so there
is no clause to read; writing it down is a specification change. The finding is recorded with
XENO-0243's three-question batch as the example.

Not a recommendation check in the gate. The measurement is why: it would turn XENO-3's sealed
P0 red and `gate verify` to exit 1. The maintainer chose the writer, and the alternative —
overriding the finding on that verdict — was weighed and is recorded.

Not a re-judgement of XENO-3. Its verdict, its artifact and its row stay as they are. The
question genuinely fails the clause and the trail now holds a case the gate is knowingly silent
about, which the review phase states rather than hides.

Not a consequence required of the free entry. Section 8 asks for consequences on the two to
four options and then for a free entry as a further one, and what a free entry leads to is
unknowable by construction. The project's own fixture reads it that way.

Not a check that a consequence is any good. Whether "the caller retries" is a real consequence
or a restatement of the option is what a person reads it for, and A90's finding is that a
reader which cannot fail is worse than none.

Not a change to `no_options: true`. It is the exception section 8 names for a question with no
options to offer, and a question that states it is exempt from all of this, as it is today.

<!-- xeno:section:constraints -->
## Constraints

The trail bounds what the gate may ask. `QuestionShape` is reached through `phaseResult` and so
through G-Schema, which runs from P0, so every check added to it is applied to every artifact
ever written. The measurement is the constraint, not an estimate of it: the consequence check
leaves 375 verdicts at exit 0 and the recommendation check does not.

What is sealed is never rewritten, which forecloses the obvious repair. XENO-3's Q-2 could be
given a recommendation in one line and that line would change its `artifacts_hash` and stale
its verdict, so the only honest routes were the writer or an override.

The two halves must share one definition. A second checker in the writer would be two
statements of what a question is, free to drift, which is the defect `Differences` was kept
away from in #201 and which `model.ChecklistResults` is shared to avoid in #242.

The writer refuses and the gate reports, and the wording differs accordingly. A refusal is read
by whoever typed the command and can say what to do; a finding is read later by whoever reads
the verdict and has to name the artifact. `shapeRefusal` already turns the first finding into a
refusal, so the gate's phrasing has to work in both.

`model.Option`'s fields are fixed by section 8. `Consequence` and `Recommended` exist and may
be read; nothing may be added, so the reason for the recommendation cannot be stored even
temporarily.

One intent, one branch, `Closes #229`, and the issue carries `wp5`.
