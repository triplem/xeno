---
intent: github.com/triplem/xeno#44
phase: 00-intake
created: "2026-09-25T21:36:10Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+fdade9d
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: c7a8676b1b784707967c28369df67e84528275894db4c02936c2170f14e6029f
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

Trivy scans the binary a release ships. Nothing looks at what the pipeline runs to
produce it, and the sharpest part of that is not a stale pin: `semantic_version`
pins semantic-release one package deep, npm resolves the rest at run time, and the
tree that results executes in the one workflow holding `contents: write`,
`issues: write` and `pull-requests: write`.

<!-- xeno:section:scope -->
## Scope

Two halves, in the order the issue asks for.

**What the release workflow may do**, which is worth more than watching it. The two
extra permissions are dropped, and the plugin configuration that could have needed
them is replaced by the form that does not.

**What is fetched and run**, which is audited on a schedule against a recorded
baseline, produced as a report like the other two scanners so it can be declared
later.

Dependabot is not adopted: the reasoning from #11 is unchanged, and its pull requests
would still collide with every change being an intent. Scanning the workflow files
for configuration problems is not done either; they are pinned to commit shas and
reviewed, which is the question that check would ask.

<!-- xeno:section:context-rationale -->
## Why this context

**The permissions were measured before they were removed.** The log of the run that
published 0.14.0 shows the plugin doing nothing to an issue or a pull request after
`Published release`, and no `released` label exists in this repository, because
`successComment: false` turns off the whole routine that comments and labels. What
remains is the failure path, where the plugin opens an issue, and that is turned off
by configuration rather than left to a permission.

**The audit was run before the check was written**, against the exact pinned set:

    19 vulnerabilities (1 low, 3 moderate, 15 high)

Almost all of them under `node_modules/npm/node_modules/`, because semantic-release
depends on npm as a library. None of them is ours to fix. A check that failed on this
would fire on its first run with no action available, which is the shape trivy.yaml
already rejects for unfixed findings, and people learn to route around.

So the check is on drift: the counts are recorded, and a twentieth finding fails
because it is a change somebody has to look at. That separates the two questions the
issue asks to keep apart, since a count that does not move is not evidence that
anything is fresh.

<!-- xeno:section:decisions -->
## Decisions

- The two extra permissions come off and the plugin is configured so that it never
  wants them, rather than kept because a future plugin might.
