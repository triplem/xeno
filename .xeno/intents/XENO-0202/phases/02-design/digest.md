---
intent: github.com/triplem/xeno#120
phase: 02-design
created: "2026-09-29T05:59:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ad0a764.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 2c0bf120da9cb563376a6beca388eb69a2fa8ace389f50ea01725bda9f04d0e4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two decisions needed the specification read closely rather than a preference. The hash covers the
effective set in a canonical rendering because Appendix B's own word is "effective set", so the bytes of
the files are the wrong subject. And output.md is not filtered because section 16 puts the filtering on
the digest — this record is the proof of why that matters, since it quotes a field name with its value
and a filter over the agent's prose would have mangled its own explanation.
