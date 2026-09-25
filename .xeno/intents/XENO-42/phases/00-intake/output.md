---
intent: github.com/triplem/xeno#42
phase: 00-intake
created: "2026-09-25T19:00:52Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+73849dd.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 933ac6118216e0c719e94c1b5c89bc754bc36ddbea2a6d67bf060da2a58e0868
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

Three passages read as spliced rather than as written. A "there" points at a sentence
that is gone, a sentence sits in the middle of a paragraph it does not belong to, and
one claim about this repository has been false since #19.

<!-- xeno:section:scope -->
## Scope

The three passages, and nothing else. No revision is counted up: these are typos in
the sense the plan means, and a revision that moves for a comma answers nothing.

<!-- xeno:section:context-rationale -->
## Why this context

What the three have in common is the finding rather than the three instances. Each
was a paragraph rewritten across two or three separate changes; the passages rewritten
whole in those same commits are intact. Editing into a sentence leaves the words
around it behind, and nobody rereads the result because the diff looks small.
