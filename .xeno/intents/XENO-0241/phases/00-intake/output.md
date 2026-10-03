---
intent: github.com/triplem/xeno#206
phase: 00-intake
created: "2026-10-03T20:43:42Z"
schema_version: "1.0"
runner_version: dev+8574810
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7070be412dca43dcc55bc1f04fd1a12a4113f5a936463d82508a037f73d87ef6
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

Section 7 puts G-Complete at "P5 and intent close" (`docs/process-definition.md:945`), and
both modes work. Neither covers an intent that simply stops: no P5, and no `xeno intent
close` either. XENO-0230 reached `main` with its verification and review phases missing
and every gate green on the phases it did write. Nothing was red and nothing was silent,
because the completeness check had no phase to run in.

The fact is already computed. `state` in `internal/runner/runner.go` returns one of three
answers for every intent — `abandoned`, `complete`, or the phase the work reached — and
`xeno intent status` prints it. What is missing is a reader that treats it as a condition
rather than as a column. All seventy post-M0 intents read `complete` today, XENO-0230
included, since its two missing phases were filled afterwards; the gap is that nobody was
told, not that the trail is currently wrong.

This is the shape the clause audit (#202) was filed to find. The clause is enforced in
both places it claims to be, and the gap is the case between them.

<!-- xeno:section:scope -->
## Scope

The issue offers three answers and names them as open. A person chose the first: the
check is enforced in CI, and the condition lives in the tool rather than in the workflow,
so that it is under test and not parsed back out of printed text.

In scope: one function in `internal/git` that lists the paths differing between two trees;
one reader in `internal/runner` that resolves those paths to intent keys and reports the
ones whose state is neither `complete` nor `abandoned`; one command, `xeno intent verify
--base REF --head REF`, that exits 1 where there are any; one step in
`.github/workflows/xeno.yml` that calls it; and the same call in both shipped CI wrapper
templates, which already pass both ends of the range and are what an adopter gets.

Out of scope: a branch that touches no intent at all. That is a violation of this
project's own rule that every change belongs to an intent, and it belongs to #120 rather
than here, so the command reports the count and exits 0.

Out of scope: the `changed-after-verdict` state, which `xeno gate verify` already reports
as a divergence, and anything that writes. The command reads and reports.

No field, no gate, no rule, and no document. G-Complete is unchanged and the gate list is
not extended; the check reads the same three-way state `intent status` has printed since
#132, at the one moment the process says a merge is decided.

<!-- xeno:section:context-rationale -->
## Why this context

The inputs are section 7's G-Complete paragraph and section 8's "Ending an intent", which
together say what completeness means and why there is no `merged` status; `state`,
`summarise` and `Status` in `internal/runner/runner.go`, which already compute the answer;
`internal/git`, because the range has to be read and that package is the only place in the
tree that starts a subprocess; and the `verify` job in `.github/workflows/xeno.yml`,
because its trail-rewrite step is the nearest thing to what this adds and a second
range-reading step should read like the first.

`internal/scaffold/files/ci-github.yml` and `ci-gitlab.yml` were read to find out whether
the check can reach an adopter, and it can: both already pass `{{.BaseRef}}` and
`{{.HeadRef}}` for `gate run`, so the new call needs no new template field.

`xeno intent status --all` was run over the whole trail before any code was written. Had
it shown intents short of P5, this change would have been a migration as much as a check,
and the scope above would be wrong. It showed seventy complete, which is why the step can
be added and required in the same commit.

`CLAUSE-READERS.md` was read and is deliberately not changed. It dates itself, says "it
is a measurement, not a specification", and states that ageing is the most a document can
do about itself. Editing a measurement after the fact to say the gap it found is now
closed would make it an index that is maintained, which it declines to be.
