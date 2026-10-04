---
intent: github.com/triplem/xeno#221
phase: 05-review
created: "2026-10-04T10:58:30Z"
schema_version: "1.0"
runner_version: dev+24becc3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a7b87bcb8345b652e04680abc3344d171e4255dc9b61ee9724f6d13e4e2cc338
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      Three, each naming what it departs from: the rename of Result.Unbindable names the
      design sentence that called the refusal a third reason beside the two it already has,
      and says what writing it showed; the changed substring in a runner test names the
      reworded refusal it follows; and the refusal that could not be transcribed against this
      intent's own P4 names #208's opening quotation, which it reproduced exactly.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      One exported field renamed, evidence.Result.Unbindable to Declined, with two callers in
      this module and four assertions in one test file, all in the same commit. The package is
      internal/, so there is nobody outside this repository to migrate. What a reader has to
      know is the behaviour beside it: an out-of-set result is declined at the attach and
      judged on an attachment, and the fifteen records that would have produced those findings
      were corrected in the same commit, which is the migration and is in the diff.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      go.yaml.in/yaml/v3 is still the one dependency. strings was added to one import block and
      is the standard library, as errors was in XENO-0242.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Three, each naming what it departs from. The
rename of `Result.Unbindable` names the design sentence that called the refusal "a third
reason beside the two it already has", which implied the field stays, and says what
writing it showed: the name described two of the three cases and both callers had a
remedy underneath the list that was right for one of them. The changed substring in
`TestStartNamesAnEntryThePipelinePublished Wrong` names the reworded refusal it follows.
And the refusal that could not be transcribed against this intent's own P4 names #208's
opening quotation, which it reproduced exactly.

**`interface-change-needs-a-migration-note` — met.** One exported field is renamed,
`evidence.Result.Unbindable` to `Declined`, with two callers in this module and four
assertions in one test file, all updated in the same commit. The package is `internal/`,
so nothing outside this repository can depend on it and there is nobody to migrate. What
a reader of this repository has to know is the behaviour beside it: an entry whose
`result` is outside section 4's set is now declined at the attach rather than recorded,
and an attachment already carrying such a value is a G-Evidence finding. Both are new
findings for a tree that has none, because the fifteen that would have produced them
were corrected in the same commit — which is the migration, and it is in the diff rather
than in a note.

**`new-dependency-needs-a-rationale` — not applicable.** `go.yaml.in/yaml/v3` is still
the one dependency. `strings` was added to one import block and is the standard library,
as `errors` was in XENO-0242.

**What the checklist cannot answer, and the review did.** Whether the one claim this
intent rests on is true: that no `artifacts_hash` covers `evidence/attached.yaml`. It
was read out of `DirHash` in the intake, not out of the prose, and then measured three
times — `gate verify` at exit 0 over 342 verdicts before the change, after the fifteen
corrections, and after the gate. The middle reading is the review's own check, because a
wrong reading would have shown fifteen divergences there and nowhere else. After the
attach the trail stands at 344 verdicts, still exit 0.

<!-- xeno:section:release-notes -->
## Release notes

**An attachment's `result` is judged.** Section 4 closes the set to `pass` and `fail`;
G-Schema has judged a declaration against it since #208, and nothing judged the record
the pipeline's result actually lands in. `evidence/attached.yaml` is the one file in a
judged phase that lies outside the `artifacts_hash`, so what it says is checked by a
gate or by nothing. G-Evidence now reports a value the document does not define, on that
file, in the words the declaration side already uses.

**And it is kept out at the point of writing.** `xeno evidence attach` declines an entry
whose result is outside the set, names it, says what to republish and exits 1, writing
neither a record nor a copy of the report. That is how all fifteen of this repository's
own attachments came to read `success`: nothing looked at the value on the way in.

**The fifteen are corrected.** `success` became `pass` in XENO-0107, XENO-0108,
XENO-0111, XENO-0121 and XENO-0200 to XENO-0210. A vocabulary correction inside a closed
set, in the one file no hash covers: `success` was typed by hand to say the run
succeeded and `pass` is what section 4 calls that. Everything else about those records
stands, `pipeline: local` included, because it is accurate about what happened.

**The order is the change.** `Verify` compares a recomputed status against the committed
one, so the check could not go in first without reporting fifteen sealed phases as
divergent. Both steps are in one commit, the correction first, and `gate verify` is at
exit 0 over the same 342 verdicts before the change, after the correction and after the
gate.

**For a reader of a verdict.** Both files that carry an evidence item are now judged:
the declaration in `output.md` by G-Schema, and the record in `attached.yaml` by
G-Evidence. Before this, the half of an item that a pipeline writes and a person can
edit without staling any verdict was the half nothing read.

**One name is more honest.** `evidence.Result.Unbindable` is `Declined`, because two of
its three reasons are about what would bind an entry and the third is about what may be
written down at all. Each reason now carries its own remedy, so the two callers no
longer offer advice that fits one case in three.

<!-- xeno:section:residual-risk -->
## Residual risk

**An attachment with no result at all, on a kind that needs one, is still unjudged.**
D-1's named exclusion, and the first thing anybody reading this should know: section 4
requires `result` on `test-report` and `build-log` because G-Test and G-Build read it,
and an attachment arriving without one leaves the same hole this intent closed one level
up. It is filed. No attachment in the trail lacks a result, so the hole is reachable
only by a pipeline that publishes one without.

**G-Test is still `not-implemented`.** It is the reader that would have caught the
fifteen three weeks ago. What this intent changes for it is that the value it will read
is now a word the document defines; what it does not change is that nobody reads it yet.

**The correction is a precedent for a hand edit.** Fifteen files, one word, read before
and after, with three readings of `gate verify` as the only safeguard. The next
vocabulary correction inside `evidence/` will be done the same way, by hand, and the
argument against a command for it — a permanent write path into sealed phases for
something that happens once — is also the argument that leaves nothing to catch a
mistake except that figure.

**`gate run` is the third caller of `Attach` and says nothing about a declined entry.**
`phase start` refuses and names it, `evidence attach` prints it and exits 1, and `gate
run` reports a provisional verdict with no reason beside it. The verdict is correct and
the reason is missing, which is one line of output and was out of scope.

**Nothing re-reads a manifest that was published wrong and later corrected.** The
declaration stays pending and somebody runs the attach again, by hand, with `gh run
download` in front of it. That is the same manual step the whole pull path has, and #225
names the fetch adapter's absence.

**The fifteen still point at transcripts somebody typed.** Bound by hashes that are
correct about bytes a person wrote. This intent corrected what they say about the run's
outcome and not what they are, and `evidence.source` for this repository reads `ci`
while all fifteen read `pipeline: local`. Accurate, and left alone.

**Two gates now read the same closed set in two places.** `EvidenceShape` for a
declaration and `attachedResult` for an attachment, deliberately not one function,
because the declaration check carries an exemption for a pending item and the attachment
check must not. The shared thing is `model.EvidenceResults`, and the test that keeps
both honest reads the set from `testdata/evidence-declaration.yaml` rather than from the
constant. If a third reader of that set appears, the argument for leaving them separate
should be re-read rather than repeated.
