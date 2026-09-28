---
intent: github.com/triplem/xeno#121
phase: 02-design
created: "2026-09-28T19:41:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c2ba6b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: be09fa54a5e123142f90b9f302a98a5a085fb6c8bf3eb4fa810352791a3600e1
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

**Both plugins go, and the plugin list keeps its remaining three.**
`@semantic-release/git` pushes; `@semantic-release/changelog` writes the file it pushes.
Commit analyzer, release notes generator and github stay, which is what derives the
version, the notes, the tag and the release.

**`CHANGELOG.md` is deleted.** Its header says it is derived from the commit history and
that editing it by hand changes the rendering and not what it renders. Kept, it would
say that about a rendering four releases stale. The content survives in every release's
notes and in the file's own history.

**`release.yml` loses `extra_plugins` and the comments that explain the old
arrangement.** The header sentence about writing `CHANGELOG.md` back, and the line
offering the `GITHUB_TOKEN` property as the reason the changelog commit starts nothing,
are replaced by what is true now: nothing pushes to `main`, and the same property is why
nothing could.

**`audit.yml` installs semantic-release alone.** Its comment says it installs what the
release installs, which is a claim about another file; it stays a claim and stays
correct by having both lists shrink together.

**`SUPPLY-CHAIN.md` loses two rows.** The table is the inventory, so a row for a package
nobody installs is worse than no row: it invites somebody to keep the pin current.

**A23 is superseded, naming what replaced it.** Not amended. The rows amended here were
right and unfinished; A23's subject, that the changelog is written back to `main`, stops
being true.

**Nothing is added.** No changelog by another route, no pull request per release, no
generated file among the assets. The release notes are the rendering, and a second one
would be a mechanism this intent has no mandate for.

<!-- xeno:section:alternatives -->
## Alternatives

**Freeze `CHANGELOG.md` where it stands.** One fewer deletion, and it leaves a generated
file telling a reader it is current when it stopped at 0.18.0. The likely repair by the
next person is to edit it by hand, which the file's own header warns against. Rejected
because a stale generated artifact is a worse record than no artifact.

**Rewrite the header to say it is no longer generated.** Keeps the history visible in
the tree and truthful. Rejected as the same problem one step removed: a file whose first
paragraph explains why the rest of it is obsolete is documentation of a decision, and
the decision is already in #121, #122 and this record.

**Remove `@semantic-release/git` and keep `@semantic-release/changelog`.** The smallest
diff that unblocks the pipeline. It leaves a plugin that writes a file nothing commits,
so the file appears in the runner's working tree, is never pushed, and is installed and
audited every release for nothing. Rejected as a step that produces nothing.

**Generate the changelog into the release assets.** A file per release beside the
binaries. It keeps a rendered changelog and adds nothing to the tree. Rejected as a new
mechanism: the release notes already carry it, and a second rendering is one more thing
to keep consistent with the first.

**Amend A23 rather than supersede it.** Consistent with the last three intents, and
wrong here. Those rows were right and unfinished; this one asserts something that stops
happening, and a qualification appended to it would contradict its own subject.

<!-- xeno:section:impact -->
## Impact

`.releaserc.json`: two plugin entries out, three remain. No `changelogFile`, no
`assets`, no `message`.

`CHANGELOG.md`: deleted, 160 lines.

`.github/workflows/release.yml`: `extra_plugins` out, header and permissions comments
corrected.

`.github/workflows/audit.yml`: two install lines out, one remains.

`SUPPLY-CHAIN.md`: two table rows out.

`ASSUMPTIONS.md`: A23 superseded.

Nothing Go changes, so `go test`, `gofmt`, `go vet` and `gate verify` are unaffected
except by the artifacts this intent writes, which they judge like any others. The `xeno`
workflow does not read any file in this diff.

What the release does after this: derives the version from the commits, publishes the
tag and the notes, builds five binaries with the version in them, checks the module set,
writes the bill of materials and the checksums, and uploads. One step fewer and no push
to `main`.

What a later reader gains: the release configuration and the branch protection stop
contradicting each other, and `SUPPLY-CHAIN.md` stops pinning two packages that are not
installed. What they lose: `git log -- CHANGELOG.md` is now the only place in the tree
that renders a release history, and it takes knowing the file existed.
