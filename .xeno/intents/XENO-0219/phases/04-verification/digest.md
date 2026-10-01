---
intent: github.com/triplem/xeno#162
phase: 04-verification
created: "2026-10-01T16:04:40Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+94d4c6c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bf253623f8acd73c74157a4f5ace3fe9401bcd1c94b6396acf2be06429d22a66
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
127 cases pass across `internal/git` and `internal/gates`, 18 packages `ok`, `gofmt` and `go vet` clean,
`gate verify` at exit 0 over 193 verdicts. The four commit types were run against this repository's own
history on a copy, over `main..HEAD`: `commit-message` with `conventional-commits` green, which is the
first time a rule has judged this project's commit convention and found it conforming;
`commit-trailer` with `Xeno-Intent` red, correctly, because this project puts the reference in the
footer and section 9 ships that trailer as an example nobody inherits unasked; `commit-signature` red,
correctly, because nothing here is signed and A21 says signing waits for publication; and the same tree
with no range red twice, each rule naming itself with the next step saying the range is an input of the
run. Gaps: no signed commit was ever tested against, so A70 is tested at the boundary and not through
it; the tests skip without git, which is the price of refusing a second dependency;
`section-implies-section` cannot judge section names where no template loads; a trailer's value is
never read; `approver-not-author` compares strings; the range is recorded nowhere, so a red commit rule
is evidence from the run and not from the record; and nothing bounds a range's size.
