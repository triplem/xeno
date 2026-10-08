---
intent: github.com/triplem/xeno#325
phase: 05-review
created: "2026-10-08T13:53:32Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a2f9ac3a4a7d90fcf39657256bf79beb2f0dfb0b75024448c56e5c7a397b5a30
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: 'The implementation is the design: two lines in the command block in the order the design named, and the reason in the paragraph the design chose over the block.'
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: 'No interface change: the commands named already exist and nothing about them moves. The skill catches up with the binary rather than the reverse.'
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: Nothing is added as a dependency.
    - result: met
      note: 'Architecture conformance lens: the skill is the plugin''s and the commands are the binary''s, and this change keeps that direction -- prose catching up with a surface, with a test holding it there. No gate, field or rule moves.'
      source: lens
---

# Review

<!-- xeno:section:release-notes -->
## Release notes

**The review skill names the two commands that write its checklist.** `xeno review answer`
and `xeno review lens` now appear in the P5 skill's command block, between the section writes
and `phase finish`, in the order a phase uses them: the rules answered first, then a lens
adding what no rule covers. Both resolve P5 themselves and take no `--phase`, which the skill
says because nothing else in that block is like that.

The paragraph that describes the entries gains the reason there are two commands rather than
one flag: a lens entry answers no rule and its arguments have no place for one, and a flag
could be passed beside a rule id — which would make the entry that keeps a lens out of
G-Policy's counted set the entry that answered a rule.

Nothing else changes. An agent following the skill now produces the structured entries the
gate counts instead of prose the gate cannot see.
