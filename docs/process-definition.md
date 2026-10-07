---
id: process-definition
title: Xeno, Process Definition v1
revision: 10
status: draft, not ratified
date: 2026-09-20
location: docs/process-definition.md
---

# Xeno, Process Definition v1

## 1. Purpose and scope

Requirements, design, code and evidence are produced separately, at different times,
partly by people and partly by models. The only thing holding them together is their
binding to an intent. Making that binding provable is the purpose of this process.

This document defines the process, not its implementation. It fixes which phases
exist, which artifacts are produced, which gates apply, and how learning happens.

Serving as an evidence source for ISO/IEC 42001 certification is the declared next
milestone. Xeno supplies evidence. It does not run a management system and does not
replace the preparation for a certification, and the mapping between process
artifacts and the requirements they speak to is itself a deliverable rather than an
assumption. It is not fully binding in v1. The places where v1 falls short are
listed in section 16.

Target environments are projects with their own infrastructure, mostly brownfield.
Process artifacts live in the project repository. Aggregation beyond a single
project is foreseen once the necessary infrastructure exists and is not part of v1.

One intent belongs to one repository. v1 targets monorepos, and the assurance that
follows is worth more than the generality it gives up: every hash in the trail
resolves against the repository it sits in, with no reference a local, network free
runner cannot check. Intents spanning several repositories return in 1.1.

## 2. Principles

1. Every artifact produced belongs identifiably to exactly one intent.
2. The repository is the source of truth. The tracker displays. The trail is therefore
   exactly as public as the repository: an open project's intents, assumptions, digests
   and named decision makers are open too, which is a property of keeping evidence where
   the code is and not a setting anybody can change.
3. The AI makes no silent assumptions. Defaults are allowed, every assumption is
   recorded and confirmed.
4. Gates are deterministic and model free. No gate calls a model.
5. At any point it is reconstructable which information a version was built on.
6. Learning is mandatory at the end of every phase and is itself versioned.
7. On the same level, a given rule beats the learned one. A rule marked `binding`
   under `given/builtin/`, `given/provider/` or `given/org/` cannot be overridden at
   all.

## 3. Intent identity

The canonical intent id is qualified: tracker host, project or repository, and
key. The short tracker key, for example `PROJ-123`, is its display form and is what
directory names use, which keeps `.xeno/intents/PROJ-123/` readable. A project has
exactly one tracker, so the short form is unambiguous inside a repository.

The ID is flat. Governance attributes are not inherited across epic levels.
Brownfield changes also require an issue, if necessary one shared issue covering
several related changes.

Only the qualified form is unique beyond one repository, and trackers differ here in
ways that matter for the adapter. A key can be tied to the repository or issued
independently of it, and some number ranges are shared with other object kinds. Jira
is the example of the decoupled case, with project wide keys and no dependency on
the code host, and it is what 1.1 has to accommodate. Which variant a system runs is
the adapter's business; this document assumes none of them. Inside a single
repository the qualified form buys nothing and is kept all the same: retrofitting it
would touch every artifact path and every stored reference.

## 4. Directory layout

All paths in this document are relative to the repository root.

```
.xeno/
  docs/
    process-definition.md
    implementation-plan.md
  config/
    project.yaml
    secrets.yaml            project additions to the secret filter, optional
    templates/              project overrides only, optional
      <n>/
        template.yaml       structure: section ids, order, required fields
        strings.en.yaml
        strings.de.yaml
    rules/
      given/provider/
      given/org/
      given/project/
      learned/provider/
      learned/org/
      learned/project/
  plugin/                 vendored, pinned
    plugin.json
    skills/
    mcp.json
    templates/            full template set, shipped
    rules/given/builtin/  shipped rule set
    secrets.yaml          shipped secret filter
  intents/
    PROJ-123/
      intent.yaml
      assumptions.yaml
      learning.yaml       closing record, written when the intent ends
      gate.yaml           G-Complete result, written when an intent is abandoned
      phases/
        00-intake/
          context-scope.yaml
        01-requirements/
        02-design/
        03-implementation/
        04-verification/
          evidence/
        05-review/
  local/                  excluded via .gitignore
```

Every phase directory contains:

| File                 | Content                                                       |
|----------------------|---------------------------------------------------------------|
| `output.md`          | the phase result                                              |
| `context.lock.yaml`  | the information base used                                     |
| `digest.md`          | shortened, secret filtered summary of the exchange            |
| `gate.yaml`          | result of all gates for this phase, and the only place a verdict lives |
| `learning.yaml`      | learning record for this phase                                |
| `cost.yaml`          | cost record for this phase                                    |

P0 additionally holds `context-scope.yaml`, described under context economy. A
phase may also carry an `evidence/` directory. That is where produced evidence
lands, most often in verification.

The intent level carries a `gate.yaml` of its own in one case: an intent that is
abandoned. G-Complete then has no phase to sit in, because an intent can be dropped in
P1 and would otherwise never meet the gate that checks how it closed. An intent that
merges meets the same gate as part of P5, and its result stays in that phase.

`.xeno/` and code go into the same merge request.

### Evidence

Builds and tests produce artifacts, and those artifacts are what a gate reads.
Three different acts hide under the same everyday words, and keeping them apart is
what makes the rest of this section readable. **Running** a test, a build or a scanner
is something Xeno never does. **Declaring** that a result exists, with its kind, its
location and its hash, is the act that binds, and it is the only one Xeno knows.
**Judging** is a gate checking that a declaration resolves and that its hash matches.
A test run locally during the work is feedback for the developer and is none of these
three.

Evidence is declared, never inferred. The declaration is an optional block in the
frontmatter of the producing phase's `output.md`, alongside the mandatory fields
defined in section 5:

```yaml
evidence:
  - id: unit-tests
    kind: <test-report|coverage|build-log|scan|sbom|other>
    result: <pass|fail>              # required for test-report and build-log
    produced_by: <command as run>
    format: <junit|trx|tap|go-test-json|other>   # optional
    sha256: <hash of the content>
    path: evidence/junit.xml        # in-repo, for small text results
    uri: <url>                      # external, for large or binary results
    job: <name>                     # which job is expected to produce it
```

An item that exists at finish time carries its `sha256` here and is sealed with the
artifact. An item that cannot exist yet, because a pipeline has to produce it, declares
only its kind and its job, and what arrives later is written to
`evidence/attached.yaml`, which sits outside the `artifacts_hash`:

```yaml
- kind: test
  job: unit
  state: attached
  result: pass                      # the job's own verdict, not Xeno's
  sha256: <hash of the content>
  uri: <url>
  pipeline: <id>
  commit: <sha>                     # what it was produced against
```

The cost is stated rather than hidden: an item attached afterwards is not bound by the
sealed artifact but by the attachment record and by the verdict that read it. `gate.yaml`
records the hashes it judged, so altering an attachment after the fact makes it disagree
with the verdict. That is weaker than sealing and it is the price of evidence that cannot
exist when the phase ends.

**Evidence is produced in CI, not on a developer machine.** `evidence.source` in
`project.yaml` takes `ci` by default and `local` as a deliberate exception, and the
value in force stands in the artifact so that a reader sees under which regime a figure
came about. The reason is that locally produced evidence cannot be checked: a JUnit
report can be written by hand and a coverage figure can be typed. A report that lives in
the CI artifact store and is referenced from there cannot, because nobody can write
into it. It also removes every tool requirement from the developer machine, which is
what makes a scanner no longer a special case, and it makes the execution environment
uniform, which ends the question of which toolchain version a number came from.

The cost is named rather than hidden: a phase that declares evidence cannot be finished
without a CI round trip. Working offline stays possible, finishing does not. And the
retention of the pipeline becomes the lifetime of the evidence, which turns the last
limitation of this section from a footnote into a property of the design.

Both kinds come from the same place, the pipeline, so what decides where an item lives
is its size and not its origin. Small textual results, a JUnit report or a coverage
summary, are copied into `evidence/` in the repository, which makes the phase self
contained and is the only way the content itself outlives CI retention. Large or binary
output stays in the artifact store and is referenced by `uri` and hash, to keep the
repository usable.

Either way the hash and the commit are in the repository, and the declaration is what
binds them. `evidence/` sits outside the phase's `artifacts_hash`, so an item is not
covered by being in the repository; it is covered because its `sha256` stands in the
declaration, the declaration stands in `output.md`, and `output.md` is inside the hash.
The chain holds for `uri` items and in-repo items alike, which is why they are described
the same way.

`result` carries what the run reported against its own threshold, and it is not a
statement by Xeno. It is also not the question whether the tool ran: a test run with
failing tests is a successful run reporting `fail`.

Which items have one follows from the producer, not from the kind. A coverage run with a
minimum, a scan with a configured threshold and a linter with a rule set all report pass
or fail, and recording that is right. A coverage report without a threshold, a software
bill of materials and anything purely descriptive report nothing, and the field is absent
there. It is required on `test-report` and `build-log` because G-Test and G-Build read
it; everywhere else it is written where it exists.

What must not be read into it is a verdict on the thing examined. A scan reporting `pass`
means the tool checked against its own threshold and was satisfied, not that the code is
clean. Whether the threshold is the right one is a question Xeno declines to answer, and
the threshold stays where it is configured.

Nothing parses the report either. Test output
differs by ecosystem, JUnit XML, TRX, TAP, the JSON a Go test run emits, and every parser
for them ages; the same reasoning already applies to scanners, whose findings Xeno
records without judging. What stops a failing test at the merge is the pipeline, which
fails on the run itself. What the trail adds is that the run happened, on which commit,
with which command, and what it returned.

`format` names the shape of the report without implying that anything reads it. It is
optional and exists so that a later version can parse where parsing is worth it, without
a schema change and without guessing from a file extension.

A gate checks that every declared item resolves and that its hash matches. Whether
anything sits in `evidence/` without a declaration is a structural question and belongs
to G-Schema, which runs from P0 and therefore covers the phases where an `evidence/`
directory may appear before the first one G-Evidence applies to. An undeclared file is
not evidence, it is a declaration somebody lost.

`pipeline`, `job` and `commit` record where a result came from and what it was produced
against. No gate reads them. Requiring the commit to be the one under review cannot hold
in a project that squashes, where after the merge no evidence would match any commit, so
what makes evidence current stays what it was: the content hash still matches, the item
sits in the phase whose verdict is being read, and the phase itself is not stale, which
G-Freshness decides. As provenance the three fields are worth having anyway, and they
may point at nothing after a squash in the same way a `uri` points at nothing once CI
retention has lapsed.

A green pipeline is not evidence in this sense. It proves something at the moment it
runs and leaves nothing in the repository, and CI retention removes the run long
before an auditor asks. Scanners are therefore declared like builds and tests:
`kind: scan`, with the command, the commit, the hash and the job's own result. What the
trail then shows is that a scan ran, which one, how it ended, and that its result
belongs to this commit. What it does not show is the threshold it was judged against, or
what counted as new; both stay in the scanner's quality gate and are enforced by the
pipeline failing, not by a gate reading a report.

External evidence can expire when CI retention lapses. The hash and the declaration
survive, the content does not. That limit is named rather than papered over.

### Template resolution

The full template set ships with the plugin under `.xeno/plugin/templates/` and is
pinned and unchanged there. `.xeno/config/templates/` holds project overrides only.

Resolution is per template id: project beats plugin. A project can replace a single
template without forking the set. The path actually used is recorded in
`context.lock.yaml`, because otherwise two projects on the same template version are
indistinguishable even though one of them overrode it.

Templates are not a standardised component type. Agent Plugins v1 covers `skills/`
and `mcp.json` only, so `templates/` sits outside the portable core. That is
harmless here: templates are read by the runner and the MCP server through the file
path, never by the plugin loader.

### Secret filter

The secret filter is a file, not a rule kind. Secret detection is pattern matching
over text and has none of the properties the rule model in section 9 is built for:
no phase binding, no precedence, no prose statement.

```yaml
# .xeno/plugin/secrets.yaml, shipped, and .xeno/config/secrets.yaml, project additions
patterns:
  - id: aws-access-key
    regex: ...
  - id: bearer-token
    regex: ...
paths_never_digested: [ "**/*.pem", ".env*" ]
```

The project file is additive. It can add patterns and paths, never remove a shipped
one, because a filter a project can switch off is not a filter. The hash over the
effective filter is `secrets_hash` in the frontmatter, so it is reconstructable which
filter a digest was written under. G-Secret and the digest writer read the same file,
which is what makes the claim that they share rules verifiable rather than stated.

## 5. Artifact schema

Two kinds of artifact exist, and they carry different fields. Phase artifacts record
how they were produced. Intent level artifacts belong to the intent as a whole and
have no phase, no template and no model behind them. Fields are frontmatter in
Markdown artifacts and top level keys in YAML artifacts.

```yaml
# every process file
intent: <qualified id>
created: <iso8601>
schema_version: <major.minor>
runner_version: <gate runner version>
plugin_version: <xeno plugin version>

# phase files add
phase: 02-design

# files produced in a session, output.md and digest.md, add
language: <ietf tag>
secrets_hash: <sha256 over the effective secret filter>
context_hash: <sha256 over context.lock.yaml>
model: <model identifier>
tool: <claude-code|codex>
tool_version: <version>

# rendered files, output.md only, add
template: design@1.4.0
strings_hash: <sha256 over the strings bundle used>
rules_hash: <sha256 over the effective rule set>

# intent.yaml adds
key: PROJ-123
status: <in-progress|abandoned>
reason: <text>              # required when status is abandoned
```

Three groups, not two, because `digest.md` is neither rendered nor written by the
agent. The agent supplies the summary text, the runner filters, hashes and writes the
file, so a digest carries no template and no strings bundle. What it does carry is
where it came from: model, tool, language and the filter it passed through.

**`schema_version` says which shape an artifact has, and it is read rather than
enforced backwards.** It exists so that a reader, and a later runner, can tell what an
artifact is without guessing from which fields happen to be present. The minor moves
when a field is added that an older reader can ignore, the major when one changes
meaning or goes away.

An artifact written before the field existed does not carry it, and its absence is the
schema that predates versioning rather than a field somebody forgot. A gate that failed
on it would invalidate every artifact behind the change that introduced it, which is
the rewriting this process avoids everywhere else. A runner reads every version it
knows and writes only the current one.

It sits beside `runner_version` and is not derived from it, because the two move for
different reasons. A runner that has released twice without touching the shape of an
artifact has one number moving and the other standing still, and a schema read out of a
tool version would report a break that never happened.

No artifact carries its own gate status. The frontmatter says how a file came about;
what somebody concluded about it afterwards is a different statement, made at a
different time by a different writer, and it lives in one place only.

### Gate result

`gate.yaml` is written by the runner and by nothing else.

```yaml
intent: <qualified id>
phase: 02-design            # absent at intent level
created: <iso8601>
runner_version: <version>
plugin_version: <version>
status: <green|red|provisional|approved|overridden>   # derived, never written directly
run_at: <iso8601>
artifacts_hash: <sha256 over the phase directory, see below>
rules_hash: <sha256 over the effective rule set at run time>
secrets_hash: <sha256 over the effective secret filter at run time>
checks:
  - gate: G-Policy
    result: <pass|fail|pending|not-implemented>
    provenance: <xeno|external>
    findings:
      - id: F-7a3c91
        file: phases/02-design/output.md
        cause: <text>
        next: <text>
        advisory: true                  # absent in the ordinary case
        decision:                       # absent while undecided
          type: <approved|overridden>
          by: <person>
          at: <iso8601>
          against: <artifacts_hash the decision was made on>
          reason: <text>
          obligation: <open|closed>     # only with type overridden
drift:                                  # empty in the ordinary case
  - file: phases/02-design/output.md
    field: <rules_hash|secrets_hash>
    artifact: <sha256 recorded in the artifact>
    gate_run: <sha256 in force at run time>
```

`pending` is the result of a check whose declared evidence a pipeline has yet to
produce, and a phase with a pending check and no undecided failure is `provisional`. The
order from loudest is red, overridden, provisional, approved, green: an undecided failure
outranks everything, and missing evidence outranks a decision about something else.

`not-implemented` exists because the gate list is a declaration and a runner may be
older than it. A gate that is skipped silently makes a green verdict mean less than it
appears to, and the gap never surfaces afterwards. The state is written, it is not a
failure, and the derived status treats it as neither pass nor fail: a phase with
unimplemented gates is green in what was checked and says which checks were not.

**Decisions sit on findings, not on phases.** Where three checks fail and somebody
accepts one of them, a phase wide verdict would hide the other two. Each finding
therefore carries its own decision, and the phase status is derived from them: no
failure is `green`, an undecided failure is `red`, all failures decided is
`approved`, and one `overridden` decision makes the phase `overridden`. The louder
state wins. Nothing writes `status` directly, which is what keeps it from disagreeing
with the findings below it.

The four values are not interchangeable. `red` means nobody has decided anything yet.
`approved` means a person assessed the failing finding before the merge and accepted
it; the finding stays failed, the approval is the governance statement this process
exists to record. `overridden` means somebody took the merge first and still owes the
artifacts, which is why it carries an obligation and `approved` does not.

**An advisory finding does not fail its check.** `advisory` marks a finding that is
reported rather than held against the phase: the check carrying it is `pass`, the
phase is `green`, and the finding is in `gate.yaml` with its id, its cause and its
remedy like any other. Where a check asks for something nobody has experience with
yet, blocking against the answer would be the wrong way round, and a check that fires
with nothing to do about it is one people learn to route around. What the field
prevents is the alternative, which is leaving the measurement out and letting the
number become decoration.

It is the exception and stays one. A finding is the thing that fails, and an advisory
one is readable only because it is rare; the clause that asks for it says so where the
check is described, and nothing else writes the field.

**Finding ids are stable across runs.** The runner rewrites `gate.yaml` on every run
and carries decisions forward, so an id must not depend on the run. It is a hash over
gate, rule id, file and cause, with an empty rule id where a gate has none: G-Schema
and G-Trace report on structure rather than on a rule, and their findings are
identified by the other three parts alone. Appendix B defines it to the byte. Where the
content a finding rests on changes, a new id appears and the old decision falls away
with the old finding. That
is the intended behaviour: a release covers what was failing when it was given, never
what fails at the same place afterwards.

Findings from an external gate are the exception. Their cause is produced by foreign
code and is stable only by that code's promise, so a decision on one does not survive
the next run. Where a project wants a lasting release for an external finding, the
honest form is a rule of its own rather than a decision that silently depends on
somebody else's wording.

Every run that rewrites `gate.yaml`, including `xeno phase finish`, carries existing
decisions forward by finding id. The file is written whole each time, so a run that
forgot them would quietly discard the one thing in it a person contributed.

Three commands write these blocks, and the runner remains the only writer:
`xeno gate approve <finding-id> --reason`, `xeno gate override <finding-id> --reason`
and `xeno obligation close <finding-id>`.

**A verdict says what it judged, by content and not by commit.** `artifacts_hash` is a
hash over the files lying directly in the phase directory, without descending into
subdirectories. Two are excluded: `gate.yaml` itself, because a hash cannot cover the
file that carries it, and `cost.yaml`, because it is written at phase completion from
session logs and may arrive after the gate ran. The rule behind both: the hash covers
what the gate judges, not what the runner writes while judging. Every file it does cover
is one Xeno wrote itself, which is what keeps the definition free of binary cases and
text detection. Appendix B defines the computation to the byte.

It is the second hash an artifact is involved in, and the two answer different
questions. `context_hash` in an artifact's frontmatter covers `context.lock.yaml` and
says what the phase was produced from. `artifacts_hash` in the verdict covers the whole
phase directory, `context.lock.yaml` included, and says what the gate judged. One looks
back at the input, the other sideways at the set being assessed.

Binding by content rather than by commit is what makes the trail survive a squash
merge. A squash replaces the commits and leaves the trees alone, so every commit hash
recorded before the merge stops resolving while every content hash keeps resolving. A
verdict that names a commit range would be worthless the moment a project squashes; one
that names the state it judged stays checkable by recomputation, in the merged branch,
years later.

**No artifact names a commit it brings about itself.** A decision, an abandonment and a
confirmed assumption are all written by a command that runs before the commit carrying
it exists, so a field for that commit could only be filled with the wrong one or amended
afterwards, which is the writing back this process avoids everywhere else. What
identifies a decision is `by`, `at`, `reason` and `against`; the commit is found through
git, which is what git is for, and it survives a squash better than a recorded hash
would.

`against` carries the `artifacts_hash` the decision was made on, and that is what makes
it verifiable. A decision whose `against` no longer matches the directory is stale by
construction, which is the intended behaviour: whoever edits an artifact after a release
loses the release.

The commit range remains an input of the run and is not recorded. It cannot be recorded
honestly, because after a squash it would point at commits that no longer exist, and a
field that is sometimes wrong is worse than no field. What a commit predicate saw is
therefore not reconstructable after a merge, which belongs with the other limitations of
commit predicates rather than in a field that suggests precision.

**Drift is recorded, not failed.** `rules_hash` and `secrets_hash` in `gate.yaml`
describe the run; the same fields in an artifact describe its authoring. They come
apart whenever a rule set is merged between the two, which the learning mechanism in
section 10 makes routine. Failing on that would invalidate every phase behind a newly
merged rule, so the difference lands in `drift` instead.

Where it is then answered depends on the phase. Drift on P0 to P4 is known by the
time P5 is written, so the renderer turns each entry into a line of the P5 checklist.
Drift on the P5 artifact itself is found by the gate run that follows the rendering
and stays in `gate.yaml`, where the reviewer meets it alongside the findings. A second
rendering pass after the gate would be the only way to get it into the artifact, and
it would buy nothing the gate result does not already show.

There is no status for "not yet checked". The absence of `gate.yaml` says that, and a
state derived from a missing file cannot go stale. It is also what tells a failed run
apart from a failing one: a runner that could not evaluate writes no verdict at all.

The derived status is what a caller sees as an exit code. Green, approved and overridden
proceed; only an undecided failure stops. A phase whose failing findings have all been
decided is the case somebody released on purpose, and stopping there would work against
the decision that was made to let it through.

G-Schema rejects a decision without the fields its type requires. A gate result may
not be edited by hand.

`context.lock.yaml` describes the information base:

```yaml
intent: <qualified id>
created: <iso8601>
runner_version: <version>
plugin_version: <version>
repo_commit: <sha>
files: [{ path: ..., sha256: ..., bytes: ... }]      # empty at P0, see below
rules_applied: [{ path: ..., version: ... }]
template_source: <plugin|project>
plugin: { version: ..., sha256: ... }
tools: []        # context tools used, each with name, version, response hash
```

Each entry of `files` carries the size of the file as well as its hash. The size is
recorded because the budget is judged against what the phase was given and not against
what the tree holds now: measured at the time of the check, a file that grew after a
verdict would move a sealed phase's standing, and every other number in this process is
judged against what the artifact recorded.

`tools` names every context tool a phase used, with its version and the hash of its
answer: in v1 the symbol index, in v2 a code graph or a comparable tool without a
schema break. It is empty where a phase ran without any, which is allowed.

**It records the context that was declared, not everything that was read.** The file is
written before the agent starts, from the context scope, and nothing stops an agent
from opening a file the profile does not name. Treating it as a measurement of what was
read would be wrong; it is the statement of what the phase was given.

It is written once per phase and not refreshed. Where the agent changes a file that the
lock lists, the recorded hash no longer matches the working tree, and that is correct:
the lock describes the input state, and a lock rewritten at the end would describe
nothing.

**At P0 it names no files.** The lock is written from the context scope, and the scope
is P0's own artifact: `xeno scope set` writes it into the phase directory, so it cannot
run before the phase has started. The order is forced, the lock is born before the thing
it resolves, and nothing fills it in afterwards: the lock sits inside the phase's
`artifacts_hash` and is the only record of what the phase was given, so writing it a
second time would rewrite a hashed artifact and destroy the answer to what changed. An
intake's `files` is empty, and the two checks that read it apply from P1 on, each saying
so where it is described. What P0 declares is binding all the same: the scope it writes
is the information base of every later phase of the intent, which is where the budget is
counted and where a file moved out from under a reading is found.

Version numbers appear in both places on purpose: the frontmatter names what was
used, `context.lock.yaml` proves it with a hash. Where the two disagree, the hash
wins and G-Supply fails.

### Language

Structure and text are separated. `template.yaml` holds stable section ids
(`acceptance-criteria`, `non-goals`, `open-questions`), their order and the
required fields. The strings bundles hold headings and guidance text. Adding a
language means adding a bundle, never touching the structure.

The process layer is always English and is never translated: YAML keys, status
values, gate names, paths, learning categories, section ids. Only artifact content
follows the project language.

```yaml
# project.yaml
language:
  artifacts: de
```

Gates match on section ids, never on headings. A gate that matches text would be
language dependent and therefore broken.

Language is part of the information base, so `language` is mandatory in every file
produced in a session and `strings_hash` in every rendered one: two runs with the
same structure but different languages produce different artifacts. If the bundle for
the configured language is missing in the template version used, G-Schema fails red.
Falling back to English silently would be exactly the kind of unstated assumption this
process excludes.

### Context economy

Brownfield repositories are large and phases repeat, so reading is the dominant
cost. Five mechanisms keep it bounded, none of them language specific:

**The context scope declares the information base, and bounds it.** P0 produces it as an
artifact of its own, and a phase reads what it names. Any extension beyond it is written
into `context.lock.yaml`, which makes visible where the reading actually happens.
Without that visibility, any optimisation is guesswork.

```yaml
# phases/00-intake/context-scope.yaml
include: [ "src/payment/**", "docs/adr/**" ]
exclude: [ "**/testdata/**" ]
links:                      # declared, never inferred
  - component: src/payment
    docs: docs/adr/0012-payments.md
budget:
  files: 120
  bytes: 400000
```

G-Schema validates the profile like any other artifact, and reports a finding where the
recorded context exceeded the budget. The finding is `advisory`, so G-Schema stays
`pass` and the phase stays green. That is deliberately a finding and not a red gate
in the sense of stopping work: it is visible, it can be decided like any other finding,
and blocking against a number nobody has experience with yet would be the wrong way
round. What it prevents is the budget quietly becoming decoration.

The budget is judged from P1 on. The comparison is against the lock of the phase being
gated, and P0's lock names no files for the reason given where the lock is described, so
the intake declares the budget and the five phases that inherit its scope are the ones
judged against it. That is what keeps one number per intent from being decoration while
the phase that writes it cannot be measured against it.

**Re-reading follows change.** `context.lock.yaml` already carries paths with
hashes, so a repeated phase knows which files changed and reads only those.

**A symbol index, not a graph.** An index of symbols with file, line and enclosing
container answers "where is X" without semantic analysis, at a fraction of the cost of
a call graph, and it is the largest single saving. It is produced by the project and
read by Xeno, which ships no indexer: every tool carries its own licence terms and its
own idea of a symbol, and which of them a project can accept is that project's decision.
A project's own toolchain already knows its languages better than a tool built around a
process can, so what Xeno fixes is the format of the answer and not the means of getting
it. What Xeno requires of the index is that it names the tool and version that produced
it and the time it was produced, because a stale index is worse than none and an index
that cannot say how old it is cannot be judged.

It is used while a phase runs and never by a gate, which is why it may be missing,
stale or wrong without the trail suffering: a verdict never depends on it. It lives
under `.xeno/local/`, is not committed, and is covered by the same retention as the
rest of the local data. What enters the repository is the `tools` entry in
`context.lock.yaml` that names it.

**Digests instead of re-derivation.** Later phases read the earlier phase's output
and digest rather than searching the codebase again. P3 works from the design, not
from a fresh scan.

**A stable prefix.** Context is assembled in order of volatility, not in the order it
was thought of: the shipped set first, then what is stable for the project, then what
changes per phase, and the current task last. Every harness worth using caches the
unchanged head of a prompt and charges a fraction for it, and the cache reaches
exactly as far as the first difference. A profile listing an intent id or a timestamp
near the front would forfeit the whole saving, which is why `context.lock.yaml`
records the assembly order and not only the set.

Links between code and documentation are declared in the context scope, not
inferred. An inferred mapping is an assumption, and assumptions in this process are
either registered or absent.

A semantic graph of calls and class relationships stays out of v1. The extension
point exists: the `tools` block in `context.lock.yaml`. Binding an external tool
there is cheaper than maintaining parsers.

Worth stating because it is easy to miss: the gates cost nothing. No gate calls a
model, so the entire governance layer is free at inference time and its cost is
runtime in CI. What costs tokens is the agent doing the work, and every lever below
acts on that.

### Rendering

`output.md` is not written freehand. The runner renders it from `template.yaml` and
the strings bundle: sections in the order the structure defines, each heading taken
from the bundle, each section carrying its id as an anchor. Only the section content
comes from the agent.

**The agent supplies content per section and the runner renders on every write.** The
file on disk is therefore always in rendered form, which is what lets the hook check a
completed artifact against its anchors, and the anchors can never be wrong because the
agent never writes them. It also means the agent needs to know neither the strings
bundle nor the configured language, which is the point of separating structure from
text. Nothing is rendered again at the end of the phase; by then the file is what it
will be.

The MCP operation is the supported path, not an enforced one. An agent can write the
file with its own tools, and no runner can prevent that. G-Schema is the backstop:
whoever writes by hand and misses an anchor fails there.

An anchor is an HTML comment on its own line before the heading, and a section runs
from its anchor to the next one or to the end of the file. There are no end markers.

```
<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria
```

A comment rather than the attribute syntax some Markdown dialects offer, because
CommonMark and the dialect a code host renders do not know `{#id}` and would print it
beside the heading as text. A comment is invisible in every renderer, is found with one
line of parsing, and is independent of the heading, which comes from the strings bundle.

Gates therefore match anchors, never headings, and the same artifact can be
re-rendered in another language without changing what it asserts. A file whose
anchors do not match its declared template version fails G-Schema.

### Acceptance criteria are identifiable

From `requirements@1.1.0` the `acceptance-criteria` section is a numbered list, and
from `verification@1.1.0` the `test-mapping` section names each criterion by its
number. Section 7 asks G-Test for the completeness of the mapping, and nothing could
answer it because nothing identified a criterion: a sentence cannot be reported as
covered or uncovered, so there was nothing to count.

The requirement is carried by the template version and not by the schema. An artifact
declaring `requirements@1.0.0` is judged as it always was. What is sealed is never
rewritten, and a check reaching backwards would re-judge every mapping a project wrote
before it adopted the convention. `template` is already in the frontmatter of every
artifact, so the anchor is read and not added.

## 6. Phase model

| Phase                | Input    | Output                                                      |
|----------------------|----------|-------------------------------------------------------------|
| P0 Intake            | issue    | `intent.yaml`, context scope, initial assumption register    |
| P1 Requirements      | P0       | specification with acceptance criteria and non goals         |
| P2 Design            | P1       | design decision, affected components                         |
| P3 Implementation    | P2       | code, draft merge request                                    |
| P4 Verification      | P3       | test evidence, mapping of acceptance criteria to tests       |
| P5 Review and merge  | P4       | approved merge request, review result including the release notes section |

**P4 proves the verification, it is not when tests get written.** The order of the
phases is the order in which the record is assembled, not a licence to write tests
afterwards. Acceptance criteria exist from P1, and an organisation that requires a test
first discipline expresses it as a rule over the commit range rather than reading it
out of the phase numbers. Saying so here costs a sentence and saves the misreading that
the table invites.

Release notes are a section of the P5 artifact, id `release-notes`, not a file of
their own. Everything that already applies then applies to them: rendering, anchors,
the strings bundle, G-Schema. Whether the section must be filled is a shipped
`checked` rule, not a property of the phase model.

Operations and maintenance are not part of v1. The phase list in the schema is kept
open so that a phase P6 can be appended later without a break.

### Working sequence

Who does what, with which command, and what lands in the repository. The process is
described everywhere else in this document by what it guarantees; this is the order in
which somebody carries it out.

**Once per repository.**

| Step | Who | Command or act | Result |
|---|---|---|---|
| Initialise | a developer | `xeno init --vendor` | `.xeno/` with config, vendored plugin, CI wrapper |
| Configure the host | somebody with repository administration | by hand, from what `xeno init` printed | protected branch, required pipeline, approval settings where the host has them |

**Once per phase**, repeated for P0 to P5.

| Step | Who | Command or act | Result |
|---|---|---|---|
| Start | a developer | `xeno phase start --intent <id> --phase <n>` | `context.lock.yaml` written, issue content read where a tracker is configured |
| Work | the agent | in the harness, through the plugin | `output.md`, the phase's own files, evidence under `evidence/` |
| Finish* | a developer | `xeno phase finish` | `digest.md`, `cost.yaml` where cost recording is available, gates run, `gate.yaml` with its `artifacts_hash` |
| Hand over* | a developer | commit and push | the phase is in the repository, complete |
| Verify | CI | the generated wrapper | enforcement report, gates recomputed and compared, result written back to the merge request |

\* In a phase that declares evidence the last two steps behave differently, see the note
below the next table.

**In a phase that declares evidence**, which in practice means P4, nobody waits for the
pipeline. `xeno phase finish` runs straight after the work and judges everything it can
decide locally, declaring each expected evidence item by kind and job. The phase is
pushed and the pipeline runs once. Nobody waits for it.

CI still writes nothing. It publishes its reports and is done, which keeps the rule in
section 12 intact rather than carving an exception out of it. A commit from a pipeline
would in any case start the next pipeline, whose last job would commit again.

The results are pulled in by the next command that writes anyway. `xeno phase start` of
the following phase first attaches whatever has since arrived for its predecessor,
re-evaluates the gates that depend on it and records the outcome in the commit it makes
regardless. No extra commit, no loop, no second writing identity. The attachment lands in
`evidence/attached.yaml`, outside the `artifacts_hash`, which yields the invariant the
whole arrangement rests on: **only the run of a person writes inside the
`artifacts_hash`, and every later writer touches only what lies outside it.**

`xeno gate run` attaches as well, before it computes, because P5 has no successor whose
start could do it. Without that, anything P5 declares itself would stay open for good and
the intent would look provisional for ever.

The attachment is the one place where a phase writes outside its own directory, and it is
kept narrow on purpose: a `phase start` may write `evidence/attached.yaml` and
`gate.yaml` in the directory of its immediate predecessor, nothing else and nowhere else.
Stated here so that it does not quietly become a general write path.

A verdict with outstanding declarations is provisional, and that is a state of its own
rather than a disagreement. The CI wrapper reports "provisional, N declarations open"
instead of failing, since the local run knew they were open and so does CI; only at the
merge request does it become a hard condition. Without that distinction every push out of
P4 would fail verification and nobody would take verification seriously again.

The waiting therefore falls where it belongs. P4 finishes without delay; whoever starts
P5 may have to wait briefly, and that is the transition where a green pipeline is wanted
anyway, since P5 is review and merge. If evidence is still missing, `phase start` refuses
with a plain reason rather than with a red gate.

What this costs belongs in the limitations: the verdict of a phase with outstanding
declarations is provisional until somebody carries on, and `xeno intent status` has to
say so, or an intent with open declarations looks finished. If nobody carries on it stays
provisional, which is honest, because then nobody judged it either.

**Where a finding fails.**

| Step | Who | Command or act | Result |
|---|---|---|---|
| Fix | a developer | rework the artifact and `xeno phase finish` again | new `artifacts_hash`, the finding disappears or persists |
| Or release it | a second person | `xeno gate approve <finding-id> --reason` | the decision on that finding, with `against` |
| Or take it anyway | a second person | `xeno gate override <finding-id> --reason` | the same, plus an open obligation |
| Later | whoever owes it | `xeno obligation close <finding-id>` | the obligation is closed and counted |

**At the end of the intent.** A merged intent meets G-Complete as part of P5, before the
merge, which is where it has to be: a gate that reports after the merge cannot gate it.
An abandoned intent runs `xeno intent close --reason`, which writes the closing learning
record and the intent level `gate.yaml`.

**The sequence is enforced, not assumed.** `xeno phase start` refuses a phase whose
predecessor holds no completed `gate.yaml` with a matching `artifacts_hash`, and one
whose predecessor is red or provisional. The next phase starts only on a decided
predecessor: green, approved or overridden. Building on a red verdict moves the failure
downstream instead of resolving it, and every failing finding has to be fixed, approved
or overridden before the work goes on, which is the human decision this process exists
to record, taken earlier rather than just before the merge. In a
single developer setting the order is kept by the person typing the commands and the
check adds nothing; it is here because the guarantee should be a property of the tool
rather than a habit, and because P2 reads P1 and P3 reads P2, so a phase started out
of order has no defined input. The same check refuses a second start of a phase that
is already running.

**The local run is not optional.** Its verdict binds the start of the next phase, and
running it is mandatory, because it is the only writer of `gate.yaml`, and a phase without a verdict in
the repository is incomplete. CI recomputes and compares; it never writes.

What is compared is what the artifacts determine: the `artifacts_hash` and the findings
that follow from the content of the phase. Findings from commit predicates are not
compared, and cannot be. The local run happens before the commit that carries
`gate.yaml` exists, so it sees a shorter range than CI does, and the two would differ on
every single run. CI computes those findings fresh and its result is the one that
counts; locally they are feedback. A phase
pushed without its `gate.yaml` fails in CI with nothing to compare against, which is a
clear failure rather than a silent one, but it is found at the end rather than during
the work. `xeno phase finish` exists so that the obligation is one command instead of a
habit.

That makes the pinned runner version a requirement rather than a convenience: every
machine that works a phase needs it, in the version `project.yaml` names, or it produces
a different hash or a different verdict than CI. `xeno init` refuses a mismatch for that
reason.

**Changes made outside the runner expire, they are not forbidden.** Editing an artifact
through the code host's web interface, or resolving a rebase conflict inside `.xeno/`,
changes the phase directory without a gate run. The `artifacts_hash` in `gate.yaml` then
no longer matches, the verdict is stale, and so is every decision whose `against` points
at it. CI finds this. The repair is `xeno phase finish` and a commit.

Everything outside `.xeno/` is unaffected. Editing a source file in the web interface
breaks no verdict; where that file belongs to an earlier phase's context, G-Freshness
raises it, which is what G-Freshness is for.

**Rebasing is safe where the trees are.** A clean rebase produces the same trees under
new parents, so every content hash still resolves and nothing has to be redone, including
a rebase onto a moved default branch. Only a conflict resolved inside `.xeno/` changes a
tree, and then the paragraph above applies. Nothing else breaks, because no
artifact records a commit that a rebase could rewrite.

**The runner keeps its own writes consistent.** A command that changes a file inside the
hashed set has to recompute the hashes it invalidates, or the tool would invalidate its
own verdicts. `gate.yaml` and `cost.yaml` are outside the phase's hash for exactly this
reason; `intent.yaml` sits inside the intent level hash and is therefore rewritten
together with it.

### Ending an intent

Not every intent ends in a merge. Requirements get withdrawn, bugs turn out to be
misreports, work gets dropped. An intent can therefore be closed as `abandoned`,
recorded in `intent.yaml` as its status with a reason and the commit that closed it,
and `.xeno/intents/` does not fill up with half finished corpses.

An intent that reaches a merge records nothing about the merge, and needs no field for
it. Deployment tooling works from commit SHAs while this side holds phases and
timestamps, and the join between them is the intent key in the merge commit message,
which a squashed merge carries too because its message is built from the merge request
title. The reading direction that matters, from a release's commits to the intents in
it, works without anything being written back after the merge.

Abandoning an intent ends with a learning record, written to
`.xeno/intents/PROJ-123/learning.yaml` in the same format as the phase records, and
G-Complete checks it in that mode. Why something was dropped is usually worth more than
why it was built, and an intent that stops in P1 has one phase record and little else.

A merged intent writes no record at this level. It already carries one per phase,
including P5, which is where the retrospective of a finished piece of work belongs.

`xeno intent close` exists for the abandoned case. It runs G-Complete against the
closing conditions and writes the intent level `gate.yaml`, because an intent dropped
in P1 never reaches P5 and would otherwise never be measured at all. A merged intent
needs no such command: G-Complete has already run as part of P5, where it had to run,
since a gate that reports after the merge cannot gate it, and nothing is added
afterwards.

### Override

A process with no valve gets switched off entirely the first time it stands in the
way of a production incident. So there is an explicit override: a merge can proceed
on a failing finding, recorded on that finding with person, reason and timestamp, and
carrying an obligation to complete the missing artifacts afterwards. The obligation
is a field, so that it can be closed and counted.

An override applies to a named finding, not to a phase. Waving through everything
that happens to be failing at once is then a series of deliberate acts, each with its
own reason, rather than one act whose scope nobody has to state.

This is a governance function, not a weakening. The alternative is not a stricter
process, it is the same bypass happening invisibly. An override is the loudest entry
in the trail and the easiest to report on, which is exactly what makes it safe to
offer.

Nothing forces the follow-up. An override stays visible as an open obligation until
the missing artifacts exist, and that visibility is the whole mechanism. A gate that
could compel the follow-up would have had to block the merge in the first place,
which is the situation the override exists to resolve.

## 7. Gates

All gates are deterministic and model free. The order is integrity first, then cost.
G-Supply runs ahead of everything because a plugin that does not match its expected
digest makes every later verdict a statement about unknown rules and unknown
templates. The
rest follow cheapest first: a missing field is found by reading one file, while
G-Rules, G-Policy and G-Complete have to resolve the whole rule set or walk every
preceding phase. A structural failure should be reported before anything has to be
resolved or traversed.

| Gate           | Checks                                                                      | From |
|----------------|-----------------------------------------------------------------------------|------|
| G-Supply       | vendored plugin matches the digest the runner carries for its own version | P0 |
| G-Schema       | fields complete, frontmatter valid, template version and strings bundle exist, no unknown file in the phase directory and nothing undeclared in `evidence/` | P0 |
| G-Trace        | every artifact bound to exactly one intent, no orphan artifacts              | P0   |
| G-Secret       | filtering against the effective secret filter, the same file the digest uses | P0   |
| G-Assumptions  | no open, unconfirmed assumption                                              | P0   |
| G-Questions    | every open question raised in any phase resolved into a decision, a confirmed assumption or a withdrawal | P5 |
| G-Learning     | `learning.yaml` present and filled                                           | P0   |
| G-Freshness    | each phase's context hash points at the current state of its predecessor, and no file a preceding phase read was changed by the change under review afterwards | P0 |
| G-Evidence     | every declared evidence item resolves and its hash matches            | P3   |
| G-Build        | declared build result present, successful and current                        | P3   |
| G-Test         | declared test result successful, mapping of acceptance criteria complete     | P4   |
| G-Rules        | rule collisions resolved, `binding` respected, `scope` matches the path      | P0   |
| G-Policy       | conformance against the effective rule set, every review checklist entry answered | P0 |
| G-Complete     | see below, two modes                                                         | P5 and intent close |

G-Rules and G-Policy sit together at the end because they resolve the same rule tree
and should do it once.

**G-Complete has two modes**, because an intent has two ways of ending, and the mode
follows from where the gate was invoked rather than from a field. Run as part of P5, it
checks that the artifacts of all preceding phases are present and green, approved or
overridden. Nothing about the merge itself, which has not happened at the point the
gate runs. Run from `xeno intent close`, it checks that the intent carries a reason and
that the closing learning record exists. An abandoned intent is not exempt from the
gate, it is measured against different conditions.

There is no `merged` status for it to read. Such a field would have to be written
before the merge to be readable by the gate that precedes it, and would then claim
something that has not happened. `abandoned` is a state somebody enters deliberately
and can therefore be recorded; reaching the end of P5 is what a completed intent looks
like.

It is the one gate with two points of invocation. In the merged case it runs as part
of P5 and its result lands in that phase. In the abandoned case it runs from
`xeno intent close` and writes the intent level `gate.yaml`, because there is no
phase left for it to belong to.

That second result carries an `artifacts_hash` of its own, computed the same way one
directory higher: the files lying directly in the intent directory, excluding the
`gate.yaml` that carries them. Not descending is what keeps `phases/` out, which is
right, because each phase already has a verdict of its own. Without it the one
verdict that closes an intent would be the only one not bound to what it judged.

Green means the run proceeds without intervention. A failing finding means a human
decides, and the decision is recorded on that finding as `approved`, not rewritten to
green. In v1 human on the loop applies throughout. Risk dependent differentiation
follows in v2.

**Gates read, they do not run.** G-Build and G-Test evaluate declared evidence
rather than starting a build or a test suite themselves. CI runs those anyway, and a
gate that re-ran them would double the pipeline while proving nothing extra. What the
gate does prove is that the declared result was recorded, with the command that produced
it, and that its content still matches what was declared.

**Repetition is normal.** A phase may be redone at any time, typically because
requirements changed mid-flight. When that happens, every later phase whose context
hash points at the old state is stale. G-Freshness marks it as such rather than
letting it silently continue to count. A stale phase is not deleted, it is re-run or
explicitly approved as still valid.

**A phase also goes stale when the code moves under it.** `context.lock.yaml` records
every file the phase was given, with its hash, so the runner can recompute those hashes
and see where the working tree has moved on. Without that check a test report from three
commits ago passes untouched, because its own content never changed; nothing else in the
process would notice.

Two limits keep the check useful rather than noisy.

It looks at the files the change under review touched, not at everything the profile
names. The wider reading is the more honest one, since a design made against code that
has since changed is genuinely questionable, but in an active repository every rebase
onto a moved default branch would set it off, and a check that fires constantly is one
people learn to ignore.

And it looks at the locks of the preceding phases, never at the lock of the phase being
gated. A phase that changes the files it read is not stale, it is working: that is what
P3 does, and without this limit every implementation phase would report itself out of
date the moment it did its job. What the check catches is the other direction, a phase
whose ground moved after it finished.

P0 stands outside the check for a third reason, which is not a limit chosen against
noise. Its lock names no files, because the scope the lock would resolve is the artifact
P0 produces, so the files an intake was given are never compared against the tree. The
rest of the intent is compared against the locks of P1 onwards, which resolve the same
scope. What this loses is a file that moved while the intake itself ran, since the next
phase's lock records it as it was by then and that is the state the comparison starts
from.

### Execution

Two verdicts, four invocation paths, one binary and one rule set. The verdicts are
what governance rests on: local binds the sequence, CI binds the merge. Two different questions hide
in that word, though, and separating them is what the fourth column does: whether a run
decides anything, and whether it can be skipped.

| Invocation              | Scope                                            | Verdict             | In the sequence |
|-------------------------|--------------------------------------------------|---------------------|---|
| Hook, on any write      | G-Secret on the file just touched                 | advisory, immediate | optional |
| Hook, on artifact write | G-Schema on the artifact just completed           | advisory, immediate | optional |
| Local run, before push  | all gates of the phase                            | binding for the next phase | **required** |
| Project CI              | all gates of the phase                            | binding             | required |

**The local run binds the sequence and cannot be skipped**, and the second half follows from
two commitments made earlier rather than being a rule of its own. The repository is the
source of truth, so the verdict has to be in it. CI verifies artifacts that exist and
produces none, so CI cannot put it there. What has to be in the repository and cannot be
written by CI can only be written before the push, and the only thing that writes it is
the local run.

The second commitment is worth keeping rather than trading away. A commit from CI would
itself be a change to the repository, and the only thing that could judge that change is
another CI run, which would write again. There is no natural end to that, and the trail
would contain entries whose author is a pipeline. Keeping CI read only costs one
obligation on the person pushing, which `xeno phase finish` reduces to a single command.

So the two statements are one statement seen from both ends: because CI does not write,
the local run must happen; because the local run happens, CI has something to recompute
and compare against. A phase pushed without its `gate.yaml` leaves CI with nothing to
compare, and it fails for that reason rather than for a failing gate.

One setting carries that last row, and it does not live in this repository: the
pipeline has to be required for a merge on the protected default branch, which GitHub
calls a required status check and other hosts call something else. Without
it a red gate produces a report that a merge walks past, and every claim in this
document about a binding verdict is a claim about a setting somebody made elsewhere.

The failure is invisible from inside. The phase was finished locally, `gate.yaml` is in
the repository with its findings, CI recomputes them, the pipeline goes red, and the
trail is complete. Only the stopping does not happen, and a project can work for months
believing it has a binding gate when what it has is a report.

`xeno enforcement check` closes as much of that as can be closed. It reads the
protected branch configuration from the host and compares it against the `enforcement`
block in `project.yaml`, which states what the project requires of its host. It is a
command of its own and never part of `xeno gate`, because it needs the network and the
gate path's freedom from it is what makes a verdict reproducible. It runs as the first
job of the same pipeline and again on a schedule, since the change worth catching is
usually the one somebody made between two merges.

Its report distinguishes a setting that is not set from one the host cannot express at
all, which is a different sentence and the one a project's records need. And it cannot
make itself binding: where the pipeline is not required for merging, this check failing
stops nothing either. What it buys is that the gap is stated instead of assumed.

The hook runs a reduced set on purpose. Resolving the rule set or walking every
preceding phase on every file write would slow the agent enough that people switch
the hooks off, and a hook nobody runs catches nothing. What it does catch is the one
class of problem that is expensive later: a secret on its way into the repository.

**These are not git hooks.** A hook here runs in the agent's harness on a tool call
and sees a file about to be written. A git hook runs on commit or on push and sees a
commit. Both appear in this process and they are unrelated: the harness hook is what
Xeno installs with its plugin, the git hook is what a project sets up for itself, and
the shipped examples for the second kind live under `examples/hooks/` precisely because
they are not part of the runner.

The two hook rows differ in their trigger for the same reason. G-Secret fires on
every write, because a secret is a secret in a half written file too. G-Schema fires
only when a write came through the MCP operation that completes an artifact, since an
artifact under construction is missing fields by definition and a hook that reports
that on every keystroke is the kind of hook people turn off.

One runner, one thin wrapper. The CI job calls nothing but
`xeno gate run --phase <n> --base <ref> --head <ref>`, generated from a template into
the host's workflow file with the runner version pinned. Both ends of the range are
passed explicitly, because a CI system may check out a merge commit it produced itself
and a runner that took that for the head would judge a commit nobody wrote. GitHub
reports both ends on the pull request event. Locally `--base` defaults to the merge base against
the default branch, which git computes offline, and `--head` to `HEAD`.

The range is a parameter rather than something the runner works out for itself,
because a runner that guesses it would return different verdicts locally and in CI
from the same repository state.

Further CI systems follow in 1.1. Since the wrapper is generated and does nothing but
pass a phase and a commit range, a second system is a second template rather than a
second integration.

### Hooks

Hooks trigger the local stage, they do not perform the checks. The check is always
the runner. Hooks fire it while the agent works, so a problem surfaces immediately
instead of at the end of the phase.

Both Claude Code and Codex support hooks, with comparable lifecycle events and
different wiring: JSON settings on one side, `hooks.json` or an inline `[hooks]`
table in TOML on the other. A plugin can ship its own lifecycle configuration, so
the hooks travel with Xeno rather than with an installation guide. Where hooks are
absent, untrusted or switched off, nothing is lost except speed of feedback.

CI runs no hooks. It calls the runner directly.

**Environment normalisation.** The hook calls a thin entry point that normalises the
environment before starting the runner. The runner only ever sees `XENO_*` and does
not know which harness it runs under.

```
XENO_PLUGIN_DATA      local data location, always .xeno/local/
XENO_HARNESS          claude-code | codex | unknown, recorded only
XENO_HARNESS_VERSION  that harness's own version, recorded only
```

**The plugin is the vendored one.** `.xeno/plugin/`, found from the git root, and
nothing else: it is the one G-Supply hashes and `context.lock.yaml` records, and the
one `xeno init --vendor` writes out of the runner. There is no override and no
resolution order.

This paragraph used to describe one, headed by a `--plugin-root` argument and an
already set `XENO_PLUGIN_ROOT`, with a client's own variable at the bottom. It was
never implemented and it is removed rather than built, because every position in it
was either unreachable or unsafe. The gate path reads the rule set and the templates
from the plugin, so a root taken from the environment makes `rules_hash`,
`strings_hash` and a rendered artifact depend on it — and a phase is judged against
the rule set its own artifact records, so a changed set does not disagree with the
trail, it stops judging it. Measured on this repository: a valid rule tree that
differs leaves `xeno gate verify` at exit 0 over 273 verdicts while G-Policy silently
reports nothing about the 75 phases that recorded the previous hash. A released
runner would catch it, because G-Supply compares the vendored tree against a digest
the binary carries; a development build carries no digest and would not.

The client's own variable was the bottom of that order and is unreachable for a
different reason: a project with no vendored plugin resolves no template, so no phase
of it renders, and `plugin_version` is absent, so G-Schema reports it. Falling back to
a plugin the client installed lets such a project fail differently rather than work.

An override may be worth having the day somebody needs one, and the day it is added
the condition is that a runner which cannot verify a plugin does not accept one from
outside the repository.

`XENO_HARNESS` and `XENO_HARNESS_VERSION` are recorded, never branched on. The
moment the runner behaves differently per harness, the tools stop being
interchangeable.

The version is in this list because `tool_version` in section 5 is a property of
the session that produced an artifact, and the harness is the only thing that
knows it. A runner that worked the value out would have to ask which harness it
was under, which is the branching the paragraph above forbids. So it is reported
and never derived, and a harness that reports nothing leaves the field absent
rather than defaulted: a plausible value in a field nobody produced says a
session happened that did not.

**Two constraints.**

1. A hook only invokes checks the runner already implements. No exclusive logic, or
   the same input would yield different results depending on which agent produced
   it.
2. Its result is advisory. The binding result is the CI run, because hook
   configuration lives on the developer machine. Organisation managed hooks are
   harder to bypass, but managed configuration cannot be assumed in every project.

## 8. Assumption register

`assumptions.yaml` lives at intent level and is carried forward across all phases.
It is a mapping, not a bare list, so that it can carry the intent level fields from
section 5 like every other process file.

```yaml
intent: <qualified id>
created: <iso8601>
updated: <iso8601>
runner_version: <version>
plugin_version: <version>
assumptions:
  - id: A-004
    phase: 02-design
    assumption: <statement>
    origin: <template-default|repo-convention|rules|user-input>
    confidence: <high|medium|low>
    status: <open|confirmed|rejected>
    confirmed_by: <person>
    rejected_by: <person>
```

Every open assumption turns the gate of its phase red. At phase boundaries with
open assumptions the process is therefore effectively human in the loop.
Confirmation is recorded in the repository and pushed, because CI cannot reliably
see the tracker.

Both decided states name their person, and exactly one of the two fields is
present: `confirmed_by` where the status is confirmed, `rejected_by` where it
is rejected. Rejecting is a statement by a person in the same way confirming
is, and the register is the record of who made it. The gate is green either
way, so nothing downstream reads the name; what reads it is somebody asking
months later who dropped an assumption and why the work went the way it did,
which is the question the trail exists to answer without asking anybody to
remember.

### Decisions are not assumptions

An assumption stands in for missing knowledge, holds provisionally and can turn out
wrong, which is why it has to be confirmed. A decision is a choice between workable
paths: it cannot become wrong, only regretted, and it needs a responsible person and a
reason rather than a confirmation. Carrying one as the other would block a gate until
somebody confirms what they just decided themselves.

Decisions therefore stay where they already are, in the `decisions` section of the
artifact of the phase in which they were made, sealed with it. Only assumptions live at
intent level, because only they have a state that changes.

### Open questions have to be resolved

**A question is asked with options, not open.** Each one carries two to four options
with their consequence, the agent's recommendation with a reason, and always a free
entry as a further option. An open question without options moves the whole of the
thinking onto the person, which is what the agent was there to take off them.

The reason belongs to the option the recommendation names and is carried there rather
than on the question. A recommendation that later moves to another option takes its
reason with it or the writer refuses, where a reason held beside the question would go
on describing the option it used to be about with nothing able to notice. It is empty
on every option but one, which is what that costs.

Two things keep that honest. Where the agent finds no options, the question says so
rather than inventing two. And the free entry is not an escape hatch but the normal case
for a domain expert: what they enter becomes the chosen option and appears in the
decision with `decided_by`, not as an aside beside it.

The `open-questions` section is sealed with its artifact, so a question cannot be
answered where it was raised. Each question therefore carries a stable key, and it is
resolved in exactly one of three ways, each with a person behind it:

- as a decision, recorded in the `decisions` section of whichever phase settles it,
  carrying `resolves: <key>` and naming who decided and why;
- as an assumption in `assumptions.yaml` that somebody confirms, which is the answer for
  the case where nobody knows: a person accepts working on a placeholder and says so;
- as a withdrawal, also a `decisions` entry, with a reason.

The third exit has to exist. Without it people invent assumptions to get a gate green,
and a register full of pretence is worse than an open question.

`G-Questions` checks from P5 that every question raised in any phase has one of the
three. It does not apply to an abandoned intent: abandoning means giving up, and a gate
that demanded answers there would only produce invented ones. Not earlier, because an open question in P1 is normal and should not hold up the
work; before a merge it is not normal any more. The staggering follows by itself: a
question never holds anything up, an unconfirmed assumption always does, and whoever
turns a question into an assumption moves the hold point forward deliberately.

## 9. Rule model

Two axes: origin and level.

```
.xeno/plugin/rules/given/builtin/           shipped, pinned
.xeno/config/rules/given/{provider,org,project}/
.xeno/config/rules/learned/{provider,org,project}/
```

| Level      | Whose rules                                | Reach                                          |
|------------|--------------------------------------------|------------------------------------------------|
| `builtin`  | Xeno itself                                 | every project everywhere                       |
| `provider` | the delivering organisation                 | all its projects, across customers             |
| `org`      | the organisation the repository belongs to  | all projects of that organisation              |
| `project`  | this repository                             | here only                                      |

The path determines the level. The frontmatter repeats it as `scope` and is checked
against the path, so a moved file cannot silently change its reach.

The shipped set lives with the plugin, pinned and hashed like the template set.
`.xeno/config/rules/` holds only the levels a project or an organisation maintains
itself. There is no project side copy of the shipped rules: overriding happens
through precedence, never through a shadow copy on the same level.

`learned/builtin/` does not exist. Nothing learns into a pinned, hashed package. The
way into the shipped set is the pull request described in section 15, a deliberate
act rather than a derivation, and G-Rules rejects the directory.

### Rule format

```yaml
id: public-interface-change-requires-migration-note
version: 4                  # change counter, not a compatibility claim
scope: provider
binding: false
kind: <checked|review>
applies_to: [02-design, 03-implementation]
statement: >
  A change to a published interface must come with a migration note.
check:                      # required for kind: checked, forbidden for kind: review
  type: section-implies-section
  when: { section: interfaces, non_empty: true }
  then: { section: migration-notes, non_empty: true }
```

`version` counts changes and claims nothing about compatibility. Any edit to a rule
can make it reject more than it did before, so a major and a minor step would be the
same thing in practice, and nothing in the runner acts on the number anyway. What
proves which rule set was in force is `rules_hash`; the counter only says how often
this file has moved since.

`binding` is only allowed under `given/`. `abstract` is required in
`learned/provider/` and `learned/org/` and holds a transferable summary without
project or code specifics.

A deterministic gate cannot evaluate prose, so a rule is one of two kinds. A
`checked` rule carries a declarative predicate and is evaluated by G-Policy, red or
green. A `review` rule carries none: it produces an item in the P5 checklist that a
person answers with a result.

The checklist is a section of the P5 artifact, `review-checklist`, rendered from the
effective rule set with the rule id as the anchor of each entry:

```yaml
review-checklist:
  - rule: public-interface-change-requires-migration-note
    result: <met|deviation|not-applicable>
    note: <text>        # required for deviation and not-applicable
```

G-Policy checks that every entry carries a result, and nothing more. Whether the
answer is a good one is not a question a deterministic gate can ask, but whether the
rule was answered at all is, and without that check a review rule is a suggestion.
Passing over one is now a recorded deviation rather than an omission nobody sees.

A rule declared `checked` without a `check`, or `review` with one, is a
configuration error and fails G-Rules.

### Predicate types

Predicates are named types implemented by the runner, never free expressions. A rule
says `type: section-implies-section` or `type: commit-message` and parameterises it;
it does not carry a regular expression of its own. A regex in a rule file would be
code without provenance, which is the line external gates exist to draw.

A commit message pattern is also reachable outside a gate, through
`xeno check commit-message`, which evaluates one message rather than a range. That is
what a client side or server side git hook calls, so the pattern exists once and is
evaluated by one implementation wherever it is checked. Copying the expression into a
hook is how a push starts rejecting what a gate accepts.

Several shipped types read the commit history of the change under review rather than
an artifact:

| Type               | Reads                                                              |
|--------------------|--------------------------------------------------------------------|
| `commit-message`   | the subject of every commit in the range, against a named shipped pattern |
| `commit-trailer`   | a required trailer on every commit in the range, `Xeno-Intent:` being the shipped case |
| `commit-signature` | a valid signature on every commit in the range                     |
| `approver-not-author` | the person who decided a finding against the authors of the range |

```yaml
# examples/rules/conventional-commits.yaml, adopted by copying into given/org/
id: conventional-commits
version: 1
scope: org
binding: false
kind: checked
applies_to: [05-review]
statement: >
  Every commit of the change under review follows Conventional Commits.
check:
  type: commit-message
  range: change-under-review
  pattern: conventional-commits     # named, shipped pattern
  exempt: [merge-commits]
```

The commit range is always passed in, never inferred by the runner, for the reason
given in section 12. The `commit-trailer` type is what binds code to an intent from
the commit side: G-Trace proves that no artifact is orphaned, and a project that
wants the other direction proven adopts the trailer rule. That direction is not part
of G-Trace itself, because a mandatory gate would impose a commit convention on every
project using Xeno, which is exactly what the example mechanism above exists to
avoid.

`commit-signature` exists for the same reason as the other two and ships the same
way, as an example rather than as a rule. Signatures are a project decision: a code
host that authenticates pushes and merges already answers most of what the signature
would prove, and the part it does not answer is named in section 16 rather than
settled by a rule nobody asked for.

`approver-not-author` is the separation of duties made checkable. Where an
organisation requires that the person releasing a change is not the person who wrote
it, the requirement is usually a sentence in a policy and a habit in a team. It is
deterministic: the decisions in `gate.yaml` carry `by`, the range carries its authors,
and the predicate compares the two sets.

Two things about it are worth saying rather than discovering. The decision is written
after the gate run that produced the finding, so the predicate is evaluated on the
next run, which the commit carrying the decision triggers anyway. And it compares two
self-asserted strings: a git author line is free text unless the commit is signed, so
the rule enforces the discipline of a team that means it and does not defeat somebody
who does not. That is still worth having, since the case it catches is the ordinary
one, an approval given by the author out of haste rather than intent.

The number of review rules is not limited. It is worth watching all the same: every
review rule adds an item to the P5 checklist, and a long checklist is where the
process starts to feel like overhead. Where a rule can be expressed as a check, it
should be.

### Shipped rules

The set shipped under `given/builtin/` is deliberately minimal and generic:
traceability of deviations, migration notes for interface changes, a rationale for
new dependencies, a filled release notes section at P5. Anything specific to a
supplier, a customer or a project belongs in that party's own layer, not in the
package.

**Patterns are shipped, expressions are not written into rules.** A rule names a
pattern from the shipped set and parameterises it. The set covers Conventional Commits
and the same form carrying an issue reference in the subject, which is what a project
needs when the reference has to survive a squash merge, since the squashed message is
built from the merge request title.

Both are regular expressions, so the distinction is not technical. It is about who
writes the expression and who answers for it. A shipped pattern has an identifier, a
version, tests in the corpus, and it changes only with a release; a project references
it. An expression in a project's own rule file is written once under time pressure,
versioned as text, untested, and slightly different in every project that needs the
same thing. Four consequences follow, and none of them is a matter of taste.

*The failure message.* A gate names file, cause and next step. With a shipped pattern
the cause is that a subject does not match a named pattern, and the page describing
that pattern is generated from the package. With a free expression the only thing left
to print is the expression, and the reader is left to work out what it wanted.

*Comparability.* Across fifty projects with shipped patterns it can be said which
checks apply where. Across fifty hand written expressions it cannot, and the promotion
of learned rules to the organisation level loses its point, because nothing is
comparable to anything.

*Cost and behaviour.* An expression can backtrack catastrophically on a long message,
and it can behave differently between runtime versions. In the shipped set that is
settled once. In fifty project rules it is unsettled fifty times, and a gate that hangs
on some inputs is not a deterministic gate.

*Provenance.* A shipped pattern moves with the plugin version, which every artifact
records through `rules_hash` and the plugin digest. An expression that was edited in a
project last month leaves the trail saying a rule was applied without saying which.

The honest cost is friction. A project with a house format has to get a pattern into
the package or build an external gate, which is slower than writing three lines in a
rule file. That is the intended trade. If it turns out that many projects need their
own formats, the finding is about the shipped set and not about this rule.

Rules that many projects want but no project should inherit unasked ship as examples
under `examples/rules/` in the distribution repository, Conventional Commits, the
`Xeno-Intent:` trailer and the signature rule among them. An example is adopted by
copying it into `given/org/` or `given/project/`, not by switching it on. There is no
`enabled` field and G-Rules needs no notion of one: a rule that is not in the tree does
not apply. The capability ships, the rule does not.

### Precedence

1. More specific beats less specific: `project` over `org` over `provider` over
   `builtin`.
2. At every level, `given` beats the matching `learned`.
3. A rule with `binding: true` under `given/builtin/`, `given/provider/` or
   `given/org/` cannot be overridden at all.

The order follows distance from the artifact. A supplier's rules travel with it
across every customer, an organisation's rules apply inside its own house, and the
repository belongs to that organisation. Where the two disagree, the owner of the
repository wins, which is also the contractual position. `binding` is what a supplier
keeps for the few statements it cannot give up regardless of customer, secrets never
reaching a repository being the obvious one.

`binding` is not allowed under `given/project/`, where it would have nothing left to
bind, and G-Rules rejects it there. Where two binding rules on different levels
collide, the gate is red and stays red. That is a matter between two organisations
and not something a gate may resolve.

G-Rules checks this deterministically. An unresolved collision is red.

In v1 the level of a learned rule is a claim, not an effect: every learned rule takes
effect only in this repository, because there is no shared repository. The claim is
what gets harvested later.

## 10. Learning

Mandatory at the end of every phase, even when the result is "no finding". The record
at intent level is the abandonment case: a merged intent writes none, because it already
has one per phase.

```yaml
intent: <qualified id>
created: <iso8601>
runner_version: <version>
plugin_version: <version>
phase: 02-design            # absent in the closing record at intent level
learnings:
  - category: <template|prompt|context-rule|project-convention>
    observation: <what was noticed>
    proposal: <concrete rule or template change>
    target: .xeno/config/rules/learned/project/<file>.yaml
```

A mapping rather than a bare list, for the reason given in section 8: a sequence has
nowhere to put the fields every process file carries.

Learning never changes behaviour directly. It produces a merge request against the
rule set. Only after merge does it take effect, and it then applies to subsequent
sessions.

`target` is the file that merge request would change. A rule the engine applies
names a path under `.xeno/config/rules/learned/project/`. Until that engine exists,
a learning whose category is `project-convention` may instead name the file a
project already reads before it acts, and the length of that file is what decides
whether a convention earns a line in it: it is sent with every request of every
session, so a line there costs more than a rule the engine loads when it applies.
Either way the route is the same, a record in a phase and a merge request somebody
reviews, and never an edit made where it was noticed.

Learning looks at its own phase only. Harvesting repetitions across intents is on the
1.1 list, because it finds nothing before the third intent and therefore delivers no
value until dogfooding is running anyway.

What is learned: templates, prompt versions, context selection rules, and project
specifics such as programming language and architecture rules. Gate thresholds, model
routing and estimation accuracy are measured in v1 but not learned.

The `abstract` field in `learned/provider/` and `learned/org/` is always written in
English, independent of the artifact language, because it is meant for later cross
project harvesting. The concrete rule below it stays in the project language.

Export across project boundaries does not happen in v1. It is an organisational act
requiring the consent of the parties involved, not a technical one.

## 11. Cost

One session per phase, one working directory per intent. `sessions` is a list
because a phase can be resumed after an interruption, not because several may run in
parallel. On phase completion the local session logs are evaluated and written to
`cost.yaml`.

```yaml
intent: <qualified id>
created: <iso8601>
runner_version: <version>
plugin_version: <version>
phase: 02-design
evidence: self-reported
tokens_in: <n>
tokens_out: <n>
tokens_cached: <n>          # where the harness reports it
cost_usd: <n>               # optional
sessions: [<id>]
```

`cost.yaml` is optional. A phase without it is complete, which is also why it sits
outside the `artifacts_hash`: a figure that arrives after the verdict must not
invalidate it.

That is one instance of a rule which holds throughout: **what is sealed is never
rewritten.** Anything learned about a phase after its verdict is either a separate
file that references it, or a comparison that is computed. It is never an edit to an
artifact whose hash carries the verdict. Drift already works this way: an artifact
authored under an earlier rule set is not touched, the difference between the hashes
it recorded and the ones in force is computed and reported. Stating the rule matters
because the convenient phrasing, marking something as outdated, would change the very
hashes the verdict rests on.

**Tokens, not money.** The token counts are what a session reports and what any of the
levers below can be judged against. Converting them into an amount depends on prices,
plans and negotiated terms, none of which belong in a repository, so `cost_usd` is
optional and the authoritative money view comes from the model access point in v2.
Input, output and the cached share are kept apart because output costs several times
input and a cached prefix a fraction of either; one combined number would hide the
effect of every lever.

`evidence: self-reported` is mandatory. The v1 numbers are indicative and must not
enter certification records without re-measurement. The field schema is cut so that
a gateway can fill it in v2 without migration.

**Requests carry the intent and the phase.** Where model access runs through a
gateway, the qualified intent id and the phase id travel with the request as
metadata. In v1 this costs nothing, since the CLI is the only caller, and it is what
turns a gateway's own figures from a sum per developer into something attributable to
a phase. It also places the same two identifiers in a third location, held by a
system nobody working a phase can write to. The metadata is asserted by the caller,
so it correlates and does not attest, and it carries identifiers only: no prompt
fragments, no file paths, no free text, and nothing that would put a customer name
into a central log with its own retention.

**What the project can turn.** Model and thinking budget are set per phase in
`project.yaml`, not once for the whole project: intake and verification rarely need
what design needs, and the model actually used is recorded in every artifact, so the
choice stays auditable rather than implicit. Enabled lenses are a cost decision too,
since each one enters the context of every phase it applies to. So is the MCP tool
surface, which is why it is fixed and small: a tool definition is paid for in every
request of every phase, whether or not it is called.

**One phase, one session.** Resumption exists for interruptions, and it is not free:
a resumed session carries its whole history back into the context. Where a phase has
been lying open, restarting it from the artifacts is usually cheaper than continuing
from the session, and the artifacts are complete by construction.

Retention and cost recording touch here, and the order matters. Pruning skips the
sessions of any phase that has no `cost.yaml` yet, and the retention period runs from
phase completion rather than from the age of the log file. Otherwise a phase left
open for longer than the retention period loses its figures quietly, which is the
worst of both: no numbers and no record that there were any.

## 12. Triggers, tools, environment

**Triggers.** The CLI, and nothing else in v1. `xeno phase start --intent PROJ-123
--phase 01-requirements` is the single entry point. A pull request never starts a
phase; the pull request is an output artifact. CI running the gates on it is not a
trigger in this sense either: it verifies artifacts that already exist, it does not
produce any.

Issue commands are deliberately absent. Something has to receive them, and every way
of doing that is a component to build and operate: a service taking webhooks, or a CI
job wired to comment events in each code host. Neither earns its keep while the
result is a phase invocation somebody can type. What remains from the tracker side is
write-back, described below, which runs from CI and needs nothing running in between.

**Two surfaces, two promises.** `xeno gate ...` never touches the network. Its input
is the state of the repository plus the commit range it was given, which is what
makes a verdict reproducible on any machine and what the governance rests on.
Everything else in the CLI may use the network: starting a phase reads the issue,
finishing one writes a comment. The distinction is not cosmetic, because a gate that
could reach a code host could also return different verdicts on different days.

**Interface.** The CLI drives the process. Nothing is operated through a UI, and v1
has no UI at all.

A read only dashboard follows in 1.1: which artifacts exist and how they connect,
which phases each intent has completed, where a stale phase blocks progress, and what
the work has cost. It is derived from the repository and holds no state of its own,
so it can neither drive the process nor disagree with it, and nothing in v1 depends
on it. Everything it would show can be read from the repository, which is what makes
it the one package that can wait without leaving a gap.

Cost will need care there. The numbers are self reported, and a dashboard makes
numbers look authoritative simply by putting them in a chart. `evidence:
self-reported` belongs wherever a cost figure is shown, not hidden in the file behind
it.

**Write-back.** After a gate run, CI writes the result to the issue the intent came
from: phase, verdict, and the findings with their decisions. It is the only
visibility for everyone who does not work from a command line, and it costs one call
from a job that already runs. In v1 the code host is GitHub; further hosts follow in
1.1, and the adapter contract stays small enough for them: resolve an intent, read an
issue, write a comment, resolve credentials. The endpoint is a configured value with a
default, since this host has a canonical address, and it stays configurable because
Enterprise Server does not use it. A second host is configuration plus an adapter and
never a change to the contract.

**Tools.** The agent is a project level choice, set in `project.yaml` and not
changed while the project runs. Both Claude Code and Codex are supported. The tool
and its version are recorded in every artifact, as a record rather than a variable.

Model, tool and version in every artifact are the raw material a project needs where
it has to keep a register of its ICT service providers. Xeno keeps no such register
and makes no claim about one. A project that maintains one links it from its own
README, and that is the whole extent of the connection.

**Full text.** Prompts, contexts and responses stay local in v1 under
`.xeno/local/`, excluded via `.gitignore`, together with the symbol index and anything
else derived rather than recorded. Retention is a project setting with a
default of 30 days, and the runner prunes what has expired on every run rather than
leaving a year of full text on someone's disk. Only `digest.md` is taken into the
repository on phase completion. In v2 storage moves to a gateway.

The digest is produced in two steps, and the split is the point. The agent writes the
summary text; the runner filters it against the effective secret filter, hashes it
and writes the file. The filtering is therefore deterministic and outside the model's
reach, so a model cannot route around it, not even by accident. The summary itself
stays marked as a derived artifact rather than a basis for decisions.

**What the gate runner reads.** Artifacts under `.xeno/`, the files a context scope
names, and the commit history of the change under review. Nothing else: the input is
the state of the repository plus an explicitly passed commit range, never a query to
the code host. Where a rule needs the range and none was passed, the gate is red
rather than quietly skipped.

**Execution environment.** The process runs entirely in the project environment,
meaning its CI, its tracker and its repository. There is no second process variant.
Two requirements for the runner follow: gate evaluation needs no network at all,
since no gate calls a model and no gate queries the code host, and the binary must be
delivered verifiably with a fixed version, a signed artifact and a checkable hash.
The agent needs the model endpoint, the gate path does not.

## 13. Distribution

Three artifacts, one distribution repository, one shared version number.

```
xeno/
  plugin.json               Agent Plugins 1.0.0
  skills/
    xeno-intake/SKILL.md
    xeno-requirements/SKILL.md
    xeno-design/SKILL.md
    xeno-implementation/SKILL.md
    xeno-verification/SKILL.md
    xeno-review/SKILL.md
    xeno-learning/SKILL.md
    xeno-lens-security/SKILL.md
    xeno-lens-privacy/SKILL.md
    xeno-lens-operations/SKILL.md
    xeno-lens-architecture/SKILL.md
  mcp.json                  declares the Xeno MCP server
  .claude-plugin/
    marketplace.json        second wrapper over the same content
  runner/                   gate runner sources
  templates/                full template set, shipped with the plugin
  rules/given/builtin/      shipped rule set
  examples/rules/           adopted by copying, never shipped into a project
  secrets.yaml              shipped secret filter
```

The distribution tree and the vendored tree are not the same set. `--vendor` copies
`plugin.json`, `skills/`, `mcp.json`, `templates/`, `rules/given/builtin/` and
`secrets.yaml`. `runner/` stays behind because the runner is delivered as a binary,
and `examples/` stays behind because an example that travels into every project is no
longer an example.

Six of the skills are the phases. `xeno-learning` is not a seventh phase: learning
happens at the end of each of the six and once more when an intent closes, so it is
cross cutting and carries its own skill for that reason.

**The phases are model invoked, and that is deliberate.** A skill is loaded because the
client judged its description to fit; a command is typed. For a lens that is the right
way round, since whether a security lens applies is a judgement about the change. For a
phase it is not: the person knows which phase they are in, and `xeno phase start` tells
them which one comes next, so what they want to say is "do the intake". MCP's prompts
primitive is the form that would carry it, and what it lacks is a second client. Claude
Code surfaces a prompt as a command, `/mcp__server__prompt`, discovered from the
connected server; Codex, as of October 2026, does not. A mechanism one client supports
is the `commands/` problem under another name, which is the thing skills exist here to
avoid, so the phases stay skills until both clients can take them. This paragraph is
what a change to that deletes.

| Artifact                   | Channel                                    | Recipient           |
|----------------------------|--------------------------------------------|---------------------|
| Plugin (skills, mcp.json)  | Agent Plugins 1.0.0, vendored from the repository | developer machine |
| Runner                     | OCI image in the host's container registry | CI                  |
| Runner                     | per platform binaries attached to the release, with their checksums | developer machine |

**The repository is public; distribution is not yet.** The two are separable and are
separated here. The repository is readable by anyone, which is what gives the host the
branch protection the sequence guarantee rests on, and the channels are still the
repository's own releases and registry and nothing else. Two consequences follow
rather than being decided separately. The marketplace wrapper stays and points at that
repository rather than at a public directory, since a client can take a marketplace
from any git repository it can reach. And the release carries checksums rather than
signatures, for the reason given under integrity below.

The runner reaches the developer machine as well as CI, because hooks and the local
run start it there. The installation route belongs in onboarding; what belongs here
is that the version is pinned in `project.yaml` and that `xeno init` refuses to run
against a different one rather than quietly using whatever is installed.

**Roles.** The phase skills already carry the classic roles: requirements is the
analyst, design is the architect, verification is QA. Modelling those again as roles
would duplicate the same thing and leave the agent unsure which skill applies.

What no phase owns are the cross cutting concerns, so those ship as lenses: security,
privacy, operations and architecture conformance. A lens is a skill a phase pulls in,
not a phase of its own, and not a role either. It declares which phases it applies
to, is enabled per project in `project.yaml`, and its output goes into the phase's
open questions, the assumption register, or the review checklist. A checklist entry
from a lens carries `source: lens` and no rule id, which is what keeps it apart from
the entries G-Policy counts: the gate checks that every entry rendered from a review
rule has been answered, and a lens cannot add to or subtract from that set. A lens
never writes an artifact of its own and never produces a gate verdict: a model based
judgement is not a deterministic check, and dressing one up as a gate would undo the
distinction this process rests on.

**Working practices are not skills.** Clean code, TDD, BDD and domain driven design
are expressed through the extension points that already exist, not through more
skills. Clean code is a rule set, its mechanical part already covered by linters
whose output is declared evidence and the rest by `review` rules. BDD is a template
variant for the requirements phase plus an evidence kind in verification. DDD is a
template variant for design plus project rules. TDD is a project setting and a rule
requiring a test in the same merge request, since the phases order artifacts and
evidence rather than the sequence of keystrokes.

The same applies to specialisations inside a phase. UI, backend and acceptance
testing all belong to verification and contribute to one artifact; they are the work
of the phase in different forms, not different perspectives on it. The coverage map
and the `kind` field on each evidence item already make visible which level a
criterion was checked at.

A skill per method or per test type would produce a matrix of thin skills, and an
agent choosing between twenty of them chooses badly. If practice shows a phase does
not run well without its own guidance, that is a demonstrated reason to add one.
Before that it is an assumption.

Lenses and any other specialist are skills rather than subagents, for two reasons, and
the second is the one that lasts. Subagents are client specific and outside Agent
Plugins 1.0.0, so building them that way would break the interchangeability of the two
agents. And v1 has nowhere to write one down: a skill is part of the vendored plugin and
hashed by G-Supply, so the trail says what was worked under, while a delegated call
leaves nothing behind, no field in the artifact, no line in `context.lock.yaml`, no model
and no brief. That would be work whose origin cannot be reconstructed, which principle 5
rules out. The first reason falls away the moment a project settles on one client; the
second does not, and it is what a later version has to solve before specialists may
become agents.

What this costs is the separate context: a skill works inside the context window of the
phase, so with four lenses active their text travels in every request instead of being
paid for once. That is acceptable here because an intent is one story and the phases are
short. Where a client does offer subagents, running a lens in a forked context is a
worthwhile optimisation and never a requirement, as long as nothing in the trail comes
to depend on it.

**MCP instead of free text, and never only MCP.** Every operation the server exposes
has a command that does the same thing, because the server is a thin shell around the
same code. The server is also not something anybody stands up: it is a subcommand of the
binary that is on the machine anyway, started by the client over standard input and
output, with no host, no port and no deployment.

That matters for the case where MCP is not available, which is not exotic: a client that
does not support it, a locked down machine, a different agent, or a person working
without one. A phase can then be carried out with commands alone. What is lost is
convenience, not evidence: the harness hook falls away, so G-Secret first applies at
phase finish rather than at every write, and the agent types commands instead of calling
operations. No gate, no artifact and no hash depends on MCP being there.

The server exposes exactly the operations that may be performed on the process: read
intent, record assumption, write artifact, fetch template, run gates locally, and query
the symbol index. Six, and the surface is a
budget rather than a list: every definition is sent with every request of every phase
whether it is called or not, so a seventh needs an argument of the kind the index has.
The agent then works with process operations rather than file paths, and the skill text
stays short.

A prompt is not a tool for this purpose, and the budget's reason does not reach one:
`prompts/list` carries a name, a title, a description and the arguments, and the
messages arrive at `prompts/get` when the prompt is invoked, so its text is fetched
rather than sent with every request. What a client puts in front of the model from the
listing is the client's own choice and is not fixed by the protocol.

**Vendoring instead of remote fetch.** `xeno init --vendor` places the plugin
pinned under `.xeno/plugin/`. It works without outbound connections, stays
inspectable because the package unit is a directory, and makes the plugin version
part of the information base. It therefore belongs in `context.lock.yaml`, not only
in an installation guide.

**Releases carry version numbers and nothing else.** No code names, and the CLI
banner stays plain:

```
xeno 0.1.0
```

**Semantic Versioning.** Everything versioned here follows SemVer 2.0.0. What counts
as a major step depends on the object, and the common thread is the same throughout:
major means a check that was green yesterday is red today.

| Object                             | Major when                                                                 |
|------------------------------------|----------------------------------------------------------------------------|
| Plugin and runner, shared number    | a gate gets stricter, a mandatory field appears, or an artifact format changes so that older artifacts can no longer go green |
| Template                            | a section id is dropped or renamed, or a section moves from optional to required |
| Artifact schema                     | a mandatory field appears or an existing field changes meaning             |

Strings bundles are not versioned separately. They follow their template and are
bound through `strings_hash`. Rules are not in the table either: their `version` is a
change counter, for the reason given in section 9.

**Version coupling.** Plugin and runner share a version number and are released
together. An agent following plugin 1.3 produces artifacts that runner 1.1 cannot
check. G-Supply therefore fails on any version difference, not only a major one: the
runner carries the digest of exactly one plugin, its own, so an upgrade is a single
act rather than a compatibility matrix. Update the runner, run `xeno init --vendor`
again, commit both together.

**Integrity.** The specification defines no signatures, no permission model, no
sandboxing and no secrets mechanism. Verification is therefore Xeno's own job, and it
needs an anchor that does not sit in the repository it is protecting.

The anchor is the runner binary. Plugin and runner are released together under one
version, so the runner carries the digest of its own plugin compiled in. G-Supply
recomputes the digest over `.xeno/plugin/` and compares. Nothing in the repository
states what the expected value is, which is the point: an expected hash stored beside
the thing it describes proves only that both were written by the same hand.

That also settles the downgrade. A vendored plugin from an earlier release has a
different digest than the one this runner expects, so it fails without any separate
version check.

This is not tamper protection and does not pretend to be. Whoever can rewrite the
vendored plugin can rewrite rules, artifacts and the CI wrapper as well. What G-Supply
gives is that a changed shipped set cannot pass unnoticed, and for that a digest the
checked side cannot influence is enough.

The release itself carries checksums and no signature while distribution stays
internal. A signature answers the question whether an artifact came from where it claims
to, and with access control on the registry it was pushed to that question is already
answered. The obvious way to sign, keyless attestation, would publish the identity of an
internal pipeline into a public transparency log to answer it a second time. When Xeno
is published, signing comes with publication and not before, and the first public
release is the first one that needs it.

## 14. Extension

Extension works on three levels, and they differ in what they cost in trust.

**Data, no code.** Templates, strings bundles, rules, secret patterns and lenses are
files. A project adds its own or overrides a shipped one by id, with the resolution
already defined. This covers most of what anyone wants to extend and costs nothing in
trust, because nothing is executed.

**Additional plugins.** Skills and MCP servers come from the plugin ecosystem. A
project can install a second plugin beside Xeno and both address the same agent.
Nothing in Xeno needs to allow for this.

**External gates.** A gate may be an external command, declared in `project.yaml`
with its path and hash, receiving and returning JSON, with the exit status deciding.
It runs only when the hash matches. Every finding it produces carries
`provenance: external` in `gate.yaml`, and G-Complete accepts it like any other gate.

The marking is the point. Without external gates the chain of trust is closed:
vendored plugin, hash, signed release. An external gate brings foreign code with
repository access into the pipeline, so the trail has to show which statement came
from Xeno and which did not. A project that does not enable them keeps the closed
chain.

**Not extensible in v1.** Predicate types are limited to what the runner implements:
an external gate covers the same ground and draws the more honest line, either Xeno
checks it with its own means or it is foreign code and marked as such. The phase list
is fixed, because phase ids appear in templates, skills, directory names and gate
assignments, and a phase nobody ships templates or skills for would leave the process
running empty at that step. Free phases are a v2 question, tied to whether a phase
without shipped templates and skills is useful at all.

## 15. Licensing and contributions

Xeno is open source under Apache 2.0. The choice is driven by the patent grant,
which matters for a tool that runs inside other organisations' pipelines.

**Results belong to the user.** No licence term applies to what is produced with the
tool. Artifacts, rules and code created in a project are the project's, without
conditions. The shipped templates and rule sets carry their own permissive notice,
separate from the code licence, so that using them triggers nothing downstream.

**Contributions are expected but not required.** Extensions and changes to the tool
itself belong back in the project: a maintained fork costs more over time than a
pull request, and every change fed back saves the next user the same detour. The
expectation is stated in the README and the procedure in `CONTRIBUTING.md`, using a
DCO sign-off rather than a CLA, which is easier to clear in corporate environments.

**What is not contributed back.** Rules, templates and artifacts from specific
projects stay with the project. Learned rules with `scope: provider` are candidates
for later harvesting between the organisations involved, as described in section 10.
That is a separate act and not an obligation towards this project. Both live in the
same repository, so the README must say plainly that they are two different things.

**Repository files.** `LICENSE` with the unmodified Apache 2.0 text, `NOTICE` with
the copyright line, an `SPDX-License-Identifier` header per source file, and a
separate notice under `templates/`.

**Who holds it.** Xeno is developed at and for javafreedom.org, which holds the
copyright and licenses it under Apache 2.0. The repository is public from v1 and the
distribution channels follow in 1.1, both of which the licence already permits without
relicensing.

Stating the holder plainly is worth more than leaving it to be inferred from a
`NOTICE` file: a tool that asks organisations to run it inside their pipelines should
say whose tool it is, and a reader who knows the origin can judge for themselves which
parts carry an opinion.

The licence makes the practical answer the same either way. The patent grant, the
absence of any claim on results and the DCO route above apply to every user
including javafreedom.org, and nothing in the process depends on who the copyright
holder is.

## 16. Deliberate limitations of v1

Sixteen entries, and they are three different kinds of statement. Read as one list they
suggest that something is wrong everywhere; sorted, they are three things worth knowing.
The numbering is continuous because other documents refer to it.

### What v1 does not do yet

These are scope, not weakness, and they appear in section 1 and in the outlook as well.

**1. No risk differentiation at gates.** Human on the loop applies uniformly.

**2. No cross project learning.** Findings stay in the project.

**3. No operations or maintenance.** The process ends at merge. Regulatory regimes aimed
at operational resilience are therefore out of reach by construction, not by omission:
they ask about running systems, and this process has nothing to say about them until a
phase for operations exists.

**4. No aggregation beyond the project.** There is no cross project view.

**5. One repository per intent.** v1 targets monorepos. A change spanning several
repositories has no representation in the process, which is a restriction rather than a
gap: it buys the assurance that every hash in the trail resolves locally.

### What a green gate does not mean

This is the group an auditor needs. The gates are strict about structure and
completeness and deliberately quiet about substance.

**6. G-Test checks completeness, not selection, and not existence.** Without impact
analysis the test scope is declared rather than derived, and no gate checks whether the
right tests were chosen. The mapping also names tests as free text, because an identifier
looks different in every ecosystem and nothing resolves it: it proves that somebody made
the connection, not that the named test exists or ever ran. Whether the run succeeded is
likewise a declared `result` and not a parsed report. A project that needs more binds an
external gate.

**7. Scan findings are declared, not judged.** A scan is evidence like a build or a test
result, bound by content hash and by the declared result. No gate evaluates its findings
against a threshold, so the trail shows that a scan ran and not what it was required to
show.

**8. A missing `result` is ambiguous outside tests and builds.** The field is written
where the producer reports against a threshold and absent where it does not, so its
absence means either that no threshold was configured or that somebody forgot to record
the outcome. On `test-report` and `build-log` the field is required and the omission
shows; everywhere else the two cases look the same. A project that needs the distinction
can state it as a rule rather than expecting the process to assume a threshold.

**9. Rule drift is marked, not blocked.** An artifact authored under an earlier rule set
stays valid; the difference appears in `drift` and as a line in the P5 checklist. Whether
the earlier assessment still holds is a human answer, and nothing forces the phase to be
redone.

**10. Overrides are visible, not enforced.** Nothing makes the follow-up happen. The
trail shows which finding was waved through and by whom, and the rest is a management
question.

### What the trail rests on

Every record has an anchor outside itself, and these are the anchors.

**11. The audit trail is only as trustworthy as the repository.** Locally executed tools
cannot be protected against tampering. G-Supply narrows this for the shipped set, since a
changed plugin cannot pass unnoticed, and the remaining anchor is the runner binary
itself: whoever replaces it wins everything. Pinning its version and pulling it by
checksum from the release is the whole of the countermeasure.

**12. Identity in the trail rests on repository access.** `approval.by`, `override.by`
and `confirmed_by` are written into files and committed. A code host authenticates the
push and the merge, which is worth something, but that evidence lives in the host's audit
log rather than in the repository, and an unsigned commit carries an author line that is
free text. Projects that need more adopt the signature example rule; Xeno does not impose
it.

**13. Two things do not survive a squash merge.** Squashing is supported and not a
special case: artifacts, verdicts, decisions and evidence are bound by content hash and
stay checkable in the merged branch. What the squash replaces is commits, so the commit
bound things break with it. Commit predicates proved what stood in review and not what is
in the history afterwards, because the commits a rule checked no longer exist. And the
link from a merge commit back to its intent is the key in the commit message, which is
text rather than a hash: where nothing enforces the message format, and on an edition
that cannot enforce it before the gate, a merge commit without the key loses that link
for good.

**14. Rule sets above the project are copied, not distributed.** The `provider` and `org`
levels exist as files in every repository, because the repository is the source of truth
and v1 has no shared store. No reconciliation takes place. Divergence is detectable
afterwards through `rules_hash` and is not prevented.

**15. Token figures are self reported.** Without a gateway there is no trace per call.
Attribution rests on the session rule, and the numbers are an indication rather than an
account.

**16. External evidence can expire.** Where CI retention lapses, the hash and the
declaration remain but the content is gone. The trail then shows what was checked, not
what it contained.

## 17. Outlook

What follows is the direction, not a decision. For v2 the decisions now exist: the
platform a phase executes on is settled in [[orchestrator-evaluation]], and
what that changes against this document is recorded in [[v2-delta]],
including the places where this document is superseded. Nothing in the sections above
has been rewritten for it, because they describe v1.

**1.1, breadth on a proven core.** Further trackers, GitHub, Jira, Bitbucket and
Azure DevOps, with the CI wrappers that go with them. Intents spanning several
repositories, where the qualified intent id starts to earn what it already costs.
Issue commands as a second trigger, once there is a reason to operate a component
that receives them. The read only dashboard. The German documentation translation
with the hash binding that keeps it from rotting. All five are held back for the same
reason: they multiply surface without answering
whether the process is proportionate, and dogfooding answers that first.

**v2, the parts that need infrastructure.**

- Gateway for model access, making cost, routing and full text storage central and
  auditable
- Direct calls to a model endpoint for the phases that do not need a coding harness,
  see below
- Aggregation beyond the project, harvesting rules with `scope: provider`
- A distribution path for the `provider` and `org` rule levels, so that they stop
  being per repository copies
- Risk dependent gates
- Live agent status, which the repository cannot answer because a working agent has
  committed nothing yet
- The dashboard UI separated from the runner binary, which 1.1 ships bundled
- Context tool in the `tools` block, then also as input for test selection in
  G-Test
- Phase P6 for operations and maintenance, and with it the connection to
  requirements about running systems rather than changes to them
- A predicate over scan findings on a normalised format such as SARIF, so that a
  threshold can be declared in the repository instead of living in a pipeline
  configuration
- A freely extensible phase list, and custom predicate types if external gates turn
  out not to cover the need

**On calling endpoints directly.** v1 drives a harness, which is why cost is self
reported, caching is somebody else's, and the model in use is a configuration the
process records rather than controls. Calling an endpoint itself would close all
three at once: measured token counts instead of parsed session logs, explicit cache
control over a prefix the process assembles anyway, batch pricing for anything with
nobody waiting, and a model choice the runner enforces rather than requests.

It is worth doing for the phases where there is nothing to drive. Intake, the digest,
the rendering of a checklist and most of verification are text in, text out, and a
harness adds a tool loop nobody needs there. It is not worth doing for design and
implementation, where the harness is the product: editing files, running tests and
iterating is exactly what it exists for, and reimplementing that inside Xeno would
turn a process tool into a coding agent.

Two things have to hold if it happens. The gate path stays network free, so endpoint
access lives on the phase side and nowhere near a verdict. And a phase driven by an
endpoint produces the same artifacts under the same schema as one driven by a
harness, recorded through the existing `tool` and `model` fields, or the audit trail
splits into two kinds and the comparison between them is lost.

## Appendix A. The full `project.yaml`

Every block below is defined in the section named beside it. The file is collected
here because it is the one artifact a reader assembles from seven places otherwise.

```yaml
# .xeno/config/project.yaml

runner_version: 0.1.0          # section 13, xeno init refuses a mismatch

language:                      # section 5
  artifacts: de                # ietf tag, the process layer stays English

agent:                         # section 12
  tool: claude-code            # claude-code | codex, fixed for the project
  model:                       # per phase, section 11
    default: <model identifier>
    "00-intake": <cheaper model identifier>
  thinking:                    # per phase, omitted means the harness default
    default: <off|low|high>

tracker:                       # section 12
  adapter: github
  base_url: https://api.github.com   # a default, overridden for Enterprise Server
  auth: { scheme: token, secret_env: XENO_TRACKER_TOKEN }

enforcement:                   # section 7, compared against the host by CI
  required_pipeline: true
  allow_bypass: false
  merge_method: no-squash
  approvals:
    required: 1
    not_by_author: true
    waived: <reason and date, where the host cannot express the requirement>

lenses:                        # section 13
  enabled: [security, privacy]

external_gates:                # section 14
  - id: house-linter
    path: tools/house-linter
    sha256: <hash, checked before every run>
    phases: [03-implementation]

evidence:                      # section 4
  source: ci                   # ci | local, ci is the default

retention:                     # section 12
  local_days: 30               # from phase completion, see section 11

templates:                     # section 4
  overrides_dir: .xeno/config/templates

index:                         # section 5, context economy
  path: <path to the symbol index the project produces>
  max_age_hours: 24            # older than this is treated as absent
```

Credentials never appear in this file. `secret_env` names the variable, the
environment holds the value.

What each field means, whether it has to be there, and what applies when it is not:

| Field | Meaning | Absent |
|---|---|---|
| `runner_version` | the runner this project is pinned to, exactly, not a lower bound | `xeno init` and every command refuse to run |
| `language.artifacts` | the language artifacts are written in, as an IETF tag; the process layer stays English regardless | `en` |
| `agent.tool` | which harness the project uses, recorded in every artifact it produces | no default; the agent layer has nothing to configure and the field in the artifact comes from the harness itself |
| `agent.model.default` | the model a phase uses unless overridden | the harness decides, and the artifact records what was used |
| `agent.model.<phase>` | an override for one phase, keyed by phase id | the default applies |
| `agent.thinking.default` | reasoning effort, since those tokens are charged as output | the harness default applies |
| `tracker` | the whole block is optional: Xeno runs without a tracker, the trigger is the CLI and write-back is comfort | no issue is read and no comment is written; a phase is started with the issue content supplied by hand |
| `tracker.adapter` | which host adapter to use | the block is incomplete and `xeno phase start` refuses rather than guessing |
| `tracker.base_url` | the host's API, which has a default and is a parameter all the same, since a self managed deployment has no canonical address | same |
| `tracker.auth` | the scheme and the environment variable holding the token, never the token | same |
| `enforcement` | what the project requires of its host, compared against it by `xeno enforcement check` | the check fails. Skipping would make it useless in exactly the projects that never wrote one, and a project that requires nothing can say so explicitly |
| `enforcement.required_pipeline` | whether a merge needs a green pipeline; this is the setting a binding verdict rests on | treated as required and reported as unmet |
| `enforcement.allow_bypass` | whether administrators may merge past the rules | treated as `false` and reported as unmet where the host allows it |
| `enforcement.merge_method` | `no-squash` where a commit predicate is active, since a squash replaces the commits it judged | not checked |
| `enforcement.approvals` | how many approvals are required and whether the author may give one | treated as unmet, unless `waived` says why the host cannot express it |
| `lenses.enabled` | which lenses run, each of which enters the context of every phase it applies to | none run |
| `external_gates` | commands whose verdict counts alongside the shipped gates, pinned by hash | none run |
| `retention.local_days` | how long session logs are kept, counted from phase completion | 30 |
| `templates.overrides_dir` | where project template overrides live | the shipped templates apply unchanged |
| `index.path` | where the symbol index the project produces is found | no index is read and the phase runs without one |
| `index.max_age_hours` | beyond this age the index is treated as absent, because a stale index is worse than none | 24 |

Three of these are worth reading twice. `tracker` is optional because nothing in the
trail depends on it. `enforcement` is not, because a missing statement about the host is
indistinguishable from a host nobody checked. And an absent `index.path` degrades
rather than fails, which is the pattern for everything that makes a phase cheaper rather
than more correct.

## Appendix B. Hashes and identifiers

Every value here is recomputed by whoever verifies the trail, so they are defined to the
byte rather than described.

**`artifacts_hash`.** Over a phase: every file lying directly in the phase directory,
without descending into subdirectories, except `gate.yaml` and `cost.yaml`. Over an
intent: every file lying directly in the intent directory, except `gate.yaml`.

Every file so covered is one Xeno wrote itself and is therefore text, which is what
allows the normalisation below to apply without any detection. That is not an
assumption but a checked property: G-Schema reports a file it does not recognise in the
phase directory, so anything else that lands there is a finding before it reaches a
hash.

1. Normalise the content: every CRLF becomes LF. Nothing else changes, and no trailing
   newline is added or removed.
2. Form one line per file: the sha256 of the normalised content as lowercase hex, two
   spaces, then the path relative to the repository root with `/` as the separator.
3. Sort the lines by path in byte order and terminate each with `\n`, the last one
   included.
4. The value is the sha256 over that stream, as lowercase hex.

The line format is the one `sha256sum` prints, so the stream can be reproduced with
standard tools; `LC_ALL=C sort` gives the required order. Nothing else enters: no
absolute paths, no file modes, no timestamps, no ownership. An empty directory
contributes no line.

Two consequences follow from it. The path is relative to the repository root, so a hash
carries the intent key and the phase inside it, and moving or renaming an intent
directory changes every verdict in it. And `evidence/` lies outside, which is why an
evidence item is bound through the `sha256` in its declaration rather than through this
value.

**`context_hash`.** The sha256 of the normalised content of `context.lock.yaml`, as
lowercase hex, normalised as in step 1 above. One file, so no line format and no path
enter: this value says what a phase was produced from, and the path it was produced
from is fixed by the phase directory that carries both files.

**`strings_hash`.** The same computation over the strings bundle the phase rendered
from, the file the `template` field resolves to. Where a project overrides a bundle,
the value covers the bundle actually used and not the one it replaced, which is what
makes the pair of fields say where the words came from.

**`secrets_hash` and `rules_hash`** are defined by the package that first writes them.
Each covers an effective set rather than one file, and how a set of several files
reduces to one value is a decision that belongs with the code that assembles the set
rather than ahead of it.

**The placeholder.** A hash field carries sixty four lowercase hex characters or the
value `by-hand`, and nothing else. `by-hand` says that no writer stood behind the
value, which is the honest state in three cases: where the field has no writer yet;
where the artifact declares `tool: manual`, because then nothing produced the artifact
either; and in `strings_hash` where the bundle version the artifact names is not the
one the repository carries, because a bundle that is gone cannot be hashed by anybody.
Elsewhere it is wrong rather than honest: a writer exists and the value was skipped.

**Finding id.** A sha256 over four fields, each terminated with `\n`: gate id, rule id,
path relative to the repository root, cause. The rule id is empty for the gates that
report on structure rather than on a rule, and its terminator is written all the same.
The id is the first six hex characters, lowercase, with `F-` in front, for example
`F-7a3c91`.

Six characters are a deliberate trade against readability in a file where the id stands
beside every decision. A collision is therefore possible in principle. Where two
findings in one `gate.yaml` produce the same id, the runner fails rather than letting
one decision cover both.

## Appendix C. About the name

Xeno refers to a recording technique in which two tracks that were never played
together, and never at the same tempo, are laid over one another afterwards. What
comes out sounds coherent even though it never happened at once. The coherence is
not in the recording, it is in the assignment.
