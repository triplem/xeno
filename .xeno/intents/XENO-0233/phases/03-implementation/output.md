---
intent: github.com/triplem/xeno#177
phase: 03-implementation
created: "2026-10-03T17:05:34Z"
schema_version: "1.0"
runner_version: dev+a18c1f3.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7f5c9523742a299d7c20ee59acfbe419119eed335414f673d631bfb777798e0c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

**`internal/plugin/plugin.go`, new.** `Dir`, the vendored directory. `Version`, reading
`.claude-plugin/plugin.json` and returning empty where there is no plugin, no manifest or no field.
`Hash`, walking the tree with `filepath.WalkDir`, skipping anything that is not a regular file, and
building Appendix B's stream with the descent that one omits. The package comment carries the
measurement that keeps the resolution order out; `Hash`'s doc comment carries the definition.

**`internal/model/model.go`.** The `PluginVersion` variable is gone, with a paragraph where it stood
saying what it was and why — a constant the release overwrote with the runner's own number.
`devVersion` is `dev`, with the reason: a literal cannot learn the tag, `debug.ReadBuildInfo` says
`(devel)`, and `0.1.0-dev` stood through twenty-nine minor releases. `Common.PluginVersion` and
`Intent.PluginVersion` are `omitempty`. `ContextLock.Plugin` and `LockPlugin` are new.

**`internal/runner/runner.go`.** `HarnessEnv`. `common()` reads `plugin.Version(r.Root)` at all three
call sites. `Start` writes the lock's plugin block where either half resolves. `agent()` prefers
`XENO_HARNESS` and falls back to `agent.tool`, including where `project.yaml` cannot be read at all.

**`.github/workflows/release.yml`.** One `-X`, with the reason in a comment above it.

**`.github/workflows/xeno.yml`.** "the runner does not branch on which harness it is": one mention
of the constant outside the tests, and no comparison against a harness name, both as greps because
what is claimed is an absence.

**`.xeno/plugin/bin/xeno-env.sh`, new and executable.** Sets `XENO_PLUGIN_DATA`, `XENO_HARNESS` and
`XENO_HARNESS_VERSION`, leaves anything already set alone, reads the client's own variables, and
says in a comment why it does not set `XENO_PLUGIN_ROOT`. `hooks/hooks.json` calls it through
`${CLAUDE_PLUGIN_ROOT}`.

**`.xeno/plugin/.claude-plugin/plugin.json`.** `0.1.0` to `0.30.0`.

**Tests.** Five in `internal/runner`; two in `internal/plugin`, one of them rewritten because it
asserted the hook's old command; `TestStampIsValidSemverBuildMetadata` became
`TestTheStampClaimsNoVersionAndItsMetadataIsWellFormed`, since its premise was the opposite of the
decision; the fixture gained a plugin; and #196's learning test was adjusted.

**Measured before the tests were written.** `plugin.Version` returns `0.30.0` and `plugin.Hash`
returns a value identical to a `sha256sum` shell pipeline over the same tree. A tree with no plugin
yields an absent field and a lock with no block. `./xeno version` reports `dev+<commit>.dirty`.

<!-- xeno:section:deviations -->
## Deviations from the design

**This intent's own P0 and P1 record `plugin_version: 0.29.2` and its later phases record
`0.30.0`.** The manifest was bumped to the newest tag, v0.29.2, early in the work; #196 merged while
this branch was open and cut v0.30.0, so the target moved and the manifest was bumped again. The
two sealed phases keep what they were written with.

That is the open clause of #177 demonstrating itself inside one intent, which is why it is left
rather than tidied: nothing keeps the manifest in step with the next release, and an intent whose
own phases straddle one is the shortest possible proof. Rewriting them would have removed the
evidence and described a write that did not happen.

**`TestStampIsValidSemverBuildMetadata` was replaced rather than amended.** Its comment said "a
version that is not parseable is worse than the constant it replaced", which is the opposite of what
the maintainer decided, so the test was rewritten under a name that says what it now checks: the
metadata after `+` is well formed and the part before it must *not* parse as a version.

**The upstream learning test was adjusted.** #196 compared a record's header against
`model.PluginVersion`, which no longer exists. It compares against `plugin.Version(f.root)` now,
which is the same assertion through the new source.

**A82 is skipped.** A83, A84 and A85 are numbered past it because A82 belongs to #196, which was
open when they were written and has since merged. Had they collided the rows would have had to be
renumbered after review, which is worse than a gap that is explained.

**One departure from the issue as filed.** #177's "done when" includes the plugin manifest and the
release agreeing, and a statement in section 16. Neither is done: the first has no mechanism this
repository can provide and the second is the process definition. Both are named in P4's gaps and the
issue keeps them.
