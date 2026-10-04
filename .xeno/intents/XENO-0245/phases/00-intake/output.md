---
intent: github.com/triplem/xeno#217
phase: 00-intake
created: "2026-10-04T11:20:46Z"
schema_version: "1.0"
runner_version: dev+dce3eab
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 10be97648005d1f95c08630b9dc28fcc3dac60db91738a2f280ec4ccacb00841
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
open_questions:
    - key: Q-1
      options:
        - consequence: 'the five inert mechanisms get an input, G-Freshness''s second half begins checking a real file list, and WP20''s baseline gets a denominator to measure a phase''s reading against. The cost is a writer and roughly three issues beyond this one, every new P0 owing a profile including this intent''s own mid-flight, and a separate judgement about what xeno''s own profile should include, which #217 puts out of scope'
          recommended: true
          text: 'Keep it, and the profile is required at P0. Section 5 stands as written and WP8 is finished: a writer produces the profile, G-Schema reports a finding where a P0 has none, and a read outside the profile is named in the phase''s digest.'
        - consequence: the smallest change, no amendment to section 5's intent, and an absent input stops looking like a check that passed on the merits. The cost is that the artifact stays specified and unproduced, so v1 ships a context economy nobody has run, WP20's baseline has no denominator, and not-implemented mislabels three checks that did run against an input that was not there
          text: Keep it, but the profile is optional. A project may have none, and the three readers report not-implemented where they have no input instead of passing.
        - consequence: 'the v1 surface shrinks to the process and the trail, which 96 intents have shown works without any of this. The cost is a change to section 5 and to the plan, both of them a person''s commit before any code; cache-friendly assembly forfeited by WP20''s own argument that it is the one lever impossible to retrofit once the assembly is scattered across the skills; and v1 shipping with no statement of what the process costs, which is the figure #117 says will decide'
          text: 'Cut it from v1. WP8, WP15 and WP20 move to 1.1 beside WP18, section 5''s context economy becomes a statement of intent for 1.1, and #217 closes as out of scope.'
        - free: true
          text: Something else, entered by the person deciding
      text: 'No intent in 96 has written a context-profile.yaml, so the profile''s three readers return nil and the lock''s files list is empty in every one of 321 verdicts: the second half of G-Freshness iterates an empty list, the budget is judged against nothing, no declared link is checked, and ChangedSince has no base. The absence looks exactly like a pass. Before deciding how absence becomes visible, the prior question is whether the context economy is something v1 means to have at all, since 96 intents have run to completion without any of it.'
    - key: Q-2
      options:
        - consequence: 'the requirement binds every new P0 and reaches no sealed one, because Verify recomputes gate status and a refusal is not a gate, so gate verify stays at exit 0 throughout and no historical verdict moves. The cost is that the requirement appears in no verdict: a reader of a later gate.yaml sees G-Schema pass and nothing saying a profile was owed, and in history a P0 that complied and one from before the rule are indistinguishable'
          recommended: true
          text: 'A refusal in phase finish at P0, where #216 and #218 already refuse. A P0 with no profile is declined with a reason and nothing is written; G-Schema keeps the profile''s content, which is the sentence section 5 gives it.'
        - consequence: the requirement is in a verdict, and A74's stated hole — a rule added after an artifact was written leaves it unjudged — is the wanted behaviour here rather than a gap. The cost is a structural requirement in the rule set, which section 10 routes through review rather than a commit; section 5 gives the profile to G-Schema and not to G-Policy, so it bends a gate's remit; and the standing rule against invented rules makes it a specification question before it is a code one
          text: A G-Policy rule, so that A74 applies for real and only phases whose rules_hash records the new set are judged.
        - consequence: one uniform check at the place section 5 implies, with no cutoff and no second mechanism. The cost is 100 decisions by a second person against XENO-0244's fifteen, 100 open obligations that intent status then owes indefinitely, and a red gate verify in between, which CLAUDE.md lists among the commands that must exit 0
          text: A G-Schema finding, and the 100 sealed P0s are released one by one with gate approve or gate override, naming a person and a reason.
        - free: true
          text: Something else, entered by the person deciding
      text: 'Given that the profile is required, where does the requirement attach? #217 assumes A74 shields the trail, and A74 binds G-Policy alone to the rule set a phase records and says nothing about G-Schema. Verify recomputes a verdict and compares it against the committed one, so a check added to a gate re-judges 100 sealed P0s with a finding they do not carry, and backfill cannot answer it: PhaseExcluded holds only gate.yaml and cost.yaml, so the profile is inside artifacts_hash and one written into a finished phase reports as artifacts changed after the verdict.'
---

# Intake

<!-- xeno:section:problem -->
## Problem

Section 5 enumerates `context-scope.yaml` among P0's artifacts, and says under context
economy what it is for: "The context scope is a budget. P0 produces it as an artifact of
its own, and a phase reads what it names." It is not something a project opts into.

**Nothing has ever produced one.** `find .xeno -name context-scope.yaml` returns nothing
across 100 finished intents and 345 verdicts, and the artifact has no writer: `section
set` writes sections of `output.md`, and the scope is a separate file with a shape of
its own. This is the family #188, #195, #208 and #220 closed four times over — a
specified artifact no command writes — and the worst member of it, because the other
four turned a phase red while this one passes.

**Five readers, all of them inert.** The scope is the only input to
`context.lock.yaml`'s `files` list, which `runner.go:426` fills from it, and everything
downstream reads that list:

- `informationBase` (`runner.go:525`) resolves include and exclude into a path, a hash
  and a size per file. It returns nil, so `files` is absent in all 346 locks.
- G-Freshness's second half (`gates.go:1148`) would report that a phase was given a file
  and the file has changed since. It iterates an empty list, for every phase of every
  intent.
- `budget` (`gates.go:492`) would report a recorded context that exceeded its declared
  number. There is no number, in 345 verdicts.
- `links` (`gates.go:455`) would report a declared code to documentation link naming a
  file that is not in the tree. No link has been declared.
- `ChangedSince` (`runner.go:191`) would tell a repeated phase which of its inputs
  moved. It reads the predecessor's empty `files` list and returns nothing, so `phase
  start` prints nothing, and a repeated phase is told to re-read everything by being
  told nothing.

**The absence reads as a pass.** Three of the five return nil on the missing file, each
with a comment saying as much: "no scope is not an error; it is a project that has not
written one", "no scope is no link", "no scope is no budget". Each sentence is true and
none of them reaches a verdict. A74's distinction is the one that applies and the one
the readers cannot draw: an empty list says a set was resolved and came out empty,
absence says there was nothing to resolve. G-Schema's only mention of the scope is
tolerance — `gates.go:681` permits the file at P0 among `KnownPhaseFiles` and nowhere
requires it — so a P0 without a scope is not merely unreported, it is admitted.

**The statement already exists, as prose.** Every one of the 100 intakes carries a
`context-rationale` section naming the files the intent read and why, and every design
and implementation phase after it works from that reading. The scope is the same
statement in a form something can act on: it is per intent, not per project, which
`informationBase` is explicit about — "a phase does not get a scope of its own, so there
is one budget per intent rather than six". So what is missing is not a judgement nobody
has made. It is the machine-readable half of one made 100 times.

**Why it was never written.** WP8 names its own blocker: "P0 cannot produce a scope
until the format is fixed." The format is fixed. `model.Profile` carries include,
exclude, links and budget against section 5's example, G-Schema validates it like any
other artifact, and the two halves of the package that depend on it — the budget check
and the link check — were built in #176 and #172 against a file that never arrived. The
blocker was cleared and the writer was never added behind it.

<!-- xeno:section:scope -->
## Scope

**In scope.** Five things, in the order they have to happen.

**The rename follows the specification commit.** `fbf8a72` renamed the artifact from the
context profile to the context scope, for the reason that commit gives: profile names a
set of settings one switches between, an intent has exactly one, and three of the file's
four fields delimit a scope while the fourth bounds it. The code follows —
`model.ContextProfile` and `model.Profile` become `ContextScope` and `Scope`, with the
five call sites at `runner.go:527`, `gates.go:455`, `492`, `681` and `model.go:146` —
and so does this intent's own artifact, which is the only one in existence. It goes
first because everything after it is named for the file, and the window is this intent:
once a merged intent carries a scope, the old name sits inside a sealed
`artifacts_hash`.

**A writer.** A command that produces `context-scope.yaml` in the intent's P0 directory.
It has to run before `phase finish` at P0, because the scope lies directly in the phase
directory and `PhaseExcluded` holds only `gate.yaml` and `cost.yaml`, so the scope is
inside `artifacts_hash`: one written after the verdict makes the phase report as
changed. That one fact fixes the writer's place in the sequence and the refusal's with
it.

**This intent writes its own, by hand, before its own P0 is finished.** The writer does
not exist while P0 runs and the scope cannot be added afterwards for the reason above,
so XENO-0245's is hand-written — which #217 anticipated as "a decision that writing it
by hand is the intent". It is the first scope in the repository, so P1 to P5 of this
intent carry the first non-empty `files` list in 346 locks, and the budget check, the
link check and G-Freshness's second half get their first live input on the intent that
gave them one.

**The staleness check honours the two limits section 5 states (#236).** It reads only
the files the change under review touched, through the `Base` and `Head` the gate
context already carries, and only the locks of the phases preceding the one being gated.
Both limits are in the specification with their reasons, and `staleReads` honours
neither, which blocked this intent at its own P1 the moment the specification commit
moved two documents its scope names. It is in scope because the refusal is not safe
without it: a mandatory scope and a check that reads the gated phase's own lock together
make every implementation intent report itself out of date for doing its job.

**`phase finish` at P0 refuses where no scope is there.** Forward only and deliberately
so. `Verify` recomputes a verdict and compares it against the committed one, so a gate
reaches backwards into 100 sealed P0s while a command refusal does not, because commands
are not re-run. D-2 records that choice against the two alternatives.

**Out of scope, each for its own reason.**

**The 100 sealed P0s.** Nothing backfills them, and nothing can: the scope is inside
`artifacts_hash`, so writing one into a finished phase reports as artifacts changed
after the verdict was written. Absence and compliance therefore look alike in history.
That is the accepted cost of D-2 and it is written into the deviations rather than left
to be found.

**Validating the scope's content.** Section 5 says G-Schema validates the scope like any
other artifact and nothing does; `gates.go:681` permits the file at P0 and no check
reads its shape. It is a judgement about a file's contents rather than its presence, and
until this intent writes one there is nothing for such a check to judge — adding it in
the same commit as the first scope would be a check whose only subject arrived with it.
Recorded as a gap and filed.

**WP8's other half.** A read outside the scope named in the phase's digest is the second
clause of the package's own "done when", and a different kind of claim: the agent's
account of what it read, where everything here is about what it was given.

**A project default for what a scope should contain.** #217 puts it out of scope and the
reason holds — `informationBase` is explicit that the scope is per intent, "so there is
one budget per intent rather than six", so each intent declares its own reading and
there is no project-wide scope to design. A shipped example in the intake template is a
WP3 question.

**WP15 and WP20.** The symbol index and the measured baseline are the packages this one
unblocks, not this one.

**The remedy a staleness finding prints is an act the runner refuses (#237).** The
finding says to read the phase again; `phase start` declines a judged phase, and the
`section set` and `phase finish` it offers instead never re-resolve the lock, which #215
made deliberate. So the finding cannot be cleared by doing what it asks, and the only
routes are a release by a second person or deleting the verdict. The honest wording
depends on what #236 leaves behind, so it is filed and not answered here.

**A budget overrun turns a phase red, and section 5 says it must not.** The check sits
in G-Schema (`gates.go:544`) and `result` makes any finding a `fail`, so a recorded
context over its budget stops the phase. Section 5 is explicit against exactly that:
"deliberately a finding and not a red gate in the sense of stopping work", because
"blocking against a number nobody has experience with yet would be the wrong way round".
The specification wins, so the code is wrong — but the runner has no notion of a finding
that does not fail its check, which is a mechanism and not a line. It is #235 rather
than absorbed here, and this intent keeps clear of it by declaring a budget with
headroom.

**The budget is frozen when P0 is judged.** The scope is inside P0's `artifacts_hash`,
so revising it afterwards makes P0 report as changed after its verdict and staleness the
predecessor hash of every phase after it. An intent that finds its budget too tight at
P3 cannot correct it. This intent met that at P1 and corrected the number before
anything was sealed; that the correction has exactly one moment is a property of the
design and is recorded here because the writer's own interface has to account for it.

**What this intent owes its own trail.** `gate verify` at exit 0 before the change and
after it, over the same count of verdicts each time, recorded as two figures rather than
one claim; the figure before is 346 verdicts at exit 0. The budget its own scope
declares is counted and not the round number in section 5's example: 43 files and
904,078 bytes resolved, declared as 50 and 1,000,000 for the writer and its tests.

<!-- xeno:section:context-rationale -->
## Why this context

#217 is read as it stands, with the two decisions taken in the session that opened this
intent and reproduced here as Q-1/D-1 and Q-2/D-2 rather than referenced, because an
issue thread is not an artifact and the trail has to carry why. One of the issue's own
claims is not carried over: it says A74 leaves the old verdicts as they are, so that
only new phases are affected. A74 is scoped to one gate, and that is checked below
rather than taken on the issue's word.

Section 5, three passages. The artifact list, for the sentence that puts
`context-scope.yaml` at P0. The context economy subsection in full, for what the scope
is a budget of, the five mechanisms, the order of volatility that the lock records
rather than only the set, and the sentence that gives G-Schema the budget finding and
says why it is a finding and not a red gate. And the phase table's P0 row, which names
the scope among what intake produces, so that its absence is not a reading of one
sentence.

`ASSUMPTIONS.md`, A74 in full and not in summary, because the cost of the chosen option
turns entirely on its scope: it binds G-Policy to the rule set a phase's artifact
records, and it says nothing about any other gate. A6 and A10 are read beside it as the
two rows that scheduled this work to WP8 and record what the lock carried before #63 and
after it.

The plan's WP8, for the package's own "done when" and for the blocker it names — "P0
cannot produce a scope until the format is fixed" — which `model.Profile` and section
5's example have fixed, and which is why the writer was never added behind it. WP20 is
read for the baseline this unblocks and for its claim that cache friendly assembly is
the one lever impossible to retrofit, which is what put the prior question to a person
at all.

`internal/runner/runner.go`, for five things. `informationBase`, which is the resolution
and carries the sentence settling the scope as per intent rather than per phase. The
`phase start` path, where the lock's `files` list is written from it. `ChangedSince`,
which reads the predecessor's list and is the mechanism a repeated phase depends on.
`Verify`, whose recompute and compare is the whole reason the requirement cannot be a
gate. And the refusals at 401, 414, 480 and 487, whose shape and voice the new one
follows rather than invents.

`internal/gates/gates.go`, for `links`, `budget` and the second half of `freshness` —
the three readers that return nil and the one that iterates an empty list — and for line
681, G-Schema's only mention of the scope, which permits the file at P0 and nowhere
requires it. That line is the difference between an artifact that is tolerated and one
that is owed.

`internal/hashing/hashing.go`, for `DirHash` and `PhaseExcluded`. Whether the scope lies
inside `artifacts_hash` decides two things at once: that backfill is impossible, and
that the writer has to run before `phase finish`. A prose reading of Appendix B is not
enough to rest both on, which is the lesson XENO-0244 recorded when the same question
went the other way for a file in a subdirectory.

`internal/model/model.go`, for `Profile`, `ContextLock.Files` and `ContextFile`, which
are the format that is already fixed, and for `MatchPath`, because the include patterns
this intent writes into its own scope have to match what the walk will actually match.

The trail is counted rather than sampled: 101 intent directories, 0 scopes, 345
verdicts, 346 `context.lock.yaml` and not one with a `files` list. #217 reports 96 and
321, which is what the figures were when it was written.

Nothing outside the repository is needed. The issue, three passages of section 5, three
rows of the assumption register, two work packages and five source files are the base.
