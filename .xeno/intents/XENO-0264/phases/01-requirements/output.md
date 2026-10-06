---
intent: github.com/triplem/xeno#258
phase: 01-requirements
created: "2026-10-06T17:29:58Z"
schema_version: "1.0"
runner_version: dev+d3983d3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f2c8e59caba230883270e549d931532a97da4bf92f0df2b415b8f12976045339
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Numbered, and P4's mapping cites the numbers — which this intent's own new checks will read, because
its P1 and P4 render at 1.1.0.

1. **`model.Option` carries `Reason`, with `omitempty`.** So no sealed artifact changes and nothing
   in the trail re-hashes.

2. **`QuestionAsked` refuses a recommended option with no reason.** On the option carrying
   `recommended: true`, not on the question. A question stating `no_options` is exempt, as it is for
   the recommendation itself.

3. **`QuestionShape` is unchanged, so the gate reads no more than it did.** Moving the recommendation
   half into the gate would fail every sealed question that predates it, which is the measurement
   `QuestionShape`'s comment carries and #229 settled.

4. **Both templates are at `1.1.0` and nothing else in either changes.** No section, no order, no
   required flag, so the anchors an artifact carries are identical before and after.

5. **G-Schema refuses a P1 artifact declaring `requirements@1.1.0` or later whose
   `acceptance-criteria` section has no numbered item.** Section 5 puts the requirement on the
   section's shape, and G-Schema is what reads an artifact's shape.

6. **G-Test refuses a P4 artifact declaring `verification@1.1.0` or later whose `test-mapping` does
   not name every criterion its P1 numbered.** Section 7 asks G-Test for "mapping of acceptance
   criteria complete", so that is where the completeness lands.

7. **Both checks read the declared template version, not the current one.** An artifact declaring
   1.0.0 is judged as it always was. All 130 P1 and P4 artifacts in the trail declare 1.0.0, so
   `gate verify` reports the same count and no verdict moves.

8. **The version comparison is numeric and not a string compare.** `1.10.0` is later than `1.9.0`,
   which a lexical compare gets wrong, and a check that silently stops applying at the tenth minor
   version is worse than none.

9. **Both checks fire on this intent.** Its P1 numbers its criteria and its P4 mapping cites them, so
   both pass here; and each is confirmed able to fail, by a test and by mutating the artifact.

10. **`testReport`'s comment is corrected.** It says the convention is "an addition to section 9",
    which is the Rule model, and carries "46 of 55 P1 artifacts", a figure superseded twice. It is
    replaced as a paragraph.

11. **The two rows of `docs/clause-readers.md` name their readers**, and the two paragraphs under the
    table say the clauses wait on nothing.

12. **The measured cost of the bump is recorded where a reader meets it**: in this intent's
    artifacts, in the pull request, and in the comment on `goneBundle`'s consequence.

13. **`go test ./...` passes, `go vet` is clean, `gofmt -l` prints nothing, `xeno gate verify` exits
    0** — and the verdict count is the one it was before, which is criterion 7's check as well.

<!-- xeno:section:non-goals -->
## Non goals

**Not #235's code.** `Advisory` on `model.Finding`, `result` ignoring an advisory finding and
`budget` marking its own are the other issue, under the other work package, on their own branch.

**Not the gate reading the reason.** `QuestionAsked` is the writer's check. The gate calls
`QuestionShape`, and the comment there carries the measurement: one sealed question in this trail
fails the recommendation requirement, so a gate that read it would re-judge the trail. #229 settled
that and this intent does not reopen it.

**Not restoring the `strings_hash` reader for an artifact declaring an older template.** There is no
way to tell a hash that is wrong from a hash that is right about a bundle the repository no longer
carries, and a check that guessed either way would be worse than the gap — a false red on every
historical artifact, or a false green that claims a reader it does not have. Recorded, not repaired.

**Not re-judging the trail.** Both checks key on the declared template version and all 130 P1 and P4
artifacts declare 1.0.0.

**Not a numbering convention stricter than section 5's.** "A numbered list" is what the clause says.
Consecutiveness, starting at one, a maximum, a minimum above one — each is an invention under the
second standing rule, and each would fail a legitimate artifact for a rule nobody wrote.

**Not a citation format for the mapping.** Section 5 says the mapping "names each criterion by its
number" and does not say how. The check reads the number as a standalone token anywhere in the
section, which accepts a table, a list and prose. A stricter reading would fail a format the clause
permits; that it also accepts a stray number is the cost, and P2 says so.

**Not a check on the P1 side from G-Test, or on the mapping from G-Schema.** One clause, one gate,
chosen by which section states the requirement. Putting both in one gate would make a green G-Test
mean something about an artifact two phases away.

**Not bumping any other template.** `intake`, `design`, `implementation` and `review` stay at 1.0.0,
because neither clause says anything about their sections and a bump costs each of them the
`strings_hash` reader for nothing.

**Not a change to any normative document.** Both clauses were committed in XENO-0262. This is the
code that follows, which is the order the first standing rule asks for.

<!-- xeno:section:constraints -->
## Constraints

**A template bump is not free, and the price is measured.** `goneBundle` in `hashes` is "the
artifact's declared template ref differs from the one the repository carries", and `recomputed` skips
`strings_hash` where it is true. So bumping `requirements` and `verification` removes that reader
from 130 sealed artifacts. Measured by corrupting one and watching the gate: red at 1.0.0, green at
1.1.0. It is one of two readers — `gate verify` still reports a divergence and exits 1, because
editing a sealed artifact moves its `artifacts_hash` and the successor's G-Freshness reads it.

**`template.Load` never consults the version an artifact declares.** It loads whatever
`template.yaml` is in the tree, so a bump changes what new artifacts declare and nothing about how
an old one renders or parses. That is why the bump is safe for the trail and why `goneBundle` is the
only place the version is compared.

**The version has to be compared numerically.** `Ref()` is `id@version` and the versions are
semver-shaped. A string compare puts `1.10.0` before `1.9.0`, so the check would stop applying at
the tenth minor version with nothing to say so.

**A P4 artifact's criteria live in another phase's file.** The mapping check has to read
`01-requirements/output.md` from the P4 gate, which `Ctx.phaseRel` allows. It also has to do nothing
where that file is missing or declares an older template, because a P4 at 1.1.0 behind a P1 at 1.0.0
is a legitimate state during the bump and for every intent already under way.

**`template.Parse` is the only honest way to read a section back.** It returns the rendered sections
by anchor, which is the form the artifact is in; matching headings would break the moment a strings
bundle changes, which is exactly what section 5 forbids a gate from doing.

**Both checks will judge this intent's own artifacts.** Its P1 and P4 render at 1.1.0, so the first
subject of a forward-only anchor is the intent that creates it. If either check is wrong about a
legitimate format, it goes red here.

**A new finding fails its gate in the ordinary way.** `result` returns `fail` for any finding and
#235's advisory field is not in this branch, so both checks stop a phase when they fire. That is
correct for these two — neither is a measurement against a number nobody has experience with — and it
is why the formats they accept have to be generous rather than tidy.

**`omitempty` on the new field, and nothing else in `Option` moves.** A sealed artifact's frontmatter
is inside `artifacts_hash`; a field that serialises only when set leaves every one of them byte
identical.
