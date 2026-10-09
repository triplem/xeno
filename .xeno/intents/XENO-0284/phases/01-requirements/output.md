---
intent: github.com/triplem/xeno#337
phase: 01-requirements
created: "2026-10-09T15:54:20Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7cba487d1d1ab66db1c29849c1e3bed4a78a4753bda721232f1f88222c38c070
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

1. **The plan precedes the trail.** `ec372ea` is the first commit after `main` on this
   branch, touches `docs/implementation-plan.md` alone, and its message records the
   exception to the first standing rule.
2. **The plan says the entry point is in scope and where.** Section 1's list of what
   waits for 1.1 names WP21; section 2 carries `### WP21 Ideation, deferred to 1.1`
   with a done-when; section 6 moves WP21 to 1.1 beside WP18; section 6.1's unsized
   row and section 8's deferred list both name it.
3. **The plan says what it is.** WP21's text says the record is a sealed directory
   beside `.xeno/intents/` holding what a phase holds, and that it is not a phase, a
   command without a record, or a P6 used first.
4. **#224 is cited and not repeated.** WP21's paragraph names #224 for the cost of a
   phase at the front and gives one sentence of it, and adds only the reason section
   12 contributes, which #224 predates.
5. **The decisions are in the trail.** The intake carries Q-1 to Q-3, each with its
   options, consequences, a recommendation with its reason and a free entry, and
   D-1 to D-3 decided by the maintainer and resolving them.
6. **Nothing else moves.** `git diff --stat main` shows the plan and the trail; the
   process definition, the register and the code are untouched.
7. **The prose holds the conventions.** Every changed line of the plan is at most 88
   characters, the heading names its section in words, and no line is a list the
   reader cannot see.
8. **The gates stay green.** `./xeno gate verify` exits 0 and `go test ./...` is
   unaffected.

<!-- xeno:section:non-goals -->
## Non goals

- **A clause in the process definition.** The record type, Appendix B's path, the template and the commands are WP21's and each is a specification change when it is built.
- **Sizing or sequencing WP21.** The plan sizes what it sequences; WP21 is deferred and sits in the unsized row.
- **Naming the key prefix, the directory or the commands.** The decision fixed the shape and left the names to the package.
- **Answering #323.** The evaluation stays open; this intent takes its one finding about the plan.
- **A register row.** Every choice here is a decision the maintainer took on options, not an assumption the agent made inside a clause's room.

<!-- xeno:section:constraints -->
## Constraints

- The plan is normative and the agent wrote it on instruction after the diff was read, which is the exception XENO-0262 and XENO-0282 set the shape of: recorded in the commit message and in the intake, and never done unprompted.
- One decision at a time, each with its consequence, its cost, a recommendation with its reason and a free entry; `xeno question record` refuses a recommendation whose reason is not on the option, which was met once and corrected.
- Prose at 88, a heading in words, Conventional Commits, `Refs #337` on the plan commit and `Closes #337` on the pull request.
- The branch carries one intent and no code.
