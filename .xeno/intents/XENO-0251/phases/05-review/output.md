---
intent: github.com/triplem/xeno#212
phase: 05-review
created: "2026-10-05T16:13:38Z"
schema_version: "1.0"
runner_version: dev+97d07da.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2626f4f5aab3fcbe111faf706b38fb19cc6e681751d37697acd86f1ea228e4c2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'One against P2''s design, which said the two gates differ by a constant and a word: they differ by a constant and two strings, and a first pass shared one weaker repair phrase for both, downgrading G-Build''s ''fix the build'' with no test asserting it. One correction inside P3 rather than a departure from the plan: the pending test first used the sealed fixture, which supplies a sha256, so it asserted a sealed item with no result — a failure, not a pending one. Both are recorded with the fact that the code was right and the suite caught neither. One finding is left rather than absorbed: #229 enlarged a clause''s reader and did not update its row in docs/clause-readers.md, which belongs to that intent and not to this commit.'
      result: deviation
      rule: deviations-are-traceable
    - note: A verdict's meaning changes for every P4 written from here, and nothing migrates because no artifact changes. G-Test reported not-implemented and now reports pass, fail or pending; the 381 pre-existing verdicts are untouched and gate verify confirms it, because Verify compares artifacts_hash, status and gate-set membership rather than each check's result. What a dependant should know is that a green G-Test means the declared result was successful and says nothing about the mapping of acceptance criteria, which is the other half of section 7's row and has no reader. That is the migration note, and docs/clause-readers.md is where it lives permanently rather than in this commit message.
      result: deviation
      rule: interface-change-needs-a-migration-note
    - note: 'None added. go.mod is untouched; the change uses what gates.go already imports. No tool was weighed either, and the one thing that would need something new is a reader for the mapping half, which needs a numbering convention rather than a dependency and is #250.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: section 7 already asks G-Test for a
declared test result and this intent implements that, so the code moves towards the
specification and no specification commit precedes it. Nothing is invented — `TestKind` is the
spelling section 4 fixes, the gate list keeps its fourteen rows, and the two things that would
have needed an addition are #250 and the numbering convention it names. The branch carries one
intent, the commit references #212, and the issue carries `wp7`.

The acceptance criteria. Nine met, one met by the commit this phase precedes. Criterion 5 is the
one the design turns on and it is a command: `gate verify` at exit 0 over 384 verdicts.

The non-goals held. No mapping half. No numbering convention. No test suite run by a gate — the
suite was run by a person and its output declared, which is what section 7 describes. No second
gate. No change to what G-Build reads, and its wording is byte-identical after a first pass had
weakened it.

What a reviewer should weigh is the trade, not the code. G-Test went from `not-implemented` to a
result over half its row. That is a green verdict meaning less than it appears to, which this
repository names as a defect class, and it was chosen deliberately over a gate that catches
nothing. The mitigation is two rows in `docs/clause-readers.md` and it is a document read by a
person. P2's alternatives has the other option written out.

**G-Test judges this intent.** P4 declared `test-report/go-test` with the output of
`go test ./...`, bound by hash, and the sealed P4 verdict carries `G-Test: pass`. The gate that
had never judged anything in this trail judges the artifact of the intent that implements it.

What this intent got wrong is in P3's deviations: a first pass gave both gates one weaker repair
string, silently downgrading G-Build's message with no test to catch it, and a first version of
the pending test supplied a hash and so asserted a failure. The code was right both times and
the suite caught neither.

What is left is in P4's gaps, in #250, and in one finding about a previous intent: #229 enlarged
a clause's reader and did not update its row in the audit, which P3's learning proposes a
convention for.

<!-- xeno:section:release-notes -->
## Release notes

**G-Test judges something.** It was `notImplemented` in the gate table and a written-but-unjudged
verdict in every P4 of this trail, while `acceptance-criteria` and `test-mapping` were required
sections written in every intent because the plan says G-Test reads them.

It now reads the first half of section 7's row, "declared test result successful". A declared
`test-report` whose result is not `pass` is a finding naming the job; one awaiting its pipeline
is `pending`, not a failure; an artifact declaring none is green. That is G-Build's behaviour
over a different kind, and it is literally G-Build's code: `declaredResults` is one body with two
callers, because the two gates differ by a constant and two strings and two copies of a
comparison is one bet taken twice.

`TestKind` is a named constant pinned to `testdata/evidence-declaration.yaml`, so a rename in the
document fails the suite rather than making the new gate inert. That guard exists because G-Build
once compared against `build`, a spelling section 4 does not define, and matched no conformant
declaration for as long as nothing checked it.

**The other half of the row has no reader, and a green G-Test now means half of it.** Before this
the gate said `not-implemented`, which caught nothing and said so. The completeness of the
mapping needs an acceptance criterion to be identifiable, and nothing identifies one: 46 of 55 P1
artifacts carry no numbered criteria. `docs/clause-readers.md` carries both halves as two rows —
the only row in that table reporting a gate that judges part of what it is asked for — and #250
carries the decision.

Nothing already written changes. `gate verify` is at exit 0 over 384 verdicts, and the 381 that
existed are intact: `Verify` compares the artifacts hash, the status and the gate set rather than
each check's result, and the trail's eighteen test-report declarations are pending with attached
results, which the gate reads through the path G-Build already had.

For anyone writing a P4: declare the test report as you did before, and a result of `fail` is now
a red phase instead of a line nobody reads. This intent's own P4 declares one — `go test ./...`,
bound by hash — and its sealed verdict carries `G-Test: pass`, which is the gate judging the
intent that implements it.

<!-- xeno:section:residual-risk -->
## Residual risk

A green G-Test now means half of section 7's row, and no verdict can say which half. That is the
risk this change creates rather than inherits. Section 5 gives a check one result from a fixed
set, so there is no place in a verdict for "judged the declared result, did not judge the
mapping"; the only statement of it is two rows in `docs/clause-readers.md`, read by a person, in
a document that is a dated measurement. A reader who takes a green G-Test for section 7's row
being satisfied is now wrong where before the gate told them it had not judged. The maintainer
weighed exactly this against a gate that catches nothing, and P2's alternatives has the other
option written out.

Nothing requires a P4 to declare a test report, so the gate will pass most of this trail in
silence. It counts from the declarations; an artifact declaring none is green, which a test
asserts deliberately. Whether a P4 owes a test report is not something section 7 says, so this is
the limit of what the gate can mean rather than a defect in it — but it means "G-Test: pass" and
"nothing was tested" are indistinguishable in a verdict.

The mapping half may never have a mechanical reader, which #250 carries. A gate matching criterion
numbers in prose is close to the thing A90 warns about, and a `review` rule answered once per
intent may be the honest ceiling. If so, the templates go on carrying two required sections for a
clause that is half read forever, and the justification the plan gives for them stays half true.

`docs/clause-readers.md` is drifting in a way this intent demonstrates twice. It now carries
three late rows against a pass dated 2026-10-03, and line 84's clause about a question's options
is already stale because #229 enlarged its reader without touching the row. P3's learning
proposes the convention; nothing enforces it, and a measurement that keeps being amended by hand
is on its way to being a register that nobody decided to keep.

The evidence this intent declared is locally produced. `go test ./...` ran on this machine and
its output was bound by hash, which `--file` is for and which the model allows; nothing proves the
run happened in CI or against this commit. So the gate's first real judgement in this trail rests
on evidence weaker than the pipeline path the plan describes, and a reviewer who wants that
stronger should say so.

What is not a risk: the 381 pre-existing verdicts, confirmed at exit 0; G-Build, whose kind,
behaviour and finding text are unchanged and asserted; and the gate path's purity, since nothing
here runs a process or opens a socket.
