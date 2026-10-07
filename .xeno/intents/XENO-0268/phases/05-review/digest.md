---
intent: github.com/triplem/xeno#238
phase: 05-review
created: "2026-10-07T08:55:25Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b30ec554c8bd117f43b88fdbecb636413d5116f9f13efe5ab44b8d8b00d7c468
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three standing rules hold, with the exception the maintainer made on instruction: the three
options were put with their consequences, the choice was theirs, the wording was then drafted
and put to them separately, and both decisions were put one at a time as section 8 asks. The
specification change is its own commit and comes first.

Both halves of #238 are closed. The three verification questions are answered in the plan's
section 9, and the answers are what makes the first branch unavailable: there is no mechanism
both clients support. The second branch is written, in two paragraphs of section 13.

The one thing to be sceptical of is named: the whole decision rests on a negative about
another project's software, read on one day. The wording limits the damage by stating the
absence as of October 2026, so a correction deletes a sentence rather than arguing with a
principle.

Two questions no rule asks are answered. Whether "until both clients can take them" is the
right condition — it makes the phases wait on the slower client, which is what
interchangeability means and also what it costs, and the alternative was put and declined. And
whether the budget paragraph claims too much — its last sentence, about what stays the
client's own choice, is what keeps the first honest.

Three rules answered, all not-applicable: the traceability rule is scoped to P3, whose two
deviations both name what they depart from; no interface changes, since the diff is two
documents and no Go file; no dependency, not even a prospective one, since the primitive is
read and declined.
