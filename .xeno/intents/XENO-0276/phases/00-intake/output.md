---
intent: github.com/triplem/xeno#325
phase: 00-intake
created: "2026-10-08T13:50:25Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 70a042ee1a779054b96b747fbd617fcb1d2f2a3a6133c8b50751a8fe301ae2f2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

> **github.com/triplem/xeno#325** — The review skill does not name the two commands that write its checklist
>
> `.xeno/plugin/skills/xeno-review/SKILL.md` lists the commands P5 is carried out with:
> `phase start`, three `section set` calls, `phase finish`, and `gate approve` and
> `gate override` for a finding. It names neither `xeno review answer` nor `xeno review lens`,
> which are the two commands that write the review checklist the phase is judged on.

The issue as `xeno phase start` read it, quoted rather than summarised.

Confirmed in the tree. The skill's command block runs from line 32 to line 36 and again at
41 and 42; `grep` for `review answer` or `review lens` over the file returns nothing. The
skill does describe the entries' shape correctly — "an entry from a lens carries `source:
lens` and no rule id" — which makes the omission harder to see rather than easier: the reader
is told what the entries are and not how to write one.

And the consequence is not cosmetic. G-Policy requires every review checklist entry to be
answered, so a phase whose checklist exists only as prose has no entries, the gate counts
nothing, and the phase passes having answered no rule. That is A73's shape exactly — a
required section carrying nothing is read by the next-step suggestion and by no gate.

This intent ran its own P5 through those two commands before writing a word of this, which is
how the omission was noticed: the skill was the wrong place to look them up.
