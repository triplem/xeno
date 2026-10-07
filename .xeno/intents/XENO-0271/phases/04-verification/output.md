---
intent: github.com/triplem/xeno#273
phase: 04-verification
created: "2026-10-07T09:58:04Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8d718d53a2af982af8c9339a24a02a51803c3c049faa0cc1d3b12b5881f3e8bb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - kind: test-report
      result: pass
      produced_by: go test ./...
      sha256: f854852e8bf0a0bc66a5d2e3369f3539822a58a1a2c1ce6458d27217b3dbdf89
      path: evidence/go-test.txt
      job: test
    - kind: build-log
      result: pass
      produced_by: go build, gofmt -l ., go vet ./..., xeno gate verify, git status over the untouched trees
      sha256: 65eb575fcd66d246fbfcecddf06a9b8177ec163e65d16b867e164ebfa2a9e7fa
      path: evidence/checks.txt
      job: checks
    - kind: other
      result: pass
      produced_by: both rename variants on throwaway copies, with the control and the false negative that preceded them
      sha256: eece881c5dce5456121703a69f3418c378d3cf1c2e335beda1afa97433b3467e
      path: evidence/variants.txt
      job: variants
    - kind: other
      result: pass
      produced_by: A33's five cells before and after, and the grep for A33 across the Go files
      sha256: 69f7d8fb4f5a287ad218edc4cfe39ea32a8465f39cd7e3af52b80200d9f1de9a
      path: evidence/row.txt
      job: row
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Ten criteria, by number, all passing. Six are read against the row, two are mechanical checks
over it, and two are the absences.

| # | What it asserts | What proves it |
|---|---|---|
| 1 | A33's assumption and reason are unchanged | `evidence/row.txt`: the five cells of the row before and after, compared cell by cell. Four are byte-identical and only the state changed |
| 2 | the state says what the first variant would cost | read in `evidence/row.txt`: `Load` keying on the phase, section 5's sentence becoming false, the override path moving |
| 3 | and what the second would cost | read: `goneBundle` for every sealed artifact and the recompute retiring |
| 4 | the figure is in it | read: 495 when the variants were measured and 499 when the row was written, both dated, with "growing with each intent" |
| 5 | it says `gate verify` notices neither | read: its own sentence, in bold, saying every one of them was reported verified under either form |
| 6 | it names the issue | read: `#273`, twice — once for the measurement and once for where the variants stay written out |
| 7 | nothing is renamed | `evidence/checks.txt`: `git status` over `internal/`, `cmd/` and `.xeno/plugin/` is empty, and the only modified file is `docs/assumptions.md` |
| 8 | no normative document changes | the same, and `docs/process-definition.md` is not in the status |
| 9 | the row is one line with five cells and keeps `approved` | `evidence/row.txt`: five fields between pipes, and the state begins `approved;` |
| 10 | the suite passes and no verdict moves | `evidence/go-test.txt`, `evidence/checks.txt`: build, `gofmt`, `go vet` clean and `xeno gate verify` matching 499 verdicts |

The measurement the row reports is itself in `evidence/variants.txt`: the control, both
variants, the code that explains the second, the count of artifacts carrying a template ref,
and the first attempt that produced a false negative.

<!-- xeno:section:results -->
## Results

**One cell changed, proved rather than claimed.** A33's five cells were compared before and
after by splitting the row on its pipes: the number, the assumption, the reason and the where
are byte-identical and only the state differs. That is the whole claim of this intent and it is
the one thing a diff of a 94-kilobyte file does not make obvious.

**The measurement the row reports, with its control.** `evidence/variants.txt` carries all of
it. On the original layout, a `strings_hash` corrupted to a hex value turns G-Schema red and
names the bundle it no longer matches. With the directories renamed and `TemplateID` made the
identity function but `id: intake` left alone, the same corruption is still caught and
`gate verify` matches 495 verdicts. With the `id:` renamed too, the same corruption goes green
and `gate verify` still matches 495. The code that explains the third result is printed beside
it: `hashes` sets `goneBundle` from `ref != t.Ref()`, and `recomputed` compares `strings_hash`
only when `!goneBundle`.

**The first attempt was a false negative and is recorded as one.** `strings_hash` was first set
to 64 zeros and the gate passed, which looked like the check being dead. It passed on the
untouched tree too, and `context_hash` set to 64 ones did the same. YAML parses an all-digit
scalar as a number, so `raw[f].(string)` fails and `hashes` skips the field. The measurement
was redone with a hex value containing letters, which is what produced the three results above.
Both the broken control and the reason are in the evidence, because a reader who repeats this
will reach for the same obvious value.

**The figure, and why the row carries two of them.** 499 artifacts carry a template ref today,
counted by listing the files that hold the field. It was 495 when the variants were measured
earlier the same day, and the difference is this intent's own phases. The row says both numbers
with the date rather than one, because a single figure would have been stale before the commit
that introduced it.

**Nothing outside the register moved.** `git status` over `internal/`, `cmd/` and
`.xeno/plugin/` is empty, and the only modified file is `docs/assumptions.md`. No normative
document, no template, no Go file.

**The reachability finding, checked.** A grep for A33 across the Go files returns one line, in
`phaseNumber`'s comment in `internal/runner/next.go`, about the `--phase` prefix.
`TemplateID`'s own comment restates A33's reason in A33's words and does not name the row. That
is what P3's deviation records and it is why the drafted decision was wrong.

**The checks.** `go test ./...` passes across 20 packages. `go build`, `gofmt -l .` outside
`vendor/` and `go vet ./...` are clean. `xeno gate verify` recomputes and matches 499 verdicts,
which is every sealed phase in the repository.

<!-- xeno:section:gaps -->
## Gaps

**A row is not a check.** Nothing stops somebody renaming the template directories tomorrow,
and the only thing that would have told them what it costs is a register row they have to
think to open. P1 ruled out the guard that would catch it — a check that each `template.yaml`'s
`id:` matches its directory name — because it would also forbid the harmless variant, so the
guard that exists is a paragraph.

**The destructive variant stays silent, and that is unchanged.** `goneBundle` cannot tell a
bundle that was renamed from one that is gone. If anybody does rename the ids, every
`strings_hash` in the trail stops being compared and the only command that recomputes the whole
trail reports every verdict verified. The row says so; nothing in the code does.

**The all-digit hash scalar is still skipped.** A `strings_hash` of 64 digits is parsed as a
number, so neither the shape check nor the recompute sees it. It is a one-in-ten-trillion
accident for a real sha256 and a reliable trap for anybody fabricating a value to test the gate,
which is what happened here. Written on the issue and in P1's non-goals, fixed nowhere.

**A33 is reachable by searching, not by reading.** `TemplateID` implements the row and does not
cite it. The finding is recorded in P3 and the fix was declined to keep a Go file out of a diff
whose claim is that nothing in the runner moves, which is a reason about this intent rather
than about the comment.

**The figures are from one machine and one day.** 495 and 499 are counts of files in this tree
on 2026-10-07. The claim they support — that the recompute retires across all of them — is a
statement about the mechanism and not about the number, and the number is there so a reader can
see the order of magnitude rather than recompute it.

**Variant A was measured for safety, not for correctness.** What was checked is that it leaves
the trail's checks alive. That it contradicts section 5 was read off the sentence, not
demonstrated: nobody ran a project override under the renamed layout to watch resolution pick
the wrong file. The reading is plain and the measurement is absent.

**Six of ten criteria are a person reading one table cell.** The cell is long and it is prose.
Whether it reads as an argument or as a wall is a judgement, and the only reader who has made
it wrote it.
