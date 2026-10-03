---
intent: github.com/triplem/xeno#172
phase: 00-intake
created: "2026-10-03T10:34:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 355b353ef38ec51d71acb083e9c539ff5f834f646d955646c053872510e63b05
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
A link naming a document that is not in the tree is skipped in silence: `informationBase` cannot hash a
file that is not there, so the one field in a profile that names a specific path rather than a pattern
is the one whose mistakes nobody reports. It is an acceptance criterion of #171 that #171 did not meet —
P3 took a different path, no test had been written, the suite was green, and P4's table found it after
P3 was sealed. It matters more than its size because everything else in a profile is a pattern, and a
pattern matching nothing is indistinguishable from a pattern whose files do not exist yet; a link is
declared precisely because section 5 forbids inferring one, so a mistyped link is the profile's only
unambiguous error. It is also the sharp edge in the way of adopting a profile here, since adopting a
mechanism whose first mistake is invisible means the first mistake is invisible. The other half, the
byte count measured after the fact rather than recorded, needs a field in the lock, and section 5
enumerates those fields — a specification change and therefore a person's commit. This intent names it,
offers the wording, and stops.
