---
intent: github.com/triplem/xeno#111
phase: 00-intake
created: "2026-09-28T17:31:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e77cc2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 503330549dcdf03e764b0b0ea05f55fd793ce1d78319c88c3924896d5ee1e27c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The scope grew by one line while it was being written: keying on the relative path only
works against a directory that is walked, so the two changes are one change. That was
the whole finding of the intake. Pairing a misplaced comment with a latent map key in
one intent is deliberate, since what they have in common is being correct against
today's data and wrong against what the specification permits.
