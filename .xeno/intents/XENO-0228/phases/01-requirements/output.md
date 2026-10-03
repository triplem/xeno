---
intent: github.com/triplem/xeno#181
phase: 01-requirements
created: "2026-10-03T12:38:28Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6cbeac4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cc63c4f3640c6a10f79cb81d7b96825a959a6a72d884e7ed435ae67e5e5fe35f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**`section set --tool-version V` writes the field into `output.md`.** The runner writes the whole
frontmatter and nobody edits a file `artifacts_hash` covers.

**`phase finish` writes a digest carrying the same value, taken from `output.md`.** Not a second
argument: the digest and the artifact beside it are one phase of one session, and the value is
already there.

**A second `phase finish` that reports nothing keeps it.** This is the criterion the whole intent
turns on, because `writeDigest` is rerun by every finish and that rerun is where the hand edit was
paid again. A phase judged twice loses nothing.

**A reported version beats the recorded one.** The case is a digest written before any section was,
where there is no `output.md` to copy from; more generally the harness saying what it is cannot be
overruled by a file.

**A later `section set` that reports nothing does not erase it.** The field describes the session,
not the invocation, so reporting it on the first write of a phase is enough.

**No report is the field absent, in both files.** Neither writer invents one, G-Schema goes on
reporting it missing, and that is the state of every artifact written before this. A35 is satisfied
rather than contradicted.

**G-Schema passes on both artifacts of a phase written with the flag, with nothing hand edited.**
Demonstrated on this intent's own six phases, which is where the claim is cheapest to check and
hardest to fake.

**The six phase skills tell the agent to report it**, because a flag nobody knows about is a field
with no writer by another route.

**Nothing else changes.** No field of section 5, no gate, no rule, no existing artifact, and no
behaviour for a repository that never passes the flag. `gate verify` stays at exit 0 and the suite
stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No fourth `XENO_*` variable.** Section 7 enumerates three and the first standing rule makes an
addition a person's commit, made before the code. This intent writes no document.

**No entry point.** The three variables section 7 specifies stay unread. That is a WP11 gap and it
is larger than this field; naming it is this intent's job, closing it is not.

**No harness detection.** The runner does not learn which harness it is under, does not call one,
and does not read a client specific variable. Section 7: the moment the runner behaves differently
per harness, the tools stop being interchangeable.

**No `project.yaml` field.** An Appendix A addition, so a specification change, and wrong anyway —
a harness version is a property of a session and a project file would state it once and be stale on
the next upgrade.

**No default and no fallback.** Absent stays absent. A35's reason is the whole argument of this
intent and inverting it for convenience would be the one way to make the change worse than the hand
edit.

**No backfill.** The artifacts already written carry the value a person typed, which is correct and
sealed.

**No change to G-Schema.** The field was always required; what was missing was a way to supply it.

<!-- xeno:section:constraints -->
## Constraints

**A35 is the test, not an obstacle.** Its reason is that a plausible value in a field nobody
produced is worse than an absent one. A version the harness reports is produced, so the row is
satisfied; a version the runner worked out would contradict it. That distinction decides every
question in this intent.

**Section 7 forbids branching on the harness** and it is the reason the value is an input rather
than a lookup. It also designs the channel that would be better, which is why the flag is written as
an override of something that does not exist yet rather than as the final answer.

**The first standing rule bounds what the agent may do here.** The better mechanism needs a
specification change and a specification change is a person's commit. The question was put before
any code was written and answered, which is why there is no document commit in front of this one.

**The digest is rewritten by every `phase finish`**, so whatever supplies its frontmatter has to
work on the second run with no new input. That is why it copies.

**`SectionSet` carries frontmatter over**, so a value reported once in a phase survives, and the
flag must not erase what an earlier write recorded.

**One writer per field.** `output.md` is where the report lands and the digest copies from it, so
there is one source per phase and not two that can disagree.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0.**

**One intent, one branch, one issue** — `181-tool-version-has-a-writer`, #181, labelled wp11.
