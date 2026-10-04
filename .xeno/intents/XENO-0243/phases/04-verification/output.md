---
intent: github.com/triplem/xeno#208
phase: 04-verification
created: "2026-10-04T08:39:07Z"
schema_version: "1.0"
runner_version: dev+49f2794.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9dc3da37b2bd1d6813ea13f4640f2bf79bba567af4fd7a3023c8ddf4e49e2764
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
| An item that exists is declared from its file and the runner computes the hash | `TestAnItemThatExistsIsDeclaredFromItsFileAndTheRunnerHashesIt`: the declared hash equals `hashing.Hex(hashing.Normalise(content))`, the copy is in `evidence/`, the item is in the frontmatter | yes, the command does not exist |
| The declared item is judged, and G-Evidence fails for a real reason | `TestADeclaredItemIsJudgedAndFailsWhenItsContentChanges`: green with the report as declared, `does not match its hash` after an edit, `does not resolve` after a removal | yes |
| An item in a store is declared with the hash the pipeline published | `TestAnItemInAStoreIsDeclaredWithTheHashThePipelinePublished`, which also asserts that no `evidence/` directory is created for it | yes |
| A pending item declares its kind and its job and is what the attach binds to | `TestAPendingItemDeclaresItsKindAndJobAndIsWhatTheAttachBindsTo`: pending, the phase provisional, then green after `phase start` of P5 attaches | yes |
| Provenance is written where given and absent where not | `TestProvenanceIsWrittenWhereGivenAndAbsentWhereNot` | yes, the two fields are not in the model |
| Every refusal refuses before anything is written | `TestADeclarationIsRefusedBeforeItIsWritten`: twelve rows, then the assertions that the phase holds no declaration and no `evidence/` | yes |
| The refusals are the gate's own judgement | `TestTheRefusalIsTheGatesOwnJudgement`, which calls `gates.EvidenceShape` on an item the command declines | yes |
| A pair is declared once, and two jobs of one kind are two items | `TestAPairIsDeclaredOnce` | yes |
| A second report does not take the first one's name | `TestDeclaringOverAnotherReportIsRefused`, which reads the first report back afterwards | yes |
| A declaration into a phase with no artifact is refused | `TestDeclaringIntoAPhaseWithNoArtifactIsRefused`, on the sentence naming `section set` | yes |
| A declaration after a verdict leaves it stale rather than rewritten | `TestADeclarationAfterAVerdictLeavesItStale`: the recorded hash unchanged, the artifact's own hash changed | yes |
| The body and the other frontmatter fields are untouched | `TestADeclarationTouchesNothingElse` | yes |
| The command line form, and all three steps of the staircase | `TestEvidenceIsDeclaredFromTheCommandLine`: 0 for both forms with the line naming the state and the hash, 1 for the pair declared twice, 2 for a phase that does not resolve | yes |
| `format` is judged against section 4's set | `TestADeclarationIsRefusedBeforeItIsWritten`'s format row, and `gate verify` at 336 verdicts unchanged | yes, the set did not exist |
| This intent's P4 carries evidence through section 6's pull path | this phase: four pending declarations written by the command, `provisional` on finish, attached at P5's start from the pipeline's own artifacts | yes |

<!-- xeno:section:results -->
## Results

**The suite, the build and the two silent tools.** Eighteen packages ok, two with no
test files, `go build -o xeno ./cmd/xeno` clean, `gofmt -l .` outside `vendor/` empty,
`go vet ./...` silent. Thirteen new tests: twelve in `internal/runner/evidence_test.go`
and one in `cmd/xeno/main_test.go`.

**`xeno gate verify` is at exit 0 over 336 verdicts**, which is the figure that matters
for the one change made to a gate. `format` joined the closed sets G-Schema compares
against, and no declaration in the trail carries a format, so nothing already committed
became a finding. Measured at 1.1 seconds for the whole trail.

**One correction carried in the open.** The implementation phase said fourteen tests in
the runner; there are twelve, and one in the entry point. The count was written while
the file was still being added to and it was never recomputed. P3 is sealed, so the
number stays wrong there and is corrected here, which is what this repository did with
the design phase's claim about the Go tree in #211. The twelve rows of the refusal table
are the count that was checked, because each is a path through `declarable`, `artifact`,
the pair check, `copyEvidence` or the shape check.

**The refusals were also run against this repository**, because the tests prove the
behaviour and a transcript proves the wording. Five of the twelve, against this phase
while it was being written:

```
$ xeno evidence declare --intent XENO-0243 --phase 04 --kind tests --job unit
refused: evidence item tests/unit has a kind section 4 does not define; section 4 fixes
the set: test-report, coverage, build-log, scan, sbom, other
  exit 1
$ xeno evidence declare --intent XENO-0243 --phase 04 --kind scan --job trivy
refused: 04-verification already declares scan/trivy: an item is identified by its kind
and its job, which is the pair an attachment is bound to, so one of two declarations of
it would stay pending with nothing saying why
  exit 1
$ xeno evidence declare --intent XENO-0243 --phase 04 --kind scan --job unit --result pass
refused: a pending item carries no --result: the run that would report one has not
happened, and the value arrives in evidence/attached.yaml as the job's own verdict
  exit 1
$ xeno evidence declare --intent XENO-0243 --phase 04 --kind test-report --job coverage \
    --uri https://example/x
refused: --uri needs --sha256: a uri is bound by its hash alone, because resolving one
needs a network the gate path never has
  exit 1
$ xeno evidence declare --intent XENO-0243 --phase 04 --kind test-report --job coverage \
    --sha256 aaaa...aaaa
refused: --sha256 names a hash and nothing it belongs to: give it with --uri, or declare
--file, which computes it
  exit 1
```

Each exits 1 and none of them wrote: the four declarations this phase carries are the
four that were asked for, and the phase has no `evidence/` directory until the attach
makes one.

**The four declarations of this phase were written by the command**, which is the first
time in this repository that an `evidence` block was not typed into a frontmatter. They
are `test-report/test`, `scan/semgrep`, `scan/trivy` and `scan/npm-audit`, each with the
command its job runs in `produced_by` and its shape in `format`, and each pending,
because the run that reports a result has not happened when the phase finishes. That
makes this phase's verdict `provisional`, which is a state of its own and not a
disagreement: `gate verify` reports it and exits 0, so the push is not blocked, and the
attachment is the next command's business.

**What the pipeline then does, and what is measured only there.** The branch is pushed
with an open pull request, which is what `pull_request` needs to trigger the three scan
workflows at all, and `xeno.yml` publishes the test report for the first time. The
attachment of all four, P4's verdict going from provisional to green with its
`artifacts_hash` unchanged, and the content of each report landing in `evidence/` are
the acceptance criterion this phase cannot assert about itself: it happens at P5's
`phase start --evidence-from`, and the review phase records what arrived. The mechanism
is asserted by `TestAPendingItemDeclaresItsKindAndJobAndIsWhatTheAttachBindsTo` against
a fixture pipeline; what the pull path has never done before is run against a real one.

**Figures.** 149 lines of runner code, 303 of tests, 49 in the entry point, 72 of entry
point tests, 23 in the model, 14 in the gates, and 79 in one workflow. The published
test report is 11 452 bytes against 495 629 for the full stream, measured on this tree
at 2 293 events over twenty packages.

<!-- xeno:section:gaps -->
## Gaps

**An attachment's `result` is still judged by nothing.** The fifteen attachments already
in the trail read `result: success`, which section 4's closed set does not contain, and
the shape check reads the frontmatter while `attached.yaml` is another file. This intent
found it and deliberately did not fix it: a check would recompute fifteen sealed phases
into a status that disagrees with the committed one, and `Verify` reports that as a
divergence, so CI would be red on history. Filed as an issue of its own, with the
fifteen listed.

**Nothing makes a verification phase declare anything.** The command exists and the
habit is what has to change, which is the same residual risk #188 recorded for the
exchange. A gate that demanded a declaration would be answered with a declaration of
something nobody ran, which is the failure the whole of section 4 is arranged to
prevent.

**`produced_by` and `format` have no reader.** Both are provenance and that is the state
section 4 describes, but it means nothing would notice a `produced_by` that names a
command nobody ran. A-002 says what the field means on a pending item and the register
carries it; no check does.

**The published report is the package level stream.** A reader who wants the output of
one failing test needs the pipeline run while it is still in retention. The deviation
names the figures; what it costs is named here rather than in a sentence about size.

**The `--uri` form has not been run against a real artifact store.** It is asserted by a
test and nothing in this repository publishes a uri item, because the three scan
workflows and the new test step all publish files. The form exists because section 4
defines it, and what would exercise it is a report too large to copy, which the four
declared here are not.

**The attach's fetch adapter is still a directory.** Downloading the pipeline's
artifacts is done by hand at P5, with `gh run download`, and the directory that comes
out of it is what `--evidence-from` reads. Section 4 describes an artifact store and WP6
describes an adapter; what exists is the offline stand-in, which is enough for the pull
path to run and is not the same thing as the adapter.

**A phase that declares evidence needs an open pull request to finish.** The three scan
workflows trigger on `pull_request`, so the round trip section 4 names cannot even start
from a pushed branch alone. That is a property of this repository's workflows rather
than of the tool, and it is written down because an intent that declared evidence and
pushed without opening the request would wait for a pipeline that never runs.

**Section 4's `id` is still unwritten**, by D-2, and so is the `sbom` kind, which
nothing in this repository produces. Both are in the section and neither is a gap this
intent left by accident.
