---
intent: github.com/triplem/xeno#181
phase: 03-implementation
created: "2026-10-03T12:39:46Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6cbeac4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5af59f78785a653b27a11a61adbef33b725fbdd8f3b787f228e870a0fc816b61
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

**`internal/runner/runner.go`.** `ToolVersion string` on `Runner`, documented as the harness
reporting itself and as an input rather than something read, with section 7's prohibition as the
reason. `SectionSet` sets `front["tool_version"]` when the value is non-empty and leaves the key
alone otherwise, so a later write that reports nothing keeps what an earlier one recorded.
`writeDigest` takes the reported value first and `recordedToolVersion(key, phase)` second.
`recordedToolVersion` is new: it reads `output.md` through `fm.ReadFront` into a one-field struct
and returns empty for every failure, which is the field absent.

**`cmd/xeno/main.go`.** `toolVersion` on `opts`, `--tool-version` registered in `parse` with the
comment that it is the one field of section 5 the runner cannot know, and
`o.r.ToolVersion = o.toolVersion` beside the two assignments already there. The usage gains a
continuation line under `common:` naming the two commands that act on it; the `section set` line
keeps its note about reading stdin, which is what it would have had to give up to carry the flag
inside 88 columns.

**`.xeno/plugin/skills/`.** The same five lines in all six phase skills, before "What the gate
refuses", naming the flag, why the runner cannot know the value, and that the digest takes it from
the artifact. `xeno-learning` has no `section set` and is untouched.

**Tests.** `internal/runner/runner_test.go` gains a section of five and a `set` helper:
`TestTheReportedToolVersionReachesBothArtifacts`,
`TestASecondFinishKeepsTheToolVersionWithoutBeingToldAgain`,
`TestAReportedVersionBeatsTheRecordedOne`, `TestWithoutAReportTheFieldStaysAbsent` and
`TestALaterSectionWriteDoesNotEraseTheVersion`. `cmd/xeno/main_test.go` gains
`TestTheToolVersionFlagReachesBothArtifacts`, which drives `section set` with the flag and
`phase finish` without it, plus `writeTemp` and `frontOf`.

**`ASSUMPTIONS.md`.** A35's last sentence replaced rather than edited into: `rules_hash` has had a
writer since the shipped rule set, so the claim that two fields remain was already stale, and the
third amendment records that the row is satisfied rather than contradicted. A80 is new.

**`README.md`.** The flag in the surface list, a paragraph on it, and the WP7 row extended.

**Demonstrated rather than only tested.** All six phases of this intent were written with the flag
and none of their twelve artifacts was hand edited. P0 went green on its first `phase finish`, which
is the first time that has happened in this repository.

<!-- xeno:section:deviations -->
## Deviations from the design

**None from the design.** Every decision P2 recorded is what the code does.

**One departure from the issue's own framing, and it is a narrowing.** #181 says A35's row is one
field out of date and this implementation acts on that: A35's closing sentence is replaced, not
annotated, because the project's convention for a paragraph changed a second time is to replace it
and read it back as a paragraph. The row had already been amended twice and misread twice.

**One thing the design did not say and the code had to decide.** The usage block could not carry
`[--tool-version V]` on the `section set` line inside 88 columns without dropping that line's note
about reading stdin. The flag went to the `common:` block instead, which is also the truer place:
`parse` reads it for every command and two act on it. Recorded here because a reader comparing the
usage against the design would otherwise find the flag somewhere the design did not put it.
