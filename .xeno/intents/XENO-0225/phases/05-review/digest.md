---
intent: github.com/triplem/xeno#172
phase: 05-review
created: "2026-10-03T10:40:54Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4153057120ef92ef429e5b32d645bef56f6933f8cbf55bd89ab56471eb571588
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three shipped review rules are answered. The criterion was met as written, quoted rather than
reinterpreted: a finding against the profile rather than a silently missing file, which closes the gap
#171 recorded and costs a second intent for one condition. The other half was honestly refused: the
byte count needs one clause in section 5's `context.lock.yaml` block — `files: [{ path, sha256, bytes }]`
with a line saying the size is recorded because the budget is judged against what the phase was given
and not against what the tree holds now — and the review carries that wording rather than a description
of it. The first decision after M0 found a home in 02-design, which cost nothing because it was about
this intent's own subject; the harder case is named in that phase's learning. What a reader should not
conclude is that a profile's links are validated: a document's existence is checked, the component is
not, and whether the document documents the component is a question section 5 forbids inferring.
Residual risk: half a link is validated, a document's existence is all this proves, the byte count
still moves after sealing, a link with no component reads badly, no profile exists here yet, and the
cost of this intent is the clearest record of what the sequence charges for a one-condition fix.
