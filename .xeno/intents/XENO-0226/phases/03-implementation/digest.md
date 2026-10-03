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
---
One field, one sum, four tests. `ContextFile.Bytes` is an `int64` with `omitempty`, taken from the
directory entry the walk already holds rather than a second stat, because two reads of one file can see
two versions of it; a declared link's document is the one place a second read is unavoidable, and the
window between hash and stat is covered by #172's finding. `budget` loses its `os.Stat` loop, sums what
the lock recorded, and produces a byte finding only where some entry recorded a size — so a lock written
before the field says nothing rather than reading as a zero-byte context. Nothing in the gate path
measures a file after this, which is a property of the path: the budget was the only place a gate asked
for a number rather than for content it hashes. Three tests are new and two are the point: a file that
grows after the phase was inside its budget produces nothing, where the same growth used to produce a
finding against a sealed phase, and a lost file keeps the size the lock recorded. Three deviations,
including one test that had to assert the opposite of what it asserted — right about the old behaviour,
and renamed rather than edited so the next reader can tell which changed.
