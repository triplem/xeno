---
intent: github.com/triplem/xeno#199
phase: 03-implementation
created: "2026-10-03T18:23:59Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55e31ce7419efcf8d06893f360bcded2857f8c5d84dd0735059076527216e27b
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

**`internal/plugin/embedded.go`, new.** `//go:embed all:embedded` with the reason for `all:`;
`shippedRoot` at `embedded/plugin`; `Shipped()` returning the sub-filesystem and whether it holds a
plugin, settled on `fs.Stat` of the manifest because `fs.Sub` succeeds on an absent path.

**`internal/plugin/embedded/PLACEHOLDER.md`, new and committed.** What the directory is for, that the
release writes `plugin/` beside it, that the order is stamp then embed then digest, and why it is not
gitignored.

**`internal/runner/init.go`.** `vendorPlugin` is one `fs.WalkDir` writing every regular file through
`createMode`. `pluginSource` returns the filesystem and a human-readable origin, preferring the
embedded copy and refusing where neither it nor `--plugin-from` is a plugin, with the refusal naming
the build rather than the file. `manifestRel`, the file that makes a directory a plugin.
`vendoredMode`, 0755 under `bin/` and 0644 elsewhere. `createMode` beside `create`, which now calls
it. `vendorTree`, `vendorFile` and `vendorRules` are gone — about sixty lines for twenty.
`InitResult.PluginFrom`.

**`cmd/xeno/main.go`.** `printInit` prints `plugin from <origin>` first, because the answer decides
whether G-Supply can pass.

**`.xeno/plugin/.claude-plugin/plugin.json`.** `0.30.0` to `0.0.0-dev`, validated with `claude plugin
validate` rather than assumed.

**`.github/workflows/release.yml`.** Three numbered steps before the ldflags line: stamp the version
into the manifest with inline `python3`; copy `.xeno/plugin/.` into `internal/plugin/embedded/plugin/`;
then compute the digest. The comment says the order is the whole of it and why a digest taken before
the stamp fails for every adopter rather than for the release.

**Tests, five, in `internal/runner/init_test.go`.** The vendored digest against the source's; the two
file sets compared as sorted lists; the entry point's mode; the refusal with no plugin anywhere; and
`PluginFrom` being set. The three existing vendor tests pass untouched.

**`ASSUMPTIONS.md`** A87 and A88, and A83's open clause closed. **`README.md`** the stamping and what
`--vendor` copies from.

**Measured end to end with a binary built the release's way.** A release in an empty repository
reports `plugin from this release`, writes 34 files, stamps `0.33.0` into the manifest, leaves
`bin/xeno-env.sh` at `rwxr-xr-x`, passes G-Supply against what it just wrote, and the adopter's first
artifact records `plugin_version: 0.33.0`.

<!-- xeno:section:deviations -->
## Deviations from the design

**One from the design, and it is an addition the design did not foresee.** `InitResult` had no field
for a note, so `PluginFrom` is a new field rather than a line in an existing one. The design said
"init prints where it vendored from" and the struct had no place to carry it.

**One from the issue as filed.** #199's "done when" says this repository's manifest "is never bumped
again". It is bumped once by this change, from `0.30.0` to `0.0.0-dev`, which is the bump that makes
the sentence true afterwards.

**Two sealed verdicts were rewritten and restored during the two intents before this one**, by a
`gate run` against a sealed phase. Not this intent's doing and recorded here because the same mistake
twice is a pattern rather than an accident: there is no read-only way to see what a gate would say
about a sealed phase, so the only safe way to look is to copy the tree first. That is this intent's
predecessor's learning and it held this time — every measurement above was taken in a copy.
