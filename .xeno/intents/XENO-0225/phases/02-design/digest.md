---
intent: github.com/triplem/xeno#172
phase: 02-design
created: "2026-10-03T10:35:36Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d3d721517b3e89eec1361e4b07f2dec3b2a1586d9e9ec2bb1cfac200dd1902db
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The check is in G-Schema beside the budget, because both read the same two files — the profile from P0,
the lock from the phase — and a second place that knows how to find a profile would be a second place
to keep right. The finding names the profile rather than the base: the base is the consequence and the
profile is the claim, and a finding on the lock would point a reader at a file that is correct about
what it was given. It is checked wherever the profile is read, which is every phase, so a mistyped link
is reported at the first gate run after it is written and at every one after until it is corrected —
where a one-phase check would report it once into a verdict nobody re-reads. The runner still skips the
link in silence and that is correct: it writes the lock, the lock describes what the phase was given,
and a file that is not there was not given. Rejected: reporting from the runner, which would make the
writer of a record its judge; a refusal, against #171's criterion and the budget's precedent; checking
the component, which may not be a directory at all; and fixing the byte count here, which needs a field
section 5 does not enumerate and is therefore a person's commit. Nothing goes to the register, because
M0 closed it this morning — this is the first intent to use the loop that replaced it rather than
describe it.
