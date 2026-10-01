---
intent: github.com/triplem/xeno#160
phase: 03-implementation
created: "2026-10-01T15:39:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bca3ab449142232b93033c1ae233ea320779713ea79a9fc6317b8d0c6edfbc45
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`internal/model` gains `ChecklistEntry` and `Output.ReviewChecklist` as `review_checklist`, with
`ChecklistResults` and `ChecklistNeedsNote` declaring section 9's three results and the two that owe a
note, so the asymmetry is a declaration rather than a condition written twice. `internal/gates` gains an
empty predicate registry and `policy`, which resolves through `internal/rules` and adds no resolution of
its own, discards the problems both calls return because an unresolvable tree is G-Rules's finding,
reports a checked rule whose type nothing implements at every phase it applies to, and counts the review
rules against the P5 checklist at the last phase. The count runs from rules to entries; an entry with no
`rule` is counted towards nothing, which is section 12's lens clause; an entry naming a rule outside the
set is a finding. A68 and A69 record the frontmatter key with why the gate keys on `rule` rather than
`source`, and the reading that every review rule is answered at P5 whatever its `applies_to`. Five
deviations, two of them about the tests: two helpers collided with names the package already had, and
the empty-tree test's first version passed while asserting almost nothing.
