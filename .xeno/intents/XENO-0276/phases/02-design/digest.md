---
intent: github.com/triplem/xeno#325
phase: 02-design
created: "2026-10-08T13:51:38Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3558d70f124ea46ee3ef5224556dabc92a24c0f6cfabb2d93e84d448d1138ea0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
One file, and the design question is placement rather than wording.

The two writers go between the sections and the finish, because that is the order a phase does
them in and the block is read top to bottom. The rules first and the lens second, matching the
sentence above them. The "why two commands and not a flag" sentence goes into the paragraph
that already describes the entries, because the absence of a rule id is a property of the entry
rather than of the command, and that paragraph is what the reader has just read.

Not added: a `--file` form, which these two commands do not have.

One section of the four, `impact`.
