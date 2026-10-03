---
intent: github.com/triplem/xeno#176
phase: 01-requirements
created: "2026-10-03T11:34:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a8207c073efa599bf222b3e706517cedbea311307d9f31eef834c306961594aa
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

**A lock written from now on carries a size per file.** Every entry of `files` has `bytes` beside
`path` and `sha256`, taken from the same walk that took the hash, so the two describe the same read
of the same file.

**The byte budget is judged against the recorded sizes.** A file that grows after a phase is sealed
does not change that phase's verdict, which is the property the specification clause was written
for and the only reason this intent exists.

**A lock that records no sizes produces no byte finding.** Not a finding against zero, and not a
pass claiming the context was inside its budget: seventy-nine intents' worth of locks record
nothing, and reading them as zero-byte contexts would be a quiet pass where there used to be a
loud one.

**A lock that records some sizes and not others is judged on what it has**, with nothing invented
for the entries that have none. That is not a state the runner can produce — it writes all or
nothing — but it is a state a hand-written lock can be in, and the check should not guess.

**The file-count budget is unchanged**, because it was already judged against the lock.

**`gate verify` stays at exit 0 over the trail.** Every lock behind this intent records no sizes,
so the byte budget has nothing to say about any of them, which is what the third criterion is for.

**Nothing else moves.** The base's contents and order are unchanged, the hash of a file is
unchanged, no rule set changes, and the suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No backfill.** Sealed locks keep what they have. A lock rewritten to carry a number nobody
recorded would describe a reading that never happened, and `artifacts_hash` covers it.

**No second use of the size.** It is summed for the budget and read by nothing else. The hash is
what detects change, and a size that also answered that question would be two answers to one
question.

**No change to the file budget.** Already judged against the lock.

**No measurement anywhere.** After this intent nothing in the gate path calls `os.Stat` for a
budget. If a reader wants to know what the context is now rather than what it was, that is a
different question and nothing in this process asks it.

**No profile for this repository.** Next, as the experiment.

<!-- xeno:section:constraints -->
## Constraints

**The specification clause is one commit old and it is the authority.** "The size is recorded
because the budget is judged against what the phase was given and not against what the tree holds
now." The code follows that sentence and nothing in this intent reinterprets it.

**Absent is not empty.** A lock with no sizes is a lock that was written before the field existed,
and the check has to be silent about it. This is A74's rule and the register closure's rule, and it
is the one way this intent could quietly make the trail look compliant.

**The lock is written once and never refreshed**, so the size is taken at `phase start` with the
hash and nothing updates it afterwards.

**The base must not change.** Which files, and in what order, is #171's design.

**No register row.** The register closed with M0; the one decision here is recorded in the design
phase.

**Short artifacts.** One field, one sum, four tests.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0.**

**One intent, one branch, one issue** — `176-the-lock-records-no-size`, #176, labelled wp8 — and
one commit before it that is not an intent, which is the specification change the first standing
rule orders first.
