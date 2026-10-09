---
intent: github.com/triplem/xeno#346
phase: 05-review
created: "2026-10-09T15:22:15Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c096cf23f1203411373e9171991d585604c8fdf8223b1451470b059dcb867713
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: P3 records none; the reads and the test strings are as the design named them.
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: The script's arguments, output and exit codes on a first run are unchanged; a second run now does what the header always said.
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: Two flags on tools the script already needed; nothing enters the binary or the pipeline.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**deviations-are-traceable — met.** P3 records none; the two reads and the two test strings are as the design named them.

**interface-change-needs-a-migration-note — not-applicable.** The script's arguments, output and exit codes on a first run are unchanged; what changes is that a second run now does what the header always said. Nothing a project runs or reads moves.

**new-dependency-needs-a-rationale — not-applicable.** Two flags on tools the script already needed; nothing enters the binary or the pipeline.

**What no rule asks.** Whether a shape test is enough for a shipped script. The alternatives section says why a fake host on `PATH` was declined for this intent and when it becomes worth having; the test's comment says what it holds and what it does not.

<!-- xeno:section:release-notes -->
## Release notes

**The label script is idempotent on a repository with more than thirty labels.** `.xeno/plugin/bin/xeno-labels.sh` read the first page of labels and no more, so on a repository past thirty its second run tried to create `xeno-approved` again and exited 1. The GitHub read now asks for the label by name and the GitLab read for a page of a hundred; a second run reports the label as existing and exits 0.

**For a project that met the exit 1.** Nothing to do: the label was created on the first run, and the second run now says so.

**The GitLab bound.** Past a hundred labels the GitLab read misses again; the line in the script says so.

<!-- xeno:section:residual-risk -->
## Residual risk

**The GitLab branch is unproved.** No `glab` here and no GitLab project; the flag and its default are from the documentation of 2026-10-09, and the first run on a GitLab project is the proof.

**A hundred labels on GitLab.** The bound is written, not checked; a project past it meets the fault this intent fixed for GitHub.

**The test holds the flags and not the read.** A fake host on `PATH` would hold the behaviour, and the design says when that becomes worth its harness.

**Belongs to nothing in the specification.** Section 12 names the script and says nothing about its read, which is right; nothing here asks for a clause.
