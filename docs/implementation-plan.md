---
id: implementation-plan
title: Xeno, Implementation Plan v1
revision: 9
status: draft, not ratified
date: 2026-09-20
location: docs/implementation-plan.md
---

# Xeno, Implementation Plan v1

Scope of this document: what gets built, in what order, and what "done" means for
each part. The process itself is defined in [[process-definition]] and is not
repeated here.

Drafts written during planning to think with do not travel into the repository. The
README, the template set and the onboarding texts are written in their own work
packages, against what the tool can actually do at that point; a draft copied in is
already out of date on its first line. The process definition and this plan are the
exception: they are results rather than scaffolding, and they live under `docs/`, at the
root of the repository rather than inside the directory the tool keeps its own state in.

**The documents under `docs/`.** There are four, and the count is rising, so what
each one decides is worth stating rather than inferring.

| Document | Decides | Does not decide |
|---|---|---|
| [[process-definition]] | what the process is, in v1 | how it gets built, and anything about v2 |
| [[implementation-plan]] | how v1 gets built, in what order, and what is deliberately deferred | the process itself |
| [[orchestrator-evaluation]] | which agent platform v2 executes a phase on, against which criteria, and what was rejected | any change to v1 |
| [[v2-delta]] | what v2 changes against the process definition, and what follows for v1 | the v1 build order |

The same rule applies to all four: they are results, they live in the repository, and a
statement made in one of them is not repeated in another. Where this plan says a thing
is decided for v2, the decision and its reasoning are in [[v2-delta]]. A fifth is
foreseeable, the conformance mapping, which currently sits inside [[v2-delta]] as a
section.

Cross references between these documents use the `id` from the frontmatter and never a
file name, so that renaming a file cannot silently break a link. The site resolves them
through a plugin that understands `[[id]]` and `[[id#anchor]]`, and it has to be one that
fails the build where a reference does not resolve: some only log a warning and render the
reference as plain text, which checks nothing. A heading that is referenced from outside
its document carries an explicit anchor through `attr_list`, because a generated slug
changes the moment a section is renumbered or reworded.

A revision counts up with every change of substance merged into `main`, and not for
typos. Once references point at an id, a reader wants to know whether the target has
changed since, and a number that moves for every comma answers nothing.

Work package numbers run in execution order.

## 1. Scope of v1

In scope: the six phase process end to end, deterministic gates in local and CI
form, template rendering with the shipped template set and both strings bundles,
English and German, the rule engine with checked and review rules, the assumption
register, learning records, indicative cost recording, a symbol index for context
economy, the agent layer for Claude Code and Codex with hook wiring, evidence
handling, the GitHub adapter with its generated wrapper, a documentation site in
English, a token economy with a measured baseline, and a conformance mapping against
ISO/IEC 42001 as the closing package.

**The target host is GitHub.** That decides more than a name in an adapter. There is a
canonical endpoint, so the base URL has a default; the adapter takes it as a parameter
all the same, because GitHub Enterprise Server does not use it and an adapter that
hard codes an address has a second host's worth of work hidden in it.

What the host offers depends on the plan a repository is on, which is the awkward part
and the one to state plainly. Required reviewers, protected branches and rulesets all
exist on GitHub and none of them is available for a private repository on the free
plan, where this project sits. So the four eyes requirement has no enforcement point
here, and neither does the protected branch that the binding verdict is supposed to rest
on. That is recorded as A27 rather than assumed away, and it is the same shape of
problem a self managed deployment on its free edition would have had, arriving through
the tier instead of the edition.

**The repository is public in v1; the distribution channels stay internal until 1.1.**
In v1 the channels are the host's releases and its container registry, what they carry
is checksummed and not signed, and the plugin is not registered in any marketplace.
What moved forward is the visibility of the repository, because the host grants branch
protection to a public repository and to no other, and that setting is what the
sequence guarantee rests on. The three obligations section 8 lists were met before the
visibility changed rather than after, which is the only order in which two of them can
be met at all.

**Licence, copyright and cost** are settled and recorded rather than open. Xeno is
developed at and for javafreedom.org, which holds the copyright and licenses under
Apache 2.0, with the repository public from v1 and the distribution channels following
in 1.1. It runs on infrastructure that is already paid for,
so the only expense that scales with the work is model usage, which is what WP13 makes
visible.

Out of scope, by decision: gateway for model access, cross project aggregation, risk
dependent gates, live agent status, a semantic call and type graph, operations and
maintenance phase.

Out of scope until 1.1, and for one reason: every item below adds surface without
telling us whether the process is proportionate, and dogfooding has to answer that
first. Further trackers and their CI wrappers. Intents spanning several repositories,
so v1 targets monorepos. Issue commands as a trigger, which would need a component to
receive them. The read only dashboard, WP18, which reads what the repository already
holds and which nothing else depends on. The German documentation translation and the
hash binding that keeps it current. Public release and everything that hangs on it,
signing included.

Artifact language is not on that list. Both strings bundles ship in v1, because the
separation of structure and text is what makes a second language cheap later, and a
mechanism proved against a test bundle is not proved.

## 2. Work packages

### WP0 Repository bootstrap

`LICENSE` (unmodified Apache 2.0), `NOTICE` with the copyright line, `README.md`,
`CONTRIBUTING.md` with the DCO procedure, a separate permissive notice under
`templates/`, `SPDX-License-Identifier` headers, and CI for the repository itself.

Also `SECURITY.md` with a reporting route, an SBOM produced at release, and a
`.gitattributes` that fixes line ending handling.

**Releases carry checksums and no signature.** While distribution is internal and the
registry has access control, the question a signature answers is already answered, and
keyless attestation would publish the identity of an internal pipeline into a public
transparency log to answer it twice. Signing arrives with publication, in the same
release, and the first public artifact is the first one that needs it.

The README states plainly that Xeno collects no telemetry, rather than leaving it to
be inferred from its absence.

**Release automation.** Xeno's own repository uses Conventional Commits, and the
pipeline derives the version, the tag, the release notes and the release from the
commit history. The shared version number of plugin and runner is set by that
tooling rather than by hand, which removes the one mistake G-Supply exists to catch
in the first place.

The tool is semantic-release, configured by `.releaserc.json` in the repository and
brought in as a CI job rather than as a dependency of this project: an action on GitHub,
the common-ci-tasks job on GitLab. A Go repository whose only dependency is one vendored
YAML library does not grow a Node dependency tree to cut a tag, and the configuration
moving between the two hosts unchanged is what makes the job replaceable rather than the
arrangement. A host's own changelog feature would cover the notes alone; it would not
cover the version, the tag and the write-back, and one tool doing all four in the same
way on every host is worth more here than the one call it saves.

**The repository's own pipeline is a deliverable that grows**, not a file somebody
adds when it is needed. It exists on day one because the protected branch needs
something it can require, and jobs arrive with the packages that produce them:

| From | Jobs |
|---|---|
| Day one, this package | format check, build, unit tests |
| Release, this package | tag triggered release, SBOM, checksums, binaries and image into the host's registries |
| WP17 | platform matrix and golden files, where line endings and path separators are settled |
| WP16 | documentation build and the Pages deployment on the default branch |
| M0 | the gate job itself, from which point Xeno checks its own repository and the hand held record in section 4 stops |
| WP10 | `xeno enforcement check` as the first job, plus the scheduled run |
| WP12 | write-back of the result to the merge request |

**Two things are called the workflow and they have nothing to do with each other.**
Xeno's own pipeline builds, tests, publishes and ships. The wrapper WP10 generates into
somebody else's repository does nothing but call `xeno gate run`. Same shape, opposite
purpose, and mixing them up puts build steps into a template or a gate call into this
pipeline.

**Pages has to be enabled for the repository.** It is an administrator's act and not on
by default. Same class of dependency as the required pipeline: outside the repository,
cheap to arrange early, annoying to discover in the week the documentation is due.

**Runners and egress.** A hosted runner has a route out, so the requirement below costs
this project nothing today. It costs any project adopting Xeno on its own
infrastructure, and it costs this one the day it moves to a self hosted runner, so the
reasoning stays rather than being deleted because the current host makes it cheap.

Outbound access is not a given on a self hosted runner. Every step has to work from what
that machine already holds, which rules out fetching toolchains from the public internet
at build time and makes the dependency mirror part of the bootstrap rather than an
afterthought.

**One issue label per work package**, `wp0` to `wp20`, created once at setup from the
list in section 2. Project labels, not group labels: the package numbers are this
project's construction plan and mean nothing in any repository that merely uses Xeno.
It is a one time act with no script behind it, because a tool that parses the plan to
automate something that happens once is more apparatus than it returns.

**Squash merges and the trailer.** Xeno's own repository references its issue in every
commit and carries the `Xeno-Intent:` trailer, adopted from `examples/rules/` once M0
exists. Where a merge request is squashed, the message that survives is built from its
title and description, so the trailer belongs there as well and not only in the
individual commits. Otherwise the rule proves what stood in review while the history
afterwards says something else, which is limitation 13 of the process definition
happening to the project that wrote it.

This governs how the tool is built. It says nothing about how projects using Xeno
commit, which is a rule they adopt or do not, see WP4.

Done when a clone builds, when the licensing position is unambiguous to a reader, and
when a merge to the default branch produces a versioned release without anyone
editing a version string.

### WP1 Artifact schema and core gates

The four field sets, common, phase, session produced and rendered, for Markdown
frontmatter and the equivalent top level keys in YAML artifacts, plus the `gate.yaml`
schema with decisions on findings. Gates G-Schema and G-Trace.

**Entries in `open-questions` and `decisions` carry stable keys, not just the sections.**
A question raised in P1 is answered in a later phase, whose `decisions` entry points back
at it with `resolves`, and a question without a key cannot be pointed at. Section ids do
not reach far enough here, and retrofitting item keys later would touch every artifact.

Decisions sit on findings and the phase status is derived from them; finding ids are
stable across runs and defined in Appendix B. Both rules are the process definition's,
section 5, and are not restated here. What this package owes them is the behaviour they
imply: a decision survives a re-run of an unchanged finding, does not survive a changed
one, and is never carried forward for a finding from an external gate, whose cause comes
from foreign code.

**The first fixtures are a deliverable of this package, written before the code.**
They come from the process definition rather than from the implementation: the example
in Appendix B, a well formed artifact of each kind, and one malformed variant per
required field. Written afterwards they would only prove that the code agrees with
itself, which is the closed circle this whole arrangement exists to avoid.

**A gate the runner does not implement yet is a state, not a silent pass.** The gate
list is a declaration and the runner grows into it: at M0 six of the fourteen do not
exist, G-Questions among them, since it takes effect from P5 and M0 proves one phase. `result: not-implemented` keeps a green verdict honest about what was actually
checked, and it keeps a gate that drops out later from disappearing unnoticed.

**What is sealed is never rewritten**, and this package is where that becomes code
rather than a sentence. Once `phase finish` has computed `artifacts_hash`, no command
writes inside it again; everything learned later goes into a file outside it or is
computed. The one sanctioned writer outside a phase's own directory is the evidence
attachment in WP6, and it is limited to two files in the immediate predecessor.

Done when a malformed or orphaned artifact fails and a well formed one passes, both
from the same binary, when a status written anywhere other than `gate.yaml` has
nowhere to go, when a decision survives a re-run of an unchanged finding and does
not survive a changed one, and when every command other than `phase finish` leaves
the `artifacts_hash` of every finished phase unchanged, which a test walking all
commands against a finished fixture checks.

The canonical intent id is qualified from the start: tracker host, project or
repository, and key. Retrofitting this later would touch every artifact path and
every stored reference, so it is cheap now and expensive at any point after.

**Hashes run over normalised content.** Line endings and path separators otherwise
make the same file hash differently on a Windows checkout and in Linux CI, which
would turn every gate that compares hashes into a coin toss. Normalisation is fixed
here, backed by `.gitattributes` in the package and a test that asserts it across
platforms.

**Schema versions are read, never rewritten.** When the artifact schema or a template
takes a major version, existing artifacts stay as they are: rewriting them would
destroy the very trail they exist for. The runner keeps reading older schema
versions, and only new phases are written in the new format. This has to be settled
before the first release, not after.

Settled in revision 9 of the process definition: `schema_version` in the common field
set, written by the runner, with an absent value read as the schema that predates the
field rather than failed.

### WP2 Template engine

`template.yaml` structure, strings bundles, the renderer that produces `output.md`
with section ids as anchors, and resolution between the shipped set and project
overrides. `template_source` written into `context.lock.yaml`.

Rendering runs on every section write rather than at the end, and the anchor is an HTML
comment; both are defined under Rendering in the process definition.

**One section is not rendered from the template alone.** The `review-checklist` in P5 is
built from the effective rule set, so the renderer needs the rule engine from WP4, which
lands four steps later. Until then P5 renders the section present and empty, which is
enough for M0, where one phase is proved rather than all six.

Done when the same artifact re-renders in a second language without changing what
it asserts, and when a missing bundle fails rather than falling back.

### WP3 Shipped template set

The six phase templates with their section ids, order and required fields, the
strings bundles in English and German, and the context profile format used by P0.

**The section count is the proportionality lever**, not the gates, which cost nothing.
Six phases with six required sections each are thirty-six fields to fill for a one line
fix. The budget is therefore three to four phase specific required sections per
template, and adding a fifth means arguing another one away.

Two sections are common to all six and optional everywhere except where noted:
`open-questions` and `decisions`. An open question is not an assumption: an assumption
is something the work proceeded on and G-Assumptions blocks while it is unconfirmed, an
open question blocks nothing and is a note to the next phase. Optional means an absent
section is fine; a section present as an empty heading is not, and G-Schema says so,
because empty headings accumulate and are worse than silence.

**A question entry carries its options.** Two to four, each with its consequence, plus
the agent's recommendation and a reason, plus a free entry that is an option like the
others rather than a field beside them. G-Schema rejects a question that has neither
options nor an explicit statement that none were found.

Entries in both sections carry stable keys, which is a schema property and is defined in
WP1 rather than here; this package only decides which sections exist.

| Phase | Required sections |
|---|---|
| P0 Intake | `problem`, `scope`, `context-rationale` |
| P1 Requirements | `acceptance-criteria`, `non-goals`, `constraints` |
| P2 Design | `decisions`, `alternatives`, `impact` |
| P3 Implementation | `changes`, `deviations` |
| P4 Verification | `test-mapping`, `results`, `gaps` |
| P5 Review | `review-checklist`, `release-notes`, `residual-risk` |

Some of these are forced rather than chosen. `acceptance-criteria` exists because G-Test
maps onto it, `test-mapping` because G-Test checks that mapping for completeness,
`review-checklist` because G-Policy checks that every review rule was answered.

The rest carry a decision worth knowing. P1 states what must hold once, written testably,
rather than separating requirements from acceptance criteria and having somebody phrase
the same thing twice. P2 promotes the common `decisions` to required, since design is
where decisions are the substance rather than a by-product, which is the case the
versioning table already knows as optional becoming required. P3 has two rather than
three because its substance is code; `deviations` is the one that matters, because a
design departed from without a record is what makes every later phase rest on a document
that is no longer true. P4 keeps `gaps` so that an acceptance criterion verified by hand
has somewhere to say so instead of being missing from the mapping.

What deliberately has no section: assumptions, which live in `assumptions.yaml`;
evidence, which is a frontmatter declaration; and a summary, which is what `digest.md`
is for.

The P5 template carries the `review-checklist` section, rendered from the effective
rule set with one entry per review rule, keyed by rule id, each with `result` and a
note required for anything other than `met`. Entries from a lens sit in the same
section with `source: lens` and no rule id, outside what G-Policy counts.

Drift lines from P0 to P4 render into the same section, since they ask a person the
same kind of question. Drift on the P5 artifact itself cannot: it is found by the gate
run that follows the rendering, and it stays in `gate.yaml` where the reviewer meets
it with the findings.

The content is written during development against what the tool can do at that point,
not carried in from a planning draft. This is the work package that decides whether
the process feels proportionate, so it is where the effort belongs.

Done when every phase renders a complete `output.md` from the shipped set alone, when
both bundles exist for every template, and when a project can override a single
template without forking the set.

### WP4 Rule engine and external gates

Rule file format, the four levels `builtin`, `provider`, `org` and `project`, the
`checked` and `review` split, the predicate types needed by the shipped rules, gates
G-Rules and G-Policy, and the minimal shipped rule set under `given/builtin/`.

Precedence is spelled out in the process definition and enforced here: more specific
beats less specific, given beats learned at the same level, `binding` above both,
`binding` rejected under `given/project/`, and `learned/builtin/` rejected outright.
The shipped set lives with the plugin, so it is covered by the plugin hash and cannot
be edited in place in a project.

Documentation consistency is a candidate for the shipped set rather than a topic of
its own: a `checked` rule expresses the derivable and mechanical part, G-Freshness
already provides the hash binding, and `review` rules are the honest admission that
the rest needs a person. Whether these ship enabled or as examples is a call for this
work package itself.

**Predicate types are named and shipped.** Rules parameterise a type, they never
carry a regular expression. Besides the section predicates the shipped set needs,
this work package delivers `commit-message`, `commit-trailer`, `commit-signature` and
`approver-not-author`, which read the commit history of the change under review, plus
the named patterns they use. Two patterns ship: Conventional Commits, and the same form
carrying an issue reference in the subject.

**`xeno check commit-message` evaluates one message instead of a range.** It is the
same pattern and the same implementation, reachable from a git hook, and it exists so
that the expression is never copied into a hook script. Its exit codes follow the same
staircase as everything else: 0 where the message matches, 1 where it does not, 2 where
it could not be evaluated, for instance because the named pattern does not exist.

Three copies of a pattern produce a push that rejects what the gate accepts, which costs
more trust than the check is worth. Where a server side hook runs on a machine without
the binary, a copy is unavoidable and the documentation says it has to be kept in step.

`approver-not-author` also reads `gate.yaml`, because it compares the `by` of a
decision against the authors of the range. It is the one predicate that makes a
separation of duties checkable rather than habitual, and it is cheap precisely because
both halves are already recorded. It evaluates on the run after the decision was
written, which the decision's own commit triggers.

G-Policy also checks the P5 checklist: every entry rendered from a review rule
carries a result, and anything other than `met` carries a note. That is the whole
extent of what a deterministic gate can say about a review rule, and without it a
review rule is a suggestion.

The range comes from `--base` and `--head`, resolved locally as the merge base against
the default branch and `HEAD`, and in CI from what the wrapper passes. The runner never
works it out for itself, because a guessed range means different verdicts locally and in
CI from the same repository state. A rule that needs a range it did not get is red.

**Two git hook templates ship under `examples/hooks/`**, with their limits written on
them. A `prepare-commit-msg` builds the message from the branch name, which is comfort
and bypassable and enough for producing one. A `pre-receive` rejects a push, which is
enforcement, and on a host that offers no push rule it is the only kind there is.
Installation is `core.hooksPath` pointing at a versioned directory, no dependency and no
extra tool; projects that already run something like husky or lefthook know what to do
with a script. `xeno init` does not touch a developer's git configuration.

**Examples rather than switches.** Conventional Commits, the `Xeno-Intent:` trailer
and the signature rule ship under `examples/rules/` and are adopted by copying into
`given/org/` or `given/project/`. No `enabled` field, no default-off state, nothing
for G-Rules to evaluate: a rule that is not in the tree does not apply. This keeps a working
practice out of the package while the capability behind it is available to anyone who
wants it.

External gates belong here too: declared in `project.yaml` with path and hash, run
only on a hash match, JSON in and out, exit status deciding, and every finding marked
`provenance: external`. Project defined predicate types are not in v1, since an
external gate covers the same ground and keeps the boundary honest.

Done when precedence, `binding` and scope-versus-path are enforced, when an
unresolved collision fails, when two colliding binding rules on different levels stay
red, when a non conforming commit subject inside the range goes red and a missing
range goes red as well, and when a modified external gate refuses to run rather than
running unnoticed.

### WP5 Assumption register and question resolution

`assumptions.yaml` as a mapping carrying the intent level fields, its lifecycle
across phases, confirmation recorded in the repository, and gate G-Assumptions.

With it the resolution of open questions, which needs no file of its own: a question
carries a key and a `decisions` entry in any later phase resolves it. The three exits and
why the third is not optional are in the process definition. The gate that checks it,
G-Questions, is not built here: it walks every phase artifact of the intent, which is
what G-Complete does, so it belongs to that package rather than to one that only has to
read a single file at intent level.

Done when an open assumption blocks a phase, when a confirmation recorded in a commit
releases it, and when a `decisions` entry two phases later is recognised as resolving a
question raised in P1.

### WP6 Evidence handling

The `evidence` declaration as an optional frontmatter block, in-repo storage for
small textual results, referencing for external ones, and gate G-Evidence.

Nothing parses a report, evidence is bound by content rather than by commit, and
staleness is G-Freshness's question rather than this gate's. All three are settled in the
Evidence section of the process definition.

What this package builds against them: the declaration with its `result` and `format`
fields, the resolution of `path` and `uri` items, and the hash comparison. Scanner output
is covered by the same mechanism, `kind: scan`, and needs no code of its own: the
threshold stays in the scanner and the pipeline enforces it, while the trail gains that a
scan ran and which one, which a green pipeline records nowhere in the repository.

**Evidence comes from CI.** `evidence.source` in `project.yaml`, `ci` by default and
`local` as the exception, recorded in the artifact. An item that exists at finish time
carries its hash in the sealed declaration; an item a pipeline has yet to produce
declares only kind and job, and what arrives later goes to `evidence/attached.yaml`,
outside `artifacts_hash`, with result, reference, hash, pipeline, job and commit.
`pipeline`, `job` and `commit` are recorded and read by no gate.

**Attachment is pulled, never pushed.** CI writes nothing. `xeno phase start` attaches
what has arrived for its immediate predecessor before anything else, and `xeno gate run`
does the same before it computes, since P5 has no successor. Both write only
`evidence/attached.yaml` and `gate.yaml` of that one predecessor. A verdict with
outstanding declarations is `provisional`, a state of its own, and `phase start` refuses
with a plain reason when what it needs has not arrived.

Done when a declared scan report resolves like any other evidence item, when an item
whose content changed is rejected, when evidence declared before a squash merge still
resolves after it, when a phase with a pending item finishes and pushes without waiting,
when the next `phase start` attaches it and carries the predecessor's verdict forward
without changing its `artifacts_hash`, when `gate run` alone does the same for P5, and
when nothing in a pipeline produces a commit.

### WP7 Runner and entry point

The CLI, the normalising entry point with `XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA`
and `XENO_HARNESS`, the resolution order, and gates G-Secret, G-Freshness, G-Build,
G-Test, G-Complete and G-Supply. Fixed version, signed artifact, checkable hash.

**Two surfaces with different promises.** `xeno gate ...` never opens a socket, and
that is a test in the platform matrix rather than an intention. The rest of the CLI,
phase start and write-back, may use the network. Keeping the boundary inside one
binary is cheaper than two binaries and only holds if it is asserted somewhere.

**`xeno phase finish` closes a phase in one act**: write the digest, write the cost
record where WP13 has landed, compute the `artifacts_hash`, run the gates, write
`gate.yaml` while carrying existing decisions forward by finding id. Carrying them
forward is where decisions get lost if nobody thinks about it, because the file is
written whole each time.

It does not render: `output.md` was rendered on every section write and is already what
it will be.

The command exists because the local run stopped being optional the moment `gate.yaml`
became part of the trail, and an obligation that is a sequence of remembered steps is one
people get wrong.

**The hash binds by content, not by commit.** Appendix B of the process definition
defines it to the byte, and that definition is normative: it covers the files lying
directly in the phase directory, excluding `gate.yaml` and `cost.yaml`, normalises them
to LF, and forms a `sha256sum` compatible stream sorted by path in byte order. Building
it from the prose elsewhere rather than from the appendix is how two implementations end
up one byte apart, which breaks the CI comparison silently. A command that writes inside
the hashed set recomputes what it invalidates.

The corpus in WP17 carries the appendix's own example as its first fixture, so a
deviation shows up as a failing test rather than as an intermittent CI disagreement.

**The commands that write decisions.** `xeno gate approve <finding-id> --reason`,
`xeno gate override <finding-id> --reason`, `xeno obligation close <finding-id>` and
`xeno intent close`. These are the reason `gate.yaml` can keep a single writer while
still recording human decisions. Without them the states exist with no way for anyone
to reach them.

**G-Supply verifies against a digest the runner carries.** Plugin and runner are
released together under one version, so the binary knows the digest of its own plugin
and recomputes it over `.xeno/plugin/`. Nothing in the repository states the expected
value, and a vendored plugin from an earlier release fails on the digest without a
separate version check. Version equality is therefore exact rather than major
compatible, which makes an upgrade one act: new runner, `xeno init --vendor` again,
both in one commit.

**Drift is recorded, not failed.** `gate.yaml` carries the `rules_hash` and
`secrets_hash` in force at run time. Where they differ from the artifact's, the
difference lands in `drift` and renders into the P5 checklist.

G-Build and G-Test read declared results rather than running builds or suites
themselves. G-Freshness compares each phase's context hash against the current state
of its predecessor, which is what keeps a redone phase from silently invalidating
everything after it. G-Complete runs in two modes and from two places: phase
completeness for a merged intent, as part of P5, and the closing conditions for an
abandoned one, from `xeno intent close`, writing the intent level `gate.yaml`. An
intent abandoned in P1 never reaches P5, so a gate bound to that phase alone would
never see the case it exists for.

G-Questions is built here for the same reason: it walks every phase artifact of the
intent rather than a single file, which is what this package already does for
G-Complete. It takes effect at P5, and it spares an abandoned intent, for the reason
given in the process definition.

**Language: Go**, for the runner and the MCP server both. The runner is started by a
hook during agent work, runs in CI, evaluates gates without a network, and ships as a
signed binary and an OCI image into environments with no preinstalled runtime. That
argues for a static binary with a fast start and no runtime dependency, which is
what Go gives with plain cross compilation and a `FROM scratch` image.

**Output and exit codes.** Findings are printed for people by default, one block per
finding with file, cause and next step, and `--format json` prints the same content for
anything that wants to consume it without parsing prose. The flag exists even though
`gate.yaml` already holds the structured result, because a hook and a pipeline step
often want the answer without opening a file.

**Exit codes are the same staircase everywhere**: 0 where the answer is positive, 1
where it is negative, 2 where the question could not be evaluated. For a gate run that
means the derived status decides, as the process definition sets out. The distinction
between 1 and 2 is what lets a pipeline route a problem to the right people, and it costs
nothing because the runner tells the two apart internally anyway. A broken runner writes
no `gate.yaml` and says on stderr what it could not do, so the write-back in WP12 reports
a failure to run rather than a red gate that never happened.

**The declarations both the runner and the documentation read** are YAML files in the
repository, embedded into the binary at build time. They are separate files by owner,
because a session works one package on one branch and should touch one file:

| File | Owner | Holds |
|---|---|---|
| `phases.yaml` | WP1 | the six phase ids, their names and the artifacts each produces |
| `gates.yaml` | WP7 | every gate with its id, its position in the order, the phase it applies from, and what it checks |
| `patterns.yaml` | WP4 | the named patterns a `commit-message` rule may reference |
| `predicates.yaml` | WP4 | the predicate types a rule may name, with their parameters |

Templates are declarations too and stay one file per template, owned by WP3, together
with the strings bundles. The shape of `template.yaml` and of a bundle is settled in that
package rather than here, because form and content are written in the same sitting.

The owner is who introduces and maintains a file, not who may touch it. `gates.yaml`
grows as packages land, since almost every one of them adds a gate, and that is the
normal case rather than a conflict: these packages describe the building of the tool and
not its use.

`predicates.yaml` exists so that the reference page WP16 generates for predicate types
comes from the same source as the runner, like every other generated page. A catalogue
that lives only in code would be the one the documentation describes from memory.

Nothing is copied into the vendored plugin. The runner reads the embedded copy, a
second copy beside it would look authoritative without being it, and a project that
wants to know which gates apply reads the generated documentation, which exists for
that. One source, one place to look.

The order in `gates.yaml` is the execution order, so changing it is a declaration change
rather than a code change, and the documentation page and the phase diagram in WP16
follow from the same file. A gate the runner does not implement yet is still listed
there, which is what `result: not-implemented` reports against.

**The decisions that would otherwise be made again in every session**, written down
because two sessions making them differently costs more than the decisions are worth:

| | |
|---|---|
| Layout | `cmd/` for the binaries, `internal/` for everything else, with the core and the adapters as separate packages so the import test in section 3 has something to check. No `pkg/` until somebody actually imports something |
| Module path | the repository's own URL, as `go.mod` states it. The value is not repeated here, because the namespace is deliberately not recorded in this document |
| YAML | `github.com/goccy/go-yaml`, pinned |
| Minimum Go version | the `go` directive in `go.mod`, which the runner installs rather than finds |

**The YAML library does not decide any hash**, which is worth saying because it looks
as if it might. Every hashed file is written exactly once, `context.lock.yaml` at phase
start and the rendered files at phase finish, and afterwards only read. CI hashes the
bytes it finds rather than re-serialising anything, and `gate.yaml`, the one file
rewritten on every run, lies outside the hash. A different library would only write
future files differently, and their hashes come from those bytes. What a library change
does cost is a diff without content, which is a nuisance and not a correctness problem.

Agent framework maturity is deliberately not a criterion. Neither the runner nor the
MCP server ever calls a model, so there is nothing in v1 that an agent SDK would
help with. Autonomous agents, if they come, are a separate component beside Xeno
rather than inside it, and MCP being language neutral means such an agent can be
written in whatever its own ecosystem favours while talking to the same server.

One toolchain, not two: a more comfortable MCP SDK elsewhere does not pay for a
second language in a single person project.

**Invocation paths and verdicts.** The hook runs G-Secret on every write and
G-Schema only when a write came through the MCP operation that completes an artifact.
The local run and CI run the full set for the phase. Two things are true of the local
run at once and both have to be built in: its verdict binds the start of the next
phase, and running it is not optional, because it is the only writer of `gate.yaml` and CI never writes. The
reduced hook scope is a requirement, not an optimisation: a hook that resolves the
rule set on every file write gets switched off, and a hook that reports missing
fields in a half written artifact gets switched off faster.

**Gate order is integrity first, then cost.** G-Supply leads, G-Rules and G-Policy
close, and the order is a declaration the runner reads rather than a sequence in
code, so the reference documentation in WP16 derives from it.

**Secret filter and digest.** The runner reads the effective filter, shipped file
plus additive project file, for both G-Secret and the digest. It filters, hashes and
writes `digest.md` from the summary text the agent supplies, so the filtering sits
outside the model's reach.

**Ending and overriding.** An intent can be closed as `abandoned` with a reason, and a
merge can proceed on a failing finding recorded as `overridden` with person, reason,
timestamp and an obligation that can be closed. Both are ordinary states of the runner,
not special cases bolted on later. A merged intent records nothing about its merge: the
join to deployment tooling runs through the intent key in the merge commit message, not
through a field that could only be written after the gate that needed it.

**Retention.** `.xeno/local/` is pruned on every run against the project's retention
setting, default 30 days, counted from phase completion. Sessions of a phase that has
no `cost.yaml` yet are skipped, so a long running phase does not lose its figures
before anyone has written them down.

**Diagnostics are a requirement, not polish.** A red gate names the file, the rule
or field at fault, and the next step. A verdict without those produces frustration
rather than governance, and for a checking tool that is half of its usability.

**The sequence is enforced.** `xeno phase start` refuses a phase whose predecessor holds
no completed `gate.yaml` with a matching `artifacts_hash`, whose predecessor is red or
provisional, or which is already running. Cheap here, load bearing later.

Done when the same binary produces the same verdict locally and in CI, when a phase
started out of order, twice or on a red predecessor is refused with the reason named,
when an approved or overridden predecessor lets it start, when an approved
finding is recorded with approver and commit rather than rewritten to green, when a
phase status cannot be written directly, when a modified or downgraded plugin fails
against the digest the runner carries, when a secret in the session text does not reach
`digest.md` regardless of what the model proposes, and when every red verdict names a
file, a cause and a next step.

### WP8 Context profile and change driven re-reading

`context-profile.yaml` as a P0 artifact with include, exclude, declared links and
budget, validated by G-Schema, plus change driven re-reading from the hashes already
in `context.lock.yaml`. Phases read the preceding digest rather than rescanning the
codebase. Reading outside the profile is recorded, not blocked.

This half needs no baseline, only hashes that already exist, and P0 cannot produce a
profile until the format is fixed. The measured half is WP15.

Done when a repeated phase reads only what changed, and when every read outside the
profile appears in `context.lock.yaml`.

### WP9 Initialisation on an existing repository

`xeno init` on a brownfield repository: what it creates, what it refuses to touch,
which `.gitignore` entries it adds, how the CI wrapper is generated, and how
`--vendor` pins the plugin. This is every user's first contact with the tool.

**It asks three questions and defaults the rest.** The tracker with its project key, the
default model, and the language of the artifacts. Everything else in `project.yaml`
ships with a defensible default that a project changes when it has a reason to. Put to a
newcomer as a questionnaire, the configuration is long enough that the first interested
party gives up around the fourth item. Adoption is decided here more than anywhere else
in this plan, and this package has had the least thought spent on it.

It also checks the locally installed runner against the version pinned in
`project.yaml` and stops with an installation instruction on any difference, not only
a major one, rather than proceeding with whatever happens to be on the machine. The
runner carries the digest of exactly one plugin, so runner and vendored plugin move
together or not at all. The installation route itself belongs in onboarding, WP16.

It also writes the `enforcement` block from what the host offers, records what the
host cannot express as waived, prints the settings the project still has to make by
hand, and says plainly that it cannot make them. Determining what the host offers needs
the API and a token; without one it writes the block with defaults, marks the edition
check as outstanding and carries on, rather than failing on the first command a new user
runs. A first contact that leaves the
project believing the gate is binding when it is not is worse than no first contact.

Two pieces of setup belong here rather than in a debugging session six weeks later: the
scheduled pipeline that runs `xeno enforcement check` daily, and a token with read
access to the protected branch settings. `xeno init` names both and creates neither.

Done when running it twice changes nothing the second time, when it never edits a file
it did not create, when a version mismatch stops it, and when the host settings a
project still has to make are on screen rather than in a document nobody opens.

### WP10 CI wrapper

One thin wrapper generated from a template, a workflow calling nothing but
`xeno gate run --phase <n> --base <ref> --head <ref>` with the runner version pinned.
Both ends of the range are passed rather than derived, because a CI system may check
out a merge commit it produced itself and judging that would mean judging a commit
nobody wrote. GitHub reports both ends on the pull request event, as
`base.sha` and `head.sha`.

The generator takes the two parameters from the start even though there is one
wrapper to generate. That is what makes the wrappers of 1.1 a template argument
rather than a rewrite.

**The host configuration is part of the deliverable.** A CI job that reports a red
verdict and does not stop the merge is a report, not a gate, and the claim that CI is
binding rests entirely on one setting outside this repository: the pipeline required
for a merge on the protected default branch. `xeno init` prints the settings the
project has to make and states that it cannot make them.

**`xeno enforcement check` compares them against the host.** It is a separate command
and the first job of the same pipeline, not part of `xeno gate`, because it needs the
network and the gate path's freedom from it is what makes a verdict reproducible. It
reads the protected branch configuration through the API with the enforcement token
from the credentials table in section 3,
compares it against the `enforcement` block in `project.yaml`, and writes a report as a
pipeline artifact. The declaration lives in `project.yaml` rather than in a file of its
own, because it is four lines and a seventh configuration file costs more than it
returns.

```yaml
# project.yaml
enforcement:
  required_pipeline: true
  allow_bypass: false
  merge_method: no-squash        # only meaningful with a commit predicate active
  approvals:
    required: 1
    not_by_author: true
    waived: "not expressible on this edition, recorded 2026-09-16"
```

Three properties decide whether this is worth anything.

It distinguishes **not set** from **not available**. A host answers `403` for a
repository whose plan does not include protected branches, and the approval fields are
then absent rather than empty; reporting that as a missing setting would send somebody
looking for a checkbox that is not there.

That distinction is not enough on its own, which is why `waived` exists. A requirement
the edition cannot express would otherwise be reported as unmet on every run forever,
and a report that always says the same thing is ignored within a fortnight. `xeno init`
therefore writes the declaration from what the host actually offers and records
everything it cannot offer as waived with a reason and a date. The unmet requirement
becomes a decision in the repository rather than a standing complaint, which is also
the form a project's records need for it.

It writes **no commit**. The report is a pipeline artifact and a job status, which is
what WP10 delivers; the merge request comment that carries it to people is added by
WP12. CI verifies artifacts that exist and does not produce any, and a second place
carrying a verdict beside `gate.yaml` is exactly what the schema forbids.

It **cannot make itself binding**. Where the pipeline is not required, this check
failing stops nothing either. That is a circularity with no way out from inside the
repository, so the value comes from visibility: the job status, later the merge request
comment, and a scheduled run, daily, whose failure shows in the pipeline overview even
when nobody is merging. The scheduled run is the most useful of the three, because the
change it catches is the one an administrator made between two merges.

Its exit codes follow the same staircase as the gate run: 0 where the host matches the
declaration, waived requirements included, 1 on a divergence, 2 where the host could not
be asked at all. The last one is the case the staircase exists for, because an expired
token and a loosened branch setting both fail the job and need entirely different people
to look at them.

The token needs read access to the protected branch settings. That belongs in the
project setup, not in a debugging session six weeks later. Whether the job token CI
provides automatically can read them is to be verified rather than assumed; in common
configurations it cannot, which is why the table names a project access token for it.

**What a host can cover, what it cannot, and where each setting lives** is documentation
rather than plan: it belongs to the adaptation part in WP16 and is written there. Two
statements belong here because they shape what this package builds.

A host can take over the checks whose subject is git metadata or the pipeline result, and
none of those whose subject is the content of an artifact; ten of the fourteen gates
therefore have no host equivalent at all. Where both could act, both should: the host
stops earlier and with an authenticated identity, the gate records what held, and a
statement that lives only in an audit log does not travel with a clone.

**A host setting can also break a gate.** The merge method decides whether the commits a
predicate judged still exist afterwards, and with squash they do not, which is limitation
13 of the process definition arriving through a setting nobody thought of as a gate.
`xeno enforcement check` therefore reads the merge method too and reports it where a
commit predicate is active in the rule set.

**The commit message format cannot be checked at push time.** GitHub has no server
side hook a repository can install and no push rule on the message, so the client side
hook is feedback that `--no-verify` removes, and the format becomes binding at the gate
and nowhere earlier. Other hosts put the same capability behind a paid tier or behind
shell access to the machine, so this is not one host's shortcoming.

This is a missing enforcement point and not a missing
gate: every gate runs on this edition, the platform simply stops less early. Projects
on a host that can do it should do it there, and the documentation says which setting
on which host.

**CI recomputes and compares, it never writes.** What is compared and what legitimately
differs is set out under Execution in the process definition. What this package builds is
the job: recompute, compare, fail on a stale or disagreeing verdict, and fail with a
clear message where a phase arrives without `gate.yaml` at all.

**Host settings are not versioned.** They change without a commit and whoever holds a
bypass right escapes them, which is the one part of the setup the repository cannot
evidence by itself. The check above narrows it and does not close it.

The wrapper reports a provisional verdict as provisional, with the number of open
declarations, rather than as a divergence. The local run knew those items were open and
so does CI; treating that as a failure would make every push out of P4 fail
verification, and verification nobody believes is worse than none. Only on the merge
request is a provisional verdict a hard condition.

Done when the wrapper produces the same verdict as a local run on the same repository
state and the same commit range, when a provisional verdict is reported as such and
does not fail the job before the merge request stage, when a repository whose pipeline
is not required for merging is reported as such rather than passing quietly, and when a
requirement the edition cannot express is reported once as waived rather than daily as
unmet.

### WP11 Agent layer

`plugin.json` per Agent Plugins 1.0.0, six phase skills plus the cross cutting
learning skill, four lenses, `mcp.json`, the MCP server exposing the process
operations (read intent, record assumption, write artifact, fetch template, run gates
locally), the summary text the runner turns into a digest, and hook wiring for both
clients.

The lenses (security, privacy, operations, architecture conformance) are skills, not
subagents and not roles, so they work in both clients. Each declares the phases it
applies to and is enabled per project. Their output goes into open questions, the
assumption register or the review checklist, never into an artifact of its own and
never into a gate verdict. A checklist entry from a lens carries `source: lens` and no
rule id, so it cannot change the set G-Policy checks for completeness.

**The plugin is served from the repository, not from a public directory.** The
marketplace wrapper stays and points at this repository, because a client can take a
marketplace from any git repository it can reach. Nothing is registered anywhere public,
and nobody should build a path for that. Whether the client accepts that URL directly is
the one thing to verify here rather than assume.

**One thing left to verify rather than assume.** It is the marketplace URL above. The
gateway was the second and is answered in section 9: it reports the deployment a request
was routed to, and it records request metadata sent under its own header names. The
record of which model was used and the attribution of cost below the project therefore
both have a source, and what remains of the exercise is sending the values from here,
which is the export below.

**The MCP tool surface is a budget, not a list.** Five operations here, six once WP15
adds the index query, and a seventh needs an argument of the kind the index has. A tool
definition is sent with every request of every phase whether it is called or not, so the
surface is a standing cost on all work the project ever does.
The same holds for each enabled lens, which is why enablement is a project decision
and not a default.

**Every MCP operation has a command behind it**, including supplying section content,
which is otherwise the one step that would force somebody to write `output.md` by hand
and get the anchors right unaided.

**Requests carry the intent and the phase, through the harness.** In v1 the runner makes
no model request; the harness does. `xeno phase start` therefore exports the qualified
intent id and the phase id into the environment the harness reads for request headers,
so that a gateway sees them. Whether each harness forwards custom headers, and under
which variable, is to be verified per harness rather than assumed, and a harness that
does not is recorded as such rather than silently producing unattributed figures.

Done when a phase can be driven from either agent with no agent specific logic in the
runner, when a model request made during a phase reaches the gateway carrying intent
and phase for every harness that supports it, when a disabled lens changes nothing but the findings, when the tool surface has
not grown without a recorded reason, and when the same phase can be carried out with
commands alone, with no MCP server and no hooks, producing artifacts a gate cannot tell
apart from the ones an agent produced.

### WP12 GitHub adapter

Reading an issue when a phase starts and writing the gate result back to the merge
request when CI finishes. One tracker in v1, selected by configuration all the same,
because the selection is what keeps the adapter from becoming the runner's assumption.

**A hosted service is the normal case and self managed is not a variant.** The base URL
has a default and is a parameter all the same, and the adapter is tested against the
hosted service it is written for. What keeps the other case open is that the address is
never a constant.

**It is not on the critical path and it is not optional.** Since the trigger is the CLI,
an intent can be carried through all six phases with the issue content copied in by
hand, which is why M1 does not wait for this package. What the adapter buys is that
nobody has to, and that the result reaches people who do not work from a command line.
That is the difference between a tool a team tolerates and one it adopts, so the
package stays in v1 and moves behind the milestone rather than out of the release.

The adapter contract stays deliberately small: resolve an intent, read an issue,
write a comment, resolve credentials. Everything else belongs in the runner.

**No issue commands.** Receiving them needs a component that listens, a webhook
service or a job wired to comment events, and neither earns its keep while the
alternative is a command somebody types. The trigger in v1 is
`xeno phase start --intent PROJ-123 --phase 01-requirements`.

Write-back stays, and this work package adds it to the wrapper WP10 generated. The
CI job runs the gates anyway, so the result reaches the issue from a job that already
exists, and the wrapper template gains one step rather than the project gaining a
component.

**Two requirements that keep 1.1 cheap.** The adapter knows no relationship between
issue key and repository, and the contract carries its endpoint and authentication
scheme as parameters rather than as constants. Jira is the case that breaks a design
built on the other assumption, and finding that out in 1.1 is fine as long as nothing
in the contract has to be redrawn to accommodate it.

Credentials come from the environment, never from `project.yaml`. The endpoint is a
setting with a default, because a self managed deployment has no canonical address, and
a second host in 1.1 needs no change to the contract.

```yaml
# project.yaml
tracker:
  adapter: github
  base_url: https://api.github.com   # a default, overridden for Enterprise Server
  auth: { scheme: token, secret_env: XENO_TRACKER_TOKEN }
```

Done when a phase started from the CLI carries the issue content into P0, and when the
gate result and the enforcement report appear on the merge request after a CI run.

### WP13 Token recording

Evaluation of local session logs into `cost.yaml`: tokens per phase, split into input,
output and the cached share, with `evidence: self-reported` mandatory.

**Tokens, not money.** The figure that can be acted on is the token count, and it is
also the one the sessions actually report. Converting it into an amount is a matter of
prices, plans and negotiated terms, none of which belong in a repository, and a central
cost view is coming from the model access point anyway. `cost_usd` therefore stays in
the schema and is optional, filled where a project wants an indication and left out
where the number would only look authoritative.

**The split is the point.** Output costs several times input, a cached prefix is a
fraction of either, and one combined number hides the effect of every lever in WP20.
Where a harness does not report the breakdown, the field is absent rather than
estimated.

Done when a phase produces a token record that is honest about its own reliability, and
when two runs of the same phase can be compared without opening the session logs.

### WP14 Learning

`learning.yaml` per phase and the closing record an abandoned intent writes at intent
level, both as mappings carrying the common fields rather than as bare lists, gate
G-Learning, and generation of the merge request against the rule set. A merged intent
writes no record at intent level; it has one per phase already.

Done when a learning record cannot be skipped, when an abandoned intent cannot close
without one, and when none of them can take effect without a merge.

### WP15 Symbol index and tuning

A symbol index over the repository and the tuning of what each phase reads. The index
answers where a name is defined so that an agent stops reading whole files to find out,
which the process definition calls the largest single saving in context economy.

**It is used at phase time and never in a gate.** That is the property everything else
follows from: it touches neither the network freedom of the gate path nor the
determinism of a verdict. So it is allowed to be missing, stale or wrong without the
trail suffering, and it is therefore optional. Where no index is available the phase
runs without one and `context.lock.yaml` records that it did.

**Tree-sitter as the engine, with no grammar shipped.** Xeno builds the index and
loads the grammars a project provides, at a location the project configures. It vendors
none of them, and that is a licence decision as much as a scope one: every grammar comes
with its own terms, and which terms a project can accept is that project's business and
not something a tool should decide by bundling.

The documentation points at the four that matter here and goes no further:

- Tree-sitter itself: <https://tree-sitter.github.io/tree-sitter/>, with the API
  overview under <https://tree-sitter.github.io/tree-sitter/using-parsers/>
- Java: <https://github.com/tree-sitter/tree-sitter-java>
- C#: <https://github.com/tree-sitter/tree-sitter-c-sharp>
- TypeScript: <https://github.com/tree-sitter/tree-sitter-typescript>
- Kotlin: <https://github.com/tree-sitter-grammars/tree-sitter-kotlin>

**Loaded at runtime, not linked in.** Grammars arrive as shared libraries and are loaded
when the index is built, which keeps the C toolchain out of Xeno's own release: the
runner stays pure Go and cross compiles per platform as before. Where a project supplies
no grammar, the index is empty for that language, and where it supplies none at all, the
phase runs without an index, which is the degradation this package is designed around
anyway.

The honest cost: out of the box there is no index. A project that wants one obtains the
grammars, checks their terms, and points the configuration at them. That is a step more
than a bundled tool would need, and it is the step that keeps other people's licence
decisions out of this repository.

**What it holds.** Symbol name, kind, file, line, and the enclosing container. Not call
relationships and not type resolution, which is the semantic graph the scope section
puts outside v1. An index says where to look; a graph answers questions, and answering
questions is what the agent is for.

**Where it lives.** `.xeno/local/index/`, never committed, covered by retention.
Derived data does not belong in the trail, CI has no use for it, and what is not
committed cannot contaminate the record.

**How it enters the record.** Through the `tools` block of `context.lock.yaml`, which
was reserved for exactly this: name, version and the hash of the answer. The artifact
then states which index, in which version, returned what, and replacing the engine later
is a tool change rather than a schema break.

**How the agent reaches it.** One MCP operation, queried instead of searching the
repository. That is the sixth tool, and WP11 requires an argument for a sixth. This is
the argument, and it is also what the budget was for.

**Freshness.** Rebuilt incrementally at phase start over what changed since the last
build, with a full build as the fallback. A stale index is worse than none, so its age
is part of every answer.

Done when the index returns correct locations for a known set of symbols in every
language a grammar was supplied for, when a missing or unreadable grammar degrades to no
index rather than to an error, and when the bytes read per phase drop against a run
without it. Bytes read are counted by the runner itself, which makes this acceptance
independent of any session log format.

### WP16 Documentation

Process definition, this plan, setup and onboarding including the runner
installation and hook trust, and the limitations of v1 stated where a reader will
actually see them.

**Structure.** Four parts, because they differ in who writes them and who reads them.

- *Getting started.* One intent played through from the issue line to a green gate in
  CI. The process definition answers what holds, not how to begin, and the first
  contact with this tool is `xeno init` in somebody else's brownfield repository.
- *Using it.* `xeno init`, triggering a phase, checking locally, the override, ending
  an intent as abandoned. Also what a project does with the model and tool fields
  where it keeps a register of its ICT service providers, which is to link its own
  register from its README. Xeno keeps none.
- *A worked intent.* One complete intent with its real artifacts, from the issue line
  to a green gate. This is the golden file fixture from WP17 rather than a second,
  prose version of it: the page shows those files, so the example stays current because
  breaking it breaks a test, and there is one place to maintain instead of two.
- *Reference.* Generated in full: commands, configuration schema, gates with the
  findings each can produce, predicate types, rule catalogue, template section ids.
  A page per gate belongs here, saying what it checks, why it goes red and what to do
  then. It derives from the same declarations the gate list and the phase diagram
  come from, so it is derived material and not prose.
- *Removing it.* Half a page, no command and no tooling. Finish or abandon the running
  intents, then take out the CI wrapper and the client wiring, which stops new intents
  while leaving everything written so far valid. Remove the vendored plugin, the
  configuration and the `.gitignore` entries if the tool is to go as well; the intent
  directories are ordinary Markdown and YAML and stay readable without it, and Appendix B
  of the process definition says how the hashes are recomputed. Nothing has to be undone
  on the host, because `xeno init` names those settings and never sets them. And the
  honest sentence rather than a promise: deleting the directories does not remove the
  history, since what was once committed stays in the old commits. A tool that comes out
  cleanly is tried more readily, and in a governance context somebody asks regardless.
- *Adapting it.* Written by hand: cutting a `checked` rule, when something is better
  as `review`, overriding a single template without forking the set, lenses, external
  gates, and adopting an example from `examples/rules/` by copying it rather than
  looking for a switch. This is where users create their own artifacts and where
  mistakes are expensive, because G-Rules only finds them afterwards.

  The host side is written here rather than in the plan: which of a host's mechanisms
  cover which gate and which only the gate can do; which settings make a verdict binding
  and where each of them lives on GitLab and on GitHub, including which tier has them;
  which settings can break a gate, the merge method above all. It is the part a project
  has to act on before its first intent, and finding it out unaided is how a team
  concludes the whole thing is not worth doing.

- *Limitations, in one place.* A generated page pulling together what the process
  definition, the delta and the evaluation each say the tool cannot do, each entry
  naming the document it comes from and the state it is in. Today that means reading
  three documents and reconciling the changes in one's head, and it is the list most
  often quoted, by people looking for weaknesses. Generated rather than written,
  because a hand maintained copy of it would be wrong within a release.
- *The documents.* The four under `docs/`, published as they are rather than
  rewritten: process definition, implementation plan, orchestrator evaluation, v2 delta.
  They are results rather than scaffolding, and they answer the questions a reader has
  before deciding whether any of this applies to them. The conformance mapping is the
  clearest case, because it will be read by people who do not open a repository, and
  the same goes for the limitations, which are the section most likely to be quoted back.
  Publishing them costs nothing, since the site is built from the repository anyway, and
  it removes the temptation to keep a second, friendlier copy of them somewhere else.

The versioning policy from the process definition sits next to the upgrade
instructions, where a user meets it: what a major step means per object, and that
artifacts are read in their old schema rather than rewritten.

**Outside contributions need an answer before the repository is public.** From 1.1 it is
open, and somebody will send a merge request with no intent, no phases and no artifacts.
Accepting it leaves a hole in the trail of the tool whose whole claim is that there are
none; refusing it makes this an open source project that does not take contributions.
The answer belongs here, in writing, before publication rather than after the first such
request: which parts of the repository are covered by the process at all, and what a
contributor is expected to do.

**English only in v1.** The translation follows in 1.1, together with the hash
binding that keeps it from rotting, and the binding gets built when the first
translated page exists rather than a year before. Where it lands then is the first
two parts: the generated pages are English anyway, since the process layer is never
translated, and the adaptation part addresses the same people who write rule files
with English keys.

This is the documentation, not the artifacts. Both strings bundles ship in v1, so a
project writes its artifacts in German from the first release.

**Tooling: MkDocs with the Material theme.** Per page frontmatter, anchor and link
checking, and Mermaid are all there without assembling them, and the alternative worth
naming is Hugo, a single static binary with no runtime in the pipeline, which fits the
rest of this project better and would have to have its checking put together by hand.
MkDocs was chosen because the checking is the point here and because this is the one
place where a Python dependency in a documentation job costs nothing that matters. The
choice belongs in this work package rather than in a later decision, because the source
hash binding and the extraction of code examples both depend on it.

**Each document carries a stable id and a revision in its frontmatter**, which the site
needs before it can publish them. A cross reference has to point at an id rather than at
a file name, or renaming a file silently breaks every link to it, which is how this has
already gone wrong once. Title, status and path are not enough for that.

**The documents part can be built before anything else in this package, and should be.**
The reference is generated from declarations that do not exist yet and the handwritten
parts describe commands that do not exist yet, but the four documents exist today and
already carry cross references between them, to sections, to limitations by number and
to files by name. None of those is checked by anything, and some are known to be wrong,
which is exactly the rot between documents this package exists to prevent. Standing the
site up early costs one job and turns every later reference into something that fails
loudly rather than quietly.

**Published from CI to Pages.** The documentation is built and
deployed by the repository's own pipeline on every merge to the default branch, so the
published site cannot drift from the source. The audience is internal for as long as
Xeno is, which changes who reads it and not what it has to say. Same principle as
everywhere else here: the repository is the source, the site is a derivative.

**Consistency is mechanical where it can be.** Documentation cannot be checked for
being right, only for being stale or self contradictory. Four measures cover most of
it:

- Anything derivable is derived. Command reference, configuration schema, gate list,
  rule catalogue, predicate type reference and template section ids are generated
  from the declarations they describe, never written by hand. A generated page cannot
  drift. The rule of thumb: when a table in the documentation carries the same
  information as a file in the repository, it gets generated.
- Code examples are extracted and compiled or executed in CI, which catches the
  example that stopped working two refactorings ago.
- Links and anchors are checked, which catches rot between documents that nobody
  notices otherwise.
- One statement, one place. The drift no tool finds is duplicated prose, so a
  decision is written once and referenced rather than repeated.

**Diagrams carry more than prose does.** The documentation leans on pictures where a
picture is clearer, which is most of the process:

- the phase flow with its gates, from issue to merge
- the gate verdict path per finding, and how the phase status is derived from the
  decisions on them: green, red, approved with approver and commit, overridden with
  its open obligation
- what an intent looks like as a graph: artifacts, their context hashes, declared
  evidence, references across repositories
- the three invocation paths and the two verdicts behind them
- rule precedence across the four levels and both origins
- what happens when a phase is redone and later phases go stale

Diagrams are Mermaid sources in the repository, never committed as images. Mermaid
needs no extra runtime, the code host renders it directly in Markdown, and the source
stays diffable, so a change to a diagram is reviewable like any other change.

The phase and gate diagram is generated from the same declarations the runner reads,
not drawn by hand. It belongs to the derived material above: a hand drawn flow is the
first thing to go stale.

Hash binding, as used for translations, applies to a few high value pairs only. It
proves that somebody looked, not that the text is right, and pasting it everywhere
turns review into clicking receipts.

Done when a merge publishes the site unattended, when no generated page is maintained
by hand, when the phase diagram changes by itself after a phase or gate declaration
changes, when every gate has a page explaining its red verdict, and when somebody
without prior knowledge gets a first intent through on their own.

### WP17 Test strategy for the runner

A tool whose purpose is deterministic checking needs unusually precise checking of
its own. Golden file tests for the renderer, a corpus per gate of artifacts that
must come out red and artifacts that must come out green, and a platform matrix that
covers Linux in v1.

**The corpus lives centrally, under `corpus/` at the repository root**, rather than in
a `testdata/` directory beside each package. It is read by more than the package it
would otherwise sit next to: the renderer tests, the gate tests and later the
documentation examples all draw from it, and a fixture that demonstrates a rule is
worth finding without knowing which package implements it.

**Linux first, Windows after.** Development and CI run on Linux, and WSL is Linux for
everything that matters here. What is deferred is the integration on a native Windows
checkout: `core.autocrlf`, path lengths, file permissions.

**Normalisation is not deferred with it.** Line endings, path separators and sort order
define every hash the tool ever computes, and introducing them later changes every
recorded value in every repository. They are built now, and they are testable now
without a Windows machine: the corpus carries fixtures with CRLF and backslash paths,
and the expected hashes are the same as for their Unix equivalents. The Windows run
adds the integration, not the rule.

The WSL assumption is worth holding loosely. In WSL everything is Linux; the case that
breaks is a native Windows checkout, and sooner or later somebody will have one.

The corpus covers commit ranges too, not only artifacts. The commit predicates are
the only checks reading something other than files under `.xeno/`, which makes them
the easiest to leave untested and the likeliest to behave differently between a local
run and CI.

Two more corpora carry their own risk. Finding id stability: the same input produces
the same id, a changed cause produces a different one, and a decision follows the
first case and not the second. And the network boundary: `xeno gate ...` completes
with no route to the outside, asserted in the matrix rather than assumed.

Done when a regression in a gate fails the build here rather than surfacing as a
wrongly green verdict in someone's project.

### WP18 Dashboard, deferred to 1.1

Specified here because the shape is decided and the decision affects what the runner
exposes, built in 1.1 because nothing in v1 depends on it. Everything it shows can be
read from the repository by hand, which is exactly why it is the package that can
wait.

Read only views derived from the repository: the artifact graph for an intent,
phases completed and stale, gate results including approvals, and cost per intent
and per phase. No state of its own, no write path, no ability to trigger anything.

**The intent list** is the entry point: everything in progress, everything merged and
everything abandoned, the last with its reason. In progress means begun and not yet
merged, judged from the artifacts, which also shows how many intents are running
alongside each other and where they are. It says nothing about whether an agent is
working right now; that is live state and stays out.

**Links out.** Each intent links to its issue in the tracker and to its repository.
The qualified intent id already carries host, project and key, so the links are
constructed rather than configured, which is what will keep the further trackers of
1.1 from becoming a special case each.

Cost figures carry their `evidence` value in the view itself. A chart lends
authority that self reported numbers have not earned.

**What v1 owes it.** The read only HTTP surface and the derivation of the views are
1.1 work, but the declarations they read are not: gate results, phases, drift entries
and cost records all have to be complete and machine readable without a dashboard to
consume them. That is a v1 property anyway and worth naming, because a format that is
only ever read by a human reader tends to grow gaps a renderer would have found.

Open overrides are a view of their own, since an override is only as useful as its
visibility: which findings were merged over, by whom, for what reason, and whether
the obligation has been closed since. Approvals sit beside them, separately, because
the two say different things and a view that merges them says neither.

Gate findings carry their provenance the same way. Ten green gates look alike no
matter how many of them came from external code, so `provenance: external` is shown
in the view and not only in `gate.yaml`.

**Packaging.** The UI ships bundled into the runner binary, which keeps the one
toolchain rule and means there is still only one artifact to deliver. It is separated
in v2 at the latest.

Because that separation is planned rather than hypothetical, the bundled UI reads
through the same interface a standalone frontend would use, a read only HTTP surface
served by the binary, and never reaches into internals. Splitting it then becomes a
packaging change instead of a rewrite.

Done when the dashboard can be deleted and rebuilt from the repository alone, when
the UI works unchanged against that interface whether it is bundled or not, and when
no figure or verdict is shown without the qualification that sits on it in the
repository.

### WP19 Conformance mapping

A table mapping every ISO/IEC 42001 requirement Xeno touches onto the artifact or
gate that supplies evidence for it, with three states per row: supplies evidence,
supports in part, outside the tool.

The third state is the important one. Without it somebody will read the mapping as a
certification package, which Xeno is not and section 1 of the process definition says
so. The work lands after M1, when the artifacts exist and the mapping can be made
against real files rather than against intentions, and it is the last package before
the v1 release.

Done when an auditor can name, for every requirement the table touches, the file to
look at, and when the list of uncovered points matches section 16 of the process
definition rather than standing beside it.

### WP20 Token economy

The levers that no other package owns, pulled in this order: measure, order, choose,
and only then trim. The context profile (WP8), the symbol index (WP15) and the digest
mechanism are already the large savings and are not repeated here.

**Baseline before levers.** After M1, a fixed reference intent is run through all six
phases and the cost records are kept. Every change below is judged against that run
and not against an intuition. The article-level advice to shorten prompts is not worth
anyone's afternoon; what is worth it is a change that applies to every phase of every
intent from then on.

**Cache friendly assembly.** Context is assembled by volatility: shipped set, then
project stable, then phase variable, then the task. `context.lock.yaml` records the
order, and the renderer never places an intent id, a timestamp or a hash near the
front. This is the one lever that costs nothing to build and pays on every request,
and it is also the one that is impossible to retrofit once the assembly is scattered
across the skills.

**Model and thinking budget per phase.** `project.yaml` sets both, with a default and
per phase overrides, and the model in use is recorded in the artifact as it already
is. Intake and verification do not need what design needs. Reasoning tokens are
charged as output, which is the expensive half, so the default for the mechanical
phases is the lower setting.

**Session discipline.** One phase, one session. The runner warns when a phase is
resumed after its session has grown beyond a configured size, because a resumed
session carries its whole history back into every request and the artifacts are a
complete enough starting point to begin again.

**Output length as an observed figure.** Templates define required sections, and a
section that invites prose is paid for in output tokens on every intent. Rendered
length per section lands in the cost record as an observation. No limit is enforced:
truncating an artifact to save tokens would trade the thing this process exists for
against its running cost.

**What does not apply.** Batch processing, despite the discount, because every phase
has a person waiting at the end of it. Explicit prompt caching through the API,
because Xeno drives a harness rather than calling a model endpoint, so the caching it
benefits from is the harness's and the only thing Xeno controls is what it hands over
and in what order. Both are named here so that nobody spends a week discovering it.

Done when the baseline intent can be re-run and compared, when the assembly order is
a declaration rather than an emergent property of the skills, when a cheaper model on
a mechanical phase is a configuration change and not a code change, and when the
measured cost per intent is stated in the documentation as a figure with its method
beside it.

## 3. Architecture

Ports and adapters, for one reason above the usual ones: it turns the tool's central
promise from a matter of discipline into a property of the build. `xeno gate` touches no
network and judges deterministically. Where the gate path cannot import the adapter
packages at all, it cannot telephone by accident, and the promise is checked by the
compiler rather than remembered by whoever writes the next feature.

**The core is pure.** Artifact schema, rendering, rule evaluation, gate evaluation and
hashing take files and parameters in and return artifacts and verdicts out. No network,
no knowledge of hosts, no clock beyond the one passed in, no reading of ambient
configuration. That is what makes the corpus in WP17 possible at all: the same input
produces the same bytes on every machine.

**The ports are the places where more than one implementation is already foreseeable**,
not every boundary somebody can imagine:

| Port | Implementations |
|---|---|
| Tracker | GitHub in v1, GitLab and Jira in 1.1 |
| Host settings | read by `xeno enforcement check`, one per host |
| Symbol index | tree-sitter in v1; the `tools` block in `context.lock.yaml` already treats the engine as replaceable |
| Session logs | one per harness, and the plan already calls their formats non contractual |

Git is deliberately not on that list. The gate path reads commit history, and tests for
it run against real temporary repositories rather than against a mock, which is both
simpler and more honest for a tool whose subject is a repository.

**A test enforces the boundary.** It walks the import graph of the gate path and fails
where an adapter package appears. Without it the rule holds until the first afternoon
somebody needs one field from the tracker inside a gate, which is exactly the afternoon
nobody is watching.

**Credentials.** Every credential comes from the environment and none appears in a file,
which the process definition already requires for the tracker. What was missing is one
place that says which ones exist:

| Credential | Purpose | Least scope | Held by | Lives in | On expiry |
|---|---|---|---|---|---|
| Tracker token | read an issue, write a comment | issue read, note write | the developer, the service account in v2 | environment, named by `tracker.auth.secret_env` | P0 and every comment fail with exit code 2 |
| Enforcement token | read protected branch settings | read on repository settings | a project access token | CI variable, masked and protected | `enforcement check` exits 2, which the staircase separates from a real divergence |
| CI job token | clone, read pipeline artifacts | what the host grants by default | the pipeline | provided by the host | not applicable |
| Gateway key | model access through the proxy | one project, the allowed model set | the developer, the run in v2 | environment | the harness fails its first request |
| Xeno's own issue token | `glab` for work on Xeno itself | API on this one project | the maintainer | environment | only affects work on Xeno |

Two rules follow from it. A token is scoped to what its row says and no further, because
the widest token in a pipeline is what an attacker aims for. And an expired token is a
different failure from a wrong setting, which is why both checks that touch the network
exit with 2 rather than 1: the people who fix the one are not the people who fix the
other.

**What this is not.** It is not a folder layout with `domain`, `application` and
`infrastructure`, and not an interface per type. Idiomatic Go does ports as small
interfaces declared where they are consumed, with concrete types everywhere else. The
testability people expect from hexagonal architecture arrives here through the purity of
the core and the golden files, not through layering, and layering added on top of that
buys indirection rather than confidence.

## 4. Working method

Xeno is built with agents, which is the only honest way to build it: a tool that
governs agentic development and was written by hand would be advice from somebody who
has not tried it. Until M0 exists, though, Xeno cannot govern its own construction, and
that gap has to be bridged deliberately rather than ignored.

**The bridge is the specification.** The process definition is normative from the first
commit, so every package is worked as if it were an intent: state what is wanted,
restate the acceptance as checkable sentences, implement, show evidence, review. The
difference is only that the record is a file in the branch rather than an artifact
under `.xeno/`, and nothing gates it but a person. From M0 the same loop runs through
the runner, and the hand held record stops.

### First steps, in order

1. **Create the repository and commit the two documents into `docs/` before any
   code.** They are the context every session loads and the reference every disagreement
   settles against.
2. **Protect the default branch and make the pipeline required for merging.** Do it on
   day one, with an empty pipeline if necessary. It is the setting the whole tool
   claims to rest on, it costs two minutes, and setting it later means working for
   weeks in a repository that does not do what the documentation says.
3. **Write `AGENTS.md` and `CLAUDE.md` at the root**, short and pointed: build and test
   commands, the line width and file layout conventions, and the three standing rules
   below. Short matters, because this file is sent with every request of every session
   for the life of the project.
4. **Choose the toolchain**, and work out what a runner without a route to the public
   internet would need, since that is the case for any project adopting Xeno on its own
   infrastructure. This is WP0 and it is the first package for a reason.
5. **Install `glab` on the machines that will do the work**, since issues are created
   from the session rather than from a browser, as described below.
6. **Take WP1 and WP7 core as the first piece of real work**, in that order, and do not
   start it in the same session as the bootstrap.

### Three standing rules for every session

**The documents are not editable by the agent.** Where the code and the specification
disagree, the specification wins until a person changes it, and a change to it is its
own commit made before the code that follows from it. This is principle P4 turned on
the construction itself. An agent that may quietly adjust the spec to match what it
wrote has removed the only control in the arrangement.

**No invented fields, gates, tools or rules.** Everything the artifacts carry is
enumerated in the process definition. An addition is a spec change first, by the rule
above. The same holds for the MCP tool surface and the gate list, both of which are
budgets rather than lists.

**Every change belongs to a work package and to an intent.** A branch carries one
intent, the commit message references its issue, and the issue carries the label of its
package. Where something needed does not belong to any package, that is a finding about
the plan and gets written down rather than absorbed.

### Tracking the work

**A work package is a label, an intent is an issue.** The distinction matters and is
easy to get wrong in the other direction. WP7 is weeks of work and many merges; an
intent in this process is one change through six phases ending in exactly one merge.
Equating a package with an intent
produces intents the model cannot carry, and it would do so in the one project that has
to demonstrate the model. Labels group, issues are the unit of work, and a package's
progress is a filtered issue list, which every host provides and epics do not.

**Issues are created just in time, from the session.** `glab issue create` from the
agent's shell, one or two ahead of the work and never sixty at once from the plan. The
plan already holds that content; a second copy in the tracker drifts from it within a
month. An issue nobody will touch this week is an intention that ages while it is
written.

The reason to create them from the session rather than in the web interface is not
convenience. The first step of working a package is turning its "Done when" sentence
into checkable statements, and those statements are what goes into the issue body. The
act of creating the issue is the act of writing the acceptance, not an administrative
step that follows it.

**`gh`, not an MCP server.** A GitHub MCP server exists and is the wrong trade here.
Every tool definition is sent with every request of every session whether it is called
or not, which is the standing cost WP20 describes for users of the tool and WP11 keeps
as a budget. For something done a few times a week, a shell command costs nothing
between uses. The token is a personal access token with API scope, in the environment,
never in the repository, which is the same rule `project.yaml` follows for the tracker
credential.

**Creating issues is not part of the adapter contract.** WP12 resolves an intent, reads
an issue, writes a comment and resolves credentials, and creating issues is
deliberately absent. It should not grow in later either: what a developer does with
`glab` is work on the project, not a step of the process, and an adapter that creates
issues is on its way to becoming a tracker front end.

**Branch and commit message carry the intent.** The branch is named after its issue,
which the host does by itself when the branch is created from the issue. The commit
subject is plain Conventional Commits and carries no issue reference:

```
feat(docs): a description
```

The reference lives in the footer, `Refs #123` on the commits of the change and
`Closes #123` on the one that finishes it, and the merge request description carries
the closing line as well.

**This rests on one host setting, and it has to be made before the convention is worth
anything.** Earlier revisions put the reference in the subject because a footer does not
survive a squash. That is true of the default and not of the host: both GitLab and
GitHub can build the squashed message from the merge request description, and then the
footer survives exactly as well as a subject does.

| Host | Setting |
|---|---|
| GitLab | *Settings, Merge requests, Squash commit message template*, set to include `%{description}` |
| GitHub | *Settings, General, Pull Requests, Squash merging*, set to *Pull request title and description* |

Without it the reference is lost at the merge and the trail loses the one link it exists
to keep, which is worse than any of the placements considered before. The setting is
therefore part of what `xeno init` prints for an administrator to arrange, alongside the
protected branch and the required pipeline.

A setting is not a commit, so no gate covers it and nothing notices it drifting back.
For this repository `scripts/github-settings.sh` applies it and prints it before and
after, which is the most a repository can do about its own configuration from inside
itself. It belongs with `xeno enforcement check`, which asks the host what it is
configured to do and is the general form of the same problem.

A merge request template carries the closing line, so that it is filled rather than
remembered.

**What the subject no longer has to do.** A code host reads a subject for closing
keywords, and `fix` is one of them, so a reference standing next to the type closed the
issue: `fix: #123 short slug` did, on this project's own issue #2, on the third of its
six commits while the work ran on for three more. A scope separated them again, so
`fix(gates): #123` did not, which was measured on #3 and is the worse property rather
than the better one, since it made the outcome depend on whether the author happened to
write a scope. In the footer the question does not arise: `Closes #123` closes because
somebody wrote it there, which is what deliberate means.

The `Xeno-Intent:` trailer stays available as a shipped example for projects whose
subject line is already spoken for. Xeno's own repository does not use it.

**Checked early, binding late.** A `prepare-commit-msg` hook from `examples/hooks/`
builds the message from the branch name, installed through `core.hooksPath` pointing at
a versioned directory. It calls `xeno check commit-message` rather than carrying the
expression, so the same pattern decides in the hook and at the gate. It is feedback,
not enforcement: `--no-verify` removes it, and nothing catches that before the gate, for
the reason given in WP10. Before M0 the convention lives in
`AGENTS.md` and nothing checks it at all, which is the honest state of it rather than a
gap to apologise for.

### How a package is worked

Start by turning the "Done when" sentence into checkable statements. Those statements
are the package's acceptance and the seed of the corpus in WP17, and writing them first
is the test first discipline the process definition asks of everyone else. A tool that
demands acceptance criteria before code and was not itself built that way would be hard
to defend and easy to notice.

Declarations before the code that reads them. Gate order, template structure, rule
format and the artifact schema are the shared truth that the runner, the documentation
generator and later the assessment all derive from. Code written before the declaration
tends to become the declaration by accident.

Golden files from the start. The artifacts are text and the gates are deterministic,
so a corpus of inputs with expected outputs is the cheapest verification available and
the one that survives refactoring. It also forces the normalisation questions, line
endings and path separators, to be answered in WP1 rather than discovered in WP17 on
another platform.

One package, one session where it fits. A resumed session carries its whole history
back into every request, which is the same cost lever WP20 describes for users of the
tool. Where a package does not fit, restarting from the branch and the written
acceptance is usually cheaper than continuing from a long session.

Load the part of the specification that governs the package, not the whole file. It is
well over a thousand lines, and sending all of it into every session for the sake of
one section is the exact habit the context profile exists to break.

### What to watch

The bottleneck is review, not production. An agent produces more code per hour than a
person can read carefully, and the packages are small enough to be reviewed in one
sitting only if they stay that way. Where a package's diff can no longer be read in one
pass, the package was too big and the plan should say so.

The second thing to watch is drift between the documents and what exists. The
specification is the control only for as long as it is true. At every milestone, the
gap between what the documents describe and what the code does gets closed in one
direction or the other, deliberately, and not left to be discovered by the next reader.

## 5. Dogfooding

From M0 onwards, Xeno is developed through Xeno, with the scope that exists at that
point and growing as each further package lands; from M1 against an effective rule
set. It is the only feedback that shows
whether P1 and P2 are cut too heavily for everyday use, and it shows it before any
breadth has been built on the assumption that they are not.

The side effect is worth having: the public repository then contains real artifacts,
which argues the case better than documentation does.

**What to watch, in this order.**

*Is the process too heavy for small changes?* Six phases with six files each for a
one line fix is out of proportion, this is the most likely place v1 needs sharpening,
and it is the question that holds back the trackers and repositories of 1.1. No
shortcut is defined, deliberately, because the shape of one should follow from
measurement rather than from anticipation. The implementation should
nonetheless keep it possible: phases are invoked individually, gates are addressed
per phase, and nothing assumes that every intent walks all six. A later fast path is
then a configuration of what exists, not a second process beside it.

*Token use per phase*, against the baseline from WP13, which is what WP20 is judged on;
WP15 is judged on bytes read, which the runner counts itself.

*Concurrency.* Two people on two intents in one repository. Artifacts sit in separate
directories, `rules/learned/` and templates do not. Git probably handles it; this is
to be observed rather than solved in advance.

## 6. Sequence

1. WP0
2. WP1, WP7 core: one hand written artifact is validated
3. WP2, WP3: a phase can be rendered and checked against its template
4. WP9, WP10: a repository can be initialised, gated and its host configuration checked
5. WP11: one agent first, the second immediately after, to expose harness leakage
   early
6. **M0**
7. WP4, WP5, WP8: the substance of the process
8. **M1**
9. WP12
10. WP6, WP13, WP14
11. WP15 and WP20, once WP13 provides a token baseline to measure against
12. WP19 last, before the release
13. WP16 and WP17 alongside throughout, not at the end

WP18 is specified but not built. It moves to 1.1 with the rest of the breadth.

**M0, the walking skeleton.** One intent, one phase, one agent, end to end with a
passing gate in CI, where green means the eight gates that exist passed and the other
five are reported as not implemented. Eight packages, and it proves the chain rather
than the process: agent, artifact, gate, pipeline, merge gate. That is where the
surprises live, and none of them are about rules. It is deliberately not a statement
about whether Xeno is any good, which is what M1 is for.

**M1, the process on one intent.** Rule engine, assumption register and context profile
in place, so an intent runs through all six phases under an effective rule set.
Dogfooding starts here. Evidence handling is not in it: a phase that declares no
evidence passes G-Evidence, and WP6 follows once there is something to attach.

**M1 does not wait for the tracker.** With the trigger on the CLI, an intent can be
carried end to end with the issue content copied in by hand. The adapter is comfort and
reach rather than function, so it sits behind the milestone and stays in the release,
for the reason given in WP12.

### 6.1 Size

No estimate in days, because there is nobody to estimate for and a number without a
basis outlives its own caveat. What follows is an ordering of magnitude, so that the
question whether v1 is a quarter or a year has an answer in this document rather than
somewhere else.

| Size | Packages |
|---|---|
| Large | WP3 templates, WP4 rule engine, WP7 runner, WP11 agent layer, WP15 symbol index, WP16 documentation, WP17 test strategy |
| Medium | WP0 bootstrap, WP1 artifact schema, WP2 rendering, WP6 evidence, WP8 context profile, WP10 CI wrapper and enforcement check, WP19 conformance, WP20 token economy |
| Small | WP5 assumptions, WP9 init, WP12 GitHub adapter, WP13 token recording, WP14 learning |
| Not sized | WP18 dashboard, specified and deferred to 1.1 |

The distribution says something the sequence hides. Of the twenty packages built for
v1, M0 needs eight and three of the seven large ones, M1 eleven and four. Neither is
an early milestone in the sense of a cheap one, and anyone reading the sequence alone
would take M1 for a third of the work when it is closer to half. What follows M1 is
measurement, tuning and writing rather than new substance.

The first honest figure comes from M0 itself, which is the earliest point at which an
estimate has anything behind it. WP13 then makes the token use measurable, and the
size classes above should be corrected against what the first phases actually took
rather than defended.

## 7. Risks

**The adapter contract hardens around GitHub.** One adapter built against one code
host is an adapter written to that host's shape, and nothing in v1 pushes back. The host
that shaped it has a canonical endpoint and issue keys tied to the repository, which are
the two things most likely to have been absorbed as if they were general. The
two countermeasures are cheap and in WP12: no assumed relationship between issue key
and repository, and endpoint plus authentication scheme as parameters rather than
constants. Neither is testable against a second system until 1.1, which is the
residual risk being accepted here knowingly.

**Dogfooding happens in one organisation on one plan.** Xeno's own repository is a
private one on the free plan, so the question whether the process is proportionate is
answered by one team under one set of constraints, and the four eyes requirement cannot
be exercised there at all, because the settings that would enforce it are not available
on that plan. The answers will be real and they will not be representative. The
countermeasure is to record the plan alongside every finding from the dogfooding stretch
rather than to correct for it.

**Finding ids are a new failure surface.** Decisions survive a re-run by id. Too
stable an id carries a release across a change that should have voided it, too
unstable an id makes every re-run lose the decisions. This needs a corpus in WP17 of
its own, not a test case.

**Staleness is the quiet failure.** G-Freshness is what keeps a redone phase from
silently invalidating everything after it. It is easy to postpone and expensive to
retrofit, because without it the trail looks complete while being wrong.

**Rule levels above the project are copies.** `provider` and `org` rules live in
every repository with no reconciliation. Divergence shows up in `rules_hash` after
the fact and not before, and the more repositories there are, the more likely it is.

**Session log formats are not contractual.** Token recording reads what the tools
happen to write, and that can change without notice. Treat WP13 as replaceable, which
is easier now that the money view belongs to the model access point and not to this
package.

**Hook trust models differ per client** and a freshly cloned repository will not run
hooks until trusted. This belongs in onboarding, not in the runner.

**Harness leakage.** The moment anything branches on `XENO_HARNESS`, tool
interchangeability is gone. Building against the second agent early is the
countermeasure.

## 8. Scope of record for 1.1

Not a plan. No sequence, no acceptance criteria, no sizing. It exists so that decisions
already taken are not rediscovered, and so that requirements raised outside this
document do not evaporate. It is revisited after the first dogfooding stretch, which is
what decides whether breadth is the right next move at all.

**Deferred by decision, recorded elsewhere in this document.** Hosts beyond GitHub,
each with its tracker adapter and its CI wrapper, GitLab next and Jira after it. GitLab
is the cheaper adapter and buys reach, since the shapes are close; Jira is the one that
tests the contract, because its project wide keys have no dependency on a code host,
which is the assumption a contract built against GitHub alone is most likely to have
absorbed. Taking the cheap one first is a decision about adoption, so the risk stays
open one release longer.

**The GitLab that is meant is the Enterprise variant.** It has the approval rules the
free edition does not, which is the difference that decides whether the four eyes
requirement has an enforcement point or is recorded as waived. An adapter written
against the free edition would be written against the weaker of the two and would have
to be extended rather than configured.

These are wrapper templates for other people's projects and not pipelines for this one.
Xeno is developed on GitHub and stays there, whatever hosts it learns to generate
wrappers for.

Intents spanning several repositories. Issue commands as a trigger, with whatever
component receives them. The dashboard, WP18, specified in section 2 and not built.
The German documentation translation together with the hash binding that keeps it
current.

**Public release, decided rather than deferred, and the repository went first.** The
repository is public from v1. What waits for 1.1 is the release: signatures, a public
distribution channel and the marketplace registration of the plugin. Three obligations
fell on v1 for it, and the first two are the kind that cannot be repaired afterwards,
which is why they were measured against the whole history before the visibility
changed rather than audited after.

*The history has to be publishable from the first commit.* Publishing a repository
publishes everything ever committed to it: internal hostnames, customer names in
examples or fixtures, anything mistaken for a test credential. Removing it later means
rewriting history that other people have already cloned. Nothing in the repository may
assume the infrastructure it is developed on, and that is a rule for v1 rather than a
cleanup task for 1.1.

*Xeno's own trail becomes public with it.* From M0 the project is developed through
Xeno, so `.xeno/intents/` fills with intents, assumptions, learning records and digests
of real sessions, with names in the `by` fields. Removing them at publication is not an
option, since the trail is the demonstration, so they are written from the start in the
knowledge that they will be read by strangers. A learning record naming a customer
cannot be unwritten later any more than a commit can.

This is not a Xeno peculiarity. Any project that publishes a repository publishes the
trail inside it, which is the point of principle 2 and worth saying out loud before a
project's first phase rather than after it. Released binaries carry none of it: a release
is the runner, the image and the plugin, and the plugin is templates, rules, skills and
the secret filter.

*Third party attribution has to be current.* The SBOM records what the release carries
and a `NOTICE` has to reflect it before anything is published. Tree-sitter grammars are
not among them, since none are shipped, which is one of the reasons not to ship them.
This one is repairable, but it is cheaper to keep current than to reconstruct.

**Raised by the adoption mapping written for the first deployment**, and recorded here
because a requirement that lives only in a downstream document is gone within a year:

- An extension block in the artifact schema, validated by G-Schema against a schema the
  project declares. Organisations carry attributes Xeno does not know, a project class
  or a risk tier among them, and today they have nowhere to sit inside the trail. This
  is the strongest of the three.
- A predicate type that checks a frontmatter value against a named set, which turns a
  catalogue of approved models or tools from a record into a rule. Decided for v1
  rather than 1.1: v2 depends on it three times over, for the approved model set, for
  the approval catalogue covering skills and agent definitions, and for the list of
  permitted MCP servers. It is cheap where it stands and a dependency if it moves.
- Risk dependent gates, already in the outlook of the process definition, given a
  concrete shape along project classes so that v2 is designed against a stated need
  rather than an assumption.

**Raised by evidence that carries a threshold.** A predicate type for "evidence of kind
X with a result is present", so that a project can require that coverage or a scan was
checked against some threshold without Xeno ever knowing the number. It would be the
first predicate to judge evidence rather than artifact text, which is why it waits for a
project that actually asks for it rather than being built on speculation.

Decided for v2, and a regulated environment is the project that asks. It stays exactly
as narrow as it is written above: the evidence item carries the job's own result, a rule
can require that a check of kind X ran and passed, and the threshold itself stays in the
scanner's quality gate, where the pipeline enforces it by failing. Xeno never reads the
report, so there is no normalised format to require, no baseline to store and no
suppression to manage. What it does not resolve: the threshold lives in pipeline
configuration, outside the trail.

**Raised by the change management side of the standards.** A release record: a command
over a commit range or a tag that reads the intent keys from the commit messages, looks
up each intent's verdict, and writes what it found as an artifact. It answers what a
release note cannot: which intents this release contains, whether one of them was
overridden, whether an obligation is still open, and whether a commit belongs to no
intent at all. The last question is the one an auditor asks, and today nothing would
notice.

It is the counterpart to the `release-notes` section rather than a replacement. The
section is one intent describing its own change, written during the work; the record
looks from a release back at the intents inside it, after the merge, which is where
change management under ISO 27001 and DORA actually starts.

Deployment stays out: environments, targets and credentials would make Xeno a deployment
tool, and it has no business knowing what "production" is.

Decided for v2 rather than 1.1, and it is what carries the handover at that boundary.
Separating development, test and production is a property of infrastructure and no
lifecycle tool can attest it; what this record can show is that what reached production
is what passed the gates. The record is sealed when the release is cut, so a later
deployment reference, stating that a release arrived in an environment at a time, is its
own append only file pointing at the record by hash rather than an edit to it. Xeno does
not deploy. It records that somebody else did.

The intent keys the record reads are put there by the project's commit message rule and
enforced at the host, which is the normal case and the precondition. What the record adds
is finding where that did not hold: commits from before the rule was introduced, an
import from another repository, a force push, an administrator with a bypass right. It
does not prove the ordinary case, it finds the exceptions, and that is what it is for.
Whether the host can enforce the rule at all is an edition question and belongs to
project setup rather than to this record.

**Raised by the quality of what a phase asks.** Five questions in the guidance text,
none of them a section, so the section budget stays untouched. P0 asks what does not
belong to this, since non goals almost never survive being written once and P0 is when
somebody still knows what they excluded. P2 asks whether this exists already and why it
is not being extended there, which in a brownfield system prevents more duplicated work
than any rule and which an agent with the symbol index can answer with evidence rather
than merely raise. P2 also asks which components are touched in which order, the answer
going into `impact`. P4 asks for the counter evidence, meaning not whether the acceptance
criteria are met but what a test would have to show for them not to be. P5 asks which
assumption this intent leaves behind for the next one, `residual-risk` being about the
risk rather than about that.

Held back because it is text in strings bundles that can arrive at any time, and because
what makes such questions good is knowing how they land, which dogfooding answers.

**Raised by decisions that keep repeating.** A second learning mode, run once per intent,
reading the `decisions` of earlier intents and reporting repetitions as
`project-convention` with a rule as its proposal. The threshold belongs in the process
definition rather than in a prompt, because a number in a prompt changes with every model
version: three identical decisions are a convention, two are coincidence. The third
answer to the same question is not a decision any more, it is a convention nobody wrote
down, and once it is a given rule the question stops being asked. It stays a proposal
with human approval, which is the difference from silently reusing an earlier answer.

Held back because it finds nothing before the third intent. It also needs read access
across intent boundaries, which nothing else has, and it would need the only fixture in
the corpus spanning three finished intents.

**Raised by requirements work.** A shipped `checked` rule over assumptions that were
never confirmed. The register is the strongest requirements instrument in the process
and the interesting number is recorded nowhere: how many assumptions from P0 and P1
survived to P5 without anybody confirming them. An assumption that carries a whole
intent unconfirmed is exactly the requirements risk that returns later as rework, and
making it visible is a rule rather than a new mechanism.

The binding of acceptance criteria upwards, to a regulatory requirement rather than
only downwards to a test, is the extension block above rather than a separate item. In
a regulated environment it is the difference between stating that something is tested
and showing that a named control is.

**Raised by the question of proportionality**, which is the one this section cannot
answer and dogfooding can.

Settled in advance, because settling it afterwards is worthless: the figures decide. From
M0 Xeno builds itself, every change to the runner runs through six phases, and the rework
rate becomes a real number about this project rather than an argument about other
people's. If it says that P1 and P2 contribute nothing on small changes, that is the
evidence for the light path, and an uncomfortable sentence about one's own product at the
same time. Agreeing now to follow the number removes the temptation to reinterpret it
later.

Today every change runs six phases. A typo in an error message therefore costs the
same process as a new payment interface. If that holds, the predictable thing happens:
people stop creating intents for small work, work past the process, and the trail
acquires holes nobody notices. A process that gets bypassed is worse than one that asks
for less.

The candidate answer is risk dependent phases rather than only risk dependent gates: a
project class or a change class determines which phases carry artifacts of their own, so
that a light path can fold P1 and P2 into one. The property to preserve is that this is
less paper for the same evidence and not less evidence. Intent, verdict, cost and
binding stay; the number of artifacts shrinks.

What decides it is measurable without anybody estimating: cost per intent, artifacts per
intent, and reopens per intent. Intents that are too large are the cause of most
expensive phases, and they are visible afterwards.

The same data answers the project management questions without Xeno managing anything.
Cycle time per phase from the timestamps, rework rate from the reopens, open obligations
per release from the release record, cost per intent from `cost.yaml`. These are reports
in the dashboard, not artifacts. The boundary holds in both directions: derive, never
manage. The moment Xeno keeps estimates, assignments or status, it is a second truth
beside the tracker, which is the mistake avoided everywhere else in this document. Of
the derived figures the rework rate is the one that earns its place, because it alone
answers whether the process produces quality at the front or merely effort at the back.

**Raised by the phase model.** A predicate over the commit range that checks whether
the test files touched by a change were committed before the implementation files they
cover. It is the only way to make a test first discipline provable without watching
somebody type, and it closes the one real gap between the phase order and an
organisation that requires tests before code. v1 states in the process definition that
P4 proves the verification rather than being when tests get written; this makes the
statement checkable for those who need it to be.

## 9. Open

To verify rather than assume, each before the code that relies on it: whether the client
accepts the instance URL for the marketplace --> unsure right now

whether the gateway reports the model that
actually served a request; whether it accepts and records request metadata --> both yes,
measured against LiteLLM 1.102.1 (#56). The deployment that served a request is in the
response headers, `x-litellm-model-name` with `x-litellm-model-id` and
`x-litellm-model-api-base`, and in the spend log row, where `model` is the served name
and `model_group` the name that was asked for. The response body carries only the name
that was asked for. What any of them names is the deployment that was routed to and not
what the provider says it served, so a recorded `model` is a routing record rather than
an attestation from the provider. Metadata arrives through headers alone, which is what
a harness can set, and the two routes that carry it are not equal. `x-litellm-tags`
lands in `request_tags`, is aggregated per tag with a count and a spend, and is
documented as gated by nothing; it is the attribution below the project and it is what
Xeno sends. `x-litellm-spend-logs-metadata` keeps a key value pair verbatim on the row
and is documented as requiring an enterprise licence, yet the proxy it was measured on
has no licence set and recorded it anyway, which is a worse foundation than a paid
feature would be: a release that enforces the gate stops the rows without an error. Per
key and per user figures need no licence either, through the two spend list endpoints
`/spend/keys` and `/spend/users` and through `/user/info` for one user, but they
aggregate above the phase and answer what a project spent rather than what a phase did.
The two list endpoints sit in `spend_tracking_routes`, so an internal user reaches them
and not only an administrator, and what each caller gets is scoped: its own rows, an
empty list where its key carries no `user_id`, and 403 rather than a filtered list where
it asks after another user. A virtual key restricted to `llm_api_routes`, which is what
a phase runs under, is refused every one of these routes outright rather than handed an
empty list, and it sets both metadata routes on its own requests all the same: a request
through one was recorded with its tags, its key value pairs and the key's own alias on
the row. The finer scoping, the empty list for a key with no `user_id`, is taken from
the documentation. The names are not free on either route. A custom header is dropped
without an error unless the proxy lists it under `extra_spend_tag_headers`, so either
the harness sends LiteLLM's own names or the proxy configuration becomes part of what an
adopter arranges.

whether each
harness forwards custom headers to it, and under which variable --> claude-code: export ANTHROPIC_CUSTOM_HEADERS="X-Custom-Header-1: value1\nX-Custom-Header-2: value2"; codex: X-Custom-Header = "MeinWert"

whether the job token CI
provides can read protected branch settings --> github via api/ cli or ui; gitlab https://docs.gitlab.com/ci/pipelines/merge_request_pipelines/#control-access-to-protected-variables-and-runners

Decided later on purpose: which of content, packaging and wiring WP11 builds per harness,
settled in dogfooding rather than in advance.

The namespace and repository name are chosen and deliberately not recorded here. The core
built by hand before M0 uses a placeholder module path, and its other choices are recorded
in its `ASSUMPTIONS.md`.
