---
intent: github.com/triplem/xeno#258
phase: 03-implementation
created: "2026-10-06T17:42:01Z"
schema_version: "1.0"
runner_version: dev+d3983d3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 49bb67319331ddd883967a96a729d0e6c712750b6156f1b78568eca9e6f79235
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

Seven files. Two lines of template, one field, three checks, two test files and a document.

**`internal/model/model.go`.** `Option` gains `Reason string` with `omitempty`, and a comment
carrying section 8's clause and the cost it names — empty on every option but one.

**`internal/gates/gates.go`, `QuestionAsked`.** After the recommendation count, the option carrying
`recommended: true` must carry a non-blank reason. The `recommended != 1` branch now `continue`s, so
a question recommending none or three gets one finding and not two: the second would name a
consequence of the first.

**`internal/gates/gates.go`, `templateAtLeast`.** `id@major.minor.patch` compared on the three
numbers. A ref naming another id, or one that does not parse, is not at the version — a check cannot
read a version it cannot find.

**`internal/gates/gates.go`, `numberedCriteria`, called from `schema`.** On a P1 artifact declaring
`requirements@1.1.0` or later, the `acceptance-criteria` section must carry a numbered item. Nothing
else is required of the numbering.

**`internal/gates/gates.go`, `mappingComplete`, called from `testReport`.** On a P4 artifact
declaring `verification@1.1.0` or later whose P1 declares `requirements@1.1.0` or later, every
number the criteria raise must be named in `test-mapping`. `namesNumber` reads the number as a token
with no digit either side, so a table cell, a list item and a sentence all count and `1` does not
match inside `10`.

**`internal/gates/gates.go`, `testReport`'s comment.** Replaced. It said the mapping half "has no
reader and is not implemented here", that the convention was "an addition to section 9" — the Rule
model — and carried "46 of 55 P1 artifacts", a figure superseded twice.

**`.xeno/plugin/templates/requirements/template.yaml` and `.../verification/template.yaml`.**
`version: 1.0.0` to `1.1.0`. One line each, and nothing else in either file: no section, no order,
no required flag, so the anchors an artifact carries are identical.

**`internal/runner/exchange_test.go` and `cmd/xeno/main_test.go`.** The `askable` fixture and its
twin in the command test gain a reason on the recommended option, because seven existing tests
asserted a writer that now refuses them — the check working on its first contact with the suite.
Three new tests: a recommendation with no reason is refused, a reason on an unrecommended option
does not satisfy the clause, and a question stating `no_options` owes none.

**`internal/gates/criteria_test.go`, new.** Ten tests over the two checks and the comparison.

**`docs/clause-readers.md`.** Three rows move from reporting no reader to naming one; a fourth row is
new for section 5's numbered list. The two paragraphs XENO-0260 wrote are replaced by what reads each
clause, the loose-token reading and its cost, the measured cost of the bumps, and where the reason
lives.

## One deduplication on the way through

`renderedSections` and the section parser this intent needed are the same function with different
receivers. Rather than a second copy, `parseSections(abs string)` holds the body and
`renderedSections` delegates: a section read by a rule and a section read by a gate cannot come to
mean different things.

## Two figures found wrong while changing them

**The clause table had thirty-nine rows and the prose said thirty-eight**, before this intent touched
either. Counted by reading the rows. Which earlier addition went unremarked is not recoverable from
the file. With the new row the count is forty, and the paragraph says both the number and that the
previous one was already wrong — the shape XENO-0260 recorded about this very document a few hours
earlier.

**`testReport`'s comment named section 9**, which is the Rule model, the same misattribution
XENO-0260 corrected in `docs/clause-readers.md` and nothing corrected here.

<!-- xeno:section:deviations -->
## Deviations from the design

**Criterion 9 is not met: neither new check fires on this intent.** P1's criterion said both would,
because the templates would be at 1.1.0 by then. They are bumped in P3, and this intent's P1 was
rendered and sealed in P1 — at `requirements@1.0.0`. Its P4 will render at `verification@1.1.0` and
find a predecessor at 1.0.0, which is exactly the case `mappingComplete` returns nothing for and
which `TestAMappingIsNotJudgedAgainstAnOlderRequirementsPhase` covers.

**The intent that introduces a template bump cannot be judged by it, and that is structural.** A
template version lives in the plugin and the plugin is changed in the implementation phase, which is
after the requirements phase has rendered and sealed its artifact. So the first subject of a
forward-only anchor is never the intent that creates it; it is the next one. P1 wrote the criterion
without asking when in the phase order the bump would happen, which was answerable at the time.

**Seven existing tests failed on the new writer requirement, which is the check working.** The
`askable` fixture in `exchange_test.go` and its twin in `main_test.go` wrote a question whose
recommended option carried no reason, and `QuestionAsked` now refuses that. Both fixtures gained a
reason. It is recorded because a reader of the diff sees seven test files touched and should know
that none of them was a test weakened to pass.

**The `recommended != 1` branch gained a `continue` that P2 did not call for.** Without it a question
recommending nothing produces two findings — "recommends 0 of its options" and "recommends an option
with no reason" — and the second is a consequence of the first with no separate repair. It is a
change to existing behaviour in a branch P2 described only as gaining a requirement.

**One deduplication that P2 did not name.** `renderedSections` and the parser the new checks needed
were the same body. `parseSections(abs string)` now holds it and `renderedSections` delegates. P2's
impact section listed the functions that would change and not this one; a second copy of a section
parser is how a rule and a gate come to disagree about what a section contains.

**The clause table's row count was wrong before this intent and the paragraph now says so.** Thirty-
nine rows against a prose figure of thirty-eight. Found by counting the rows while adding one, which
is the only way it could be found, and it is the gap XENO-0260 recorded about this document earlier
today. The correction names both numbers rather than quietly writing forty.

**The first mutation of `templateAtLeast` caught the wrong case.** It was meant to prove the numeric
comparison by making it lexical, and the crude version also accepted `one.two`, so the test failed on
a malformed ref rather than on `1.10.0`. A test asserting the claim directly was added —
`TestTenthMinorVersionIsLaterThanTheNinth`, which also asserts its own premise that a string compare
would order them wrongly — and then a precise mutation was run against it. A mutation that fails for
the wrong reason is a green light nobody earned.
