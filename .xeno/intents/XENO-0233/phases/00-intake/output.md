---
intent: github.com/triplem/xeno#177
phase: 00-intake
created: "2026-10-03T16:59:17Z"
schema_version: "1.0"
runner_version: dev+b54626e.dirty
plugin_version: 0.29.2
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9003e4040aee2e86ef95618faf10339d3ae60bea39abfa5356539c0a8ed37453
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

Two version fields in the common header of every artifact, and neither said anything true.

**`plugin_version` was a constant that a release overwrote with the runner's own number.**
`model.PluginVersion = devVersion`, with a comment saying there was no plugin yet to have a
version — and there is one, with a manifest the code never read. The release's ldflags set it from
the same variable as `runner_version`, so a released 0.28.0 runner working in a project whose
vendored plugin was 0.26.0 recorded `0.28.0`. Section 5 wants the pair for one purpose: "the
frontmatter names what was used, `context.lock.yaml` proves it with a hash. Where the two disagree,
the hash wins and G-Supply fails." A number that cannot disagree with the runner can never be proved
wrong, so the field had no reader and no possible finding.

**`runner_version` claimed `0.1.0-dev` through twenty-nine minor releases.** The base was a literal
in the source, and a literal cannot learn the newest tag: `debug.ReadBuildInfo` reports `(devel)`
for the main module and the tag is known only to git, at build time. So every artifact in this
trail recorded `0.1.0-dev+<commit>` while the project shipped v0.29.2 — not a stale number but a
wrong one, in the field whose only job is to say which binary wrote an artifact.

**And the lock carried neither half of the proof.** Section 5 enumerates
`plugin: { version, sha256 }` in `context.lock.yaml`, which is the "proves it with a hash" half of
the sentence above. It was never written, so there was nothing for G-Supply to compare even if that
gate existed.

**Three defences against a swapped plugin, all absent at once.** That is why #183's resolution order
could not be built. Section 7 ranks a `--plugin-root` argument and `XENO_PLUGIN_ROOT` above the
vendored tree, and `internal/gates` reads `internal/rules` and `internal/template`, both of which
resolve from that tree. Measured on a copy of this repository: a valid rule tree that differs leaves
`gate verify` at exit 0 over 273 verdicts while G-Policy silently stops judging all 75 phases that
recorded the previous hash, because A74 judges a phase only against the set its own artifact names.
Not two clones disagreeing about a verdict — one of them stops judging and reports success.

**What section 7 does specify and nothing read.** `XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA` and
`XENO_HARNESS` are the normalised environment, and `grep os.Getenv` over `internal/` and `cmd/`
returned the enforcement token and the harness version added a week ago. The entry point that would
set them did not exist, so `XENO_HARNESS` reached no artifact and `tool` came from `project.yaml`
alone — a project's declaration of what it uses, standing in for what was running.

<!-- xeno:section:scope -->
## Scope

**In scope — the whole of #177 that this repository can settle.** `plugin_version` read from the
vendored plugin's own manifest, absent where no plugin is vendored. The release's ldflags no longer
setting it. The lock's `plugin: { version, sha256 }` block, already enumerated in section 5, with
the tree hash defined to the byte in the package that writes it. `runner_version` claiming no number
on a development build. The manifest bumped to the newest tag. And the trail's own shapes explained
where a reader meets them.

**In scope of #183 — the parts that cannot make a verdict depend on the environment.** The entry
point section 7 describes, shipped with the plugin and called by its hook, normalising
`XENO_PLUGIN_DATA`, `XENO_HARNESS` and `XENO_HARNESS_VERSION`, and reading the client's own
variables, which is the one place allowed to know which harness it is. `XENO_HARNESS` recorded into
section 5's `tool` field, beating the project's declaration where both are set. A check that nothing
branches on it.

**Out of scope, and each for its own reason.**

The resolution order, `--plugin-root` and `XENO_PLUGIN_ROOT`. The measurement above is why: with
G-Supply unimplemented, the lock's block new and nothing reading it back, an override would remove
the only thing that could catch a swapped plugin — and do it silently. Three shapes were put to the
maintainer and this one chosen; the order waits for a section 7 decision, which is a person's
commit.

`XENO_PLUGIN_DATA` read by the runner. Section 7 says the value is "always `.xeno/local/`", so
threading it through five packages to arrive at the constant they already use buys nothing. The
entry point sets and exports it, which is the half that matters.

The manifest tracking the next release. It is a literal in a data file read at run time, only a
release knows the number, and a release here cannot write a file — #121 removed the plugin that
would commit back to main. #177's open clause, and a decision between a check that is red from the
release until somebody bumps and the plugin versioning independently.

A statement in section 16. #177 asks for the trail's own version shapes to be explained where that
section lists what the record does not mean. That is the process definition; the README carries it
instead and says why.

Backfilling. Every artifact behind this records `0.1.0-dev` and `plugin_version: 0.1.0-dev`, sealed.

<!-- xeno:section:context-rationale -->
## Why this context

Section 5 is read for the two fields and for the sentence that explains why both are there — the
frontmatter names what was used and the lock proves it with a hash — and for the lock's own block,
which enumerates `plugin: { version, sha256 }` and so needs no document change to write.

Section 7's environment normalisation is read for the four variables and for the resolution order,
and the order is read twice: once for what it specifies, and once against `internal/gates`'
dependency graph, which is what makes it unbuildable here. The hooks subsection is read for the two
constraints on a hook, which is what a thin entry point is not.

Section 13 is read for the one shared version number across the three artifacts, which is the ask
the manifest bump answers and the next release will break again.

Appendix B is read for what it defines and what it does not. It defines `artifacts_hash`,
`context_hash` and `strings_hash`, and says of the other two that "`secrets_hash` and `rules_hash`
are defined by the package that first writes them" — which is the precedent for defining the plugin
tree hash in `internal/plugin` rather than leaving it described.

A74 is read as the mechanism behind the measurement: a phase is judged only against the rule set its
own artifact records, so a changed effective set produces silence rather than red.

A35 is read for absent-is-not-a-default, which decides what `plugin_version` does where no plugin is
vendored, and A81 for the precedent on precedence between a flag and a variable.

`internal/model/model.go` is read for the two version variables and `vcsStamp`, `internal/hashing`
for the primitives and for the non-descending `DirHash` the plugin hash cannot use,
`internal/runner/runner.go` for `common()` and `agent()`, and `.github/workflows/release.yml` for
the ldflags line that set both fields from one value.

`internal/plugin/plugin_test.go` is read because it asserted the hook's command, and the entry point
changes it.

Measured rather than read: `debug.ReadBuildInfo().Main.Version` is `(devel)`, `git describe` says
v0.29.2, and the plugin tree hash matches a `sha256sum` shell pipeline byte for byte.
