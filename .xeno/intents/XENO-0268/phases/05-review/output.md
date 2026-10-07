---
intent: github.com/triplem/xeno#238
phase: 05-review
created: "2026-10-07T08:54:33Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b30ec554c8bd117f43b88fdbecb636413d5116f9f13efe5ab44b8d8b00d7c468
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'The rule applies to 03-implementation and that phase recorded two deviations, both of which name what they depart from. The register row is A99 rather than A98 and names P1''s criterion 11, which asked for one row without saying which number; A98 exists on the unmerged branch for #228 and a row''s number is cited from other rows, so a collision after both merge is not renameable. The intent key was passed explicitly rather than derived and names no criterion, because it departs from nothing this intent wrote: xeno intent start derived XENO-0267, the key the #228 branch already holds, the derived intent was removed before any phase started, and the finding is in P0''s learning.yaml. Marked not-applicable rather than met because the rule is scoped to the implementation phase and this is the review''s answer about it; both deviations are traceable and neither is a note.'
      result: not-applicable
      rule: deviations-are-traceable
    - note: 'No interface changes. Nothing in the runner, the CLI, the gates, the artifacts, the templates or the skills moves: the diff against main is two documents and no Go file, and xeno gate verify reports 481 verdicts verified. What an adopter meets is two paragraphs of specification and one entry in the plan. The MCP surface is the same six operations, so nothing that calls it changes either. There is nothing to migrate.'
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: 'None added. go.mod and go.sum are untouched and nothing compiles differently. The decision is about a protocol primitive that is not adopted, so not even a prospective dependency is taken on: the MCP prompts primitive is read from its specification and declined, and the plugin gains no file.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** Two normative documents are changed, which the first standing
rule bars the agent from doing. The exception is the one the maintainer made on instruction:
the three options were put with their consequences and costs, the choice was theirs, the
wording of all three passages was then drafted and put to them, and "write them as drafted" is
what was acted on. Both decisions were put one at a time, which section 8 asks for and
XENO-0243 is the counter-example to. The specification change is its own commit and comes
first. Nothing is invented: no field, no gate, no rule, and the MCP surface is the same six
operations. The change belongs to WP11 and to intent XENO-0268 for issue #238.

**What #238 asked for, and what it got.** Either the phases as person-invoked commands through
a mechanism both clients support, with three questions answered in the plan's section 9, or the
decision not to recorded with its reason and the model invocation stated as deliberate for
phases as well as lenses. The three questions are answered, the first branch turned out to be
unavailable because the answers say it is, and the second is written. Both halves of the issue
are closed, which is the test for closing it.

**The one thing a reviewer should be sceptical of.** The whole decision rests on a negative
about another project's software, read on one day from that project's issue tracker and
changelog. If Codex does surface MCP prompts somehow — behind a flag, in a release between the
check and the merge, documented somewhere neither source covers — then the first branch was
available and this intent took the wrong one. P4's gaps say this in those words. What limits
the damage is the wording: the paragraph states the absence as of October 2026, so a
correction deletes a sentence rather than arguing with a principle.

**The question no rule asks: is "until both clients can take them" the right condition?** It
makes the phases wait on the slower client, which is what interchangeability means and also
what it costs. The alternative, shipping for the faster one and accepting a difference, was
put to the maintainer with the plan's sentence about per-harness packaging in its favour and
was declined. Worth naming because the condition will be read as obvious once it is in the
document, and it was a choice.

**The question after that: does the budget paragraph claim too much?** It says a prompt's text
is fetched rather than sent with every request, which the protocol fixes, and then says what a
client puts in front of the model from the listing is the client's own choice. The second
sentence is what keeps the first honest. Without it the paragraph would read as "prompts are
free", and six phase prompts whose names and descriptions a client keeps resident are not free
— smaller than six tool definitions and not zero.

**What is deliberately not here.** No prompts, no `commands/`, no change to any skill, no gate
or field recording how a phase was invoked, and no scheduled re-check of another project's
changelog. P1's non-goals carry the reason for each. The last is the one worth repeating: a
check nobody runs is worse than a dated sentence that says when it was true.

**Two numbering deviations, both recorded.** The register row is A99 because A98 is on the
unmerged branch for #228, and a row's number is cited from other rows so a collision is not
renameable. The intent key was passed explicitly because `intent start` derived XENO-0267,
which that branch already holds; the derived intent was removed before any phase started. The
second is a finding about the runner and sits in P0's `learning.yaml` rather than being
absorbed.

<!-- xeno:section:release-notes -->
## Release notes

The six phases stay model invoked skills, and section 13 now says that is deliberate rather
than leaving a reader to work out whether it was a choice. A skill is loaded because the client
judged its description to fit; a command is typed. That is the right way round for a lens,
since whether a security lens applies is a judgement about the change, and it is not for a
phase, since the person knows which phase they are in.

What the phases are waiting for is named: MCP's prompts primitive, which is the protocol's own
form for something a person invokes and is not one client's packaging. Claude Code surfaces a
prompt as `/mcp__<server>__<prompt>`, discovered from the connected server. Codex, as of
October 2026, does not. A mechanism one client supports is the packaging problem under another
name, which is the thing skills exist here to avoid, so the phases stay skills until both
clients can take them. The paragraph is written so that the day that changes, the sentence is
false and has to go.

The MCP budget clause gains a paragraph saying what the budget does not cover. A tool
definition is sent with every request of every phase whether it is called or not, which is why
the surface is a budget; a prompt is not, because `prompts/list` carries only a name, a title,
a description and the arguments while the messages arrive at `prompts/get` when the prompt is
invoked. What a client keeps in context from the listing is the client's own choice and is not
fixed by the protocol, so the paragraph claims the saving the protocol carries and no more.

The implementation plan's section 9 carries all three answers with what each was read from,
including the issue number and date behind the Codex one, so that a reader rechecking has the
thread rather than the conclusion.

Nothing else changes. No skill, no template, no rule, no gate, no field, and the MCP surface
is the same six operations it was. `docs/assumptions.md` gains A99 with the decision, who took
it and when.

<!-- xeno:section:residual-risk -->
## Residual risk

**The decision rests on a dated negative about software this project does not control.** If
Codex surfaces MCP prompts and the check missed it, #238's first branch was available and this
intent took the second. The check was made on 2026-10-07 from that project's own issue tracker
and changelog, with a positive control in the other client to show the method finds the feature
where it exists, and the paragraph says October 2026 rather than asserting a permanent fact.
What remains is the risk no method covers: a feature behind an undocumented flag, or a release
between the check and the merge.

**Nothing watches for the condition changing.** The wording puts the correction in a diff when
somebody comes to make it, and nowhere when nobody does. There is no shape for a check here —
this repository cannot watch another project's changelog — so the honest alternative to a dated
sentence is no sentence, which is the state #238 found.

**The budget claim is half measured and says so.** The protocol fixes where a prompt's text
lives and not what a client keeps resident from the listing. If a client did hold every
prompt's name and description in every request, six phase prompts would be six standing costs
after all. Smaller than six tool definitions, not zero, and the paragraph's last sentence is
what stops it being read as zero.

**The decision is taken without the thing it is about.** The MCP server does not exist. That is
the right way round for deciding not to build something and it is still a decision made against
documentation rather than against a surface somebody has used. When WP11 builds the server the
question is worth re-reading, and nothing forces that either.

**The register is discontinuous on this branch.** A97 is followed by A99, because A98 is on the
unmerged branch for #228. If #228 merges the gap closes; if it is abandoned the register skips
a number permanently, which is visible as a gap rather than as two rows with one number.

**An intent key can be derived twice.** `xeno intent start` counts from what the branch can
see, so two branches open at once both derive the same key, and the key sits inside
`artifacts_hash` and inside the merge commit that names it. This intent avoided it by passing
`--intent`; the next one will only avoid it if somebody remembers. The finding is in P0's
`learning.yaml` and the lasting form is a refusal or a warning in `intent start`, which is a
runner change nobody has scheduled.

**For a person, not for the code.** Whether a reader of section 13 agrees that "the person
knows which phase they are in" is the whole difference. It is the premise the paragraph rests
on, it is stated and not argued, and somebody who works differently — several intents open,
phases picked up out of order — would read it as less obvious than it sounds.
