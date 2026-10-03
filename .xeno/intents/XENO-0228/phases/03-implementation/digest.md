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
---
`ToolVersion` is a field on `Runner`, documented as an input because section 7 forbids the runner
asking a harness anything. `SectionSet` sets the frontmatter key only when the value is non-empty,
so the carry-over keeps what an earlier write recorded. `writeDigest` takes the reported value
first and `recordedToolVersion` second; that function is new, reads `output.md` through
`fm.ReadFront` into a one-field struct, and answers empty for a missing artifact, missing
frontmatter or missing field alike. `cmd/xeno` gains the flag, its field on `opts`, and the
assignment beside `EvidenceFrom`; the usage documents it on the `common:` block, which gained a
continuation line, because `parse` reads it for every command and two act on it — the `section set`
line would have had to give up its note about stdin to carry it inside 88 columns, which is the one
decision the design did not make and this phase did. All six phase skills gained the same five
lines. Six tests: five in the runner for both artifacts from one report, the second finish, a report
beating the record, absence in both files, and a later write not erasing; one at the surface driving
the flag through the command line. A35's closing sentence is replaced rather than edited into,
because the row had been amended twice and was already one field stale, and A80 is new. The change
is also demonstrated rather than only tested: all six phases of this intent were written with the
flag, none of their twelve artifacts was hand edited, and P0 went green on its first `phase finish`
for the first time in this repository.
