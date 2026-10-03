---
intent: github.com/triplem/xeno#199
phase: 00-intake
created: "2026-10-03T18:22:00Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7ca8bac9f0250f071c334a8e74051de9fd2160eb2954227c1096ec8b7e465d98
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

G-Supply shipped one change ago, anchored to a plugin a release cannot deliver, and `init --vendor`
has been copying an incomplete tree since before that.

**The anchor points at something the release cannot produce.** Section 13 puts the digest of the
plugin in the binary, the digest is over `.xeno/plugin/`, and the way that directory comes to exist
in a project is `xeno init --vendor` — which read `r.PluginSource`, a path on disk defaulting to
`.xeno/plugin` itself. A released binary in a fresh repository had nothing to copy from. The gate
was correct and unsatisfiable.

**And what it copied was not the plugin.** `vendorPlugin` named its parts: two trees one level deep,
two single files, and the rule set one level deeper. Two things were in the tree and in none of the
calls. `secrets.yaml`, which the function's own comment lists as part of section 13's `--vendor` set
and which `internal/secrets` reads. And `bin/xeno-env.sh`, the entry point added two changes ago,
which the plugin's own `hooks.json` names. Measured: 32 files vendored where the plugin has 34.

**Either omission alone breaks the gate for every adopter.** The digest covers the whole tree, so a
vendored tree missing two files cannot match, and G-Supply would have failed on every phase of every
project that installed the release. The gate shipped before the vendoring that makes it satisfiable,
and nothing caught it because the tests about vendoring listed the parts the code listed.

**The manifest was the other half of the same mechanism.** Section 13 asks for one shared version
number across runner and plugin. The literal had been bumped by hand three times in one session and
was overtaken by a release twice before the branch that bumped it merged — once inside an intent,
whose first two phases record one number and its later phases another. A release cannot commit the
number back to main, because #121 removed the plugin that would, for the reason `release.yml`
records: a push made with `GITHUB_TOKEN` reports no required check, so the required check is never
satisfied.

**What makes both solvable at once is that a release can write to what it ships.** It cannot write
to main and it does not need to. The copy an adopter receives is the only copy the coupling has to
hold for, and it is the same copy the digest has to be taken over.

<!-- xeno:section:scope -->
## Scope

**In scope.** The plugin embedded in the binaries, so `init --vendor` has a source that is the bytes
the digest was taken over. The release stamping the version into that copy, and taking the digest
afterwards. `init --vendor` walking the plugin instead of naming its parts, which fixes both
omissions by construction. `bin/` written executable, because an `embed.FS` cannot carry a mode. The
repository's own manifest set to `0.0.0-dev`, a placeholder that cannot drift. And a test comparing
the vendored tree's digest against the source's, so the next directory the plugin gains cannot
quietly go missing.

**Out of scope, and each for its own reason.**

A check that the manifest matches the newest tag. It is what the maintainer was offered and declined,
for the reason the npm audit baseline already records about a job that fires with no action
available: a check red from every release until somebody types a number the release already knows is
one people learn to route around. Stamping removes the need for it rather than enforcing it.

The resolution order. Still #183's open clause and still a section 7 decision. This change makes the
anchor deliverable, which is what that decision rests on, and it does not take it.

A signature over the shipped plugin. Section 13 disclaims one and says when that changes: "signing
comes with publication and not before."

Verifying that a published binary carries the digest of the tag's tree. The release computes it over
its own checkout, which is the right tree by construction, and a later check would need the
published artifact. Named in the verification phase of the intent before this one and still open.

`mcp.json`. Section 13's `--vendor` list names it and the plugin does not carry one, because a client
reading a declaration of a server that does not exist fails at startup. The walk copies whatever is
there, so it arrives the day it exists.

<!-- xeno:section:context-rationale -->
## Why this context

Section 13 is read for the three sentences this intent turns on: the shared version number and the
release together, the digest compiled into the runner, and "Nothing in the repository states what
the expected value is." The first is what stamping satisfies, the second is what embedding makes
deliverable, and the third is why the digest is taken after the stamp rather than before.

Section 13's `--vendor` list is read for what a vendored plugin is supposed to contain, which is
where `secrets.yaml` was already named and not copied.

`internal/runner/init.go` is read for `vendorPlugin` and the three walkers under it, for
`PluginSource`, and for `create`, which writes `0o644` and is why a mode has to be a parameter now.

`internal/plugin/plugin.go` is read for `Hash` and `Dir`: the digest is over paths relative to the
repository root, so a tree hashed anywhere else produces a different value, which fixes where the
stamping has to happen.

`internal/secrets/secrets.go` is read for `Shipped`, which points at the file that was not being
vendored.

`.xeno/plugin/hooks/hooks.json` is read for the path it names, which is the other file that was not
being vendored.

`.github/workflows/release.yml` is read for the ldflags line and for where three steps have to go
before it.

The Go documentation for `embed` is read for two facts that decide the design: a directive cannot
reach outside its own package directory, which is why the release copies the tree in rather than the
package pointing at it, and an `embed.FS` reports every file read-only, which is why the executable
bit is set by `init` rather than carried.

Measured rather than read: `claude plugin validate` accepts `0.0.0-dev`; `all:` is needed for the
manifest's dotted directory; and a vendored tree's digest equals the source's, which is the property
the whole change exists for.
