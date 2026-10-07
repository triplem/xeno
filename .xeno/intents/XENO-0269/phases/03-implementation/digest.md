---
intent: github.com/triplem/xeno#207
phase: 03-implementation
created: "2026-10-07T09:08:17Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 91456688e97f3921b20cfb81360977163e07dfca0dff93fdfc2aa2aec36f8dc7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
One paragraph in section 12: the triple is a declaration and not a measurement, where each of
the three comes from, what G-Schema does with them, why the runner cannot corroborate one, why
a check against the second copy is refused, what the gateway's record is and why it is out of
reach, and how the register is read.

Two paragraphs and one count in `docs/clause-readers.md`: explanation 129 to 130, why the
clause has no row, and why it is reader-shaped in one direction only that is already taken.

One row, A100, with the decision, the two options not chosen, the three writers as the code has
them, `plugin_version` as the contrast, and the correction to the draft.

Three deviations. The approved wording said `model` comes from `project.yaml` or the harness
and the code reads it from `project.yaml` alone; the phrase was inherited from the issue, and
the specification commit was amended rather than corrected afterwards because it had not been
pushed. The paragraph was replaced a second time to drop two em dashes the document does not
use. And the row is A100, because A98 and A99 are on the unmerged branches for #228 and #238
and a row's number is cited from other rows.

No code, no gate. Build, tests, gofmt, vet and gate verify pass.
