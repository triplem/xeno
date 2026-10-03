---
intent: github.com/triplem/xeno#172
phase: 04-verification
created: "2026-10-03T10:39:22Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8d484db13807a7920c9c83f16c11d5a057989aaa9e36f72a11c76e282569ccc0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
431 cases pass, `gofmt` and `go vet` clean, `gate verify` at exit 0 over 228 verdicts. The real mistake
was made on a copy and reported: a profile with one link whose document was typed
`docs/process-defintion.md` — the transposition somebody actually makes — produced a finding naming the
profile, the component and the path as written, so a reader sees the typo rather than being told to go
looking. Before this intent the same profile produced nothing at all. The base is unchanged: #171's
test for a declared link's document joining it last with its hash still passes, which is the criterion
this intent was most at risk of breaking, since the check and the resolution read the same field for
different purposes. It does not block. Gaps: the byte count still moves after a phase is sealed, which
is the half this intent cannot fix and which needs one clause of section 5; a link with no component
produces an awkward sentence, where the alternative would leave its document unchecked; nothing checks
that a component means anything, so half a link is validated and half is not, deliberately; a
document's existence is all that is checked, since section 5 forbids inferring the mapping; and there
is still no profile here, so the experiment this unblocks is the first thing to read these findings in
anger.
