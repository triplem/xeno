---
intent: github.com/triplem/xeno#332
phase: 02-design
created: "2026-10-09T13:32:25Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c68b98d38e5edd4b8d2723f9cb72f99570612587c08dcbbfd95e71d339907c85
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**Two constants change and nothing else about how they are read.** `ApprovedLabel` and
`ApprovedWord` hold the new strings; `Approval()` keeps its comparison, so `/Xeno
approved.` and `/xeno approved:` are one first line as `Approved.` and `approved:` were.
The refusal and the intake's sentence print the constants and so say the new names
without a word of their own changing; the one wording that moves is "the word", since
`/xeno approved` is two.

**The script lives where the entry point lives.** `.xeno/plugin/bin/xeno-labels.sh`, a
POSIX shell script in `xeno-env.sh`'s shape: a header saying what it is and why, `set
-eu`, the host as the first argument and the project as the second, `gh label create`
or `glab label create` with the name, a colour and a description that says what the
label starts, and a read first so that a label already there is reported and left
alone. Running it twice changes nothing the second time, which is the property
`github-settings.sh` has and the test asserts by shape rather than by running the host.
It is vendored by `init` with the executable bit, as `bin/` already is.

**The description is the clause in one line.** `An issue xeno intent start may start;
needs a comment beginning /xeno approved as well`, so that a reader of the label list
sees the other half.

**init names it, in the shape of the other five.** One string in `Manual`: `create the
label xeno-approved, with .xeno/plugin/bin/xeno-labels.sh or by hand; an issue becomes
an intent only with it and a comment beginning /xeno approved`. Named and not created,
like the other five, for the reason the comment above the list gives: a command that
claimed to have made a host setting would be lying.

**The fixtures move by search and replace, and one test checks the old names are gone.**
Every `"approved"` label and `approved\n` first line in the five test files becomes the
new string; `identity_test.go` gains a case asserting that the old pair is now two
missing halves, so that the rename cannot silently become an addition.

**The host is relabelled at the merge, by hand.** Create `xeno-approved`, add it to the
fourteen issues, remove `approved` from them, delete the label; four `gh` calls recorded
on the pull request. At the merge and not before, because until the new binary is on
`main` a start on this repository reads the old names.

**A107.** The comparison unchanged, the label created and never read back, the old names
gone rather than accepted, and the host act recorded on the pull request.

<!-- xeno:section:alternatives -->
## Alternatives

**Both pairs accepted for a release.** Softer for projects mid-flight. Declined: the clause fixes two names, and a trail where two forms counted is a trail where a reader has to know the date.

**The script in `scripts/` only.** This repository's, not the deliverable; the maintainer asked for the deliverable and the clause now names it.

**`xeno init` creating the label.** Through the adapter, a fifth operation on a contract WP12 fixes at four; and `init` names host settings and never sets them. Declined.

**A label description without the comment half.** Shorter, and a reader of the label list would set the label and wonder why the start refused. Declined.

<!-- xeno:section:impact -->
## Impact

**For a project.** Two names to use instead of two others, one script to run once, one line more in `init`'s output. An issue approved under the old names is refused until a person writes the new comment; the label is a relabel.

**For this repository.** Fourteen issues relabelled at the merge; the open ones with an `approved` comment need a `/xeno approved` one before their intents can start.

**For the trail.** Nothing sealed changes. Four intakes quote the old names and are right about their day.

**For the plugin.** One more file under `bin/`, so the vendored tree and the digest a release compiles move with the next release, as any plugin change does.
