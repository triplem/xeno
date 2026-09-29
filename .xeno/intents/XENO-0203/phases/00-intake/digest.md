---
intent: github.com/triplem/xeno#127
phase: 00-intake
created: "2026-09-29T06:53:06Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57acf68.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f62fb0f00238ff490c56429da4e509c26393a95c23ce1ab066c8444f4915f5fe
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The objection in #127 was about the rule set and not the mechanism, and separating those two was most
of the intake. What the measurement then decided is that the obvious selection rule, drop every rule that
uses a device section 4 cannot express, throws away the two most valuable rules in the set. The criterion
that survived is empirical: does the rule fire on this repository's own text. Two did, and both would have
been damaging.
