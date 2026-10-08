---
intent: github.com/triplem/xeno#325
phase: 01-requirements
created: "2026-10-08T13:51:01Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5f9b8f51b0c47c24c41f19bad2673859bbe6b438b2aaa573570ef3fd4e048143
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **The skill's command block names both writers.** `xeno review answer` and `xeno review
   lens` appear in `.xeno/plugin/skills/xeno-review/SKILL.md` with their arguments, in the
   order a phase uses them: the rules answered first, then a lens adding what no rule covers.

2. **It says why there are two commands and not one flag.** A lens entry takes no rule id,
   and the impossibility is structural rather than checked — that is the sentence a reader
   needs in order not to look for `--source` on the answer.

3. **The existing paragraph about `source: lens` stays.** It is correct, and this intent adds
   the writer rather than rewriting the description.

4. **The plugin tests still hold.** `internal/plugin`'s surface test requires every command a
   skill names to exist in the dispatch table, so the two new lines are held against the
   binary and a renamed command breaks the skill rather than drifting from it.

5. **Nothing outside the plugin changes.** No gate, no artifact field, no runner behaviour:
   the defect is a missing sentence and the fix is that sentence.
