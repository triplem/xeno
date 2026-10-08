---
intent: github.com/triplem/xeno#95
phase: 04-verification
created: "2026-10-08T13:41:45Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2ce6548ddb505638d865bf506d3d075bda8b61a008d046433b643610dafa7f6c
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
| 1. renovate is configured | `renovate.json` exists at the root; `python3 -c "json.load(open('renovate.json'))"` parses it; it declares `$schema` as renovate's own schema reference and names the three managers this repository has pins for | met |
| 2. the dashboard is the default | `prCreation: "approval"`, `dependencyDashboardApproval: true` on every rule, `automerge: false`, and no `schedule` key anywhere. Read out of the committed file rather than from a run | met |
| 3. a breaking bump is not proposed silently | each rule's `description` says that a pinned sha and its row in `docs/supply-chain.md` move together, so the constraint is in the file a person opens. The automatic path for those managers is the dashboard, which is the same answer from the other side | met |
| 4. nothing is bumped | `git diff main..HEAD --stat` touches `renovate.json` and `.xeno/`, no workflow and no `go.mod`; `go test ./internal/model/` passes, which is #317's two tests saying the pins and the page still agree | met |
| 5. what this does not reach is written down | the intake's scope says switching dependabot off is a host setting, and `renovate.json` does not pretend to reach it | met |

<!-- xeno:section:results -->
## Results

Five criteria, five met, and the checks are the ones a reader can run.

    python3 -c "import json;json.load(open('renovate.json'))"   renovate.json parses
    go test ./internal/model/                                    ok
    go vet ./...                                                 clean
    gofmt -l . | grep -v '^vendor/'                              nothing
    ./xeno gate verify                                           verified 523 verdicts

The fourth is the one worth naming: it is #317's pair of tests, and they pass, which is the
evidence for criterion 4 — the pins in the tree and the rows of `docs/supply-chain.md` still
agree, because nothing bumped either.

**What was not checked, and cannot be here.** Renovate itself never ran. Nothing in this
repository can run it: it needs the app installed on the host and a token, which is the same
class of thing as the protected branch settings. So criteria 1 to 3 are checked against the
committed configuration and renovate's published schema, not against its behaviour. The
first real dashboard is the first evidence of the behaviour, and it will arrive after this
merge rather than before it.
