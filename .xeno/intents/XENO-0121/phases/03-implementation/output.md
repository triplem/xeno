---
intent: github.com/triplem/xeno#121
phase: 03-implementation
created: "2026-09-28T19:43:35Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c2ba6b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 685731465bd31ccbbdacc7cc83a27a4f1ec1ec9ef77f46dd0fc8d459408c17f2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`.releaserc.json`. `@semantic-release/changelog` and `@semantic-release/git` are gone,
with the `changelogFile`, the `assets` list and the release commit message that belonged
to them. Three plugins remain: commit analyzer, release notes generator, github, the
last with its two comment routines still off.

`CHANGELOG.md`. Deleted, 160 lines. What it rendered is in the notes of every release
and in the history of the file.

`.github/workflows/release.yml`. The `extra_plugins` block is gone, so the action
installs semantic-release alone. The header paragraph no longer says the pipeline writes
the changelog back, and a new one says nothing writes to `main`, why the arrangement
could not survive #86's protection, and where the changelog went. The line that offered
the `GITHUB_TOKEN` property as the reason the changelog commit starts nothing is gone:
the property is true and is now the reason the commit could not pass, so leaving it as a
reassurance would hand the next reader a correct fact supporting the wrong conclusion.

`.github/workflows/audit.yml`. One `npm install` line instead of four, and the comment
says one package since #121 rather than claiming three lines that no longer exist.

`SUPPLY-CHAIN.md`. Two rows out of the table. A pin on a package nobody installs invites
somebody to keep it current.

`ASSUMPTIONS.md`. A23 superseded, naming what came true and what replaced it, with the
row it replaced still readable. The `## Built` paragraph no longer lists the changelog
among what the release pipeline does, because that section describes the present.

Verified by hand rather than by a test: `.releaserc.json` parses as JSON, both workflows
parse as YAML, and nothing outside this intent's own artifacts names either plugin or
the file.

<!-- xeno:section:deviations -->
## Deviations from the design

None in the change. Both plugins, the file, the four files that named them and A23 are
as P2 decided.

One thing P2 did not list and the sweep found: the `## Built` section of
`ASSUMPTIONS.md` also described the pipeline as writing the changelog back. P2 named A23
and not that paragraph. It is corrected here because `## Built` is a statement about the
present, and leaving it would have left the same wrong sentence one heading away from
the row that supersedes it.

The order of the record is the order of the work, and the specification came before
both. #122 changed the plan, in its own commit and its own pull request, because squash
is the only merge method here and a bundled change would have collapsed the ordering the
first standing rule is about. P0 to P2 were written before any configuration changed,
and the configuration changed inside P3.
