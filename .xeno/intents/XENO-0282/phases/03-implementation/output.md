---
intent: github.com/triplem/xeno#332
phase: 03-implementation
created: "2026-10-09T13:41:10Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bf0be6c8b2238c3d99de255c56fc938a2a0f519e82e29b796656b815152dcc80
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Thirteen files, 70 lines in and 41 out, plus the clause that preceded them in `b431e42`.

**The constants: `internal/model/identity.go`.** `ApprovedLabel` is `xeno-approved` and
`ApprovedWord` is `/xeno approved`, with the reason on the constants; the comparison in
`Approval()` is unchanged and its comment names the new forms. `runner.go`'s refusal drops
"the word" and prints the constant, since the first line is two words now.

**The script: `.xeno/plugin/bin/xeno-labels.sh`.** POSIX shell in the entry point's shape,
executable, `set -eu`, host and project as arguments, a read before the write so a label
already there is reported and left alone, `gh label create` or `glab label create` with
the clause's other half as the description, and a usage line for a third host.
`TestTheLabelScriptCreatesTheLabelTheClauseNames` in `plugin_test.go` holds that it is
there, executable, and carries the two constants' values and both commands.

**The line in init: `internal/runner/init.go`.** A sixth entry in `Manual`, built from
the two constants so that it cannot drift from them, naming the script;
`TestInitNamesWhatItCannotDo` looks for the label and the script's name.

**The fixtures.** Six test files spell the new names: `tracker_test.go`'s fake host and
refusal strings, `mcp_test.go`, both adapters' `tracker_test.go`, `main_test.go` and
`identity_test.go`, which gains the case *the names before #332, now two missing
halves*. The first pass replaced too much and too little: it hit a gate decision
`"approved"` in `tracker_test.go` and missed a capitalised label and the GitLab
fixture, which the suite reported and the second pass corrected.

**The documents.** `docs/commands.md`'s paragraph names both forms, why they carry the
tool's name, and the script; `docs/assumptions.md` gains A107.

**Measured on the branch.** `go test ./...` green after the second pass, `gofmt` and `go
vet` clean, `sh -n` on the script clean, `xeno gate verify` 558 verdicts. The script was
not run against a host from this session; the host act is the merge's.

<!-- xeno:section:deviations -->
## Deviations from the design

None. The design named the constants, the script beside the entry point, the init line, the fixtures with one new case, A107 and the host act at the merge, and each is as written. The second pass over the fixtures is a correction inside the phase and not a departure from the design.
