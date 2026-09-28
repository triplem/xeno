---
intent: github.com/triplem/xeno#118
phase: 00-intake
created: "2026-09-28T20:11:37Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57e9207.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 0a88e621dc3c7ffb30fea5f22d432cdc2ea7a9189e39fc86daaaa4f821f6296b
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

An intent key is the issue number padded to four digits, so `.xeno/intents/` sorts by
the order the issues were filed. That is not the order the work happened, and after
fifty-five intents the difference is no longer theoretical: XENO-0107 was started after
XENO-0111 and before XENO-0121, and a directory listing puts it two places before
either.

So the question a reader asks first of a record of work — what happened, in what order —
is the one the layout answers worst. Every intent already carries the answer.
`intent.yaml` records `created`, and it has done since the first intent; nothing reads
it for this purpose.

There is no way to ask, either. `xeno intent status` requires `--intent` and reports the
phases of one intent. Nothing lists intents at all, so the order of work is available
only by opening fifty-five files or by sorting a directory that sorts by something else.

The padding makes this worse rather than better in the one way it was supposed to help.
It was chosen for sorting alone, and the sort it produces is over issue numbers, so the
only thing the padding guarantees is that the wrong order is stable.

Two intents of this session show the shape. XENO-0107 and XENO-0121 were both started on
2026-09-28, after XENO-0111 of the same day; sorted by key the session reads 0107, 0108,
0111, 0121, and sorted by `created` it reads 0108, 0111, 0107, 0121, which is what
happened.

<!-- xeno:section:scope -->
## Scope

`xeno intent status` without `--intent` lists every intent in the order it was created,
one line each, from the `created` field already in every `intent.yaml`. With `--intent`
it reports the phases of one intent exactly as it does today.

The key scheme changes for intents created from here on. A key is `XENO-` and the next
number in a sequence of its own, beginning at `XENO-0200`. The issue an intent belongs
to stays where it already is, the `intent` field of `intent.yaml`, which carries the
full reference rather than a number.

`CLAUDE.md` records the new rule and what happens to the old keys. The fifty-five
existing intents keep theirs, because the key sits inside `artifacts_hash` and inside
the merge commits that name them.

This intent is the first member of the sequence it introduces, `XENO-0200`, rather than
`XENO-0118` under a rule it retires.

Not a rename. Nothing existing moves, and no verdict is recomputed.

Not a second sort order in the tree. The directory still sorts by name; the tool is what
answers the question, and the design says why that is the right place rather than a
shortcoming.

Not a new command. `intent status` exists; its argument becomes optional.

Nothing in `docs/`. Neither normative document states the form of a key, only that a key
exists and is the join in a merge commit message. Both remain true.

<!-- xeno:section:context-rationale -->
## Why this context

**The key stops carrying the issue number because it was never the right carrier.** A
key is an identifier the tool treats as opaque: nothing in the code derives it, and the
only check is that `intent.yaml`'s `key` matches the directory it lies in. Meanwhile the
issue is recorded properly one line above, as `github.com/triplem/xeno#118`, with the
host and the repository in it. The number in the key was a second, lossier copy of that,
and it was paying for the copy with the sort order.

**The sequence starts at `XENO-0200` so that no new key can collide with an old one.**
The existing keys run up to `XENO-0121` because issue numbers reach that far, and a
sequence starting at fifty-six would land on `XENO-0056`, which exists. Starting at two
hundred leaves a gap that can never be filled, and a gap is cheaper than a rule about
which scheme a key belongs to. It also means the scheme a key follows is legible from
the number alone.

**Nothing is renamed, and that is not a compromise.** `CLAUDE.md` already states why:
the key sits inside `artifacts_hash` and inside the merge commits that name an intent,
so renaming one changes every verdict in it. A change that rewrote fifty-five intents
would have to recompute fifty-five sets of verdicts and could not leave the merge
commits consistent with them at all. Forward-only is the only shape available, and the
gap at two hundred is what keeps the two schemes apart without a rule.

**The tool answers the question rather than the file system.** A directory sorts by
name, and choosing a name so that the sort comes out right is what produced this issue.
`created` is recorded, so the order of work is a query and not a property of a string,
and a query can be asked in other orders later without renaming anything.

**`intent status` grows an optional argument rather than gaining a sibling.** The
command surface is a budget rather than a list. Asking about all intents instead of one
is the same question at a different scope, and `gate verify` already takes `--intent`
optionally in exactly this way.

**This intent is `XENO-0200`.** Numbering it after its own issue would have made the
last key of the retired scheme the intent that retires it, and a reader of
`.xeno/intents/` would find the change explained in the wrong half of the directory.
