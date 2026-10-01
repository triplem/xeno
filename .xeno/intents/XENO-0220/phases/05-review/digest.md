---
intent: github.com/triplem/xeno#165
phase: 05-review
created: "2026-10-01T16:38:25Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4f94129.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a2e14b1b722a4c2d8feee1698fa3c9a696d307660e895b09ce80f039662806be
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The first review checklist in this repository with entries in it, because this is the piece that shipped
the rules: `deviations-are-traceable` met, since the six deviations of P3 each name what they depart
from; `interface-change-needs-a-migration-note` met, since G-Policy's behaviour changes for every
project and A74 says what a dependant does, which is nothing; `new-dependency-needs-a-rationale`
not-applicable, since none was added. Four questions beyond the rules: the set is refusable, which is
why it is thin; the guard both protects twenty-four sealed verdicts and hides a rule added after an
artifact was written; nothing forbidden was written, with one predicate type admitted by section 9's own
scoping and the plan's open call answered as A72 rather than deferred; and the wording is answerable,
though only by the person who wrote it. Residual risk: the guard's hole is a specification question
about staleness; section 9's one worked rule example cannot ship as written, because the templates have
neither section it names; nothing stops a project editing the shipped set in place while G-Supply is
unimplemented; the set is thin enough that nothing in it would have caught any finding of the last
twenty intents; the hooks were read and never run; and a fifth rule would be either specific, which the
layer forbids, or weaker than the four there are.
