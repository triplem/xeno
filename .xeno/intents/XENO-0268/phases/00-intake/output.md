---
intent: github.com/triplem/xeno#238
phase: 00-intake
created: "2026-10-07T08:46:49Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1825f0083f8f612b2ab88144bf5ce95f6c106f6ee5ef972cff6c2c642b58bb90
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

The plugin ships seven skills under `.xeno/plugin/skills/` and no commands. #238's argument is
that a phase is the wrong thing to invoke by description: a skill is model invoked, loaded
because the client judged its description to fit, and a command is person invoked, typed by
somebody. The six phase process is driven phase by phase by a person who knows which phase
they are in, and `xeno phase start` already names the next one. So what the person wants to
say is "do the intake", and today that is either a sentence they hope matches a skill's
description or seven CLI invocations they type themselves.

That is an argument the skills do not answer rather than a gap in them. A lens is correctly
model invoked, because whether a security lens applies is a judgement about the change. A
phase is not.

## What the issue asked to verify, and what the answers are

#238 named three things to verify in the practice of the plan's section 9, each before the
code that rests on it, and said that none of it was blocked on a decision because the MCP
server does not exist yet. All three are now answered.

**Claude Code surfaces MCP prompts as commands.** As `/mcp__<server>__<prompt>`, discovered
from the connected server rather than declared, and resolved alongside skills and
`.claude/commands/` files where a short name collides. The protocol's own prompts page says
prompts are "user-controlled" and "typically ... triggered through user-initiated commands in
the user interface", with slash commands as its illustration, so the mechanism is what #238
took it to be.

**Codex does not.** `/mcp` inspects tools and resources. The request to expose MCP prompts as
slash commands, openai/codex#8342 of 2025-12-19, is closed as duplicate, and nothing in the
Codex changelog implements `prompts/list` or `prompts/get` as commands — what it carries for
MCP through 2026 is the 2026-07-28 protocol revision, paginated tool discovery, the removal of
`codex mcp-server`, and request verification. The answer is dated: it is a statement about
October 2026 and not about the protocol.

**A prompt is not charged per request the way a tool definition is.** `prompts/list` carries a
name, a title, a description and the arguments; the messages arrive at `prompts/get` when the
prompt is invoked. So the budget's reason — that every definition is sent with every request
of every phase whether it is called or not — does not reach a prompt. What a client puts in
front of the model from the listing is the client's own choice and is not fixed by the
protocol, which is the part that stays unverified for any particular client.

## What that leaves

#238's first branch requires "a mechanism both clients support". There is none. MCP prompts
are the only candidate that is not one client's packaging, and they work in one client. So the
first branch is unavailable, and the second is what is left: the decision not to, recorded
with its reason, and the skills' model invocation stated as deliberate for phases as well as
for lenses.

The maintainer was put the three options with their consequences on 2026-10-07 and chose that
one. The wording of all three passages was then drafted, put to them, and written on
instruction, which is the exception to the first standing rule and is named in the
specification commit.

<!-- xeno:section:scope -->
## Scope

In scope are three passages, drafted and approved before any of them was written.

A paragraph in section 13 of `docs/process-definition.md`, after the one that says six of the
skills are the phases, saying that their model invocation is deliberate, what the difference
between a skill and a command is, why it is the right way round for a lens and not for a
phase, and what the phases are waiting for. It is worded as a fact a change deletes rather
than a permission a change satisfies, so that the day Codex ships prompts the sentence is
false and has to go.

A paragraph appended to section 13's MCP budget clause, saying that a prompt is not a tool for
that purpose and why: `prompts/list` carries the name and the description, the messages arrive
at `prompts/get`. #238 asked for this explicitly, on the ground that the budget sentence
otherwise leaves a reader to infer that prompts are tools.

An entry in section 9 of `docs/implementation-plan.md`, in the `--> answer` style the other
verification points use, carrying all three answers with what was read for each.

In scope is one row in `docs/assumptions.md`, because the decision outlives this intent and
the fact behind it is dated: a reader in a year needs to know that Codex was checked in
October 2026 and what would reopen the question.

## Out of scope

Out of scope is building the prompts for Claude Code alone. It was one of the three options
put to the maintainer, with the plan's own sentence in its favour — which of content,
packaging and wiring WP11 builds per harness is "settled in dogfooding rather than in
advance" — and it was not chosen. It also waits on the MCP server, which does not exist.

Out of scope is `commands/` in Agent Plugins format. That is one client's packaging and would
need a second mechanism for the other, which is the thing section 13 gives as the reason the
lenses are skills rather than subagents. #238 rules it out in its own second heading.

Out of scope is any change to the skills themselves. Their descriptions are what a client
matches, and nothing here argues they match badly; the argument is about who invokes a phase,
not about how well the description reads.

Out of scope is the sixth and seventh MCP operation and anything else about the tool surface.
The budget paragraph gains a statement about prompts and loses nothing: the six operations
stay six, and a seventh still needs an argument of the kind the index has.

Out of scope is re-checking Codex on a schedule. Nothing in this repository can watch another
project's changelog, and a check that nobody runs is worse than a dated sentence that says
when it was true.

<!-- xeno:section:context-rationale -->
## Why this context

Six files, 360034 bytes.

`docs/process-definition.md` carries both passages that land in it: section 13, where the
skills are enumerated and where the MCP surface is called a budget, and the lenses' stated
reason for being skills rather than subagents, "so they work in both clients", which is the
clause the whole decision turns on.

`docs/implementation-plan.md` carries section 9, where the answers go, in a style it is worth
matching rather than inventing; WP11, whose scope this is and whose "done when" asks for a
phase that can be driven from either agent with no agent specific logic in the runner; and the
sentence leaving per-harness packaging to dogfooding, which is the argument for the option
that was not chosen and had to be put with the others.

`docs/assumptions.md` is where the row lands, read first for the test a row has to meet.

`.xeno/plugin/skills/xeno-intake/SKILL.md` is one of the seven skills the issue is about, read
to see what a phase skill actually says and to confirm that nothing in it is being changed:
its description is the thing a client matches, and the decision is about who invokes it.

`.xeno/plugin/.claude-plugin/plugin.json` is read to confirm what the plugin declares and does
not: no `commands/`, which is the state #238 describes in its first sentence.

`CLAUDE.md` carries the standing rules this intent runs under: the first, which the documents
were written under an instruction against; and the one about a decision put to a person being
put one at a time, which is why the three options were one question and the wording a second.

The links block declares `.xeno/plugin/skills` against the process definition, because what
the skills are for is stated there and reading the tree alone would make the question look
like a packaging choice rather than the interchangeability one it is.

Not in scope: the MCP server, which does not exist, and `internal/plugin`, because nothing in
the runner changes. The prompts primitive was read from the protocol specification and the two
clients' documentation rather than from this tree, which has neither.
