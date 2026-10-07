---
intent: github.com/triplem/xeno#238
phase: 03-implementation
created: "2026-10-07T08:50:26Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2c5aecf86c229ba30b504a088d6422674119fdbe61ad1470f4ab79e793ccca32
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`docs/process-definition.md`, section 13, two paragraphs.

After the paragraph that says six of the skills are the phases: **The phases are model
invoked, and that is deliberate.** It states the difference it rests on, a skill loaded
because the client judged its description to fit against a command that is typed; says the
arrangement is right for a lens and not for a phase, with the reason for each; names MCP's
prompts primitive as the form that would carry it and the second client it lacks, with Claude
Code's `/mcp__server__prompt` and Codex as of October 2026; and ends by saying that a
mechanism one client supports is the `commands/` problem under another name, so the phases
stay skills until both clients can take them, and that the paragraph is what a change to that
deletes.

Beside the MCP budget clause: a prompt is not a tool for that purpose and the budget's reason
does not reach one, because `prompts/list` carries a name, a title, a description and the
arguments while the messages arrive at `prompts/get` when the prompt is invoked. The last
sentence says what stays the client's own choice and is not fixed by the protocol, which is
the half that was not measured.

`docs/implementation-plan.md`, section 9, one entry. The three questions, then after `-->` the
three answers: Claude Code yes with the name and how it is discovered; Codex no as of October
2026, with openai/codex#8342 of 2025-12-19 closed as duplicate and what the changelog does
carry; and no, a prompt is not charged per request, with where the text arrives and what the
client decides.

`docs/assumptions.md`, one row. A99 carries the decision, what was read for each of the three
answers, the date Codex was checked, the option that was not chosen with the argument that was
in its favour, and what stays unmeasured. Its state is `approved (#238)` and names who decided
and when, because this one was put to a person and answered rather than read off a document.

Both documents were committed first, on their own, as `CLAUDE.md` requires, and the exception
to the first standing rule is named in that commit's message. The register row and the
artifacts follow in the second commit.

Nothing else is touched. No Go file, no skill, no template, no rule. `go test ./...`, `gofmt`,
`go vet` and `xeno gate verify` all pass, and the verdict count is the trail's as it stands.

<!-- xeno:section:deviations -->
## Deviations from the design

Two, and both are about numbering rather than about anything the documents say.

**The register row is A99 and not A98.** P1's criterion 11 asked for one row and did not say
which number. A98 exists, on the unmerged branch for #228, and the sequence here is the next
number of what this branch can see, which is A97. Taking A98 would have produced two rows with
one number the moment both merge, and a row's number is cited from other rows — A95 cites A77,
A97 cites A89 — so a collision is not renameable afterwards. A99 leaves A98 to the branch that
already wrote it. If #228 is abandoned the register skips a number, which is the cheaper of the
two outcomes and visible as a gap rather than as two A98s.

**The intent key was passed explicitly.** `xeno intent start` without `--intent` derived
XENO-0267, which is the key the #228 branch already holds; the sequence is the next number of
what the repository tree holds and a branch cut from main cannot see an unmerged intent. The
derived intent was removed and XENO-0268 created in its place, before any phase started, so
nothing was sealed under the wrong key. P0's `learning.yaml` carries this: the lasting form is
for `intent start` to refuse or warn when the derived key exists on another branch, which is a
runner change and a finding about the plan rather than a step in this intent.

Nothing else departs from P2. Both paragraphs are where the design put them, the plan's entry
is one entry rather than three, and the row says what stays unmeasured.
