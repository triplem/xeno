---
intent: github.com/triplem/xeno#260
phase: 02-design
created: "2026-10-06T07:24:00Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a4050616e5e139a3adefa5866231ad2e25205dcda6f0f853e39ec14b3c16223b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

The `advisories` prose is replaced rather than amended. It currently narrates an upgrade from
24.2.9 to 25.0.9 and asserts that `postcss-selector-parser` went with it, which is false today;
editing into a paragraph whose central claim has failed is how the words around it get left behind,
which is the conventions' own reason for replacing rather than editing.

The new prose answers three questions in this order: what is there now, why the file drifted, and
what was tried. The order matters because a reader arrives at this file after a failing job and
wants the count first; the drift explains why their pull request failed for nothing they did; and
the attempted fixes are what stops the next reader repeating the experiment.

Both directions go in the same sentence. "high 30 to 10, moderate 1 to 2" rather than two
sentences, because the file's rule is about a number rising and the honest statement is that one
rose while another fell by twenty — splitting them invites quoting the half that flatters.

A `drift` field is not added. The temptation is a machine-readable note that the tree is unpinned,
and the second standing rule makes an added field a specification change where the artifacts are
concerned; this file is CI configuration rather than an artifact, so the rule does not reach it —
but nothing reads a new field either, and a sentence in the prose is read by the person who needs
it. Prose where the reader is a person, which is the same judgement #255 made three times.

`measured` becomes 2026-10-06 and `toolchain` stays `semantic-release@25.0.9`, now with the fact
that it is also the latest. That one line is what makes the next reader's first question —
"can we just upgrade?" — answerable from the file instead of from an experiment.

No register row. A44 already records the arrangement and the arrangement is unchanged; what moved
is the measurement A44 points at, and putting numbers in the register as well as in the file would
be two places to keep true. The drift finding is a learning, which section 10 routes.

`Refs #260`. The issue asks that the audit pass *and* that the route be recorded; this does the
first and records this route, and the lockfile question is the unfinished half, so `Closes` would
be a claim the diff does not support.

<!-- xeno:section:alternatives -->
## Alternatives

Committing a lockfile instead was the stronger route and is the open half of #260. It is the only
one that stops the drift, because the counts would become a function of this repository rather than
of npm's registry. It is not this intent for two reasons: it would record the same 10 and 2 at pin
time, so it needs this measurement anyway, and committing a dependency tree for a toolchain this
repository installs rather than ships is a decision under A28 that wants taking on its own terms
rather than as a side effect of unblocking a queue.

Narrowing the check to critical and high was considered. It would have passed today — high fell —
and it gives up exactly what A44 wanted the moderate count for, which is noticing that something
changed. A check narrowed until it stops failing is the thing that pattern produces.

Narrowing it to advisories with an applicable fix was considered and is more interesting, because
it is measurable: five of the twelve say `fixAvailable: true`. But `npm audit fix` resolves none of
them, so the five are not actually actionable, and a check keyed on npm's own claim about fixability
would be keyed on a field that is wrong here in every instance.

Downgrading `semantic-release` to 15.14.0 is npm's own advice for seven of the twelve and is
rejected on A44 and A28 together, and on the obvious: a major downgrade of the release machinery to
silence advisories in a tool that only ever runs in CI.

Leaving the baseline and marking the audit job non-blocking was considered. It would unblock
everything immediately and it converts a check into a notification, which is the end state of every
check nobody can act on. A44's whole reason for the baseline shape was to avoid that.

Setting `high` to 10 and leaving `moderate` at 1 was considered — tighten only, loosen nothing, and
let the job keep failing on the one advisory. It is the most conservative reading and it is
dishonest about what was measured: the moderate count is 2, and a baseline that records 1 is not a
baseline but a wish.

<!-- xeno:section:impact -->
## Impact

One file, four values and one paragraph. `.github/npm-audit-baseline.json` is the whole diff.

The audit job starts passing, which unblocks #259 and everything behind it. That is the point and
it is also the only behavioural effect: nothing else reads this file, no Go code touches it, and it
is outside `artifacts_hash`, so the trail is indifferent.

The check gets stricter overall. The high bound drops from 30 to 10, so twenty high advisories that
could have arrived unreported now cannot; the moderate bound rises by one, which is the single
advisory that actually appeared. A reader who sees only "a number was raised" has the story
backwards, and the commit message is where that is corrected.

What the file gains that it did not have is the answer to its own next question. Its prose explains
an upgrade; it has never said that there is nothing left to upgrade to, or that `npm audit fix`
achieves nothing, or that the tree is unpinned and therefore drifts. The next person to meet a
failing audit reads the file rather than repeating the experiment, which cost one background install
of two trees to establish.

What does not change is that it will drift again. The mechanism is untouched: the next upstream
publication can fail an unrelated pull request exactly as this one did, and the only defence
recorded is a date. The prose says so rather than letting the new numbers read as a bound, and the
lockfile question stays open on #260 where it can be decided on its merits.

The honest cost is that this is the second time in three days that the baseline has needed a
deliberate act — #189 raised it, #260 re-measures it — and nobody has decided whether a file
requiring that cadence should be a file. Two data points is not a trend; it is enough to be worth
the sentence.
