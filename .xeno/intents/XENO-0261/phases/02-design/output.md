---
intent: github.com/triplem/xeno#237
phase: 02-design
created: "2026-10-06T12:03:51Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2b29bdf50a360308cd85e7e3b79b7b9836c914a164bd386bce9f5910c4f3cebb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: The two staleness remedies about a sealed lock name section 7's two routes — starting the earlier phase over, and a second person approving the finding on the phase being gated as still valid — each with the reason section set and phase finish will not clear it. The unreadable-file remedy keeps make it readable.
      rationale: 'Section 7 already enumerates both routes, so this makes a string agree with the specification rather than asking the specification to change, and both are acts the runner carries out. The approval is phrased as a second person''s act and not as a command, because Runner.red in next.go already prints gate approve and its comment says it offers neither way out as a command: releasing a finding is a second person''s statement. A reader of gate.yaml has no suggestion beside the field, which is why the route is named in the string at all. The alternatives were teaching phase start to re-resolve a lock, which contradicts #215 and needs a section 11 commit, and dropping the instruction, which replaces one exception to #216 and #218 with another.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The re-run comes first and the approval second.** Section 7 lists them in that order — "re-run or
explicitly approved as still valid" — and the order is also the honest one: the re-run is the act
that makes the staleness untrue, and the approval is the act that accepts it. A remedy leading with
the approval would read as "this is fine", which is a judgement the remedy is not entitled to make.

**The approval names who does it rather than what to type.** "or a second person approves it on
`<gated phase>` as still valid." `Runner.red` already prints `gate approve` and its comment says why
it offers neither way out as a command: releasing a finding is a statement by a second person, and a
runner that put the command in the reader's hand would nudge towards the decision it exists to
record. Saying "a second person approves it" is true for a reader of `gate.yaml`, who has no
suggestion beside the field, and is not the nudge that comment refuses.

**Each remedy says why the obvious route fails, in the clause that follows it.** "`section set` and
`phase finish` will not clear it: the lock keeps what the phase was given." Without that, a person
reaches for the first route the refusal offers, runs two commands and finds the finding unchanged,
which is the second of the two refusals the issue counts.

**The re-run route says what it discards.** "starting `<earlier phase>` over discards it and every
phase after it." It is the sentence that stops this being a smaller version of the same fault, and it
is also the reason the approval is worth naming at all: without the cost stated, a reader would take
the re-run as the cheap option.

**The unreadable case keeps its remedy and gains a reason.** "make it readable and run the gate
again; the tree is wrong here, not the lock." It is the one of the three that already worked, and the
added clause is for the reader who meets all three and would otherwise read the difference as an
oversight. It also stops a later pass "fixing" it into the other two.

**The tests assert the routes and the phases, not the strings.** Each of the three gets a test naming
what its remedy must mention — the earlier phase, the gated phase, that `phase finish` does not clear
it, that the re-run discards. A test matching the whole literal would break on every rewording and
pass on a route silently dropped, which is the wrong way round for a string whose fault was that
nothing read it.

**Nothing in the check's conditions is touched.** The changes are inside the three `finding` calls.
That is a deliberate boundary rather than a happy accident: #236 settled when the check fires, and an
intent that adjusted both would make it impossible to tell which change caused a later difference in
behaviour.

<!-- xeno:section:alternatives -->
## Alternatives

**Teach `phase start` a flag that re-resolves the lock.** The only route that would make the remedy
as written true, and the only one that offers a stale phase a refresh rather than a choice between
accepting and discarding. It rewrites an artifact inside `artifacts_hash` and destroys the only
record of what the phase was given, which are #215's two reasons and are restated in the code at the
refusal. Reopening a closed decision is a person's and a section 11 commit precedes it. Rejected as
out of scope, not as wrong: it is the one shape that would give the process a third verb.

**Drop the instruction and state only the fact.** The cheapest, and it stops the remedy being false.
Every other finding and refusal in this runner carries a remedy that works, so this makes a new
exception rather than closing one, and it leaves a person to discover the release route — which is
the fault reported. Rejected.

**Name `gate approve` with its flags in the remedy.** It is what a reader of `gate.yaml` would most
like, and it is what `next.go` deliberately declines to do a few lines away. Two parts of one runner
disagreeing about whether to hand over that command is worse than either answer, and the function
that declined it gave a reason this intent has no standing to overturn. Rejected, and the phrasing
"a second person approves it" is what the rejection produced.

**Put the release first, as the primary route.** It is what a person will almost always do, and #237
itself suggests it — "either the remedy names the release as the first route". Section 7 lists the
re-run first, and a remedy that leads with the approval tells a reader the staleness is acceptable
before they have looked. Rejected on order rather than content; both routes are named either way.

**Give all three findings the same remedy.** Symmetrical and shorter. It would replace a remedy that
works with one that asks a second person to release a finding caused by a file permission. Rejected:
the issue is about a remedy naming an act the runner refuses, and this would create one naming an act
nobody needs.

**Fix every remedy in `gates.go` in one pass.** Several others may name acts that do not work;
nothing has checked. It is #202's shape of work — a pass that finds eight things of one kind — and
#202's own finding is that such a pass costs about an intent each and one that fixed as it went would
stop at the first. Rejected as an intent of its own, and worth filing.

**Make the remedy a structured field rather than text.** A route the runner could offer as a command,
read by `next.go` instead of composed into prose. Section 5 enumerates `next: <text>` and nothing
else, so it is a new field and a specification change, for a benefit three strings do not justify.
Rejected.

<!-- xeno:section:impact -->
## Impact

**`internal/gates/gates.go`, inside `staleReads`.** Three `finding` calls. The two about a sealed
lock gain the earlier phase's name, the gated phase's name, the reason `phase finish` will not clear
the finding, and what the re-run discards. The third keeps "make it readable" and gains a clause
saying which kind of fault it is. No change to the loop, the guard, the hashing, the causes or the
sort.

**`internal/gates/staleness_test.go`.** Three tests added, one per case, asserting what each remedy
names. The existing three are untouched, so the two limits #236 implemented keep the readers they
were given.

**What a person meets now.** A stale finding that names the two acts section 7 names, says which
phase each applies to, and says why the route the refusal offers first does not work. The two
refusals the issue counts become none: following either named route reaches an act the runner carries
out, and following the approval route clears the finding, which is the issue's "Done when".

**What a reader of `gate.yaml` meets.** The same sentence, with no suggestion beside it. That reader
is why the approval is in the string at all rather than left to `next.go`, and why it is phrased as a
second person's act rather than as a command.

**What gets longer.** Two remedies that were one clause each are now three. A remedy is read once by
somebody who is stuck, which is the case that justifies the length; the alternative was a short
remedy that sent them to a refusal.

**No normative document, no field, no gate, no template, no rule.** Section 7 already names both
routes. `model.Finding` is unchanged, so `gate.yaml`'s shape is unchanged and nothing in the trail
re-hashes.

**Sealed artifacts are unaffected.** No existing `gate.yaml` carries a staleness finding — the check
has produced one only inside XENO-0245, which #236 was filed from — so nothing in the trail holds a
remedy that this makes historical. A finding already written would keep its old text, correctly:
`next` is sealed with the verdict it belongs to.

**Adopters get it with the runner.** This is `internal/gates`, so it ships in the binary rather than
in the plugin, and an adopter sees the new wording on the next release with no action of their own.
