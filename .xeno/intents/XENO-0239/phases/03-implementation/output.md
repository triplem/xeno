---
intent: github.com/triplem/xeno#213
phase: 03-implementation
created: "2026-10-03T20:05:25Z"
schema_version: "1.0"
runner_version: dev+427eeb7
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 30b823acf0445f5797f28ab2ded39f833496c13d834f384e36dc6b7f014fd7a6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`docs/implementation-plan.md`, WP8, in its own commit ahead of this intent. The paragraph
that said "Reading outside the profile is recorded, not blocked" is replaced by one that
says the record goes in the phase's digest and not in `context.lock.yaml`, and gives
section 5's reason for it: the lock is written before the agent starts, states what the
phase was given, and one rewritten at the end would describe nothing. The replacement
names the record as the agent's own account, because nothing in the harness reports what
was opened.

The done-when changes with it: from "every read outside the profile appears in
`context.lock.yaml`" to "a read outside the profile is named in the digest of the phase
that made it".

`ASSUMPTIONS.md` gains A92, carrying the contradiction, the digest's four properties, all
three resolutions, and why there is no gate.

Nothing else. No code, no field, no gate, and the process definition is untouched.

<!-- xeno:section:deviations -->
## Deviations from the design

**One correction, caught before the commit was pushed.** The first draft of the WP8
paragraph said the digest "is read at P5", and that is not what the process definition
says. What it says is "Later phases read the earlier phase's output and digest rather than
searching the codebase again", which is a different and better-grounded reader. The
intake's context-rationale carries the original claim and is sealed, so the correction is
here and in the paragraph that shipped.

It matters because the resolution rests on the digest having a reader at all. "Read at P5"
would have put a clause in a normative document on a property I had not verified — which
is the failure this intent exists to fix, reproduced in the fix.

**The replacement is longer than what it replaced**, four sentences against one. The
project's convention is to replace rather than edit into a paragraph and to read it back as
prose, and read back as prose it needs the reason: WP8 is where somebody meets this clause,
and a pointer to section 5 is what let the contradiction survive.

Nothing else deviates. The maintainer chose resolution B of three and that is what was
carried out.
