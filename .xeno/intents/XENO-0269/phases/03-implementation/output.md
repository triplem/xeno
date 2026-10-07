---
intent: github.com/triplem/xeno#207
phase: 03-implementation
created: "2026-10-07T09:08:05Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 91456688e97f3921b20cfb81360977163e07dfca0dff93fdfc2aa2aec36f8dc7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`docs/process-definition.md`, section 12, one paragraph. It goes after the one that calls the
triple the raw material a project needs for a provider register: **The triple is a declaration,
not a measurement.** It names where each of the three comes from — `--tool-version` or
`XENO_HARNESS_VERSION` for the version, `project.yaml`'s default for `model`, `project.yaml`
with `XENO_HARNESS` over the top for `tool`; says G-Schema checks presence and shape and
nothing more; says the runner cannot corroborate one because section 7 keeps it from branching
on the harness; says a check against the only other copy in the repository would hide the gap
rather than close it; names the gateway's record, what the plan says it is worth and why it is
out of reach; and ends on how the register is read.

`docs/clause-readers.md`, the explanation count from 129 to 130, with two paragraphs beside
#267's. The first says which clause moved it and why it has no row: what the clause states is
an absence, and a reader of the code cannot violate an absence by writing the wrong thing. The
second says it is reader-shaped in one direction only and that the direction is already taken —
building the refused check would leave the clause true and the gap invisible, which is the
thing #207 warns about, and what would make the clause false is a gateway and a figure read
against it, both outside v1.

`docs/assumptions.md`, one row. A100 carries the decision, who took it and when, the two
options not chosen with their costs, the three writers as the code has them, `plugin_version`
as the contrast with its anchor, and the correction to the drafted wording.

The specification paragraph was committed first and on its own, as `CLAUDE.md` requires. The
count, the row and the artifacts follow in the second commit.

Nothing else is touched. No Go file, no gate, no field, no template, no rule, no skill.
`go test ./...`, `gofmt -l .` outside `vendor/`, `go vet ./...` and `xeno gate verify` all pass.

<!-- xeno:section:deviations -->
## Deviations from the design

Three, and the first is the one that matters.

**The approved wording was wrong about `model`, and was corrected before the branch went
anywhere.** The draft put to the maintainer said `model` comes from `project.yaml` or from the
harness. `agent()` reads `model` from `agent.model.default` alone; `XENO_HARNESS` overrides
`tool` and nothing else, and the comment above the function says both are project level and
not among the fields that come from the harness. The phrase was inherited from #207, which says
the same thing, and agreeing with an issue is not reading the code. It departs from P1's
criterion 2, which asks that each of the three be described as the code has it, and the
criterion is met by the paragraph as it stands. The specification commit was amended rather
than followed by a correction, because it had not been pushed and a wrong sentence in the trail
would have outlived the fix. P0's `learning.yaml` carries the general form: drafted wording that
states what the code does cites the function before it is put to a person.

**The paragraph was replaced a second time to drop two em dashes.** The document uses one in
132 kilobytes, so two more would have been a style this file does not have. Replaced as a whole
paragraph rather than edited at the dashes, which is what `CLAUDE.md` asks for a paragraph being
changed a second time, and read back whole afterwards. It departs from nothing: the approved
wording's meaning is unchanged and the punctuation was never what was approved.

**The register row is A100.** P1's criterion 10 asked for one row without saying which number.
A98 and A99 are on the unmerged branches for #228 and #238, so the next number this branch can
see is A98 and taking it would produce a collision that is not renameable, because a row's
number is cited from other rows. A100 leaves both to the branches that wrote them.

Nothing else departs from P2. The paragraph is where the design put it, the clause is counted
and not enumerated, and the gateway is named without being assigned.
