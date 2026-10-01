---
intent: github.com/triplem/xeno#156
phase: 04-verification
created: "2026-10-01T14:41:21Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5068ff7a029de2b6bdbc793100dd1bcaf4ad7f7bd28b16c3654692134ab51ce0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Every criterion of 01-requirements is a claim about a file, so each one maps to the command
that reads the thing the file describes rather than to a test. There are no tests for this
change, which is the point of the gap recorded below.

| Criterion | Checked by |
|---|---|
| "Not built" says only what is not built | `sed -n '/^## Not built/,$p' ASSUMPTIONS.md \| grep -ciE "secret filter\|digest writer"` |
| G-Supply and G-Secret carry what they wait on | read back as a paragraph against the gate table in `internal/gates/gates.go` |
| "Built" accounts for what postdates its list | the added paragraph against `cmd/xeno/main.go`, `internal/cost`, `internal/host`, `internal/index` and A62 to A65 |
| The command block is the command surface | the block's commands against the dispatch table, `diff` of the two sorted lists |
| The trail section is true of the trail | the phase directories counted, and the first six-phase key read off them |
| A row per package with tests, every named test existing | `grep -rl "func <name>(" --include='*_test.go'` per test, `ls` per package |
| Nothing regresses | `go build`, `go test ./...`, `gofmt -l`, `go vet ./...`, `./xeno gate verify` |
| The gap is written down once | #156 exists, carries the finding, and neither file absorbs it |
| The figures reach #117 | the comment on #117, with the `gate run` timing taken on a copy |

The two criteria with no mechanical check are the two that are judgments: whether the
paragraphs read as paragraphs, and whether the finding is stated once rather than absorbed.
Both are read back by a person, which is what 05-review is for, and neither is claimed as
verified here.

<!-- xeno:section:results -->
## Results

**The register's "Not built" no longer names the built.** The grep over the section returns
0 for both "secret filter" and "digest writer".

**The command block equals the dispatch table, with one expected difference.** The sorted
list from the README and the sorted keys of `commands` in `cmd/xeno/main.go` differ in one
line: `version`, which the README carries and the table does not, because `run` answers it
before dispatch. Every other command and every flag matches, and nothing in the README is
absent from the binary.

**The trail is fifty-one intents of one phase and twenty of six**, plus this one, which shows
as four while it is running. The first six-phase key is `XENO-0107`, which is the boundary the
paragraph names.

**Every test named in a new row exists.** `TestNoProfileRecordsNoInformationBase`,
`TestSectionSetWritesTheContextHashOfTheLockBesideIt`, `TestFinishWritesTheDigestFromTheSummary`
and `TestTheDigestIsFilteredAndSaysWhichFilter` in `internal/runner/runner_test.go`;
`TestExportPrintsTheEnvironmentAndNothingElse` in `cmd/xeno/main_test.go`;
`TestNoIndexBlockMeansNoIndex` and `TestTheIndexPathIsRelativeToTheRepository` in the same
runner file. Every package named as a row's test carries test files: `internal/host` one,
its two adapters one and two, `internal/cost` one, `internal/index` one, `internal/secrets`
two, `cmd/xeno` two.

**The suite is clean.** `go build` succeeds, sixteen packages report `ok` or no test files
with no `FAIL`, `gofmt -l` outside `vendor/` lists nothing, `go vet ./...` is silent.

**`gate verify` is green over the whole trail**: 175 verdicts, exit 0, three runs at 240, 269
and 219 ms. The count was 171 before this intent and the four new phases account for the
difference.

**`gate run` was timed on a throwaway copy**, as the standing request requires, at 5, 3 and
3 ms for 03-implementation. The working tree after the measurement shows only the two
corrected files and this intent's own directory, so no sealed phase was rewritten to take a
timing.

**Each phase of this intent was judged green on its own artifacts**, 00 to 03, with G-Supply,
G-Secret, G-Rules and G-Policy reported as `not-implemented` throughout, which is the honest
state and not a pass.

<!-- xeno:section:gaps -->
## Gaps

**Nothing here is a test, so nothing here runs again.** Every check in the results section is
a command somebody ran once, by hand, today. The next change to `cmd/xeno/main.go` will not
fail because the README's block no longer matches it, and that is the same property that let
the trail paragraph go false for twenty intents. This intent closes the instance and leaves
the class open, by decision rather than by oversight.

**The two judgments are unverified by anything mechanical.** Whether the three replaced
paragraphs read as paragraphs, and whether the ownership finding is stated once rather than
absorbed into a file, are read by a person in 05-review. Neither can be asserted here and
neither is.

**The counts are true as of this commit and not after it.** Fifty-one and twenty, 175
verdicts, sixteen packages: each is a measurement, and the paragraphs that carry a number into
either file were deliberately written to carry a boundary instead. The numbers in these
artifacts are allowed to age because an artifact is a record of a moment and the README is
not.

**`gate run` was timed on 03-implementation of this intent rather than on a phase typical of
the trail.** Three milliseconds is the cost of judging this phase's artifacts on a copy of
this tree, and a phase with more findings, more evidence or a larger information base would
cost more. The figure is useful as an order of magnitude against the section count and should
not be read as a per-phase constant.

**G-Supply, G-Secret, G-Rules and G-Policy did not run.** Four of the thirteen gates are
`not-implemented`, so this intent's green means nine gates passed. For a change to two
Markdown files that is close to the full relevant set, since the four unimplemented ones
concern the plugin, the hook and the rule set, but the verdict should be read as nine of
thirteen rather than as unqualified.

**The figures for #117 are derivable only.** Cost per intent still has no writer in a form
these artifacts can carry — #65 is the record of that — so what reaches #117 is sections,
words, lines and timings, and the token figure the plan's proportionality question actually
wants is absent again.
