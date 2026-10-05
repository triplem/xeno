---
intent: github.com/triplem/xeno#229
phase: 03-implementation
created: "2026-10-05T15:29:05Z"
schema_version: "1.0"
runner_version: dev+2d997e0.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cbaea306c782118a36a55d7654b02e3cf8cd22bb4e27ff9df15a37c8ff8881bd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`internal/gates/gates.go`. `QuestionShape` counts bare options alongside proper ones and free
ones, and reports "question Q-1 has 2 of 3 options with no consequence" where any proper option
has none; the free entry is skipped with the reason beside the `continue`. `QuestionAsked` is new
beside it: it calls `QuestionShape` and adds the recommendation check, exactly one, with the
count in the message. `QuestionShape`'s doc comment gains the paragraph explaining why the
recommendation is not there, naming XENO-3's Q-2 and the `DIVERGENT` line.

`internal/runner/exchange.go`. `RecordQuestion` calls `QuestionAsked` instead of
`QuestionShape`, with a comment at the call site saying why the writer is stricter than the gate.

`internal/gates/questions_test.go`, new, six tests. A bare option is a finding naming the count;
a well-formed question with a bare free entry is not; the gate is silent on the recommendation
and the writer is not, which is the asymmetry pinned as an assertion; two recommendations are
refused; `QuestionAsked` reports both kinds at once, which is what proves it adds to
`QuestionShape` rather than replacing it; `no_options` is exempt from both.

`internal/runner/exchange_test.go`, five tests on the writer's path, with two fixtures: `bare`
is `askable` with the consequences removed, and `unrecommended` is XENO-3's Q-2's shape — every
consequence present, nothing marked.

`internal/runner/runner_test.go`. The `question` fixture gains consequences on its two proper
options. Three tests failed on it — `TestUnresolvedQuestionTurnsP5RedAndLeavesP1Green`,
`TestQuestionResolvedTwoPhasesLater` and `TestQuestionResolvedByConfirmedAssumption` — and all
three are about a question being resolved rather than about its shape, so the fixture was
relying on the gap this intent closes. Correcting it leaves the three testing what they are
named for.

The measurement was run again after the real change rather than trusting the probe: `gate
verify` at exit 0 over 378 verdicts, which is the 375 that existed plus this intent's three
judged phases. That is criterion 5, and the whole design rests on it.

The gate had no question tests at all before this. `QuestionShape` was covered only through the
writer, in `internal/runner`, which is why a change to the gate's half of it broke three tests
in a package that was not the one being changed.

<!-- xeno:section:deviations -->
## Deviations from the design

One deviation from the acceptance criteria, and it is an addition to the work rather than a
reduction. Criterion 7 said the `askable` fixture must still pass unchanged, which it does — but
a different fixture, `question` in `internal/runner/runner_test.go`, did not, and the criteria
did not anticipate it. Its two proper options carried no consequence, so three tests about
question resolution went red on a shape finding. The fixture now carries consequences.

That is a correction and worth being plain about rather than filing under housekeeping: the
fixture was the shape section 8 forbids, and three tests had been passing over it since WP5.
Nothing was weakened to make the suite green — the three tests assert what they always did, and
the input is now the shape the section asks for.

One departure from the design, which named two places for the explanation of the split and
produced three. P2 said both functions carry it and so does the call site; writing it, the
call-site comment turned out to be the one that matters most, because `exchange.go` is where a
reader sees `QuestionAsked` chosen over `QuestionShape` and has no other reason to look for one.
The count is the same; which of the three is load bearing is not what the design expected.

No deviation on the measurement, and it was re-run rather than inherited. The probe during P0
used a hand-edited `QuestionShape`; the figure in criterion 5 comes from the real change, at
exit 0 over 378 verdicts.

Two things are left undone deliberately and both need a person, which the scope and non-goals
already said and which the review phase will carry as issues rather than as gaps: no field for
the recommendation's reason, and no rule that questions are put in sequence.

One thing was found and left. `internal/gates` had no test for `QuestionShape` at all — its only
coverage was through the writer in another package, which is why changing the gate's half broke
three tests in `internal/runner` and none in `internal/gates`. The new file closes it for
questions; whether the other exported shape functions are in the same position is not this
intent's to check, and is worth someone's pass.
