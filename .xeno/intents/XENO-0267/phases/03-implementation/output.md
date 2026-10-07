---
intent: github.com/triplem/xeno#228
phase: 03-implementation
created: "2026-10-07T06:34:25Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3f920b60852ab2c93fbe492852418050bc810966579e4e750f52e466938fd3f6
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

`internal/runner/next.go`, two suggestions and two comments.

The `not-started` case of `next` now reads "start 01-requirements in a fresh session, so that
its context is what context.lock.yaml says it was given", and keeps `xeno phase start` as its
command. The comment above it says which clauses the sentence rests on — section 6 for the
digest a phase works from, section 5 for what the lock states — and why no harness is named:
section 7 records `XENO_HARNESS` and never branches on it, and a person working with commands
alone has no such command at all.

The last phase's case now ends "...and let the review and the pipeline run. Nothing of this
intent's context is read again, so the next one begins best in a fresh session". It still
carries no command, because nothing it names is one. The comment says what the next intent
starts from instead: the tree and the scope its own intake declares.

The staleness suffix still appends cleanly. `Next` adds ". Its predecessor changed after this
phase started, so read it again first" to whatever text it got, and the new clause ends a
sentence rather than a phrase, so the two read as two sentences and not as one run on.

`internal/runner/runner_test.go`, two assertions.

`TestTheSuggestionMovesOnAndBackWithTheVerdict` already asserted that a decided phase points
at its successor. It now also asserts "fresh session", "context.lock.yaml" and "was given" in
the text, and that `/clear` is not in it. The three strings hold the reason and not only the
phrase, because a hint whose reason has been edited away is followed once; the fourth is the
cheapest guard against section 7's rule being broken here later.

`TestAfterP5TheNextStepIsNotACommand` already asserted that nothing is offered after P5 and
that the merge is named. It now also asserts "fresh session" and "read again", which are the
hint and its own reason — the end of an intent has a different one from the start of a phase.

`docs/assumptions.md`, one row. A98 carries both decisions #228 asks for beyond the sentence,
with the clause behind each: section 13's one vendored plugin for the skill text, and the
plan's WP20 for the figure, along with the harder reason a gate was refused — the only signal
is the cost ledger, which is gitignored local data outside `artifacts_hash` by design. Its
state is `open`, because what the row asks a person for is a yes to a reading rather than a
choice between paths, and the cost half is accepted until WP20 either way.

`go build`, `go test ./...`, `gofmt -l .` and `go vet ./...` all pass. `xeno gate verify` is
P4's evidence and runs there.

<!-- xeno:section:deviations -->
## Deviations from the design

One, and it is a correction to this intent's own plan rather than to anything the documents
say.

**The register row is `open` rather than `approved`.** P1's criterion 10 asked for the row and
said what it carries; it did not say which state, and the states this file uses are a record of
what a person has said. Nobody has said anything about either half yet: the reading is the
runner's, drawn from section 13 and from the plan's WP20. Writing `approved` would have put a
yes in a person's mouth, and that is the one thing the state column exists to report. So the
row says `open`, names what it is asking for — a yes to the reading, not a choice between
paths — and the cost half carries `accepted until WP20` inside it, since that package can
carry the figure whatever the answer to the first half is.

Nothing else departs from P2. No document was edited, no harness is named anywhere, no hook
exists, no gate changed, and the suggestion for a running, red, provisional or stale phase is
untouched.
