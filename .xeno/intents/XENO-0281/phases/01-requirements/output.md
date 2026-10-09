---
intent: github.com/triplem/xeno#343
phase: 01-requirements
created: "2026-10-09T13:19:47Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dab5681c0c92e8ac60980031fa1827678b27b3feb4b598bb98bf134b714c5466
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

1. **The pin is a commit's pseudo-version.** `gosec.yml` installs
   `github.com/securego/gosec/v2/cmd/gosec@v2.29.1-0.20261009120814-7b1b5cebe007`, and
   its comment says why a commit and when it moves to a release.
2. **The scanner reads 1.27.2's export data.** gosec installed at that pseudo-version and
   run with the workflow's flags against the tree reports no type errors, four accepted
   findings and zero issues; `go version -m` on the binary shows `x/tools v0.51.0`.
3. **The row moves with the pin.** `docs/supply-chain.md`'s gosec row names the
   pseudo-version and `go test ./internal/model/` is green in both directions.
4. **A104 records it.** The status cell names the commit pin and the release that ends it.
5. **The check is green where it was red.** The `gosec` check passes on this intent's pull
   request under Go 1.27.2.
6. **Nothing else moves.** The diff outside `.xeno/intents/` is one workflow line and its
   comment, one table row, one register cell; suite, format, vet and `gate verify` green.

<!-- xeno:section:non-goals -->
## Non goals

- **Pinning Go to a patch**, which would make the run green today and freeze CI on 1.27.2 against D-5.
- **A tools module** carrying gosec with a bumped `x/tools`: a second `go.mod` in the tree for a line that a pseudo-version answers.
- **Waiting for 2.30.0**, with every merge red meanwhile.
- **Any change to the policy or the four acceptances.**

<!-- xeno:section:constraints -->
## Constraints

- Every tool pinned exactly and on the pin table, held both ways by the test; a pseudo-version is exact.
- The toolchain stays pinned by major, D-5.
- The job id and the required check `gosec` do not change.
- Prose at 88, a paragraph changed twice is replaced.
