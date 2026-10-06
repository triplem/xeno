---
intent: github.com/triplem/xeno#260
phase: 03-implementation
created: "2026-10-06T07:24:54Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e94aa77d744264e8b8fa54915700bf6cc990504c8b658aee67c2ecdb3fb80f98
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One file, five values and one paragraph. `.github/npm-audit-baseline.json` is the whole diff.

`vulnerabilities` becomes `critical: 0, high: 10, moderate: 2, low: 0`. `measured` becomes
2026-10-06. `toolchain` keeps `semantic-release@25.0.9` and gains "and also the latest published
version", which is the line that answers the next reader's first question from the file instead of
from an experiment.

`advisories` is replaced rather than amended. The old paragraph narrated the 24.2.9 to 25.0.9
upgrade and said `postcss-selector-parser` had gone with it, which is false today; editing into a
paragraph whose central claim has failed is how the words around it get left behind.

The new paragraph answers three questions in the order a reader arriving from a failing job needs
them. What is there now: twelve advisories, named, and both directions in one sentence — high from
30 to 10, moderate from 1 to 2, with `postcss-selector-parser` returning through a freshly resolved
tree. Why it drifted: the tree is resolved at install time, `npm init -y` and `npm install
semantic-release@<pinned>` with no lockfile anywhere here, so the numbers are a property of npm's
registry at the moment of the run and are a dated measurement rather than a bound. What was tried:
25.0.9 is the latest so `@latest` gives an identical tree, `npm audit fix` resolves none of the
twelve, the five marked `fixAvailable` are not fixable because the constraint is the parent's range,
and npm's suggestion for the other seven is a major downgrade that A44 declined and A28 prevents.
It closes by naming the lockfile question as the open half of #260.

**The check was reproduced locally against the real report rather than assumed.** The job's
comparison — each count against the recorded one, failing on a rise — was run over the artifact CI
uploaded from the failing run, with the new baseline: nothing rises, so the job passes. That is as
close as this repository can get to verifying `audit.yml` without running it.

The JSON parses and keeps the same five keys it had.

<!-- xeno:section:deviations -->
## Deviations from the design

No deviation from the design. The replacement rather than amendment, the three questions in that
order, both directions in one sentence, the `toolchain` line, the absence of a `drift` field, the
absence of a register row, and `Refs` rather than `Closes` all landed as P2 specified.

One addition beyond the design, and it is the only thing here that counts as verification. P2 said
criterion 7 could only be met on the pull request, because nothing in this repository runs
`audit.yml`. That is still true of the job, and the *comparison* the job performs turns out to be
reproducible: the failing run uploads its `npm-audit.json` as an artifact, and running the
baseline's own rule over it with the new numbers shows nothing rising. So the criterion is still
met on the pull request, and the phase is no longer relying on that alone.

One thing worth recording about the figures. Both the CI run and a local install of the pinned
version produced `high 10, moderate 2`, and so did an install of `@latest` — three measurements
agreeing is why the numbers went into the file without hedging. The local pair is also what
established that there is nothing to upgrade to, which is the half of the question the old prose
could not answer and the new one can.

Nothing else departs. No change to `audit.yml`, no lockfile, no narrowing of the check, no
downgrade, no register row, and no normative document.
