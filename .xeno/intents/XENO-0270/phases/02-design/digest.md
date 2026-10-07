---
intent: github.com/triplem/xeno#231
phase: 02-design
created: "2026-10-07T09:40:26Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 92f2ab415eaf7e0eeb82de72e8f2c1f7a29150f3fcca730f282ba5ee7529c4b8
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The README opens with section 1's problem rather than with the runner's parts, because a reader
who recognises the sentence knows in one line whether this applies to them. Four commands, not
six and not twenty-seven: `init`, `intent start`, `phase start`, `phase finish`, in the order
somebody runs them on a first intent. `gate run` is deliberately not among them, because a
reader who types it first learns the wrong thing about when a verdict happens.

The reference page carries `usage` and no prose of its own, and the test compares the fenced
block rather than the file, so the page can still say where its text comes from. The index is
`docs/README.md` and not `index.md`, because GitHub renders it when somebody opens the
directory and WP16 can rename it when there is a generator. It marks the two normative
documents in their own entries, not in a preamble a reader following a link never sees, and
groups the rest by the errand a reader arrives with.

The moved material goes to one page, verbatim, because it all answers one question: what this
repository's own trail shows and why it reads oddly. The no-network-call paragraph is dropped
rather than moved, since copying a normative sentence is what the index entry exists to avoid.

D-1 records the maintainer's instruction and what it corrected in the issue.

The alternatives name WP16's generator as the fix and this test as a guard, so that WP16 does
not inherit the guard as a decision.
