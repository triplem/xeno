---
intent: github.com/triplem/xeno#247
phase: 03-implementation
created: "2026-10-05T19:39:30Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a8e4506fd5b5135e62a53f3b2022ea4825236e835d4c939ab78b6092dc95c257
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

Three files, one of them new, and no code.

`docs/clause-readers.md` gains a row for section 8's third requirement:

| § | clause | reader |
|---|---|---|
| 8 | the recommendation carries a reason | **no field to carry it**; nothing reads it (#247) |

"No field" comes before "nothing reads it" because the first is why the second cannot be fixed by
a gate. The tool-requirement count moves 38 to 39 and the table counts 39.

The mapping row is untouched and still says `nothing`. The paragraph under the table gains a
sentence naming `examples/rules/mapping-is-complete.yaml`, what adopting it buys — a reader once
per intent rather than once per criterion — and why the reader column still says `nothing`: a rule
nobody has enabled fails nothing. That is the prohibition criterion 2 set, one row away from the
error #254 corrected.

`CLAUDE.md` gains one paragraph under Conventions, no new heading, nine lines. A decision put to a
person is put one at a time, with section 8's own reason, XENO-0243 as the instance, the
dependency as what makes a batch wrong, the free entry as the second reason, and "nothing checks
this (#248)". The file goes from 78 lines to 87.

`examples/rules/mapping-is-complete.yaml` is new, the fourth example rule, `scope: org`,
`binding: false`, `kind: review`, over P4 and P5. Its header carries the measurement — 46 of 55 P1
artifacts with no numbered criteria — what adopting it buys and does not, and A72's reason for
`examples/` rather than `given/builtin/`. Its statement follows section 9's sixth clause, which
bounds the mapping to completeness and says it is "not selection, and not existence".

**The example is inert and that was measured rather than reasoned.** `rules.Load` walks
`.xeno/plugin/rules/` and the project config, so nothing under `examples/` can reach the effective
set — and `gate verify` is at exit 0 over 408 verdicts, which is the 405 that existed plus this
intent's three judged phases, with `rules_hash` unchanged in every one. A grep confirms no artifact
in the trail names the new rule.

<!-- xeno:section:deviations -->
## Deviations from the design

No deviation from the design. The row's wording and ordering, the mapping row staying untouched,
the sentence under the table, the paragraph's placement under Conventions with no new heading, and
the example's scope, kind and header all landed as P2 specified.

One thing P2 named and this phase had to weigh again at the keyboard: the `CLAUDE.md` paragraph is
nine lines in a file whose last instruction is "Keep it short". It earns them by carrying the
reason rather than the rule — "ask one at a time" alone reads as etiquette, and what makes it a
rule is that XENO-0243's later options assumed an unanswered first question, which is checkable in
the trail. A three-line version was written first and said what to do without saying why, which is
the shape of convention this repository has repeatedly found gets ignored; the audit document
exists because of exactly that.

One correction inside the phase. The sentence under the table first read "what a project can do
about it" as a heading fragment and then repeated the figures the paragraph below already carries.
It was replaced whole and read back as a paragraph, which is what the conventions ask of prose
being changed a second time, and the figures stayed where they were.

No criterion was falsified this time, which is worth noting after the last two intents. Criterion
8 in particular held exactly as predicted: nothing under `examples/` reaches `rules.Load`, and the
408 verdicts prove it rather than the reasoning.
