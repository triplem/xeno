---
intent: github.com/triplem/xeno#177
phase: 01-requirements
created: "2026-10-03T17:00:45Z"
schema_version: "1.0"
runner_version: dev+b54626e.dirty
plugin_version: 0.29.2
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dc26eac4a874edb2d5754dae74387e71f4d20d19bd3ec7eeeab5737731fa8438
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**`plugin_version` is the vendored plugin's own, read from its manifest.** Not a constant, and not
the runner's number: the field names the plugin an artifact was rendered from, which is what makes a
disagreement with the lock possible at all.

**Absent where no plugin is vendored**, and section 5 requiring the field in every process file is
what makes that a G-Schema finding — a repository with no plugin rendered from none, which is the
honest thing for its artifacts to say.

**The release sets the runner's version and not the plugin's.** One `-X` where there were two.

**The lock carries `plugin: { version, sha256 }`**, which section 5 already enumerates, absent
together where there is no plugin because a version with no hash is half a claim.

**The plugin tree hash is defined to the byte**, in the package that writes it, which is the
precedent Appendix B sets for `secrets_hash` and `rules_hash`. Every file at any depth, one line of
`<sha256 of normalised content>  <path>`, sorted by path in byte order, hashed.

**A development build's `runner_version` claims no version.** `dev+<commit>[.dirty]`, because a
literal cannot learn the newest tag and `0.1.0-dev` was wrong rather than stale.

**The manifest declares the newest tag**, 0.29.2, where it declared 0.1.0.

**The entry point normalises what section 7 lists and not `XENO_PLUGIN_ROOT`**, with the measurement
as the stated reason, and the plugin's hook calls it plugin-relative so an installing project runs
its own copy.

**`XENO_HARNESS` reaches section 5's `tool` field** and beats `agent.tool` where both are set;
where nothing is exported the project's declaration stands, which is what A35's second amendment
settled and this does not undo.

**Nothing branches on the harness**, checked in the job whose name is the required context: one
constant and one read of it, and no comparison against a harness name.

**The trail is unaffected.** Every sealed artifact keeps the header it was written with, `gate
verify` stays at exit 0, and the suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No `--plugin-root` and no `XENO_PLUGIN_ROOT`.** The measurement is the reason and it is in the
entry point, in A84 and in #183: with G-Supply unimplemented and the lock's block new, an override
would remove the only thing that could catch a swapped plugin, silently. Three shapes were put to
the maintainer and this one chosen.

**No `XENO_PLUGIN_DATA` read by the runner.** Section 7 says the value is always `.xeno/local/`, so
a read that must yield the constant five packages already use buys nothing. The entry point sets and
exports it.

**No mechanism keeping the manifest in step with the next release.** It is a literal in a data file
read at run time, only a release knows the number, and a release here cannot write a file. #177's
open clause.

**No statement in section 16.** #177 asks for the trail's version shapes to be explained there. That
is the process definition; the README carries it and says why.

**No G-Supply.** The lock now records what that gate would compare and nothing reads it back — a
field with a writer and no reader, which is the way round that leaves the question answerable later.

**No backfill.** Every artifact behind this records `0.1.0-dev` and a constant `plugin_version`,
sealed.

**No change to `agent.model`.** Only `tool` gains a second source, because only `tool` has one.

<!-- xeno:section:constraints -->
## Constraints

**A verdict is reproducible from the repository alone.** A42, and the property that decided the
scope: nothing here may make `rules_hash`, a rendered artifact or a verdict depend on the
environment.

**Absent is not a default.** A35, for `plugin_version` where no plugin is vendored, and for
`XENO_HARNESS_VERSION` where nothing reports it.

**Appendix B's reason applies to a hash it does not define.** "A hash described rather than defined
is how two implementations end up one byte apart," so the plugin tree hash is written to the byte in
the package that writes it, which is what the appendix says about the two it defers.

**Section 7 forbids branching on the harness** and the check that enforces it has to be in the
required job, not beside it, which is #186's lesson.

**A literal cannot learn the tag.** `debug.ReadBuildInfo` reports `(devel)` for the main module, so
a development build either maintains a number by hand or claims none.

**One `-X` per fact.** A release knows its own version and not the version of a plugin in somebody
else's repository.

**The hook runs what an installing project has.** Plugin-relative and never this repository's tree.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0, with the exit code
captured rather than piped.**

**One intent, one branch, two issues** — `177-the-plugin-says-its-own-version`, labelled wp7 and
wp11, based on main. A83, A84 and A85 skip A82, which belongs to the open #196 and would otherwise
collide.
