---
intent: github.com/triplem/xeno#325
phase: 03-implementation
created: "2026-10-08T13:52:05Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a96e165f89ee9b6678fa792e8dd20201e359f67ffe44078d65dcb5b421b5b247
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One file, `.xeno/plugin/skills/xeno-review/SKILL.md`, in two places.

**The command block gains the two writers**, between the sections and the finish:

    xeno review answer RULE --intent KEY --result R [--note TEXT]
    xeno review lens   --intent KEY --result R --note TEXT

with the sentence under them that the answers come before the lens, and that both resolve P5
themselves and take no `--phase` — which is true of these two and of nothing else in the
block, so a reader who copies the pattern would otherwise add one.

**The paragraph about the entries gains the reason there are two commands.** `review answer`
takes the rule it answers and `review lens` takes none, because a lens entry answers no rule
and its arguments have no place for one. A flag on the first could be passed beside a rule
id, and then the entry that keeps a lens out of G-Policy's counted set would be the entry that
answered a rule.

Nothing else moves. No gate, no field, no runner behaviour, and the paragraph that already
describes `source: lens` is unchanged because it was already right.

`internal/plugin`'s surface test passes, which is the part that matters: it requires every
command a skill names to exist in the dispatch table, so these two lines are now held against
the binary rather than being prose beside it.
