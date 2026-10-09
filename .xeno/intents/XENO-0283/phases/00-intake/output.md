---
intent: github.com/triplem/xeno#346
phase: 00-intake
created: "2026-10-09T15:10:06Z"
schema_version: "1.0"
runner_version: dev+9590797
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 87865ec7e33a1a31be6160e7b377f4d1859407074bdbcb17ae6af3dff0b16932
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

Approved by @triplem on 2026-10-09T14:20:50Z: this is a nice extension

> **github.com/triplem/xeno#346** — The label script's read stops at thirty labels, so its second run fails
>
> Found by running it, 2026-10-09, minutes after #345 merged.
>
> `.xeno/plugin/bin/xeno-labels.sh github triplem/xeno` created the label on its first run
> and on its second tried to create it again, which `gh` refused: *label with name
> "xeno-approved" already exists*. The script is meant to change nothing the second time,
> and its test asserts the shape of that and not the behaviour.
>
> ## Why
>
> The read before the write is `gh label list --json name`, which returns the first
> thirty labels and no more. This repository has thirty-three, ten of GitHub's, twenty-one
> work packages, the old `approved` and the new one, so the new one was on the page the
> read never asked for. `glab label list` pages at twenty by default and has the same
> fault waiting.
>
> ## What fixes it
>
> `--limit` on `gh`, high enough to be the whole list, and `--per-page` on `glab`, or a
> read by name rather than by list: `gh label list --search xeno-approved` returns the
> match alone, and `glab` has no search but takes a per-page. A test that exercises the
> idempotence would need a host; what can be held without one is that the read names a
> limit or a search, which is the shape of the fault.
>
> ## Done when
>
> The second run of the script against a repository with more than thirty labels reports
> the label as existing and exits 0, on both hosts' branches, and the plugin test holds
> the read's shape.
>

The fault is in the tree and was reproduced against the host before this phase was
written: `sh .xeno/plugin/bin/xeno-labels.sh github triplem/xeno` on `main` at `9590797`
printed *label with name "xeno-approved" already exists; use `--force` to update its
color and description* and exited 1, with the label present on the repository. The
host carries thirty-two labels today, the old `approved` having been deleted at #345's
merge, so the label is still beyond the first page.

**Where the read is.** Line 31 of the script, `gh label list --repo "$project" --json
name --jq '.[].name' | grep -qx "$label"`; `gh label list` documents `--limit` with a
default of thirty and `--search`, which matches names and descriptions and sorts by
best match. Line 38, `glab label list --repo "$project" --output json`, has the same
shape; `glab`'s documentation gives `--per-page` a default of thirty, not the twenty
the issue says, and no search flag. The test beside it,
`TestTheLabelScriptCreatesTheLabelTheClauseNames` in `internal/plugin/plugin_test.go`,
holds that the script is there, executable, and carries the label constant, the word
and both create commands; nothing in it reads the read.

**The fault was named before it was met.** XENO-0282's verification wrote, under gaps,
*a host that paginates past the label would create a duplicate, which both hosts refuse
anyway*. The refusal is what happened; what it missed is that a refused create under
`set -e` is an exit 1, so the script is not idempotent on any repository with more
labels than one page, and this repository is one.

<!-- xeno:section:scope -->
## Scope

This intent makes the script's read before the write ask for the whole answer on both
hosts, and makes the plugin test hold that shape.

- **The GitHub read.** `gh label list --search "$label"` in place of the unbounded list,
  keeping the exact-match `grep -qx`, because `--search` matches substrings and
  descriptions and the description of this label contains its own name.
- **The GitLab read.** `--per-page 100` on `glab label list`, the most the host's API
  returns on one page, since `glab` has no search; the bound is written on the line.
- **The test.** `TestTheLabelScriptCreatesTheLabelTheClauseNames` gains the two read
  flags to its list of strings the script must carry, which is the shape of the fault
  and all a test without a host can hold.
- **The proof.** The second run against this repository, which has more than thirty
  labels, reports the label as existing and exits 0; recorded in P4.

What it does not do.

**It does not page.** A loop over pages would make the GitLab branch correct past a
hundred labels at the cost of a parser in POSIX shell; a hundred labels is well past
what a tracker that uses them for work packages carries, and the bound is a comment on
the line rather than a silent edge.

**It does not test against a host.** A test that ran `gh` would make the suite depend
on the network and on a credential, which section 7 keeps out of the gate path; the
test holds the read's shape and P4 holds the run.

**It does not touch the constants, the init line or the clause.** Section 12 names the
script and says what it does; it does not say how it reads, so nothing in the
specification moves.

**It does not correct XENO-0282.** Its gap paragraph is sealed and was right as far as
it went.

<!-- xeno:section:context-rationale -->
## Why this context

- Issue #346 with its approval, and the reproduction against the host on `main`.
- `.xeno/plugin/bin/xeno-labels.sh`, the two reads and the test's strings.
- `internal/plugin/plugin_test.go`, the test that holds the script's shape.
- `internal/model/identity.go`, for the constants the test reads the label from.
- `docs/process-definition.md` section 12, which names the script, to confirm it says
  nothing about how it reads.
- `docs/commands.md`, the paragraph naming the script, for whether it needs a word.
- `.xeno/intents/XENO-0282/phases/04-verification/output.md`, the gap that named this.
- `gh label list --help` and `glab`'s `docs/source/label/list.md`, for the flags and
  their defaults.
