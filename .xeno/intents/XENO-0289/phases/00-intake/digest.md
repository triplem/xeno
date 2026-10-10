---
intent: github.com/triplem/xeno#359
phase: 00-intake
created: "2026-10-10T15:36:51Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5b4221e73ab7771aaa9c890736b44a0e52630307b21a3215aec8e9316bb1a788
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
P0 fixes the intent as the page and nothing else. #359 asks for a page defining the
labels and the states they stand for, asks who acts on each and when it is removed, and
asks two questions about what becomes of `xeno-approved` when the work it authorised
stops. The page is writable now, because it records what is in use and defers to section
12 for the one label the runner reads. The lifecycle is a section 12 clause and therefore
the maintainer's commit, so the first question is on the issue with its options, their
consequences and their costs, and `xeno-needs-decision` is applied while it waits; the
second follows from it and is not asked beside it. The title's second half, trigger
labels, stays with #338, whose answer is a specification change nobody has made and whose
label does not exist. The scope is eleven files, ten of them extant and 64,727 bytes, with
the normative two and the assumptions page out and the four paragraphs the page needs from
them quoted in the intake. #359 carries no work package label and WP16 is the nearest fit,
which is a finding about the plan and is the phase's learning record rather than something
absorbed.
