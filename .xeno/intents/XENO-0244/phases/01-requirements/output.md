---
intent: github.com/triplem/xeno#221
phase: 01-requirements
created: "2026-10-04T09:28:24Z"
schema_version: "1.0"
runner_version: dev+24becc3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6784c959221bf4600bd84653f30aee81a6ad0262663b878af724d215c9c76a15
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      resolves: Q-1
      chosen: Correct the fifteen, then check unconditionally. success becomes pass in fifteen attached.yaml files, evidence attach declines a result outside the set at the point of writing, and the gate judges every attachment with no date in it.
      rationale: 'Three facts decide it. The file being corrected lies in a subdirectory no artifacts_hash covers, since DirHash does not descend and section 4 puts evidence/ outside the hash; the sha256 each attachment records is of the report rather than of the result word, and gate.yaml keeps no copy of the value, checked against XENO-0210''s verdict; and success and pass mean the same thing in the sentence the field stands in, so this is a vocabulary correction inside a closed set and not a change to what any run reported. With the fifteen reading pass the new check finds nothing in them, so gate verify stays at exit 0 and the gate needs no cutoff and nobody owes fifteen decisions. Against the alternatives: a date would be the first grandfather clause in this project and would leave G-Test reading success on the phases that have it; releasing them would cost fifteen decisions by a second person and leave fifteen open obligations for a defect with a four line fix; and the refusal alone would close the inflow while leaving the one file that can be edited without staling a verdict unchecked, which is the other half of the bargain #208 struck when it paid for attaching outside the artifacts_hash with the hash gate.yaml records. The cost is accepted and recorded: all fifteen read pipeline: local, so there is no pipeline to republish from and the trail carries a value a person corrected.'
      decided_by: triplem
      proposed_by: claude-opus-5
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**The fifteen read `pass`, and nothing else about them changes.** Fifteen
`evidence/attached.yaml` files, one word each. Every other field stays as it is,
including `pipeline: local`, every `sha256` still matches the report it points at, and
every path still resolves. The files are read before and after rather than edited by a
pattern nobody looked at.

**No hash moves and no verdict is contradicted.** `xeno gate verify` is at exit 0 over
the whole trail after the correction, with the same verdict count as before it, and the
`artifacts_hash` of each of the fifteen phases is unchanged. This is the criterion the
order of the work exists for, and it is recorded as a figure taken three times: before
the change, after the correction, and after the gate.

**An attachment carrying a result outside section 4's set is a finding.** In G-Evidence,
on `evidence/attached.yaml`, which is the file somebody would have had to edit for the
record to be in that state. The wording follows `EvidenceShape`'s for the same judgement
on the declaration side — the set named, and that the value is what the run reported
against its own threshold — rather than inventing a second phrasing for one rule.

**An absent result on an attachment is not a finding.** Section 4 requires the field on
`test-report` and `build-log` and has it absent where a producer reports nothing, and
the attachment inherits that: `other/trivy-db`, which this repository's pipeline
publishes without a result, must stay green. The check judges a value and never an
absence.

**`evidence attach` declines an entry whose result is outside the set.** It is counted
as pending, the entry is not written, and the reason is returned in `Unbindable` so that
`evidence attach`, `phase start` and `gate run` each say it in their own voice, which is
what that field exists for. The refusal applies to a file backed entry as well as to one
with a `uri`, because the result is read off the manifest in both cases.

**The declined entry leaves nothing behind.** No record in `attached.yaml`, no file
copied into `evidence/`, and the declaration stays pending so that a republished
manifest can bind it later.

**The exit code staircase is unchanged.** `xeno evidence attach` exits 1 where something
was declined, which it already does for the two reasons it has, and 0 where everything
bound. Nothing new appears on the staircase.

**The fifteen phases stay green throughout.** After the correction the new check finds
nothing in them, so their recomputed status equals their committed one and `Verify`
reports no divergence. Asserted over the real trail rather than only over a fixture,
because the fifteen are the reason this intent exists.

**A gate finding is produced where one is deserved.** A fixture whose attachment carries
a word outside the set is red on G-Evidence, which is the check failing for a real
reason, and the same fixture with `pass` is green. Both directions, by test.

**Nothing else changes.** No field of section 5, no new gate, no rule, no template, no
document, and no declaration anywhere in the trail. `evidence/` keeps every file it has.

**The usual gates of this repository.** `gofmt` and `go vet` silent, the suite green,
the build clean, prose at 88 columns, SPDX on any new file, and a test for the refusal,
for the gate finding, for the absent result staying green, and for the fifteen verifying
clean.

<!-- xeno:section:non-goals -->
## Non goals

**Requiring a result on an attached `test-report` or `build-log`.** G-Test and G-Build
read the field, which is why section 4 requires it on those two kinds, and an attachment
that arrives without one leaves the same hole one level down. It is not in D-1, no
attachment in the trail lacks a result, and widening the scope while correcting fifteen
sealed files is how a correction turns into a redesign. Recorded as a gap and filed as
its own issue.

**G-Test.** It is the reader that would have caught this and it is `not-implemented`.
Implementing it means deciding what a test report's result means for a phase's verdict,
which is a work package's argument and not this intent's.

**The `pipeline: local` provenance of the fifteen.** `evidence.source` for this
repository is `ci`, so the value records a regime the project does not configure. It is
nonetheless accurate about what happened, and correcting it would be rewriting the
history rather than its vocabulary. It stays, and #208's problem section is where it is
described.

**The hand written transcripts the fifteen point at.** Still what they are: a report
somebody typed, bound by a hash. The hash is correct and the content is what it is;
nothing here re-produces evidence that cannot be re-produced.

**The declarations in `output.md`.** Inside `artifacts_hash`, with verdicts over them.
They read `kind: test-report` and `job: go-test` and carry no result at all, which is
correct for a pending item, so there is nothing to correct in them even if they could be
touched.

**A migration, or any code that corrects attachments.** Fifteen files, one word, once. A
command for it would be a write path into sealed phases that exists forever to do
something that happens once, which is the opposite of what section 6 keeps narrow.

**Judging anything else in `attached.yaml`.** `state`, `pipeline` and `commit` are read
by no gate, deliberately, and section 4 says so. Adding a check for them because one is
being added for `result` would be the invented rule the standing rules forbid.

<!-- xeno:section:constraints -->
## Constraints

**The order is the decision.** The correction goes in before the check, in the same
commit, and the artifact says why: `Verify` compares a recomputed status against the
committed one, so a check that arrives first makes fifteen sealed phases divergent and
CI red on history. Nothing in this intent may reorder those two steps.

**What is sealed is never rewritten.** The one file being edited is the one section 4
puts outside the `artifacts_hash`, in a subdirectory `DirHash` does not descend into,
carrying a `sha256` of a report rather than of itself. That property is asserted by this
intent's own `gate verify` figures rather than assumed from the prose.

**The documents are not editable here.** Section 4 already closes the set and already
says what the attachment block carries. This intent adds a reader for a clause, which is
the shape #202's audit asked for, and changes no sentence.

**One authority on shape.** The gate's finding and the attach's refusal are one rule in
two places, and the places are different by necessity: the attach keeps a value out and
the gate reports one that got in. Both compare against `model.EvidenceResults` and
neither holds a list of its own.

**The refusal declines rather than fails.** `Attach` is called from three places, and an
error would stop a `phase start` that is in the middle of carrying a predecessor's
verdict forward. The existing mechanism is a decline with the reason returned and the
item left pending, and this is a third reason beside the two it has.

**Evidence comes from CI.** The refusal sits in the path a published manifest arrives
through, so a pipeline that publishes a word outside the set is told at the point of
writing rather than in a verdict afterwards. This intent's own P4 runs that path against
the four real manifests.

**One dependency.** Nothing new is imported. The set is `internal/model`'s, the finding
is `internal/gates`'s, the decline is `internal/evidence`'s.

**Prose at 88 columns, Go at whatever `gofmt` produces.** Both rules whose reader is a
person, which #211 wrote down knowingly.
