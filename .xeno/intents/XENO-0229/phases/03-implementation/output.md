---
intent: github.com/triplem/xeno#184
phase: 03-implementation
created: "2026-10-03T12:55:57Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1786f11f1345383fa4a4d23ced8b76490af7555f31177d9a55e859db4638914b
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

**The specification, already committed and alone.** `e8f68b1`: section 7's block gains
`XENO_HARNESS_VERSION`, recorded only, the four entries re-aligned to the longest name, and a
paragraph saying why a version belongs in a list of things recorded and never branched on. The
`XENO_HARNESS` paragraph is replaced rather than edited into, which is this project's convention for
a paragraph changed a second time. The plan's WP7 sentence names the same four.

**`internal/runner/runner.go`.** `HarnessVersionEnv = "XENO_HARNESS_VERSION"`, exported, with the
comment that it is the only `XENO_*` variable anything in the binary reads and that the three beside
it in section 7's list are specified and read by nothing, naming #183. `New` sets
`ToolVersion: os.Getenv(HarnessVersionEnv)`. The field's own comment is replaced: it said the value
is an input rather than something read, which was true for one commit and is now half the story.

**`cmd/xeno/main.go`.** `o.r.ToolVersion = o.toolVersion` becomes conditional on the flag being
non-empty, with the comment that an argument is somebody saying it about this run while the variable
is whatever set the session up. The flag's help names the variable it overrides. The usage's
`common:` block gains a second continuation line.

**`internal/runner/runner_test.go`.** `newFixture` calls `t.Setenv(HarnessVersionEnv, "")` before
constructing anything and delegates the construction to `reopen`, which is new and exists because
the one test about the variable sets it and then builds the runner again.
`TestTheHarnessVersionIsReadFromTheEnvironment` asserts both artifacts, reached through the variable.

**`cmd/xeno/main_test.go`.** `TestTheFlagBeatsTheHarnessVersionVariable` sets the variable to
`9.9.9`, writes one section with no flag and reads `9.9.9` back, then writes a second with
`--tool-version 2.1.276` and reads that back.

**`.xeno/plugin/skills/`.** The same paragraph in all six phase skills, naming both channels and
showing `export XENO_HARNESS_VERSION=2.1.276` beside the per-run form.

**`ASSUMPTIONS.md`.** A80 marked superseded in part, naming the half that failed. A81 for the two
channels, their order, and why the tests have to clear the variable.

**`README.md`.** Both channels, and the two test names in the WP7 row.

**Measured, not only tested.** A throwaway tree with the variable exported and no flag anywhere
carries `tool_version` in both artifacts of a phase; this intent's own six phases were written the
same way.

<!-- xeno:section:deviations -->
## Deviations from the design

**None from the design.** Every decision P2 recorded is what the code does.

**One from A80, and it is the point of the intent rather than a departure in it.** That row said the
better mechanism was out of the agent's reach and recorded the flag as a channel rather than the
channel. It is marked superseded in part instead of replaced, because the half about the design was
right and became this implementation unchanged; what failed was treating the document as fixed. A
row that was useful and partly wrong is more use amended than rewritten, and which half failed is
the thing worth keeping.

**One widening the specification commit took on.** The fourth name is longer than the three above
it, so all four lines of the block were re-aligned. That is three lines changed for a reason that is
typography, in a document the agent may not edit except by approval, and it is recorded here so that
a reader comparing the diff against the approval finds the reason rather than wondering.
