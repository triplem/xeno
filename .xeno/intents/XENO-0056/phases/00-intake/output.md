---
intent: github.com/triplem/xeno#56
phase: 00-intake
created: "2026-09-26T11:14:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f5114fa
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 8aaa5d10685e4dae2389f196abd112a9ea923e9194362ff66ef3520a8a1c8f2f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

Section 9 of the implementation plan holds four things to verify rather than assume,
each before the code that relies on it. The gateway is two of them: whether it reports
the model that actually served a request rather than only the virtual name that was
asked for, and whether it accepts and records request metadata. The annotation said
they would be sorted out later, and WP11 repeated them as the second of its two
unverified points. The record of which model was used depends on the first and the
attribution of cost below the project on the second, so both stand in front of the
agent layer.

<!-- xeno:section:scope -->
## Scope

The plan's two questions, answered by measurement against a running LiteLLM proxy, and
the annotation in section 9 replaced by what was measured. WP11's paragraph loses the
gateway and keeps the marketplace URL, because the two documents would otherwise
disagree about how many things are left to verify.

Nothing is built here. The export that sends intent and phase from `xeno phase start`,
the header names it has to use and the harnesses that carry them are the work the
answer makes possible, and they belong to WP11 rather than to this intent.

<!-- xeno:section:context-rationale -->
## Why this context

**The answer had to be produced, not looked up.** The proxy ran with a single
deployment whose public name equalled its upstream model, which is the one arrangement
in which the two names cannot be told apart. A second deployment was registered whose
public name differed from the model behind it, the probes ran through it, and it was
removed again; without it every observation would have been consistent with the
gateway reporting only what was asked for.

**The negative result is the load bearing one.** A request carrying `X-Xeno-Intent` and
`X-Xeno-Phase` was recorded with no metadata and no tags, and no error anywhere. An
export written from the plan's own wording would have produced exactly that: rows with
no attribution in them, discovered whenever somebody first queried them. The names
LiteLLM reads are therefore part of the answer and not an implementation detail of it.

**What the limit is.** The gateway names the deployment it routed to, taken from its
own configuration, and the provider's returned model id is not surfaced beside it. The
process definition requires `model` in every session produced artifact and does not say
which of the two identifiers it means, which is a sentence for a person to write in
that document rather than an assumption to record here.
