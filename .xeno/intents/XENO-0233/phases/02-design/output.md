---
intent: github.com/triplem/xeno#177
phase: 02-design
created: "2026-10-03T17:04:55Z"
schema_version: "1.0"
runner_version: dev+a18c1f3.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 596bd4101e3eeddc62a7257123f25d602347f3cf6db75010a4d0290e5946cbf7
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

**`internal/plugin` stops being a test-only package.** It held tests that check the shipped tree and
no code; it now holds the three things the tree can be asked — where it is, what version it declares,
and the hash over it. That is the package the question belongs to, and `internal/model` could not
answer it because `model` has no file access and is the layer that names things.

**`Dir` is the vendored directory and the resolution order is absent**, with the measurement written
into the package comment as the reason. A reader who comes looking for `XENO_PLUGIN_ROOT` finds why
it is not there rather than an omission.

**The tree hash is defined in the package that writes it**, which Appendix B's own words make the
precedent: "`secrets_hash` and `rules_hash` are defined by the package that first writes them." The
definition is in the doc comment, to the byte, because the appendix's reason for defining the others
applies unchanged — a hash described rather than defined is how two implementations end up one byte
apart.

**It descends, where `hashing.DirHash` does not.** A phase directory's subdirectory is evidence and
a plugin's subdirectories are the plugin, so the computation is Appendix B's with the descent that
one deliberately omits. Cross-checked against a `sha256sum` shell pipeline, which is how
`artifacts_hash` was checked.

**`model.PluginVersion` is removed rather than left unused.** It was the constant; leaving it would
leave something for the next ldflags line to set.

**`Common.PluginVersion` and `Intent.PluginVersion` gain `omitempty`.** Absent is not a default, and
section 5 requiring the field in every process file is what turns the absence into a G-Schema finding
rather than something to paper over.

**The fixture gains a vendored plugin.** Every repository has one after `xeno init --vendor`, so a
fixture without one was not a repository. The absence is asserted in its own test instead of being
the condition of every other.

**`devVersion` is `dev`.** A literal cannot learn the newest tag, so a development build claims none.
The release's ldflags is the only thing that knows the number and it still sets it.

**`XENO_HARNESS` beats `agent.tool`.** A project declares what it uses; the variable is what is
running, and a session under a harness the project did not name is the case the pair exists to make
visible.

**The entry point is a shell script in the plugin, called plugin-relative.** `${CLAUDE_PLUGIN_ROOT}`
and not this repository's tree, or an installing project runs a path that exists only here.

<!-- xeno:section:alternatives -->
## Alternatives

**Implementing section 7's resolution order.** Refused on the measurement, which is the whole
argument of this intent: `internal/gates` reads the vendored rules and templates, so an override
makes `rules_hash` depend on the environment, and A74 then makes the effect silence rather than
disagreement. Three shapes were put to the maintainer with the numbers and this one chosen. It is a
decision for section 7 and a person's commit, not something to build quietly.

**Keeping `plugin_version` as a constant and fixing only the ldflags.** Refused: the field would
still be a number nothing produced. The ldflags line was the second problem, not the first.

**Reading the plugin's version from `marketplace.json` instead.** Refused — that file is the
distribution wrapper and names the plugin's source, where `plugin.json` is the plugin's own manifest
and is what A77 established the client reads.

**Reusing `hashing.DirHash` for the plugin.** Refused because it does not descend, so it would have
hashed `secrets.yaml` alone and called it the tree. A hash over one file of a five-directory tree is
worse than no hash, because it looks like one.

**Leaving the plugin hash out until G-Supply exists.** Refused. Section 5 enumerates the block, the
appendix delegates the definition to the writer, and a field with a writer and no reader leaves the
question answerable later — where a reader with no field has nothing to answer it with.

**Maintaining `0.1.0-dev` by hand, or deriving it in a build script.** Both refused, and the
maintainer chose between them and claiming nothing: a maintained literal drifts, and a script means
the documented `go build` stops being the build. `dev` removes the lie instead of maintaining it.

**`XENO_HARNESS` losing to `agent.tool`.** Refused: that would make the project's declaration
unfalsifiable, which is the same defect `plugin_version` had.

**Having the runner read `CLAUDE_PLUGIN_ROOT` to find the plugin.** Refused by section 7's sentence
that the runner does not know which harness it runs under. The entry point reads it, which is what
an entry point is for.

<!-- xeno:section:impact -->
## Impact

**`internal/plugin/plugin.go`, new.** `Dir`, `Version` and `Hash`, with the package comment carrying
the measurement that keeps the resolution order out and the doc comment on `Hash` carrying the
definition.

**`internal/model/model.go`.** `PluginVersion` the variable is gone and its absence is explained
where it stood; `devVersion` is `dev` with the reason; `Common.PluginVersion` and
`Intent.PluginVersion` are `omitempty`; `ContextLock.Plugin` and `LockPlugin` are new.

**`internal/runner/runner.go`.** `common()` reads `plugin.Version(r.Root)`; `Start` writes the lock's
plugin block where either half is there; `agent()` prefers `XENO_HARNESS`; `HarnessEnv` is a new
exported constant beside `HarnessVersionEnv`.

**`.github/workflows/release.yml`.** One `-X` where there were two, with the reason.

**`.github/workflows/xeno.yml`.** A step checking that the harness is recorded and not branched on:
one mention of the constant outside the tests, and no comparison against a harness name.

**`.xeno/plugin/bin/xeno-env.sh`, new**, and `hooks/hooks.json` calling it through
`${CLAUDE_PLUGIN_ROOT}`.

**`.xeno/plugin/.claude-plugin/plugin.json`.** `0.1.0` to `0.30.0`.

**Tests.** Five in `internal/runner` for the plugin version, the lock's block in both states and the
harness precedence; two in `internal/plugin` for the entry point and the hook's command; the
version tests retargeted, since one of them asserted a semver shape the decision removes; the
fixture gaining a plugin; and #196's learning test adjusted, because it compared against the
constant.

**`README.md` and `ASSUMPTIONS.md`.** The two version shapes explained where a reader meets them,
and A83, A84, A85 — numbered past A82, which belongs to #196 and would otherwise have collided.

**No gate, no rule, no document.** `gate verify` reports 275 at exit 0 and every sealed artifact
keeps the header it was written with.
