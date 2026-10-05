---
intent: github.com/triplem/xeno#229
phase: 05-review
created: "2026-10-05T15:31:15Z"
schema_version: "1.0"
runner_version: dev+2d997e0.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 40fa9cb3fb27de8c40b4a3021afd6e67b2ab7d75602f3d4f8ac6cfd16e977acd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Two, each naming what it departs from. Against the acceptance criteria: criterion 7 required the askable fixture to pass unchanged and it does, but a second fixture the criteria did not anticipate, question in runner_test.go, carried two options with no consequence and failed three tests about question resolution; it now carries consequences, and P3''s deviations says plainly that nothing was weakened to make the suite green. Against P2''s design: it named two places for the explanation of the split and three were written, with the call site in exchange.go turning out to be the load-bearing one, because that is where a reader sees QuestionAsked chosen over QuestionShape.'
      result: deviation
      rule: deviations-are-traceable
    - note: 'An interface outside this intent changes and nothing migrates, because what it rejects was never valid. QuestionShape keeps its name, signature and every caller, and gains a check; QuestionAsked is new. The dependants are callers of QuestionShape, which are the gate paths and, until this commit, the writer — the writer now calls QuestionAsked, one line, in this same commit. What changes for anyone outside is that xeno question record refuses a question section 8 already forbade: bare options, or no recommendation. There is no migration because there is no valid prior usage to migrate, and the refusal names the count so the repair is visible from the message.'
      result: deviation
      rule: interface-change-needs-a-migration-note
    - note: 'None added. go.mod is untouched; the change uses fmt and strings, both already imported by gates.go, and the tests use strings and testing. No tool was weighed either: the one thing that would have needed something new is a reader for question sequencing, which needs a timestamp field rather than a dependency, and that is #248.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: section 8 already asks for the
consequence and the recommendation, so the code moves towards the specification and no
specification commit precedes this. Nothing is invented — `Consequence` and `Recommended`
existed and stopped being decoration, and the two things that would have needed a new field or
a new rule are #247 and #248 rather than code. The branch carries one intent, the commit
references #229, and the issue carries `wp5`.

The acceptance criteria. Ten met, one met by the commit this phase precedes. Criterion 5 is the
one the design turns on and it is a command: `gate verify` at exit 0 over 378 verdicts, re-run
against the real change rather than against P0's probe.

The non-goals held. No field for the reason. No sequence rule. No recommendation check in the
gate. No re-judgement of XENO-3, whose verdict, artifact and row are untouched. No consequence
required of the free entry. No judgement of whether a consequence is any good.

What a reviewer should check is the asymmetry and whether it is the right trade. The gate reads
two of section 8's three requirements and the writer reads the third, so a hand-written question
that recommends nothing passes every gate. The alternative was releasing a finding on XENO-3's
sealed P0, which the maintainer weighed and rejected on cost; both are in P2's alternatives with
the measurement that forced the choice.

What this intent got wrong is in P3's deviations. A fixture in `internal/runner/runner_test.go`
had carried the shape section 8 forbids since WP5, three tests passed over it, and the criteria
did not anticipate it. Nothing was weakened to make the suite green: the fixture now carries
consequences and the three tests assert what they always did.

What is left is in P4's gaps and in two issues. The recommendation has no reader in the gate;
the reason has no field at all; sequence may be uncheckable without a timestamp; and whether
`EvidenceShape` and `DecisionShape` are in the position `QuestionShape` was in — exported for a
writer and tested only through it — is nobody's yet.

<!-- xeno:section:release-notes -->
## Release notes

Section 8 asks three things of a question's options. Two of them are read now, where one was.

> Each one carries two to four options with their consequence, the agent's recommendation
> with a reason, and always a free entry as a further option.

**The consequence is read by the gate.** A question whose proper options do not all say what
taking them leads to is a finding, naming the count: "question Q-1 has 2 of 3 options with no
consequence". The free entry is exempt, because it stands for an answer nobody has written yet.

**The recommendation is read by the writer.** `xeno question record` refuses a question that
recommends no option, and one that recommends more than one — section 8 says "the agent's
recommendation", and a list with everything starred has made no choice.

So a question with four bare options and no recommendation, which was well formed as far as
anything could tell, is now refused at the moment of asking.

**The recommendation is deliberately not in the gate, and that is the one thing to know.**
`QuestionShape` is reached through G-Schema, which runs from P0, so a check added to it judges
every artifact ever written. XENO-3's Q-2 recommends no option — three options, each with a
consequence, none marked — and with the check in the gate, `xeno gate verify` reports
"DIVERGENT XENO-3 00-intake: committed status green, recomputed red" and exits 1. What is
sealed is never rewritten, so the choice was to release that finding on a pre-M0 verdict or to
enforce the clause on questions written from now on. The second was taken. Both functions and
the call site say so, and a test asserts the silence so that moving the check fails a test
rather than CI.

The consequence check cost nothing against history, which was measured rather than assumed:
every proper option in the trail already carries one, and `gate verify` stays at exit 0.

For anyone writing a question: carry a consequence on each option you offer, recommend exactly
one, and leave the free entry bare. `no_options: true` is unaffected — a question with nothing
to offer has nothing to recommend.

The third requirement still has no reader and no field. The reason for the recommendation can
only live inside the question's text or inside a consequence, which is #247, and that nothing
says questions are put one at a time is #248. Both need a change to a document the agent cannot
edit.

<!-- xeno:section:residual-risk -->
## Residual risk

The asymmetry is the whole of the risk and it is permanent until somebody decides otherwise. A
hand-written artifact carrying a question that recommends nothing passes every gate, so the
guarantee holds for questions written through the command and not for the artifact. That is
narrower than before — two of three requirements now have a gate reader where one did — and it
is the same shape #242 left behind for the review checklist, which is worth noticing as a
pattern rather than twice as a coincidence: when a check cannot go in the gate, it goes in the
writer, and the artifact stops being the thing that is guaranteed.

The justification is a fact about today's trail and the test asserts the code. If XENO-3's Q-2
were released, or the intent closed differently, the recommendation check could move into the
gate and nothing would prompt anyone to try;
`TestTheGateIsSilentOnTheRecommendation` would then fail and read as an obstacle rather than as
an invitation. Its comment says what to check before trusting the failure, which is the best a
test can do about a reason that lives outside it.

A fixture had carried the forbidden shape since WP5 and nothing noticed for five work packages.
What found it was a stricter gate, not a reader, and the same could be true of the other
exported shape functions: `EvidenceShape` and `DecisionShape` are exported for a writer and
tested only through it, in another package, and nobody has looked. Named in P4's gaps rather
than guessed at here.

#247 is the uncomfortable one. The clause's third requirement is unreadable and also unwritable,
so every recommendation in this trail either carries its reason in prose or carries none and
nothing can tell which. This intent improves the clause's coverage while leaving its most
load-bearing third in the worst state a requirement can be in, and a reader of
`docs/clause-readers.md` should see section 8 as improved rather than closed.

What is not a risk: the 375 pre-existing verdicts, which `gate verify` confirms at exit 0; the
`askable` fixture, which passes unchanged; and every existing caller of `QuestionShape`, which
keeps its name, signature and behaviour on everything history contains.
