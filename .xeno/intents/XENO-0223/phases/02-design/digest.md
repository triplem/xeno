---
intent: github.com/triplem/xeno#171
phase: 02-design
created: "2026-10-03T09:44:58Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1f2113f8f0ae5296458c1203fe589e62031cb815b6482da598da981a2f466f32
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The changed set is printed and never recorded: section 5's enumeration has no field for it, and the
same section says the lock records what was declared rather than what was read, so a recorded
derivation would describe neither and could drift from its inputs. `repo_commit` comes from
`internal/git`, with absence a state rather than an invented value. `rules_applied` is the effective
set with each rule's path and version, written by the same resolution that writes the hash so the two
cannot disagree, and absent where the set is empty — the distinction A74 rests on. The budget is
checked in G-Schema against the lock, because the profile declares it and the lock records what was
given; the finding names both numbers and nothing blocks, which is what the section asks for in the
sentence calling it deliberately a finding. Bytes are measured from the tree when the check runs, not
stored, and a missing file is G-Freshness's finding rather than a second report of one cause. A
declared link adds its document to the base and a missing target is a finding against the profile. The
order is the profile's `include` order with paths sorted inside each pattern: volatility is the
project's judgement and the tie-break is what the hash needs. And the runner still reads nothing for
the agent — "reads only what changed" is a report, because the knowing is the agent's and the
knowledge is the runner's to hand over.
