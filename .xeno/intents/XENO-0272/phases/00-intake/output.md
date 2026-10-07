---
intent: github.com/triplem/xeno#277
phase: 00-intake
created: "2026-10-07T13:27:21Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5436dd20ed5cf7566de5405d0183f1d7f12c4232bf50d3802350373cb25af122
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

`hashes` in `internal/gates/gates.go` read the artifact's own `tool` field and relaxed every
hash check on the strength of it:

    byHand := raw["tool"] == "manual"
    ...
    honest := byHand || model.OneOf(f, model.WriterlessHash) || (f == "strings_hash" && goneBundle)

With `byHand` true, `model.HashPlaceholder` was accepted in every field of
`model.HashFields` — `context_hash` and `strings_hash` included, the two the gate can
otherwise recompute against the file they name. The comment's reason was sound: nothing
produced a manual artifact, so none of its hashes had a writer. What it assumed is that the
field is true.

Section 12 now says the triple is a declaration and not a measurement, and that nothing
corroborates `model`, `tool` or `tool_version`. So the one gate that acted on the triple was
relaxing itself on a declaration, which is sharper than the provider register #207 was about:
there the consequence stops at how a reader reads a register, here it decided a verdict.

## What the term actually did, measured

Four artifacts in the whole trail declare `tool: manual`: the `output.md` and `digest.md` of
XENO-1's and XENO-2's intakes, hand written on 2026-09-21 and 2026-09-22 before M0. Dropping
the term turns exactly those two phases red, with four findings, all `context_hash says
by-hand where a writer exists`.

Nothing else moves, and the reason is that the other two terms already cover the honest cases
without consulting the artifact. `secrets_hash` and `rules_hash` are in
`model.WriterlessHash`. `strings_hash` is covered by `goneBundle` for these two artifacts
anyway, because both carry `template: intake@0.1.0` and the shipped set answers `intake@1.0.0`.
So the `tool` term had exactly one effect in this repository: it exempted `context_hash` for
two pre-M0 intakes.

## The uncomfortable half

Both of those phases have a `context.lock.yaml` beside them. The file is there and it is
hashable today, so the finding is correct rather than spurious: a writer does exist now. What
is also true is that no writer existed when the artifact was sealed, and section 11 forbids
rewriting a sealed artifact to satisfy a check, so the four findings cannot be fixed. They can
only be released or left.

## The decision

#277 put three shapes: a clause in section 12 saying the exemption rests on a declared field,
a narrower exemption, or accepting it with the reason written down. The maintainer was given
all three with the measurement above and chose the narrower exemption on 2026-10-07, then
approved the four findings with a reason naming the date each artifact was written and why it
cannot be corrected.

<!-- xeno:section:scope -->
## Scope

In scope is `hashes` in `internal/gates/gates.go`. The `byHand` variable goes, and `honest`
keeps the two terms that are facts about the tree: a field nothing writes yet, and a strings
bundle the repository no longer carries. The comment above it says what was removed and why,
since the reason it was there is sound and only its premise was wrong.

In scope is the doc comment on `hashShape`, which said the placeholder is honest "in two cases
and wrong in a third" and gave `tool: manual` as one of the two. It is replaced as a paragraph
rather than edited at the clause, because it is being changed a second time and because the
count it opens with moves.

In scope is one case in `TestHashFieldShape`. "by-hand passes a manual artifact" asserted
exactly the behaviour being removed and is inverted, with a comment saying why the field is not
evidence. The test's own heading is unchanged and still right: the placeholder is honest only
where nothing could have written the value.

In scope are the four findings the change produces and their release. They were approved by the
maintainer on 2026-10-07, two per intent, with a reason naming the date the artifact was
written, that the lock beside it is hashable today so the finding is correct rather than
spurious, and that section 11 forbids rewriting a sealed artifact to satisfy it.
`gate approve` rather than `gate override`: an override carries an obligation, and there is
nothing anybody can close.

In scope is one row in `docs/assumptions.md`, because the decision outlives this intent: the
next person to find a check that would be easier with a declaration needs the reason it was
taken out, and the four approvals need to be findable from somewhere other than two gate files.

## Out of scope

Out of scope is any change to the process definition. Section 12's paragraph on the triple
says nothing corroborates it, which stays true and is the clause this change acts on rather
than amends. A sentence saying one gate used to act on the triple would be a note about
history in a normative document.

Out of scope is correcting the four artifacts. Writing the real `context_hash` into them would
satisfy the check and rewrite two sealed phases, which section 11 forbids in as many words.
That is why the findings are released rather than fixed.

Out of scope is anything about `model` or `tool_version`. Neither is read by any gate; `tool`
was the only one of the three that decided anything, which is what made this worth a change
rather than a clause.

Out of scope is `goneBundle`, which stays. It is a fact about the tree — the shipped set either
answers the recorded ref or it does not — and it is the term that keeps the two pre-M0
artifacts' `strings_hash` honest without their saying anything about themselves.

Out of scope is the `tool` field itself. Section 12 requires it, a project with a provider
register needs it, and nothing here says a recorded value is false. What changes is that no
verdict depends on one.

<!-- xeno:section:context-rationale -->
## Why this context

Six files, 377368 bytes.

`internal/gates/gates.go` holds `hashes`, `recomputed` and the `hashShape` doc comment, which
is the whole of the change. Read rather than recalled: the measurement of what the `tool` term
did rests on which fields `model.WriterlessHash` already covers and on when `goneBundle` is
set, and both are in this file.

`internal/gates/schema_test.go` holds `TestHashFieldShape`, whose six cases are the
specification of this behaviour in code. One of them asserted the thing being removed, which
is how the change was found to be a change and not a refactor.

`internal/model/model.go` holds `HashFields`, `WriterlessHash` and `HashPlaceholder`. Reading
it is what established that dropping the `tool` term leaves `secrets_hash` and `rules_hash`
untouched, which is half of why the blast radius is two phases and not the trail.

`docs/process-definition.md` carries section 12's paragraph on the triple, added for #207
earlier today, which is the clause that makes this a defect rather than a quirk: a gate acting
on a field the document calls a declaration. It also carries section 11's "what is sealed is
never rewritten", which is why the four findings are released and not fixed, and Appendix B,
which defines the placeholder.

`docs/assumptions.md` is where the row lands, read first for the test a row has to meet.

`CLAUDE.md` carries the standing rules: the first, which is why the specification is acted on
rather than amended, and the one about a negative result, which applies because the claim that
nothing else in the trail moves is an absence.

The links block declares `internal/gates` against the process definition, because what the
change asserts is that a gate may not read a clause's declaration, and that is a statement
about both.

Not in scope: `.xeno/intents/XENO-1` and `XENO-2`, whose artifacts are the subject but are not
read for their content — what matters is their `tool` and `context_hash` fields, which the
measurement reports, and nothing in them is edited.
