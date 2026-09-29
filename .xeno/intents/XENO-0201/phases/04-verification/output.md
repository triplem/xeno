---
intent: github.com/triplem/xeno#120
phase: 04-verification
created: "2026-09-28T20:35:17Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 9f98e596b5914c1541794cfd924dbcc3b90fda4594cfa045bd4bf12ba4e5a57c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
evidence:
  - kind: test-report
    job: go-test
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it |
|---|---|
| AC1 `phase finish --summary` writes the digest | `TestFinishWritesTheDigestFromTheSummary`, and this intent's own 03-implementation digest, which the runner wrote |
| AC2 the fields it carries, and the three it must not | the same test asserts all ten present and `template`, `strings_hash` and `rules_hash` absent; `TestTheDigestCarriesTheFieldOrderOfSectionFive` asserts the order |
| AC3 no `secrets_hash`, and the gate says so | `TestTheWrittenDigestCarriesNoSecretsHashAndTheGateSaysSo`, which asserts both the absence and the finding |
| AC4 no summary, no change | `TestFinishWithoutASummaryWritesNoDigest`, and every pre-existing test in the package passing with an empty summary |
| AC5 a written digest reaches the same verdict as a typed one with the same content | this intent's 03-implementation: red on `secrets_hash` and `tool_version`, green once a hand filled those two. Argued from that rather than asserted in a test |
| AC6 `model` and `tool` in `output.md` | `TestSectionSetWritesTheModelAndToolTheProjectRecords`, and this intent's own output files |
| AC7 no `agent` block guesses nothing | `TestNoAgentBlockLeavesBothFieldsAbsent`, over a missing file, a missing block and an empty one |
| AC8 A35 names three fields | read: the row lists `tool_version`, `secrets_hash` and `rules_hash`, and records what gained a writer |
| AC9 everything green stays green | `go test ./...`, `gate verify` over 85 verdicts, and the sixty hand written digests untouched |

**What could not be measured.** The four intents before this ran their new tests against the tree
without the change and reported which failed. That is not available here: `Finish` gained a
parameter, so the tests do not compile against the old runner, and the transcript carries the
compiler error rather than a row of failures. What the tree shows instead is stronger than a test
run: no digest in this repository was written by a runner before this change, because there was no
writer to do it.

<!-- xeno:section:results -->
## Results

`go test ./...` passes on every package, including every test written before this
change, which is AC4: `Finish` with an empty summary behaves as it always did.

`gofmt -l .` outside `vendor/` prints nothing, `go vet ./...` is silent, and `./xeno
gate verify` recomputes 85 verdicts and matches.

The result worth recording is that the first digest the runner ever wrote turned its
phase red. It omits `secrets_hash` and `tool_version`, G-Schema requires both of a file
produced in a session, and the phase went green only after a hand filled those two. That
is the writer working: the hand written digests of sixty intents were green because a
hand fills what the tool refuses to invent, and the honest writer produces a file the
gates reject until those two fields have sources. The learning of P3 records it, and it
is not a regression to be fixed but the shape of a half closed gap.

What did close: the frontmatter of a digest, its field order, its body and eight of its
ten fields are now produced by the runner, and `model` and `tool` appear in `output.md`
from the same source. Of A35's five writerless fields, three remain.

Read rather than executed: the A35 row and the usage line. No gate reads either.

<!-- xeno:section:gaps -->
## Gaps

**The runner writes a digest it cannot make green.** `secrets_hash` needs the filter,
`tool_version` needs the harness, and until both exist a digest written by `phase
finish` is red on two fields and a hand has to finish it. The gap #120's scoping named
is half closed, and the half that remains is the half that needs new work packages
rather than a struct field.

**Nothing filters.** Section 5 gives the runner two jobs and this implements one. A
digest written by this runner is a verbatim copy of what the agent supplied, so a
summary carrying a secret carries it into the tree. That was already true of every hand
written digest, and it is now true of a file the runner's name is on, which is worse in
exactly one way: it looks produced.

**AC5 is argued, not tested.** No test compares a written digest with a typed one of the
same content; the evidence is this intent's own phase going red and then green. A test
would need to fabricate the two writerless fields, which is the thing the writer exists
not to do.

**The before and after could not be measured.** `Finish` changed shape, so the new tests
do not compile against the old tree. Four intents in a row have reported which tests
fail without the change and this one cannot, which is a limit of the method rather than
of the change.

**`tool_version` is still absent from the config that carries `tool`.** Section 12's
block names the tool and the model and not the version, which is defensible — a version
is a property of a session — and it means the field has no home short of the harness.
Whoever builds WP11 decides where it comes from.

**The evidence is a local run.** Sixth intent in a row, same reason.
