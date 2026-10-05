---
intent: github.com/triplem/xeno#212
phase: 03-implementation
created: "2026-10-05T16:03:25Z"
schema_version: "1.0"
runner_version: dev+97d07da.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 219b2f41d75e353c56630368345bf6f08b2c77142e2045c18cde328a3598c962
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

`internal/gates/gates.go`. The table row becomes `{"G-Test", 4, testReport}`. `TestKind` is a
new constant beside `BuildKind`, with the comment pointing at that one's history rather than
repeating it. `declaredResults(c, kind, label, repair)` is the shared body and `build` and
`testReport` are its two callers; `testReport` carries the paragraph explaining that the mapping
half has no reader, with the figures and the reason.

The repair string is each caller's rather than one phrase for both. G-Build keeps "fix the build
and let the pipeline produce a new result" exactly as it was, and G-Test says "fix the suite",
because a build and a suite are fixed in different places. A first pass shared one weaker phrase
for both and that was a silent degradation of G-Build's message, which nothing asserted — it
was caught by reading the diff rather than by the suite.

`internal/gates/evidence_test.go` gains five tests and `TestKind` on the `declaration` struct,
asserted against the document the same way `BuildKind` is: a rename in
`testdata/evidence-declaration.yaml` now fails here rather than making the new gate inert, which
is the defect that file exists to catch.

`internal/gates/testdata/evidence-declaration.yaml` gains `test_kind: test-report`, and the
comment above it now explains both constants rather than one.

`docs/clause-readers.md` gains two rows for one sentence of section 7, because its halves have
different readers; the tool-requirement count moves 35 to 37, and two paragraphs say the rows
arrived in #212, why a green G-Test now means half its row, and why the mapping half has no
reader — with the figures and the date.

#250 carries the mapping half, with the measurement and the three things deciding it involves:
whether criteria get a numbering convention, whether the check belongs in a gate at all or in a
`review` rule, and whether the trail is re-judged or the check is forward-only.

The trail was verified again after the real change rather than resting on the probe: `gate
verify` at exit 0 over 384 verdicts, which is the 381 that existed plus this intent's three
judged phases.

<!-- xeno:section:deviations -->
## Deviations from the design

One deviation from the design, which said the two gates differ by "a constant and a word". They
differ by a constant and two strings. The first pass gave both the same weaker next step, "fix
it and let the pipeline produce a new result", which quietly downgraded G-Build's message from
"fix the build" — a behaviour change to an existing gate that no test asserted and the suite did
not catch. The repair is now the caller's, G-Build's wording is byte-identical to what it was,
and the design's sentence was wrong about the shape of the thing being shared.

One correction inside this phase rather than a deviation from the plan.
`TestAPendingTestReportIsPendingAndNotAFailure` was first written with the `sealed` fixture and
failed, because `sealed` supplies a sha256 and `Pending()` is `SHA256 == ""` — so the fixture was
a sealed item with no result, which is a failure and not a pending one. The test asserted the
wrong thing and the code was right. The fixture now declares no hash, which is how every
test-report in this trail is written, and the comment says so.

That mattered beyond the test. It is the second time in this intent that the pending path was
misread: P0's audit script read the declaration's own `result`, found seventeen empty, and
predicted the result half could not be implemented. The probe disproved it and the real reason is
now checked rather than asserted — 17 of the 18 declarations are attached with `result: pass`,
which was confirmed by reading `evidence/attached.yaml` rather than inferred from `gate verify`
passing.

One thing found and left, because it belongs to the intent that caused it rather than to this
one. `docs/clause-readers.md` line 84 reads "a question carries two to four options and one free
entry | G-Questions", and #229 made that clause's reader larger — the consequence is read by
G-Schema's shape check and the recommendation by the writer. That row is now incomplete and
#229 did not update it, which is the same miss this intent is adding two rows to fix for a
different clause. Noted rather than absorbed: it is one line, and putting it in this commit would
hide a finding about the previous one.
