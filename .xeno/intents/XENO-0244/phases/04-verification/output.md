---
intent: github.com/triplem/xeno#221
phase: 04-verification
created: "2026-10-04T09:39:52Z"
schema_version: "1.0"
runner_version: dev+24becc3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 312534f21640d4a3d373eb4d7dd4f6cbc12f49e126442156f4dadb57823993b8
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - format: go-test-json
      job: test
      kind: test-report
      produced_by: go test -json ./...
    - format: other
      job: semgrep
      kind: scan
      produced_by: semgrep --config .semgrep/golang.yaml --error --metrics off --json -o semgrep.json cmd internal
    - format: other
      job: trivy
      kind: scan
      produced_by: trivy rootfs --trivy-config trivy.yaml --format json xeno
    - format: other
      job: npm-audit
      kind: scan
      produced_by: npm audit --json, over what the release installs
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it | Fails without the change |
|---|---|---|
| The fifteen read `pass`, and nothing else about them changes | the diff: fifteen files, fifteen insertions, fifteen deletions, each a `result:` line; `pipeline: local` and every `sha256` untouched | yes |
| No hash moves and no verdict is contradicted | `xeno gate verify` at exit 0 over 342 verdicts, taken before the change, after the correction and after the gate | yes, and this is the figure the order of the work exists for |
| An attachment carrying a result outside the set is a finding | `TestAnAttachedResultOutsideTheSetIsAFinding`: G-Evidence `fail`, the cause naming the value, the finding on `evidence/attached.yaml` | yes |
| Every documented value passes | `TestEveryDocumentedResultIsAcceptedOnAnAttachment`, looping over `testdata/evidence-declaration.yaml` | yes |
| An absent result on an attachment is not a finding | `TestAnAttachmentWithNoResultIsNotAFinding`, which is `other/trivy-db`'s case | yes, a check that demanded one would fail a conformant pipeline |
| `evidence attach` declines an entry whose result is outside the set | `TestAResultOutsideTheSetIsDeclined`: nothing attached, one pending, one sentence | yes |
| The declined entry leaves nothing behind | the same test: no record in `attached.yaml`, no report copied into `evidence/` | yes |
| The reason names the entry, the value and what to republish | the same test, on `scan/trivy`, `success` and `Republish` | yes |
| Every documented value is recorded as published, `fail` included | `TestEveryDocumentedResultIsRecorded` | yes |
| An entry whose producer reports nothing is still recorded | `TestAnEntryWithNoResultIsStillRecorded` | yes |
| The refusal runs against a real manifest | the transcript in the results section, against this phase | yes |
| The fifteen phases stay green throughout | `gate verify` at exit 0 with the verdict count unchanged, and this intent's own `intent verify` | yes |
| The exit code staircase is unchanged | `evidence attach` at 1 where something was declined, 0 where everything bound, 2 where it could not run; all three in the transcript | yes |
| Nothing else changes | no new gate, rule, template, document or declaration; `gate verify` at 342 verdicts in all three readings | yes |

<!-- xeno:section:results -->
## Results

**The figure the order of the work exists for, taken three times.** `xeno gate verify`
at exit 0 over **342 verdicts** in all three readings: before the change, after the
fifteen corrections, and after the gate was added. The verdict count is identical, no
divergence is reported, and every one of the fifteen phases keeps its `artifacts_hash`.
That is the reading of Appendix B confirmed rather than assumed: `artifacts_hash` is
computed over the files lying directly in a phase directory, so `evidence/attached.yaml`
is outside it. Had that been wrong, the middle reading would have reported fifteen
divergences.

Timings beside it, on an idle machine: 1 322 and 1 306 ms before, 1 222 and 1 336 ms
after the correction, 1 296 and 1 280 ms after the gate. The spread between readings of
the same tree is wider than the difference between trees, which is the state XENO-0242
and XENO-0243 both recorded for this figure.

**The correction.** Fifteen files, fifteen insertions, fifteen deletions, each a
`result:` line: XENO-0107, XENO-0108, XENO-0111, XENO-0121 and XENO-0200 to XENO-0210.
All fifteen were read before the edit and are identical in shape, one entry each with
exactly one `result:` line, which is what made a whole line anchor safe. `pipeline:
local` and every `sha256` are untouched, and G-Evidence resolving each one afterwards is
what proves the reports still match their hashes.

**The suite, the build and the two silent tools.** Eighteen packages ok, two with no
test files, `go build -o xeno ./cmd/xeno` clean, `gofmt -l .` outside `vendor/` empty,
`go vet ./...` silent. Six new tests, three in `internal/evidence/attach_test.go` and
three in `internal/gates/evidence_test.go`. Code: +50 / -14 in `attach.go`, +26 in
`gates.go`, +7 / -4 in `main.go`, +6 / -4 in `runner.go`. Tests: +87 / -8 and +44, plus
the one changed substring in `runner_test.go`.

**The refusal was run against this phase's own declarations**, with a hand written
manifest standing in for a pipeline that published one entry wrong and one unbindable:

```
$ xeno evidence attach --intent XENO-0244 --phase 04 --from DIR
attached 0, pending 4; run xeno gate run to carry the verdict forward
  not attached: scan/semgrep was published with result success, which section 4 does not
  define; the set is pass and fail, and the value is what the run reported against its own
  threshold. Republish it with one of the two
  not attached: scan/trivy was published with a uri and no sha256, and a uri is bound by
  its hash alone, because resolving one needs a network the gate path never has; republish
  it with the sha256
  exit 1
```

Both reasons carry their own remedy, which is what the rename of the field was for, and
neither entry was recorded: the phase directory held `context.lock.yaml` and `output.md`
and no `evidence/` afterwards. The four declarations stayed pending, so the real
pipeline can still bind them.

**The four declarations of this phase were written by the command** —
`test-report/test`, `scan/semgrep`, `scan/trivy`, `scan/npm-audit` — which makes this
the second intent to declare evidence and the first to do it as a habit rather than as
the thing being built. The phase finishes `provisional` and the pipeline is what
resolves it.

**What this phase cannot assert about itself.** That the real manifests bind through the
new check: the four entries CI publishes all carry `pass`, so the path the refusal
guards is the one that must stay open, and the review phase records what arrived. The
declined case is asserted by test and by the transcript above; the accepted case is
asserted by test and then by this intent's own attachment.

<!-- xeno:section:gaps -->
## Gaps

**An attachment with no result at all, on a kind that needs one, is still unjudged.**
Section 4 requires `result` on `test-report` and `build-log` because G-Test and G-Build
read it, and this intent judges a value while leaving an absence alone — deliberately,
because the absence is legitimate for every other kind and no attachment in the trail
lacks one. The asymmetry is therefore in the right place on the declaration side and
incomplete on the attachment side. Filed as its own issue, as D-1's scope requires.

**G-Test is still `not-implemented`.** It is the reader that would have caught the
fifteen, and it stays absent. The result of a `test-report` attachment is now a word the
document defines, which is exactly the groundwork G-Test needs and is as far as this
intent goes.

**The fifteen still point at transcripts somebody typed.** Bound by hashes that are
correct about bytes a person wrote, with `pipeline: local` recording a regime this
repository does not configure. The vocabulary is corrected and the provenance is not,
because the provenance is accurate.

**Nothing re-reads a manifest the pipeline has already published.** The refusal keeps a
bad entry out at the point of writing, and an entry that was recorded before this change
is caught by the gate. What is not covered is a manifest published wrong and then
corrected in the artifact store: the declaration stays pending and somebody has to run
the attach again, which is the same manual step as every other part of the pull path
(#225 names the fetch adapter's absence).

**The three callers of `Attach` print the declined list in three voices, and only two
were corrected.** `gate run` calls it as well, through the same path, and says nothing
about a declined entry: `cmdGateRun` reports the verdict and the pending count. The
verdict is correct — the phase is provisional — but the reason is not printed there as
it is in `phase start` and `evidence attach`. Named rather than fixed, because the third
voice is a line of output and D-1's scope is the gate and the refusal.

**The correction was a hand edit.** Fifteen files, one word, read before and after, with
the figures to show nothing else moved. There is no command for it and this intent
argued there should not be. What that leaves is a precedent: the next vocabulary
correction inside `evidence/` will be done the same way, by hand, and the only thing
standing between it and a mistake is the three readings of `gate verify`.
