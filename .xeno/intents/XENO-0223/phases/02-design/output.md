---
intent: github.com/triplem/xeno#171
phase: 02-design
created: "2026-10-03T09:44:58Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1f2113f8f0ae5296458c1203fe589e62031cb815b6482da598da981a2f466f32
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The changed set is printed and never recorded.** Section 5 enumerates the lock's fields and
there is none for it; the same section says the lock "records the context that was declared, not
everything that was read", so a record of what changed would be a record of neither. It is derived
at `phase start` from the predecessor's lock and the tree, printed beside the next step, and
derivable again by anyone holding both. A derivation that is recorded can drift from its inputs; a
derivation that is printed cannot.

**`repo_commit` comes from `internal/git`, and its absence is a state.** One more read of the
local clone, in the package that already owns the subprocess. A directory that is not a repository
and a repository with no commit both record nothing, because a lock saying which commit a phase
ran against is a claim and an invented value would be a false one.

**`rules_applied` is the effective set with each rule's path and version, and absent where the set
is empty.** It answers the question `rules_hash` cannot: which rules, and which revision of each.
The same `rules.Load` and `rules.Effective` that write the hash produce the list, so the two cannot
disagree about what applied. An empty set writes no field, which keeps the distinction A74 rests
on: a phase with no rules and a phase that resolved an empty set look different in the file.

**The budget is checked in G-Schema against the lock, not against the profile.** The profile
declares the budget and the lock records what the phase was given, so the comparison is between
the two files and belongs where the artifact's own fields are judged. The finding names both
numbers and the phase stays decidable, which is what section 5 asks for in the sentence that calls
it deliberately a finding.

**Bytes are measured from the tree at the time of the check, not stored.** The lock records paths
and hashes, so a byte count has to come from the files. Where a file the lock names is gone, the
count skips it and the finding says so rather than failing: a missing file is G-Freshness's
finding, and two gates reporting one cause is the thing #160 avoided.

**A declared link adds its document to the base, and a missing one is a finding against the
profile.** The link is the one place in the profile that names a specific file rather than a
pattern, which makes an absent target a typo rather than an empty match. The document enters the
base with its hash like anything else, so G-Freshness guards it and a later phase sees it in the
changed set.

**The order is the profile's `include` order, with paths sorted inside each pattern.** Section 5
asks for an order of volatility and says the lock records the assembly order rather than only the
set. The project writes its patterns from stable to volatile and the lock follows, which puts the
judgement where the knowledge is. The tie-break keeps two runs byte-identical, which the hash
needs. A file matched by two patterns takes the position of the first, because a stable prefix is
decided by the first thing that claims it.

**Nothing in the runner reads for the agent, so "reads only what changed" is a report.** The runner
writes files and judges them; it has never read source for a model and does not start now. What it
can do is say which files of the declared base moved, which is what the section means by a repeated
phase knowing what changed — the knowing is the agent's, and the knowledge is the runner's to hand
over.

<!-- xeno:section:alternatives -->
## Alternatives

**Record the changed set in the lock.** A `changed_since_predecessor` list would be readable later
and would survive the session. Rejected twice over: section 5's field list has no such entry, and a
recorded derivation drifts from its inputs the moment anything else is rewritten. The second
reason is the one that would matter even if the enumeration allowed it.

**Compute the changed set against the previous run of the same phase rather than the predecessor
phase.** A repeated phase is literally the same phase run again, which is what "repeated" says.
Rejected because a phase is written once and sealed: the thing that runs again in this process is
the next phase over a tree that moved, and the lock of the phase before it is the only record of
what was given. Where a phase genuinely is re-run, its own lock is the comparison and the same code
answers it.

**Take `repo_commit` from the environment in CI.** The wrapper knows the commit and passing it in
would avoid a subprocess. Rejected: the lock would then say different things locally and in CI for
the same tree, and the runner does not take facts about the repository from its caller — the same
argument that keeps the commit range an explicit input rather than an inferred one.

**Write `rules_applied` as an empty list where no rule is in force.** Simpler, and a reader sees
the field everywhere. Rejected for the distinction A74 is built on: an empty list says a set was
resolved and was empty, absence says there was nothing to resolve. Both are true states of a
repository and they are not the same state.

**Make the budget a red gate.** It is a budget, and a budget nothing enforces is the decoration the
section warns about. Rejected because the section says the opposite in the same breath and gives the
reason: nobody has experience with the number yet. A finding is visible, decidable and recorded,
which is what the mechanism needs before anyone can argue about the figure.

**Store byte counts in the lock so the budget can be checked without the tree.** It would make the
check self-contained. Rejected as a field section 5 does not enumerate, and as a number that would
be wrong the moment a file changed — the hash already covers that, and the count can be taken when
it is needed.

**Order `files` by the profile's include order without a tie-break.** Shorter, and the order inside
a pattern would be whatever the walk produced. Rejected because the lock is hashed: a walk order
that depends on a filesystem would make two runs over one tree disagree, which is the property
every hash in this project exists to prevent.

**Give the agent the changed files' contents rather than their paths.** The saving would be real and
immediate. Rejected because the runner has never read source for a model, and because the lock is a
statement about what was declared: a runner that handed over content would be deciding what the
phase reads, which is the agent's decision recorded in the profile.

<!-- xeno:section:impact -->
## Impact

**The lock becomes readable rather than only verifiable.** `rules_hash` lets somebody who holds the
tree recompute what applied; `rules_applied` lets somebody who holds the file read it. Those are
different audiences, and section 5 asks for both — the hash in the frontmatter and the list in the
lock.

**A phase gains a sentence it did not have.** Starting a phase after a tree moved now says which
files of the declared base moved. In a repository with no profile it says nothing at all, which is
every project today and this one.

**Two of section 5's four missing fields stop being missing**, and the two that remain have an owner
each: `plugin` with G-Supply, `tools` with WP15. A reader comparing the section with a lock now
finds two absences, both explained.

**The budget stops being decoration.** A project that declares one and exceeds it learns so at the
gate, with both numbers in the finding. Nothing blocks, which means the first projects to use it
generate the experience the section says nobody has yet.

**A declared link does something.** Until now the field was parsed and ignored, which made the
paragraph forbidding inference describe a feature that did not exist. A profile that names a
document now gets it, hashed, guarded by G-Freshness, and visible in the changed set.

**The adoption question lands on your desk with figures.** Writing a profile for this repository
makes G-Freshness's second half live: an earlier phase whose base a later phase changed carries a
finding that needs a person to release. The demonstration measures how many such findings a normal
intent would produce here, which is the number the decision turns on and which nobody has had.

**Nothing in the trail moves.** The new fields appear in locks written from now on; the twenty-three
intents behind this one carry what they carried, and their `files` are still empty because no
profile existed when they ran.
