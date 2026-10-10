---
intent: github.com/triplem/xeno#355
phase: 01-requirements
created: "2026-10-10T12:13:16Z"
schema_version: "1.0"
runner_version: dev+90c7227.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cd40d0813e2c9424330be0f81c028ff8b8dc3061ce8c262c3d7192fe1661d336
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Thirteen criteria, and the two that matter most are the two the issues did not contain.

Criterion 2 is the claim the whole design rests on: that `docker buildx imagetools create`
copies a manifest index with its digest intact. If that is false the mirror is a second
pin rather than the same one, `ARG BASE` and the released image disagree, and the decision
on #355 does not hold. It is written as a criterion to be measured against a registry
rather than as a premise, because the local docker has no buildx plugin and documentation
is not evidence. P4 is where it is answered; a wrong answer is a design change and not an
implementation fix.

Criterion 10 came out of reading the release job. `IMAGE` is written into `GITHUB_ENV` by
the build step, and a dispatch skips the build step, so the image step as it stands would
resolve its registry to the empty string. Neither issue mentions it: both were written
about why a step failed, not about what a second trigger into that step would find
missing. The learning generalises it — a second trigger is a requirement about what a job
carries between its steps and not only about the conditions on them.

The rest divide cleanly. Criteria 1, 3, 4 and 5 are the mirror: fetched from one registry,
the digest written once and read at run time, `ARG BASE` untouched so that #351's manager
and a hand build keep working, and Docker Hub reached on a miss and nothing else.
Criteria 6 to 9 are the way back: a dispatch reaches the image step, carries the bytes
that release already published and verifies them against its `SHA256SUMS`, fails for a
version with no release, and cuts nothing. Criteria 11 to 13 are the page, the test that
holds the page to the tree, and the gate suite.

Criterion 4 is the narrowest and the one most easily lost: `git diff main -- Dockerfile`
empty. Everything the decision on #355 bought — the regex manager, the hand build, one pin
rather than two — is paid for by the release reading that file rather than restating it.

Three sections, no open question, no decision; all four were taken before P0.
