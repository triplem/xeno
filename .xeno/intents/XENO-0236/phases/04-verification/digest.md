---
intent: github.com/triplem/xeno#183
phase: 04-verification
created: "2026-10-03T19:08:44Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cf91720178eae3b115b4fddd1933d4db0a5d1c2cb88e3728d98b21a3c659874f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Most criteria are held by reading a file, which is what a change to a document means, and two by a
command: `grep -c 'XENO_PLUGIN_ROOT\|plugin-root' docs/implementation-plan.md` is 0, and the test
asserting the entry point exports no plugin root still fails if it does — the only part of this
change with a mechanism behind it. The specification commit is two files under `docs/`, 32
insertions, 11 deletions, no `.go` and no `.yml`. The straggler grep finds five files mentioning the
mechanism and every one explains its removal. Eighteen packages ok, `gofmt` and `go vet` clean,
`gate verify` 285 at exit 0 — which is the criterion a document change with no mechanism behind it
should be judged by, and the one that would have caught an accidental behaviour change. Nothing
needed measuring, because there is no behaviour to exercise; the two measurements the argument rests
on were taken in the exchange that produced the decision and are in the document rather than only
here. Six gaps: P3's justification for the comment rewrap cites a convention the maintainer retired
two exchanges later, and the sealed phase keeps it, which is correct and is what judging a phase
against what was in force means; nothing prevents the clause being reinstated except that the
mechanism would have to arrive with it; `XENO_PLUGIN_DATA` is still listed and read by nothing, so
the list still holds one entry in the state the removed one was in; G-Complete still runs only at P5
and is still nobody's issue; the condition kept in section 7 has no test and cannot have one, which
is a category #202 warns against counting as a gap and which this intent added one of deliberately;
and the removal was argued from two measurements taken here, which the document carries so a later
reader cannot over-read it as a general claim.
