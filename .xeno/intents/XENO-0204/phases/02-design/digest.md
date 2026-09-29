---
intent: github.com/triplem/xeno#128
phase: 02-design
created: "2026-09-29T09:30:41Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+38430d4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7a491f4b36247b02e42c027c7571efbb29b5c9c677e07ebfd6ad484d99851f9d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two rejected alternatives are the ones worth the record. Excluding .xeno/ would have removed the noise and
the subject at once, since a digest is exactly where a leaked secret would land. And pointing the scan at
the project's own filter sounds like one definition of a secret everywhere, when it is upstream's rules with
the precision devices stripped out: the scan would inherit a redactor's widening and find nothing the runner
had not already caught.
