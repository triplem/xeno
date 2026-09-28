---
intent: github.com/triplem/xeno#118
phase: 00-intake
created: "2026-09-28T20:11:37Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57e9207.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 0a88e621dc3c7ffb30fea5f22d432cdc2ea7a9189e39fc86daaaa4f821f6296b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The issue asked whether the numbering could be an improvement and the intake found that
the sort order was never the key's job: created is already recorded in every intent.yaml
and nothing reads it. So the change is two things that look unrelated and are not, a
listing that answers the question and a key that stops trying to. The startpoint of 0200
is what keeps the two schemes apart without a rule, since a sequence beginning at
fifty-six would land on a key that exists.
