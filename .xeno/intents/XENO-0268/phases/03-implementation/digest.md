---
intent: github.com/triplem/xeno#238
phase: 03-implementation
created: "2026-10-07T08:50:36Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2c5aecf86c229ba30b504a088d6422674119fdbe61ad1470f4ab79e793ccca32
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two paragraphs in section 13 of the process definition: one saying the phases' model
invocation is deliberate, with the skill-versus-command difference, why a lens is the other
way round, the prompts primitive and the client that lacks it, and the sentence that makes the
paragraph deletable; one beside the MCP budget clause saying a prompt is not a tool for that
purpose and naming what stays the client's own choice.

One entry in the plan's section 9 with the three questions and the three answers, each with
what was read for it.

One row in the register, A99, carrying the decision, the date Codex was checked, the option
not chosen with the argument in its favour, and what stays unmeasured. Approved by the
maintainer on 2026-10-07.

Both documents were committed first and on their own, with the exception to the first standing
rule named in that commit's message.

Two deviations, both numbering. The row is A99 rather than A98, because A98 is on the unmerged
#228 branch and a row's number is cited from other rows, so a collision is not renameable. And
the intent key was passed explicitly, because `intent start` derived XENO-0267, which that
branch already holds; the derived intent was removed before any phase started and the finding
is in P0's learning record.

No code, no skill, no gate. Build, tests, gofmt, vet and gate verify pass.
