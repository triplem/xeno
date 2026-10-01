---
intent: github.com/triplem/xeno#160
phase: 04-verification
created: "2026-10-01T15:42:13Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e9fb6d2693ca8c24c8ad60edc56f51de5fa6f5164f40ed713f333e0b0b9f21ed
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

Each criterion is a fixture in `internal/gates/policy_test.go`, and five of them were also run
against the real binary on a throwaway copy of this repository, which the results record.

| Criterion | Test |
|---|---|
| A repository with no rule tree is green at every phase | `TestNoRuleTreeIsGreenForPolicy`, which walks all six |
| A review rule that applies to P5 with no entry is red | `TestAReviewRuleWithNoEntryIsRed` |
| An answered review rule is green | `TestAnAnsweredReviewRulePasses` |
| An entry with no result is red | `TestWhatAnEntryOwes/no_result` |
| A result outside the three is red | `TestWhatAnEntryOwes/a_result_outside_the_three` |
| `deviation` and `not-applicable` owe a note, `met` does not | the same table's four remaining cases |
| A lens entry neither completes nor breaks the count | `TestLensEntriesDoNotAnswerARule`, with four of them against one unanswered rule |
| A lens entry still owes a result, and is named by its source | `TestALensEntryStillOwesAResult` |
| An entry for a rule outside the effective set is red | `TestAnEntryForARuleOutsideTheSetIsRed` |
| A checked rule with no implementation is red, naming rule and type | `TestACheckedRuleWithNoImplementationIsRed` |
| Only at the phases the rule names | `TestACheckedRuleIsReadAtThePhasesItNames` |
| `G-Policy` is no longer `not-implemented` | `TestGPolicyIsNoLongerReportedAsNotImplemented` |
| The tree is resolved once, and a broken tree is not reported twice | `TestAnUnresolvableTreeIsNotReportedTwice` |
| Nothing else moves | the whole suite, `gofmt`, `go vet`, `./xeno gate verify` |

The criterion with no test of its own is that the gate reads the resolver rather than resolving
again. Nothing asserts the absence of a second implementation; what is asserted is the
consequence that made it matter, which is the broken tree reported once. The direct form of it
would be a test of the package's import list, and A42 is the precedent for putting that in CI
rather than in a test.

<!-- xeno:section:results -->
## Results

**The suite is green.** 94 cases pass in `internal/gates` alone, 17 packages report `ok` or no
test files, nothing fails, `gofmt -l` outside `vendor/` lists nothing, `go vet ./...` is silent.
`./xeno gate verify` is at exit 0 over 187 verdicts, which is 183 on `main` plus the four phases
of this intent judged so far.

**The gate was exercised against the real binary on a copy of this repository**, in five steps,
using XENO-0217's own sealed phases as the artifacts:

*A review rule with no answer.* `release-notes-are-filled` under `given/org/`, applying to
`05-review`, turned that phase red with `F-cca1ab`: "review rule release-notes-are-filled has no
checklist entry", next step "answer it in review_checklist with met, deviation or
not-applicable".

*The same rule answered.* A `review_checklist` entry with `result: deviation` and a note turned
the phase green. That is the mechanism section 9 describes working end to end: a rule a gate
cannot evaluate, answered by a person, recorded in the artifact, counted by the gate.

*The note removed.* The same entry without its note turned the phase red again with `F-dc8382`:
`checklist entry "release-notes-are-filled" is deviation and carries no note`. So the asymmetry
holds in the binary and not only in the table test.

*A checked rule at a phase it names.* The section 9 example rule, `section-implies-section`
under `given/org/` applying to `02-design`, turned that phase red with `F-05937e` naming the
rule and the type, while G-Rules stayed green on the same tree — the tree is well formed and the
predicate is unimplemented, which are two different statements and came out as two different
gates.

*What the copy also showed.* 02-design of XENO-0217 carries `rules_hash: by-hand`, because it
was written before #158's writer existed. So a phase whose rule set is judged by G-Policy today
can carry a `rules_hash` that records nothing about which set applied. That is A66's division
seen from the other side, and it is in the gaps.

**Findings are stable and sorted.** `policy` sorts by cause, so two runs over one tree produce
the same verdict in the same order, which is what makes a finding id stable across runs.

**Three quiet rows where there were four.** A phase now reports G-Supply, G-Secret and G-Test as
`not-implemented` and eleven gates running of fourteen.

<!-- xeno:section:gaps -->
## Gaps

**Entries written into a phase before P5 are ignored without a word.** `reviewChecklist` runs
only at the last phase, so a `review_checklist` in a design artifact's frontmatter is read by
nobody and reported by nobody. It is the right place for the check and the wrong silence: the
honest behaviour would be a finding that a checklist appears where no checklist belongs, and
that is not in section 9's enumeration, so it is a candidate for the same proposal A67 already
names for `applies_to`.

**A phase judged against a rule set can carry a `rules_hash` that proves nothing.** On the copy,
02-design of XENO-0217 was judged red against a checked rule while carrying
`rules_hash: by-hand`, because that artifact predates #158's writer. Every artifact sealed before
this week is in that position, so for the whole existing trail the two gates that read rules can
disagree with the field that is supposed to record which rules applied. Nothing is wrong in the
code; what is missing is any statement in the artifacts themselves about which of the two eras
they belong to.

**Nothing checks that an answer is true, by design, and nothing says so where a reader will
see it.** Section 9 is explicit that this is the limit of a deterministic gate, and section 16
lists what a green gate does not mean. A `met` on every entry with no note is a green verdict and
an unexamined review, and that sentence belongs in section 16's list rather than only in this
artifact — which is a spec change and therefore a person's.

**A lens that forges a rule id is indistinguishable from an agent that answers one.** A68 records
it as the honest limit: both are self-asserted strings in the same file. It becomes real when
WP11 ships lenses, and the thing that would make it checkable — who wrote which entry — is not
something the artifact carries.

**No checklist has been written by anything but a test and me.** The fixtures assert what the
gate reads; nothing asserts that what an agent would write is what the gate reads. The shipped
set piece is where that gets its first outside writer, and the first finding worded badly will
be found there.

**The predicate registry has one key path and no test behind it.** A type that is in the map is
evaluated by the function behind it, and the map is empty, so the branch that calls a predicate
is written and unexercised. The next piece covers it by existing; until then this is code with
no test.

**The resolution is run twice per gate run.** G-Rules resolves the tree and G-Policy resolves it
again, in the same `gate run`, because each gate is a function of the context and nothing caches.
Section 7 says the two gates should resolve the tree once; what this piece honours is that they
resolve it the same way, not that they do it once. On this repository that is two walks of an
empty directory, so it is a correctness argument about a future cost and not a measured one.
