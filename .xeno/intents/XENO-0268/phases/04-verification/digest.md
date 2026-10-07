---
intent: github.com/triplem/xeno#238
phase: 04-verification
created: "2026-10-07T08:54:11Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 94f1a8e73955042401ee08cb6dcd08f8e2a476de6e0836d77cc9f810b893420e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Twelve criteria, all passing: eight read by a person, because the deliverable is prose and a
check that a paragraph exists says nothing about whether it says the thing, and four checks.

The three answers are in `evidence/verification.txt` with the page each came from. Claude Code
surfaces a prompt as `/mcp__<server>__<prompt>`, discovered from the connected server, and the
protocol agrees about the shape. Codex does not, checked 2026-10-07: `/mcp` inspects tools and
resources, openai/codex#8342 of 2025-12-19 is closed as duplicate, and the changelog's MCP
work through 2026 is elsewhere. And `prompts/list` returns the name and the description while
`prompts/get` returns the messages, so a prompt's text is not sent with every request.

The negative was checked the way the convention asks: the same question asked of the client
that has the feature returns a documented mechanism, so the method finds it where it exists,
and the issue's own thread was read for its state.

`go test ./...` passes across 20 packages, `internal/secrets` at 132.5 s. Build, `gofmt`,
`go vet` clean. `xeno gate verify` matches 481 verdicts. The diff against main is two
documents and no Go file.

The gaps are the honest half. The load bearing fact is a dated negative about somebody else's
software and nothing here would notice if it changed; the third answer is half measured,
because the protocol does not fix what a client keeps in context from the listing; eight
criteria are a person reading prose; the decision is taken without the MCP server it is about;
and A98 is missing from the register on this branch.
