---
intent: github.com/triplem/xeno#256
phase: 04-verification
created: "2026-10-06T08:39:27Z"
schema_version: "1.0"
runner_version: dev+5163d1b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 12177efa8bf7c7fa3a5d22eb11cd70e5a20c3fc292ff1136c3313c3cffe34da1
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
      sha256: 9f5ceacbfa9aa00f69072f978577b72d54c5d1aee61170b4ba906d6145809ad9
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

No test is added and none could be. The change is one paragraph in a file no gate, rule or Go file
reads — it is prose sent to a model, which is the only mechanism a convention of this kind has.

What stands in for it is reading, three counts and a re-verification:

| criterion | what answers it |
|---|---|
| 1, one paragraph under Conventions, no heading | `git diff`: ten lines added, no `##` among them |
| 2, the rule as an act | reading it: "recreate the thing and run the check again before writing the finding down" |
| 3, the reason framed on sealing | the clause beginning "It matters more here than elsewhere", which is the argument for the addition being in this file |
| 4, the instance named and checkable | the three facts, each re-verified from the tree rather than recalled |
| 5, says nothing checks it | the closing "Nothing checks this (#256)", the form the sequencing paragraph uses |
| 6, no heading, no list, still short | 97 lines against 87; ten added for one convention |
| 7, no row in `docs/clause-readers.md` | `git diff --stat` names one file |
| 8, no normative document | the same |
| 9, the trail | `./xeno gate verify` |
| 10, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 11, prose within 88 | an `awk` pass over the file |
| 12, one commit | the commit itself |

**Criterion 4 is the one with any content, and it was done the way the paragraph asks.** A sentence
about verifying a negative resting on a remembered negative would be its own counter-example, so all
three facts were taken from the tree while P3 ran: `grep -n` puts the entry at `.gitignore:10`;
`git log -S "internal/plugin/embedded/plugin/" -- .gitignore` puts it in `154bb09`, which is #200;
and `grep -rl "not gitignored" .xeno/intents/XENO-0249/` returns `04-verification/output.md` and
`05-review/output.md`, which is where the claim is sealed.

Criterion 6 is a judgement rather than a measurement and this phase reports the figure rather than
claiming the judgement was checked. Ten lines on 87, for a convention earned by one instance. A
reader who thinks that is too much for a file asking to stay short is making the same judgement with
the same number.

The suite is run and proves nothing about this change, which this phase says rather than letting a
green run stand as evidence. No Go code reads `CLAUDE.md`.

**The evidence was declared from a settled file**: the suite ran to completion, the process was
confirmed gone, and the hash was read twice and compared before anything declared it — the method
XENO-0255's mistake taught.

<!-- xeno:section:results -->
## Results

Eleven criteria met, one pending until the commit.

1. Met. One paragraph under Conventions, ten lines, no heading added.

2. Met. "So recreate the thing and run the check again before writing the finding down."

3. Met. "It matters more here than elsewhere: a finding goes into an artifact and is sealed with it,
   so a wrong one is permanent rather than corrected, with the correction somewhere a reader of that
   phase will not be."

4. Met, and re-verified from the tree. `.gitignore:10` carries the entry; `git log -S` puts it in
   `154bb09`, which is #200; the false claim is in XENO-0249's `04-verification/output.md` and
   `05-review/output.md`.

5. Met. "Nothing checks this (#256)."

6. Met. 97 lines against 87. Ten for one convention, which is a judgement reported as a figure.

7. Met. No row in `docs/clause-readers.md`, for the third intent running on the same reasoning.

8. Met. No normative document; `git diff --stat` names `CLAUDE.md` alone.

9. Met. `./xeno gate verify` exits 0 with the 423 verdicts that existed before this intent intact and
   the rest its own judged phases.

10. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints nothing,
    and `go test ./...` is 18 packages `ok` with 0 failures — declared from a file whose writer had
    exited and whose hash was read twice and compared.

11. Met. Nothing in `CLAUDE.md` exceeds 88 columns, after a reflow that P3's deviations records.

12. Pending. The commit comes after this phase is judged.

**What the issue asked for and which branch this took.** #256's done-when offered two: the convention
written down somewhere an agent reads before investigating, or the decision not to write it down
recorded with the reason. This takes the first, and the reason it is defensible is the framing rather
than the content — the generic rule would have been the first line in `CLAUDE.md` about nothing
specific to this project, and the sealing consequence is what makes it belong. P2's alternatives
carries the case for the other branch, which was a real one: a learning in `learning.yaml` is at
least honest about having no reader.

<!-- xeno:section:gaps -->
## Gaps

Nothing checks it, which the paragraph says and which is the whole of its enforcement. A session that
writes an unverified negative into an artifact breaks the convention with no finding, no refusal and
no verdict; the only reader is a model reading four sentences before it investigates. Whether the next
false negative is caught is not measurable from here and no verdict will record either outcome.

`CLAUDE.md` now carries two conventions whose reader is a model and whose compliance is invisible.
The sequencing one arrived in #259 and this is the second, both from this session's own mistakes. That
is not a defect in either, and it is a direction: a file sent with every request, asking to stay
short, accumulating rules nothing can check. Its last line is the only thing arguing against the next
addition and nothing enforces that either.

The instance will age. #201 is a closed issue and XENO-0249's phases are sealed, so the example stays
accurate — but if `internal/plugin/embedded/plugin` ever stops being gitignored the paragraph will
read as though it were describing a live state rather than a corrected claim. It says "recorded ...
when the entry was at `.gitignore:10` all along", which is past tense and should survive, and nothing
will report it if it does not.

The second instance is in the issue and not in the file. #212's intake read the wrong field on
`test-report` and concluded the opposite of the truth, which is the same shape by a different
mechanism — a tool's "nothing" taken for "nothing anywhere". A reader of `CLAUDE.md` gets one example
and will generalise from it to paths rather than to fields, which is narrower than the class. The
issue has both; the paragraph names the class in one clause and leans on the path instance.

Nothing was done about the three learnings this convention joins. #229's, #212's and XENO-0254's sit
in their phases' `learning.yaml` with nothing reading them, which is what #256 observed about its own
origin. One of them became this paragraph; the other two did not, and section 10's route for them is
unchanged and unused.
