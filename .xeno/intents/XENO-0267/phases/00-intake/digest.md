---
intent: github.com/triplem/xeno#228
phase: 00-intake
created: "2026-10-06T20:27:13Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 16c0ce93519790623d8eff27159bcacc2367285cc38a15a6224b81fa6e3aa464
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The runner says nothing about starting a phase in a fresh session, and the context economy of
the process assumes one: a phase that inherits its predecessors' sessions works from a context
`context.lock.yaml` does not describe, and the digest it was given stops being load bearing.
The gap is visible already, in the cost ledger's `none` lines and in the seven per cent the
cost package records. What is missing is a sentence, in the one place built for sentences
nobody has to act on.

Two sentences go into `internal/runner/next.go`: the `not-started` suggestion says the phase
is started in a fresh session and why, and the last phase's suggestion says the same for the
session that is ending. Section 7 forbids the runner naming a harness's command, so neither
does.

The issue's two other candidates are refused with their reasons, and both reasons outlive this
intent. The plugin's skill text may not carry `/clear` either: section 13 ships one vendored
plugin for both clients, and a skill is read by the model while clearing is the person's act.
The cost figure comparing declared context against spend belongs to WP20, which owns session
discipline and the baseline a figure is read against.

Out of scope: any change to the documents, a hook that ends a session, and enforcement of any
kind.
