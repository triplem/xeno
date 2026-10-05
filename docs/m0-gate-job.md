# The M0 gate job: Xeno verifies its own repository

This is a record of how M0 was reached once, in the tree and with the tooling of the
day, not a procedure to follow now.

The implementation plan uses M0 for two things. In the milestone table it is the gate
job, the point from which Xeno checks its own repository; that is what this guide
builds. In section 6 it is the walking skeleton, one intent, one phase and one agent end
to end through template, rendering, `xeno init`, CI wrapper and merge gate, which still
needs WP2, WP3, WP9, WP10 and WP11. This takes one intent, one phase and one CI job.
Every
block below was run end to end before it was written down.

The implementation plan targets GitHub, which is where this repository is, so the
workflow file and the module path are bound to the host it is actually on. That was not
always so: the plan named a self managed GitLab until #36, and A18 scheduled the module
path to move with it. A18 is void.

## 1. Lay out the repository

```
<repo root>/
  .github/workflows/xeno.yml
  docs/process-definition.md      renamed from xeno-process-definition-v1.md
  docs/implementation-plan.md     renamed from xeno-implementation-plan-v1.md
  docs/orchestrator-evaluation.md
  docs/v2-delta.md                renamed from v2-delta-against-v1.md
  docs/assumptions.md             renamed from ASSUMPTIONS.md
  docs/m0-gate-job.md             this guide, renamed from M0.md
  cmd/  internal/  vendor/  go.mod  README.md  .gitignore
```

Build once locally:

```sh
go test ./...
go build -o xeno ./cmd/xeno
```

## 2. Open the issue

Open issue #1 in the repository, "M0: Xeno verifies its own repository". Its number is
the intent. Replace `OWNER/REPO` below with the repository's path, identically in every
file: G-Trace compares the `intent` field across them.

## 3. Create the intent

`.xeno/intents/XENO-1/intent.yaml`

```yaml
intent: github.com/OWNER/REPO#1
key: XENO-1
status: in-progress
created: 2026-09-21T09:00:00Z
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
```

`.xeno/intents/XENO-1/assumptions.yaml`

```yaml
intent: github.com/OWNER/REPO#1
created: 2026-09-21T09:00:00Z
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
assumptions: []
```

## 4. Run P0

```sh
./xeno phase start --intent XENO-1 --phase 00
```

Then write the three files of the phase into `.xeno/intents/XENO-1/phases/00-intake/`.

`output.md`

```markdown
---
intent: github.com/OWNER/REPO#1
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
```

`digest.md`

```markdown
---
intent: github.com/OWNER/REPO#1
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
---
Written by hand before M0; there was no session to summarise.
```

`learning.yaml`

```yaml
intent: github.com/OWNER/REPO#1
phase: 00-intake
created: 2026-09-21T09:00:00Z
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
no_finding: true
```

The fields marked `by-hand`, `tool: manual` and `model: none` are honest rather than
filled in: nothing produced them. G-Schema checks that they are present, and G-Supply
and G-Secret, which would check what they mean, are not implemented yet and say so in
the verdict.

```sh
./xeno phase finish --intent XENO-1 --phase 00   # green, with four not-implemented
./xeno gate verify                              # verified 1 verdicts, exit 0
```

## 5. Commit and push

One commit with everything, carrying the intent in a trailer:

```sh
git add -A
git commit -m "M0: Xeno verifies its own repository" -m "Xeno-Intent: XENO-1"
git push
```

The workflow runs the tests, builds the runner and runs `xeno gate verify`. When it is
green, the gate job stands. The walking skeleton does not yet.

## From here

Each work package is opened as an intent and its intake runs through Xeno; the rest of
the work is recorded by hand in `docs/assumptions.md`, as before. That is less than
every change running through Xeno, and deliberately so: G-Freshness checks only its
first half until WP8 (A6), so no phase beyond P0 runs in this process before that gap
is closed, and the hand kept record ends with the walking skeleton rather than with
this job. A red verdict stops the next phase (A12), which is where proportionality will
show itself first.

WP0 is the first such intent, `XENO-2`. Its intake is in the repository without
`context.lock.yaml` and `gate.yaml`: replace `OWNER/REPO` in its files, write its
`digest.md` as in step 4 above, then

```sh
./xeno phase start  --intent XENO-2 --phase 00
./xeno phase finish --intent XENO-2 --phase 00
```

and commit with both trailers, `Xeno-Intent: XENO-2` and a sign-off.

The digest is easy to forget here because this section lists the two files the runner
writes and step 4 lists the three a person does. It was forgotten on the first run, and
the verdict was green because at that point no gate looked for the file. G-Schema does
now, which is the more useful half of the repair: the instruction above can be misread
again, the gate cannot.
