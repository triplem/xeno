---
intent: github.com/triplem/xeno#212
phase: 01-requirements
created: "2026-10-05T15:48:43Z"
schema_version: "1.0"
runner_version: dev+09e2aa6
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2b794112bcad47eb2752aeb83d3ac56c66dc14bde2df49c4fe0e0d07d07201af
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

1. G-Test is no longer `notImplemented` in the gate table, and reports `pass`, `fail` or
   `pending` at P4 like G-Build does.

2. A declared `test-report` whose result is not `pass` is a finding naming the job. One
   awaiting its artifact is `pending`, not failed, which is the path G-Build already has and
   A60 explains.

3. G-Build and G-Test share the comparison rather than each carrying a copy of it; the only
   difference between them is the kind.

4. `TestKind` is a named constant beside `BuildKind`, so the spelling is defined in one place
   and G-Schema's closed set is the only authority on it.

5. `./xeno gate verify` exits 0 with the 381 verdicts that exist now intact, growing only by
   this intent's own judged phases. This is the criterion the design turns on and it is a
   command.

6. `docs/clause-readers.md` gains a row for section 7's G-Test clause, naming the result half's
   reader and saying the mapping half has none, with its tool-requirement count moving from 35
   to 36 and a sentence saying this row arrived after the pass — the file's own precedent.

7. The mapping half is an issue of its own, carrying both measurements: 46 of 55 P1 artifacts
   with no numbered criteria, and that identifying a criterion is a convention the documents do
   not have.

8. Tests cover: a declared test report with `pass` is green, with `fail` is a finding, pending
   with nothing attached is `pending`, pending with an attached `fail` is a finding, a
   `build-log` is not read by G-Test and a `test-report` is not read by G-Build, and an artifact
   declaring neither leaves both green.

9. `go build`, `go test ./...`, `go vet ./...` pass and `gofmt -l` outside `vendor/` prints
   nothing.

10. One commit, `Closes #212`, and the issue carries `wp7`.

<!-- xeno:section:non-goals -->
## Non goals

Not the mapping half. It needs an acceptance criterion to be identifiable, which no document
provides, and a check over 46 artifacts that never had one. Its own issue with the figures.

Not a numbering convention for acceptance criteria. Requiring them to be a numbered list is an
addition to what section 9 says, so a specification change first and a person's commit.

Not running a test suite. The plan says G-Build and G-Test read declared results rather than
running anything, and the gate path opens no socket and starts no process.

Not a second gate. The gate list is a budget by section 7, and section 7 gives G-Test one row
with two halves in it.

Not a change to what G-Build reads. `BuildKind` stays `build-log` and G-Build's behaviour on
every existing artifact is unchanged, which the tests assert rather than assume.

Not a claim that G-Test is now read in full. A green G-Test will have judged the result half
only, and the audit row is where that is written down rather than in a commit message nobody
reads twice.

Not a re-judgement of anything. The result half was probed against the whole trail before this
intent was planned, at exit 0, and criterion 5 is the same command run again.

<!-- xeno:section:constraints -->
## Constraints

The trail bounds the check, and the bound was measured rather than estimated. The result half
passes over 381 verdicts because the pending-and-attached path carries the results; anything
stricter than G-Build's own comparison has to be re-measured before it is written.

A60 binds the pending case. "A pending evidence item owes no `result`", so an item whose
artifact has not arrived is `pending` and not a failure, and a gate that read the declaration's
empty `result` as a non-pass would fire on seventeen sealed artifacts — which is exactly what
the first audit script predicted and the probe disproved.

Section 5's `not-implemented` is what G-Test gives up. It means "a check a runner did not
perform", and leaving it there was the alternative the maintainer weighed; once the gate reports
a result, the honest statement of what it did not judge has to live somewhere, and
`docs/clause-readers.md` is the only place in this repository whose purpose is holding it.

The audit document is dated. It records a pass made on 2026-10-03 against `d19a1ca` and says a
reader named by symbol is wrong the day the symbol is renamed; a row added now has to say it
arrived later, as the one existing late row does, or the document starts claiming a pass it did
not make.

The gate list is a budget. Section 7 fixes the fourteen gates and their phases, so this intent
may change what a row reads and not how many rows there are.

One intent, one branch, `Closes #212`, and the issue carries `wp7`.
