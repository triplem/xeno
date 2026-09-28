---
intent: github.com/triplem/xeno#118
phase: 03-implementation
created: "2026-09-28T20:15:46Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57e9207.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 9c382f5800735d9bbee76a2276fa7ee7ccc2390182a13d25be29bb343eb76ad2
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

`internal/runner/runner.go`. `IntentsRoot`, the `IntentSummary` type, `Intents()` and
`summarise()`.

`Intents()` reads the intent directories, summarises each and sorts by `created` with
the key as the tie break. A missing directory returns an empty slice and no error,
because a repository that has started no intent is not in error. `summarise()` reports
rather than fails: an unreadable `intent.yaml` or a missing `created` becomes a row
carrying the reason, and the furthest phase with a verdict comes from the existing
`Status` path.

`cmd/xeno/main.go`. `intent status` loses `needsKey`, `cmdIntentStatus` branches on the
empty key, `cmdIntentList` prints the rows, and the usage line shows the argument as
optional and says what happens without it.

`CLAUDE.md`. The key paragraph is replaced: a key is the next number of a sequence
beginning at `XENO-0200`, the issue lives in the `intent` field, and the old keys stay
for the reason the paragraph already gave, with the gap explained.

`internal/runner/runner_test.go`. Six tests and two helpers. The order over keys and
dates that disagree, two intents of one day, the undated row, the stability of ties, the
empty repository, and the row naming the furthest phase.

**The defect the criteria caught.** The first version truncated `created` to its date
inside `summarise`, then sorted on the truncated value. Two intents of one day therefore
fell through to the key tie break and came out in the order the listing exists to
correct: XENO-0107 printed before XENO-0108. The whole value is kept and sorted, the
command prints the first ten characters, and both the runner and the command carry a
comment saying why, because the truncation looked correct and produced the original bug.

This intent is `XENO-0200`, the first key of the sequence it introduces.

<!-- xeno:section:deviations -->
## Deviations from the design

None in the shape. The listing, the sort, the tie break, the reporting of an unreadable
intent, the optional argument and the replaced paragraph are as P2 decided them.

Two corrections during the work, both caught by the criteria rather than by the
compiler.

The date truncation described in the changes section. AC3 names XENO-0107 after
XENO-0111, and the first version put it before XENO-0108; the criterion was written from
the issue and it is what failed.

One test asserted the undated intent sorts first, and the fixture's own intent carries
no `created` either, so two undated rows tie and break on the key. The assertion became
that the undated row sorts before a dated one, which is what AC4 actually asks. A test
wrong about the fixture rather than about the code.

The order of the record is the order of the work. P0 to P2 before any code, the code
inside P3.
