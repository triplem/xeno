---
intent: github.com/triplem/xeno#346
phase: 03-implementation
created: "2026-10-09T15:20:13Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ceb7dfaa54dfd27113b53314d8533549915dcf6429fd43000b7baa477b8918ab
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

Two files, 18 lines in and 5 out.

**The script: `.xeno/plugin/bin/xeno-labels.sh`.** The GitHub read is `gh label list
--search "$label" --repo "$project" --json name --jq '.[].name' | grep -qx "$label"`,
the search ahead of the repository so that the line reads as what it asks; the exact
match on the names stays. The GitLab read is `glab label list --per-page 100 --repo
"$project" --output json`, the JSON name match unchanged. The header gains a paragraph
saying what the idempotence rests on, why the two reads differ, and that the GitLab
branch meets the same fault past a hundred labels, with #346 as the reference.

**The test: `internal/plugin/plugin_test.go`.** `TestTheLabelScriptCreatesTheLabelTheClauseNames`
requires `gh label list --search` and `glab label list --per-page` among the strings
the script carries, and its comment says why a test without a host holds the shape of
the read and not its answer. Written first: on `main`'s script it reported both
strings missing, which is the failure acceptance criterion 4 asks for.

**Measured on the branch.** `go test ./...` green, `gofmt -l` empty, `go vet ./...`
clean, `sh -n` on the script clean. The second run against this repository, with
thirty-two labels, printed `xeno-approved exists on triplem/xeno; nothing changed` and
exited 0, where `main`'s script exited 1 with the host's refusal; `gh label list
--search xeno-approved` on this repository returns the one name. The `gate verify`
figure is P4's.

<!-- xeno:section:deviations -->
## Deviations from the design

None. The design named the search on the GitHub read with the exact match kept, the page of a hundred on the GitLab read with the bound on the line, and the two strings in the test, and each is as written.
