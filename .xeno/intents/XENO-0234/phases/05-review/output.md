---
intent: github.com/triplem/xeno#183
phase: 05-review
created: "2026-10-03T17:33:59Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a6884dd47b63307b54fc9f91659ebf97baf0581faaa428c1cb262eef500dc136
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
      Three: the ldflags line split across three assignments because one line with both -X paths
      runs to 101 columns; ASSUMPTIONS.md's stale gate list corrected for G-Rules and G-Policy as
      well, which the design did not ask for; and two sealed verdicts rewritten during the work and
      restored from git.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      A gate that reported not-implemented now reports a judgement, so a changed vendored plugin
      becomes red on every phase from P0 for anybody on a released binary, where it was silent.
      Nothing migrates: sealed verdicts keep what they recorded and gate verify stays at exit 0,
      because it compares a phase's status and both results derive to green.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Three, each naming what it departs from: the ldflags line split
across three assignments because one line with both `-X` paths runs to 101 columns; `ASSUMPTIONS.md`'s
stale list corrected for G-Rules and G-Policy as well as G-Supply, which the design did not ask for;
and two sealed verdicts rewritten during the work and restored.

**`interface-change-needs-a-migration-note` — met.** A gate that reported `not-implemented` now
reports a judgement, which changes what a verdict means for anybody running a released binary: a
changed vendored plugin becomes red on every phase from P0, where it was silent. Nothing migrates —
sealed verdicts keep what they recorded and `gate verify` stays at exit 0 across the difference,
because it compares a phase's status and both results derive to green. The README says why a reader
of this trail meets `not-implemented` everywhere, which is the part somebody will ask about.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added.

**Beyond the three rules.**

*Does the gate do what section 13 asks, and only that?* Yes, and the test that proves it is the one
asserting the anchor is not readable from the tree. The lock's `plugin` block, added by the intent
before this one, was the most plausible wrong answer and is refused by the sentence that makes the
arrangement worth having: an expected hash stored beside the thing it describes proves only that both
were written by the same hand.

*Was anything invented?* No result value, no field, no gate, no rule. The no-anchor case reuses
`not-implemented`, which section 5 defines for a check a runner did not perform, and the reason it
fits is the reason that state exists — a silent skip makes a green verdict mean less than it appears
to.

*Is A42 intact?* Yes, checked as the `verify` job checks it. The gate path gained
`internal/plugin`, which reads files and `internal/hashing`.

*What does this change about the recommendation it came from?* The recommendation was to build this
before Shape 3, so that the containment Shape 3 needs is self-checking rather than trusted. It is,
for a released binary. Under a development build the anchor is absent and a swapped plugin is still
silence, so the containment is only as good as the binary running it — which is worth knowing before
section 7's decision rather than after.

*What did the review find that the phases before it had not?* Nothing new, and the verification phase
found the largest gap while writing it: a released binary cannot vendor the plugin it carries the
digest of. That is a distribution question rather than a gate question and it wants its own issue.

*Could this change a verdict behind it?* No. `gate verify` is 279 at exit 0 and every sealed phase
keeps the `not-implemented` it recorded, which was true of the binary that wrote it.

<!-- xeno:section:release-notes -->
## Release notes

**G-Supply is implemented.** It recomputes the digest over `.xeno/plugin/` and compares it against
one the release compiled into the binary.

| what it finds | what it says |
|---|---|
| the plugin the runner was released with | `pass` |
| a plugin that differs by any amount | `fail`, naming the digest found and the one expected |
| no vendored plugin | `fail`, saying to run `xeno init --vendor` |
| no digest in the binary | `not-implemented` |

**The anchor is the binary and nothing in the repository.** Section 13's reason: "an expected hash
stored beside the thing it describes proves only that both were written by the same hand." So the
gate reads neither the lock's `plugin.sha256` nor any version — an earlier plugin fails on the digest
alone, which is what makes an upgrade one act: new runner, `xeno init --vendor`, one commit.

**A development build reports `not-implemented`**, because `go build` compiles in no digest. That is
section 5's state for a check a runner did not perform, and it is what every phase of this
repository's own trail records. A released binary over the same tree reports `pass`.

**It runs first, from P0.** Section 7 puts it there because a plugin that does not match its expected
digest makes every later verdict a statement about unknown rules and unknown templates.

**What this is not.** Section 13: "This is not tamper protection and does not pretend to be. Whoever
can rewrite the vendored plugin can rewrite rules, artifacts and the CI wrapper as well. What
G-Supply gives is that a changed shipped set cannot pass unnoticed."

<!-- xeno:section:residual-risk -->
## Residual risk

**A released binary cannot vendor the plugin it carries the digest of.** `xeno init --vendor` copies
from the repository's own tree, so an adopter running a release in a fresh repository has nothing to
copy, and this gate is anchored to something the release cannot deliver. The largest remaining piece,
found in the verification phase, and it wants its own issue. Until then the gate is correct and the
path to satisfying it exists only in a repository that already has the plugin.

**The first real test is the next release.** Every measurement here used a locally built binary with
the release's flags. A mis-assembled ldflags line would compile in an empty anchor and switch the
gate off for that release silently — the script's non-zero exit guards a missing plugin and guards
nothing against a quoting mistake one line later.

**Under a development build the gate is inert**, which is most of what happens in this repository. The
containment that Shape 3's resolution order would rely on is therefore only as good as the binary
running it, and that is worth knowing before section 7's decision rather than after.

**A changed plugin turns every phase red at once**, including sealed ones, for anybody on a released
binary. That is section 7's intent — a mismatch makes every later verdict a statement about unknown
rules — and it means an adopter who edits the vendored tree meets a repository-wide red rather than a
local one. The finding says what to do; the volume will still be a surprise.

**The digest covers whitespace.** Two vendored copies of one released plugin differing in a comment
have different digests, so this gate is strict in a way a version comparison would not be. Section 13
asks for exactly that; it is recorded because the first person to hit it will think it a bug.

**Nothing confirms a published binary carries the tag's digest.** The release computes it over its own
checkout, which is right by construction, and no later check recomputes it from the tag against the
artifact.

**What is not a risk.** Any verdict behind this intent, and any repository running a development
build, whose verdicts are unchanged.
