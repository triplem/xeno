---
intent: github.com/triplem/xeno#324
phase: 05-review
created: "2026-10-08T13:49:23Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 07c93cb934f41be084d452c8b22fd1e38279ed10483592fda9e1b9ef5ae1b1bb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: 'The implementation is what the design named: a commented line with its reason in the GitHub template, and a test that asserts both the presence and the two absences.'
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: No interface change. A generated file gains a comment; no command, flag, artifact field or template id moves, and a project that regenerates its wrapper gets the comment and nothing else.
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: No dependency of any kind is added.
    - result: met
      note: 'Operations lens: the change is a sentence an operator reads at the moment they need it, and the failure it describes -- checkout refused on a permission -- is the one that is hardest to diagnose from its symptom. Nothing in the gate path moves.'
      source: lens
---

# Review

<!-- xeno:section:release-notes -->
## Release notes

**The generated GitHub pipeline now names the container user.** The workflow `xeno init`
writes carries a commented `options: --user <uid>` under the image it names, with the reason:
the host mounts its own work directory into the container and chowns nothing,
`actions/checkout` runs inside it, and the image's user has to be able to write there.

A hosted runner needs nothing — that directory belongs to uid 1001 and the image runs as
1001. A self-hosted runner, or a controller deployment, may own it as another uid, and then
checkout fails on a permission before any gate runs. Uncomment the line with the uid your
runner uses.

The line is deliberately commented and no uid is defaulted beyond the one already true: a
number written here would be right on one runner and wrong on every other. GitLab's wrapper
carries nothing, because that executor leaves the build directory world writable and the
problem does not arise.
