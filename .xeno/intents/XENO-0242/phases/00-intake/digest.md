---
intent: github.com/triplem/xeno#188
phase: 00-intake
created: "2026-10-03T21:16:07Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4963a3601877c4ea1fde451f9cd88cd5eabdd3c473fba3f35a3aab488b0c5fd4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Section 5 carries open questions and decisions structurally, section 8 gives a question
three exits through a person, G-Questions checks them from P5, and the loop has never
once run end to end: four intents raised a structured `open_questions` block, no
intent has ever written a `decisions` block, and `decided_by` does not occur anywhere
under `.xeno/`. Two things keep it that way. G-Questions is vacuous where nothing was
asked, because it verifies that a raised question is resolved and nothing declares
how many there should have been, so the greenest possible trail is the one that asked
nothing; and the fields have no writer, so a decision taken in conversation lands
wherever prose is cheapest, invisible to a query and with nobody's name in the field
built for it. In scope, after Q-1: `xeno question record` and `xeno decision record`
writing the two frontmatter blocks the way `assumption record` writes the register,
refusing what the gates refuse, and one line in `xeno intent status` for an intent that
reached P5 with no question and no decision in any phase — a figure, not a finding,
because a phase with nothing to ask is the ordinary case. Out of scope: strictness
about `proposed_by` against `decided_by`, which no section says today and which the
first standing rule therefore makes the maintainer's commit; any quota or finding
for a phase that asked nothing, which would be answered with invented questions;
the sealed phases of XENO-0227 to XENO-0229 and the hand-kept D-1 to D-7; and the
two other unwritten frontmatter blocks, `evidence` being #208 and `review_checklist`
being WP7's. Q-1 is raised in this phase and settled in P1 by a decision carrying
`decided_by`, because the done-when of #188 is that the count of that field in the
trail is not zero.
