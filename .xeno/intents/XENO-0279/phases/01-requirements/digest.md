---
intent: github.com/triplem/xeno#94
phase: 01-requirements
created: "2026-10-08T20:14:31Z"
schema_version: "1.0"
runner_version: dev+30b1dea
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 25758f0d043cb5342c96f8255e06e0fd9387349617b6752ddcdcd9144bc7e090
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Twelve acceptance criteria, each a state of the tree: the policy file that fails on any
finding, four accepted gosec findings with their reasons and no more, tests outside gosec
and inside errcheck and staticcheck, two workflows in the scanners' shape with the
judgement of the vulnerability report written in a step, semgrep gone from the tree and
from every sentence outside the trail, the pin table held in both directions, the
required checks renamed on the host before the merge, renovate seeing the two new pins,
the suite and gates green with no behaviour change, and two register rows. Non-goals:
depguard, a period of running both, trivy, attachments, gosec on tests, more linters.
