---
intent: github.com/triplem/xeno#145
phase: 03-implementation
created: "2026-09-30T05:51:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4e41c6b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b2e517c05c97cd3f1ad9bb0b36e37b1b4abbec665de85be14e5d8c180723ab0a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
One scaffold file and one table row, which is what #98 promised: WrapperHost now describes a host
completely — the two range expressions, the tracker adapter and address, and the squash setting in that
host's words. The last was a function branching on the host id for one draft, and moving it to the table is
what keeps a third host from editing code. Both defaults are gone, cmd/xeno's flag and init.go's fill-in, and
the host is resolved once before either file is generated so the wrapper and the tracker block agree by
construction. Twelve existing tests gained Host: "github" because the flag is now required; their assertions
are untouched, which is the distinction from the rule the last intent wrote about not editing tests to pass.
The generated GitLab YAML parses and its folded script line resolves to one command with both range flags.
