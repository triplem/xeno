---
intent: github.com/triplem/xeno#337
phase: 05-review
created: "2026-10-09T16:00:53Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 39a35105fc83a6801cdb5509755bf0170130afbf109289ca38b5d98b97479bf6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: P3 records none; the trim and the count correction are written inside their phases.
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: A deferred package in the plan; nothing a project runs, reads or carries changes until WP21 is built.
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: A document change; nothing enters the binary, the plugin or the pipeline.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**deviations-are-traceable — met.** P3 records none; the trim of the #224 paragraph and the correction of the line count both happened inside their phases and are written there.

**interface-change-needs-a-migration-note — not-applicable.** No interface moves: the plan gains a deferred package and nothing a project runs, reads or carries changes until WP21 is built.

**new-dependency-needs-a-rationale — not-applicable.** A document change; nothing enters the binary, the plugin or the pipeline.

**What no rule asks.** Whether the agent should have written a normative document at all. The first standing rule says no, the maintainer instructed it after reading the diff, and the commit message, the intake and this phase each say so; the exception is a trace and not a precedent, and the plan's text does not say it was agent-written because a reader of the plan is owed the content and not its provenance.

<!-- xeno:section:release-notes -->
## Release notes

**The plan says how a greenfield project reaches its first issue.** Ideation is WP21, a sealed record beside the intents holding the vision and the plan an idea becomes and producing the issues the approval act then takes one at a time. It is not a phase, for the reasons #224 gives and one section 12 adds; it is specified in section 2 now and built in 1.1, beside WP18.

**What is unchanged.** The process definition, the runner and the way issues are written here: by hand, approved one at a time, until WP21 exists.

**For a reader of the plan.** WP21 is named in section 1's deferred list, section 6's sentence and size table, and section 8's scope of record, each where WP18 already was.

<!-- xeno:section:residual-risk -->
## Residual risk

**The shape is decided before its specification.** WP21's record type, path, template and commands are each a process definition change when built, and that change may find a cost the plan's paragraph does not name; the plan would then be edited again, by the same rule.

**The decisions nest.** D-2 and D-3 were framed after D-1 was answered; a reader who would have chosen a phase or a command finds the later two decided inside a choice they would not have made, and the intake records the sequence so that this is visible.

**The adapter contract widens in 1.1.** A fifth operation on a contract WP12 fixes at four, named as the decision WP12 says it is and not taken here.

**Belongs to the specification.** All of it: section 6, Appendix B, the template set and the command list, in that order, when WP21 is built.
