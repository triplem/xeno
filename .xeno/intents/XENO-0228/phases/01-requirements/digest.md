---
intent: github.com/triplem/xeno#181
phase: 01-requirements
created: "2026-10-03T12:38:28Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6cbeac4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cc63c4f3640c6a10f79cb81d7b96825a959a6a72d884e7ed435ae67e5e5fe35f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`section set --tool-version V` writes the field into `output.md` and `phase finish` writes a digest
carrying the same value, taken from the artifact beside it rather than from a second argument. The
criterion the intent turns on is the second `phase finish`: `writeDigest` is rerun by every finish
and that rerun is where the hand edit was paid again, so a phase judged twice must lose nothing. A
reported version beats a recorded one, for a digest written before any section was; a later section
write that reports nothing does not erase what an earlier one recorded, because the field describes
the session and not the invocation; and no report at all is the field absent in both files, which is
A35 satisfied rather than contradicted — the row's reason is that a plausible value in a field
nobody produced is worse than an absent one, and a reported one is produced. The six phase skills
tell the agent to report it, since a flag nobody knows about is a field with no writer by another
route. Out of scope: a fourth `XENO_*` variable and the entry point behind it, both of which need a
change to a list section 7 enumerates and therefore a person's commit; harness detection, which
section 7 forbids; a `project.yaml` field, which would be stale by design; and any default, since
inverting A35 for convenience is the one way to make this worse than the edit it replaces.
