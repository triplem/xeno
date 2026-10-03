---
intent: github.com/triplem/xeno#183
phase: 04-verification
created: "2026-10-03T17:33:12Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1cf9b338ccb2a59043a6b1ecfcd73396764429e651ad108967aa6f7c6cb5a991
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Every criterion maps to a test, and the one doing the work is that the anchor is not readable from the
repository: the test writes the correct digest into four plausible places, including inside the
plugin's own manifest directory, and requires a failure anyway — every other test passes under an
implementation that reads the lock, so this is what separates this gate from the obvious wrong one.
The earlier-plugin test differs only in a `version:` line, so a gate reading the version could answer
differently and this one does not. The no-anchor test asserts both the result and the absence of
findings, because a check that did not run must not carry a complaint. Eighteen packages ok with
seven new cases, `gofmt` and `go vet` clean, `gate verify` 279 at exit 0, and A42 intact — the gate
path still reaches neither the network nor the index. Measured with a binary built the way the
release builds one: a matching tree passes, an appended comment fails naming both digests, the plugin
moved away fails saying `init --vendor`, and the development build reports `not-implemented`; the
released binary verifies this whole trail at exit 0, because `gate verify` compares a phase's status
and both results derive to green. Six gaps, the largest found while writing them: a released binary
cannot vendor the plugin it carries the digest of, since `init --vendor` copies from the repository's
own tree — the gate is anchored to something the release cannot deliver, which wants its own issue.
Then: nothing confirms a published binary carries the tag's digest; the gate is untested against a
real release and a mis-assembled ldflags line would switch it off silently; `gate verify` compares
status rather than per-check results; the resolution order's containment is only as good as the
binary running it, since a dev build has no anchor; and no signature, which section 13 disclaims.
