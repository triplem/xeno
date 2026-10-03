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
---
`internal/plugin` is new: `Dir`, `Version` reading the manifest and returning empty where there is
no plugin, and `Hash` walking the tree with the descent Appendix B's phase computation omits, with
the definition in the doc comment and the measurement in the package comment. `model` loses the
`PluginVersion` variable, with a paragraph where it stood saying what it was; `devVersion` is `dev`;
both headers gain `omitempty`; `ContextLock.Plugin` and `LockPlugin` are new. The runner reads
`plugin.Version` at all three call sites, writes the lock's block where either half resolves, and
prefers `XENO_HARNESS` over `agent.tool` including where `project.yaml` cannot be read. The release
has one `-X`. The `verify` job gains a step asserting the harness is recorded and not branched on,
as two greps, because what is claimed is an absence. The plugin ships an executable entry point,
called through `${CLAUDE_PLUGIN_ROOT}`, which says in a comment why it does not set
`XENO_PLUGIN_ROOT`. Everything was measured before the tests: `0.30.0` from the manifest, a tree
hash identical to a `sha256sum` pipeline, an absent field and no lock block without a plugin, and
`dev+<commit>.dirty` from the binary. Five deviations, the first of which is evidence rather than
untidiness: this intent's P0 and P1 record `0.29.2` and its later phases `0.30.0`, because #196
merged and cut a release while the branch was open — the open clause of #177 proving itself inside
the intent that documents it, left in place because rewriting it would describe a write that did not
happen.
