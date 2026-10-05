---
intent: github.com/triplem/xeno#239
phase: 00-intake
created: "2026-10-05T08:50:03Z"
schema_version: "1.0"
runner_version: dev+3f5ad34
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 979be27f5ad9f6b22a4c94ab541886b525aa6271b342f3586b2baeb5385cc87d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

The repository root holds nine documents. Five of them have a reason to be there:
`README.md` is the entry point, and `CONTRIBUTING.md`, `SECURITY.md`, `LICENSE` and
`NOTICE` are files the host surfaces from the root by convention. The other four have
none: `ASSUMPTIONS.md`, `SUPPLY-CHAIN.md`, `CLAUSE-READERS.md` and `M0.md`.

`docs/` is where this project's prose lives. It holds the two normative documents and two
reference ones, so a reader looking for a document looks there, and four of the nine are
not there to be found.

`M0.md` shows the shape of the problem most plainly: its own first code block prints the
repository layout it is part of, lists `docs/process-definition.md` and the three beside
it, and sits outside that directory while doing so.

The cost of leaving it is not untidiness. WP16 is where link checking arrives, and a link
checker over a tree whose documents are split between two places by no rule is a checker
that has to be told the exception rather than the rule.

<!-- xeno:section:scope -->
## Scope

In scope is moving four documents under `docs/` and making every reference to them that
can be corrected resolve again. `ASSUMPTIONS.md`, `SUPPLY-CHAIN.md` and `CLAUSE-READERS.md`
keep their names in lowercase; `M0.md` becomes `docs/m0-gate-job.md`, because every sibling
under `docs/` is a descriptive hyphenated phrase rather than an identifier and the file is
a walkthrough of standing up the gate job. The moved M0 document gains one line saying it
records how something was done once, which is what a reader who finds a walkthrough in a
documentation directory will otherwise not expect.

In scope is one row in `ASSUMPTIONS.md` for the references this cannot correct. The trail
holds sealed mentions of all four old paths, inside artifacts whose `artifacts_hash`
covers them, and rewriting one stales every verdict over it.

In scope is both issues in one commit. #239 and #222 are the same move over overlapping
files, and four moves in four commits pay the reference sweep four times; #239 says so and
`Closes` both.

Out of scope are the five documents that stay. `README.md` is the entry point and the other
four are host surfaced, and while the host would surface them from `docs/` as well, moving
them buys nothing against a convention every reader already has.

Out of scope is a link checker. That is WP16's, and this intent is the tree it will run
over rather than the run.

Out of scope is rewriting the sealed references. What is sealed is never rewritten, which
is the rule #217 met over its own artifact and resolved the same way.

Out of scope, and the one reference this intent leaves broken knowingly, is
`docs/implementation-plan.md` line 2009: "its other choices are recorded in its
`ASSUMPTIONS.md`". The plan is normative and the first standing rule puts it beyond the
agent, so a correction there is a person's commit made before the code rather than part of
this one. It is a filename in inline code and not a link, which is why it is a staleness to
record and not a break to chase; the row in the register covers it with the sealed ones.
Neither normative document says where a non-normative document lives, so nothing in this
move contradicts a specification and no specification commit is required to precede it.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the four documents themselves, the files that name them, and the count of
which of those references can be corrected.

The two normative documents are read rather than quoted, because this intent's scope claims
neither of them governs where a non-normative document lives, and that claim is checkable
either way. It held for the process definition and failed for the plan: one line of the
plan names `ASSUMPTIONS.md`, which is how the one reference this move cannot correct was
found rather than left for a reader.

`CLAUDE.md`, `CONTRIBUTING.md` and `README.md` are in scope because each names at least one
of the four by path and all three are correctable. The workflow files are in scope for the
same reason, in comments rather than in a path any step resolves, which is worth knowing
before deciding whether a stale one breaks a run: it does not.

`internal/runner/exchange.go` is the only Go file in scope, and it is here for a comment
rather than for code. Nothing in this repository opens any of the four documents by path,
which is the fact that makes this a prose change: the move cannot break a build, and the
test suite is therefore evidence about the sweep and not about the rename.
