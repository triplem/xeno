---
intent: github.com/triplem/xeno#325
phase: 04-verification
created: "2026-10-08T13:53:09Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a4be30412ac22b4c844a1a0b72b8826281f5ff30e8242f265933321f1241f60f
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
| 1. the block names both writers | read the committed skill: `xeno review answer RULE …` and `xeno review lens …` sit between the section writes and `phase finish`, with the sentence that the answers come first | met |
| 2. it says why two commands and not a flag | the paragraph above the block carries it: a lens entry answers no rule, its arguments have no place for one, and a flag could be passed beside a rule id | met |
| 3. the `source: lens` paragraph stays | unchanged; `git diff` over the file touches the sentence before it and the command block, and leaves that paragraph alone | met |
| 4. the plugin test holds the lines | `TestEveryCommandASkillNamesExists` reads the dispatch table out of `main.go` and every `xeno …` line out of every skill. Renaming the new line to `xeno review note` fails it: *xeno-review names `xeno review note`, which the dispatch table does not have*. Restored, it passes | met |
| 5. nothing outside the plugin changes | `git diff --stat` over the merge: one file under `.xeno/plugin/`, plus this intent's own trail | met |

<!-- xeno:section:results -->
## Results

Five criteria, five met, and the fourth needed two attempts to check — which is the result
worth recording.

The first sabotage renamed the command to `xeno review lensX` and the test passed. That looked
like criterion 4 being false. It was the check that was wrong: the pattern a skill's commands
are read with is `^ {4}xeno ([a-z-]+)(?: ([a-z-]+))?`, so an uppercase letter ends the match
and `lensX` was read as `lens`, which the dispatch table does have. Renaming to `xeno review
note`, which is lowercase throughout and absent from the table, fails exactly as it should:

    plugin_test.go:162: xeno-review names `xeno review note`, which the dispatch
    table does not have

So the test does hold the new lines, and the near-miss is a property of the pattern rather
than of the skill. A command whose name differs from a real one only by case would pass
unnoticed; that is a finding about the pattern, recorded and not fixed here, because a skill
naming a command in mixed case is not a defect this intent met.

    go test ./internal/plugin/   ok
    go test ./internal/runner/   ok
    gofmt -l . | grep -v vendor/ nothing
