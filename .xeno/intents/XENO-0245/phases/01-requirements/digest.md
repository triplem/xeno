---
intent: github.com/triplem/xeno#217
phase: 01-requirements
created: "2026-10-04T20:03:48Z"
schema_version: "1.0"
runner_version: dev+fbf8a72.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bdf2e5a91353448be6d7f0b3d6e7f667140d475784c750057f738cc4de0943ca
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Requirements turn the two decisions into criteria and settle the writer's shape by
precedent rather than by invention. D-1 records keeping the context economy in v1 with
the scope required at P0, against making it optional and against cutting WP8, WP15 and
WP20 to 1.1; the deciding facts are that WP8's own blocker, "P0 cannot produce a scope
until the format is fixed", is already cleared by `model.Scope`, and that WP20 calls
cache friendly assembly the one lever impossible to retrofit. D-2 records the
requirement attaching as a refusal in `phase finish` rather than as a gate, and carries
the correction that forced the question: #217 says A74 keeps sealed verdicts from being
re-judged, A74 binds G-Policy alone to the rule set a phase records, and `Verify`
recomputes and compares every phase's status — so a gate would re-judge 100 sealed P0s,
and backfill is impossible because `PhaseExcluded` holds only `gate.yaml` and
`cost.yaml` and the scope is therefore inside `artifacts_hash`. A refusal is the only
forward-only option, because commands are not re-run where verdicts are recomputed. The
writer reads its entry as a document rather than as flags, which is `question record`'s
argument applied to patterns, links and two budget numbers, and it reports what the
patterns resolve to — a criterion taken straight from this intent's own P0, where the
figure was established by hand with `wc` and the first budget chosen was wrong and could
only be corrected because nothing was sealed yet. A scope written after its phase has a
verdict behaves as `section set` does, so no second mechanism is invented for a file
that happens to be YAML. Three mechanisms that have never had an input get a criterion
each: the lock's `files` list, the changed-file report at `phase start`, and the link
check. Non-goals, each with its reason: the 100 sealed P0s, which nothing can backfill,
so absence and compliance look alike in history; #235, the budget overrun that turns a
phase red against section 5's explicit words, which needs a notion of a finding that
does not fail its check and is therefore WP1's argument; validating the scope's content
in G-Schema, which had no subject until this intent; WP8's other half; a project
default, since the scope is per intent; a migration command for two files; and WP15 and
WP20. One question is deliberately left to P2 rather than assumed: whether the
information base is what a phase was given to read or also what the intent will change,
because an intent naming a file it then edits has moved an input of its own earlier
phases and G-Freshness will correctly say so from P3 on. This intent will be the first
to see it. Constraints: the writer and the refusal ship in one commit, the writer runs
before `phase finish` because the scope is inside the hash, nothing sealed is rewritten
and the figures say so rather than the prose, `model.Scope` is the one authority on the
format, the refusal names its own remedy, and no new gate, gate result, rule or
oddly-behaving finding appears. The lock for this phase resolved 43 files and 904,078
bytes against the declared 50 and 1,000,000. The phase was then reopened twice, both
times from what that first resolution caused rather than from a change of mind. The
maintainer's commit fbf8a72 renamed the artifact to the context scope, so the criteria
carry the rename of model.ContextProfile and model.Profile and of this intent's own
file. And that commit moved two documents the scope names, which turned this phase red
through a check that read its own lock: Q-3 put whether #236 belongs here or to an
intent of its own, and the answer widens this one, so three criteria now cover section
5's two limits and a fourth asserts that recomputing this phase stops needing the
release it was given. One non-goal is withdrawn rather than carried: P0 left to P2 the
question of what a scope should contain so that staleness stays quiet, and section 5
answers it against the check instead — "a phase that changes the files it read is not
stale, it is working" — so there was nothing for an intent to decide.
