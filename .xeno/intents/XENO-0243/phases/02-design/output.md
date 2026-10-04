---
intent: github.com/triplem/xeno#208
phase: 02-design
created: "2026-10-04T08:28:10Z"
schema_version: "1.0"
runner_version: dev+49f2794
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cdf4164ab7f4a154712181361be4dfe7ddb99d8879ba480dcd046e2003257874
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The command is `xeno evidence declare`, in section 4's own verb.** The surface spells
a writer `record` four times — learning, question, decision, assumption — and this one
does not join them, because section 4 names three acts and keeps them apart on purpose:
running, which Xeno never does, declaring, which is the act that binds, and judging,
which is a gate's. `evidence attach` is already the second verb in this family and binds
what a pipeline published to a declaration. A fifth `record` would make the one act the
document calls the only one Xeno knows indistinguishable from the four that write prose
into frontmatter.

**Three forms, one command.** The declaration has three states in section 4 — an item
that exists in the repository, an item in a store, an item a pipeline owes — and they
differ by which two fields are filled. Three subcommands would make a reader choose a
verb for a distinction the section draws by size and timing, and the refusals that keep
them apart are worth more than the names would be: `--file` with `--uri` is refused,
`--sha256` with `--file` is refused, and `--result` on a pending item is refused.

**The hash is computed, never given, and there is no flag for it on the `--file` form.**
This is the property the issue asks for in the words of the `CLAUSE-READERS.md` row it
cites: a command that changes a hashed file recomputes what it invalidates, which until
now `SectionSet` did for `context_hash` and nothing did in general. The declared item is
bound by its `sha256` and nothing else, so a typed hash would be a binding somebody
chose rather than one the bytes produced. `--sha256` exists only with `--uri`, where the
bytes are not here and the pipeline published the hash; that asymmetry is the whole of
what `uri` costs and `evidence.unbindable` already says it in the words the refusal
reuses.

**The copy and the declaration are one act.** `--file` copies the report into the
phase's `evidence/` under its own name and declares `path: evidence/<name>` in the same
run. Two commands, or a flag that declared a path somebody had already copied, would
make G-Schema's clause about an undeclared file in `evidence/` reachable by using the
tool correctly. Where the target exists already with different bytes the command refuses
rather than overwriting, because a file in `evidence/` is bound by a hash in some
declaration and the overwrite would break a verdict that is not this item's.

**The pair is the identity, and it is declared once.** `Collect` and `Attach` both key
an attachment to a declaration by kind and job. A second declaration of the same pair is
refused naming the first, because the attach would bind one of the two and leave the
other pending for ever with nothing saying why. That is also the answer to Q-2's `id`:
the identity exists, it is two fields, and a third would be a name for it that no reader
uses.

**Refusals call the gate's own check.** `evidenceShape` becomes `EvidenceShape`, as
`QuestionShape` and `DecisionShape` were exported for the same reason, and the command
runs it over the item it is about to write. Nothing may be well formed on the way in and
malformed in the file, and the command holds no second opinion about what a declaration
is.

**`format` joins the closed sets the gate judges.** Section 4 enumerates five values for
it, and a set the document closes with nothing comparing against it is how `kind: build`
survived in the gate for as long as it did, which #198's `BuildKind` comment records. No
declaration in the trail carries a `format`, so nothing already committed can become a
finding: the check is safe to add in the same commit as the field, which is the only
moment it is safe.

**The writer amends one frontmatter field and renders nothing.** `exchange.go` already
answers this for the two blocks beside `evidence`: going through `SectionSet` would
re-render the body from the template, which is the template's business and not a
declaration's. The same two helpers are reused — `r.artifact` to read the three views of
`output.md`, `r.amendFront` to write one field back with the body byte for byte — and
`frontmatterOrder` already places `evidence` between `decisions` and `review_checklist`.

**It says nothing about `gate.yaml`.** A declaration after a verdict changes the
artifact and so changes its hash, and the next `phase finish` reports that the phase
changed after its verdict was written. That is #216's state, and it is what `section
set`, `question record` and `learning record` already do. A command that refused a
judged phase would be inventing a rule; one that rewrote the verdict would be writing a
verdict nobody asked for.

**The declaration's fields go in section 4's order.** `EvidenceItem` gains `produced_by`
and `format` and is reordered to the order the section prints: kind, result,
produced_by, format, sha256, path, uri, job. Field order is a decision this repository
already takes seriously enough to keep a list for in `frontmatterOrder`, for the reason
that a reader compares an artifact against the document and not against a struct.

**The test report is published by the job that already runs the tests.** `xeno.yml`'s
`verify` job gains `go test -json` into a file, a readable summary printed from it, a
`manifest.yaml` in the shape the three scan workflows write, an upload, and the test
failure raised again at the end. The pattern is copied rather than invented, and the
reason it has that shape is the one `semgrep.yml` records: the run worth reading is the
one that failed, so the step that fails cannot be the step that stops the report from
being published.

<!-- xeno:section:alternatives -->
## Alternatives

**`xeno evidence record`, for symmetry with the other four writers.** Rejected on the
previous section's first paragraph: the symmetry is with `attach`, not with `record`,
because both belong to the one mechanism section 4 defines and the section's whole
method is keeping declaring apart from running and from judging.

**A flag on `section set`, so that a phase writes its evidence as it writes its
sections.** It is the shortest change and it was rejected for two reasons. `section set`
re-renders the body from the template on every write, so a declaration would be coupled
to a template resolution that has nothing to do with it, and the flag would be ignored
on every other section, which is a flag that is wrong six times out of seven.
`exchange.go` took the same decision two weeks ago for `open_questions` and `decisions`.

**`--path` instead of `--file`, naming the location inside `evidence/` directly.** It
reads closer to the field being written, and it was rejected because it describes the
result rather than the input: somebody would have to copy the file first, and the copy
and the declaration being one act is what keeps an undeclared file out of `evidence/`.
`--from` was rejected too — it is the attach's flag for the directory standing in for
the artifact store, and one word for a file here and a directory there is how a reader
learns to read a flag twice.

**`--sha256` accepted on the `--file` form, for the case where somebody already has the
hash.** Rejected: the file is there, so the hash is a fact and not an input, and a given
one can only agree or disagree with the bytes. Agreeing, it is noise; disagreeing, the
command would have to decide which of the two to write, and whichever it chose would be
a declaration somebody can make say what they want. That is the one thing the field
exists to prevent.

**Inferring the declaration from the manifest, which is what the issue asks about.** Q-3
and D-3. Rejected because section 4 says evidence is declared and never inferred, and
the first standing rule makes a contradiction of that a document change in its own
commit before any code.

**Declaring a result on a pending item, from the local run.** Rejected: the pipeline's
run is the one being declared, and a local result would be a claim about a run that has
not happened. The value arrives with the attachment as the job's own verdict, which is
what `attached.yaml` carries and what G-Test and G-Build read.

**Judging the attachment's `result` against the closed set in the same commit, since the
set is being extended for `format` anyway.** Rejected, and this is the closest call. It
is four lines, it would catch the one defect this intent found in the existing trail,
and `Verify` would then report fifteen sealed phases as divergent, because it compares a
recomputed status against the committed one. CI would be red on history rather than on
the change. The finding is recorded and filed; the fix is a decision about fifteen
verdicts.

**Publishing the test report as JUnit XML through a converter, rather than as the JSON
`go test` emits.** Rejected: it adds a tool to a workflow in order to produce a format
nothing parses. Section 4 is explicit that nothing parses the report and that `format`
exists so that a later version could, and `go-test-json` is one of the five values the
section names.

**Declaring the test report with `--file` from a local `go test` run instead of pending
from CI.** It would have finished the intent in one push, with no round trip and no open
pull request at P4. Rejected because `evidence.source` for this repository reads `ci`,
and a locally produced report is the exception the field exists to mark, not the
default. The fifteen declarations already in the trail took exactly that shortcut, with
`pipeline: local` and a hand written transcript behind the hash, and reproducing it in
the intent filed to end it would have been the clearest possible way to leave the gap
open.

<!-- xeno:section:impact -->
## Impact

**One new file and four touched.** `internal/runner/evidence.go` holds the writer,
beside `exchange.go`, which it borrows `artifact` and `amendFront` from.
`internal/model/model.go` gains two fields on `EvidenceItem`, the field order of section
4, and the closed set for `format`. `internal/gates/gates.go` exports `EvidenceShape`
and judges `format` against that set. `cmd/xeno/main.go` gains the subcommand, seven
flags and a usage line. `.github/workflows/xeno.yml` gains the report, the manifest and
the upload.

**No gate is added or removed, and no verdict in the trail changes.** The `format` check
is the only new finding anything can produce, and nothing in the repository declares a
`format`, so `gate verify` stays at exit 0 over every verdict. The field order changes
what a newly written item looks like and no file that already exists, because the block
is re-marshalled only in the phase a declaration is being written into.

**`model.EvidenceItem` is read in four places** — `evidenceShape`, `Collect`,
`evidence.Attach` and `Pending` — and every one of them reads kind, job, result, sha256,
path or uri. The two new fields are written and read by nothing, which is the state
section 4 describes for `format` and A-002 records for `produced_by`.

**The exit code staircase gains one command and no step.** 0 where the item was
declared, 1 on each of the eight refusals, 2 where the flags could not be parsed or the
phase does not resolve. Asserted through `run` in `cmd/xeno`, which is how #110 made the
staircase testable without executing the binary.

**The write path does not widen.** The command writes `output.md` of the phase named
and, on the `--file` form, one file under that phase's `evidence/`. It never writes
`attached.yaml`, never writes `gate.yaml`, and never touches another phase, so section
6's sentence about a phase start writing two files in its immediate predecessor stands
untouched.

**A phase that declares evidence is provisional until somebody carries on.** This
intent's P4 will be, which is what section 6 designed and what `verifyCode` exits 0 on.
The cost is the one section 4 names: a CI round trip before P5 can start, and an open
pull request for the scan workflows to trigger at all. If nobody carries on, the verdict
stays provisional, which is honest, because then nobody judged it.

**`xeno.yml`'s `verify` job changes shape, and it is a required check.** The test step
becomes `continue-on-error` with its outcome read by the manifest step and raised again
at the end, which means a failing suite fails the job one step later than it does today.
Everything between those two steps — the build, the title check, `gate verify`, the
trail guard — runs on a tree whose tests failed, which is a readable report rather than
a truncated one and is the arrangement `semgrep.yml` already uses for the same reason.

**What a reader of the trail gains.** A verification phase whose test report is a file a
pipeline wrote, bound by a hash the runner computed, declared by a command, with the run
id and the commit beside it in `attached.yaml`. Before this, the strongest thing P4
could say was a sentence in prose and a transcript somebody pasted.

**What nobody gains yet.** `produced_by` and `format` have no reader, the attachment's
`result` is still unjudged, and nothing makes a verification phase declare anything at
all. The last of those is the same shape as #188's residual risk: the command exists and
the habit is what has to change.
