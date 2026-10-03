---
intent: github.com/triplem/xeno#171
phase: 01-requirements
created: "2026-10-03T09:43:58Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55593de266e98b4c1610f888597ae98188f0be680bdaaf28a635dbc0a463423f
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

**A repository with no profile behaves exactly as it does today.** No base, an empty `files`, no
budget finding, no changed set to report, and the same verdict with the same rows. Twenty-three
intents are in that state and none of their locks may change meaning.

**The lock carries `repo_commit`.** The commit the phase was started against, as the repository
reports it. A repository that is not a git repository, or has no commit yet, records nothing rather
than an invented value.

**The lock carries `rules_applied`, one entry per rule of the effective set**, each with the path
it was read from and its own `version` counter. Where no rule is in force the field is absent
rather than an empty list, because an absent field is a phase with no rules and an empty list is a
claim about a set.

**`rules_applied` and `rules_hash` agree.** The same resolution answers both, so a lock listing
three rules and a frontmatter hash over four would be a contradiction nothing could explain.

**A recorded context over its budget is a G-Schema finding and the phase is not blocked by it.**
The finding names the budget and what was recorded — files against `files`, bytes against `bytes` —
and the phase is red in the sense every other finding makes it red: decidable, releasable, visible.
Where no budget is declared there is nothing to exceed and no finding.

**A declared link's document is in the information base.** A profile whose `links` name a document
puts that document in the base with its hash, whether or not `include` matches it, and a link
naming a document that does not exist is a finding against the profile rather than a silently
missing file.

**The `files` order follows the profile's `include` order.** A file matched by the first pattern
comes before a file matched by the second, and within one pattern the order is by path, so the
result is deterministic and the project decides what is stable. The lock therefore records an
assembly order rather than an alphabet.

**`xeno phase start` names what changed.** Starting a phase whose predecessor recorded a base
prints the files of that base whose hashes no longer match the tree, and says nothing where nothing
changed. The comparison is between the predecessor's lock and the tree, so it needs no new field
and no second record.

**Nothing is written that section 5 does not enumerate.** No "what changed" field, no `plugin`
block, no `tools`. The report is output and the two absent fields stay absent.

**The mechanism is demonstrated against this repository on a copy.** A profile written for this
tree, a phase started, the base resolved, the budget exceeded on purpose once, a link followed, the
changed set reported after an edit — with the figures recorded, because the adoption question needs
them.

**Nothing else moves.** No rule set changes, so no `rules_hash` changes. No sealed artifact is
rewritten; the new fields appear in locks written from now on. `./xeno gate verify` stays at exit 0
over the trail and the suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No `plugin: { version, sha256 }`.** It is what G-Supply compares against, the gate is
unimplemented, and a field written for a reader that does not exist is a field nobody maintains.

**No `tools`.** WP15 owes it, #155 settled its shape, and inventing it here would mean deciding the
hash definition that issue already assigned to the piece that writes it.

**No blocking on the budget.** Section 5 says it is a finding and not a red gate in that sense, and
says why: blocking against a number nobody has experience with yet is the wrong way round.

**No inferred link.** The specification forbids it in the paragraph that declares the field.

**No profile for this repository.** The mechanism is exercised on a copy. Adopting it makes
G-Freshness's second half live against every intent here, and a finding on an earlier phase needs a
person to release it — an agent that wrote the finding is not the second person the override is for.

**No measurement of the saving.** What re-reading only the changed files saves in tokens is WP15's
and WP20's, and this half of WP8 was scoped in the plan as the one that needs no baseline.

**No change to what `files` means.** It stays the declared base and not a record of what was read,
which is the sentence section 5 is most explicit about.

<!-- xeno:section:constraints -->
## Constraints

**Section 5's enumeration is a budget.** The lock's fields are the ones written out there plus the
two section 7 needs, and this piece may add only from that list. The changed set is therefore a
report and not a record, which is also the honest shape: it is derived from two locks and the tree,
and recording a derivation would let it drift from its inputs.

**The lock is written once and never refreshed.** Section 5 says so and says why. Everything added
here is written at `phase start` and nothing rewrites it afterwards.

**The order has to be deterministic, because the lock is hashed.** The profile's `include` order
gives the volatility order the section asks for and a tie-break by path inside each pattern keeps
two runs identical.

**A budget finding is a G-Schema finding, which means it is subject to a decision like any other.**
It names the numbers on both sides, because a finding that says "over budget" without saying by how
much is a finding nobody can act on.

**An absent field means absent.** No empty list where a project declared nothing, no zero where
nothing was counted: the distinction between "no rules were in force" and "a set was resolved and
it was empty" is the one A74 is built on, and it holds here too.

**No new dependency**, 88 columns, SPDX, `gofmt`, `go vet`, the suite, and `./xeno gate verify` at
exit 0.

**One intent, one branch, one issue.** `171-the-lock-and-the-changed-set`, #171, labelled wp8, off a
`main` that carries WP4 and the agent layer.
