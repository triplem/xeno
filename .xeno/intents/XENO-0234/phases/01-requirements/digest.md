---
intent: github.com/triplem/xeno#183
phase: 01-requirements
created: "2026-10-03T17:31:04Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bfcc88d7238601c5b5088ad8a0cd3c57625ecd4e88c7c2c71569a93bc8c0cb3a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
A matching plugin passes and the digest is recomputed over the tree rather than read from it; a
one-byte difference fails with a finding naming both digests, because one saying only that they
differ leaves the reader to compute two values the gate has; an earlier plugin fails on the digest
alone, which is how section 13 settles the downgrade without a version check; no plugin at all with
an anchored runner fails rather than reporting not-implemented, since the check ran and what it was
anchored to is absent; and a binary with no digest reports `not-implemented` with no findings, which
is a development build and is this repository. The criterion doing the real work is that the anchor
is not readable from the repository: a file naming the right digest, anywhere including inside the
plugin, must not rescue a changed tree — every other criterion is satisfied by an implementation that
reads its expectation from the tree, so this is the one that separates them. The gate runs first and
from P0; the release computes the digest with the same code the gate compares with, because a digest
computed differently is a red verdict on every project that installs the release; a released binary
reports `pass` where the dev build reports `not-implemented` and `gate verify` stays at exit 0 across
that, since both derive to green. Out of scope: the resolution order, which this was recommended
ahead of so the containment is not built with the thing it contains; embedding the plugin, a real gap
A58 is adjacent to; reading the lock's block or the two version halves, which section 13 excludes
because both are in the tree; and signatures, which section 13 disclaims.
