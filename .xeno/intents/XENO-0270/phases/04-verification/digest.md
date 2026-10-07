---
intent: github.com/triplem/xeno#231
phase: 04-verification
created: "2026-10-07T09:50:11Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6f6d371592ec84b7f94c9ef537209cb49c657e986ebffc9260a74b5b90ef9474
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Twelve criteria, all passing: seven read by a person, because the deliverable is a document,
and five checks.

The reference test was checked both ways: one line of the page's block changed, the test failed
naming the page, restored and it passed. Nothing is lost, counted rather than claimed: 14 of
the old README's 19 paragraphs are carried verbatim and the other five are accounted for one by
one. The command list is strictly larger, 27 against 21, with the six that had drifted now in
it and nothing dropped. The index covers all ten other files under `docs/`, both normative
documents carry the word in their own entry, and all eighteen relative links across the four
files resolve against the tree.

Two claims in the new prose were wrong and were corrected before this phase: the index said
starting a phase is the one command that may use the network, where section 12 says everything
else in the CLI may; and the two v2 drafts were described as retrospective comparisons where
both are architecture records. Opening the files is what caught them.

`go test ./...` passes across 20 packages, build, `gofmt` and `vet` clean, `gate verify`
matches 499 verdicts, and no normative document is among the files touched.

The gaps are the honest half. Whether the README reads well is a judgement nothing checks.
Three of the four files are held by nothing, and a file added under `docs/` tomorrow will not
appear in the index silently. The per command notes are the same unheld second copy the intent
set out to remove, one level down. The coverage table moved unexamined, deliberately. WP16
inherits three pages and a guard it may read as the design. And two wrong claims were written
inside this intent, caught by a method nothing enforces.
