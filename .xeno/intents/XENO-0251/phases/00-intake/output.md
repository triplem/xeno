---
intent: github.com/triplem/xeno#212
phase: 00-intake
created: "2026-10-05T15:47:42Z"
schema_version: "1.0"
runner_version: dev+09e2aa6
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3554d5f57bd958789fe8d158e61a3c94fcfdb92d9a9dd150138adae660c13a9d
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

G-Test is `{"G-Test", 4, notImplemented}` and has been a written-but-unjudged verdict in every
P4 of this trail. Section 7's row asks it for two things: "declared test result successful,
mapping of acceptance criteria complete".

It is not blocked the way G-Secret is. G-Secret waits on a harness hook from WP11 that does not
exist; the plan says G-Build and G-Test "read declared results rather than running builds or
suites themselves", and G-Build works.

The templates were shaped around it. The plan says `acceptance-criteria` exists because G-Test
maps onto it and `test-mapping` because G-Test checks that mapping for completeness. Both
sections are required, both are written in every intent, and nothing has ever read either.

#212 put the two halves the other way round from how they measure, and both of its premises
are now stale. It says the mapping half is reachable and the result half is blocked by #208.

**#208 is closed**, so declared evidence has a writing command, and the trail holds twenty
declarations of `kind: test-report` across eighteen P4 artifacts. A G-Build mirror over that
kind was built as a probe and run: `gate verify` stays at exit 0 over 381 verdicts, because the
pending-and-attached path already carries the results. So the result half costs nothing against
history.

**The mapping half is the unreachable one.** A gate would have to identify an acceptance
criterion to check that P4's mapping covers it, and nothing identifies one: of 55 P1 artifacts,
46 have no numbered criteria at all, and the nine that do are the most recent. Even one of those
nine, XENO-0248, would fire — its criterion 4 is not mentioned in P4's mapping by number. So a
mechanical completeness check re-judges most of the trail and needs a numbering convention that
does not exist, which is an addition to the documents before it is anything else.

Both figures were measured rather than reasoned about, which is what inverted the issue.

<!-- xeno:section:scope -->
## Scope

In scope is the result half of G-Test: a declared `test-report` whose result is not `pass` is a
finding, with the pending-and-attached path G-Build already has, so an item awaiting its
artifact is `pending` rather than failed. G-Test stops reporting `not-implemented`.

In scope is sharing the comparison with G-Build rather than copying it. The two differ by one
constant, and two near-identical gate bodies free to drift is what `BuildKind` exists to
prevent one level down.

In scope is `TestKind` as a named constant beside `BuildKind`, for the reason that one records:
the gate read `build` for as long as nothing checked the closed set and matched no conformant
declaration the whole time.

In scope is a row in `docs/clause-readers.md` saying the mapping half has no reader, which is
what that document is for, and the count in its table moving from 35 to 36. The file's own
precedent is followed: it already carries one row that arrived after the pass, left visible
rather than absorbed, and says why.

In scope is being explicit that a green G-Test judged one of the two things section 7 asks it
for. That is the cost of this change and it is the reason the audit row is part of it rather
than a follow-up: a gate that goes from `not-implemented` to `pass` having read half its row is
the "green verdict meaning less than it appears to" this repository keeps naming, and the only
honest mitigation available is to write down which half.

Out of scope is the mapping half. It needs a way to identify an acceptance criterion, which is a
convention the documents do not have, and a check over 46 artifacts that never had one. Its own
issue, with the measurement.

Out of scope is a numbering convention for acceptance criteria. Requiring criteria to be a
numbered list is an addition to what section 9 says, so a specification change first and a
person's commit by the first standing rule.

Out of scope is running a test suite. The plan is explicit that G-Build and G-Test read declared
results rather than running anything, and the gate path makes no network call and starts no
process.

Out of scope is a second gate. The gate list is a budget by section 7, and the two halves are
one row of it.

No normative document is touched. Section 7 already asks for the result half and this intent
implements it, so the code moves towards the specification and no specification commit precedes
this.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the gate table, the gate this one mirrors, the kinds it reads, and the trail it
would be applied to.

`internal/gates/gates.go` is read for `build` and for `BuildKind`. The first is the shape the
result half takes, down to the pending-and-attached handling that turns out to be why history
survives the check; the second is a comment about a defect — the gate matching no conformant
declaration for as long as nothing checked the spelling — which is the argument for naming
`TestKind` rather than writing the string twice.

Both normative documents are read for section 7's row and for the plan's sentence about reading
declared results rather than running suites, and for section 9's sixth clause, which bounds the
mapping half to completeness, "not selection, and not existence". Reading the last one is what
settles that the mapping half is about criteria being covered rather than tests being real, and
therefore that identifying a criterion is the whole difficulty.

`docs/clause-readers.md` is read for its table's format, its count, and its own handling of a
row that arrived late. It gives the precedent this intent follows and the reason to follow it:
a pass that quietly absorbed its own miss would be the thing #202 warns about.

`internal/gates/evidence_test.go` is read for how a declared and attached evidence item is set
up in a fixture, because the result half's tests need the same shapes G-Build's do and a second
way of building them would be a second account of what an evidence item is.

The trail is read twice by script and once by probe, because this intent's shape turns entirely
on two counts. The first found twenty `test-report` declarations across eighteen artifacts with
seventeen carrying no result in the declaration, which looked like the result half was
unimplementable; the probe then showed `gate verify` at exit 0, because those items are pending
with attached results, and the declaration's own `result` was the wrong field to count. The
second found 46 of 55 P1 artifacts with no numbered criteria. Without the probe this intent
would have implemented the mapping half and broken the trail, which is the inverse of what the
issue proposed and of what the first script suggested.
