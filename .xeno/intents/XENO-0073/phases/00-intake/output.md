---
intent: github.com/triplem/xeno#73
phase: 00-intake
created: "2026-09-26T16:20:12Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+bcb6994
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: f501ae2c85cbedbcd4b7e82dfac7bab954b966e2321cf9109597ec3fde636af4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

G-Schema checked the four hash fields for presence, and since #75 for shape: sixty four
lowercase hex characters or the placeholder where nothing could have written the value.
It checked nothing about what the characters are. A hash of the right shape over the
wrong content passed, which is the case the trail rests on, because `context_hash` says
what a phase was produced from and a stale one says it about a file that has since
changed.

The claim the process makes is that a verifier recomputes what an artifact asserts. It
held for `artifacts_hash`, which `xeno gate verify` recomputes, and for the finding id,
which is derived. It did not hold for anything in an artifact's own frontmatter.

<!-- xeno:section:scope -->
## Scope

The two hashes Appendix B defines, recomputed and compared: `context_hash` against the
lock file beside the artifact, `strings_hash` against the bundle the phase rendered from
where the repository still carries that version. `secrets_hash` and `rules_hash` have no
writer and therefore nothing to compare against, which the appendix says rather than
leaving it to be noticed.

Not a new gate. G-Schema already reads the frontmatter and judges its shape, so the
value check sits beside the shape check, which is A51.

<!-- xeno:section:context-rationale -->
## Why this context

**The trail was measured before the check was written.** 36 artifacts carry a
`context_hash` that matches the lock file beside them, 2 carry the placeholder from M0,
20 carry a `strings_hash` that matches the bundle and 18 name `intake@0.1.0`, a version
the plugin no longer has. So nothing in the repository turns red, and that was known
rather than hoped for: a check that reddens the trail it is added to has to come with a
decision about the trail, and this one does not.

**G-Schema rather than G-Freshness.** Both are hashes and only one of them is about a
relationship between phases. What is compared here is a field against a file in the same
phase directory, which is G-Schema's subject; G-Freshness compares a phase against its
predecessor, and `context_hash` would have been the one thing it judged inside a phase.
A51 records the choice.

**The fixtures were the finding in miniature.** The runner fixture carried
`context_hash` as a repeated hex digit and the gates corpus carried `1111…`, both of the
right shape over content they did not cover, and every phase they built passed. They
carry the real hash now, computed from the lock file the phase was started with, which
is also why the fixture had to start the phase before writing the artifact rather than
after.