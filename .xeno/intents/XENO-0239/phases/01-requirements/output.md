---
intent: github.com/triplem/xeno#213
phase: 01-requirements
created: "2026-10-03T20:04:35Z"
schema_version: "1.0"
runner_version: dev+427eeb7
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3e149f289fe109858057709904d94fa3fab095ad1eaf733f09fd83773aae2ce5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

WP8 no longer asks for anything section 5 forbids. Its paragraph and its done-when put the
record of an out-of-profile read in the phase's digest, and say in the plan itself why the
lock cannot carry it, so that the next reader does not reopen the question.

The reason is stated where the clause is, not left as a pointer: a reader of WP8 learns
that the lock is written before the agent starts and sealed, and that a lock rewritten at
the end would describe nothing.

The record is named as what it is. The plan says it is the agent's own account, because
nothing in the harness reports what was opened. A clause that oversold its own strength
would be worse than the contradiction it replaced.

`ASSUMPTIONS.md` records the resolution with both alternatives and why they lost, so that
adding the field to section 5 is a decision somebody takes again rather than an idea that
looks new.

The specification change is its own commit and precedes this intent's work, which is the
first standing rule. Section 5 is untouched.

<!-- xeno:section:non-goals -->
## Non goals

No change to the process definition. It wins by the first standing rule, and here it also
carries the better reasoning.

No gate, and no field. Nothing judges the digest's prose, and a check that an out-of-profile
read had been declared would have to know what was read, which is the thing nothing reports.
A reader that cannot fail is worse than none, which is A90.

No code. `ContextLock` already matches section 5 and needs no change; what changed is a
completion condition the code never implemented.

Not WP8's first half. "A repeated phase reads only what changed" is still unmet and gets its
own issue now that the contradiction is settled.

<!-- xeno:section:constraints -->
## Constraints

The documents are not editable by the agent, so the resolution was put to the maintainer as
three options with their consequences, and the one chosen is the one carried out. The other
two are recorded rather than discarded.

A plan paragraph changed for the second time is replaced rather than edited into, and read
back as a paragraph. This is that case, and the replacement runs longer than what it
replaced because the reason now travels with the clause.

Prose at 88, tables exempt — which is now what `CLAUDE.md` says, since #210.

The digest's properties are the whole load-bearing part: written at `phase finish`, inside
`artifacts_hash`, outside `context_hash`, and read by the phases that follow. If any of
those were false the resolution would not hold, so they were read rather than recalled.
