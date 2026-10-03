---
intent: github.com/triplem/xeno#181
phase: 00-intake
created: "2026-10-03T12:37:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6cbeac4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b80e09fa0248b51d78079575cca8ef77914d046bf79b23ae78c021df2acbee25
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`tool_version` is required by G-Schema and written by no command, so every phase of every intent
here has gone red on it once and been corrected by hand — twice per phase, in `output.md` and in
`digest.md`, which is the order of three hundred edits to files inside `artifacts_hash`. The channel
the specification designed for it was never built: A35 assigns the field to the harness, section 7
designs a normalising entry point where "the runner only ever sees `XENO_*`" and lists three
variables, and the runner reads none of them — `grep os.Getenv` over `internal/` and `cmd/` returns
the enforcement token and nothing else. So the one harness fact a gate requires has nowhere to
arrive from, and A35's choice of absent over guessed is right while leaving nothing able to produce
a value. It is paid twice because the writers differ: `SectionSet` carries an existing artifact's
frontmatter over, so an edit there survives, while `writeDigest` builds its own and is rerun by
every `phase finish`, destroying the edit each time. In scope: `--tool-version` on `section set`, the
digest copying it out of `output.md`, absence unchanged, and the six skills reporting it. Out of
scope and each for a reason: a fourth `XENO_*` variable, which is a change to a list section 7
enumerates and therefore a person's commit; building the entry point, which is a WP11 gap of its
own; deriving the version in the runner, which is the harness branching section 7 forbids; and a
`project.yaml` field, which would be an Appendix A change and stale by design.
