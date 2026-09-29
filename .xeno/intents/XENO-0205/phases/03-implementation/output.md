---
intent: github.com/triplem/xeno#65
phase: 03-implementation
created: "2026-09-29T11:39:29Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+edd2b2e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 447634d6097c1a887a145ab0c3ab894fef3023405b7cceef71d76bdbdf6e91af
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`internal/cost`. `TranscriptTotals` sums the usage of every assistant message of a JSONL
transcript, skipping a record without usage and a half written last line, because a
transcript is written while a session runs. `Append` and `Read` hold the ledger.
`ForPhase` totals a phase as the sum of differences between consecutive lines of the
same session, so the first line of a session contributes nothing: what it counts was
spent before the ledger existed. `LivePhase` reads the intent and the phase out of the
shell file `phase start` writes.

`cmd/xeno`. `cost turn` reads a hook's JSON on standard input, takes `transcript_path`
and `session_id`, and appends a line naming whatever phase is open, or `none`. Every
failure path returns zero: unreadable JSON, an absent transcript, one with no usage, an
unwritable ledger.

`internal/runner`. `Finish` calls `writeCost` before it evaluates, and `writeCost`
writes nothing where the ledger has nothing for the phase. `internal/model` gains
`Cost`, the schema of section 11 without `cost_usd`, which that section makes optional
and v2 authoritative.

`.xeno/plugin/hooks/README.md` and `.claude/settings.json`. The hook, and this
repository wired to it, so the thing being built is used by the project building it. A
`Stop` hook, once per turn.

**What it produced, running against this session's own transcript.** Two invocations of
the command the hook runs appended two lines naming `03-implementation` of this intent,
with cumulative output tokens of 895,900 and 909,832. The difference, 13,932, is what
the turn between them cost, and it is the first token figure this repository has ever
recorded.

**One repair that belongs to no work package.** The comment above `GateRun` was
duplicated, both lines identical, introduced by the insertion that added `writeDigest`
in #125: that edit used the line as an anchor and re-appended it. It is a comment and it
was wrong in the way #111 was about, so it is removed here rather than left for somebody
to read twice. The deviations section says so.

<!-- xeno:section:deviations -->
## Deviations from the design

None in the design. The hook, the ledger's cumulative lines, the `none` for an
unattributed turn, the failure-open command and the absent file for a phase with nothing
are as P2 decided, and the measurement that drove them was taken before P0.

The order of the work was right this time: the 7% figure was measured while answering
whether the writer was possible at all, and P0 to P2 were written before any code. That
is the first intent since XENO-0201 where the code did not precede its own design.

**A repair carried in this intent that is not this intent's subject.** The duplicated
`GateRun` comment described above. The third standing rule says every change belongs to
a work package and an intent, and this one belongs to neither: it is a defect I
introduced in #125 and found while editing the same neighbourhood. Carrying it here is a
judgement that a separate issue and pull request for a duplicated comment line costs
more than it records. It is in its own commit so that it can be read on its own.

**The hook is not yet firing for the session that wrote it.** `.claude/settings.json` is
read when a session starts, so the two ledger lines above came from invoking the command
by hand with a real hook payload. The mechanism is exercised; the wiring is not, and the
verification says which.
