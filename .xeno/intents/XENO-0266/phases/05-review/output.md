---
intent: github.com/triplem/xeno#267
phase: 05-review
created: "2026-10-06T20:11:41Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3f824f5c03369fcccd5f297706b775cedacaa16c84df93300f44e98174c07de4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: deviation
      note: 'Two, both in P3. (1) docs/clause-readers.md got a paragraph where P1''s criterion 12 asked for three rows: the enumerated table is tool requirements, defined there as clauses the runner, a gate or CI can satisfy or violate, and none of the three new clauses can be violated by code because each states a consequence of the order phase start and scope set run in. They are explanations in that document''s four kinds, so the count moves 126 to 129 and a paragraph names them. Three rows would have claimed something fails if they are broken. (2) budget''s doc comment was corrected although it asserted nothing wrong, because its existing text about an unrecorded size being silence rather than zero bytes leads a reader with #267 in mind to conclude the recorded flag is what suppresses the check at P0, when the empty list is. Four wordings also changed on contact with the file and are in the specification commit''s message; no criterion was falsified.'
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: 'No interface changes. No artifact gains a field, no gate gains a check, no command changes its behaviour or its refusals, and the binary behaves identically: gate verify reports 475 verdicts verified before and after, and no test file is in the diff against main. What an adopter meets is three paragraphs of prose, so there is nothing to migrate and nothing to do. The one thing worth knowing is in the release notes rather than a migration note: a declared context budget is judged from P1 on and always has been, and section 5 now says so.'
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: None added, go.mod and go.sum are untouched, and nothing compiles differently. The diff is three paragraphs and one YAML comment in docs/process-definition.md, three doc comments across internal/runner and internal/gates, and one paragraph plus one number in docs/clause-readers.md.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** A normative document is touched, which the first standing rule bars
the agent from doing. The exception is the one the maintainer made on instruction: the wording was
drafted, the two alternatives were put with their costs, the decision was theirs, and "write them
as drafted" is what was acted on. The specification change is its own commit and comes before the
one that follows from it, which is the order the rule asks for even in the exception. It is the
third such exception in two days and is recorded as one in the commit message, in P0's problem
section and here, because an exception nobody counts is a habit.

Nothing is invented. No artifact gains a field, no gate gains a check, no command changes. The
three paragraphs describe the runner as it already is, which is the one shape of specification
change that cannot drift from the code.

The work belongs to WP8 and to XENO-0266, and both commits reference #267.

**What a reviewer should look at first.** Whether the three paragraphs are true. They are prose
about a mechanism, so nothing can fail if one of them is wrong, and the only reader is somebody
holding the document against `Start`, `ScopeSet` and the two checks. The claim is: `phase start`
resolves the lock from a scope that `scope set` cannot have written yet at P0, and nothing writes
the lock again. `runner.go:449` and `scope.go:65` are the two lines that make it so.

**Whether a paragraph was placed into an argument it splits.** Three placements, three sets of
neighbours, and this is the fault the last three specification commits each made once. Section 5's
lock subsection: the new paragraph follows the two it is a consequence of and precedes one that
begins a new subject. Section 5's budget clause: the new paragraph follows the clause complete,
and the clause itself was rewritten in XENO-0263, so it was read back whole rather than edited
into. Section 7: the new paragraph follows both limits and says it is not a third, and "Two limits
keep the check useful rather than noisy" is unchanged.

**Whether the deviation in `docs/clause-readers.md` is the right one.** P1 asked for three rows and
got a paragraph. The argument is in P3's deviations and it is about that document's own taxonomy: a
row in the enumerated table claims code can violate the clause, and none of these three can be.
A reviewer who disagrees should say so, because the alternative is three rows reading `nothing` in
the reader column, which is a different and also defensible answer.

**What is not here.** No test, because nothing changes behaviour; no `budget` result saying it did
not apply, because that is a value outside the four A4 and A42 fix; and P0's lock does not start
recording files, which is the decision #267 resolved and not an omission.

<!-- xeno:section:release-notes -->
## Release notes

**Nothing to adopt and one thing to know.** No binary behaviour changes, so a project that updates
gets the same verdicts on the same artifacts; `gate verify` reports 475 verdicts verified before
and after.

What a reader of the specification learns is that a context budget is judged from P1 on. A project
that declares one in `phases/00-intake/context-scope.yaml` has always had it judged against five
phases and never against the intake, and section 5 now says so where the clause is and gives the
reason where the lock is described. Nobody's figures move; what moves is whether the figure can be
read correctly.

The same holds for the staleness half of G-Freshness. An intake's reading was never compared
against the tree, and section 7 now says so, says it is not one of the two limits chosen against
noise, and names what escapes: a file that moved while the intake itself ran, since the next
phase's lock records it as it was by then.

**For somebody writing a tool against the trail.** `files` is absent from every P0
`context.lock.yaml` — 0 of 122 on this branch — and absent from 49 locks at each later phase,
which are the intents that ran before `scope set` existed. The schema block in section 5 now
carries that warning on the `files` line itself.

<!-- xeno:section:residual-risk -->
## Residual risk

**The three paragraphs have no reader, and one of them could become false.** "An intake's `files`
is empty" is true because of the order `phase start` and `scope set` run in, and nothing in the
repository would notice if that order changed. The mitigation is phrasing and not a check: the
sentence states the fact rather than permitting it, so adopting either of #267's candidates means
deleting it, and the specification commit that opens the change is visible as a deletion.
`docs/clause-readers.md` records this as the one of the three that is reader-shaped and
deliberately unread.

**The explanation count in `docs/clause-readers.md` moved by hand, 126 to 129.** That is the fault
XENO-0260 and XENO-0264 each recorded about that same file — a prose figure beside the structure
it describes, with nothing checking that the two agree — and this intent adds a third instance
rather than fixing it. The fix is a check in the verify job and belongs to the learning those two
already filed. A reader who counts the explanations cannot: they were never enumerated, so the
figure rests on the pass of 2026-10-03 plus three.

**The intake is still not measured against its own budget**, which is the decision. What is not
known is the size of what is given up, because the honest measurement does not exist: resolving an
old intake's patterns today describes today's tree and not the one the phase was given, which is
the same reason section 5 records a size per file instead of measuring.

**A comment can still be true somewhere and false where it is read.** Two were, in two packages,
and were found by counting the trail rather than by reading either one. The learning recorded at
P3 proposes the convention and section 10 routes it through review, so nothing takes effect here;
a third instance would be found the same way, which is to say by accident.

**The exception to the first standing rule is now the third in two days.** Each is recorded, which
is the only control there is. The risk is not that a wording was wrong — four adjustments on
contact with the file were found and are written down — but that the recording stops being read.
