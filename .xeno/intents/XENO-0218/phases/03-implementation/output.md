---
intent: github.com/triplem/xeno#160
phase: 03-implementation
created: "2026-10-01T15:39:13Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bca3ab449142232b93033c1ae233ea320779713ea79a9fc6317b8d0c6edfbc45
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Two files changed in `internal/gates` and `internal/model`, one line in the runner, one new
test file of 196 lines, two register rows. 164 lines added across the existing files.

**`internal/model`.** `ChecklistEntry` with `Rule`, `Result`, `Note` and `Source`, and
`Output.ReviewChecklist` as `review_checklist,omitempty` beside the three lists already there.
`ChecklistResults` holds the three results section 9 fixes and `ChecklistNeedsNote` the two that
owe a note, so the asymmetry — a `met` needs none — is a declaration rather than a condition
written twice. The comment on `Rule` says why its absence is what the gate keys on.

**`internal/runner`.** One line: `review_checklist` added to `frontmatterOrder`, so the key is
written where section 5's order puts it rather than sorted alphabetically by the map.

**`internal/gates`, the registry.** `predicateFn` and an empty `predicates` map. The comment
says what it is for and that filling it is the next piece's only change here. A checked rule
whose type is not a key produces a finding; a type that is a key is evaluated by the function
behind it, which is the line the next piece will exercise.

**`internal/gates`, `policy`.** Replaces `notImplemented` in the table. Resolves the tree
through `rules.Load` and `rules.Effective` and adds no resolution of its own. Discards the
problems both return, because an unresolvable tree is G-Rules's finding and reporting it twice
in one verdict would say one broken file is two. Calls `checkedRules` at every phase and
`reviewChecklist` only at the last, then sorts the findings by cause so a verdict reads the same
on two runs.

**`checkedRules`.** For every checked rule in the effective set whose `applies_to` names this
phase: the type is looked up, an absent implementation is a finding naming the rule and the type
with the next step of writing it as `review` until the type ships, and a present one is called.

**`reviewChecklist`.** Reads the P5 artifact's frontmatter and returns nothing where it cannot,
because the artifact's absence is G-Schema's finding. Builds the set of review rule ids from the
effective set, then walks the entries: no result, a result outside the three, or a `deviation` or
`not-applicable` with no note, each a finding naming the entry. An entry with no `rule` is
counted towards nothing and the walk moves on, which is section 12's lens clause. An entry whose
rule is not a review rule of the set is a finding. Then every review rule with no answer is a
finding, which is the direction the whole gate turns on.

**`entryName`** names an entry that may have no rule id: by its rule where it has one, by its
source where it has only that, and as "with no rule" where it has neither. A finding that cannot
name what it is about is a finding nobody can act on.

**Tests, `internal/gates/policy_test.go`.** An empty tree green at all six phases. A review rule
with no entry red. An answered rule green. A table over what an entry owes: no result, a result
outside the three, a `deviation` and a `not-applicable` without a note, and the two that pass —
a `deviation` with a note and a `met` without one. Four lens entries leaving a missing answer
missing. A lens entry owing a result and being named by its source. An entry for a rule nobody
has. A checked rule with no implementation red, and the same rule at a phase it does not name
green. The table no longer reporting `not-implemented`. And a tree G-Rules rejects passing here
rather than being reported twice.

**`ASSUMPTIONS.md`, A68 and A69.** A68 is the frontmatter key, the entry's fields, why the gate
keys on `rule` rather than `source`, and why the rendered section is not derived. A69 is the
reading that every review rule is answered at P5 whatever its `applies_to`, with the alternative
named, plus the finding for an entry outside the set.

<!-- xeno:section:deviations -->
## Deviations from the design

**Discarding the resolver's problems is a decision the design did not take.** 02-design says the
gate calls `Load` and `Effective` and says nothing about what it does with what they return.
Writing it made the question unavoidable: a misplaced `scope` would otherwise be reported by
G-Rules and again by G-Policy, as two findings about one file in one verdict. They are discarded
here, with the reason in the code and a test that pins it. It is the right behaviour and it was
decided in the wrong phase.

**Two helpers collided with names already in the package.** `reviewRule` and `causes` both
existed in the gates tests — the first from #158's own test file, the second in
`schema_test.go`, where it also prefixes the file name. The new file now uses `aReviewRule` and
the existing `causes`, which is better than a second spelling of the same helper, and it is
recorded because the first version of the test file was written without reading the package's
existing test helpers.

**`TestNoRuleTreeIsGreenForPolicy` was written twice.** Its first version built a fixture it
threw away and assigned a result onto an unrelated value to reuse a variable, which passed and
asserted less than it appeared to. Rewritten to walk all six phases against an empty tree, which
is what the criterion says. Recorded because a test that passes while asserting the wrong thing
is the one kind of test that makes a suite worse.

**The gate is correct at P0 to P4 for a reason the criterion did not state.** The criterion says
the gate is quiet where it has nothing to say. What makes it quiet is that `reviewChecklist` is
called only at the last phase, so the four earlier phases cannot report a checklist finding even
if an artifact there carried entries. That is deliberate — the checklist is a P5 artifact — but
it means entries written into an earlier phase's frontmatter are silently ignored rather than
reported, which nothing in the criteria covers and which is in the gaps.

**No `xeno` command writes an entry, so every test writes YAML by hand.** That is the state of
all four structured lists and it is in the non-goals, but it has a consequence worth naming: the
fixtures assert what the gate reads, and nothing asserts that what an agent would plausibly
write is what the gate reads. The first checklist written by something other than this test file
will be in the shipped set piece.
