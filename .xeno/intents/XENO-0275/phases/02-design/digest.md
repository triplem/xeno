---
intent: github.com/triplem/xeno#324
phase: 02-design
created: "2026-10-08T13:46:50Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fe01733c917a299e8973c17db55a1691110f151d0ad7dde15af460f416382187
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two lines in one template and a test in the file that already renders every wrapper.

The design question was whether the line is present or commented, and the answer follows from
what this repository can know: an uncommented `--user` needs a number, and the only number
available is the one that is already right on the hosted runner, so the line would be
redundant there and wrong anywhere else. Commented, it offers the name of the knob and the
reason to reach for it, which is all a generated file can honestly give.

Not a `project.yaml` field: Appendix A enumerates that file and the uid belongs to the
runner rather than to the project.

One section of the four, `impact`.
