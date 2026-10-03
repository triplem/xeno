---
intent: github.com/triplem/xeno#176
phase: 03-implementation
created: "2026-10-03T11:39:31Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+bdf4e26.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b147b33e9b26367089427003a15d379a28a00b78e7b1729b5b1cf660d487427a
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

One field, one sum, four tests, and the commit before this one that made the field legal.

**`internal/model`.** `ContextFile.Bytes`, an `int64`, `omitempty`. Its comment quotes the clause
the previous commit added and says what `omitempty` buys: a lock written before the field is
indistinguishable from one whose files are all empty, which is why the check asks whether any entry
carries a size rather than whether the sum is zero.

**`internal/runner`.** The size comes from the directory entry the walk already holds, through
`d.Info()`, rather than from a second `os.Stat`: two reads of one file can see two versions of it,
and the recorded size exists to describe the same read as the hash. A declared link's document is
the one place where a second read is unavoidable, since it is not part of the walk, and a document
that disappears between the hash and the stat is skipped — the same finding #172 added covers it.

**`internal/gates`.** `budget` loses its `os.Stat` loop. It sums `f.Bytes` and tracks whether any
entry recorded one; the finding is produced only where something did. The comment above the
function no longer describes measuring from the tree, because nothing in the gate path does: the
budget was the only place a gate asked for a number rather than for content it hashes.

**Tests.** The fixture now writes a size per file, which is what the two existing byte tests needed
to keep meaning what they meant. Three cases are new, and two of them are the point of the intent.
A file that grows after the phase was inside its budget produces nothing, where before the same
growth produced a finding against a sealed phase. A file the lock names and the tree has lost keeps
the size the lock recorded, so losing a file cannot move a verdict in either direction — the
inverse of the first, and a change from the behaviour before, where a lost file quietly reduced the
total. And a lock with two files and no sizes produces the file-count finding and nothing about
bytes, which is the criterion that protects seventy-nine intents' worth of locks.

**What the trail looks like from here.** The lock of this intent's own 03-implementation carries
`repo_commit`, `rules_applied` with four rules, and `bytes` on every file of its base — which is
empty here, because this repository still writes no profile. The demonstration below is where the
field is visible.

<!-- xeno:section:deviations -->
## Deviations from the design

**A test changed meaning rather than being extended.** `TestAMissingFileIsNotCountedAndNotReported
Twice` asserted that a file the lock names and the tree has lost is skipped in the sum. That was
right when the sum came from the tree and is wrong now: the lock recorded the size, so losing the
file afterwards cannot change what the phase was given. It is now
`TestALostFileKeepsTheSizeTheLockRecorded` and asserts the opposite, which is the behaviour the
specification clause asks for. Recorded because a reviewer seeing an assertion inverted will want
to know whether the behaviour or the expectation was wrong, and this time it was the behaviour —
deliberately.

**A declared link's document needs the second read the rest of the base avoids.** The walk holds
an entry for every file it matched and `d.Info()` is free; a link's document is resolved by path
outside the walk, so its size comes from `os.Stat`. That is the one place in this change where the
hash and the size are two reads, and the window between them is covered by the finding #172 added
for a document that is not there — but the honest statement is that a document replaced in that
window would be hashed as one version and sized as another.

**The budget's comment had to be rewritten rather than amended.** It described measuring from the
tree and gave the reason — the lock carries no sizes — which is now false. The convention for a
paragraph being changed a second time is to replace it, and this is the second time: #171 wrote it,
#176 replaced it.
