---
intent: github.com/triplem/xeno#105
phase: 00-intake
created: "2026-09-27T13:34:26Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0b3a547.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 577a525f5ddfe1805ef87af1880810f67393ea4d6928b7b9a81c195cbca34826
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

G-Build filters declarations on `kind != "build"`. Section 4 fixes the closed set as
`test-report`, `coverage`, `build-log`, `scan`, `sbom`, `other`, and there is no `build`
in it, so the gate matches nothing a conformant artifact declares. A declared build
carrying `result: fail` passes G-Build. The gate is not weak, it is inert, and the
verdict it contributes to says a build was judged when nothing was.

The set is unchecked everywhere, which is why the mismatch could persist unnoticed. That
gap hides a second: section 4 requires `result` on `test-report` and `build-log`, because
G-Test and G-Build read it, and no gate checks that either. Limitation 8 of section 16
argues a missing `result` is unambiguous on exactly those two kinds, so an unchecked
`result` makes a stated limitation untrue rather than merely unenforced.

<!-- xeno:section:scope -->
## Scope

G-Build reads `build-log`. The closed set of `kind` and the requirement of `result` on the
two kinds that carry it become G-Schema findings, beside the enumerations that gate
already judges. The sets move to `internal/model`, where `LearningCategories` and
`AssumptionOrigins` are.

The fixtures move to the spec's spelling with the code. They carry `kind: test` today,
which is not in the set either, so they would fail the new check and would have hidden the
old one for as long as they stood.

Not G-Test, which stays `not-implemented`. It reads `result` on `test-report` and the
acceptance criteria mapping, and that is WP6's other half rather than this repair.

<!-- xeno:section:context-rationale -->
## Why this context

**The gate was tested against itself.** `build()`'s loop body is at nought per cent
coverage: no test has ever put a build item in front of the gate, and the tests that do
declare evidence use `kind: test`, which the specification does not define either. Code
and fixtures agreed with each other and neither agreed with section 4. That is the shape
A25 already met from the other side, where a promise about routing was checked at the
verdict rather than at the path, and the answer here is the same one: the closed set is
judged where an artifact is read, so a spelling nobody defined cannot pass through a gate
by being spelled consistently.

**A closed set is checked once, where the artifact is read.** G-Schema already judges the
enumerations of section 10 and section 8. Putting `kind` anywhere else would give the
evidence declaration two readers with two opinions, and G-Build would be judging shape
rather than reading a result, which section 7 says is what a gate does not do.

**`result` is required on two kinds and written where it exists on the rest.** Section 4
is explicit that the field follows the producer rather than the kind, so the check is not
that `result` is present but that it is present on `test-report` and `build-log`, and that
where it is present anywhere it reads `pass` or `fail`. A gate demanding it everywhere
would contradict the sentence that lets a bill of materials report nothing.

**What this does not buy.** G-Build still reads a declared result and does not judge a
build. Limitation 6 of section 16 stands untouched: whether the run happened, on which
commit, with which command, is what the trail adds, and the pipeline is what stops a
failing build at the merge.
