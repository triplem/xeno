---
intent: github.com/triplem/xeno#346
phase: 04-verification
created: "2026-10-09T15:21:56Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bc7175d9db48c2eb28e2563a82c514cd8877cd2a3081c0803f1eaed8784277cd
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
| 1. the GitHub read asks by name | the line reads `gh label list --search "$label" … \| grep -qx "$label"`; `gh label list --search xeno-approved` on this repository returned the one name, and `grep -qx` is unchanged | met |
| 2. the GitLab read names its page | the line carries `--per-page 100` and the header says the bound and why; `glab` is not installed here and the flag and its default of thirty are from its `docs/source/label/list.md` | met by reading, see gaps |
| 3. the second run exits 0 | `sh .xeno/plugin/bin/xeno-labels.sh github triplem/xeno` with the label present and thirty-two labels on the repository: `xeno-approved exists on triplem/xeno; nothing changed`, exit 0; the same command on `main`'s script exited 1 with the host's refusal, run minutes earlier | met |
| 4. the test holds the read's shape | `TestTheLabelScriptCreatesTheLabelTheClauseNames` on `main`'s script: two strings reported missing, `FAIL`; on this branch, `ok` | met |
| 5. nothing else moves | `git diff --stat main`: the script and the test, 18 in and 5 out, and the trail; the constants, the description, the colour and the usage line are untouched | met |
| 6. suite and gates | `go test ./...` exit 0, `gofmt -l` empty, `go vet` clean, `sh -n` clean, `gate verify` 577 | met |

<!-- xeno:section:results -->
## Results

Six criteria; six met, one of them by reading rather than by a run.

    go test ./...                                  ok, exit 0
    gofmt -l, go vet                               clean
    sh -n xeno-labels.sh                           clean
    sh xeno-labels.sh github triplem/xeno          exists; nothing changed, exit 0 (was exit 1 on main)
    gh label list --search xeno-approved           one name
    ./xeno gate verify                             verified 577 verdicts

**The fault and the fix were both measured against the host**, which XENO-0282 could
not do because the label did not exist yet. The run on `main`'s script is the issue's
reproduction; the run on this branch is the done-when.

**Timings, for #117.** `gate verify` 48 s over 577 verdicts, from 36 s over 558 at
XENO-0282 on the same day; the per-verdict figure moved from 65 ms to 83 ms and
the count by nineteen, so the total did not grow with the count alone.

<!-- xeno:section:gaps -->
## Gaps

**The GitLab branch is read, not run.** `glab` is not on this machine and no GitLab project is available; the flag, its spelling and its default of thirty are from the tool's documentation on 2026-10-09. A project on GitLab is the first run.

**A hundred is the GitLab bound.** Past a hundred labels the GitLab read misses again and the script exits 1 as it did here; the line says so and nothing checks it.

**The test holds the flags, not the behaviour.** A later edit that kept `--search` and dropped the exact match, or moved the flag to a different command, would pass it; the test says what it holds.

**The timing moved more than the count.** Nineteen verdicts more and twelve seconds more, on one machine, one run each; whether the tree or the machine is why is for #117's series and not for this phase.
