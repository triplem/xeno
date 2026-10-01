---
intent: github.com/triplem/xeno#158
phase: 05-review
created: "2026-10-01T15:16:46Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 03a7821821bd722f378c447f45c6c9a1d87f12d54a65fae42cab3c5a6a6805c3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Reviewed clause by clause against section 9: the two axes, the path deciding the level, `scope` checked
against it, `binding` under `given/` only and not under `given/project/`, `learned/builtin/` rejected,
the two kind errors, `abstract` above the project level, the three precedence rules, and the refusal to
resolve a binding collision. Three additions beyond the section are recorded in A67 with the line they
are drawn on, and the one candidate addition is refused as a specification change. Nothing is invented;
the one thing that had to be decided is the hash, which Appendix B delegates and A66 defines. The
release: G-Rules implemented, `internal/rules` new, `rules_hash` with a writer, and a project with no
rules unaffected. Residual risk, in order of weight: the first outside reader of a finding from this
gate will be whoever writes the shipped set; `rules_hash` is written and unchecked by design; the
`by-hand` honesty rule is a release behind the writer for the second field now; `kind: checked` is
indistinguishable from `review` in effect until the predicates land, which a project adopting a rule in
this window would not guess; the `version` requirement is the weakest of A67's three and is written to
be argued with; and no fixture carries all eight axis combinations at once.
