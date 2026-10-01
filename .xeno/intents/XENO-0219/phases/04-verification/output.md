---
intent: github.com/triplem/xeno#162
phase: 04-verification
created: "2026-10-01T16:03:38Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+94d4c6c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bf253623f8acd73c74157a4f5ace3fe9401bcd1c94b6396acf2be06429d22a66
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

Every criterion is a test, and five of them were also run against the real binary over this
repository's own history, which the results record.

| Criterion | Test |
|---|---|
| The implication holds and nothing else | `TestSectionImpliesSection`, four states including a whitespace-only antecedent |
| A section the template does not have is a configuration error | `TestASectionTheTemplateDoesNotHaveIsAConfigurationError` |
| A check naming no section is a configuration error | `TestACheckWithNoSectionsIsAConfigurationError` |
| Every subject judged against a named pattern | `TestCommitMessageJudgesEverySubject` |
| The second pattern wants a reference | `TestTheSecondShippedPatternWantsAReference`, `TestTheWithIssuePattern` over ten subjects |
| An unknown pattern says what is shipped | `TestAnUnknownPatternNameSaysWhatIsShipped` |
| Merge commits exempt where the rule says so, judged where it does not | `TestMergeCommitsAreExemptWhereTheRuleSaysSo` |
| A trailer on every commit, by key, case insensitively | `TestCommitTrailerRequiresItOnEveryCommit`, `TestTrailersAreReadByKeyAndCaseInsensitively` |
| A signature that verifies on every commit | `TestCommitSignatureRequiresAVerifiedSignature`, `TestAnUnsignedCommitIsNotSigned` over every letter git gives |
| An approval by an author of the range is red | `TestApproverNotAuthor`, by email and by name |
| An approval by somebody else, and no decisions, are green | the same test, and `TestApproverNotAuthorWithNoDecisionsPasses` |
| A rule that needs a range and did not get one is red | `TestARuleThatNeedsARangeAndDidNotGetOneIsRed`, once per type |
| A git failure is a finding | `TestARefThatDoesNotResolveSaysSo`, `TestADirectoryThatIsNoRepositoryIsAnError` |
| One `git log` per gate run | `TestTheRangeIsReadOncePerGateRun`, five rules and five findings |
| The registry answers for all five and for no sixth | `TestTheFiveTypesOfSectionNineAreRegistered`, `TestACheckedRuleWithNoImplementationIsRed` |
| Both patterns reachable the same two ways | `TestBothPatternsAreNamedInTheListing`, and `CheckMessage` is the one implementation both call |
| Nothing else moves | the whole suite, `gofmt`, `go vet`, `./xeno gate verify` |

Two criteria have no test of their own. That the range is read oldest first is asserted in
`TestTheRangeIsReadOldestFirst` rather than as a criterion, because nothing in the specification
requires an order and the predicates are order-independent; it is pinned so a change to the
format string shows. And that a subject survives the format string is
`TestASubjectWithAwkwardCharactersSurvives`, which is a property of the reader rather than a
criterion of the gate — and the one that would have caught the NUL problem if it had been
written first.

<!-- xeno:section:results -->
## Results

**The suite is green.** 127 cases pass across `internal/git` and `internal/gates`, 18 packages
report `ok` or no test files, nothing fails, `gofmt -l` outside `vendor/` lists nothing,
`go vet ./...` is silent. `./xeno gate verify` is at exit 0 over 193 verdicts.

**The four commit types were run against this repository's own history.** On a copy, with the
range `main..HEAD`, which is the one commit this branch carries:

*`commit-message` with `conventional-commits`.* Green. The subject of
`feat(rules): G-Policy counts the answer to a review rule` matches the pattern, which is the
first time a rule has judged this project's own commit convention and found it conforming.

*`commit-trailer` with `Xeno-Intent`.* Red, `F-6beedb`, naming the commit by its short hash and
subject and saying it carries no `Xeno-Intent` trailer. That is correct: this project puts the
reference in the footer as `Refs` and `Closes`, not as that trailer, and section 9 ships the
trailer as an example precisely because no project should inherit it unasked.

*`commit-signature`.* Red, `F-3759f3`, same commit, no signature that verifies. Also correct —
nothing here is signed, and A21 records that signing waits for publication.

*The same tree with no range.* Both commit rules red, each naming itself and reporting
`no commit range: base "", head ""`, with the next step saying the range is an input of the run
and never inferred. Two rules, two findings, one for each rule that needed it.

**`section-implies-section` was not run on this repository**, because no rule naming it exists
here and writing one would have meant inventing a rule to test a predicate. It is covered by
four fixture states plus the two configuration errors.

**The findings name a commit the way a person would.** Eight characters of hash and the subject
in quotes, so a reader recognises the commit without resolving it, and the file named is the
rule that asked rather than the commit, because the rule is what a reader can change.

**One `git log` per gate run, measured by the finding count.** Five rules naming
`commit-message` over a range with one bad subject produce five findings and one subprocess; the
error case produces one finding per rule and one attempt.

**Three quiet rows, unchanged.** G-Supply, G-Secret and G-Test remain `not-implemented`. This
piece adds no gate and removes none.

<!-- xeno:section:gaps -->
## Gaps

**No signed commit was ever tested against.** Every signature case is either an unsigned commit
from a real repository or a `Commit` value with a letter set by hand. The path where git reports
`G` has never run, because signing a commit in a test needs a key, and A21 says this project does
not sign yet. So A70's decision is tested at the boundary and not through it.

**The tests need git on the path and skip without it.** That is the cost of refusing a second
dependency, recorded in the design. A machine without git reports a pass over skipped tests,
which is the weakest possible signal, and the CI image has git because it checks out the
repository — so the skip protects a developer's machine and hides nothing in the pipeline.

**`section-implies-section` cannot be judged where there is no template.** The check against the
phase's template is skipped when the template will not load, so a repository with no vendored
plugin resolves the rule and reports nothing about its section names. That is deliberate — a
rule's sections cannot be checked against a template that is not there — and it means the
configuration-error finding is unreachable in exactly the repositories least likely to have got
the names right.

**A trailer's value is never read.** `Xeno-Intent: whatever-somebody-typed` satisfies the rule.
Section 9 asks for a required trailer and nothing more, and the alternatives record why checking
the value needs an expression; what it means in practice is that the trailer proves a habit
rather than a reference.

**`approver-not-author` compares strings and says so.** An approver who writes their name
differently from their git author line passes, and one who writes it the same way is caught.
Section 9 states the limit; this gap exists to make sure a reader of a green verdict from that
rule knows what they are holding.

**The range is still not recorded anywhere.** A verdict from any of the four commit types depends
on `--base` and `--head`, which the artifact deliberately does not carry, so re-running
`gate verify` later cannot reproduce it and does not try — `verify` recomputes hashes and
compares, and nothing in it re-evaluates a predicate. What that means is that a red commit rule
is evidence from the run that produced it, not from the record.

**Nothing bounds a range.** A rule named over a range of ten thousand commits reads all of them
into memory in one `git log`. Nobody has measured it, no limit exists, and the honest statement
is that the range this project passes is always a branch.

**The two git test helpers are near-duplicates in different packages.** Twenty lines each,
deliberately not shared, recorded here because it is the kind of thing a reviewer flags and a
maintainer later consolidates with a third package nobody wanted.
