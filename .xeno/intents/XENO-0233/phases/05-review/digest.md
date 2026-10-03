---
intent: github.com/triplem/xeno#177
phase: 05-review
created: "2026-10-03T17:07:16Z"
schema_version: "1.0"
runner_version: dev+a18c1f3.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b3618a008e8dbde7cf2d88d214c759931c131bf39bc1bf056d8c3e2506c41fb3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three shipped review rules are answered, and the migration note is the substance: two fields
change what they say in every artifact written from here, `runner_version` in shape and
`plugin_version` in source and value, and a repository with no vendored plugin now produces a
G-Schema finding where it produced a plausible number — correct by section 5 and a behaviour change
somebody will meet as a red verdict. `deviations-are-traceable` is met with five, the first being
evidence rather than untidiness: P0 and P1 record 0.29.2 and the later phases 0.30.0, because #196
merged and cut a release mid-branch. Beyond the rules: the decision was put to the maintainer twice,
and the second time was the useful one, because sourcing a field correctly from a wrong literal still
reports a wrong number and nothing in the first exchange would have caught it. Nothing here makes a
verdict depend on the environment, which is the property the whole scope was chosen to protect, and
the check that nothing branches on the harness is in the required job. The trail is safe at 275 and
exit 0; one sealed lock was modified during the work by a mistyped `phase start`, caught by verify as
a divergence and restored, which is the contrast #193 was about — a modification is loud. Four header
fields now come from four sources and none is a constant, where two of four were. What remains is
that neither version is checked against anything, because G-Supply does not exist, and that the
manifest will be stale again at the next release.
