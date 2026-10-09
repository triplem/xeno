---
intent: github.com/triplem/xeno#332
phase: 00-intake
created: "2026-10-09T13:29:30Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 88cb1b0a8c81ca8541cb63839792418cc750f3d7e6ff82def86b8123ca136f2b
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

Approved by @triplem on 2026-10-09T11:43:51Z: Use /xeno approved in the message (as first line) Use xeno-approved as the label

> **github.com/triplem/xeno#332** — Is approved the right label or should we use xeno-approved
>
> To distinguish to already existing labels in a brownfield project?
>
> If yes, all issues in xeno which are named approved should be re labeled.

The issue as it stood at 2026-10-09T13:29:30Z, read by `xeno phase start` and quoted rather than summarised. What this phase concludes about it belongs below.

The issue as `xeno phase start` read it, with the maintainer's answer above the quote,
read from the host by the clause this intent changes: the label `approved` and a comment
beginning `approved` were both there, under the old names, which is the last time either
name starts anything.

**The analysis the decision rests on**, posted on the issue on 2026-10-08 and restated
here because `phase start` quotes the body alone. The label and the word are fixed by
section 12, held in `model.ApprovedLabel` and `model.ApprovedWord`, printed by the
refusal and named in `docs/commands.md`. A brownfield tracker may already carry an
`approved` label that means approved for a sprint or by legal, and an intent started by
it would be #95's defect with a different cause; a bare `approved` as a comment's first
line is also what one person writes to another about a design on a labelled issue. The
`@xeno` form the issue proposed is a mention, and a GitHub account `xeno` has existed
since 2009. The maintainer decided: the label is `xeno-approved` and the comment's first
line is `/xeno approved`, and the earlier comments stand: a script in the deliverable
creates the label, and #323's trigger labels are not labels Xeno asks for.

**The specification moved first.** Section 12's paragraph names both and says why, and
names who creates the label, a person with the plugin's `bin/xeno-labels.sh` or by hand,
which `xeno init` lists among the settings it does not make; commit `b431e42`, revision
10 to 11, written by the agent on the maintainer's instruction and recorded here so that
the exception to the first standing rule stays one.

**What the tree says.** One constant for each name, read by `Issue.Approval()` and
printed by `approved()`'s refusal and by the intake's sentence; thirty-odd fixture strings
in five test files that spell the old names; `docs/commands.md`'s paragraph on what
`intent start` refuses; `init`'s `Manual` list of five host settings, none of them a
label; the plugin's `bin/` with one script, the entry point, and a test that it is there
and executable. The host carries `approved` on fourteen issues, twelve open, and no label
beginning `xeno-`.

<!-- xeno:section:scope -->
## Scope

This intent renames both halves of approval to carry the tool's name, gives the label a
script that creates it and a line in `init` that names it, and relabels this
repository's issues.

- **The constants.** `model.ApprovedLabel` becomes `xeno-approved` and `model.ApprovedWord`
  becomes `/xeno approved`; `Approval()` compares the first line the same way, trimmed,
  case ignored, trailing punctuation dropped. Every refusal, sentence and document that
  spells either name says the new one.
- **The script.** `.xeno/plugin/bin/xeno-labels.sh`, shipped with the plugin and vendored
  by `xeno init`, creates the label on GitHub with `gh` or on GitLab with `glab`, with a
  description that says what it starts, and changes nothing where it exists. A test
  beside the entry point's holds that it is there and executable.
- **The line in init.** `InitResult.Manual` gains the label as a setting a person makes
  on the host, naming the script, which is where `init` already lists the branch
  protection and the squash setting.
- **The fixtures and the documents.** Every test fixture that spells `approved` as a label
  or a first line moves to the new names; `docs/commands.md`'s paragraph does the same;
  the intake skill says the comment's form.
- **The register.** A107 for what the clause left to the implementation: the comparison
  on the first line unchanged, the label created and never read back, and the old names
  gone rather than accepted beside the new.
- **This repository's issues.** `xeno-approved` created on the host, added to every issue
  carrying `approved`, and `approved` removed, when this merges; a host act, recorded on
  the pull request, which the maintainer or this session performs since labels take a
  right the host grants.

What it does not do.

**It does not accept the old names beside the new.** A runner that read both would
leave a project unable to say which approval counted; the clause fixes two names and
the constants hold two. Issues approved under the old word need a new comment, which
the pull request says.

**It does not read the label back.** Whether the host carries the label is a question
for the person who runs the script; `enforcement check` compares the enforcement block
and a label is not in it, and adding one is Appendix A.

**It does not add labels for #323's mechanisms.** The analysis said why: trigger labels
are the design section 12 refuses, and none of the three issues has a use case yet.

**It does not touch the sealed trail.** XENO-0278's intake says "the label `approved`"
and XENO-0279's to XENO-0281's sentences quote the word; they record what the names were.

<!-- xeno:section:context-rationale -->
## Why this context

- Issue #332 with its five comments: the question, the four follow-ups, the analysis of
  2026-10-08 and the maintainer's decision of 2026-10-09.
- `docs/process-definition.md` section 12, the clause as changed in `b431e42`, which
  fixes the two names and names the script and the `init` line.
- `internal/model/identity.go` and its test, the constants and the comparison.
- `internal/runner/tracker.go`, `runner.go` and their tests, the refusal and the intake's
  sentence and the fake host that spells the names.
- `internal/runner/init.go` and `init_test.go`, the `Manual` list and the test that
  reads it; `cmd/xeno/main.go`, where the list is printed.
- `internal/mcp/mcp_test.go` and `internal/host/github/tracker_test.go`, the other two
  fixtures that spell the names.
- `internal/plugin/plugin_test.go` and `.xeno/plugin/bin/xeno-env.sh`, for the shape a
  shipped script and its test have.
- `docs/commands.md`, the paragraph on the refusal; `docs/assumptions.md`, A103 as the
  nearest row; `scripts/github-settings.sh`, for the shape of a host setting reported
  and set, which the label script borrows.
