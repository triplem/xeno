---
intent: github.com/triplem/xeno#254
phase: 00-intake
created: "2026-10-05T18:56:24Z"
schema_version: "1.0"
runner_version: dev+e471bbb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 33b31309b5f15745460e2d648e32e31e19821e51502a2fe58c3ca17772107ddc
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The audit's one row about questions describes a third of section 8's sentence and names a gate
that does not read it: the shape check runs through G-Schema from P0, which is the fact #229
turned on, and the row says G-Questions.

#229 made it wrong without touching it, three days after the pass was made. #212 found it while
adding two rows of its own and left it, which is why this is its own intent.
