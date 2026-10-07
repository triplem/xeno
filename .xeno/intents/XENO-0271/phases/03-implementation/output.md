---
intent: github.com/triplem/xeno#273
phase: 03-implementation
created: "2026-10-07T09:56:51Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b9cbd042396f4cb333816d8c98cf3ba16a0572bdd60b365674b3909fd6866412
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

`docs/assumptions.md`, one row, one cell. A33's state column keeps `approved` and gains what
spelling the prefix out would cost.

It says that the directory name and the template id are separate things, naming
`Resolved.Ref()` and the `id:` field, because that is the distinction the whole answer turns on
and the one a reader of the issue would not have. It says what renaming the directory alone
does: harmless to the trail, with the two figures that establish it — 495 verdicts matching and
a corrupted `strings_hash` still caught — and then its real cost, that `Load` keys on the phase,
section 5's per-id sentence becomes false, and a project override moves. It says what renaming
the `id:` does: `goneBundle` for every sealed artifact and the recompute retired across the 495
that carry a template ref, counted on 2026-10-07 and growing. It says in its own sentence that
`gate verify` reports all 495 verified under either form. It says what the prefix buys, which
is `ls` ordering, and that `template.yaml` already carries `phase: 00-intake`. It names the
maintainer's decision with its date and says both variants stay written out in #273.

The assumption and the reason are untouched, and the row is one line with five cells as every
row in that table is.

Nothing else changes. No Go file, no template, no normative document, no gate, no artifact.
`gofmt -l .` outside `vendor/` is clean, `go vet ./...` passes, `go test ./...` passes, and
`xeno gate verify` recomputes and matches 498 verdicts.

<!-- xeno:section:deviations -->
## Deviations from the design

One, and it is a correction made inside P2 rather than a departure from it.

**A decision was drafted on a false premise and rewritten.** The reason for not touching any
code comment was first written as "`TemplateID`'s comment already cites A33, so the row is
reachable from the function". A grep for A33 across the Go files returned one line, in
`phaseNumber`'s comment in `internal/runner/next.go`, which is about the `--phase` prefix and
not about the template id. `TemplateID`'s own comment restates A33's reason in A33's words and
does not name the row. The decision is unchanged — no comment is touched — and its reason now
says what is true, and names the reachability as a finding this intent leaves standing. It
departs from nothing: P1 set no criterion about comments and its constraints put no Go file in
the diff, which still holds.

It is worth recording rather than quietly fixing because of what caught it. The claim was about
the tree, it was plausible, and it was written in a decisions section where a reader would take
it for checked. What found it was running the grep while writing the next paragraph, which is
the same accident that found the error in XENO-0269's approved wording. P0's learning record of
that intent proposed that drafted specification wording cite the code it describes; this is the
same defect one document down, in an artifact rather than in the specification, and nothing
proposed covers it.

Nothing else departs from P2. The measurement is in the state column, `approved` stands, both
variants are named, the figure is a dated count, the `gate verify` sentence is its own, the two
incidental findings stayed on the issue, and no code comment changed.
