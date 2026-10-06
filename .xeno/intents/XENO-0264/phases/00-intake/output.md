---
intent: github.com/triplem/xeno#258
phase: 00-intake
created: "2026-10-06T17:28:18Z"
schema_version: "1.0"
runner_version: dev+d3983d3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 22639752e3827890b6dd12e29d7821e6ac156cbb3047eab683c70e30fde49511
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

Two clauses of section 5 and section 8 now enumerate what the artifacts may carry, and nothing reads
either. `docs/clause-readers.md` says so in two rows: the recommendation's reason has "**no field to
carry it**; nothing reads it", and G-Test's mapping half has "**nothing**".

**Section 8 now says where the reason lives and `model.Option` has no field for it.** The struct
carries `Text`, `Consequence`, `Recommended` and `Free`. `QuestionAsked` requires exactly one
recommended option and cannot require a reason, because there is nowhere for one to be. So a question
written today satisfies the half of section 8 that has a reader and not the half that does not, and
nothing distinguishes the two.

**Section 5 now says a criterion is identifiable from `requirements@1.1.0` and a mapping names it
from `verification@1.1.0`, and the plugin ships 1.0.0 of both.** The templates are at
`.xeno/plugin/templates/requirements/template.yaml` and `.../verification/template.yaml`, both
`version: 1.0.0`, and every one of the 65 P1 artifacts in the trail declares `requirements@1.0.0`.
So the clause names a version that does not exist and binds nothing. `testReport` in
`internal/gates/gates.go` carries a comment saying the mapping half "has no reader and is not
implemented here", with the measurement and the reason — and the reason it gives is now out of date,
because the numbering convention it waits on exists.

**Section 5's budget clause has the same shape and is the other issue.** #235's amendment landed in
the same session; its code is separate and this intent does not touch it.

## What the gate comment still says, and should not

`testReport`'s comment says the convention "is an addition to section 9 and so a person's commit".
Section 9 is the Rule model; the convention went into section 5, which XENO-0260 corrected in
`docs/clause-readers.md` and nothing corrected here. The comment also carries "46 of 55 P1
artifacts", a figure of 2026-10-05 superseded twice.

## The cost of the template bump, measured rather than assumed

Bumping a template's version changes what G-Schema can check about every artifact declaring the old
one. `hashes` computes `goneBundle` as "the artifact's declared template ref differs from the one the
repository carries", and `recomputed` skips `strings_hash` where it is true, because "a bundle the
repository no longer carries cannot be hashed by anybody".

Measured on a scratch copy on 2026-10-06, by corrupting one sealed P1 artifact's `strings_hash` to a
value that is not the bundle's:

| template version | G-Schema on the corrupted artifact |
|---|---|
| `requirements@1.0.0`, as shipped | **red** — "strings_hash does not match the strings bundle of requirements@1.0.0" |
| `requirements@1.1.0` | **green** |

So the bump silently removes that reader from all 65 sealed P1 artifacts, and from all 65 P4 artifacts
if `verification` is bumped too. Nothing announces it: `gate verify` reports 459 verdicts verified
either way.

**It is one of two readers and not the only one.** The same experiment showed `gate verify`
reporting `DIVERGENT XENO-0258 02-design: committed status green, recomputed red` and exiting 1,
because editing a sealed artifact moves its `artifacts_hash` and the successor's G-Freshness reads
the predecessor hash. Tampering is still caught; what is lost is the field-level reader that names
which field moved.

<!-- xeno:section:scope -->
## Scope

In scope is `Reason` on `model.Option`, with `omitempty`, and `QuestionAsked` requiring it on the
option that carries `recommended: true`. The writer only, where the recommendation check already
sits: `QuestionShape`'s own comment records why that half is not in the gate, and the reason holds
for this one unchanged.

In scope is bumping `.xeno/plugin/templates/requirements/template.yaml` and
`.../verification/template.yaml` to `1.1.0`. Nothing else in either template changes — not a
section, not an order, not a required flag — so the anchors an artifact carries are the same before
and after.

In scope are two checks, one per clause and each in the gate its section implies. Section 5 says the
`acceptance-criteria` section is a numbered list, which is a requirement on an artifact's shape, so
G-Schema reads it on a P1 artifact declaring `requirements@1.1.0` or later. Section 7 asks G-Test for
"mapping of acceptance criteria complete", so G-Test reads it on a P4 artifact declaring
`verification@1.1.0` or later.

In scope is both checks applying from the declared template version and not from the current one. An
artifact declaring 1.0.0 is judged as it always was, which is what section 5 says and what makes
nothing in the trail re-judged.

In scope is this intent's own artifacts being the first judged by both checks, because its P1 and P4
render at 1.1.0. That is the forward-only anchor meeting its first subject, and if either check is
wrong about a legitimate format it fails here rather than on somebody else's work.

In scope are the two rows of `docs/clause-readers.md`, which move from reporting no reader to naming
one, and the two paragraphs under the table that say what each clause waits on — it is no longer a
commit.

In scope is `testReport`'s comment, which says the convention is "an addition to section 9" and
carries a figure of 2026-10-05. Both are wrong now and the comment is the thing a reader of the gate
meets.

In scope is recording the measured cost of the bump. 130 sealed artifacts lose the `strings_hash`
reader and nothing announces it, which belongs in this intent's artifacts and in the pull request
rather than in a surprise six months from now.

Out of scope is #235's code. `Advisory` on `model.Finding`, `result` ignoring an advisory finding and
`budget` marking its own are the other issue under the other work package, and the third standing
rule gives them their own branch.

Out of scope is making the gate read the reason. `QuestionAsked` is the writer and `QuestionShape` is
what the gate calls; moving the recommendation half up would fail every sealed question that predates
it, which is the measurement `QuestionShape`'s comment carries and #229 settled.

Out of scope is re-judging the trail. Neither check applies to an artifact declaring 1.0.0, which is
all 130 of them.

Out of scope is restoring the `strings_hash` reader for an artifact declaring an older template.
There is no way to tell a hash that is wrong from a hash that is right about a bundle the repository
no longer has, and guessing either way is worse than the gap. It is recorded, not repaired.

Out of scope is a numbering convention stricter than "a numbered list". Consecutiveness, starting at
one, and a maximum are not in section 5, and the second standing rule makes each of them an
invention.

No normative document is touched. Both clauses were committed in XENO-0262 and this is the code that
follows them, which is the order the first standing rule asks for.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the two clauses, the three files that have to change to read them, and the two that
decide how far the change may go.

`docs/process-definition.md` is read for both clauses as committed. Section 8's paragraph puts the
reason on the option the recommendation names, which is what `QuestionAsked` has to require and
where. Section 5's "Acceptance criteria are identifiable" gives the two template versions and says
the requirement is carried by the template version and not by the schema, which is what makes the
checks apply from a declared version rather than from the current one. Section 7's G-Test row is read
to confirm the mapping check belongs there and the numbering check does not.

`internal/gates/gates.go` is read for four things it already does and one it says it cannot.
`QuestionAsked` and `QuestionShape`, and the comment on the second recording why the recommendation
half is the writer's and not the gate's. `hashes` and `recomputed`, for `goneBundle` — which is what
makes a template bump cost something, and the reason this intent measured it rather than asserting
that a version number is free. `testReport`, whose comment says the mapping half is unimplemented and
why, with two statements that are now wrong. And `result`, read to confirm that a new finding from
either check fails its gate in the ordinary way, since #235's advisory field is not in this branch.

`internal/model/model.go` is read for `Option`, which gains a field, and for `HashFields` and
`WriterlessHash`, which is where the `strings_hash` reader lives that the bump will cost. It is also
where `Phases` and `TemplateID` are, which the two checks need to find a P4's predecessor.

`internal/template/template.go` is read for `Load`, `Ref` and `Parse`. `Load` resolves whatever
`template.yaml` is in the tree and does not consult the version an artifact declares, which is why a
bump does not break rendering and why `goneBundle` is the only place the version is compared. `Parse`
returns the rendered sections by anchor, which is how both checks read a criterion and a mapping back
out of an artifact.

`docs/clause-readers.md` is read for the two rows and the two paragraphs XENO-0260 wrote under the
table. It is the record of which clauses have readers, and this intent is the first thing in three
intents that moves a reader column rather than explaining why it cannot be moved.

`CLAUDE.md` is read for the three standing rules — the first because this is the code that follows a
specification commit and therefore comes after it, the second because every field and check here was
enumerated first, the third because #235's code is a different branch — and for the convention on
verifying a negative, which is why the cost of the template bump was measured on a scratch copy by
corrupting a sealed artifact and watching the gate, twice, rather than reasoned from `goneBundle`.

The measurement that decided the scope of the risk was run and not recalled: a sealed P1 artifact's
`strings_hash` was corrupted, G-Schema went red at `requirements@1.0.0` and green at
`requirements@1.1.0`, and `gate verify` reported a divergence and exited 1 in both cases. The first
attempt at that experiment used a hash of sixty-four digits, which YAML types as a number so the
check skipped it silently and both runs looked green; the value was changed to one containing letters
and the experiment re-run. That is the shape #263's convention is about, found inside this intent's
own evidence rather than in the code.
