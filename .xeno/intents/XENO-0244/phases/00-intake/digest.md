---
intent: github.com/triplem/xeno#221
phase: 00-intake
created: "2026-10-04T09:27:41Z"
schema_version: "1.0"
runner_version: dev+24becc3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dc170d12d8252cf7f8c1b6b81fec391976cfd0330e069c7fdd282fb81b1b5ed6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake establishes the one property the whole intent rests on and records the
decision that was already taken. Section 4 closes the set of an evidence `result` to
`pass` and `fail`; `EvidenceShape` judges a declaration against it since #208, and
nothing judges an attachment, because the shape check reads frontmatter and
`evidence/attached.yaml` is another file. Fifteen attachments in the trail read
`success` — XENO-0107, 0108, 0111, 0121 and 0200 to 0210, all `test-report`/`go-test`,
all `pipeline: local`, all typed by hand on both sides. It is not cosmetic: `build()`
compares an attached result against `"pass"`, so an attached `build-log` carrying
`success` is reported as a build that did not succeed, and G-Test, which would read the
fifteen that exist, is `not-implemented`, so the defect is invisible today and wrong in
both directions the moment either reader has an input. The way in is also still open,
since `Attach` copies the manifest's value without looking at it while already declining
what it cannot bind. What makes a fix orderable rather than impossible: `evidence/` lies
in a subdirectory and `DirHash` does not descend, the recorded `sha256` is of the report
rather than of the word, and `gate.yaml` keeps no copy of the value, checked against
XENO-0210's verdict. So the correction moves no hash and contradicts no verdict, and the
check is addable after it and not before — `Verify` compares a recomputed status against
the committed one. Q-1 was put with four options and settled by the maintainer as D-1 in
the comment thread of #221 before this intent existed, which the learning record notes
as the second such lag in two intents. Scope: correct the fifteen, judge an attachment's
result in G-Evidence, decline an out-of-set result in `Attach`. Out of scope and filed
instead: requiring a result on an attached `test-report` or `build-log`, which is the
same asymmetry one level down; G-Test; the `pipeline: local` provenance, which is
accurate about what happened; and the declarations themselves, which are inside
`artifacts_hash`. Read: #221 with its decision, section 4's evidence subsection,
Appendix B with `DirHash` beside it, `gates.go` for `evidence`, `Collect`, `build` and
`EvidenceShape`, `attach.go` for `Attach` and `unbindable`, `Verify`, all fifteen files,
and the four workflow manifests this intent's own P4 will declare.
