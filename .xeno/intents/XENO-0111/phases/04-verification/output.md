---
intent: github.com/triplem/xeno#111
phase: 04-verification
created: "2026-09-28T17:38:14Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e77cc2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: dc6f05ec8ea3d3ae2f64cd4bb703a8e57229d44e197456e533f237bdff495f63
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
evidence:
  - kind: test-report
    job: go-test
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it |
|---|---|
| AC1 the comment documents `printInit` | read, and `gofmt` enforces the blank line that separates it from `printEnforcement`'s. `go doc` is not run in CI, so this one is checked by a reader |
| AC2 a pending declaration declares no file | `TestAPendingDeclarationDeclaresNoFile`: the undeclared file is reported and nothing is keyed on a basename of nothing |
| AC3 a nested declaration names one file | `TestADeclarationUnderASubdirectoryDeclaresThatFileAlone`: `logs/a.txt` declared, `a.txt` at the top still reported |
| AC4 same basename, different directories | the same test, from the other side: the top level file is not taken as declared by the nested declaration |
| AC5 no finding for a directory of declared files | `TestADirectoryOfDeclaredFilesIsNoFinding` |
| AC6 everything green stays green | `go test ./...` on every package, and `./xeno gate verify` over 61 verdicts, which includes every phase of every intent this repository carries |
| AC7 a path outside `evidence/` declares nothing inside | `TestAPathOutsideEvidenceDeclaresNothingInside` |

AC1 is the one no gate reaches. It is proved by reading the file, and the blank line that made
the defect possible is now the thing `gofmt` would remove the fix by restoring.

<!-- xeno:section:results -->
## Results

`go test ./...` passes on every package. `gofmt -l .` outside `vendor/` prints nothing,
`go vet ./...` is silent, and `./xeno gate verify` recomputes 61 verdicts and matches.

Three of the four new tests fail against the tree with `internal/gates/gates.go`
reverted, and the transcript carries that run: a nested declaration marking a top level
file as declared, a directory of declared files reported, and a path outside `evidence/`
taken to declare one inside it.

The fourth passes there. The entry keyed `.` was inert, exactly as the issue said, so
the test guards against its return rather than proving its removal. Reporting it as a
caught defect would have been the easy mistake of this phase.

The comment has no test. `go doc` renders it under `printInit` and nothing in CI renders
godoc, so AC1 rests on a reading of the file.

<!-- xeno:section:gaps -->
## Gaps

**AC1 is verified by a reader.** Nothing in the pipeline renders godoc, so a comment
attached to the wrong function is exactly as invisible after this change as before it.
What has changed is one comment; what produced it has not. A check that parses the file
and reports a doc comment separated from its function by nothing would close the class,
and it is not written.

**The evidence is a local run again.** `.github/workflows/xeno.yml` uploads no artifact,
so the declared test report is a transcript produced on this machine and attached
through the stand in directory. `attached.yaml` records a hash and a state and nothing
about what produced the bytes. Same finding as XENO-0111's predecessor, unchanged, and
it belongs to whichever package owns evidence publication.

**Nothing exercises a nested `evidence/` outside the tests.** `Attach` writes one flat
level, so the case the fix is for still has no producer. The tests are the only place
the nesting exists, which is what makes them the whole value of the change and also
means the fix is unproven against a real nested layout.

**The walk is not measured.** One `Stat` per file per gate run on directories holding a
handful of files. Stated in P2, still not measured, and a phase with a large `evidence/`
would be the first to notice.
