---
intent: github.com/triplem/xeno#213
phase: 05-review
created: "2026-10-03T20:08:05Z"
schema_version: "1.0"
runner_version: dev+b2b1f0f.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a042109f8e7846577888a085cd4217a93974c7bb449183335909ef684de3e7c3
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
      One, in the implementation phase's deviations section: the draft of the WP8 paragraph said
      the digest "is read at P5", and the process definition says later phases read the earlier
      phase's output and digest. Caught before the commit was pushed, and it matters because a
      normative clause nearly shipped on a property that had not been checked, which is the
      failure this intent exists to fix.
  - rule: interface-change-needs-a-migration-note
    result: not-applicable
    note: >-
      No interface changes. One paragraph of the implementation plan and one row of
      ASSUMPTIONS.md; no code, no command, no field, no gate and no artifact shape. ContextLock
      already matched section 5 and is untouched, so nothing anybody depended on moved.
  - rule: new-dependency-needs-a-rationale
    result: met
    note: >-
      None added, and none was a candidate. The resolution uses an artifact that already exists
      rather than a mechanism, which is the reason it was preferred over dropping the clause.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. The specification change is its own commit, made before this
intent's work, and the maintainer chose the resolution from three put to them — which is
what "until a person changes it" requires. No field, gate, tool or rule is added. The
branch carries one intent, the commit references #213, and the issue carries `wp8`.

The process definition is untouched and still wins. It was the document with the better
reasoning as well as the privileged one, so the rule and the merits agreed here. They will
not always.

The acceptance criteria. WP8 asks for nothing section 5 forbids; the reason travels with
the clause rather than pointing at it; the record is named as the agent's own account;
A92 carries the resolution with both alternatives and why there is no gate.

The non-goals held. No section 5, no gate, no field, no code, and WP8's first half is left
for its own issue rather than folded in.

Conventions. Prose at 88, tables exempt, per #210. The paragraph was replaced rather than
edited into and read back as prose. The commit subject is a Conventional Commit at 61
characters after the first draft ran to 87, and the body is at 72.

One correction is in the implementation's deviations with its evidence: the draft claimed
the digest is read at P5 and the definition says later phases read it.

<!-- xeno:section:release-notes -->
## Release notes

WP8 and section 5 disagreed about `context.lock.yaml`. WP8 asked for every read outside the
context profile to appear in it; section 5 says the lock records what was declared and that
treating it as a measurement of what was read would be wrong. Both are normative, so the
plan's condition could not be met.

The record moves to the phase's digest, which is written at `phase finish`, sealed by
`artifacts_hash`, outside `context_hash`, and read by the phases after it. WP8 now carries
section 5's reason with it and says the record is the agent's own account, because nothing
in the harness reports what was opened.

No gate, no field, no code. A92 records the resolution and the two alternatives: the field
in section 5, and dropping the clause on A89's precedent, which was the closest call.

WP8's first half, "a repeated phase reads only what changed", is still unmet and gets its
own issue.

<!-- xeno:section:residual-risk -->
## Residual risk

The clause is unenforceable and now says so. An agent that reads outside its profile and
writes nothing leaves no trace, so the budget remains bounded by what the profile declared.
That is a weak record deliberately, and the alternative on the table was no record at all.

The deeper risk is the one the learning record names. #202's pass sorted every normative
clause by its reader in the code and never compared two clauses about the same artifact, so
a direct contradiction survived a pass built to find unread rules. There may be others of
the same shape, and nothing here looks for them: this intent fixed the instance, not the
class.

The resolution rests on four properties of the digest. Three are structural and will hold;
the fourth, that later phases read it, is a statement about how phases are worked rather
than something a gate asserts. If that stopped being true the record would still be written
and nobody would be reading it.

Section 5 and WP8 agreeing now is a state nothing maintains. Both are prose, neither
references the other by anchor, and a later edit to section 5's paragraph could reopen the
contradiction silently — which is how this one arrived.
