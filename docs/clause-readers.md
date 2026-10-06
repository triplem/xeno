<!-- SPDX-License-Identifier: Apache-2.0 -->

# The clauses and their readers

One pass over the process definition and the implementation plan, listing what each
normative clause requires and what reads it. The output is this list. Fixing what it
finds is not part of it, for the reason #202 gives: eight findings of this shape cost
about an intent each, and a pass that fixed as it went would stop at the first.

It is a measurement, not a specification. Where it and the documents disagree, the
documents win. The pass was made on 2026-10-03 against `d19a1ca`, and a reader named by
symbol is wrong the day the symbol is renamed with nothing to say so, which is the most
a document can do about its own ageing.

## How it was made

219 candidate sentences were extracted from both documents by searching for `never`,
`always`, `must`, `cannot`, `has to`, `is required`, `refuses`, `fails red` and `is
recorded`, outside tables and code blocks. Each was read and placed in one of four
kinds, and for the first kind the reader was looked up in the code rather than recalled.

**The four kinds.** A clause is only interesting here if something could read it.

- **Tool requirement** — the runner, a gate or CI can satisfy or violate it. These are
  listed below with their reader.
- **Architectural property** — true by how the code is arranged, with nothing asserting
  it. A change could break it silently. Listed, because this is where the gaps are.
- **Addressed to a person** — "the v1 numbers must not enter certification records
  without re-measurement", most of section 10, the limitations. Not a gap. Counted, not
  enumerated: #202 warns that calling these unread produces a list nobody trusts.
- **Explanation** — prose using a normative word about something other than a
  requirement: "a reader cannot tell", "it was never implemented". The largest group.

| kind | count |
|---|---|
| tool requirement | 39 |
| architectural property | 11 |
| addressed to a person | 48 |
| explanation | 126 |

## Tool requirements, and what reads each

A reader is named as the thing that would fail if the clause were violated. "test" means
only a test asserts it; a gate or a refusal is stronger, because it acts on a repository
rather than on a fixture.

One row arrived after the pass. The review-rule coverage clause was missed, and this
intent's own review phase went red on it, which is why the implementation phase's digest
says thirty-four and this table says thirty-five. It is left visible because a pass that
quietly absorbed its own miss would be the thing #202 warns about.

Two more arrived later still, in #212, and they are one sentence of section 7 split in
two because its halves have different readers. G-Test's row asks for a declared test
result and for the completeness of the mapping of acceptance criteria; the first is read
from #212 and the second by nothing, so a green G-Test now means half of its row.
Splitting one clause over two rows is why the count is thirty-seven and not thirty-six,
and it is the only row here reporting a gate that judges part of what it is asked for.
The alternative was to leave G-Test at `not-implemented`, which was honest and caught
nothing; #212 weighed the two and this document is where the cost was put.

Two of the rows below are a correction rather than a late addition, and the
difference matters to what this pass claims. #229 gave two more of section 8's
requirements a reader and did not touch the row describing them, so until #254 the
table said a question's shape was read by G-Questions alone. It is read by G-Schema,
which calls the shape check through `phaseResult` from P0 — the fact #229 itself
turned on, since a check added there reaches every artifact ever written. One row
describing a third of the clause and naming the wrong gate became two describing all
of it and naming what reads each, which is why the count is thirty-eight. The pass was
not incomplete here; it was overtaken, and then wrong.

**Why the mapping half has no reader**, and what a project can do about it. A rule
answered by a person is available as `examples/rules/mapping-is-complete.yaml`, which is
not enabled here or anywhere: adopting it gives the clause a reader once per intent rather
than once per criterion, which is weaker than the gate section 7 asks for and is the
ceiling until a criterion is identifiable. The reader column above says `nothing` because
an example nobody has enabled fails nothing.

It needs an acceptance criterion to be identifiable, so that a gate can say the mapping
covers it, and the convention that would make one identifiable is an addition to
**section 5** — the Language and Rendering subsections, where `template.yaml` holds the
section ids, their order and the required fields. It is therefore a specification
change. This paragraph said section 9 until #258, and section 9 is the Rule model.
Searching the specification for "acceptance" returns 572 and 692 in section 5, 710 and
713 in section 6's phase table and 942 in section 7's G-Test row, and nothing in all of
section 9. A wrong number here is wrong at the moment the document is read, which is
before somebody writes the commit, and #258 and its first comment had both copied it.

**The check needs two things identifiable and the measurement had counted one.** A
completeness check reads a P4 `test-mapping` and asks whether it covers every criterion
P1 raised, so it needs criteria that can be named and a mapping that names them.
Measured on 2026-10-06 over the trail as it stood before the intent that measured it,
which is what the counts of 64 are:

| | artifacts | identifiable | not |
|---|---|---|---|
| P1 `acceptance-criteria`, as a numbered list | 64 | 19 | 45 |
| P4 `test-mapping`, citing a criterion by number | 64 | 7 | 57 |

The cost of applying the check to the trail is **57 verdicts**, because a completeness
check is judged at P4. This intent's own artifacts are not in those counts and move both
the right way — its P1 numbers its criteria and its P4 mapping cites the numbers, so the
P1 figure becomes 20 of 65 the moment it lands. That is the whole of what a forward-only
anchor does, visible one intent at a time. The figure of 2026-10-05 — 55 P1 artifacts,
46 with no numbered criteria, and one of the nine that had them already failing a check
by number — stands as a dated measurement rather than being corrected: the trail grew by
nine intents between the two dates, so part of the difference is growth and part is
which population was counted, and a reader shown only the newer figure cannot tell
which.

**Decided in #258, and what it waits on.** The anchor is the template version: section 5
requires numbered criteria from `requirements@1.1.0` and a numbered mapping from
`verification@1.1.0`, and the check judges only artifacts declaring 1.1.0 or later.
Nothing is re-judged and no anchor is invented — all 64 P1 artifacts declare
`requirements@1.0.0` today and `template` is already a frontmatter field, read by
`model.Template`, so the equivalent of G-Policy's `judgedUnder` that a gate's own checks
were said to lack is already in every artifact. It waits on the section 5 commit, which
the first standing rule makes a person's. The reader column still says `nothing`,
because a decision is not a reader and nothing fails today if a mapping covers three of
five criteria; A90 is why that is left as it is.

**Why the recommendation's reason has no field**, and what was decided about it. Section
8 asks for "the agent's recommendation with a reason" and `model.Option` carries `Text`,
`Consequence`, `Recommended` and `Free`, so the reason can only live inside the
question's text or inside a consequence, where it is indistinguishable from what it
shares the field with. #247 argued both homes without settling: "with a reason" reads as
one reason per question, which puts it beside `Text` on `Question`; a reason is about
the option chosen, which puts it on `Option` and costs a field empty on every option but
one.

#258 settled it on `Option`, required by `QuestionAsked` wherever `recommended: true`.
The reason and the recommendation then cannot drift apart — a recommendation that moves
to another option takes its reason with it or the writer refuses, where a reason held
beside the question would go on describing the option it used to be about with nothing
able to notice. A conditionally required field is not new here: `ChecklistEntry.Note` is
required for two of three results and optional for the third. This waits on the section
8 commit, and the row goes on saying there is no field until one exists.

| § | clause | reader |
|---|---|---|
| 2 | every assumption is recorded and confirmed | G-Assumptions |
| 2 | a `binding` rule under `given/` cannot be overridden | `internal/rules` precedence, G-Rules |
| 4 | an intent dropped in P1 meets G-Complete at close | `IntentClose`, G-Complete's second mode |
| 4 | nothing undeclared in `evidence/` | G-Schema |
| 4 | `result` is required on `test-report` and `build-log` | `model.ResultRequiredKinds`, G-Build |
| 5 | every declared evidence item resolves and its hash matches | G-Evidence |
| 5 | required fields present in every artifact | G-Schema, `testdata/required-fields.yaml` |
| 5 | `schema_version` is read, never enforced backwards | G-Schema, by omission |
| 5 | a hash field carries 64 hex characters or the placeholder | G-Schema |
| 5 | finding ids do not depend on the run | `hashing.FindingID`, test |
| 5 | `artifacts_hash` excludes `gate.yaml` and `cost.yaml` | `hashing.PhaseExcluded` |
| 5 | the budget is judged against the size the lock recorded | G-Schema's budget check |
| 5 | a missing strings bundle fails red | G-Schema |
| 5 | gates match section ids, never headings | `template.Parse`, test |
| 6 | a question carries two to four options, each with its consequence, and one free entry | G-Schema's shape check, G-Questions |
| 8 | a question recommends exactly one of its options | `QuestionAsked`, the writer only; no gate reads it |
| 8 | the recommendation carries a reason | **no field to carry it**; nothing reads it (#247) |
| 6 | a decision carries id, chosen, rationale and `decided_by` | G-Questions' decision shape |
| 6 | `phase start` refuses a predecessor with no completed verdict | `predecessorAllowsStart`, test |
| 6 | `phase start` refuses a second start of a running phase | the run marker, test |
| 6 | `phase start` refuses where declared evidence is missing | `Start`, test |
| 6 | every failing finding is fixed, approved or overridden first | `Decided`, `predecessorAllowsStart` |
| 6 | a decision is recorded on the finding, not rewritten to green | `Decide`, carry-forward test |
| 6 | a decision names a person and a reason | `Decide` refusal |
| 6 | a decision against a stale verdict is refused | `rewriteStatus`, test |
| 7 | G-Test: declared test result successful | `gates.testReport`, `TestKind`, from #212 |
| 7 | G-Test: mapping of acceptance criteria complete | **nothing**; see below |
| 7 | G-Supply compares the vendored tree against the binary's digest | `gates.supply` |
| 7 | G-Freshness reads the predecessors' locks, never the phase's own | `gates.freshness` |
| 7 | an unimplemented gate is written, not skipped | `notImplemented` |
| 7 | `xeno gate` never touches the network | the `verify` job's dependency grep |
| 7 | no gate reads the symbol index | the `verify` job's dependency grep |
| 7 | `XENO_HARNESS` is recorded, never branched on | the `verify` job's grep |
| 7 | the plugin is the vendored one, with no override | `plugin.Dir`, test |
| 6 | every review rule is answered met, deviation or not-applicable | G-Policy |
| 9 | `scope` is checked against the path | G-Rules |
| 9 | predicates are named types, never expressions | the `predicates` table, finding for anything else |
| 9 | `abstract` is required in `learned/provider|org` | `rules.go:247` |
| 13 | `xeno init` refuses a version mismatch | `init.go:64` |

## Architectural properties with nothing asserting them

Each is true today because of how the code is arranged. None would fail a test or a gate
if a later change broke it. This is the list #202 was filed to produce.

| § | property | what would notice |
|---|---|---|
| 4 | Xeno never runs a test, a build or a scanner | nothing — no `exec.Command` outside `internal/external` and `internal/git`, by habit |
| 4 | evidence is declared, never inferred | nothing |
| 5 | a command that changes a hashed file recomputes what it invalidates | nothing in general; `SectionSet` does it for `context_hash` |
| 6 | CI recomputes and compares, and never writes | nothing — the wrapper happens to call only `gate verify` |
| 7 | the local run cannot be skipped | nothing in the repository; it rests on the required check |
| 9 | the commit range is always passed in, never inferred | nothing — `Base`/`Head` are inputs by construction |
| 10 | learning never changes behaviour directly | nothing — the route is a convention |
| 11 | what is sealed is never rewritten — **modification** | `gate verify` reports a divergence |
| 11 | what is sealed is never rewritten — **deletion** | the `verify` job's trail guard, since #193 |
| 12 | a pull request never starts a phase | nothing |
| 12 | the secret filter is outside the model's reach | nothing — true because the runner does the filtering |

Two of these were unguarded until this session and are listed with their new readers,
because the pair is the clearest illustration of the difference: a modification diverges
loudly and a deletion was silent until something was written to notice it.

## What this pass found, and where it went

**#205 — `XENO_PLUGIN_DATA` is specified as "always `.xeno/local/`" and read by
nothing.** The entry point exports it; five packages use the constant instead. A reader
cannot tell a variable with no reader from one whose reader is a constant. Recorded in
A84 and still the state of the one remaining entry in section 7's list.

**#206 — G-Complete runs only at P5**, so an intent that never reaches P5 is never
checked for completeness. XENO-0230 reached `main` with its verification and review
phases missing and every gate green. Known since that intent, and until #206
filed nowhere.

**#208 — declared evidence has no writing command.** Section 5 puts the declaration in
the frontmatter of `output.md`, which the runner writes and `artifacts_hash` covers, so
`evidence attach` has nothing to attach against and no verification phase in this trail
has ever carried evidence. G-Evidence passes correctly over an empty set. Found while
running this intent rather than by the pass.

**Section 13's `--vendor` list names `mcp.json`** and the plugin carries none,
deliberately, because a client reading a declaration of a server that does not exist
fails at startup. The walk copies it the day it exists. A clause whose reader is correct
to find nothing.

**#207 — section 12 requires model, tool and version in every artifact** as the raw
material for a provider register. `tool_version` had no writer until #181 and
`plugin_version` was a constant until #177; both now have one, and nothing checks that
what they say is true. A self-report with a reader that cannot fail.

## What this pass is not

It does not cover `CLAUDE.md` or `CONTRIBUTING.md`, which carry conventions rather than
the process. One of those — the 88-column rule — was row eight of #202's own table and
is being settled separately.

It does not judge whether a clause should have a reader. "Xeno never runs a test" needs
none if nobody is tempted; it needs one the day somebody adds a convenience. That triage
is the next step and it belongs to whoever reads this.
