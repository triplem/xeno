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
---
`internal/plugin/embedded.go` carries the directive with `all:` for the manifest's dotted directory,
`shippedRoot` at `embedded/plugin`, and `Shipped()` settling on `fs.Stat` of the manifest because
`fs.Sub` succeeds on an absent path. The placeholder is committed, saying what the directory is for
and that the order is stamp, embed, digest. `vendorPlugin` becomes one `fs.WalkDir` writing every
regular file through `createMode`; `pluginSource` prefers the embedded copy, names its origin and
refuses where neither it nor `--plugin-from` is a plugin, naming the build rather than the file;
`vendoredMode` is 0755 under `bin/`; the three old walkers are gone, about sixty lines for twenty.
`printInit` says where the plugin came from, because the answer decides whether G-Supply can pass.
The manifest is `0.0.0-dev`, validated with `claude plugin validate` rather than assumed. The release
gains three numbered steps before the ldflags line, with the comment saying why a digest taken before
the stamp fails for every adopter rather than for the release. Five new tests and the three existing
vendor tests passing untouched, which is the evidence that the walk copies what the list copied.
Measured end to end with a binary built the release's way: a release in an empty repository reports
`plugin from this release`, writes 34 files, stamps `0.33.0`, leaves the entry point executable,
passes G-Supply against what it just wrote, and the adopter's first artifact records `plugin_version:
0.33.0`. Three deviations: `PluginFrom` is a new field because the struct had nowhere to carry it;
the manifest is bumped once to make "never bumped again" true; and two sealed verdicts were rewritten
in the intents before this one, which is why every measurement here was taken in a copy.
