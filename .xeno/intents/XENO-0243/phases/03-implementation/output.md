---
intent: github.com/triplem/xeno#208
phase: 03-implementation
created: "2026-10-04T08:37:54Z"
schema_version: "1.0"
runner_version: dev+49f2794.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 511064cfb3dfb60d5f964bbe909fcd3db977d33ef80cc6b0c4f801f7a425261d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

**`internal/runner/evidence.go`, new, the writer.** `DeclareEvidence(key, phase, file,
item)` takes the item as section 4 describes it and the local file separately, because
the file is the input and `path` is a result. The order inside it is the one the
refusals depend on: the flags are judged against each other, the artifact is read, the
pair is checked against what the phase already declares, the file is read and hashed,
the gate's shape check runs, and only then is anything written. A refusal therefore
leaves no copy in `evidence/` and no half declared item, which is asserted by the table
of twelve refusals ending on the assertion that the phase holds neither.

`declarable` holds what no gate can hold: the ways of arriving at a declaration that
cannot be written. `--file` with `--uri`, a hash beside the bytes, a `uri` with nothing
binding it, a hash with nothing it belongs to, a pending item with no job, and a pending
item carrying a result. Everything about the item itself is `EvidenceShape`'s.

`copyEvidence` writes the report beside the artifact and refuses to write over a file
whose content differs, because a file in `evidence/` is bound by a hash in some
declaration and nothing here can tell whose. Identical content is not a conflict: the
declaration being written binds exactly the bytes already there.

**`internal/model/model.go`.** `EvidenceItem` gains `produced_by` and `format` and is
reordered into section 4's order, with a comment saying why there is no `id`.
`EvidenceFormats` is section 4's five shapes of a report, as a closed set beside the two
that were already there.

**`internal/gates/gates.go`.** `evidenceShape` becomes `EvidenceShape`, for the reason
`QuestionShape` was exported, and judges `format` against the new set where one is
given. An absent format is not a finding; the field is optional in the section that
defines it.

**`cmd/xeno/main.go`.** `evidence declare` with seven flags, and the usage line in two
rows because the three forms do not fit one. `--file` is the report itself here and is
the only command where that flag means something other than a message to read; `--from`
was the alternative and is the attach's word for a directory. The line printed says
which of the three states the item is in, and for a bound item it prints the hash,
because that is the thing that binds and nobody typed it.

**`.github/workflows/xeno.yml`.** The `verify` job's test step becomes `go test -json`
into a file, with `continue-on-error` and the failure raised again as the last step of
the job. A summary step prints a line per package and writes the package level stream
beside the full one; a manifest step writes the entry in the shape the three scan
workflows write, and an upload publishes both. The full report is half a megabyte, an
attachment is copied into the repository, and section 4 decides where an item lives by
its size, so what the trail keeps is the package level record: that the run happened,
with which command, and what it returned, per package. The pattern and its reason are
`semgrep.yml`'s.

**Tests.** `internal/runner/evidence_test.go`, fourteen of them: the three forms, the
hash being the hash of the content, the copy landing beside the artifact, G-Evidence
passing and then failing both ways when the report is edited and when it is removed, the
pending item being bound by the attach, provenance written and never defaulted, the
twelve refusals, the pair declared once, the overwrite refused, the artifact that does
not exist yet, the stale verdict after a declaration, and the body and frontmatter left
untouched. One of them calls `gates.EvidenceShape` directly to assert that the command's
refusals are the gate's own judgement and not a second opinion. `cmd/xeno/main_test.go`
gains the command line form of the two common cases and all three steps of the exit code
staircase.

<!-- xeno:section:deviations -->
## Deviations from the design

**The published test report is the package level stream and not the whole one.** The
design said `go test -json` into a file and did not ask what size that file is. Measured
on this tree: 495 629 bytes over 2 293 lines for the full stream, 11 452 bytes over 90
lines for the events that name no test, which are the ones carrying each package's
result. An attachment is copied into the repository and would arrive in every intent
that declares a test report, so the full stream would put half a megabyte of log text
inside `artifacts_hash` for a record that says what twelve kilobytes say. Section 4
decides where an item lives by its size, and this is that sentence applied one level
down, to which part of a report is worth keeping. Both files are uploaded, so the
pipeline's reader loses nothing; the manifest names the small one. The cost is stated
rather than hidden: the trail says which packages ran and how they ended, and a reader
who wants the output of a particular test needs the pipeline run while it is still in
retention.

**The summary step also writes the filtered report.** One step, two outputs, where the
design implied two steps. It is one pass over the stream and splitting it would mean
parsing the same file twice; the step's name says what a reader sees and its comment
says what it writes.

**Twelve refusals, where the design counted eight.** The design's impact section said "1
on each of the eight refusals", which counted the flag combinations and treated the
shape check as one. The implementation refuses on twelve distinct paths: seven flag
combinations, the phase with no artifact, the pair already declared, the report that
cannot be read, the copy that would land on other content, and the shape check, whose
four causes are the kind, a missing result, a result outside the set and a format
outside it. The count in the design was wrong rather than the design: nothing was added
that it did not describe, and the test table is the authority because it has one row per
path.

**Nothing else.** The verb, the three forms, the single authority on shape, the hash
computed and never given, the copy and the declaration as one act, the two provenance
fields, the field order, and the silence about `gate.yaml` are as the design phase
decided them.
