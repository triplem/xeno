---
intent: github.com/triplem/xeno#186
phase: 05-review
created: "2026-10-03T13:28:45Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6c4a1aa.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 64ccbc0f0f44c01357348c8187012014bf4a4eaf1a2f19e6122e1d1f9bd96454
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three shipped review rules are answered, and two of them earn their keep here.
`interface-change-needs-a-migration-note` is met because renaming three check contexts is a
migration: the protection setting must be applied after the merge, since a required context that
never reports blocks the merge that would create it, and two open pull requests report the old names
and need a rebase — a change correct in the diff and wrong in the sequence, so the sequence is in the
release notes, the commit and the description. `new-dependency-needs-a-rationale` is met rather than
not-applicable, because `actions/setup-node` is new to this repository and the reason is the engine
`semantic-release@25` declares. `deviations-are-traceable` is met with three. Beyond the rules: the
thing is not yet enforced and the review says so plainly — this intent makes enforcement expressible
and one command afterwards binds it, so until then the `CONTRIBUTING.md` rule is a rule and not a
mechanism. The baseline was read before it was raised, which is the rule that file sets about itself,
and the reading changed the answer: a newer release reports 30 where the old reports 35 and drops the
one advisory with supply-chain consequence, so the number moved because the toolchain moved. A hook
was asked for and refused on both of section 7's constraints, written into the design's alternatives
rather than answered in conversation only. And the investigation found two defects the issue did not
have, neither visible from the workflow that failed: one context required of five, and three sharing
a name.
