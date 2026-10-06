---
intent: github.com/triplem/xeno#267
phase: 04-verification
created: "2026-10-06T20:10:15Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 825c4746e322ee6e68f365a4588f13a2d956385a755e34ae09f6000c7ea2b22b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - kind: test-report
      result: pass
      produced_by: go test ./..., gofmt -l ., go vet ./..., xeno gate verify
      sha256: 8e7932ff1063ba23a2564692417f568fb6762b18e0d1c16655147302a4fb07ca
      path: evidence/go-test.txt
      job: test
    - kind: other
      produced_by: the 0-of-122 re-measurement over the trail
      sha256: 05059fb8980a827ab23de5d8972f6a0fa3e9635530a9adb70e2a40f9fa115c6c
      path: evidence/lock-files.txt
      job: checks
    - kind: other
      produced_by: the three paragraphs and three comments, read in the files
      sha256: 3cf778f837434294443732492af8cc581c2ee46e64d53651c61452b59ef6f4a0
      path: evidence/checks.txt
      job: paragraphs
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Fourteen criteria, by number, all passing. Nine of them are read by a person, because the
deliverable is prose and a check that asserts a paragraph exists asserts nothing about whether it
is true.

| # | what it asserts | how it is checked | result |
|---|---|---|---|
| 1 | section 5 says `files` is empty at P0 | the paragraph at `process-definition.md:579`; `evidence/checks.txt` | pass |
| 2 | it says why the order is forced | the same paragraph, naming `scope set` and the phase directory | pass |
| 3 | it says why nothing fills the lock afterwards | the same paragraph, giving #215's own reason | pass |
| 4 | it names the two checks and says they apply from P1 on | the same paragraph; each of the two says so where described | pass |
| 5 | it says what P0 still binds | the same paragraph, final sentence | pass |
| 6 | the schema block's `files` line points at it | `process-definition.md:552` | pass |
| 7 | the budget clause names the comparison | the paragraph at `process-definition.md:650` | pass |
| 8 | the staleness clause says P0's exclusion is not a third limit | `process-definition.md:1057`; "Two limits" is unchanged at 1043 | pass |
| 9 | it says what escapes | the same paragraph, final sentence | pass |
| 10 | `informationBase` no longer says a phase older than the rule | `runner.go:549`; two paragraphs where there was one | pass |
| 11 | `staleReads` no longer says it either | `gates.go:1387` | pass |
| 12 | `docs/clause-readers.md` carries the three clauses | a paragraph, not three rows; the count moves 126 to 129 | pass, with a deviation |
| 13 | no behaviour changes | `go test ./...`, `gofmt`, `go vet`, `gate verify` at 475 verdicts; no test file in the diff | pass |
| 14 | the specification commit is its own, and first | `git log`: `3f1fab3` touches only `process-definition.md`, `5dbba24` follows | pass |

Criterion 12 passes against the issue's "done when" and not against its own wording. P1 asked for
three rows and the document's taxonomy asks for a paragraph; P3's deviations says why, and the
row count stays at forty-one.

<!-- xeno:section:results -->
## Results

## The measurement, re-taken

#267's table was measured on 2026-10-06 over 117 P0 locks. It was re-taken on this branch, over
the trail as it now stands, because a claim this intent's whole argument rests on should not be
inherited from the issue that raised it:

| | locks | with a `files` list | without |
|---|---|---|---|
| P0 | 122 | **0** | 122 |
| P1 | 71 | 22 | 49 |
| P2 | 71 | 22 | 49 |
| P3 | 71 | 22 | 49 |
| P4 | 71 | 22 | 49 |
| P5 | 70 | 21 | 49 |

The grew-by-five difference from the issue's figures is five intents, not a different method.
`evidence/lock-files.txt` carries the run.

**The negative result is a positive one somewhere.** `CLAUDE.md` asks for exactly this before a
finding is sealed: a tool asked about something absent answers as it does about something that
does not match. The same `grep -q '^files:'` that returns nothing for 122 P0 locks returns
something for 22 locks at each later phase, and the 49 without are the intents that ran before
`scope set` existed. So the key is there to be found and the P0 answer is an absence rather than
a silence — which is also visible in this intent's own five locks, one without and four with.

## The paragraphs, as they read in the file

All three were drafted, approved, written, and then read back as paragraphs rather than as
diffs — which found three sentences that needed changing and are recorded in the specification
commit's message. The fourth adjustment was mechanical: the budget paragraph's final sentence
ran to 94 characters after the rewording and was rewrapped.

A fifth would have been a placement mistake and was avoided by reading the neighbours first.
Section 5's lock subsection has two paragraphs before the new one — what the lock records, and
that it is not refreshed — and one after, about version numbers appearing in two places. The new
paragraph is a consequence of the second and would read as a qualification of the first if it came
earlier, and the one after begins a different subject, so nothing is split.

## Nothing changes behaviour

`go test ./...` passes with no test file in the diff against `main`. `gofmt -l .` prints nothing
outside `vendor/`, `go vet ./...` is clean, and `xeno gate verify` reports 475 verdicts verified
and exits 0. `evidence/go-test.txt` carries all four.

That is the claim and the whole of it: three paragraphs, three comments and one index paragraph,
and a reader of any of them learns something the runner already did.

<!-- xeno:section:gaps -->
## Gaps

**The intake is still not measured against its own budget.** That is the decision and not a gap in
carrying it out, and it is written down in three places now rather than nowhere. What remains
unmeasured is whether it matters: nothing records how often a P0's resolved scope would have
exceeded the budget declared beside it, and the figure is derivable from the trail — each P0's
`context-scope.yaml` against a resolution of its own patterns — but a resolution taken today
describes today's tree and not the one the intake was given. So the honest version of that
measurement does not exist, which is the same reason section 5's clause on `bytes` was written.

**The staleness check still does not run locally.** `staleReads` returns nothing without a commit
range, and section 12 forbids working one out, so the half of G-Freshness this intent wrote a
paragraph about is exercised in CI and not in the run that binds the sequence. #235 is the issue
about a check that did not run; XENO-0265 closed the half of it that had a shape, and the other
half is untouched here.

**Nothing checks that a comment's claim is true of the phase it is read at.** The learning
recorded at P3 proposes the convention — a comment explaining an absent value names which phases
the rule covers — and section 10 routes it through a merge request against the rule set rather
than taking effect on being noticed. Two instances in two packages were found by counting the
trail, not by reading either comment, and a third would be found the same way.

**`docs/clause-readers.md` still has a prose count beside a table nothing counts.** The
explanation count moves 126 to 129 by hand, which is the fault XENO-0260 and XENO-0264 both
recorded about that file. This intent adds one more figure of that kind rather than fixing it,
because the fix is a check in the verify job and belongs to the learning those two already filed.
