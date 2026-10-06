---
intent: github.com/triplem/xeno#260
phase: 05-review
created: "2026-10-06T07:30:48Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0338e0aaa27bc11aace13ef5ba6cf8cb972005d659a44a2b1ab3d01149b44fdf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'One addition beyond P2 design, and it is the only verification in the intent. P2 said criterion 7 could be met only on the pull request, because nothing here runs audit.yml. True of the job, and the comparison it performs turned out to be reproducible: the failing run uploads its npm-audit.json as an artifact, so the baseline own rule was run over it with the new numbers and nothing rises. The criterion is still met on the pull request and the phase no longer rests on that alone. Nothing else departed: the replacement rather than amendment of the prose, both directions in one sentence, the toolchain line, the absence of a drift field and of a register row, and Refs rather than Closes, all landed as specified.'
      result: deviation
      rule: deviations-are-traceable
    - note: No interface changes. Four numbers, a date and a paragraph in a JSON file that no Go code reads, that lies outside artifacts_hash, and that only .github/workflows/audit.yml consumes — and the job reads the same five keys it read before. Nothing outside this repository depends on it. What changes for a person is that the audit stops failing on pull requests that did not cause it, and that the file now answers whether the packages can be bumped or fixed, which it previously could not.
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: 'None added and no code changed; go.mod is untouched. The change runs in the opposite direction from a dependency decision: it records that the one pinned npm toolchain cannot be moved, because 25.0.9 is already the latest and npm only suggestion for seven of the twelve advisories is a major downgrade that A44 declined and A28 pinning exists to prevent. The dependency question that remains is whether to commit a lockfile so the transitive tree stops moving, which is the open half of #260 and a decision under A28 rather than part of this.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: the plan's WP0 asks for the audit and
fixes none of its numbers, so re-measuring needs no specification commit. Nothing is invented — the
file keeps its five keys, no field is added, and the one temptation was a machine-readable note
about the drift, which went into the prose because the reader is a person. The branch carries one
intent, the commit references the issue, and #260 carries `wp0`.

The acceptance criteria. Ten met, one met on the pull request, which is where the audit job runs.

The non-goals held. No lockfile, no narrowing of the check, no downgrade, no change to `audit.yml`,
no resolution of any advisory, no refresh cadence adopted, and no register row — A44 records the
arrangement and the arrangement is unchanged.

**What a reviewer should check is the direction.** One number rises and the commit message has to
be read for both: high falls 30 to 10, moderate rises 1 to 2. The file's own rule is that raising a
number is a deliberate act belonging in the message that raises it, so the message is part of the
work and a reader who finds only the improvement quoted has been handled rather than informed.

**What was verified and what was not.** The job's comparison was reproduced over the
`npm-audit.json` the failing run uploaded, with the new baseline: nothing rises. The job itself —
its install, its pin resolution, its report generation — runs in CI and nothing here can exercise
it, which P4's gaps states rather than letting the reproduction stand for more than it is.

The question that prompted this was answered with a measurement rather than an opinion. Bumping and
fixing are both unavailable: 25.0.9 is already the latest, `@latest` resolves to an identical tree
with identical counts, and `npm audit fix` moves nothing. That is now in the file so the next reader
does not repeat the experiment, and it makes A44's "none of them is this project's to fix" measured
rather than reasoned.

What is left open is named in P4's gaps and on the issue: the drift itself, which only a lockfile
would end; the job's unverifiability from here; and whether a file needing two deliberate
re-measurements in three days should have a refresh rule or be a different instrument.

<!-- xeno:section:release-notes -->
## Release notes

The npm audit baseline is re-measured, and the audit job stops failing on pull requests that did
not cause it.

| | was, 2026-10-03 | now, 2026-10-06 |
|---|---|---|
| critical | 0 | 0 |
| high | **30** | **10** |
| moderate | **1** | **2** |
| low | 0 | 0 |

**Most of this is a tightening.** The high bound falls by twenty as upstream published fixes, so
twenty high advisories that could previously have arrived unreported now cannot. One number rises:
moderate, by one, which is `postcss-selector-parser` returning — the old prose said it had gone with
the 24.2.9 to 25.0.9 upgrade, and a freshly resolved tree reaches it again.

Nothing in this repository changed between the two measurements. The job blocked #259, whose diff
is three documentation and example files.

**Why it drifts, now recorded in the file.** `audit.yml` runs `npm init -y` and
`npm install semantic-release@<pinned>` with no lockfile anywhere here, so transitive dependencies
resolve fresh on every run. The counts are a property of npm's registry at the moment of the run
rather than of this repository, and the file now says to treat them as a dated measurement and not
as a bound.

**Bumping and fixing are both unavailable, and the file says so.** `semantic-release@25.0.9` is
already the latest published version, so an install of `@latest` produces an identical tree and
identical counts; `npm audit fix` resolves none of the twelve, leaving high 10 and moderate 2 with
`semantic-release` still at 25.0.9. The five advisories npm marks `fixAvailable` are not fixable in
practice, because without a lockfile every install already takes the newest transitives the parents'
ranges allow. For the other seven npm's only suggestion is `semantic-release@15.14.0`, a major
downgrade, which A44 declined and A28's pinning exists to prevent.

So A44's "none of them is this project's to fix" is measured rather than reasoned, and true of all
twelve today. The next person to meet a failing audit reads the file instead of repeating the
experiment.

**What this does not do.** It does not stop the drift: the mechanism is untouched and the next
upstream publication can fail an unrelated pull request the same way. Whether the tree should be
pinned so the numbers stop moving without a commit is the open half of **#260**, which is why this
is `Refs` and not `Closes`.

<!-- xeno:section:residual-risk -->
## Residual risk

One number rises, and a reader who takes that alone has the change backwards. The mitigation is
entirely in the commit message and the file's prose, both of which say both directions; neither is
enforced, and the file's own rule — that raising a number is a deliberate act belonging in the
message that raises it — has no reader but a person. If somebody later raises a count without
saying so, nothing catches it.

The drift is untouched and will recur. This is a dated measurement, not a bound, and the next
upstream publication can fail an unrelated pull request exactly as this one did. The only route that
ends it is a lockfile, which is the open half of #260 and a decision under A28; until then the
defence is that the prose now explains the failure to whoever meets it.

**This is the second deliberate re-measurement in three days.** #189 raised the baseline and this
re-measures it. A file needing that cadence is either a file with a refresh rule or the wrong
instrument, and nobody has decided which. Two data points is not a trend, and the risk is that the
third arrives unremarked because the second was handled smoothly.

The job is unverified from here. What was reproduced is its comparison, over an artifact it
produced; the install, the pin resolution and the report generation run in CI. A change to any of
them would fail in a way nothing local would see until a pull request went red — the same limit
#201 recorded about the release workflow, and it has not improved.

The ten high advisories are counted and not examined. None has a fix this repository can apply,
which is measured; whether any is reachable in the way the release actually uses `semantic-release`
is a question nothing here asks, and A44's position that the count is the signal is unexamined
rather than wrong.

Nothing records the path between the two measurements. High fell from 30 to 10 over three days and
the file says it fell, not when or in how many steps, so a later reader cannot distinguish a steady
improvement from a spike and a recovery. The Actions history and the uploaded artifacts hold it for
as long as they are retained.

What is not a risk: the verdicts that existed before this intent, confirmed intact; anything
executable, since no Go code reads this file and it lies outside `artifacts_hash`; and the audit's
own shape, which is unchanged — what moved is the numbers it compares against.
