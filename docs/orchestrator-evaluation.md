---
id: orchestrator-evaluation
title: Xeno, Orchestrator Evaluation for v2
revision: 2
status: draft, not ratified
date: 2026-09-20
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
