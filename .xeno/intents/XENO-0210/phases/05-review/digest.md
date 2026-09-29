---
intent: github.com/triplem/xeno#110
phase: 05-review
created: "2026-09-29T19:14:14Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cf4b09894211fd0a1b0d591484d3c32beda49ef4a917dfff4eeb6b0fab824a93
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The risk worth carrying is the one the change created: main is now the line that matters most, since it
supplies the two writers in an order nothing checks, and passing them the wrong way round would send verdicts
to standard error while every test in the package passed. The other is unchanged and honest — twenty printers
untested, and no test anywhere takes a real provisional phase through to an exit code.
