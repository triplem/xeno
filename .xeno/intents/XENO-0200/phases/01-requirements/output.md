---
intent: github.com/triplem/xeno#118
phase: 01-requirements
created: "2026-09-28T20:12:16Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57e9207.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 3dd14f396a7b2d0a555d80475e620ece548e2d7bc267d23e1251e6949d6a31e8
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**AC1.** `xeno intent status` without `--intent` lists every intent under
`.xeno/intents/`, one line each, in ascending order of the `created` field of its
`intent.yaml`.

**AC2.** `xeno intent status --intent KEY` prints exactly what it prints today: one line
per phase, with the staleness note where it applies, and the suggested next step unless
`--no-next`.

**AC3.** The listing puts XENO-0107 after XENO-0111, which is when it was started, and
the fifty-five intents of this repository come out in the order they were created rather
than the order their issues were filed.

**AC4.** An intent whose `intent.yaml` is unreadable or carries no `created` is listed
rather than skipped, and says so in its line. A record that cannot be dated is a finding
for somebody to look at, not a row to hide.

**AC5.** Two intents sharing a `created` value come out in a stable order, so that two
runs of the command over one tree print the same thing.

**AC6.** The listing needs no intent to exist. A repository with no `.xeno/intents/`
prints nothing and exits zero rather than reporting an error.

**AC7.** `CLAUDE.md` states that a key is `XENO-` and the next number of a sequence
beginning at `XENO-0200`, that the issue lives in the `intent` field, and that intents
created before this keep their keys for the reason it already gives.

**AC8.** Nothing existing is renamed and no verdict changes. `gate verify` matches every
verdict in the repository, and this intent's own key is `XENO-0200`.

**AC9.** The usage text names the optional argument, since a command whose flag became
optional and does not say so is a feature nobody finds.

<!-- xeno:section:non-goals -->
## Non goals

A rename of anything. The fifty-five existing intents keep their keys and their
verdicts, and no migration is written, offered or planned.

A second command. `intent status` takes the argument optionally; nothing named `intent
list` appears, and the command surface stays the size it is.

Sorting by anything else. Not by status, not by phase reached, not in reverse. `created`
answers the question the issue asked and another order is another issue.

A machine readable form. No `--json`, no stable column contract. The output is for a
person, and a format somebody parses is a promise this intent does not make.

The directory layout. `.xeno/intents/<KEY>/` stays, one directory per key, and the key
still has to match the directory it lies in.

`docs/`. Neither normative document states the form of a key, so neither changes.

What a key should encode beyond a number. No date, no phase, no work package. The
sequence is a sequence.

<!-- xeno:section:constraints -->
## Constraints

The key is opaque to the tool and must stay so. Nothing derives it, nothing parses it,
and the one check is that `intent.yaml`'s `key` matches its directory. A listing that
read meaning out of the number would make the scheme load bearing in exactly the way
this change is removing.

`created` is read and never written. The listing reports what each intent recorded at
its creation; it does not repair, backfill or infer a date for one that lacks it.

The existing output of `intent status --intent` is a contract in practice. It is what
every phase of every intent in this repository has been read with, so the one argument
form must print what it printed before, to the column.

No new dependency and no new field. The listing reads a field that exists, in files that
exist, with the YAML reader already there.

`CLAUDE.md` is the file that states the convention, and it is short by instruction. The
new rule replaces the old paragraph rather than being added beside it.

One intent, one issue. The issue carries no work package label, which is itself worth
noting: the key scheme belongs to no package and the listing belongs to the runner. The
commits reference #118.
