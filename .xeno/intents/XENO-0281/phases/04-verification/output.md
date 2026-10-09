---
intent: github.com/triplem/xeno#343
phase: 04-verification
created: "2026-10-09T13:24:41Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: eeb653f26fde12cef38d1fb21509ff48a6be48b32bb6194326dddb55030c9e58
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
| 1. the pin is a commit's pseudo-version with its reason | `gosec.yml` read: the install line and the six-line comment above it | met |
| 2. the scanner reads 1.27.2's export data | gosec installed at the pseudo-version, run with the workflow's flags against the tree: no type errors, `Nosec: 4`, `Issues: 0`; `go version -m` shows `x/tools v0.51.0`; the local toolchain is 1.27.1, so the 1.27.2 half is proved by the pull request's run, criterion 5 | met locally; the runner's run is the proof |
| 3. the row moves with the pin | `go test ./internal/model/` green; the row carries `v2.29.1-0.…` and the workflow the same string | met |
| 4. A104 records it | the status cell read back | met |
| 5. the check is green where it was red | the `gosec` check on this intent's pull request under Go 1.27.2 | recorded on the pull request after this phase is written; red yesterday on #342 and on `main` |
| 6. nothing else moves | `git diff --stat` outside `.xeno/intents/`: three files, 9 in, 5 out; `go test ./...`, `gofmt -l`, `go vet`, `gate verify` 558 | met |

<!-- xeno:section:results -->
## Results

Six criteria; five met on the branch and the sixth, the green check under 1.27.2, is the pull request's to show, since the local toolchain is 1.27.1 and cannot reproduce the failure.

    gosec at the pseudo-version      no type errors, Nosec 4, Issues 0
    go version -m gosec              golang.org/x/tools v0.51.0
    go test ./...                    ok, exit 0
    gofmt -l, go vet                 clean
    ./xeno gate verify               verified 558 verdicts

**Timings, for #117.** `gate verify` 36 s over 558, unchanged per verdict.

<!-- xeno:section:gaps -->
## Gaps

**The failure was not reproduced locally.** The machine runs Go 1.27.1, where 2.29.0 works; the diagnosis rests on the runner log's error text and the version difference between the green and the red run, and the fix is proved by the pull request's check and not here.

**The pin is a moving target by design.** A commit on main until 2.30.0; renovate proposes the release, and until then the row says a commit.
