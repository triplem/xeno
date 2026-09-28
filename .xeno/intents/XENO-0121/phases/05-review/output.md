---
intent: github.com/triplem/xeno#121
phase: 05-review
created: "2026-09-28T19:45:41Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c2ba6b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: ca1be237f71140b8f00593b0de397bf156dc01faa7f344a4e23367c4430ceb61
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

**The specification changed first, by a person, in its own commit.** #122 rewrote the
release automation sentence and carried nothing else, because squash is the only merge
method here and a bundled change would have collapsed the ordering the first standing
rule is about. Nothing under `docs/` is in this diff.

**Nothing was invented.** Two plugins, one file, one workflow input and two table rows
are gone. No changelog by another route, no pull request per release, no file among the
release assets.

**The protection was not touched.** `scripts/github-settings.sh` differs from `main` by
zero lines, and no secret is referenced anywhere in the diff. The bypass and the token
were refused in
#121 and not revisited.

**Every place that named the removed mechanism was found.** One grep over every
Markdown, JSON, YAML and shell file outside `vendor/` returns a single line, the
superseded row. The design's list of files was one short and the grep caught it, which
is recorded in the deviations rather than absorbed.

**A23 is superseded rather than amended**, because its subject stopped being true. The
three intents before this amended rows that were right and unfinished, and reaching for
the same move here would have appended a qualification to a sentence contradicting it.

**The comments describe the present.** The `GITHUB_TOKEN` property that once explained
why the changelog commit was harmless now explains why it could not pass, and it is
written that way round. Leaving it would have been #111's defect in a workflow file.

**Every verdict still matches.** `gate verify` over 73.

**The phases were written in their order.** P0 to P2 before any configuration changed,
the change inside P3, the sweep and the transcript inside P4.

**What the checklist cannot cover is named in the mapping.** AC9, that a release is cut,
is a property of the next push and appears as a prediction rather than as evidence.

<!-- xeno:section:release-notes -->
## Release notes

Releases work again. Nothing has been published since v0.18.0 on 2026-09-26, because the
release workflow failed on every merge: `@semantic-release/git` pushed a `CHANGELOG.md`
commit to `main`, and the protection set in #86 refuses a direct push that cannot carry
the required `verify` check.

The changelog is no longer written back to the tree. `@semantic-release/changelog` and
`@semantic-release/git` are out of `.releaserc.json`, and `CHANGELOG.md` is removed
rather than left frozen four releases short of the truth. The release notes of each
GitHub release carry what the file rendered, and `git log -- CHANGELOG.md` still has
what it rendered before.

Everything else about a release is unchanged: the version and the notes come from the
Conventional Commits history, the tag is published, five platform binaries carry the
version through ldflags, the module set is checked, the bill of materials and
`SHA256SUMS` are written and uploaded.

The branch protection is untouched and no credential was added. The alternative was a
bypass actor or a token able to push past a required check, which would have reopened
what #86 closed.

`SUPPLY-CHAIN.md` and the audit workflow no longer pin two packages nobody installs.

<!-- xeno:section:residual-risk -->
## Residual risk

**The fix is unproved until this merges.** Every gate in this intent is green and none
of them reads a file the intent changed. The first evidence that the pipeline works is
the release run on the merge of this pull request, and if it fails, the diagnosis in
#121 was wrong about the cause rather than incomplete.

**Nothing local reads the release configuration.** `.releaserc.json` and the workflow
files are checked here for being parseable and for not naming what was removed. A plugin
list that is valid and wrong is green, and this intent adds no check for that.

**`SUPPLY-CHAIN.md` stays hand maintained against three other files.** Four places named
the plugin set and the sweep found a fifth mention. Nothing makes them agree, and the
document whose purpose is to prevent supply chain drift drifts by hand.

**The release history left the tree.** A reader who never knew `CHANGELOG.md` existed
will not find it in `git log`, and the releases page is now the only place the rendering
is visible without knowing what to look for.

**Any future step wanting to write to `main` from a pipeline meets the same wall.** That
is #86's arrangement and this change accepts it rather than working around it. The next
one to hit it should read A23 and this record before proposing a bypass.

**Accepted with the five named.** None blocks the change. Each is smaller than the state
it replaces, which was a pipeline that failed on every merge and a repository that had
not released in two days.
