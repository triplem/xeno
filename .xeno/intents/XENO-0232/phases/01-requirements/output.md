---
intent: github.com/triplem/xeno#195
phase: 01-requirements
created: "2026-10-03T14:20:17Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+b54626e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ad789d1ae40c5aac79c835fa51791fdc0b18398d9e7100a226bfa4dbc9786b84
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

**The record's header is the runner's.** The four artifacts of one phase agree about which build
produced them, which is the whole point and is asserted as the agreement rather than against a
literal: the record's `runner_version` is compared with the `gate.yaml` the same binary wrote beside
it.

**The four keys are arguments.** `--category`, `--observation`, `--proposal`, `--target`, which are
section 10's four and no fifth.

**`--no-finding` writes the empty record**, with the same header, because section 10 calls it the
honest empty case and a file left out is a G-Learning finding.

**Entries accumulate.** A second call appends and the order is the order they were recorded in. A
phase that learned two things says so twice.

**Without `--phase` the record is the intent's own**, which is what `intent close` reads, and it
carries no `phase` field.

**A category outside section 10's four is refused**, with the set named, before anything is read.

**An entry missing any of its keys is refused**, naming the flag, and whitespace is not a value.

**`--no-finding` beside an entry is refused**, and an entry on a record that states `no_finding` is
refused, and `--no-finding` on a record that carries entries is refused. Either way one of the two
statements is already on disk, and a writer that resolved it silently would be choosing which was
meant.

**A refused call leaves no file.** Everything is checked before anything is written.

**G-Learning passes on what the command writes**, unchanged. The gate is right and was never the
gap.

**A hand-written record still works.** Thirty intents' worth exist, sealed.

**Nothing else changes.** No gate, no rule, no field of section 5, no existing artifact.
`gate verify` stays at exit 0 and the suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No content.** No template for an observation, no default proposal, no suggested category. Section
10 routes a learning through a merge request against the rule set so that it takes effect after
review; a runner that drafted the text would be proposing rules, and the review would be of its
proposal rather than of the phase's.

**No change to G-Learning.** It checks the header fields, the two permitted body keys, the four
entry keys and the closed category set, and all of that stays. The writer refuses the same things
earlier, which is where a refusal is cheaper, and the gate remains the thing that judges a record
nobody wrote with the command.

**No `plugin_version`.** A constant in every artifact, including the three the runner wrote, and it
does not match the manifest the plugin now carries. #177.

**No backfill.** The hand-written records are sealed inside `artifacts_hash`.

**No sealed-phase guard.** `SectionSet` has none and the reason holds: the write makes `verify`
report a divergence and `phase finish` judges the phase again, so refusing would refuse the first
half of a re-judgement.

**No second file and no new field.** The record is where section 10 puts it, with the keys section
10 gives it.

**No learning for the closing intent written automatically.** `intent close` reads the intent level
record and does not write one; what this adds is a writer somebody calls, not an inference.

<!-- xeno:section:constraints -->
## Constraints

**Section 10 fixes the shape and the gate already reads it.** Four categories, four keys, the empty
record. Nothing here may add a fifth key or a fifth category, and the writer's checks have to be the
same set the gate checks or the two disagree at the first addition.

**The content is the agent's and the header is the runner's.** That split is `section set`'s and is
reused rather than retaken. It is also what makes the command possible at all: an observation about
a phase can come from nothing but whoever did it.

**The record is inside `artifacts_hash`.** So the header is written at the moment of writing and not
carried over from a previous call — it describes the binary doing this write.

**Both levels, one writer.** Section 10 owes a record at the end of every phase and once more when
an intent closes, and `intent close` reads the second one. A command that served only the phase
would leave the intent level record with the problem this intent is about.

**Refuse before reading.** A category outside the four should never reach the file, so the argument
check comes before the record is opened and a refused call leaves nothing behind.

**A contradiction is refused, not resolved.** `no_finding` and an entry cannot both be true, and the
one already on disk is somebody's statement.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0.**

**One intent, one branch, one issue** — `195-learning-has-a-writer`, #195, labelled wp14 — based on
main and not on another branch.
