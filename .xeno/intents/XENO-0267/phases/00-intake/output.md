---
intent: github.com/triplem/xeno#228
phase: 00-intake
created: "2026-10-06T20:26:22Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 16c0ce93519790623d8eff27159bcacc2367285cc38a15a6224b81fa6e3aa464
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

Nothing in the tool suggests starting a phase in a fresh session, and the context economy of
the process assumes somebody does. Section 6 has a phase work from its predecessor's digest
rather than from a fresh scan; section 5 has `context.lock.yaml` state what the phase was
given. Both describe a phase that starts from a known base. A session that carries P0 to P3
into P4 gives that phase a context the lock does not describe, and the digest — which exists
so the next phase re-derives nothing — becomes redundant rather than load bearing.

The gap is already measured, in this repository and in the one place that could measure it.
`internal/cost/cost.go` says of attribution that "filtering a transcript by each phase's own
window captures about seven per cent of what a session spent, because the work is done before
`phase start` is called", and the ledger records a turn with no phase open as `none` rather
than attributing it. So the ledger does not hide an uncleared session; it shows one, as the
`none` lines it cannot attribute.

What is missing is the sentence. `internal/runner/next.go` holds the whole of the suggestion
machinery, documented as "never an action": it reads the working sequence against one intent,
names the step and leaves it unmade, and where the next step is nobody's subcommand it prints
a sentence with no command. Three suggestions matter here. `next` answers `not-started` with
"start " plus the phase and `xeno phase start`, which is what a person sees after
`intent start` and after every `phase finish` that went green. For the last phase it answers
"P5 is decided, so the merge is next. That is not a xeno command: commit, push, and let the
review and the pipeline run", which is the end of an intent. Neither mentions the session.

## What constrains the wording

Section 7 records `XENO_HARNESS` and never branches on it, and says why: "the moment the
runner behaves differently per harness, the tools stop being interchangeable". A CI check in
`xeno.yml` asserts the name appears in one constant and in no condition. `/clear` is Claude
Code's command, Codex has its own, and a person working with commands alone has none, so the
runner may not print any of them. A harness neutral sentence is the only form the runner can
carry, and it is enough: the person reading it knows their own client.

## The second half of the issue

#228 asks for two more answers beyond the sentence. Whether the plugin's skill text should
carry the harness specific command, which its own table leaves at "if anywhere". And whether
the cost record should report the gap between what a phase declared it was given and what its
session actually spent, or whether that is a later measurement. Both are decided in this
intent rather than left, because an issue closed with one of its three answers given reads as
answered.

<!-- xeno:section:scope -->
## Scope

In scope are two sentences in `internal/runner/next.go`, and nothing else in the runner.

The `not-started` suggestion says that the phase is started in a fresh session, and why: so
that its context is what `context.lock.yaml` says it was given. It keeps `xeno phase start`
as its command, because the sentence changes what a person does before typing it and not
what they type.

The last phase's suggestion, which is the end of an intent, says the same for the session
that is ending. It carries no command and already names a step that is nobody's subcommand,
so a second such sentence sits in the place the machinery was built for. #228's comment asks
for this half as "just a hint, no requirement", which is what a suggestion is.

In scope are the tests for both, in `internal/runner/runner_test.go`, beside the existing
assertions on the same two suggestions.

In scope is one row in `docs/assumptions.md` carrying the two decisions #228 asks for beyond
the sentence, because both outlive this intent: the skill text stays free of a harness
command, and the gap between declared context and actual spend is WP20's measurement.

## Out of scope

Out of scope is any change to the documents. The suggestion machinery adds no field, no gate
and no rule — its own comment says so — and section 6 already fixes the sequence this reads.
A sentence that reports the sequence is not a change to it, so there is nothing for the
specification to say that it does not already say.

Out of scope is the harness specific command in the plugin's skill text. #228's table offers
it on the ground that "the plugin is per harness and may name one", and in this tree it is
not: section 13 ships one plugin, vendored under `.xeno/plugin/`, hashed whole by G-Supply,
and the lenses are skills rather than subagents precisely "so they work in both clients". A
`/clear` in that tree would be wrong under Codex and hashed into the digest either way. The
skills are also read by the model, and clearing a session is the person's act, not the
model's, so the sentence would be addressed to the wrong reader as well as to the wrong
client.

Out of scope is a hook that ends a session at `phase finish`. #228 lists it and prices it
itself: a hook that closes somebody's session is the kind of help nobody asks for twice.

Out of scope is a figure in the cost record comparing declared context against spend. The
implementation plan puts it in WP20 under session discipline, "one phase, one session", with
the runner warning when a resumed session has grown past a configured size, and WP20 is the
package that owns measurement and holds the baseline such a figure is read against. Building
it here would be the measurement without the baseline, and WP20's own rule is measure first.

Out of scope is enforcement of any kind. Nothing here is a gate, and a phase started in a
reused session stays green, which is the point of a suggestion.

<!-- xeno:section:context-rationale -->
## Why this context

Eight files, 410657 bytes, in order of how close they sit to the change.

`internal/runner/next.go` is the file being changed. It holds `Suggestion`, with the comment
that fixes what a suggestion may be, and the three answers of `next` that this intent touches
or deliberately leaves alone.

`cmd/xeno/main.go` holds `suggest`, which prints the text, the command and the owed lines,
and `cmdPhaseFinish` and `cmdIntentStart`, which are the two callers that put a person in
front of the `not-started` suggestion. It is read to confirm that the sentence reaches a
person on the path #228 names, and that `--no-next` is the one way to silence it.

`internal/cost/cost.go` carries the measurement the issue reaches for. Its package comment
holds the seven per cent figure and the reason for it, and `NoPhase` is how an unattributed
turn is recorded rather than guessed. It is what makes the cost half of the issue answerable
without building anything.

`docs/process-definition.md` is read for four clauses: section 5 on what `context.lock.yaml`
states, section 6 on the working sequence the suggestion reports, section 7 on `XENO_HARNESS`
being recorded and never branched on, which decides the wording, and section 13 on the one
vendored plugin and the skills that work in both clients, which decides where the harness
command may not go.

`docs/implementation-plan.md` is read for WP20's session discipline, which already owns "one
phase, one session" and the warning on a resumed session, and for WP11, whose scope this
intent sits in. It is what makes the cost figure a scheduled answer rather than an open one.

`docs/assumptions.md` is where the row lands, and is read first for the test a row has to
meet: that the fact outlives the intent that found it.

`.xeno/plugin/skills/xeno-intake/SKILL.md` is the candidate #228's table names. It is read
rather than assumed, because the argument for leaving it alone rests on what the skill text
addresses and who reads it.

`CLAUDE.md` carries the standing rules: that the documents are not the agent's to edit, which
is why the specification is out of scope and the reason for that is stated rather than
implied, and that every change belongs to a work package and an intent.

The links block declares `internal/runner` against the process definition, because the
sequence the suggestion reports is section 6's and the sentence it gains is constrained by
section 7, and reading the file without them gets the wording wrong.
