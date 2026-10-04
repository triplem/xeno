---
intent: github.com/triplem/xeno#217
phase: 01-requirements
created: "2026-10-04T19:09:28Z"
schema_version: "1.0"
runner_version: dev+dce3eab
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bdf2e5a91353448be6d7f0b3d6e7f667140d475784c750057f738cc4de0943ca
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
open_questions:
    - key: Q-3
      options:
        - consequence: the writer, the refusal and the fix that makes the refusal safe land in one commit, which is what this phase's own constraint on the writer and the refusal was already reaching for, and the commit convention takes two keywords for it. The cost is that re-judging P0 changes its artifacts_hash, so P1's recorded predecessor hash goes stale and needs a release — a correct finding this time, since P1 really was written against an earlier P0
          recommended: true
          text: 'Widen XENO-0245 to carry #236 as well. P0''s scope and P1''s acceptance criteria grow, both phases are re-judged, and the branch closes both issues.'
        - consequence: 'one concern per intent, as the standing rule prefers, and it is self-healing: once the first limit is in, the rebase noise that would otherwise hit this intent on a moved default branch is exactly what that limit was written to silence. The cost is a full six-phase intent for close to a two-line fix, this intent sitting open across it, and its P1 approved against documents that will have moved again'
          text: '#236 as its own intent, run first, with XENO-0245 waiting at P2.'
        - consequence: 'something lands today and nothing unsafe becomes mandatory. The cost is that this phase records the writer and the refusal shipping in one commit as a constraint, so withdrawing it means changing this phase and re-judging anyway, for less result, and #217''s done-when stays unmet — leaving a writer nothing asks for, which is how the artifact came to be specified and unwritten in the first place'
          text: 'Ship the writer now and defer the refusal to a later intent, after #236.'
        - free: true
          text: Something else, entered by the person deciding
      text: 'The specification commit that renamed the artifact moved two documents this intent''s scope names, G-Freshness read P1''s own lock, P1 went red and A12 refused P2 on a red predecessor. Section 5 states two limits on that check — only the files the change under review touched, and never the lock of the phase being gated — and staleReads honours neither, so the refusal this intent adds is not safe to land without the fix: a mandatory scope plus a check that reads the gated phase''s own lock makes every implementation intent report itself out of date for doing its job. Which intent carries #236?'
decisions:
    - chosen: 'Keep the context economy in v1, and the profile is required at P0. Section 5 stands as written: a writer produces context-profile.yaml, a P0 without one is refused, and the readers that depend on it get their first input.'
      decided_by: triplem
      id: D-1
      proposed_by: claude-opus-5
      rationale: 'The profile is not a feature of its own but the only input to context.lock.yaml''s files list, which five mechanisms read, so cutting it cuts WP8, WP15 and WP20 together. Two things decided it against the cheaper options. WP8''s own blocker, ''P0 cannot produce a profile until the format is fixed'', is already cleared by model.Profile and section 5''s example, so the package is a writer away from its done-when rather than a package away. And WP20 states that cache friendly assembly is ''the one lever that costs nothing to build and pays on every request, and it is also the one that is impossible to retrofit once the assembly is scattered across the skills'', which is the only piece of this carrying a deadline. Against making it optional: the artifact would stay specified and unproduced, so v1 would ship a context economy nobody had run, WP20''s baseline would have no denominator, and not-implemented would mislabel three checks that do run against an input that is not there. Against cutting it to 1.1: changes to section 5 and to the plan before any code, the one lever that cannot be retrofitted forfeited, and v1 shipping with no statement of what the process costs, which is the figure #117 says will decide. The cost accepted and recorded: every new P0 owes a profile, and the 100 sealed ones cannot be made to.'
      resolves: Q-1
    - chosen: 'The requirement attaches as a refusal in phase finish at P0: a P0 with no profile is declined with a reason that names the writer, and nothing is written. G-Schema keeps the profile''s content, which is the sentence section 5 gives it.'
      decided_by: triplem
      id: D-2
      proposed_by: claude-opus-5
      rationale: 'Put to the maintainer because #217''s own statement of cost was wrong, and the correction is the whole of the reason. The issue says A74 keeps sealed verdicts from being re-judged; A74 binds G-Policy alone to the rule set a phase''s artifact records and says nothing about G-Schema, while Verify recomputes every phase''s status and compares it against the committed one. So a G-Schema check re-judges 100 sealed P0 phases with a finding they do not carry, and backfill cannot answer it: PhaseExcluded holds only gate.yaml and cost.yaml, so the profile is inside artifacts_hash and one written into a finished phase reports as artifacts changed after the verdict. That is the inverse of XENO-0244, which was orderable only because evidence/attached.yaml lies in a subdirectory DirHash does not descend into. The maintainer''s requirement was that the profile be required in the future, and a refusal is the only one of the three that is forward only, because commands are not re-run where verdicts are recomputed, so gate verify stays at exit 0 over the whole trail and no historical verdict moves. Against a G-Policy rule, which A74 would genuinely cover: it puts a structural requirement in the rule set, which section 10 routes through review rather than a commit, and section 5 gives the profile to G-Schema rather than to G-Policy. Against a G-Schema finding with the 100 released one by one: 100 decisions by a second person against XENO-0244''s fifteen, 100 open obligations intent status would then owe indefinitely, and a red gate verify in between, which CLAUDE.md lists among the commands that must exit 0. The cost accepted and recorded: the requirement appears in no verdict, so a reader of a later gate.yaml sees G-Schema pass and nothing saying a profile was owed, and in the history a P0 that complied and one from before the rule are indistinguishable.'
      resolves: Q-2
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**An intent's P0 can produce a scope with a command.** The scope is written into the
intent's P0 directory, with the header fields every process file carries, from an entry
the command reads rather than from flags. The argument is `question record`'s and it
holds here for the same reason: include patterns, exclude patterns, a link per component
and two budget numbers do not fit flags without inventing a separator, and the shape
read is then the shape written. Unknown fields are refused rather than dropped, which is
the standing rule about invented fields applied where a typo would otherwise pass as
silence.

**The writer says what the patterns resolve to.** It reports the file count and the byte
total the include and exclude patterns select, using the same resolution `phase start`
will use, so that the budget is set from a counted figure. This criterion comes from
this intent's own P0: the figure was established with `wc` by hand, the first number
chosen was wrong, and the budget cannot be revised after P0 is judged.

**A scope written after its phase has a verdict says so.** The behaviour is `section
set`'s and not a new rule: the write is allowed, the phase then reports as changed after
its verdict, and the next `phase finish` judges it again. Nothing new appears on the
staircase and no second mechanism is invented for a file that happens to be YAML.

**`phase finish` at P0 refuses where there is no scope.** The reason names section 5 and
the remedy names the writer, in the voice the refusals at `runner.go:401`, `414`, `480`
and `487` already use. Exit 1, nothing written, the phase stays running.

**The refusal reaches no sealed phase.** `xeno gate verify` is at exit 0 over the whole
trail after the change with the same verdict count as before it, and the
`artifacts_hash` and committed status of every one of the 100 historical P0 phases are
unchanged. This is the criterion D-2 exists for and it is recorded as a figure before
and after rather than as a claim.

**A phase after P0 is given what the scope names.** `context.lock.yaml` for P1 onward
carries a `files` list resolved from the scope, with a path, a hash and a size each, in
the scope's own include order. Asserted on this intent's own trail, which is the first
in the repository to have one, and not only on a fixture.

**A repeated phase is told what moved.** With a scope in place, `phase start` names the
files that changed since the predecessor's lock, which is the half of WP8 that has never
had an input. A test gives it one and asserts the list.

**A declared link that names a file not in the tree is a finding.** The check exists and
has never run. A fixture with a link to a missing document is red on G-Schema and the
same fixture with the document present is green. Both directions, by test.

**An absent scope is not a finding on a phase that already has a verdict.** The
requirement is in the command and not in the gate, so recomputing any sealed phase
produces exactly the checks it produced before. Asserted by `gate verify` over the real
trail and by a fixture phase with no scope recomputing unchanged.

**The budget is not made to block this intent.** #235 is open and a budget overrun still
turns a phase red against section 5's words. Nothing here changes that, and this
intent's own declared budget keeps clear of it: 50 files and 1,000,000 bytes against 43
and 904,078 resolved.

**The rename reaches every name the artifact has.** `model.ContextProfile` and
`model.Profile` are `ContextScope` and `Scope`, the five call sites follow, this
intent's own artifact is `context-scope.yaml`, and no occurrence of the old name is left
in the tree outside `docs/v2-delta.md`, which is a different document and not in this
change. The specification commit `fbf8a72` is the cause and the code carries no opinion
of its own about the name.

**The staleness check reads only what the change touched.** A phase whose lock lists a
file the commit range did not touch is not reported, which is section 5's first limit
and the one that keeps a rebase onto a moved default branch quiet. `Base` and `Head` are
already on the gate context; the check uses them. A test gives it a lock listing two
files, touches one, and asserts one finding.

**The staleness check never reads the lock of the phase being gated.** Section 5's
second limit, and the one that blocked this intent: judging P1 reads P0's lock alone. A
test judges a phase that changed a file its own lock lists and asserts it is green,
which is the case the specification describes as "working" rather than stale.

**This intent's own P1 is green on the merits once the limits are in.** F-31d9b3 and F-03bb2f were
released only because the check read P1's own lock. With the second
limit honoured, recomputing P1 produces neither, so the release becomes unnecessary
rather than load-bearing. Asserted by `gate verify` over this intent's own trail, which
is the evidence that the fix addresses what actually happened.

**Nothing else changes.** No field of section 5, no new gate, no rule, no template, no
document, and no artifact in any sealed phase.

**The usual gates of this repository.** `gofmt` and `go vet` silent, the suite green,
the build clean, `./xeno gate verify` at exit 0, prose at 88 columns, SPDX on any new
file, and a test for the writer, for the refusal, for the resolved lock, for the
changed-file report and for the link finding.

<!-- xeno:section:non-goals -->
## Non goals

**The 100 sealed P0 phases.** Nothing backfills them and nothing can. `PhaseExcluded`
holds only `gate.yaml` and `cost.yaml`, so `context-scope.yaml` lies inside
`artifacts_hash`, and a scope written into a finished phase reports as artifacts changed
after the verdict was written. Absence and compliance therefore look alike in the
history, which is the accepted cost of D-2 and is written down rather than left to be
found.

**Making a budget overrun stop being red.** #235. The check sits in G-Schema and
`result` fails any check carrying a finding, where section 5 asks for a finding that
does not stop work. The specification wins, so the code is wrong, but the runner has no
notion of a finding that does not fail its check and the shapes available are each a
change to what a verdict means. That is WP1's argument and the process definition's, not
a call site to pick while adding a writer.

**Validating the scope's content in G-Schema.** Section 5 says it validates the scope
like any other artifact and nothing does; `gates.go:681` permits the file at P0 and no
check reads its shape. Until this intent there was no scope for such a check to judge,
and adding it in the same commit as the first two scopes would be a check whose only
subjects arrived with it. The command refusing unknown fields on the way in is the half
that belongs here. Filed as a gap.

**WP8's other half.** A read outside the scope named in the phase's digest is the second
clause of the package's own done-when and a different kind of claim: the agent's account
of what it read, where everything here is about what it was given. It also needs the
question below settled first.

**Deciding what a scope should contain so that staleness stays quiet.** P0 left this to
P2, on the assumption that an intent naming a file it then edits had a judgement to make
about what belongs in its scope. Section 5 already answers it and the answer is the
check's, not the scope's: "A phase that changes the files it read is not stale, it is
working: that is what P3 does." So there is nothing for an intent to decide and nothing
to write in the documentation; there is a check that reads one lock too many, which is
#236 and is now in scope. The assumption is withdrawn rather than carried.

**A project default for what a scope contains.** `informationBase` is explicit that the
scope is per intent, "so there is one budget per intent rather than six", so each intent
declares its own reading and there is no project-wide scope to design. A shipped example
in the intake template is WP3's question.

**A command that fixes or migrates an existing scope.** Two scopes exist in the world
and both were written in this intent. A write path into sealed phases that exists
forever to do something that has happened twice is the opposite of what the command
surface keeps narrow.

**WP15 and WP20.** The symbol index and the measured baseline are the packages this one
unblocks. Neither is started here, and the baseline in particular wants a scope that has
been used by more than the intent that invented it.

**Any change to what a verdict means.** No new gate, no new gate result, no rule, and no
finding that behaves unlike every other finding. The requirement lands in a command for
exactly that reason.

<!-- xeno:section:constraints -->
## Constraints

**The writer and the refusal ship in one commit.** A P0 required to produce something no
command can produce is worse than either half alone, and a writer with nothing asking
for its output is how this artifact came to be specified and unwritten in the first
place. Nothing in this intent may land one without the other.

**The requirement is in a command, never in a gate.** `Verify` recomputes a verdict and
compares it against the committed one, so a check added to a gate re-judges every phase
already sealed, which is what A74 was written about after the shipped rule set turned
twenty-four sealed P5 phases red. A74 is no help here: it binds G-Policy alone to the
rule set a phase's artifact records and says nothing about any other gate. A command
refusal reaches forward only because commands are not re-run.

**The writer runs before `phase finish`.** The scope is inside the phase's
`artifacts_hash`, so one written after the verdict makes the phase report as changed and
staleness the predecessor hash of every phase after it. That single fact places the
writer in the sequence and places the refusal at `phase finish` rather than anywhere
earlier, and it is why the budget has exactly one moment at which it can be set.

**What is sealed is not rewritten.** No file in any of the 100 finished intents is
touched. This intent asserts that with its own `gate verify` figures before and after
rather than reasoning from the prose, which is the practice XENO-0244 established when
the same question about `artifacts_hash` went the other way for a file in a
subdirectory.

**The documents are not editable here.** Section 5 already puts the scope at P0, already
says it is a budget P0 produces, and already says what it carries. This intent adds the
writer and the reader for clauses that have had neither, which is the shape #202's audit
asked for, and changes no sentence. Where the code and the specification disagree, as
they do on a budget overrun being red, the specification wins and the disagreement is
filed.

**One authority on the format.** `model.Scope` is it. The writer refuses what it cannot
decode into that type and the readers keep decoding into it; neither grows a field, and
neither holds a parallel idea of what a scope is.

**The refusal's wording carries its own remedy.** Each of the runner's existing refusals
names what to run next, and this one names the writer. A refusal that states a rule
without the command that satisfies it is how a person ends up hand editing an artifact,
which is the practice the last four intents have been closing.

**One dependency.** Nothing new is imported. The walk, the matching, the hashing and the
YAML are all already in the tree.

**This intent lives under its own rule.** Its P0 carries the first scope in the
repository and its P1 onward are the first locks with a `files` list, so every mechanism
it builds is exercised on its own trail before it is asserted on a fixture. Where that
produces a finding, as the staleness check will once P3 edits a file P1 was given, the
finding is correct and is answered rather than suppressed.

**Prose at 88 columns, Go at whatever `gofmt` produces.** Both rules whose reader is a
person, which #211 wrote down knowingly.
