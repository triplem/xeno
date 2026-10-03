---
intent: github.com/triplem/xeno#210
phase: 04-verification
created: "2026-10-03T19:43:46Z"
schema_version: "1.0"
runner_version: dev+9fcc639.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7a897d9b87abbf8562ef136ef587e18b791f6fbeb7f8eb186bd2fa87505bc7ed
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Three claims in the new wording can be checked, and each maps to one measurement.

"Go source has no width rule beyond `gofmt`" maps to running `gofmt -l` over a file with a
112-character line: it is silent and exits 0. The formatter does not wrap, so a width rule
for Go cannot come from it.

"Markdown prose wraps at 88; tables and code blocks do not" maps to a width pass over every
tracked Markdown file, counting lines outside tables and fenced blocks.

"`CONTRIBUTING.md` explains" the 72 maps to reading the passage rather than trusting the
pointer: a convention that delegates its reason is only as good as the delegation.

The project's suite is run because the commit touches the repository, not because a
paragraph could break it.

<!-- xeno:section:results -->
## Results

`gofmt -l` is silent on a file whose third line is 112 characters, and exits 0. A width
rule for Go has no reader and could not have one without a tool this project does not
carry.

The width pass at `9fcc639`: **1 772 Go lines over 88, in 58 of 64 files**, the longest
217 characters, in `internal/runner/runner_test.go`. Outside `.xeno/`, **90 Markdown prose
lines over 88 in 7 files** — 35 in the implementation plan, 27 in the process definition,
10 in the v2 delta, 9 in `ASSUMPTIONS.md`, 4 in the README, 4 in the embed placeholder, 1
in `SUPPLY-CHAIN.md`. Every one of the 62 in the two normative documents is at 89 or 90.
Sealed phase artifacts hold a further 9 732 such lines across 406 files, written before the
prose convention settled and immutable now.

`CONTRIBUTING.md:63` carries the reason, under "Write the description at 72 characters":
the description becomes the squashed commit body and GitHub rewraps it to that width, "so
anything wider is reflowed into something nobody wrote". The pointer is good.

The suite: `go build` succeeds, `go test ./...` is ok, `go vet` is silent, `gofmt -l`
outside `vendor/` prints nothing, and `./xeno gate verify` reports 307 verdicts at exit 0.

<!-- xeno:section:gaps -->
## Gaps

None of the three widths is checked. That is the deliberate position of A91 and the reason
is A90's: a check that cannot fail is worse than none, and both candidate checks need a
decision first — which of 1 772 Go lines get rewrapped, and whether a Markdown pass can
tell a long link from prose. So this intent replaces a wrong statement with a right one and
leaves all three rules in the hands of whoever is reading.

The 90 live prose lines over 88 are not fixed. The new wording makes them violations of a
rule that is now correctly stated, where before they were violations of one that was wrong,
and rewrapping two normative documents by one or two columns per line is a change to files
the agent cannot edit.

The 9 732 sealed lines cannot be fixed at all. They sit inside `artifacts_hash`, and A74
judges a phase against the rule set its own artifact records, so reflowing them would
rewrite the trail to satisfy a convention that postdates them.

The Markdown figure is sensitive to how the pass treats fenced blocks and table rows. The
numbers above exclude both; a different reading would move the README's four and the
placeholder's four, which are inside or adjacent to usage blocks. The 62 in the two
normative documents are prose and are not sensitive.
