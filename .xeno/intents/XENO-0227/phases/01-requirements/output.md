---
intent: github.com/triplem/xeno#179
phase: 01-requirements
created: "2026-10-03T12:14:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fb0dee9dfb504fd16659ab4ece640d32f23d88d149d27f6dff84f482aad36a73
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

**`xeno intent start --for ISSUE` writes an `intent.yaml` and nothing else.** The seven fields of
section 5 and no eighth: the qualified id, the key, `status: in-progress`, `created`,
`schema_version`, `runner_version` and `plugin_version`.

**The qualified id is the host, the repository and the key.** `--for 179` in this repository writes
`github.com/triplem/xeno#179`, with the repository from `tracker.project` and the host from
`tracker.base_url`.

**`--for` also takes a longer prefix of the id**, up to the whole thing, so a repository with no
tracker block can give `git.example/group/proj#4` and a repository with one can name another
repository on the same host. What a person leaves out is what the configuration holds.

**`created` is the runner's clock and the version fields are the runner's own**, so an intent's
`runner_version` reads the same string as the `gate.yaml` of its first phase. That is the half of
#177 this intent closes.

**The key continues the sequence on disk.** Run in this repository the command chooses `XENO-0227`,
because the highest key is `XENO-0226`, and it pads as that key is padded.

**`--intent` names the key instead**, which is what the first intent of a repository needs, since
there is no sequence to continue.

**Running it twice refuses rather than overwriting**, and the first intent's id survives the second
attempt. `intent.yaml` is inside `artifacts_hash`, so a second write would change a verdict that
named the first.

**Without `--for` it refuses and names the flag.** The issue is the one input, so its absence
cannot be defaulted.

**A hand-written `intent.yaml` still works.** Seventy-nine of them exist; `phase start` reads one
as before and this command does not touch them.

**`xeno intent status` lists the new intent like any other**, with its date and no problem
reported.

**The next-step suggestion for an intent that does not exist names the command.** It said no
command creates one, which this intent makes untrue.

**`gate verify` stays at exit 0 over the trail**, and the suite, `gofmt`, `go vet` stay green.

<!-- xeno:section:non-goals -->
## Non goals

**No `assumptions.yaml`.** Nothing refuses for its absence and not every intent here has one. A
command that wrote an empty register would be producing an artifact to satisfy section 4's list
rather than a reader, and `assumption record` already refuses with a sentence that names the file.

**No backfill and no rewriting.** The seventy-nine hand-written intents keep their fields, their
`0.1.0-dev` and their keys. Their keys are inside `artifacts_hash` and inside the merge commits
that name them, which is D-7's reasoning and it has not changed.

**No network.** `--for` writes the key it was given into the id and never asks the host whether
that issue exists. Reading issue content into P0 is the open half of WP12 and belongs there.

**No new field in `intent.yaml` and no new status.** The seven fields are section 5's and the two
statuses are section 5's; A20 decided there is no `merged`.

**No new configuration field.** Nothing is added to Appendix A's `tracker` block: the host comes
from `base_url` because that is the only field in it that names a machine, and the key prefix comes
from the directory because no field names one.

**No change to how an intent is closed or judged.** `intent close` writes `abandoned` as before and
G-Complete reads the file as before.

**No validation of the issue key's shape.** Section 3 says the key need not be a number and names
Jira as the case, so what is checked is the one thing that breaks a reader: a space, which YAML
reads as the start of a comment.

<!-- xeno:section:constraints -->
## Constraints

**The specification is the authority and it does not change for this.** Section 5's field list is
what the command writes, section 3's three parts are what the id is made of, and Appendix A's
`tracker` block is what the derivation may read. Nothing here needs a document change, which is why
there is no commit before this one.

**The `tracker` block is optional.** Appendix A says so, and the reason given is that nothing in
the trail depends on it. So every path through the derivation has to survive its absence, and the
way through is `--for` with the whole id.

**A name is not an address.** `tracker.adapter` is `github` or `gitlab`, and mapping either to a
host would be a table that is correct until the first self-managed deployment — which is the case
`base_url` exists for. The host is read off `base_url`.

**Absent is not guessed.** Where there is no sequence to continue, or where the directory holds two
prefixes, the command refuses and asks for `--intent` rather than choosing a prefix. That is the
same rule `qualified()` states about the id: a guessed value makes every comparison against it
meaningless.

**The two key schemes coexist.** Keys below `XENO-0121` are issue numbers and the current sequence
starts at `XENO-0200`; the gap is how a reader tells them apart (D-7). The derivation takes the
highest number whatever its width, so the two need no special case.

**One writer per artifact.** `intent.yaml` had none and now has one; `IntentClose` remains the only
other writer, and `model.Intent`'s field order stays as it is, because that order is what keeps a
close from reordering a file inside the intent-level hash.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0.**

**One intent, one branch, one issue** — `179-the-intent-is-created-by-a-command`, #179, labelled
wp7.
