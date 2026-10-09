---
intent: github.com/triplem/xeno#95
phase: 02-design
created: "2026-10-09T16:49:47Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f6d8680bc355e51dbd4bbe037c45ec6e673f547e62e0518ab48017c1f3e75426
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Two files, three appended managers, one removed key, one changed line.

The design question that took the longest is the smallest edit: whether the `Dockerfile`'s base
image reference gains its tag. Matching the digest alone keeps the change inside
`renovate.json`, which is tidier, and it does not work — a digest names no stream, so renovate
has nothing to ask for a newer digest of and falls back to `latest`, which is not the image the
page calls 13-slim. The tag also closes a gap that has nothing to do with any bot:
`docs/supply-chain.md` names a tag the tree does not carry, so the page could not be checked
against the file.

The native managers were both considered and both rejected for the same reason, which is the
reason this intent exists: a manager that resolves nothing reads as coverage. `dockerfile`
would depend on renovate's handling of an `ARG`/`FROM` pair, which this repository cannot test
and would not notice losing; `pip_requirements` would find no requirements file at all.

Whether the test suite could hold these regexes was settled against: the dialect is
JavaScript's, so a Go test would check a translation, and a translation that passes while the
original fails is the exact undetected gap being closed. The check is a script whose output is
P4's evidence.

The `Dockerfile` edit's effect on `verify` was read out of the pattern before the edit was
made. `supply_chain_test.go` extracts the digest with a group whose character class excludes
`:`, so inserting `:13-slim` moves what the first group captures and leaves the second — the
digest, the only one stored — alone. That is why the test file is in the context scope: so the
answer came from the pattern rather than from running it and hoping.

Two sections of the four, no question, no decision.
