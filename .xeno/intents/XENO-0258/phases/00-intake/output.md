---
intent: github.com/triplem/xeno#256
phase: 00-intake
created: "2026-10-06T08:34:02Z"
schema_version: "1.0"
runner_version: dev+5163d1b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8e8a33c45af8d19a088cd51c8a17c446489e675d7130bd8c93159b80ef470454
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

A finding written into a phase's artifact is sealed with it, and a false one cannot be corrected
where it was made. This session produced one and it propagated further than it should have.

#201's verification and review phases recorded that `internal/plugin/embedded/plugin` is not
gitignored. It is, at `.gitignore:10`, added by #200 with a comment. The check ran
`git check-ignore` against the path moments after the rehearsal copy had been deleted, and
`git check-ignore` reports nothing for a path that does not exist — so an absent directory read as
an absent rule. The claim reached a verification phase, a review phase, a pull request and two
reports before anything tested it with the path there. XENO-0249's P3, P4 and P5 carry it
permanently; the correction lives in #254, in a comment on PR #246 and in this intent, none of
which is where somebody reading XENO-0249 is looking.

**The class is wider than one command and the tell is the same every time.** A tool asked about
something that is not there answers the way it answers about something that does not match:
`git check-ignore` exits 1 with no output for both, `test -f` is false for both, `grep -l` and
`find` print nothing for both. The dangerous moment is immediately after a cleanup step, which is
exactly when a rehearsal or a scratch copy has just been removed and exactly when somebody is
writing up what they found.

A second instance in the same session, different mechanism and the same shape. #212's intake ran a
script reading each declared `test-report`'s own `result` field, found seventeen of eighteen empty,
and concluded the result half of G-Test could not be implemented without re-judging the trail. The
opposite was true: those items are pending with results in `evidence/attached.yaml`, which
G-Build's existing path already reads. The script's "nothing" meant "not in the field I read" and
was taken for "nothing anywhere". That one was caught because the next step was to build the gate
and run `gate verify`; the gitignore one had no next step to catch it.

**What makes this this project's problem rather than general advice** is where a finding goes.
Section 11's rule is that what is sealed is never rewritten, and a finding belongs in an artifact,
so an unverified negative does not become a mistake somebody corrects — it becomes a permanent
sentence in the trail with the correction somewhere else. That is a property of this process and
not of investigation at large.

It is recorded as XENO-0254's P0 learning, which is why nothing has happened to it: section 10
routes a learning through a merge request against the rule set after review, and nothing has
proposed one. #229's and #212's learnings sit in the same place and nothing reads those either.

<!-- xeno:section:scope -->
## Scope

In scope is one paragraph in `CLAUDE.md`, under Conventions, beside the one about putting a decision
one at a time. It says that a negative result from a check over a path or a tree is evidence only
when the thing checked exists, and it says why that matters here rather than in general: a finding
goes into an artifact that is sealed, so an unverified negative becomes permanent.

In scope is the framing being about sealing. The generic version — verify your checks — is advice
about investigation and does not belong in a file carrying this project's context; what belongs is
the consequence this process has and others do not, which is that the mistake cannot be corrected
where it was made.

In scope is the gitignore instance by name, because a convention with an instance is checkable and
one without is a preference. The second instance, #212's field read, stays in the issue: one example
earns the lines and two is the issue's job.

Out of scope is a mechanism. Nothing can check this — a gate would have to know which of a
command's two silences it received, and `docs/clause-readers.md` is full of rows recording that a
clause's reader is a person. This joins them knowingly and the paragraph says nothing checks it, as
the sequencing paragraph does.

Out of scope is a row in `docs/clause-readers.md`. That document is a pass over the two normative
documents and a convention in `CLAUDE.md` is a clause in neither — the same reason #255 declined a
row for the sequencing convention, and declining it again keeps that consistent.

Out of scope is repairing XENO-0249. What is sealed is never rewritten; the correction already sits
in #254, in a comment on PR #246, and now here, and #254's residual risk records that a reader of
that intent meets the claim with nothing beside it.

Out of scope is a general convention about tools. `test -f`, `grep -l` and `find` share the
ambiguity and naming all of them would turn a paragraph into a list; the paragraph names the class
and one instance, and the issue holds the table.

No normative document is touched. `CLAUDE.md` is the conventions file — #210 changed it, #259 added
to it — and the first standing rule binds the process definition and the plan, neither of which says
anything about how a finding is arrived at.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the file being written into, the rule that makes the mistake permanent, and the
precedent for a convention whose reader is a person.

`CLAUDE.md` is read for its length and its last line, because this intent adds to a file asking not
to be added to. It is 87 lines and nine of them arrived in #259 this session. That is the argument
the design has to answer: whether this paragraph earns its place, and the test is whether a session
reading it before investigating would behave differently — which for the gitignore instance it
would, because the check that produced the claim took three seconds to redo.

`docs/process-definition.md` is read for section 11's sealing rule and for where a finding lives,
because the whole case for putting this in `CLAUDE.md` rather than leaving it as general advice
rests on the consequence being specific to this process. If a finding were a comment on a pull
request the lesson would be ordinary carefulness; it is a sentence in an artifact covered by
`artifacts_hash`, and that is what makes it a convention.

`docs/assumptions.md` is read for A90, which is the row this joins: a reader that cannot fail is
worse than none, which is why the paragraph says nothing checks it rather than proposing something
that would.

`docs/clause-readers.md` is read as the precedent and as the boundary. It is the project's catalogue
of rules whose reader is a person, which is what this becomes; and it is a pass over the two
normative documents, which is why this gets no row there — the same line #255 drew for the
sequencing convention, and drawing it the same way twice is what keeps it a rule rather than a
preference.

`docs/implementation-plan.md` is read to confirm nothing in it governs how a finding is reached, so
no specification commit precedes this.

The instance itself is read from the trail rather than recalled: XENO-0249's P3 deviations, P4 gaps
and P5 residual risk all carry the false claim, and `.gitignore:10` carries the entry that disproves
it. Both were checked again while writing this, because a paragraph about verifying a negative
should not rest on a remembered one.
