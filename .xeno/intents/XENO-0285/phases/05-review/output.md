---
intent: github.com/triplem/xeno#95
phase: 05-review
created: "2026-10-09T16:55:53Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f1a1f1398f98da435598ac7cd90046900a81fdee066520e302ad20de942ee297
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'The implementation is the design, and the three figures the design named were checked against the file rather than asserted: customManagers six to nine, description plus ten lines, commitBody present in HEAD and absent here. No deviation. The one thing weighed and not taken, a committed script that re-runs the manager check, is recorded as a gap in P4 rather than slipped in, because it is a tool the design did not propose.'
      result: met
      rule: deviations-are-traceable
    - note: No interface change. Three regex managers appended to a configuration, one key removed from it, and a tag added to a Dockerfile ARG whose digest is unchanged, so FROM resolves the manifest it resolved before. No command, flag, artifact field or template moves, and nothing has ever run the previous configuration to migrate from.
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: No new dependency. go.mod is untouched, no action is added and no tool is installed. Three pins this repository already carries become visible to a bot that was already configured, and the image's base is the same digest it was, named by the tag docs/supply-chain.md already called it.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
    - note: 'Supply chain lens: this change widens what a third party will propose moving. A28 pins to what a release demonstrably used, and a bot proposes the newest by construction, so three more pins now have a proposer. XENO-0277 accepted that tension for the pins it reached and it is the same tension here, one notch wider -- most sharply for the image''s base, where the tag now names a stream renovate will offer newer digests of. Accepted on the same ground: nothing automerges, every proposal is read, and a pin nobody can be told has moved is the worse failure. Named rather than resolved.'
      result: deviation
      source: lens
---

# Review

<!-- xeno:section:release-notes -->
## Release notes

**Renovate can now see every pin this repository records as movable.** Three were invisible
to it: the image's base in the `Dockerfile`, `gitleaks` in `gitleaks.yml`, and `zensical` in
`docs.yml`. Each is a row of `docs/supply-chain.md` — a version this project pinned
deliberately — and each sat in a line no manager matched, so nothing could ever report that
it had moved. A regex manager now reaches each of the three, bringing the configuration from
six to nine.

That was the argument for renovate over dependabot in the first place: most of what this
project pins is a version written inside workflow text rather than in a manifest, and the
plain managers see none of it. Three pins were still outside, which left the argument
unfinished.

**The base image reference now names its tag.** `ARG BASE=docker.io/library/debian@sha256:…`
became `…/debian:13-slim@sha256:…`. The digest is unchanged, so the image builds from the
same bytes; the tag is there because `docs/supply-chain.md` already called this pin `debian`
13-slim and the file did not say so, which left the page unconfirmable against the tree and
left a datasource with no stream to follow.

**Nothing is bumped.** Every version in the tree after this change is the version that was
there before. This makes pins visible, not newer.

**One proposal will arrive broken, deliberately.** Renovate can move `GITLEAKS_VERSION` and
cannot recompute the `GITLEAKS_SHA256` beside it, so a gitleaks bump fails its install step's
checksum until somebody recomputes it. That is the trade this configuration already makes for
actions, whose bumped sha needs a `docs/supply-chain.md` row moved by hand. Loud and waiting
for a person beats a pin going stale with nothing to notice. Nothing automerges, so no broken
proposal can land on its own.

**`commitBody` is gone.** It carried `Refs #95`, which this change closes. `CONTRIBUTING.md`
puts the issue of the work in the footer and a dependency bump is the work of no issue, and
the squash takes the pull request description, so a branch commit's body never reached main
anyway. The configuration's own `description` says this, because a key that is absent cannot
explain itself to the next reader.

**Still switched off.** None of this runs until `RENOVATE_TOKEN` is in the repository's
secrets, which only a maintainer can add. That is #350, with what to check on the first run.

<!-- xeno:section:residual-risk -->
## Residual risk

**The datasources are unexercised.** That `pypi` answers for `zensical`, that
`github-releases` with `extractVersionTemplate` turns tag `v8.30.1` into `8.30.1`, and that
the `docker` datasource follows `13-slim` — all three are read off renovate's documented
behaviour, not observed. The regexes were run as renovate's own dialect by a script that is
not renovate. A behaviour the two disagree about would not show up here, and the first real
run is where it would. #350 carries that check.

**A wrong `extractVersionTemplate` would be quiet.** If it failed to strip the `v`, gitleaks
would be compared against a version that does not exist and the manager would simply never
propose — which looks exactly like a pin that has not moved. This is the failure mode the
intent is about, reappearing one level down. The script shows the current value extracting
correctly; it cannot show the comparison succeeding, because that needs the datasource.

**Nothing re-runs the manager check.** Recorded as a gap in P4. The three pins are reached
today; a workflow line rewritten later can silently unreach one, and the next person has to
think to look.

**The gitleaks proposal's failure has to stay legible.** It fails on a checksum, which reads
like a tampered download rather than like a bump that needs a hand-recomputed hash. The
reason is in the manager's `description` and in `renovate.json`'s top-level `description`, so
it is in the file rather than only in this artifact — but somebody meeting a red
`sha256sum -c` for the first time will reasonably wonder about the download before reading
the configuration.

**What is not a risk.** The image. The digest is byte for byte unchanged, and a digest wins
over a tag, so `FROM $BASE` resolves what it resolved before. If the tag and the digest ever
disagreed, the digest would decide and the build would not move.
