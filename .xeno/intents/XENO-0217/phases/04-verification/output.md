---
intent: github.com/triplem/xeno#158
phase: 04-verification
created: "2026-10-01T15:15:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2ea9174303a58847db930b3d5825924e254a41f7f6b8cced3a4c6ccdf4dcad33
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Every criterion of 01-requirements maps to a test, because every one of them is a tree and a
verdict over it. The fixtures are built in the test rather than read from `testdata`, which is
what the rest of this repository does today and is noted as a deviation from WP17's corpus.

| Criterion | Test |
|---|---|
| A repository with no rule tree is green | `TestNoRuleTreeIsNoRulesAndNoProblems`, `TestNoRuleTreePasses` |
| A well formed tree resolves in precedence order | `TestPrecedenceIsSpecificityThenGivenOverLearned`, `TestAWellFormedTreePasses` |
| `binding` beats specificity | `TestBindingBeatsAMoreSpecificRule` |
| Two binding rules on different levels stay red and unresolved | `TestTwoBindingRulesOnDifferentLevelsStayUnresolved`, `TestABindingCollisionReachesTheVerdict` |
| `scope` against the path is red | `TestTheConfigurationErrorsSectionNineNames/scope_against_the_path`, `TestAMisplacedScopeIsRedAndNamesFileCauseAndNextStep` |
| `binding` outside `given/`, and under `given/project/`, as two findings | the same table's two binding cases |
| `learned/builtin/` is red, on the directory | `TestLearnedBuiltinIsRejectedOnTheDirectory` |
| `checked` without a `check`, `review` with one | the same table's two kind cases |
| A duplicate id at one level is red | `TestADuplicateIdAtOneLevelIsAProblem` |
| `abstract` required above the project level | the same table's abstract case |
| `rules_hash` is byte exact and reproducible | `TestTheHashCoversTheContentAndNotTheFiles`, `TestKeyOrderInACheckDoesNotReachTheHash`, `TestOnlyTheEffectiveSetReachesTheHash` |
| Its definition is checked against a pipeline, not against itself | `TestTheHashIsTheSha256OfTheRendering`, and the rendering's shape in `TestTheRenderingIsOneSortedLinePerRule` |
| The runner writes the hash it computed | `TestSectionSetWritesTheHashOfTheEffectiveRuleSet` |
| G-Rules is no longer `not-implemented` | `TestGRulesIsNoLongerReportedAsNotImplemented` |
| Every red verdict names a file, a cause and a next step | asserted for every case in the configuration error table, and at the verdict in `TestAMisplacedScopeIsRedAndNamesFileCauseAndNextStep` |
| Nothing else moves | `./xeno gate verify` over the whole trail |

The one criterion with no test is the one about two meanings of `rules_hash` divided by this
commit, which is a statement about the trail rather than about the code. It was checked by
reading the artifacts of this intent, where the division turns out to fall inside the intent
itself.

<!-- xeno:section:results -->
## Results

**The suite is green.** 203 test cases pass across the three packages this piece touched, 17
packages report `ok` or no test files, and nothing fails. `gofmt -l` outside `vendor/` lists
nothing and `go vet ./...` is silent.

**`./xeno gate verify` is at exit 0 over the whole trail**, 175 verdicts, which is 171 before
this intent and four of its own phases. No sealed artifact's `artifacts_hash` moved, which is
what the criterion asked and what the placeholder decision in A66 was taken to protect.

**G-Rules reports a verdict in this repository.** Over the tree as it stands, with no rule
files anywhere, it passes, which is the state the criterion put first. The gate table's row no
longer reads `not-implemented`, so a phase here now reports ten gates running rather than nine.

**The collision was exercised against the real binary on a throwaway copy of the tree.** Two
binding rules with one id, at `given/provider/` and `given/org/`, turned 03-implementation red
with one finding, `F-e02b63`, whose cause names the id and both files and whose next step is
that a gate may not resolve a binding collision. Removing one of the two returned the phase to
green, and the copy was then discarded, so no sealed phase of this repository was judged against
an invented rule tree.

**The hash a phase recorded is reproducible by hand.** On that copy, with one rule in
`given/provider/`, the phase recorded
`c721ad48d50f6c298d8cd88bc34b1dc388d3d7feb38e9cc1f14cb72a00c2ccbc`. Piping the rendering A66
defines through `sha256sum`, as nine tab separated fields ending in an empty check field,
produces the same value. That is the claim Appendix B's byte exactness actually makes, checked
against a pipeline rather than against the implementation.

**The division of the field's two meanings falls inside this intent.** 00-intake, 01-requirements
and 02-design carry `rules_hash: by-hand`, because the harness wrote them before the writer
existed; 03-implementation and 04-verification carry
`e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`, which is the hash of the
empty rendering and the correct statement for a repository that resolved its tree and found no
rules. The boundary 00-intake predicted turns out to run through the middle of the intent that
predicted it.

**Each phase of this intent was judged green on its own artifacts**, 00 to 03, with G-Supply,
G-Secret, G-Test, G-Policy and now not G-Rules among the `not-implemented` rows: four remain
where there were five.

<!-- xeno:section:gaps -->
## Gaps

**Nothing checks `rules_hash` after it is written.** G-Schema recomputes `context_hash` and
`strings_hash` and deliberately does not recompute this one, for the reason A66 gives: the
field names the rule set as it was, not a file in the tree now. So a wrong value in the field
is caught by nobody, exactly as for `secrets_hash`. What the field gives a reader is
reconstructability where they hold the tree of that commit, not a check.

**The placeholder is still accepted in a new artifact.** `rules_hash` stays in
`WriterlessHash`, so an artifact written by hand today can say `by-hand` and pass, although a
writer now exists. Tightening it would turn seventy-one sealed intents red, which A62 measured
on the same field; the honest statement is that the honesty rule is one release behind the
writer and that this is the second field in that position.

**No predicate is evaluated, so a `check` is unvalidated.** A rule may name a type that does
not exist, or parameterise one wrongly, and resolve and hash without complaint. The type
registry belongs to the next piece, and until it lands a `checked` rule is a rule whose check
is read and never run — which is also why `kind: checked` is not yet more useful than
`review`.

**`applies_to` is not checked against the phases.** A typo silently disables a rule for the
phase it was meant for. A67 records why it is not a finding here and that it is the proposal
worth making, and until that proposal lands this is a real hole in a file a person writes by
hand.

**No rule has ever been written by anyone but this test suite.** Every tree exercised here was
built by a test or by me on a copy. The shipped set is the first rule tree a person will write,
and the fixtures in this piece are what it will be judged against — which means the first real
use of this code is also the first chance for its findings to be read by somebody who did not
write them.

**The resolution has no test for a four level tree with mixed origins.** Precedence is tested
pairwise and by level, and `more` is simple enough that the pairwise cases cover its branches,
but there is no fixture with all eight combinations present at once. A cheap test to add and
not added here.

**The fixtures are not where WP17 wants them.** They build trees in `t.TempDir()` rather than
living under `corpus/`, so the per-gate corpus that package specifies will have to absorb them.
