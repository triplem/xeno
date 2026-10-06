---
intent: github.com/triplem/xeno#237
phase: 00-intake
created: "2026-10-06T12:01:32Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bc38359aca6d5e793eb2fc4275850dd774c78530dced031448c31fa01b4a1a68
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

The staleness half of G-Freshness raises three findings, and each prints a remedy. The third reads:

    read the phase again for what changed; the lock records what it was given,
    not what is there now

Reading the phase again is the one act the runner refuses:

    refused: 01-requirements has a verdict; redo the work with section set and
    phase finish, or remove .xeno/intents/K/phases/01-requirements to start it over

**And the first route that refusal offers does not clear the finding.** `section set` and
`phase finish` re-render the artifact and recompute its hash; `context.lock.yaml` is written only by
`phase start`, and `Start` refuses a phase that has a verdict for two reasons #215 gives and the code
restates at `runner.go:409` — the lock is inside `artifacts_hash`, so a second start rewrites a
sealed artifact, and the lock is the only record of what the phase was given, which is the answer to
"what changed". So the lock keeps the hashes it was born with by design, and a person following the
remedy is refused, then offered an act that leaves the finding exactly where it was.

The second route the refusal offers does clear it, by removing the phase directory, which discards
the phase rather than refreshing it.

So a person following the printed remedy is turned away twice before reaching the one thing that
works, and the thing that works is not in the remedy. Every other refusal and finding in this runner
carries a remedy that does what it says — #216 and #218 were careful about that — so this is the
exception and not the pattern.

## What the specification already says

Section 7, under the staleness clause: "A stale phase is not deleted, it is re-run or explicitly
approved as still valid." Two routes, both named, and the remedy names neither. `xeno gate approve`
is the second of them and the runner has carried it since the gate command existed.

## Why this was not a wording fix when it was filed

#237 held it against #236, and correctly. Before #236, `staleReads` looked at every entry of every
lock including the phase being gated, so it fired on a phase that had just done its job, and for
that kind of finding "re-run or explicitly approved" is the wrong instruction — the phase is not
stale, the check is wrong. #236 is closed and the code now carries both of section 5's limits:
`for i := 0; i < idx` excludes the phase being gated, and `if !touched[f.Path] { continue }` narrows
to the change under review. What the check raises now is only what section 7 describes, a phase whose
ground moved after it finished, and for that the specification's two routes are the right answer.

## The three findings are not alike

Two are about a lock that cannot be refreshed and one is not. A file that is **gone** and a file that
has **changed since** are both the ground moving under a sealed phase. A file that **cannot be read**
is a condition of this machine — a permission, a broken link, a filesystem — and "make it readable"
is already a remedy the runner carries out, because the next `gate run` then finds the file and the
finding does not recur. Giving all three the same instruction would make one of them worse.

<!-- xeno:section:scope -->
## Scope

In scope are the three remedy strings in `staleReads`. The two about a sealed lock name the two
routes section 7 names: approving the finding on the phase being gated as still valid, and starting
the earlier phase over. Each also says why `section set` and `phase finish` will not clear it, because
that is the route the refusal offers and the one a person reaches for first.

In scope is the third remedy staying what it is. An unreadable file is a condition of the machine and
not of the trail; `make it readable` already works, and the next `gate run` clears the finding without
anybody approving anything. What it gains is one clause saying which of the two it is, so that a
reader meeting three findings does not read the odd one as an oversight.

In scope is naming `gate approve` rather than describing it. The fault reported is that a person
following the remedy is refused twice before finding the route that works, and a remedy that says
"approve it" without saying with what repeats the fault in a smaller way.

In scope is the remedy naming which phase to approve on and which to start over. They are different
phases: the finding is raised while gating one phase and is about the lock of an earlier one, so a
remedy naming neither leaves the person to work out that `gate approve` wants the later of the two.

In scope is saying what starting the earlier phase over costs. Every phase after it built on its
verdict, so the route is not a refresh — and a remedy that offers it without saying so is the same
trap one step further along.

In scope are tests. The two limits have tests from #236 and the remedies have none, which is how
three strings came to say something the runner refuses.

Out of scope is `phase start` learning to re-resolve a lock. It would make the old remedy true and it
contradicts a closed decision: #215 refused it because the lock is inside `artifacts_hash` and is the
only record of what the phase was given. Reopening that is a section 11 question and a person's
commit before any code.

Out of scope is the budget check stopping the phase. #235 is the other wp8 issue and is about
`result` turning any finding into a red gate; it reaches this one only in that an approved staleness
finding is how a phase gets past it today. Fixing one does not touch the other.

Out of scope is anything about when the check fires. #236 settled both limits and this intent changes
no condition, no loop bound and no comparison — only what three findings say to do.

No normative document is touched. Section 7 already names both routes; this makes a string agree with
it, which is the first standing rule working in the direction it usually does not get to.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the function that prints the remedies, the function that refuses them, the clause that
names what a remedy should say, and the tests that did not catch it.

`internal/gates/gates.go` is read for `staleReads` and for the three `finding` calls inside it. The
loop bound and the `touched` guard are read too, because the honest remedy depends on which findings
the check can now raise: with #236's limits in place it raises only what section 7 calls a phase
whose ground moved after it finished, and that is what makes the specification's two routes the right
instruction rather than the wrong one.

`internal/gates/staleness_test.go` is read for what is already asserted and what is not. Three tests
cover the two limits and the finding naming its phase and its file; none asserts anything about the
remedy. A string with no reader is how it came to name an act the runner declines, and it is the same
shape as several rows of `docs/clause-readers.md`.

`internal/runner/runner.go` is read for `Start` and for the comment at line 409, which is the reason
the remedy cannot be made true by doing what it says. It is also read for `predecessorAllowsStart`,
because "stopping work" in this runner means a red predecessor blocking the next phase, and that is
what an approval releases.

`docs/process-definition.md` is read for section 7's staleness clause — "A stale phase is not
deleted, it is re-run or explicitly approved as still valid" — which is the sentence the new remedies
are made of, and for the two limits whose reasons #236 implemented. Also for section 5's gate result,
where the statuses are enumerated and where an approved finding makes the phase `approved` rather
than green, so that the remedy promises a release and not a disappearance.

`docs/assumptions.md` is read for A12, the rule that the next phase starts only on a decided
predecessor, which is what makes approving the finding an act with an effect rather than a note.

`CLAUDE.md` is read for the three standing rules and for the two conventions this intent runs under.
The decision between naming section 7's routes, teaching `phase start` to re-lock, and dropping the
instruction was put as one question with each option's consequence and cost and a recommendation. And
the claim that nothing asserts the remedies was taken from a search over the test file rather than
from recalling what the tests cover, because a search for an absent assertion and a search with a
wrong pattern return the same nothing.

The refusal was reproduced rather than quoted from the issue. #237 quotes it as offering "remove
.../gate.yaml", and the text in the tree names the phase directory: #225 changed it, because from
then on an artifact left behind is itself a phase under way and removing the verdict alone is not
enough. The remedy being written had to agree with the refusal as it is now and not as the issue
recorded it.
