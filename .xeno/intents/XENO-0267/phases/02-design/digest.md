---
intent: github.com/triplem/xeno#228
phase: 02-design
created: "2026-10-06T20:29:40Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fec25d92f5563bdedd73c107a1312efed0de0812ecfc6d8968d1acd269899d1a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The sentence hangs on the phase state rather than on the command that asked, because `Next`
knows nothing about its caller and a sentence that appeared only after `phase finish` would
make `Suggestion` something other than a reading of section 6's table. So it lands on the
`not-started` answer, which is what a person sees after a finish that went green and after
`intent start`, and it keeps `xeno phase start` as its command. It says why in the lock's own
terms, because a reason a person can check against a file they already have is followed more
than once.

The end-of-intent sentence goes on the last phase's answer, which already carries no command,
and names what is not read again: the next intent starts from the tree and the scope its own
intake declares.

Neither reads `XENO_HARNESS`; there is no condition to add and the CI check on the constant is
untouched.

Four alternatives were refused with their costs: the harness's command in the plugin's skill
text, which breaks the interchangeability the skills exist for and addresses a reader who
cannot carry it out; a hook, which #228 prices itself; the cost figure now rather than in WP20,
which would be the lever before the baseline; and a gate, which would have to read gitignored
local data to reach a verdict. Saying nothing is the state the issue found.

Impact is two strings, two assertions and one register row. No artifact gains a field, no gate
changes, and every sealed verdict in the trail verifies as before.
