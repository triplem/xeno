---
intent: github.com/triplem/xeno#165
phase: 00-intake
created: "2026-10-01T16:22:26Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+46d4cb2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 623ad5f80509828133015f2559fd8231e45bd9d6f460ca7e515d8d8d0c6220ce
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Three pieces of WP4 are built and no rule has ever applied: `given/builtin/` is empty, so the gates are
green over nothing. The plan names four rules — traceability of deviations, migration notes for
interface changes, a rationale for new dependencies, a filled release notes section at P5 — plus three
examples nobody should inherit unasked and two hook templates. Two findings change the shape of the
piece. `vendorPlugin` copies only the templates, so a shipped set would never reach a project that ran
`xeno init`; that has been true since WP9 and went unnoticed because there were no rules to leave
behind. And three of the four rules cannot be `checked`: the six templates have no `interfaces` section,
no `migration-notes` section and nothing about dependencies, so those rules have nothing to read and
are `review` rules, which section 9 calls the honest admission that the rest needs a person. The fourth
needs `section-non-empty`, which does not exist and which nothing else covers either, since `Missing`
is read only by the next-step suggestion — section 9 scopes the section predicates to what the set
needs, so the absence is a gap in this package rather than a wall. And the set is the first text here
written to be acted on by somebody who did not build it, which two earlier reviews deferred to this
piece.
