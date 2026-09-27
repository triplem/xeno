---
intent: github.com/triplem/xeno#98
phase: 00-intake
created: "2026-09-27T09:44:22Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+5ddcaab
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: d4ea6131918e594e50d2b201c4499f46579418402a1f66e52d8673d568ee4c11
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

Two functions in the runner were long string literals with positional arguments:
`renderGitHub` at 46 lines with three, and `projectYAML` at 49 with four. Both write
files into somebody else's repository, the CI wrapper and the initial configuration, and
both are the first thing an adopter sees.

Two costs followed. An organisation with its own defaults could change neither without
forking the runner, which is the issue's point. And a positional argument in the wrong
place produced a wrapper that looked right and judged the wrong commit, which is the
cost nobody had paid yet.

<!-- xeno:section:scope -->
## Scope

`internal/scaffold`: the two files with `text/template` and named fields, a project
override under `.xeno/config/scaffold/`, and a generated file that names the copy it
came from.

Not the plugin, and not the template package. The `.gitignore` lines stay where they
are: a list of one entry is not a template.

<!-- xeno:section:context-rationale -->
## Why this context

**The plugin is the obvious home and cannot be one.** `xeno init` is what creates
`.xeno/` and what vendors the plugin, so a file it needs in order to run cannot live in
what it produces. G-Supply says the same from the other side: the digest it checks
belongs to a plugin that does not exist when the wrapper is written. So the defaults are
embedded in the runner and the override is read from the project, which is available on
every run after the first.

**A scaffold is not a template, and the packages say so.** A template renders an
artifact a gate then judges, and carries an id, a version and a strings hash for exactly
that reason. A scaffold is a starting point somebody edits afterwards. Sharing the
resolution code would have been fifteen lines saved and an invitation to give the
wrapper a version nobody reads.

**The generated file names its source because an override can break the tool.** A
wrapper that does not call `xeno gate run` produces a green pipeline that judges
nothing. That is the adopter's business, but it should not be silent, and the line costs
nothing: `context.lock.yaml` records `template_source` for the same reason (A34).

**Named fields are the half that pays without any override.** The wrapper took three
positional arguments and the configuration four. Both now take a struct, so the failure
mode where a fourth argument lands in the third position and the file still looks
plausible is gone rather than documented.