---
intent: github.com/triplem/xeno#217
phase: 03-implementation
created: "2026-10-04T20:36:45Z"
schema_version: "1.0"
runner_version: dev+fbf8a72.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fa4eebc3bd6e89347edd11cd090532cd5b2dc1b141e15f80714cc0331cd891fe
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

**The rename, first and mechanically.** `model.ContextScope` carries
`context-scope.yaml`, `model.Scope` is the type, and the five readers follow at
`runner.go`, `gates.go` three times and `model.go`. The comments and the local variables
follow the document rather than keeping a name it no longer uses, and the test names
with it. `docs/v2-delta.md` is untouched, as the design said: a different document and
not in this change. This intent's own artifact was renamed and P0 re-judged, which
staled P1's predecessor hash and then P2's, so both were re-started and re-judged in
turn — the cascade the design predicted, paid while nothing is merged.

**`internal/runner/scope.go`, new.** `ScopeSet` decodes the entry into `model.Scope`
with `KnownFields(true)`, refuses a header a person wrote, refuses an empty `include`,
refuses a P0 that has not started, writes the artifact with the runner's own header, and
returns what the patterns resolve to. `readScope` and `resolveScope` are the two halves
`informationBase` was, extracted so the writer and `phase start` resolve once and cannot
disagree; `informationBase` is now the two of them, nine lines, and the walk moved
unchanged rather than being rewritten in the move.

**The refusal, first statement in `Finish`.** At P0, where the scope is absent, with the
reason naming section 5 and the remedy naming `xeno scope set`. Before `writeDigest` and
`writeCost`, so a refused finish writes nothing: the test asserts that by the summary's
text being absent from the digest rather than by the file being absent, because the
fixture writes a digest when it writes the artifact.

**Section 5's second limit: `i < idx`.** One character in `staleReads`, and the whole of
what unblocked this intent. The specification's sentence is in the comment above it, in
its own words, with a line asking the next reader not to restore the inclusive bound as
a tidy-up.

**Section 5's first limit: `changedPaths`.** A new helper beside `staleReads` turning
`Ctx.Base` and `Ctx.Head` into the set of paths the change under review touched, through
`git.Paths`, which the gate path already uses for `Commits`. A lock entry outside the
set is skipped. nil and empty are kept apart deliberately and the comment says so: nil
is "no change under review to narrow to" and stops the check, empty is "the range
touched nothing" and is a real answer. That is A74's distinction applied to an input.

**`xeno scope set`, in the table and in the help.** `needsKey` and not `needsPhase`,
beside `section set` in the help text for the command it most resembles. It prints the
figure and the sentence about the one moment a budget can be set, because that sentence
is what this intent's own P0 learned the hard way.

**Both fixtures now carry a scope.** `newFixture` in `internal/runner` and `repo` in
`cmd/xeno` each write one, with the same argument the vendored `plugin.json` already
carries there: every intent has one from this commit, so a fixture without one is a case
to assert on its own rather than the condition of every other test. Without it 88 tests
were refused, which is the measure of how many phases finish a P0.

**Four tests for the two limits**, in `internal/gates/staleness_test.go`, using the
`gitRepo` and `gitCommit` helpers that were already there for the commit predicates: the
gated phase's own lock is not read; a preceding phase's moved input is a finding that
names the phase and the file; only what the range touched is reported, with the
untouched file moved first and committed so that the range genuinely excludes it; and no
range reports nothing.

**Five tests for the writer and the refusal**, in `internal/runner/scope_test.go`: the
artifact written with the runner's header and the figure reported; the reported figure
equal to the one the lock records; three refusals in a table; the P0 refusal naming the
artifact and the remedy and writing neither verdict nor digest; and a phase after P0 not
being asked for a scope it never has. One more in `cmd/xeno` for the command end to end.

**One existing test changed rather than deleted.**
`TestGivenFilesAreComparedAgainstTheTree` asserted that a phase changing a file its own
lock lists goes red, which is the behaviour section 5 calls working rather than stale.
It now asserts green, and the comment says it asserted the opposite until #236 and that
this is how the defect lived. `TestChangedSinceIsSilentWithoutAProfile` became
`TestChangedSinceIsSilentWhereTheScopeMatchesNothing`: its premise, a repository with no
scope, is unreachable now that a P0 cannot be finished without one, so the silence
belongs to a scope that resolved and came out empty.

**The figures.** `gofmt` and `go vet` silent, eighteen packages green, `./xeno gate
verify` at exit 0 over 348 verdicts before the change and after it. This intent's own P1
is green with no release on it, which is the criterion P1 set for itself: the two
findings it was given were raised only because the check read its own lock.

<!-- xeno:section:deviations -->
## Deviations from the design

**The scope's own cascade was paid three times, not once.** The design said renaming
this intent's artifact would re-judge P0 and called it free. It re-judged P0, which
staled P1, which staled P2, and P3's lock had to be re-resolved after that. Each step
was mechanical and none needed a release, because a re-started phase re-resolves its
lock against the committed tree, but the design understated it as one re-judge rather
than a chain the length of the phases already finished. The cost is recorded here
because the next intent to rename an artifact of its own will pay the same chain, and
#237 is where the awkwardness lives: a phase cannot re-resolve a lock without its
verdict being removed first.

**88 tests were refused by the refusal, which the design did not predict.** It named the
two fixtures as a consequence and not as a number. The number is the interesting part:
it is how many tests finish a P0, and therefore how much of the suite would have had to
change if the requirement had gone into a gate instead, where no fixture change could
have answered it.

**The staleness half now has no coverage outside a git repository.** Both limits need a
commit range, so the tests moved to `internal/gates`, where the `gitRepo` helper already
was, and they skip where git is not on the path, as that helper does. The runner's own
test for the information base keeps the half of its assertions that need no range.
Nothing is untested, but the coverage sits one package away from the lock it is about.

**`xeno scope set` has not been run against this repository's own intent.** The command
is tested in two packages and was deliberately not used to rewrite XENO-0245's own
scope: the header it writes carries a fresh `created`, which would change P0's
`artifacts_hash` and start the cascade above for no gain. The first real use is the next
intent, which P1 named as the measure of whether the writer is usable.
