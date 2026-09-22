---
intent: github.com/triplem/xeno#1
phase: 00-intake
created: 2026-09-21T09:00:00Z
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: by-hand
model: none
tool: manual
tool_version: "0"
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Xeno verifies its own repository in CI. From this point every change to Xeno runs
through Xeno.

## Scope

A CI job running `xeno gate verify` on every push and pull request.

## Non goals

Any other phase, any adapter, any agent.
