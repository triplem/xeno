---
intent: github.com/triplem/xeno#221
phase: 02-design
created: "2026-10-04T09:29:36Z"
schema_version: "1.0"
runner_version: dev+24becc3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b33f0b251d52f99ef3f1a7ce158db071a54c72d467f914e1ebd75423b9e80983
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The check goes in G-Evidence, not in G-Schema.** G-Schema judges the shape of one
file's frontmatter, which is where a declaration lives; an attachment lives in another
file and is only meaningful once it has been paired with the declaration it binds.
`Collect` does that pairing and `evidence` is the gate that reads the result of it, so
the check sits where the record is already resolved. Putting it in G-Schema would mean a
second reader of `attached.yaml` in a gate that has never opened that file.

**The finding is placed on `attached.yaml`.** `evidence` already distinguishes the two
files it reports about — a declaration's failure is `output.md`'s and an attachment's is
`attached.yaml`'s, "the file somebody would have had to edit for it to be in that state"
— and the new finding follows that sentence rather than restating it.

**The wording is `EvidenceShape`'s.** The same rule on the declaration side says the
kind and job of the item, that section 4 fixes the set, names its two values, and adds
that the value is what the run reported against its own threshold. One rule reported in
two places has to read the same in both, or a reader concludes they are two rules.

**A value is judged and an absence is not.** Section 4 requires `result` on
`test-report` and `build-log` and has it absent where the producer reports nothing. This
repository's vulnerability scan publishes exactly such an entry, `other/trivy-db` with a
file and no result, so a check that demanded one would turn a conformant pipeline into a
finding. The asymmetry is the same one `EvidenceShape` carries and it is deliberate on
both sides.

**The refusal in `Attach` declines rather than errors.** `Attach` is called from `phase
start`, `gate run` and `evidence attach`. An error returned would abort a `phase start`
that is in the middle of carrying a predecessor's verdict forward, so the mechanism
already there is the right one: the entry is not written, it is counted as pending, and
the sentence goes into `Unbindable` for the caller to print. This is a third reason
beside the two that field was built for, which is why no new mechanism appears.

**The result check applies before the two binding checks.** `unbindable` today answers
only for an entry with no file, because a file backed entry is bound by the bytes that
were copied. A result outside the set is wrong whichever way the entry is bound, so the
check is lifted out and runs for every entry, with the uri and hash reasons left where
they are. Order matters in the message a person sees: an entry with a bad result and no
hash is told about the result, because that is the one the publisher can fix in the step
that wrote it.

**Nothing re-reads the manifest to validate it separately.** A validating pass over
`manifest.yaml` before the binding loop would be a second place that knows what an entry
must look like. The check is where the value is copied from the manifest into the
record.

**The correction of the fifteen is a hand edit, recorded as one.** One word in fifteen
files in a directory outside every hash, done once, with the files read before and
after. The alternative was a command, and a write path into sealed phases that exists
forever to do something that happens once is the opposite of what section 6 keeps
narrow. The deviations section of the implementation phase names the fifteen keys and
the figures that prove nothing else moved.

**The order of the two steps is in the commit, not only in the prose.** The correction
and the gate land together, so no commit exists in which the check is present and the
fifteen still read `success`. A reader bisecting the history never meets a tree where
`gate verify` reports fifteen divergences.

<!-- xeno:section:alternatives -->
## Alternatives

**The check in G-Schema, beside the declaration's.** It would put one rule in one
function, which is the obvious shape. Rejected because G-Schema reads frontmatter and
the attachment is a different file whose meaning depends on the declaration it pairs
with; the gate that pairs them already exists and already reports about both files.

**A shared function called from both gates.** Tempting while the rule is one comparison
against one set. Rejected because the two are not the same judgement: `EvidenceShape`
reads a declaration that may legitimately be pending and exempt, and the attachment
check reads a record whose whole purpose is to carry the run's verdict. A helper that
both call would have to take the exemption as a parameter, which is two rules in one
function rather than one rule in two places. The set they compare against is the shared
thing, and it is already shared.

**Requiring `result` on an attached `test-report` and `build-log` as well.** The
strongest of the rejected options and the one that is filed rather than dropped: it is
the same clause one level down, G-Test and G-Build read the field, and today nothing
would notice an attachment arriving without it. Out of scope by D-1, and extending a
correction of fifteen sealed files with a criterion nobody asked for is how the order
this intent depends on gets lost.

**Correcting the fifteen with a command.** `xeno evidence repair`, or a flag on
something. Rejected: fifteen files, one word, once, and the command would outlive the
need by the whole life of the project while being a write path into phases that are
sealed.

**Rewriting the fifteen as `fail` where the run had actually failed.** Nothing suggests
any of them had. Every one records a local `go test` run that passed, and reading
`success` as anything but `pass` would be inventing a fact about a run nobody can
re-produce.

**Leaving the fifteen and judging from a date.** Q-1's second option. Rejected by D-1,
and worth naming here for the reason the decision gives: a cutoff in a gate is a rule
every later reader has to understand, and it would leave `G-Test` reading `success` on
exactly the phases that have it.

**Adding the refusal only.** Q-1's fourth option, and the cheapest. Rejected by D-1
because `attached.yaml` is the one file in a judged phase that can be edited without
making any verdict stale, so the gate reading it is the half that cannot be skipped.
#208 paid for attaching outside the `artifacts_hash` with the hash `gate.yaml` records;
this is the rest of that bargain.

**A test over the real trail only, rather than a fixture.** The fifteen are the reason
this intent exists, so asserting against them is tempting. Rejected as the only
assertion: a test that reads the repository it lives in passes for as long as nobody
changes the repository, and the gate has to be shown failing, which no state of this
trail may produce. Both are done — a fixture for the finding and its absence, and the
trail's `gate verify` as a figure.

<!-- xeno:section:impact -->
## Impact

**Three files, and fifteen one word corrections.** `internal/evidence/attach.go` gains
the result check and lifts `unbindable` out of its file-less branch.
`internal/gates/gates.go` gains four lines in `evidence`. `internal/model/model.go` is
untouched: `EvidenceResults` is already there and already shared, which is what makes
this a reader for a clause rather than a new rule.

**No verdict in the trail changes, and that is measured rather than argued.** `gate
verify` over 339 verdicts at exit 0 before the change, after the correction and after
the gate. The fifteen phases keep their `artifacts_hash`, because the file corrected
lies in a subdirectory `DirHash` does not descend into.

**One gate can now fail where it could not.** G-Evidence over a phase whose attachment
carries a word outside the set. The failure is new behaviour for a tree that does not
exist anywhere yet, which is the point of adding it in the same commit as the
correction.

**One command can now decline where it could not.** `evidence attach` exits 1 and names
the entry, as it already does for an entry it cannot bind. A pipeline publishing a bad
result learns it at the point of writing rather than from a verdict later, and the
declaration stays pending so a republished manifest can still bind it.

**What a reader of a verdict gains.** The two files that carry an evidence item are now
both judged: `output.md`'s declaration by G-Schema since #208, and `attached.yaml`'s
record by G-Evidence. Before this, the file outside the `artifacts_hash` was the only
part of a phase that could be changed without any gate noticing what it said.

**What nobody gains yet.** An attachment with no result at all on a kind that needs one
is still unjudged; G-Test is still `not-implemented`; and the fifteen still point at
transcripts somebody typed, bound by hashes that are correct about bytes a person wrote.

**Risk, and where it is paid.** The correction is a hand edit of fifteen committed
files. If the reading of Appendix B were wrong — if `artifacts_hash` did descend into
`evidence/` — `gate verify` would report fifteen divergences immediately and in CI,
before the branch could merge, which is the cheapest possible way for that mistake to
surface. The figure is taken three times for that reason, and the middle reading is the
one that proves it.
