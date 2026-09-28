---
intent: github.com/triplem/xeno#120
phase: 01-requirements
created: "2026-09-28T20:29:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 89683c8e0ca2d5930d0d525efb6b11d3cae49bca86273feb302e58ad1bceb8ab
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
AC3 and AC4 are the two that constrain rather than describe: the digest carries no
secrets_hash because there is nothing to filter with, and a finish without a summary
behaves exactly as before because section 5 refuses to make the tool path mandatory. AC7
is the one that will be got wrong by anybody implementing this quickly, since a missing
agent block has to leave two fields absent rather than fall back to a plausible default.
