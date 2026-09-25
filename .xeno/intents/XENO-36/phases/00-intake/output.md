---
intent: github.com/triplem/xeno#36
phase: 00-intake
created: "2026-09-25T16:06:09Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+69a7600.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 83666dafd60c5af5a95691284f86f81c0de4ab7aeb1463e0c2c5e1c7526f889b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

Seven files describe a host this project is not on. The plan names a self managed
GitLab Community Edition as the target, the process definition follows it in five
places, and a set of assumptions promises work that a move there would bring.

<!-- xeno:section:scope -->
## Scope

The two normative documents, the v2 delta, and everything written in expectation of
the move: A18, which is void, the "thrown away at the move" clauses of A21 and A24,
one paragraph of SUPPLY-CHAIN.md and the header comments of two workflows.

No code. WP9 and WP10 follow this in their own change, because the specification is
the control and a change to it is its own commit made before the code that follows.

<!-- xeno:section:context-rationale -->
## Why this context

The target host paragraph decides more than a name, so the reasoning attached to it
has to be rewritten rather than search and replaced. A self managed instance has no
canonical endpoint, which is why a base URL was a required setting; GitHub has one,
and an adapter still takes it as a parameter for Enterprise Server. Approval rules
were absent on the Community Edition; on GitHub they exist and depend on the tier,
which is the general form of what A27 records for this repository.

A27 itself stands. Branch protection needs a paid plan for a private repository, and
that is a property of the tier rather than of the host, so the sequence guarantee
still rests on discipline rather than on a setting.
