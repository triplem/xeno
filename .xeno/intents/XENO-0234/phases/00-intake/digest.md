---
intent: github.com/triplem/xeno#183
phase: 00-intake
created: "2026-10-03T17:30:24Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 411b3f7c1e2d855f01d3eb145085cdd16425870fb556c1b2c945145dea399b6e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
G-Supply is the first gate section 7 lists and the reason it gives for the order — a plugin that does
not match its expected digest makes every later verdict a statement about unknown rules and unknown
templates — and it was written as `not-implemented`, so every verdict here is that statement. It is
also the gate the last two intents kept arriving at: #177's `plugin_version` was a constant whose
only purpose, by section 5's own sentence, is to be compared with the lock's hash by this gate, and
#183's resolution order was unbuildable because the three defences against a swapped plugin were all
absent — the field a constant, the lock without a `plugin` block, and this gate missing. Two were
closed and this is the third. What section 13 asks for is specific and is not what the lock does:
the anchor is the binary, "nothing in the repository states what the expected value is, which is the
point: an expected hash stored beside the thing it describes proves only that both were written by
the same hand." So the lock's `plugin.sha256` is not the anchor and must not be read as one, and the
downgrade needs no version check because an earlier plugin has a different digest. The case the
specification does not name is a build carrying no digest, which is every artifact here: reporting a
pass would be exactly what section 5 defines `not-implemented` to prevent. Out of scope: the
resolution order, which this gate was recommended ahead of; embedding the plugin in the runner,
which A58 already records and which is a different gap; reading the lock's two halves against each
other, which section 13 excludes; and G-Secret, which waits on a hook.
