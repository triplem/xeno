---
intent: github.com/triplem/xeno#332
phase: 04-verification
created: "2026-10-09T13:41:10Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2dce87f210c6874687296b4411ced30e1dc1c769e9c4d7d6df01db7265a97ba7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| criterion | how it was checked | result |
|---|---|---|
| 1. the clause precedes the code | `git log`: `b431e42` first after `main`, revision 11, names both forms, the script and the init line | met |
| 2. constants say the new names and accept nothing else | `identity_test.go`: the new case with `approved` and `approved\n…` reports two missing; the `both` case with `Xeno-Approved` and `/Xeno approved.` reports none | met |
| 3. refusal and sentence print the new names | `TestAnIssueNobodyApprovedDoesNotBecomeAnIntent` asserts `the label xeno-approved` and `a comment whose first line is /xeno approved`; `TestTheIntakeSaysWhereNoApprovalWasFound` the sentence | met |
| 4. the script is shipped, executable, idempotent | `TestTheLabelScriptCreatesTheLabelTheClauseNames`; `sh -n` clean; the read-before-write is in the text and was not run against a host | met for presence and shape; idempotence by reading, see gaps |
| 5. init names the label | `TestInitNamesWhatItCannotDo` looks for the constant and the script's name in `Manual` | met |
| 6. every fixture and document moved | `grep -rn '"approved"' --include='*_test.go'` finds only gate decisions and the new regression case; `docs/commands.md` names both forms | met |
| 7. A107 | present | met |
| 8. suite and gates | `go test ./...` exit 0, `gofmt -l` empty, `go vet` clean, `gate verify` 558 | met |
| 9. the host relabelled at the merge | done at the merge and recorded on the pull request | after this phase |

<!-- xeno:section:results -->
## Results

Nine criteria; eight met on the branch and the ninth is the merge's.

    go test ./...                 ok, exit 0 (second run; the first found three fixture faults)
    gofmt -l, go vet              clean
    sh -n xeno-labels.sh          clean
    ./xeno gate verify            verified 558 verdicts

**The old pair is refused, measured.** The new case in `identity_test.go` gives an issue labelled `approved` with a comment beginning `approved` and gets two missing halves, which is what makes this a rename and not an addition.

**Timings, for #117.** `gate verify` 36 s over 558, unchanged per verdict.

<!-- xeno:section:gaps -->
## Gaps

**The script was not run against a host.** Its two branches were read, not executed; `gh label create` and `glab label create` have the flags written here on 2026-10-09, and the first run at the merge is the proof. The GitLab branch has no host to be proved on at all.

**Idempotence rests on a read.** `gh label list` and `glab label list` answer whether the label exists; a host that paginates past the label would create a duplicate, which both hosts refuse anyway.

**Issues approved under the old word need a new comment.** Nine open issues carry `approved` comments of the old form; relabelling moves the label and not the comments, and each needs `/xeno approved` before its intent can start.

**The sealed intakes quote the old names**, correctly for their day, and a reader who meets XENO-0278's "the label `approved`" has to know the date.
