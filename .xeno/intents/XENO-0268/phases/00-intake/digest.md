---
intent: github.com/triplem/xeno#238
phase: 00-intake
created: "2026-10-07T08:47:31Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1825f0083f8f612b2ab88144bf5ce95f6c106f6ee5ef972cff6c2c642b58bb90
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The plugin ships seven model invoked skills and no commands, and #238's argument is that a
phase is the wrong thing to invoke by description: the person knows which phase they are in,
so "do the intake" is a command and not a judgement about the change the way a lens is.

All three of the issue's verification points are answered. Claude Code surfaces MCP prompts as
`/mcp__<server>__<prompt>`, discovered from the connected server. Codex does not, as of
October 2026: the request for it is closed as duplicate and nothing in the changelog
implements it. And a prompt is not charged per request as a tool definition is, because
`prompts/list` carries the name and the description while the messages arrive at
`prompts/get`.

So #238's first branch, a mechanism both clients support, is unavailable, and the second is
what is left. The maintainer was put the three options with their consequences and chose
recording the decision not to; the wording was then drafted, approved and written on
instruction.

In scope: two paragraphs in section 13, one entry in the plan's section 9, and one row in the
register carrying the dated fact. Out of scope: prompts for one client, `commands/`, any
change to the skills, anything else about the tool surface, and a scheduled re-check of
another project's changelog.
