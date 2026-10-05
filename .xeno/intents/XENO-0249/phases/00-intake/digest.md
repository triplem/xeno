---
intent: github.com/triplem/xeno#201
phase: 00-intake
created: "2026-10-05T12:48:01Z"
schema_version: "1.0"
runner_version: dev+f2f65d5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1a371a30a09d5386056df231bc930934ab318f051e9feebeedd0f7d863d0b8ea
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The release hashes one tree and ships another, and the two agree only because `cp -R` ran a
moment earlier. A copy that succeeds while producing a different tree fails G-Supply for
every adopter of that release, on every phase, with two hex strings and nothing to act on.

The hash cannot compare them: it keys each line on the path relative to the repository root,
so the same bytes at two locations differ by construction. That is the one piece of work.
