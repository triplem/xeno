---
intent: github.com/triplem/xeno#355
phase: 03-implementation
created: "2026-10-10T13:12:11Z"
schema_version: "1.0"
runner_version: dev+b9beef5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a18fbb9eef30fc310d01c4800521cb39248878e647d3df0934e82abb4f3f56bd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Two commits. `docs(supply-chain)` first and alone, as the approval on #355 and the
standing rule both ask, then `fix(ci)` with the four workflow edits that follow from it.
121 lines added and 14 removed in `release.yml`, and nothing at all in the `Dockerfile`.

The dispatch path was run rather than reviewed. Outside the repository, with the real
token and v0.60.2's real assets: the download succeeds, `sha256sum -c` prints
`xeno-0.60.2-linux-amd64: OK`, the `chmod` leaves `-rwxr-xr-x`, the build context holds
the binary and nothing else, and the binary reports `xeno 0.60.2`. Criterion 8's two
failure modes were run too — `v9.9.9` exits 1 with "release not found", an asset pattern
naming a version the release does not carry exits 1 with "no assets match the file
pattern" — and both stop under `set -eu` before the `docker login`, so nothing is pushed.
The base extraction was run against the real `Dockerfile` and resolves to the digest, the
tag `13-slim` and the mirror reference `ghcr.io/triplem/xeno/base@sha256:a29215f6…`.

One deviation, and a test caused it. The design meant to widen the `docker` row to say it
copies as well as builds and pushes. `exempt` in `internal/model/supply_chain_test.go` is
keyed on that row's exact first cell, so moving the words would have unexempted the row,
failed the test, and made the documents commit a code change as well. The sentence went
into the prose beside the fetch path, which is where the reason lives anyway, and the
documents commit stayed what the rule asks for. The learning generalises it: before a
documents-only commit, grep for the document's text and not only for its path.

The gates pass on the tree as it stands rather than on one the change was reverted out of:
`gofmt -l` empty outside `vendor/`, `go vet ./...` clean, `go test -count=1 ./...` clean,
`xeno gate verify` at exit 0 over 594 verdicts, and `xeno check commit-message` at exit 0
on both commits. The workflow parses, and its steps and conditions were read back out of
the parsed YAML rather than off the diff.

Two sections, no open question, no decision.
