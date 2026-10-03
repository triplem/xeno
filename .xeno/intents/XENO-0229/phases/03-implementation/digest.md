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
---
The specification is `e8f68b1`, committed alone and before this: section 7's block gains
`XENO_HARNESS_VERSION` with the four entries re-aligned to the longest name, the `XENO_HARNESS`
paragraph replaced rather than edited into, a paragraph saying why a version belongs in a list of
things recorded and never branched on, and the plan's WP7 sentence naming the same four.
`HarnessVersionEnv` is a new exported constant whose comment says it is the only `XENO_*` variable
the binary reads and that the other three are specified and read by nothing; `New` initialises
`ToolVersion` from it, and the field's comment is replaced because it said the value is an input
rather than something read, which was true for one commit. In `cmd/xeno` the assignment becomes
conditional on the flag being non-empty, the help names the variable it overrides, and the usage
gains a line. `newFixture` clears the variable before constructing anything and delegates to a new
`reopen`, which the one variable test uses to build the runner after setting it; the surface test
sets `9.9.9`, reads it back from an artifact written with no flag, then writes a second section with
`--tool-version 2.1.276` and reads that. All six skills name both channels and show the export. A80
is marked superseded in part, naming the half that failed, and A81 records the order and why the
tests must clear the variable. Three deviations, all recorded: none from the design, one from A80
which is the intent's purpose, and the block's re-alignment, which is typography inside a document
the agent may only touch by approval and is written down so the diff needs no explaining.
