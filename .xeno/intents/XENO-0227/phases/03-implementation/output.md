---
intent: github.com/triplem/xeno#179
phase: 03-implementation
created: "2026-10-03T12:16:09Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5ea6e23d68f0b6364a25a8de901a8e01bec476814ed39e7fa5bf61bdaebbfe9b
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

**`internal/model/identity.go`, new.** `Tracker` with `adapter`, `project` and `base_url`;
`Qualified(issue)`, which splits the argument at the last `#`, completes the project from
`tracker.project` and the host from `base_url`, and refuses an id holding a space; `host()`, which
parses `base_url` and strips a leading `api.`; `NextKey(names)`, which takes the shared prefix and
the number after the highest, padded as that highest key is; and `splitKey`, which reads a name as a
prefix and a number and reports whether it is a key at all.

**`internal/model/model.go`.** `Tracker Tracker` on `Project`, with the comment saying why the
block has one shape rather than one per reader.

**`internal/runner/runner.go`.** `IntentStart(key, issue)` reads `project.yaml`, builds the id,
derives the key where none was given, refuses where the directory exists, and writes the record.
`nextKey()` lists the intent directories and hands the names to `model.NextKey`; a missing intents
directory is the same case as an empty one. `projectConfig` is a constant now, and the four readers
of that path in the file use it.

**`internal/runner/enforcement.go`.** The tracker block of its anonymous struct is `model.Tracker`.

**`internal/runner/next.go`.** The suggestion for an intent that does not exist names
`xeno intent start` and says that `assumptions.yaml` is still written by hand.

**`cmd/xeno/main.go`.** The usage line, the `issue` field on `opts`, the `--for` flag, the table
entry with `needsKey` false, and `cmdIntentStart`, which reports the error through `o.report` for
the staircase, prints the key and the id, and sets `o.key` so that the next-step suggestion is
about the intent that was just created.

**Tests.** `internal/model/identity_test.go`: four table tests over the derivations and their
refusals, including a Jira-shaped key, a self-managed deployment, the two key schemes side by side
and the two prefixes case. `internal/runner/runner_test.go`: a section of seven, covering the
derived record field by field, the version matching the `gate.yaml` beside it, a given key, the
second run, the missing issue, the no-tracker-block path, and the new intent being listed and its
phase starting. `cmd/xeno/main_test.go`: two at the surface, for the command running without
`--intent` and naming the key it chose, and for the staircase on a second run and on a missing
`--for`.

**`README.md`** gains the command, a paragraph on what it derives, and the test names in the WP7
row. **`ASSUMPTIONS.md`** gains A78 and A79 and a clause in "Built".

<!-- xeno:section:deviations -->
## Deviations from the design

**`--intent` is optional, where the issue made it required.** Recorded at intake, decided at design
and repeated here because it is the one place the code does not do what the issue says. The issue's
own argument — that a field the runner knows should not be typed — applies to the key, and the
sequence it would be derived from is on disk. `--intent` is kept for the case the derivation cannot
serve, which is a repository with no intents yet.

**Two files were touched that the design did not have to touch.** `model.Project` gained the
`Tracker` type and `enforcement.go` now reads it, rather than the new code declaring the same three
fields a second time. It is a consolidation and not a refactor: the fields, their tags and the
behaviour of `EnforcementCheck` are unchanged, and the existing tests of it pass untouched.

**`projectConfig` is a new constant.** The path was spelled out four times in `runner.go` and the
new reader wanted it twice, to read and to name the file in an error. The other two spellings, in
`init.go` and `enforcement.go`, are local constants already and were left alone.

**Nothing else deviates.** No gate, no rule, no field of section 5, no field of Appendix A, and no
existing intent.
