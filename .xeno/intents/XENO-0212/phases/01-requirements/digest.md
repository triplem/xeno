---
intent: github.com/triplem/xeno#143
phase: 01-requirements
created: "2026-09-29T20:27:46Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+530ce03.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ba1a872477adce4f1ee25484fb3db595a2b73ebab64e8353e3a1dc29bb58aa95
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Seven criteria. Three are structural — the port returns requirements, the GitHub spellings live in one package
and nowhere else, the runner selects rather than defaults — and one is the removal of Protection and the
evaluate half of #99's table. Criterion 5 is the load bearing one and is written as a constraint on method:
the existing assertions in enforcement_test.go are the only evidence that this was a move, so a case that
will not pass is a finding recorded in deviations rather than an expected string to edit. A42 is checked
rather than assumed, because a new package importing net/http is exactly how the gate path's freedom from the
network gets broken by accident. Tracker is excluded and the reason is recorded: #104 omits it from the
done-when this intent answers, and an interface with no adapter and no caller is what #97 declined to write.
