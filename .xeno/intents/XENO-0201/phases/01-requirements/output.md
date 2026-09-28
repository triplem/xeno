---
intent: github.com/triplem/xeno#120
phase: 01-requirements
created: "2026-09-28T20:29:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 89683c8e0ca2d5930d0d525efb6b11d3cae49bca86273feb302e58ad1bceb8ab
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

**AC1.** `xeno phase finish --summary FILE` writes `digest.md` in the phase directory:
the frontmatter section 5 requires of a file produced in a session, then the summary
text. It reads stdin where `--summary` is given without a path, as `section set` does.

**AC2.** The digest it writes carries `intent`, `phase`, `created`, `schema_version`,
`runner_version`, `plugin_version`, `language`, `context_hash`, `model` and `tool`, in
the order section 5 lists them, and carries no `template`, no `strings_hash` and no
`rules_hash`, because a digest is neither rendered nor covered by a rule set.

**AC3.** It carries no `secrets_hash`. There is no filter, so the field is absent rather
than holding a hash of nothing, and G-Schema reports it missing.

**AC4.** `phase finish` without a summary behaves exactly as it does today: it writes no
digest, judges whatever is there, and reports the file missing where it is missing.

**AC5.** A digest the runner writes passes G-Schema and G-Trace on every field they read
except the ones with no writer, and a phase whose digest was written this way reaches
the same verdict as one whose digest was written by hand with the same content.

**AC6.** `model` and `tool` are read from the `agent` block of `project.yaml` and
written into `output.md` by `section set` as well.

**AC7.** A project file with no `agent` block, or one naming neither field, leaves both
absent rather than writing a guess. The same holds for a project file that cannot be
read at all.

**AC8.** A35 names the three fields that still have no writer, `tool_version`,
`secrets_hash` and `rules_hash`, and records which two gained one.

**AC9.** Everything green stays green: `gate verify` matches every verdict, and the
sixty intents whose digests were written by hand are untouched.

<!-- xeno:section:non-goals -->
## Non goals

The filter. No pattern file is shipped, added or invented, `secrets_hash` gets no
writer, and nothing in a digest is redacted. Section 5 says the runner filters; this
intent makes the runner write, and the filtering waits for the file it filters against.

`tool_version`. A property of the session, absent from the block that carries `tool`,
and A35 keeps it.

`rules_hash`. WP4.

A digest command of its own. The working sequence puts the digest at `phase finish` and
this follows it.

Rewriting the digests that exist. Sixty intents carry hand written ones, inside
`artifacts_hash`, and nothing regenerates them.

Enforcement of any kind. Nothing is forbidden, nothing is blocked, and #120's question
stays open with its specification change still outstanding.

A summary the runner composes. The agent supplies the text; the runner does not
summarise, shorten or rewrite it. Section 5 divides the labour and this respects the
division.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes. Section 5 and the working sequence already say what this
builds, and they say it about a runner that did not do it.

The summary stays optional. Section 5 says the supported path is not an enforced one, so
a `phase finish` that refused a phase without a summary would contradict the sentence
being implemented.

The frontmatter of a digest is section 5's, minus what a rendered file adds. Three
groups, not two: the digest is in the session group and not the rendered one, so a
template or a strings hash on it would be wrong rather than merely unnecessary.

The field order is the order section 5 lists, which the existing `frontmatterOrder`
already encodes. One order for both artifacts, or a reader comparing them has to know
which writer produced which.

A missing or unreadable project file must not fail a phase. The language already
degrades to a default there, and the two new fields degrade to absent.

No new dependency, no new gate, no new field. Every field written is one section 5
enumerates.

One intent, one issue. The commits reference #120, which is where the scoping that chose
this work is recorded.
