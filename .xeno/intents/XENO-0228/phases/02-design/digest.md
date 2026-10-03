---
intent: github.com/triplem/xeno#181
phase: 02-design
created: "2026-10-03T12:39:12Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6cbeac4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e8bf10ef8058c02b7658172f52832a0a2377a7acd7b4eeb47890d512f6fe1567
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`ToolVersion` is a field on `Runner`, set by `parse` for every command, joining `Base`, `Head` and
`EvidenceFrom` as an input of a run rather than an argument of a method. `SectionSet` writes it only
when reported and never erases, so the carry-over of an existing artifact's frontmatter means one
report per phase is enough. `writeDigest` prefers a report and otherwise copies from the phase's
`output.md` through `recordedToolVersion`, which answers empty for a missing artifact, missing
frontmatter or missing field alike; the copy lives in `writeDigest` because that is the function
every `phase finish` reruns, so any other home would leave a path that writes a digest without it.
Six alternatives refused. `XENO_HARNESS_VERSION` is the right long term answer and needs a change to
the list section 7 enumerates, which the first standing rule makes a person's commit — put to the
maintainer before any code was written and answered, and it becomes this flag's fallback rather than
its replacement. The runner asking the harness is the branching section 7 forbids. A `project.yaml`
field would be an Appendix A change and stale after the next upgrade, turning an absent field into a
confidently wrong one. A second input to `phase finish` moves the hand edit rather than removing it.
Defaulting to `runner_version` is the exact failure A35 names. And relaxing G-Schema would have
hidden the gap in every project instead of this one.
