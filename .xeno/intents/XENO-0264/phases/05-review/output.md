---
intent: github.com/triplem/xeno#258
phase: 05-review
created: "2026-10-06T17:46:47Z"
schema_version: "1.0"
runner_version: dev+d3983d3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 42b599919f6c6a80b4f1ad118644ef0727caf5d97baf03f05f07a8b0b77ffc49
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: deviation
      note: 'Seven, in P3. (1) criterion 9 is not met: neither new check fires on this intent, because the template bump lands in the implementation phase and this intent''s P1 sealed at requirements@1.0.0 three phases earlier — the intent that introduces a forward-only anchor is structurally the one it cannot judge, which is the P3 learning. (2) seven existing tests failed on the new writer requirement and both fixtures gained a reason rather than the check being narrowed. (3) the recommended!=1 branch gained a continue P2 did not call for, so a question recommending nothing gets one finding and not two. (4) renderedSections and the parser the new checks needed were the same body, so parseSections now holds it and renderedSections delegates — a deduplication P2''s impact section did not list. (5) the clause table had thirty-nine rows against a prose figure of thirty-eight before this intent, found by counting while adding one; the paragraph names both numbers. (6) testReport''s comment named section 9, the Rule model, the same misattribution XENO-0260 corrected elsewhere. (7) the first mutation of templateAtLeast failed for the wrong reason, so a test asserting the claim directly was added and the mutation made precise.'
    - rule: interface-change-needs-a-migration-note
      result: deviation
      note: 'Two changes an adopter meets, and neither needs an action of theirs. The shipped templates go to 1.1.0, so a project vendoring the next release renders new P1 and P4 artifacts at that version and the two new checks begin to apply to them; artifacts already written declare 1.0.0 and are judged as they always were, which is the clause''s own anchor and is why nothing migrates. And model.Option gains Reason with omitempty, so no artifact already written changes shape and a question written by this runner now carries a reason where one is recommended. Two things worth a note rather than a migration: a project that forked the requirements or verification template under .xeno/config/templates/ keeps its own version and so opts out of both checks silently, which the per-id override in section 5 permits and nothing announces; and the plugin digest moves, so a new context.lock.yaml records a different plugin.sha256 from every sealed one, which G-Supply does not compare because it is not-implemented.'
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: 'None added and go.mod is untouched. One standard library import, strconv, for the numeric version comparison; the alternative was a string compare, which puts 1.10.0 below 1.9.0 and would have made the check stop applying at the tenth minor version with nothing to say so. No regexp or semver library: the refs are id@major.minor.patch and three Atoi calls answer the question.'
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched: both clauses were committed in
XENO-0262 and this is the code that follows them, which is the order the first standing rule asks
for and the first time this session that order has run forwards. Nothing is invented — every field,
version and check here was enumerated in section 5 or section 8 first, and the one thing that looked
like it needed inventing, a forward-only anchor for a gate's own checks, turned out to be the
template version every artifact already declares. The work belongs to #258 under `wp5`, on a branch
carrying one intent, and the commit references the issue without closing it: the issue also tracks
#248's clause, which closed without code, and #235's amendment has its own.

**A criterion is reported failed rather than reinterpreted.** Criterion 9 said both checks would fire
on this intent's own artifacts. They do not: the template bump lands in the implementation phase and
this intent's P1 sealed at `requirements@1.0.0` three phases earlier. The temptation was to read "this
intent" as "this branch" and call the scratch-copy run sufficient. P4 says fail, and the structural
reason is in P3's deviations: a convention carried by a shipped template reaches the next intent and
never the one that ships it.

**Every check was confirmed able to fail, and one mutation was redone because it passed for the wrong
reason.** Four mutations, four failures. The first attempt at the version comparison was crude enough
that the test failed on a malformed ref rather than on `1.10.0`, which is the thing A90 objects to
wearing a green badge; a test asserting the claim directly was added, including its own premise that a
string compare would order the two wrongly, and then the mutation was made precise.

**Both checks were proved able to pass as well as to fail.** A mapping naming all three criteria in a
table, and the same in prose, both leave G-Test green. Section 5 says the mapping "names each
criterion by its number" and does not say how, so a check that only ever went red would be reading a
clause nobody wrote.

**The cost of the bump was measured before it was accepted, and is written where a reader meets it.**
130 sealed artifacts lose the `strings_hash` recomputation; red at 1.0.0 and green at 1.1.0 on a
corrupted value; tampering still caught by `artifacts_hash` and the successor's freshness, exit 1. It
is in the code beside the consequence, in `docs/clause-readers.md`, in this intent's artifacts and in
the pull request, because the trail is not where somebody wondering why a sealed artifact stopped
being checked will look.

**Seven failing tests were read as the check working and not as tests to narrow.** The `askable`
fixture wrote a question whose recommended option carried no reason, which is precisely what section 8
now forbids. Both fixtures gained a reason. A reader of a diff touching seven test files should be
able to tell that from the alternative, and P3's deviations says which it was.

**Two figures were found wrong while being changed.** The clause table had thirty-nine rows against a
prose figure of thirty-eight, before this intent; and `testReport`'s comment named section 9 and
carried a superseded measurement. Both are corrected and the first names both numbers rather than
quietly writing forty — it is the gap XENO-0260 recorded about this very document a few hours earlier,
arriving in the document that catalogues such gaps.

**A duplicate was removed rather than added.** The section parser the new checks needed already
existed as `renderedSections`. One body now, with `renderedSections` delegating, so a section read by
a rule and a section read by a gate cannot come to mean different things.

**#235's code is deliberately not here.** `Advisory` on `model.Finding`, `result` and `budget` are
the other issue under the other work package, and both touch `gates.go` and `model.go` as this does —
which is a reason to keep them apart rather than to merge them.

<!-- xeno:section:release-notes -->
## Release notes

**Two clauses that nothing read now have readers.** `docs/clause-readers.md` reported "**no field to
carry it**; nothing reads it" for section 8's reason and "**nothing**" for G-Test's mapping half. Both
rows now name what would fail.

**The recommendation's reason is a field on the option.** `model.Option` gains `Reason` with
`omitempty`, and `QuestionAsked` refuses a recommended option that carries none. On the option and not
on the question, which is what section 8 now says and what makes the reason and the recommendation
unable to drift apart: a recommendation that moves elsewhere takes its reason with it or the writer
refuses. It is the writer's and not a gate's, because one sealed question in this trail already fails
the recommendation requirement and a gate reading either half would re-judge the trail (#229).

**An acceptance criterion is identifiable, and the mapping is judged against it.**
`.xeno/plugin/templates/requirements/template.yaml` and `.../verification/template.yaml` go to
`1.1.0`, one line each and nothing else. `gates.numberedCriteria` reads section 5's numbered list from
G-Schema, on the phase that can still fix it; `gates.mappingComplete` reads section 7's completeness
from G-Test, naming every criterion the mapping missed. **A green G-Test now means its whole row** for
an artifact at `verification@1.1.0`, where it has meant half of it since #212.

**Nothing in the trail is re-judged.** Both checks key on the version the artifact declares, compared
numerically so that `1.10.0` is later than `1.9.0`. All 69 P1 and 68 P4 artifacts declare 1.0.0,
`gate verify` is green at the same count, and an intent already under way keeps a P1 at 1.0.0 behind a
P4 at 1.1.0 without either check firing.

**The mapping check reads a number as a token and not a format.** A table cell, a list item and a
sentence all name a criterion, which is as much as section 5 says. The cost is that a stray number
satisfies it, so it can report a false green and never a false red — the right way round for a finding
that stops a phase.

## What the template bumps cost, measured

`hashes` treats a bumped template version as a bundle that is gone, so `recomputed` stops checking
`strings_hash`. **130 sealed P1 and P4 artifacts lose that reader**, and nothing announces it:
corrupting one sealed artifact's `strings_hash` gives G-Schema red at `requirements@1.0.0` and green
at 1.1.0, with `gate verify` reporting the same verdict count both times.

It is one of two readers. The same experiment had `gate verify` report `DIVERGENT` and exit 1, because
editing a sealed artifact moves its `artifacts_hash` and the successor's G-Freshness reads the
predecessor hash. There is no honest repair: a recorded hash that differs from the current bundle's
cannot be told apart from a right one about a bundle that has been replaced, and the two need opposite
answers. `docs/clause-readers.md` now carries the measurement.

## Also in here

`testReport`'s comment said the mapping half "has no reader", named section 9 — the Rule model — and
carried a figure superseded twice. Replaced. The clause table had thirty-nine rows against a prose
figure of thirty-eight before this intent; with the new row for section 5's numbered list it is forty,
and the paragraph says which was wrong. And the section parser the new checks needed already existed,
so there is one body rather than two.

**Not here:** #235's code, which is the other issue; and a gate for the reason, which #229 settled.

<!-- xeno:section:residual-risk -->
## Residual risk

**130 sealed artifacts lost a reader, permanently and quietly.** The `strings_hash` recomputation no
longer applies to any P1 or P4 artifact declaring 1.0.0. Measured, recorded in three places, and not
repairable without guessing. Tampering is still caught by `artifacts_hash` and the successor's
freshness, so what is lost is the field-level reader that names which field moved.

**Neither new check has judged a real artifact.** Criterion 9 is unmet: this intent's P1 sealed at
1.0.0 before the bump landed, so the first subject of either check is the next intent. They are proved
by thirteen tests and by a scratch intent, which is good evidence and not the trail they were written
for.

**The mapping check can report a false green.** A number anywhere in `test-mapping` as a token
satisfies it. Deliberate, because section 5 does not say how a criterion is named and a stricter
reading would fail a format the clause permits — but a mapping that mentions "3" in passing covers
criterion 3, and nothing distinguishes the two.

**Nothing requires the numbering to be coherent.** `1. 1. 1.` satisfies `numberedCriteria` and a
mapping naming `1` then satisfies completeness for all of them. Consecutiveness and uniqueness are not
in section 5 and the second standing rule forbids inventing them, so the gap is real and closing it is
a clause and not a commit.

**A hand-written question owes no reason.** `QuestionAsked` is the writer and no gate reads either
half of section 8's recommendation clause. Section 5 names G-Schema as the backstop for whoever writes
by hand; for this clause it is not one, and that is inherited from #229 rather than decided here.

**An adopter who forked the `requirements` template opts out silently.** The per-id override section 5
describes keeps whatever version a project's own template declares, so a fork at 1.0.0 gets neither
check and nothing says so.

**The plugin digest moved and nothing compares it.** Two template files changed, so a new
`context.lock.yaml` records a different `plugin.sha256` from every sealed one. G-Supply is
`not-implemented`; when it is, it will have to read the lock's recorded value rather than the current
tree's.

**`docs/clause-readers.md`'s count will drift again.** It is a number in prose beside the rows it
describes, with nothing checking that the two agree — which is how it came to say thirty-eight over
thirty-nine. This intent corrects it and adds a row, and the next addition has the same nothing
watching it. Twice recorded today, in XENO-0260 and here.
