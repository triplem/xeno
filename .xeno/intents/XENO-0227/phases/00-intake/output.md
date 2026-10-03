---
intent: github.com/triplem/xeno#179
phase: 00-intake
created: "2026-10-03T12:12:58Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1f18dba6bf50164a47bf524f262249a8e2f5a8c84a5271e4e61901c303ea3a21
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

`intent.yaml` is the one artifact of this process that no command writes. Section 4 lists it in
the intent directory and section 5 gives it a shape; every other artifact has a writer.
`context.lock.yaml` and `output.md` come from `phase start` and `section set`, `digest.md` and
`gate.yaml` from `phase finish`, `learning.yaml` from whoever answers the phase, `cost.yaml` from
`cost turn`. This one is typed by hand before `phase start`, because `qualified()` reads it and
refuses the phase without it.

**Three of its seven fields are guessed where the runner knows them.** `created` is a timestamp a
person types. `runner_version` and `plugin_version` are copied from memory, so every `intent.yaml`
in this repository records `0.1.0-dev` while the artifacts beside it record
`0.1.0-dev+<commit>.dirty`. Two strings for one build in one directory, which is #177 seen from the
intent level.

**`status: in-progress` is written by hand although the runner owns the other value.** `intent
close` writes `abandoned`, and the process definition has no `merged` status by design (A20), so
the only value a person ever types is the one the runner would have defaulted to.

**And the file sits inside `artifacts_hash`,** so a typo in a hand-written field is sealed with the
phase. `gate verify` recomputes the hash and would report the correction as a divergence, which
means a mistyped timestamp is permanent in the same way a verdict is.

What is left after the derivable fields are taken away is the issue. `github.com/triplem/xeno#179`
is the host, the repository and the key, and nothing in the key or in the configuration gives the
key: `tracker.project` holds the repository and `tracker.base_url` names the host. So the issue is
an input and everything else, including the intent's own key, follows from what is already written
down.

<!-- xeno:section:scope -->
## Scope

**In scope.** A command that creates an intent, `xeno intent start --for ISSUE [--intent KEY]`. It
takes the issue, derives the qualified id from the tracker block where the id is not given whole,
derives the key from the sequence the intents directory holds where `--intent` does not name one,
and writes `created`, `status`, `schema_version` and the two version fields from the runner itself.
It refuses where the intent directory exists, as `phase start` refuses a running phase. The
next-step suggestion for an intent that does not exist names the command instead of saying no
command creates one.

**Out of scope, and each for its own reason.**

`assumptions.yaml`. Section 4 lists it beside `intent.yaml`, but nothing refuses for its absence —
the gate reads it with the error discarded and `assumption record` refuses with a sentence naming
the file — and the intents in this repository do not all have one. A command that wrote an empty
register would be creating an artifact to satisfy a list rather than a reader.

Rewriting the intents that exist. Seventy-nine of them are hand written, their keys sit inside
`artifacts_hash` and inside the merge commits that name them, and D-7 already settled that a key
is not renamed for tidiness. A hand-written `intent.yaml` goes on working unchanged, which is a
requirement of this intent rather than a concession.

Reading the issue from the host. `--for` takes a number and writes it into the id; whether that
issue exists is a question for the tracker half of WP12, which is open. Nothing in the trail
depends on the answer, and a command that called out would put the network into the one path that
has to work offline.

A `merged` status. A20 decided there is none, and the second standing rule puts a new value in the
specification before the code.

**The deliberate change to the issue as filed.** The issue writes the command as `xeno intent start
--intent KEY --for ISSUE`, with the key required. The key is derivable — this repository's own
convention is "the next number of a sequence of its own, padded to four digits", and that sequence
is on disk — so requiring it would leave one of the two guessed inputs guessed. `--intent` stays,
optional, because the first intent of a repository has no sequence to continue.

<!-- xeno:section:context-rationale -->
## Why this context

Section 3 is read for what the qualified id is made of: tracker host, project or repository, and
key. It is the only statement of the three parts, it says the key need not be a number and names
Jira as the case where it is not, and it is therefore what the derivation has to agree with.

Section 5 is read for `intent.yaml`'s field list — the common header plus `key`, `status` and a
`reason` when the status is abandoned — because the command writes exactly that set and the second
standing rule forbids an eighth field.

Appendix A is read for the `tracker` block, for two facts that decide the derivation: the block is
optional, so every path through it has to survive its absence; and the fields it has are `adapter`,
`base_url`, `auth` and, in this repository, `project`. None of them is the id's host, which is why
the host has to come from `base_url` rather than from a name.

`internal/runner/runner.go` is read for `qualified()`, which is the reader this writer has to
satisfy, and for `IntentClose`, which is the only other writer of `intent.yaml` and the reason
`model.Intent`'s field order is what it is. `common()` is read for where the runner already puts
`created` and the two version fields into an artifact, because this writes the same values into a
file that until now had them typed.

`internal/runner/next.go` is read for the one sentence that says no command creates an intent,
which this intent makes untrue.

`internal/runner/enforcement.go` is read because it is the only existing reader of the tracker
block, and its anonymous struct is the shape the new reader would otherwise duplicate.

`CLAUDE.md` is read for the key convention and for A20 and D-7, which between them say what a key
is, why the sequence starts at `XENO-0200`, and why the keys below `XENO-0121` are left alone.

Nothing outside the repository is needed. The one question that would have changed the shape of
this — whether the key is an input or a derivation — is answered in the repository, by the
seventy-nine keys that already follow one sequence.
