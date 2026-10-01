---
intent: github.com/triplem/xeno#165
phase: 04-verification
created: "2026-10-01T16:35:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4f94129.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1386cc1d282af791746414f39b25f1ca2b62015ef53919cbbb5b9ca3aea29cfa
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Most criteria here are about files a person reads, so the mapping has two columns of a different
kind: what a test asserts, and what was read back by hand.

| Criterion | Checked by |
|---|---|
| Four rules, each passing the gate that reads them | `TestTheShippedSetResolves`, which counts files against rules in force |
| None names a supplier, a customer or a project | `TestTheShippedSetNamesNothingSpecific`, twelve words, and a reading of all four statements |
| Each is well formed for a reader | `TestEveryShippedRuleIsWellFormedForAReader`: level, scope, not binding, eight words, a full stop, a phase |
| A checked rule names an implemented type | `TestAShippedCheckedRuleNamesAnImplementedType`, and that at least one is checked |
| `release-notes-are-filled` is evaluated and red over an empty section | `TestTheShippedReleaseNotesRuleIsEvaluated`, reading the file from where it ships |
| `section-non-empty` is in the registry and nowhere else | `TestTheRegisteredTypesAreTheOnesTheDocumentsAllow`, with the count asserted |
| Its four states and two configuration errors | `TestSectionNonEmpty`, `TestSectionNonEmptyRefusesAnUnknownSectionAndNoSection` |
| `xeno init` carries the rule tree, and what arrives resolves | `TestVendorPutsTheShippedRuleSetInTheRepository` |
| `init` twice changes nothing | `TestInitIsIdempotent`, which was already there and still passes |
| The examples resolve where a reader puts them | `TestTheExamplesResolveWhereAReaderPutsThem`, all four copied into `given/org/` |
| An example at the wrong level says so | `TestAnExampleAtTheWrongLevelSaysSo`, one finding naming claim and path |
| Nothing reads `examples/` and no `enabled` field exists | `grep -r enabled` over the tree, and the resolver's own fields |
| The hooks carry their limits in their own text | read back; `sh -n` on both |
| The documentation call is answered in the register | A72 |
| This repository answers its own review rules | this intent's 05-review, which is the next phase |
| Every finding read as a stranger would | read back, recorded in the results |
| Nothing else moves | the suite, `gofmt`, `go vet`, `./xeno gate verify` |

The criterion with no test and no reading is the last structural one: that `rules_hash` now carries
something. It is observed rather than asserted — this intent's own 03-implementation records
`ce2250ab`, where every artifact before the rules existed records the hash of an empty rendering.

<!-- xeno:section:results -->
## Results

**The suite is green.** 382 cases pass across the tree, nothing fails, `gofmt -l` outside
`vendor/` lists nothing, `go vet ./...` is silent. `./xeno gate verify` is at exit 0 over 199
verdicts.

**The divergence this piece created is gone.** Before the guard, writing the four rules turned
twenty-four sealed P5 phases red: `committed status green, recomputed red`, every six-phase intent
in the repository. After A74's rule, `gate verify` is clean again, and the two artifacts of this
intent written after the set exists carry its hash where everything older carries the hash of an
empty rendering or the placeholder.

**The set resolves and is in force.** Four files, four rules, no finding from G-Rules, and this
intent's 03-implementation records `ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721`
— the first `rules_hash` in this repository that is a hash over rules rather than over nothing.

**`release-notes-are-filled` was evaluated against the file as shipped.** The test copies the rule
from `.xeno/plugin/rules/given/builtin/` rather than restating it, builds a P5 artifact with the
review template, and gets green over filled notes and red over empty ones, naming
`release-notes is empty`. One checked rule, proved against the real file.

**`xeno init --vendor` carries the tree.** Into an empty repository: four rule files under
`.xeno/plugin/rules/given/builtin/`, resolving with no problem and no collision, with as many
rules in force as files vendored. Without `--vendor`, nothing is copied. Run twice, the second run
changes nothing, which is WP9's own criterion and still holds.

**The examples resolve where a reader puts them.** All four copied into `given/org/` load with no
problem and every one claims `scope: org`. One copied into `given/project/` produces exactly one
finding, naming both the claim and the path: that is the gate teaching the two axes, and it is the
only teaching material this repository has until WP16.

**Nothing reads `examples/` and no `enabled` field exists anywhere.** A grep over the rule tree,
the examples and the resolver finds none, which is section 9's "a rule that is not in the tree does
not apply" holding by construction rather than by a comment.

**Every finding these rules can produce was read back as a stranger would read it**, which is the
criterion two earlier reviews deferred here. Three of the four rules produce no finding of their
own — a review rule's output is its statement in a checklist — so what was read is the statements,
and two were rewritten: `interface-change-needs-a-migration-note` gained "something outside this
intent depends on", because the first version would have fired on every internal rename; and
`new-dependency-needs-a-rationale` lost the word "justified", because a rule asking whether a
dependency is justified asks for a verdict rather than for the reasoning. The checked rule's red
message — `release-notes is empty`, next step "write release-notes, or record a deviation against
the rule" — names the section and the two ways out, and the configuration error names the
template and lists its sections. The hook templates were read for whether somebody could install
them from the text alone, which is why both carry the `core.hooksPath` lines.

**The two statements that did not change** are `deviations-are-traceable`, whose "names what it
departs from" is the whole point and reads as an instruction, and `release-notes-are-filled`,
which says what the section is for and then lets the predicate do one thing.

<!-- xeno:section:gaps -->
## Gaps

**A rule added after an artifact was written does not apply to it until something renders it
again.** A74's guard reads the hash the artifact recorded, and the hash is written at render time,
so a project could write its P5 and then add a rule and have the rule not apply to that phase. The
next section write re-records the hash and both values are visible in the artifact, and neither of
those is a check. Closing it means deciding whether a rule set change makes a phase stale, which is
G-Freshness's vocabulary and a specification question.

**Three review rules, one checked rule, because the templates have no sections for the other
two.** `interface-change-needs-a-migration-note` is written out in section 9 as a checked rule over
an `interfaces` section implying a `migration-notes` section, and the shipped template set has
neither. So the specification's one worked rule example cannot be shipped as written. That is a
finding about WP3's template set and not about the predicates, and it is the strongest argument
for adding those two sections — which this piece deliberately did not do.

**The three review rules have never been answered by anybody.** This intent's own P5 is the first
answer and I am writing it, which is the same closed loop two earlier reviews named. The difference
is that this time the statements had to be answerable rather than plausible, and two were rewritten
for it. The loop closes properly the first time somebody else answers them.

**The forbidden-word test is a tripwire, not a check.** Twelve words. A statement that is specific
without naming a language, a tool or a project passes it, and a legitimate statement that needed
one of those words would fail it. The property it stands in for — nothing specific reaches every
project everywhere — has no mechanical form.

**The hooks are read and not run.** Both parse under `sh -n`; neither has been executed.
`pre-receive` reads a ref update protocol nothing here produces, and `prepare-commit-msg` would
need a commit in flight. They are examples, and an example that was never run may not run.

**`xeno init` vendors the rule tree and nothing verifies what it vendored afterwards.** G-Supply
is still `not-implemented`, so a project that edits the shipped rules in place gets no finding at
all. The comment in each file says the set is not edited here; nothing enforces it, and the
sentence in section 9 — "cannot be edited in place" — is true of intent rather than of mechanism.

**The set is minimal to the point of being thin.** Four rules, three of which ask a person to
confirm something they probably did anyway. Nothing in them would have caught any finding this
project has recorded in twenty intents. That is what "deliberately minimal and generic" buys, and
whether it buys anything is a question the first adopter answers, not this test suite.
