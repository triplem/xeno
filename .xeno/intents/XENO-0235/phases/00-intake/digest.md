---
intent: github.com/triplem/xeno#199
phase: 00-intake
created: "2026-10-03T18:22:00Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7ca8bac9f0250f071c334a8e74051de9fd2160eb2954227c1096ec8b7e465d98
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
G-Supply shipped one change ago anchored to a plugin a release cannot deliver: the digest is over
`.xeno/plugin/`, that directory comes from `init --vendor`, and vendoring read a path on disk, so a
released binary in a fresh repository had nothing to copy. The gate was correct and unsatisfiable.
And what vendoring copied was not the plugin — it named its parts, and two things were in the tree
and in none of the calls: `secrets.yaml`, which the function's own comment lists as part of section
13's `--vendor` set, and `bin/xeno-env.sh`, the entry point the plugin's own `hooks.json` names.
Thirty-two files vendored where the plugin has thirty-four, and either omission alone makes the
digest unmatchable, so G-Supply would have failed for every adopter on every phase. Nothing caught it
because the tests about vendoring listed the parts the code listed. The manifest is the other half of
the same mechanism: section 13 asks for one shared number, the literal was bumped by hand three times
in one session and overtaken by a release twice, and a release cannot commit it back to main because
#121 removed the plugin that would. What makes both solvable at once is that a release can write to
what it ships — the copy an adopter receives is the only copy the coupling has to hold for and the
same copy the digest must be taken over. In scope: embed, stamp, digest in that order; a walk instead
of a list; `bin/` executable; `0.0.0-dev` in the repository; and a test comparing digests rather than
listing files. Out of scope: a check that the manifest matches the tag, which was offered and
declined; the resolution order, still section 7's; a signature, which section 13 disclaims; and
confirming a published binary carries the tag's digest.
