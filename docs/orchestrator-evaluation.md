---
id: orchestrator-evaluation
title: Xeno, Orchestrator Evaluation for v2
revision: 3
status: draft, not ratified
date: 2026-10-08
location: docs/orchestrator-evaluation.md
---

# Xeno, Orchestrator Evaluation for v2

## 1. Purpose

v2 makes Xeno multi-agent and delegates the execution of a phase to an agent
platform. This document records which platform was chosen, against which
criteria, what was rejected and why, and under which conditions the decision
is revisited. It exists so that the question is not reopened from memory, and
so that a reader who disagrees can see what the disagreement is about.

It is an architecture decision record, not a plan. What follows from it is in
the implementation plan.

Section 9 records a second decision of the same kind, taken later and about a
different kind of tool: a specification framework, evaluated and declined. It
sits here because the form is the same one — a candidate pinned to a version,
the rows it changes, the reason that decides it, and the conditions under which
the answer is revisited — and because a tool declined is only useful to a reader
who can find the reason. The page is the record of external tools this project
evaluated, of which the orchestrator was the first.

## 2. What the platform has to do

The requirement split into two layers during the design of v2, and the split
decides the whole evaluation.

**Layer 2, disposition across phases.** Which phase comes next, what is waiting
for an approval, what resumes after an interruption. v2 answers this from the
repository: `xeno intent status` is a pure function over the intent directory,
a phase starts with a fresh conversation, and the orchestrator carries no
content that is not in the repository. There is no durable orchestration state
outside git, by construction.

**Layer 1, execution of one phase.** Sandboxed execution, sub-agent delegation
with separate context, a stop hook that can block completion until a
deterministic check passes, MCP, skills with levels and integrity, action
confirmation, an event stream, and model access through a gateway.

The consequence is that the platform is chosen for layer 1 alone. Anything a
candidate offers for layer 2 is not a benefit here. It is a second source of
truth competing with the repository.

## 3. Criteria

| # | Criterion | Why it carries weight |
|---|---|---|
| C1 | Self-hostable without a vendor cloud | The target may be self managed. Nothing in the trail may depend on a hosted service. |
| C2 | Model agnostic through an OpenAI compatible gateway | Model access, cost and routing belong to the existing LiteLLM proxy. A platform bound to one vendor's models moves that decision out of the organisation. |
| C3 | Deterministic completion barrier | A phase must not be declarable as finished while its gate is red. Without a blocking hook, the process depends on a model's cooperation. |
| C4 | Sub-agent delegation with separate context | Phase agents delegate to specialists. Separate context is what makes delegation cheaper than a single long conversation rather than more expensive. |
| C5 | Sandboxed execution as the default | The agent runs unattended. Isolation is not a hardening step to add later. |
| C6 | MCP and a skill model with levels | The artifact write path and the process operations are MCP. Skills need repository and organisation scope, not per developer machine. |
| C7 | Event stream and external tracing | Operations needs observability. Evidence stays in the repository, so this is an operational criterion, not an evidentiary one. |
| C8 | Open source, license compatible with Apache 2.0 distribution | Xeno is published from 1.1. |
| C9 | Replaceable behind a port | The largest risk is not the choice. It is a choice that cannot be reversed. |

C1 to C3 are exclusion criteria. A candidate failing one of them is out
regardless of how it scores elsewhere.

## 4. Candidates

### 4.1 OpenHands, chosen

The Software Agent SDK with the Agent Server covers C1 to C7 without an
exception. Relevant properties:

- Agent Server exposes agent execution, conversations, tools and workspaces
  over REST and WebSocket, deployable on own infrastructure.
- Model access runs through LiteLLM, so the existing proxy is the gateway
  rather than a parallel path.
- Stop hooks run when the agent tries to finish and can block completion until
  repository specific checks pass. This is where `xeno phase finish` is called,
  and it is the reason the runner keeps the verdict.
- The TaskToolSet creates sub-agents that run synchronously, return a result
  and can be resumed by task id. Sub-agents are definable as files with
  frontmatter, which lets their definitions live in the repository and be
  hashed.
- Skills exist at repository, organisation, user and global scope, with keyword
  and path triggers. v2 restricts this to the first two, which the runner
  checks.
- Confirmation policy and security analyzers classify actions before execution,
  including analyzers that work without network and without a model.
- The conversation event log is append only and typed, and OpenTelemetry
  tracing is available for operations.
- ACP delegation to other harnesses exists as a fallback path, unused in v2 but
  available if the direct model path has to be abandoned.

`rajistics-demo/sdlc-automation-github-demo` is this choice running rather than an
alternative to it: it drives the same Agent Server through the OpenHands Automations
API and starts each run from a label on an issue, which is the design section 12's
"Issue commands are deliberately absent" paragraph declines, and it is evidence for
that paragraph rather than against it, because the thing receiving the label is a
hosted automation and not anything in the repository — whose tree carries no workflows
at all. #330 took the half of this that needs no receiver: `xeno-approved` is a
precondition `xeno intent start` reads off an issue it has already fetched, not an
event anything listens for.

### 4.2 Claude Agent SDK, rejected, best fallback

The same agent loop that powers Claude Code, embeddable in an own program, with
tool execution, permission gating, MCP, hooks, sub-agents and provider routing
over Anthropic, Bedrock or Vertex. Conceptually the closest match to what v2
needs, and the concepts map almost one to one.

Rejected on C2 and C5. Model access is confined to Anthropic's own routes,
which conflicts with a house-wide gateway and with the ability to place cheap
models on cheap phases. Sandbox, server and any operator surface would have to
be built, which is precisely the part that costs operations rather than
development.

Kept as the named fallback. Because the concepts match, a switch would be an
adapter behind the orchestrator port rather than a redesign, provided the port
contract in section 7 stays free of OpenHands specifics.

### 4.3 Inner loop coding agents, rejected

OpenCode, Cline, Aider, Goose, Continue. Rejected on the shape of the tool
rather than on quality. They assume a developer at a terminal or in an IDE. v2
requires work that proceeds while nobody is watching, which is a server with a
sandbox, not a client.

### 4.4 Durable workflow engines, rejected

LangGraph with a Postgres checkpointer, LangSmith Deployment, Temporal. They
solve layer 2 better than anything else available: checkpointing at every step,
`interrupt()` as a first class primitive for human approval, replay from an
earlier checkpoint.

Rejected because layer 2 is already solved, and solved in the one place that
needs no additional system. The repository is the durable store, `gate.yaml`
is the checkpoint, the phase boundary is the interrupt, `xeno intent status` is
the resume, and `xeno phase reopen` is the replay. Adding an engine would put a
second authoritative record of process state beside the one the whole tool is
built to defend.

Two further reasons. The control flow is a fixed six state machine, so the
expressive power of a graph runtime buys nothing that a directory listing does
not already answer. And the runner is Go, so the library would sit only in the
orchestrator, which v2 deliberately keeps thin and remote.

Temporal is recorded as the candidate to revisit if 2.1 introduces many
concurrent intents and a real scheduler becomes necessary. The problem then is
durable execution rather than graph topology, and it would sit beside Xeno, not
inside it.

### 4.5 General multi-agent frameworks, rejected

AutoGen and the Microsoft Agent Framework, CrewAI, Google ADK, Mastra. They
contribute orchestration vocabulary for layer 2 and nothing for layer 1: no
coding agent, no sandbox, no skill model, no tool surface for working on code.

## 5. Decision

OpenHands is the execution platform for a phase. The repository remains the
orchestration state. The runner remains the only authority on a verdict.

The division of labour in one sentence: OpenHands owns the work, Xeno owns the
record, and every crossing of that line leaves an entry in the repository.

### 5.1 Two operating modes, not two harnesses

An earlier draft of this evaluation proposed dropping the v1 harnesses and
supporting OpenHands alone. That framing was wrong, because what is expensive
about a second harness is never the agent. It is what the runner has to know
about it: a second hook wiring, a second plugin packaging, a second session log
format for cost accounting.

v2 therefore has two operating modes and one code path.

**Orchestrated.** Agent Server, orchestrator port, sub-agents, stop hook,
action control, complete `run.yaml`.

**Manual.** No orchestrator. A person, or whatever agent a developer happens to
use, works the phase through the commands and the MCP server. The runner does
not know who is typing and has no reason to care.

The manual mode is not a second code path. It is the absence of the first, and
v1 already requires it as an acceptance criterion for the agent layer: the same
phase must be workable with commands alone, with no MCP server and no hooks,
producing artifacts a gate cannot tell apart. What v2 drops is harness specific
support, not harness specific usage. Phase instructions and sub-agent
definitions are files in the repository either way.

What manual mode honestly loses: no stop hook, so the secret gate first applies
at phase finish; no action record; no delegation tree; and token figures only
as far as the gateway can attribute them. The `run.yaml` carries a `mode` field
and omits the fields of the other mode rather than filling them, so the trail
does not split into two kinds. A rule may require `mode: orchestrated` for a
given project class.

The mode is therefore a property of a phase and not of an intent. In a project that
has an orchestrator, somebody can still work one phase by hand, and the runner could
not prevent it if it wanted to, since it does not know who is typing. That is the
point rather than a gap: what matters is that the trail says which regime each phase
came about under, and that a rule can forbid the weaker one where it matters.

### 5.2a Two releases, not one {#two-releases-not-one}

The tiers below are also the delivery order. v2.0 is the local tier: sub-agents, the
stop hook, action control, `run.yaml`, catalogue and model resolution, measured cost
through the gateway. Hold points block for the length of a session and approvals stay
commands, because with nobody waiting overnight there is nothing to resume. v2.1 is
the central tier: a run that ends at a hold point and resumes later, approvals in the
tracker, event triggers, the event store with its retention, and the dashboard.

The reason is dogfooding. What made v1 credible was that it could be used while it was
being built, and v2 as originally cut could not be, because it presumes an operated
platform before the first phase runs. v2.0 delivers most of the value against nothing
but a container runtime and the gateway. The break between the two is small, since the
orchestrator port does not change and only its endpoint does.

### 5.2 Three deployment tiers

Adoption has to stay cheap, and the process has to be testable on a laptop.

| Tier | Needs | Waiting at a hold point |
|---|---|---|
| None | Runner and gateway | Not applicable, manual mode |
| Local | Agent Server with a process or container sandbox | Blocking, for the length of a session |
| Central | Agent Server per environment, one sandbox per conversation | The run ends and resumes later |

The orchestrator block in `project.yaml` is optional, exactly as the tracker
block is today. Absent, Xeno runs as v1 did. The same configuration with a
different endpoint moves a project between the upper two tiers.

### 5.3 Model selection

Two levels with a clean division. The gateway decides what is permitted, per
project or per key, which is an organisational decision. Xeno decides what is
used, per phase and per sub-agent, and passes it explicitly on every run. No
service side default participates in the decision.

Resolution has four steps, first match wins: sub-agent, phase, project, shipped
default. The shipped default names no model and refuses to start when the
project has set none. A tool that picks a model for you produces exactly the
silent assumption principle 3 rules out, and model names age faster than
anything else in this document.

The floor is cheap and every escalation is an explicit line. That is the same
shape as lens enablement: cost is opt-in, so each expensive decision can be
read, defended and removed in review. Reasoning effort follows the same rule,
off by default and raised only where the model was raised, since those tokens
bill as output.

`project.yaml` names the gateway's virtual model name, not a provider string,
so routing knowledge does not migrate into the repository and rot there. The
`run.yaml` records both the requested name and the model the gateway reports as
having served the request. A divergence between the two is a finding.

Authority sits outside Xeno. The project's fact sheet holds the approved
models, `project.yaml` references it and mirrors the permitted set with a
review date. The mirror exists because the 1.1 predicate has to check the set
without network access, and mirroring rather than distributing is the pattern
v1 already uses for rule levels above the project. A stale mirror is drift: a
line in the P5 checklist, not a blocked gate. The reference is an identifier
plus an optional base URI set at organisation level, so that a repository which
later goes public carries no internal hostname.

## 6. Risks and mitigations

**Architecture churn.** OpenHands is in transition from V0 to V1, V0 was
deprecated in April 2026, and much published material describes the old
architecture. Mitigation: pin versions, treat the platform docs as the only
reference, and expect API movement during v2 development.

**Open core with a commercial control plane.** Role based access, SAML, audit
logs, usage monitoring and budget enforcement live in the enterprise tier
rather than in the open core. Mitigation: Xeno's evidence is in the repository
and does not depend on that tier. Access control and cost control are solved
through the organisation's own gateway and infrastructure. This has to stay
true as the product develops, and is therefore a condition to re-check.

**License.** The documentation states explicitly that each public repository
carries its own license and that one should not assume a single license across
the ecosystem. Secondary sources disagree about which applies. Mitigation: a
legal review of the specific repositories consumed, completed before v2 work
starts rather than before publication.

**Lock-in.** Mitigated by section 7 and by naming a fallback in 4.2.

## 7. The orchestrator port

The port has five operations and no more. The gate path does not import it, and
the same import graph test that protects the gate path from the tracker adapter
covers it.

| Operation | Contract |
|---|---|
| start | Start a run for one phase with a context reference, return a run id |
| finish | End a run, return its result |
| result | Fetch the result of a run |
| events | Fetch the raw event stream of a run for digest and hash |
| abort | Terminate a run |

Two rules keep the port honest. No parameter names a concept specific to one
platform. And the contract is reviewed once against the fallback in 4.2 before
it is frozen, not to switch, but to detect anything OpenHands specific that has
leaked into it.

## 8. Conditions for revisiting

- Governance capabilities required by v2 move from the open core into a
  commercial tier.
- The license review contradicts the assumption in 6.
- 2.1 introduces concurrent intents at a scale that requires a scheduler. The
  candidate is then Temporal, evaluated as an addition rather than a
  replacement.
- The direct model path proves unusable and ACP delegation becomes necessary,
  which reopens the decision recorded against `agent.tool`.

## 9. A specification framework, declined

### 9.1 What was evaluated

OpenSpec, shipped on npm as `@fission-ai/openspec`, MIT licensed, binary `openspec`,
**pinned at 1.14.1**, published 2026-10-05. Fifty-two versions lie between the first
release on 2025-09-06 and that one, which is about one a week; the registry figures were
read on 2026-10-08 and the behaviour described below from the repository at tag
`v1.14.1` rather than from its documentation.

It calls itself an AI-native system for spec-driven development. Its four stated
principles are fluid not rigid, with no phase gates, iterative not waterfall, easy not
complex, and brownfield-first. Its unit of work is a change directory carrying a
proposal, a design, a task list and delta specifications; `openspec archive` folds the
deltas into `openspec/specs/<capability>/spec.md` and moves the directory under
`changes/archive/YYYY-MM-DD-<name>/`. The four commands that drive the workflow are
typed into the agent rather than into a terminal.

How the evaluation was carried out, candidate by candidate and source by source, is
issue #223. This section is the decision and the reason for it.

### 9.2 The four rows

| What | Here | With OpenSpec 1.14.1 | Changed |
|---|---|---|---|
| the unit of work | an intent, six phases, one issue, one branch | a change directory, named rather than keyed, bound to no issue and with no phases, since "no phase gates" is a stated principle | replaced |
| where the record lives | `.xeno/intents/KEY/`, sealed by `artifacts_hash` | `openspec/changes/<name>/`, then an archive directory; Markdown with no hash of any kind, editable after the fact by design | replaced, and unsealed |
| who enforces it | fourteen gates, deterministic and model free | `openspec validate` for structure and `--archived` for ticked checkboxes; past that, skill text and a person, with the documentation saying that everything below is convention and not enforcement | replaced, and weaker |
| who may edit the specification | a person, in a commit of its own | the agent, as the normal operation | **inverted** |

All four change. The fourth does not merely change, it reverses.

### 9.3 Decision

No, and the fourth row decides it on its own.

OpenSpec's `propose` step has the agent write the specification, and its FAQ instructs
that where hand-written code and the specification disagree and the code is right, the
delta specification is updated to match what shipped. That is the specification being
corrected to the implementation, by the implementer, as the documented happy path. The
first of this project's three standing rules is the opposite sentence: the documents are
not editable by the agent, the specification wins until a person changes it, and a
change to it is its own commit made before the code that follows from it. The rule
exists because an agent that can edit the specification it is judged against has no
specification, only a record of what it did. There is no setting that turns this off,
because it is not a setting; it is what the product is for.

Two coherent answers to one problem, and only one of them can hold in one repository.

It is not declined for quality. Its agent contract — a JSON shape per command, a shared
diagnostic envelope, an exit-code table — is better documented than most of what it
competes with, and 9.6 takes the one mechanism in it this project had not thought of.

### 9.4 What the measurement said

The decision rests on 9.3. The measurement is the smaller half of the case, and it is
recorded because it is what a later reader would otherwise have to redo.

**It replaces one section of one template.** The six shipped templates declare 28
section slots over 18 distinct ids. OpenSpec's fixed vocabulary is four Markdown markers
and three filenames, and it speaks to exactly one of those ids, `acceptance-criteria`.
It has nothing for the other seventeen, nothing for the frontmatter of section 5,
nothing for `gate.yaml`, `context.lock.yaml`, `cost.yaml`, `learning.yaml`,
`assumptions.yaml` or `evidence/attached.yaml`, and nothing for any of the five hashes.
What it would contribute inside `acceptance-criteria` is RFC 2119 and Given/When/Then,
both of which are adoptable by writing them there and installing nothing.

**Interoperability is zero, and was re-measured.** The readers of `openspec/specs/` are
its own CLI and its forks. A third-party product that advertises OpenSpec support was
checked on 2026-10-08 and shells out to `openspec list` and `openspec validate` rather
than parsing the format, which is the CLI again and not a second reader. Against that,
the artifacts here are Markdown with YAML frontmatter and plain YAML, Appendix B defines
every hash to the byte, and the schema is a published deliverable: a documented format
with a stated hash construction is a stronger interoperability claim than conformance to
a single-implementation one.

**The surface gets larger, not smaller.** Nothing in the gate path, the hash chain or
the rule engine has a counterpart in a tool whose own rules its documentation calls
unenforceable. Acquired, in exchange for that one template section: a Node runtime on
every machine that works a phase; a weekly-moving dependency in a project with one
vendored dependency and a rule that a second is a decision rather than a step; an
instruction surface under `.claude/`, rewritten by a globally installed CLI and
therefore outside the plugin digest G-Supply checks; an outbound registry query and
anonymous telemetry to switch off and write down, in a project whose gate path reaches
no network at all; and a second object called a rule, which is prose injected into a
prompt, in a project where a rule is a predicate G-Policy evaluates.

### 9.5 Conditions for revisiting

- OpenSpec grows a content hash over a change directory and a verdict record. The
  sealing objection falls away, which leaves the first standing rule alone and makes a
  narrower argument worth running again.
- A second, independent implementation reads `openspec/specs/`: an auditor's tool, a
  code host, a certification body. Zero and non-zero are different calculations.
- The format passes to a body that maintains it rather than to the vendor that ships it.
- `acceptance-criteria` proves too loose in practice. The remedy is then RFC 2119 and
  Given/When/Then inside that section, written by a person in a commit of its own, and
  not a framework.

What would not change the answer is OpenSpec becoming better at what it does. The fourth
row is a disagreement about who owns the specification, and a better tool does not
settle it.

### 9.6 The one mechanism worth taking

`openspec validate --archived` is a sweep over finished work. It lists every directory
under `changes/archive/`, counts the task checkboxes of each one, and exits 1 where a
change has an unticked box or a task file it could not read. It checks nothing else: an
archived change's delta specifications are not validated, because they were applied at
archive time. That narrowness is what makes it cheap enough for the pre-commit hook its
own help text names it for.

The hole it covers is the interesting part. `openspec archive` does warn about
incomplete tasks, but `--yes` turns that refusal into a line of output, and nothing
looks at the archive again afterwards, because the ordinary discovery of changes
excludes it. Stated without the tool, the mechanism is this: a deliberate bypass taken
at the moment of finishing is re-asked later by something that sweeps what has already
finished.

This project has that mechanism, in a stricter form, spread over three places rather
than one. The bypass is recorded instead of warned: `xeno gate override` writes a
decision of type `overridden` carrying `obligation: open` onto the finding, inside a
verdict `artifacts_hash` seals, where OpenSpec's `--yes` leaves nothing behind at all.
The refusals sit before the ending rather than after it: G-Complete fails a P5 whose
preceding phases are red or provisional, G-Questions fails an open question unresolved
at P5, and `xeno intent verify --base --head` fails a merge range holding an intent that
reached neither a decided P5 nor `xeno intent close`. Between them those three refuse
what `--archived` refuses, which is a unit declared finished with work still open inside
it. And the obligation an override leaves is read back: it is printed as owed in the
next-step block of any command run for that intent, including after P5 is decided, until
`xeno obligation close` closes it.

What is deliberately absent is the refusal after the fact. `xeno intent verify` exits 0
on a range whose intent carries an open obligation, and no gate reads the field. That is
section 6 of [the process definition](process-definition.md) holding its position, not a
gap: nothing forces the follow-up, because a gate that could compel it would have had to
block the merge in the first place, which is the situation an override exists to
resolve. The gate-shaped version of `--archived` is therefore not a gate this project
can have. It would undo the override.

What is left is a reader, and it is already owned. An open obligation is visible one
intent at a time and nowhere across intents; the across-intents view is WP18 in [the
implementation plan](implementation-plan.md), which names open overrides as a view of
its own — which findings were merged over, by whom, for what reason, and whether the
obligation has been closed since — and defers it to 1.1 because everything it shows can
be read from the repository by hand. Nothing further is owed here, and the figure says
why the wait costs nothing: on 2026-10-08 `xeno gate verify` counts 519 verdicts in
this trail, and not one of them carries an override, so the view would have nothing to
show.

## Sources

- https://docs.openhands.dev/overview/introduction
- https://docs.openhands.dev/sdk
- https://docs.openhands.dev/sdk/guides/task-tool-set
- https://docs.openhands.dev/sdk/guides/security
- https://docs.openhands.dev/sdk/arch/events
- https://docs.openhands.dev/openhands/usage/customization/repository
- https://docs.openhands.dev/overview/skills
- https://docs.langchain.com/oss/python/langgraph/persistence/
- https://temporal.io/blog/temporal-langgraph-plugin-durable-execution

For section 9, read on 2026-10-08: https://github.com/Fission-AI/OpenSpec at tag
`v1.14.1` for the behaviour, in particular `src/commands/validate.ts`,
`src/utils/task-progress.ts` and `src/core/archive.ts`, and the documentation beside
them, `docs/cli.md`, `docs/agent-contract.md`, `docs/concepts.md` and `docs/faq.md`;
https://registry.npmjs.org/@fission-ai/openspec for the version, the release dates and
the licence. The evaluation itself, with the namesakes set aside and every source
quoted, is issue #223.

For the sentence in 4.1, read on 2026-10-10:
https://github.com/rajistics-demo/sdlc-automation-github-demo at commit
`2ddf6c91791d95b8de8ff38b683241743ec12049`, dated 2026-10-07, which is a commit and
not a tag because the repository publishes no releases. Read there:
`.github/labels.json` for the four trigger labels and the four status labels beside
them; `automations/github/openhands-{context,build,review,qa}/` for the presets and
prompts the Automations API registers, with
`scripts/automations/register_github_automations.py` for the registration; and
`.github/ISSUE_TEMPLATE/`, `.github/pull_request_template.md` and `openspec/` for the
rest of the design the sentence does not take. The tree holds no `.github/workflows`
directory, which is what makes "something has to receive them" literal there. Nine
stars, one author and no licence file.
