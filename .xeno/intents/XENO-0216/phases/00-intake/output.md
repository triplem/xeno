---
intent: github.com/triplem/xeno#156
phase: 00-intake
created: "2026-10-01T14:34:54Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8f7531e869730170e6dba53236684807fa7b080d5dd92eef13871018546816a2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

`ASSUMPTIONS.md` and `README.md` are the two records this repository keeps by hand, and both
describe a runner that has moved on. They are also the two files read first by anyone
orienting, which is what makes their drift expensive rather than untidy: a stale sentence in
them is read as current by someone who has nothing else to compare it against.

**The worst of it is the README's trail section.** It says the intents stop after P0 until
WP8, because G-Freshness is half implemented until then and no phase beyond P0 could be
judged honestly. A6 closed with #63, both halves are implemented, and the trail in this
repository holds twenty intents that run all six phases and finish green, from `XENO-0107`
to `XENO-0215`. A reader of that paragraph concludes that a tool whose whole claim is a
judged trail has never judged a phase past intake. The file says the opposite of what sits
beside it in the same tree.

**The assumption register's "Not built" paragraph lists what is built.** The secret filter
and the digest writer are both in `internal/secrets` and wired into `phase finish`, where
the digest is filtered before it is hashed and A62 defines the hash over the filter. The
same paragraph calls WP12 the GitHub adapter, where the label and the landed work are
GitLab, and names as still to verify a question section 9 of the plan answers per harness
and #58 implemented.

**Both lists are additive and neither was added to.** The register's "Built" enumerates the
runner commands and stops before `section set`, the three `assumption` commands and
`cost turn`. The README's command block stops at the same place and also predates
`--export`, `--summary` and `intent status --all`. Its coverage table ends at WP10 and the
early packages, while WP8, WP11's export, WP12, WP13, WP15 and WP17 each have tests in the
tree and no row.

**The gap behind all of it.** Neither file belongs to a work package. WP16 is the
documentation: the four documents under `docs/`, the generated reference and the published
site. `ASSUMPTIONS.md` is the record section 4 of the plan asks to be kept by hand until
Xeno governs its own construction, and the README is the front door. The plan says what
goes in them and names nobody who keeps them current, so both are maintained by whoever
notices — which is how they came to describe a tool two months behind. That is a finding
about the plan rather than a task inside one, and the third standing rule says it gets
written down rather than absorbed.

<!-- xeno:section:scope -->
## Scope

**In scope.** The two files, corrected rather than annotated. `ASSUMPTIONS.md`: the "Not
built" paragraph replaced as a paragraph, and one paragraph added to "Built" for what
postdates its list. `README.md`: the opening scope sentence, the command block brought level
with `cmd/xeno/main.go`, the trail section replaced, and the coverage table extended with a
row per package that has tests here and none. And this issue, which is where the gap behind
the drift is written down.

**Out of scope, and each for its own reason.**

The four documents under `docs/`. The first standing rule puts them beyond the agent, and
nothing here needs them changed: every correction is a record catching up with the tree, not
a disagreement with the specification. Where this intent found the plan wanting, in that no
package owns the two files, it writes a finding and stops.

Naming an owner for them. That is the obvious fix and it is a plan change, which under the
first standing rule is a person's commit made before the code that follows from it. This
intent may not write it and should not pre-empt its shape.

WP16's substance. The site, the generated reference, the per-gate pages and the limitations
page are the package, and none of them is touched here. The two files corrected here are not
part of what WP16 publishes, which is precisely the gap.

Backfilling the figures. The register's own history is not reconstructed; what is corrected
is its present tense. An entry that was true when written stays as it is, because the state
column is the record of when it was answered and rewriting it would lose that.

The rest of the staleness survey. The audit that found this also produced a per-package
reading of what is open, which belongs in a reply to the question that asked for it and not
in a file: it is true on the day it is written and the plan's own sections are where that
question is answered durably.

<!-- xeno:section:context-rationale -->
## Why this context

The information base is the two files under correction plus what would contradict them, and
the point of the reading was that nothing in it is taken from either file: a record cannot
be checked against itself.

`ASSUMPTIONS.md` is read for A6, A20 and A62, which are the three entries the corrections
turn on. A6 is closed with #63, which is what makes the README's trail paragraph false. A20
is closed with no action and still holds for the intents up to `XENO-0106`, which is what
keeps half of that paragraph true and is why it is replaced rather than deleted. A62 is the
precedent for a hash defined by the package that writes it and is cited for the secret
filter, which that paragraph says does not exist.

`docs/implementation-plan.md` is read for the done-when of every package and for section 6's
sequence, because "what is open" is a claim about acceptance criteria and not about files.
WP16 is read for what the documentation package actually covers, which is how the finding
about ownership was established rather than assumed.

The tree is read where a record makes a claim about it: `internal/gates/gates.go` for which
gates stand as `not-implemented`, `internal/secrets` and `internal/runner/runner.go` for the
filter and the digest, `internal/host` for the branch rules port and its two adapters,
`internal/cost`, `internal/index`, and `cmd/xeno/main.go` for the command surface, which is
the authority the README's block is brought level with.

`.xeno/intents/` is counted rather than described: fifty-one intents hold one phase and
twenty hold six, and the boundary is `XENO-0107`. That count is what the trail paragraph was
checked against, and the replacement deliberately states the two shapes rather than the
numbers, because a count in a README is stale at the next intent.

The issue labels are read for which packages have findings recorded against them, and
section 9 of the plan for which verification points are answered, which is what dates the
register's remark about harness headers.

Nothing outside the repository is needed. Both files are claims about this tree, so the tree
is the authority, and the one thing a host could have told us — whether #58 closed — is in
the issue list that was read anyway.
