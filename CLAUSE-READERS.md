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
| tool requirement | 35 |
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
| 6 | a question carries two to four options and one free entry | G-Questions |
| 6 | a decision carries id, chosen, rationale and `decided_by` | G-Questions' decision shape |
| 6 | `phase start` refuses a predecessor with no completed verdict | `predecessorAllowsStart`, test |
| 6 | `phase start` refuses a second start of a running phase | the run marker, test |
| 6 | `phase start` refuses where declared evidence is missing | `Start`, test |
| 6 | every failing finding is fixed, approved or overridden first | `Decided`, `predecessorAllowsStart` |
| 6 | a decision is recorded on the finding, not rewritten to green | `Decide`, carry-forward test |
| 6 | a decision names a person and a reason | `Decide` refusal |
| 6 | a decision against a stale verdict is refused | `rewriteStatus`, test |
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
