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
---
`xeno.yml` gains `fetch-depth: 0` on the checkout, a step "the trail is not rewritten" before
`format`, and a step "the title is the commit subject, so it is a Conventional Commit" after
`build`. The trail step refuses a base it cannot find with `git cat-file -e "$BASE^{commit}"` and
the words "Refusing rather than passing on no evidence", then runs
`git diff --diff-filter=DR --name-only "$BASE" HEAD -- .xeno/intents/` and prints the paths and the
rule where the list is not empty; its comment carries the measurement that produced it. The title
step pipes `$TITLE` into `xeno check commit-message --pattern conventional-commits` and says, on
failure, that the title becomes the squashed subject and that a prose one releases nothing silently.
Both are guarded by the event name and both take their string through `env`, which is the decision
the neighbouring step had already made and commented. `CONTRIBUTING.md` gains two paragraphs: the
trail only grows, with what catches an edit and what caught nothing; and the title is the subject,
with the host's appended reference, which pattern is which, and the five days. Everything was tested
in a throwaway worktree first — a committed deletion reports 31 paths, an addition 0, a directory
added and dropped in two commits 0 — and that third case is what decided the two-dot form, which
reasoning would have got right only by accident. Two deviations: the missing-base refusal is its own
branch with its own message rather than folded in, and the title pipeline is split across two lines
because one line runs to 102 columns and a shell wrap is load bearing.
