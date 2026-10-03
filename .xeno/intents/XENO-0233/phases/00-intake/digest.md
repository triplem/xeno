---
intent: github.com/triplem/xeno#177
phase: 00-intake
created: "2026-10-03T17:00:03Z"
schema_version: "1.0"
runner_version: dev+b54626e.dirty
plugin_version: 0.29.2
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9003e4040aee2e86ef95618faf10339d3ae60bea39abfa5356539c0a8ed37453
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two version fields, neither saying anything true. `plugin_version` was a constant that the release's
ldflags overwrote with the runner's own number, so a released 0.28.0 runner in a project whose
vendored plugin was 0.26.0 recorded 0.28.0 — and section 5 wants the pair for one purpose, that the
frontmatter names what was used and the lock proves it with a hash, so a number which cannot
disagree with the runner can never be proved wrong. `runner_version` claimed `0.1.0-dev` through
twenty-nine minor releases, because a literal cannot learn the newest tag: `debug.ReadBuildInfo`
reports `(devel)` and only git knows it, at build time. The lock carried neither half of the proof,
though section 5 enumerates `plugin: { version, sha256 }` in it. Three defences against a swapped
plugin absent at once, which is why #183's resolution order could not be built: section 7 ranks
`--plugin-root` and `XENO_PLUGIN_ROOT` above the vendored tree, `internal/gates` reads the vendored
rules and templates, and a copy of this repository shows that a valid rule tree which differs leaves
`gate verify` at exit 0 over 273 verdicts while G-Policy silently stops judging all 75 phases that
recorded the previous hash — not two clones disagreeing, one of them stopping and reporting success.
In scope: the whole of #177 this repository can settle, and the parts of #183 that cannot make a
verdict depend on the environment — the entry point, `XENO_HARNESS` recorded into `tool`, and a
check that nothing branches on it. Out of scope with reasons: the resolution order, which waits on a
section 7 decision; `XENO_PLUGIN_DATA` read by the runner, since section 7 pins it to one value; the
manifest tracking the next release, which no mechanism here can do; and a statement in section 16.
