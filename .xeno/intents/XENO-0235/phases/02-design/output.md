---
intent: github.com/triplem/xeno#199
phase: 02-design
created: "2026-10-03T18:23:23Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0d268a4d8c7863265a55d18197a3bfff2c16ec75be64f01202ba8e091150295e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The release copies the tree into the package, rather than the package pointing at the tree.** A
`go:embed` directive cannot reach outside its own package directory, so this is not a preference. The
target is `internal/plugin/embedded/plugin/`, beside a committed placeholder rather than over it, so
the presence of that one path is what distinguishes a release from a build.

**The placeholder is committed and not gitignored.** `go:embed` needs the directory to exist at
compile time, and a clean checkout has to build.

**`all:` on the directive.** The plugin's manifest lives under `.claude-plugin/`, and embed skips
names beginning with a dot without it. Measured rather than assumed.

**`Shipped()` returns the sub-filesystem and a boolean**, and settles the question on the manifest's
presence rather than on the directory's, because `fs.Sub` succeeds on a path that is not there.

**`vendorPlugin` is one `fs.WalkDir` over whatever the source is.** Three functions become one, each
of which knew a tree name and a depth, and two of which had been silently incomplete. A walk cannot
go stale the next time the plugin gains a directory.

**The source is the embedded copy where there is one, the directory otherwise.** A release carries
the bytes its digest was taken over, so vendoring from a directory would produce a tree G-Supply
fails. A build that carries neither refuses and names `--plugin-from`.

**`create` gains a mode through `createMode`** rather than growing a parameter everywhere. One caller
needs a mode and the other twenty do not.

**`vendoredMode` is a function of the path**, 0755 under `bin/` and 0644 elsewhere, and not of the
source's mode — which does not survive an `embed.FS` and would therefore make a release and a
development build vendor different trees.

**`InitResult` gains `PluginFrom` and `init` prints it.** The answer decides whether G-Supply can
pass, so it is the one thing about vendoring worth saying on every run.

**The release stamps with `python3` inline.** The runner cannot do it — it would be a command that
edits the plugin, which is the one thing `--vendor` exists to copy — and the workflow already uses
`python3` for the audit comparison.

**The repository's manifest is `0.0.0-dev`.** The same answer A85 gives for `runner_version`, for the
same reason, and validated against the client rather than assumed.

<!-- xeno:section:alternatives -->
## Alternatives

**Keeping the list and adding the two missing entries.** Refused: it had been wrong twice, and the
second time was two changes ago by the same hand that wrote the list. A third entry would be correct
until the next directory.

**Pointing `go:embed` at `../../.xeno/plugin`.** Not available — a directive cannot leave its package
directory. This is the constraint that shapes the whole design and it is worth recording as a
constraint rather than a choice.

**Moving the plugin into the Go package and generating `.xeno/plugin/` from it.** Refused: A77
settled the distribution root at `.xeno/plugin/` against the installed client, the marketplace
wrapper points there, and this repository renders from it. Two copies in git, or a generate step
before every build.

**A release asset the adopter downloads.** Refused: it puts the network in `init`, and the plugin
would arrive from a different place than the digest was computed over, which is the mistake this
change exists to remove.

**Gitignoring the embed directory.** Refused because `go:embed` fails to compile when its target is
absent, so a clean checkout would not build.

**Stamping the version in the runner, at `init --vendor` time.** Refused: the digest is taken at
release time, so a version written later changes the tree the gate was anchored to. Stamp and digest
have to happen in that order and in one place.

**A check that the manifest matches the newest tag.** Offered to the maintainer and declined. The
reason is already in this repository, in the npm audit baseline's own comment: a check that fires
with no action available is one people learn to route around, and this one would be red from every
release until somebody typed a number the release already knew.

**Copying the source's file mode.** Refused: an `embed.FS` reports read-only for everything, so a
release and a development build would vendor trees that differ in mode, and one of them would have a
hook it cannot run.

<!-- xeno:section:impact -->
## Impact

**`internal/plugin/embedded.go`, new.** The directive, `shippedRoot`, and `Shipped()`.

**`internal/plugin/embedded/PLACEHOLDER.md`, new.** What the directory is for, that a release writes
`plugin/` beside it, and why it is not gitignored.

**`internal/runner/init.go`.** `vendorPlugin` is one walk; `pluginSource` chooses and names the
source; `manifestRel` and `vendoredMode` are new; `createMode` carries the mode and `create` calls
it; `vendorTree`, `vendorFile` and `vendorRules` are gone. `InitResult.PluginFrom`.

**`cmd/xeno/main.go`.** `printInit` says where the plugin came from.

**`.xeno/plugin/.claude-plugin/plugin.json`.** `0.30.0` to `0.0.0-dev`.

**`.github/workflows/release.yml`.** Three steps before the ldflags line, numbered and in order,
with the reason the order matters stated where somebody editing them will read it.

**Tests, five in `internal/runner/init_test.go`.** The vendored tree's digest against the source's;
every file of the plugin vendored, compared as two sorted lists so a failure names what is missing;
the entry point executable; a build with no plugin and no source refusing and naming
`--plugin-from`; and `PluginFrom` being reported. The three existing vendor tests pass unchanged,
which is the evidence that the walk copies what the list copied.

**`ASSUMPTIONS.md`.** A87 and A88, and A83's open clause closed — it said nothing kept the manifest
in step, which is no longer true.

**`README.md`.** What the stamping means for a reader and what `init --vendor` now copies from.

**No document change, no gate change, no change to the digest's definition.**
