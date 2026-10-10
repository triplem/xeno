---
intent: github.com/triplem/xeno#359
phase: 02-design
created: "2026-10-10T15:41:16Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 07629c2e1bb22f407de5f3d9578eabcfdc832a9f4ee25581d68769a45281b930
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
`docs/labels.md`, a page of its own, indexed in `docs/README.md` under "Running it" and
named in `zensical.toml`'s nav after "Commands". The alternatives were a section of
`CONTRIBUTING.md`, of `docs/commands.md` and of `docs/the-trail-in-this-repository.md`,
and each ran into the same thing: the content has two halves with different standing —
what Xeno reads on any tracker, and what this repository has agreed among itself — and
every existing page has already chosen a side, so putting the content in one of them makes
one half look like the other. That is the confusion #359 reports. The page is grouped by
what reads a label rather than by what a label is about, so that `xeno-approved` and
`xeno-needs-decision` are not placed side by side on the strength of a shared prefix, and
each group's table carries #359's questions as its columns, with "who sets" and "who
clears" separate because for one label they are different people. The page opens by saying
that no label holds an intent's state, because a table of labels as states invites the
reading that setting one changes something. The second label stays out of
`.xeno/plugin/bin/xeno-labels.sh`: that script creates what Xeno requires of an adopter's
tracker, and a habit of this repository in it would widen its promise. Three files change
and no code does. The learning record is that nothing ties a new page to the site's
hand-written nav, so an orphan page builds green.
