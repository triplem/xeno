---
intent: github.com/triplem/xeno#19
phase: 00-intake
created: 2026-09-24T18:11:27Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b43f326d64e7109fe30a5b1b640650298caaeb188eace6de0d0a5f182404a749
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #19. The only dependency this project has is unmaintained, and the package it
came from now lives somewhere else.

## Scope

`gopkg.in/yaml.v3` becomes `go.yaml.in/yaml/v3` at v3.0.5, with the imports, the
vendor directory, `NOTICE` and `SUPPLY-CHAIN.md` that follow from it.

This changes what a release carries rather than how it is built. The binary is
different afterwards and so is the bill of materials published beside it, which is what
makes it an intent rather than maintenance.

## Non goals

Anything the newer version offers. Moving and writing against it are separate changes.

Removing the dependency. One vendored YAML library is a reasonable thing to have and
this intent is not the place to reopen that.

## What is expected to be untrue by the end

`SUPPLY-CHAIN.md` explains why `vendor/modules.txt` carries no go annotation for this
dependency: the old module declares no `go` directive. The new one declares `go 1.16`,
so the annotation appears and the explanation stops holding. It was written two changes
ago. Noted here so the correction is part of the work rather than a later surprise.
