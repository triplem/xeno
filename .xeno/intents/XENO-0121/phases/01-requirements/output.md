---
intent: github.com/triplem/xeno#121
phase: 01-requirements
created: "2026-09-28T19:40:34Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c2ba6b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 7cc3f1156050155642a51617479adf404b87f930f77d5e03016c8d1a6d952d4c
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

**AC1.** `.releaserc.json` carries neither `@semantic-release/changelog` nor
`@semantic-release/git`, and the plugin list is otherwise unchanged: commit analyzer,
release notes generator, github.

**AC2.** `release.yml` passes no `extra_plugins`, and its comments describe what the
pipeline does now. No comment offers the `GITHUB_TOKEN` property as a reason the
changelog commit is safe.

**AC3.** `audit.yml` installs neither plugin, and `SUPPLY-CHAIN.md` pins neither.
Nothing in the repository names a version of either.

**AC4.** `CHANGELOG.md` is gone from the tree, and nothing references it.
`.releaserc.json` carries no `changelogFile` and no `assets`.

**AC5.** A23 is superseded, with what replaced it named, and the row it replaced still
readable.

**AC6.** The release job's own gates are untouched: format, vet, test and verify run
before semantic-release as they did, and the build, module set, bill of materials and
upload steps still key off `new_release_published`.

**AC7.** Branch protection is unchanged and no secret is added.
`scripts/github-settings.sh` is not touched.

**AC8.** Everything green stays green. `go test ./...` passes, `gate verify` matches
every verdict, and the `xeno` workflow is unaffected because none of this is Go.

**AC9.** The next merge to `main` publishes a release. That one is provable only after
the merge, and the verification phase says so rather than claiming it.

<!-- xeno:section:non-goals -->
## Non goals

The protection, and any credential. #121 refused the bypass and the token, and this
intent does not revisit it. `scripts/github-settings.sh` stays as #86 wrote it.

A changelog by another route. No pull request per release, no separate branch, no file
generated into the release assets. The release notes carry it, which is the decision,
and adding a second rendering would be a new mechanism rather than the removal of a
broken one.

`@semantic-release/github`'s own configuration. Its two comment routines stay off, A43
stays as it is, and `contents: write` stays the only permission.

The version derivation. Conventional Commits still decide the number, and the tag, the
notes, the binaries, the bill of materials and the checksums are unchanged.

The releases already published. v0.18.0 and before keep their notes and their assets,
and the history of `CHANGELOG.md` keeps what it rendered.

Whatever v0.19.0 turns out to contain. This intent makes a release possible again; it
does not choose what the first one carries.

<!-- xeno:section:constraints -->
## Constraints

The plan is already changed and must not be changed again. #122 rewrote the one
sentence, by a person, in its own commit. Nothing under `docs/` is in this intent's
diff.

No invented mechanism. This removes two plugins, a file and a block of workflow input.
Adding anything to replace them would be a decision, and the decision was taken in #121.

`SUPPLY-CHAIN.md` and `audit.yml` must agree with `.releaserc.json` after this. Three
places name the plugin set and a version left behind in any of them is a pin on
something nobody installs, which is exactly the drift the document exists to prevent.

A23 is superseded rather than amended, because its subject stops being true. The rows
amended in this repository were right and unfinished; this one was right and is now
wrong.

The comments must describe the present. A fact stated as a reason for something that no
longer happens is the defect #111 fixed in godoc, and it is cheaper to avoid than to
find.

One intent, one work package. The issue carries `wp0`, the branch carries this intent,
the commits reference #121.
