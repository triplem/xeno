---
intent: github.com/triplem/xeno#190
phase: 03-implementation
created: "2026-10-03T14:03:40Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 381601e7f8d136985a0383d63e0c9de2ee965b4e6a44a85aab5677ff402d2554
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

**`.github/workflows/xeno.yml`, three changes.**

`fetch-depth: 0` on the checkout, with the comment that the step below compares against the base of
the pull request and a shallow clone does not have that commit.

A step **"the trail is not rewritten"**, after the ci-skip step and before `format`. Its comment
carries the measurement: an edit is caught by `gate verify`, a deletion by nothing, and `rm -rf` on
one intent exits 0 with six fewer verdicts. It refuses a base it cannot find —
`git cat-file -e "$BASE^{commit}"` — with "Refusing rather than passing on no evidence", and
otherwise runs `git diff --diff-filter=DR --name-only "$BASE" HEAD -- .xeno/intents/`, printing the
paths and the rule where the list is not empty.

A step **"the title is the commit subject, so it is a Conventional Commit"**, after `build` because
it uses that binary. It pipes `$TITLE` into `xeno check commit-message --pattern
conventional-commits` and, on failure, says that the title becomes the squashed subject and that a
prose title releases nothing silently. Its comment records why the plain pattern and not the one
with the reference, and the five days as the reason the step exists.

Both are `if: github.event_name == 'pull_request'`, and both take their string through `env`, which
is the decision the neighbouring step had already made and commented.

**`CONTRIBUTING.md`, two paragraphs.** Under "Every change is an intent": the trail only grows,
section 7's rule, what catches an edit, what caught nothing until #193, and that a branch which
corrects itself is not refused. Under "Commit messages", immediately after the subject rule it
qualifies: the title *is* that subject, the host appends the reference so
`feat(docs): a description` arrives as `feat(docs): a description (#3)`, which is what the second
shipped pattern is for, and a prose title releases nothing silently.

**Tested against real cases before this was written**, in a throwaway worktree off main:
a committed deletion of `XENO-0225` reports 31 paths; an addition reports 0; a directory added and
then dropped in two commits reports 0. And the title step's command accepts `fix(ci): a real title`
and refuses `The audit gate binds`.

**No Go file, no gate, no rule, no document.**

<!-- xeno:section:deviations -->
## Deviations from the design

**None from the design.** Every decision P2 recorded is what the code does, including the two step
positions and the reason for each.

**One thing the design did not have to settle and the implementation did.** The trail step's refusal
of a missing base is phrased as its own `if` with its own two-line message rather than folded into
the comparison. Folding it in would have been three lines shorter and would have made the one case
this intent is most worried about — a comparison against nothing — share its output with the case
where everything is fine.

**One wrapping deviation.** The title step's pipeline is split across two lines with a trailing `|`,
because `./xeno check commit-message --pattern conventional-commits` on one line with the `if` and
the `printf` runs to 102 columns. The project wraps at 88 in code as well as in prose, and a shell
pipeline is the one construction here where the wrap is load bearing rather than cosmetic: a broken
continuation changes what runs.
