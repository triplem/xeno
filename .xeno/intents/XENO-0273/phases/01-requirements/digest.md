---
intent: github.com/triplem/xeno#153
phase: 01-requirements
created: "2026-10-07T14:23:34Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 43bea369197e1537629b226d47cb453062c32d080bb1a8b8efd4754d42dcc774
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Ten criteria. Two are the page carrying the example: its first fenced block is byte-identical
to `example-symbols.yaml`, and a test fails when they part, checked by making them part. Four
are the page's prose: it says nothing the example does not, it says where the format is fixed
and that the two configuration keys are in Appendix A rather than section 5, it says a project
produces the index and Xeno ships no indexer, and it says an absent, stale, unreadable or
malformed index is not an error. Four are the rest: an entry in the index, the example
unchanged, no normative document changed, and the suite plus every link resolving.

The non goals are the restraint: no prose reference page, which #153's own last paragraph
argues against; no edit to the example, which is a fixture and was #151's deliverable; no move
out of `testdata`; no site; no link checker; no new facts; and no generated page, since WP16
owns that and what is here is the same stopgap `docs/commands.md` carries.

The binding constraints are the first standing rule, the rule that a fact is written down once,
and the rule about a negative result — which applies to the claim that the prose adds nothing,
so it is checked sentence by sentence against the example, section 5 and Appendix A.
