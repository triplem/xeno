---
intent: github.com/triplem/xeno#238
phase: 01-requirements
created: "2026-10-07T08:47:48Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5d074f8a5e5d69d29d375adc7da6b9d26ddd128cf66163ddaac2c30d10ebc5bf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Numbered, and P4's mapping cites the numbers.

1. **Section 13 says the phases' model invocation is deliberate**, in the place a reader
   learns that six of the skills are the phases, so somebody asking #238's question finds the
   answer where the question occurs to them.

2. **It states the difference it rests on**: a skill is loaded because the client judged its
   description to fit, a command is typed. Without that sentence the paragraph asserts a
   preference.

3. **It says why the same arrangement is right for a lens and not for a phase.** Whether a
   security lens applies is a judgement about the change; which phase somebody is in is not.

4. **It names what the phases are waiting for**, with the mechanism and the client that lacks
   it, so the decision reads as dated rather than as a principle.

5. **It is worded as a fact a change deletes**, not as a permission a change satisfies. The
   day Codex ships prompts the sentence is false, and a deleted sentence is visible in a diff
   where an unchanged permissive one is not.

6. **The budget paragraph says a prompt is not a tool for that purpose, and why.**
   `prompts/list` carries the name and the description, the messages arrive at `prompts/get`.
   #238 asked for exactly this, because the budget sentence otherwise leaves a reader to infer
   that prompts are tools.

7. **It says what stays the client's choice.** What a client puts in front of the model from
   the listing is not fixed by the protocol, so the paragraph does not claim more than was
   measured.

8. **The plan's section 9 carries all three answers**, in the `--> answer` style the other
   verification points use, each with what was read for it, including the dates and the issue
   number that makes the Codex answer checkable.

9. **The six operations stay six.** Nothing is added to or removed from the MCP surface, and a
   seventh still needs an argument of the kind the index has.

10. **No skill, template, rule or line of Go changes.** `go test ./...`, `gofmt`, `go vet` and
    `xeno gate verify` pass, and the verdict count is the trail's as it stands.

11. **`docs/assumptions.md` carries one row** with the decision, the date the Codex fact was
    checked, what was read for it, and what reopens the question.

12. **The specification change is its own commit, before anything else**, and the exception to
    the first standing rule is named in its message.

<!-- xeno:section:non-goals -->
## Non goals

**No prompts, in either client.** Not for Claude Code alone, which was put as an option and
not chosen, and not as a mechanism gated on which client is connected, which would be the
runner behaving differently per harness that section 7 forbids. The MCP server does not exist
either, so there is nothing to add them to.

**No `commands/` directory in the plugin.** That is Claude Code's format. #238 rules it out
under its own heading, and section 13 gives the same reason for the lenses being skills.

**No change to a skill's text or description.** The decision is about who invokes a phase. A
skill whose description reads badly is a different finding and nobody has made it.

**No change to the tool surface.** The budget paragraph gains a statement about what the
budget does not cover. It gains no operation, loses none, and the sentence about a seventh
needing an argument stands untouched.

**No gate, no field, no rule.** Nothing in an artifact records how a phase was invoked, and
nothing could: the runner is the same whether a person typed the command or an agent did,
which is what WP11's "artifacts a gate cannot tell apart" requires.

**No scheduled re-check of Codex.** The answer is dated in the document and in the register
row, which is the honest form. A check that watches another project's changelog is not
something this repository can run, and one that nobody runs is worse than a sentence that says
when it was true.

**No reopening of the decision inside this intent.** The three options were put once, with
their consequences and costs, and answered. Section 8's rule is one decision at a time; asking
again under a different heading is the batch it forbids with the letter satisfied.

**No new issue for the MCP server.** WP11 already owns it and the plan already lists it as not
built. A second record of the same work would be the kind of addition the third standing rule
asks to be written down rather than absorbed, and there is nothing here that no package owns.

<!-- xeno:section:constraints -->
## Constraints

**The first standing rule, and the exception.** The documents are not editable by the agent.
Both documents changed here, under an instruction: the three options were put to the
maintainer with their consequences and costs on 2026-10-07, the choice was theirs, the wording
of all three passages was then drafted and put to them, and "write them as drafted" is what
was acted on. The exception is named in the specification commit's message, because a trace is
what stops it becoming a habit.

**The specification commit comes first**, and here it is the only commit that carries a
document change; the register row and the artifacts follow in the second.

**A negative result is evidence only when the thing checked was there to be found.** The Codex
answer is a negative, and it is the load bearing one. It is not taken from a single search: the
issue that asks for the feature was read and its state confirmed, the changelog was read for
what it does say about MCP, and the positive case was established first in the other client so
that the method is known to find the thing when it is there. The sentence in the document says
October 2026 rather than asserting a permanent fact.

**No invented fields, gates, tools or rules.** Nothing is added to an artifact, to the gate
list or to the MCP surface.

**Prose wraps at 88 characters**, tables and code blocks do not, and a paragraph changed a
second time is replaced rather than edited into. Both new paragraphs were rewrapped as whole
paragraphs after the first attempt left two lines over the width, and were read back as
paragraphs rather than as a diff.

**Headings name their section in words.** The new paragraph in section 13 is a bold lead-in in
the register that section uses, not a numbered subsection.

**One dependency.** Untouched; nothing compiles differently and no Go file is in the diff.

**Every change belongs to a work package and to an intent.** WP11, intent XENO-0268 for issue
#238. The key was passed explicitly, because the sequence derives from the tree and this
branch cannot see XENO-0267 on the unmerged branch for #228; that is recorded as a learning in
P0 rather than absorbed.
