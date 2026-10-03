---
intent: github.com/triplem/xeno#177
phase: 04-verification
created: "2026-10-03T17:06:24Z"
schema_version: "1.0"
runner_version: dev+a18c1f3.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9200ab5300824cae228c25689a9a8a0fd07021b5b7dc9d06f9087a0d4bd9a08c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Most criteria are held by tests and three by files or commands, which is what a change spanning the
runner, a workflow and a shipped script means. The plugin version is asserted against a manifest
declaring `0.26.0` and against not being the runner's; absence is asserted in both halves, the empty
field and the G-Schema finding, because absence without the finding would be a quiet gap; the lock's
block is asserted in both states and against its own header, which is the comparison G-Supply will
make. The stamp test asserts that the part before `+` does *not* parse as a version, which is the
decision, and that the metadata after it still does, which is what was worth keeping from the test it
replaced. The entry point is asserted not to export `XENO_PLUGIN_ROOT`, which is the assertion that
would catch somebody adding it without reading why it is absent. Eighteen packages ok, `gofmt` and
`go vet` clean, `gate verify` 275 at exit 0 with the code captured rather than piped — three times
today a `| head -1` reported `head`'s status and once hid a sealed artifact this branch had modified.
The tree hash agrees byte for byte with a shell pipeline implementing the written definition
independently. This intent's own four header fields come from four sources and none is a constant,
where two of the four were before it. Seven gaps: the resolution order, which is #183's larger half
and a section 7 decision; the manifest, bumped twice in one session because a release landed
mid-branch; no statement in section 16; G-Supply still absent so the lock's block has no reader; the
hash pinning the tree rather than the plugin's identity; `plugin_version` now falsifiable and
unfalsified; and the entry point asserted in shape and not in behaviour.
