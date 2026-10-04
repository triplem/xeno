---
intent: github.com/triplem/xeno#208
phase: 03-implementation
created: "2026-10-04T08:38:38Z"
schema_version: "1.0"
runner_version: dev+49f2794.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 511064cfb3dfb60d5f964bbe909fcd3db977d33ef80cc6b0c4f801f7a425261d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The implementation adds one file, touches four and a workflow.
`internal/runner/evidence.go` holds `DeclareEvidence`, which judges the flags against
each other, reads the artifact, refuses a pair the phase already declares, reads and
hashes the report, runs the gate's own shape check and only then writes anything, so
that a refusal leaves no copy in `evidence/` and no half declared item. `declarable`
holds the seven flag combinations no gate can judge; `copyEvidence` refuses to write
over a file whose content differs, because a file there is bound by a hash in some
declaration and nothing in the command can tell whose. `model.EvidenceItem` gains
`produced_by` and `format`, is reordered into section 4's order and keeps a comment
about why there is no `id`; `EvidenceFormats` is the section's five shapes of a report.
`evidenceShape` becomes `EvidenceShape` and judges a given format against that set.
`cmd/xeno/main.go` gains `evidence declare` with seven flags, and the line it prints
names the state the item is in and, for a bound one, the hash nobody typed. `xeno.yml`'s
test step becomes `go test -json` into a file with `continue-on-error` and the failure
raised again as the last step of the job, followed by a summary that prints one line per
package and writes the package level stream, a manifest in the shape the three scan
workflows write, and an upload of both. Fourteen tests in the runner and one in the
entry point, including G-Evidence failing both ways once a declaration exists — the
edited report and the removed one — which is the gate failing for a real reason for the
first time. Three deviations are recorded rather than absorbed: the published report is
the package level stream, 11 452 bytes against 495 629, because an attachment is copied
into the repository and section 4 decides where an item lives by its size; the summary
step writes that file as well as printing it, one pass over one stream; and there are
twelve refusal paths where the design had counted eight, which is a figure a design
phase had no way to know. Build, `gofmt`, `go vet` and the whole suite are clean, and
`gate verify` is at exit 0 over 336 verdicts, so adding `format` to the closed sets
changed no committed verdict. Read in this phase: `exchange.go` for the two helpers,
`gates.go` around `evidenceShape` and `BuildKind`, `hashing.go`, `attach.go`,
`runner_test.go` for the fixture, `main_test.go` for `invoke` and `repo`, and
`semgrep.yml` and `xeno.yml` for the publishing pattern and where it goes.
