---
intent: github.com/triplem/xeno#195
phase: 04-verification
created: "2026-10-03T14:23:34Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+b54626e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ea5acde82962207262a2c3f87b322b0aa72a98003f070796a460fd128782e892
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Every criterion maps to a test, and the one that matters is asserted as an agreement rather than a
literal: `learning.yaml`'s `runner_version` against the `gate.yaml` the same binary wrote beside it,
because a pinned string would pass on a build that stamped neither. The intent level test asserts
not only that the record exists at intent level but that it was *not* also written into a phase,
which is the assertion that would catch a `learningPath` that defaulted. The refusal table is seven
cases — a category outside the four, each missing key, whitespace as a target, nothing at all, and
`--no-finding` beside an entry — each checking the wording and that no file was left behind.
Eighteen packages ok with fourteen new cases, `gofmt` and `go vet` clean, `gate verify` 271 at exit
0. Measured on a throwaway tree first: the command's record carries `0.1.0-dev+b54626e.dirty` where
every hand-written record in this repository carries `0.1.0-dev`, both forms work, all seven
refusals fire, G-Learning passes. And this intent's own trail is the demonstration — every learning
record written by the command, five phases each carrying `G-Learning pass` with a header matching
its verdict, so a phase directory now shows all four artifacts agreeing. Six gaps: `plugin_version`
is still a constant and wrong in all four together (#177); thirty intents keep the typed header,
sealed, so the repository holds two provenances for one field; nothing makes the agent use the
command, since G-Learning checks the header is present and not that it is true; the skill's text is
counted rather than read; the content is unverifiable by design; and `--no-finding` cannot be told
from a phase nobody thought about.
