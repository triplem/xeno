---
intent: github.com/triplem/xeno#188
phase: 01-requirements
created: "2026-10-03T21:18:04Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 686062591ec40cb405416cf8ec3a2fd36b2f34d6c6cd75b5ea79d17b03750e94
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Q-1 is settled: the two writing commands and the figure, decided by the maintainer,
with the reasoning recorded as D-1 in this phase's frontmatter with `decided_by` on
it — the first `decisions` entry in this repository's trail and the thing #188's
done-when counts. The acceptance is that `question record` writes an `open_questions`
entry into the running phase's frontmatter and refuses, before writing, everything
`questionShape` would find afterwards: no key, a key already used in that intent,
fewer than two or more than four options, anything but exactly one free entry,
with `no_options` the exception section 8 names. `decision record` writes the
flat entry with the person on it, `--withdraw` being section 8's third exit, and
`--by` is never defaulted because A35's rule applies hardest to the one field the
register exists for. Neither command refuses a judged phase and neither touches
`gate.yaml`: a write after a verdict invalidates it exactly as a section write does,
which is the mechanism #216 already built. `intent status` marks an intent that
reached 05-review with no question and no decision in any phase, in the listing as
a row mark and in the one-intent form as a line after the table, computed once on
the walk `summarise` already does. It is a figure and not a finding: a phase with
nothing to ask is ordinary, an intent reaching the merge having asked nothing is the
measurement. Out of scope and each with its reason: strictness about `proposed_by`
against `decided_by`, which no section states; any quota; any change to G-Questions,
which is correct about what it checks; writers for `evidence` and `review_checklist`,
being #208 and WP7's; an intent-level file for questions, which section 8 rules out;
and backfilling the sealed phases.
