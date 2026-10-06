---
intent: github.com/triplem/xeno#260
phase: 05-review
created: "2026-10-06T09:25:32Z"
schema_version: "1.0"
runner_version: dev+081da51.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f0e946c06da79257f2bad15865bf4d8ffa6108c6e6b1aee408a794f6a6c30edd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: deviation
      note: Four, all recorded in P3 and all caught by running something rather than by reading. (1) --omit=dev where the action passes --only=prod, named in P1's constraints and decided in P2; both flags were run against the same checkout and produced the same 28 advisories in the same split. (2) low rises 0 to 1 and P2's impact section wrote the change as 0/10/2/0 to 2/23/2/1 without remarking that an empty severity class acquires an entry; the advisory is diff, the check is unaffected, and the commit message names all four classes rather than three. (3) the tree is checked out into the workspace rather than /tmp, because actions/checkout refuses a path outside it. (4) the sha's sed pattern was first written for '- uses:' and release.yml carries a bare 'uses:'; running both against the file is what found it, and an empty ref would have checked out a moving default branch. No criterion was falsified; criterion 10 is open and P4's gaps say so.
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: 'No interface changes. The diff is one CI workflow, one JSON baseline and two documents; no Go source, no command, no flag, no artifact field, no gate, no template and no rule. Nothing shipped by xeno init --vendor is touched, so no adopter sees anything: the audit job is this repository''s own CI. What changes for a maintainer is which tree a number describes, and the baseline''s comment and the job''s header both say so where they are read.'
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: None added. go.mod is untouched and the one vendored dependency is unchanged. The job does now check out a second repository, which is not a dependency of this project but is a second pinned reference in its CI — and it is the reference release.yml already pins, read out of that file rather than written again, so the count of identifiers this repository maintains does not rise. The 533-package Node tree is installed for the length of the job and shipped by nothing, which is the distinction docs/supply-chain.md draws between what the pipeline runs and what the release carries.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched: `docs/process-definition.md` and
`docs/implementation-plan.md` are unchanged and neither says which tree a CI job installs, so the
first rule is not engaged. Nothing is invented: no field, no gate, no tool, no rule — the second
standing rule's budget is untouched, and `docs/supply-chain.md` and `docs/assumptions.md` describe
the repository rather than specify the process. The work belongs to #260, labelled `wp0`, on a
branch carrying one intent, and the commit references it.

**The decision was put one at a time, with its consequences and a recommendation.** #260 offered
three routes. The measurement found a fourth and retired one of the three, and the question put was
one question with three options, each carrying what it costs, and the recommendation named with its
reason. The `CLAUDE.md` convention asks for exactly that, and XENO-0243's failure mode — later
questions assuming an answer to the first — did not arise, because the other two decisions this
session holds were not asked alongside it.

**Every negative result was re-measured with the thing present.** This is the convention #263 added,
and this intent is the first to run under it. `handlebars` being absent from the audit's tree,
`@semantic-release/changelog` being absent from it, the 531-of-532 identity with the action's
lockfile and the equality of `--only=prod` and `--omit=dev` were each taken from a tree that had been
installed, not from reading a manifest or a workflow. The `- uses:` sed pattern that matched nothing
is the clearest instance: it would have reached CI as an empty ref, and the only thing that found it
was running it against the file.

**The claim being repaired was itself an unverified negative, which is why it survived.** The audit's
header asserted that its install matched the release's, and nothing could see otherwise from inside
this repository: the two workflows read consistently, because one of them delegates. That is the P0
learning and it generalises past this job — a described install standing in for a performed one.

**What a reader of the baseline file will see first.** A number raised from 12 to 28 with two
criticals appearing. The file's own rule puts the explanation in the commit message, so the message
leads with the correction and names all four severity classes, including `low` going from 0 to 1,
which P3 records as a deviation because the design wrote only three.

**Three places carried the same false sentence and all three are replaced as paragraphs.** The audit
header, the baseline comment and the supply-chain prose had each been amended once before, which is
the case the convention is about — a small diff is exactly when the words left around it are not
noticed. Each was read back as prose afterwards, and `docs/supply-chain.md` was checked for the
88-column rule, where the only over-long lines are table rows and one pre-existing line about
semgrep.

**What this intent deliberately did not do, and why that is not an omission.** It fixes no advisory,
bumps no sha and commits no lockfile. The first two are outside this repository's reach, measured
rather than assumed; the third is answered by the measurement rather than deferred, since the shipped
tree is already pinned at the action's sha. The sha bump is the real follow-on and it needs the
number this intent produces before it can be judged.

<!-- xeno:section:release-notes -->
## Release notes

**The npm audit now reads the tree the release installs.** It did not. The release installs
semantic-release through `cycjimmy/semantic-release-action`, which runs `npm ci` against its own
committed lockfile and installs the pinned version on top; the audit started from `npm init -y` and
resolved every transitive fresh. Measured on 2026-10-06, those are different trees — 12 advisories
and no critical against 28 and two — and the audit's was the smaller one. It now checks the action
out at the sha `release.yml` pins and performs the action's own steps, taking both the sha and the
version out of that file rather than keeping copies.

**The baseline rises from 0/10/2/0 to 2/23/2/1, and the rise is a correction.** Nothing in the
pipeline changed between the two measurements. The two criticals are `handlebars`, JavaScript
injection via AST type confusion, and `tar`, arbitrary file creation via hardlink path traversal,
both held at their vulnerable versions by the action's lockfile. Sixteen advisories in total were
invisible to the old tree, among them `@semantic-release/changelog` and `@semantic-release/git` —
the two plugins #121 removed from the workflow's plugin list, which the action installs anyway as
its own dependencies.

**The counts stop drifting.** 531 of the tree's 532 packages are verbatim from the action's
lockfile, the single difference being `semantic-release` itself at the pinned version, so the
figures are now a function of two identifiers in this repository. A pull request touching three
Markdown files can no longer fail this job because npm published overnight, which is what #260 was
filed about.

**#260's lockfile question is answered rather than left open.** The shipped tree is already pinned,
by the action's commit sha, under A28. A `package-lock.json` committed here would pin a third tree
that nothing installs.

**`docs/supply-chain.md` and A44 are corrected.** The document repeated the false claim and said the
audit installs three packages where it has installed one since #121; its table said
semantic-release 24.2.9 against a pin of 25.0.9. A44's row records that the evidence moved under it:
approved on ten high advisories, it now has to hold for two criticals, and its reasoning does —
none of the twenty-eight is this project's to fix, and the sixteen new ones move only when the
action's sha does.

**What is deliberately not here.** No advisory is fixed and none can be from this repository. The
action's sha is not bumped: a newer lockfile may well drop both criticals, and that is a change to
what cuts every release, which A28 judges by what demonstrably works. It wants its own intent, and
this one produces the number to judge it against.

<!-- xeno:section:residual-risk -->
## Residual risk

**The CI run has not happened.** Criterion 10 is open, which P4's gaps section states rather than
rounds off. Every figure was measured under node 26.9.0 and npm 12.0.2 where CI pins node 24, and
531 of 532 packages come from a lockfile no node version re-resolves — so the exposure is one
install, which added no path the lockfile did not already hold. Small and not nil, and if the run
disagrees the figures are wrong and not the design.

**`actions/checkout` has not been exercised with `repository`, `ref` and `path` together.** The
reproduction fetched the same sha with `git fetch --depth 1`, which yields the same tree by a
different command. A public repository, a 40-character sha and a relative path are all documented
inputs, and the failure would be loud. That is the whole mitigation.

**Two criticals are now on every run, with nothing to do about them.** The intended effect, and the
hazard A44 names from the other side: a number nobody can act on is a number people stop reading.
It is bounded rather than removed — the counts can now only move on a commit here, so the next move
is attributable. What is not bounded is habituation, and the deferred sha bump is the thing that
would actually retire it.

**The baseline's prose and its counts can disagree and nothing notices.** The comment names all
twenty-eight advisories so a reader can diff by hand what four counts cannot distinguish; a critical
replacing a critical still reads as no change. Whoever next raises a number has to update both, and
only a person will catch it if they do not. This is a new instance of the shape
`docs/clause-readers.md` catalogues, and it does not get a row there, because that document is a
pass over the two normative documents and this is a comment in a JSON file.

**A reader of XENO-0230 meets the false claim with nothing beside it.** Three of its phases assert
that the audit installs exactly as the release does, and section 11 means they stay. The correction
is here, in #260 and in `audit.yml`'s header, none of which is where somebody reading that intent is
looking. The same residual risk XENO-0254 recorded about the gitignore claim, by the same mechanism.

**`release-toolchain/` sits in the workspace while the job runs.** Nothing reads it but the two
steps that build and audit it, and nothing uploads it. No check asserts that, and a later step added
to this job would meet a tree of 533 packages beside the repository.

**Staleness is still unwatched.** Correct and stable counts are not evidence that the pins are the
ones to be on. Unchanged by this intent, out of its scope, and worth saying beside a baseline that
now looks authoritative. #44 holds the question; the sha bump is where it is next asked.
