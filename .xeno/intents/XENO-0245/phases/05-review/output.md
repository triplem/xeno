---
intent: github.com/triplem/xeno#217
phase: 05-review
created: "2026-10-04T21:30:33Z"
schema_version: "1.0"
runner_version: dev+fbf8a72.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5d55a606dca30b56b82067774643103210ecc33d28d011947aab97b82a40c5d0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Four, each naming what it departs from: the cascade names P2''s sentence calling the artifact rename free while nothing is merged, and says it is a chain the length of the phases already finished; the 88 refused tests name the design treating the fixtures as a consequence rather than a number; the staleness coverage moving to internal/gates names the package it left and why; and xeno scope set not being run against this repository''s own intent names the cascade as its reason.'
      result: met
      rule: deviations-are-traceable
    - note: 'Two exported identifiers renamed, model.ContextProfile to ContextScope and model.Profile to Scope, with five readers in this module and all in the same commit. The package is internal/, so there is nobody outside this repository to migrate. The migration that exists is the artifact''s name: a repository holding a context-profile.yaml at P0 has it ignored from this commit and its P0 then refused for having no scope. No such repository exists and the remedy is one git mv, which is why it is a note rather than code.'
      result: met
      rule: interface-change-needs-a-migration-note
    - note: go.yaml.in/yaml/v3 is still the one dependency. scope.go imports only packages already in use, and staleReads gained internal/git, which gates.go already imported for Commits.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter. What follows is the
reasoning behind each, and then the thing the checklist cannot ask about.

**`deviations-are-traceable` — met.** Four, each naming what it departs from. The
cascade names P2's sentence calling the artifact rename "free while nothing is merged"
and says what building it showed: it is a chain the length of the phases already
finished, not one re-judge. The 88 refused tests name the design's treatment of the
fixtures as a consequence rather than a number. The staleness coverage moving to
`internal/gates` names the package it left and why. And `xeno scope set` not being run
against this repository's own intent names the cascade as the reason, which is the first
deviation again, arrived at from the other side.

**`interface-change-needs-a-migration-note` — met, and the note is about the file and
not the identifiers.** Two exported identifiers were renamed, `model.ContextProfile` to
`ContextScope` and `model.Profile` to `Scope`, with five readers in this module and all
of them in the same commit. The package is `internal/`, so nothing outside this
repository can depend on them and there is nobody to migrate. The migration that does
exist is the artifact's name: a repository holding a `context-profile.yaml` at P0 has it
ignored from this commit, because the constant names the other file, and its P0 is then
refused for having no scope. No such repository exists — this one's single scope was
renamed in the same commit, and the hundred finished intents have none at all — but an
adopter mid-intent would meet it, and the remedy is one `git mv` rather than a command,
which is why it is a note and not code.

**`new-dependency-needs-a-rationale` — not applicable.** `go.yaml.in/yaml/v3` is still
the one dependency. `internal/runner/scope.go` imports `bytes`, `os`, `path/filepath`,
`sort`, `strings`, the yaml package and three internal ones, all already in use;
`staleReads` gained `internal/git`, which `gates.go` already imported for `Commits`.

**What the checklist cannot answer, and the review did.** Whether the claim this intent
rests on is true: that a requirement in a command reaches no sealed phase where one in a
gate would re-judge all hundred. It was read out of `Verify` in the intake rather than
out of the prose — recompute and compare, with A74 scoped to G-Policy alone and not to
G-Schema — and then measured at every step: `gate verify` at exit 0 over 346 verdicts
before the change, 348 after the rename and its cascade, 349 after the implementation,
350 now. A wrong reading would have shown a hundred divergences at the moment the
refusal landed, and it showed none.

**And the second thing the review checked, because this intent is the first to be able
to.** Whether the staleness fix addresses what actually happened rather than something
adjacent. P1 of this intent was released twice, by the maintainer, on findings the
inclusive loop bound raised against P1's own lock. Those releases are gone and P1 is
green with neither an approval nor an override on it, recomputed by `gate verify` rather
than asserted. That is the one criterion in P1 that could only be met by this intent's
own trail.

<!-- xeno:section:release-notes -->
## Release notes

**A context scope is written by a command, and P0 cannot be finished without one.**
`xeno scope set --intent KEY` reads the scope on stdin or from `--file`, writes it into
the intent's P0 with the runner's own header, and reports what its patterns resolve to
in files and bytes, so that the budget is set from a counted figure. `xeno phase finish`
at P0 refuses where no scope is there, naming the artifact and the command. Section 5
has required the artifact from the beginning; nothing produced it for a hundred intents.

**The artifact is `context-scope.yaml`, renamed from `context-profile.yaml`.** The
specification commit `fbf8a72` made the change and this follows it into the code:
`model.ContextScope` and `model.Scope`. A repository holding the old filename should
`git mv` it; from this commit the old name is not read.

**The staleness half of G-Freshness honours the two limits section 5 states.** It reads
only the files the change under review touched, through the commit range the run was
given, and only the locks of the phases preceding the one being gated. A phase that
changes a file it read is working rather than stale, which is the specification's own
sentence and was the opposite of what the code did. Without a commit range the half
reports nothing, so in practice it runs in CI.

**What a project sees first.** The next intent's P0 is refused until it writes a scope.
Writing one makes four mechanisms live that have had no input since they were built: the
lock's `files` list, the budget finding, the declared link check, and the report of what
changed since the predecessor's phase. A scope with no `include` is refused, because one
that resolves to nothing reaches every reader as the absence it was meant to replace.

**Set the budget from the figure the writer reports, and get it right first time.** The
scope lies inside P0's `artifacts_hash`, so once that phase is judged the budget cannot
be revised without making the verdict and every predecessor hash after it stale. The
command says so when it prints the figure.

<!-- xeno:section:residual-risk -->
## Residual risk

**The staleness half passes silently where it has no commit range, which is every local
run.** The logic is covered by four tests with a real repository and a real range, so it
is not untested; it is unexercised outside CI. A check that reports nothing because it
had no input looks exactly like one that passed on the merits, which is the defect this
whole intent was opened about, now present in a narrower place. It cannot be reported as
anything else until a check that did not run has a result, which is #235. This is the
largest residual risk here and it is named in the gate's own comment, in P2's decisions,
in P4's gaps and here.

**#235 now has three callers and no owner.** A budget overrun still turns a phase red
against section 5's explicit words; a P0 with no scope is refused by a command and
therefore appears in no verdict; and the staleness half cannot say it did not run. All
three want one mechanism, a finding that is visible and decidable without failing its
check, and that is a change to what a verdict means. Until it exists, each of the three
is a place where the trail says less than it seems to.

**The writer has never written a scope anybody kept.** Six tests cover it and the only
scope in existence was typed by hand before the command existed. The first real use is
the next intent, and the things most likely to be wrong are the ones tests do not reach:
whether the figure arrives in time to be useful, and whether a person sets a budget from
it or ignores the line.

**A scope naming files the intent will edit still produces findings in CI, correctly,
and nobody has seen one.** With both limits in place, a change that touches a file an
earlier phase read is reported against that earlier phase, which section 5 calls "the
more honest one" and intends. This intent's own CI run is the first that can produce it,
and the behaviour when it does — whether a person reads it as a defect or as the
expected consequence of implementing — is unmeasured.

**The cascade is a cost every intent that renames its own artifact will pay.** Four
phases were re-judged here, each mechanical and none needing a release, because a
re-started phase re-resolves its lock. The awkwardness is #237: the route runs through
removing a verdict, and the remedy a staleness finding prints is an act `phase start`
refuses.

**`CLAUSE-READERS.md` names symbols this change renamed**, so the audit is stale in the
way it predicted of itself. Nothing depends on it, and re-running it is its own act.
