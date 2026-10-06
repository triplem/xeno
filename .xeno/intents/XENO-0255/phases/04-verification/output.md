---
intent: github.com/triplem/xeno#247
phase: 04-verification
created: "2026-10-05T19:46:27Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 30b5914677d7092092d9024f346bb0d3003f82bbfa25f335d13d69d10496fd39
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
      sha256: 1d942db0fbd377219af9b7bc99fd2fb884295eeb05b53f3fd842d5b0df1b54fa
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

No test is added and none could be. The change is a row in a document nothing reads, a paragraph in
the file every session reads and nothing checks, and a rule file nothing loads. A90 records the
first as deliberate; the second is read by a model rather than a program; the third is inert by
construction.

What stands in for it is reading, four counts and one command:

| criterion | what answers it |
|---|---|
| 1, the reason's row and the count | reading the row; the count line and a script counting rows between the two headings, both 39 |
| 2, the mapping row untouched | `git diff` over `docs/clause-readers.md`, which does not contain the mapping row, and does contain the new sentence |
| 3, the `CLAUDE.md` paragraph | reading it; `wc -l` at 87 against 78; and that it sits under Conventions with no heading added |
| 4, the example rule's shape | reading it against `documentation-follows-the-change.yaml`: `scope: org`, `binding: false`, `kind: review`, a header saying why it is an example and how to adopt it |
| 5, the rule names its ceiling | the header, which says once per intent rather than once per criterion, and that adopting it does not make the clause mechanically covered |
| 6, no normative document touched | `git diff --stat` |
| 7, nothing enabled, no code | `grep -rl mapping-is-complete .xeno/` finds nothing, and `rules.Load` walks `.xeno/plugin/rules/` and the project config, not `examples/` |
| 8, the trail | `./xeno gate verify` at exit 0, with `rules_hash` unchanged in every verdict |
| 9, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 10, prose within 88 | an `awk` pass over both documents outside tables and code blocks |
| 11, one commit | the commit itself |

**Criterion 8 needed a command rather than an argument.** "A file under `examples/` cannot reach the
effective set" is a claim about `rules.Load`'s walk; the way to test it is that `rules_hash` is
recorded in every artifact and `gate verify` recomputes it, so an example that leaked would change
the hash and every verdict with it.

Criterion 2 is a prohibition, and the diff is how it is read: the mapping row not appearing is
stronger evidence than reading it and finding it unchanged.

**This phase was started over, and the reason is a defect worth the space.** Its first attempt
declared `test-report/go-test` while a background `go test` was still writing the file. The
declaration bound the hash of a partial run; the suite then finished, the file grew, and G-Evidence
reported "evidence test-report/go-test content does not match its hash — the file changed after it
was declared". `gate verify` carried it as `DIVERGENT XENO-0255 04-verification`. `DeclareEvidence`
refuses a second declaration of the same kind and job, so the honest route was removing this phase
and P5 and starting both over, which is the escape #225 made explicit a few intents ago. This time
the suite ran to completion first and the file was stable before anything hashed it.

The suite is run and proves nothing about the three changes, which this phase says rather than
letting a green run stand as evidence. It would pass identically if the row were wrong, the
paragraph absent and the rule malformed.

<!-- xeno:section:results -->
## Results

Ten criteria met, one pending until the commit.

1. Met. The row reads "the recommendation carries a reason | **no field to carry it**; nothing reads
   it (#247)", the count line says 39, and the table counts 39.

2. Met, and it is a prohibition. `git diff` over `docs/clause-readers.md` does not contain the
   mapping row; the new sentence naming the example and saying why the reader column still says
   `nothing` is present.

3. Met. One paragraph under Conventions, no heading added, `CLAUDE.md` at 87 lines against 78.

4. Met. `scope: org`, `binding: false`, `kind: review`,
   `applies_to: [04-verification, 05-review]`, with a header in the shape the four beside it have.

5. Met. The header says it is answered once per intent rather than once per criterion, and that
   adopting it does not make the clause mechanically covered.

6. Met. `git diff --stat` names `CLAUDE.md` and `docs/clause-readers.md`; one file is added under
   `examples/rules/`. No normative document appears.

7. Met. `grep -rl mapping-is-complete .xeno/` finds nothing and no Go file changed.

8. Met. `./xeno gate verify` exits 0 with the verdicts that existed before this intent intact and
   `rules_hash` unchanged in every one, which is the proof that the example did not reach the
   effective set.

9. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
   nothing, and `go test ./...` is 18 packages `ok` with 0 failures — the run this phase declares,
   complete and stable before it was hashed.

10. Met. An `awk` pass over both documents outside tables and code blocks finds nothing over 88.

11. Pending. The commit comes after this phase is judged, and its footer takes a keyword each.

**This phase and P5 were started over, and the record of why belongs here rather than only in a
learning.** The first attempt declared the test report while a background `go test` was still
writing the file, so the declaration bound the hash of a partial run. The suite finished, the file
grew, and G-Evidence reported that the content does not match its hash with the exact next step —
"the file changed after it was declared". `gate verify` surfaced it as
`DIVERGENT XENO-0255 04-verification`.

Nothing recovered that cheaply. `DeclareEvidence` refuses a second declaration of the same kind and
job, with a reason: one of two declarations would stay pending with nothing saying why. Making the
file match the declared hash would mean reconstructing a truncated test run, which is not something
to do to evidence. Approving the finding would be recording that the evidence is untrustworthy and
proceeding, when the cause was a mistake with a trivial fix. So both phases were removed and redone
— the escape #225 made explicit — and this time the suite ran to completion and its hash was taken
from a file nothing was still writing.

The gate caught it immediately and said what had happened in one sentence. That is the half of this
worth keeping.

<!-- xeno:section:gaps -->
## Gaps

Two of the three changes are read by nobody who is not already looking. The audit row sits in a
document nothing reads, which is A90's own subject; the example rule sits where nothing loads it by
design. Only the `CLAUDE.md` paragraph has a mechanism behind it — the file is sent with every
request — which makes it the one change here that will alter behaviour and the one nothing can
verify did.

Nothing checks the convention, and the paragraph says so. A session that batches three questions
breaks it with no finding, no refusal and no verdict; the only reader is the model reading the file
before it asks.

The example rule has never been loaded. Nothing under `examples/` reaches `rules.Load`, which is
what makes it safe and also means a malformed field would sit unreported until a project copies it
into `.xeno/config/rules/`. The four rules beside it have been in that position since A72, and a
project adopting it is the first thing that would validate it.

Whether a person can usefully answer it once per intent is unknown. An intent with eleven criteria
compresses a real reading task into one checklist entry, and it may turn out the honest answer is
always `deviation` with a note, which would make it a worse instrument than no rule. Only adoption
would show that.

The three mechanisms the issues asked for are absent and #258 is the only thing tracking them. A
reader finding #247 closed will not know a field is still wanted unless they follow the link, which
is a document pointing at a document.

**Nothing stops the evidence mistake this phase made from recurring.** `evidence declare --file`
hashes whatever is on disk at that moment and cannot know a writer is still running; the gate
catches it afterwards, which is correct and is two commands late. A declaration of a file being
written is indistinguishable, at declare time, from a declaration of a finished one. Whether that
wants a guard — and what a guard could even look for — is not this intent's to decide, and the
learning records the method that avoids it instead.
