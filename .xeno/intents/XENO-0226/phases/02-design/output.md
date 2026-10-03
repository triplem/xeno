---
intent: github.com/triplem/xeno#176
phase: 02-design
created: "2026-10-03T11:35:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d45cd1ea2159b08b459e93f537886c91ae1ee3bb26b0178890fd18554ebc367a
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

**`bytes` is an `int64` on `ContextFile`, `omitempty`.** The size of a file as the walk saw it.
`omitempty` is what makes a lock written before the field indistinguishable from one whose files
are all empty — which is a state nothing in this repository has and which the check has to treat
the same way either way: nothing to sum means nothing to say.

**The size and the hash are taken from the same walk.** `informationBase` already has the entry
open; the size comes from the directory entry rather than a second `os.Stat`, so the two numbers
describe one read of one file and cannot disagree about which version they saw.

**The check sums what the lock recorded and measures nothing.** `budget` loses its `os.Stat` loop.
A lock with no sizes sums to zero, and zero is the signal to say nothing rather than a context of
zero bytes: the condition is on whether any entry carries a size, not on the sum. That distinction
is the whole of the third criterion and it is one line.

**A lock with some sizes and not others is judged on the sum of what it has.** No entry is
estimated and no finding mentions the gap. The runner cannot produce that lock; a person editing
one can, and a check that guessed would be guessing about a file it cannot see.

**Nothing in the gate path measures a file after this.** The budget was the only place a gate
asked the tree for a number rather than for a hash, and that is now true of nothing. It is worth
saying because it is a property of the gate path rather than of this check: a verdict is computed
from artifacts and hashes, and a size was the one exception.

<!-- xeno:section:alternatives -->
## Alternatives

**Keep measuring, and exclude sealed phases.** The check could ask whether the phase has a verdict
and skip the budget where it has one. Rejected: it makes a gate behave differently on a
recomputation than on the first run, which is the property `gate verify` exists to deny — the same
binary, the same tree, the same verdict.

**Record one total for the whole base rather than a size per file.** Smaller, and the budget only
ever needs the sum. Rejected because the specification clause says each entry carries the size of
the file, and because a per-file number survives a file leaving the base: a reader comparing two
locks can see which file grew, where a total says only that something did.

**Treat an absent size as zero and report against the sum.** One fewer condition. Rejected on the
third criterion: seventy-nine intents' locks record nothing, they would all read as zero-byte
contexts, and every one of them would pass its budget quietly. A quiet pass where a loud finding
used to be is the worst available outcome, and it is the one a careless implementation produces.

**Backfill the sizes into existing locks.** It would make the trail uniform and the check
unconditional. Rejected for the reason every rewrite in this project is rejected: the lock describes
what the phase was given, `artifacts_hash` covers it, and a number nobody recorded is not evidence.

**Take the size with a second `os.Stat` after the hash.** Simpler to read. Rejected because the walk
already carries the entry, and two reads of one file can see two versions of it — which is exactly
the inconsistency the recorded size exists to prevent.

<!-- xeno:section:impact -->
## Impact

**A sealed phase's standing stops depending on the tree.** After this, every number a verdict rests
on is one the artifact recorded. That was already true of hashes, of the rule set and of the
information base; the byte budget was the exception, and it is now the rule.

**The gate path asks the tree for nothing but content it hashes.** A property worth stating once:
what a gate reads is artifacts and the files they name, and the one place it asked for a measurement
is gone.

**Every lock from here is three fields per file where it was two.** A lock for a twenty-three file
base grows by about a line and a half per file in diff terms and nothing in bytes that matters. The
trail behind this intent is unchanged.

**The second of the two conditions on the profile experiment is met.** #172 closed the first. The
experiment — one intent, a narrow profile, the releases counted — is now unblocked.

**One thing a reader of an old lock still cannot tell.** Whether its context was inside its budget.
The sizes were never recorded, so the question is unanswerable for seventy-nine intents, and the
check says nothing rather than guessing. That is the honest state and it is the cost of having
measured before recording.
