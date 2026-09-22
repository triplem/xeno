---
intent: github.com/triplem/xeno#2
phase: 00-intake
created: 2026-09-22T09:00:00Z
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
open_questions:
  - key: Q-1
    text: Where is the DCO sign-off enforced?
    options:
      - text: A commit-msg hook from examples/hooks/, installed through core.hooksPath
        consequence: feedback only; --no-verify removes it and nothing behind it notices
      - text: A checked rule over the commit range, judged by G-Policy at the gate
        consequence: binding where it counts, at the merge, and it needs WP4
        recommended: true
      - text: The code host's own mechanism, a push rule or a DCO app
        consequence: enforced outside the trail, so the gate cannot say that it held
      - text: Something else
        free: true
---

# Intake

Xeno verifies its own repository in CI. From this point every change to Xeno runs
through Xeno.

## Scope

A CI job running `xeno gate verify` on every push and pull request.

## Non goals

Any other phase, any adapter, any agent.
