---
intent: github.com/triplem/xeno#171
phase: 03-implementation
created: "2026-10-03T09:55:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55886c6db971e14e9d7d66caba001b5f07b9b519f2378be207f0992a724c3ae3
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

Four mechanisms, one new exported function, five files changed, no new dependency.

**`internal/model`.** `ContextLock.RepoCommit` and `ContextLock.RulesApplied`, with
`AppliedRule` carrying the path and the version counter section 5 writes. Both `omitempty`,
because absence is a state: a lock naming no commit is a repository that has none, and absent
`rules_applied` is a phase with nothing to resolve where an empty list would say a set resolved to
nothing. The comment on `Files` now says what its order means.

**`internal/git`.** `Head(root)`, one `git rev-parse HEAD`, returning an empty string for every way
it can fail — not a repository, no commit yet, no git — because the caller records a fact about the
repository and an invented one would be a false claim.

**`internal/runner`, the base.** `informationBase` buckets the walk by the first `include` pattern
that claims each file and emits the buckets in the profile's own order, sorting paths inside each.
So the project decides volatility by the order it writes its patterns, the lock records that
decision, and two runs over one tree stay byte-identical, which the hash needs. A file two patterns
match takes the position of the first. Then every declared link's document joins the base with its
hash, last, whether or not `include` matched it: a link is the most specific thing in a profile and
the most likely to move, and a declaration that put nothing in the base was ornamental.

**`internal/runner`, the lock.** `Start` writes `repo_commit` from `headCommit` and `rules_applied`
from `rulesApplied`, which calls the same `rules.Load` and `rules.Effective` that produce
`rules_hash` — so a lock listing three rules beside a hash over four is not a state this code can
reach.

**`internal/runner`, the changed set.** `ChangedSince(key, phase)` reads the preceding phase's lock
and compares every path in its base against the tree: a different hash is listed, a missing file is
listed as gone, and the order is the base's own. Nothing is recorded, because section 5's field
list has no entry for it and the same section says the lock states what was declared rather than
what was read.

**`cmd/xeno`.** `phase start` prints the changed set under one line naming the phase it is measured
against, and prints nothing where nothing moved — which is also what a repository with no profile
gets, since it declared no base to move.

**`internal/gates`.** `budget(c)` compares the profile's `budget` against the lock's `files`,
reporting the count against `files` and the measured size against `bytes`, each naming both numbers
and pointing at the profile. It is appended to G-Schema's findings, which is the gate section 5
names and which makes it decidable like any other finding rather than a block. Bytes are measured
from the tree when the check runs, since the lock carries hashes and no sizes; a file the lock names
and the tree has lost is skipped, because its absence is G-Freshness's finding and one cause
reported twice is what #160 avoided.

**Tests.** `internal/runner`: the include order, and that reversing the profile reverses the base; a
declared link's document in the base with a hash; `repo_commit` absent in a fixture that is no
repository; `rules_applied` with the path and the version, agreeing with the resolution, and absent
where no rule is in force; the changed set after an edit, after a removal, with no predecessor and
with no profile. `internal/gates`: no profile and no budget are no finding; over the file budget and
over the byte budget, each naming both numbers; within budget, nothing; a lost file neither counted
nor reported twice; and the finding arriving through G-Schema rather than through a function nobody
calls.

**One existing test changed.** `TestGivenFilesAreComparedAgainstTheTree` asserted the alphabetical
order of the base. It was right about the old behaviour and wrong about section 5, so it now asserts
the include order with the reason in a comment.

<!-- xeno:section:deviations -->
## Deviations from the design

**A test from an earlier intent had to be corrected rather than extended.**
`TestGivenFilesAreComparedAgainstTheTree` asserted that the information base comes out in
alphabetical order. It was a true statement about the runner and a false one about section 5, which
asks the lock to record an order of volatility. It now asserts the include order, with the reason in
a comment. Recorded because a reviewer seeing an assertion inverted will want to know whether the
behaviour or the expectation was wrong, and it was the expectation.

**The byte budget from section 5's own example is far too small for a repository like this one.**
The example in the section declares `bytes: 400000`. A profile written for this repository with four
include patterns resolves 23 files and **709,921 bytes**, so the first thing the demonstration did
was exceed the example figure by three quarters. The mechanism is right and the illustration is
misleading: a project copying the example will be over budget on its first phase. That is a finding
about the document, raised here rather than fixed, and it is the kind of thing the section itself
predicted when it said nobody has experience with the number yet.

**The changed set cannot be seen until the earlier phase's findings are decided.** With a profile in
force, editing files of the base turns the earlier phase red through G-Freshness, and `phase start`
refuses to start the next phase on a red predecessor — so the report that tells a phase what moved
is unreachable until somebody has released the findings about the same files. In the demonstration
that took four approvals before the next phase would start at all. It is the sequence working as
specified and it is a worse sequence than I expected: the mechanism that saves reading is gated
behind a decision about the reading.

**`ChangedSince` reports a removal as a path with "(gone)" appended**, which puts a word inside what
is otherwise a list of paths. The alternative was a second list, which the one-line report cannot
carry, or dropping removals, which would be silent about the case most likely to matter. It is a
string in output rather than in an artifact, so nothing parses it, and it is ugly.

**The budget's byte count is measured from the tree, which means it moves after the phase.** A phase
inside its budget when it started can be over it a week later because a file grew, and `gate verify`
would then report a finding against a sealed phase. The lock carries no sizes, so the alternative
was a field section 5 does not enumerate. It is in the gaps, and it is the same class of problem
A74 solved for rules with the recorded hash — which this field cannot use, because it has none.
