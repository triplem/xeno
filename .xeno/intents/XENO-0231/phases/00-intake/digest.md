---
intent: github.com/triplem/xeno#190
phase: 00-intake
created: "2026-10-03T14:01:26Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b3fda305f67b232e62b645bdada281d943abd01d66119e731303dde41999d457
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two conventions this repository states and does not check. The pull request title is the commit
subject — `squash_merge_commit_title` is `PR_TITLE` and squash is the only method — so the subject
written on a branch is discarded and the title is the only place `CLAUDE.md`'s rule has any effect;
fifteen merges since `v0.28.0` carried prose titles, semantic-release read sixteen commits and
concluded "no release", and nothing ran for five days over WP4, the rule engine, external gates, the
agent layer, the context lock and the symbol index. `xeno check commit-message` ships the pattern
and is invoked nowhere. And a sealed intent can be deleted with nothing noticing: editing a sealed
`output.md` is caught, `verify` reporting DIVERGENT and exiting 1, while `rm -rf` on one intent takes
259 verdicts to 253 and exits 0, because nothing records how many there should be — a worse
violation of "what is sealed is never rewritten" than an edit, since an edit at least diverges. Both
are the shape of the last three intents: something stated in one place, and nothing comparing the
repository against it. The second has already cost a phase, XENO-0230 having reached main without
its verification or review while every gate was green. In scope: two steps in the `xeno` workflow,
so both are inside the one required check; the title against the shipped plain pattern; the trail
compared against the pull request's base as a tree diff, so a directory created and dropped inside
one branch is not reported; and both conventions written where a contributor reads them. Out of
scope: modification, which `verify` catches; re-sealing, which the working sequence provides for
deliberately; the fifteen missed releases, a maintainer's decision; G-Complete's P5-only scope; and
a hook, for the standing reason.
