---
intent: github.com/triplem/xeno#208
phase: 05-review
created: "2026-10-04T08:47:11Z"
schema_version: "1.0"
runner_version: dev+49f2794.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bc732b8a221d9cfb93ad730806787b3edf48cb78107ea7c2819355be8579429d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      Three, each naming what it departs from: the published report being the package level
      stream names the design sentence that asked for go test -json into a file without
      asking how large that file is, with the two figures that decided it; the summary step
      writing that file as well as printing it names the design's implied second step; and
      the refusal count names the impact section's eight against the twelve paths there are.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      One command and seven flags are added, two fields appear in EvidenceItem, evidenceShape
      becomes exported, and nothing is taken away or behaves differently. The one visible
      difference for a repository that never runs the command is that a declaration carrying
      a format outside section 4's five is now a finding where it was silent, and nothing in
      this trail carries a format, so no verdict changes.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added. go.yaml.in/yaml/v3 is still the one dependency and is not imported
      by the new file, which uses internal/hashing, internal/gates, internal/model and the
      standard library. The workflow's new steps use python3, which three workflows already
      use and the runner image carries.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Three, each naming what it departs from. The
published report being the package level stream names the design sentence that said `go
test -json` into a file without asking how large that file is, and gives the two figures
that decided it. The summary step writing the filtered report as well as printing it
names the design's implied second step and says why one pass over one stream is
preferred. The refusal count names the design's impact section, which said eight, and
the correction is that there are twelve paths; the test table is the authority because
it has one row per path.

**`interface-change-needs-a-migration-note` — met.** One command and seven flags are
added, two fields appear in a struct, one unexported function becomes exported, and
nothing is taken away or behaves differently. A repository that never runs `evidence
declare` sees one difference: a declaration carrying a `format` outside section 4's five
is now a G-Schema finding where it was silent, and no declaration in this trail carries
a `format` at all, so no verdict changes. `gates.evidenceShape` becoming
`gates.EvidenceShape` is internal to the module and has one caller outside its own
package, the new command. The field order of `model.EvidenceItem` changes what a newly
written item looks like and no file that exists: an item is re-marshalled only in the
phase a declaration is being written into.

**`new-dependency-needs-a-rationale` — not applicable.** Nothing was added.
`go.yaml.in/yaml/v3` is still the one dependency and is not imported by the new file;
the writer uses `internal/hashing`, `internal/gates`, `internal/model` and the standard
library. The workflow's new steps use `python3`, which the three scan workflows already
use and which the runner image carries.

**What the checklist cannot answer, and the review did.** The one change to a gate was
checked against the whole trail rather than against this intent: `format` joins a closed
set, and `gate verify` recomputes 336 verdicts at exit 0, so nothing already committed
became a finding. That is the check the review actually turns on, because a gate change
is the one kind of change in this repository whose cost is paid by phases nobody is
working on.

<!-- xeno:section:release-notes -->
## Release notes

**`xeno evidence declare` writes the declaration section 4 defines.** Until now the
block had no writer, so a verification phase that wanted to carry a test report had to
type one into the frontmatter of a file inside `artifacts_hash` — which fifteen phases
of this trail did, pointing at transcripts somebody pasted, until the practice stopped a
week ago.

    xeno evidence declare --intent KEY --phase 04 --kind test-report --job test \
      --result pass --format go-test-json --file go-test.json

Three forms, which are section 4's three states. With `--file` the report is copied into
the phase's `evidence/` and the command computes its `sha256`; there is no flag that
supplies that hash, because the item is bound by it and nothing else. With `--uri` and
`--sha256` an item too large to copy is declared where it lies. With neither, an item a
pipeline has yet to produce declares its kind and its job, which is the pair `xeno
evidence attach` binds the result to when it arrives.

Twelve refusals, all of them before anything is written, and the four about shape are
G-Schema's own check called by the command, so nothing can be well formed on the way in
and malformed in the file.

**`produced_by` and `format` are written for the first time.** Two fields section 4
enumerates and the model never carried. `format` is also judged now, against the five
shapes the section names.

**CI publishes the test report it already produced.** `xeno.yml` writes `go test -json`
to a file, prints a line per package, and uploads the report with a manifest in the
shape the three scan workflows use. The package level stream is what the trail keeps,
eleven kilobytes against the full report's four hundred and eighty-four.

**The pull path of section 6 has run.** This intent's own verification phase declared
four items, finished `provisional`, was pushed, and had all four attached from the
pipeline's artifacts at the review phase's start: the test report and the semgrep, trivy
and npm-audit scans, each with its hash, its run id and the commit it was produced
against, and P4's `artifacts_hash` unchanged by any of it. That is the first attachment
from a real pipeline in this repository, and the first verification phase whose evidence
is a file something other than a person wrote.

**For a reader of a verdict.** G-Evidence now has an input, and it can fail: a declared
report that is edited or removed makes the phase red, which is asserted both ways by
test. Before this it passed on every phase in the trail by having nothing to read.

<!-- xeno:section:residual-risk -->
## Residual risk

**Nothing makes a verification phase declare anything.** The command exists, the pull
path works, and the next intent can finish six green phases having declared nothing at
all. This is #188's residual risk in a second field, and it has the same shape of
answer: a gate that demanded a declaration would be met with a declaration of something
nobody ran, which is what the whole of section 4 is arranged against. What would notice
is a figure, and `xeno intent status` is where the one for the exchange went.

**An attachment's `result` is still judged by nothing**, and fifteen in the trail carry
a word section 4 does not define. Filed as #221 with the three routes out, because the
fix is a decision about fifteen verdicts: `Verify` compares a recomputed status against
the committed one, so a check added later is red on history. The lesson is in P4's
learning record — a closed set gets its check in the commit that introduces the field,
which is the only moment it is free, and this intent could do that for `format` and not
for the attachment.

**`produced_by` is a self-report with no reader.** Nothing would notice a declaration
naming a command nobody ran, which is exactly #207's finding about `model`, `tool` and
`tool_version` arriving in a third field. A-002 records what the field means on a
pending item and the register carries the reading; no check does.

**The round trip needs an open pull request.** The three scan workflows trigger on
`pull_request`, so an intent that declares pending evidence and pushes a branch without
opening one waits for a pipeline that never runs. That is a property of this
repository's workflows, not of the tool, and it is now in P4's gaps rather than in
somebody's memory.

**The test step of a required check changed shape.** `continue-on-error` with the
failure raised as the last step of the job means a failing suite fails the job later
than it did, and the steps between run on a tree whose tests failed. It is the
arrangement `semgrep.yml` has used for weeks and it was exercised on this very pull
request, where `intent verify` failed mid-job as designed for an intent whose P5 was not
written yet; the report was published before it. What has not been exercised is the case
where the suite itself fails, and the step that raises it again is the one thing between
that and a green job.

**The `--uri` form has not met a real artifact store.** It is asserted by test and
nothing in this repository publishes a uri item, because everything published is small
enough to copy. The form exists because section 4 defines it.

**Downloading the artifacts is still a person's job.** `gh run download` four times at
P5, into four directories that `--evidence-from` and `evidence attach` read. WP6
describes a fetch adapter; what exists is the offline stand-in, and the gap between them
is the one step of the pull path that is not a command.
