---
intent: github.com/triplem/xeno#143
phase: 03-implementation
created: "2026-09-29T20:36:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+530ce03.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ad97e4880c94000dde64a2250d1e3c5bae247a79bdcdf528553ed1b15ee6ba59
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
internal/host holds BranchRules and selects by project.yaml with no default; internal/host/github holds every
GitHub spelling, with transport and decode moved bodily and the four evaluate branches split into named
builders; internal/enforcement loses Protection, ReviewsExpressible and the transport, keeps the filter, the
waive logic and the merge_method line, and gains the names as exported constants so a misspelling does not
compile. The report for this repository is byte identical to main's apart from the timestamp, which is the
strongest available evidence that the move preserved behaviour. One real regression was caught in the writing:
returning nothing where the host offers no protection gave the right state and lost the reason it carries, and
every state assertion would have passed. Criterion 4's reporting half could not be built as written and the
alternative P2 recorded was built instead.
