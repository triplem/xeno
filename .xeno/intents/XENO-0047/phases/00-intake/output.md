---
intent: github.com/triplem/xeno#47
phase: 00-intake
created: "2026-09-26T15:15:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+a336153
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 03a00eb8903bb40f9bcd700024f05c8fa13148d45b5815364a6bd64e7c6cf886
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

This repository was private on a tier where branch protection, rulesets and required
reviewers do not exist: the API answered `403 Upgrade to GitHub Pro or make this
repository public` to every question about them. First step 2 of section 4 calls that
setting the one the whole tool claims to rest on, so the sequence guarantee rested on
discipline instead, `xeno enforcement check` reported it as not available, and A27
recorded the gap as a property of the plan rather than of the host.

The constraint was the tier and the tier followed from the visibility. What had kept the
question open was policy rather than licence: whether the copyright holder's rules
allowed publication.

<!-- xeno:section:scope -->
## Scope

The visibility, and the two sentences in the documents that tied it to distribution. The
repository is public from v1; the release still carries checksums rather than
signatures, the channels are still the host's releases and registry, and the plugin is
in no marketplace, all of which wait for 1.1 as before.

The setting itself is the maintainer's to flip, after this commit rather than before it.
What follows the flip is one measurement, `xeno enforcement check` against the host, and
then A27 and the security contact half of A17 close against what it reports rather than
against an expectation.

<!-- xeno:section:context-rationale -->
## Why this context

**The policy coupling went away rather than being resolved.** #47 concluded that the
blocker was the copyright holder's policy, not the content of the repository. #81
changed the holder to javafreedom.org, which removed the question the issue could not
answer from inside.

**The obligations were measured before the visibility changed, because two of them
cannot be met afterwards.** Section 8 requires a history publishable from the first
commit and a trail written in the knowledge that strangers read it. Both were checked
over the whole history and not only the current tree: no mail address outside `noreply`
identities, no string of the shape of a key ever committed, no internal hostname, no
private address, and in 252 trail files no personal name and no decision attributed to
anyone but "the maintainer". The third obligation, current third party attribution, is
one vendored dependency and a `NOTICE` that names it.

**Visibility and distribution were one sentence and are two decisions.** Publishing the
repository gives the host a setting to enforce; publishing a release means signatures, a
channel and a marketplace entry, and none of those moved. Writing them as one sentence
is what made this look like a larger decision than it was.

**What the gate path shows now is what the host reports.** Until the flip, not available
was the honest state and was built as a third value for exactly this case. After it, the
same command answers from the same API and the report says met or unmet. The difference
is measured rather than assumed, which is the acceptance criterion the issue set itself.