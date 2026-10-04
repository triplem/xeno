---
intent: github.com/triplem/xeno#221
phase: 03-implementation
created: "2026-10-04T09:39:17Z"
schema_version: "1.0"
runner_version: dev+24becc3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d001d5c46f78496c10bb57a7ca07db2b99176f67c91de60c44f34fe68a5d83ad
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

**The fifteen corrections, first.** One word in fifteen `evidence/attached.yaml` files:
XENO-0107, XENO-0108, XENO-0111, XENO-0121 and XENO-0200 to XENO-0210. All fifteen were
read first and are identical in shape — one entry, `test-report`/`go-test`, `state:
attached`, a `sha256`, `path: evidence/go-test.txt`, `pipeline: local` and a commit — so
each has exactly one `result:` line and the edit is anchored on the whole line. Fifteen
files, fifteen insertions, fifteen deletions. Nothing else in them is touched, and
`pipeline: local` stays because it is accurate.

**`internal/gates/gates.go`.** `attachedResult` judges a recorded value against
`model.EvidenceResults` and returns the finding on `evidence/attached.yaml`; `evidence`
calls it in the branch where an attachment has been resolved. A value is judged and an
absence is not, and the comment says which entry of this repository's own pipeline that
asymmetry protects. The wording follows `EvidenceShape`'s for the same rule on the
declaration side.

**`internal/evidence/attach.go`.** `unrecordable` answers why an entry cannot be written
down at all, as against `unbindable`, which answers what would bind one. It runs for
every entry and before the binding checks, because a result is wrong whichever way an
entry is bound and because it is the one a publisher fixes in the step that wrote it.
`Result.Declined` is what `Result.Unbindable` was.

**`internal/runner/runner.go` and `cmd/xeno/main.go`,** the two callers that read that
field. Both had a remedy of their own — "republish with a sha256" — which was right for
one of the three reasons the moment there were three. Each reason sentence now carries
its own remedy and the callers print it: `phase start` refuses naming the entries and
says that waiting will not help, and `evidence attach` names them and exits 1 with no
line of advice of its own underneath.

**Tests.** Three in `internal/evidence/attach_test.go`: a result outside the set is
declined, names the entry and the value, says what to republish, and leaves neither a
record nor a copied report behind; an entry whose producer reports nothing is still
recorded, with `other/trivy-db` named as the real case; and every value the document
defines is recorded as published, `fail` included. Three in
`internal/gates/evidence_test.go`: an attachment carrying `success` is a finding on
`attached.yaml`, every documented value passes, and an attachment with no result at all
is not a finding. The sets come from `testdata/evidence-declaration.yaml` rather than
from the constants, which is that file's own reason for existing.

**One test string.** `TestStartNamesAnEntryThePipelinePublishedWrong` asserts the advice
is in the refusal; the sentence was reworded, so the expected substring follows it.

<!-- xeno:section:deviations -->
## Deviations from the design

**`Result.Unbindable` is renamed to `Result.Declined`, which the design did not call
for.** The design said the refusal would be "a third reason beside the two it already
has", which implied the field stays as it is. Writing it showed the name was the reason
the callers were wrong: `Unbindable` describes two of the three cases, and an entry
whose result is a word section 4 does not define is recordable and declined rather than
unbindable. Both callers had written their own remedy underneath the list — "republish
with a sha256" — which is right for one reason out of three, and a field whose name says
one thing invites exactly that. So the field says what the list is, the sentences carry
their own remedies, and the callers print them. It is an internal rename with two call
sites and four assertions in one test file; the review checklist carries it as the
migration note.

**One expected substring in `internal/runner/runner_test.go` changed case.** The
refusal's closing sentence was reworded from "Republish it with a sha256, or declare it
differently; waiting will not help." to "Waiting will not help; republish it as each
line says.", so the test that asserts the advice is present follows the wording. The
assertion is unchanged in what it checks.

**The refusal could not be transcribed against this repository's own P4 at the time it
was written.** `xeno evidence attach` against a hand written manifest exited 2 with
"open .xeno/intents/XENO-0244/phases/04-verification/output.md: no such file or
directory", which is the error #208 was filed about: the phase had no artifact yet
because its sections had not been written. The transcript is therefore taken in the
verification phase, where the artifact exists, rather than here. It is recorded because
the error is the issue's own opening quotation and meeting it again is worth a line.

**Nothing else.** The check's place, the finding's file, the wording taken from
`EvidenceShape`, judging a value and not an absence, the decline rather than an error,
the order of the two steps, and the hand edit of the fifteen are as the design decided
them.
