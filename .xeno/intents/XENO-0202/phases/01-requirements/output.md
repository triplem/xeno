---
intent: github.com/triplem/xeno#120
phase: 01-requirements
created: "2026-09-29T05:58:06Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ad0a764.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 16f274be80b0e51ad08d5df00a705f3fe95ead57f76063694e72f658c33da07c
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

**AC1.** `.xeno/plugin/secrets.yaml` exists and carries `patterns`, each with an `id`
and a `regex`, and `paths_never_digested`, each a glob, in the shape section 4 gives.

**AC2.** The effective filter is the shipped file plus `.xeno/config/secrets.yaml` where
a project has one. A project adds patterns and paths and removes none, and a project
pattern reusing a shipped id leaves the shipped one in force.

**AC3.** `secrets_hash` is sixty four lowercase hex characters over the effective set,
computed from the set's content rather than the files' bytes: a comment edited in either
file leaves the value unchanged, and a pattern added to either changes it.

**AC4.** Two filters with the same patterns in a different order hash the same. The
value does not depend on the order patterns are written in or on which file they came
from.

**AC5.** `phase finish --summary` redacts the summary against the effective patterns
before writing the digest, and the digest carries `secrets_hash`. A summary carrying a
value a pattern matches reaches the file with that value replaced, and the rest of the
line intact.

**AC6.** `section set` writes `secrets_hash` into `output.md`.

**AC7.** No shipped filter means no filtering and no hash: the field is absent and the
summary is copied as it is. That is the state a repository is in before the plugin is
vendored, and it must not refuse a phase.

**AC8.** A35 names the two fields that still have no writer, `tool_version` and
`rules_hash`.

**AC9.** The `by-hand` honesty rule is unchanged, and the reason is recorded with the
number measured: eighty-seven verdicts diverge if it is tightened.

**AC10.** Everything green stays green. `gate verify` matches every verdict, and the
sixty intents carrying `secrets_hash: by-hand` are untouched.

<!-- xeno:section:non-goals -->
## Non goals

G-Secret. The gate that scans artifacts for secrets stays `not-implemented`. Section 4
says it reads the same file this writes against, which is what makes the shared rules
verifiable later; what the gate does when it finds a secret is its own decision and its
own work.

`rules_hash`, and WP4 with it.

`tool_version`. No home in section 12's block, so WP11 decides.

`paths_never_digested` as behaviour. The globs are shipped, enter the effective set and
enter the hash. Nothing acts on them, because the runner does not read the files a
summary talks about and has no text to exclude.

A large pattern set. The shipped file stays close to section 4's example. A project adds
what it needs, which is the documented route.

Redaction of anything but the summary. `output.md` is the agent's prose and is not
filtered; section 16 puts the filtering on the digest, which is the file that leaves a
session.

Tightening the `by-hand` rule. Measured, refused, recorded.

Rewriting sealed artifacts. Sixty carry `by-hand` inside their hashes and stay as they
are.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes. Section 4 defines the file, section 16 says who filters,
and Appendix B delegates the computation. All three already say what this builds.

The computation is a decision and is recorded as an assumption. Appendix B refuses to
fix it, so inventing it silently would leave the one number nobody could reconstruct.

A project may only widen the filter. Section 4 says a filter a project can switch off is
not a filter, so the union is the only shape available and matching stays a disjunction.

The hash covers the set, not the files. Two projects with the same effective filter
written in different orders, or with different comments, must agree, or `secrets_hash`
stops meaning which filter a digest passed through.

Redaction must not destroy the sentence around it. A digest is read by a person, so a
match is replaced and the text around it is left, which also makes it visible that
something was removed.

A missing filter must not refuse a phase. The field is absent, as it is for any field
with no source, and `by-hand` is not written in its place by the runner.

No new dependency. Regexes are `regexp`, globs are the matcher the model package already
has.

One intent, one issue. The commits reference #120, which stays open.
