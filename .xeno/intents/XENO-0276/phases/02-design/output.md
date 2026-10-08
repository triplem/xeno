---
intent: github.com/triplem/xeno#325
phase: 02-design
created: "2026-10-08T13:51:38Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3558d70f124ea46ee3ef5224556dabc92a24c0f6cfabb2d93e84d448d1138ea0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:impact -->
## Impact

One file, and the question is where the two lines go rather than what they say.

The command block runs in the order a phase does: `phase start`, the sections, `phase
finish`. The two writers belong between the sections and the finish, because the checklist is
written before the phase is judged and because the block is read top to bottom by somebody
carrying the phase out. The rules come first and the lens second, which is the order of the
sentence above them: a rule is answered, then a lens adds what no rule covers.

The paragraph that already describes the entries is where the "why two commands" sentence
goes, not into the command block. That paragraph is what a reader has just read when they
reach the commands, and the reason an entry has no rule id is a property of the entry rather
than of the command.

What is deliberately not added: a `--file` form. `review answer` and `review lens` take their
note as an argument, and showing a flag they do not have would send a reader to `--help` to
find out.

Nothing outside `.xeno/plugin/skills/xeno-review/SKILL.md` moves. The commands exist, the
gate reads what they write, and `internal/plugin`'s surface test is what holds the new lines
to the dispatch table.
