---
intent: github.com/triplem/xeno#332
phase: 01-requirements
created: "2026-10-09T13:32:25Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6fbca72005afe6e491ed005b0dcfcd9a5d886a04be045cdb32f703bfcd80c92b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **The clause precedes the code.** `b431e42` is the first commit after `main` on this
   branch and names both new names, the script and the `init` line; revision 11.
2. **The constants say the new names and nothing accepts the old.** `ApprovedLabel` is
   `xeno-approved`, `ApprovedWord` is `/xeno approved`; `Issue.Approval()` on an issue
   labelled `approved` with a comment `approved` reports both halves missing, and on one
   labelled `xeno-approved` with `/xeno approved` reports neither; the comparison stays
   trimmed, case ignored, trailing punctuation dropped.
3. **The refusal and the sentence print the new names.** `intent start` on an issue
   missing either names `the label xeno-approved` or `a comment whose first line is
   /xeno approved`; the intake's "No approval was found" sentence names the same.
4. **The script is shipped, executable and idempotent.** `.xeno/plugin/bin/xeno-labels.sh`
   exists with the executable bit, takes a host and a project, creates `xeno-approved`
   with a description through `gh` or `glab`, and exits 0 without changing anything where
   the label exists; a test beside the entry point's holds the first two, and `sh -n`
   the syntax.
5. **init names the label.** `InitResult.Manual` carries a line naming `xeno-approved`
   and the script, and `TestInitNamesWhatItCannotDo` holds it.
6. **Every fixture and document moved.** `grep -rn` for `"approved"` as a label or a
   comment's first line in `*_test.go` finds nothing; `docs/commands.md` and the intake
   skill name the new forms.
7. **The register carries the decision.** A107.
8. **The suite and the gates stay green.** `go test ./...`, `gofmt -l`, `go vet ./...`,
   `./xeno gate verify` exit 0.
9. **The host is relabelled at the merge.** `xeno-approved` exists on the repository,
   every issue that carried `approved` carries it, and `approved` is gone; the pull
   request records the read-back.

<!-- xeno:section:non-goals -->
## Non goals

- **Accepting the old names beside the new.** The clause fixes two names; a runner reading four leaves a reader unsure which counted.
- **Reading the label back from the host.** A label is not in the enforcement block, and adding one is Appendix A.
- **Labels for #323's mechanisms.** No use case, and trigger labels are the design section 12 refuses.
- **Configurable names.** Rejected on #330 and again in the analysis.
- **Rewriting any sealed intake.** They record the names of their day.

<!-- xeno:section:constraints -->
## Constraints

- Section 12's clause as changed is the specification; the constants are its only home in code and nothing reads a configuration for them.
- `internal/gates` reaches no host; nothing here touches a gate.
- The script is a shipped file of the plugin, so it is vendored by `init`, executable under `bin/`, and part of the digest G-Supply compares on a release.
- One vendored dependency; this adds none, and the script needs only `gh` or `glab`, which a person who sets labels has.
- Prose at 88, commit subjects as Conventional Commits, a paragraph changed twice is replaced.
