---
intent: github.com/triplem/xeno#179
phase: 05-review
created: "2026-10-03T12:19:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6a7702e7d5515a764c7dc5a928a6bb28343da58e73467655b20b484a3212a44b
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
      Two deviations, each naming what it departs from: `--intent` optional where the issue wrote it
      required, and the tracker block consolidated where the design did not ask for it. The first is
      recorded in all four phases that had a reason to mention it, because it departs from the issue
      being implemented rather than from a convention.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      The surface gains a command and takes nothing away. A hand-written `intent.yaml` is read
      exactly as before, which is a criterion of this intent rather than a note about it, and the
      seventy-nine that exist are untouched.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added. `net/url` and `strconv` are the standard library and the one vendored
      dependency is unchanged.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Two deviations, each naming what it departs from: `--intent`
optional where the issue wrote it required, and the tracker block consolidated where the design did
not ask for it. The first is recorded in all four phases that had a reason to mention it, because it
is a departure from the issue being implemented and not from a convention.

**`interface-change-needs-a-migration-note` — met.** The surface gains a command and takes nothing
away. A hand-written `intent.yaml` is read exactly as before, which is a criterion of this intent
rather than a note about it, and the seventy-nine that exist are untouched. `--intent` is optional
on the new command and required everywhere it was required before.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added. `net/url` and `strconv`
are the standard library, and the one vendored dependency is unchanged.

**Beyond the three rules.**

*Was the departure from the issue handled the way this project handles a disagreement?* The issue
fixed the command's signature with `--intent` required. The question was raised before any code was
written, answered by the maintainer, recorded in the intake's scope as a deliberate departure, and
carried through requirements, design and implementation. It is in the pull request description as
well, because the issue is what a reader will have read first.

*Is anything invented?* No field in `intent.yaml`, no field in Appendix A, no gate, no rule, no
status. The two derivations read values that are already in the repository — one in `base_url`, one
in the directory names — and both refuse rather than guessing where those are silent. That is the
second standing rule's test and A78 and A79 are the rows it asks for.

*Could this change a verdict behind it?* No. The gate list and the rule set are untouched, so
nothing recomputes differently; `gate verify` reports 241 verdicts at exit 0, against 237 before,
and the four are this intent's own phases. The existing intents are not read or rewritten by
anything here.

*What did the consolidation cost?* `model.Project` gained a type and `enforcement.go` reads it. The
fields, the tags and the behaviour are the same, and that package's tests pass untouched. It was
done because the alternative was a second declaration of one block, which is how two spellings of
one field begin.

*What arrived while this ran?* The `tool_version` gap: the field is required by G-Schema, the runner
has no source for it, and the digest is rewritten by every `phase finish`, so both artifacts of all
five phases needed the same hand edit. It is this intent's implementation learning and it wants an
issue, because it is paid twice per phase by every intent this repository has ever run.

<!-- xeno:section:release-notes -->
## Release notes

**`xeno intent start` creates an intent.** It takes the issue — `xeno intent start --for 179`, or
`--for github.com/triplem/xeno#179` where the configuration cannot supply the rest — and writes
`intent.yaml`. Everything else is derived: the key continues the sequence the intents directory
holds, the host comes from `tracker.base_url` and the repository from `tracker.project`, `created`
is the runner's clock, the two version fields are the runner's own, and `status` is `in-progress`.

**`--intent KEY` names the key instead of deriving it,** which is what a repository with no intents
yet needs.

**It refuses where the intent directory already exists**, because `intent.yaml` sits inside the
intent-level `artifacts_hash` and a second write would change a verdict that named the first one.

**A hand-written `intent.yaml` is unchanged and still works.** Nothing in this release reads or
rewrites an existing one.

**An intent's `runner_version` now says which binary wrote it.** Hand-written records carry
`0.1.0-dev` beside artifacts carrying `0.1.0-dev+<commit>`; a created one carries the same string as
the `gate.yaml` of its first phase.

<!-- xeno:section:residual-risk -->
## Residual risk

**The host derivation is one line and nobody here has a GitHub Enterprise or self-managed GitLab
deployment to point it at.** It is stated as A78 and it is a convenience: `--for` takes the whole
qualified id, so a deployment the derivation gets wrong is worked around by the argument while the
row is corrected. The failure is visible rather than quiet — the id is written into `intent.yaml`
and read back by every phase of the intent.

**A wrong issue number is still sealed.** Nothing checks that the issue exists, so `--for` with a
typo in it writes an id for an issue nobody filed and `phase start` proceeds. That is exactly the
state before this change; what has changed is that three other fields can no longer be mistyped.

**The derived key is only as good as the directory.** A repository whose intent directories hold two
prefixes gets a refusal and has to name the key, which is the right outcome but one nobody has met
yet: this repository has one prefix and two padding widths, and the widths are what the tests cover.

**`tool_version` remains a hand edit.** Not caused here and not worsened, and now written down as a
learning. Until it is addressed, every phase of every intent goes red once on a field the runner
cannot write.

**What is not a risk.** The verdicts behind this intent: no gate, no rule and no existing artifact
is touched, and `gate verify` is 0 over all 241. And the tracker block consolidation: one type
where there were two declarations, with the existing tests of the only prior reader unchanged and
passing.
