---
intent: github.com/triplem/xeno#69
phase: 00-intake
created: "2026-09-26T13:21:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+9313e10
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 6af600ae156a956c7962a57b9a7bef92c2f62e590aab8321b7400cafba6d5178
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

The register had grown to 49 rows and nobody had read it end to end since. Fourteen rows
said `open, explained` while nothing in them was pending: the phrase was true while the
work was in flight and stayed behind when the work finished. Eight more said `open` for
something answered and scheduled, which reads as a question nobody got to. One was
closed and said so only in prose. And three statements in the file had stopped being
true: an adapter that changed host, a `Not built` list naming three things the section
above it calls built, and a count of standing verification points that two answers had
overtaken.

<!-- xeno:section:scope -->
## Scope

The state column, and the three errors. The fourteen read `approved` with their
explanations kept; the eight name the package that closes them; A20 reads closed; the
header says what each state means so the column can be read without this commit.

Not the eight rows that want a decision rather than an edit. A5, A15, A16, A17, A22,
A24, A26 and A27 are being put to the maintainer one at a time, and A27 is #47 already.

<!-- xeno:section:context-rationale -->
## Why this context

**Nothing here approved itself.** The fourteen were confirmed by the maintainer, in one
answer, after being listed with what each would mean. The register's own header reserves
the yes for a person and that is exactly why the batch had to be asked rather than
inferred from the work being done.

**Six of the eight scheduled rows got an issue, two did not.** A gap in code that
already ships is discoverable only from this file, and a file is not a work queue: #63
to #68 carry G-Freshness's missing half, G-Learning's shape, the absent cost record, the
routing WP4 has to honour, G-Complete's P5 mode and the two enforcement settings nobody
compares. A13 and A46 describe work their package will reach anyway, so they name the
package and stop there rather than filling the tracker.

**The errors are the reason to reread rather than to rewrite.** None of the three was
introduced carelessly: each was true when written and was overtaken by a decision
recorded somewhere else, #36 for the adapter, WP9 and WP10 for the built list, #56 for
the count. That is the failure mode of a hand kept record, and it is an argument for
rereading it on a schedule rather than for keeping it differently.