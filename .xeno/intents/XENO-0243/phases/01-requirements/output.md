---
intent: github.com/triplem/xeno#208
phase: 01-requirements
created: "2026-10-04T07:37:51Z"
schema_version: "1.0"
runner_version: dev+49f2794
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f88a71e86c72c348270949328dbf227bc0e9cb85426abffbcce273ba402c1894
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      resolves: Q-1
      chosen: The writer, section 4's two unwritten provenance fields, this intent's P4 declaring four pending items, and a report and manifest published by xeno.yml on the pattern the three scan workflows already use.
      rationale: 'The issue''s done-when names a verification phase carrying its test report, and the three workflows that publish today publish scans, so the words are satisfied only where a test report exists to declare. The command alone would ship unexercised against a pipeline, which is the state every other half-built part of this mechanism is already in: the handling, the gate and the attach were built in WP6 and have never had an input. Against the smaller option it costs a restructured test step in a required workflow, which proves itself only on the push, and an open pull request at P4 so that pull_request triggers the scans; both are the ordinary cost of the arrangement section 6 describes, paid once here rather than deferred again.'
      decided_by: triplem
      proposed_by: claude-opus-5
    - id: D-2
      resolves: Q-2
      chosen: format and produced_by. Both are added to the model and written from a flag each.
      rationale: WP6 names both as what the package builds and neither has ever existed in the code, so this is a specified field gaining a writer and not an invented one. produced_by is the command as run, which is what a reader of a hash a year later wants and the one thing the hash does not say. id is left out because the attach and G-Evidence key an item by kind and job, and a second identity both readers ignore is how two keys for one thing come to disagree.
      decided_by: triplem
      proposed_by: claude-opus-5
    - id: D-3
      resolves: Q-3
      chosen: Declaration stays explicit. --evidence-from keeps attaching against declarations only, and a manifest entry nothing declared is ignored as it is today.
      rationale: 'Section 4 says evidence is declared, never inferred, and the first standing rule makes a contradiction of it a document change in its own commit before any code. Inferring would also answer in code a question the documents have already answered, and would put declarations in a trail that no person made. The cost is accepted: a report published and never declared stays unused, which is a phase that did not ask for it rather than a defect.'
      decided_by: triplem
      proposed_by: claude-opus-5
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**An item that exists is declared from its file, and the runner computes the hash.**
`xeno evidence declare --intent K --phase NN --kind test-report --job unit --result pass
--file PATH` copies the file into the phase's `evidence/` under its own name, writes
`path: evidence/<name>` and the `sha256` of what was copied into the `evidence` block of
that phase's `output.md`, and leaves the body and every other frontmatter field as it
found them. No flag takes a hash for a file, because the one field that binds the item
is the one nobody may type.

**An item in an artifact store is declared with its hash, and without one it is
refused.** `--uri URL --sha256 HEX` writes both as given. `--uri` without `--sha256` is
refused in the words `evidence.unbindable` already uses for the same defect on the
attach side: a uri is bound by its hash alone, because resolving one needs a network the
gate path never has.

**An item a pipeline owes declares its kind and its job and nothing else.** No `--file`,
no `--uri`, no `--result`: the run that would report one has not happened, and the value
arrives in `evidence/attached.yaml` as the job's own verdict. `--result` on a pending
declaration is refused for that reason. What is written is the pair `evidence attach`
looks an entry up by, so the declaration is what makes the pull path able to find
anything at all.

**It refuses before writing whatever G-Schema would report afterwards.** A kind outside
section 4's six; a result outside `pass` and `fail`; a `test-report` or a `build-log`
that exists and carries no result; a format outside section 4's five. The judgement is
the gate's own check called by the command, not a second opinion about shape, which is
how `question record` and `decision record` refuse. A refusal names which of these it
is, writes nothing, and exits 1.

**A pair is declared once.** A second `--kind K --job J` for the same phase is refused,
naming the declaration that is already there. `evidence attach` and `Collect` key an
attachment to a declaration by that pair, so two of them make an attachment ambiguous
and one of the two would silently never bind.

**The three forms are exclusive.** `--file` with `--uri`, and `--sha256` with `--file`,
are each refused with the reason rather than with a usage line: an item lives in the
repository or in a store according to its size, and a hash is computed where the bytes
are there to compute it from.

**`produced_by` and `format` are written where given and absent where not.** Both are
section 4 fields that the model has never carried. Nothing reads them, which is what
section 4 says of `format` in as many words, and nothing defaults them: an empty
`produced_by` is a field nobody filled, which A35 holds to be better than a plausible
value nobody produced.

**A declaration into a phase with no artifact is refused.** The sentence is `question
record`'s: the artifact is created by its first section write. A command that created it
here would write a frontmatter of whatever the runner could work out alone.

**A write after a verdict invalidates the verdict, as a section write does.** The
command does not look at `gate.yaml`, does not refuse a judged phase and rewrites
nothing. The next `phase finish` reports that the phase changed after its verdict was
written, which is #216's mechanism and the answer `section set`, `question record` and
`learning record` already give.

**Nothing in `evidence/` is left undeclared by the act that put it there.** The copy and
the declaration are one command, so G-Schema's clause about an undeclared file in
`evidence/` cannot be broken by using this command, which is the reason the copy is not
a separate step.

**`xeno.yml` publishes the test report it already produces.** The `verify` job writes
`go test -json ./...` to a file, prints a readable summary of it, writes a
`manifest.yaml` naming `kind: test-report`, `job: test`, `result` from the step's own
outcome, `format: go-test-json`, the run id and the commit, uploads both, and raises a
test failure again at the end so that a red suite still fails the job. The pattern is
`semgrep.yml`'s, which exists for the same reason: the run worth reading is the one that
failed.

**This intent's own verification phase carries evidence, through the path section 6
designed.** P4 declares four pending items — the test report and the three scans the
workflows publish — and finishes `provisional`, which `gate verify` reports and exits 0
on, so the push is not blocked. The branch is pushed, the pull request opened, the
pipeline publishes, and P5's `phase start --evidence-from DIR` attaches what arrived
into `evidence/attached.yaml` outside `artifacts_hash` and carries P4's verdict forward
to green without its `artifacts_hash` changing.

**G-Evidence can fail for a real reason.** With an item declared and resolved, editing
the copied file makes the gate report that the content does not match its hash, and
removing it makes the gate report that it does not resolve. Both are asserted by a test
rather than by the trail, because a trail that demonstrated them would be a repository
with a red verdict in it.

**Nothing else changes.** No gate is added or removed, no rule, no template, no
document, and no existing artifact. `format` joins the closed sets G-Schema already
judges, which no declaration in the trail carries, so no committed verdict changes:
`./xeno gate verify` stays at exit 0 over every verdict in the repository.

**The usual gates of this repository.** `gofmt` and `go vet` silent, the suite green,
the build clean, SPDX on any new file, prose at 88 columns, and a test for each refusal
and each of the three forms.

<!-- xeno:section:non-goals -->
## Non goals

**Judging an attachment's `result` against section 4's closed set.** The fifteen
attachments in the trail read `result: success`, a word the section does not define. A
check for it would turn fifteen sealed phases red, and `Verify` compares a recomputed
status against the committed one, so every one of them would be reported as a divergence
and CI would be red on history rather than on this change. The finding is recorded in
this intent and filed as an issue of its own; what to do about fifteen verdicts is a
decision, not a line in this commit.

**Inferring a declaration from a manifest.** D-3. `phase start --evidence-from DIR`
keeps attaching against declarations only, and an entry nothing declared stays unused.
Section 4 says evidence is declared and never inferred, and `CLAUSE-READERS.md` lists
that clause among the properties with nothing asserting it, which is where it stays: a
reader for it is a separate argument from a writer for the declaration.

**A gate, a rule or a required field for `produced_by` or `format`.** Both are
provenance and both are optional in the section that defines them. A gate invented for
them would be the invented rule the standing rules forbid, and a required field would
make every declaration already in the trail incomplete.

**`id` as a seventh field.** Q-2's third option. The attach and G-Evidence key an item
by kind and job, so an id would be a second identity for the same thing that both
readers would keep ignoring.

**Parsing a report.** Section 4 settles it: nothing parses one, `format` exists so that
a later version could without a schema change, and the result a gate reads is the one
the run reported about itself.

**Running anything.** The command declares and copies. It does not run a test, a build
or a scanner, and the one thing it computes is a hash of bytes it was handed.

**Rewriting the fifteen phases that declared by hand.** Their declarations sit inside
`artifacts_hash` with verdicts over them, and a rewrite changes every verdict in its
intent. They are described in P0's problem section and left alone.

**Attaching from a real artifact store.** The fetch adapter stays what it is, a
directory with a `manifest.yaml` standing in for the store, which is how the pipeline's
artifacts are downloaded in this intent too. Nothing here reaches the network.

**A second evidence path for the MCP surface.** The surface is a budget of six
operations and this is a CLI subcommand. A seventh operation needs the argument section
7 asks for, and declaring evidence has not been shown to need one.

**Attaching more than the declared four in this intent.** `other/trivy-db` is published
by the vulnerability scan and is left undeclared, which is the state an entry nothing
declares is designed to have: ignored, not an error.

<!-- xeno:section:constraints -->
## Constraints

**The documents are not editable here.** Section 4 fixes the nine fields, the two closed
sets, which kinds carry a result and where an item lives. Everything this intent writes
is one of those fields, and the two the model gains are enumerated there. No section
changes, so nothing waits on a commit this intent may not make.

**What is sealed is never rewritten.** The command writes into a phase's `output.md`,
which is inside `artifacts_hash`, so it is a tool for a phase somebody is working on. A
phase already judged is not refused and not repaired either: the write makes the verdict
stale and the next `phase finish` says so, which is the only mechanism this repository
has for that and the one #216 built.

**Only the run of a person writes inside `artifacts_hash`.** The declaration is written
by a command somebody runs; the attachment is written by a later command and lands in
`evidence/attached.yaml`, outside the hash. The new command does not touch
`attached.yaml`, and the attach does not touch the declaration. Neither widens the
narrow write path section 6 fixes for a phase start, which is two files in the immediate
predecessor and nothing else.

**One authority on shape.** The refusals call the gate's own check. A command with its
own opinion about what a declaration is would let something be well formed on the way in
and malformed in the file, which is the argument `exchange.go` already records.

**Evidence comes from CI.** `evidence.source` is `ci` for this repository, so the items
this intent declares are items a pipeline produces and the local run is not the producer
of any of them. The `--file` form exists for the small textual result section 4 has
copied into `evidence/`, and what it copies is a file pulled from the artifact store,
not one a developer typed. The command cannot tell the two apart; the regime stands in
`context.lock.yaml`, and A-001 records that this is as far as the tool can go.

**A phase that declares evidence cannot be finished without a CI round trip.** Section 4
names the cost and this intent pays it: P4 is provisional on the push, the pipeline runs
once, and P5's start attaches. Nobody waits at P4, which is the whole point of the
arrangement, and the branch has to carry an open pull request for the scan workflows to
trigger at all.

**The exit code staircase.** 0 where the declaration was written, 1 on a refusal with a
reason, 2 where the command could not run. The same staircase every other subcommand
follows, asserted through `run` rather than by executing the binary.

**One dependency.** Nothing new is imported. The hash comes from `internal/hashing`, the
frontmatter from `internal/fm` and `internal/runner`, the shape from `internal/gates`.

**Prose at 88 columns, Go at whatever `gofmt` produces, SPDX on new files.** The width
rule has no checker and is a rule whose reader is a person, which #211 wrote down
knowingly.
