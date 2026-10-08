---
intent: github.com/triplem/xeno#95
phase: 00-intake
created: "2026-10-08T14:40:29Z"
schema_version: "1.0"
runner_version: dev+5f12412
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2d0fe57bf21259ff40c12425220caa7dc1d0c58036e65a071721522135f5dd3e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake of #95 again, after XENO-0274 shipped a configuration that contradicts the answers
this issue carries.

The issue body proposes; its comments decide. The maintainer analysis names six kinds of
pinned thing and three constraints, and the answers given today are self-hosted, no automerge,
dashboard plus pull requests. XENO-0274 read the body alone — which is all `phase start` carries
— and got two of the three wrong, missed the DCO entirely, and left out the `customManagers`
that reach the versions written inside workflow text, which is most of what this project pins.

That is this phase learning record: a phase start that quotes the body and nothing else lets a
phase proceed believing the body is the whole of what was said.

One open question is recorded rather than assumed: whether a bot may sign off under the DCO.
G-Questions will not let P5 be decided until it is answered, which is the mechanism the first
intent did not use.

Two sections and one question, of the five this template defines.
