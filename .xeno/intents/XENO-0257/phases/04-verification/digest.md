---
intent: github.com/triplem/xeno#257
phase: 04-verification
created: "2026-10-06T07:58:51Z"
schema_version: "1.0"
runner_version: dev+77a35b1.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 502737ad99e3ee986a40d94604daac626cde2b5df989e11a6dcc9c5694f62994
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Eleven criteria met, one pending until the commit. The script that compares the candidate against the
shipped set is the verification that mattered: ids, versions, phases and headings, all six templates.

What the intent does not know is whether six sections are enough, which needs an intent run through
the set — and the fixture is what makes that cheap rather than what answers it.
