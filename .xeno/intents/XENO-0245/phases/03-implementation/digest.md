---
intent: github.com/triplem/xeno#217
phase: 03-implementation
created: "2026-10-04T20:59:15Z"
schema_version: "1.0"
runner_version: dev+fbf8a72.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fa4eebc3bd6e89347edd11cd090532cd5b2dc1b141e15f80714cc0331cd891fe
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Implementation follows P2's order and the order mattered. The rename went first:
model.ContextScope carries context-scope.yaml, model.Scope is the type, five readers
follow, and the comments, locals and test names follow the document rather than keeping
a name it no longer uses. docs/v2-delta.md is untouched by design. Renaming this
intent's own artifact re-judged P0, which staled P1, which staled P2, and P3's lock was
re-resolved after that: the design called it one re-judge and it is a chain the length
of the phases already finished, paid while nothing is merged and recorded as a
deviation. internal/runner/scope.go is new and holds ScopeSet, readScope and
resolveScope. ScopeSet decodes into model.Scope with KnownFields(true), refuses a header
a person wrote by A35, refuses an empty include because a scope resolving to nothing
reaches every reader as the absent one it replaces, writes the artifact with the
runner's header, and returns what the patterns resolve to. readScope and resolveScope
are the two halves informationBase was, extracted so the writer and phase start resolve
once and cannot disagree, and informationBase is now the two of them in nine lines with
the walk moved unchanged. The refusal is the first statement in Finish, at P0, before
writeDigest and writeCost, so a refused finish writes nothing; its test asserts that by
the summary text being absent from the digest rather than by the file being absent,
because the fixture writes a digest when it writes the artifact. Section 5's second
limit is i < idx in staleReads, one character and the whole of what unblocked this
intent, with the specification's sentence in the comment and a line asking the next
reader not to restore the inclusive bound. Its first limit is changedPaths, a helper
turning Ctx.Base and Ctx.Head into the set of paths the change touched through
git.Paths, with nil and empty kept apart deliberately: nil stops the check because there
was no change under review to narrow to, and empty is a real answer that finds nothing
stale, which is A74's distinction applied to an input. xeno scope set is in the table
with needsKey and not needsPhase, and in the help beside section set. Both fixtures now
carry a scope, with the vendored plugin.json's argument: without it 88 tests were
refused, which is the measure of how many tests finish a P0 and of what a gate-based
requirement would have cost instead. Four tests cover the two limits, in internal/gates
where the gitRepo helper already was, and five cover the writer and the refusal, with
one more in cmd/xeno for the command end to end. One existing test changed rather than
being deleted: TestGivenFilesAreComparedAgainstTheTree asserted that a phase changing a
file its own lock lists goes red, which is what section 5 calls working rather than
stale, and its comment now says it asserted the opposite until #236. One was renamed
because its premise became unreachable: a repository with no scope cannot finish a P0,
so the silence belongs to a scope that resolved and came out empty. Figures: gofmt and
go vet silent, eighteen packages green, gate verify at exit 0 over 348 verdicts before
and after, and this intent's own P1 green with no release on it, which is the criterion
P1 set itself. Four deviations recorded, among them that xeno scope set has deliberately
not been run against this repository's own intent, because the header it writes carries
a fresh created and would start the cascade again for no gain.
