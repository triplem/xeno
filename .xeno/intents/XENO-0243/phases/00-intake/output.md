---
intent: github.com/triplem/xeno#208
phase: 00-intake
created: "2026-10-04T07:35:31Z"
schema_version: "1.0"
runner_version: dev+49f2794
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c84fb91310fc1166c95dfb7283857d31876d2b7e1496a127edb430a0f04d1cb3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
open_questions:
    - key: Q-1
      text: '#208 asks for a declaring command and says the loop has never run end to end. How far does this intent go: the writer alone, the writer with real evidence in its own trail, or that plus a publishing step so that a test report exists to declare?'
      options:
        - text: The writer, section 4's two unwritten provenance fields, this intent's P4 declaring four pending items, and a report and manifest published by xeno.yml on the pattern the three scan workflows already use.
          consequence: the issue's own words are satisfied, a verification phase carries its test report rather than naming it in prose, G-Evidence judges four real items and the attach resolves something; the cost is that a required workflow's test step is restructured to continue-on-error with the failure raised again at the end, it proves itself only on the push, and the pull request has to be open at P4 for pull_request to trigger the scans
          recommended: true
        - text: The writer, and P4 declares only the three scans the workflows already publish.
          consequence: the pull path runs for real with no change to any workflow, and the trail carries its first attached evidence; the cost is that every item is kind scan, so the test report stays named in prose, which is the half of the issue's done-when that mentions one
        - text: The writer and its tests, with nothing declared in this intent's own trail.
          consequence: the cheapest and the lowest risk, no CI round trip, one push; the cost is that "no verification phase has ever attached evidence" survives the commit filed to end it, and the command ships without ever having been run against a pipeline
        - text: Something else, entered by the person deciding
          free: true
    - key: Q-2
      text: Section 4's evidence block lists id, produced_by and format; model.EvidenceItem carries none of the three and no gate reads them. Which of them does the writer write?
      options:
        - text: format and produced_by. Both are added to the model and written from a flag each.
          consequence: a declaration says which command produced the item and in what shape, which is what makes a hash readable a year later, and both are enumerated in section 4, so neither is an invented field; the cost is two fields round-tripping through every artifact's frontmatter with nothing else writing them, and a CI manifest carries neither, so an attached item leaves them as the declaration set them
          recommended: true
        - text: 'None. The writer writes what the model already has: kind, job, result, path, uri and sha256.'
          consequence: 'the smallest change, and the command writes exactly what the gates read; the cost is that three specified fields keep no writer, so the gap #208 names shrinks rather than closes and produced_by stays in prose'
        - text: All three, with id assigned the way Q-n and D-n are.
          consequence: the block as section 4 prints it, field for field; the cost is a second identity beside kind and job for the same item, which the attach and G-Evidence would both keep ignoring, and two keys for one thing is how they come to disagree
        - text: Something else, entered by the person deciding
          free: true
    - key: Q-3
      text: '#208 asks it alongside: should phase start --evidence-from DIR declare what it finds in the manifest, or does declaration stay explicit?'
      options:
        - text: Declaration stays explicit. --evidence-from keeps attaching against declarations only, and a manifest entry nothing declared is ignored as it is today.
          consequence: section 4's clause that evidence is declared and never inferred holds, and the new command is the only way an item enters a trail; the cost is that somebody has to declare before anything a pipeline publishes is of use, so a report published and never declared is silently unused
          recommended: true
        - text: Infer it. A manifest entry with no declaration gets one written for it.
          consequence: nothing a pipeline publishes is lost and P4 needs no foresight about which jobs will run; the cost is that it contradicts a normative clause, so by the first standing rule it is a document change in its own commit first, and the trail would then carry declarations no person made
        - text: Something else, entered by the person deciding
          free: true
---

# Intake

<!-- xeno:section:problem -->
## Problem

Section 4 keeps three acts apart. **Running** a test, a build or a scanner is something
Xeno never does. **Declaring** that a result exists, with its kind, its location and its
hash, is the act that binds and the only one Xeno knows. **Judging** is a gate checking
that a declaration resolves and that its hash matches. The declaration is a block in the
frontmatter of the producing phase's `output.md`, which `xeno section set` writes and
`artifacts_hash` covers.

**Nothing writes the block.** `assumption record` writes the register, `learning record`
writes the learning file, `question record` and `decision record` write the two
frontmatter blocks beside this one since #188. For evidence there is `evidence attach`,
which binds what a pipeline published to a declaration that is already there, and no
command that puts one there. So the act section 4 calls the only one Xeno knows is the
one act the runner cannot perform.

**What happened instead is the measurement.** Fifteen of the forty-seven verification
phases in this trail carry a declaration, every one of them `test-report`/`go-test`,
every one of them attached, and every attachment reads `pipeline: local` with `result:
success`. The newest is XENO-0210 on 2026-09-29; nothing since declares anything, which
is thirty-two verification phases. Both halves of each of the fifteen were typed: the
declaration into the frontmatter of a file inside `artifacts_hash`, and the attachment
into `evidence/attached.yaml`, which lies outside it. The content they point at is a
hand curated transcript with commentary in it, not a report a run emitted. That is the
case section 4 gives as its reason for sourcing evidence from CI — "a JUnit report can
be written by hand" — arrived at from the inside, and `evidence.source` for this
repository reads `ci`.

**`result: success` is not a word section 4 defines.** The closed set is `pass` and
`fail`, G-Schema judges a declaration against it, and nothing judges an attachment,
because the shape check reads the frontmatter and `attached.yaml` is a different file.
Fifteen records carry a value no reader of this repository can act on, in the one file a
sealed phase can be edited in without making any verdict stale.

**And nothing has ever been attached from a pipeline.** Three workflows publish a
manifest in the shape `evidence attach` reads — `scan/semgrep`, `scan/trivy` with
`other/trivy-db` beside it, and `scan/npm-audit` — on every pull request, each as an
uploaded artifact. The pull path of section 6, where `phase start` of the next phase
attaches what has arrived for its predecessor, has never once run. It cannot: there is
nothing to attach against, because a declaration is what the attach looks for.

**The two gates over this are vacuous, and correctly so.** G-Evidence checks that every
declared item resolves and its hash matches; over an empty set it passes. G-Build reads
`result` on `build-log` items; there are none anywhere in the trail. A gate that cannot
fail because its input is never produced is the audit's own category in #202, arrived at
from the other direction: #202 looked for clauses with no reader, and this is a reader
with no input.

WP6 built the handling, the gate and the attach. What is missing is the writer, the same
gap `tool_version` had before #181 and `plugin_version` before #177. Both of those ended
the same way: a field specified, carried by every artifact, and filled by hand until a
command filled it.

<!-- xeno:section:scope -->
## Scope

**In scope.** One writing command, two fields, and this intent's own verification phase
carrying real evidence through the path section 6 designed for it.

`xeno evidence declare` writes one item into the `evidence` block of a phase's
`output.md`. Three forms, which are section 4's three states and not three commands: an
item that exists now is declared from the file, which is copied into `evidence/` and
whose `sha256` the command computes; an item that lives in an artifact store is declared
with its `uri` and the hash the pipeline published; an item a pipeline has yet to
produce declares only its kind and its job, which is what makes `evidence attach` able
to find it later. The command refuses before writing whatever G-Schema would report
afterwards, by calling the gate's own shape check rather than holding a second opinion
about it, which is how `question record` and `decision record` are built.

`produced_by` and `format` are added to the model. Section 4 enumerates both, WP6 names
them as what the package builds, and neither has ever been written: `produced_by` is the
command as run, which is the one thing a reader of a hash two years later wants and the
one thing the hash does not say, and `format` names the shape of the report without
implying that anything parses it.

The verification phase of this intent declares four pending items and attaches them from
the pipeline: the test report, and the three scans the workflows already publish. The
test report needs a publishing step, because `go test ./...` in `xeno.yml` writes no
file; it is added on the pattern the three scan workflows already use, a report, a
manifest, an upload, and the failure raised again at the end.

**Out of scope, each for its own reason.**

Judging an attachment's `result` against section 4's closed set. The fifteen attachments
in the trail carry `success`, so a check for it would turn fifteen sealed phases red,
and `gate verify` compares a recomputed status against the committed one: the divergence
would be reported for every one of them and CI would be red on history rather than on
this change. The finding is recorded in this phase and filed as an issue of its own; the
fix is somebody's deliberate decision about fifteen verdicts, not a line in this commit.

Inferring a declaration from a manifest. `phase start --evidence-from DIR` keeps
attaching against declarations only. Section 4 says evidence is declared, never
inferred, `CLAUSE-READERS.md` lists that clause with nothing asserting it, and inferring
here would answer in code a question the documents have already answered. Q-3 and D-3.

A reader for `produced_by` or `format`. Both are provenance. No gate reads them, which
is what section 4 says of `format` in as many words, and a gate invented for them would
be the invented rule the standing rules forbid.

The fifteen phases already in the trail. Their declarations are inside `artifacts_hash`
with verdicts written over them, so rewriting one changes every verdict in its intent.
They stay as they are and this phase's problem section is where they are described.

**What this intent owes its own trail.** Its P4 is the first verification phase in this
repository whose evidence is a file a pipeline wrote, bound by a hash the runner
computed, with the declaration made by a command. The acceptance criteria say so in the
form a test can read, and the pull path is exercised rather than described: P4 finishes
provisional, the branch is pushed, the pipeline publishes, and P5's start attaches.

<!-- xeno:section:context-rationale -->
## Why this context

#208 is read in full, with its claims checked against the tree rather than carried over.
Two of them hold: the declaration has no writer, and `evidence attach` has nothing to
attach against. One does not. "No verification phase in this trail has ever attached
evidence" was written from XENO-0237, where the attach was run before the sections
existed and failed on a missing `output.md`; fifteen earlier phases do carry an
attachment, by hand, and the problem section replaces the sentence with the count. The
issue's own framing of itself as a gate that cannot fail stands either way, because what
those fifteen attached is a transcript somebody typed.

Section 4's evidence subsection is read whole, as the specification of what may be
written: the three acts, the declaration's nine fields, the two closed sets, which items
carry a `result` and which do not, where an item lives according to its size rather than
its origin, and the sentence that `evidence/` sits outside `artifacts_hash` while the
declaration that binds it does not. It is also where the regime comes from: evidence is
produced in CI, `local` is the deliberate exception, and this repository is configured
`ci`.

Section 6's two tables and the note under them are read for the sequence this intent has
to walk rather than describe: `phase finish` judges what it can decide locally, the push
happens without waiting, the pipeline publishes, and the next `phase start` attaches.
The note also fixes what a phase start may write in its predecessor's directory, two
files and no more, which is why nothing in this intent widens that path.

Sections 5 and 9 are read for where the block sits in the frontmatter and for
G-Evidence, G-Build and G-Schema's clause about `evidence/`. Section 7 is read for the
exit code staircase the new command joins and for the MCP surface, which is a budget of
six operations and is not touched: a CLI subcommand is not an MCP tool.

`internal/evidence/attach.go` is read first among the code, because it is the reader of
what the new command writes. Its `unbindable` is the authority on why a `uri` without a
hash cannot be recorded, and the same reasoning decides what the declare refuses.

`internal/gates/gates.go` is read for `evidenceShape`, which is the check the command
calls so that it cannot write what the gate would report, for `evidence` and `build`,
which are the gates that then have an input, for `undeclaredEvidence`, which is why a
file copied into `evidence/` has to be declared in the same act, and for `Collect`,
which keys an attachment to a declaration by kind and job and is the reason a duplicate
pair is refused.

`internal/runner/exchange.go` is read as the pattern. It is two weeks old, it writes
frontmatter blocks of the same kind, and it already answers the three questions a writer
of a sealed artifact has: amend one field rather than re-render the body, refuse through
the gate's own check, and say nothing about `gate.yaml` because a write after a verdict
is #216's state and not this command's business.

`internal/runner/runner.go` is read for `SectionSet`, which is the one writer of
`output.md` and the one place a hashed field is recomputed on every write, and for
`frontmatterOrder`, which already places `evidence` between `decisions` and
`review_checklist`.

`internal/model/model.go` is read for `EvidenceItem`, `Attached` and `Output`, and for
the two closed sets. The item carries six of section 4's nine fields; `id`,
`produced_by` and `format` have never existed in the code, which is the half of the gap
that is not the missing command.

`internal/runner/runner.go:Verify` is read for one reason: it compares a recomputed
status against the committed one, which is what makes judging the fifteen old
attachments a change to history rather than a change to the tool, and is why that
finding leaves this intent as an issue.

`.github/workflows/semgrep.yml`, `trivy.yml` and `audit.yml` are read for the manifest
they publish and for the pattern the test report's publishing step copies: the report,
the manifest written from the step's own outcome, the upload, and the failure raised
again at the end so that a scan which fails is still readable. `xeno.yml` is read for
where that step goes and for what the `verify` job does with a provisional verdict.

`CLAUSE-READERS.md` is read for the two rows this intent moves and for the one it does
not: "evidence is declared, never inferred" stays a property with no reader,
deliberately, and Q-3 is why.

Nothing outside the repository is needed. The issue, the two documents, nine source
files and four workflows are the whole base.
