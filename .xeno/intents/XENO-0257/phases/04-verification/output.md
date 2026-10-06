---
intent: github.com/triplem/xeno#257
phase: 04-verification
created: "2026-10-06T07:58:20Z"
schema_version: "1.0"
runner_version: dev+77a35b1.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 502737ad99e3ee986a40d94604daac626cde2b5df989e11a6dcc9c5694f62994
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - format: other
      job: go-test
      kind: test-report
      path: evidence/go-test.txt
      produced_by: go test ./...
      result: pass
      sha256: 6baf5f923d00931217e29de2c18ba061f36a60dfa261766c1c0821c0bff4d342
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

No Go test is added and none would mean anything. The change is thirteen files under `examples/`
that nothing loads: `template.Load` reads `.xeno/config/templates` and `.xeno/plugin/templates` and
never this directory. A test would assert that a file this repository does not read says what it
says.

What stands in for it is one script, three counts and a reading:

| criterion | what answers it |
|---|---|
| 1, six directories, both files each | `find examples/templates -type f`: six `template.yaml`, six `strings.en.yaml`, one README |
| 2, ids and headers match the shipped set | a script parsing both and asserting the section id lists are identical and `id`, `version` and `phase` match — for all six |
| 3, six required and eleven not | the same script, counting required flags from the parsed files: `problem`, `acceptance-criteria`, `changes`, `test-mapping`, `results`, `release-notes` |
| 4, `release-notes` kept and says what reads it | reading the review override's header, which names `release-notes-are-filled` and `section-non-empty` |
| 5, the frontmatter point | the same header, saying G-Policy reads `review_checklist` whatever the section's flag is |
| 6, kept by decision distinguished from kept by a reader | the requirements and verification headers, which say "by decision, not by a reader" and name #258 |
| 7, each file says what it drops and why | reading all six headers |
| 8, the README | reading it: the nine-intent figures, the measurement, how to adopt a subset, what to collect |
| 9, everything parses and headings cover the ids | the same script, which also asserts every id the template defines has a heading in its bundle |
| 10, nothing adopted | `ls .xeno/config/templates` — absent — and `./xeno gate verify` at exit 0 with the verdict count unchanged by the fixture |
| 11, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 12, one commit | the commit itself |

**Criterion 2 is the one that needed a script rather than eyes.** The twelve files were generated
from the shipped set, so the ids are right by construction; the script checks that claim
independently by parsing both trees and comparing, because "generated from" is an assertion about a
script that also needs testing. It asserts three things per template — id lists identical, header
fields equal, every id has a heading — and all six pass.

Criterion 10 is two checks because the fixture's inertness has two halves: nothing reads
`examples/`, which is a property of `Load`, and nothing in this repository has adopted it, which is
a property of the working tree. The first is read in the code and the second is `ls` plus the
verdict count.

**The evidence was declared from a settled file.** The suite ran to completion, the process was
confirmed gone, and the hash was taken twice and compared before anything declared it — the method
XENO-0255's mistake taught, where a hash was bound to a log still being written and cost two
phases.

The suite proves nothing about this change and this phase says so rather than letting a green run
stand as evidence. No Go code reads a template under `examples/`.

<!-- xeno:section:results -->
## Results

Eleven criteria met, one pending until the commit.

1. Met. Thirteen files: six `template.yaml`, six `strings.en.yaml`, one README.

2. Met. The script asserts, for all six, that the section id lists are identical to the shipped
   templates' and that `id`, `version` and `phase` match. All six pass.

3. Met. Six required — `problem`, `acceptance-criteria`, `changes`, `test-mapping`, `results`,
   `release-notes` — counted from the parsed files; eleven `required: false`.

4. Met. The review override's header says **READ BY A GATE** and names `release-notes-are-filled`
   and its `section-non-empty` check as the only section id any rule names.

5. Met. The same header says the `review_checklist` frontmatter is read by G-Policy whatever the
   section's flag is, so the prose becomes optional and the answers do not.

6. Met. The requirements and verification headers say "by decision, not by a reader" and name #258.

7. Met. Each of the six headers lists what it drops with a reason for each.

8. Met. The README carries the nine-intent figures, the keeps and their grounds, the measurement
   that one of seventeen is read by a gate, how to adopt a subset, what to collect, and what not to.

9. Met. Everything parses, and every id each template defines has a heading in its bundle.

10. Met, in both halves. `.xeno/config/templates` does not exist, and `./xeno gate verify` exits 0
    with the verdict count moving only by this intent's own judged phases. Nothing under
    `examples/` can be loaded: `template.Load` reads the project config and the plugin and never
    this directory.

11. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
    nothing, and `go test ./...` is 18 packages `ok` with 0 failures.

12. Pending. The commit comes after this phase is judged.

**The figure the fixture rests on, restated because it is the result and not a detail.** Of the
seventeen required sections, one is read by a gate and sixteen are read by the next-step suggestion.
Established three ways: no Go file outside the tests names a section by id, only `release-notes`
appears across the shipped rules and the examples, and `sectionNonEmpty`'s own comment says the
general case and cites A73. The first of those is the one that could have been wrong — the apparent
hit was `Scope` in `internal/rules/rules.go`, which is a rule's scope field and not the `scope`
section, and taking it at face value would have made the candidate wrong about what it may touch.

<!-- xeno:section:gaps -->
## Gaps

**Nobody has run it, which is the whole of what this intent does not know.** The candidate parses,
its ids match, and the one section a gate reads is still required. Whether six sections produce an
artifact a person can review in six months is exactly the question the fixture exists to ask, and
asking it needs an intent run through the reduced set. The README says so; this phase cannot.

It can age with nothing to notice. If a shipped template gains a section the candidate will not
have it; if `release-notes-are-filled` stops being the only rule naming a section, the header that
says it is the only one becomes false. Nothing reads these files — which is what makes them safe and
also what makes them unverifiable — so the defence is the date in the README, the same defence
`docs/clause-readers.md` has and the same one #254 had to repair there.

The eleven drops are a judgement and the reasons are assertions. "Non goals bound a design, and a
change with no design has none to bound" is the kind of sentence that sounds right and has not been
tested against an intent that turned out to need its non-goals. A reviewer can disagree file by
file, which is why each reason is in the file rather than pooled in the README, and disagreement is
cheaper than the measurement.

The measurement the fixture enables is not protected from the mistake this session made. The README
says not to measure a timing, and nothing enforces that; the next person collecting figures can
still read `context.lock.yaml` and report agent latency as process cost, as four comments on #117
did. A paragraph is the whole guard.

Whether a shortcut should exist at all is untouched. This provides a candidate and the method; the
plan's position is that the shape follows from measurement, and a measurement of one intent against
one candidate is one data point against nine. The README asks for the comparison and does not say how
many runs would settle it, because nothing here knows.
