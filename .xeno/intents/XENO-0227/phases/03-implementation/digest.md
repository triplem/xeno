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
---
`internal/model/identity.go` is new and holds the derivations as pure functions: `Tracker` with
Appendix A's three fields, `Qualified`, which splits `--for` at the last `#` and completes the
project and the host from the configuration, `host`, which reads `base_url` and strips a leading
`api.`, and `NextKey`, which takes the shared prefix and the number after the highest, padded as
that key is. `IntentStart` in the runner reads `project.yaml`, builds the id, derives the key where
none was given, refuses where the directory exists, and writes the record; `nextKey` lists the
intent directories for it and treats a missing directory as an empty one. `Tracker` moved onto
`model.Project` so that `EnforcementCheck`, the only existing reader of the block, stops declaring
the same three fields, and `projectConfig` is now a constant where the path was spelled out four
times. `cmd/xeno` gains the usage line, the `--for` flag, a table entry with `needsKey` false and
`cmdIntentStart`, which prints the key it chose and sets it so the suggestion is about the new
intent; `next.go` stops saying that no command creates one. Thirteen tests across three packages:
the derivations and their refusals as tables, the record field by field, its version against the
`gate.yaml` beside it, a given key, a second run, a missing issue, the no-tracker path, the listing
and the phase start, and the exit-code staircase at the surface. One deviation from the issue, which
is `--intent` being optional, and one consolidation the design did not require, which is the tracker
block having one shape.
