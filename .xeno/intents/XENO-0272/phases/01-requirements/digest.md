---
intent: github.com/triplem/xeno#277
phase: 01-requirements
created: "2026-10-07T13:28:54Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5973737622e49269831018c24f084e3145350abeea617e0ff79f83d3c1acba79
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Thirteen criteria. Five are the change: `honest` no longer reads `tool`, no gate reads any of
the triple to decide anything, the comment says what was removed and why, `hashShape`'s doc
comment describes what exists, and the test case that asserted the old behaviour asserts the
new one. Five are the consequence: exactly two phases change verdict on exactly four findings,
both read `approved`, the releases are approvals so nothing is owed, `gate verify` exits zero
again, and no artifact's content is edited. Three are the absences: no normative document
changes, the register carries a row, and the suite passes.

The non goals are the restraint: no specification change, no correcting the four artifacts, no
override, nothing about `model` or `tool_version`, `goneBundle` stays, the `tool` field stays
required, no new gate or finding kind, and no check comparing `tool` against `project.yaml`,
which #207 refused in the document.

The binding constraints are section 11, which is why the findings are released rather than
fixed and which this change bumps against by rewriting two `gate.yaml` files; that a release is
a second person's statement, which is why the reasons were drafted and approved before the
commands were run; and the rule about a negative result, since "nothing else in the trail
moves" is an absence and is established by running `gate verify` over all of it.
