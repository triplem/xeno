---
intent: github.com/triplem/xeno#267
phase: 00-intake
created: "2026-10-06T19:55:37Z"
schema_version: "1.0"
runner_version: dev+1d61fa2
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d9dab5eb586badbf5d808ef5b68c697c1d992fd2902db37e9482ec069bc306f0
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

`phase start` writes `context.lock.yaml` from the context scope, and at P0 the scope does not
exist yet: `xeno scope set` writes `phases/00-intake/context-scope.yaml` and refuses to run
before the phase has started —

    refused: 00-intake has not started; run xeno phase start --intent K --phase 00 first

So the order is forced. The lock is born before the thing it resolves, and nothing fills it
afterwards, because #215 made the lock write-once: it sits inside `artifacts_hash` and it is the
only record of what the phase was given. **No P0 `context.lock.yaml` in this trail records a
`files` list** — 0 of 117 when #267 was filed, measured by testing each lock for the key.

Three readers therefore read nothing at P0. `staleReads` iterates `lock.Files`, so an intake can
never be found stale. `budget` sums `f.Bytes` over the same list and compares its length against
the declared count, so `recorded` stays false and the check reports nothing at the one phase that
declares a budget. And `ChangedSince` prints the predecessor's files, which for P1 is an empty
list.

## What the specification says, and does not say

Section 5 says of the lock: "The file is written before the agent starts, from the context scope"
and "It is written once per phase and not refreshed." Both are true and neither mentions that at
P0 the scope is not there to be read, so a reader of section 5 has every reason to expect an
intake's `files` to be populated and no way to find out from the document that it never is.

Section 5 says of the budget check that what it prevents "is the budget quietly becoming
decoration", and section 7 gives the staleness half two limits with a reason each. Neither says
the intake is outside them.

## The decision taken

#267 put three candidates: `scope set` completing the lock, `phase start` taking the scope as an
argument at P0, or recording in the specification that the two checks do not apply there. The
maintainer chose the third on 2026-10-06.

It is defensible for a reason the issue understated. The scope is one artifact per intent and
every later phase resolves the same one, so the budget is judged five times per intent and a
moved file is found against the locks of P1 onwards. What is not judged is the intake against the
budget the intake declares. The alternatives each buy that one phase at a cost: completing the
lock in `scope set` makes one command the writer of two artifacts and needs an ordering guard,
because `context_hash` is rewritten from the lock on every section render and a scope set after
the sections would leave a stale hash and a red G-Schema two commands later, which is #225 again;
taking the scope at `phase start` makes it an input to starting rather than P0's product, which
is a change to what the intake is.

<!-- xeno:section:scope -->
## Scope

In scope are three passages of `docs/process-definition.md`, drafted on the issue and approved by
the maintainer before any of them was written:

A paragraph where the lock is described, section 5, saying that at P0 `files` is empty, why the
order that produces it is forced, and that the two checks reading it are judged from P1 on. The
schema block above it gains a comment on the `files` line pointing at that paragraph, because
somebody reading the shape of the artifact is who needs it.

A paragraph under the budget clause, section 5's context economy, saying the check is judged from
P1 on and naming what the intake still binds: one budget per intent and five phases inherit the
scope it declares.

A paragraph under the staleness half of G-Freshness, section 7, saying P0 stands outside the
check for a reason that is not one of the two limits chosen against noise, and what the
arrangement loses — a file that moved while the intake itself ran, since the next phase's lock
records it as it was by then.

In scope are two code comments that assert the opposite of those paragraphs and are provably
wrong. `informationBase` in `internal/runner/runner.go` says an empty list is "a phase started
before that rule rather than a project opting out", citing #217's rule that a P0 cannot be
finished without a scope. `staleReads` in `internal/gates/gates.go` says the same thing in the
same words. Both are true for P1 to P5 and false for P0, where the empty list is every intake
ever run and not a phase older than a rule.

In scope is a row per clause in `docs/clause-readers.md` where the table's shape asks for one,
since three clauses are being added to a document that table indexes.

## Out of scope

Out of scope is any change that makes P0's lock record files. That is the decision, not an
omission: the two candidates are written down in #267 and stay available if the figures later
argue for them.

Out of scope is a check result that says a check did not run. #235 asked for that shape and
XENO-0265 closed the advisory half of it; a `budget` reporting "not applicable at P0" would be a
new value in a set A4 and A42 fix, and so a specification change of its own.

Out of scope is closing #267's measurement as a test. Nothing here changes behaviour, so there is
nothing to assert that the existing suite does not already cover.

<!-- xeno:section:context-rationale -->
## Why this context

Six files, 321021 bytes, in order of volatility.

`docs/process-definition.md` is the artifact being changed and the only normative statement of
what the lock holds, what the budget check compares and what the staleness half is limited to. All
three passages land in it.

`internal/runner/runner.go` holds `Start`, which writes the lock, `informationBase`, which resolves
it from the scope, `ChangedSince`, which reads the predecessor's `files`, and `SectionSet`, whose
rewriting of `context_hash` on every render is what makes the first of #267's candidates need a
guard. One of the two wrong comments is here.

`internal/runner/scope.go` holds `ScopeSet` and the refusal that forces the order, and
`resolveScope`, which both writers share.

`internal/gates/gates.go` holds `budget` and `staleReads`, the two checks the paragraphs have to
describe honestly. The other wrong comment is here.

`docs/clause-readers.md` indexes section clauses against the code that reads them, so three new
clauses are three candidate rows.

`CLAUDE.md` carries the standing rules this intent runs under: the specification change is its own
commit before the code that follows, and the first rule bars the agent from the documents, which
is why the wording was drafted and approved before anything was typed into the file.

`docs/assumptions.md` is deliberately not in the scope. Nothing here is an assumption taken while
building; the one open judgement, which of #267's three candidates to take, was put to a person
and answered.

The links block declares `internal/runner` against the process definition, because the forced
order the paragraphs explain is a property of two commands in that package and of the artifact
section 5 defines, and reading either alone gets the explanation wrong.
