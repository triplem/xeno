---
intent: github.com/triplem/xeno#118
phase: 05-review
created: "2026-09-28T20:17:22Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57e9207.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: fdab7c47f754a64a9be5b69569f53d88dc5d764f837446fa0bd60fdd161f69bb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed.** Neither normative document states the form of a key:
they say a key exists and that it is the join in a merge commit message, and both remain
true. The form was stated in `CLAUDE.md` alone, which is where it changed.

**Nothing was invented.** No field, no gate, no dependency, no command. `intent status`
took a required argument and now takes it optionally, which is what `gate verify`
already does.

**Nothing was renamed and no verdict moved.** `gate verify` over 79. The fifty-five
intents created under the old scheme keep their keys, their hashes and the merge commits
that name them.

**The key stopped carrying a fact that is recorded better elsewhere.** `intent.yaml`'s
`intent` field gives the host, the repository and the issue; the key gave a number, and
paid for it with the sort order.

**The sequence starts where it cannot collide.** Old keys reach `XENO-0121`, the new
sequence begins at `XENO-0200`, and the gap is the only marker needed. A boundary at the
next free number would have moved when an issue reached it.

**This intent is the first member of its own sequence.** `XENO-0200`, not `XENO-0118`,
so the change is explained in the half of the directory it belongs to.

**The order comes from a recorded field rather than from a name.** Which is the whole
argument: a name sorts one way, and the next question about order will not need a
rename.

**A criterion caught the defect the change would otherwise have shipped.** The first
implementation sorted a truncated date and reproduced the issue inside its own fix. That
is in the deviations, in the digest and in a learning, rather than quietly corrected.

**The phases were written in their order.** P0 to P2 before any code, the code inside
P3.

<!-- xeno:section:release-notes -->
## Release notes

`xeno intent status` without `--intent` lists every intent in the order it was created:
one row with the date, the key, the intent's own status, and the furthest phase carrying
a verdict. With `--intent` it reports one intent's phases exactly as before.

The order comes from the `created` field that every `intent.yaml` has always recorded.
Sorted by key, this repository reads as its issues were filed; sorted by creation, it
reads as the work happened.

Intent keys are now a sequence of their own, beginning at `XENO-0200`. The issue an
intent belongs to is the `intent` field of `intent.yaml`, which carries the host and the
repository with it.

Keys assigned before this stay as they are, up to `XENO-0121`. The key sits inside
`artifacts_hash` and inside the merge commits that name an intent, so renaming one would
change every verdict in it. The gap between `0121` and `0200` is how a reader tells
which scheme a key follows.

An intent whose `intent.yaml` cannot be read, or which records no creation date, appears
in the listing with the reason rather than being left out.

<!-- xeno:section:residual-risk -->
## Residual risk

**Nothing enforces the new scheme.** The key is opaque to the tool, deliberately, so an
intent created tomorrow as `XENO-0119` would be accepted and would sit in the middle of
the old range. The rule lives in `CLAUDE.md` and in no check. That is the same class of
gap as the one this session has met twice, a convention with no gate behind it, and it
is the first thing to close if the scheme matters.

**Nothing allocates the next number.** It is read off the listing and incremented by
hand. A gap in the sequence would never be noticed, and two branches could take the same
number before either merged; the collision would surface as a directory that already
exists, which is a poor error for a good reason.

**Two schemes in the tree forever.** The gap makes them legible and no rule makes them
uniform. A reader who does not know the history sees a range of keys that stops at 0121
and resumes at 0200 and has to be told why, which is what `CLAUDE.md` now does and what
nothing else does.

**The listing shows a problem no gate judges.** An unreadable `intent.yaml` becomes a
row with a reason, and nothing fails. No gate looks inside an intent directory outside a
phase, which is #109, still open; this change makes the state visible to a person
without judging it.

**Creation, not commencement.** An intent created and started much later sorts by its
creation. P2 rejected reading the first phase's lock because an intent with no phases
would have no place in the list, and this imprecision is the price.

**Accepted with the five named.** None blocks the change, and each is smaller than the
state it replaces, which was a record of work whose only listing order was the order its
issues were filed.
