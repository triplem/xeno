---
intent: github.com/triplem/xeno#207
phase: 00-intake
created: "2026-10-07T09:04:32Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f976cf5fdb4a104187d89a283f5ba966ce4731df032eb8bfd1b8c8f044f8a134
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

Section 12 requires `model`, `tool` and `tool_version` in every artifact and calls the three
the raw material a project needs where it has to keep a register of its ICT service providers.
All three have writers now: `tool_version` since #181, `XENO_HARNESS` reaching `tool` under
A84. Nothing corroborates any of them.

`r.ToolVersion` is `--tool-version` or `XENO_HARNESS_VERSION`, whichever the caller gave, and
`recordedToolVersion` reads it back out of the artifact for the digest rather than asking
again. `agent()` reads `tool` from `project.yaml`'s `agent.tool` and `model` from
`agent.model.default`, with `XENO_HARNESS` overriding `tool` alone. G-Schema's `sessionFields`
requires all three to be present, and `hashes` checks the shape of the hash fields beside them
and says nothing about these.

So the register the triple feeds cannot distinguish a reported value from a verified one, and
#207 found this by the clause audit under the heading "a reader that cannot fail". It also
named what shape it is: `plugin_version` sat in G-Schema's required set for weeks, filled by a
constant, enforced and meaningless. The three fields are better than that — they have real
writers — and the gap is the same one.

## Why corroboration is not available

Section 7 keeps the runner from branching on the harness, with the reason in the document: the
moment the runner behaves differently per harness, the tools stop being interchangeable. A CI
check asserts the name appears in one constant and in no condition. So the runner cannot ask
the harness anything about itself beyond what the harness volunteered.

The only other copy of the three inside the repository is `project.yaml`, which is the same
declaration in a second place, and #207 says plainly what a check against that would be: "a
check that compares the field against another copy of the same self-report, which would make
the gap invisible instead of closing it".

One record does come from another hand, and it is worth naming rather than leaving a reader to
conclude nothing exists. The plan's section 9 records what a gateway reports: the deployment a
request was routed to, in `x-litellm-model-name` and in the spend log row, and it says what that
is worth — "a recorded `model` is a routing record rather than an attestation from the
provider". Two things keep it out of reach. The runner makes no model request in v1; the harness
does. And `xeno gate ...` never touches the network, which is what makes a verdict reproducible
on any machine.

## What is left

#207's two branches are corroboration or a statement. Corroboration is unavailable for the
reasons above, so the statement is what is left: that the triple is a declaration, where each
value comes from, why nothing can check it, and that the register is read in that light. The
issue puts it in section 12 and in the register, which is where it goes.

The maintainer was put the three options with their consequences on 2026-10-07 and chose this
one. The wording was then drafted, put to them and written on instruction, which is the
exception to the first standing rule and is named in the specification commit.

<!-- xeno:section:scope -->
## Scope

In scope is one paragraph in section 12 of `docs/process-definition.md`, after the one that
calls the triple the raw material for a provider register, drafted and approved before it was
written. It says the triple is a declaration and not a measurement, where each of the three
values comes from, that G-Schema checks presence and shape and nothing more, why the runner
cannot corroborate one without breaking section 7, that a check against `project.yaml` would
hide the gap rather than close it, what the gateway's record is and why it is out of reach,
and that a register built on the triple says what an agent declared.

In scope is `docs/clause-readers.md`. The paragraph is an explanation in that document's four
kinds — nothing can violate it by writing the wrong thing, because what it states is an
absence — so it is counted and not enumerated, and the count moves from 129 to 130 with a
paragraph saying which clause moved it and why it has no row. That is the shape #267's three
clauses were given in the same document.

In scope is one row in `docs/assumptions.md`, because the decision outlives this intent and
names the contrast that makes it readable: `plugin_version` has an anchor the checked side
cannot influence, compiled into the runner, and the triple has none.

## Out of scope

Out of scope is any check comparing an artifact's `tool` or `model` against `project.yaml`.
#207 names this as the thing that should not happen, and the reason holds: it is the same
declaration in a second place, so a green result would mean the two copies agree and nothing
else, while the gap it was built to close would stop being visible.

Out of scope is the gateway comparison. It is the one corroboration that would be worth
something, and it needs a gateway, which section 12 puts in v2, plus a figure outside the gates,
since `xeno gate ...` never touches the network. It was put to the maintainer as an option with
that cost and not chosen. It belongs to whichever package owns measurement when a gateway
exists, and this intent does not assign it.

Out of scope is anything about `plugin_version`, `runner_version` or `secrets_hash`. They are
named in the register row as the contrast, because `plugin_version` is the field that had this
shape and now has an anchor, and naming it is what makes the row legible. None of them changes.

Out of scope is a new gate result or finding saying a field is unverified. That would be a
value outside the set A4 and A42 fix and would fire on every artifact forever, which is #235's
shape and was answered there.

Out of scope is removing any of the three fields. Section 12 requires them and they are the
raw material a project with a register needs; what was wrong is what a reader took them for,
not that they are recorded.

<!-- xeno:section:context-rationale -->
## Why this context

Seven files, 536085 bytes.

`docs/process-definition.md` carries the paragraph and the three clauses it rests on: section
12's requirement of the triple and its provider-register sentence, section 7's rule that
`XENO_HARNESS` is recorded and never branched on, and section 12's own promise that
`xeno gate ...` never touches the network, which is what puts the gateway out of reach.

`docs/implementation-plan.md` carries section 9's answer about the gateway, which is the only
place in this repository that says what a gateway's record of a model is worth: the deployment
a request was routed to rather than an attestation from the provider. The paragraph's claim
about the one record from another hand rests on it, and getting it wrong in either direction
would be either a false limitation or a false promise.

`internal/runner/runner.go` holds the three writers, read rather than recalled: `ToolVersion`
and `recordedToolVersion` for the version, `agent()` for `tool` and `model`. It is what
corrected the drafted wording, which said `model` comes from `project.yaml` or the harness
where `agent()` reads it from `project.yaml` alone.

`internal/gates/gates.go` holds `sessionFields`, which is G-Schema's whole involvement with the
triple, and `hashes`, which is the check beside it that does recompute something. Reading both
is what makes "presence and shape and nothing more" a statement about the code.

`docs/clause-readers.md` is where the new clause is counted, read first for the taxonomy and for
how #267's three clauses were handled in the same position.

`docs/assumptions.md` is where the row lands, read for the test a row has to meet.

`CLAUDE.md` carries the standing rules: the first, which the paragraph was written under an
instruction against; the one about a decision put to a person one at a time, which is why the
options and the wording were two questions; and the one about a negative result, which applies
here because the claim is an absence.

The links block declares `internal/gates` against the process definition, because the sentence
about what G-Schema checks is a statement about both and reading either alone gets it wrong.
