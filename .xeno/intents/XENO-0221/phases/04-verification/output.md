---
intent: github.com/triplem/xeno#167
phase: 04-verification
created: "2026-10-01T16:57:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f8aae81434f84d87589f815a574a6a3c07d808a3c752fe10d40e15e727079997
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

| Criterion | Test |
|---|---|
| A project that declares nothing is unaffected | `TestNothingDeclaredRunsNothing`, `TestNoDeclarationLeavesTheVerdictAsItWas` |
| A matching gate runs, the exit status is the verdict, the check carries id and provenance | `TestADeclaredGateRunsAndPasses`, `TestANonZeroExitIsAFail` |
| Exit zero with findings is a pass carrying them | `TestExitZeroWithFindingsIsAPassCarryingThem` |
| Findings get their ids the same way as every other | `TestAnExternalGateReachesTheVerdictAndKeepsNoDecision` |
| A mismatch does not run and is red, naming the mismatch | `TestAModifiedGateRefusesToRun`, which also asserts the command left no trace |
| A declaration that cannot be run is red, naming what happened | `TestAMissingOrUnrunnableGateIsAFinding` |
| An answer that is not the agreed JSON is red | `TestAnAnswerThatIsNotTheAgreedJSONIsAFail`, three shapes |
| A non-zero exit with no findings still produces one | `TestANonZeroExitWithNothingToSayStillProducesAFinding` |
| A command that does not finish is red, naming the limit | `TestAGateThatDoesNotFinishIsKilled` |
| A decision on an external finding does not survive | `TestAnExternalGateReachesTheVerdictAndKeepsNoDecision`, the second half |
| A gate runs only at the phases it names, and a declaration for none runs nowhere | `TestAGateRunsOnlyAtThePhasesItNames`, `TestADeclarationForNoPhaseRunsNowhereAndSaysSo` |
| A wrong declaration reads as a configuration mistake | `TestADeclarationThatIsWrongReadsAsAConfigurationMistake`, six cases, each naming `project.yaml` |
| `G-Complete` accepts an external check like any other | not tested here, and not new: it reads the checks it is given, and the phase going red in the runner test is the same path |
| Nothing else moves | the suite, `gofmt`, `go vet`, `./xeno gate verify` |

Two things are asserted that no criterion asked for. The contract's input is read back out of the
process in `TestTheCommandIsGivenTheIntentThePhaseAndWhereToLook`, because A76 is a promise to
somebody else's tool and a promise nothing checks is a comment. And a two thousand character cause
is bounded in `TestAForeignAnswerIsBounded`, because what a foreign command prints is not this
project's to control.

<!-- xeno:section:results -->
## Results

**The suite is green.** 407 cases pass across the tree, nothing fails, `gofmt -l` outside
`vendor/` lists nothing, `go vet ./...` is silent, `./xeno gate verify` is at exit 0 over 204
verdicts.

**`internal/external` is 1.2 seconds, of which one second is the hung gate.** Before `WaitDelay`
the same test took thirty: the process was killed on time and the read waited for the pipe a
grandchild still held. That is the measurement behind A75 and the reason the row says what it
says.

**A modified gate did not run.** The test writes a command, declares its hash, replaces the file
with one that would touch a marker, and judges: the check is red naming both hashes, and the
marker does not exist. That is WP4's done-when — "a modified external gate refuses to run rather
than running unnoticed" — with the second half asserted rather than assumed.

**An external check reaches the verdict and keeps no decision.** Through the runner, with a
command declared in `project.yaml`: the check carries `house-linter` and `provenance: external`,
its finding carries an id, the phase is red, `gate approve` records the release, and the next
`gate run` produces the same finding with the same id and no decision. #66 wrote that rule and
the comment at the verdict said it was checked there so that "a second path into it, an external
gate above all, meets the same rule as the first". The second path exists and meets it.

**Four ways a foreign command can behave badly are each a fail with a cause.** Not valid JSON, a
list where an object belongs, a finding with no cause, and a non-zero exit with nothing to say.
The last one is the one that matters structurally: `Status` refuses a failing check with no
finding, so without the synthesised one a project's tool could make a verdict unrepresentable and
refuse the run.

**Six wrong declarations each name `.xeno/config/project.yaml`.** No id, no path, no hash, a hash
that is not sixty-four lowercase hex characters, an upper-case hash, and a phase that is not one
of the six. A project's configuration mistake reads as a configuration mistake rather than as a
gate that would not start.

**The contract is checked against itself from the outside.** The fake gate writes its stdin to a
file and the test unmarshals it into `Input`, asserting all five fields. A76 is a promise to
somebody else's tool, and this is the only thing in the repository that holds it.

**A project that declares none produces no external check**, and the closure is nil rather than
empty, so the common case adds nothing to the verdict path.

<!-- xeno:section:gaps -->
## Gaps

**There is no sandbox and this piece did not build one.** The command runs with the pipeline's
environment, the pipeline's credentials and write access to the repository. Section 14 says the
marking is the point, so what the trail gains is honesty rather than containment, and a project
enabling an external gate has decided to run its own code in its own pipeline. Nothing here
reduces that and nothing pretends to.

**`gate verify` runs the commands too.** It recomputes verdicts, so a project with an external
gate executes its own tool once per phase in the verification step as well as in the run. That
follows from an external gate being a gate, it is written in the design's impact, and nobody has
measured what it costs on a trail the size of this one.

**No external gate has ever run outside a test.** Every other piece of WP4 was exercised against
this repository on a throwaway copy; this one was not, because declaring a gate here means
writing a tool for this project to run on itself, which is content. So what is proven is the
mechanism against fake commands written by the person who wrote the mechanism, and the first real
gate is the first independent reading of the contract.

**A76 is a promise with one holder.** The input shape is asserted by one test reading the process's
stdin back. Nothing generates documentation from it, nothing versions it, and a project's tool
that agrees with it today has no way to notice if it changes — the honest version of that is the
generated reference WP16 specifies.

**The timeout is one number for every gate.** Sixty seconds, not configurable, with no per-gate
override. A project whose linter legitimately takes longer has to argue with A75 rather than with
its configuration. That is deliberate for a first version and it is the first thing a real adopter
will ask for.

**Findings are not deduplicated and not bounded in number.** A tool that answers ten thousand
findings puts ten thousand findings in `gate.yaml`, each with an id. Each one is clipped in length
and nothing clips the list.

**A gate's own output on standard error is discarded.** `Output` captures stdout and the error
text only on a failure to start. A tool that explains itself on stderr says it into nothing, which
will look like a bug to whoever wrote the tool.

**The declaration's `phases` is validated against the six and nothing else.** A gate declared for
every phase runs six times per intent, which is a project's choice and reads as one; but nothing
warns that an expensive gate at P0 is paid for at every phase.
