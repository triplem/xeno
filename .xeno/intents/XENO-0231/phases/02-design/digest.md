---
intent: github.com/triplem/xeno#190
phase: 02-design
created: "2026-10-03T14:02:55Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: adb7d3cc26947b41663591761a817795bee68537284fdfa03ebfae8d07bd0fae
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Both steps go in `xeno.yml`'s `verify` job, because its id is the context `main` requires and a step
beside it would be advisory — #186 is the cost of the other choice, measured at six merges. The
title step runs after `build` and calls `xeno check commit-message`, so one rule has one
implementation; a regular expression in the workflow would be a third statement of it beside
`CLAUDE.md`'s prose and the plugin's pattern. The trail step runs early, before a minute of Go,
because its subject is a destructive mistake, and it compares `git diff --diff-filter=DR` between the
base and `HEAD` — two-dot so it compares trees, since a directory created and dropped inside one
branch never reached main; `HEAD` because on a pull request that is the merge result and so the tree
the merge would produce. `D` and `R` and not `M`: `verify` catches modification and names the phase,
and the path is inside `artifacts_hash` so a rename is a rewrite. A base that is not in the clone is
its own failure with its own message, because the alternative prints what a clean run prints. The
checkout gains `fetch-depth: 0`; both strings come through `env` for the reason the neighbouring step
already records. Eight alternatives refused, four of them because they would have created a second
statement of one rule — an expression beside the pattern, a workflow beside the required context, a
hash beside what git knows, a default base beside the one the pull request names.
