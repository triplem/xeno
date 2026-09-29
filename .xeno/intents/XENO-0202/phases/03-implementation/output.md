---
intent: github.com/triplem/xeno#120
phase: 03-implementation
created: "2026-09-29T06:02:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ad0a764.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: ebf8354b5fd84dd1b92b18350f4d0e41ca4ab06eb64a67df8813b96955eb281c
context_hash: 7fa895719de738c30f00205587534293d8f40cd60e5a0ecb40f04b02ad8910b9
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`.xeno/plugin/secrets.yaml`. The shipped filter: five patterns and five never digested
paths, with the reason the set is small written into the file. Every pattern is a shape
with a fixed prefix or a fixed envelope, which is why it can be matched without a
detector; entropy based detection is a different technique and the comment says so.

`internal/secrets`. `Filter`, `Load`, `Empty`, `Hash` and `Redact`, in about a hundred
lines with its own tests. A package of its own because section 4 says G-Secret and the
digest writer read the same file: a gate written later reads this rather than copying
it.

`Load` unions the shipped file and a project's additions, sorts the result and compiles
it. A missing shipped file yields an empty filter and no error. A regex that will not
compile drops that one pattern rather than the filter, and the pattern still enters the
hash, so what was in force stays reconstructable even where one line of it never fired.

`Hash` renders the effective set canonically, `pattern\t<id>\t<regex>\n` then
`path\t<glob>\n`, both sorted, and hashes that. An empty filter has no hash, so the
caller writes no field.

`Redact` replaces each match with `[redacted: <id>]`, which names the rule that fired
and leaves the line.

`internal/runner`. `writeDigest` loads the filter, redacts the summary and writes
`secrets_hash`. `SectionSet` writes the field into `output.md` and redacts nothing,
because section 16 puts the filtering on the digest.

`ASSUMPTIONS.md`. A35 amended to two fields. A62 records the computation Appendix B
delegates, and carries the measurement that decided the `by-hand` rule: removing
`secrets_hash` from `WriterlessHash` diverges all eighty-seven verdicts, because sixty
intents were sealed with `by-hand` when it was the only honest value.

Twelve tests: nine on the package and three through the runner. The one worth naming
asserts that the shipped filter catches each shape it claims and leaves ordinary prose
alone, because a filter that redacts prose teaches a reader to distrust the redaction.

This intent's own artifacts from here carry a `secrets_hash` the runner computed, and
their digests passed through the filter.

<!-- xeno:section:deviations -->
## Deviations from the design

None in the shape. The package, the union, the canonical hash, the named redaction, the
absent field for an empty filter and the untouched honesty rule are as P2 decided them.

Two things the design did not foresee, both inside decisions it had taken.

A regex that will not compile. P2 said nothing about it, and refusing to load would stop
every phase of a repository over one line of a project file. The pattern is dropped, the
rest of the filter stands, and the bad pattern still enters the hash so the field keeps
saying what was in force. It follows from the decision that a missing filter must not
refuse a phase.

The helper that fills the fields with no writer in this record's own frontmatter matched
a field name quoted in the prose, because an acceptance criterion in this intent
contains a field name with its value. It now reads the frontmatter block alone. That is
a note about the scaffolding of these artifacts rather than about the change, and it
belongs here because the same prose is what `SectionSet` deliberately does not redact.

The order of the record is the order of the work.
