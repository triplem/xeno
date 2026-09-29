---
intent: github.com/triplem/xeno#136
phase: 05-review
created: "2026-09-29T18:10:38Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4b87535.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bd3b84de73ddf17db5daaa8fd704fc278f0e5e34a1e54be64f551ba20161b587
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Five residual risks for a change of case, and the one to carry is that the convention now lives in a comment
in cmd/xeno rather than in CLAUDE.md, where the citation that caused the problem still sits undistinguished
from prose. The other is structural rather than about this change: three intents are queued on one another
now, each asked for while the last was open, and they merge in order.
