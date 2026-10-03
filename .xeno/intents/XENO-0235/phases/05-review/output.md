---
intent: github.com/triplem/xeno#199
phase: 05-review
created: "2026-10-03T18:25:56Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a99f2eca6567f7882c45d8f3b507c8c418f32d1ec829a4c706b7e24d226c7c15
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      Three: PluginFrom as a new field because InitResult had nowhere to carry a note; the manifest
      bumped once from 0.30.0 to 0.0.0-dev, which is the bump that makes "never bumped again" true;
      and two sealed verdicts rewritten in the two intents before this one, which is why every
      measurement here was taken in a copy.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      init --vendor changes what it copies, from a named list to everything, and where it copies
      from, preferring the binary over a directory. An existing repository is untouched because
      create keeps what exists, which is the property that made the walk safe to substitute, and an
      adopter's plugin_version becomes the release's number.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      embed and io/fs are the standard library.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Three, each naming what it departs from: `PluginFrom` as a new
field because `InitResult` had nowhere to carry a note; the manifest bumped once, from `0.30.0` to
`0.0.0-dev`, which is the bump that makes the issue's "never bumped again" true afterwards; and two
sealed verdicts rewritten in the two intents before this one, recorded because the same mistake twice
is a pattern and because this intent's measurements were all taken in copies as a result.

**`interface-change-needs-a-migration-note` — met, and it is the substance.** `init --vendor` changes
what it copies, from a named list to everything, and where it copies from, preferring the binary over
a directory. A release now vendors its own plugin; a development build refuses where it has neither.
An existing repository is untouched — `create` keeps what exists, so a second run still changes
nothing, which is the property that made the walk safe to substitute. And `plugin_version` in an
adopter's artifacts becomes the release's number where it would have been whatever a stale manifest
said.

**`new-dependency-needs-a-rationale` — not-applicable.** `embed` and `io/fs` are the standard
library.

**Beyond the three rules.**

*Is the gate satisfiable now?* Yes, and it was not when it shipped one change ago. That is the point
of this intent and it was measured rather than argued: a release binary in an empty repository
vendors 34 files and `G-Supply` passes against them. The gate was correct and unsatisfiable, which is
a worse state than unimplemented, because `not-implemented` says so and a failing gate blames the
adopter.

*Did the ordering get recorded where it matters?* Stamp, embed, digest. It is the one thing here whose
failure lands on somebody else — a digest taken before the stamp passes for the release and fails
every project that installs it — and the reason sits between the steps in the workflow rather than in
a document, because that is where somebody editing them reads.

*Was anything invented?* No field in an artifact, no gate, no rule, no result value. `PluginFrom` is
a field of an in-memory struct that `init` prints, not of anything the trail carries.

*What did the verification phase find that the design had not?* That the embed target is tracked, so a
failed release could leave a generated tree in a dirty checkout and nothing ignored it. Fixed in this
intent, which is the right place: it was found before the change shipped rather than after.

*What does this leave for the decision it was done in service of?* The resolution order rests on
G-Supply being able to fail red, and that now means something for an adopter: the anchor is
deliverable and the gate is satisfiable. Under a development build the gate is still inert, which was
said before and is unchanged.

*Could this change a verdict behind it?* No. `gate verify` is 285 at exit 0 and every sealed artifact
keeps what it recorded.

<!-- xeno:section:release-notes -->
## Release notes

**`xeno init --vendor` copies the plugin out of the binary.** A released runner carries the plugin it
was built with, so a project with no plugin gets one, and it is the same bytes the digest G-Supply
compares was taken over.

    plugin from this release

**It copies all of it.** The vendoring named its parts before, and two were missing from the list:
`secrets.yaml`, which the secret filter reads, and `bin/xeno-env.sh`, which the plugin's own hook
invokes. A tree two files short of the released one can never match the digest, so G-Supply would
have failed for every adopter on every phase.

**The entry point arrives executable**, which an embedded copy cannot carry and `init` therefore
sets.

**A development build carries no plugin and says so**, naming `--plugin-from`. It will not vendor
something plausible.

**The release stamps its version into the plugin it ships.** Section 13 asks for one shared number
across runner and plugin; a release cannot commit one back to `main`, so it writes it into the copy
an adopter receives. An adopter's `plugin_version` is therefore the release's number.

**This repository's own manifest says `0.0.0-dev`** and is not bumped again. It had been bumped by
hand three times in one session and overtaken by a release twice.

**What a release must keep in this order:** stamp the version, embed the tree, then take the digest.
A digest taken before the stamp describes bytes nobody receives.

<!-- xeno:section:residual-risk -->
## Residual risk

**The release path is untested by a release.** Every measurement used a locally built binary and the
release's steps run by hand in a copy. Two mistakes there would be quiet: a stamp that did not happen
leaves `0.0.0-dev` in an adopter's artifacts, and a digest taken before the stamp fails for every
adopter rather than for the release. The ordering comment between the steps is the whole of the
protection, and the first real exercise is the next release.

**A partial copy would go unnoticed.** The tree is copied with `cp -R` and the digest is taken over
`.xeno/plugin/`, so a copy that failed halfway would produce a digest that does not describe the
embedded tree. Three lines comparing the two after copying would close it and are not written.

**A development build cannot exercise the embedded path.** `Shipped()` is false, so the branch every
adopter depends on runs only when somebody builds the release's way by hand. That was done once and
is not in CI.

**Nothing confirms a published binary carries the digest of the tag's tree.** Carried over from the
intent before this one and now larger, because the binary carries the tree as well as the digest.

**A released binary fails G-Supply in this repository.** Correct — the plugin here is a development
copy — and it means the only way to exercise a release against this trail is to vendor the release's
plugin over this one's, which is not a change anybody should commit.

**`mcp.json` is still absent**, so an adopter's plugin declares no server. The walk copies it the day
it exists, which is why the walk is the fix rather than a longer list.

**The embed target is generated and now ignored**, which is the right state and one a reader should
know about: `internal/plugin/embedded/plugin/` can exist locally after a release build and must not
be committed.
