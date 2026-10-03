---
intent: github.com/triplem/xeno#202
phase: 04-verification
created: "2026-10-03T19:28:32Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5a45d2c4deb4f0d087fe3ce0a8f46bf295c22541acc7721f5096508fe7589daa
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

The deliverable is a document, so there is no unit test of it. What can be checked is
checked, and each check maps to one acceptance criterion.

"A reader is looked up, not recalled" maps to a lookup of every symbol the file names:
`ResultRequiredKinds`, `PhaseExcluded`, `FindingID`, `predecessorAllowsStart`,
`rewriteStatus`, `notImplemented`, `IntentClose`, `Decided`, plus the claim that
`exec.Command` appears outside tests in two packages.

"The list is a file in the repository" maps to the file being tracked by the commit and
`gate verify` recomputing the trail that contains it.

"Prose wraps at 88, tables do not" maps to a width check over the file excluding table
rows.

The project's own suite — build, test, vet, gofmt, `gate verify` — is run because the
commit touches the repository, not because the document could break it.

<!-- xeno:section:results -->
## Results

Every symbol the file names as a reader exists in the tree at `d19a1ca`. The lookup
returned a definition for all eight, and `exec.Command` outside tests appears in exactly
`internal/git/git.go` and `internal/external/external.go`, which is what the first
architectural row claims.

The width check over `CLAUSE-READERS.md` reports no prose line over 88 characters. Two
reflows were needed: the first draft went in from a heredoc unwrapped, and the dating
sentence added after the requirements phase overran by one line.

`go build -o xeno ./cmd/xeno` succeeds. `go test ./...` is ok for all twenty packages
with tests. `go vet ./...` is silent. `gofmt -l .` outside `vendor/` prints nothing.
`./xeno gate verify` reports "verified 301 verdicts" and exits 0.

G-Supply reports `not-implemented` on every phase of this intent, which A86 says is the
state of every artifact written by `go build`.

<!-- xeno:section:gaps -->
## Gaps

Nothing checks that the file stays true. A reader named by symbol is wrong the day the
symbol is renamed, and no test compares the file against the tree. This is the same shape
as the problem the file documents, it is stated in the file's own opening, and A90 records
why a gate was rejected rather than deferred.

The counts are a reading, not a parse. The extraction was mechanical and the placement of
each of 219 sentences into one of four kinds was a judgement, so a second pass would
likely move a handful of sentences between "addressed to a person" and "explanation". The
two enumerated kinds are the ones that were checked individually.

The evidence of this phase is not attached. Section 5 declares evidence in the frontmatter
of `output.md`, which the runner writes and which sits inside `artifacts_hash`, and no
command declares an item — so attaching would mean editing a hashed artifact by hand. The
runs are named above instead. Every verification phase in this trail is in the same state,
and it is a finding about the plan rather than about this intent.

Section 12's triple being self-reported, G-Complete running only at P5, and
`XENO_PLUGIN_DATA` having no reader are findings of the pass and are not addressed here,
by the non-goals.
