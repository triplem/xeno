---
intent: github.com/triplem/xeno#324
phase: 04-verification
created: "2026-10-08T13:48:57Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 94d31936a06e7ada2eabd5194359e0eb252cb40db55365476e9bc8c0e2a1b7f3
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
| 1. the rendered wrapper carries the line | `xeno init --host github` in a throwaway repository, then read the `container:` block out of the generated `.github/workflows/xeno-gate.yml`: the commented `options: --user 1001` is there, under the image | met |
| 2. the reason is beside it | the same six lines above it say what the uid is for and that a hosted runner needs nothing, read in the rendered file rather than in the template | met |
| 3. it still parses as YAML | `yaml.safe_load` over the rendered file: parses | met |
| 4. a test asserts it | `TestOnlyTheGitHubWrapperOffersTheContainerUser`, beside the one that asserts every wrapper names the image. It also asserts the line stays commented, and that no other host's wrapper mentions `--user` | met |
| 5. GitLab is unchanged | `xeno init --host gitlab` in a second throwaway repository: `grep -c user .gitlab-ci.yml` is 0, and the test asserts the same for every host that is not github | met |

<!-- xeno:section:results -->
## Results

Five criteria, five met.

    go test ./internal/runner/     ok
    go vet ./...                   clean
    gofmt -l . | grep -v vendor/   nothing
    xeno init --host github        the block carries the commented line
    xeno init --host gitlab        no mention of a user
    yaml.safe_load                 the rendered workflow parses

**One thing worth recording about the check itself.** The first run of criterion 1 failed and
the failure was mine: I read the rendered wrapper from a binary built before the template
changed, because the templates are embedded in the binary. The test passed at the same moment,
which is exactly the gap between a test over the template and a check over what an adopter
gets. Rebuilt, both agree.

**What was not checked.** No container job ran on any runner, hosted or self-hosted, so the
claim that uid 1001 matches a hosted runner's work directory is still read out of the runner's
sources rather than observed — the same open confirmation #305 left behind. This intent does
not close it and does not pretend to: what it delivers is the sentence an adopter needs, not
the proof that 1001 is right.
