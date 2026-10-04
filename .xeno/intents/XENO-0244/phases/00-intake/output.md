---
intent: github.com/triplem/xeno#221
phase: 00-intake
created: "2026-10-04T09:26:39Z"
schema_version: "1.0"
runner_version: dev+24becc3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dc170d12d8252cf7f8c1b6b81fec391976cfd0330e069c7fdd282fb81b1b5ed6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
open_questions:
    - key: Q-1
      text: 'Fifteen attachments in the trail read result: success, which section 4''s closed set does not contain, and nothing judges an attachment''s result. A check added as it stands turns fifteen sealed phases divergent, because Verify compares a recomputed status against the committed one. What happens to the fifteen, and what applies to everything after them?'
      options:
        - text: Correct the fifteen, then check unconditionally. success becomes pass in fifteen attached.yaml files, evidence attach declines a result outside the set at the point of writing, and the gate judges every attachment with no date in it.
          consequence: 'gate verify stays at exit 0 throughout, because the file corrected is in a subdirectory no artifacts_hash covers and the recorded sha256 is of the report rather than of the word; the check is then uniform and nobody owes fifteen decisions. The cost is that it edits a record of a run nobody can reproduce, since all fifteen read pipeline: local, so the trail carries a value a person corrected rather than one a job emitted, and the intent has to say so in its deviations'
          recommended: true
        - text: Judge from a date. The check applies only to attachments written after a stated point and the fifteen stand as the historical record they are.
          consequence: nothing sealed is touched at all; the cost is the first rule in this project with a grandfather clause in it, a cutoff every later reader of the gate has to understand, and G-Test still reading success on those fifteen phases and finding neither value it knows
        - text: Release the fifteen. The check applies to everything and each finding is cleared with gate approve or gate override, naming a person and a reason.
          consequence: the process's own mechanism for a known defect is used and each release is attributable; the cost is fifteen decisions by a second person, the wrong value left standing in the file, and an override leaving an open obligation per phase, fifteen of them, which intent status then owes indefinitely
        - text: Add the refusal only, with no gate. Nothing judges what is already there.
          consequence: the smallest change and no verdict moves anywhere, and nothing written from now on can carry a word the section does not define; the cost is that the fifteen stay unjudged and attached.yaml, the one file in a judged phase that can be edited without making any verdict stale, still reaches G-Evidence unchecked
        - text: Something else, entered by the person deciding
          free: true
---

# Intake

<!-- xeno:section:problem -->
## Problem

Section 4 closes the set: an evidence `result` is `pass` or `fail`, and it is what the
run reported against its own threshold. G-Schema judges a **declaration** against that
set, through `EvidenceShape`, since #208. Nothing judges an **attachment**, because the
shape check reads the frontmatter of `output.md` and `evidence/attached.yaml` is a
different file — the one file in a judged phase that lies outside `artifacts_hash` and
can therefore be edited without making any verdict stale.

**Fifteen attachments in the trail read `result: success`.** XENO-0107, XENO-0108,
XENO-0111, XENO-0121 and XENO-0200 to XENO-0210, every one of them
`test-report`/`go-test` in `04-verification`, every one with `pipeline: local`, all
written by hand on both sides before any command existed to write either. `success` is
not a word any reader of this repository can act on.

**And a reader would not merely ignore it.** `build()` reads an attached item's `result`
and compares it against `"pass"`, so an attached `build-log` carrying `success` is
reported as a build that did not succeed: a green run turned into a red verdict by a
word. No `build-log` exists anywhere in the trail, so this has never fired, and G-Test —
which would read the fifteen that do exist — is `not-implemented`. The defect is
therefore invisible today and wrong in both directions the moment either reader has an
input. That is the same shape as the gap #208 closed, one level down: a value nothing
checks on the way in, in a field a gate reads.

**The way in is still open.** `evidence.Attach` copies `m.Result` from the manifest into
the record without looking at it. It already declines what it cannot bind — an entry
with a `uri` and no hash, an entry with neither — and returns the reason so each caller
says it in its own voice. A result outside the set is the same kind of defect and is not
declined.

**Why this was not fixed in the intent that found it.** `Verify` compares a recomputed
status against the committed one. A check added while the fifteen still read `success`
would recompute fifteen sealed phases into a status that disagrees with their verdict,
and CI would be red on history rather than on the change. So the trail's own age is what
makes the gate unaddable, and #208 added the `format` check in the same commit as the
field for exactly that reason — the only moment it is free. Three weeks separate the two
cases.

**What a correction can touch, and what it cannot.** `evidence/attached.yaml` lies in a
subdirectory, and `DirHash` computes `artifacts_hash` over the files lying directly in
the phase directory without descending, which is what section 4 says in prose:
`evidence/` sits outside the hash. The `sha256` each attachment records is the hash of
the report it points at, not of the result word beside it, and `gate.yaml` keeps no copy
of the value — XENO-0210's verdict carries its `artifacts_hash` and no occurrence of
`success` anywhere. So `success` can become `pass` in fifteen files without moving one
hash and without contradicting one recorded verdict, which is what makes the check
addable after it and not before.

<!-- xeno:section:scope -->
## Scope

**In scope.** Three things, in the order they have to happen.

**The fifteen are corrected.** `success` becomes `pass` in fifteen `attached.yaml`
files. It is a vocabulary correction inside a closed set, not a change to what any run
reported: `success` was typed by hand to say the run succeeded and `pass` is what
section 4 calls that. It touches no hash, as the problem section establishes, and it
goes first because every verdict stays reproducible only in that order.

**The gate judges an attachment's `result`.** In G-Evidence, where the attachment has
already been resolved, rather than in G-Schema, which reads one file's frontmatter. The
finding belongs to `attached.yaml`, which is the file somebody would have had to edit
for the record to be in that state, and that is already how G-Evidence places its
findings about an attachment.

**`evidence attach` declines a result outside the set.** The mechanism exists and this
is a third reason beside the two it already has: an entry it cannot bind is declined,
not recorded, with the reason returned so the caller says it in its own voice. A value
nothing checks on the way in is how all fifteen got there, and the gate alone would
report them afterwards rather than keeping them out.

**Out of scope, each for its own reason.**

**Requiring a result on an attached `test-report` or `build-log`.** The declaration side
requires it on those two kinds, because G-Test and G-Build read it, and the same
argument applies one level down to an attachment that arrives without one. It is not in
the decision this intent carries out, no attachment in the trail lacks a result, and
extending the scope while correcting fifteen sealed files is how a correction becomes a
redesign. Recorded as a gap and filed.

**G-Test.** Still `not-implemented`, and the reader that would have caught this. Giving
it an implementation is a work package's worth of decisions about what a test report's
result means for a phase, and it is not this intent.

**The `pipeline: local` provenance of the fifteen.** It records a regime this repository
does not configure, `evidence.source` being `ci`. The value is accurate about what
happened and correcting it would be rewriting history rather than its vocabulary. It
stays.

**Rewriting the fifteen declarations.** They sit in `output.md` inside `artifacts_hash`
with verdicts over them. Nothing here touches them; the correction is in the file no
hash covers, which is the whole reason this is doable at all.

**Anything that makes a phase declare evidence.** The first residual risk of XENO-0243
and still true.

**What this intent owes its own trail.** `gate verify` at exit 0 before the change,
after the correction and after the gate, over the same number of verdicts each time,
recorded as three figures rather than one claim. Its own P4 declares the four pending
items the workflows publish, so the new refusal path runs against a real manifest on the
way in.

<!-- xeno:section:context-rationale -->
## Why this context

#221 is read as it stands, with the decision recorded in its comment thread, which is
this intent's own Q-1 and D-1: the question was put with four options and their
consequences, and a person chose correcting the fifteen before adding an unconditional
check. The reasoning behind that choice is reproduced in the problem section rather than
referenced, because an issue comment is not an artifact and the trail has to carry why.

Section 4's evidence subsection, again and for a different sentence than #208 needed:
the two closed sets, which kinds carry a result, that `result` is the job's own verdict
and never a statement by Xeno, that `evidence/` sits outside the `artifacts_hash`, and
what the attachment block carries. The last of those is the specification of the record
being corrected.

Appendix B is read for the one property the correction rests on: `artifacts_hash` is
computed over the files lying directly in the phase directory, so a subdirectory is
outside it. `internal/hashing/hashing.go:DirHash` is read beside it, because a prose
reading of Appendix B is not enough to touch fifteen sealed phases on.

`internal/gates/gates.go` is read for `evidence`, which resolves a declaration against
its attachment and is where the check goes; for `Collect`, which pairs the two by kind
and job; for `build`, which is the reader that would turn `success` into a red verdict
and is the evidence that this is not cosmetic; and for `EvidenceShape`, which is the
same judgement on the declaration side and whose wording the new finding should match
rather than invent.

`internal/evidence/attach.go` is read for `Attach` and `unbindable`: where the result is
copied from the manifest, and the shape a refusal takes there — a decline with a reason
returned, counted as pending, rather than an error that stops the run.

`internal/runner/runner.go:Verify` is read once more for the comparison that made this
expensive: committed status against recomputed status, which is why the order of the two
steps is the whole decision.

The fifteen files themselves are read, all of them, to establish that the only value in
them outside the set is this one word and that every one is `test-report`/`go-test` with
a hash and a path that still resolve. A correction applied to a file nobody read is how
a vocabulary fix becomes a data loss.

`xeno.yml`, `semgrep.yml`, `trivy.yml` and `audit.yml` are read for what their manifests
publish, because this intent's own P4 declares those four items and the new refusal sits
in the path they arrive through.

Nothing outside the repository is needed. #221's comment thread, two sections of the
process definition, Appendix B and four source files are the base.
