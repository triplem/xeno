---
intent: github.com/triplem/xeno#238
phase: 02-design
created: "2026-10-07T08:48:42Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2c83e90d151443825741d512fefaf78d29e4a958c37662227121ba0ceb8cd7ec
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - chosen: 'Record the decision not to expose the phases as person-invoked commands, rather than building MCP prompts for the one client that supports them: a paragraph in section 13 saying the skills'' model invocation is deliberate for phases as well as lenses, a paragraph beside the MCP budget clause saying a prompt is not a tool for that purpose, and all three of the issue''s verification answers in the plan''s section 9.'
      decided_by: Markus M. May
      id: D-1
      proposed_by: claude-opus-5
      rationale: 'The issue''s first branch requires a mechanism both clients support and there is none. MCP prompts are the only candidate that is not one client''s packaging, and the verification found them in one client only: Claude Code surfaces a prompt as /mcp__server__prompt, discovered from the connected server, while Codex as of October 2026 does not — openai/codex#8342 of 2025-12-19 is closed as duplicate and nothing in the changelog implements prompts/list or prompts/get as commands. A mechanism one client supports is the commands/ problem under another name, which is the reason section 13 gives for the lenses being skills rather than subagents. The second option, prompts for Claude Code alone, has the plan''s sentence leaving per-harness packaging to dogfooding in its favour and was put with that argument; it was not chosen, and it waits on an MCP server that does not exist. The third, answering the three questions and leaving the issue open, was put with the cost that the next reader has the facts and not the decision. The third verification answer is recorded whichever way the first two went, because the budget sentence otherwise leaves a reader to infer that prompts are tools: prompts/list carries a name, a title, a description and the arguments, and the messages arrive at prompts/get on invocation.'
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**Where the paragraph about model invocation goes.** Immediately after the one that says six
of the skills are the phases and that `xeno-learning` is not a seventh. That paragraph is
where a reader learns the phases are skills, so it is where the question "why is a phase
invoked by description" occurs to them. The alternative was the **Roles** paragraph two
screens down, which is about what the skills carry rather than how they are reached.

**The paragraph is worded as a fact a change deletes.** "Codex, as of October 2026, does not"
and "the phases stay skills until both clients can take them", followed by the sentence saying
the paragraph is what a change to that deletes. XENO-0266 reached the same convention for a
specification clause recording a consequence rather than requiring something of a tool: a
deleted sentence is a decision visible in a diff, an unchanged permissive one is a silence.
The alternative wording — that phases may be exposed as prompts once both clients support them
— reads as a permission that a future change satisfies while leaving the current claim
standing.

**The budget statement is a paragraph beside the budget clause, not a clause inside it.** The
existing paragraph is one argument: six operations, a definition sent with every request, a
seventh needs an argument. Inserting a sentence about prompts into it would make the reader
hold two subjects at once, and the paragraph has already been written once for the reason it
gives. Beside it, the new paragraph answers the one question a reader of the budget has about
prompts and stops.

**It names what was not measured.** "What a client puts in front of the model from the listing
is the client's own choice and is not fixed by the protocol." The protocol fixes where the text
lives; it does not fix whether a client keeps a prompt's name and description in context. That
half is unverified for any particular client, and a paragraph claiming the whole saving would
be the stronger claim the evidence does not carry.

**All three answers go in the plan's section 9 and nowhere else.** #238 asked for them there.
The entries in that section are a question and an answer after `-->`, so the three questions
are one entry rather than three, because they were one verification exercise and the third
answer only matters if the first two have the shape they do.

**The Codex answer carries its issue number and date.** openai/codex#8342, opened 2025-12-19,
closed as duplicate. A reader who wants to recheck needs the thread, not the conclusion, and
the conclusion is the half that will go stale.

**One row in `docs/assumptions.md`, not a decision entry in this phase alone.** The fact is
dated and the decision outlives the intent: the next person reaching for prompts needs to know
that Codex was checked, when, and what reopens it. A decision sealed only in this P2 would be
read by somebody who already knew to look here.

<!-- xeno:section:alternatives -->
## Alternatives

**MCP prompts for both clients, which is what #238 wanted.** The right answer if it existed.
Prompts are MCP's own primitive for a person-invoked operation, they are not one client's
packaging, and they cost nothing against the tool budget. What they lack is Codex: `/mcp`
there inspects tools and resources, the request to expose prompts as slash commands is closed
as duplicate, and the changelog's MCP work through 2026 is elsewhere. Unavailable, and the
paragraph is worded so that the day it becomes available the sentence has to go.

**MCP prompts for Claude Code alone.** Put to the maintainer as the second option, with the
plan's own sentence in its favour: which of content, packaging and wiring WP11 builds per
harness is "settled in dogfooding rather than in advance", so a per-harness entry point is not
forbidden. Two things cost it. A person on Codex types `/mcp__xeno__intake` and gets nothing,
which is the `commands/` problem #238 named under another name rather than a different one.
And it waits on the MCP server, which does not exist, so choosing it would have left #238 open
with the work parked. Not chosen.

**`commands/` in Agent Plugins format.** The shortest path, and the one the plugin's own format
offers. #238 rejects it in its own second heading and section 13 supplies the reason it
rejects: the lenses are skills rather than subagents so that they work in both clients, and a
`commands/` directory is one client's packaging that would need a second mechanism for the
other. It also sits inside the vendored tree that G-Supply hashes whole, so the one-client
artifact travels into every project.

**Answer the three questions and leave #238 open.** The third option put. It gets the facts
written down where they can go without touching a normative document, and it costs the
decision: the next reader finds that Codex was checked and not what was concluded, and the
issue stays open with one of its three parts done. Not chosen, and the facts went into the
plan rather than only the register because that is where #238 asked for them.

**A gate or a field recording how a phase was invoked.** Never put as an option, and worth
writing down as refused rather than unconsidered. WP11's "done when" requires artifacts a gate
cannot tell apart from the ones an agent produced, which is the opposite of recording the
invocation, and section 7 forbids the runner behaving differently per harness. There is no
shape for this that the process allows.

**Say nothing.** The state #238 found: seven skills, no commands, and a reader left to work
out whether model invocation for a phase was a choice or an oversight. It was a choice, and
nothing said so.

<!-- xeno:section:impact -->
## Impact

`docs/process-definition.md`, section 13: two paragraphs, 18 lines. One after the enumeration
of the skills, one beside the MCP budget clause. Nothing is deleted and no existing sentence is
reworded.

`docs/implementation-plan.md`, section 9: one entry, 10 lines, carrying three questions and
three answers in the style the other verification points use.

`docs/assumptions.md`: one row, with the decision, the date the Codex fact was checked, what
was read for it and what reopens it.

Nothing else. No Go file, no skill, no template, no rule, no gate, and the MCP surface is the
same six operations it was.

**For a reader of the documents.** The question #238 asks is answered where it occurs to them,
with its reason and its date, and the budget clause no longer leaves prompts to inference. A
reader who wants prompts now learns in one paragraph why there are none and what would change
it.

**For the trail.** Nothing moves. `xeno gate verify` recomputes every sealed verdict and
matches, because the documents are not inside anybody's `artifacts_hash` — they are read by
context scopes, and the locks that recorded a hash of the process definition belong to phases
that are finished and are compared against what they recorded rather than against today's
tree.

**For WP11.** One of its two open verification points closes, and the "done when" is not
affected: a phase can still be driven from either agent with no agent specific logic in the
runner, which is what the skills already do. What changes is that the entry point question has
an answer instead of being open by omission.

**What is left undone, deliberately.** The MCP server. The plan lists it as not built and WP11
owns it; this intent decides what it will not carry rather than building any of it.
