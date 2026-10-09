---
intent: github.com/triplem/xeno#95
phase: 05-review
created: "2026-10-09T16:56:29Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f1a1f1398f98da435598ac7cd90046900a81fdee066520e302ad20de942ee297
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
The review of a change that makes three pins visible and moves none of them.

Three rules answered, two not-applicable and one met. No interface moves and no dependency is
added: what looks like a dependency change is the same debian digest with the tag the
supply-chain page already called it by. The supply chain lens is recorded as a deviation
rather than as met, for the reason XENO-0277 recorded the same one: A28 pins to what a release
demonstrably used and a bot proposes the newest, and three more pins now have a proposer. One
notch wider than before, most sharply for the image's base, where the tag names a stream
renovate will offer newer digests of. Accepted on the same ground and named rather than
resolved.

The residual risk worth reading twice is that the datasources are unexercised. The regexes
were run in renovate's dialect by a script that is not renovate, so every claim here is about
what the configuration says and none is about what renovate does with it. A wrong
`extractVersionTemplate` would be quiet in exactly the way this intent exists to fix: gitleaks
compared against a version that does not exist would never propose, which looks like a pin
that has not moved.

Two sections of the five, three rules and one lens entry, no question, no decision.
