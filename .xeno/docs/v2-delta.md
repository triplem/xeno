---
id: v2-delta
title: Xeno v2, Delta against Process Definition v1
revision: 2
status: draft, not ratified
date: 2026-09-20
location: .xeno/docs/v2-delta.md
---

# Xeno v2, Delta against Process Definition v1

## 0. What this document is

v2 makes Xeno multi-agent and moves the execution of a phase onto an agent
platform. This document records what that changes in the process definition,
section by section, so that the next revision of that document can be written
from a list rather than from memory. It records decisions, not their
implementation.

The platform decision itself and the alternatives rejected are in
[[orchestrator-evaluation]].

The seam every decision below was measured against: **OpenHands owns the work,
Xeno owns the record, and every crossing of that seam leaves an entry in the
repository.** Where v1 contradicted that seam, v1 changes.

## 1. What does not change

The parts that carry the tool are untouched, and it is worth saying so before
the list of changes suggests otherwise.

Every artifact belongs to exactly one intent. The repository is the source of
truth. Gates are deterministic and model free. It stays reconstructable which
information a version was built on. Learning is mandatory and versioned. The
precedence of given over learned rules. One intent, one repository. The
network free gate path, and the architecture rule that keeps it network free by
making the adapters unimportable from it.

## 2. Principles, section 2

**Principle 2 is extended, not replaced.** The trail is exactly as public as the
repository. The raw event stream is not part of the trail and is subject to the
access rules of the service that holds it. Without this addition the sentence
becomes untrue the moment a central event store exists, and it is a sentence
much of the rest of the document rests on.

**Principle 3 gains a consequence.** No justification text attached to a
decision without a confirming human. This held implicitly in v1 because only
people typed. With an orchestrator that can propose a reopen or an abandon, it
has to be written down. It applies to approve, override, reopen and intent
close. Learning records are observations rather than decisions, stay agent
written and carry `source: agent`; the closing record of an abandoned intent is
the exception, because abandoning is a decision.

**Principle 4 gains a boundary, not an exception.** Confirmation policy and
security analyzers, including the LLM based one, control execution before a
tool call. They do not judge work and produce no verdict. They are therefore
not gates, and principle 4 is intact.

## 3. Phase model, section 6

**Order is enforced by the runner.** `xeno phase start` refuses a phase whose
predecessor has no completed `gate.yaml` with a matching `artifacts_hash`. In
v1 the order was guaranteed by a person typing the commands. With an
orchestrator it would otherwise depend on a model.

**Completion is barred by a stop hook.** The OpenHands stop hook calls
`xeno phase finish`, and its result decides. No agent can declare a phase
finished whose gate is red, and the verdict stays with the runner.

**Hold points at every phase boundary, approval required after P1, P2 and P5**,
configurable per project. Approval where a person decides something no gate can
check: what is built, how it is cut, whether it ships. P0, P3 and P4 are
covered by gates, tests and evidence. A red gate always stops, independently of
the configuration. This list is the natural attachment point for risk dependent
gates: it becomes a property of the project class.

**A hold point ends the run.** The state is in the repository, the conversation
is not continued. Blocking would hold a sandbox and a context window open
overnight and would move process state into the memory of a service.

**Resumption computes rather than remembers.** `xeno intent status` is a pure
function over the intent directory, read on the intent's own branch, since that
is where its phases are written; the dashboard reads the same way and a merged
intent is read from the default branch. There is no stored position that could go
stale, and therefore no reconciliation between two systems.

**A phase restarts with a fresh conversation.** The orchestrator is rebuilt from
`intent.yaml`, the status, the approval record and the open assumptions; the
phase takes its context solely from `context.lock.yaml`. The rule behind it:
the orchestrator may pass nothing into a phase that is not in the repository.
It disposes, it carries no content. An interrupted phase is discarded and
restarted rather than continued, because a continued conversation is exactly
the state in which artifact and record diverge. On resumption the lock file is
checked against the tree; a divergence writes a new revision, which is what
G-Freshness is for.

**A phase has one phase agent that may delegate.** The phase agent owns the
sub-conversation, the context and the authorship of the artifact. Delegation
creates task conversations with their own context whose result returns as text.
Delegation depth is two: orchestrator, phase agent, specialist. Parallel work
inside a phase is deferred, because it requires separate clones and a merge rule,
which is a separate mechanism with separate failure modes.

**A resolved question is checked against the whole context.** Once a question has been
resolved, a specialist reads the answer against the specification, the design, the
assumption register and the other decisions, and reports contradictions. The result goes
into open questions or the review checklist, never into a verdict, because otherwise an
AI would end up judging a person's answer wrong, and principle 4 says no.

The value lies less in the answer than in its consequences: a decision taken in P3 can
refute an assumption from P1, and nobody notices. That is what such a pass finds.

In high project classes the lenses run over the answer as well, which is cheap now that
they are sub-agents with their own context and their own briefs. Repetition is not the
same thing: asking one model twice mostly yields the same answer at twice the price,
which is one prejudice twice rather than a second opinion. Where a second instance is
genuinely wanted, it uses a different model.

How deep the checking goes follows the project class and the phase, never a
self-assessment. An agent that rates its own question low risk has an interest in doing
so. Where finer control is wanted, it hangs off something objective, such as whether the
resolution touches an interface, a migration or anything security related, which is
checkable in a way that a free text risk level is not.

One specialist is worth naming rather than leaving to a project to discover. In
P1 an adversarial reader, whose only brief is to find ways of satisfying the
specification while missing the intent, is the delegation that pays back
fastest. People can do this and rarely do, because it is thankless work. Its
output goes into open questions and non goals, never into a verdict, and non
goals are the line in P1 that almost never survives being written once.

**Lenses become sub-agents.** v1 made them skills explicitly because subagents
were client specific and outside Agent Plugins 1.0.0. That reason is gone. They
still write no artifact of their own and produce no verdict, and a checklist
entry from a lens still carries `source: lens` and no rule id. As a side effect
a lens now pays for its context once instead of riding in every request of the
phase.

**`xeno phase reopen --phase <n> --reason` is new.** A return to an earlier
phase is allowed and requires human approval. The following phases become
outdated by that act, and nothing is written into them to say so: their
recorded input hash no longer matches what their predecessor now holds, which
is what makes them outdated and what `xeno intent status` reports. Without the
command an orchestrator would simply rewrite artifacts and the verdicts of
later phases would age unnoticed.

**The orchestrator may propose an abandon, not perform it.** The abandon is the
only act that ends without a result, and v1 rightly values its learning record
above the ones written while building. A model allowed to give up has a cheap
exit from every difficult phase.

**A reason has two parts.** An anchor that exists in the repository, being a
finding id, an unmet acceptance criterion or a refuted assumption, and a
sentence. No reopen without an anchor. The sentence may be drafted by an agent
and counts only once confirmed: `proposed_by` carries the agent identity,
`stated_by` the person. This is the assumption register's mechanism applied to
a decision about the state of the work, which is what it structurally is.

**Order of work is taken from the tracker, and priority stays there.** When several
intents are open and capacity is short, the orchestrator has to choose which run to
start next, and it asks the tracker rather than keeping a ranking of its own. Priority is
a management decision, it changes more often than any artifact, and it sits next to
everything else that does not run through Xeno. In `intent.yaml` it would be wrong the
day after the commit and a change to it would have to pass through the very process it
is meant to steer. Reprioritising in the tracker therefore takes effect immediately and
without a commit, and Xeno records only the order in which work actually happened, which
follows from the timestamps anyway.

What Xeno does contribute to prioritisation is the part it usually lacks, namely real
figures: cost per intent, cycle time per phase, rework rate, open obligations per
release. Whoever prioritises today estimates effort; with those numbers the effort is
known for comparable intents.

**Concurrency is carried by branches, exclusivity by a run marker.** Each intent
gets its own branch at the start of P0 rather than when P3 needs a merge
request, so two intents work on different branches and cannot overwrite each
other's intent directory. Within an intent at most one phase runs at a time,
because a phase whose predecessor is still writing has no defined input. A run
marker in the intent directory, set by `xeno phase start` and removed by
`xeno phase finish`, carries the run id and a timestamp and can be broken by a
command when a run has died. It lives in the repository rather than in the
service, because it would otherwise be the one piece of process state that does
not survive a restart of the service, and it sits outside `artifacts_hash`,
since a lock that changed the verdict it protects would be absurd. A long lived
intent branch makes G-Freshness busier, which is what G-Freshness is for.

**Two limits per phase, iterations and tokens.** Set in `project.yaml` and
differentiated per phase, because P3 legitimately consumes a multiple of P1.
Two rather than one, because they catch different failures: the loop, and the
agent that reads the same large file again and again. Both count cumulatively per
phase across attempts, not per attempt, or a restart would reset them and the
limit would bind nothing; the number of attempts is capped as well. Reaching a
limit ends the run, and no `output.md` and no `gate.yaml` are written, so the
existing rule for an interrupted phase applies and the phase is discarded and
restarted rather than continued. The abort is surfaced in the tracker and goes into
`learning.yaml`. A phase that hits the same limit twice says something about the
cut of the intent or the choice of model, and that is what the learning
register is for. How often an infrastructure failure is retried is an
operational parameter, not a process decision.

## 4. Artifacts, sections 4 and 5

**What is sealed is never rewritten.** Later knowledge about a finished phase is
either a separate file that references it, or a comparison that is computed. It
is never an edit to an artifact whose hash already carries a verdict. Without
this rule, every "mark it as outdated" in this document would quietly change
the very hashes the verdict rests on. With it, the answer is usually that
nothing is written at all: a phase is outdated when the current
`artifacts_hash` of its predecessor is no longer the one it recorded as input,
and that is a function over data already present, surfaced by
`xeno intent status` and the dashboard. This is the same shape as rule drift in
v1, and the same shape as computing the intent status and the cost report
rather than storing them.

**`run.yaml` is new**, one per phase directory. It records how the artifact came
about: mode, conversation and task identifiers, the delegation tree, the hash
of each agent definition, requested and served model, the skills and MCP servers that were
actually used, the aggregated action record, and the hash of the raw event
stream. It lies directly in the phase directory and
is therefore inside `artifacts_hash` without any new rule, unlike `gate.yaml`
and `cost.yaml`, which the runner writes itself. It is mandatory in both modes.
In manual mode it holds little beyond the mode and the model, and that is the
point: an almost empty `run.yaml` states that nobody can say more about how this
phase came about, which is a different claim from the file being missing.

It carries two named sections, because it answers two different questions.
`provenance` is what went in: platform, orchestrator, agents, their definitions
and models, the skills and servers that were used. `execution` is what came
out: how the run ended, how many actions there were, which ones a person
confirmed or rejected, and the hash of the raw stream. One file rather than two,
since both would land in the same `artifacts_hash` anyway and a second file buys
nothing but a second name. The schema is in appendix A.

It is the largest schema addition of v2 and the place where principle 5 becomes
redeemable under multi-agent conditions at all.

**Sub-agent definitions are artifacts.** Shipped ones under `.xeno/plugin/`,
project ones under `.xeno/config/`, resolved by id like templates and rules,
hashed by G-Supply. Who worked with what brief is part of the information base.

**The orchestrator is one of them.** Its prompt determines the disposition, so it
belongs in the repository on the same terms as the agents it drives, as a file
based agent definition with its hash in `run.yaml`. Where a platform cannot take
its definition from the repository, the fallback is a reference to another
repository pinned to a fixed version, never an untracked setting in a service.
What stays outside either way is the deployment side, being endpoint, credentials
and sandbox image: what the orchestrator does is repository, where it runs is
configuration.

**The release record is new**, and it is decided for v2 rather than 1.1. It names which
intents are in a release, which of them were overridden, which obligations are still
open, and which commits belong to no intent at all. It is sealed when the release is cut
and never edited afterwards, so anything learned later points at it by hash, per the
rule at the head of this section.

**Who commits what, over an intent.** Two writing identities, and only one of them can
create an artifact.

| Phase | Commit | By | Touches |
|---|---|---|---|
| P0 | after `phase finish` | the run | intent directory, artifact, `gate.yaml`, `cost.yaml`, `run.yaml` |
| P1 | after `phase finish` | the run | artifacts |
| | after approval in the tracker | the runner | the approval record only |
| P2 | after `phase finish` | the run | artifacts, attaching what arrived for P1 |
| | after approval | the runner | the approval record only |
| P3 | after `phase finish` | the run | code and artifacts |
| P4 | after `phase finish` | the run | artifacts, attaching what arrived for P3 |
| | again if there are findings | the run | artifacts |
| P5 | after `phase finish` | the run | review result and release notes, attaching what arrived for P4 |
| | after approval | the runner | the approval record only |
| | merge | a person or the platform | one commit on the default branch |

A run commit carries the responsible human as author, the service account as committer,
the agent and its model in `Co-authored-by`, and the run id in a trailer. A runner commit
for an approval carries the identity authenticated in the tracker as author and the
service account as committer, and touches nothing inside the `artifacts_hash`. CI commits
nothing at all.

The number of commits therefore says nothing about the number of responsible parties. An
intent has seven to ten commits, two kinds of writer, and exactly one kind that can
produce an artifact. A squash collapses all of it into one commit on the default branch,
which is harmless because the record lives in the files and not in the history.

**An intent never spans the version boundary.** v2 is adopted after a release, with
every intent closed and every phase finished. There is no migration of in flight
intents, no artifact retrofitted with a `run.yaml`, and no phase whose earlier half
was judged under a different schema. The rule costs a release cycle and removes an
entire class of questions.

**The run id is the clamp between the two integration surfaces.** The orchestrator port
starts runs and the MCP server takes writes, and nothing otherwise connects them, so an
agent started for one intent can write into another's directory. The write path does not
help, because it only checks that a write goes through it and not who it is for. The
runner issues the run id at `phase start` and passes it into the run, the MCP server is
started for exactly one intent and phase, and it refuses any write that does not belong
to them. This is the only place in the design where trust between orchestrator and agent
would otherwise be implicit, and the failure it prevents is the embarrassing one: an
artifact that is formally correct, passes every gate and belongs to the wrong intent.

**Every artifact is written through the MCP operation**, whoever calls it.
Intent binding, schema and anchors hang on one code location, and anything
written into the phase directory past it becomes a G-Schema finding rather than
a silent error. This is also why manual mode is not a second code path.

**`agent.tool` disappears** and is replaced by an `orchestrator` block with a
model per agent. The `tool` and `model` fields in artifacts remain, fed from
the run rather than from configuration.

**The event adapter is the brittlest code in v2, and it is treated as such.** It parses
a foreign schema, in a product mid architecture change, and everything in `provenance`
exists only through that interpretation. The danger is not a change but a silent one: a
renamed event is not a crash, it is an omission, and the result is a `run.yaml` whose
rejected actions section is empty because nothing was recognised rather than because
nothing was rejected. That file looks complete, passes every gate and states something
false about a run, which is worse than a missing file because a missing file deceives
nobody.

Two consequences. Knowledge of the foreign schema sits in one adapter behind the
orchestrator port and does not seep into the code that writes `run.yaml`. And an unknown
event type produces a finding in the phase rather than a silent skip, with the adapter's
schema version recorded in `run.yaml` so that it stays answerable later which runs were
interpreted under which reading. Not an abort: the work is done and should not be lost.

**A replay adapter ships as a test double.** The strength of v1 is that the gate path is
testable offline; the orchestrator port is not, and without a substitute every test hangs
off a running platform, which is exactly the code that would then never get tests, namely
`run.yaml` generation, action aggregation and cost splitting. Recordings arise from every
real run anyway and only need keeping. Keeping one per supported platform version turns
the previous paragraph's problem into a test failure rather than a trail failure.

## 5. Evidence, logging, systems, section 12

**Evidence is the repository, exclusively.** The event stream is operations. It
is condensed per phase into the filtered digest, and the artifact carries the
sha256 of the raw stream with its identifiers and time window. Whoever needs
the raw stream fetches it and checks the hash against the repository. In manual
mode the field is absent rather than empty.

**Run state goes into a database in v2.1, never into commits.** Everything that arises
while a run is going on and then changes, namely runs and their outcome, attempts,
action events, cost, evidence as it arrives and pending approvals, is operating state
rather than record. Git is the wrong tool for it: every write is a commit, every commit
starts a pipeline, and every concurrent write is a rebase. That single cause is behind
most of the awkwardness in this section.

The repository keeps what carries or explains a verdict: artifacts, `gate.yaml`,
approval records, digests, and a condensed provenance record. Exactly one operation
crosses from one to the other, the closing of a state, and it writes only the condensed
form together with a hash over the raw material. A second path would mean two truths,
and then the database is a regression rather than an improvement. The rule that keeps it
honest is blunt: if an auditor could ask about it, it does not live in the database.

What this removes rather than solves: `run.yaml` no longer has to be written from the
sandbox mid run, an attempt counter survives a discarded phase because it was never in
the repository, an arriving report needs no commit from outside, and a provisional
verdict disappears because the repository only receives a verdict once there is one.

Not in v2.0. On the local tier one run at a time in one process needs no more than a file
beside the run. v2.1 is where concurrency, overnight hold points and several intents at
once actually arrive, and it is very likely the same system as the event store below.

**Systems required for the central tier**, which is one of three, the others needing a
container runtime and the gateway or nothing at all beyond the runner: Agent Server with
a sandbox runtime,
the LiteLLM proxy as the gateway, a conversation and event store with a
retention period, a registry for runner and sandbox images, a secret store for
the sandbox, and the automation server for event triggers. An OpenTelemetry
collector with a backend is recommended and is explicitly operations, not
evidence. The lower tiers need correspondingly less, down to nothing beyond
runner and gateway.

**Two retention clocks.** Repository evidence lives as long as the repository.
The raw stream has three cases, because the tiers put it in different places. In manual
mode there is no stream and the thirty days under `.xeno/local/` stay as they are. On
the local tier the stream lies beside the run and keeps the same thirty days. On the
central tier it goes to the store with a period per project class, twelve months by
default, set in `project.yaml`. The hash in `run.yaml` is the same anchor in all three,
which is what lets the tiers differ without the trail differing.

**Token figures stop being self reported, from two sources.** The runtime
supplies the split, since a phase is one sub-conversation and each sub-agent
has its own usage identifier. The gateway supplies the total per project, which
makes the split checkable rather than merely plausible: the sum of a period's
phases against the gateway's total, the difference being manual work.
`evidence` takes three values accordingly, `runtime` for a split, `gateway` for
a total, `self-reported` only where neither exists. `cost.yaml` is still written
by the runner at `xeno phase finish`, now from metrics fetched over the
orchestrator port rather than from parsed session logs, and it stays outside
`artifacts_hash` for the same reason as before. It carries one line per agent,
keyed by the identifiers in `run.yaml`, split into input, output and cached
share, because one combined number hides the effect of every lever. The intent
level view is `xeno cost report --intent`, computed from those files rather
than stored as a second truth.

**Every model request carries correlation metadata.** Intent key, phase id, run
id, agent name and mode travel with the request, so that the same identifiers
appear in the repository, in the agent server's events and in the gateway's
logs. Three consequences follow, and the third is the one that makes it worth
doing. The gateway can then attribute a request to a phase and an agent rather
than only to a project, which removes the reconciliation above. The same
identifiers go on the OpenTelemetry span attributes, so an operational trace
lines up with an artifact. And the gateway's log becomes an independent record
of what ran, held by a system nobody working a phase can write to.

The metadata is asserted by the caller, so it correlates, it does not attest. It
carries identifiers only: no prompt fragments, no file paths, no free text, and
nothing that would put a customer name into a central log with its own
retention.

**Action control.** `ConfirmRisky` with an ensemble of the deterministic
analyzers and the LLM analyzer, threshold HIGH, MEDIUM in sandboxes with
network access. Recorded per phase in aggregate: policy, analyzer
configuration, counts of actions, confirmations and rejections, plus every
rejected and every HIGH action in full with reason and time. Every rejection
also goes into `learning.yaml`, because it is the only signal in which a person
contradicts an agent before the damage.

**The role follows the command, not a switch.** Phase work runs in the sandbox: `phase
start`, `phase finish`, the gate run, writing artifacts. Disposition runs where the
orchestrator is: `intent status`, choosing the next phase, starting a run, taking an
approval. One binary with two roles rather than two binaries, or the gate path would
have two places it could live in.

The consequence worth stating, because the convenient design is the other one: the
sandbox clones the repository and pushes itself, and shares no filesystem with the
orchestrator. Handing a working directory in works locally and falls apart centrally.
This is also where three earlier decisions meet, namely that the sandbox needs push
rights on the intent branch, that status is computed rather than passed, and that the
orchestrator carries no content.

**Xeno requires a state, not a packaging.** The stop hook calls `phase finish`, so the
runner has to exist inside the sandbox. How it gets there is the project's business:
compiled into its own image, pulled from an internal artifact store, distributed by
configuration management, mounted, or simply present in a process sandbox. Xeno
publishes a binary and a checksum and prescribes no delivery form, because hardened
environments reject a copy from a foreign registry and because a sandbox is not
necessarily a container in the first place.

What is prescribed is the check: before the first command the orchestrator asks the
runner for its version and aborts if it does not satisfy the pin. That one sentence
replaces the whole packaging question and holds for a container, a VM and a process
alike. A reference image ships as an example for the entry tier, explicitly neither
hardened nor a prerequisite.

The pin is a floor, not a range, and this replaces v1, which pins the runner exactly
through `runner_version` and refuses on any difference. Compatibility needs a minimum, reproducibility needs a
record, and the record already exists in the artifact. An upper bound would add nothing
that `schema_version` does not already cover, since a runner that cannot read a schema
refuses anyway, and it would force a commit in every project on every minor release. The
price is that two projects can judge the same tree under different runner versions,
which is acceptable exactly because the artifact says which one, the same logic as rule
drift. Where that is not enough, a rule can require an exact match against an approved
catalogue entry, through the same predicate type as models and skills.

**The sandbox holds the minimum, and holds it briefly.** It gets a gateway token,
read access to the repository, push rights to the intent branch alone, and egress to
an allowlist covering the package registries and the gateway. Everything else is
denied by default, which makes model access the only route out that can carry
anything, and that route runs through the proxy with a log.

Credentials are issued per run and expire with it. A long lived service account secret
inside a sandbox driven by a model is the kind of credential that turns up later in a
postmortem, and short lived issuance is the one difference that matters when it does.

The agent holds no tracker token. Comments, approval requests and status go through the
runner's tracker port, which already exists, and the agent reads the issue because the
runner placed it in the context at the start. That removes a credential from the sandbox
entirely and makes the trail more consistent, since the component that writes artifacts
is also the one that writes to the tracker.

Push rights end at the intent branch. The default branch is protected and unreachable
from the sandbox, so the blast radius of a compromised run is one branch that a person
has to approve before any of it moves.

Secrets are injected from the secret store as environment and never written into the
workspace, and they are masked in the event stream, which is the dangerous place because
it lives for months. `run.yaml` records which secret names were available and never a
value. G-Secret covers the repository and does not help here.

## 6. Triggers, section 12

**The CLI is no longer the only trigger.** Events may create an intent, start
P0 and answer an approval. Everything else stays a deliberate command.

**A run is triggered by an authorised act, not by content.** Creating an issue
or commenting on one starts nothing. What starts a run is an act only somebody
holding the rights for it can perform in the tracker, a label being the obvious
form. This is the same hardening OpenHands applied to its own workflows after
working through the injection chain, and it closes the path on which a stranger
starts a run. The run reads the issue as of the moment it was triggered rather
than a thread that keeps developing while the agent works, which
`context.lock.yaml` requires anyway. Issue content enters the context as
delimited input and never as instruction. None of this is protection in itself.
The protection is the sandbox, the action control and the gate; this narrows
the surface they have to cover.

**An incoming event is authenticated before it is read.** The decision that only an
authorised act starts a run means nothing if the endpoint receiving the event accepts
whatever reaches it: anybody who learns the address could start a run and the whole
decision would be bypassed. The automation server verifies the webhook secret, or the
signature where the tracker provides one, and discards anything that fails before
looking at its content. The same holds between the services: orchestrator to agent
server, runner to tracker port and agent server to gateway each authenticate, and each
of those credentials joins the table in [[implementation-plan]] with its own row.

**No agent ever runs in CI.** CI stays read only and recomputes. The reason is
not only the v1 argument about commits written by pipelines. It is the attack
chain OpenHands had to harden its own workflows against: prompt injection,
execution in CI, cache poisoning, theft of publish secrets from a more
privileged workflow.

**A phase approval is given in the tracker**, on the issue or the merge request.
That is where the people are who do not work from a command line, the act is
authenticated there and appears in the host's audit log. This covers the phase
boundary only. Decisions about findings, so approving one or overriding it,
stay commands as they are in v1, and hold points exist only in orchestrated
mode: in manual mode there is nothing to hold and nothing to release. `by` carries the
identity the adapter delivered; the commit is made by the service account and
says so. Limitation 12 shrinks rather than grows, at the price that part of the
record sits outside the repository, which belongs in the limitations.

Approval is a port, like tracker and host settings, so that the 2.1 move to a
graphical surface is an adapter rather than a change to the process.

## 7. Distribution, section 13

Dropped: `plugin.json` per Agent Plugins 1.0.0, the shipped hook wiring for
Claude Code and Codex, and session log parsing per harness. These are the three
expensive pieces of supporting two harnesses.

Added: packaging for the OpenHands plugin format, file based sub-agent
definitions, and G-Supply coverage for those definitions.

Unchanged: vendoring rather than remote fetch, the shared version number of
plugin and runner, the runner binary as the integrity anchor, and internal
distribution until publication.

## 8. Catalogue and rules, sections 9 and 14

**One catalogue for four kinds.** Skills, sub-agent definitions, MCP servers and the
execution environment, the last as a digest where there are images and as an identifier
of the build path where there are not
, share one approval catalogue, each with a hash and a state: approved, in trial
and restricted to project classes without production exposure, or revoked.
Checked by the predicate type "value against a named set", which moves from the
1.1 list into v1, since three decisions here depend on it and it is cheap where
it stands. Everything vendored is additionally covered by G-Supply. The
catalogue is mirrored into the repository like the model set and for the same
reason: the gate path reads it without a network.

**A second predicate type asks for evidence with a result.** "Evidence of kind X is
present and passed" is all it needs to be, which is what the v1 outlook already said
before this document elaborated it into something larger. Whether a finding is critical,
and what counts as new relative to the merge base, is the scanner's quality gate and is
enforced by the pipeline failing. Xeno never reads a report, needs no normalised format,
stores no baseline and manages no suppressions, and a project whose tooling cannot emit
SARIF is excluded by nothing.

The residual is the same one as with approval rules: the threshold lives in pipeline
configuration, outside the trail, where it can be weakened without leaving a trace. The
platform enforces, Xeno records, and 10.7 says so.

**The configuration is reported against two references.** A shipped default
configuration exists, and the project's own is compared against it, field by field,
each divergence a line with the shipped value, the set value and a `reason` written by
whoever set it. No approval is required, because most divergences are ordinary, higher
limits in P3 being the obvious case; what an assessor wants is not that nobody diverges
but that every divergence is attributable. The second reference is what the project
class requires, since a shipped default states what Xeno considers sensible and not
what anybody demands. Only with both columns does the report answer the question
actually asked, which is not why this differs from the shipped file but whether it is
permitted for this class. The comparison itself is the same comparison of recorded
hashes already used for rule drift and for catalogue revocation.

**Skill levels are restricted to repository and organisation**, the latter as a
pinned copy in the repository. User skills and the global registry are refused
by the runner. A skill that exists on one machine breaks reproducibility.

**Before and after are recorded separately.** `context.lock.yaml` holds what was
available to a phase, `run.yaml` what was actually used in it, derived from the
events. The split runs along that line and not along the kind of thing, so it
holds for skills and for MCP servers alike. Without it, triggered skills make
principle 5 unredeemable.

**Revocation is drift.** It blocks future use, and existing artifacts are not
touched: an artifact records the hash of every skill, agent definition and
server it used, so whether one of them is still approved is answered by
comparing that hash against the catalogue. Same mechanism as rule drift and as
the review date on the mirrored model set, and no third concept.

**MCP tool budget moves from the project to the sub-agent.** Each agent carries
only its own surface. The rule that an additional operation needs an argument
stays, now per agent. Remote servers are permitted in the phase path and
excluded from the gate path.

## 9. Limitations, section 16

Resolved: 15, token figures, for the orchestrated mode.

Changed: 1, risk differentiation now has an attachment point in the list of
approval required phases. 12, identity now rests on an authenticated act in the
tracker rather than on a free text author line, at the price of evidence that
lives outside the repository.

New, and each of these belongs in the document rather than in a footnote:

- The raw event stream expires by configuration. After that the trail shows
  what was recorded, not what was said.
- Manual mode records less. The secret gate applies later, there is no action
  record and no delegation tree.
- The catalogue is checked for membership, not for currency. Whether the
  approved set is still the right one is a human answer.
- Disposition across phases is model driven. Order, completion and approval are
  not, which is what the first three decisions in section 3 are for.

## 10. Regulatory fit, SDLC scope only

A first cross check against ISO/IEC 42001, ISO/IEC 27001 and the DORA
regulatory technical standards, limited to the software development lifecycle.
Organisational measures are out of scope here by intent.

This belongs in a document of its own, the one that later feeds the conformance
mapping work package. Keeping it here for now is a matter of sequencing, not of
substance, and the split is an implementation question.

### 10.1 The boundary

Xeno ends at the merge. Deployment, operation, detection, response, recovery,
resilience testing and incident reporting lie outside by construction, which v1
states as limitation 3 of section 16. In ISO/IEC 42001 terms the deployment and the
operation and monitoring controls of the AI system life cycle lie beyond the
phase boundary as well.

This is a boundary, not an omission, and it is written down so that it reads as
one.

### 10.2 The answer to the agent risk

The DORA RTS requires measures mitigating the risk of unintentional alteration
and intentional manipulation of ICT systems during development, maintenance and
deployment into production. That is the agent risk verbatim, and it is the
question an assessor asks first once agents write code. The answer is spread
across this document, so it stands here once, in one place.

1. Execution is sandboxed, and the sandbox holds only short lived credentials
   scoped to the intent branch, which holds on every tier rather than only on
   the central one.
2. Actions are classified before execution, and anything above the threshold
   requires a person, with the deterministic analyzers working without network
   and without a model.
3. The verdict before a merge is deterministic, local, network free and model
   free, and the stop hook means an agent cannot declare a phase finished whose
   gate is red.
4. No agent runs in CI. CI stays read only and recomputes what the runner
   already decided.
5. Where a run is started by an event rather than by a command, it starts only
   from an act by somebody holding the rights for it, and it reads the
   triggering content as of that moment and as data.
6. There is exactly one write path for artifacts, and anything written past it
   becomes a finding rather than a silent change.

7. Taken together, the attack surface of an agent is smaller than that of a
   developer with a laptop: push rights to one branch, short lived credentials,
   no tracker token, and writing only through one operation.

The property that makes this an answer rather than a list: not one of the seven
depends on a model behaving well. The seventh also fixes the yardstick, since the
question is not how dangerous an agent is but what it is being compared with, and the
comparison is the status quo rather than an ideal procedure.

### 10.3 Release record

Change management asks about the change that reached production, not about a
single merge request, which is where the v1 record ends. The release record in
section 4 closes that, and it is also what carries the handover at the boundary
described in 10.5.

### 10.4 Security findings

v1 attaches scanner output as evidence without recording how the check ended, so
a scan that failed and a scan that passed look alike in the trail. The predicate
in section 8 closes that: a rule can require that a check of kind X ran and
passed, and the merge is barred by the pipeline in any case.

What this does not resolve, and it is more than v1 admitted: Xeno evidences that
a check with a threshold ran and how it ended, never what the threshold was or
whether the scanner looked in the right places. Both live in pipeline
configuration. Limitation 7 narrows, it does not disappear.

### 10.5 Environments

Xeno does not know environments and will not. Separating development, test and
production is a property of infrastructure, evidenced by infrastructure
configuration, and no lifecycle tool can attest it.

What Xeno can evidence is that what reached production is what passed the
gates. That is the release record, plus a deployment reference as a deferred
extension: an external event stating that a release arrived in an environment
at a time. The release record is sealed when the release is cut and is never
edited afterwards, so the reference is its own append only file pointing at the
release by hash, per the rule in section 4. It is the only one of the three
"later knowledge" cases in this document that needs a write at all, because the
information comes from outside and exists nowhere else. Xeno does not deploy.
It records that somebody else did.

### 10.6 Segregation of duties

In orchestrated mode the author of a commit is the service account, so the
existing rule that an approval may not come from the author is formally always
satisfied and substantively empty.

The explicit rule: the person approving a phase must not be the person
approving the merge request, and both are authenticated humans. For that rule
to bite, the three fields a commit already has are each used for what they are
good for. The author line names the responsible human, being whoever approved
the phase or, where there was no hold point, the owner of the intent, which is
what platform approval rules evaluate. The committer is the service account,
which is simply true. `Co-authored-by` names the agent and its model, which
puts the AI contribution into `git log` and `git blame` rather than only into an
artifact, and a further trailer carries the run id so that a commit points at
its `run.yaml`.

Trailers are free text in the commit body and no approval rule evaluates them,
which is precisely why the responsible human belongs in the author line and not
there. Squash merges keep trailers only if the squash message carries them over,
which is a project setting and one v1 already treats as a special case.

GitLab Enterprise can enforce the rule at platform level through approval rules,
so Xeno records rather than enforces, which is the same division of labour as
with tracker identity in section 6.

The residual risk is that the rule lives in platform configuration, outside the
repository, where it can be changed without leaving a trace in the trail.
Recording the effective approval rule alongside a release would close that, and
is an option rather than a decision.

### 10.7 What remains uncovered

**Conformance is a property of the configuration, not of the tool.** Every one of
the six points in 10.2 rests on settings a project can weaken without anything
turning red: thresholds, which phases require approval, allowlists, the project
class itself. The configuration report in section 8 makes that visible against
the shipped defaults and against what the class requires, which is the most a
tool can do. Making a configuration visible is not the same as making it right.

Everything after the merge. The topology of environments. Data for the
development of AI systems, if the organisation builds AI systems rather than
merely developing with AI, which is a separate evidence chain. And the register
of ICT third party providers, which Xeno keeps none of and should keep none of,
while supplying raw material for it with the model, version and tool recorded
per artifact.

## 11. Deferred

**2.0 and 2.1 are two releases**, split along the deployment tiers, with the local
tier first so that v2 can be dogfooded before a platform is operated. The split and
its reasoning are in [[orchestrator-evaluation#two-releases-not-one]].

**2.1:** parallel agents within a phase, including separate clones and a merge rule.
Approval through a graphical surface, subject to verifying that it can deliver
an authenticated identity. The read only dashboard, which shows intents and
verdicts where Agent Canvas shows runs.

**2.2 or later:** intents spanning several repositories. The condition is a
statement about what binds the verdicts of two repositories together, because
the sentence "every hash resolves locally and without network" falls with it.

**Open, not decided:** whether platform metadata can carry the intent and the phase at
all. The request is made by OpenHands rather than by Xeno, and whether arbitrary metadata
is passed through to the gateway is a question to the platform. If it is not, the split
still comes from the runtime metrics and only attribution below the project in the
gateway falls away.

**Open, not decided:** version binding for the platform. v1 binds runner and plugin
by version and digest, while the agent server, the SDK and the sandbox image are bound
nowhere, although they change what an agent does. `run.yaml` records them from the
start, so the decision can be taken later against real data rather than in advance.

**Conditional:** a durable execution engine such as Temporal, if concurrent
intents ever require a real scheduler. It would sit beside Xeno, never inside
it, and it is not a replacement for the repository as the process state.

## 12. For the adoption guide

Two things belong in the guide that does not exist yet, and nowhere else, because in
the process definition they would look like a tool binding that is not intended.

Server side scanners produce no file. A tool that analyses on a server reports nothing
into a pipeline unless a step asks it to, so the job that writes the result into the
evidence item has to be added deliberately, with one example per common tool.

`.xeno/intents/` becomes a hot path in a monorepo. Many branches create many new
subdirectories that all land on the default branch, and `CODEOWNERS` rules on that path
fire on every merge request. The conflicts themselves are empty, since each intent only
touches its own directory. A paragraph in the guide is the whole answer; partitioning by
year or area would take an afternoon and solves a problem nobody has reported.

## 13. To verify before implementation

Two points, both recorded in the agent layer work package rather than here, so
that they are read by whoever builds that part:

- Whether the gateway reports the model that actually served a request, rather
  than the requested virtual name. The record of what was used depends on it.
- Whether the client accepts the instance URL for the plugin marketplace
  directly.

## 14. Consequences for v1

This is not a work package assignment. That is re-evaluated once v1 stands,
because planning against unbuilt code is guessing, and the v1 plan already says
as much about its own 1.1 list. What follows is narrower and worth something
now: places where v1 would otherwise build something v2 discards, or omit
something v2 would have to retrofit expensively.

**Enforce the phase order in v1.** Refusing a phase whose predecessor has no
completed verdict is a local, deterministic check that costs almost nothing
today and carries the whole of v2 later. There is no reason to wait for it.

**Pull the predicate type forward.** "Frontmatter value against a named set"
sits on the 1.1 list. v2 needs it three times, for the model set, the approval
catalogue and the MCP server list. Building it in v1 is cheap and removes a
dependency from v2.

**Treat the single write path as load bearing.** v1 already routes artifact
creation through the MCP operation with a command behind it. v2 stands or falls
on it, because it is what keeps intent binding, schema and anchors at one code
location while several agents work. It should be built as a property, not as a
convenience.

**Introduce the correlation identifiers in v1.** Intent key and phase id passed
as metadata on every model request cost nothing when the CLI is the only
caller, and they mean that a single developer working through the proxy already
produces attributable numbers. Without them, v1's cost figures stay a sum per
developer.

**WP11 splits into three layers, and only the third is expensive.** Content, so
phase instructions, sub-agent definitions and lenses, is files in the
repository and harness neutral. Packaging, so that a client finds that content
by itself, costs a little per format and is the difference between a developer
using Xeno and a developer setting Xeno up. Wiring, so hooks, is the only layer
that means real work per harness, and what it buys is G-Secret at write time
rather than at phase finish.

Content and packaging stay in both versions, because single developer use
continues in v2 through the manual mode. Wiring becomes justification bearing,
and the justification is best found in dogfooding: if G-Secret regularly fires
only at the end of a phase, the wiring has earned its price. Before that it is
a presumption. How many harnesses to package for is therefore a smaller
question than it looked, and it can wait.

**Keep session log parsing, but expect it to shrink.** If single developers work
through the gateway too, the gateway supplies their numbers as well and the
per-harness parsing falls away entirely. That depends on whether the proxy keys
allow attribution below the project, which is a question to the proxy
configuration rather than to Xeno.

## Appendix A. `run.yaml`

One file per phase directory, inside `artifacts_hash`, mandatory in both modes.
Fields of the section that does not apply to the mode are absent rather than
empty.

```yaml
schema_version: 2
run: 01JB8F3K2Q                    # correlation id, also sent as request metadata
mode: orchestrated                 # orchestrated | manual
attempt: 1

provenance:
  platform:
    agent_server: 1.14.2           # recorded, not yet bound, see section 11
    runner: {version: 2.0.1, sha256: 3e77...}
    sandbox:
      type: container              # container | vm | process
      id: sha256:9f3c...           # only where the environment supplies one
  orchestrator:
    definition: .xeno/config/agents/orchestrator.md
    sha256: 4d1a...
    ref: null                      # set when the definition is pinned elsewhere
  conversation: conv_01JB8F...
  agents:
    - task: task_01JB8G...
      role: phase-agent            # phase-agent | specialist | lens
      name: verification
      definition_sha256: 7b22...
      model_requested: sonnet-5
      model_served: sonnet-5
      thinking: off
      parent: null
    - task: task_01JB8H...
      role: lens
      name: security
      definition_sha256: c019...
      model_requested: haiku-4-5
      model_served: haiku-4-5
      thinking: off
      parent: task_01JB8G...
  skills:                          # used, not merely available
    - id: acceptance-mapping
      sha256: 1f9e...
  mcp_servers:                     # used, not merely available
    - name: xeno
      transport: stdio
      tools_called: [read artifact, write artifact, run gates]

execution:
  started_at: 2026-09-19T08:14:03Z
  ended_at: 2026-09-19T08:51:40Z
  outcome: completed               # completed | aborted-limit | aborted-error
  iterations: 34
  limits:
    iterations: 60                 # cumulative across attempts
    tokens: 4000000
    attempts: 2
  actions:
    policy: confirm-risky
    analyzers: [pattern, llm]
    threshold: high
    total: 212
    confirmed: 4
    rejected: 1
    notable:
      - at: 2026-09-19T08:33:11Z
        action: bash
        risk: high
        decision: rejected
        by: m.example
        reason: deleting the migration directory was not asked for
  secrets_available: [gateway_token, git_push_token]   # names only, never values
  events:
    sha256: a77c...
    window: [2026-09-19T08:14:03Z, 2026-09-19T08:51:40Z]
    store: conv_01JB8F...
```

`notable` carries every rejected action and every action classified at or above
the threshold, in full. Everything else is counted and not listed, because the
interesting subset is the one where a person intervened. Every rejection also
goes into `learning.yaml`.

In manual mode `provenance` holds the platform block empty, no conversation, one
agent entry with `role: phase-agent` and whatever model is known, and `execution`
holds only the timestamps and the outcome.

## Appendix B. `project.yaml`, v2 additions

`xeno init` still asks three questions and defaults the rest, but in v2 they are the
tracker with its project key, the project class and the model set, since class and model
set are what everything else in this appendix hangs off. The language of the artifacts
drops to a default. The v1 version of the three is in [[implementation-plan]].

Everything in the v1 appendix stays. What follows is added. The `orchestrator`
block is optional, and its absence is what makes the manual mode the whole of the
tool, which is the entry tier.

```yaml
schema_version: 2

project:
  key: PROJ
  class: regulated                 # drives risk dependent gates and approvals
  fact_sheet:
    id: FS-2026-014                # authority for the approved models
    reviewed: 2026-09-19           # drift once it ages, a P5 checklist line

runner:
  min_version: 2.0.0               # a floor, never a range

evidence:
  source: ci                       # ci | local

models:
  allowed: [haiku-4-5, sonnet-5, opus-5, fable-5-1]   # mirrored from the fact sheet
  default: haiku-4-5               # the floor, deliberately cheap
  by_phase:
    "01": fable-5-1                # every escalation is one readable line
  by_agent:
    impl-backend: opus-5
  thinking:
    default: off
    by_phase:
      "01": medium

catalogue:
  path: .xeno/config/catalogue.yaml  # skills, agent definitions, MCP servers, with hashes

retention:
  raw_events: 12m                  # the hash in run.yaml outlives it

orchestrator:
  endpoint: https://agents.example.internal   # or: local
  definition: .xeno/config/agents/orchestrator.md
  approvals:
    phases: ["01", "02", "05"]     # a red gate always holds, independently
    surface: tracker               # tracker | cli
  limits:
    default: {iterations: 60, tokens: 4000000, attempts: 2}
    by_phase:
      "03": {iterations: 160, tokens: 12000000}
  triggers:
    intent_create: label:xeno-intake     # an act by somebody with the rights
    approval: label:xeno-approved
```

Four notes on reading it. `models.allowed` is a mirror and not the authority; the
fact sheet is, and the mirror exists so the gate path needs no network.
`limits` count cumulatively per phase across attempts. `approvals.phases` is the
list of hold points that require a person, not the list of hold points that
exist. And `triggers` names acts rather than content, which is what keeps a
stranger from starting a run.
