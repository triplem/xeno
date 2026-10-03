---
intent: github.com/triplem/xeno#195
phase: 00-intake
created: "2026-10-03T14:19:19Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+b54626e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dd6466b301c625dece60e0a2a0f287bcd31c658ce077a7e74e5f711bab22c72d
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

`learning.yaml` was the last artifact of this process that no command wrote, and one phase directory
shows what that cost:

    output.md      runner_version: 0.1.0-dev+74553ad.dirty
    digest.md      runner_version: 0.1.0-dev+74553ad.dirty
    gate.yaml      runner_version: 0.1.0-dev+74553ad.dirty
    learning.yaml  runner_version: 0.1.0-dev

Four artifacts, one phase, one directory, one build. Three say which binary wrote them and the
fourth says `0.1.0-dev`, because a person typed it.

**This is #179's argument one file over.** That issue gave `intent.yaml` a writer for exactly this
reason, in these words: "`created` is a timestamp a person types; `runner_version` and
`plugin_version` are copied from memory, so every `intent.yaml` in this repository records
`0.1.0-dev` while the artifacts beside it record `0.1.0-dev+<commit>.dirty`. Two strings for one
build in one directory." The sentence needs no amendment to apply here, and `learning.yaml` was not
in that issue's scope.

**It is inside `artifacts_hash`.** `KnownPhaseFiles` lists it and Appendix B excludes only
`gate.yaml` and `cost.yaml`, so the typed value is sealed with the verdict and a later correction
reads as a divergence. Every learning record in this repository is in that state, which is thirty
intents' worth.

**The writerless artifact is also the one the process cares most about keeping honest.** Section 10
routes a learning through a merge request against the rule set so that it takes effect after review
rather than on being noticed. The record is the input to that route, and it was the one file whose
header nobody could check against anything.

**It is the fourth of these found this morning.** A check that could not run before the merge, a
convention that could not survive it, a rule with no reader, and now an artifact with no writer. The
maintainer found this one by reading the frontmatter and asking why the exported harness version was
not in it — which it correctly is not, because section 5 scopes `tool_version` to the two files
produced in a session. What was wrong was the field beside it.

<!-- xeno:section:scope -->
## Scope

**In scope.** A command that writes the record: the four keys of section 10 as arguments, the header
from the runner, and the phase optional so that the intent level record `intent close` reads has the
same writer. The empty case said rather than left out. Entries accumulating, so a phase that learned
two things does not retype the first. The category and the four keys refused before anything is
read. The learning skill telling the agent to call it instead of writing the file, and the intake
skill with it.

**Out of scope, and each for its own reason.**

The content. No template for an observation and no suggested proposal. Section 10 routes a learning
through review precisely so that it does not take effect on being noticed, and a runner that drafted
the text would be proposing rules.

`plugin_version`. It reads `0.1.0-dev` in all four files, including the three the runner wrote,
because `model.PluginVersion` is a constant whose comment says there is no plugin yet to have a
commit — and `.xeno/plugin/.claude-plugin/plugin.json` now declares `"version": "0.1.0"`, which it
does not even match. That is #177 and it stays there.

Backfilling the thirty intents' worth of hand-written records. Their hashes are sealed and a record
rewritten to carry a version nobody stamped would describe a write that never happened.

A guard on a sealed phase. `SectionSet` has none, for a reason that holds here: the write makes
`gate verify` report a divergence and `phase finish` is what judges the phase again, so a refusal
would refuse the first half of a re-judgement.

The remaining writerless field. `secrets_hash` has a writer and `rules_hash` has a writer; after
this nothing in section 5 is left without one except `plugin_version`, which is a constant rather
than an absence.

<!-- xeno:section:context-rationale -->
## Why this context

Section 10 is read for the record's shape and for what it is for: the four categories, the four keys
of an entry, the empty record, and the route through a merge request against the rule set. The last
of those decides the scope — the content stays the agent's.

Section 5 is read for the field groups, twice. Once to confirm that `learning.yaml` carries the
common header and not the session fields, so `tool_version` is correctly absent from it and the
maintainer's question had a different answer than it looked. Once for the six fields the header is.

Appendix B and `model.KnownPhaseFiles` are read together to establish that the record is inside
`artifacts_hash`: the appendix excludes `gate.yaml` and `cost.yaml` and nothing else, so the typed
header is sealed with the verdict.

`internal/gates/gates.go`'s `learningFindings` and `learningEntries` are read for what G-Learning
already checks after the fact — the header fields, the two permitted body keys, the four entry keys,
the closed category set — because a writer that refuses the same things before the file exists is
the point, and because nothing about the gate should need to change.

`internal/runner/runner.go`'s `RecordAssumption` is read as the closest existing command: an
argument per field, a closed set refused with the set named, and a record appended to rather than
replaced. `SectionSet` is read for the split this follows and for its lack of a sealed-phase guard.
`common()` is read for the six fields the header is made of.

`internal/gates/gates.go` around `completeInReview` is read for the intent level record's path,
because `intent close` reads one and it needs the same writer rather than a second one.

#179 is read as the precedent and for its wording, which this intent reuses rather than restates.

Nothing outside the repository is needed.
