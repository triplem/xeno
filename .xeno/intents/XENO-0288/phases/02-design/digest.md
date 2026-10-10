---
intent: github.com/triplem/xeno#334
phase: 02-design
created: "2026-10-10T15:43:14Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4a116381f0666bacf836f8b5b1182c8d9d4b7ca308030fc68a5fbf8ee99af393
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
The design of one section. Ten construction decisions, eight alternatives rejected, four
places in one file, and two assumptions marked as assumptions.

It is section 10 of the page and not a subsection of 9, because section 9 is about a
specification framework and this is about a methodology that compares with the process
itself. Six subsections in 9's order, with 10.4 holding #323's three shapes in the place
9.4's measurement holds, because the decision does not rest on the arithmetic and a reader
who meets the numbers first will think it does.

The third row is settled against the code rather than against either document. The guide
says a sensor result is advisory in this release, two lines below a paragraph that describes
gate-fired sensors; the reference says blocking is enforced for `fire_on: gate`. The
implementation agrees with the reference, so the row names the two functions and the hook
with their lines in an 8,520-line file, and states the three ways the override is refused,
because an enforcement claim without its refusal path is the flattering half.

One alternative was rejected for a reason worth carrying forward. Option two of the question
on #334 named its cost as blob OIDs not being stable across a rebase that changes a tree.
AI-DLC's own reference claims the opposite property for that mechanism — attribution "keys
on blob content (OIDs), not commit ancestry" — so a rebase that leaves trees alone leaves
every OID alone, and one that changes a tree reports `drifted`, which is the right answer.
D-1 does not move, because the reason that decides it is `docs/process-definition.md:480`
and not the rebase property. What the section writes is what the mechanism does, with :480
as the reason. Repeating a cost the source contradicts, inside an artifact that seals it, is
the thing `CLAUDE.md`'s negative-result rule exists to prevent in a different dress, and the
learning of this phase is the general form.

Two assumptions. That a pinned section is the right form although AI-DLC moves at a tag a
day and the pin is stale as a description of `main` within days — taken as right for the
reason #323 gives, with the obligation that every cell say where it was read. And that the
hosted sample compares on the same rows without a reading of its own, which is why the line
it gets is held to its README at the pinned commit and makes no claim about its gates, its
sensors or its approvals: incomplete rather than wrong is the trade a line is for.

No open question. Three sections of the four this template defines, plus the learning.
