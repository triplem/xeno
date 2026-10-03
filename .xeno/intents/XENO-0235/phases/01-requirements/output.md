---
intent: github.com/triplem/xeno#199
phase: 01-requirements
created: "2026-10-03T18:22:38Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 924f00a8283ae31a812cf1e4b876b4348b6339523a0bd9b1440ddd996d41cf32
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**A released binary in an empty repository vendors the plugin.** No plugin on disk, no
`--plugin-from`, and `init --vendor` writes the tree from the binary and says where it came from.

**That tree's digest is the one the binary carries**, so G-Supply passes against what the release
just installed. This is the property the whole change exists for and the one that was impossible
before it.

**The adopter's `plugin_version` is the release's number.** Stamped into the shipped copy, so section
13's one shared version number holds for anybody who receives a release.

**The digest is taken after the stamp.** Taken before it, no adopter's tree could match — the
ordering is the one thing in this change that fails silently if it is wrong.

**Every file of the plugin is vendored, at any depth, with nothing named.** `secrets.yaml` and
`bin/xeno-env.sh` arrive, and so does whatever the plugin gains next.

**The vendored entry point is executable**, because the plugin's hook invokes it and an `embed.FS`
reports every file read-only whatever was committed.

**A development build carries no plugin and says so**, naming `--plugin-from` rather than the file it
could not find. Same answer `ExpectedDigest` gives in the same situation: this build is not a release
and will not pretend to be one.

**The embedded copy wins over the directory where there is one.** A release must not be talked into
vendoring a tree lying about on disk, because the digest it carries is not that tree's.

**The repository's own manifest says `0.0.0-dev`** and is never bumped again. `claude plugin
validate` accepts it.

**The test compares digests, not file lists.** A list passes while a byte differs, and the digest is
what the gate compares.

**Nothing in this repository's trail changes.** `gate verify` stays at exit 0 and every sealed
artifact keeps what it recorded.

<!-- xeno:section:non-goals -->
## Non goals

**No check that the manifest matches the newest tag.** Offered and declined, for the reason the npm
audit baseline already states about a job that fires with no action available. Stamping removes the
need for the check rather than enforcing the bump.

**No resolution order.** #183's open clause and a section 7 decision. This makes the anchor
deliverable, which that decision rests on, and does not take it.

**No signature over the shipped plugin.** Section 13 disclaims it: "signing comes with publication
and not before."

**No check that a published binary carries the tag's digest.** The release computes it over its own
checkout, right by construction, and a later check needs the published artifact. Still open from the
intent before this one.

**No `mcp.json`.** Section 13's list names it and the plugin carries none, because a client reading a
declaration of a server that does not exist fails at startup. The walk copies what is there, so it
arrives the day it does.

**No change to what the digest covers.** Still the whole tree, defined in `internal/plugin`.

**No embedding of anything else.** The scaffold files `init` writes are already embedded and A58
records why; this adds the plugin and nothing beside it.

<!-- xeno:section:constraints -->
## Constraints

**Stamp, embed, digest, in that order.** The digest must be over the bytes an adopter receives. Any
other order produces a value no vendored tree can match, and it fails for every adopter on every
phase rather than for the release that made the mistake.

**The digest includes the path.** `plugin.Hash` writes `<sum>  .xeno/plugin/<rel>` per file, so a
tree hashed anywhere but at `.xeno/plugin/` of a repository root gives a different value. The
stamping therefore happens in the release's own checkout, in place.

**A `go:embed` directive cannot reach outside its package directory.** So the release copies the tree
into the package rather than the package pointing at the tree, and the directory has to exist at
compile time for a clean checkout to build — which is why the placeholder is committed.

**An `embed.FS` carries no modes.** Every file reports read-only, so an executable bit is set on the
way out or not at all.

**`--vendor` must not succeed with a partial tree.** A tree that is nearly the plugin is worse than a
refusal, because the gate fails later and somewhere else.

**One shared version number is section 13's ask**, and a release cannot commit to main. So the copy
it ships is where the number goes.

**Absent is not a default.** A build that carries no plugin refuses rather than vendoring something
plausible, which is A35's rule and A86's answer for the sibling case.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0, with the exit code
captured and not piped.**

**One intent, one branch, one issue** — `199-the-release-ships-the-plugin`, #199, labelled wp7, on
main.
