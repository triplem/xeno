---
intent: github.com/triplem/xeno#38
phase: 00-intake
created: "2026-09-25T17:13:44Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+403d78f.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 5e1d8c7fda68e15bedb0bf9658491b091b1acec76e28e595ec0f070ed0615092
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

A repository cannot be set up to use Xeno. There is no `xeno init`, no
`project.yaml`, and no wrapper for a pipeline to call. The plan calls this every
user's first contact with the tool and says it has had the least thought spent on it.

<!-- xeno:section:scope -->
## Scope

WP9 without the host API, in this change: the configuration with three settings and
defaults for the rest, `--vendor` pinning the plugin, the version check that stops on
any difference, idempotence, and printing the settings a person still has to make.

The `enforcement` block is written with defaults and the edition check marked
outstanding, which the plan explicitly provides for where no token is available.

WP10 follows in its own change: the wrapper generator, `xeno enforcement check`
against the API, and the range parameters on `gate run`.

<!-- xeno:section:context-rationale -->
## Why this context

The two halves split at the network. Everything in this one is decidable from the
repository and testable without a host; everything in the other needs an API and a
token, and this repository is the negative case for it, since A27 means the host
answers 403 to every question about protected branches.

Splitting there also keeps the first contact honest. A command that fails because a
token is missing would be a poor first thing for anybody to run, and the plan says so.
