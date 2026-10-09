---
intent: github.com/triplem/xeno#337
phase: 04-verification
created: "2026-10-09T16:00:52Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0620fa450cb0eef4f26c81e0be3f4110b19cd7c1bdd55b38724d06a724051086
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| criterion | how it was checked | result |
|---|---|---|
| 1. the plan precedes the trail | `git log main..HEAD`: `ec372ea` first, one file, 53 in and 6 out; its message carries the exception | met |
| 2. in scope and where | `grep -n WP21 docs/implementation-plan.md`: five lines, at section 1's deferred list, the section 2 heading, the section 6 sentence, the 6.1 unsized row and section 8's deferred list | met |
| 3. what it is | the shape paragraph says a sealed directory beside `.xeno/intents/` holding the four files a phase holds, not a phase, with section 12's reason against a phase before an issue; the command without a record is excluded by the sentence on gated documents | met |
| 4. #224 cited and not repeated | `grep -n '#224'`: one line, "for the reasons #224 gives, which are not repeated here", followed by one sentence of it and then section 12's reason | met |
| 5. decisions in the trail | the intake's frontmatter: three `key: Q-`, three `free: true`, three `id: D-`, each decision `decided_by: triplem` and `resolves` its question | met |
| 6. nothing else moves | `git diff --stat main..HEAD` outside the trail: the plan alone; `git status` clean apart from the intent directory | met |
| 7. conventions | every added line of the plan within 88 columns by `awk`; the heading is words; the sentence after the size table stays true at twenty built for v1 | met |
| 8. gates green | `./xeno gate verify` verified 583 verdicts, exit 0; `go test ./...` exit 0 | met |

<!-- xeno:section:results -->
## Results

Eight criteria, eight met.

    git log main..HEAD                 ec372ea, docs/implementation-plan.md alone, then the trail
    grep -n WP21 docs/implementation-plan.md    5 lines: 105, 1384, 1720, 1752, 1835
    grep -n '#224' docs/implementation-plan.md  1 line, 1399
    intake frontmatter                 Q-1..Q-3 with free entries, D-1..D-3 by triplem
    added lines over 88 columns        none
    go test ./...                      ok, exit 0
    ./xeno gate verify                 verified 583 verdicts, exit 0

**One correction inside the trail.** P3's first judging said 54 lines in; the diff
after the trim is 53, and the section was re-set and the phase re-judged before P4
started, so no phase recorded a predecessor hash that moved.

**Timings, for #117.** `gate verify` 49 s over 583 verdicts with nothing else running;
a run alongside `go test ./...` read 69 s over 582, which is the machine and not the
trail, and is reported so that the figure is not taken for a jump.

<!-- xeno:section:gaps -->
## Gaps

**Nothing here is executable.** The deliverable is a document saying what a package is and when it is done; whether the shape holds is proved by building WP21 in 1.1, and the first thing that build does is the process definition change the plan names.

**The done-when's third property has no method yet.** "A figure #117 can place beside an intent's" assumes a record's cost is counted the way an intent's is, and nothing says how a record without phases is counted; that is WP21's to say.

**The adapter write is a widening the plan names and does not decide.** WP12 fixes four operations and calls a fifth a decision; WP21 says the widening is that decision and leaves it to the build.

**Three decisions by one person in one sitting.** Each was put on its own with its options, but the second and third were framed after the first was answered, by the same agent; a reader who disagrees with D-1 finds D-2 and D-3 framed inside it, which is the shape one-at-a-time has.
