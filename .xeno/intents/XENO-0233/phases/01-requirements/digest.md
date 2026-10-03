---
intent: github.com/triplem/xeno#177
phase: 01-requirements
created: "2026-10-03T17:00:45Z"
schema_version: "1.0"
runner_version: dev+b54626e.dirty
plugin_version: 0.29.2
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dc26eac4a874edb2d5754dae74387e71f4d20d19bd3ec7eeeab5737731fa8438
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`plugin_version` is the vendored plugin's own, read from its manifest, absent where no plugin is
vendored — and section 5 requiring the field in every process file is what makes that absence a
G-Schema finding rather than something to paper over. The release sets the runner's version and not
the plugin's, one `-X` where there were two. The lock carries `plugin: { version, sha256 }`, which
section 5 already enumerates, absent together where there is no plugin because a version with no
hash is half a claim; the tree hash is defined to the byte in the package that writes it, which is
the precedent Appendix B sets for `secrets_hash` and `rules_hash`. A development build's
`runner_version` claims no version — `dev+<commit>` — because a literal cannot learn the newest tag
and `0.1.0-dev` was wrong rather than stale, and the manifest declares 0.29.2 where it declared
0.1.0. The entry point normalises what section 7 lists and deliberately not `XENO_PLUGIN_ROOT`, with
the measurement as the reason, and the hook calls it plugin-relative. `XENO_HARNESS` reaches the
`tool` field and beats `agent.tool` where both are set, with nothing branching on it, checked in the
job whose name is the required context. The trail is unaffected: sealed artifacts keep their
headers, because a header lives in the file and verify recomputes verdicts. Out of scope with
reasons: the resolution order, awaiting a section 7 decision; `XENO_PLUGIN_DATA` read by the runner,
which section 7 pins to one value; the manifest tracking the next release, which no mechanism here
can do; a statement in section 16; and G-Supply, for which the lock now records what it would
compare.
