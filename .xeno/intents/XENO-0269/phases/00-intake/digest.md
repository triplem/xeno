---
intent: github.com/triplem/xeno#207
phase: 00-intake
created: "2026-10-07T09:05:14Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f976cf5fdb4a104187d89a283f5ba966ce4731df032eb8bfd1b8c8f044f8a134
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Section 12 requires `model`, `tool` and `tool_version` in every artifact and calls them the
raw material for a provider register. All three have writers and nothing corroborates any of
them: `--tool-version` or `XENO_HARNESS_VERSION` for the version, `project.yaml` for `model`,
`project.yaml` with `XENO_HARNESS` over the top for `tool`, and G-Schema requiring presence
and checking shape beside them.

Corroboration is unavailable. Section 7 keeps the runner from branching on the harness, so it
cannot ask; `project.yaml` is the same declaration in a second place, and #207 names a check
against it as the thing that would hide the gap; and the one record from another hand, a
gateway's, is both outside the repository and outside the gate path, which never touches the
network. The plan's section 9 also says what it would be worth: a routing record rather than
an attestation.

So the statement is what is left. One paragraph in section 12, approved and written on
instruction, one count moved in `docs/clause-readers.md` with the reason it has no row, and one
register row naming `plugin_version` as the contrast.

Out of scope: a check against `project.yaml`, the gateway comparison, a new gate result for an
unverified field, and removing any of the three fields.
