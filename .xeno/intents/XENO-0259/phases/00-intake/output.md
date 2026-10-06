---
intent: github.com/triplem/xeno#260
phase: 00-intake
created: "2026-10-06T09:11:46Z"
schema_version: "1.0"
runner_version: dev+081da51
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e256d8656e0a16728fb828f6dd54e15be693bbbf84fddb5ce509cd8c84cee391
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

The `audit` workflow exists to answer what the release pipeline runs, as opposed to what the
release ships, and `docs/supply-chain.md` splits the two that way: trivy reads the binary, the
audit reads the tree that produces it. The audit's own comment states the premise it rests on —
semantic-release is "installed here exactly as the release installs it, and audited".

**It is not.** Measured on 2026-10-06 under node 26.9.0 and npm 12.0.2, the two installs produce
different trees and the audit's is the smaller one:

| tree | critical | high | moderate | low | total |
|---|---|---|---|---|---|
| `npm init -y` then `npm install semantic-release@25.0.9`, which is `audit.yml` | 0 | 10 | 2 | 0 | 12 |
| what `cycjimmy/semantic-release-action@b12c8f6` installs | 2 | 23 | 2 | 1 | 28 |

The release tree is a strict superset: every advisory the audit reports is in it, and sixteen more
are not in the audit. Two of the sixteen are critical — `handlebars`, JavaScript injection via AST
type confusion, and `tar`, arbitrary file creation via hardlink path traversal. Four more are the
semantic-release plugins themselves, among them `@semantic-release/changelog` and
`@semantic-release/git`, the two that #121 removed from the workflow's plugin list. `release.yml`
already says of a third plugin that "absent from the list is not absent from the machine"; it is
true of these two as well, and the audit has never looked where they are.

## Why the trees differ

The action does not install semantic-release the way the audit does. `index.js` runs

    npm --loglevel error ci --only=prod

in the action's own checkout, against the `package-lock.json` the action commits — 534 entries,
pinning `npm@11.6.2`, `handlebars@4.7.8`, `lodash@4.17.21` and `js-yaml@4.1.1` among them. Only
then does `src/installSpecifyingVersionSemantic.task.js` run `npm install
semantic-release@25.0.9` on top of that tree.

The audit starts from `npm init -y`, so npm resolves every transitive dependency fresh and takes
the newest each parent's range allows. That is why `handlebars` is absent from it: a newer
`conventional-changelog-writer` no longer reaches it, and the action's lockfile holds the older one
that does. **The pin is what keeps the vulnerable versions, and the audit's unpinned resolution is
what hides them.**

## The drift was in the wrong tree

#260 was filed because the counts moved with nothing in this repository to cause it, and reasoned
that the tree resolves fresh on every run. That is true of the tree the audit builds. It is not
true of the tree the release installs: comparing the 532 installed packages path by path against
the action's lockfile, **531 are verbatim from it, and the single difference is `semantic-release`
itself, 25.0.2 from the lockfile raised to the pinned 25.0.9.** The one floating install adds no
package the lockfile did not already resolve.

So the shipped tree is already pinned, by the action's commit sha, which is what A28 asks for and
`docs/supply-chain.md` records. The thing that drifts is a tree that exists only inside the audit
job and that nothing ships. #260's second route — commit a lockfile here so the counts stop moving
— would therefore produce a third tree, pinned to this repository and to neither of the other two,
and the audit would still not be measuring what runs.

## What is wrong, stated once

The audit reports zero critical advisories for a pipeline that runs two. A check whose number is
false in the optimistic direction is worse than the one A90 objects to: a reader that cannot fail
tells you nothing, and this one tells you something untrue.

<!-- xeno:section:scope -->
## Scope

In scope is `audit.yml` installing the tree the action installs: a second checkout of
`cycjimmy/semantic-release-action` at the sha read out of `release.yml`, `npm ci --omit=dev` in it,
then `npm install semantic-release@<pin>` on top, which is the action's own three steps in the
action's own order. The sha is read and not copied, for the reason #187 gives about the version: two
copies of one identifier is how an audit comes to examine a tree nothing ships, which is the fault
being repaired here rather than a risk being guarded against.

In scope is the baseline rising from 12 to 28 with two criticals in it, and the comment in
`.github/npm-audit-baseline.json` saying which tree the figures are of and how it is built. The
file's own rule is that raising a number is a deliberate act belonging in the commit message that
raises it; this raise is a correction of what was measured rather than a worsening of what runs, and
the commit message has to say so in one sentence or a reader takes it for the opposite.

In scope is A44's wording. The assumption says the counts are not this project's to fix, which was
written about ten high advisories and now has to be said about two criticals as well. Whether that
is still the right answer is the question the design has to put, not assume.

In scope is correcting `docs/supply-chain.md`. It carries the same false claim — the audit "installs
the three pinned semantic-release packages exactly as the release does" — and two further errors
beside it: the packages are one since #121, and the table's semantic-release row says 24.2.9 where
the pin is 25.0.9.

Out of scope is fixing any advisory. None of the twenty-eight is reachable from here: the sixteen
new ones are inside the action's committed lockfile, which only a bump of the action's sha can move,
and the twelve old ones were measured on 2026-10-06 as unfixable by `npm audit fix` with
`semantic-release@15.14.0` the only alternative npm offers. This intent changes what is measured and
not what is installed.

Out of scope is bumping `cycjimmy/semantic-release-action`. A newer sha would carry a newer lockfile
and might drop the criticals, and that is worth knowing; it is also a change to what cuts every
release, judged against A28's rule that what is wanted is the set that demonstrably works. It wants
its own intent with the release that demonstrates it, and this one has to land first or there is no
number to judge the bump against.

Out of scope is committing a lockfile in this repository. The problem section measures why: the
shipped tree is already pinned by the action's sha, and a lockfile here would pin a third tree that
nothing installs. #260's own framing of that route is answered rather than deferred.

Out of scope is staleness. Nothing here says whether the pinned action or the pinned version is
still the one to be on, and #44 keeps that apart on purpose. A correct count that does not move is
no more evidence of currency than a wrong one was.

No normative document is touched. The first standing rule binds `docs/process-definition.md` and
`docs/implementation-plan.md`; neither says anything about which tree an audit job installs.
`docs/supply-chain.md` and the baseline file describe the repository rather than specify the
process, and A44 is a row in the assumption register, which is where a decision of this kind is
recorded.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the two workflows that disagree, the file that records the number, the document that
repeats the claim, and the two assumptions the change sits between.

`.github/workflows/audit.yml` is read for the premise under examination. Its header is the clearest
statement of what the job is for — "What the pipeline runs, as opposed to what the release ships" —
and of the claim that fails: "It is installed here exactly as the release installs it, and audited."
The header also carries #187's rule that the version is read out of the release's workflow rather
than copied, which is the pattern the sha has to follow.

`.github/workflows/release.yml` is read because it is where the install actually happens, and the
install is one line: `cycjimmy/semantic-release-action` at a sha, with `semantic_version: 25.0.9`.
What that action does with those inputs is not in this repository, so it was read from the action at
the pinned sha rather than assumed — `index.js`, `src/index.js` and
`src/installSpecifyingVersionSemantic.task.js`, which are the three steps the scope now describes.

`.github/npm-audit-baseline.json` is read for its own rule about itself and for what it claims. Its
comment is where the figures and their date live, and it is the file whose numbers this intent
raises; it also carries #260's open half in its last sentence, which this intent closes.

`docs/supply-chain.md` is read as the document of record for every pin, and it is where the false
claim is repeated in prose rather than in a comment. It is also the file that draws the trivy and
audit boundary this intent depends on: if the Node tree were shipped rather than run, two criticals
would be a release problem and not an audit one.

`docs/assumptions.md` is read for A28 and A44, which this change sits between and which turn out to
interact. A28 pins the action to a sha so that two runs a month apart use the same tooling; A44
baselines the audit's counts because none of them is this project's to fix. The interaction is the
finding: A28's pin is what holds the vulnerable versions in place, and until now the audit measured
a tree where A28 did not apply, so neither assumption's reader could see the other.

`docs/process-definition.md` is read to confirm that nothing in it governs this. Section 7's gate
list and the artifact schema say nothing about an audit job, so no specification commit precedes
this one and the first standing rule is not engaged.

`CLAUDE.md` is read for the two conventions this intent is run under. A decision put to a person is
put one at a time, which is why the choice between re-measuring the wrong tree, measuring the right
one and narrowing the check was put as one question with its consequences. And a negative result is
evidence only when the thing checked was there to be found, which is why "`handlebars` is not in the
audited tree" was not written down from a grep over a tree that had not been built: both trees were
installed and audited, and the 532 paths compared against the lockfile, before any of it was
claimed.
