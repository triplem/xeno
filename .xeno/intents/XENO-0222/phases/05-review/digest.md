---
intent: github.com/triplem/xeno#169
phase: 05-review
created: "2026-10-03T09:12:24Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1e1f1961edcb9f201b63b965185442abdded9b0d79b03f954bde04649f14a489
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three shipped review rules answered: deviations traceable, the interface note met because `--vendor`
is additive and A77 plus the release notes carry the manifests' contract, and no dependency added.
Beyond them: the slice stayed a slice, with the tempting non-goal — a stub `mcp.json` to make the next
piece smaller — refused because a client reading it fails at startup; the specification was followed
everywhere except two paths, where it describes a third party's format and the third party disagrees,
so A77 records what the client requires and the correction is raised rather than taken; every name in
the skills is checked against the tree, and the one skill that described a verdict nobody will see now
says the gate reports `not-implemented`; and the piece learned something the plan should know, which is
that the harness prices the surface at ~606 tokens always-on and ~620–830 per invocation where WP11
argues about a standing cost with no figure. Residual risk: nothing shows the skills help and they were
written by the agent they instruct, with the next intents as the first measurement of their kind; six
hundred tokens a session against an unknown benefit; section 13 now carries a known error; vendored
skills are pinned; one harness hides all leakage; the hook that ships is the easy half; and
`mcp.json` is in the document's vendor list and not in the tree.
