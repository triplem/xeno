---
intent: github.com/triplem/xeno#257
phase: 02-design
created: "2026-10-06T07:53:39Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1854ba1f9621b50a769dcddc5868407b93829f5f51862a37197afad069939e1c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

Six of seventeen stay required: `problem`, `acceptance-criteria`, `changes`, `test-mapping`,
`results`, `release-notes`. They are the chain from why to what was done to whether it worked, plus
the two a gate reads or will read.

Each of the eleven drops has a reason and the reason is in the file. `scope` and
`context-rationale` go because for a small change the problem states the scope and the information
base is the default; `non-goals` and `constraints` because they bound a design that a one-line fix
does not have; `decisions`, `alternatives` and `impact` because there is no design to record;
`deviations` because a change with no plan beyond its criteria cannot depart from one; `gaps` and
`residual-risk` because they describe what a larger change leaves behind; and `review-checklist`
because its answers live in frontmatter that G-Policy requires regardless.

`required: false` rather than deletion, everywhere. The flag is the only thing `Missing` reads, so
it is the entire lever; and keeping every id defined means `section set` still accepts all
seventeen, a rule naming one cannot find it missing, and an author who wants a dropped section can
simply write it. A candidate that deleted sections would be a different template rather than a
shorter path through the same one.

`release-notes` stays required and says why in its own file. It is the one section any rule names,
through `section-non-empty`, so a candidate that dropped it would ship a template failing a shipped
checked rule at P5. A fixture that did that would be a trap, and the file says so where somebody
editing the candidate will read it.

`acceptance-criteria` and `test-mapping` stay required by decision, and their files distinguish
that from being read. Nothing reads them today; #258 settled that section 9 will make a criterion
identifiable and a gate will read the mapping, so dropping them would foreclose a decision already
taken. Saying "by decision, not by a reader" is what keeps the next reader from believing a gate
exists.

`examples/templates/` beside `examples/rules/`, and each header takes that directory's shape: what
it is, why it is not enabled, how to adopt it. A72 is the precedent — a review rule in the shipped
set reaches every adopter, and a template in the shipped set reaches them the same way.

A README rather than a comment in each file for the method. The per-file headers say what that file
drops; the question of how to try the candidate and what to measure is one answer for the set, and
repeating it six times is how six copies drift.

<!-- xeno:section:alternatives -->
## Alternatives

Keeping only what a gate reads was the minimal candidate: `release-notes` alone, with sixteen
optional. It is the honest floor and it is not a candidate anybody would adopt — the artifacts stop
being a record and the process becomes a verdict with a changelog. The fixture is meant to be tried,
so it has to be defensible rather than extremal.

Dropping whole phases rather than sections was considered and is outside the lever. Section 7 fixes
the six phases and their gates; a template can make a section optional and cannot remove a phase,
and attempting it through an empty template would leave a phase that starts, writes nothing and
finishes, which the gates would judge as a phase that wrote nothing.

A single `small-change` template set rather than six overrides was considered. `Load` resolves per
id and there is no notion of a named variant, so one set is what the mechanism supports; a variant
selector would be a new field and a specification change, which is what the fixture exists to avoid
needing.

Shipping the candidate in `.xeno/plugin/templates/` with the flags changed was rejected on A72's
reasoning. It reaches every adopter, and this candidate is an argument nobody has tested — the
shipped set is not where an untested argument belongs.

Adopting it in this repository to produce the first measurement was the tempting one and is rejected
on comparability. Nine intents were measured against the seventeen; switching mid-session makes the
tenth incomparable with all of them, and the comparison is the only thing the fixture is for. The
measurement wants a deliberate run after this merges, which the README asks for.

Deleting the dropped sections instead of flagging them was considered and rejected on three
counts: `section set` would refuse ids the shipped process accepts, a rule naming a deleted section
would report it unknown rather than empty, and an author who wanted one back would have to edit the
template instead of writing the section.

A `checked` rule asserting the candidate's own shape was considered — something that fails if a
project adopts it and then leaves `release-notes` empty. `release-notes-are-filled` already does
exactly that, which is the point: the shipped rule is what makes the one real constraint
self-enforcing, and a second rule would be a second statement of it.

<!-- xeno:section:impact -->
## Impact

Thirteen files added under `examples/`, nothing else. Six `template.yaml`, six `strings.en.yaml`,
one README. No shipped template changes, no code, no normative document.

Nothing executable changes and nothing can. `Load` reads `.xeno/config/templates` and
`.xeno/plugin/templates`, so a directory under `examples/` is inert by construction — the same
property `examples/rules/` has had since A72, and the same one `gate verify` proves by leaving the
verdict count alone.

What becomes possible is the measurement the plan asks for. A project — including this one, later
and deliberately — copies six directories into `.xeno/config/templates/` and runs an intent, and
the figures are comparable with the nine on #117 because the only thing that changed is which
sections `Missing` asks for. That is the whole of the fixture's purpose and it is one `cp` away.

What the candidate asserts, and a reader can disagree with file by file, is that eleven of the
seventeen describe a change larger than the one a shortcut is for. Each file carries its own reason,
so disagreement lands on a sentence rather than on the set.

The honest limit is that nobody has run it. The candidate parses, its flags are the only difference
from the shipped set, and the one section a gate reads is still required — but whether six sections
produce an artifact a person can review in six months is not something this intent learns. The
README says that trying it is how the question is answered, which is the plan's own position.

The second limit is that the fixture can age without anybody noticing. The shipped templates may
gain a section, or `release-notes-are-filled` may stop being the only rule naming a section, and the
candidate would then be describing a set that no longer exists — exactly what
`docs/clause-readers.md` keeps doing and what #254 had to correct. Nothing reads these files, so
nothing will report it; the README carries the date and the measurement it was built against, which
is the same defence that document has.
