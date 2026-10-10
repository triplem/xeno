---
intent: github.com/triplem/xeno#334
phase: 04-verification
created: "2026-10-10T16:01:35Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8f88e6c3c30dcd672ba8341f312a4023d3794cf11cee0c3d6d07107f431c1fe5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Verification of a document, which has no suite, so the work was to say for each criterion
what kind of answer it got and to put an address on every claim the section makes about
somebody else's repository.

Twelve of fourteen criteria met, one met in its operative clause only, one not answerable by
this tree. The four gates `CLAUDE.md` names are green: `gofmt -l .` outside `vendor/` prints
zero lines, `go vet ./...` exits 0, `go test ./...` exits 0 — run twice, because the first
run's exit code was lost and twenty `ok` lines with no `FAIL` is not an exit code — and
`./xeno gate verify` exits 0 over 607 verdicts. The `gofmt` pipeline's own status is 1
because `grep` matched nothing, which is the trap worth naming rather than the result.

The cell-by-cell table is the substance of this phase. Every AI-DLC claim in section 10 has
an address at commit `6a378b53c0a4fe0641ed7d8de8dfff94264d5b6a`, read from a tarball of that
exact commit extracted to disk, so each line number is a line number in the pinned tree.
Fifty-odd claims, including the two contradicting sentences at `:140` and `:101`, the four
places in `core/tools/aidlc-state.ts` and `core/hooks/aidlc-run-sensors.ts` that settle which
is current, and line 5 of each of the six shipped sensor manifests.

Three figures from #323 moved and the section carries the current ones: 115 audit event
types rather than 117, 5,123 stars rather than 5,084 and dated to the day, and a third row
that is a settled account rather than "sensors that may be advisory". One claim from #323 did
not reproduce and is therefore not in the section: nothing at the pin says a later agent may
revise an earlier artifact. What the tree says instead is that reviewed outputs remain frozen
and that a person's jump back offers keep, modify or redo, so the fourth row is written from
those two and makes the weaker, checkable claim.

Four negative claims, each probed with a positive control in the same command and both
answers recorded. Nothing in `core/` binds an intent to a tracker — one hit, which is a prose
citation of AI-DLC's own issues, against a control of five files for `intents.json`. Nothing
in `core/` hashes the record — no hits, against a control of three files for `Unit Source
Fingerprint`. No file here is misformatted — zero lines, against the same binary printing the
path of a deliberately misformatted file. Nothing normative moved — a zero-line diff, against
`334 9` for the file that did change, in the same breath.

Criterion 9 is met in the clause that required something and not in the sentence that
illustrated it. The section names four mechanisms, takes none, and points at #340. It does
not point at #339, AI-DLC's reviewer agent, which is a real omission; and it should not point
at #338, which came from the demo of 4.1 and not from this reading. The omission is recorded
as a gap rather than repaired, because P3 is decided and the document committed at `295bcc4`,
so an edit now stales P3 and P4 is already running and cannot be started again to pick up a
redo. What is lost is one pointer and #339 is still reachable through #323.

Five gaps. The missing #339 paragraph. The hosted sample's own five rows, unread by anybody.
The section being a claim about one commit in a project that moves faster, with no repair for
the day AI-DLC fixes its own contradiction and the third row becomes an argument about
something no longer there. Nothing re-checking the width of this page's prose. And #334's
missing work-package label, which is written down in this trail and nowhere a reader of the
plan would find it.

One learning, about this phase's own discovery: an acceptance criterion is one sentence, and
an illustration beside the requirement reads as part of it.

Three sections of the five this template defines. No open question and no decision in this
phase.
