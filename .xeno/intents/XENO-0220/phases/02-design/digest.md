---
intent: github.com/triplem/xeno#165
phase: 02-design
created: "2026-10-01T16:24:45Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+46d4cb2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fe47bd10d628135679c86db23e83da99ca958c09ca4df7e9f6ad878c97ace407
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Four rules whose kinds follow the templates rather than their statements: `deviations-are-traceable`
(review, P3), `interface-change-needs-a-migration-note` (review, P2 and P3, where section 9 writes it
as checked against sections the shipped templates do not have), `new-dependency-needs-a-rationale`
(review, P2 and P3) and `release-notes-are-filled` (checked, P5, `section-non-empty`). That one new
predicate is `sectionImpliesSection` without an antecedent, admitted by the same sentence: this package
delivers the section predicates the shipped set needs. Three review rules to one checked rule is a
finding about the templates, not the predicates — the two missing sections are WP3's to add. The
documentation consistency call, which the plan leaves here, is answered: it ships as an example, because
the mechanical part has nothing to point at until WP16 generates the pages and because a review rule in
`given/builtin/` would be answered `not-applicable` for ever by every project whose documentation lives
elsewhere. `xeno init` vendors the rule tree as it vendors the templates. The examples are complete
rules carrying `scope: org`, so copying one into the wrong level produces a `scope` finding — the gate
teaching the layout better than a comment could. The hooks call the binary rather than carrying a
pattern, because three copies of a pattern produce a push that rejects what the gate accepts.
