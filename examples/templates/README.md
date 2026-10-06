<!-- SPDX-License-Identifier: Apache-2.0 -->

# A candidate reduced section set

Six template overrides that keep **6 of the 17 required sections**. Not shipped, not
enabled anywhere, and not a recommendation. It exists to be tried and measured.

## Why it exists

`#117` has nine six-phase intents measured. The record is a fixed cost:

| | across nine intents |
|---|---|
| required sections written | **17**, every time |
| files under `.xeno/` | 32–33 |
| lines added under `.xeno/` | **1,350 – 1,499** |
| lines changed outside `.xeno/` | **13 – 440** |

The record spans 11% and what it describes spans a factor of thirty-four. XENO-0254 is
the extreme: thirteen lines of one document behind about 1,400 lines of record.

So proportionality is not a ratio a shortcut could improve by a percentage — the
numerator does not move with the change. The only lever is **which of the seventeen a
small change may omit**, and the implementation plan says the shape of a shortcut should
follow from measurement rather than be anticipated. This is the thing to measure.

## What it keeps, and on what grounds

| section | phase | why it stays required |
|---|---|---|
| `release-notes` | review | **a gate reads it.** `release-notes-are-filled` is a shipped `checked` rule whose `section-non-empty` check names this id — the only section id any rule names |
| `acceptance-criteria` | requirements | by decision, not by a reader: #258 settled that section 9 will make a criterion identifiable and a gate will read the mapping against it |
| `test-mapping` | verification | the same decision, from the other side |
| `problem` | intake | why the change exists, which nothing else carries |
| `changes` | implementation | what was done |
| `results` | verification | whether it worked |

The eleven it drops each carry their reason in the file that drops them.

## The measurement behind it

Of the seventeen required sections, **one is read by a gate and sixteen are read by the
next-step suggestion.**

No Go file outside the tests names any section by id — the one apparent hit, `Scope` in
`internal/rules/rules.go`, is a rule's own scope field and not the `scope` section.
Across the shipped rules and the examples, exactly one section id appears:
`release-notes`. `sectionNonEmpty`'s own comment says the general case out loud and
cites A73: *"a required section carrying nothing is read only by the next-step
suggestion and by no gate"*.

**G-Policy is not a counter-example.** Its checklist requirement reads the
`review_checklist` frontmatter list, not the `review-checklist` section, so making that
section optional does not make the answers optional. The review override says so.

Counted on 2026-10-06, against the seventeen that nine intents were measured over.

## How it works, and why nothing is deleted

Every section the shipped template defines is still defined here, with the same `id`,
`version` and `phase`. Only the `required` flag moves.

That is the whole of the lever: `Resolved.Missing` reports required sections that carry
nothing, and nothing else reads the flag. Keeping the ids means `xeno section set` still
accepts all seventeen, a rule naming one cannot find it unknown, and an author who wants
a dropped section simply writes it.

## Trying it

```sh
cp -R examples/templates/. .xeno/config/templates/
```

A project template beats the shipped one **per id**, so copying a subset is fine — take
`intake/` alone if that is the experiment. `strings.en.yaml` must come with each
directory: a missing bundle is an error and never a fall back to another language, so a
project on another language writes its own.

Nothing here is enabled by its presence under `examples/`. `template.Load` reads
`.xeno/config/templates` and `.xeno/plugin/templates` and never this directory.

## What to measure, and against what

Run one intent through the reduced set and collect what `#117` collects:

- required sections written
- files under `.xeno/`
- lines added under `.xeno/`
- words across `output.md`, `digest.md` and `learning.yaml`
- lines changed outside `.xeno/`

Compare against the nine already recorded. The comparison is valid because the only
thing that changed is which sections `Missing` asks for.

**Do not measure a timing.** Four comments on `#117` reported per-phase elapsed figures
from `context.lock.yaml` and `gate.yaml` before noticing they measure the agent's
session shape rather than the process: batching a phase's commands into one shell call
gives 0s, splitting them gives minutes, for the same work. The counts are what the trail
can honestly report.

## What this does not claim

Nobody has run it. The candidate parses, its ids match the shipped set, and the one
section a gate reads is still required — but whether six sections produce an artifact a
person can review in six months is exactly the question, and only running it answers
that.

It can also age without anybody noticing. If the shipped templates gain a section, or if
`release-notes-are-filled` stops being the only rule naming one, this describes a set
that no longer exists. Nothing reads these files, so nothing will report it. The date
above is the defence, which is the same one `docs/clause-readers.md` has.
