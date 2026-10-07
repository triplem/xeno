---
intent: github.com/triplem/xeno#273
phase: 00-intake
created: "2026-10-07T09:52:56Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 531321fa2b5bed268d9849ad359230ee142f9f13db61771c9c7dc853ff1eefab
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

#273 asks whether the shipped template directories should carry the phase as a prefix, so
that `intake` reads `00-intake` or `P0-intake`, and ends "Does this make sense?". It is a
question put to be answered rather than a defect, and the answer needed measuring because the
two ways of doing it behave differently and only one of them is obvious.

## The directory name and the template id are not the same thing

`Resolved.Ref()` is `Template.ID + "@" + Template.Version`, read from the `id:` field inside
`template.yaml`. The directory name is only the path `template.Load` resolves, and
`model.TemplateID` maps a phase to it by stripping the ordering prefix, which is A33. So a
rename has two forms and they are not variants of one change.

**Renaming the directory alone** leaves `id: intake`, so every artifact goes on recording
`template: intake@1.0.0`. Measured on a throwaway copy with `TemplateID` made the identity
function: `gate verify` passes all 495 verdicts and a `strings_hash` corrupted to a hex value
in a sealed artifact is still caught, naming the bundle it no longer matches. Harmless to the
trail, and it contradicts section 5. "Resolution is per template id: project beats plugin. A
project can replace a single template without forking the set" — under this form `Load` is
called with the phase and finds the directory by phase name, so resolution is per phase, the
id is decorative, and a project override moves to `.xeno/config/templates/00-intake/`.

**Renaming the directory and the `id:` with it** is what the issue is actually asking for, and
it is the destructive one. `gate verify` still passes all 495 verdicts, because `artifacts_hash`
does not cover the template ref, and the same corruption now goes green. Every sealed artifact
records an old id, the shipped set answers a new one, so `hashes` sets `goneBundle` and stops
recomputing `strings_hash` — the branch that exists because "a bundle the repository no longer
carries cannot be hashed by anybody". 495 artifacts record a template ref and every one of them
is under an old id, so the check retires across the whole trail and nothing says so.

## What the prefix would buy

`ls` ordering, and agreement with `phases/00-intake/` as it reads everywhere else.
`template.yaml` already carries `phase: 00-intake`, so a reader who opens the directory learns
the phase from the file, and the only reader who sees the bare directory name is somebody
listing the tree.

## The answer, and why it is worth writing down

The measurement was put on the issue on 2026-10-07 with a recommendation to leave the
directories alone, and the maintainer chose that and asked for the reason recorded rather than
left in a closed issue. A33 says the prefix "says nothing a template needs" and names the
alternative it was chosen over, a mapping table. What it does not say is what spelling the
prefix out would cost, and that is now a figure rather than a preference.

<!-- xeno:section:scope -->
## Scope

In scope is A33 in `docs/assumptions.md`, amended rather than replaced. Its assumption and its
reason stand: the template id is the phase without its ordering prefix, the prefix orders the
phases and says nothing a template needs, and the alternative was a mapping table. What its
state column gains is what spelling the prefix out would cost, measured: that renaming the
directory alone makes resolution per phase and contradicts section 5's per-id sentence, and
that renaming the `id:` with it retires the `strings_hash` recompute over every artifact in the
trail while `gate verify` goes on reporting them all verified.

That is the whole of the change. One row, one cell.

## Out of scope

Out of scope is renaming anything. The directories stay, `model.TemplateID` stays, and
`template.yaml`'s `id:` fields stay. The measurement is on the issue with both variants and
their costs, so either stays available if somebody later decides the `ls` ordering is worth
what it costs.

Out of scope is any change to the process definition. Section 5's per-id resolution sentence is
cited as the thing one variant would contradict, not amended; it is correct as it stands and
the variant is what would have been wrong.

Out of scope are the two things the measurement found on the way, both written on the issue and
neither this intent's. A `strings_hash` whose value is 64 digits with no letter is parsed by
YAML as a number, so `raw[f].(string)` fails and the field is skipped by both the shape check
and the recompute; it is how the first attempt at this measurement went wrong, and a real
sha256 being all digits is a one-in-ten-trillion accident. And `goneBundle` cannot tell a
bundle that is gone from one that was renamed, which is why the second variant is quiet; that
is the behaviour the branch was written for and a change to it needs an argument this issue
does not supply.

Out of scope is a check that `template.yaml`'s `id:` matches its directory name. It would make
the second variant impossible to do by halves and it would also forbid the first, which is a
decision about what a template id is rather than a guard, and nobody has asked for it.

Out of scope is `docs/clause-readers.md`. Nothing is added to a normative document, so there is
no clause to count.

<!-- xeno:section:context-rationale -->
## Why this context

Six files, 369156 bytes.

`docs/assumptions.md` holds A33, which is the row being amended, and is read first for the test
a row has to meet and for how a state column carries a correction: A95 and A97 are the
precedents for a cell that records what was learned after the row was written.

`docs/process-definition.md` carries section 5's template resolution clause, "Resolution is per
template id: project beats plugin", which is the sentence one variant contradicts and the
reason that variant is not merely cosmetic. It also carries Appendix B, which is what
`strings_hash` is recomputed against.

`internal/template/template.go` holds `Resolved.Ref()` and `Load`. Reading them is what
separated the two variants: the ref comes from the `id:` field and the directory is only the
lookup path, and an answer written without that distinction would have been wrong in whichever
direction it guessed.

`internal/model/model.go` holds `TemplateID`, which is A33 in code, and `HashFields`, which is
the set the recompute walks.

`internal/gates/gates.go` holds `hashes` and `recomputed`, which is where `goneBundle` decides
whether `strings_hash` is compared at all. That branch is the whole of the second variant's
cost and the measurement is a statement about it.

`CLAUDE.md` carries the standing rules: the first, which is why the specification is cited and
not touched, and the one about a negative result being evidence only when the thing checked was
there to be found, which is exactly what went wrong on the first attempt and is recorded.

The links block declares `internal/template` against the process definition, because the
per-id resolution sentence is a statement about that package and reading either alone gets the
first variant's cost wrong.

Not in scope: `.xeno/plugin/templates/`, because nothing in it changes, and the trail, because
no artifact moves.
