---
intent: github.com/triplem/xeno#228
phase: 05-review
created: "2026-10-07T06:41:55Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b9e8695e59946aa922c6d82bb2740a86b6fd729aa052f9c7999b98ae1c85611b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three standing rules hold: no normative document is touched, nothing is invented, and the
change belongs to WP11 and to intent XENO-0267 for issue #228. `docs/assumptions.md` is the
open register and A98 is a decision that outlives this intent.

All three of #228's asks are answered. The `not-started` suggestion says a phase is started in
a fresh session and why; the last phase's suggestion carries the end-of-intent hint the issue's
comment asked for as a hint and not a requirement; and the two decisions the issue left live
are taken and recorded in A98 with the clause behind each.

One of them goes the other way from the issue's own table. #228 offers the plugin's skill text
as the place for a harness specific command, on the ground that the plugin is per harness; in
this tree section 13 ships one vendored plugin for both clients. The row says plainly that what
it wants from a person is a yes to that reading, and the plan leaves per-harness packaging open
to dogfooding, so the no costs a line in a SKILL.md if it comes.

Three rules answered: one deviation, the register row being `open` rather than `approved`
against P1's criterion 10, and two not-applicable.

The residual risk is the hint itself. Nothing enforces or measures it, the figure that would
belongs to WP20, and A98 is the only thing joining the two — so the real risk is that nobody
reads the row. The reading of section 13 could also be wrong, and the absence of a harness
command is checked over three trees and not over the documents.
