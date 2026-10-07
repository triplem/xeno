---
intent: github.com/triplem/xeno#273
phase: 04-verification
created: "2026-10-07T10:00:20Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8d718d53a2af982af8c9339a24a02a51803c3c049faa0cc1d3b12b5881f3e8bb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Ten criteria, all passing. A33's five cells were compared before and after by splitting the row
on its pipes: four byte-identical, only the state changed, which is the whole claim and the one
thing a diff of a 94-kilobyte file does not make obvious.

The measurement the row reports is in the evidence with its control: on the original layout a
`strings_hash` corrupted to a hex value turns G-Schema red; with the directories renamed and the
`id:` left alone it is still caught and `gate verify` matches 495; with the `id:` renamed too it
goes green and `gate verify` still matches 495. The code that explains the third is printed
beside it.

The first attempt was a false negative and is recorded as one. 64 zeros in `strings_hash` passed
on the untouched tree as well, because YAML parses an all-digit scalar as a number and
`raw[f].(string)` fails, so the field is skipped. Redone with hex containing letters.

The row carries two figures, 495 and 499, both dated: the second is after this intent's own
phases were sealed, and one number would have been stale before the commit that introduced it.

Nothing outside the register moved. Suite, build, `gofmt`, `vet` clean and `gate verify` matches
499 verdicts.

The gaps are the honest half. A row is not a check and the guard that would be one was ruled out
because it forbids the harmless variant too. The destructive variant stays silent.
The all-digit scalar is still skipped. A33 is reachable by searching rather than by reading the
function that implements it. Variant A's contradiction of section 5 was read off the sentence
and not demonstrated. And six of the ten criteria are a person reading one long table cell.
