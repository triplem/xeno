---
intent: github.com/triplem/xeno#111
phase: 02-design
created: "2026-09-28T17:33:06Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e77cc2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: e02a19406aa6874c48bd01e7c970e421ee70e3219f29cad3beba057342256bf9
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The alternative worth writing down is the one the issue could be read as asking for: key
on the relative path and keep reading one level. It fails an acceptance criterion rather
than a taste test, which is why P1 was worth writing first. The second rejected variant,
forbidding subdirectories outright, was rejected on authority: section 4 permits them,
so refusing them is a specification change and not a code change.
