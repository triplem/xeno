---
intent: github.com/triplem/xeno#235
phase: 03-implementation
created: "2026-10-06T15:22:22Z"
schema_version: "1.0"
runner_version: dev+8decfdc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 42ec6f58a70c0df11e04c8e81e9756271cc99e29e1ecc08738fa299abfa9bd50
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

One file, `docs/process-definition.md`, three places in section 5, and no code anywhere. Sixteen
lines added and one replaced.

**The `gate.yaml` findings block carries the key.** Between `next` and `decision`:

    advisory: true                  # absent in the ordinary case

Aligned at column 40, where the block's two existing comments start, and saying when the key is
absent, which is what both of them say.

**Two paragraphs after the status derivation.** The first says what an advisory finding is — the
check carrying it is `pass`, the phase is `green`, and the finding is in `gate.yaml` with its id,
its cause and its remedy like any other — and why the exception exists: blocking against a number
nobody has experience with would be the wrong way round, and a check that fires with nothing to do
about it is one people learn to route around. The second bounds it: a finding is the thing that
fails, an advisory one is readable only because it is rare, the clause that asks for one says so
where its check is described, and nothing else writes the field.

**One sentence in the Context economy subsection**, inside the paragraph that carries the budget
clause: "The finding is `advisory`, so G-Schema stays `pass` and the phase stays green." It is the
one line that makes the clause and its mechanism meet, which is the fault being repaired — they
have lived two hundred lines apart since the document was written.

## What is not here

No `Advisory` on `model.Finding`. No change to `result` or `budget`. No tests. No fifth check
result. No touch to `drift` or to section 16's ninth limitation. No row in
`docs/clause-readers.md`. `git diff --stat` names one file outside the trail.

## One departure from the drafted wording

**The paragraph is placed after the four-values paragraph, where the draft said "after 'Decisions
sit on findings, not on phases'".** The status derivation runs over two paragraphs and the second
completes the first — it is what explains why `red`, `approved`, `overridden` and `green` are not
interchangeable. Taking the draft literally would have inserted a clause about a finding that does
not fail into the middle of the argument about findings that do.

The prose is committed as drafted. The placement is the document's.

That is the third time in two intents that wording drafted in an issue comment needed adjusting on
contact with the file, after the section 8 heading that would have deleted a sentence and the
section 5 text that cited this repository's trail. A draft reads as prose in a comment and as a
position in a document, and the second is only visible once it is there.

<!-- xeno:section:deviations -->
## Deviations from the design

**The placement differs from the draft, and the prose does not.** Recorded in the changes section
with the reason: the draft's "after 'Decisions sit on findings, not on phases'" splits a
two-paragraph argument. P2 decided the wording would be committed as drafted, and it was; what the
draft said about where it goes was wrong about the document rather than about the clause.

**Three for three, which is now a pattern rather than three accidents.** XENO-0262 found that the
drafted section 8 heading would have deleted an existing sentence, and that the drafted section 5
text cited this repository's trail where the specification cites no repository at all. This intent
found a placement that splits an argument. All three were found by putting the text into the file
and reading it there, and none of them was visible in the issue comment the person approved. The P3
learning proposes that a drafted specification sentence is placed in the document and read in place
before it is offered for approval, so that what the person approves is what lands.

**One commit, as XENO-0262 settled.** The first standing rule asks a specification change to be
"its own commit"; every merge here is a squash, so two commits on a branch arrive as one and the
property is one of `main`. This commit carries one specification change and no code, which is the
substance of the rule. No criterion here asked for two, because P1 was written after XENO-0262 had
measured it — the one place in this session where a phase learned from the previous intent rather
than repeating it.

**The specification now describes a key no runner writes.** `advisory` is in no struct and nothing
reads it. Deliberate, and the second such key this session put into the document after
`requirements@1.1.0`. Named so that a reader finds it a sequence rather than an oversight: the
document leads its implementation by one intent each time, and nothing bounds how long.
