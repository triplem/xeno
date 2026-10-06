---
intent: github.com/triplem/xeno#260
phase: 00-intake
created: "2026-10-06T07:21:45Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 61f47006d73a7c6c68bf9d169eb4e65cc9239201ffe31f19501df4912da29ff4
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

The audit job fails on every pull request and nothing in this repository changed to cause it. It
blocked #259, whose diff is three documentation and example files.

```
above the recorded baseline: moderate: 2 against 1 recorded
counts: {"critical": 0, "high": 10, "moderate": 2, "low": 0}
```

| | baseline, measured 2026-10-03 | run of 2026-10-06 |
|---|---|---|
| critical | 0 | 0 |
| high | **30** | **10** |
| moderate | **1** | **2** |
| low | 0 | 0 |

High fell by twenty and moderate rose by one. The check fails on a count rising, so a
twenty-advisory improvement and a one-advisory regression together read as a failure, and the
high bound is now so loose that twenty new high advisories could arrive unreported.

The new moderate is `postcss-selector-parser`, quadratic complexity in flat selector parsing.
The baseline's own prose says that one went away with the 25.0.9 upgrade — "picomatch,
postcss-selector-parser and three of the five brace-expansion advisories are gone with it" — so
it has returned through a freshly resolved tree.

**It drifted with no commit because the tree is resolved at install time.** `audit.yml` does
`npm init -y` and `npm install semantic-release@<pinned>`, with no lockfile anywhere in the
repository, so transitive dependencies resolve fresh on every run and the counts move when
upstream publishes. The baseline measures a moving target on npm's schedule rather than this
repository's.

**Bumping and fixing were tested and neither is available.** `npm view semantic-release version`
is 25.0.9, which is exactly what is pinned, so there is nothing to bump to; an install of
`@latest` produces an identical tree and identical counts. `npm audit fix` changes nothing —
high 10 and moderate 2 before and after, with `semantic-release` still 25.0.9. The five
advisories marked `fixAvailable: true` are not fixable in practice, because with no lockfile every
install already takes the newest transitives the parents' ranges allow, so the constraint is the
parent's range. For the other seven npm's only suggestion is `semantic-release@15.14.0`, a major
downgrade, which A44 declined and A28's pinning exists to prevent.

So A44's "none of them is this project's to fix" is now measured rather than reasoned, and true of
all twelve today.

<!-- xeno:section:scope -->
## Scope

In scope is re-measuring `.github/npm-audit-baseline.json`: `high` 30 to 10, `moderate` 1 to 2,
`measured` to 2026-10-06, and the `advisories` prose rewritten to describe the twelve that are
there now rather than the thirty-one that were.

In scope is recording in that prose that bumping and fixing were tested and neither is available,
with the figures. The file's existing prose explains a version upgrade; what it has never carried
is the measurement that there is nothing left to upgrade to.

In scope is the commit message carrying the deliberate act, because the file asks for exactly
that: "Raising a number here is a deliberate act and belongs in the commit message that raises
it." One number rises and one falls by twenty, and the message has to say both so that a reader
does not take a loosening for the whole of it.

Out of scope is the lockfile. Whether the counts should become a function of this repository
rather than of npm's publishing is the open half of #260, it is separable, and the counts it would
record at pin time are the same 10 and 2 — so this intent is the common prefix of both routes and
forecloses neither.

Out of scope is narrowing what the check asserts. Failing only on critical and high, or only on an
advisory with an applicable fix, is the third route on #260 and gives up signal A44 wanted; this
intent does not take it and does not prevent it.

Out of scope is downgrading `semantic-release`. A44 declined it, A28 pins the toolchain
deliberately, and npm's suggestion of 15.14.0 is a major downgrade of the release machinery.

Out of scope is anything about the ten high advisories beyond counting them. None has a fix this
repository can apply, which is now measured, and the baseline's purpose is to notice a change
rather than to resolve one.

Therefore `Refs #260` and not `Closes`. The issue's own done-when asks that the audit pass on a
pull request that changed nothing relevant to it *and* that the route taken be recorded; this does
the first and records this route, while the lockfile question stays open on the issue.

No normative document is touched. The baseline is a measurement and says so, and the plan's WP0
asks for the audit without fixing its numbers.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the file being changed, the job that reads it, the pin it is measured against, and
the two register rows that govern both.

`.github/npm-audit-baseline.json` is read for its own rule about itself, which is the reason this
is an intent rather than an edit: "Raising a number here is a deliberate act and belongs in the
commit message that raises it." Reading it also supplies the prose that has to be replaced rather
than amended — it names advisories that are gone and says `postcss-selector-parser` went with the
upgrade, which today is false.

`.github/workflows/audit.yml` is read for what the check actually compares and how the tree is
built. The second is the finding: `npm init -y` and `npm install semantic-release@<pinned>` with no
lockfile, so the counts are a property of npm's registry at the moment of the run. Nothing in the
baseline's prose says that, and it is why the file drifts without a commit.

`.github/workflows/release.yml` is read for the pin the audit job derives its version from, so
that the new prose names the toolchain it was measured against rather than the toolchain somebody
assumes.

`docs/assumptions.md` is read for A44 and A28. A44 is the row this file embodies — the audit fails
on a count rising, not on a count being non-zero, because none of these is this project's to fix
and a check firing with no action available is one people route around. A28 is why a downgrade is
not on the table. Both are load bearing for the new prose and both are quoted from rather than
recalled.

`docs/implementation-plan.md` is read for WP0's account of the audit, to confirm that the plan asks
for the check and fixes none of its numbers, so re-measuring needs no specification commit.

The counts themselves come from two places and both are recorded rather than trusted: the failing
CI run's own output, and a local reproduction of the same install on node 26.9.0 and npm 12.0.2,
which produced identical figures for the pinned version and for `@latest` and showed `npm audit
fix` resolving none of the twelve. The local run is what turned "we could probably just bump it"
from a plausible answer into a measured no.
