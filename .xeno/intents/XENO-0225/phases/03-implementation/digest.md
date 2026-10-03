---
intent: github.com/triplem/xeno#172
phase: 03-implementation
created: "2026-10-03T10:37:37Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3ace547907c89324f3f52248d6c7f55f55f9f96c2f6b83a9ae9c910e5b1e1ffd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
One function, three tests, two files, 55 lines. `links(c)` reads the profile from P0, walks its links
and reports one finding per declared document that is not in the tree, naming the component, the path
and the profile, with a next step saying to correct the path or take the link out because a declared
link is a claim about a file. A link with no `docs` declares nothing to find and is skipped. One line
appends it to `schema`, so both profile checks arrive through the gate section 5 names. The comment
carries the argument rather than the mechanics: the profile's one unambiguous error, the finding naming
the claim rather than the consequence, and the runner being right to skip the link in silence because
it records what the phase was given. The tests reuse #171's budget fixture with one more field. What is
not here is the byte count: section 5 writes the lock's `files` as a path and a hash, so a size is a
field the specification does not have, and the review carries the sentence a person would add. This is
the first implementation phase of twenty-eight with no deviation, for three visible reasons — one
condition, a criterion rather than a document as its specification, and a shape that already existed.
