---
intent: github.com/triplem/xeno#118
phase: 02-design
created: "2026-09-28T20:13:04Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57e9207.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: bf5553b0f8ce9aa480c721926d8bb81ec79ae5f7709b85d7ec147d7d55ead87b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The runner gains `Intents()`, and `cmdIntentStatus` branches on the key.** The listing
is a runner operation because it reads artifacts, and the command decides which of the
two questions was asked. `needsKey` becomes false for `intent status` alone, and the
command reports the missing key itself where it needs one, which is what `gate verify`
already does with the same flag.

**One row per intent: created, key, status, phase reached.** The date rather than the
timestamp, because the order is what matters and a column of identical times reads worse
than a column of dates. `status` is the intent's own, from `intent.yaml`, so an
abandoned intent says so. The phase reached is the furthest one carrying a verdict,
which is what tells a reader where an intent stopped without opening it.

**Sorted by `created`, ties broken by key.** A stable order, so two runs over one tree
print the same thing, and the tie break is the only place the key is still allowed to
decide anything.

**An intent whose date cannot be read is listed with the reason in place of the date.**
Not skipped, not sorted to the end silently: it sorts to the end because an empty string
sorts first and the sort is ascending, so it is ordered deliberately rather than by
accident of the comparison.

**No intent directory at all prints nothing and exits zero.** A repository that has not
started one is not in error, and `Intents()` returns an empty slice rather than a
failure.

**The key rule in `CLAUDE.md` is replaced, not extended.** The paragraph says what a key
is now, that the sequence starts at `XENO-0200`, where the issue lives, and that earlier
intents keep their keys for the reason already given. The file is short by instruction,
so the old sentence goes rather than acquiring a qualification.

**This intent is `XENO-0200`.** The first member of the sequence it introduces.

<!-- xeno:section:alternatives -->
## Alternatives

**A separate `xeno intent list`.** Clearer in the usage text and one command more.
Rejected because the command surface is a budget and the two questions differ only in
scope; `gate verify` already takes `--intent` optionally for the same reason, so a
sibling here would be the odd one out.

**Sort the directory instead, by renaming.** What the issue's first reading suggests:
make the key sort correctly. Rejected on cost and on kind. Renaming changes
`artifacts_hash` and the merge commits of fifty-five intents, and it would fix the order
once while leaving the next question — by status, by package, by phase reached — needing
another rename. A name sorts one way.

**Start the new sequence at `XENO-0122`, the next free number.** Tighter, and it puts
the two schemes adjacent with nothing marking the boundary. A key would then need a rule
to be interpreted, and the boundary would move again the moment an issue reached 122.
Two hundred leaves a gap that can never be filled, which is the cheapest possible
marker.

**Keep the issue number in the key and add the sequence.** `XENO-0200-0118`. Both facts
in the name, at the price of a longer key in every directory, every merge commit and
every verdict, to carry a number the `intent` field already carries with its host and
repository. Rejected as a third copy of one fact.

**Read `created` from the phase lock or the git history instead of `intent.yaml`.** Both
available and both worse: the lock belongs to a phase and not an intent, and git history
dates the commit rather than the work. `intent.yaml` records what the intent asserts
about itself, which is the thing being listed.

**Sort by the first phase's `created` rather than the intent's.** Marginally closer to
when work began, since an intent can be created and started later. Rejected because it
makes the order depend on a file in a subdirectory that may not exist, and an intent
with no phases would then have no place in the list at all.

<!-- xeno:section:impact -->
## Impact

`internal/runner`: a new `Intents()` returning one record per intent directory, and the
type it returns. It reads `intent.yaml` and the phase verdicts already readable through
the existing status path.

`cmd/xeno/main.go`: `intent status` stops requiring a key, `cmdIntentStatus` branches,
and the usage line shows the argument as optional.

`CLAUDE.md`: the key paragraph replaced.

Tests: the order over a fixture with keys and dates that disagree, the unreadable date,
the empty tree, the tie break, and the single intent form printing what it printed
before.

What a reader gains: one command answers what happened in what order, and the answer
comes from a field rather than from a naming scheme. What they lose: the issue number is
no longer visible in a directory listing, and finding which issue an intent belongs to
means reading its `intent.yaml` or the listing this change adds.

Nothing existing moves. Fifty-five intents keep their keys, their hashes and their
verdicts, and `gate verify` is the check on that.

The cost that will be paid later: two schemes in one directory forever. The gap at two
hundred makes them legible, and no rule makes them uniform.
